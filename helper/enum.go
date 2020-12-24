package helper

import "github.com/fatih/color"

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

const leafSeed = "  $@3O⸮2#%!42?+%        "

var (
	leafSeedSize = len(leafSeed)
)