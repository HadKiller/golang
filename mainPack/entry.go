package main

import (
	"fmt"
	tree "mainPack/myNode"
)

func main() {


	tree.NLR(tree.CreateTree())
	fmt.Print("\n")
	tree.LNR(tree.CreateTree())
	fmt.Print("\n")
	tree.LRN(tree.CreateTree())
}
