package main

import (
	"fmt"
)

// Element Element
type Element struct {
	Style *BaseValue[Style]
	Event *BaseValue[Event]

	Text *BaseValue[string]
	Item *BaseValue[[]*BaseValue[Element]]
}

// Style Style
type Style struct {
	Width  *BaseValue[string]
	Height *BaseValue[string]

	FlexDirection      *BaseValue[string]
	FlexWrap           *BaseValue[string]
	FlexJustifyContent *BaseValue[string]
	AlignItems         *BaseValue[string]
	AlignContent       *BaseValue[string]

	Order      *BaseValue[string]
	FlexBasis  *BaseValue[string]
	FlexGrow   *BaseValue[string]
	FlexShrink *BaseValue[string]
	AlignSelf  *BaseValue[string]
}

// Event Event
type Event struct {
	OnClick *BaseValue[func()]
}

// Button Button
type Button struct {
	Element *BaseValue[Element]
	Type    *BaseValue[string]
	Icon    *BaseValue[string]
	Plain   *BaseValue[bool]
	Round   *BaseValue[bool]
	Circle  *BaseValue[bool]
}

// Render Render
func (e Button) Render() *BaseValue[Element] {
	return e.Element
}

func init() {
	fmt.Println(&Element{
		Style: Value(Style{
			Width:    Value("100%"),
			FlexWrap: Value("no-warp"),
		}),
		Text: Value("hello world!"),
		Item: Value([]*BaseValue[Element]{
			Value(Element{
				Text: Value("ok"),
			}),
			Value(Element{
				Text: Value("!!"),
			}),
			Button{
				Element: Value(Element{
					Text: Value("按钮"),
					Event: Value(Event{
						OnClick: Value(func() {
							fmt.Println("!?")
						}),
					}),
				}),
			}.Render(),
		}),
	})
}
