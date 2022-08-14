package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taoyuan_carpark/logging"
	"taoyuan_carpark/service"
	"time"
)

const defaultPort = ":8080"

func main() {
	app := service.NewGin()

	service.Init(app)

	srv := &http.Server{
		Addr:    defaultPort,
		Handler: app,
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Print(zerolog.InfoLevel, "", "Server Fail")
			quit <- syscall.SIGINT
		}
	}()

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Print(zerolog.InfoLevel, "", "Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		errMsg := fmt.Sprintf("Server Shutdown: %s", err)
		logging.Print(zerolog.ErrorLevel, "", errMsg)
	}

	logging.Print(zerolog.InfoLevel, "", "Server exiting")
}
