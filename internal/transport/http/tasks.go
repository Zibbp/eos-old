package http

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/internal/errors"
	"github.com/zibbp/eos/internal/redis"
	"github.com/zibbp/eos/internal/tasks"
	"github.com/zibbp/eos/internal/utils"
)

type StartVideoScannerTaskRequest struct {
	Type utils.ScanType `json:"type" validate:"required,oneof=full quick"`
}

type StartVideoGenerateThumbnailsTaskRequest struct {
	VideoID string `json:"video_id" validate:"required"`
}

type StartVideoDownloadThumbnailsRequest struct {
	VideoID string `json:"video_id" validate:"required"`
}

func (h *Handler) StartVideoScannerTask(c echo.Context) error {

	var request StartVideoScannerTaskRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	task, err := tasks.NewVideoStartScannerTask(request.Type)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	info, err := redis.GetAsynqClient().Client.Enqueue(task, asynq.Queue(string(utils.ScannerQueue)), asynq.MaxRetry(2))
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}
	log.Info().Msgf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	response := map[string]interface{}{
		"task_id": info.ID,
	}

	return c.JSON(200, response)
}

func (h *Handler) StartVideoGenerateThumbnailsTask(c echo.Context) error {

	var request StartVideoGenerateThumbnailsTaskRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	task, err := tasks.NewVideoGenerateThumbnailsTask(request.VideoID)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	info, err := redis.GetAsynqClient().Client.Enqueue(task, asynq.Queue(string(utils.ThumbnailGeneratorQueue)), asynq.Unique(time.Hour))
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}
	log.Info().Msgf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	response := map[string]interface{}{
		"task_id": info.ID,
	}

	return c.JSON(200, response)
}

func (h *Handler) StartVideoDownloadThumbnailsTask(c echo.Context) error {
	task, err := tasks.NewVideoStartProcessTask(utils.DownloadThumbnails)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	info, err := redis.GetAsynqClient().Client.Enqueue(task, asynq.Queue(string(utils.ThumbnailGeneratorQueue)), asynq.Unique(time.Hour))
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}
	log.Info().Msgf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	response := map[string]interface{}{
		"task_id": info.ID,
	}

	return c.JSON(200, response)
}
