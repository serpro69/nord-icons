package main

import (
	"encoding/xml"
	"log"
	"os"
)

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
