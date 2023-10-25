package playback

import (
	"context"

	"github.com/zibbp/eos/ent"
	"github.com/zibbp/eos/ent/playback"
	"github.com/zibbp/eos/internal/database"
)

func UpdateProgress(videoId string, timestamp int, status string, ctx context.Context) error {
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

	newStatus := playback.StatusInProgress
	if status == "finished" {
		newStatus = playback.StatusFinished
	}

	if pb != nil {
		_, err := pb.Update().SetTimestamp(timestamp).SetStatus(newStatus).Save(ctx)
		if err != nil {
			return err
		}
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

func UpdateStatus(videoId string, status string, ctx context.Context) error {
	newStatus := playback.StatusInProgress
	if status == "finished" {
		newStatus = playback.StatusFinished
	}
	_, err := database.DB().Client.Playback.Update().Where(playback.VideoID(videoId)).SetStatus(newStatus).Save(ctx)
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
