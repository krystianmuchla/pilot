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
