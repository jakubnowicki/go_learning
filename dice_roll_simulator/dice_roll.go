package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll_dice() int {
	rand.Seed(time.Now().UnixNano())
	var dice_value = rand.Intn(5) + 1
	return dice_value
}

func main() {
	var value = roll_dice()
	fmt.Println(value)
}
