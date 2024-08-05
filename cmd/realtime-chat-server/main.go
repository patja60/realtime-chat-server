package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patja60/realtime-chat-server/internal/app/auth"
	"github.com/spf13/viper"

	database "github.com/patja60/realtime-chat-server/pkg"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("fatal error config file: %s", err)
	}

	// Bind environment variables
	viper.BindEnv("db.host", "DB_HOST")
	viper.BindEnv("db.port", "DB_PORT")
	viper.BindEnv("db.user", "DB_USER")
	viper.BindEnv("db.password", "DB_PASSWORD")
	viper.BindEnv("db.name", "DB_NAME")
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
}

func main() {
	initConfig()
	// init db
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	// init repo
	authRepo := auth.NewAuthRepository(*db)

	// init usecase
	authUsecase := auth.NewAuthUsecase(authRepo)

	// init handler -> inject usecase
	authHandler := auth.NewAuthHandler(authUsecase)

	// init router -> use handler
	r := mux.NewRouter()

	r.HandleFunc("/api/signup", authHandler.Signup).Methods("POST")
	r.HandleFunc("/api/signin", authHandler.Signin).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
