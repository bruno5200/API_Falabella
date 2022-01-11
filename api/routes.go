package api

import (
	"github.com/bruno5200/API-Falabella/beer"
	"github.com/gofiber/fiber/v2"
)

func RoutesUp(app *fiber.App) {
	app.Get("/beers", beer.BeerList)
	app.Post("/beers", beer.NewBeer)
	app.Get("/beers/:beerID", beer.FindBeer)
	app.Get("/beers/:beerID/boxprice", beer.BeerBox)
}
