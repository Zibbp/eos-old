package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/ent"
	"github.com/zibbp/eos/internal/errors"
	"github.com/zibbp/eos/internal/video"
)

type VideoService interface {
	GetVideos(c echo.Context, limit int, offset int, channelId string) (video.Pagination, error)
	GetVideo(c echo.Context, id string) (*ent.Video, error)
	CreateVideo(c echo.Context, channelID string, videoDto *video.Video) (*ent.Video, error)
	UpdateVideo(c echo.Context, vID string, cID string, videoDto *video.Video) (*ent.Video, error)
	DeleteVideo(c echo.Context, id string) error
	GetVideosByChannelID(c echo.Context, channelID string) ([]*ent.Video, error)
	GetRandomVideos(c echo.Context, limit int) ([]*ent.Video, error)
	SearchVideos(c echo.Context, limit int, offset int, query string) (video.Pagination, error)
}

type CreateVideoRequest struct {
	ID            string  `json:"id"`
	ChannelID     string  `json:"channel_id" validate:"required"`
	Title         string  `json:"title" validate:"required"`
	Description   string  `json:"description"`
	UploadDate    string  `json:"upload_date" validate:"required"`
	Uploader      string  `json:"uploader"`
	Duration      int64   `json:"duration" validate:"required"`
	ViewCount     int64   `json:"view_count" validate:"required"`
	LikeCount     int64   `json:"like_count" validate:"required"`
	DislikeCount  int64   `json:"dislike_count"`
	Format        string  `json:"format"`
	Width         int64   `json:"width"`
	Height        int64   `json:"height"`
	Resolution    string  `json:"resolution"`
	FPS           float64 `json:"fps"`
	AudioCodec    string  `json:"audio_codec"`
	VideoCodec    string  `json:"video_codec"`
	ABR           float64 `json:"abr"`
	VBR           float64 `json:"vbr"`
	Epoch         int64   `json:"epoch"`
	CommentCount  int64   `json:"comment_count"`
	Tags          string  `json:"tags"`
	Categories    string  `json:"categories"`
	VideoPath     string  `json:"video_path" validate:"required"`
	ThumbnailPath string  `json:"thumbnail_path" validate:"required"`
	JSONPath      string  `json:"json_path" validate:"required"`
	CaptionPath   string  `json:"caption_path"`
	Path          string  `json:"path" validate:"required"`
}

func (h *Handler) GetVideos(c echo.Context) error {

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

	channelId := c.QueryParam("channel_id")

	videos, err := h.Service.VideoService.GetVideos(c, limit, offset, channelId)
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, videos)
}

func (h *Handler) SearchVideos(c echo.Context) error {

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

	q := c.QueryParam("q")
	if q == "" {
		return echo.NewHTTPError(400, errors.New(400, "invalid query param q"))
	}

	videos, err := h.Service.VideoService.SearchVideos(c, limit, offset, q)
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, videos)
}

func (h *Handler) GetVideo(c echo.Context) error {
	video, err := h.Service.VideoService.GetVideo(c, c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}
	return c.JSON(200, video)
}

func (h *Handler) CreateVideo(c echo.Context) error {
	var videoDto CreateVideoRequest
	if err := c.Bind(&videoDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	if err := c.Validate(videoDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	cID := videoDto.ChannelID

	// Date
	parsedUploadDate, err := time.Parse("20060102", videoDto.UploadDate)
	if err != nil {
		return echo.NewHTTPError(400, errors.New(400, fmt.Sprintf("invalid upload date: %s", err.Error())))
	}

	video := &video.Video{
		ID:            videoDto.ID,
		Title:         videoDto.Title,
		Description:   videoDto.Description,
		UploadDate:    parsedUploadDate,
		Uploader:      videoDto.Uploader,
		Duration:      videoDto.Duration,
		ViewCount:     videoDto.ViewCount,
		LikeCount:     videoDto.LikeCount,
		DislikeCount:  videoDto.DislikeCount,
		Format:        videoDto.Format,
		Width:         videoDto.Width,
		Height:        videoDto.Height,
		Resolution:    videoDto.Resolution,
		FPS:           videoDto.FPS,
		AudioCodec:    videoDto.AudioCodec,
		VideoCodec:    videoDto.VideoCodec,
		ABR:           videoDto.ABR,
		VBR:           videoDto.VBR,
		Epoch:         videoDto.Epoch,
		CommentCount:  videoDto.CommentCount,
		Tags:          videoDto.Tags,
		Categories:    videoDto.Categories,
		VideoPath:     videoDto.VideoPath,
		ThumbnailPath: videoDto.ThumbnailPath,
		JSONPath:      videoDto.JSONPath,
		CaptionPath:   videoDto.CaptionPath,
		Path:          videoDto.Path,
	}

	if video.ID == "" {
		video.ID = uuid.New().String()
	}

	v, err := h.Service.VideoService.CreateVideo(c, cID, video)
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, v)
}

func (h *Handler) UpdateVideo(c echo.Context) error {
	var videoDto CreateVideoRequest
	if err := c.Bind(&videoDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	if err := c.Validate(videoDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	vID := c.Param("id")

	cID := videoDto.ChannelID

	// Date
	parsedUploadDate, err := time.Parse("20060102", videoDto.UploadDate)
	if err != nil {
		return echo.NewHTTPError(400, errors.New(400, fmt.Sprintf("invalid upload date: %s", err.Error())))
	}

	video := &video.Video{
		Title:         videoDto.Title,
		Description:   videoDto.Description,
		UploadDate:    parsedUploadDate,
		Uploader:      videoDto.Uploader,
		Duration:      videoDto.Duration,
		ViewCount:     videoDto.ViewCount,
		LikeCount:     videoDto.LikeCount,
		DislikeCount:  videoDto.DislikeCount,
		Format:        videoDto.Format,
		Width:         videoDto.Width,
		Height:        videoDto.Height,
		Resolution:    videoDto.Resolution,
		FPS:           videoDto.FPS,
		AudioCodec:    videoDto.AudioCodec,
		VideoCodec:    videoDto.VideoCodec,
		ABR:           videoDto.ABR,
		VBR:           videoDto.VBR,
		Epoch:         videoDto.Epoch,
		CommentCount:  videoDto.CommentCount,
		Tags:          videoDto.Tags,
		Categories:    videoDto.Categories,
		VideoPath:     videoDto.VideoPath,
		ThumbnailPath: videoDto.ThumbnailPath,
		JSONPath:      videoDto.JSONPath,
		CaptionPath:   videoDto.CaptionPath,
		Path:          videoDto.Path,
	}

	v, err := h.Service.VideoService.UpdateVideo(c, vID, cID, video)
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, v)

}

func (h *Handler) DeleteVideo(c echo.Context) error {
	err := h.Service.VideoService.DeleteVideo(c, c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}
	return c.JSON(200, "deleted")
}

func (h *Handler) GetVideosByChannelID(c echo.Context) error {
	videos, err := h.Service.VideoService.GetVideosByChannelID(c, c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}
	return c.JSON(200, videos)
}

func (h *Handler) GetRandomVideos(c echo.Context) error {
	queryNumber := c.QueryParam("number")
	if queryNumber == "" {
		queryNumber = "10"
	}
	number, err := strconv.Atoi(queryNumber)
	if err != nil {
		return echo.NewHTTPError(400, errors.New(400, "invalid limit"))
	}

	videos, err := h.Service.VideoService.GetRandomVideos(c, number)
	if err != nil {
		return echo.NewHTTPError(500, errors.New(500, err.Error()))
	}
	return c.JSON(200, videos)
}
