package printer

import (
	"fmt"
	"github.com/fatih/color"
)

var green = color.New(color.FgHiGreen).SprintFunc()

func (c *Printer) Success(a ...interface{}) {
	if c.Enable {
		fmt.Println(green(a...))
		return
	}
	fmt.Println(a...)
}
