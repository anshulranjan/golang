package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//this program allow the user to guess the number between 1 and 10. The user will get 5 attempt.
//the guess number will also change after each failed attempt.
//the user will get special prize if he guess the number in first attempt.
func input(a int) int {
	var input string
	fmt.Printf("Let's guess the number between 1 and 10. Attempt %d", a)
	fmt.Println()
	fmt.Scanln(&input)
	number, err := strconv.Atoi(input)
	if number < 0 && number > 10 {
		fmt.Println("The number should be between 0 and 10")
		return -1
	}
	if err != nil {
		return -1
	}
	return number
}
func main() {
	rand.Seed(time.Now().UnixNano())
	guess := rand.Intn(10)
	fmt.Println("Welcome to the game")
	fmt.Println("You have to guess the right number")
	fmt.Println("The number will be between 0 and 10")
	fmt.Println("You will have 3 attempt to guess the number")
	turn := 0
	for ; turn < 3; turn++ {
		outputnumber := input(turn + 1)
		if outputnumber == -1 {
			continue
		}
		if guess == outputnumber {
			if turn == 0 {
				fmt.Println("Congrats!!! Since you guess in first attempt you won special prize")
			} else {
				fmt.Println("Congrats!!! You Won")
			}
			break
		} else {
			guess = rand.Intn(10)
		}
	}
	if turn == 3 {
		fmt.Println("Sorry You lost. Try Again")
	}
}
