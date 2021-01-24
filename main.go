package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Hotsukai/lifegame/components"
)

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
	lifegame := components.NewLifeGame(height, width, initRate, interval)
	lifegame.MainLoop()
}
