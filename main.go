package main

import (
	"fmt"

	"github.com/frkntplglu/emir-backend/internal/handlers"
	"github.com/frkntplglu/emir-backend/internal/repositories"
	"github.com/frkntplglu/emir-backend/internal/services"
	. "github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-fiber/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", "database-1.cpngaskb6zpb.us-east-1.rds.amazonaws.com", "postgres", "milo2023", "postgres", 5432)
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

	app.Use(cors.New())

	// Swagger
	sw := CreateNewSwagger("5000 Words API", "1.0")
	AddEndpoints(handlers.UserSwaggerEndpoints)
	AddEndpoints(handlers.WordSwaggerEndpoints)

	swagger.SwaggerHandler(app, sw.GenerateDocs(), swagger.Config{Prefix: "/swagger"})

	// Repositories
	wordRepository := repositories.NewWordRepository(db)
	userRepository := repositories.NewUserRepository(db)
	savedWordRepository := repositories.NewSavedWordRepository(db)
	quizRepository := repositories.NewQuizRepository(db)
	questionRepository := repositories.NewQuestionRepository(db)

	// Services
	wordService := services.NewWordService(*wordRepository)
	userService := services.NewUserService(*userRepository)
	savedWordService := services.NewSavedWordService(*savedWordRepository)
	quizService := services.NewQuizService(*quizRepository)
	questionService := services.NewQuestionService(*questionRepository)

	// Handlers
	wordHandler := handlers.NewWordHandler(*wordService)
	userHandler := handlers.NewUserHandler(*userService)
	savedWordHandler := handlers.NewSavedWordHandler(*savedWordService, *wordService)
	quizHandler := handlers.NewQuizHandler(*quizService)
	questionHandler := handlers.NewQuestionHandler(*questionService)

	// Routes
	wordHandler.SetRoutes(app)
	userHandler.SetRoutes(app)
	savedWordHandler.SetRoutes(app)
	quizHandler.SetRoutes(app)
	questionHandler.SetRoutes(app)

	// App Starting
	app.Listen("0.0.0.0:9000")

}
