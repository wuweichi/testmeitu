package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// TextStats holds comprehensive statistics about analyzed text
type TextStats struct {
	TotalChars      int
	TotalWords      int
	TotalLines      int
	TotalSentences  int
	UniqueWords     int
	WordFrequency   map[string]int
	CharFrequency   map[rune]int
	LongestWord     string
	ShortestWord    string
	AvgWordLength   float64
	AvgSentenceLen  float64
	ReadingTime     float64
	SpeakingTime    float64
	TopWords        []WordCount
	TopChars        []CharCount
}

// WordCount represents a word and its frequency
type WordCount struct {
	Word  string
	Count int
}

// CharCount represents a character and its frequency
type CharCount struct {
	Char  rune
	Count int
}

// ByCount implements sort.Interface for []WordCount based on Count
type ByCount []WordCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

// ByCharCount implements sort.Interface for []CharCount based on Count
type ByCharCount []CharCount

func (a ByCharCount) Len() int           { return len(a) }
func (a ByCharCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCharCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

// TextAnalyzer is the main analyzer struct
type TextAnalyzer struct {
	stats TextStats
}

// NewTextAnalyzer creates a new TextAnalyzer instance
func NewTextAnalyzer() *TextAnalyzer {
	return &TextAnalyzer{
		stats: TextStats{
			WordFrequency: make(map[string]int),
			CharFrequency: make(map[rune]int),
		},
	}
}

// AnalyzeText performs comprehensive text analysis
func (ta *TextAnalyzer) AnalyzeText(text string) {
	// Reset stats
	ta.stats = TextStats{
		WordFrequency: make(map[string]int),
		CharFrequency: make(map[rune]int),
	}

	// Basic character analysis
	ta.analyzeCharacters(text)

	// Word analysis
	ta.analyzeWords(text)

	// Sentence analysis
	ta.analyzeSentences(text)

	// Calculate derived statistics
	ta.calculateDerivedStats()

	// Sort top words and characters
	ta.sortTopItems()
}

func (ta *TextAnalyzer) analyzeCharacters(text string) {
	ta.stats.TotalChars = len(text)
	
	for _, char := range text {
		if unicode.IsSpace(char) {
			continue
		}
		ta.stats.CharFrequency[char]++
	}
}

func (ta *TextAnalyzer) analyzeWords(text string) {
	// Split text into words using regex to handle punctuation
	re := regexp.MustCompile(`[\w']+`)
	words := re.FindAllString(text, -1)
	
	ta.stats.TotalWords = len(words)
	
	if len(words) == 0 {
		ta.stats.LongestWord = ""
		ta.stats.ShortestWord = ""
		return
	}
	
	ta.stats.LongestWord = words[0]
	ta.stats.ShortestWord = words[0]
	
	uniqueWords := make(map[string]bool)
	totalWordLength := 0
	
	for _, word := range words {
		word = strings.ToLower(word)
		ta.stats.WordFrequency[word]++
		uniqueWords[word] = true
		
		wordLength := len(word)
		totalWordLength += wordLength
		
		if wordLength > len(ta.stats.LongestWord) {
			ta.stats.LongestWord = word
		}
		if wordLength < len(ta.stats.ShortestWord) {
			ta.stats.ShortestWord = word
		}
	}
	
	ta.stats.UniqueWords = len(uniqueWords)
	ta.stats.AvgWordLength = float64(totalWordLength) / float64(len(words))
}

func (ta *TextAnalyzer) analyzeSentences(text string) {
	// Split text into lines
	lines := strings.Split(text, "\n")
	ta.stats.TotalLines = len(lines)
	
	// Split text into sentences using regex
	re := regexp.MustCompile(`[.!?]+`)
	sentences := re.Split(text, -1)
	
	// Filter out empty sentences
	validSentences := []string{}
	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		if len(sentence) > 0 {
			validSentences = append(validSentences, sentence)
		}
	}
	
	ta.stats.TotalSentences = len(validSentences)
	
	if len(validSentences) > 0 {
		totalSentenceLength := 0
		for _, sentence := range validSentences {
			totalSentenceLength += len(sentence)
		}
		ta.stats.AvgSentenceLen = float64(totalSentenceLength) / float64(len(validSentences))
	}
}

func (ta *TextAnalyzer) calculateDerivedStats() {
	// Calculate reading time (average reading speed: 200 words per minute)
	if ta.stats.TotalWords > 0 {
		ta.stats.ReadingTime = float64(ta.stats.TotalWords) / 200.0
	}
	
	// Calculate speaking time (average speaking speed: 150 words per minute)
	if ta.stats.TotalWords > 0 {
		ta.stats.SpeakingTime = float64(ta.stats.TotalWords) / 150.0
	}
}

func (ta *TextAnalyzer) sortTopItems() {
	// Sort top 10 words
	wordCounts := make([]WordCount, 0, len(ta.stats.WordFrequency))
	for word, count := range ta.stats.WordFrequency {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}
	sort.Sort(ByCount(wordCounts))
	
	if len(wordCounts) > 10 {
		ta.stats.TopWords = wordCounts[:10]
	} else {
		ta.stats.TopWords = wordCounts
	}
	
	// Sort top 10 characters
	charCounts := make([]CharCount, 0, len(ta.stats.CharFrequency))
	for char, count := range ta.stats.CharFrequency {
		charCounts = append(charCounts, CharCount{Char: char, Count: count})
	}
	sort.Sort(ByCharCount(charCounts))
	
	if len(charCounts) > 10 {
		ta.stats.TopChars = charCounts[:10]
	} else {
		ta.stats.TopChars = charCounts
	}
}

// GetStats returns the current statistics
func (ta *TextAnalyzer) GetStats() TextStats {
	return ta.stats
}

// DisplayStats prints statistics in a formatted way
func (ta *TextAnalyzer) DisplayStats() {
	stats := ta.stats
	
	fmt.Println("=== TEXT ANALYSIS RESULTS ===")
	fmt.Printf("Total Characters: %d\n", stats.TotalChars)
	fmt.Printf("Total Words: %d\n", stats.TotalWords)
	fmt.Printf("Total Lines: %d\n", stats.TotalLines)
	fmt.Printf("Total Sentences: %d\n", stats.TotalSentences)
	fmt.Printf("Unique Words: %d\n", stats.UniqueWords)
	fmt.Printf("Longest Word: '%s' (%d characters)\n", stats.LongestWord, len(stats.LongestWord))
	fmt.Printf("Shortest Word: '%s' (%d characters)\n", stats.ShortestWord, len(stats.ShortestWord))
	fmt.Printf("Average Word Length: %.2f characters\n", stats.AvgWordLength)
	fmt.Printf("Average Sentence Length: %.2f characters\n", stats.AvgSentenceLen)
	fmt.Printf("Estimated Reading Time: %.2f minutes\n", stats.ReadingTime)
	fmt.Printf("Estimated Speaking Time: %.2f minutes\n", stats.SpeakingTime)
	
	fmt.Println("\n=== TOP 10 MOST FREQUENT WORDS ===")
	for i, wc := range stats.TopWords {
		fmt.Printf("%d. '%s' - %d occurrences\n", i+1, wc.Word, wc.Count)
	}
	
	fmt.Println("\n=== TOP 10 MOST FREQUENT CHARACTERS ===")
	for i, cc := range stats.TopChars {
		fmt.Printf("%d. '%c' - %d occurrences\n", i+1, cc.Char, cc.Count)
	}
}

// FileProcessor handles file operations
type FileProcessor struct {
	analyzer *TextAnalyzer
}

// NewFileProcessor creates a new FileProcessor instance
func NewFileProcessor() *FileProcessor {
	return &FileProcessor{
		analyzer: NewTextAnalyzer(),
	}
}

// ProcessFile reads and analyzes a text file
func (fp *FileProcessor) ProcessFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	var content strings.Builder
	
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}
	
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	
	fp.analyzer.AnalyzeText(content.String())
	return nil
}

// GetAnalyzer returns the text analyzer
func (fp *FileProcessor) GetAnalyzer() *TextAnalyzer {
	return fp.analyzer
}

// InteractiveMode provides an interactive text analysis interface
func InteractiveMode() {
	fmt.Println("=== INTERACTIVE TEXT ANALYZER ===")
	fmt.Println("Enter your text (type 'END' on a new line to finish):")
	
	scanner := bufio.NewScanner(os.Stdin)
	var content strings.Builder
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "END" {
			break
		}
		content.WriteString(line)
		content.WriteString("\n")
	}
	
	analyzer := NewTextAnalyzer()
	analyzer.AnalyzeText(content.String())
	analyzer.DisplayStats()
}

// BatchMode processes multiple files
func BatchMode(filenames []string) {
	fmt.Printf("=== BATCH PROCESSING %d FILES ===\n", len(filenames))
	
	for i, filename := range filenames {
		fmt.Printf("\nProcessing file %d/%d: %s\n", i+1, len(filenames), filename)
		
		processor := NewFileProcessor()
		err := processor.ProcessFile(filename)
		if err != nil {
			fmt.Printf("Error processing file: %v\n", err)
			continue
		}
		
		processor.GetAnalyzer().DisplayStats()
	}
}

// GenerateSampleText creates sample text for testing
func GenerateSampleText() string {
	return `The quick brown fox jumps over the lazy dog. This sentence contains all letters of the English alphabet.

Programming is the process of creating a set of instructions that tell a computer how to perform a task. Programming can be done using many programming languages such as Java, Python, C++, and Go.

Go, also known as Golang, is a programming language designed at Google. It is statically typed and compiled. Go is known for its simplicity, efficiency, and concurrency features.

Text analysis involves examining text to extract meaningful information. This can include word frequency analysis, character counting, sentiment analysis, and more. Advanced text analyzers can provide insights into writing style, readability, and content structure.

Natural language processing (NLP) is a field of computer science concerned with the interactions between computers and human language. NLP techniques are used in many applications including search engines, machine translation, and text analysis tools.`
}

// DemoMode runs a demonstration with sample text
func DemoMode() {
	fmt.Println("=== DEMONSTRATION MODE ===")
	fmt.Println("Analyzing sample text...")
	
	sampleText := GenerateSampleText()
	analyzer := NewTextAnalyzer()
	analyzer.AnalyzeText(sampleText)
	analyzer.DisplayStats()
}

// Help displays usage information
func Help() {
	fmt.Println("=== TEXT ANALYZER HELP ===")
	fmt.Println("Usage:")
	fmt.Println("  -i, --interactive  Enter interactive mode")
	fmt.Println("  -f, --file <file>  Analyze a specific file")
	fmt.Println("  -b, --batch <files> Analyze multiple files (comma-separated)")
	fmt.Println("  -d, --demo         Run demonstration with sample text")
	fmt.Println("  -h, --help         Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  ./text_analyzer -i")
	fmt.Println("  ./text_analyzer -f document.txt")
	fmt.Println("  ./text_analyzer -b file1.txt,file2.txt,file3.txt")
	fmt.Println("  ./text_analyzer -d")
}

func main() {
	if len(os.Args) < 2 {
		Help()
		return
	}
	
	switch os.Args[1] {
	case "-i", "--interactive":
		InteractiveMode()
	case "-f", "--file":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please specify a filename")
			os.Exit(1)
		}
		processor := NewFileProcessor()
		err := processor.ProcessFile(os.Args[2])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		processor.GetAnalyzer().DisplayStats()
	case "-b", "--batch":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please specify filenames")
			os.Exit(1)
		}
		filenames := strings.Split(os.Args[2], ",")
		BatchMode(filenames)
	case "-d", "--demo":
		DemoMode()
	case "-h", "--help":
		Help()
	default:
		fmt.Printf("Unknown option: %s\n", os.Args[1])
		Help()
		os.Exit(1)
	}
}