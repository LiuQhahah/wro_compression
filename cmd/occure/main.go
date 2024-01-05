package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	rootCmd := cobra.Command{
		Short: "count the number of occurrences of each word in its input",
		Use:   "occure [flags] [file]",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := checkArgs(args); err != nil {
				fmt.Printf("checkArgs() error = %v, wantErr %v", err, true)
			}

			readFile(args[0])
		},
	}
	rootCmd.Execute()
}

func readFile(path string) error {
	fileInfos, err := os.Stat(path)
	if err != nil {
		return errors.New("please input right file path")
	}
	if fileInfos.IsDir() {
		return errors.New("please input right file path instead of dir")
	}
	byteFile, err := os.ReadFile(path)
	if err != nil {
		return errors.New("read file error")
	}
	var charMap = make(map[byte]int)
	for _, v := range byteFile {
		charMap[v]++
	}
	for k, v := range charMap {
		fmt.Printf("%c:%d\n", k, v)
	}
	return nil
}

func checkArgs(args []string) error {
	if len(args) != 1 {
		return errors.New("please input right file path")
	}
	return nil
}
