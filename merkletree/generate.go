package merkletree

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func GenerateTree(filename string, shardSize int) (*Branch, error) {

	// find file

	item, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	if item.IsDir() {
		return nil, fmt.Errorf("location must be a file not a directory: %s", filename)
	}

	// setup io

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, shardSize)
	reader := bufio.NewReader(file)

	leaves := []*Branch{}

	// find leaves

	fmt.Printf("Sharding file %s\n", filename)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, nil
		}

		hash := sha256.Sum256(buffer)
		leaf := NewLeaf(hash[:])

		leaves = append(leaves, leaf)
	}
	fmt.Printf("%d shards of size %d bytes found\n", len(leaves), shardSize)

	fmt.Println("Generating Merkle Tree")

	// generate layers
	layers := [][]*Branch{leaves}

	for len(layers[len(layers)-1]) > 1 {
		layer := []*Branch{}
		lastLength := len(layers[len(layers)-1])

		for i := 0; i < lastLength; i += 2 {

			// new branch node
			branch := &Branch{
				Hash:   nil,
				LChild: layers[len(layers)-1][i],
				RChild: nil,
			}

			// if layer is of odd length, duplicate last node
			if i == lastLength-1 {
				branch.RChild = branch.LChild
			} else {
				branch.RChild = layers[len(layers)-1][i+1]
			}

			err = branch.generateHash()
			if err != nil {
				return nil, err
			}

			layer = append(layer, branch)
		}

		layers = append(layers, layer)
	}

	return layers[len(layers)-1][0], nil
}
