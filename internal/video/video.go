package video

import (
	"context"
	"fmt"
	"math"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/ent"
	entChannel "github.com/zibbp/eos/ent/channel"
	entVideo "github.com/zibbp/eos/ent/video"
	"github.com/zibbp/eos/internal/database"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type Video struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	UploadDate    time.Time `json:"upload_date"`
	Uploader      string    `json:"uploader"`
	Duration      int64     `json:"duration"`
	ViewCount     int64     `json:"view_count"`
	LikeCount     int64     `json:"like_count"`
	DislikeCount  int64     `json:"dislike_count"`
	Format        string    `json:"format"`
	Width         int64     `json:"width"`
	Height        int64     `json:"height"`
	Resolution    string    `json:"resolution"`
	FPS           float64   `json:"fps"`
	AudioCodec    string    `json:"audio_codec"`
	VideoCodec    string    `json:"video_codec"`
	ABR           float64   `json:"abr"`
	VBR           float64   `json:"vbr"`
	Epoch         int64     `json:"epoch"`
	CommentCount  int64     `json:"comment_count"`
	Tags          string    `json:"tags"`
	Categories    string    `json:"categories"`
	VideoPath     string    `json:"video_path"`
	ThumbnailPath string    `json:"thumbnail_path"`
	JSONPath      string    `json:"json_path"`
	CaptionPath   string    `json:"caption_path"`
	Path          string    `json:"path"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

type Chapter struct {
	ID        string  `json:"id"`
	StartTime float64 `json:"start_time"`
	Title     string  `json:"title"`
	EndTime   float64 `json:"end_time"`
	VideoID   string  `json:"video_id"`
}

type Pagination struct {
	Offset     int          `json:"offset"`
	Limit      int          `json:"limit"`
	TotalCount int          `json:"total_count"`
	Pages      int          `json:"pages"`
	Data       []*ent.Video `json:"data"`
}

func (s *Service) GetVideos(c echo.Context, limit int, offset int, channelId string) (Pagination, error) {
	var pagination Pagination

	// Query builder
	query := database.DB().Client.Video.Query()

	// Filter by channel id
	if channelId != "" {
		query = query.Where(entVideo.HasChannelWith(entChannel.ID(channelId)))
	}

	v, err := query.Order(ent.Desc(entVideo.FieldUploadDate)).Limit(limit).Offset(offset).All(c.Request().Context())
	if err != nil {
		return pagination, fmt.Errorf("failed to get videos: %v", err)
	}

	// Get total count
	totalCountQuery := database.DB().Client.Video.Query()
	if channelId != "" {
		totalCountQuery = totalCountQuery.Where(entVideo.HasChannelWith(entChannel.ID(channelId)))
	}
	totalCount, err := totalCountQuery.Count(context.Background())
	if err != nil {
		return pagination, fmt.Errorf("failed to get total count: %v", err)
	}

	pagination.Limit = limit
	pagination.Offset = offset
	pagination.TotalCount = totalCount
	pagination.Pages = int(math.Ceil(float64(totalCount) / float64(limit)))
	pagination.Data = v

	return pagination, nil
}

func (s *Service) SearchVideos(c echo.Context, limit int, offset int, query string) (Pagination, error) {
	var pagination Pagination

	// Query builder
	q := database.DB().Client.Video.Query()

	// Filter by channel id
	if query != "" {
		q = q.Where(entVideo.TitleContainsFold(query))
	}

	v, err := q.Order(ent.Desc(entVideo.FieldUploadDate)).Limit(limit).Offset(offset).All(c.Request().Context())
	if err != nil {
		return pagination, fmt.Errorf("failed to get videos: %v", err)
	}

	// Get total count
	totalCountQuery := database.DB().Client.Video.Query()
	if query != "" {
		totalCountQuery = totalCountQuery.Where(entVideo.TitleContainsFold(query))
	}
	totalCount, err := totalCountQuery.Count(context.Background())
	if err != nil {
		return pagination, fmt.Errorf("failed to get total count: %v", err)
	}

	pagination.Limit = limit
	pagination.Offset = offset
	pagination.TotalCount = totalCount
	pagination.Pages = int(math.Ceil(float64(totalCount) / float64(limit)))
	pagination.Data = v

	return pagination, nil
}

func (s *Service) GetVideo(c echo.Context, id string) (*ent.Video, error) {
	video, err := database.DB().Client.Video.Query().Where(entVideo.ID(id)).WithChannel().WithChapters().Only(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get video: %v", err)
	}

	return video, nil
}

func (s *Service) CreateVideo(c echo.Context, channelID string, videoDto *Video) (*ent.Video, error) {

	video, err := database.DB().Client.Video.Create().SetID(videoDto.ID).SetChannelID(channelID).SetTitle(videoDto.Title).SetDescription(videoDto.Description).SetUploadDate(videoDto.UploadDate).SetUploader(videoDto.Uploader).SetDuration(videoDto.Duration).SetViewCount(videoDto.ViewCount).SetLikeCount(videoDto.LikeCount).SetDislikeCount(videoDto.DislikeCount).SetFormat(videoDto.Format).SetWidth(videoDto.Width).SetHeight(videoDto.Height).SetResolution(videoDto.Resolution).SetFps(videoDto.FPS).SetAudioCodec(videoDto.AudioCodec).SetVideoCodec(videoDto.VideoCodec).SetAbr(videoDto.ABR).SetVbr(videoDto.VBR).SetEpoch(videoDto.Epoch).SetCommentCount(videoDto.CommentCount).SetTags(videoDto.Tags).SetCategories(videoDto.Categories).SetVideoPath(videoDto.VideoPath).SetThumbnailPath(videoDto.ThumbnailPath).SetJSONPath(videoDto.JSONPath).SetCaptionPath(videoDto.CaptionPath).SetPath(videoDto.Path).Save(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to create video: %v", err)
	}

	return video, nil
}

func (s *Service) GetRandomVideos(c echo.Context, number int) ([]*ent.Video, error) {
	videos, err := database.DB().Client.Video.Query().Order(func(s *sql.Selector) {
		s.OrderBy("RANDOM()")
	}).Limit(number).All(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get videos: %v", err)
	}

	return videos, nil
}

func (s *Service) UpdateVideo(c echo.Context, id string, channelID string, videoDto *Video) (*ent.Video, error) {

	video, err := database.DB().Client.Video.UpdateOneID(id).SetChannelID(channelID).SetTitle(videoDto.Title).SetDescription(videoDto.Description).SetUploadDate(videoDto.UploadDate).SetUploader(videoDto.Uploader).SetDuration(videoDto.Duration).SetViewCount(videoDto.ViewCount).SetLikeCount(videoDto.LikeCount).SetDislikeCount(videoDto.DislikeCount).SetFormat(videoDto.Format).SetWidth(videoDto.Width).SetHeight(videoDto.Height).SetResolution(videoDto.Resolution).SetFps(videoDto.FPS).SetAudioCodec(videoDto.AudioCodec).SetVideoCodec(videoDto.VideoCodec).SetAbr(videoDto.ABR).SetVbr(videoDto.VBR).SetEpoch(videoDto.Epoch).SetCommentCount(videoDto.CommentCount).SetTags(videoDto.Tags).SetCategories(videoDto.Categories).SetVideoPath(videoDto.VideoPath).SetThumbnailPath(videoDto.ThumbnailPath).SetJSONPath(videoDto.JSONPath).SetCaptionPath(videoDto.CaptionPath).SetPath(videoDto.Path).Save(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to update video: %v", err)
	}

	return video, nil
}

func (s *Service) DeleteVideo(c echo.Context, id string) error {
	err := database.DB().Client.Video.DeleteOneID(id).Exec(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to delete video: %v", err)
	}

	return nil
}

func (s *Service) GetVideosByChannelID(c echo.Context, channelID string) ([]*ent.Video, error) {
	videos, err := database.DB().Client.Video.Query().Where(entVideo.HasChannelWith(entChannel.ID(channelID))).All(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get videos by channel id: %v", err)
	}

	return videos, nil
}

func ScannerGetVideos() ([]*ent.Video, error) {
	videos, err := database.DB().Client.Video.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get videos: %v", err)
	}

	return videos, nil
}

func ScannerCreateVideo(videoDto *Video, channelID string) error {
	_, err := database.DB().Client.Video.Create().SetID(videoDto.ID).SetChannelID(channelID).SetTitle(videoDto.Title).SetDescription(videoDto.Description).SetUploadDate(videoDto.UploadDate).SetUploader(videoDto.Uploader).SetDuration(videoDto.Duration).SetViewCount(videoDto.ViewCount).SetLikeCount(videoDto.LikeCount).SetDislikeCount(videoDto.DislikeCount).SetFormat(videoDto.Format).SetWidth(videoDto.Width).SetHeight(videoDto.Height).SetResolution(videoDto.Resolution).SetFps(videoDto.FPS).SetAudioCodec(videoDto.AudioCodec).SetVideoCodec(videoDto.VideoCodec).SetAbr(videoDto.ABR).SetVbr(videoDto.VBR).SetEpoch(videoDto.Epoch).SetCommentCount(videoDto.CommentCount).SetTags(videoDto.Tags).SetCategories(videoDto.Categories).SetVideoPath(videoDto.VideoPath).SetThumbnailPath(videoDto.ThumbnailPath).SetJSONPath(videoDto.JSONPath).SetCaptionPath(videoDto.CaptionPath).SetPath(videoDto.Path).Save(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create video: %v", err)
	}

	return nil
}

func ScannerCreateChapter(chapterDto *Chapter, videoID string) error {
	_, err := database.DB().Client.Chapter.Create().SetID(chapterDto.ID).SetVideoID(videoID).SetStartTime(chapterDto.StartTime).SetEndTime(chapterDto.EndTime).SetTitle(chapterDto.Title).SetTitle(chapterDto.Title).Save(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create chapter: %v", err)
	}

	return nil
}

func ScannerGetVideo(videoID string) (*ent.Video, error) {
	video, err := database.DB().Client.Video.Query().Where(entVideo.ID(videoID)).Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get video: %v", err)
	}

	return video, nil
}
