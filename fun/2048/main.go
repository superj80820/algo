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
				fmt.Println("UP!")
				ch <- enum.UP
			case "s":
				fmt.Println("DOWN!")
				ch <- enum.DOWN
			case "a":
				fmt.Println("LEFT!")
				ch <- enum.LEFT
			case "d":
				fmt.Println("RIGHT!")
				ch <- enum.RIGHT
			}
		}
	}()

	gameHandler := handler.CreateGameHandler()
	gameHandler.NewGame(30, 30)
	gameHandler.PrintBoard()

	for {
		action := <-ch
		fmt.Println("-------")
		gameHandler.Process(action)
		gameHandler.PrintBoard()
	}

}
