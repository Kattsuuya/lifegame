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

func createFieldFrame(height int, width int) *Field {
	field := new(Field)
	field.height = height
	field.width = width
	field.status = make([][]bool, height)
	for h := 0; h < height; h++ {
		field.status[h] = make([]bool, width)
	}
	return field
}

func (field *Field) initFieldStatus(initRate float64) *Field {
	rand.Seed(time.Now().UnixNano())
	for h := 0; h < field.height; h++ {
		for w := 0; w < field.width; w++ {
			if rand.Float64() < initRate {
				field.status[h][w] = true
			} else {
				field.status[h][w] = false
			}
		}
	}
	return field
}

func (field *Field) printField() {
	for h := 0; h < field.height; h++ {
		for w := 0; w < field.width; w++ {
			if field.status[h][w] == true {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println("")
	}
}

//LifeGame ライフゲーム
type LifeGame struct {
	currentField   *Field
	history        []Field
	intervalSecond int
}

func newLifeGame(height int, width int, initRate float64, interval int) *LifeGame {
	lifeGame := new(LifeGame)
	gamefield := createFieldFrame(height, width).initFieldStatus(initRate)

	lifeGame.currentField = gamefield
	lifeGame.intervalSecond = interval
	return lifeGame
}

// func nextFrame(game *LifeGame) {
// 	nextFrame := createFieldFrame(game.currentField.height, game.currentField.width)
// }
func mainLoop(game *LifeGame) {

	time.Sleep(time.Duration(game.intervalSecond) * time.Second)
}

func parseCommandLine() (int, int, float64, int) {
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
	interval, _ := strconv.Atoi(args[4])
	return height, width, initRate, interval
}

func main() {
	height, width, initRate, interval := parseCommandLine()
	lifegame := newLifeGame(height, width, initRate, interval)
	lifegame.currentField.printField()
}
