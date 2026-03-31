package logic

type Mouse interface {
	Move(x int, y int) error
	Scroll(x int, y int) error
	PressButton1() error
	ReleaseButton1() error
	PressButton2() error
	ReleaseButton2() error
	PressButton3() error
	ReleaseButton3() error
}
