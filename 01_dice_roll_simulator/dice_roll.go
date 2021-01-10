package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll_dice() int {
	rand.Seed(time.Now().UnixNano())
	var dice_value = rand.Intn(6) + 1
	return dice_value
}

func main() {
	fmt.Println("Press 1 to roll the dice or 0 to exit:")
	var answer string
	fmt.Scanln(&answer)
	if answer == "1" {
		var value = roll_dice()
		fmt.Println(value)
	}
}
