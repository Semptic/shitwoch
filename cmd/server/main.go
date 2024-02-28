package main

import (
	"fmt"
	"net/http"
	"os"
  "time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Semptic/shitwoch/cmd/server/templates"
)

func main() {
	host := os.Getenv("IP")
	port := os.Getenv("PORT")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "4000"
	}

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", IndexHandler)

	app.Logger.Fatal(app.Start(fmt.Sprintf("%s:%s", host, port)))
}

func IndexHandler(ctx echo.Context) error {
  weekday := time.Now().Weekday() 

  wednesday := weekday == time.Wednesday

	return Render(ctx, http.StatusOK, view.Index(wednesday))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
