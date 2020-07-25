package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func main() {

	fmt.Println("                     GUESS THE FRUIT                 ")
	fmt.Println()
	fmt.Println()

	someWords := `apple banana mango strawberry
	orange grape pineapple apricot lemon coconut watermelon
	cherry papaya berry peach lychee muskmelon`

	words := strings.Split(someWords, " ")

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	word := words[r.Intn(len(words))]

	for i := 0; i < len(word); i++ {
		fmt.Print("_")
		fmt.Print(" ")
	}
	fmt.Println()

	wordLen := len(word)

	guessedWord := ""
	allowedAttempts := len(word) + 2
	correctGuess := false

	for allowedAttempts != 0 && correctGuess != true {
		fmt.Println()
		fmt.Println()

		fmt.Print("Enter a letter: ")
		allowedAttempts--
		var input string

		fmt.Scanf("%s", &input)
		fmt.Println()

		if !unicode.IsLetter([]rune(input)[0]) {
			fmt.Println("Enter only a alphabet")
			continue
		} else if len(string(input)) > 1 {
			fmt.Println("Enter only single alphabet")
			continue
		} else if strings.Contains(guessedWord, input) {
			fmt.Println("Alphabet already guessed")
			continue
		} else {

			guessedWord += string(input)
			count := 0
			for _, val := range word {
				if strings.Contains(guessedWord, string(val)) {
					count++
					fmt.Print(string(val))
					fmt.Print(" ")
				} else {
					fmt.Print("_")
					fmt.Print(" ")
				}
			}

			if wordLen == count {
				correctGuess = true
			}

		}
	}

	if correctGuess {
		fmt.Println()
		fmt.Println("You guessed the word correctly 😃")
	} else {
		fmt.Println()
		fmt.Println("You ran out of attempts ☹")
	}
}
