package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/ent"
	"github.com/zibbp/eos/internal/channel"
	"github.com/zibbp/eos/internal/errors"
)

type ChannelService interface {
	GetChannels(c echo.Context) ([]*ent.Channel, error)
	GetChannel(c echo.Context, id string) (*ent.Channel, error)
	CreateChannel(c echo.Context, channelDto *channel.Channel) (*ent.Channel, error)
	UpdateChannel(c echo.Context, channelDto *channel.Channel) (*ent.Channel, error)
	DeleteChannel(c echo.Context, id string) error
	GetChannelByName(c echo.Context, name string) (*ent.Channel, error)
}

type CreateChannelRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	ImagePath   string `json:"image_path" validate:"required"`
}

func (h *Handler) GetChannels(c echo.Context) error {
	channels, err := h.Service.ChannelService.GetChannels(c)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, channels)
}

func (h *Handler) GetChannel(c echo.Context) error {
	channel, err := h.Service.ChannelService.GetChannel(c, c.Param("id"))
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, channel)
}

func (h *Handler) GetChannelByName(c echo.Context) error {
	name := c.Param("name")
	channel, err := h.Service.ChannelService.GetChannelByName(c, name)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, channel)
}

func (h *Handler) CreateChannel(c echo.Context) error {
	var channelDto CreateChannelRequest
	if err := c.Bind(&channelDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	if err := c.Validate(channelDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	channel := &channel.Channel{
		ID:          channelDto.ID,
		Name:        channelDto.Name,
		Description: channelDto.Description,
		ImagePath:   channelDto.ImagePath,
	}

	if channel.ID == "" {
		channel.ID = uuid.New().String()
	}

	ch, err := h.Service.ChannelService.CreateChannel(c, channel)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, ch)
}

func (h *Handler) UpdateChannel(c echo.Context) error {
	var channelDto CreateChannelRequest
	if err := c.Bind(&channelDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	if err := c.Validate(channelDto); err != nil {
		return echo.NewHTTPError(400, errors.New(400, err.Error()))
	}

	channel := &channel.Channel{
		ID:          c.Param("id"),
		Name:        channelDto.Name,
		Description: channelDto.Description,
		ImagePath:   channelDto.ImagePath,
	}

	ch, err := h.Service.ChannelService.UpdateChannel(c, channel)
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, ch)
}

func (h *Handler) DeleteChannel(c echo.Context) error {
	err := h.Service.ChannelService.DeleteChannel(c, c.Param("id"))
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, "ok")
}
