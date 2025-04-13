package api

import (
	"fmt"
	"math/rand/v2"
	"number_guessing_game/constants"
	"strings"
)

type Server struct {
	Cammand       string
	SuffixCammand string
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {

	fmt.Print(constants.WelcomeMessage)
	s.StartGame()
	return nil
}

func (s *Server) StartGame() {
	fmt.Print(constants.SelectDifficultyLevel)
	var input int
	if _, err := fmt.Scanln(&input); err != nil {

		fmt.Println(constants.ErrParseDifficultyLevel)
	}
	if input < 1 || input > 3 {
		fmt.Println(constants.DifficultyRangeError)
		return
	}

	fmt.Printf(constants.SelectedGame, constants.Levels[input])
	sysGuess := rand.IntN(100)
	if sysGuess == 0 {
		sysGuess = 1
	}
	for i := range constants.LevelsChances[input] {
		guessedNumber := 0
		fmt.Print("\nEnter your guess: ")
		if _, err := fmt.Scanln(&guessedNumber); err != nil {
			fmt.Println(constants.ErrParseGuess)
		}
		if guessedNumber < 1 || guessedNumber > 100 {
			fmt.Println("Enter a number between 1-100")
			break
		}
		if sysGuess < guessedNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n", guessedNumber)
		} else if sysGuess > guessedNumber {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guessedNumber)
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i+1)
			break
		}

	}

	fmt.Println("\nWant to Play Again: Y/N")
	var wantToPlayAgain string
	if _, err := fmt.Scanln(&wantToPlayAgain); err != nil {
		fmt.Println("Error parsinf response")
	}
	fmt.Println()
	if strings.ToLower(wantToPlayAgain) == "y" {
		s.StartGame()
	}
}
