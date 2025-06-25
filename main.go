package main

import (
	"github.com/Leeroyakbar/bowlnow-backend/configs"
	"github.com/Leeroyakbar/bowlnow-backend/handlers"
	"github.com/Leeroyakbar/bowlnow-backend/repositories"
	"github.com/Leeroyakbar/bowlnow-backend/routes"
	"github.com/Leeroyakbar/bowlnow-backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	configs.InitDB()

	db := configs.DB
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // izinkan React Vite
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	routes.RegisterUserRoutes(app, handler)
	app.Listen(":3000")
}
