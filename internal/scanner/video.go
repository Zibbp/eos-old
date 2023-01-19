package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/internal/channel"
	"github.com/zibbp/eos/internal/comment"
	"github.com/zibbp/eos/internal/kv"
	"github.com/zibbp/eos/internal/utils"
	"github.com/zibbp/eos/internal/video"
)

type VideoInfo struct {
	ID             string          `json:"id"`
	Channel        string          `json:"channel"`
	ChannelID      string          `json:"channel_id"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	Uploader       string          `json:"uploader"`
	Duration       int64           `json:"duration"`
	ViewCount      int64           `json:"view_count"`
	LikeCount      int64           `json:"like_count"`
	DislikeCount   int64           `json:"dislike_count"`
	Format         string          `json:"format"`
	Width          int64           `json:"width"`
	Height         int64           `json:"height"`
	Resolution     string          `json:"resolution"`
	FPS            float64         `json:"fps"`
	VideoCodec     string          `json:"vcodec"`
	VBR            float64         `json:"vbr"`
	AudioCodec     string          `json:"acodec"`
	ABR            float64         `json:"abr"`
	Epoch          int64           `json:"epoch"`
	CommentCount   int64           `json:"comment_count"`
	VideoPath      string          `json:"video_path"`
	ThumbnailPath  string          `json:"thumbnail_path"`
	JsonPath       string          `json:"json_path"`
	SubtitlePath   string          `json:"subtitle_path"`
	UploadDate     string          `json:"upload_date"`
	TempUploadDate time.Time       `json:"temp_upload_date"`
	Path           string          `json:"path"`
	Comments       []Comment       `json:"comments"`
	Chapters       []video.Chapter `json:"chapters"`
}

type Comment struct {
	ID               string    `json:"id"`
	Text             string    `json:"text"`
	Timestamp        int64     `json:"timestamp"`
	LikeCount        int64     `json:"like_count"`
	IsFavorited      bool      `json:"is_favorited"`
	Author           string    `json:"author"`
	AuthorID         string    `json:"author_id"`
	AuthorThumbnail  string    `json:"author_thumbnail"`
	AuthorIsUploader bool      `json:"author_is_uploader"`
	Parent           string    `json:"parent"`
	VideoID          string    `json:"video_id"`
	Replies          []Comment `json:"replies"`
}

func (s *Service) StartScanner() error {
	log.Info().Msg("starting scanner")

	// Get root channel folder to process
	channelDirs, err := utils.GetFoldersInDir("/videos")
	if err != nil {
		kv.DB().Set("scanner", "stopped")
		log.Error().Err(err).Msg("failed to get channel dirs")
		return fmt.Errorf("failed to get channel dirs: %w", err)
	}

	// Fetch existing videos
	existingVideos, err := video.ScannerGetVideos()
	if err != nil {
		kv.DB().Set("scanner", "stopped")
		log.Error().Err(err).Msg("failed to get videos")
		return fmt.Errorf("failed to get videos: %w", err)
	}

	log.Info().Msgf("found %d channel dirs", len(channelDirs))

	var existingVideoPaths []string
	var existingVideoIDs []string
	for _, existingVideo := range existingVideos {
		existingVideoPaths = append(existingVideoPaths, existingVideo.Path)
		existingVideoIDs = append(existingVideoIDs, existingVideo.ID)
	}

	for _, channelDir := range channelDirs {
		s.ProcessChannelDirectory(channelDir, existingVideoPaths, existingVideoIDs)

	}

	log.Info().Msg("finished scanning")

	// Set scanner to stopped
	kv.DB().Set("scanner", "stopped")
	return nil
}

func (s *Service) ProcessChannelDirectory(channelDir string, existingVideoPaths []string, existingVideoIDs []string) {
	log.Info().Msgf("processing channel dir: %s", channelDir)
	channelDirPath := fmt.Sprintf("/videos/%s", channelDir)
	httpPath := fmt.Sprintf("/%s", channelDir)
	// Get video folders in channel dir
	videoDirs, err := utils.GetFoldersInDir(channelDirPath)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get video dirs in channel dir %s", channelDir)
	}

	log.Info().Msgf("found %d video dirs in channel dir %s", len(videoDirs), channelDir)

	// Folders that need to be skipped
	ytDlpChannelDir := fmt.Sprintf("%s-NA-%s-Videos", channelDir, channelDir)
	ytDlpPlaylistFolder := fmt.Sprintf("%s-NA-YT-Archive", channelDir)

	for i, videoDir := range videoDirs {
		// Skip some yt-dlp folders
		if videoDir == ytDlpChannelDir || videoDir == ytDlpPlaylistFolder {
			continue
		}

		// Quick check if video dir is already in db by checking if the path exists
		videoDirPath := fmt.Sprintf("%s/%s", channelDirPath, videoDir)
		if utils.StringInSlice(videoDirPath, existingVideoPaths) {
			log.Debug().Msgf("video dir %s already exists in db", videoDir)
			continue
		}

		log.Debug().Msgf("Processing video directory %s", videoDir)

		var importVid VideoInfo

		// Get files in video dir
		videoFiles, err := utils.GetFilesInDir(fmt.Sprintf("%s/%s", channelDirPath, videoDir))
		if err != nil {
			log.Error().Err(err).Msgf("failed to get files in video dir %s", videoDir)
			continue
		}

		// Process video files
		for _, videoFile := range videoFiles {
			// Get file extension
			fileSlice := strings.LastIndex(videoFile, ".")
			fileExt := videoFile[fileSlice+1:]

			switch fileExt {
			case "json":
				// Ensure the correct json file is being processed
				if strings.Contains(videoFile, "info.json") {
					jsonData, err := os.ReadFile(fmt.Sprintf("%s/%s/%s", channelDirPath, videoDir, videoFile))
					if err != nil {
						log.Error().Err(err).Msgf("failed to read json file %s", videoFile)
						continue
					}

					// Unmarshal json data
					err = json.Unmarshal(jsonData, &importVid)
					if err != nil {
						log.Error().Err(err).Msgf("failed to unmarshal json file %s", videoFile)
						continue
					}
					// Set json path
					importVid.JsonPath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
				}
			case "mkv":
				// Set video path
				importVid.VideoPath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
			case "webp":
				// Set thumbnail path
				importVid.ThumbnailPath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
			case "jpg":
				// Set thumbnail path
				importVid.ThumbnailPath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
			case "vtt":
				// Set subtitle path
				importVid.SubtitlePath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
			case "srt":
				// Set subtitle path
				importVid.SubtitlePath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
			case "mp4":
				// Set video path
				importVid.VideoPath = fmt.Sprintf("%s/%s/%s", httpPath, videoDir, videoFile)
			}

		}

		// Set path
		importVid.Path = fmt.Sprintf("%s/%s", channelDirPath, videoDir)

		// Check if json file was found
		if importVid.JsonPath == "" {
			log.Debug().Msgf("json file not found in video dir %s", videoDir)
			continue
		}

		// Get the first video in the loop to check for channel
		if i == 0 {
			// Get channel
			dbChannel, err := channel.ScannerGetChannel(importVid.ChannelID)
			if err != nil {
				log.Error().Err(err).Msgf("failed to get channel %s", importVid.ChannelID)
				continue
			}
			if dbChannel == nil {
				// Create channel
				channelDto := &channel.Channel{
					ID:        importVid.ChannelID,
					Name:      importVid.Channel,
					ImagePath: fmt.Sprintf("%s/%s.jpg", httpPath, channelDir),
				}
				_, err = channel.ScannerCreateChannel(channelDto)
				if err != nil {
					log.Error().Err(err).Msgf("failed to create channel %s", importVid.ChannelID)
					continue
				}
			}
		}

		// Check if video is already in db by checking ID
		// This catches if the video is already in the db but the path is different
		if utils.StringInSlice(importVid.ID, existingVideoIDs) {
			log.Debug().Msgf("video %s already exists in db", importVid.ID)
			continue
		}

		// Parse Date
		parsedUploadDate, err := time.Parse("20060102", importVid.UploadDate)
		if err != nil {
			log.Error().Err(err).Msgf("failed to parse upload date %s for video %s", importVid.UploadDate, importVid.ID)
			continue
		}

		// Create video in database
		videoDto := &video.Video{
			ID:           importVid.ID,
			Title:        importVid.Title,
			Description:  importVid.Description,
			UploadDate:   parsedUploadDate,
			Uploader:     importVid.Uploader,
			Duration:     importVid.Duration,
			ViewCount:    importVid.ViewCount,
			LikeCount:    importVid.LikeCount,
			DislikeCount: importVid.DislikeCount,
			Format:       importVid.Format,
			Width:        importVid.Width,
			Height:       importVid.Height,
			Resolution:   importVid.Resolution,
			FPS:          importVid.FPS,
			AudioCodec:   importVid.AudioCodec,
			VideoCodec:   importVid.VideoCodec,
			ABR:          importVid.ABR,
			VBR:          importVid.VBR,
			Epoch:        importVid.Epoch,
			CommentCount: importVid.CommentCount,
			// Tags:          importVid.Tags,
			// Categories:    importVid.Categories,
			VideoPath:     importVid.VideoPath,
			ThumbnailPath: importVid.ThumbnailPath,
			JSONPath:      importVid.JsonPath,
			CaptionPath:   importVid.SubtitlePath,
			Path:          importVid.Path,
		}

		// Import video
		err = video.ScannerCreateVideo(videoDto, importVid.ChannelID)
		if err != nil {
			log.Error().Err(err).Msgf("failed to create video %s", importVid.ID)
			continue
		}
		log.Info().Msgf("imported video %s", importVid.ID)

		// Chapters
		if len(importVid.Chapters) > 0 {
			for _, chapter := range importVid.Chapters {
				parsedChapter := video.Chapter{
					ID:        uuid.New().String(),
					StartTime: chapter.StartTime,
					EndTime:   chapter.EndTime,
					Title:     chapter.Title,
					VideoID:   importVid.ID,
				}
				err := video.ScannerCreateChapter(&parsedChapter, importVid.ID)
				if err != nil {
					log.Error().Err(err).Msgf("failed to create chapter %s for video %s", parsedChapter.ID, importVid.ID)
					continue
				}
			}
			log.Info().Msgf("imported chapters for video %s", importVid.ID)
		}

		// Comments
		if len(importVid.Comments) > 0 {
			for _, vidComment := range importVid.Comments {
				parsedComment := comment.Comment{
					ID:               vidComment.ID,
					Text:             vidComment.Text,
					LikeCount:        vidComment.LikeCount,
					IsFavorited:      vidComment.IsFavorited,
					Author:           vidComment.Author,
					AuthorID:         vidComment.AuthorID,
					AuthorThumbnail:  vidComment.AuthorThumbnail,
					AuthorIsUploader: vidComment.AuthorIsUploader,
					Parent:           vidComment.Parent,
				}
				// Parse Epoch to time
				parsedTime := time.Unix(vidComment.Timestamp, 0)
				if err != nil {
					log.Error().Err(err).Msgf("failed to parse timestamp %s for comment %s", vidComment.Timestamp, vidComment.ID)
					continue
				}
				parsedComment.Timestamp = parsedTime
				err = comment.ScannerCreateComment(&parsedComment, importVid.ID)
				if err != nil {
					log.Error().Err(err).Msgf("failed to create comment %s for video %s", parsedComment.ID, importVid.ID)
					continue
				}

			}

			log.Info().Msgf("imported comments for video %s", importVid.ID)
		}

	}

}
