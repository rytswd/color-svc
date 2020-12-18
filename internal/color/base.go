// Package color is a utility to provide colors.
//
// Ref: https://developer.mozilla.org/en-US/docs/Web/CSS/color_value
package color

// Color holds color information.
type Color struct {
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}

// AllColors returns all available colors.
var AllColors []Color

func init() {
	for _, red := range RedishColors {
		AllColors = append(AllColors, Color(red))
	}
	for _, green := range GreenishColors {
		AllColors = append(AllColors, Color(green))
	}
	for _, blue := range BluishColors {
		AllColors = append(AllColors, Color(blue))
	}
	for _, yellow := range YellowishColors {
		AllColors = append(AllColors, Color(yellow))
	}
}

// EnabledColors creates a slice containing only selection of colors.
func EnabledColors(red, green, blue, yellow bool) []Color {
	r := []Color{}
	if red {
		for _, red := range RedishColors {
			r = append(r, Color(red))
		}
	}
	if green {
		for _, green := range GreenishColors {
			r = append(r, Color(green))
		}
	}
	if blue {
		for _, blue := range BluishColors {
			r = append(r, Color(blue))
		}
	}
	if yellow {
		for _, yellow := range YellowishColors {
			r = append(r, Color(yellow))
		}
	}
	return r
}
