package main

import (
	"fmt"
	"math/rand"
)

func welcome() int {
	var s int
	fmt.Printf("Please select the difficulty level:\n1. Easy|(10 live)\n2. Medium(|5 live)\n3. Hard|(3 live)")
	fmt.Println("Enter number level:")
	fmt.Scan(&s)
	return s
}

func choice_dif(s int) int {
	var Easy int = 1
	var Medium int = 2
	var Hard int = 3

	switch s {
	case 1:
		return Easy
	case 2:
		return Medium
	case 3:
		return Hard

	default:
		return Medium
	}
}

func randomBetween(min, max int) int {
	random_number := rand.Intn(max-min+1) + min
	return random_number
}

func main() {
	level := welcome()
	lives := choice_dif(level)
	target_number := randomBetween(1, 100)

	fmt.Println("Guess a number between 1 and 100")

	for lives > 0 {
		var guess int
		fmt.Println("Enter u guess:")
		fmt.Scan(&guess)
		if guess == target_number {
			fmt.Println("Nice! u win")
			return
		} else if guess < target_number {
			fmt.Println("Too Low!")
		} else {
			fmt.Println("Too High!")
		}

		lives--
		fmt.Printf("You have %d lives left.\n", lives)
	}

	fmt.Println("Sorry u lives out. The number was:", target_number)
}
