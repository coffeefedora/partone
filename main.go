package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coffeefedora/partone/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	port string = ":9000"
)

func main() {
	stopChan := make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// each controller in the controllers file is responsible
	// for defining its supported routes.
	// this routine will load them up
	loadRoutes(r)

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

	defer cancel()
	log.Println("Server gracefully stopped!")
}

func loadRoutes(r *chi.Mux) {
	r.Mount("/", controllers.RootRouter)
	controllers.StaticInitializeRoute(r) // serve up public css, js, etc.
}
