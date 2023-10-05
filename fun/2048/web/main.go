package main

import (
	"fmt"
	"syscall/js"

	"github.com/superj80820/algo/fun/2048/enum"
	"github.com/superj80820/algo/fun/2048/handler"

	"github.com/pkg/errors"
)

var keyMap = map[int]interface{}{
	38: enum.UP,
	40: enum.DOWN,
	37: enum.LEFT,
	39: enum.RIGHT,

	87: enum.UP,
	83: enum.DOWN,
	65: enum.LEFT,
	68: enum.RIGHT,

	48: enum.ZERO,
	49: enum.ONE,
	50: enum.TWO,
	51: enum.THREE,
	52: enum.FOUR,
	53: enum.FIVE,
	54: enum.SIX,
	55: enum.SEVEN,
	56: enum.EIGHT,
	57: enum.NIGHT,
}

func main() {
	fmt.Println("Hello World from Golang")

	js.Global().Set("sendKey", sendKey())

	gameHandler := handler.GetSingleTonGameHandler()
	gameHandler.NewGame(4)

	var curPrintStr string
	curPrintStr += gameHandler.ScoreToHTMLString() + "\n"
	curPrintStr += gameHandler.BoardToHTMLString()
	if err := printDOM(curPrintStr); err != nil {
		fmt.Println("get error: ", err.Error())
	}

	select {}
}

func sendKey() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		gameHandler := handler.GetSingleTonGameHandler()
		var curPrintStr string
		touchKey := args[0].Int()
		fmt.Println("get key: ", touchKey)
		switch k := keyMap[touchKey].(type) {
		case enum.Action:
			if gameHandler.CheckWin() || !gameHandler.CheckAvailable() {
				return nil
			}
			gameHandler.Process(k)
			curPrintStr += gameHandler.ScoreToHTMLString() + "\n"
			curPrintStr += gameHandler.BoardToHTMLString()
			if gameHandler.CheckWin() {
				curPrintStr += "\n\n\u00A0\u00A0\u00A0You Win!"
			}
			if !gameHandler.CheckAvailable() {
				curPrintStr += "\n\n\u00A0\u00A0\u00A0You lose!"
			}
			if err := printDOM(curPrintStr); err != nil {
				fmt.Println("get error: ", err.Error())
			}
		case enum.Number:
			gameHandler.NewGame(int(k))
			curPrintStr += gameHandler.ScoreToHTMLString() + "\n"
			curPrintStr += gameHandler.BoardToHTMLString()
			if err := printDOM(curPrintStr); err != nil {
				fmt.Println("get error: ", err.Error())
			}
		}
		return nil
	})
}

func printDOM(val string) error {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return errors.New("document is not defined")
	}
	containerEl := jsDoc.Call("getElementById", "board")
	if !containerEl.Truthy() {
		return errors.New("board is not find")
	}
	containerEl.Set("innerText", val)
	return nil
}
