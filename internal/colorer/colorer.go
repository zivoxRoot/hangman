package colorer

import (
	"os"
)

// Colorer has a 'ColorMode' property, and has methods to render colors.
type Colorer struct {
	ColorMode string
}

// NewColorer creates a new instance of 'Colorer', it automatically determines the 'ColorMode' property of 'Colorer'
func NewColorer() *Colorer {
	colorer := Colorer{}

	colorSupport := getColorSupport()
	colorer.ColorMode = colorSupport

	return &colorer
}

// getColorSupport gets the user's terminal's supported colors.
func getColorSupport() string {
	term := os.Getenv("COLORTERM")
	return term
}
