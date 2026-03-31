package linux

const (
	KEY_ESC        = 1
	KEY_1          = 2
	KEY_2          = 3
	KEY_3          = 4
	KEY_4          = 5
	KEY_5          = 6
	KEY_6          = 7
	KEY_7          = 8
	KEY_8          = 9
	KEY_9          = 10
	KEY_0          = 11
	KEY_MINUS      = 12
	KEY_EQUAL      = 13
	KEY_BACKSPACE  = 14
	KEY_TAB        = 15
	KEY_Q          = 16
	KEY_W          = 17
	KEY_E          = 18
	KEY_R          = 19
	KEY_T          = 20
	KEY_Y          = 21
	KEY_U          = 22
	KEY_I          = 23
	KEY_O          = 24
	KEY_P          = 25
	KEY_LEFTBRACE  = 26
	KEY_RIGHTBRACE = 27
	KEY_ENTER      = 28
	KEY_A          = 30
	KEY_S          = 31
	KEY_D          = 32
	KEY_F          = 33
	KEY_G          = 34
	KEY_H          = 35
	KEY_J          = 36
	KEY_K          = 37
	KEY_L          = 38
	KEY_SEMICOLON  = 39
	KEY_APOSTROPHE = 40
	KEY_GRAVE      = 41
	KEY_LEFTSHIFT  = 42
	KEY_BACKSLASH  = 43
	KEY_Z          = 44
	KEY_X          = 45
	KEY_C          = 46
	KEY_V          = 47
	KEY_B          = 48
	KEY_N          = 49
	KEY_M          = 50
	KEY_COMMA      = 51
	KEY_DOT        = 52
	KEY_SLASH      = 53
	KEY_SPACE      = 57
	KEY_CAPSLOCK   = 58
	KEY_F1         = 59
	KEY_F2         = 60
	KEY_F3         = 61
	KEY_F4         = 62
	KEY_F5         = 63
	KEY_F6         = 64
	KEY_F7         = 65
	KEY_F8         = 66
	KEY_F9         = 67
	KEY_F10        = 68
	KEY_F11        = 87
	KEY_F12        = 88
	KEY_RIGHTALT   = 100
	KEY_HOME       = 102
	KEY_UP         = 103
	KEY_PAGEUP     = 104
	KEY_LEFT       = 105
	KEY_RIGHT      = 106
	KEY_END        = 107
	KEY_DOWN       = 108
	KEY_PAGEDOWN   = 109
	KEY_INSERT     = 110
	KEY_DELETE     = 111
)

type Keyboard struct {
	InputDevice
}

func NewKeyboard(name string) (*Keyboard, error) {
	inputDevice, err := NewInputDevice()
	if err != nil {
		return nil, err
	}
	keyboard := Keyboard{InputDevice: inputDevice}
	if err = keyboard.initializeKeys(); err != nil {
		_ = keyboard.Close()
		return nil, err
	}
	if err = keyboard.initialize(name); err != nil {
		_ = keyboard.Close()
		return nil, err
	}
	return &keyboard, nil
}

func (keyboard *Keyboard) PressSpace() error {
	if err := keyboard.emitKeyEvent(KEY_SPACE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseSpace() error {
	if err := keyboard.emitKeyEvent(KEY_SPACE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressLeftShift() error {
	if err := keyboard.emitKeyEvent(KEY_LEFTSHIFT, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseLeftShift() error {
	if err := keyboard.emitKeyEvent(KEY_LEFTSHIFT, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressRightAlt() error {
	if err := keyboard.emitKeyEvent(KEY_RIGHTALT, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseRightAlt() error {
	if err := keyboard.emitKeyEvent(KEY_RIGHTALT, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressApostrophe() error {
	if err := keyboard.emitKeyEvent(KEY_APOSTROPHE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseApostrophe() error {
	if err := keyboard.emitKeyEvent(KEY_APOSTROPHE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressComma() error {
	if err := keyboard.emitKeyEvent(KEY_COMMA, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseComma() error {
	if err := keyboard.emitKeyEvent(KEY_COMMA, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressMinus() error {
	if err := keyboard.emitKeyEvent(KEY_MINUS, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseMinus() error {
	if err := keyboard.emitKeyEvent(KEY_MINUS, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressDot() error {
	if err := keyboard.emitKeyEvent(KEY_DOT, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseDot() error {
	if err := keyboard.emitKeyEvent(KEY_DOT, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressSlash() error {
	if err := keyboard.emitKeyEvent(KEY_SLASH, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseSlash() error {
	if err := keyboard.emitKeyEvent(KEY_SLASH, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press0() error {
	if err := keyboard.emitKeyEvent(KEY_0, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release0() error {
	if err := keyboard.emitKeyEvent(KEY_0, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press1() error {
	if err := keyboard.emitKeyEvent(KEY_1, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release1() error {
	if err := keyboard.emitKeyEvent(KEY_1, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press2() error {
	if err := keyboard.emitKeyEvent(KEY_2, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release2() error {
	if err := keyboard.emitKeyEvent(KEY_2, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press3() error {
	if err := keyboard.emitKeyEvent(KEY_3, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release3() error {
	if err := keyboard.emitKeyEvent(KEY_3, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press4() error {
	if err := keyboard.emitKeyEvent(KEY_4, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release4() error {
	if err := keyboard.emitKeyEvent(KEY_4, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press5() error {
	if err := keyboard.emitKeyEvent(KEY_5, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release5() error {
	if err := keyboard.emitKeyEvent(KEY_5, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press6() error {
	if err := keyboard.emitKeyEvent(KEY_6, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release6() error {
	if err := keyboard.emitKeyEvent(KEY_6, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press7() error {
	if err := keyboard.emitKeyEvent(KEY_7, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release7() error {
	if err := keyboard.emitKeyEvent(KEY_7, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press8() error {
	if err := keyboard.emitKeyEvent(KEY_8, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release8() error {
	if err := keyboard.emitKeyEvent(KEY_8, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Press9() error {
	if err := keyboard.emitKeyEvent(KEY_9, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) Release9() error {
	if err := keyboard.emitKeyEvent(KEY_9, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressSemicolon() error {
	if err := keyboard.emitKeyEvent(KEY_SEMICOLON, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseSemicolon() error {
	if err := keyboard.emitKeyEvent(KEY_SEMICOLON, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressEqual() error {
	if err := keyboard.emitKeyEvent(KEY_EQUAL, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseEqual() error {
	if err := keyboard.emitKeyEvent(KEY_EQUAL, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressA() error {
	if err := keyboard.emitKeyEvent(KEY_A, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseA() error {
	if err := keyboard.emitKeyEvent(KEY_A, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressB() error {
	if err := keyboard.emitKeyEvent(KEY_B, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseB() error {
	if err := keyboard.emitKeyEvent(KEY_B, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressC() error {
	if err := keyboard.emitKeyEvent(KEY_C, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseC() error {
	if err := keyboard.emitKeyEvent(KEY_C, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressD() error {
	if err := keyboard.emitKeyEvent(KEY_D, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseD() error {
	if err := keyboard.emitKeyEvent(KEY_D, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressE() error {
	if err := keyboard.emitKeyEvent(KEY_E, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseE() error {
	if err := keyboard.emitKeyEvent(KEY_E, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF() error {
	if err := keyboard.emitKeyEvent(KEY_F, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF() error {
	if err := keyboard.emitKeyEvent(KEY_F, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressG() error {
	if err := keyboard.emitKeyEvent(KEY_G, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseG() error {
	if err := keyboard.emitKeyEvent(KEY_G, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressH() error {
	if err := keyboard.emitKeyEvent(KEY_H, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseH() error {
	if err := keyboard.emitKeyEvent(KEY_H, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressI() error {
	if err := keyboard.emitKeyEvent(KEY_I, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseI() error {
	if err := keyboard.emitKeyEvent(KEY_I, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressJ() error {
	if err := keyboard.emitKeyEvent(KEY_J, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseJ() error {
	if err := keyboard.emitKeyEvent(KEY_J, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressK() error {
	if err := keyboard.emitKeyEvent(KEY_K, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseK() error {
	if err := keyboard.emitKeyEvent(KEY_K, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressL() error {
	if err := keyboard.emitKeyEvent(KEY_L, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseL() error {
	if err := keyboard.emitKeyEvent(KEY_L, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressM() error {
	if err := keyboard.emitKeyEvent(KEY_M, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseM() error {
	if err := keyboard.emitKeyEvent(KEY_M, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressN() error {
	if err := keyboard.emitKeyEvent(KEY_N, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseN() error {
	if err := keyboard.emitKeyEvent(KEY_N, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressO() error {
	if err := keyboard.emitKeyEvent(KEY_O, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseO() error {
	if err := keyboard.emitKeyEvent(KEY_O, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressP() error {
	if err := keyboard.emitKeyEvent(KEY_P, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseP() error {
	if err := keyboard.emitKeyEvent(KEY_P, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressQ() error {
	if err := keyboard.emitKeyEvent(KEY_Q, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseQ() error {
	if err := keyboard.emitKeyEvent(KEY_Q, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressR() error {
	if err := keyboard.emitKeyEvent(KEY_R, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseR() error {
	if err := keyboard.emitKeyEvent(KEY_R, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressS() error {
	if err := keyboard.emitKeyEvent(KEY_S, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseS() error {
	if err := keyboard.emitKeyEvent(KEY_S, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressT() error {
	if err := keyboard.emitKeyEvent(KEY_T, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseT() error {
	if err := keyboard.emitKeyEvent(KEY_T, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressU() error {
	if err := keyboard.emitKeyEvent(KEY_U, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseU() error {
	if err := keyboard.emitKeyEvent(KEY_U, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressV() error {
	if err := keyboard.emitKeyEvent(KEY_V, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseV() error {
	if err := keyboard.emitKeyEvent(KEY_V, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressW() error {
	if err := keyboard.emitKeyEvent(KEY_W, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseW() error {
	if err := keyboard.emitKeyEvent(KEY_W, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressX() error {
	if err := keyboard.emitKeyEvent(KEY_X, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseX() error {
	if err := keyboard.emitKeyEvent(KEY_X, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressY() error {
	if err := keyboard.emitKeyEvent(KEY_Y, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseY() error {
	if err := keyboard.emitKeyEvent(KEY_Y, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressZ() error {
	if err := keyboard.emitKeyEvent(KEY_Z, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseZ() error {
	if err := keyboard.emitKeyEvent(KEY_Z, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressLeftBrace() error {
	if err := keyboard.emitKeyEvent(KEY_LEFTBRACE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseLeftBrace() error {
	if err := keyboard.emitKeyEvent(KEY_LEFTBRACE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressBackslash() error {
	if err := keyboard.emitKeyEvent(KEY_BACKSLASH, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseBackslash() error {
	if err := keyboard.emitKeyEvent(KEY_BACKSLASH, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressRightBrace() error {
	if err := keyboard.emitKeyEvent(KEY_RIGHTBRACE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseRightBrace() error {
	if err := keyboard.emitKeyEvent(KEY_RIGHTBRACE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressGrave() error {
	if err := keyboard.emitKeyEvent(KEY_GRAVE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseGrave() error {
	if err := keyboard.emitKeyEvent(KEY_GRAVE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressEsc() error {
	if err := keyboard.emitKeyEvent(KEY_ESC, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseEsc() error {
	if err := keyboard.emitKeyEvent(KEY_ESC, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF1() error {
	if err := keyboard.emitKeyEvent(KEY_F1, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF1() error {
	if err := keyboard.emitKeyEvent(KEY_F1, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF2() error {
	if err := keyboard.emitKeyEvent(KEY_F2, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF2() error {
	if err := keyboard.emitKeyEvent(KEY_F2, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF3() error {
	if err := keyboard.emitKeyEvent(KEY_F3, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF3() error {
	if err := keyboard.emitKeyEvent(KEY_F3, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF4() error {
	if err := keyboard.emitKeyEvent(KEY_F4, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF4() error {
	if err := keyboard.emitKeyEvent(KEY_F4, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF5() error {
	if err := keyboard.emitKeyEvent(KEY_F5, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF5() error {
	if err := keyboard.emitKeyEvent(KEY_F5, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF6() error {
	if err := keyboard.emitKeyEvent(KEY_F6, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF6() error {
	if err := keyboard.emitKeyEvent(KEY_F6, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF7() error {
	if err := keyboard.emitKeyEvent(KEY_F7, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF7() error {
	if err := keyboard.emitKeyEvent(KEY_F7, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF8() error {
	if err := keyboard.emitKeyEvent(KEY_F8, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF8() error {
	if err := keyboard.emitKeyEvent(KEY_F8, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF9() error {
	if err := keyboard.emitKeyEvent(KEY_F9, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF9() error {
	if err := keyboard.emitKeyEvent(KEY_F9, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF10() error {
	if err := keyboard.emitKeyEvent(KEY_F10, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF10() error {
	if err := keyboard.emitKeyEvent(KEY_F10, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF11() error {
	if err := keyboard.emitKeyEvent(KEY_F11, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF11() error {
	if err := keyboard.emitKeyEvent(KEY_F11, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressF12() error {
	if err := keyboard.emitKeyEvent(KEY_F12, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseF12() error {
	if err := keyboard.emitKeyEvent(KEY_F12, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressHome() error {
	if err := keyboard.emitKeyEvent(KEY_HOME, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseHome() error {
	if err := keyboard.emitKeyEvent(KEY_HOME, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressEnd() error {
	if err := keyboard.emitKeyEvent(KEY_END, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseEnd() error {
	if err := keyboard.emitKeyEvent(KEY_END, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressInsert() error {
	if err := keyboard.emitKeyEvent(KEY_INSERT, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseInsert() error {
	if err := keyboard.emitKeyEvent(KEY_INSERT, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressDelete() error {
	if err := keyboard.emitKeyEvent(KEY_DELETE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseDelete() error {
	if err := keyboard.emitKeyEvent(KEY_DELETE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressBackspace() error {
	if err := keyboard.emitKeyEvent(KEY_BACKSPACE, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseBackspace() error {
	if err := keyboard.emitKeyEvent(KEY_BACKSPACE, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressEnter() error {
	if err := keyboard.emitKeyEvent(KEY_ENTER, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseEnter() error {
	if err := keyboard.emitKeyEvent(KEY_ENTER, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressPageUp() error {
	if err := keyboard.emitKeyEvent(KEY_PAGEUP, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleasePageUp() error {
	if err := keyboard.emitKeyEvent(KEY_PAGEUP, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressPageDown() error {
	if err := keyboard.emitKeyEvent(KEY_PAGEDOWN, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleasePageDown() error {
	if err := keyboard.emitKeyEvent(KEY_PAGEDOWN, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressUp() error {
	if err := keyboard.emitKeyEvent(KEY_UP, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseUp() error {
	if err := keyboard.emitKeyEvent(KEY_UP, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressDown() error {
	if err := keyboard.emitKeyEvent(KEY_DOWN, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseDown() error {
	if err := keyboard.emitKeyEvent(KEY_DOWN, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressLeft() error {
	if err := keyboard.emitKeyEvent(KEY_LEFT, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseLeft() error {
	if err := keyboard.emitKeyEvent(KEY_LEFT, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressRight() error {
	if err := keyboard.emitKeyEvent(KEY_RIGHT, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseRight() error {
	if err := keyboard.emitKeyEvent(KEY_RIGHT, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressCapslock() error {
	if err := keyboard.emitKeyEvent(KEY_CAPSLOCK, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseCapslock() error {
	if err := keyboard.emitKeyEvent(KEY_CAPSLOCK, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) PressTab() error {
	if err := keyboard.emitKeyEvent(KEY_TAB, keyPressed); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) ReleaseTab() error {
	if err := keyboard.emitKeyEvent(KEY_TAB, keyReleased); err != nil {
		return err
	}
	return keyboard.emitSynEvent()
}

func (keyboard *Keyboard) initializeKeys() error {
	if err := Ioctl(keyboard.file.Fd(), UI_SET_EVBIT, uintptr(EV_KEY)); err != nil {
		return err
	}
	for _, key := range []uint16{
		KEY_ESC,
		KEY_1,
		KEY_2,
		KEY_3,
		KEY_4,
		KEY_5,
		KEY_6,
		KEY_7,
		KEY_8,
		KEY_9,
		KEY_0,
		KEY_MINUS,
		KEY_EQUAL,
		KEY_BACKSPACE,
		KEY_TAB,
		KEY_Q,
		KEY_W,
		KEY_E,
		KEY_R,
		KEY_T,
		KEY_Y,
		KEY_U,
		KEY_I,
		KEY_O,
		KEY_P,
		KEY_LEFTBRACE,
		KEY_RIGHTBRACE,
		KEY_ENTER,
		KEY_A,
		KEY_S,
		KEY_D,
		KEY_F,
		KEY_G,
		KEY_H,
		KEY_J,
		KEY_K,
		KEY_L,
		KEY_SEMICOLON,
		KEY_APOSTROPHE,
		KEY_GRAVE,
		KEY_LEFTSHIFT,
		KEY_BACKSLASH,
		KEY_Z,
		KEY_X,
		KEY_C,
		KEY_V,
		KEY_B,
		KEY_N,
		KEY_M,
		KEY_COMMA,
		KEY_DOT,
		KEY_SLASH,
		KEY_SPACE,
		KEY_CAPSLOCK,
		KEY_F1,
		KEY_F2,
		KEY_F3,
		KEY_F4,
		KEY_F5,
		KEY_F6,
		KEY_F7,
		KEY_F8,
		KEY_F9,
		KEY_F10,
		KEY_F11,
		KEY_F12,
		KEY_RIGHTALT,
		KEY_HOME,
		KEY_UP,
		KEY_PAGEUP,
		KEY_LEFT,
		KEY_RIGHT,
		KEY_END,
		KEY_DOWN,
		KEY_PAGEDOWN,
		KEY_INSERT,
		KEY_DELETE,
	} {
		if err := Ioctl(keyboard.file.Fd(), UI_SET_KEYBIT, uintptr(key)); err != nil {
			return err
		}
	}
	return nil
}
