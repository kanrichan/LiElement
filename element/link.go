package element

import "vue-tui/tui"

// Link Link
type Link struct {
	Element   *tui.BaseValue[tui.Element]
	Type      *tui.BaseValue[string]
	Underline *tui.BaseValue[bool]
	Disabled  *tui.BaseValue[bool]
	Href      *tui.BaseValue[string]
	Icon      *tui.BaseValue[string]
}

// Render Render
func (e *Link) Render() *tui.BaseValue[tui.Element] {
	return e.Element
}
