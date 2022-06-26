package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := readFiles("candy-icons/apps/scalable")

	for _, v := range files {
		fmt.Println(v)
	}
}

func readFiles(in string) map[os.DirEntry]string {
	f, err := os.Open(in)
	if err != nil {
		fmt.Println(err)
		panic("Error while opening input dir")
	}

	files, err := f.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		panic("Error while reading input dir")
	}

	path, err := filepath.Abs(in)
	if err != nil {
		fmt.Println(err)
		panic("Error getting absolute path")
	}

	out := make(map[os.DirEntry]string)

	for _, file := range files {
		if extension(file.Name()) == "svg" && file.Type().IsRegular() {
			out[file] = filepath.Join(path, file.Name())
		}
	}

	return out
}

func extension(fileName string) string {
	split := strings.Split(fileName, ".")
	return split[len(split)-1]
}
