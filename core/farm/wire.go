// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package farm

import (
	"github.com/google/wire"
	"reinvest/core"
	"reinvest/core/farm/config"
	"reinvest/printer"
	"reinvest/token"
)

func NewFarm() (Farm, func(), error) {
	panic(wire.Build(InitFarm, config.NewFarmConfig, core.Connect, config.NewPool, config.NewPrivateKey, config.NewNetWork, token.NewTokenBasic, printer.NewPrinter))
}
