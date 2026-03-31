package linux

import (
	"encoding/binary"
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	EV_SYN     = 0x00
	EV_KEY     = 0x01
	EV_REL     = 0x02
	SYN_REPORT = 0
)

var emitMutex = sync.Mutex{}
var evSyncMinInterval = 5 * time.Millisecond
var lastEvSynTime = time.Now()

type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

func NewInputEvent(typ, code uint16, value int32) *InputEvent {
	t := syscall.Timeval{Sec: 0, Usec: 0}
	return &InputEvent{Time: t, Type: typ, Code: code, Value: value}
}

func (inputEvent *InputEvent) Emit(deviceFile *os.File) error {
	emitMutex.Lock()
	defer emitMutex.Unlock()
	if inputEvent.Type == EV_SYN {
		time.Sleep(evSyncMinInterval - time.Since(lastEvSynTime))
		defer func() { lastEvSynTime = time.Now() }()
	}
	return binary.Write(deviceFile, binary.LittleEndian, inputEvent)
}
