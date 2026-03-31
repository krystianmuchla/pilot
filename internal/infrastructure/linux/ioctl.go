package linux

import "syscall"

func Ioctl(fd, req, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, req, arg)
	if err != 0 {
		return err
	}
	return nil
}
