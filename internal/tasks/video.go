package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/internal/redis"
	"github.com/zibbp/eos/internal/scanner"
	"github.com/zibbp/eos/internal/utils"
	"github.com/zibbp/eos/internal/video"
)

const (
	TypeVideoGenerateThumbnails = "video:generate_thumbnails"
	TypeVideoStartProcess       = "video:start_process"
	TypeVideoDownloadThumbnails = "video:download_thumbnails"
)

type VideoGenerateThumbnailsTask struct {
	VideoID string
}

type VideoStartProcessTask struct {
	Task utils.VideoProcessTask
}

type VideoDownloadThumbnailsTask struct {
	VideoID string
}

func NewVideoStartProcessTask(task utils.VideoProcessTask) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoStartProcessTask{
		Task: task,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoStartProcess, payload, asynq.MaxRetry(1), asynq.Queue(string(utils.ScannerQueue)), asynq.Timeout(0)), nil
}

func NewVideoGenerateThumbnailsTask(videoID string) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoGenerateThumbnailsTask{
		VideoID: videoID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoGenerateThumbnails, payload, asynq.MaxRetry(2), asynq.Queue(string(utils.ThumbnailGeneratorQueue)), asynq.Timeout(0)), nil
}

func NewVideoDownloadThumbnailsTask(videoID string) (*asynq.Task, error) {
	payload, err := json.Marshal(VideoDownloadThumbnailsTask{
		VideoID: videoID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeVideoDownloadThumbnails, payload, asynq.MaxRetry(0), asynq.Queue(string(utils.ScannerQueue)), asynq.Timeout(0)), nil
}

func HandleVideoStartProcessTask(ctx context.Context, t *asynq.Task) error {

	var p VideoStartProcessTask
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Info().Str("task", TypeVideoStartProcess).Msgf("starting video process: %s", p.Task)

	// get videos
	videos, err := video.ScannerGetVideos()
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoStartProcess).Msgf("failed to get videos")
	}

	for _, video := range videos {
		switch p.Task {
		case utils.DownloadThumbnails:
			// download thumbnails
			task, err := NewVideoDownloadThumbnailsTask(video.ID)
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoStartProcess).Msgf("failed to create task: %s", TypeVideoDownloadThumbnails)
				return err
			}
			videoTask, err := redis.GetAsynqClient().Client.Enqueue(task)
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoStartProcess).Msgf("failed to enqueue task: %s", TypeVideoDownloadThumbnails)
				return err
			}

			log.Debug().Str("task", TypeVideoStartProcess).Msgf("enqueued task: %s, video: %s", TypeVideoDownloadThumbnails, videoTask.ID)
		}
	}

	return nil
}

func HandleVideoGenerateThumbnailsTask(ctx context.Context, t *asynq.Task) error {

	var p VideoGenerateThumbnailsTask
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Info().Str("task", TypeVideoGenerateThumbnails).Msgf("generating thumbnails for video: %s", p.VideoID)

	// get video
	dbVideo, err := video.ScannerGetVideo(p.VideoID)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to get video: %s", p.VideoID)
		return err
	}

	err = utils.CreateDirectory(fmt.Sprintf("/tmp/%s", p.VideoID))
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to create directory: %s", p.VideoID)
		return err
	}

	// // ffprobe
	// ffprobeArgs := []string{
	// 	"-v",
	// 	"error",
	// 	"-select_streams",
	// 	"v:0",
	// 	"-show_entries",
	// 	"stream=codec_name",
	// 	"-of",
	// 	"default=noprint_wrappers=1:nokey=1",
	// }

	// // get video codec
	// videoCodec, err := utils.ExecuteFFprobeCommand(ffprobeArgs, dbVideo.VideoPath)
	// if err != nil {
	// 	log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to get video codec for video: %s", p.VideoID)
	// 	_ = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", p.VideoID))
	// 	return err
	// }
	// // trim any whitespace
	// videoCodec = []byte(string(videoCodec)[:len(videoCodec)-1])
	// fmt.Println(string(videoCodec))

	// hardwareDecodeAvailable := false
	// inputCodec := ""

	// switch string(videoCodec) {
	// case "h264":
	// 	hardwareDecodeAvailable = true
	// 	inputCodec = "h264_qsv"
	// case "vp9":
	// 	hardwareDecodeAvailable = true
	// 	inputCodec = "vp9_qsv"
	// case "av1":
	// 	hardwareDecodeAvailable = true
	// 	inputCodec = "av1_qsv"
	// }

	// preFfmpegArgs := []string{
	// 	"-hwaccel",
	// 	"qsv",
	// }

	// if hardwareDecodeAvailable {
	// 	preFfmpegArgs = append(preFfmpegArgs, "-c:v", inputCodec)
	// }

	// // -vf "hwdownload,format=nv12,fps=1/10,scale=160:90"

	// ffmpegArgs := []string{
	// 	"-vf",
	// 	"hwdownload,format=nv12,fps=1/10,scale=160:90",
	// 	fmt.Sprintf("/tmp/%s/thumbnail%%05d.jpg", p.VideoID),
	// }

	// CPU
	preFfmpegArgs := []string{}
	ffmpegArgs := []string{
		"-vf",
		"fps=1/10,scale=160:90:flags=lanczos",
		"-q:v",
		"3",
		fmt.Sprintf("/tmp/%s/thumbnail%%05d.jpg", p.VideoID),
	}

	// execute ffmpeg to generate thumbnails
	err = utils.ExecuteFFmpegCommand(dbVideo.VideoPath, preFfmpegArgs, ffmpegArgs)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to generate thumbnails for video: %s", p.VideoID)
		_ = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", p.VideoID))
		return err
	}

	// generate storyboard
	args := []string{
		"-geometry",
		"+0+0",
		"-tile",
		"5x",
	}

	err = utils.GenerateStoryboardImage(args, fmt.Sprintf("/tmp/%s/thumbnail*.jpg", p.VideoID), fmt.Sprintf("/tmp/%s/storyboard.jpg", p.VideoID))
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to generate storyboard for video: %s", p.VideoID)
		_ = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", p.VideoID))
		return err
	}

	// move image to video directory
	err = utils.MoveFile(fmt.Sprintf("/tmp/%s/storyboard.jpg", p.VideoID), fmt.Sprintf("%s/thumbnails.jpg", dbVideo.Path))
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to move storyboard for video: %s", p.VideoID)
		_ = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", p.VideoID))
		return err
	}

	// remove tmp directory
	err = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", p.VideoID))
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to remove tmp directory: %s", p.VideoID)
		return err
	}

	log.Info().Str("task", TypeVideoGenerateThumbnails).Msgf("finished generating thumbnails for video: %s", p.VideoID)

	return nil
}

func HandleVideoDownloadThumbnailsTask(ctx context.Context, t *asynq.Task) error {
	var p VideoDownloadThumbnailsTask
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	log.Info().Str("task", TypeVideoDownloadThumbnails).Msgf("downloading thumbnails for video: %s", p.VideoID)

	// get video
	dbVideo, err := video.ScannerGetVideo(p.VideoID)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to get video: %s", p.VideoID)
		return err
	}

	// get video info
	info, err := scanner.GetVideoInfo(dbVideo.JSONPath)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to get video info for video: %s", p.VideoID)
		return err
	}
	// vtt thumbnails
	for _, format := range info.Formats {
		// sb0 is the best format for thumbnails (160x90)
		if format.FormatID == "sb0" {
			// skip if there are no row or columns
			if format.Rows == nil || format.Columns == nil || format.FPS == nil {
				log.Info().Str("task", TypeVideoDownloadThumbnails).Msgf("skipping video %s, no rows, columns, or fps", info.ID)
				return nil
			}
			// create tmp dir
			err = utils.CreateDirectory(fmt.Sprintf("/tmp/%s", info.ID))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to create directory: %s", info.ID)
				return err
			}

			for i, fragment := range format.Fragments {
				err = utils.DownloadFile(fragment.URL, fmt.Sprintf("/tmp/%s/thumbnail%04d.jpg", info.ID, i))
				if err != nil {
					log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to download thumbnail: %s", fragment.URL)

					// allow this to fail and succeed job
					err = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", info.ID))
					if err != nil {
						log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to remove tmp directory: %s", info.ID)
						continue
					}
					return nil
				}
			}

			// generate storyboard
			args := []string{
				"-geometry",
				"+0+0",
				"-tile",
				"1x",
			}
			err = utils.GenerateStoryboardImage(args, fmt.Sprintf("/tmp/%s/thumbnail*.jpg", info.ID), fmt.Sprintf("/tmp/%s/thumbnails.jpg", info.ID))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to generate storyboard for video: %s", info.ID)
				return err
			}

			// move image to video directory
			err = utils.MoveFile(fmt.Sprintf("/tmp/%s/thumbnails.jpg", info.ID), fmt.Sprintf("%s/thumbnails.jpg", info.Path))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to move storyboard for video: %s", info.ID)
				return err
			}

			// remove tmp directory
			err = utils.RemoveDirectory(fmt.Sprintf("/tmp/%s", info.ID))
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to remove tmp directory: %s", info.ID)
				return err
			}

			// calculate interval of thumbnails
			interval := float64(format.Fragments[0].Duration) / float64(*format.Rows**format.Columns)

			// update video
			_, err = dbVideo.Update().SetThumbnailsPath(fmt.Sprintf("%s/%s", dbVideo.Path, "thumbnails.jpg")).SetThumbnailsWidth(int(*format.Width)).SetThumbnailsHeight(int(*format.Height)).SetThumbnailsInterval(interval).SetThumbnailsRows(int(*format.Rows)).Save(ctx)
			if err != nil {
				log.Error().Err(err).Str("task", TypeVideoDownloadThumbnails).Msgf("failed to update video %s", info.ID)
				return err
			}

			log.Info().Str("task", TypeVideoDownloadThumbnails).Msgf("generated thumbnails for video %s", info.ID)
		}
	}

	return nil
}
