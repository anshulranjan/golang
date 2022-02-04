package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func input(a int) int {
	var input string
	fmt.Printf("Let's guess the number between 1 and 10. Attempt %d", a)
	fmt.Println()
	fmt.Scanln(&input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return number
}
func main() {
	rand.Seed(time.Now().UnixNano())
	givennumber := input(1)
	guess := 10
	if givennumber < 0 {
		fmt.Println("Please choose a postive number")
		return
	}
	turn := 0
	for ; turn < 4; turn++ {
		n := rand.Intn(guess + 1)
		if n == givennumber {
			if turn == 0 {
				fmt.Println("Congrats!!! Since you guess in first attempt you won special prize")
			} else {
				fmt.Println("Congrats!!! You Won")
			}
			break
		} else {
			givennumber = input(turn + 2)
		}
	}
	if turn == 4 {
		fmt.Println("Sorry You lost. Try Again")
	}
}
