package exercise4

import (
	"fmt"
	"math/rand"
)

func DisplayMenu() {
	option := 0

	for {
		fmt.Println("| Guess The Number |")
		fmt.Println("[1] Play")
		fmt.Println("[0] Exit")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Starting game...")
			play()
		case 0:
			fmt.Println("Exiting the game.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func play() {
	var randomNumber int = rand.Intn(101)
	guess := -1
	fmt.Println("I have selected a number between 0 and 100. Can you guess it?")

	for i := 0; i < 10; i++ {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess < randomNumber {
			fmt.Printf("The correct number is greater than %d\n", guess)
		} else if guess > randomNumber {
			fmt.Printf("The correct number is lesser than %d\n", guess)
		} else {
			fmt.Println("Congratulations! You guessed the number.")
			return
		}
	}
	fmt.Printf("Sorry, you didn't guess the number. The correct number was %d\n", randomNumber)
}
