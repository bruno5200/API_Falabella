package main

import (
	"errors"

	"github.com/bruno5200/API-Falabella/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, ok := err.(*fiber.Error); ok {
				return errors.New("this is managed error")
			}
			return errors.New("this is managed error")
		},
	})
	//middlewares
	app.Use(recover.New())
	app.Use(logger.New())

	api.RoutesUp(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bender te da la Bienvenida a su api de cervezas.")
	})
	app.Listen(":3000")
}

/*      Beer
		type object
		properties:
        - Id: integer `1`
        - Name: string `Golden`
        - Brewery: string `Kross`
        - Country: string `Chile`
        - Price: number	`10.5`
        - Currency: string `EUR`
*/

/*
GET		/beers: devuelve una lista de las cervezas
POST	/beers: crea nueva cerveza
GET		/beers/{beerID}: devuelve una cerveza específica
GET		/beers/{beerID}/boxprice: precio de el sixpack en la moneda solicitada
*/

/*		Beerbox
		type: object
		properties:
		priceTotal: number `63.0`
*/
