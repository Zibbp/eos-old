package playback

import (
	"context"

	"github.com/zibbp/eos/ent"
	"github.com/zibbp/eos/ent/playback"
	"github.com/zibbp/eos/internal/database"
	"github.com/zibbp/eos/internal/utils"
)

func UpdateProgress(videoId string, timestamp int, ctx context.Context) error {
	pb, err := database.DB().Client.Playback.Query().Where(playback.VideoID(videoId)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			_, err := database.DB().Client.Playback.Create().SetVideoID(videoId).SetTimestamp(timestamp).Save(ctx)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	if pb != nil {
		pb.Update().SetTimestamp(timestamp).Save(ctx)
	}

	return nil
}

func GetProgress(videoId string, ctx context.Context) (*ent.Playback, error) {
	pb, err := database.DB().Client.Playback.Query().Where(playback.VideoID(videoId)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return pb, nil
}

func GetAllProgress(ctx context.Context) ([]*ent.Playback, error) {
	pb, err := database.DB().Client.Playback.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return pb, nil
}

func UpdateStatus(videoId string, status utils.PlaybackStatus, ctx context.Context) error {
	_, err := database.DB().Client.Playback.Update().Where(playback.VideoID(videoId)).SetStatus(status).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProgress(videoId string, ctx context.Context) error {
	_, err := database.DB().Client.Playback.Delete().Where(playback.VideoID(videoId)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
