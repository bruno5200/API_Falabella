package beer

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type beer struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}
type pay struct {
	Currency string `json:"currency"`
	Quantity int    `json:"quantity"`
}
type allBeers []beer

var beers = allBeers{
	{
		Id:       1,
		Name:     "Golden",
		Brewery:  "Kross",
		Country:  "Chile",
		Price:    10.5,
		Currency: "EUR",
	},
}

func BeerList(c *fiber.Ctx) error {
	if len(beers) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "No hay cervezas registradas",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success":     true,
		"description": "Operación exitosa",
		"beers":       beers,
	})
}
func NewBeer(c *fiber.Ctx) error {
	var newBeer beer

	b := new(beer)

	if err := c.BodyParser(b); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success":     false,
			"description": "Request Invalida",
		})
	}

	err := json.Unmarshal(c.Body(), &newBeer)

	newBeer.Id = len(beers) + 1

	beers = append(beers, newBeer)

	if err != nil {
		return c.Status(409).JSON(&fiber.Map{
			"success":     false,
			"description": "El ID de la cerveza ya existe.",
		})
	}

	c.JSON(newBeer)

	return c.Status(201).JSON(&fiber.Map{
		"success":     true,
		"description": "Cerveza creada.",
		"beer":        newBeer,
	})
}

func FindBeer(c *fiber.Ctx) error {
	Id := c.Params("beerID")

	beerId, err := StringToInt(Id)

	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success":     false,
			"description": "El Id de la cerveza no existe.",
		})
	}

	if len(beers) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "No hay cervezas registradas.",
		})
	}

	var newBeer beer
	for _, beer := range beers {
		if beer.Id == beerId {
			newBeer = beer
		}
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"beer":    newBeer,
	})
}
func BeerBox(c *fiber.Ctx) error {

	var newPay pay

	var beersQuantity = 6.0

	b := new(beer)

	err := json.Unmarshal(c.Body(), &newPay)

	if err != nil {
		return c.Status(409).JSON(&fiber.Map{
			"success":     false,
			"description": "Desconocido",
			"error":       err,
		})
	}

	Id := c.Params("beerID")

	beerId, err := StringToInt(Id)

	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success":     false,
			"description": "El Id de la cerveza no existe.",
		})
	}

	if beerId == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success":     false,
			"description": "El Id de la cerveza no existe.",
		})
	}

	if err := c.BodyParser(b); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success":     false,
			"description": "Request Invalida",
		})
	}

	if len(beers) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "No hay cervezas registradas.",
		})
	}

	var newBeerPack beer
	for _, beer := range beers {
		if beer.Id == beerId {
			newBeerPack = beer
		}
	}

	total := newBeerPack.Price * float32(newPay.Quantity) * float32(beersQuantity)

	if newBeerPack.Currency != newPay.Currency {
		var en = newBeerPack.Currency + newPay.Currency
		var de = newPay.Currency + newBeerPack.Currency
		switch en {
		case "USDEUR":
			total = total * 0.881505
		case "USDGBP":
			total = total * 0.734555
		case "USDCAD":
			total = total * 1.264435
		case "USDPLN":
			total = total * 4.00301
		case "USDCLP":
			total = total * 835.599323
		case "USDBOB":
			total = total * 6.888193
		}
		switch de {
		case "EURUSD":
			total = total / 0.881505
		case "GBPUSD":
			total = total / 0.734555
		case "CADUSD":
			total = total / 1.264435
		case "PLNUSD":
			total = total / 4.00301
		case "CLPUSD":
			total = total / 835.599323
		case "BOBUSD":
			total = total / 6.888193
		}

	}

	return c.Status(200).JSON(&fiber.Map{
		"success":     true,
		"description": "Operación Exitosa.",
		"priceTotal:": total,
		"currency":    newPay.Currency,
	})
}

func StringToInt(s string) (int, error) {
	var beerId int

	_, err := fmt.Sscan(s, &beerId)

	if err != nil {
		return 0, err
	}
	return beerId, err
}
