package shapes

import (
	"errors"
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/shape"
)

type Picture struct {
	shapesMap  map[string]*Shape
	shapesList []*Shape
}

func NewPicture() *Picture {
	return &Picture{shapesList: make([]*Shape, 0), shapesMap: make(map[string]*Shape)}
}

func (p *Picture) MovePicture(dx, dy float64) {
	for i := range p.shapesList {
		p.shapesList[i].Move(dx, dy)
	}
}

func (p *Picture) AddShape(id string, color gfx.Color, sb *shape.IShapeBehavior) error {
	_, exists := p.shapesMap[id]
	if exists {
		return errors.New("shape with this id already exist")
	}

	newShape := NewShape(id, &color, sb)
	p.shapesMap[id] = newShape
	p.shapesList = append(p.shapesList, newShape)

	return nil
}

func (p *Picture) DeleteShape(id string) {
	for i, listElement := range p.shapesList {
		if listElement == p.shapesMap[id] {
			p.shapesList = append(p.shapesList[:i], p.shapesList[i+1:]...)
		}
	}
	delete(p.shapesMap, id)
}

func (p *Picture) List() []string {
	var result []string

	for i, s := range p.shapesList {
		result = append(result, fmt.Sprintf("%d %s %s %s", i+1, s.GetShapeName(), s.GetId(), s.GetShapeInfo()))
	}

	return result
}

func (p *Picture) Draw(canvas gfx.ICanvas) {
	for _, s := range p.shapesList {
		s.Draw(canvas)
	}
}

func (p *Picture) GetShapeById(id string) (*Shape, error) {
	s, exists := p.shapesMap[id]
	if !exists {
		return nil, errors.New("shape with this id is not exist")
	}

	return s, nil
}
