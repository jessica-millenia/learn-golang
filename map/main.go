package main

import "fmt"

func main() {
	colors := map[string]string{
		"black":  "#000000",
		"yellow": "#ffff00",
		"white":  "#ffffff",
	}
	colors["blue"] = "#0000ff"

	fmt.Println("yellow hex color is ", colors["yellow"])
	delete(colors, "yellow")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "hex color is", hex)
	}
}
