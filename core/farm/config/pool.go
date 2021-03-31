package config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
)

func NewPool() (poolID PoolID, cf func(), err error) {

	poolIdValidate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("Invalid Pool ID")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Pool ID",
		Validate: poolIdValidate,
	}
	var result string
	result, err = prompt.Run()
	if err != nil {

		return
	}
	var number int
	number, err = strconv.Atoi(result)
	if err != nil {
		return
	}
	poolID = PoolID(number)
	cf = func(){

	}
	return
}

func Pause() {
	fmt.Print("Press Any Key to Exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	return
}
