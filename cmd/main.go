package main

import (
	"jti-test/internal/configs"
	"jti-test/internal/services"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Oauth2 Services
	services.InitialzeGoogleOauth()

	// Initialize RDB
	configs.InitializeDatabase()

	configs.Setup()

	// Routes for the application
	configs.InitializeFiber(viper.GetString("port"), configs.GetDB())

	log.Println("Started running on http://localhost:" + viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}
