package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := readFiles("candy-icons/apps/scalable")

	fmt.Println(unmarshalSvg(files[0]))
}

func unmarshalSvg(filePath string) svg {
	var icon svg
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(data, &icon)
	if err != nil {
		log.Fatal(err)
	}
	return icon
}

type rawXML struct {
	Inner []byte `xml:",innerxml"`
}

type svg struct {
	XMLName        xml.Name         `xml:"svg"`
	LinearGradient []linearGradient `xml:"linearGradient"`
	Path           []rawXML         `xml:"path"`
}

type linearGradient struct {
	Stop []stop `xml:"stop"`
}

type stop struct {
	Offset      string `xml:"offset,attr"`
	StopOpacity string `xml:"stop-opacity,attr"`
	Style       string `xml:"style,attr"`
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
