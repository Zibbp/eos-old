package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

type EosMetricsCollector struct{}

// Metrics
var (
	totalVideos = prometheus.NewDesc(
		prometheus.BuildFQName("eos", "video", "count"),
		"Total number of videos",
		nil, nil,
	)

	totalChannels = prometheus.NewDesc(
		prometheus.BuildFQName("eos", "channel", "count"),
		"Total number of channels",
		nil, nil,
	)

	channelVideoCount = prometheus.NewDesc(
		prometheus.BuildFQName("eos", "channel", "video_count"),
		"Number of videos in a channel",
		[]string{"channel"}, nil,
	)
)

func (emc *EosMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(emc, ch)
}

func (emc *EosMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	// Gather data
	videoCount, err := getVideoCount()
	if err != nil {
		log.Error().Err(err).Msg("failed to get video count")
	}

	ch <- prometheus.MustNewConstMetric(totalVideos, prometheus.GaugeValue, float64(videoCount))

	channelCount, err := getChannelCount()
	if err != nil {
		log.Error().Err(err).Msg("failed to get channel count")
	}

	ch <- prometheus.MustNewConstMetric(totalChannels, prometheus.GaugeValue, float64(channelCount))

	channelVideoCountMap, err := getChannelVideoCount()
	if err != nil {
		log.Error().Err(err).Msg("failed to get channel video count")
	}

	for channel, count := range channelVideoCountMap {
		ch <- prometheus.MustNewConstMetric(channelVideoCount, prometheus.GaugeValue, float64(count), channel)
	}
}

func NewEosMetricsCollector() *EosMetricsCollector {
	return &EosMetricsCollector{}
}
