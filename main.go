package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/khodemobin/learn_go/router"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	serverMux := mux.NewRouter()

	//register routes
	router.RegisterRoutes(serverMux, logger)

	server := &http.Server{
		Addr:         ":" + getPort(),
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	logger.Println("Recieved terminate , graceful shoutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if cancel != nil {
		panic(cancel)
	}
	server.Shutdown(ctx)
}

func getPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		return "8000"
	}

	return port
}
