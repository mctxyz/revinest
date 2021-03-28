package printer

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
)

type Printer struct {
	Enable bool
}

func  NewPrinter() *Printer {
	sysType := runtime.GOOS

	if sysType == "windows" {
		return &Printer{Enable: false}

	}
	return &Printer{Enable: true}
}

var yellow = color.New(color.FgHiYellow).SprintFunc()

func (c *Printer) Warning(a ...interface{}) {
	if c.Enable {
		fmt.Println(yellow(a...))
		return
	}
	fmt.Println(a...)
}
