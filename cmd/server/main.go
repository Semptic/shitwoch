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
	app.GET("/is_shitwoch", ShitwochHandler)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}

func isShitwoch(timezone string) bool {
	if timezone != "" {
		loc, _ := time.LoadLocation(timezone)
		weekday := time.Now().In(loc).Weekday()

		return weekday == time.Wednesday
	}

	weekday := time.Now().Weekday()

	return weekday == time.Wednesday
}

func ShitwochHandler(ctx echo.Context) error {
	timezone := ctx.QueryParam("timezone")

	return Render(ctx, http.StatusOK, view.Shitwoch(isShitwoch(timezone)))
}

func IndexHandler(ctx echo.Context) error {
	return Render(ctx, http.StatusOK, view.Index(isShitwoch("UTC")))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
