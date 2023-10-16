package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/ent"
	"github.com/zibbp/eos/internal/channel"
	"github.com/zibbp/eos/internal/comment"
	"github.com/zibbp/eos/internal/redis"
	"github.com/zibbp/eos/internal/scanner"
	"github.com/zibbp/eos/internal/utils"
	"github.com/zibbp/eos/internal/video"
)

const (
	TypeVideoStartScanner = "video:start_scanner"
	TypeVideoScanChannel  = "video:scan_channel"
	TypeVideoProcess      = "video:process"
)

type VideoStartScannerPayload struct {
	Type utils.ScanType
}

type VideoScanChannelPayload struct {
	ChannelName string
	Type        utils.ScanType
}

type VideoProcessPayload struct {
	ChannelDirectoryPath string
	VideoName            string
	HttpPath             string
	Type                 utils.ScanType
	Channels             []*ent.Channel
	ExistingVideoIDs     []string
}

func NewVideoStartScannerTask(scanType utils.ScanType) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoStartScannerPayload{Type: scanType})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoStartScanner, payload), nil
}

func NewVideoScanChannelTask(channelName string, scanType utils.ScanType) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoScanChannelPayload{ChannelName: channelName, Type: scanType})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoScanChannel, payload), nil
}

func NewVideoProcessTask(channelDirectoryPath string, videoName string, httpPath string, scanType utils.ScanType, channels []*ent.Channel, existingVideoIDs []string) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoProcessPayload{ChannelDirectoryPath: channelDirectoryPath, VideoName: videoName, Type: scanType, Channels: channels, HttpPath: httpPath, ExistingVideoIDs: existingVideoIDs})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoProcess, payload), nil
}

func HandleVideoStartScannerTask(ctx context.Context, t *asynq.Task) error {
	log.Info().Str("task", TypeVideoStartScanner).Msg("running task")

	var p VideoStartScannerPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	channelDirectories, err := utils.GetFoldersInDir(scanner.VideoDirectory)
	if err != nil {
		return err
	}

	log.Info().Str("task", TypeVideoStartScanner).Msgf("scanned %d channel directories", len(channelDirectories))

	// create tasks for each channel
	for _, channelDirectory := range channelDirectories {
		// create task for channel
		task, err := NewVideoScanChannelTask(channelDirectory, p.Type)
		if err != nil {
			return err
		}

		// enqueue task
		channelTask, err := redis.GetAsynqClient().Client.Enqueue(task, asynq.Queue(string(utils.ScannerQueue)), asynq.MaxRetry(2))
		if err != nil {
			return err
		}

		log.Info().Str("task", TypeVideoStartScanner).Msgf("enqueued task %s for channel %s", channelTask.Type, channelDirectory)
	}

	return nil
}

func HandleVideoScanChannelTask(ctx context.Context, t *asynq.Task) error {

	var p VideoScanChannelPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Info().Str("task", TypeVideoScanChannel).Msgf("running task for channel %s", p.ChannelName)

	channelDirectoryPath := fmt.Sprintf("%s/%s", scanner.VideoDirectory, p.ChannelName)
	httpPath := fmt.Sprintf("/%s", p.ChannelName)

	// Get video folders in channel directory
	videoDirectories, err := utils.GetFoldersInDir(channelDirectoryPath)
	if err != nil {
		return err
	}

	log.Info().Str("task", TypeVideoScanChannel).Msgf("found %d videos in channel %s", len(videoDirectories), p.ChannelName)

	// fetch channels
	channels, err := channel.ScannerGetChannels()
	if err != nil {
		return err
	}

	// fetch existing videos
	existingVideos, err := video.ScannerGetVideos()
	if err != nil {
		return err
	}

	var existingVideoIDs []string
	var existingVideoPaths []string
	for _, existingVideo := range existingVideos {
		existingVideoIDs = append(existingVideoIDs, existingVideo.ID)
		existingVideoPaths = append(existingVideoPaths, existingVideo.Path)
	}

	// loop through video folders
	for _, videoDirectory := range videoDirectories {
		// check if video path is already in database
		videoDirectoryPath := fmt.Sprintf("%s/%s", channelDirectoryPath, videoDirectory)
		if utils.StringInSlice(videoDirectoryPath, existingVideoPaths) {
			log.Debug().Str("task", TypeVideoScanChannel).Msgf("video %s already exists", videoDirectoryPath)
			continue
		}

		// enqueue task to process video
		processVideoTask, err := NewVideoProcessTask(channelDirectoryPath, videoDirectory, httpPath, p.Type, channels, existingVideoIDs)
		if err != nil {
			return err
		}

		processVideoTaskInfo, err := redis.GetAsynqClient().Client.Enqueue(processVideoTask, asynq.Queue(string(utils.ScannerQueue)), asynq.MaxRetry(2))
		if err != nil {
			return err
		}

		log.Info().Str("task", TypeVideoScanChannel).Msgf("enqueued task %s for video %s", processVideoTaskInfo.Type, videoDirectory)
	}

	return nil
}

func HandleVideoProcessTask(ctx context.Context, t *asynq.Task) error {

	var p VideoProcessPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to unmarshal payload")
		return err
	}

	// TODO: check if video dir already in DB? or do this in the channel scanner

	log.Info().Str("task", TypeVideoProcess).Msgf("processing video %s", p.VideoName)

	videoDirectoryPath := fmt.Sprintf("%s/%s", p.ChannelDirectoryPath, p.VideoName)

	var importVideo scanner.VideoInfo

	// Get files in video directory
	videoFiles, err := utils.GetFilesInDir(videoDirectoryPath)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to get files in directory")
		return err
	}

	// loop through video files
	for _, videoFile := range videoFiles {
		// get file extension
		fileExtension := utils.GetFileExtension(videoFile)
		switch fileExtension {
		case "json":
			// parse info json file
			jsonData, err := os.ReadFile(fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to parse json")
				return err
			}
			err = json.Unmarshal(jsonData, &importVideo)
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to unmarshal payload")
				return err
			}

			importVideo.JsonPath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		case "mkv":
			importVideo.VideoPath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		case "webp":
			importVideo.ThumbnailPath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		case "jpg":
			importVideo.ThumbnailPath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		case "vtt":
			importVideo.SubtitlePath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		case "srt":
			importVideo.SubtitlePath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		case "mp4":
			importVideo.VideoPath = fmt.Sprintf("%s/%s", videoDirectoryPath, videoFile)
		}

	}

	// set path
	importVideo.Path = fmt.Sprintf("%s/%s", p.ChannelDirectoryPath, p.VideoName)

	// Check if json file was found
	if importVideo.JsonPath == "" {
		log.Error().Str("task", TypeVideoProcess).Msgf("json file not found for video %s", p.VideoName)
		return fmt.Errorf("json file not found")
	}

	// check if channel exists
	dbChannel, err := channel.ScannerGetChannel(importVideo.ChannelID)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to get channel %s", importVideo.ChannelID)
		return err
	}

	if dbChannel == nil {
		log.Error().Str("task", TypeVideoProcess).Msgf("channel %s not found for video %s", importVideo.ChannelID, p.VideoName)

		// create channel now to prevent blocking
		channelDTO := &channel.Channel{
			ID:        importVideo.ChannelID,
			Name:      importVideo.Channel,
			ImagePath: fmt.Sprintf("%s/%s.jpg", p.HttpPath, path.Base(p.ChannelDirectoryPath)),
		}
		_, err := channel.ScannerCreateChannel(channelDTO)
		if err != nil {
			log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to create channel %s", importVideo.ChannelID)
			return err
		}
	}

	// check if video already exists
	if utils.StringInSlice(importVideo.ID, p.ExistingVideoIDs) {
		log.Debug().Str("task", TypeVideoProcess).Msgf("video %s already exists", importVideo.ID)
		return nil
	}

	// parse date
	parsedUploadDate, err := time.Parse("20060102", importVideo.UploadDate)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to parse upload date '%s' for %s", importVideo.UploadDate, importVideo.ID)
		return err
	}

	videoDto := &video.Video{
		ID:           importVideo.ID,
		Title:        importVideo.Title,
		Description:  importVideo.Description,
		UploadDate:   parsedUploadDate,
		Uploader:     importVideo.Uploader,
		Duration:     importVideo.Duration,
		ViewCount:    importVideo.ViewCount,
		LikeCount:    importVideo.LikeCount,
		DislikeCount: importVideo.DislikeCount,
		Format:       importVideo.Format,
		Width:        importVideo.Width,
		Height:       importVideo.Height,
		Resolution:   importVideo.Resolution,
		FPS:          importVideo.FPS,
		AudioCodec:   importVideo.AudioCodec,
		VideoCodec:   importVideo.VideoCodec,
		ABR:          importVideo.ABR,
		VBR:          importVideo.VBR,
		Epoch:        importVideo.Epoch,
		CommentCount: importVideo.CommentCount,
		// Tags:          importVideo.Tags,
		// Categories:    importVideo.Categories,
		VideoPath:     importVideo.VideoPath,
		ThumbnailPath: importVideo.ThumbnailPath,
		JSONPath:      importVideo.JsonPath,
		CaptionPath:   importVideo.SubtitlePath,
		Path:          importVideo.Path,
	}

	// create video
	err = video.ScannerCreateVideo(videoDto, importVideo.ChannelID)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to create video")
		return err
	}

	log.Info().Str("task", TypeVideoProcess).Msgf("imported video %s", importVideo.ID)

	// chapters
	if len(importVideo.Chapters) > 0 {
		for _, chapter := range importVideo.Chapters {
			parsedChapter := video.Chapter{
				ID:        uuid.New().String(),
				StartTime: chapter.StartTime,
				EndTime:   chapter.EndTime,
				Title:     chapter.Title,
				VideoID:   importVideo.ID,
			}
			err := video.ScannerCreateChapter(&parsedChapter, importVideo.ID)
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to create chapter")
				continue
			}
		}
		log.Info().Str("task", TypeVideoProcess).Msgf("imported %d chapters for video %s", len(importVideo.Chapters), importVideo.ID)
	}

	// comments
	if len(importVideo.Comments) > 0 {
		for _, videoComment := range importVideo.Comments {
			parsedComment := comment.Comment{
				ID:               videoComment.ID,
				Text:             videoComment.Text,
				LikeCount:        videoComment.LikeCount,
				IsFavorited:      videoComment.IsFavorited,
				Author:           videoComment.Author,
				AuthorID:         videoComment.AuthorID,
				AuthorThumbnail:  videoComment.AuthorThumbnail,
				AuthorIsUploader: videoComment.AuthorIsUploader,
				Parent:           videoComment.Parent,
			}
			// parse epoch to time]
			parsedTime := time.Unix(videoComment.Timestamp, 0)
			parsedComment.Timestamp = parsedTime
			err = comment.ScannerCreateComment(&parsedComment, importVideo.ID)
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msg("failed to create comment")
				continue
			}
		}
		log.Info().Str("task", TypeVideoProcess).Msgf("imported %d comments for video %s", len(importVideo.Comments), importVideo.ID)
	}

	// vtt thumbnails
	for _, format := range importVideo.Formats {
		if format.FormatID == "sb0" {
			// create tmp dir
			err = utils.CreateDirectory(fmt.Sprintf("/tmp/%s", importVideo.ID))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to create directory: %s", importVideo.ID)
				return err
			}

			for i, fragment := range format.Fragments {
				err = utils.DownloadFile(fragment.URL, fmt.Sprintf("/tmp/%s/thumbnail%04d.jpg", importVideo.ID, i))
				if err != nil {
					log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to download thumbnail: %s", fragment.URL)
					continue
				}
			}

			// generate storyboard
			err = utils.GenerateStoryboardImage(fmt.Sprintf("/tmp/%s/thumbnail*.jpg", importVideo.ID), fmt.Sprintf("/tmp/%s/thumbnails.jpg", importVideo.ID))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to generate storyboard for video: %s", importVideo.ID)
				return err
			}

			// move image to video directory
			err = utils.MoveFile(fmt.Sprintf("/tmp/%s/thumbnails.jpg", importVideo.ID), fmt.Sprintf("%s/thumbnails.jpg", importVideo.Path))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to move storyboard for video: %s", importVideo.ID)
				return err
			}

			// remove tmp directory
			err = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", importVideo.ID))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoProcess).Msgf("failed to remove tmp directory: %s", importVideo.ID)
				return err
			}

			log.Info().Str("task", TypeVideoProcess).Msgf("generated thumbnails for video %s", importVideo.ID)
		}
	}

	return nil
}
