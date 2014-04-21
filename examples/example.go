package main

import (
	"fmt"
	"os"

	"github.com/hailiang/go-shapefile"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: go run examples/example.go [map/bou2_4p.shp]")
	} else {
		file := os.Args[1]
		shapeFile := shapefile.Open(file)
		fmt.Printf("%#v\n", shapeFile)

		for k, v := range shapeFile.Shape(0).Attrs {
			fmt.Printf("%s = %#v\n", k, v)
		}
	}
}
