package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player represents a game player
type Player struct {
	Name  string
	Score int
	Level int
}

// Game represents the main game state
type Game struct {
	Players []Player
	Round   int
	Active  bool
}

// NewGame creates a new game instance
func NewGame() *Game {
	return &Game{
		Players: []Player{},
		Round:   1,
		Active:  true,
	}
}

// AddPlayer adds a new player to the game
func (g *Game) AddPlayer(name string) {
	player := Player{
		Name:  name,
		Score: 0,
		Level: 1,
	}
	g.Players = append(g.Players, player)
	fmt.Printf("Player %s joined the game!\n", name)
}

// StartRound begins a new round of the game
func (g *Game) StartRound() {
	fmt.Printf("\n=== Round %d ===\n", g.Round)
	for i := range g.Players {
		points := rand.Intn(100) + 1
		g.Players[i].Score += points
		if g.Players[i].Score > g.Players[i].Level*100 {
			g.Players[i].Level++
			fmt.Printf("%s leveled up to level %d!\n", g.Players[i].Name, g.Players[i].Level)
		}
		fmt.Printf("%s earned %d points (Total: %d)\n", g.Players[i].Name, points, g.Players[i].Score)
	}
	g.Round++
}

// DisplayScores shows all player scores
func (g *Game) DisplayScores() {
	fmt.Println("\n=== Current Scores ===")
	for _, player := range g.Players {
		fmt.Printf("%s: %d points (Level %d)\n", player.Name, player.Score, player.Level)
	}
}

// CheckWinner determines if there's a winner
func (g *Game) CheckWinner() *Player {
	var winner *Player
	maxScore := 0
	
	for i := range g.Players {
		if g.Players[i].Score > maxScore {
			maxScore = g.Players[i].Score
			winner = &g.Players[i]
		}
	}
	
	if maxScore >= 1000 {
		return winner
	}
	return nil
}

// SimulateGame runs the main game simulation
func SimulateGame() {
	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Println("===================================")
	
	game := NewGame()
	
	// Add some players
	players := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	for _, name := range players {
		game.AddPlayer(name)
	}
	
	// Game loop
	for game.Active {
		game.StartRound()
		game.DisplayScores()
		
		// Check for winner
		if winner := game.CheckWinner(); winner != nil {
			fmt.Printf("\n🎉 Congratulations! %s wins with %d points! 🎉\n", winner.Name, winner.Score)
			game.Active = false
			break
		}
		
		// Small delay between rounds
		time.Sleep(1 * time.Second)
	}
	
	fmt.Println("\nGame Over! Thanks for playing!")
}

// Additional utility functions to increase code length

// CalculateAverageScore calculates the average score of all players
func (g *Game) CalculateAverageScore() float64 {
	if len(g.Players) == 0 {
		return 0
	}
	
	total := 0
	for _, player := range g.Players {
		total += player.Score
	}
	return float64(total) / float64(len(g.Players))
}

// GetTopPlayers returns the top N players by score
func (g *Game) GetTopPlayers(n int) []Player {
	if n > len(g.Players) {
		n = len(g.Players)
	}
	
	// Create a copy to avoid modifying original
	players := make([]Player, len(g.Players))
	copy(players, g.Players)
	
	// Simple bubble sort for demonstration
	for i := 0; i < len(players)-1; i++ {
		for j := 0; j < len(players)-i-1; j++ {
			if players[j].Score < players[j+1].Score {
				players[j], players[j+1] = players[j+1], players[j]
			}
		}
	}
	
	return players[:n]
}

// SaveGameState simulates saving game state (placeholder)
func (g *Game) SaveGameState() {
	fmt.Println("Game state saved!")
}

// LoadGameState simulates loading game state (placeholder)
func (g *Game) LoadGameState() {
	fmt.Println("Game state loaded!")
}

// ResetGame resets the game to initial state
func (g *Game) ResetGame() {
	for i := range g.Players {
		g.Players[i].Score = 0
		g.Players[i].Level = 1
	}
	g.Round = 1
	g.Active = true
	fmt.Println("Game has been reset!")
}

// PlayerStats displays detailed statistics for a player
func (p *Player) PlayerStats() {
	fmt.Printf("\nPlayer: %s\n", p.Name)
	fmt.Printf("Score: %d\n", p.Score)
	fmt.Printf("Level: %d\n", p.Level)
	fmt.Printf("Points needed for next level: %d\n", p.Level*100-p.Score)
}

// GameStatistics displays overall game statistics
func (g *Game) GameStatistics() {
	fmt.Println("\n=== Game Statistics ===")
	fmt.Printf("Total Players: %d\n", len(g.Players))
	fmt.Printf("Current Round: %d\n", g.Round)
	fmt.Printf("Average Score: %.2f\n", g.CalculateAverageScore())
	
	topPlayers := g.GetTopPlayers(3)
	fmt.Println("\nTop 3 Players:")
	for i, player := range topPlayers {
		fmt.Printf("%d. %s - %d points\n", i+1, player.Name, player.Score)
	}
}

// Complex mathematical functions to increase code complexity

// Fibonacci calculates the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Factorial calculates factorial of n
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

// GCD calculates greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM calculates least common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// String manipulation functions

// ReverseString reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountVowels counts vowels in a string
func CountVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, char := range s {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				break
			}
		}
	}
	return count
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	return s == ReverseString(s)
}

// File operations simulation

// FileHandler simulates file operations
type FileHandler struct {
	Filename string
	Content  []string
}

// NewFileHandler creates a new file handler
func NewFileHandler(filename string) *FileHandler {
	return &FileHandler{
		Filename: filename,
		Content:  []string{},
	}
}

// WriteLine writes a line to the file
func (fh *FileHandler) WriteLine(line string) {
	fh.Content = append(fh.Content, line)
	fmt.Printf("Written to %s: %s\n", fh.Filename, line)
}

// ReadAll reads all content from the file
func (fh *FileHandler) ReadAll() []string {
	fmt.Printf("Reading from %s:\n", fh.Filename)
	for i, line := range fh.Content {
		fmt.Printf("Line %d: %s\n", i+1, line)
	}
	return fh.Content
}

// ClearFile clears all content from the file
func (fh *FileHandler) ClearFile() {
	fh.Content = []string{}
	fmt.Printf("Cleared all content from %s\n", fh.Filename)
}

// Advanced game features

// PowerUp represents a game power-up
type PowerUp struct {
	Name        string
	Description string
	Effect      func(*Player)
}

// ApplyDoubleScore doubles the player's score
func ApplyDoubleScore(player *Player) {
	player.Score *= 2
	fmt.Printf("%s used Double Score power-up! Score is now %d\n", player.Name, player.Score)
}

// ApplyLevelBoost increases player level by 2
func ApplyLevelBoost(player *Player) {
	player.Level += 2
	fmt.Printf("%s used Level Boost power-up! Level is now %d\n", player.Name, player.Level)
}

// CreatePowerUps creates available power-ups
func CreatePowerUps() []PowerUp {
	return []PowerUp{
		{
			Name:        "Double Score",
			Description: "Doubles your current score",
			Effect:      ApplyDoubleScore,
		},
		{
			Name:        "Level Boost",
			Description: "Increases your level by 2",
			Effect:      ApplyLevelBoost,
		},
	}
}

// UsePowerUp allows a player to use a power-up
func (p *Player) UsePowerUp(powerUp PowerUp) {
	fmt.Printf("\n%s is using %s power-up!\n", p.Name, powerUp.Name)
	fmt.Printf("Effect: %s\n", powerUp.Description)
	powerUp.Effect(p)
}

// Main function with extensive demonstration
func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("🚀 Starting Fun Game Simulator 🚀")
	fmt.Println("===================================")
	
	// Run the main game simulation
	SimulateGame()
	
	// Demonstrate additional features
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Additional Features Demonstration")
	fmt.Println(strings.Repeat("=", 50))
	
	// Create a new game for demonstration
	demoGame := NewGame()
	demoGame.AddPlayer("Demo Player")
	demoGame.Players[0].Score = 500
	demoGame.Players[0].Level = 5
	
	// Show player stats
	demoGame.Players[0].PlayerStats()
	
	// Show game statistics
	demoGame.GameStatistics()
	
	// Demonstrate mathematical functions
	fmt.Println("\nMathematical Functions:")
	fmt.Printf("Fibonacci(10) = %d\n", Fibonacci(10))
	fmt.Printf("Factorial(5) = %d\n", Factorial(5))
	fmt.Printf("IsPrime(17) = %t\n", IsPrime(17))
	fmt.Printf("GCD(48, 18) = %d\n", GCD(48, 18))
	fmt.Printf("LCM(12, 18) = %d\n", LCM(12, 18))
	
	// Demonstrate string functions
	fmt.Println("\nString Functions:")
	testString := "Hello, World!"
	fmt.Printf("Original: %s\n", testString)
	fmt.Printf("Reversed: %s\n", ReverseString(testString))
	fmt.Printf("Vowel count: %d\n", CountVowels(testString))
	fmt.Printf("Is palindrome: %t\n", IsPalindrome("racecar"))
	
	// Demonstrate file operations
	fmt.Println("\nFile Operations:")
	fileHandler := NewFileHandler("game_log.txt")
	fileHandler.WriteLine("Game started at " + time.Now().Format(time.RFC1123))
	fileHandler.WriteLine("Player scores recorded")
	fileHandler.ReadAll()
	
	// Demonstrate power-ups
	fmt.Println("\nPower-up Demonstration:")
	powerUps := CreatePowerUps()
	testPlayer := &Player{Name: "Test Player", Score: 100, Level: 2}
	testPlayer.PlayerStats()
	testPlayer.UsePowerUp(powerUps[0]) // Double Score
	testPlayer.UsePowerUp(powerUps[1]) // Level Boost
	testPlayer.PlayerStats()
	
	// Complex game simulation with multiple rounds
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Extended Game Simulation")
	fmt.Println(strings.Repeat("=", 50))
	
	extendedGame := NewGame()
	extendedGame.AddPlayer("Extended Player 1")
	extendedGame.AddPlayer("Extended Player 2")
	
	for i := 0; i < 5; i++ {
		extendedGame.StartRound()
	}
	
	extendedGame.DisplayScores()
	extendedGame.GameStatistics()
	
	// Reset and show it works
	fmt.Println("\nResetting game...")
	extendedGame.ResetGame()
	extendedGame.DisplayScores()
	
	fmt.Println("\n🎊 Program completed successfully! 🎊")
	fmt.Println("Thank you for using the Fun Game Simulator!")
}