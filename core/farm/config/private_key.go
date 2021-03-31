package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"reinvest/utils"
	"strings"
)

func NewPrivateKey() (privateKey string, cf func(), err error) {
	cf = func(){

	}
	privateValidate := func(input string) error {
		input = strings.Replace(input, " ", "", -1)
		input = strings.Replace(input, "\n", "", -1)
		wallet, err := utils.CheckPrivateKey(input)
		if err != nil {
			return fmt.Errorf("Invaild Private Key ")

		}
		if wallet == "" {
			return errors.New("Invaild Private Key ")
		}
		return nil
	}
	pk := promptui.Prompt{
		Label:    "Private Key",
		Validate: privateValidate,
		Mask:     '*',
	}
	privateKey, err = pk.Run()

	return
}
