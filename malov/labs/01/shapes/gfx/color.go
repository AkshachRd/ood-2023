package gfx

import (
	"fmt"
	"regexp"
)

const HexColorRegex = "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"

type Color struct {
	hex string
}

func NewColor(r, g, b int) *Color {
	return &Color{rgbToHex(r, g, b)}
}

func NewColorFromHex(hex string) *Color {
	return &Color{hex}
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%x%x%x", r, g, b)
}

func IsValidHexColor(str string) bool {
	match, _ := regexp.Match(HexColorRegex, []byte(str))
	return match
}

func (c *Color) GetHex() string {
	return c.hex
}
