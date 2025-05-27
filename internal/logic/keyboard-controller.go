package logic

import "github.com/bendahl/uinput"

const keyboardName = "pilot-keyboard"

type KeyboardController struct {
	keyboard uinput.Keyboard
}

func NewKeyboardController() (*KeyboardController, error) {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte(keyboardName))
	if err != nil {
		return nil, err
	}
	return &KeyboardController{keyboard}, nil
}

func (controller *KeyboardController) EnterKey(keyCode byte) error {
	switch keyCode {
	case 32:
		return controller.enterKey(uinput.KeySpace)
	case 33:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key1)
	case 34:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyApostrophe)
	case 35:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key3)
	case 36:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key4)
	case 37:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key5)
	case 38:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key7)
	case 39:
		return controller.enterKey(uinput.KeyApostrophe)
	case 40:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key9)
	case 41:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key0)
	case 42:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key8)
	case 43:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyEqual)
	case 44:
		return controller.enterKey(uinput.KeyComma)
	case 45:
		return controller.enterKey(uinput.KeyMinus)
	case 46:
		return controller.enterKey(uinput.KeyDot)
	case 47:
		return controller.enterKey(uinput.KeySlash)
	case 48:
		return controller.enterKey(uinput.Key0)
	case 49:
		return controller.enterKey(uinput.Key1)
	case 50:
		return controller.enterKey(uinput.Key2)
	case 51:
		return controller.enterKey(uinput.Key3)
	case 52:
		return controller.enterKey(uinput.Key4)
	case 53:
		return controller.enterKey(uinput.Key5)
	case 54:
		return controller.enterKey(uinput.Key6)
	case 55:
		return controller.enterKey(uinput.Key7)
	case 56:
		return controller.enterKey(uinput.Key8)
	case 57:
		return controller.enterKey(uinput.Key9)
	case 58:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeySemicolon)
	case 59:
		return controller.enterKey(uinput.KeySemicolon)
	case 60:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyComma)
	case 61:
		return controller.enterKey(uinput.KeyEqual)
	case 62:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyDot)
	case 63:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeySlash)
	case 64:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key2)
	case 65:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyA)
	case 66:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyB)
	case 67:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyC)
	case 68:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyD)
	case 69:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyE)
	case 70:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyF)
	case 71:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyG)
	case 72:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyH)
	case 73:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyI)
	case 74:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyJ)
	case 75:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyK)
	case 76:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyL)
	case 77:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyM)
	case 78:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyN)
	case 79:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyO)
	case 80:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyP)
	case 81:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyQ)
	case 82:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyR)
	case 83:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyS)
	case 84:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyT)
	case 85:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyU)
	case 86:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyV)
	case 87:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyW)
	case 88:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyX)
	case 89:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyY)
	case 90:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyZ)
	case 91:
		return controller.enterKey(uinput.KeyLeftbrace)
	case 92:
		return controller.enterKey(uinput.KeyBackslash)
	case 93:
		return controller.enterKey(uinput.KeyRightbrace)
	case 94:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.Key6)
	case 95:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyMinus)
	case 96:
		return controller.enterKey(uinput.KeyGrave)
	case 97:
		return controller.enterKey(uinput.KeyA)
	case 98:
		return controller.enterKey(uinput.KeyB)
	case 99:
		return controller.enterKey(uinput.KeyC)
	case 100:
		return controller.enterKey(uinput.KeyD)
	case 101:
		return controller.enterKey(uinput.KeyE)
	case 102:
		return controller.enterKey(uinput.KeyF)
	case 103:
		return controller.enterKey(uinput.KeyG)
	case 104:
		return controller.enterKey(uinput.KeyH)
	case 105:
		return controller.enterKey(uinput.KeyI)
	case 106:
		return controller.enterKey(uinput.KeyJ)
	case 107:
		return controller.enterKey(uinput.KeyK)
	case 108:
		return controller.enterKey(uinput.KeyL)
	case 109:
		return controller.enterKey(uinput.KeyM)
	case 110:
		return controller.enterKey(uinput.KeyN)
	case 111:
		return controller.enterKey(uinput.KeyO)
	case 112:
		return controller.enterKey(uinput.KeyP)
	case 113:
		return controller.enterKey(uinput.KeyQ)
	case 114:
		return controller.enterKey(uinput.KeyR)
	case 115:
		return controller.enterKey(uinput.KeyS)
	case 116:
		return controller.enterKey(uinput.KeyT)
	case 117:
		return controller.enterKey(uinput.KeyU)
	case 118:
		return controller.enterKey(uinput.KeyV)
	case 119:
		return controller.enterKey(uinput.KeyW)
	case 120:
		return controller.enterKey(uinput.KeyX)
	case 121:
		return controller.enterKey(uinput.KeyY)
	case 122:
		return controller.enterKey(uinput.KeyZ)
	case 123:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyLeftbrace)
	case 124:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyBackslash)
	case 125:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyRightbrace)
	case 126:
		return controller.enterDoubleKey(uinput.KeyLeftshift, uinput.KeyGrave)
	case 127:
		return controller.enterKey(uinput.KeyEsc)
	case 128:
		return controller.enterKey(uinput.KeyF1)
	case 129:
		return controller.enterKey(uinput.KeyF2)
	case 130:
		return controller.enterKey(uinput.KeyF3)
	case 131:
		return controller.enterKey(uinput.KeyF4)
	case 132:
		return controller.enterKey(uinput.KeyF5)
	case 133:
		return controller.enterKey(uinput.KeyF6)
	case 134:
		return controller.enterKey(uinput.KeyF7)
	case 135:
		return controller.enterKey(uinput.KeyF8)
	case 136:
		return controller.enterKey(uinput.KeyF9)
	case 137:
		return controller.enterKey(uinput.KeyF10)
	case 138:
		return controller.enterKey(uinput.KeyF11)
	case 139:
		return controller.enterKey(uinput.KeyF12)
	case 140:
		return controller.enterKey(uinput.KeyHome)
	case 141:
		return controller.enterKey(uinput.KeyEnd)
	case 142:
		return controller.enterKey(uinput.KeyInsert)
	case 143:
		return controller.enterKey(uinput.KeyDelete)
	case 144:
		return controller.enterKey(uinput.KeyBackspace)
	case 145:
		return controller.enterKey(uinput.KeyEnter)
	case 146:
		return controller.enterKey(uinput.KeyPageup)
	case 147:
		return controller.enterKey(uinput.KeyPagedown)
	case 148:
		return controller.enterKey(uinput.KeyUp)
	case 149:
		return controller.enterKey(uinput.KeyDown)
	case 150:
		return controller.enterKey(uinput.KeyLeft)
	case 151:
		return controller.enterKey(uinput.KeyRight)
	case 152:
		return controller.enterKey(uinput.KeyCapslock)
	case 153:
		return controller.enterKey(uinput.KeyTab)
	case 161:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyA)
	case 163:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyL)
	case 166:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyS)
	case 172:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyX)
	case 175:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyZ)
	case 177:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyA)
	case 179:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyL)
	case 182:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyS)
	case 188:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyX)
	case 191:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyZ)
	case 198:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyC)
	case 202:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyE)
	case 209:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyN)
	case 211:
		return controller.enterTripleKey(uinput.KeyLeftshift, uinput.KeyRightalt, uinput.KeyO)
	case 230:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyC)
	case 234:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyE)
	case 241:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyN)
	case 243:
		return controller.enterDoubleKey(uinput.KeyRightalt, uinput.KeyO)
	}
	return nil
}

func (controller *KeyboardController) enterKey(key int) error {
	return controller.keyboard.KeyPress(key)
}

func (controller *KeyboardController) enterDoubleKey(firstKey int, secondKey int) error {
	err := controller.keyboard.KeyDown(firstKey)
	if err != nil {
		return err
	}
	err = controller.keyboard.KeyPress(secondKey)
	if err != nil {
		return err
	}
	return controller.keyboard.KeyUp(firstKey)
}

func (controller *KeyboardController) enterTripleKey(firstKey int, secondKey int, thirdKey int) error {
	err := controller.keyboard.KeyDown(firstKey)
	if err != nil {
		return err
	}
	err = controller.keyboard.KeyDown(secondKey)
	if err != nil {
		return err
	}
	err = controller.keyboard.KeyPress(thirdKey)
	if err != nil {
		return err
	}
	err = controller.keyboard.KeyUp(secondKey)
	if err != nil {
		return err
	}
	return controller.keyboard.KeyUp(firstKey)
}

func (controller *KeyboardController) Stop() error {
	return controller.keyboard.Close()
}
