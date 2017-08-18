package main

import (
	"github.com/arthall/go-mazing/maze"
)

func main() {
	g := maze.RandomWalk(40, 40)
	//g.DisplayUnicode()
	g.DisplayImage(true, true)
}
