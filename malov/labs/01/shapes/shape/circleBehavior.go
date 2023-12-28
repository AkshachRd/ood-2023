package shape

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
)

type CircleBehavior struct {
	center point.Point
	radius float64
}

func NewCircleBehavior(center *point.Point, radius float64) *CircleBehavior {
	return &CircleBehavior{*center, radius}
}

func (cb *CircleBehavior) GetShapeInfo() string {
	return fmt.Sprintf("%f %f %f", cb.center.X, cb.center.Y, cb.radius)
}

func (cb *CircleBehavior) Move(dx, dy float64) {
	cb.center.X += dx
	cb.center.Y += dy
}

func (cb *CircleBehavior) Draw(canvas gfx.ICanvas) {
	canvas.DrawEllipse(cb.center.X, cb.center.Y, cb.radius, cb.radius)
}

func (cb *CircleBehavior) GetShapeName() string {
	return "circle"
}
