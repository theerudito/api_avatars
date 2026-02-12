package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func router(app *fiber.App) {

	allowedOrigins := map[string]bool{
		"http://localhost:5173": true,
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowOriginsFunc: func(origin string) bool {
			fmt.Println(origin)
			return allowedOrigins[origin]
		},
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/avatars/:option", ImageController)

}
