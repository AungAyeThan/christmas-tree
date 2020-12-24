package helper

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)


// Tree definition
type Tree struct {
	treeName string
	leaves  [][]string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewTree(width int, treeName string) *Tree {
	height := width / 2

	leaves := make([][]string, height)

	for i := 0; i < height; i++ {
		leaves[i] = newLevelLeaves(width, " ")
		if i == 0 {
			leaves[i][width/2] = "★"
			continue
		}

		leaves[i][height-i] = "/"
		leaves[i][height+i] = "\\"
		for j := height - i + 1; j < height+i; j++ {
			leaves[i][j] = leafContent()
		}
	}

	leaves = append(leaves, bottomLeaves(width, "^"), bottomLeaves(width, " "))

	return &Tree{
		leaves:  leaves,
		treeName: treeName,
	}
}

// Print : prints the whole tree at once
func (tree *Tree) Print() {
	fmt.Println()

	for _, leaves := range tree.leaves {
		for _, leaf := range leaves {
			switch leaf {
			case "★":
				yellowColor.Print(leaf)
			default:
				randomColor().Print(leaf)
			}
		}

		fmt.Println()
	}

	fmt.Println()


	redColor.Print(" Merry Christmas! ")
	yellowColor.Print("everybody")
	redColor.Print(" stay safe")
	yellowColor.Print(" and wash your hands")

	fmt.Println()
	fmt.Println()
}


func leafContent() string {
	return string(leafSeed[rand.Intn(leafSeedSize)])
}

func newLevelLeaves(size int, content string) []string {
	target := make([]string, size)

	for i := 0; i < size; i++ {
		target[i] = content
	}

	return target
}

func bottomLeaves(size int, content string) []string {
	target := newLevelLeaves(size, content)
	target[0] = " "

	center := size / 2
	target[center] = " "
	target[center+1] = " "
	target[center-1] = "|"
	target[center+2] = "|"

	return target
}

func randomColor() *color.Color {
	var (
		yellowColor = color.New(color.FgYellow)
		redColor    = color.New(color.FgRed)

		colors = []*color.Color{
			color.New(color.FgGreen),
			yellowColor,
			redColor,
			color.New(color.FgBlue),
			color.New(color.FgCyan),
			color.New(color.FgMagenta),
		}

		colorSize = len(colors)
	)

	return colors[rand.Intn(colorSize)]
}
