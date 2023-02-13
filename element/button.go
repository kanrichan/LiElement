package element

import (
	"vue-tui/tui"
)

// Button Button
type Button struct {
	Element *tui.BaseValue[tui.Element]
	Type    *tui.BaseValue[string]
	Icon    *tui.BaseValue[string]
	Plain   *tui.BaseValue[bool]
	Round   *tui.BaseValue[bool]
	Circle  *tui.BaseValue[bool]
}

// Render Render
func (e Button) Render() *tui.BaseValue[tui.Element] {
	return e.Element
}
