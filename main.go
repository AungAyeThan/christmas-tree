package main

import (
	"flag"
	"time"
	christmas "toturials/helper"
)

func main() {
	var team string
	var cond bool
	flag.StringVar(&team, "Team Unicorn ", "Aung Aye Than", "")
	flag.BoolVar(&cond, "live", false, "Keep printing the tree.")

	flag.Parse()

	masTree := christmas.NewTree(40, team)

	for true {
		masTree.Print()
		if !cond {
			return
		}
		time.Sleep(2 * time.Second)
	}
}