package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var colores = []string{
	"2f4f4f", // DarkSlateGray
	"556b2f", // DarkOliveGreen
	"8b0000", // DarkRed
	"800080", // Purple
	"b22222", // Firebrick
	"a52a2a", // Brown
	"5f9ea0", // CadetBlue
	"4b0082", // Indigo
	"483d8b", // DarkSlateBlue
	"6a5acd", // SlateBlue
	"708090", // SlateGray
	"2e8b57", // SeaGreen
	"3cb371", // MediumSeaGreen
	"228b22", // ForestGreen
	"cd5c5c", // IndianRed
	"c71585", // MediumVioletRed
	"9932cc", // DarkOrchid
	"8a2be2", // BlueViolet
	"ba55d3", // MediumOrchid
	"d2691e", // Chocolate
	"ff6347", // Tomato
	"ff4500", // OrangeRed
	"800000", // Maroon
	"4169e1", // RoyalBlue
	"00008b", // DarkBlue
	"191970", // MidnightBlue
	"000080", // Navy
	"8b4513", // SaddleBrown
	"7f7f7f", // Gray
}

func main() {

	app := fiber.New()

	app.Get("/avatar", func(c *fiber.Ctx) error {

		nombre := c.Query("nombre")

		if nombre == "" {
			return c.Status(fiber.StatusBadRequest).SendString("El parámetro 'nombre' es obligatorio")
		}

		avatarBytes := CrearAvatar(nombre)

		c.Set("Content-Type", "image/svg+xml")

		return c.Send(avatarBytes)

	})

	port := os.Getenv("PortServer")

	if port == "" {
		port = "2000"
	}

	log.Fatal(app.Listen(":" + port))

}

func CrearAvatar(nombre string) []byte {

	rand.Seed(time.Now().UnixNano())
	bg := colores[rand.Intn(len(colores))]

	words := strings.Fields(nombre)
	initials := ""
	for _, w := range words {
		if len(w) > 0 {
			initials += strings.ToUpper(string(w[0]))
		}
	}

	if len(initials) > 2 {
		initials = initials[:2]
	}

	svg := fmt.Sprintf(
		`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 100 100">
			<rect width="100" height="100" rx="12" fill="#%s"/>
			<text x="50" y="50"
				text-anchor="middle"
				dominant-baseline="central"
				font-family="Arial, Helvetica, sans-serif"
				font-size="40"
				font-weight="600"
				fill="#ffffff">%s</text>
		</svg>`,
		128, 128, bg, initials,
	)

	return []byte(svg)
}
