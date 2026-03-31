package linux

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	keyReleased = 0
	keyPressed  = 1
)

type InputDevice struct {
	file        *os.File
	initialized bool
}

func NewInputDevice() (inputDevice InputDevice, err error) {
	file, err := os.OpenFile("/dev/uinput", syscall.O_WRONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		return inputDevice, err
	}
	return InputDevice{file: file}, nil
}

func (device *InputDevice) Close() error {
	if device.file == nil {
		return nil
	}
	if device.initialized {
		if err := Ioctl(device.file.Fd(), UI_DEV_DESTROY, 0); err != nil {
			return err
		}
	}
	return device.file.Close()
}

func (device *InputDevice) initialize(name string) error {
	var uinputUserDev UinputUserDev
	copy(uinputUserDev.Name[:], name)
	uinputUserDev.ID.Bustype = BUS_USB
	uinputUserDev.ID.Vendor = 0x1234
	uinputUserDev.ID.Product = 0x5678
	uinputUserDev.ID.Version = 1
	if err := Ioctl(device.file.Fd(), UI_DEV_SETUP, uintptr(unsafe.Pointer(&uinputUserDev))); err != nil {
		return err
	}
	if err := Ioctl(device.file.Fd(), UI_DEV_CREATE, 0); err != nil {
		return err
	}
	device.initialized = true
	return nil
}

func (device *InputDevice) emitKeyEvent(code uint16, value int32) error {
	return NewInputEvent(EV_KEY, code, value).Emit(device.file)
}

func (device *InputDevice) emitSynEvent() error {
	return NewInputEvent(EV_SYN, SYN_REPORT, 0).Emit(device.file)
}
