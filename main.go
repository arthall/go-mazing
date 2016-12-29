package main

import (
	"github.com/arthall/go-mazing/maze"
)

func main() {
	g := maze.RandomWalk(40, 40)
	root := g.GetCell(0, 0)
	dis := g.Distance(root)
	farthest := g.GetFarthestCell(dis)
	root = farthest
	dis = g.Distance(root)
	farthest = g.GetFarthestCell(dis)

	g.AddDistances(dis)
	dis = g.Path(farthest)
	g.ClearDistances()
	//g.DisplayUnicode()
	g.AddDistances(dis)
	//g.DisplayUnicode()
	g.DisplayImage(false, true)
}
