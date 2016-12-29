package maze

func (g Grid2d)Distance(root *Cell) map[coordinate]int {
	distances := make(map[coordinate]int, 1)
	distances[root.position] = 0
	frontier := make([]coordinate, 0)
	frontier = append(frontier, root.position)

	for len(frontier) > 0 {
		newFrontier := make([]coordinate, 0)
		for _, cell := range frontier {
			for _, link := range (g.GetCell(cell.x, cell.y).links) {
				_, ok := distances[link]
				if !ok {
					distances[link] = distances[cell] + 1
					frontier = append(frontier, link)
					newFrontier = append(newFrontier, link)
				}
			}
		}
		frontier = newFrontier
	}

	return distances
}

func (g Grid2d)GetFarthestCell(distances map[coordinate]int) *Cell {
	var max_distance int
	var farthest_cell coordinate

	max_distance = -1

	for k, v := range distances {
		if v > max_distance {
			farthest_cell = k
			max_distance = v
		}
	}

	return g.GetCell(farthest_cell.x, farthest_cell.y)
}

func (g Grid2d)Path(goal *Cell) map[coordinate]int {
	current := goal
	distances := make(map[coordinate]int, 1)
	distances[current.position] = current.contents

	for current.contents > 0 {
		for _, pos := range current.links {
			cell := g.GetCell(pos.x, pos.y)
			if cell.contents < current.contents {
				distances[cell.position] = cell.contents
				current = cell
				break
			}
		}
	}
	return distances
}
