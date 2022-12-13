package day13

import (
    "sort"

	"github.com/hpompecki/aoc_2022/input"
	"golang.org/x/exp/utf8string"
)

type Node struct {
    val int
    children []*Node
    parent *Node
}

func NewNode(parent *Node) *Node {
    return &Node{val: -1, parent: parent, children: make([]*Node, 0)}
}

func DecoderKey(fileName string) int {
    packets := make([]string, 0)
	for l := range input.Read(fileName) {
        if l != "" {
            packets = append(packets, l)
        }
    }
    packets = append(packets, "[[2]]", "[[6]]")

    sort.Slice(packets, func(i, j int) bool {
        return rightOrder(packets[i], packets[j])
    })

    var m, n int
    for i, s := range packets {
        if s == "[[2]]" {
            m = i + 1
        }
        if s == "[[6]]" {
            n = i + 1
        }
    }

    return m * n
}

func RightOrderIndices(fileName string) int {
	var (
		left  string
		right string
		pair  int
		sum   int
	)

	for l := range input.Read(fileName) {
		if left == "" {
			left = l
			continue
		}
		if right == "" {
			right = l
			continue
		}
		pair++
		if rightOrder(left, right) {
			sum += pair
		}
		left = ""
		right = ""
	}
	return sum
}

func rightOrder(left, right string) bool {
    leftRoot := NewNode(nil)
    rightRoot := NewNode(nil)

    buildNodes(leftRoot, left)
    buildNodes(rightRoot, right)

    return checkOrder(leftRoot, rightRoot) == "left"
}

func checkOrder(left, right *Node) string {
    switch {
    case left.val >= 0 && right.val >= 0 && left.val < right.val:
        return "left"
    case left.val >= 0 && right.val >= 0 && left.val > right.val:
        return "right"
    case left.val >= 0 && right.val >= 0 && left.val == right.val:
        return "eq"
    case left.val >= 0 && right.val == -1:
        newNode := NewNode(left)
        newNode.val = left.val
        left.val = -1
        left.children = append(left.children, newNode)
        return checkOrder(left, right)
    case left.val == -1 && right.val >= 0:
        newNode := NewNode(right)
        newNode.val = right.val
        right.val = -1
        right.children = append(right.children, newNode)
        return checkOrder(left, right)
    case left.val == -1 && right.val == -1:
        for i := 0; (i < len(left.children) && i < len(right.children)); i++ {
            if o := checkOrder(left.children[i], right.children[i]); o != "eq" {
                return o
            }
        }
        switch {
        case len(left.children) < len(right.children):
            return "left"
        case len(left.children) > len(right.children):
            return "right"
        default:
            return "eq"
        }
    }
    return "eq"
}

func buildNodes(root *Node, input string) {
    inputU := utf8string.NewString(input)

    current := root
    for i := 1; i < inputU.RuneCount(); i++ {
        switch {
        case inputU.At(i) == '[':
            newNode := NewNode(current)
            current.children = append(current.children, newNode)
            current = newNode
        case inputU.At(i) == ']':
            current = current.parent
        case inputU.At(i) == ',':
        case inputU.At(i) == '0' && inputU.At(i-1) >= '1' && inputU.At(i-1) <= '9':
            current.children[len(current.children) - 1].val *= 10 
        default:
            newNode := NewNode(current)
            newNode.val = int(inputU.At(i) - '0')
            current.children = append(current.children, newNode) 
        }
    }
}
