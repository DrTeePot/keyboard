// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package evdev

import (
	"github.com/jteeuwen/evdev"
	"github.com/jteeuwen/keyboard"
)

// Keyboard is a keyboard backend for libevdev.
type Keyboard struct {
	*keyboard.Base
	mods keyboard.Modifier // Active key modifiers.
}

// New constructs a new keyboard instance.
func New() keyboard.Keyboard {
	kb := new(Keyboard)
	kb.Base = keyboard.NewBase()
	return kb
}

func (k *Keyboard) testMod(code uint16, value int32, a, b uint16, mod keyboard.Modifier) bool {
	if code != a && code != b {
		return false
	}

	if value == 0 {
		k.mods &^= mod
	} else {
		k.mods |= mod
	}

	return true
}

func (kb *Keyboard) Poll(v interface{}) {
	evt, ok := v.(evdev.Event)
	if !ok || evt.Type != evdev.EvKeys {
		return // Not meant for us.
	}

	switch {
	case kb.testMod(evt.Code, evt.Value, evdev.KeyLeftAlt, evdev.KeyRightAlt, keyboard.ModAlt):
		return
	case kb.testMod(evt.Code, evt.Value, evdev.KeyLeftShift, evdev.KeyRightShift, keyboard.ModShift):
		return
	case kb.testMod(evt.Code, evt.Value, evdev.KeyLeftCtrl, evdev.KeyRightCtrl, keyboard.ModCtrl):
		return
	case kb.testMod(evt.Code, evt.Value, evdev.KeyLeftMeta, evdev.KeyRightMeta, keyboard.ModSuper):
		return
	}

	//var key keyboard.Key
	var k keyboard.Key

	switch evt.Code {
	case evdev.KeyA:
		k = keyboard.KeyA
	case evdev.KeyB:
		k = keyboard.KeyB
	case evdev.KeyC:
		k = keyboard.KeyC
	case evdev.KeyD:
		k = keyboard.KeyD
	case evdev.KeyE:
		k = keyboard.KeyE
	case evdev.KeyF:
		k = keyboard.KeyF
	case evdev.KeyG:
		k = keyboard.KeyG
	case evdev.KeyH:
		k = keyboard.KeyH
	case evdev.KeyI:
		k = keyboard.KeyI
	case evdev.KeyJ:
		k = keyboard.KeyJ
	case evdev.KeyK:
		k = keyboard.KeyK
	case evdev.KeyL:
		k = keyboard.KeyL
	case evdev.KeyM:
		k = keyboard.KeyM
	case evdev.KeyN:
		k = keyboard.KeyN
	case evdev.KeyO:
		k = keyboard.KeyO
	case evdev.KeyP:
		k = keyboard.KeyP
	case evdev.KeyQ:
		k = keyboard.KeyQ
	case evdev.KeyR:
		k = keyboard.KeyR
	case evdev.KeyS:
		k = keyboard.KeyS
	case evdev.KeyT:
		k = keyboard.KeyT
	case evdev.KeyU:
		k = keyboard.KeyU
	case evdev.KeyV:
		k = keyboard.KeyV
	case evdev.KeyW:
		k = keyboard.KeyW
	case evdev.KeyX:
		k = keyboard.KeyX
	case evdev.KeyY:
		k = keyboard.KeyY
	case evdev.KeyZ:
		k = keyboard.KeyZ
	case evdev.Key0:
		k = keyboard.Key0
	case evdev.Key1:
		k = keyboard.Key1
	case evdev.Key2:
		k = keyboard.Key2
	case evdev.Key3:
		k = keyboard.Key3
	case evdev.Key4:
		k = keyboard.Key4
	case evdev.Key5:
		k = keyboard.Key5
	case evdev.Key6:
		k = keyboard.Key6
	case evdev.Key7:
		k = keyboard.Key7
	case evdev.Key8:
		k = keyboard.Key8
	case evdev.Key9:
		k = keyboard.Key9
	case evdev.KeyGrave:
		k = keyboard.KeyGraveAccent
	case evdev.KeyComma:
		k = keyboard.KeyComma
	case evdev.KeyDot:
		k = keyboard.KeyPeriod
	case evdev.KeySlash:
		k = keyboard.KeySlash
	case evdev.KeySemiColon:
		k = keyboard.KeySemicolon
	case evdev.KeyApostrophe:
		k = keyboard.KeyApostrophe
	case evdev.KeyMinus:
		k = keyboard.KeyMinus
	case evdev.KeyEqual:
		k = keyboard.KeyEqual
	case evdev.KeyLeftBrace:
		k = keyboard.KeyLeftBracket
	case evdev.KeyRightBrace:
		k = keyboard.KeyRightBracket
	case evdev.KeyBackSlash:
		k = keyboard.KeyBackslash

	case evdev.KeySpace:
		k = keyboard.KeySpace
	case evdev.KeyEscape:
		k = keyboard.KeyEscape
	case evdev.KeyF1:
		k = keyboard.KeyF1
	case evdev.KeyF2:
		k = keyboard.KeyF2
	case evdev.KeyF3:
		k = keyboard.KeyF3
	case evdev.KeyF4:
		k = keyboard.KeyF4
	case evdev.KeyF5:
		k = keyboard.KeyF5
	case evdev.KeyF6:
		k = keyboard.KeyF6
	case evdev.KeyF7:
		k = keyboard.KeyF7
	case evdev.KeyF8:
		k = keyboard.KeyF8
	case evdev.KeyF9:
		k = keyboard.KeyF9
	case evdev.KeyF10:
		k = keyboard.KeyF10
	case evdev.KeyF11:
		k = keyboard.KeyF11
	case evdev.KeyF12:
		k = keyboard.KeyF12
	case evdev.KeyF13:
		k = keyboard.KeyF13
	case evdev.KeyF14:
		k = keyboard.KeyF14
	case evdev.KeyF15:
		k = keyboard.KeyF15
	case evdev.KeyF16:
		k = keyboard.KeyF16
	case evdev.KeyF17:
		k = keyboard.KeyF17
	case evdev.KeyF18:
		k = keyboard.KeyF18
	case evdev.KeyF19:
		k = keyboard.KeyF19
	case evdev.KeyF20:
		k = keyboard.KeyF20
	case evdev.KeyF21:
		k = keyboard.KeyF21
	case evdev.KeyF22:
		k = keyboard.KeyF22
	case evdev.KeyF23:
		k = keyboard.KeyF23
	case evdev.KeyF24:
		k = keyboard.KeyF24
	case evdev.KeyUp:
		k = keyboard.KeyUp
	case evdev.KeyDown:
		k = keyboard.KeyDown
	case evdev.KeyLeft:
		k = keyboard.KeyLeft
	case evdev.KeyRight:
		k = keyboard.KeyRight
	case evdev.KeyTab:
		k = keyboard.KeyTab
	case evdev.KeyEnter:
		k = keyboard.KeyEnter
	case evdev.KeyBackSpace:
		k = keyboard.KeyBackspace
	case evdev.KeyInsert:
		k = keyboard.KeyInsert
	case evdev.KeyDelete:
		k = keyboard.KeyDelete
	case evdev.KeyPageUp:
		k = keyboard.KeyPageUp
	case evdev.KeyPageDown:
		k = keyboard.KeyPageDown
	case evdev.KeyHome:
		k = keyboard.KeyHome
	case evdev.KeyEnd:
		k = keyboard.KeyEnd
	case evdev.KeyKP0:
		k = keyboard.KeyKp0
	case evdev.KeyKP1:
		k = keyboard.KeyKp1
	case evdev.KeyKP2:
		k = keyboard.KeyKp2
	case evdev.KeyKP3:
		k = keyboard.KeyKp3
	case evdev.KeyKP4:
		k = keyboard.KeyKp4
	case evdev.KeyKP5:
		k = keyboard.KeyKp5
	case evdev.KeyKP6:
		k = keyboard.KeyKp6
	case evdev.KeyKP7:
		k = keyboard.KeyKp7
	case evdev.KeyKP8:
		k = keyboard.KeyKp8
	case evdev.KeyKP9:
		k = keyboard.KeyKp9
	case evdev.KeyKPSlash:
		k = keyboard.KeyKpDivide
	case evdev.KeyKPAsterisk:
		k = keyboard.KeyKpMultiply
	case evdev.KeyKPMinus:
		k = keyboard.KeyKpSubtract
	case evdev.KeyKPPlus:
		k = keyboard.KeyKpAdd
	case evdev.KeyKPDot:
		k = keyboard.KeyKpDecimal
	case evdev.KeyKPEqual:
		k = keyboard.KeyKpEqual
	case evdev.KeyKPEnter:
		k = keyboard.KeyKpEnter
	case evdev.KeyNumLock:
		k = keyboard.KeyNumLock
	case evdev.KeyCapsLock:
		k = keyboard.KeyCapsLock
	case evdev.KeyScrollLock:
		k = keyboard.KeyScrollLock
	case evdev.KeyPause:
		k = keyboard.KeyPause
	case evdev.KeyMenu:
		k = keyboard.KeyMenu
	default:
		k = keyboard.KeyUnknown
	}

	if evt.Value == 1 {
		kb.RecordKey(k, kb.mods)
	}
}
