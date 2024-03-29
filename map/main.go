package main

import "fmt"

func main() {
	// Ways to declare map
	// #1
	colors := map[string]string{
		"red":  "#FF0000",
		"blue": "#0000FF",
	}

	// #2
	var colors2 map[string]string

	// #3
	colors3 := make(map[string]string)

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	// Assign value
	colors3["white"] = "#FFFFFF"
	fmt.Println(colors3)

	// Delete value
	delete(colors3, "white")

	fmt.Println(colors3)

	colors["green"] = "#00FF00"

	// Iterate map
	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("Hex code for color %v is %v\n", key, value)
	}
}
