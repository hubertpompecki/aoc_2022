package day8

import (
	"strconv"

	"github.com/hpompecki/aoc_2022/input"
)

func CountVisibleTrees(fileName string) int {
	grid := populateGrid(fileName)

	var c int
	for x := range grid {
		for y := range grid[x] {
			if visible(grid, x, y) {
				c++
			}
		}
	}

	return c
}

func HighestScenicScore(fileName string) int {
	grid := populateGrid(fileName)

	var s int
	for x := range grid {
		for y := range grid[x] {
			if score := scenicScore(grid, x, y); score > s {
				s = score
			}
		}
	}

	return s
}

func populateGrid(fileName string) [][]int {
	grid := make([][]int, 0)
	for l := range input.Read(fileName) {
		grid = append(grid, make([]int, 0))
		for _, r := range l {
			n, _ := strconv.Atoi(string(r))
			grid[len(grid)-1] = append(grid[len(grid)-1], n)
		}
	}
	return grid
}

func scenicScore(grid [][]int, x, y int) int {
	return scoreUp(grid, x, y) * scoreRight(grid, x, y) * scoreDown(grid, x, y) * scoreLeft(grid, x, y)
}

func scoreUp(grid [][]int, x, y int) int {
	var s int
	h := grid[x][y]
	for x2 := x - 1; x2 >= 0; x2-- {
		s++
		if grid[x2][y] >= h {
			break
		}
	}
	return s
}

func scoreRight(grid [][]int, x, y int) int {
	var s int
	h := grid[x][y]
	for y2 := y + 1; y2 < len(grid[0]); y2++ {
		s++
		if grid[x][y2] >= h {
			break
		}
	}
	return s
}

func scoreDown(grid [][]int, x, y int) int {
	var s int
	h := grid[x][y]
	for x2 := x + 1; x2 < len(grid); x2++ {
		s++
		if grid[x2][y] >= h {
			break
		}
	}
	return s
}

func scoreLeft(grid [][]int, x, y int) int {
	var s int
	h := grid[x][y]
	for y2 := y - 1; y2 >= 0; y2-- {
		s++
		if grid[x][y2] >= h {
			break
		}
	}
	return s
}

func visible(grid [][]int, x, y int) bool {
	return allLowerUp(grid, x, y) || allLowerRight(grid, x, y) || allLowerDown(grid, x, y) || allLowerLeft(grid, x, y)
}

func allLowerUp(grid [][]int, x, y int) bool {
	h := grid[x][y]
	for x2 := x - 1; x2 >= 0; x2-- {
		if grid[x2][y] >= h {
			return false
		}
	}
	return true
}

func allLowerRight(grid [][]int, x, y int) bool {
	h := grid[x][y]
	for y2 := y + 1; y2 < len(grid[0]); y2++ {
		if grid[x][y2] >= h {
			return false
		}
	}
	return true
}

func allLowerDown(grid [][]int, x, y int) bool {
	h := grid[x][y]
	for x2 := x + 1; x2 < len(grid); x2++ {
		if grid[x2][y] >= h {
			return false
		}
	}
	return true
}

func allLowerLeft(grid [][]int, x, y int) bool {
	h := grid[x][y]
	for y2 := y - 1; y2 >= 0; y2-- {
		if grid[x][y2] >= h {
			return false
		}
	}
	return true
}
