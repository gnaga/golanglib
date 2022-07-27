package main

import "fmt"

func main() {
	colors := map [string]string{
		"red":"#ff000",
		"green":"#asdff",
	}
	printMap(colors)
}

func printMap (c map[string]string) {
	for col,hex := range c {
		fmt.Println("color:",col, "hex:",hex)
	}
}