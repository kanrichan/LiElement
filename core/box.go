package core

type box struct {
	X int
	Y int
	W int
	H int
}

func (box1 box) IsOverlap(box2 box) bool {
	return !(box1.X+box1.W < box2.X ||
		box1.Y+box1.H < box2.Y ||
		box2.X+box2.W < box1.X ||
		box2.Y+box2.H < box1.Y)
}
