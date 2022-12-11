package day11

import (
	"sort"
	"strconv"
	"strings"

	"github.com/hpompecki/aoc_2022/input"
)

type Monkey struct {
	items       []int
	operation   func(int) int
	divisor     int
	trueTarget  int
	falseTarget int
	inspects    int
}

func (m *Monkey) Act(reduce bool) {
	for _, i := range m.items {
		i = m.operation(i)
		if reduce {
			i = i / 3
		}
		m.inspects++
		if i%m.divisor == 0 {
			monkeys[m.trueTarget].items = append(monkeys[m.trueTarget].items, i)
		} else {
			//
			monkeys[m.falseTarget].items = append(monkeys[m.falseTarget].items, i%productOfDivisors)
		}
	}
	m.items = make([]int, 0)
}

var monkeys []*Monkey
var productOfDivisors int

func MonkeyBusiness(fileName string) int {
	buildMonkeys(fileName)
	for i := 1; i <= 20; i++ {
		for _, m := range monkeys {
			m.Act(true)
		}
	}
	return businessLevel()
}

func MonkeyBusiness2(fileName string) int {
	buildMonkeys(fileName)
	for i := 1; i <= 10000; i++ {
		for _, m := range monkeys {
			m.Act(false)
		}
	}
	return businessLevel()
}

func buildMonkeys(fileName string) {
	monkeys = make([]*Monkey, 0)
	current := 0
	for l := range input.Read(fileName) {
		l = strings.TrimSpace(l)
		switch {
		case strings.HasPrefix(l, "Monkey"):
			l = strings.TrimSuffix(l, ":")
			current, _ = strconv.Atoi(strings.Split(l, " ")[1])
			monkeys = append(monkeys, &Monkey{})
		case strings.HasPrefix(l, "Starting items: "):
			l = strings.TrimPrefix(l, "Starting items: ")
			for _, s := range strings.Split(l, ", ") {
				n, _ := strconv.Atoi(s)
				monkeys[current].items = append(monkeys[current].items, n)
			}
		case strings.HasPrefix(l, "Operation: new = old * old"):
			monkeys[current].operation = func(i int) int {
				return i * i
			}
		case strings.HasPrefix(l, "Operation: new = old * "):
			l = strings.TrimPrefix(l, "Operation: new = old * ")
			n, _ := strconv.Atoi(l)
			monkeys[current].operation = func(i int) int {
				return i * n
			}
		case strings.HasPrefix(l, "Operation: new = old + "):
			l = strings.TrimPrefix(l, "Operation: new = old + ")
			n, _ := strconv.Atoi(l)
			monkeys[current].operation = func(i int) int {
				return i + n
			}
		case strings.HasPrefix(l, "Test: divisible by "):
			l = strings.TrimPrefix(l, "Test: divisible by ")
			n, _ := strconv.Atoi(l)
			monkeys[current].divisor = n
		case strings.HasPrefix(l, "If true: throw to monkey "):
			l = strings.TrimPrefix(l, "If true: throw to monkey ")
			n, _ := strconv.Atoi(l)
			monkeys[current].trueTarget = n
		case strings.HasPrefix(l, "If false: throw to monkey "):
			l = strings.TrimPrefix(l, "If false: throw to monkey ")
			n, _ := strconv.Atoi(l)
			monkeys[current].falseTarget = n
		default:
			continue
		}
	}

	productOfDivisors = 1
	for _, m := range monkeys {
		productOfDivisors *= m.divisor
	}
}

func businessLevel() int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})
	return monkeys[0].inspects * monkeys[1].inspects
}
