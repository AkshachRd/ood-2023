package shapes

import (
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/shape"
)

type Shape struct {
	shapeBehavior shape.IShapeBehavior
	color         gfx.Color
	id            string
}

func NewShape(id string, color *gfx.Color, sb *shape.IShapeBehavior) *Shape {
	return &Shape{id: id, color: *color, shapeBehavior: *sb}
}

func (s *Shape) GetShapeName() string {
	return s.shapeBehavior.GetShapeName()
}

func (s *Shape) GetShapeInfo() string {
	return s.color.GetHex() + " " + s.shapeBehavior.GetShapeInfo()
}

func (s *Shape) Move(dx, dy float64) {
	s.shapeBehavior.Move(dx, dy)
}

func (s *Shape) Draw(canvas gfx.ICanvas) {
	canvas.SetColor(s.color)
	s.shapeBehavior.Draw(canvas)
}

func (s *Shape) SetColor(color *gfx.Color) {
	s.color = *color
}

func (s *Shape) SetShapeBehavior(sb shape.IShapeBehavior) {
	s.shapeBehavior = sb
}

func (s *Shape) GetColor() *gfx.Color {
	return &s.color
}

func (s *Shape) GetId() string {
	return s.id
}
