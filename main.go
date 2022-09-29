package main

import (
	"fmt"
	"os"

	"github.com/t02smith/merkle-tree/merkletree"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no file passed as argument")
		return
	}

	root, err := merkletree.GenerateTree(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	rootHash, err := (*root).Hash()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Root hash: %x\n", rootHash)
}
