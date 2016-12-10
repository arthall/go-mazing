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
