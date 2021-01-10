package main

import (
	"fmt"
	"math/rand"
	"time"
	"bytes"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+-=,.?"

func random_character() string {
	rand.Seed(time.Now().UnixNano())
	return string(charset[rand.Intn(len(charset))])
}

func main() {
	fmt.Println("Enter the length of password:")
	var password_length int
	fmt.Scanln(&password_length)
	var password bytes.Buffer
	for i := 0; i < password_length; i++ {
		var c string = random_character()
		password.WriteString(c)
	}
	fmt.Println(password.String())
}
