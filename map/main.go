package main

import "fmt"

var p = fmt.Println

func main() {
	colors := make(map[string]int)
	// colors := map[string]int{
	// 	"red":    1,
	// 	"yellow": 2,
	// }
	colors["red"] = 1
	colors["blue"] = 2
	colors["yellow"] = 5
	printMap(colors)
}

func printMap(c map[string]int) {
	for color, digit := range c {
		p("digit for", color, "is", digit)
	}
}
