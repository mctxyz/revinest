package printer

import (
	"fmt"
	"github.com/fatih/color"
)

var blue = color.New(color.FgHiBlue).SprintFunc()

func (c *Printer) Info(a ...interface{}) {
	if c.Enable {
		fmt.Println(blue(a...))
		return
	}
	fmt.Println(a...)
}
