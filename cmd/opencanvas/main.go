package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/irth/opencanvas/pkg/opencanvas"
	opencanvasRender "github.com/irth/opencanvas/pkg/renderer"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: opencanvas <input file> <output file>")
		os.Exit(1)
	}

	inFileName := os.Args[1]
	outFileName := os.Args[2]

	inFile, err := os.Open(inFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer inFile.Close()

	outFile, err := os.Create(outFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	var canvas opencanvas.Canvas
	err = json.NewDecoder(inFile).Decode(&canvas)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", canvas.BoundingBox())

	err = opencanvasRender.Render(outFile, &canvas)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
