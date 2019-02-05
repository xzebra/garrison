package utils

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
)

func OutputCursor() {
	fmt.Fprintf(color.Output, "%s", cursorSign)
}

func OutputError(err error) {
	fmt.Fprintf(color.Output, "%s%s\n", errSign, err.Error())
}
