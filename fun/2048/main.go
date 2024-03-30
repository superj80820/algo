package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/superj80820/algo/fun/2048/enum"
	"github.com/superj80820/algo/fun/2048/handler"
)

func main() {
	setGameCh := make(chan string)
	isSetDoneCh := make(chan bool)
	actionCh := make(chan enum.Action)

	go func() {
	LOOP:
		for {
			select {
			case isSetDone := <-isSetDoneCh:
				if isSetDone {
					break LOOP
				}
				fmt.Println("board size: ")
				Action := bufio.NewScanner(os.Stdin)
				Action.Scan()
				setGameCh <- Action.Text()
			}
		}
		for {
			Action := bufio.NewScanner(os.Stdin)
			Action.Scan()
			switch strings.ToLower(Action.Text()) {
			case "w":
				fmt.Println("UP!")
				actionCh <- enum.UP
			case "s":
				fmt.Println("DOWN!")
				actionCh <- enum.DOWN
			case "a":
				fmt.Println("LEFT!")
				actionCh <- enum.LEFT
			case "d":
				fmt.Println("RIGHT!")
				actionCh <- enum.RIGHT
			}
		}
	}()

	var boardSize int
	for {
		isSetDoneCh <- false
		var err error
		boardSize, err = strconv.Atoi(<-setGameCh)
		if err != nil {
			fmt.Println("Set fail, please retry")
			continue
		}
		if boardSize <= 1 {
			fmt.Println("Set fail, please greater than 1")
			continue
		}
		break
	}
	isSetDoneCh <- true
	gameHandler := handler.CreateGameHandler()
	gameHandler.NewGame(boardSize)
	gameHandler.PrintBoard()

	for {
		action := <-actionCh
		fmt.Println("-------")
		gameHandler.Process(action)
		gameHandler.PrintBoard()
		if gameHandler.CheckWin() {
			fmt.Println("You Win!")
			os.Exit(0)
		}
		if !gameHandler.CheckAvailable() {
			fmt.Println("You lose!")
			os.Exit(0)
		}
	}

}
