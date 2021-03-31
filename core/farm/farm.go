package farm

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reinvest/core/farm/config"
	"reinvest/core/farm/mdex"
	"reinvest/core/farm/mdex_boardroom_mdx"
	"reinvest/core/farm/mdex_boardroom_wht"
	"reinvest/core/farm/pancake"
	"reinvest/printer"
	"reinvest/token"
)

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
	if farmConfig.NetWork.Name == "Mdex-BoardRoom-WHT" {
		return mdex_boardroom_wht.NewMedxBoardRoomFarmWHT(farmConfig, client, tokenBasic, printer), config.Pause, nil
	}
	if farmConfig.NetWork.Name == "Mdex-BoardRoom-MDX" {
		return mdex_boardroom_mdx.NewMedxBoardRoomFarmMDX(farmConfig, client, tokenBasic, printer), config.Pause, nil
	}
	if farmConfig.NetWork.Name == "Pancake" {
		return pancake.NewPanckeFarm(farmConfig, client, tokenBasic, printer), config.Pause, nil
	}

	return nil, config.Pause, errors.New("Unknown Network")
}
