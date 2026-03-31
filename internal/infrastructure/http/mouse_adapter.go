package http

import "pilot/internal/logic"

type MouseAdapter struct {
	mouse logic.Mouse
}

func NewMouseAdapter(mouse logic.Mouse) MouseAdapter {
	return MouseAdapter{mouse: mouse}
}

func (adapter MouseAdapter) Move(x int, y int) error {
	return adapter.mouse.Move(x, y)
}

func (adapter MouseAdapter) ClickButton1() error {
	if err := adapter.mouse.PressButton1(); err != nil {
		return err
	}
	return adapter.mouse.ReleaseButton1()
}

func (adapter MouseAdapter) Scroll(x int, y int) error {
	return adapter.mouse.Scroll(x, -y)
}

func (adapter MouseAdapter) PressButton1() error {
	return adapter.mouse.PressButton1()
}

func (adapter MouseAdapter) ReleaseButton1() error {
	return adapter.mouse.ReleaseButton1()
}

func (adapter MouseAdapter) PressButton2() error {
	return adapter.mouse.PressButton2()
}

func (adapter MouseAdapter) ReleaseButton2() error {
	return adapter.mouse.ReleaseButton2()
}

func (adapter MouseAdapter) PressButton3() error {
	return adapter.mouse.PressButton3()
}

func (adapter MouseAdapter) ReleaseButton3() error {
	return adapter.mouse.ReleaseButton3()
}
