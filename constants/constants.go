package constants

var (
	Levels = map[int]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}
	LevelsChances = map[int]int{
		1: 10,
		2: 5,
		3: 3,
	}
)

const (
	SelectedGame = `Great! You have selected the %s difficulty level.
Let's start the game!
`
	WelcomeMessage = `Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.
	
`
SelectDifficultyLevel = `Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
	
Enter your choice: `
	ErrParseDifficultyLevel = "Error parsing difficulty level"
	ErrParseGuess           = "Error parsing guess"
	DifficultyRangeError    = "Difficulty should be between 1-3"
	CongoText               = "Congratulations! You guessed the correct number in %d attempts."
)
