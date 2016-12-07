package main

import (
	"github.com/arthall/go-mazing/maze"
)

func main() {
	g := maze.Sidewinder(20, 20)
	g.Display2()
}
