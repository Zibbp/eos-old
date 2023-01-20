package metrics

import (
	"context"
	"fmt"

	entChannel "github.com/zibbp/eos/ent/channel"
	entVideo "github.com/zibbp/eos/ent/video"
	"github.com/zibbp/eos/internal/database"
)

func getVideoCount() (int, error) {
	count, err := database.DB().Client.Video.Query().Count(context.Background())
	if err != nil {
		return 0, fmt.Errorf("failed to get video count: %v", err)
	}

	return count, nil
}

func getChannelCount() (int, error) {
	count, err := database.DB().Client.Channel.Query().Count(context.Background())
	if err != nil {
		return 0, fmt.Errorf("failed to get channel count: %v", err)
	}

	return count, nil
}

func getChannelVideoCount() (map[string]int, error) {
	channels, err := database.DB().Client.Channel.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get channels: %v", err)
	}

	channelVideoCount := make(map[string]int)
	for _, channel := range channels {
		count, err := database.DB().Client.Video.Query().Where(entVideo.HasChannelWith(entChannel.ID(channel.ID))).Count(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get video count for channel %s: %v", channel.ID, err)
		}

		channelVideoCount[channel.Name] = count
	}

	return channelVideoCount, nil
}
