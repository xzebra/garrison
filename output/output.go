package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	errSign    = color.RedString("[!] ")
	infoSign   = color.YellowString("[?] ")
	addSign    = color.GreenString("[+] ")
	removeSign = color.RedString("[-] ")
	cursorSign = color.BlueString(":: ")

	StatusOFF = color.RedString("OFF")
	StatusON  = color.RedString("ON")
)

func PrintTitle() {
	fmt.Println()
	fmt.Println(" ▄▄ •  ▄▄▄· ▄▄▄  ▄▄▄  ▪  .▄▄ ·        ▐  ▄")
	fmt.Println("▐█ ▀ ▪▐█ ▀█ ▀▄ █·▀▄ █·██ ▐█ ▀. ▪     •█▌▐█")
	fmt.Println("▄█ ▀█▄▄█▀▀█ ▐▀▀▄ ▐▀▀▄ ▐█·▄▀▀▀█▄ ▄█▀▄ ▐█▐▐▌")
	fmt.Println("▐█▄▪▐█▐█ ▪▐▌▐█•█▌▐█•█▌▐█▌▐█▄▪▐█▐█▌.▐▌██▐█▌")
	fmt.Println("·▀▀▀▀  ▀  ▀ .▀  ▀.▀  ▀▀▀▀ ▀▀▀▀  ▀█▄▀▪▀▀ █▪")
	fmt.Println()
}

func Cursor() {
	fmt.Fprintf(color.Output, "%s", cursorSign)
}

func Error(err error) {
	fmt.Fprintf(color.Output, "%s%s\n", errSign, err.Error())
}

func Added(str string) {
	fmt.Fprintf(color.Output, "%s%s\n", addSign, str)
}

func Removed(str string) {
	fmt.Fprintf(color.Output, "%s%s\n", removeSign, str)
}

func Info(str string) {
	fmt.Fprintf(color.Output, "%s%s\n", infoSign, str)
}
