package scanner

import (
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/internal/channel"
	"github.com/zibbp/eos/internal/kv"
	"github.com/zibbp/eos/internal/video"
)

type Service struct {
	ChannelService *channel.Service
	VideoService   *video.Service
}

func NewService(channelService *channel.Service, videoService *video.Service) *Service {
	return &Service{
		ChannelService: channelService,
		VideoService:   videoService,
	}
}

func (s *Service) Scan() (string, error) {
	// Check if scanner is already running
	checkScanner, ok := kv.DB().Get("scanner")
	if ok && checkScanner == "running" {
		return "scanner is already running", nil
	}

	// Set scanner to running
	kv.DB().Set("scanner", "running")

	// go routine to scan channels
	go func() {
		err := s.StartScanner()
		if err != nil {
			log.Error().Err(err).Msg("failed to start scanner")
		}
	}()

	return "scanner started", nil
}
