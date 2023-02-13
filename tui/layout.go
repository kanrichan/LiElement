package tui

import "sort"

type container struct {
	x uint64
	y uint64
	w uint64
	h uint64

	element *BaseValue[Element]

	father   *container
	children []*container
}

func (c *container) screen(element *BaseValue[Element]) {
	c.element = element
}

func (c *container) allocate() {
	if c.element.Get().Item == nil || len(c.element.Get().Item.Get()) == 0 {
		return
	}
	c.children = nil
	for i := range c.element.Get().Item.Get() {
		cc := &container{
			x:       0,
			y:       0,
			w:       0,
			h:       0,
			element: c.element.Get().Item.Get()[i],
			father:  c,
		}
		if c.element != nil && c.element != nil && c.element.Get().Style != nil {
			cc.w = MeasureLength(c.element.Get().Style.Get().Width, c.w)
			cc.h = MeasureLength(c.element.Get().Style.Get().Height, c.h)
		}
		c.children = append(c.children, cc)
	}
	switch {
	default:
		fallthrough
	case c.element.Get().Style == nil, c.element.Get().Style.Get().Display == nil,
		c.element.Get().Style.Get().Display.Get() == "":

	case c.element.Get().Style.Get().Display.Get() == "flex":
		c.flex()
	case c.element.Get().Style.Get().Display.Get() == "grid":

	}
}

func (c *container) flex() {
	var sorted = make([]*container, len(c.children))
	copy(sorted, c.children)
	sort.SliceStable(sorted, func(i, j int) bool {
		var orderI int
		var orderJ int
		if sorted[i] != nil && sorted[i].element != nil && sorted[i].element.Get().Style != nil &&
			sorted[i].element.Get().Style.Get().Order == nil {
			orderI = sorted[i].element.Get().Style.Get().Order.Get()
		}
		if sorted[j] != nil && sorted[j].element != nil && sorted[j].element.Get().Style != nil &&
			sorted[j].element.Get().Style.Get().Order == nil {
			orderJ = sorted[j].element.Get().Style.Get().Order.Get()
		}
		return orderI < orderJ
	})
	var (
		getx func(int) uint64
		setx func(int, uint64)
		gety func(int) uint64
		sety func(int, uint64)
		getw func(int) uint64
		setw func(int, uint64)
		geth func(int) uint64
		seth func(int, uint64)
		w    func() uint64
		h    func() uint64
	)
	switch {
	default:
		fallthrough
	case c.element == nil, c.element.Get().Style == nil:
		fallthrough
	case c.element.Get().Style.Get().FlexDirection == nil,
		c.element.Get().Style.Get().FlexDirection.Get() == "row":
		getx = func(i int) uint64 { return sorted[i].x }
		setx = func(i int, x uint64) { sorted[i].x = x }
		gety = func(i int) uint64 { return sorted[i].y }
		sety = func(i int, y uint64) { sorted[i].y = y }
		getw = func(i int) uint64 { return sorted[i].w }
		setw = func(i int, w uint64) { sorted[i].w = w }
		geth = func(i int) uint64 { return sorted[i].h }
		seth = func(i int, h uint64) { sorted[i].h = h }
		w = func() uint64 { return c.w }
		h = func() uint64 { return c.h }
	case c.element.Get().Style.Get().FlexDirection.Get() == "row-reverse":
		getx = func(i int) uint64 { return sorted[i].w - sorted[i].x }
		setx = func(i int, x uint64) { sorted[i].x = x - sorted[i].w }
		gety = func(i int) uint64 { return sorted[i].y }
		sety = func(i int, y uint64) { sorted[i].y = y }
		getw = func(i int) uint64 { return sorted[i].w }
		setw = func(i int, w uint64) { sorted[i].w = w }
		geth = func(i int) uint64 { return sorted[i].h }
		seth = func(i int, h uint64) { sorted[i].h = h }
		w = func() uint64 { return c.w }
		h = func() uint64 { return c.h }
	case c.element.Get().Style.Get().FlexDirection.Get() == "column":
		getx = func(i int) uint64 { return sorted[i].y }
		setx = func(i int, x uint64) { sorted[i].y = x }
		gety = func(i int) uint64 { return sorted[i].x }
		sety = func(i int, y uint64) { sorted[i].x = y }
		getw = func(i int) uint64 { return sorted[i].h }
		setw = func(i int, w uint64) { sorted[i].h = w }
		geth = func(i int) uint64 { return sorted[i].w }
		seth = func(i int, h uint64) { sorted[i].w = h }
		w = func() uint64 { return c.h }
		h = func() uint64 { return c.w }
	case c.element.Get().Style.Get().FlexDirection.Get() == "column-reverse":
		getx = func(i int) uint64 { return sorted[i].h - sorted[i].y }
		setx = func(i int, x uint64) { sorted[i].y = x - sorted[i].h }
		gety = func(i int) uint64 { return sorted[i].x }
		sety = func(i int, y uint64) { sorted[i].x = y }
		getw = func(i int) uint64 { return sorted[i].h }
		setw = func(i int, w uint64) { sorted[i].h = w }
		geth = func(i int) uint64 { return sorted[i].w }
		seth = func(i int, h uint64) { sorted[i].w = h }
		w = func() uint64 { return c.h }
		h = func() uint64 { return c.w }
	}
	for i := range sorted {
		if sorted[i] != nil && sorted[i].element != nil && sorted[i].element.Get().Style != nil &&
			sorted[i].element.Get().Style.Get().FlexBasis != nil {
			setw(i, MeasureLength(sorted[i].element.Get().Style.Get().FlexBasis, w()))
		}
	}
	var (
		table  = make(map[int]int)
		bundle = make(map[int][]int)
		sumw   = make(map[int]uint64)
		sumg   = make(map[int]uint64)
		sums   = make(map[int]uint64)
	)
	var volume int
	sumw[volume] = 0
	sumg[volume] = 0
	sums[volume] = 0
	for i := range sorted {
		switch {
		default:
			fallthrough
		case sorted[i] == nil, sorted[i].element == nil, sorted[i].element.Get().Style == nil:
			fallthrough
		case sorted[i].element.Get().Style.Get().FlexWrap == nil,
			sorted[i].element.Get().Style.Get().FlexWrap.Get() == "nowrap":
			break
		case sorted[i].element.Get().Style.Get().FlexWrap.Get() == "warp",
			sorted[i].element.Get().Style.Get().FlexWrap.Get() == "warp-reverse":
			if sumw[i]+getw(i) > c.w {
				volume++
				sumw[volume] = 0
				sumg[volume] = 0
				sums[volume] = 0
			}
		}
		table[i] = volume
		bundle[volume] = append(bundle[volume], i)
		sumw[volume] += getw(i)
		if sorted[i] != nil && sorted[i].element != nil && sorted[i].element.Get().Style != nil {
			if sorted[i].element.Get().Style.Get().FlexGrow != nil {
				sumg[volume] += sorted[i].element.Get().Style.Get().FlexGrow.Get()
			}
			if sorted[i].element.Get().Style.Get().FlexShrink != nil {
				sums[volume] += sorted[i].element.Get().Style.Get().FlexShrink.Get()
			}
		}
	}
	for k, v := range sumg {
		if v > 0 && sumw[k] > w() {
			for _, i := range bundle[k] {
				if sorted[i] != nil && sorted[i].element != nil && sorted[i].element.Get().Style != nil &&
					sorted[i].element.Get().Style.Get().FlexGrow != nil {
					setw(i, getw(i)+(sorted[i].element.Get().Style.Get().FlexGrow.Get()/v)*(sumw[k]-w()))
				}
			}
		}
	}
	for k, v := range sumg {
		if v > 0 && sumw[k] < w() {
			for _, i := range bundle[k] {
				if sorted[i] != nil && sorted[i].element != nil && sorted[i].element.Get().Style != nil &&
					sorted[i].element.Get().Style.Get().FlexShrink != nil {
					setw(i, getw(i)+(sorted[i].element.Get().Style.Get().FlexShrink.Get()/v)*(sumw[k]-w()))
				}
			}
		}
	}
}
