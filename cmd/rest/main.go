package main

import (
	"fmt"

	"github.com/frkntplglu/emir-backend/internal/config"
	"github.com/frkntplglu/emir-backend/internal/handlers"
	"github.com/frkntplglu/emir-backend/internal/repositories"
	"github.com/frkntplglu/emir-backend/internal/services"
	. "github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-fiber/swagger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect to database")
	}

	_, err = db.DB()
	if err != nil {
		fmt.Println("failed to connect to database")
	}

	app := fiber.New(fiber.Config{
		ProxyHeader: fiber.HeaderXForwardedFor,
	})

	// Swagger
	sw := CreateNewSwagger("5000 Words API", "1.0")
	AddEndpoints(handlers.UserSwaggerEndpoints)
	AddEndpoints(handlers.WordSwaggerEndpoints)

	swagger.SwaggerHandler(app, sw.GenerateDocs(), swagger.Config{Prefix: "/swagger"})

	// Repositories
	wordRepository := repositories.NewWordRepository(db)
	userRepository := repositories.NewUserRepository(db)
	savedWordRepository := repositories.NewSavedWordRepository(db)

	// Services
	wordService := services.NewWordService(*wordRepository)
	userService := services.NewUserService(*userRepository)
	savedWordService := services.NewSavedWordService(*savedWordRepository)

	// Handlers
	wordHandler := handlers.NewWordHandler(*wordService)
	userHandler := handlers.NewUserHandler(*userService)
	savedWordHandler := handlers.NewSavedWordHandler(*savedWordService, *wordService)

	// Routes
	wordHandler.SetRoutes(app)
	userHandler.SetRoutes(app)
	savedWordHandler.SetRoutes(app)

	// App Starting
	app.Listen(":9000")

}
