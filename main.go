package main

//Libreria para imprimir por consola
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	//Armo el ruteo, le digo que cuando ingrese a la ruta home, me redirija a la carpeta public.
	app.Static("/", "./public")
	app.Listen(":3000")
	fmt.Print("Server on PORT 3000")
}
