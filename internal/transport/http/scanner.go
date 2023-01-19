package http

import (
	"github.com/labstack/echo/v4"
	"github.com/zibbp/eos/internal/errors"
)

type ScannerService interface {
	Scan() (string, error)
}

func (h *Handler) Scan(c echo.Context) error {
	msg, err := h.Service.ScannerService.Scan()
	if err != nil {
		return c.JSON(500, errors.New(500, err.Error()))
	}

	return c.JSON(200, msg)
}
