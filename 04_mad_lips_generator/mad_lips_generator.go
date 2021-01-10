package main

import (
	"fmt"
)

func main() {
	var color string
	var noun string
	var name string
	fmt.Println("Enter a color:")
	fmt.Scanln(&color)
	fmt.Println("Enter a plural noun:")
	fmt.Scanln(&noun)
	fmt.Println("Enter a name:")
	fmt.Scanln(&name)
	fmt.Printf("Roses are %s\n", color)
	fmt.Printf("%s are blue\n", noun)
	fmt.Printf("I love %s", name)
}
