package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

// Metrics
var (
	totalVideos = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "video_count",
		Help: "Total number of videos",
	})
	totalChannels = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "channel_count",
		Help: "Total number of channels",
	})
	channelVideoCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "channel_video_count",
		Help: "Total number of videos per channel",
	}, []string{"channel"})
)

func GatherMetrics() *prometheus.Registry {
	// Gather data
	videoCount, err := getVideoCount()
	if err != nil {
		log.Error().Err(err).Msg("failed to get video count")
		totalVideos.Set(0)
	}
	totalVideos.Set(float64(videoCount))
	channelCount, err := getChannelCount()
	if err != nil {
		log.Error().Err(err).Msg("failed to get channel count")
		totalChannels.Set(0)
	}
	totalChannels.Set(float64(channelCount))
	channelVideoCountMap, err := getChannelVideoCount()
	if err != nil {
		log.Error().Err(err).Msg("failed to get channel video count")
	}
	for channel, count := range channelVideoCountMap {
		channelVideoCount.WithLabelValues(channel).Set(float64(count))
	}

	r := prometheus.NewRegistry()
	r.MustRegister(totalVideos, totalChannels, channelVideoCount)
	return r
}
