package main

import (
	"bufio"
	"fmt"
	"github.com/beevik/etree"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	for _, dir := range []string{
		"apps/scalable",
		"devices/scalable",
		"mimetypes/scalable",
		"places/16",
		"places/48",
		"preferences/scalable",
	} {
		for _, file := range readFiles(fmt.Sprintf("candy-icons/%v", dir)) {
			outDir := fmt.Sprintf("out/nord-icons/%v", dir)
			doc := createXmlDoc(file.Second)
			var writer *bufio.Writer
			if err := os.MkdirAll(outDir, os.FileMode(0755)); err != nil {
				log.Fatal(err)
			}
			if outFile, err := os.Create(fmt.Sprintf("%v/%v", outDir, file.First)); err != nil {
				log.Fatal(err)
			} else {
				writer = bufio.NewWriter(outFile)
			}
			if _, err := doc.WriteTo(writer); err != nil {
				log.Fatal(err)
			}
			if err := writer.Flush(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

type Pair[F, S any] struct {
	First  F
	Second S
}

func createXmlDoc(path string) *etree.Document {
	log.Println("Process file:", path)
	doc := etree.NewDocument()

	if err := doc.ReadFromFile(path); err != nil {
		log.Fatal(err)
	}

	root := doc.SelectElement("svg")

	for _, gradient := range root.FindElements("linearGradient") {
		p := randomPalette()
		var colors []string
		for _, s := range gradient.FindElements("stop") {
			var color string
			var style *etree.Attr
			switch {
			case s.SelectAttr("style") != nil:
				style = s.SelectAttr("style")
			case s.SelectAttr("stop-color") != nil:
				style = s.SelectAttr("stop-color")
			}
			for {
				if len(p.colors()) == len(colors) {
					color = p.randomColor()
					colors = []string{}
					colors = append(colors, color)
					break
				} else if contains(colors, color) {
					continue
				} else {
					color = p.randomColor()
					colors = append(colors, color)
					break
				}
			}
			style.Value = fmt.Sprintf("stop-color:%v", color)
		}
	}

	return doc
}

func contains(s []string, ele string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, ele)
	return i < len(s) && s[i] == ele
}

func readFiles(in string) []Pair[string, string] {
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
	var out []Pair[string, string]

	for _, file := range files {
		if extension(file.Name()) == "svg" && file.Type().IsRegular() {
			p := Pair[string, string]{
				First:  file.Name(),
				Second: filepath.Join(path, file.Name()),
			}
			out = append(out, p)
		}
	}

	return out
}

func extension(fileName string) string {
	split := strings.Split(fileName, ".")
	return split[len(split)-1]
}
