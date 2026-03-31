package linux

const (
	BTN_LEFT   = 0x110
	BTN_RIGHT  = 0x111
	BTN_MIDDLE = 0x112
	REL_X      = 0x00
	REL_Y      = 0x01
	REL_HWHEEL = 0x06
	REL_WHEEL  = 0x08
)

type Mouse struct {
	InputDevice
}

func NewMouse(name string) (*Mouse, error) {
	inputDevice, err := NewInputDevice()
	if err != nil {
		return nil, err
	}
	mouse := Mouse{InputDevice: inputDevice}
	if err = mouse.initializeKeys(); err != nil {
		_ = mouse.Close()
		return nil, err
	}
	if err = mouse.initializeRels(); err != nil {
		_ = mouse.Close()
		return nil, err
	}
	if err = mouse.initialize(name); err != nil {
		_ = mouse.Close()
		return nil, err
	}
	return &mouse, nil
}

func (mouse *Mouse) Move(x int, y int) error {
	if err := mouse.emitRelEvent(REL_X, int32(x)); err != nil {
		return err
	}
	if err := mouse.emitRelEvent(REL_Y, int32(y)); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) Scroll(x int, y int) error {
	if err := mouse.emitRelEvent(REL_HWHEEL, int32(x)); err != nil {
		return err
	}
	if err := mouse.emitRelEvent(REL_WHEEL, int32(y)); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) PressButton1() error {
	if err := mouse.emitKeyEvent(BTN_LEFT, keyPressed); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) ReleaseButton1() error {
	if err := mouse.emitKeyEvent(BTN_LEFT, keyReleased); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) PressButton2() error {
	if err := mouse.emitKeyEvent(BTN_MIDDLE, keyPressed); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) ReleaseButton2() error {
	if err := mouse.emitKeyEvent(BTN_MIDDLE, keyReleased); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) PressButton3() error {
	if err := mouse.emitKeyEvent(BTN_RIGHT, keyPressed); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) ReleaseButton3() error {
	if err := mouse.emitKeyEvent(BTN_RIGHT, keyReleased); err != nil {
		return err
	}
	return mouse.emitSynEvent()
}

func (mouse *Mouse) initializeKeys() error {
	if err := Ioctl(mouse.file.Fd(), UI_SET_EVBIT, uintptr(EV_KEY)); err != nil {
		return err
	}
	for _, key := range []uint16{BTN_LEFT, BTN_RIGHT, BTN_MIDDLE} {
		if err := Ioctl(mouse.file.Fd(), UI_SET_KEYBIT, uintptr(key)); err != nil {
			return err
		}
	}
	return nil
}

func (mouse *Mouse) initializeRels() error {
	if err := Ioctl(mouse.file.Fd(), UI_SET_EVBIT, uintptr(EV_REL)); err != nil {
		return err
	}
	for _, rel := range []uint16{REL_X, REL_Y, REL_HWHEEL, REL_WHEEL} {
		if err := Ioctl(mouse.file.Fd(), UI_SET_RELBIT, uintptr(rel)); err != nil {
			return err
		}
	}
	return nil
}

func (mouse *Mouse) emitRelEvent(code uint16, value int32) error {
	return NewInputEvent(EV_REL, code, value).Emit(mouse.file)
}
