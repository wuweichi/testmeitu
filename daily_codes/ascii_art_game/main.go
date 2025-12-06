package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ASCII art database
var asciiArts = map[string]string{
	"cat": `
 /\_/\  
( o.o ) 
 > ^ <  
`,
	"dog": `
 / \__
 (    @\___
 /         O
/   (_____/
/_____/   U
`,
	"dragon": `
                    __====-_  _-====___
            _--^^^#####//      \\\#####^^^--_
         _-^##########// (    ) \\\##########^-_
        -############//  |\^^/|  \\\############-
      _/############//   (@::@)   \\\############\_ 
     /#############((     \\//     ))#############\ 
    -###############\\    (oo)    //###############-
   -#################\\  / VV \  //#################-
  -###################\\/      \//###################-
 _#/|##########/\######(   /\   )######/\##########|\#_
 |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
 `  |/  V  V  `   V  \#\| |  | |/#/  V   '  V  V  \|  '
    `   `  `      `   / | |  | | \   '      '  '   '
                     (  | |  | |  )
                    __\ | |  | | /__
                   (vvv(VVV)(VVV)vvv)
`,
	"tree": `
      *
     ***
    *****
   *******
  *********
 ***********
     |||
     |||
`,
	"house": `
      /\
     /  \
    /    \
   /______\
   |      |
   |      |
   |______|
   |  |  |
   |__|__|
`,
}

// Game state structure
type GameState struct {
	Score      int
	Level      int
	Lives      int
	TimeLimit  int
	StartTime  time.Time
	CurrentArt string
}

// Initialize game
func NewGame() *GameState {
	return &GameState{
		Score:      0,
		Level:      1,
		Lives:      3,
		TimeLimit:  30,
		StartTime:  time.Now(),
		CurrentArt: "",
	}
}

// Display game header
func displayHeader(gs *GameState) {
	fmt.Println("========================================")
	fmt.Printf("ASCII Art Game | Score: %d | Level: %d | Lives: %d\n", gs.Score, gs.Level, gs.Lives)
	fmt.Println("========================================")
}

// Display ASCII art
func displayArt(name string) {
	if art, exists := asciiArts[name]; exists {
		fmt.Println(art)
	} else {
		fmt.Println("Art not found!")
	}
}

// Get random art name
func getRandomArt() string {
	keys := make([]string, 0, len(asciiArts))
	for k := range asciiArts {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

// Quiz game
func playQuiz(gs *GameState) bool {
	artName := getRandomArt()
	gs.CurrentArt = artName
	
	fmt.Println("\nGuess the ASCII art name:")
	displayArt(artName)
	fmt.Print("Your guess: ")
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	if strings.EqualFold(input, artName) {
		gs.Score += 10 * gs.Level
		fmt.Printf("Correct! +%d points\n", 10*gs.Level)
		return true
	} else {
		gs.Lives--
		fmt.Printf("Wrong! It was '%s'. Lives remaining: %d\n", artName, gs.Lives)
		return false
	}
}

// Drawing game
func playDrawing(gs *GameState) bool {
	targetArt := getRandomArt()
	fmt.Printf("\nDraw '%s' using ASCII characters. Type 'done' when finished.\n", targetArt)
	fmt.Println("Example: use characters like *, -, |, /, \\, etc.")
	
	reader := bufio.NewReader(os.Stdin)
	var drawing strings.Builder
	
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		if strings.EqualFold(line, "done") {
			break
		}
		
		drawing.WriteString(line + "\n")
	}
	
	fmt.Println("\nYour drawing:")
	fmt.Println(drawing.String())
	fmt.Println("\nOriginal art:")
	displayArt(targetArt)
	
	fmt.Print("Rate your drawing (1-10): ")
	ratingInput, _ := reader.ReadString('\n')
	ratingInput = strings.TrimSpace(ratingInput)
	rating, err := strconv.Atoi(ratingInput)
	
	if err != nil || rating < 1 || rating > 10 {
		rating = 5
	}
	
	points := rating * gs.Level
	gs.Score += points
	fmt.Printf("You earned %d points!\n", points)
	return true
}

// Memory game
func playMemory(gs *GameState) bool {
	artName := getRandomArt()
	fmt.Println("\nMemorize this ASCII art for 3 seconds:")
	displayArt(artName)
	
	time.Sleep(3 * time.Second)
	
	// Clear screen
	fmt.Print("\033[2J\033[H")
	
	fmt.Println("Now answer the questions:")
	
	questions := []struct {
		question string
		answer   string
	}{
		{"What was the art name?", artName},
		{"How many lines did it have?", strconv.Itoa(strings.Count(asciiArts[artName], "\n"))},
		{"Did it contain the character '/'? (yes/no)", func() string {
			if strings.Contains(asciiArts[artName], "/") {
				return "yes"
			}
			return "no"
		}()},
	}
	
	correct := 0
	reader := bufio.NewReader(os.Stdin)
	
	for _, q := range questions {
		fmt.Printf("%s: ", q.question)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if strings.EqualFold(input, q.answer) {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong! Correct answer: %s\n", q.answer)
		}
	}
	
	points := correct * 5 * gs.Level
	gs.Score += points
	fmt.Printf("You got %d/%d correct! +%d points\n", correct, len(questions), points)
	return correct > 0
}

// Art gallery
func viewGallery() {
	fmt.Println("\n=== ASCII Art Gallery ===")
	for name, art := range asciiArts {
		fmt.Printf("\n%s:\n%s", name, art)
	}
}

// Add custom art
func addCustomArt() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter art name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	if name == "" {
		fmt.Println("Invalid name!")
		return
	}
	
	fmt.Println("Enter your ASCII art (type 'END' on a new line to finish):")
	var art strings.Builder
	
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		
		if line == "END" {
			break
		}
		
		art.WriteString(line + "\n")
	}
	
	asciiArts[name] = art.String()
	fmt.Printf("Art '%s' added successfully!\n", name)
}

// Save game
func saveGame(gs *GameState) {
	file, err := os.Create("savegame.txt")
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	defer file.Close()
	
	fmt.Fprintf(file, "%d\n%d\n%d\n", gs.Score, gs.Level, gs.Lives)
	fmt.Println("Game saved successfully!")
}

// Load game
func loadGame(gs *GameState) {
	file, err := os.Open("savegame.txt")
	if err != nil {
		fmt.Println("No saved game found.")
		return
	}
	defer file.Close()
	
	var score, level, lives int
	_, err = fmt.Fscanf(file, "%d\n%d\n%d\n", &score, &level, &lives)
	if err != nil {
		fmt.Println("Error loading game:", err)
		return
	}
	
	gs.Score = score
	gs.Level = level
	gs.Lives = lives
	fmt.Println("Game loaded successfully!")
}

// Main menu
func showMenu() {
	fmt.Println("\n=== Main Menu ===")
	fmt.Println("1. Play Quiz Game")
	fmt.Println("2. Play Drawing Game")
	fmt.Println("3. Play Memory Game")
	fmt.Println("4. View Art Gallery")
	fmt.Println("5. Add Custom Art")
	fmt.Println("6. Save Game")
	fmt.Println("7. Load Game")
	fmt.Println("8. Exit")
	fmt.Print("Choose an option: ")
}

// Main function
func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	// Create game state
	game := NewGame()
	
	fmt.Println("Welcome to ASCII Art Game!")
	fmt.Println("Test your knowledge and creativity with ASCII art.")
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		displayHeader(game)
		showMenu()
		
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch input {
		case "1":
			if !playQuiz(game) && game.Lives <= 0 {
				fmt.Println("\nGame Over! You ran out of lives.")
				return
			}
			if game.Score >= game.Level*100 {
				game.Level++
				fmt.Printf("\nLevel up! Now at level %d\n", game.Level)
			}
		case "2":
			playDrawing(game)
		case "3":
			playMemory(game)
		case "4":
			viewGallery()
		case "5":
			addCustomArt()
		case "6":
			saveGame(game)
		case "7":
			loadGame(game)
		case "8":
			fmt.Printf("\nThanks for playing! Final score: %d\n", game.Score)
			return
		default:
			fmt.Println("Invalid option!")
		}
		
		// Check time limit
		if time.Since(game.StartTime).Seconds() > float64(game.TimeLimit*60) {
			fmt.Println("\nTime's up! Game over.")
			fmt.Printf("Final score: %d\n", game.Score)
			return
		}
	}
}
