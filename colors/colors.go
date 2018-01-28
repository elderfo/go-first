package colors

import (
	"errors"
)

// RetrievalColor is a type of color
type RetrievalColor int

const (
	// Purple related to the color purple
	Purple RetrievalColor = 0
	// Red related to the color red
	Red RetrievalColor = 1
	// Blue related to the color blue
	Blue RetrievalColor = 2
)

// GetColor returns a color
func GetColor(retrievalColor RetrievalColor) (color string, err error) {
	switch retrievalColor {
	case Purple:
		return "purple", nil
	case Red:
		return "red", nil
	case Blue:
		return "blue", nil
	default:
		return "", errors.New("Invalid argument")
	}
}
