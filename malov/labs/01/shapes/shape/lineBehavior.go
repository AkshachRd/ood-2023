package shape

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
)

type LineBehavior struct {
	startPoint point.Point
	endPoint   point.Point
}

func NewLineBehavior(startPoint, endPoint *point.Point) *LineBehavior {
	return &LineBehavior{*startPoint, *endPoint}
}

func (lb *LineBehavior) GetShapeInfo() string {
	return fmt.Sprintf("%f %f %f %f", lb.startPoint.X, lb.startPoint.Y, lb.endPoint.X, lb.endPoint.Y)
}

func (lb *LineBehavior) Move(dx, dy float64) {
	lb.startPoint.X += dx
	lb.startPoint.Y += dy
	lb.endPoint.X += dx
	lb.endPoint.Y += dy
}

func (lb *LineBehavior) Draw(canvas gfx.ICanvas) {
	canvas.MoveTo(lb.startPoint.X, lb.startPoint.Y)
	canvas.LineTo(lb.endPoint.X, lb.endPoint.Y)
}

func (lb *LineBehavior) GetShapeName() string {
	return "line"
}
