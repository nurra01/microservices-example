package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"services/user/handlers"
	"services/user/kafka"
	"services/user/logger"
	"services/user/redis"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// init a logger
	log := logger.NewLogger()

	// load .env file
	err := godotenv.Load("user-service.env")
	if err != nil {
		log.Debug("failed to load user-service.env file")
	}

	port := os.Getenv("PORT") // service run port
	if port == "" {
		port = "8080" // if missing env variable, use default
	}

	// configure kafka writer
	w1, w2, err := kafka.Configure()
	if err != nil {
		log.Fatal(err)
	}
	// close kafka writer and connection when exit
	defer w1.Close()
	defer w2.Close()

	// configure redis client
	err = redis.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	// create the handlers
	uh := handlers.NewUserHandler(log)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// POST requests
	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/user/register", uh.Register)
	postR.Use(uh.MiddlewareValidateRegisterUser)

	// GET requests
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/user/verify/{id}", uh.Verify)

	// CORS handler
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	})
	ch := crs.Handler(sm)

	// create a new server
	s := http.Server{
		Addr:         fmt.Sprintf(":%s", port), // configure the bind address
		Handler:      ch,                       // set the default handler
		ReadTimeout:  5 * time.Second,          // max time to read request from the client
		WriteTimeout: 10 * time.Second,         // max time to write response to the client
		IdleTimeout:  120 * time.Second,        // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Infof("Starting server on port %v", port)

		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Info("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
