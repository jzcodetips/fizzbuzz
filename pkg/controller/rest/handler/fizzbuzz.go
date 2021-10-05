package handler

import (
	"errors"
	"fizzbuzz/pkg/core/metrics"
	"fizzbuzz/pkg/core/service"
	"fizzbuzz/pkg/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// errors.
var (
	ErrServiceFizzBuzzRequired = errors.New("service fizzbuzz is required")
)

// FizzBuzz handler struct.
type FizzBuzz struct {
	service *service.FizzBuzz
}

// NewFizzBuzz returns a new instance of FizzBuzz.
func NewFizzBuzz(service *service.FizzBuzz) (*FizzBuzz, error) {
	if service == nil {
		return nil, ErrServiceFizzBuzzRequired
	}

	f := &FizzBuzz{
		service: service,
	}

	return f, nil
}

// Handle FizzBuzz.
func (f *FizzBuzz) Handle(ctx echo.Context) error {
	p, err := parseQueryParam(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entity.H{"message": err.Error()})
	}

	metrics.CollectFizzBuzz("/fizzbuzz", *p)

	return ctx.JSON(http.StatusOK, f.service.Process(p.Int1, p.Int2, p.Limit, p.Str1, p.Str2))
}

func parseQueryParam(ctx echo.Context) (*entity.FizzBuzzParams, error) {
	int1, err := strconv.Atoi(ctx.QueryParam("int1"))
	if err != nil {
		return nil, errors.New("int1 is required and should be an integer")
	}

	int2, err := strconv.Atoi(ctx.QueryParam("int2"))
	if err != nil {
		return nil, errors.New("int2 is required and should be an integer")
	}

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		return nil, errors.New("limit is required and should be an integer")
	}

	str1 := ctx.QueryParam("str1")
	if str1 == "" {
		return nil, errors.New("str1 is required")
	}

	str2 := ctx.QueryParam("str2")
	if str2 == "" {
		return nil, errors.New("str2 is required")
	}

	return &entity.FizzBuzzParams{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}, nil
}
