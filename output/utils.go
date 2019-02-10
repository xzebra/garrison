package output

import "strings"

func Center(input string, width int) string {
	var output string
	dif := width - len(input)
	side := dif / 2
	if dif%2 != 0 {
		output += " "
	}
	output += strings.Repeat(" ", side)
	output += input
	output += strings.Repeat(" ", side)
	return output
}

func CenterRight(input string, width int) string {
	dif := width - len(input)
	return strings.Repeat(" ", dif) + input
}
