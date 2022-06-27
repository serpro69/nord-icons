package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := readFiles("candy-icons/apps/scalable")

	file, err := os.Open(files[0])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var buf bytes.Buffer

	decoder := xml.NewDecoder(file)
	encoder := xml.NewEncoder(&buf)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error getting token: %v\n", err)
			break
		}

		switch v := token.(type) {
		case xml.StartElement:
			if v.Name.Local == "linearGradient" {
				var gradient linearGradient
				if err = decoder.DecodeElement(&gradient, &v); err != nil {
					log.Fatal(err)
				}
				for i := range gradient.Stop {
					gradient.Stop[i].Style = fmt.Sprintf("stop-color:%v", randomColor())
				}
				if err = encoder.EncodeElement(gradient, v); err != nil {
					log.Fatal(err)
				}
				continue
			}
		}

		if err := encoder.EncodeToken(xml.CopyToken(token)); err != nil {
			log.Fatal(err)
		}
	}

	if err := encoder.Flush(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
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
