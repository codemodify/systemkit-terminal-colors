// +build windows

package main

import (
	"log"

	colors "github.com/codemodify/systemkit-terminal-colors"
)

func main() {
	log.Println()
	log.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	log.Println()

	log.Println(colors.ColoredText("Test - 01 - Foreground.Default", colors.Foreground.Default))
	log.Println(colors.ColoredText("Test - 01 - Foreground.Black", colors.Foreground.Black).String() + "<- BLACK here")
	log.Println(colors.ColoredText("Test - 01 - Foreground.Red", colors.Foreground.Red))
	log.Println(colors.ColoredText("Test - 01 - Foreground.Green", colors.Foreground.Green))
	log.Println(colors.ColoredText("Test - 01 - Foreground.Yellow", colors.Foreground.Yellow))
	log.Println(colors.ColoredText("Test - 01 - Foreground.Blue", colors.Foreground.Blue))
	log.Println(colors.ColoredText("Test - 01 - Foreground.Magenta", colors.Foreground.Magenta))
	log.Println(colors.ColoredText("Test - 01 - Foreground.Cyan", colors.Foreground.Cyan))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightGray", colors.Foreground.LightGray))
	log.Println(colors.ColoredText("Test - 01 - Foreground.DarkGray", colors.Foreground.DarkGray))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightRed", colors.Foreground.LightRed))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightGreen", colors.Foreground.LightGreen))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightYellow", colors.Foreground.LightYellow))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightBlue", colors.Foreground.LightBlue))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightMagenta", colors.Foreground.LightMagenta))
	log.Println(colors.ColoredText("Test - 01 - Foreground.LightCyan", colors.Foreground.LightCyan))
	log.Println(colors.ColoredText("Test - 01 - Foreground.White", colors.Foreground.White))

	log.Println()
	log.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	log.Println()

	log.Println(colors.ColoredText("Test - 02  RED", colors.Foreground.Red))
	log.Println(colors.ColoredText("Test - 02  RED - BOLD", colors.Foreground.Red).Bold())
	log.Println(colors.ColoredText("Test - 02  RED - BOLD - BLINK", colors.Foreground.Red).Bold().Blink())
	log.Println(colors.ColoredText("Test - 02  RED - BOLD - BLINK - UNDERLINE", colors.Foreground.Red).Bold().Blink().Underline())

	log.Println()
	log.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	log.Println()

	log.Println(colors.ColoredTextWithBG("Test - 03 - BLACK on YELLOW", colors.Foreground.Black, colors.Background.LightYellow))

	log.Println()
	log.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	log.Println()
}
