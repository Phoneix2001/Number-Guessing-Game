# Number Guessing Game

**Number Guessing Game** is a fun and interactive command-line game written in Go. The game challenges you to guess a randomly generated number within a specified range and limited attempts based on the selected difficulty level.

---

## 🎮 Features

- 🧩 Three difficulty levels: Easy, Medium, and Hard  
- 🔢 Random number generation between 1 and 100  
- 🎯 Feedback on each guess (higher or lower)  
- 🔄 Option to replay the game after completion  

---

## 📦 Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)

---

## 🔧 Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/number-guessing-game.git
   cd number-guessing-game
   ```

2. **Build the binary:**

   ```bash
   go build -o number-guessing-game
   ```

3. *(Optional)* **Install globally:**

   ```bash
   sudo mv number-guessing-game /usr/local/bin/
   ```

   You can now use `number-guessing-game` from anywhere on your system.

---

## 🛠️ Usage

1. **Run the game:**

   ```bash
   ./number-guessing-game
   ```

2. **Follow the prompts:**

   - Select a difficulty level:
     - Easy: 10 chances
     - Medium: 5 chances
     - Hard: 3 chances
   - Enter your guesses to find the correct number.
   - Receive feedback on whether the number is higher or lower than your guess.

3. **Replay the game:**

   After completing a game, you can choose to play again by entering `Y` when prompted.

---

## 💡 Example Workflow

```bash
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2
Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 30
Congratulations! You guessed the correct number in 3 attempts.

Want to Play Again: Y/N
```

---

## 🌐 Project Source

This project is assigned from and submitted to [roadmap.sh](https://roadmap.sh/projects/number-guessing-game).
