package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

var choices = []string{"rock", "paper", "scissors"}

func pick_value() string {
	rand.Seed(time.Now().UnixNano())
	random_index := rand.Intn(len(choices))
	return choices[random_index]
}

func check_result(player, computer string) (string, int, int) {
	if player == computer {
		return "Tie!", 0, 0
	}

	player_value, _ := search_array(choices, player)
	computer_value, _ := search_array(choices, computer)

	result := player_value - computer_value

	if result == -1 || result == 2 {
		return fmt.Sprintf("Computer scores!\n%s beats %s", computer, player), 0, 1
	}

	if result == -2 || result == 1 {
		return fmt.Sprintf("Player scores!\n%s beats %s", player, computer), 1, 0
	}

	return "Error!", 0, 0
}

func search_array(array []string, value string) (int, bool) {
	for i, _ := range array {
		if array[i] == value {
			return i, true
		}
	}

	return 0, false
}

func main() {
	fmt.Println("Rock, Paper or Scissors?")
	computer_answer := pick_value()
	var player_answer string
	fmt.Scanln(&player_answer)
	player_answer = strings.ToLower(player_answer)
	result, _, _ := check_result(player_answer, computer_answer)
	fmt.Println(result)
}
