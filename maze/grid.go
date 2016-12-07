package maze

import (
	"fmt"
	"strings"
	"math/rand"
)

type Maze interface{
	initialize()
	size()
	randomCell()
}

type Grid2d struct {
	rows, cols int
	grid [][]Cell
}

func (g *Grid2d) prepare() [][]Cell {
	t := make([][]Cell, g.rows)
	for i := 0; i < g.rows; i++ {
		t[i] = make([]Cell, g.cols)
		for j := 0; j < g.cols; j++ {
			t[i][j] = Cell{position: coordinate{i, j}}
		}
	}
	return t
}

func (g *Grid2d) configure() {
	for i, row := range(g.grid) {
		for j := range(row) {
			g.grid[i][j].north = coordinate{i -1, j}
			g.grid[i][j].east = coordinate{i, j + 1}
			g.grid[i][j].south = coordinate{i + 1, j}
			g.grid[i][j].west = coordinate{i, j - 1}
		}
	}
}

func (g *Grid2d) initialize(rows int, cols int) {
	g.rows = rows
	g.cols = cols
	g.grid = g.prepare()
	g.configure()
}

func (g *Grid2d) size() int {
	return g.rows * g.cols
}

func (g *Grid2d) randomCell() *Cell {
	x := rand.Intn(g.rows)
	y := rand.Intn(g.cols)
	return &g.grid[x][y]
}

func (g *Grid2d) getCell(x int, y int) *Cell {
	if x < 0 || x > g.rows - 1 {
		return nil
	}
	if y < 0 || y > g.cols - 1 {
		return nil
	}
	return &g.grid[x][y]
}

func (g *Grid2d) exists(pos coordinate) bool {
	if g.getCell(pos.x, pos.y) ==  nil {
		return false
	}
	return true
}

func (g *Grid2d) display() {
	fmt.Println("+" + strings.Repeat("---+", g.cols))

	empty := "   "
	for _, row := range(g.grid) {
		middle := make([]byte, 0)
		bottom := make([]byte, 0)

		middle = append(middle, '|')
		bottom = append(bottom, '+')
		for _, cell := range(row) {
			middle = append(middle, []byte(empty)...)
			if cell.isLinked(cell.east) {
				middle =append(middle, ' ')
			} else {
				middle = append(middle, '|')
			}

			if cell.isLinked(cell.south) {
				bottom = append(bottom, []byte(empty)...)
			} else {
				bottom = append(bottom, []byte("---")...)
			}
			bottom = append(bottom, '+')
		}
		fmt.Println(string(middle))
		fmt.Println(string(bottom))
	}
}


func (g *Grid2d) calculateIntersection(c Cell) int {
	result := 0
	if !c.isLinked(c.east) {
		result += 1
	}
	if !c.isLinked(c.south) {
		result += 2
	}
	if g.exists(c.east) {
		east := g.getCell(c.east.x, c.east.y)
		if !east.isLinked(east.south) {
			result += 4
		}
	}
	if g.exists(c.south) {
		south := g.getCell(c.south.x, c.south.y)
		if !south.isLinked(south.east) {
			result += 8
		}
	}
	return result
}

func (g *Grid2d) Display2() {
	intersections := []rune(" ╵╴┘╶└─┴╷│┐┤┌├┬┼")
	vWall := intersections[9]
	hWall := strings.Repeat(string(intersections[6]), 3)
	empty := strings.Repeat(string(intersections[0]), 3)

	// Top of maze
	top := make([]rune, 0)
	top = append(top, intersections[12])
	for _, cell := range(g.grid[0]) {
		top = append(top, []rune(hWall)...)
		if g.getCell(cell.east.x, cell.east.y) != nil {
			if cell.isLinked(cell.east) {
				top = append(top, intersections[6])
			} else {
				top = append(top, intersections[14])
			}
		}
	}
	top = append(top, intersections[10])
	fmt.Println(string(top))

	// Rest of maze
	for r, row := range(g.grid) {
		middle := make([]rune, 0)
		bottom := make([]rune, 0)

		middle = append(middle, vWall)
		if r == g.rows - 1 {
			bottom = append(bottom, intersections[5])
		} else {
			if g.getCell(r, 0).isLinked(g.getCell(r, 0).south){
				bottom = append(bottom, vWall)
			} else {
				bottom = append(bottom, intersections[13])
			}

		}
		for _, cell := range(row) {
			// Middle of the row
			middle = append(middle, []rune(empty)...)
			if cell.isLinked(cell.east) {
				middle =append(middle, ' ')
			} else {
				middle = append(middle, vWall)
			}

			// bottom edge of row
			if cell.isLinked(cell.south) {
				bottom = append(bottom, []rune(empty)...)
			} else {
				bottom = append(bottom, []rune(hWall)...)
			}
			bottom = append(bottom, intersections[g.calculateIntersection(cell)])
		}
		fmt.Println(string(middle))
		fmt.Println(string(bottom))
	}
}