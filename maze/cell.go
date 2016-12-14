package maze


type coordinate struct {
	x, y int
}

// var numbers = []rune("⓪①②③④⑤⑥⑦⑧⑨⑩⑪⑫⑬⑭⑮⑯⑰⑱⑲⑳㉑㉒㉓㉔㉕㉖㉗㉘㉙㉚㉛㉜㉝㉞㉟㊱㊲㊳㊴")
var numbers = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Cell struct {
	position coordinate
	north    coordinate
	south    coordinate
	east     coordinate
	west     coordinate
	contents int
	links    []coordinate
}

func (c Cell) isLinked(neighbor coordinate) bool {
	for _, cell := range (c.links) {
		if cell.x == neighbor.x && cell.y == neighbor.y {
			return true
		}
	}
	return false
}

func (c *Cell) link(neighbor *Cell, bidirectional bool) {
	c.links = append(c.links, neighbor.position)
	if bidirectional {
		neighbor.link(c, false)
	}
}

func (c *Cell) unlink(neighbor *Cell, bidirection bool) {
	for i, pos := range (c.links) {
		if pos.x == neighbor.position.x && pos.y == neighbor.position.y {
			c.links = append(c.links[:i], c.links[i + 1:]...)
		}
	}

	if bidirection {
		neighbor.unlink(c, false)
	}
}