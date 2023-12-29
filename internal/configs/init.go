package configs

import (
	"fmt"
	"jti-test/controller"
	"jti-test/internal/repository"
	"jti-test/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitializeViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

var (
	phoneNumberRepo    repository.PhoneNumberRepository
	providerRepo       repository.ProviderRepository
	phoneNumberService services.PhoneNumberService
	providerService    services.ProviderService
)

func Setup() {
	CreateRepositories()
	CreateServices()
	CreateController()
}

func CreateRepositories() {
	phoneNumberRepo = repository.NewPhoneNumberRepository(GetDB())
	providerRepo = repository.NewProviderRepository(GetDB())
}

func CreateServices() {
	phoneNumberService = services.NewPhoneNumberService(phoneNumberRepo)
	providerService = services.NewProviderService(providerRepo)
}

func CreateController() {
	controller.NewPhoneNumberController(phoneNumberService, providerService)
}

func InitializeFiber(port string, db *gorm.DB) {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws", websocket.New(HandleConnections))

	app.Get("/", services.HandleMain)
	app.Get("/form", services.HandleFormPage)
	app.Get("/dashboard", services.HandleDashboardPage)
	app.Get("/update-phonenumber", services.HandleUpdatePage)

	app.Get("/login-gl", services.HandleGoogleLogin)
	app.Get("/callback-gl", services.CallBackFromGoogle)
	app.Post("/add-phonenumber", controller.AddNewPhoneNumber)
	app.Get("/provider", controller.GetAllProvider)
	app.Get("/phonenumber", controller.GetAllPhoneNumber)
	app.Post("/phonenumber/:id", controller.UpdatePhoneNumber)
	app.Delete("/phonenumber/:id", controller.DeletePhoneNumber)
	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
