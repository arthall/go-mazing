package main

import (
	"github.com/arthall/go-mazing/maze"
)

func main() {
	g := maze.Sidewinder(40, 40)
	root := g.GetCell(0, 0)
	dis := g.Distance(root)
	g.AddDistances(dis)
	//g.DisplayUnicode()
	g.DisplayImage(false)
}
