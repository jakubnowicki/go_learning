package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var dictionary [][]string

func getRandomElement(choices []string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(choices))
	return choices[randomIndex]
}

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text [][]string

	for scanner.Scan() {
		tmpText := strings.Split(scanner.Text(), ";")
		text = append(text, tmpText)
	}

	file.Close()

	var phrase string

	for _, line := range text {
		phrase = phrase + getRandomElement(line)
	}
	fmt.Println(phrase)
}
