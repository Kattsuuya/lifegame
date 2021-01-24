package components

import (
	"fmt"
	"math/rand"
	"time"
)

//Field ゲームフィールド
type Field struct {
	height int
	width  int
	status [][]int
}

//CreateFieldFrame フレームを作成
func CreateFieldFrame(height int, width int) *Field {
	field := new(Field)
	field.height = height
	field.width = width
	field.status = make([][]int, height)
	for h := 0; h < height; h++ {
		field.status[h] = make([]int, width)
	}
	return field
}

//InitFieldStatus フィールドを乱数で初期化
func (field *Field) InitFieldStatus(initRate float64) *Field {
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
