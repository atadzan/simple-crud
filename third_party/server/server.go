package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(app *fiber.App, port uint16) {
	idleConnClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnClosed)
	}()

	// Run server.
	if err := app.Listen(fmt.Sprint(":", port)); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnClosed
}
