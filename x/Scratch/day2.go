package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	os.Open("README.md")
	color.BgHiYellow
}
