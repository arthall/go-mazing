package maze

import (
	"math/rand"
	"time"
)

func Sidewinder(x int, y int) Grid2d {
	rand.Seed(time.Now().UTC().UnixNano())
	g := Grid2d{}
	g.initialize(x, y)

	for i, row := range g.grid {
		run := make([]*Cell, 0)
		for j := range row {
			cell := g.getCell(i, j)
			run = append(run, cell)
			east := g.getCell(cell.east.x, cell.east.y)
			atEasternEdge := (east == nil)
			atNorthernEdge := (g.getCell(cell.north.x, cell.north.y) == nil)

			shouldClose := atEasternEdge || (!atNorthernEdge && rand.Intn(2) == 0)

			if shouldClose {
				member := run[rand.Intn(len(run))]
				north := g.getCell(member.north.x, member.north.y)
				if north != nil {
					member.link(north, true)
				}
				run = nil
			} else {
				cell.link(east, true)
			}
		}
	}
	return g
}
