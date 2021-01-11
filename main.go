package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/khodemobin/learn_go/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	goodbyeHandler := handlers.NewGoodbye(logger)
	productsHandler := handlers.NewProducts(logger)

	serverMux := http.NewServeMux()

	serverMux.Handle("/", helloHandler)
	serverMux.Handle("/products", productsHandler)

	serverMux.Handle("/bye", goodbyeHandler)

	server := &http.Server{
		Addr:         ":8000",
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

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	if err != nil {
		panic(err)
	}
	server.Shutdown(tc)
}
