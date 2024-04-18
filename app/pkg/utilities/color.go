package utilities

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateHexColor - функция для генерация color
func GenerateHexColor() string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate three random integers for RGB values
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	// Format the RGB values as a hexadecimal color string
	hexColor := fmt.Sprintf("#%02X%02X%02X", r, g, b)
	return hexColor
}
