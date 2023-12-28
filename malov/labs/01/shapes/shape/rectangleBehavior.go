package shape

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
)

type RectangleBehavior struct {
	topLeft point.Point
	width   float64
	height  float64
}

func NewRectangleBehavior(topLeft *point.Point, width, height float64) *RectangleBehavior {
	return &RectangleBehavior{*topLeft, width, height}
}

func (lb *RectangleBehavior) GetShapeInfo() string {
	return fmt.Sprintf("%f %f %f %f", lb.topLeft.X, lb.topLeft.Y, lb.width, lb.height)
}

func (lb *RectangleBehavior) Move(dx, dy float64) {
	lb.topLeft.X += dx
	lb.topLeft.Y += dy
}

func (lb *RectangleBehavior) Draw(canvas gfx.ICanvas) {
	canvas.MoveTo(lb.topLeft.X, lb.topLeft.Y)
	canvas.LineTo(lb.topLeft.X+lb.width, lb.topLeft.Y)
	canvas.LineTo(lb.topLeft.X, lb.topLeft.Y+lb.height)

	canvas.MoveTo(lb.topLeft.X+lb.width, lb.topLeft.Y+lb.height)
	canvas.LineTo(lb.topLeft.X+lb.width, lb.topLeft.Y)
	canvas.LineTo(lb.topLeft.X, lb.topLeft.Y+lb.height)
}

func (lb *RectangleBehavior) GetShapeName() string {
	return "rectangle"
}
