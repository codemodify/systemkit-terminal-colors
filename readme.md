# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Colors in Terminals, X-Platform
[![](https://img.shields.io/github/v/release/codemodify/systemkit-terminal-colors?style=flat-square)](https://github.com/codemodify/systemkit-terminal-colors/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-terminal-colors?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-terminal-colors?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-terminal-colors/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-terminal-colors?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-terminal-colors?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-terminal-colors)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-terminal-colors)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-terminal-colors?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-terminal-colors?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-terminal-colors?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-terminal-colors?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-terminal-colors?style=flat-square)

### What it offers
- Simple, robust ANSI colored text in terminals
- Yes it works on Windows

### How does it differ compared to
- github.com/logrusorgru/aurora
	- too complicated
	- does not work on Windows

- github.com/fatih/color
	- discontinued

### Credits
- github.com/mattn/go-colorable - for kick-ass fixed colors on Windows without a ton of third party dependencies

### Samples
```go
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
```

![](https://raw.githubusercontent.com/codemodify/systemkit-terminal-colors/master/.helper-files/dox/colors-unix.gif)
![](https://raw.githubusercontent.com/codemodify/systemkit-terminal-colors/master/.helper-files/dox/colors-windows.png)
