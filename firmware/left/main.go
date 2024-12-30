package main

import (
	"context"
	_ "embed"
	"log"
	"machine"
	"machine/usb"

	keyboard "github.com/sago35/tinygo-keyboard"
	jp "github.com/sago35/tinygo-keyboard/keycodes/japanese"
)

func main() {
	usb.Product = "Lunakey Pico tinygo-keyboard 0.1.0"
	usb.VendorID = 0x5954
	usb.ProductID = 0x0003

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	d := keyboard.New()

	colPins := []machine.Pin{
		machine.GPIO21,
		machine.GPIO20,
		machine.GPIO19,
		machine.GPIO18,
		machine.GPIO17,
		machine.GPIO16,
	}

	rowPins := []machine.Pin{
		machine.GPIO12,
		machine.GPIO13,
		machine.GPIO14,
		machine.GPIO15,
	}

	d.AddMatrixKeyboard(colPins, rowPins, [][]keyboard.Keycode{
		{
			jp.KeyTab, jp.KeyQ, jp.KeyW, jp.KeyE, jp.KeyR, jp.KeyT,
			jp.KeyLeftCtrl, jp.KeyA, jp.KeyS, jp.KeyD, jp.KeyF, jp.KeyG,
			jp.KeyLeftShift, jp.KeyZ, jp.KeyX, jp.KeyC, jp.KeyV, jp.KeyB,
			0, 0, jp.KeyEsc, jp.KeyLeftAlt, jp.KeyMod1, jp.KeySpace,
		},
		{
			jp.KeyTab, jp.KeyQ, jp.KeyW, jp.KeyE, jp.KeyR, jp.KeyT,
			jp.KeyLeftCtrl, jp.KeyA, jp.KeyS, jp.KeyD, jp.KeyF, jp.KeyG,
			jp.KeyLeftShift, jp.KeyZ, jp.KeyX, jp.KeyC, jp.KeyV, jp.KeyB,
			0, 0, jp.KeyEsc, jp.KeyLeftAlt, jp.KeyMod1, jp.KeySpace,
		},
	})

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{TX: machine.GPIO0, RX: machine.GPIO1})
	d.AddUartKeyboard(24, uart, [][]keyboard.Keycode{
		{
			jp.KeyY, jp.KeyU, jp.KeyI, jp.KeyO, jp.KeyP, jp.KeyAt,
			jp.KeyH, jp.KeyJ, jp.KeyK, jp.KeyL, jp.KeySemicolon, jp.KeyColon,
			jp.KeyN, jp.KeyM, jp.KeyComma, jp.KeyPeriod, jp.KeySlash, jp.KeyBackslash,
			jp.KeySpace, jp.KeyMod2, jp.KeyHiragana, jp.KeyTo1, 0, 0,
		},
		{
			jp.KeyY, jp.KeyU, jp.KeyI, jp.KeyO, jp.KeyP, jp.KeyAt,
			jp.KeyH, jp.KeyJ, jp.KeyK, jp.KeyL, jp.KeySemicolon, jp.KeyColon,
			jp.KeyN, jp.KeyM, jp.KeyComma, jp.KeyPeriod, jp.KeySlash, jp.KeyBackslash,
			jp.KeySpace, jp.KeyMod2, jp.KeyHiragana, jp.KeyTo1, 0, 0,
		},
	})

	// override ctrl-h to BackSpace
	//d.OverrideCtrlH()

	loadKeyboardDef()

	return d.Loop(context.Background())
}
