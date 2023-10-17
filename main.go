package main

//Libreria para imprimir por consola
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	//Armo el ruteo, le digo que cuando ingrese a la ruta home, me redirija a la carpeta public.
	app.Static("/", "./client/dist")
	// le mando un objeto del tipo  contexto que tiene el resp y el request
	// la funcion puerde retornar un error
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "usuarios desde el backend",
		})
	})
	app.Listen(":3000")
	fmt.Print("Server on PORT 3000")
}
