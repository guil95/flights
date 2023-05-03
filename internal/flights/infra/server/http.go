package server

import (
	"context"
	"net/http"

	"github.com/guil95/flights/internal/flights"
	"github.com/labstack/echo/v4"
)

type FlightsRequest struct {
	Route       [][]string `json:"route"`
	FlightsList [][]string `json:"flights"`
}

func HandleFlights(service flights.Flights, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req FlightsRequest

		if (&echo.DefaultBinder{}).BindHeaders(c, &req) != nil {
			return c.NoContent(http.StatusUnprocessableEntity)
		}

		if c.Bind(&req) != nil {
			return c.NoContent(http.StatusUnprocessableEntity)
		}

		if len(req.Route) == 0 || len(req.FlightsList) == 0 {
			return c.NoContent(http.StatusUnprocessableEntity)
		}

		f := flights.Flights{FlightsList: req.FlightsList, Route: req.Route}

		path, err := f.CalculateRoute()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, path)
	}
}
