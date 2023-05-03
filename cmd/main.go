package main

import (
	"context"

	"github.com/guil95/flights/internal/flights"
	"github.com/guil95/flights/internal/flights/infra/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()

	e.POST("/calculate", server.HandleFlights(flights.Flights{}, ctx))

	e.Logger.Fatal(e.Start(":8080"))
}
