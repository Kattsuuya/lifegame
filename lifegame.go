package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Cell セル
type Cell struct {
}

//Field ゲームフィールド
type Field struct {
	height int
	width  int
	status [][]bool
}

func newField(height int, width int, initRate float64) *Field {
	field := new(Field)
	field.height = height
	field.width = width
	fieldTable := make([][]bool, height)
	rand.Seed(time.Now().UnixNano())
	for h := 0; h < height; h++ {
		fieldTable[h] = make([]bool, width)
		for w := 0; w < width; w++ {
			if rand.Float64() < initRate {
				fieldTable[h][w] = true
			} else {
				fieldTable[h][w] = false
			}
		}
	}
	field.status = fieldTable
	return field
}

//LifeGame ライフゲーム
type LifeGame struct {
	gameField *Field
	history   []Field
	interval  float64
}

func newLifeGame(height int, width int, initRate float64, interval float64) *LifeGame {
	lifeGame := new(LifeGame)
	gamefield := newField(height, width, initRate)
	lifeGame.gameField = gamefield
	lifeGame.interval = interval
	return lifeGame
}

func parseCommandLine() (int, int, float64, float64) {
	args := os.Args
	for _, arg := range args {
		if arg == "-h" || arg == "--help" || len(args) == 1 {
			fmt.Println(`
positional arguments:
	height      Field height
	width       Field width
	init_rate   Percentage of surviving cells of the first generation
	interval    Time to evolve to the next generatio
optional arguments:
	-h, --help  show this help message and exit`)
			os.Exit(0)
		}
	}
	height, _ := strconv.Atoi(args[1])
	width, _ := strconv.Atoi(args[2])
	initRate, _ := strconv.ParseFloat(args[3], 64)
	interval, _ := strconv.ParseFloat(args[4], 64)
	return height, width, initRate, interval
}

func main() {
	height, width, initRate, interval := parseCommandLine()
	lifegame := newLifeGame(height, width, initRate, interval)
	fmt.Println(lifegame.gameField.status)
}
