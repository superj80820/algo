package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/superj80820/algo/fun/2048/enum"
	"github.com/superj80820/algo/fun/2048/handler"
)

func main() {
	ch := make(chan enum.Action)

	go func() {
		for {
			Action := bufio.NewScanner(os.Stdin)
			Action.Scan()
			switch strings.ToLower(Action.Text()) {
			case "w":
				ch <- enum.UP
			case "s":
				ch <- enum.DOWN
			case "a":
				ch <- enum.LEFT
			case "d":
				ch <- enum.RIGHT
			}
		}
	}()

	gameHandler := handler.CreateGameHandler()
	gameHandler.NewGame(4, 4)
	gameHandler.PrintBoard()

	for {
		action := <-ch
		fmt.Println("-------")
		gameHandler.Process(action)
		gameHandler.PrintBoard()
	}

}
