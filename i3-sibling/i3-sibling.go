package main

import (
	"fmt"
	"log"
	"os"

	i3ipc "github.com/brunnre8/i3ipc-go"
	"github.com/davecgh/go-spew/spew"
)

const (
	DirNext = "next"
	DirPrev = "prev"
)

func main() {
	spew.Config.MaxDepth = 2
	spew.Config.Indent = "\n\t"

	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s <prev|next>\n", os.Args[0])
	}
	dir := os.Args[1]

	if !((dir == DirNext) || (dir == DirPrev)) {
		log.Fatal("Expecting \"prev\" or \"next\" as argument")
	}

	ipcsocket, err := i3ipc.GetIPCSocket()

	if err != nil {
		log.Fatal("Error connecting to i3", err)
	}

	tree, err := ipcsocket.GetTree()
	if err != nil {
		log.Fatal("Error fetching tree", err)
	}

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

	idx := 0
	for idx = 0; idx < numSiblings; idx++ {
		if focused.ID == siblings[idx].ID {
			break
		}
	}

	var nextIdx int

	if dir == DirPrev {
		nextIdx = (idx - 1) % numSiblings
	} else {
		nextIdx = (idx + 1) % numSiblings
	}

	targetNode := siblings[nextIdx]

	log.Printf("Focusing node %x\n", targetNode.ID)

	ok, err := ipcsocket.Command(fmt.Sprintf("[con_id=%v] focus", targetNode.ID))
	if !ok {
		log.Fatal("Focus command returned error:  ", err)
	}
}
