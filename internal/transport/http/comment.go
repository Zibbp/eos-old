package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/internal/comment"
	"github.com/zibbp/eos/internal/errors"
)

type CommentService interface {
	GetComments(c echo.Context, limit int, offset int, videoId string) (comment.Pagination, error)
}

func (h *Handler) GetComments(c echo.Context) error {
	queryLimit := c.QueryParam("limit")
	if queryLimit == "" {
		queryLimit = "10"
	}
	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		return echo.NewHTTPError(400, errors.New(400, "invalid limit"))
	}
	queryOffset := c.QueryParam("offset")
	if queryOffset == "" {
		queryOffset = "0"
	}
	offset, err := strconv.Atoi(queryOffset)
	if err != nil {
		return echo.NewHTTPError(400, errors.New(400, "invalid offset"))
	}

	videoId := c.QueryParam("video_id")

	if videoId == "" {
		return echo.NewHTTPError(400, errors.New(400, "invalid video_id"))
	}

	pagination, err := h.Service.CommentService.GetComments(c, limit, offset, videoId)
	if err != nil {
		return err
	}

	return c.JSON(200, pagination)
}
