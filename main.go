package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tonit/toni/cmd"
)

func main() {
	var colored = color.HiCyanString("Toni CLI")
	fmt.Println(">> " + colored + " <<")
	cmd.Execute()
}
