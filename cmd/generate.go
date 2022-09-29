/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t02smith/merkle-tree/merkletree"
)

var (
	filename, serializeDir, serializeTo string
	shardSize                           int
	doSerialize                         bool
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate a new merkle tree",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if filename == "" {
			fmt.Println("No file passed")
			return
		}

		tree, err := merkletree.GenerateTree(filename, shardSize)
		if err != nil {
			fmt.Println(err)
		} else {
			hash, err := tree.GetHash()
			if err != nil {
				fmt.Printf("Error getting root hash: %s\n", err)
			} else {
				fmt.Printf("Root hash: %x\n", hash)
			}

		}

		if doSerialize {
			tree.Serialize(serializeDir, serializeTo)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&filename, "filename", "f", "", "Pass the location of the file to generate a tree of")
	generateCmd.Flags().IntVarP(&shardSize, "shard-size", "s", merkletree.SHARD_SIZE, "The size in bytes of each shard. Smaller shards will take longer but changes to files will be easier to find.")

	generateCmd.Flags().BoolVarP(&doSerialize, "serialize", "z", false, "Store the resulting object in a file")
	generateCmd.Flags().StringVarP(&serializeDir, "serialize-dir", "o", ".", "Where to store the serialized object")
	generateCmd.Flags().StringVarP(&serializeTo, "serialize-type", "t", "json", "Type of file to serialize to [json/gob]")
}
