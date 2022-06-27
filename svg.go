package main

type linearGradient struct {
	Stop []stop `xml:"stop"`
}

type stop struct {
	Style string `xml:"style,attr"`
}
