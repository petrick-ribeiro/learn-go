/* 
 *  Refactoring the API and implementing the gorilla/mux framework.
 */

package main

import (
	"coffee-shop/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// Create the handlers
	ph := handlers.NewProducts(l)

	// Create a new serve mux and register the handlers
	// sm := http.NewServeMux()
  sm := mux.NewRouter()

	// sm.Handle("/", ph)
  getRouter := sm.Methods(http.MethodGet).Subrouter()
  getRouter.HandleFunc("/", ph.GetProducts)

  putRouter := sm.Methods(http.MethodPut).Subrouter()
  putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)

  postRouter := sm.Methods(http.MethodPost).Subrouter()
  postRouter.HandleFunc("/", ph.AddProduct)

	// Create the Server
	srv := http.Server{
		Addr:         ":9090",           // configure the bind addr
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read the request
		WriteTimeout: 10 * time.Second,  // max time to write response
		IdleTimeout:  120 * time.Second, // max time for the connection using TCP Keep-Alive
	}

	// Start the Server
	go func() {
		log.Println("Starting server on port :9090")

		err := srv.ListenAndServe()
		if err != nil {
			log.Printf("Error starting the server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Gracefully shutdown the Server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
