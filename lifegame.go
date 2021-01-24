package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Field ゲームフィールド
type Field struct {
	height int
	width  int
	status [][]int
}

func createFieldFrame(height int, width int) *Field {
	field := new(Field)
	field.height = height
	field.width = width
	field.status = make([][]int, height)
	for h := 0; h < height; h++ {
		field.status[h] = make([]int, width)
	}
	return field
}

func (field *Field) initFieldStatus(initRate float64) *Field {
	rand.Seed(time.Now().UnixNano())
	for h := 0; h < field.height; h++ {
		for w := 0; w < field.width; w++ {
			if rand.Float64() < initRate {
				field.status[h][w] = 1
			} else {
				field.status[h][w] = 0
			}
		}
	}
	return field
}

func (field *Field) printField() {
	for h := 0; h < field.height; h++ {
		for w := 0; w < field.width; w++ {
			if field.status[h][w] == 1 {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println("")
	}
}

func (field *Field) countAroundLife(h int, w int) int {
	result := 0
	for hIterator := -1; hIterator <= 1; hIterator++ {
		for wIterator := -1; wIterator <= 1; wIterator++ {
			if hIterator == 0 && wIterator == 0 {
				continue
			}
			if h+hIterator >= 0 && w+wIterator >= 0 && h+hIterator < field.height && w+wIterator < field.width {
				result += field.status[h+hIterator][w+wIterator]
			}
		}
	}
	return result
}

//LifeGame ライフゲーム
type LifeGame struct {
	currentField   *Field
	lastField      *Field
	intervalSecond int
}

func newLifeGame(height int, width int, initRate float64, interval int) *LifeGame {
	lifeGame := new(LifeGame)
	gamefield := createFieldFrame(height, width).initFieldStatus(initRate)

	lifeGame.currentField = gamefield
	lifeGame.intervalSecond = interval
	return lifeGame
}

func (game *LifeGame) nextFrame() (*Field, *Field) {
	nextFrame := createFieldFrame(game.currentField.height, game.currentField.width)
	for h := 0; h < nextFrame.height; h++ {
		for w := 0; w < nextFrame.width; w++ {
			AroundLifeCount := game.currentField.countAroundLife(h, w)
			var nextCell int
			if game.currentField.status[h][w] == 1 {
				if AroundLifeCount == 2 || AroundLifeCount == 3 {
					nextCell = 1
				} else {
					nextCell = 0
				}
			} else if game.currentField.status[h][w] == 0 {
				if AroundLifeCount == 3 {
					// fmt.Println("debug", h, w, AroundLifeCount)
					nextCell = 1
				} else {
					nextCell = 0
				}
			} else {
				fmt.Println("エラーが発生しました")
			}
			nextFrame.status[h][w] = nextCell

		}
	}
	return nextFrame, game.currentField
}

func (game *LifeGame) isChange(lastField *Field) bool {
	for hIndex, line := range game.currentField.status {
		for wIndex, cell := range line {
			if cell == lastField.status[hIndex][wIndex] {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func (game *LifeGame) mainLoop() {
	i := 1
	for {
		fmt.Println("step", i)
		game.currentField.printField()
		game.currentField, game.lastField = game.nextFrame()
		time.Sleep(time.Duration(game.intervalSecond) * time.Second)
		if !game.isChange(game.lastField) {
			break
		}
		fmt.Printf("\033[%dA", game.currentField.height+1)
		i++
	}
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
	lifegame.mainLoop()
}
