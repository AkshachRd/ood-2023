package gfx

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/point"
	"os"
)

const StrokeWidth = 4

type ICanvas interface {
	MoveTo(x, y float64)
	SetColor(color Color)
	LineTo(x, y float64)
	DrawEllipse(cx, cy, rx, ry float64)
	DrawText(x, y, fontSize float64, text string)
}

type Canvas struct {
	CurrentColor Color
	DrawPoint    point.Point
	OutputSvg    string
}

func NewCanvas() *Canvas {
	return &Canvas{CurrentColor: *NewColorFromHex("#000000"), DrawPoint: point.Point{}}
}

func (c *Canvas) Save(outputFileName string) error {
	file, err := os.Create(outputFileName)
	defer file.Close()

	if err != nil {
		fmt.Println("Invalid output file", err)
		return err
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\">\n")
	if err != nil {
		fmt.Println("Error writing opening SVG tag:", err)
		return err
	}
	_, err = w.WriteString(c.OutputSvg + "\n")
	if err != nil {
		fmt.Println("Error writing SVG content:", err)
		return err
	}
	_, err = w.WriteString("</svg>")
	if err != nil {
		fmt.Println("Error writing closing SVG tag:", err)
		return err
	}

	if err != nil {
		return errors.New("can't write to output file")
	}

	w.Flush()

	return nil
}

func (c *Canvas) SetColor(color Color) {
	c.CurrentColor = color
}

func (c *Canvas) MoveTo(x, y float64) {
	c.DrawPoint = point.Point{x, y}
}

func (c *Canvas) LineTo(x, y float64) {
	c.OutputSvg += fmt.Sprintf(
		"<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" style=\"stroke:%s;stroke-width:%d\" />",
		c.DrawPoint.X,
		c.DrawPoint.Y,
		x,
		y,
		c.CurrentColor.GetHex(),
		StrokeWidth,
	)
}

func (c *Canvas) DrawEllipse(cx, cy, rx, ry float64) {
	c.OutputSvg += fmt.Sprintf(
		"<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" ry=\"%f\" style=\"stroke:%s;stroke-width:%d;fill:none\" />",
		cx,
		cy,
		rx,
		ry,
		c.CurrentColor.GetHex(),
		StrokeWidth,
	)
}

func (c *Canvas) DrawText(x, y, fontSize float64, text string) {
	c.OutputSvg += fmt.Sprintf(
		"<text x=\"%f\" y=\"%f\" font-size=\"%f\" fill=\"%s\">%s</text>",
		x,
		y,
		fontSize,
		c.CurrentColor.GetHex(),
		text,
	)
}
