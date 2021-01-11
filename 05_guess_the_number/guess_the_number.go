package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check_value(guess, value int) (string, bool) {
	if guess == value {
		return fmt.Sprintf("You guessed right, the number is %d!", value), true
	}

	if guess > value {
		return "Your guess is too high!", false
	} else {
		return "Your guess is too low!", false
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var value = rand.Intn(10) + 1
	var result bool
	var message string
	for i := 0; i < 3; i++ {
		fmt.Println("Make a guess:")
		var guess int
		fmt.Scanln(&guess)
		message, result = check_value(guess, value)
		if i < 2 || result {
			fmt.Println(message)
		}
		if result { break }
	}
	if !result {
		fmt.Printf("You loose! Correct answer was %d", value)
	}
}
