package main

import (
  "net"
	"fmt"
  "log"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Semptic/shitwoch/cmd/server/templates"
)

type State struct {
  db *GeoIpDb
}

func main() {
	host := os.Getenv("IP")
	port := os.Getenv("PORT")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "4000"
	}

  db, err := NewDb()

  if err != nil {
    log.Panic(err)
  }

  defer db.Close()

  state := State{db: db}

	app := echo.New()

  // app.IPExtractor = echo.ExtractIPDirect()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", state.IndexHandler)
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

func (state *State) IndexHandler(ctx echo.Context) error {
  ipString := ctx.RealIP()
  ip := net.ParseIP(ipString)

  timezone, err := state.db.TimezoneFromIp(ip)

  if err != nil {
    log.Println(err)
  }

  if timezone == "" {
    timezone = "UTC"
  }

  log.Println("Timezone:", timezone)

	return Render(ctx, http.StatusOK, view.Index(isShitwoch(timezone)))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
