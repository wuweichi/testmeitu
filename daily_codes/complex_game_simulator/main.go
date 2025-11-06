package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name     string
	Health   int
	Strength int
	Agility  int
}

func (c *Character) Attack(target *Character) int {
	damage := c.Strength + rand.Intn(10)
	target.Health -= damage
	return damage
}

func (c *Character) IsAlive() bool {
	return c.Health > 0
}

func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 创建角色
	player := Character{
		Name:     "Hero",
		Health:   100,
		Strength: 20,
		Agility:  15,
	}

	enemy := Character{
		Name:     "Monster",
		Health:   80,
		Strength: 18,
		Agility:  10,
	}

	// 游戏循环
	round := 1
	for player.IsAlive() && enemy.IsAlive() {
		fmt.Printf("\n--- Round %d ---\n", round)
		
		// 玩家攻击
		if player.IsAlive() {
			damage := player.Attack(&enemy)
			fmt.Printf("%s attacks %s for %d damage!\n", player.Name, enemy.Name, damage)
			fmt.Printf("%s health: %d\n", enemy.Name, enemy.Health)
		}
		
		// 敌人攻击
		if enemy.IsAlive() {
			damage := enemy.Attack(&player)
			fmt.Printf("%s attacks %s for %d damage!\n", enemy.Name, player.Name, damage)
			fmt.Printf("%s health: %d\n", player.Name, player.Health)
		}
		
		round++
		time.Sleep(1 * time.Second)
	}

	// 游戏结束
	fmt.Println("\n--- Game Over ---")
	if player.IsAlive() {
		fmt.Printf("%s wins!\n", player.Name)
	} else {
		fmt.Printf("%s wins!\n", enemy.Name)
	}
}