package main

import (
	"fmt"
	_ "goproject/custompkg"

	"github.com/guptarohit/asciigraph"
)

func main() {
	// custompkg.PrintCustom()
	data := []float64{3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
