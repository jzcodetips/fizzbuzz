package handler

import (
	"fizzbuzz/pkg/core/metrics"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Metrics handler.
func Metrics(ctx echo.Context) error {
	c, err := metrics.GetCollectedByName("/fizzbuzz")
	if err != nil {
		return ctx.String(http.StatusOK, err.Error())
	}

	return ctx.JSON(http.StatusOK, c)
}
