package shapes

import (
	"bufio"
	"fmt"
	canvas2 "github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"github.com/AkshachRd/ood-2023/labs/01/shapes/shape"
	"io"
	"strconv"
	"strings"
)

type PictureController struct {
	picture    *Picture
	canvas     canvas2.ICanvas
	input      io.Reader
	output     io.Writer
	commandMap map[string]func([]string)
}

func NewPictureController(picture *Picture, canvas canvas2.ICanvas, input io.Reader, output io.Writer) *PictureController {
	controller := &PictureController{
		picture: picture,
		canvas:  canvas,
		input:   input,
		output:  output,
	}

	controller.commandMap = map[string]func([]string){
		"AddShape":    controller.HandleAddShape,
		"MoveShape":   controller.HandleMoveShape,
		"MovePicture": controller.HandleMovePicture,
		"DeleteShape": controller.HandleDeleteShape,
		"ChangeColor": controller.HandleChangeColor,
		"ChangeShape": controller.HandleChangeShape,
		"DrawShape":   controller.HandleDrawShape,
		"DrawPicture": controller.HandleDrawPicture,
		"List":        controller.HandleList,
	}

	return controller
}

func (c *PictureController) HandleCommand() (bool, error) {
	commandLine, err := bufio.NewReader(c.input).ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return false, err
		}
		return false, nil
	}

	args := strings.Fields(commandLine)
	if len(args) == 0 {
		return false, nil
	}

	command := args[0]
	handler, found := c.commandMap[command]
	if found {
		handler(args[1:])
		return true, nil
	}

	return false, nil
}

func (c *PictureController) HandleAddShape(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id> <hex color>")
		return
	}

	id := args[0]
	color := args[1]

	if id == "" || !canvas2.IsValidHexColor(color) {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id> <hex color>")
		return
	}

	var sbf shape.ShapeBehaviorFactory
	behavior, err := sbf.CreateBehaviorFromArgs(args[2:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.picture.AddShape(id, *canvas2.NewColorFromHex(color), &behavior)
}

func (c *PictureController) HandleMovePicture(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <delta x> <delta y>")
		return
	}

	dxStr := args[0]
	dyStr := args[1]

	dx, err1 := strconv.ParseFloat(dxStr, 64)
	dy, err2 := strconv.ParseFloat(dyStr, 64)

	if err1 != nil || err2 != nil {
		fmt.Fprintln(c.output, "can not parse numbers from arguments")
		return
	}

	c.picture.MovePicture(dx, dy)
}

func (c *PictureController) HandleMoveShape(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id> <delta x> <delta y>")
		return
	}

	id := args[0]
	dxStr := args[1]
	dyStr := args[2]

	dx, err1 := strconv.ParseFloat(dxStr, 64)
	dy, err2 := strconv.ParseFloat(dyStr, 64)

	if id == "" || err1 != nil || err2 != nil {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id> <delta x> <delta y>")
		return
	}

	shape, err := c.picture.GetShapeById(id)
	if err == nil {
		shape.Move(dx, dy)
	}
}

func (c *PictureController) HandleDeleteShape(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id>")
		return
	}

	id := args[0]

	c.picture.DeleteShape(id)
}

func (c *PictureController) HandleChangeColor(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id> <hex color>")
		return
	}

	id := args[0]
	color := args[1]

	if id == "" || !canvas2.IsValidHexColor(color) {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id> <hex color>")
		return
	}

	s, _ := c.picture.GetShapeById(id)
	s.SetColor(canvas2.NewColorFromHex(color))
}

func (c *PictureController) HandleChangeShape(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id>")
		return
	}

	id := args[0]

	if id == "" {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id>")
		return
	}

	var sbf shape.ShapeBehaviorFactory
	behavior, err := sbf.CreateBehaviorFromArgs(args[1:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, err := c.picture.GetShapeById(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s.SetShapeBehavior(behavior)
}

func (c *PictureController) HandleList(args []string) {
	shapeInfo := c.picture.List()
	for _, info := range shapeInfo {
		fmt.Fprintln(c.output, info)
	}
}

func (c *PictureController) HandleDrawPicture(args []string) {
	c.picture.Draw(c.canvas)
}

func (c *PictureController) HandleDrawShape(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id>")
		return
	}

	id := args[0]

	if id == "" {
		fmt.Fprintln(c.output, "invalid command arguments, template: <id>")
		return
	}

	shape, _ := c.picture.GetShapeById(id)
	if shape != nil {
		shape.Draw(c.canvas)
	}
}
