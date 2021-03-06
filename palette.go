package main

import (
	"math/rand"
	"strings"
	"time"
)

type palette string

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// https://www.nordtheme.com/docs/colors-and-palettes
const (
	PolarNight palette = "#2E3440 #3B4252 #434C5E #4C566A"
	SnowStorm  palette = "#D8DEE9 #E5E9F0 #ECEFF4"
	Frost      palette = "#8FBCBB #88C0D0 #81A1C1 #5E81AC"
	Aurora     palette = "#BF616A #D08770 #EBCB8B #A3BE8C #B48EAD"
)

func (p palette) colors() []string {
	return strings.Split(string(p), " ")
}

func (p palette) randomColor() string {
	colors := p.colors()
	i := r.Intn(len(colors))
	return colors[i]
}

func randomPalette() palette {
	all := []palette{PolarNight, SnowStorm, Frost, Aurora}
	i := r.Intn(len(all))
	return all[i]
}

func allColors() []string {
	var all []string
	all = PolarNight.colors()
	all = append(all, SnowStorm.colors()...)
	all = append(all, Frost.colors()...)
	return append(all, Aurora.colors()...)
}

func randomColor() string {
	colors := allColors()
	i := r.Intn(len(colors))
	return colors[i]
}
