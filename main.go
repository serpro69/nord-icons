package main

import (
	"bufio"
	"fmt"
	"github.com/beevik/etree"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := readFiles("candy-icons/apps/scalable")

	doc := createXmlDoc(files[0])

	outFile, _ := os.Create("out/test.svg")
	writer := bufio.NewWriter(outFile)
	if _, err := doc.WriteTo(writer); err != nil {
		log.Fatal(err)
	}
	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func createXmlDoc(path string) *etree.Document {
	doc := etree.NewDocument()

	if err := doc.ReadFromFile(path); err != nil {
		log.Fatal(err)
	}

	root := doc.SelectElement("svg")

	for _, gradient := range root.FindElements("linearGradient") {
		for _, s := range gradient.FindElements("stop") {
			style := s.SelectAttr("style")
			style.Value = fmt.Sprintf("stop-color:%v", randomColor())
		}
	}

	return doc
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
