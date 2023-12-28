package shape

import (
	"errors"
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
	"strconv"
)

type ShapeBehaviorFactory struct{}

func (ssf *ShapeBehaviorFactory) CreateBehaviorFromArgs(args []string) (IShapeBehavior, error) {
	shapeTypeStr := args[0]
	shapeType := ssf.MapStringToShapeType(shapeTypeStr)
	behaviorArgs := args[1:]

	switch shapeType {
	case Circle:
		return ssf.CreateCircleBehavior(behaviorArgs)
	case Line:
		return ssf.CreateLineBehavior(behaviorArgs)
	case Rectangle:
		return ssf.CreateRectangleBehavior(behaviorArgs)
	case Text:
		return ssf.CreateTextBehavior(behaviorArgs)
	case Triangle:
		return ssf.CreateTriangleBehavior(behaviorArgs)
	default:
		return nil, errors.New("unknown shape")
	}
}

func (ssf *ShapeBehaviorFactory) MapStringToShapeType(str string) ShapeType {
	switch str {
	case "circle":
		return Circle
	case "line":
		return Line
	case "rectangle":
		return Rectangle
	case "text":
		return Text
	case "triangle":
		return Triangle
	default:
		msg := fmt.Sprintf("unknown shape type: %s", str)
		panic(msg)
	}
}

func (ssf *ShapeBehaviorFactory) CreateCircleBehavior(args []string) (*CircleBehavior, error) {
	cx, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return &CircleBehavior{}, errors.New("can not parse center x")
	}

	cy, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return &CircleBehavior{}, errors.New("can not parse center y")
	}

	radius, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return &CircleBehavior{}, errors.New("can not parse radius")
	}

	return NewCircleBehavior(&point.Point{X: cx, Y: cy}, radius), nil
}

func (ssf *ShapeBehaviorFactory) CreateLineBehavior(args []string) (*LineBehavior, error) {
	x1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return &LineBehavior{}, errors.New("can not parse x1")
	}

	y1, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return &LineBehavior{}, errors.New("can not parse y1")
	}

	x2, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return &LineBehavior{}, errors.New("can not parse x2")
	}

	y2, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		return &LineBehavior{}, errors.New("can not parse y2")
	}

	return NewLineBehavior(&point.Point{X: x1, Y: y1}, &point.Point{X: x2, Y: y2}), nil
}

func (ssf *ShapeBehaviorFactory) CreateTextBehavior(args []string) (*TextBehavior, error) {
	x, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return &TextBehavior{}, errors.New("can not parse x")
	}

	y, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return &TextBehavior{}, errors.New("can not parse y")
	}

	fontSize, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return &TextBehavior{}, errors.New("can not parse font size")
	}

	text := args[3]

	if text == "" {
		return &TextBehavior{}, errors.New("empty text")
	}

	return NewTextBehavior(&point.Point{X: x, Y: y}, fontSize, text), nil
}

func (ssf *ShapeBehaviorFactory) CreateTriangleBehavior(argsStr []string) (*TriangleBehavior, error) {
	var args [6]float64

	for i, argStr := range argsStr {
		arg, err := strconv.ParseFloat(argStr, 64)
		if err != nil {
			return &TriangleBehavior{}, errors.New("can not parse triangle argument")
		}
		args[i] = arg
	}

	return NewTriangleBehavior(&[3]point.Point{
		{args[0], args[1]},
		{args[2], args[3]},
		{args[4], args[5]},
	}), nil
}

func (ssf *ShapeBehaviorFactory) CreateRectangleBehavior(argsStr []string) (*RectangleBehavior, error) {
	var args [4]float64

	for i, argStr := range argsStr {
		arg, err := strconv.ParseFloat(argStr, 64)
		if err != nil {
			return &RectangleBehavior{}, errors.New("can not parse rectangle argument")
		}
		args[i] = arg
	}

	return NewRectangleBehavior(&point.Point{X: args[0], Y: args[1]}, args[2], args[3]), nil
}

type ShapeType int

const (
	Circle ShapeType = iota
	Line
	Rectangle
	Text
	Triangle
)
