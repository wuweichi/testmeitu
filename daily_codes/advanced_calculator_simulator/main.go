package main

import (
    "bufio"
    "fmt"
    "math"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
    "unicode"
)

// Constants for the program
const (
    Version = "1.0.0"
    MaxHistory = 100
    MaxVariables = 50
    MaxFunctions = 20
)

// Struct to represent a variable
type Variable struct {
    Name  string
    Value float64
}

// Struct to represent a function
type Function struct {
    Name     string
    Params   []string
    Body     string
}

// Struct to represent the calculator state
type Calculator struct {
    Variables   map[string]float64
    Functions   map[string]Function
    History     []string
    HistoryPos  int
    Random      *rand.Rand
}

// Initialize a new calculator
func NewCalculator() *Calculator {
    calc := &Calculator{
        Variables: make(map[string]float64),
        Functions: make(map[string]Function),
        History:   make([]string, 0, MaxHistory),
        Random:    rand.New(rand.NewSource(time.Now().UnixNano())),
    }
    // Initialize built-in variables
    calc.Variables["pi"] = math.Pi
    calc.Variables["e"] = math.E
    calc.Variables["phi"] = (1 + math.Sqrt(5)) / 2
    // Initialize built-in functions (as placeholders)
    calc.Functions["sin"] = Function{Name: "sin", Params: []string{"x"}, Body: ""}
    calc.Functions["cos"] = Function{Name: "cos", Params: []string{"x"}, Body: ""}
    calc.Functions["tan"] = Function{Name: "tan", Params: []string{"x"}, Body: ""}
    calc.Functions["log"] = Function{Name: "log", Params: []string{"x"}, Body: ""}
    calc.Functions["sqrt"] = Function{Name: "sqrt", Params: []string{"x"}, Body: ""}
    return calc
}

// Add an entry to history
func (c *Calculator) AddHistory(entry string) {
    if len(c.History) >= MaxHistory {
        c.History = c.History[1:]
    }
    c.History = append(c.History, entry)
    c.HistoryPos = len(c.History)
}

// Display history
func (c *Calculator) ShowHistory() {
    fmt.Println("Calculation History:")
    for i, entry := range c.History {
        fmt.Printf("%d: %s\n", i+1, entry)
    }
}

// Evaluate a mathematical expression
func (c *Calculator) Evaluate(expr string) (float64, error) {
    // Remove whitespace
    expr = strings.ReplaceAll(expr, " ", "")
    if expr == "" {
        return 0, fmt.Errorf("empty expression")
    }
    // Check for variable assignment
    if strings.Contains(expr, "=") {
        parts := strings.Split(expr, "=")
        if len(parts) != 2 {
            return 0, fmt.Errorf("invalid assignment syntax")
        }
        varName := strings.TrimSpace(parts[0])
        varExpr := strings.TrimSpace(parts[1])
        // Validate variable name
        if !isValidIdentifier(varName) {
            return 0, fmt.Errorf("invalid variable name: %s", varName)
        }
        // Evaluate the expression
        value, err := c.Evaluate(varExpr)
        if err != nil {
            return 0, err
        }
        c.Variables[varName] = value
        c.AddHistory(fmt.Sprintf("%s = %v", varName, value))
        return value, nil
    }
    // Check for function definition (simplified)
    if strings.HasPrefix(expr, "func") {
        // Placeholder for function definition
        return 0, fmt.Errorf("function definition not fully implemented in this version")
    }
    // Evaluate basic arithmetic and functions
    return c.evaluateExpression(expr)
}

// Helper to evaluate an expression recursively
func (c *Calculator) evaluateExpression(expr string) (float64, error) {
    // Handle parentheses
    for strings.Contains(expr, "(") && strings.Contains(expr, ")") {
        start := strings.LastIndex(expr, "(")
        end := strings.Index(expr[start:], ")") + start
        if end <= start {
            return 0, fmt.Errorf("mismatched parentheses")
        }
        subExpr := expr[start+1 : end]
        subResult, err := c.evaluateExpression(subExpr)
        if err != nil {
            return 0, err
        }
        expr = expr[:start] + fmt.Sprintf("%v", subResult) + expr[end+1:]
    }
    // Handle functions (simplified)
    for funcName, funcDef := range c.Functions {
        if strings.Contains(expr, funcName+"(") {
            // Placeholder: just call built-in math functions
            // In a full implementation, parse arguments and evaluate
            return 0, fmt.Errorf("function evaluation not fully implemented")
        }
    }
    // Handle basic arithmetic: +, -, *, /, ^
    // This is a simplified evaluator; a real one would use a proper parser
    // For demonstration, we'll do a simple left-to-right evaluation with precedence
    // Split by operators (simplified)
    tokens := tokenize(expr)
    if len(tokens) == 0 {
        return 0, fmt.Errorf("no tokens in expression")
    }
    // Evaluate numbers and variables
    values := make([]float64, 0, len(tokens))
    operators := make([]string, 0, len(tokens)-1)
    for _, token := range tokens {
        if isOperator(token) {
            operators = append(operators, token)
        } else {
            value, err := c.parseValue(token)
            if err != nil {
                return 0, err
            }
            values = append(values, value)
        }
    }
    if len(values) != len(operators)+1 {
        return 0, fmt.Errorf("invalid expression syntax")
    }
    // Apply operators with precedence (simplified: ^, */, +-)
    result := values[0]
    for i := 0; i < len(operators); i++ {
        nextVal := values[i+1]
        switch operators[i] {
        case "^":
            result = math.Pow(result, nextVal)
        case "*":
            result *= nextVal
        case "/":
            if nextVal == 0 {
                return 0, fmt.Errorf("division by zero")
            }
            result /= nextVal
        case "+":
            result += nextVal
        case "-":
            result -= nextVal
        default:
            return 0, fmt.Errorf("unknown operator: %s", operators[i])
        }
    }
    return result, nil
}

// Parse a token to a float64 value
func (c *Calculator) parseValue(token string) (float64, error) {
    // Check if it's a number
    if val, err := strconv.ParseFloat(token, 64); err == nil {
        return val, nil
    }
    // Check if it's a variable
    if val, ok := c.Variables[token]; ok {
        return val, nil
    }
    // Check for special constants
    switch token {
    case "pi":
        return math.Pi, nil
    case "e":
        return math.E, nil
    case "phi":
        return (1 + math.Sqrt(5)) / 2, nil
    }
    return 0, fmt.Errorf("unknown token: %s", token)
}

// Tokenize an expression (simplified)
func tokenize(expr string) []string {
    var tokens []string
    var current strings.Builder
    for _, ch := range expr {
        if isOperator(string(ch)) {
            if current.Len() > 0 {
                tokens = append(tokens, current.String())
                current.Reset()
            }
            tokens = append(tokens, string(ch))
        } else {
            current.WriteRune(ch)
        }
    }
    if current.Len() > 0 {
        tokens = append(tokens, current.String())
    }
    return tokens
}

// Check if a string is an operator
func isOperator(s string) bool {
    return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}

// Check if a string is a valid identifier for variables/functions
func isValidIdentifier(s string) bool {
    if len(s) == 0 || !unicode.IsLetter(rune(s[0])) {
        return false
    }
    for _, ch := range s {
        if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
            return false
        }
    }
    return true
}

// Display all variables
func (c *Calculator) ShowVariables() {
    fmt.Println("Variables:")
    for name, value := range c.Variables {
        fmt.Printf("  %s = %v\n", name, value)
    }
}

// Display all functions
func (c *Calculator) ShowFunctions() {
    fmt.Println("Functions:")
    for name, fn := range c.Functions {
        fmt.Printf("  %s(%s)\n", name, strings.Join(fn.Params, ", "))
    }
}

// Generate a random math problem
func (c *Calculator) GenerateProblem() string {
    problems := []string{
        "What is 2 + 2?",
        "Calculate the area of a circle with radius 5.",
        "Solve for x: 3x + 5 = 20",
        "What is the square root of 144?",
        "Compute 10! (factorial).",
    }
    return problems[c.Random.Intn(len(problems))]
}

// Main interactive loop
func (c *Calculator) Run() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Advanced Calculator Simulator v%s\n", Version)
    fmt.Println("Type 'help' for commands, 'exit' to quit.")
    for {
        fmt.Print(">> ")
        if !scanner.Scan() {
            break
        }
        input := scanner.Text()
        input = strings.TrimSpace(input)
        if input == "" {
            continue
        }
        if input == "exit" || input == "quit" {
            fmt.Println("Goodbye!")
            break
        }
        switch input {
        case "help":
            c.ShowHelp()
        case "history":
            c.ShowHistory()
        case "variables":
            c.ShowVariables()
        case "functions":
            c.ShowFunctions()
        case "problem":
            fmt.Println("Random Problem:", c.GenerateProblem())
        case "clear":
            c.History = make([]string, 0, MaxHistory)
            fmt.Println("History cleared.")
        case "version":
            fmt.Printf("Version: %s\n", Version)
        default:
            result, err := c.Evaluate(input)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Result: %v\n", result)
            }
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
    }
}

// Display help information
func (c *Calculator) ShowHelp() {
    fmt.Println("Available commands:")
    fmt.Println("  <expression>   - Evaluate a mathematical expression (e.g., 2+3*4)")
    fmt.Println("  var = expr     - Assign a value to a variable (e.g., x = 10)")
    fmt.Println("  help           - Show this help message")
    fmt.Println("  history        - Show calculation history")
    fmt.Println("  variables      - List all variables")
    fmt.Println("  functions      - List all functions")
    fmt.Println("  problem        - Generate a random math problem")
    fmt.Println("  clear          - Clear history")
    fmt.Println("  version        - Show program version")
    fmt.Println("  exit or quit   - Exit the program")
    fmt.Println("\nExamples:")
    fmt.Println("  >> 3 + 4 * 2")
    fmt.Println("  >> a = 5")
    fmt.Println("  >> b = a^2")
    fmt.Println("  >> sin(pi/2)")
}

// Main function
func main() {
    calc := NewCalculator()
    calc.Run()
}
