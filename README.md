# ssd1306 font

### This is  lite ssd1306 font library using TinyGo. Tested on Raspberry Pi Pico.


Example:
```
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

```




#### FONT_6x8 
![FONT_6x8](https://user-images.githubusercontent.com/20000097/186149922-b45a203e-44cd-478e-97d9-5f99e885c639.png)

#### FONT_7x10
![FONT_7x10](https://user-images.githubusercontent.com/20000097/186149937-4c95dc76-6f37-4675-b36c-3ef3fbb3ea5c.png)

#### FONT_11x18
![FONT_11x18](https://user-images.githubusercontent.com/20000097/186149933-1420af08-3221-4cf7-9ee4-13f2e20f6c41.png)

#### FONT_16x26
![FONT_16x26](https://user-images.githubusercontent.com/20000097/186149928-0c716beb-da9b-4dd5-b72f-5d681adaa802.png)

