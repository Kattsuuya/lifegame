package components

import (
	"fmt"
	"time"
)

//LifeGame ライフゲーム
type LifeGame struct {
	currentField   *Field
	lastField      *Field
	intervalSecond int
}

//NewLifeGame ゲームを新規作成
func NewLifeGame(height int, width int, initRate float64, interval int) *LifeGame {
	lifeGame := new(LifeGame)
	gamefield := CreateFieldFrame(height, width).InitFieldStatus(initRate)

	lifeGame.currentField = gamefield
	lifeGame.intervalSecond = interval
	return lifeGame
}

func (game *LifeGame) nextFrame() (*Field, *Field) {
	nextFrame := CreateFieldFrame(game.currentField.height, game.currentField.width)
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

//MainLoop ゲームのメインのループ
func (game *LifeGame) MainLoop() {
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
