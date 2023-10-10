package main

import "fmt"

func main() {

	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	printMap(colours)
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
