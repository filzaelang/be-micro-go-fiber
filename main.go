package main

import (
	"log"
	"net/http"
	"os"

	"github.com/filzaelang/be-micro-go-fiber/models"
	"github.com/filzaelang/be-micro-go-fiber/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetArticles(context *fiber.Ctx) error {
	articleModels := &[]models.Article{}

	err := r.DB.Find(articleModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get articles"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "articles fetched successfully",
		"data":    articleModels,
	})
	return nil
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	userModels := &[]models.User{}

	err := r.DB.Find(userModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users fetched successfully",
		"data":    userModels,
	})
	return nil
}

func (r *Repository) GetPartai(context *fiber.Ctx) error {
	partaiModels := &[]models.Partai{}

	err := r.DB.Debug().Find(partaiModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "partai fetched successfully",
		"data":    partaiModels,
	})
	return nil
}

func (r *Repository) GetPaslon(context *fiber.Ctx) error {
	paslonModels := &[]models.Paslon{}

	err := r.DB.Debug().Find(paslonModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "paslon fetched successfully",
		"data":    paslonModels,
	})
	return nil
}

func (r *Repository) GetVoter(context *fiber.Ctx) error {
	voterModels := &[]models.Voter{}

	err := r.DB.Debug().Find(voterModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "voter fetched successfully",
		"data":    voterModels,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/articles", r.GetArticles)
	api.Get("/users", r.GetUsers)
	api.Get("/partai", r.GetPartai)
	api.Get("/paslon", r.GetPaslon)
	api.Get("/voter", r.GetVoter)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateArticles(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
