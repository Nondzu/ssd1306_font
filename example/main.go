package main

import (
	"machine"
	"time"

	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

func main() {

	time.Sleep(time.Millisecond * 100) // Please wait some time after turning on the device to properly initialize the display
	machine.I2C0.Configure(machine.I2CConfig{Frequency: 400000})

	// Display
	dev := ssd1306.NewI2C(machine.I2C0)
	dev.Configure(ssd1306.Config{Width: 128, Height: 64, Address: 0x3C, VccState: ssd1306.SWITCHCAPVCC})
	dev.ClearBuffer()
	dev.ClearDisplay()

	//font library init
	display := font.NewDisplay(dev)
	display.Configure(font.Config{FontType: font.FONT_7x10}) //set font here

	display.YPos = 20                 // set position Y
	display.XPos = 0                  // set position X
	display.PrintText("HELLO WORLD!") // print text

	for {

	}
}
