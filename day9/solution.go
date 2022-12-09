package day9

import (
    "math"
    "golang.org/x/exp/slices"
	"strconv"
	"strings"
	"github.com/hpompecki/aoc_2022/input"
)

type Vector struct {
    x, y int
}

var visited = make([][2]int, 0)
var visitedLong = make([][2]int, 0)

func VisitedCount(fileName string) int {
    h := &Vector{}
    t := &Vector{}
    for l := range input.Read(fileName) {
        processInstruction(l, h, t)
    }
    return len(visited) 
}

func VisitedCountLong(fileName string) int {
    knots := [10]*Vector{}
    for i := range knots {
        knots[i] = &Vector{}
    }
    for l := range input.Read(fileName) {
        processInstructionLong(l, knots)
    }
    return len(visitedLong) 
}

func processInstruction(i string, h *Vector, t *Vector) {
    split := strings.Split(i, " ")
    v, _ := strconv.Atoi(split[1])
    for i := 1; i <= v; i++ {
        switch split[0] {
        case "U":
            h.y++
        case "R":
            h.x++
        case "D":
            h.y--
        case "L":
            h.x--
        }
        moveTail(t, h)
        if !slices.Contains(visited, [2]int{t.x, t.y}) {
            visited = append(visited, [2]int{t.x, t.y})
        }
    }
}

func processInstructionLong(i string, k [10]*Vector) {
    split := strings.Split(i, " ")
    v, _ := strconv.Atoi(split[1])
    for i := 1; i <= v; i++ {
        switch split[0] {
        case "U":
            k[0].y++
        case "R":
            k[0].x++
        case "D":
            k[0].y--
        case "L":
            k[0].x--
        }
        for i := range k {
            if i == 0 {
                continue
            }
            moveTail(k[i], k[i-1])
        }
        if !slices.Contains(visitedLong, [2]int{k[9].x, k[9].y}) {
            visitedLong = append(visitedLong, [2]int{k[9].x, k[9].y})
        }
    }
}

func moveTail(t *Vector, h *Vector) {
    if math.Abs(float64(t.x - h.x)) > 1 || math.Abs(float64(t.y - h.y)) > 1 {
        if math.Abs(float64(t.x - h.x)) == 1 {
            t.x = h.x
        }
        if math.Abs(float64(t.y - h.y)) == 1 {
            t.y = h.y
        }
        if math.Abs(float64(t.x - h.x)) == 2 {
            t.x = (h.x + t.x) / 2
        }
        if math.Abs(float64(t.y - h.y)) == 2 {
            t.y = (h.y + t.y) / 2
        }
    }
}
