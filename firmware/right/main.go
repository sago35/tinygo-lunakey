package main

import (
	"context"
	"log"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
)

func main() {
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
			0x0000, 0x0001, 0x0002, 0x0003, 0x0004, 0x0005,
			0x0006, 0x0007, 0x0008, 0x0009, 0x000A, 0x000B,
			0x000C, 0x000D, 0x000E, 0x000F, 0x0010, 0x0011,
			0x0012, 0x0013, 0x0014, 0x0015, 0x0016, 0x0017,
		},
	})

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{TX: machine.GPIO0, RX: machine.GPIO1})
	d.Keyboard = &keyboard.UartTxKeyboard{
		Uart: uart,
	}

	return d.Loop(context.Background())
}
