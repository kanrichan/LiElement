package tui

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

	Display *BaseValue[string]

	FlexDirection      *BaseValue[string]
	FlexWrap           *BaseValue[string]
	FlexJustifyContent *BaseValue[string]
	AlignItems         *BaseValue[string]
	AlignContent       *BaseValue[string]

	Order      *BaseValue[int]
	FlexBasis  *BaseValue[string]
	FlexGrow   *BaseValue[uint64]
	FlexShrink *BaseValue[uint64]
	AlignSelf  *BaseValue[string]
}

// Event Event
type Event struct {
	OnClick *BaseValue[func()]
}
