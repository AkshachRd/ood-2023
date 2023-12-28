package shapes

import (
	"fmt"
	canvas "github.com/AkshachRd/ood-2023/labs/01/shapes/gfx"
	"os"
)

type Args struct {
	OutputFilename string
}

func ParseArgs(args []string) (Args, error) {
	if len(args) != 2 {
		return Args{}, fmt.Errorf("Invalid arguments format\nUsage: program.exe <output file name>")
	}
	return Args{OutputFilename: args[1]}, nil
}

func Main() {
	args, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	picture := NewPicture()
	canvas := canvas.NewCanvas()
	controller := NewPictureController(picture, canvas, os.Stdin, os.Stdout)

	for {
		fmt.Print("\n> ")

		if ok, err := controller.HandleCommand(); !ok {
			if err != nil {
				break
			}

			fmt.Println("Unknown command!")
		}
	}

	if err = canvas.Save(args.OutputFilename); err != nil {
		fmt.Println("Error saving to SVG:", err)
		os.Exit(1)
	}
}
