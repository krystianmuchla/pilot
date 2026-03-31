package http

import (
	"pilot/internal/logic"
)

type KeyboardAdapter struct {
	keyboard logic.Keyboard
}

func NewKeyboardAdapter(keyboard logic.Keyboard) KeyboardAdapter {
	return KeyboardAdapter{keyboard: keyboard}
}

func (adapter KeyboardAdapter) EnterKey(keyCode byte) error {
	switch keyCode {
	case 32:
		return adapter.exec(adapter.keyboard.PressSpace, adapter.keyboard.ReleaseSpace)
	case 33:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press1,
			adapter.keyboard.Release1,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 34:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressApostrophe,
			adapter.keyboard.ReleaseApostrophe,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 35:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press3,
			adapter.keyboard.Release3,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 36:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press4,
			adapter.keyboard.Release4,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 37:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press5,
			adapter.keyboard.Release5,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 38:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press7,
			adapter.keyboard.Release7,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 39:
		return adapter.exec(adapter.keyboard.PressApostrophe, adapter.keyboard.ReleaseApostrophe)
	case 40:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press9,
			adapter.keyboard.Release9,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 41:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press0,
			adapter.keyboard.Release0,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 42:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press8,
			adapter.keyboard.Release8,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 43:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressEqual,
			adapter.keyboard.ReleaseEqual,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 44:
		return adapter.exec(adapter.keyboard.PressComma, adapter.keyboard.ReleaseComma)
	case 45:
		return adapter.exec(adapter.keyboard.PressMinus, adapter.keyboard.ReleaseMinus)
	case 46:
		return adapter.exec(adapter.keyboard.PressDot, adapter.keyboard.ReleaseDot)
	case 47:
		return adapter.exec(adapter.keyboard.PressSlash, adapter.keyboard.ReleaseSlash)
	case 48:
		return adapter.exec(adapter.keyboard.Press0, adapter.keyboard.Release0)
	case 49:
		return adapter.exec(adapter.keyboard.Press1, adapter.keyboard.Release1)
	case 50:
		return adapter.exec(adapter.keyboard.Press2, adapter.keyboard.Release2)
	case 51:
		return adapter.exec(adapter.keyboard.Press3, adapter.keyboard.Release3)
	case 52:
		return adapter.exec(adapter.keyboard.Press4, adapter.keyboard.Release4)
	case 53:
		return adapter.exec(adapter.keyboard.Press5, adapter.keyboard.Release5)
	case 54:
		return adapter.exec(adapter.keyboard.Press6, adapter.keyboard.Release6)
	case 55:
		return adapter.exec(adapter.keyboard.Press7, adapter.keyboard.Release7)
	case 56:
		return adapter.exec(adapter.keyboard.Press8, adapter.keyboard.Release8)
	case 57:
		return adapter.exec(adapter.keyboard.Press9, adapter.keyboard.Release9)
	case 58:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressSemicolon,
			adapter.keyboard.ReleaseSemicolon,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 59:
		return adapter.exec(adapter.keyboard.PressSemicolon, adapter.keyboard.ReleaseSemicolon)
	case 60:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressComma,
			adapter.keyboard.ReleaseComma,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 61:
		return adapter.exec(adapter.keyboard.PressEqual, adapter.keyboard.ReleaseEqual)
	case 62:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressDot,
			adapter.keyboard.ReleaseDot,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 63:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressSlash,
			adapter.keyboard.ReleaseSlash,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 64:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press2,
			adapter.keyboard.Release2,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 65:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressA,
			adapter.keyboard.ReleaseA,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 66:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressB,
			adapter.keyboard.ReleaseB,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 67:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressC,
			adapter.keyboard.ReleaseC,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 68:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressD,
			adapter.keyboard.ReleaseD,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 69:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressE,
			adapter.keyboard.ReleaseE,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 70:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressF,
			adapter.keyboard.ReleaseF,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 71:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressG,
			adapter.keyboard.ReleaseG,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 72:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressH,
			adapter.keyboard.ReleaseH,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 73:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressI,
			adapter.keyboard.ReleaseI,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 74:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressJ,
			adapter.keyboard.ReleaseJ,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 75:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressK,
			adapter.keyboard.ReleaseK,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 76:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressL,
			adapter.keyboard.ReleaseL,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 77:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressM,
			adapter.keyboard.ReleaseM,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 78:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressN,
			adapter.keyboard.ReleaseN,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 79:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressO,
			adapter.keyboard.ReleaseO,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 80:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressP,
			adapter.keyboard.ReleaseP,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 81:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressQ,
			adapter.keyboard.ReleaseQ,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 82:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressR,
			adapter.keyboard.ReleaseR,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 83:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressS,
			adapter.keyboard.ReleaseS,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 84:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressT,
			adapter.keyboard.ReleaseT,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 85:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressU,
			adapter.keyboard.ReleaseU,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 86:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressV,
			adapter.keyboard.ReleaseV,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 87:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressW,
			adapter.keyboard.ReleaseW,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 88:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressX,
			adapter.keyboard.ReleaseX,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 89:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressY,
			adapter.keyboard.ReleaseY,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 90:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressZ,
			adapter.keyboard.ReleaseZ,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 91:
		return adapter.exec(adapter.keyboard.PressLeftBrace, adapter.keyboard.ReleaseLeftBrace)
	case 92:
		return adapter.exec(adapter.keyboard.PressBackslash, adapter.keyboard.ReleaseBackslash)
	case 93:
		return adapter.exec(adapter.keyboard.PressRightBrace, adapter.keyboard.ReleaseRightBrace)
	case 94:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.Press6,
			adapter.keyboard.Release6,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 95:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressMinus,
			adapter.keyboard.ReleaseMinus,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 96:
		return adapter.exec(adapter.keyboard.PressGrave, adapter.keyboard.ReleaseGrave)
	case 97:
		return adapter.exec(adapter.keyboard.PressA, adapter.keyboard.ReleaseA)
	case 98:
		return adapter.exec(adapter.keyboard.PressB, adapter.keyboard.ReleaseB)
	case 99:
		return adapter.exec(adapter.keyboard.PressC, adapter.keyboard.ReleaseC)
	case 100:
		return adapter.exec(adapter.keyboard.PressD, adapter.keyboard.ReleaseD)
	case 101:
		return adapter.exec(adapter.keyboard.PressE, adapter.keyboard.ReleaseE)
	case 102:
		return adapter.exec(adapter.keyboard.PressF, adapter.keyboard.ReleaseF)
	case 103:
		return adapter.exec(adapter.keyboard.PressG, adapter.keyboard.ReleaseG)
	case 104:
		return adapter.exec(adapter.keyboard.PressH, adapter.keyboard.ReleaseH)
	case 105:
		return adapter.exec(adapter.keyboard.PressI, adapter.keyboard.ReleaseI)
	case 106:
		return adapter.exec(adapter.keyboard.PressJ, adapter.keyboard.ReleaseJ)
	case 107:
		return adapter.exec(adapter.keyboard.PressK, adapter.keyboard.ReleaseK)
	case 108:
		return adapter.exec(adapter.keyboard.PressL, adapter.keyboard.ReleaseL)
	case 109:
		return adapter.exec(adapter.keyboard.PressM, adapter.keyboard.ReleaseM)
	case 110:
		return adapter.exec(adapter.keyboard.PressN, adapter.keyboard.ReleaseN)
	case 111:
		return adapter.exec(adapter.keyboard.PressO, adapter.keyboard.ReleaseO)
	case 112:
		return adapter.exec(adapter.keyboard.PressP, adapter.keyboard.ReleaseP)
	case 113:
		return adapter.exec(adapter.keyboard.PressQ, adapter.keyboard.ReleaseQ)
	case 114:
		return adapter.exec(adapter.keyboard.PressR, adapter.keyboard.ReleaseR)
	case 115:
		return adapter.exec(adapter.keyboard.PressS, adapter.keyboard.ReleaseS)
	case 116:
		return adapter.exec(adapter.keyboard.PressT, adapter.keyboard.ReleaseT)
	case 117:
		return adapter.exec(adapter.keyboard.PressU, adapter.keyboard.ReleaseU)
	case 118:
		return adapter.exec(adapter.keyboard.PressV, adapter.keyboard.ReleaseV)
	case 119:
		return adapter.exec(adapter.keyboard.PressW, adapter.keyboard.ReleaseW)
	case 120:
		return adapter.exec(adapter.keyboard.PressX, adapter.keyboard.ReleaseX)
	case 121:
		return adapter.exec(adapter.keyboard.PressY, adapter.keyboard.ReleaseY)
	case 122:
		return adapter.exec(adapter.keyboard.PressZ, adapter.keyboard.ReleaseZ)
	case 123:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressLeftBrace,
			adapter.keyboard.ReleaseLeftBrace,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 124:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressBackslash,
			adapter.keyboard.ReleaseBackslash,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 125:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightBrace,
			adapter.keyboard.ReleaseRightBrace,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 126:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressGrave,
			adapter.keyboard.ReleaseGrave,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 127:
		return adapter.exec(adapter.keyboard.PressEsc, adapter.keyboard.ReleaseEsc)
	case 128:
		return adapter.exec(adapter.keyboard.PressF1, adapter.keyboard.ReleaseF1)
	case 129:
		return adapter.exec(adapter.keyboard.PressF2, adapter.keyboard.ReleaseF2)
	case 130:
		return adapter.exec(adapter.keyboard.PressF3, adapter.keyboard.ReleaseF3)
	case 131:
		return adapter.exec(adapter.keyboard.PressF4, adapter.keyboard.ReleaseF4)
	case 132:
		return adapter.exec(adapter.keyboard.PressF5, adapter.keyboard.ReleaseF5)
	case 133:
		return adapter.exec(adapter.keyboard.PressF6, adapter.keyboard.ReleaseF6)
	case 134:
		return adapter.exec(adapter.keyboard.PressF7, adapter.keyboard.ReleaseF7)
	case 135:
		return adapter.exec(adapter.keyboard.PressF8, adapter.keyboard.ReleaseF8)
	case 136:
		return adapter.exec(adapter.keyboard.PressF9, adapter.keyboard.ReleaseF9)
	case 137:
		return adapter.exec(adapter.keyboard.PressF10, adapter.keyboard.ReleaseF10)
	case 138:
		return adapter.exec(adapter.keyboard.PressF11, adapter.keyboard.ReleaseF11)
	case 139:
		return adapter.exec(adapter.keyboard.PressF12, adapter.keyboard.ReleaseF12)
	case 140:
		return adapter.exec(adapter.keyboard.PressHome, adapter.keyboard.ReleaseHome)
	case 141:
		return adapter.exec(adapter.keyboard.PressEnd, adapter.keyboard.ReleaseEnd)
	case 142:
		return adapter.exec(adapter.keyboard.PressInsert, adapter.keyboard.ReleaseInsert)
	case 143:
		return adapter.exec(adapter.keyboard.PressDelete, adapter.keyboard.ReleaseDelete)
	case 144:
		return adapter.exec(adapter.keyboard.PressBackspace, adapter.keyboard.ReleaseBackspace)
	case 145:
		return adapter.exec(adapter.keyboard.PressEnter, adapter.keyboard.ReleaseEnter)
	case 146:
		return adapter.exec(adapter.keyboard.PressPageUp, adapter.keyboard.ReleasePageUp)
	case 147:
		return adapter.exec(adapter.keyboard.PressPageDown, adapter.keyboard.ReleasePageDown)
	case 148:
		return adapter.exec(adapter.keyboard.PressUp, adapter.keyboard.ReleaseUp)
	case 149:
		return adapter.exec(adapter.keyboard.PressDown, adapter.keyboard.ReleaseDown)
	case 150:
		return adapter.exec(adapter.keyboard.PressLeft, adapter.keyboard.ReleaseLeft)
	case 151:
		return adapter.exec(adapter.keyboard.PressRight, adapter.keyboard.ReleaseRight)
	case 152:
		return adapter.exec(adapter.keyboard.PressCapslock, adapter.keyboard.ReleaseCapslock)
	case 153:
		return adapter.exec(adapter.keyboard.PressTab, adapter.keyboard.ReleaseTab)
	case 161:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressA,
			adapter.keyboard.ReleaseA,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 163:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressL,
			adapter.keyboard.ReleaseL,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 166:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressS,
			adapter.keyboard.ReleaseS,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 172:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressX,
			adapter.keyboard.ReleaseX,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 175:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressZ,
			adapter.keyboard.ReleaseZ,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 177:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressA,
			adapter.keyboard.ReleaseA,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 179:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressL,
			adapter.keyboard.ReleaseL,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 182:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressS,
			adapter.keyboard.ReleaseS,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 188:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressX,
			adapter.keyboard.ReleaseX,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 191:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressZ,
			adapter.keyboard.ReleaseZ,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 198:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressC,
			adapter.keyboard.ReleaseC,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 202:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressE,
			adapter.keyboard.ReleaseE,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 209:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressN,
			adapter.keyboard.ReleaseN,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 211:
		return adapter.exec(
			adapter.keyboard.PressLeftShift,
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressO,
			adapter.keyboard.ReleaseO,
			adapter.keyboard.ReleaseRightAlt,
			adapter.keyboard.ReleaseLeftShift,
		)
	case 230:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressC,
			adapter.keyboard.ReleaseC,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 234:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressE,
			adapter.keyboard.ReleaseE,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 241:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressN,
			adapter.keyboard.ReleaseN,
			adapter.keyboard.ReleaseRightAlt,
		)
	case 243:
		return adapter.exec(
			adapter.keyboard.PressRightAlt,
			adapter.keyboard.PressO,
			adapter.keyboard.ReleaseO,
			adapter.keyboard.ReleaseRightAlt,
		)
	}
	return nil
}

func (adapter KeyboardAdapter) exec(actions ...func() error) error {
	for _, action := range actions {
		if err := action(); err != nil {
			return err
		}
	}
	return nil
}
