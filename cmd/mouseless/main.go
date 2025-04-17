package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	sx, sy := robotgo.GetScreenSize()
	rect := robotgo.GetScreenRect()
	fmt.Println("display id", robotgo.DisplayID)
	fmt.Println("Hello world", sx, sy)
	fmt.Println("rect", rect)
}
