package comment

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/ent"
	entComment "github.com/zibbp/eos/ent/comment"
	entVideo "github.com/zibbp/eos/ent/video"
	"github.com/zibbp/eos/internal/database"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type Comment struct {
	ID               string    `json:"id"`
	Text             string    `json:"text"`
	Timestamp        time.Time `json:"timestamp"`
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

type Pagination struct {
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
	Pages      int       `json:"pages"`
	Data       []Comment `json:"data"`
}

func (s *Service) GetComments(c echo.Context, limit int, offset int, videoId string) (Pagination, error) {
	var pagination Pagination

	// Get all comments
	com, err := database.DB().Client.Comment.Query().Where(entComment.HasVideoWith(entVideo.ID(videoId))).Order(ent.Desc(entComment.FieldLikeCount)).All(c.Request().Context())
	if err != nil {
		return pagination, fmt.Errorf("failed to get comments: %v", err)
	}

	var comments []Comment
	for _, comment := range com {
		comments = append(comments, Comment{
			ID:               comment.ID,
			Text:             comment.Text,
			Timestamp:        comment.Timestamp,
			LikeCount:        comment.LikeCount,
			IsFavorited:      comment.IsFavorited,
			Author:           comment.Author,
			AuthorID:         comment.AuthorID,
			AuthorThumbnail:  comment.AuthorThumbnail,
			AuthorIsUploader: comment.AuthorIsUploader,
			Parent:           comment.Parent,
		})
	}

	// Add child comments to parent comments
	for i, comment := range comments {
		if comment.Parent != "root" {
			for j, parentComment := range comments {
				if parentComment.ID == comment.Parent {
					comments[j].Replies = append(comments[j].Replies, comments[i])
				}
			}
		}
	}

	// Pagination
	pagination.Limit = limit
	pagination.Offset = offset
	pagination.TotalCount = len(comments)
	pagination.Pages = int(math.Ceil(float64(len(comments)) / float64(limit)))
	pagination.Data = comments[offset : offset+limit]

	return pagination, nil

	// // Query builder
	// query := database.DB().Client.Comment.Query()

	// // Filter by video id
	// if videoId != "" {
	// 	query = query.Where(entComment.HasVideoWith(entVideo.ID(videoId)))
	// }
	// com, err := query.Order(ent.Desc(entComment.FieldLikeCount)).Limit(limit).Offset(offset).All(c.Request().Context())
	// if err != nil {
	// 	return pagination, fmt.Errorf("failed to get comments: %v", err)
	// }

	// // Get total count
	// totalCountQuery := database.DB().Client.Comment.Query()
	// if videoId != "" {
	// 	totalCountQuery = totalCountQuery.Where(entComment.HasVideoWith(entVideo.ID(videoId)))
	// }
	// totalCount, err := totalCountQuery.Count(c.Request().Context())
	// if err != nil {
	// 	return pagination, fmt.Errorf("failed to get total count: %v", err)
	// }

	// pagination.Limit = limit
	// pagination.Offset = offset
	// pagination.TotalCount = totalCount
	// pagination.Pages = int(math.Ceil(float64(totalCount) / float64(limit)))
	// pagination.Data = com

	// return pagination, nil
}

func ScannerCreateComment(commentDto *Comment, videoID string) error {
	_, err := database.DB().Client.Comment.Create().
		SetID(commentDto.ID).
		SetText(commentDto.Text).
		SetTimestamp(commentDto.Timestamp).
		SetLikeCount(commentDto.LikeCount).
		SetIsFavorited(commentDto.IsFavorited).
		SetAuthor(commentDto.Author).
		SetAuthorID(commentDto.AuthorID).
		SetAuthorThumbnail(commentDto.AuthorThumbnail).
		SetAuthorIsUploader(commentDto.AuthorIsUploader).
		SetParent(commentDto.Parent).
		SetVideoID(videoID).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create comment: %v", err)
	}
	return nil
}
