package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/internal/errors"
	"github.com/zibbp/eos/internal/playback"
)

type PlaybackRequest struct {
	// VideoID   string               `json:"video_id" validate:"required"`
	Timestamp int    `json:"timestamp" validate:"required"`
	Status    string `json:"status" validate:"required" enum:"in_progress,finished"`
}

func UpdateProgress(c echo.Context) error {
	// get video id from parameter
	videoID := c.Param("video_id")
	if videoID == "" {
		return echo.NewHTTPError(400, errors.New(400, "video_id is required"))
	}

	var req PlaybackRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	// validate request
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	fmt.Println(req.Status)

	err := playback.UpdateProgress(videoID, req.Timestamp, req.Status, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, nil)
}

func GetProgress(c echo.Context) error {
	videoID := c.Param("video_id")
	if videoID == "" {
		return echo.NewHTTPError(400, errors.New(400, "video_id is required"))
	}

	pb, err := playback.GetProgress(videoID, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, pb)
}

func GetAllProgress(c echo.Context) error {
	pb, err := playback.GetAllProgress(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, pb)
}

func DeleteProgress(c echo.Context) error {
	videoID := c.Param("video_id")
	if videoID == "" {
		return echo.NewHTTPError(400, errors.New(400, "video_id is required"))
	}

	err := playback.DeleteProgress(videoID, c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, nil)
}
