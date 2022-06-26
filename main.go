package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := readFiles("candy-icons/apps/scalable")

	icon := unmarshalSvg(files[0])

	icon.LinearGradient[0].Stop[0].Style = fmt.Sprintf("stop-color:%v", randomColor())

	fmt.Println(icon)
}

func readFiles(in string) []string {
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

	//out := make(map[os.DirEntry]string)
	var out []string

	for _, file := range files {
		if extension(file.Name()) == "svg" && file.Type().IsRegular() {
			out = append(out, filepath.Join(path, file.Name()))
		}
	}

	return out
}

func extension(fileName string) string {
	split := strings.Split(fileName, ".")
	return split[len(split)-1]
}
