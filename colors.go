package colors

import (
	"fmt"
	"log"
	"os"
)

//
// https://misc.flogisoft.com/bash/tip_colors_and_formatting
//
// "ESC  [  ${attr}; ${clbg}; ${clfg}m TEXT"
//

type ColorValue int
type formatValue string

const (
	formattingNormal    formatValue = "0"
	formattingBold      formatValue = "1"
	formattingUnderline formatValue = "4"
	formattingBlink     formatValue = "5"

	escape = "\033"
	reset  = escape + "[0m"

	ColorFG_Default      ColorValue = 39
	ColorFG_Black        ColorValue = 30
	ColorFG_Red          ColorValue = 31
	ColorFG_Green        ColorValue = 32
	ColorFG_Yellow       ColorValue = 33
	ColorFG_Blue         ColorValue = 34
	ColorFG_Magenta      ColorValue = 35
	ColorFG_Cyan         ColorValue = 36
	ColorFG_LightGray    ColorValue = 37
	ColorFG_DarkGray     ColorValue = 90
	ColorFG_LightRed     ColorValue = 91
	ColorFG_LightGreen   ColorValue = 92
	ColorFG_LightYellow  ColorValue = 93
	ColorFG_LightBlue    ColorValue = 94
	ColorFG_LightMagenta ColorValue = 95
	ColorFG_LightCyan    ColorValue = 96
	ColorFG_White        ColorValue = 97

	ColorBG_Default      ColorValue = 49
	ColorBG_Black        ColorValue = 40
	ColorBG_Red          ColorValue = 41
	ColorBG_Green        ColorValue = 42
	ColorBG_Yellow       ColorValue = 43
	ColorBG_Blue         ColorValue = 44
	ColorBG_Magenta      ColorValue = 45
	ColorBG_Cyan         ColorValue = 46
	ColorBG_LightGray    ColorValue = 47
	ColorBG_DarkGray     ColorValue = 100
	ColorBG_LightRed     ColorValue = 101
	ColorBG_LightGreen   ColorValue = 102
	ColorBG_LightYellow  ColorValue = 103
	ColorBG_LightBlue    ColorValue = 104
	ColorBG_LightMagenta ColorValue = 105
	ColorBG_LightCyan    ColorValue = 106
	ColorBG_White        ColorValue = 107
)

type Colors struct {
	Default      ColorValue
	Black        ColorValue
	Red          ColorValue
	Green        ColorValue
	Yellow       ColorValue
	Blue         ColorValue
	Magenta      ColorValue
	Cyan         ColorValue
	LightGray    ColorValue
	DarkGray     ColorValue
	LightRed     ColorValue
	LightGreen   ColorValue
	LightYellow  ColorValue
	LightBlue    ColorValue
	LightMagenta ColorValue
	LightCyan    ColorValue
	White        ColorValue
}

var Foreground = Colors{
	Default:      ColorFG_Default,
	Black:        ColorFG_Black,
	Red:          ColorFG_Red,
	Green:        ColorFG_Green,
	Yellow:       ColorFG_Yellow,
	Blue:         ColorFG_Blue,
	Magenta:      ColorFG_Magenta,
	Cyan:         ColorFG_Cyan,
	LightGray:    ColorFG_LightGray,
	DarkGray:     ColorFG_DarkGray,
	LightRed:     ColorFG_LightRed,
	LightGreen:   ColorFG_LightGreen,
	LightYellow:  ColorFG_LightYellow,
	LightBlue:    ColorFG_LightBlue,
	LightMagenta: ColorFG_LightMagenta,
	LightCyan:    ColorFG_LightCyan,
	White:        ColorFG_White,
}

var Background = Colors{
	Default:      ColorBG_Default,
	Black:        ColorBG_Black,
	Red:          ColorBG_Red,
	Green:        ColorBG_Green,
	Yellow:       ColorBG_Yellow,
	Blue:         ColorBG_Blue,
	Magenta:      ColorBG_Magenta,
	Cyan:         ColorBG_Cyan,
	LightGray:    ColorBG_LightGray,
	DarkGray:     ColorBG_DarkGray,
	LightRed:     ColorBG_LightRed,
	LightGreen:   ColorBG_LightGreen,
	LightYellow:  ColorBG_LightYellow,
	LightBlue:    ColorBG_LightBlue,
	LightMagenta: ColorBG_LightMagenta,
	LightCyan:    ColorBG_LightCyan,
	White:        ColorBG_White,
}

// ColoredText -
func ColoredText(text string, fg ColorValue) *StringExtender {
	return &StringExtender{
		format: formattingNormal,
		bg:     Background.Default,
		fg:     fg,
		text:   text,
	}
}

// ColoredTextWithBG -
func ColoredTextWithBG(text string, fg ColorValue, bg ColorValue) *StringExtender {
	return &StringExtender{
		format: formattingNormal,
		bg:     bg,
		fg:     fg,
		text:   text,
	}
}

// StringExtender -
type StringExtender struct {
	format formatValue
	bg     ColorValue
	fg     ColorValue
	text   string
}

// String - `Stringer` interface
func (thisRef StringExtender) String() string {
	dbug := fmt.Sprintf(
		"%s[%s;%d;%dm%s%s",
		escape,
		thisRef.format,
		thisRef.bg,
		thisRef.fg,
		thisRef.text,
		reset,
	)

	return dbug
}

// Bold -
func (thisRef *StringExtender) Bold() *StringExtender {
	thisRef.format = thisRef.format + ";" + formattingBold
	return thisRef
}

// Underline -
func (thisRef *StringExtender) Underline() *StringExtender {
	thisRef.format = thisRef.format + ";" + formattingUnderline
	return thisRef
}

// Blink -
func (thisRef *StringExtender) Blink() *StringExtender {
	thisRef.format = thisRef.format + ";" + formattingBlink
	return thisRef
}

func init() {
	log.SetOutput(NewColorable(os.Stdout))
	log.SetFlags(0)
}
