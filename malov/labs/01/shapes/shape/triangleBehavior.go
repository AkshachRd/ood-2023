package shape

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
)

type TriangleBehavior struct {
	vertexes [3]point.Point
}

func NewTriangleBehavior(vertexes *[3]point.Point) *TriangleBehavior {
	return &TriangleBehavior{*vertexes}
}

func (lb *TriangleBehavior) GetShapeInfo() string {
	var result string
	for _, v := range lb.vertexes {
		result += fmt.Sprintf(" %f %f", v.X, v.Y)
	}
	return result
}

func (lb *TriangleBehavior) Move(dx, dy float64) {
	for i := range lb.vertexes {
		lb.vertexes[i].X += dx
		lb.vertexes[i].Y += dy
	}
}

func (lb *TriangleBehavior) Draw(canvas gfx.ICanvas) {
	canvas.MoveTo(lb.vertexes[0].X, lb.vertexes[0].Y)

	for _, v := range lb.vertexes[1:] {
		canvas.LineTo(v.X, v.Y)
		canvas.MoveTo(v.X, v.Y)
	}

	canvas.LineTo(lb.vertexes[0].X, lb.vertexes[0].Y)
}

func (lb *TriangleBehavior) GetShapeName() string {
	return "triangle"
}
