package day12

import (
	"github.com/hpompecki/aoc_2022/input"
	"golang.org/x/exp/slices"
)

type Point struct {
	x, y int
}

var (
	hmap   [][]rune
	start  Point
	target Point
)

func ShortestPathFromBestStart(fileName string) int {
	buildMap(fileName)

	currentPoints := []Point{target}
	visitedPoints := []Point{target}
	steps := 0

	for !reachedLowest(currentPoints) {
		newCurrentPoints := make([]Point, 0)
		for _, p := range currentPoints {
			if p.y+1 < len(hmap) && hmap[p.y+1][p.x] >= hmap[p.y][p.x]-1 && !slices.Contains(visitedPoints, Point{x: p.x, y: p.y + 1}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x, y: p.y + 1})
				visitedPoints = append(visitedPoints, Point{x: p.x, y: p.y + 1})
			}
			if p.x+1 < len(hmap[p.y]) && hmap[p.y][p.x+1] >= hmap[p.y][p.x]-1 && !slices.Contains(visitedPoints, Point{x: p.x + 1, y: p.y}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x + 1, y: p.y})
				visitedPoints = append(visitedPoints, Point{x: p.x + 1, y: p.y})
			}
			if p.y-1 >= 0 && hmap[p.y-1][p.x] >= hmap[p.y][p.x]-1 && !slices.Contains(visitedPoints, Point{x: p.x, y: p.y - 1}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x, y: p.y - 1})
				visitedPoints = append(visitedPoints, Point{x: p.x, y: p.y - 1})
			}
			if p.x-1 >= 0 && hmap[p.y][p.x-1] >= hmap[p.y][p.x]-1 && !slices.Contains(visitedPoints, Point{x: p.x - 1, y: p.y}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x - 1, y: p.y})
				visitedPoints = append(visitedPoints, Point{x: p.x - 1, y: p.y})
			}
		}
		currentPoints = newCurrentPoints
		steps++
	}

	return steps
}

func reachedLowest(currentPoints []Point) bool {
	for _, p := range currentPoints {
		if hmap[p.y][p.x] == 'a' {
			return true
		}
	}
	return false
}

func ShortestPath(fileName string) int {
	buildMap(fileName)

	currentPoints := []Point{start}
	visitedPoints := []Point{start}
	steps := 0

	for !slices.Contains(currentPoints, target) {
		newCurrentPoints := make([]Point, 0)
		for _, p := range currentPoints {
			if p.y+1 < len(hmap) && hmap[p.y+1][p.x] <= hmap[p.y][p.x]+1 && !slices.Contains(visitedPoints, Point{x: p.x, y: p.y + 1}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x, y: p.y + 1})
				visitedPoints = append(visitedPoints, Point{x: p.x, y: p.y + 1})
			}
			if p.x+1 < len(hmap[p.y]) && hmap[p.y][p.x+1] <= hmap[p.y][p.x]+1 && !slices.Contains(visitedPoints, Point{x: p.x + 1, y: p.y}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x + 1, y: p.y})
				visitedPoints = append(visitedPoints, Point{x: p.x + 1, y: p.y})
			}
			if p.y-1 >= 0 && hmap[p.y-1][p.x] <= hmap[p.y][p.x]+1 && !slices.Contains(visitedPoints, Point{x: p.x, y: p.y - 1}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x, y: p.y - 1})
				visitedPoints = append(visitedPoints, Point{x: p.x, y: p.y - 1})
			}
			if p.x-1 >= 0 && hmap[p.y][p.x-1] <= hmap[p.y][p.x]+1 && !slices.Contains(visitedPoints, Point{x: p.x - 1, y: p.y}) {
				newCurrentPoints = append(newCurrentPoints, Point{x: p.x - 1, y: p.y})
				visitedPoints = append(visitedPoints, Point{x: p.x - 1, y: p.y})
			}
		}
		currentPoints = newCurrentPoints
		steps++
	}

	return steps
}

func buildMap(fileName string) {
	hmap = make([][]rune, 0)
	for l := range input.Read(fileName) {
		row := make([]rune, 0)
		for i, char := range l {
			if char == 'S' {
				start = Point{x: i, y: len(hmap)}
				char = 'a'
			}
			if char == 'E' {
				target = Point{x: i, y: len(hmap)}
				char = 'z'
			}
			row = append(row, char)
		}
		hmap = append(hmap, row)
	}
}
