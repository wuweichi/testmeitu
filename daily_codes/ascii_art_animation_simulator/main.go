package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// ==================== 数据结构定义 ====================
type Vector2D struct {
	X, Y float64
}

func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.X + other.X, v.Y + other.Y}
}

func (v Vector2D) Subtract(other Vector2D) Vector2D {
	return Vector2D{v.X - other.X, v.Y - other.Y}
}

func (v Vector2D) Multiply(scalar float64) Vector2D {
	return Vector2D{v.X * scalar, v.Y * scalar}
}

func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2D) Normalize() Vector2D {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2D{0, 0}
	}
	return Vector2D{v.X / mag, v.Y / mag}
}

type Particle struct {
	Position Vector2D
	Velocity Vector2D
	Char    rune
	Color   string
	Life    int
}

func NewParticle(x, y float64, char rune, color string, life int) *Particle {
	return &Particle{
		Position: Vector2D{x, y},
		Velocity: Vector2D{rand.Float64()*2 - 1, rand.Float64()*2 - 1},
		Char:    char,
		Color:   color,
		Life:    life,
	}
}

type AnimationSystem struct {
	Particles []*Particle
	Width     int
	Height    int
	Mutex     sync.Mutex
}

func NewAnimationSystem(width, height int) *AnimationSystem {
	return &AnimationSystem{
		Particles: make([]*Particle, 0),
		Width:     width,
		Height:    height,
	}
}

func (as *AnimationSystem) AddParticle(p *Particle) {
	as.Mutex.Lock()
	defer as.Mutex.Unlock()
	as.Particles = append(as.Particles, p)
}

func (as *AnimationSystem) Update() {
	as.Mutex.Lock()
	defer as.Mutex.Unlock()
	aliveParticles := make([]*Particle, 0, len(as.Particles))
	for _, p := range as.Particles {
		p.Life--
		if p.Life <= 0 {
			continue
		}
		p.Position = p.Position.Add(p.Velocity)
		if p.Position.X < 0 || p.Position.X >= float64(as.Width) {
			p.Velocity.X = -p.Velocity.X
		}
		if p.Position.Y < 0 || p.Position.Y >= float64(as.Height) {
			p.Velocity.Y = -p.Velocity.Y
		}
		aliveParticles = append(aliveParticles, p)
	}
	as.Particles = aliveParticles
}

func (as *AnimationSystem) Render() string {
	as.Mutex.Lock()
	defer as.Mutex.Unlock()
	grid := make([][]rune, as.Height)
	for i := range grid {
		grid[i] = make([]rune, as.Width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	colorMap := make([][]string, as.Height)
	for i := range colorMap {
		colorMap[i] = make([]string, as.Width)
		for j := range colorMap[i] {
			colorMap[i][j] = "\033[0m"
		}
	}
	for _, p := range as.Particles {
		x, y := int(p.Position.X), int(p.Position.Y)
		if x >= 0 && x < as.Width && y >= 0 && y < as.Height {
			grid[y][x] = p.Char
			colorMap[y][x] = p.Color
		}
	}
	var sb strings.Builder
	for y := 0; y < as.Height; y++ {
		for x := 0; x < as.Width; x++ {
			sb.WriteString(colorMap[y][x])
			sb.WriteRune(grid[y][x])
		}
		sb.WriteString("\033[0m\n")
	}
	return sb.String()
}

// ==================== 工具函数 ====================
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getTerminalSize() (int, int, error) {
	var width, height int
	if runtime.GOOS == "windows" {
		cmd := exec.Command("mode", "con")
		output, err := cmd.Output()
		if err != nil {
			return 80, 24, err
		}
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Columns:") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					width, _ = strconv.Atoi(fields[1])
				}
			}
			if strings.Contains(line, "Lines:") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					height, _ = strconv.Atoi(fields[1])
				}
			}
		}
	} else {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		output, err := cmd.Output()
		if err != nil {
			return 80, 24, err
		}
		dimensions := strings.Split(strings.TrimSpace(string(output)), " ")
		if len(dimensions) >= 2 {
			height, _ = strconv.Atoi(dimensions[0])
			width, _ = strconv.Atoi(dimensions[1])
		}
	}
	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}
	return width, height, nil
}

func randomColor() string {
	colors := []string{
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[91m", // Bright Red
		"\033[92m", // Bright Green
		"\033[93m", // Bright Yellow
		"\033[94m", // Bright Blue
		"\033[95m", // Bright Magenta
		"\033[96m", // Bright Cyan
	}
	return colors[rand.Intn(len(colors))]
}

func randomChar() rune {
	chars := []rune("@#$%&*+=~^<>?/|\\.:;,\"'`!")
	return chars[rand.Intn(len(chars))]
}

// ==================== 动画模式 ====================
type AnimationMode int

const (
	ModeFireworks AnimationMode = iota
	ModeRain
	ModeSwarm
	ModeTextScroller
)

func (as *AnimationSystem) runFireworks() {
	for i := 0; i < 50; i++ {
		x := rand.Float64() * float64(as.Width)
		y := float64(as.Height - 1)
		char := randomChar()
		color := randomColor()
		life := rand.Intn(100) + 50
		as.AddParticle(NewParticle(x, y, char, color, life))
	}
}

func (as *AnimationSystem) runRain() {
	for i := 0; i < 20; i++ {
		x := rand.Float64() * float64(as.Width)
		y := 0.0
		char := '|'
		color := "\033[36m"
		life := rand.Intn(50) + 30
		velocity := Vector2D{0, rand.Float64()*2 + 1}
		p := NewParticle(x, y, char, color, life)
		p.Velocity = velocity
		as.AddParticle(p)
	}
}

func (as *AnimationSystem) runSwarm() {
	centerX := float64(as.Width) / 2
	centerY := float64(as.Height) / 2
	for i := 0; i < 30; i++ {
		angle := rand.Float64() * 2 * math.Pi
		distance := rand.Float64() * 10
		x := centerX + math.Cos(angle)*distance
		y := centerY + math.Sin(angle)*distance
		char := '*' if rand.Intn(2) == 0 else 'o'
		color := "\033[33m"
		life := rand.Intn(200) + 100
		velocity := Vector2D{rand.Float64()*2 - 1, rand.Float64()*2 - 1}
		p := NewParticle(x, y, char, color, life)
		p.Velocity = velocity
		as.AddParticle(p)
	}
}

func (as *AnimationSystem) runTextScroller(text string) {
	x := float64(as.Width)
	y := float64(as.Height / 2)
	for _, ch := range text {
		if utf8.RuneLen(ch) > 1 {
			continue
		}
		color := randomColor()
		life := 150
		velocity := Vector2D{-1, 0}
		p := NewParticle(x, y, ch, color, life)
		p.Velocity = velocity
		as.AddParticle(p)
		x += 1.5
	}
}

// ==================== 主程序 ====================
func main() {
	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 获取终端尺寸
	width, height, err := getTerminalSize()
	if err != nil {
		width, height = 80, 24
		fmt.Printf("Warning: Could not get terminal size, using default %dx%d\n", width, height)
	}
	// 调整尺寸以避免边界问题
	width = int(float64(width) * 0.9)
	height = int(float64(height) * 0.8)
	if width < 20 {
		width = 20
	}
	if height < 10 {
		height = 10
	}

	// 创建动画系统
	animSystem := NewAnimationSystem(width, height)

	// 动画模式循环
	mode := ModeFireworks
	modeDuration := 300 // 每个模式的帧数
	frameCount := 0
	textIndex := 0
	texts := []string{
		"ASCII ART ANIMATION SIMULATOR",
		"GO LANG IS AWESOME!",
		"CREATIVE CODING FUN",
		"PARTICLES IN MOTION",
		"TERMINAL MAGIC",
	}

	// 主循环
	for {
		clearScreen()

		// 根据模式添加新粒子
		switch mode {
		case ModeFireworks:
			if frameCount%10 == 0 {
				animSystem.runFireworks()
			}
		case ModeRain:
			if frameCount%5 == 0 {
				animSystem.runRain()
			}
		case ModeSwarm:
			if frameCount%15 == 0 {
				animSystem.runSwarm()
			}
		case ModeTextScroller:
			if frameCount%20 == 0 {
				animSystem.runTextScroller(texts[textIndex])
			}
		}

		// 更新和渲染
		animSystem.Update()
		renderOutput := animSystem.Render()

		// 显示信息
		info := fmt.Sprintf("Mode: %d | Particles: %d | Frame: %d | Size: %dx%d\n",
			mode, len(animSystem.Particles), frameCount, width, height)
		renderOutput = info + renderOutput
		fmt.Print(renderOutput)

		// 模式切换
		frameCount++
		if frameCount >= modeDuration {
			frameCount = 0
			mode = (mode + 1) % 4
			if mode == ModeTextScroller {
				textIndex = (textIndex + 1) % len(texts)
			}
			// 清空粒子以便新模式开始
			animSystem.Particles = make([]*Particle, 0)
		}

		// 控制帧率
		time.Sleep(50 * time.Millisecond)

		// 简单退出条件（例如运行一定时间后退出，实际中可改为按键检测）
		if frameCount > 1000 {
			fmt.Println("Animation finished. Exiting...")
			break
		}
	}
}
