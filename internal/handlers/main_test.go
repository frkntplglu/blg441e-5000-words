package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/frkntplglu/emir-backend/internal/repositories"
	"github.com/frkntplglu/emir-backend/internal/services"
	container "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestingSetup struct {
	testDb             *gorm.DB
	wordHandler        *WordHandler
	quizHandler        *QuizHandler
	questionHandler    *QuestionHandler
	wordService        *services.WordService
	quizService        *services.QuizService
	questionService    *services.QuestionService
	wordRepository     *repositories.WordRepository
	quizRepository     *repositories.QuizRepository
	questionRepository *repositories.QuestionRepository
	cleanup            func()
}

var testingSetup TestingSetup

func TestMain(m *testing.M) {
	initDbQuery, err := os.ReadFile("../../db.sql")
	if err != nil {
		log.Fatalf("Error reading db.sql file: %v", err)
	}
	postgresPort := nat.Port("5432/tcp")
	postgresContainer, err := container.GenericContainer(context.Background(),
		container.GenericContainerRequest{
			ContainerRequest: container.ContainerRequest{
				Image:        "postgres:14-alpine",
				ExposedPorts: []string{postgresPort.Port()},
				Env: map[string]string{
					"POSTGRES_PASSWORD": "POSTGRES_PASSWORD",
					"POSTGRES_USER":     "POSTGRES_USER",
					"POSTGRES_DB":       "POSTGRES_DB",
				},
				WaitingFor: wait.ForAll(
					wait.ForLog("database system is ready to accept connections"),
					wait.ForListeningPort(postgresPort),
				),
				SkipReaper: true,
			},
			Started: true,
			Reuse:   false,
		})

	if err != nil {
		log.Fatalf("GenericContainer error: %v", err)
	}

	hostPort, err := postgresContainer.MappedPort(context.Background(), postgresPort)
	if err != nil {
		log.Fatalf("Container host port map error: %v", err)
	}
	containerPort, _ := strconv.Atoi(hostPort.Port())

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", "127.0.0.1", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB", containerPort)

	testingSetup.testDb, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Connection failed error: %v", err.Error())
	}

	execQuery := testingSetup.testDb.Exec(string(initDbQuery))
	if execQuery.Error != nil {
		log.Fatalf("Init db query exec error: %v", execQuery.Error)
	}

	testingSetup.wordRepository = repositories.NewWordRepository(testingSetup.testDb)
	testingSetup.quizRepository = repositories.NewQuizRepository(testingSetup.testDb)
	testingSetup.questionRepository = repositories.NewQuestionRepository(testingSetup.testDb)
	var wordRepository = repositories.NewWordRepository(testingSetup.testDb)
	var quizRepository = repositories.NewQuizRepository(testingSetup.testDb)
	var QuestionRepository = repositories.NewQuestionRepository(testingSetup.testDb)

	testingSetup.wordService = services.NewWordService(*wordRepository)
	testingSetup.quizService = services.NewQuizService(*quizRepository)
	testingSetup.questionService = services.NewQuestionService(*QuestionRepository)

	testingSetup.wordHandler = NewWordHandler(*testingSetup.wordService)
	testingSetup.quizHandler = NewQuizHandler(*testingSetup.quizService)
	testingSetup.questionHandler = NewQuestionHandler(*testingSetup.questionService)

	os.Exit(m.Run())
}
