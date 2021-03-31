package mdex_boardroom_wht

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reinvest/core/farm/config"
	"reinvest/printer"
	"reinvest/token"
)

type MdexFarm struct {
	Printer         *printer.Printer
	FarmConfig      *config.FarmConfig
	Client          *ethclient.Client
	FarmInfo        *PoolInfo
	TokenBasic      *token.TokenBasic
	TokenAInfo      *token.Token
	TokenBInfo      *token.Token
	RewardTokenInfo *token.Token
	LpTokenInfo     *token.Token
}

type PendingReward struct {
	Amount      *big.Int
	TokenAmount *big.Int
}

func NewMedxBoardRoomFarmWHT(farmConfig *config.FarmConfig, client *ethclient.Client, tokenBasic *token.TokenBasic, printer *printer.Printer) *MdexFarm {
	return &MdexFarm{
		FarmConfig: farmConfig,
		Client:     client,
		TokenBasic: tokenBasic,
		Printer:    printer,
	}
}

