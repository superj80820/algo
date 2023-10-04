package main

import (
	"fmt"
	"syscall/js"

	"github.com/superj80820/algo/fun/2048/enum"
	"github.com/superj80820/algo/fun/2048/handler"
)

var keyMap = map[int]interface{}{
	37: enum.LEFT,
	38: enum.UP,
	39: enum.RIGHT,
	40: enum.DOWN,

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

	printDOM(gameHandler.ToHTMLString())

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
			gameHandler.Process(k)
			curPrintStr += gameHandler.ToHTMLString()
			if gameHandler.CheckWin() {
				curPrintStr += "\n\n\u00A0\u00A0\u00A0You Win!"
			}
			if !gameHandler.CheckAvailable() {
				curPrintStr += "\n\n\u00A0\u00A0\u00A0You lose!"
			}
			printDOM(curPrintStr)
		case enum.Number:
			gameHandler.NewGame(int(k))
			curPrintStr += gameHandler.ToHTMLString()
			printDOM(curPrintStr)
		}
		return nil
	})
}

func printDOM(val string) string {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return "document is not defined"
	}
	containerEl := jsDoc.Call("getElementById", "board")
	if !containerEl.Truthy() {
		return "board is not find"
	}
	// p := jsDoc.Call("createElement", "p")
	// p.Set("innerText", val)
	// containerEl.Call("append", p)
	containerEl.Set("innerText", val)
	return "nil"
}
