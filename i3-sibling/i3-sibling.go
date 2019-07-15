package main

import (
	"fmt"
	"log"
	"os"

	i3ipc "github.com/brunnre8/i3ipc-go"
)

const (
	DirNext = "next"
	DirPrev = "prev"
)

func main() {
	// Argument parsing
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s <prev|next>\n", os.Args[0])
	}
	dir := os.Args[1]

	if !((dir == DirNext) || (dir == DirPrev)) {
		log.Fatal("Expecting \"prev\" or \"next\" as argument")
	}

	// Get i3 socket and parse tree
	ipcsocket, err := i3ipc.GetIPCSocket()

	if err != nil {
		log.Fatal("Error connecting to i3", err)
	}

	tree, err := ipcsocket.GetTree()
	if err != nil {
		log.Fatal("Error fetching tree", err)
	}

	// Find the focused node and its siblings
	focused := tree.FindFocused()
	parent := focused.Parent

	if parent.Type != "con" {
		log.Fatalf("Non-con parent (%s), aborting.\n", parent.Type)
	}

	if len(parent.Nodes) <= 1 {
		log.Fatalf("Parent does not have >1 children, aborting.\n")
	}

	siblings := parent.Nodes
	numSiblings := len(siblings)

	// Find our index into the siblings, then select the next/prev sibling
	idx := 0
	for idx = 0; idx < numSiblings; idx++ {
		if focused.ID == siblings[idx].ID {
			break
		}
	}

	var nextIdx int

	if dir == DirPrev {
		nextIdx = (idx - 1)
		if nextIdx < 0 {
			nextIdx += numSiblings
		}
	} else {
		nextIdx = (idx + 1) % numSiblings
	}

	targetNode := siblings[nextIdx]

	ok, err := ipcsocket.Command(fmt.Sprintf("[con_id=%v] focus", targetNode.ID))
	if !ok {
		log.Fatal("Focus command returned error:  ", err)
	}
}
