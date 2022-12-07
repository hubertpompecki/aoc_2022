package day7

import (
	"strconv"
	"strings"

	"github.com/hpompecki/aoc_2022/input"
)

type Node struct {
	name     string
	size     int
	children []*Node
	parent   *Node
}

func NewNode(n string, p *Node) *Node {
	return &Node{
		name:     n,
		parent:   p,
		children: make([]*Node, 0),
	}
}

func (n *Node) GetChild(name string) *Node {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	return nil
}

func (n *Node) AddChild(c *Node) {
	n.children = append(n.children, c)
}

type FileSystem struct {
	root       *Node
	currentDir *Node
}

func SmallDirsTotal(fileName string) int {
	// build tree
	fs := &FileSystem{}

	for line := range input.Read(fileName) {
		processLine(fs, line)
	}

	populateSizes(fs.root)

	return sumSmallSizes(fs.root)
}

func DirToDelete(fileName string) int {
	// build tree
	fs := &FileSystem{}

	for line := range input.Read(fileName) {
		processLine(fs, line)
	}

	populateSizes(fs.root)

	targetSize := fs.root.size - 40000000

	return dirToDeleteSize(fs.root, targetSize)
}

func sumSmallSizes(n *Node) int {
	var size int
	if n.size <= 100000 && len(n.children) > 0 {
		size += n.size
	}
	for _, c := range n.children {
		size += sumSmallSizes(c)
	}
	return size
}

func dirToDeleteSize(n *Node, targetSize int) int {
	size := 70000000
	if n.size >= targetSize && n.size < size && len(n.children) > 0 {
		size = n.size
	}
	for _, c := range n.children {
		childSize := dirToDeleteSize(c, targetSize)
		if childSize >= targetSize && childSize < size && len(c.children) > 0 {
			size = childSize
		}
	}
	return size
}

func processLine(fs *FileSystem, l string) {
	if l == "$ ls" {
		// noop
		return
	}

	if l == "$ cd .." {
		fs.currentDir = fs.currentDir.parent
		return
	}

	if strings.HasPrefix(l, "$ cd ") {
		name := strings.TrimPrefix(l, "$ cd ")
		if fs.currentDir == nil {
			fs.currentDir = NewNode(name, nil)
			fs.root = fs.currentDir
		} else {
			fs.currentDir = fs.currentDir.GetChild(name)
		}
		return
	}

	if strings.HasPrefix(l, "dir ") {
		name := strings.TrimPrefix(l, "dir ")
		fs.currentDir.AddChild(NewNode(name, fs.currentDir))
		return
	}

	// file
	splits := strings.Split(l, " ")
	size, _ := strconv.Atoi(splits[0])
	name := splits[1]
	fs.currentDir.AddChild(&Node{name: name, size: size})
}

func populateSizes(n *Node) int {
	for _, c := range n.children {
		n.size += populateSizes(c)
	}
	return n.size
}
