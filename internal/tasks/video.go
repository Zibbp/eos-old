package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/internal/utils"
	"github.com/zibbp/eos/internal/video"
)

const (
	TypeVideoGenerateThumbnails = "video:generate_thumbnails"
)

type VideoGenerateThumbnailsTask struct {
	VideoID string
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

	ffmpegArgs := []string{
		"-vf",
		"fps=1/10,scale=160:90:flags=fast_bilinear",
		"-q:v",
		"2",
		fmt.Sprintf("/tmp/%s/thumbnail%%05d.jpg", p.VideoID),
	}

	// execute ffmpeg to generate thumbnails
	err = utils.ExecuteFFmpegCommand(dbVideo.VideoPath, ffmpegArgs)
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to generate thumbnails for video: %s", p.VideoID)
		return err
	}

	// generate storyboard
	err = utils.GenerateStoryboardImage(fmt.Sprintf("/tmp/%s/thumbnail*.jpg", p.VideoID), fmt.Sprintf("/tmp/%s/storyboard.jpg", p.VideoID))
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to generate storyboard for video: %s", p.VideoID)
		return err
	}

	// move image to video directory
	err = utils.MoveFile(fmt.Sprintf("/tmp/%s/storyboard.jpg", p.VideoID), fmt.Sprintf("%s/thumbnails.jpg", dbVideo.Path))
	if err != nil {
		log.Error().Err(err).Str("task", TypeVideoGenerateThumbnails).Msgf("failed to move storyboard for video: %s", p.VideoID)
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
