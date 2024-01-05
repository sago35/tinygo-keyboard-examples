package main

import (
	"context"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	d := keyboard.New()

	gpioPins := []machine.Pin{
		machine.WIO_KEY_A,
		machine.WIO_KEY_B,
		machine.WIO_KEY_C,
	}

	for c := range gpioPins {
		gpioPins[c].Configure(machine.PinConfig{Mode: machine.PinInput})
	}

	d.AddGpioKeyboard(gpioPins, [][]keyboard.Keycode{
		{
			// layer 0
			jp.KeyA,
			jp.KeyB,
			jp.KeyMod1,
		},
		{
			// layer 1
			jp.Key1,
			jp.KeyBackspace,
			jp.KeyMod1,
		},
	})

	d.Loop(context.Background())
}
