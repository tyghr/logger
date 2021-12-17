package logger

import "fmt"

func Colored(color, text string) string {
	var iC uint8
	switch color {
	case "black":
		iC = 30
	case "red":
		iC = 31
	case "green":
		iC = 32
	case "yellow":
		iC = 33
	case "blue":
		iC = 34
	case "magenta":
		iC = 35
	case "cyan":
		iC = 36
	case "white":
		iC = 37
	case "r_black":
		iC = 40
	case "r_red":
		iC = 41
	case "r_green":
		iC = 42
	case "r_yellow":
		iC = 43
	case "r_blue":
		iC = 44
	case "r_magenta":
		iC = 45
	case "r_cyan":
		iC = 46
	case "r_white":
		iC = 47
	case "l_black":
		iC = 90
	case "l_red":
		iC = 91
	case "l_green":
		iC = 92
	case "l_yellow":
		iC = 93
	case "l_blue":
		iC = 94
	case "l_magenta":
		iC = 95
	case "l_cyan":
		iC = 96
	case "l_white":
		iC = 97
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(iC), text)
}
