package farm

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reinvest/core/farm/config"
	"reinvest/core/farm/mdex"
	"reinvest/core/farm/pancake"
	"reinvest/printer"
	"reinvest/token"
)

//var mdexProvider = wire.NewSet(pdf_service.New, wire.Bind(new(pdf.PdfCoreServer), new(*pdf_service.PdfServices)))

type Farm interface {
	Harvest() (*big.Int, error)
	SwapRewardToPairWithRetry(rewardAmount *big.Int, tryCount int) (*big.Int, *big.Int, string, string, error)
	AddLiquidityWithRetry(wishA *big.Int, wishB *big.Int, tokenAAddress, tokenBAddress string, tryCount int) (string, error)
	Reinvest() (*big.Int, string, error)
	LpToken() *token.Token
	RewardToken() *token.Token
	Start() error
}

func InitFarm(farmConfig *config.FarmConfig, client *ethclient.Client, tokenBasic *token.TokenBasic, printer *printer.Printer) (Farm, func(), error) {
	if farmConfig.NetWork.Name == "Mdex" {
		return mdex.NewMdexFarm(farmConfig, client, tokenBasic, printer), config.Pause, nil
	}
	if farmConfig.NetWork.Name == "Pancake" {
		return pancake.NewPanckeFarm(farmConfig, client, tokenBasic, printer), config.Pause, nil
	}

	return nil, config.Pause, errors.New("Unknown Network")
}
