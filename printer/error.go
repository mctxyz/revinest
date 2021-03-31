package printer

import (
	"fmt"
	"github.com/fatih/color"
)

var red = color.New(color.FgHiRed).SprintFunc()

func (c *Printer) Error(a ...interface{}) {
	if c.Enable {
		fmt.Println(red(a...))
		return
	}
	fmt.Println(a...)
}
