package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	colors[10] = "#ffffff"

	delete(colors, 10)

	fmt.Println(colors)
}
