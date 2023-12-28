package shape

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
)

type TextBehavior struct {
	topLeft  point.Point
	fontSize float64
	text     string
}

func NewTextBehavior(topLeft *point.Point, fontSize float64, text string) *TextBehavior {
	return &TextBehavior{*topLeft, fontSize, text}
}

func (lb *TextBehavior) GetShapeInfo() string {
	return fmt.Sprintf("%f %f %f %s", lb.topLeft.X, lb.topLeft.Y, lb.fontSize, lb.text)
}

func (lb *TextBehavior) Move(dx, dy float64) {
	lb.topLeft.X += dx
	lb.topLeft.Y += dy
}

func (lb *TextBehavior) Draw(canvas gfx.ICanvas) {
	canvas.DrawText(lb.topLeft.X, lb.topLeft.Y, lb.fontSize, lb.text)
}

func (lb *TextBehavior) GetShapeName() string {
	return "text"
}
