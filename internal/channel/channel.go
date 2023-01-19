package channel

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/ent"
	entChannel "github.com/zibbp/eos/ent/channel"
	"github.com/zibbp/eos/internal/database"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type Channel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (s *Service) GetChannels(c echo.Context) ([]*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.Query().Order(ent.Desc(entChannel.FieldCreatedAt)).All(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get channels: %v", err)
	}

	return ch, nil
}

func (s *Service) GetChannel(c echo.Context, id string) (*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.Query().Where(entChannel.ID(id)).Only(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %v", err)
	}

	return ch, nil
}

func (s *Service) GetChannelByName(c echo.Context, name string) (*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.Query().Where(entChannel.Name(name)).Only(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %v", err)
	}

	return ch, nil
}

func (s *Service) CreateChannel(c echo.Context, channelDto *Channel) (*ent.Channel, error) {

	ch, err := database.DB().Client.Channel.Create().SetID(channelDto.ID).SetName(channelDto.Name).SetDescription(channelDto.Description).SetImagePath(channelDto.ImagePath).Save(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %v", err)
	}

	return ch, nil

}

func (s *Service) UpdateChannel(c echo.Context, channelDto *Channel) (*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.UpdateOneID(channelDto.ID).SetName(channelDto.Name).SetDescription(channelDto.Description).SetImagePath(channelDto.ImagePath).Save(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("failed to update channel: %v", err)
	}

	return ch, nil
}

func (s *Service) DeleteChannel(c echo.Context, id string) error {
	err := database.DB().Client.Channel.DeleteOneID(id).Exec(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to delete channel: %v", err)
	}

	return nil
}

func ScannerGetChannels() ([]*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.Query().All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get channels: %v", err)
	}

	return ch, nil
}

func ScannerGetChannel(id string) (*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.Query().Where(entChannel.ID(id)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get channel: %v", err)
	}

	return ch, nil
}

func ScannerCreateChannel(channelDto *Channel) (*ent.Channel, error) {
	ch, err := database.DB().Client.Channel.Create().SetID(channelDto.ID).SetName(channelDto.Name).SetDescription(channelDto.Description).SetImagePath(channelDto.ImagePath).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %v", err)
	}

	return ch, nil

}
