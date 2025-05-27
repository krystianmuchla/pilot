package logic

import (
	"github.com/bendahl/uinput"
)

const mouseName = "pilot-mouse"

type MouseController struct {
	mouse uinput.Mouse
}

func NewMouseController() (*MouseController, error) {
	mouse, err := uinput.CreateMouse("/dev/uinput", []byte(mouseName))
	if err != nil {
		return nil, err
	}
	return &MouseController{mouse}, nil
}

func (controller *MouseController) Move(x int, y int) error {
	err := controller.mouse.Move(int32(x), int32(y))
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) ClickButton1() error {
	err := controller.mouse.LeftClick()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) Scroll(x int, y int) error {
	err := controller.mouse.Wheel(true, int32(x))
	if err != nil {
		return err
	}
	err = controller.mouse.Wheel(false, int32(y*-1))
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) PressButton1() error {
	err := controller.mouse.LeftPress()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) ReleaseButton1() error {
	err := controller.mouse.LeftRelease()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) PressButton2() error {
	err := controller.mouse.MiddlePress()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) ReleaseButton2() error {
	err := controller.mouse.MiddleRelease()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) PressButton3() error {
	err := controller.mouse.RightPress()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) ReleaseButton3() error {
	err := controller.mouse.RightRelease()
	if err != nil {
		return err
	}
	return nil
}

func (controller *MouseController) Stop() error {
	return controller.mouse.Close()
}
