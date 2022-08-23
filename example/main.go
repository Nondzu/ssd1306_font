package main

import (
	"machine"
	"time"

	"github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

func main() {

	time.Sleep(time.Millisecond * 100)
	gp1 := machine.GP15
	gp1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.I2C0.Configure(machine.I2CConfig{Frequency: 400000})

	// Display
	dev := ssd1306.NewI2C(machine.I2C0)

	dev.Configure(ssd1306.Config{Width: 128, Height: 64, Address: 0x3C, VccState: ssd1306.SWITCHCAPVCC})

	dev.ClearBuffer()
	dev.ClearDisplay()
	display := ssd1306_font.NewDisplay(dev)

	display.Configure(ssd1306_font.Config{FontType: ssd1306_font.FONT_11x18})

	display.YPos = 20
	display.XPos = 2
	display.PrintText("HELLO WORLD :)")

	for {

	}
}
