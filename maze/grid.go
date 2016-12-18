package maze

import (
	"fmt"
	"strings"
	"math/rand"
	"image"
	"image/png"
	"image/draw"
	"os"
	"image/color"
	"github.com/dustin/go-heatmap/schemes"
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

func (g *Grid2d) GetCell(x int, y int) *Cell {
	if x < 0 || x > g.rows - 1 {
		return nil
	}
	if y < 0 || y > g.cols - 1 {
		return nil
	}
	return &g.grid[x][y]
}

func (g *Grid2d) exists(pos coordinate) bool {
	if g.GetCell(pos.x, pos.y) ==  nil {
		return false
	}
	return true
}

func (g *Grid2d) displayAscii() {
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
		east := g.GetCell(c.east.x, c.east.y)
		if !east.isLinked(east.south) {
			result += 4
		}
	}
	if g.exists(c.south) {
		south := g.GetCell(c.south.x, c.south.y)
		if !south.isLinked(south.east) {
			result += 8
		}
	}
	return result
}


func (g *Grid2d) calculateTile(c Cell) int {
	result := 0

	if c.isLinked(c.north) {
		result += 1
	}
	if c.isLinked(c.east) {
		result += 2
	}
	if c.isLinked(c.south) {
		result += 4
	}
	if c.isLinked(c.west) {
		result += 8
	}

	return result
}

func (g *Grid2d) DisplayUnicode() {
	intersections := []rune(" ╵╴┘╶└─┴╷│┐┤┌├┬┼")
	vWall := intersections[9]
	hWall := strings.Repeat(string(intersections[6]), 3)
	empty := strings.Repeat(string(intersections[0]), 3)

	// Top of maze
	top := make([]rune, 0)
	top = append(top, intersections[12])
	for _, cell := range(g.grid[0]) {
		top = append(top, []rune(hWall)...)
		if g.GetCell(cell.east.x, cell.east.y) != nil {
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
			if g.GetCell(r, 0).isLinked(g.GetCell(r, 0).south){
				bottom = append(bottom, vWall)
			} else {
				bottom = append(bottom, intersections[13])
			}
		}

		for _, cell := range(row) {
			// Middle of the row
			middle = append(middle, intersections[0])
			if cell.contents != 0 {
				middle = append(middle, numbers[cell.contents])
			} else {
				middle = append(middle, intersections[0])
			}
			middle = append(middle, intersections[0])
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

func(g *Grid2d) DisplayImage(withHeatmap bool) {
	cellSize := 30
	myimage := image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{g.rows * cellSize, g.cols * cellSize}})
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	file, err := os.Open("mazetiles.png")
	if err != nil {
		fmt.Println("Could open tiles.")
	}
	tiles, err := png.Decode(file)
	if err != nil {
		fmt.Println("Could decode tiles.")
	}
	defer file.Close()

	// This loop just fills the image with random data
	for x, row := range(g.grid) {
		for y, cell := range(row) {
			rect := image.Rect(y * cellSize, x * cellSize, y * cellSize +  cellSize, x * cellSize + cellSize)
			if withHeatmap {
				draw.Draw(myimage, rect, &image.Uniform{schemes.Classic[cell.contents]}, image.ZP, draw.Src)
			}
			tileNumber := g.calculateTile(cell)
			point := image.Point{tileNumber * 30, 0}
			draw.Draw(myimage, rect, tiles, point, draw.Over)
		}
	}

	myfile, _ := os.Create("test.png")
	defer myfile.Close()
	png.Encode(myfile, myimage)
}

//func ShowImage(m image.Image) {
//	var buf bytes.Buffer
//	err := png.Encode(&buf, m)
//	if err != nil {
//		panic(err)
//	}
//	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
//	fmt.Println("IMAGE:" + enc)
//}

func (g *Grid2d) AddDistances(dis map[coordinate]int) {
	for pos, val := range dis {
		g.GetCell(pos.x, pos.y).contents = val
	}
}