package alpaca

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reinvest/core/farm/alpaca/contracts"
)

type SwapFactory struct {
	Client       *ethclient.Client
	SwapContract *contracts.AplpacaFactory
	Farm         *AlpacaFarm
	Address      string
}

func NewSwapFactory(address string, client *ethclient.Client, farm *AlpacaFarm) (*SwapFactory, error) {
	swapFactory, err := contracts.NewAplpacaFactory(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}
	return &SwapFactory{
		Address:      address,
		Client:       client,
		SwapContract: swapFactory,
		Farm:         farm,
	}, nil

}

func (c *SwapFactory) Calc(amountB *big.Int, slippage float64) *big.Int {
	minAmount := new(big.Float).SetInt(amountB)

	min := new(big.Float).Mul(minAmount, big.NewFloat(1.00-slippage))
	a := big.NewInt(0)
	res, _ := min.Int(a)
	return res
}
