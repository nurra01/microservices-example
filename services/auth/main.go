package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"services/auth/db"
	"services/auth/handlers"
	"services/auth/logger"
	"services/auth/redis"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// init a logger
	log := logger.NewLogger()

	// load .env file
	err := godotenv.Load("auth-service.env")
	if err != nil {
		log.Fatal("failed to load auth-service.env file")
	}

	port := os.Getenv("PORT") // service run port
	if port == "" {
		port = "8081"
	}

	// check secret key env variable
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("missing JWT_SECRET env variable")
	}

	// configure db connection
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// configure redis client
	err = redis.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	// create handlers
	ah := handlers.NewAuthHandler(log)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// POST requests
	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.Handle("/user/login", ah.MiddlewareValidateLogin(http.HandlerFunc(ah.Login)))

	// GET requests
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.Handle("/auth/token", ah.MiddlewareValidateAccessToken(http.HandlerFunc(ah.AccessToken)))
	getR.Handle("/users/{id}/profile", ah.MiddlewareValidateUserProfile(http.HandlerFunc(ah.UserProfile)))

	// CORS handler
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Origin", "Authorization"},
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
