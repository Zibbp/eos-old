package scanner

import (
	"time"

	"github.com/zibbp/eos/internal/video"
)

const (
	VideoDirectory = "/videos"
)

type VideoInfo struct {
	ID             string          `json:"id"`
	Channel        string          `json:"channel"`
	ChannelID      string          `json:"channel_id"`
	Title          string          `json:"title"`
	Formats        []Format        `json:"formats"`
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
	Type           string          `json:"_type"`
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

type Format struct {
	FormatID   string     `json:"format_id"`
	FormatNote string     `json:"format_note"`
	Width      *int64     `json:"width,omitempty"`
	Height     *int64     `json:"height,omitempty"`
	FPS        *float64   `json:"fps,omitempty"`
	Rows       *int64     `json:"rows,omitempty"`
	Columns    *int64     `json:"columns,omitempty"`
	Fragments  []Fragment `json:"fragments,omitempty"`
}

type Fragment struct {
	URL      string  `json:"url"`
	Duration float64 `json:"duration"`
}
