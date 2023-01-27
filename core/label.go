package core

import (
	"fmt"
	"image/color"
	"strings"
)

type BaseLabel struct {
	Class string
	ID    string

	father   *BaseLabel
	children []*BaseLabel
	order    int

	box box
}

type StylePositon int

const (
	Static StylePositon = iota
	Absolute
	Relative
	Fixed
	Inherit
)

type Style struct {
	Positon StylePositon
	Left    int
	Right   int
	Top     int
	Bottom  int
	Width   int
	Height  int

	Margin       int
	MarginTop    int
	MarginBottom int
	MarginLeft   int
	MarginRight  int

	Overflow int64

	Color           color.Color
	BackGroundColor color.Color
}

type Text struct {
	BaseLabel

	Text  string
	Style Style
}

func (text *Text) Render() {
	if text.Style.Width != 0 {
		text.box.W = text.Style.Width
	} else {
		text.box.W = len(text.Text)
	}
	if text.Style.Margin != 0 {
		text.box.W += text.Style.Margin * 2
	} else {
		text.box.W += text.Style.MarginLeft + text.Style.MarginRight
	}
	if text.Style.Height != 0 {
		text.box.H = text.Style.Height
	} else {
		text.box.H = strings.Count(text.Text, "\n")
	}
	if text.Style.Margin != 0 {
		text.box.H += text.Style.Margin * 2
	} else {
		text.box.H += text.Style.MarginTop + text.Style.MarginBottom
	}
	switch text.Style.Positon {
	case Static:
		text.box.X = 0
		text.box.Y = 0
	case Absolute:
		text.box.X = text.father.box.X
		text.box.Y = text.father.box.Y
	case Relative:
		if text.order == 0 {
			text.box.X = text.father.box.X
			text.box.Y = text.father.box.Y
		} else {
			if text.father.box.W-text.children[text.order-1].box.W-(text.children[text.order-1].box.X-text.father.box.X) > text.box.W {
				text.box.X = text.children[text.order-1].box.X + text.children[text.order-1].box.W
				text.box.Y = text.children[text.order-1].box.Y
			} else {
				text.box.X = text.father.box.X
				text.box.Y = text.children[text.order-1].box.Y + 1
			}
		}
	case Fixed:

	case Inherit:

	}
	if text.box.W > text.father.box.W {
		text.box.W = text.father.box.W
	}
	if text.box.H > text.father.box.H {
		text.box.H = text.father.box.H
	}
}

func init() {
	xxx := &Text{
		Text: "",
		Style: Style{
			Color: color.White,
		},
	}
	fmt.Println(xxx)
}
