package pancake

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Swap struct {
	Client *ethclient.Client
	Farm   *PancakeFarm
}

//func NewSwap(client *ethclient.Client, SwapTokenAddress string, farm *PancakeFarm) (*Swap, error) {
//	swapContract, err := contracts.NewSwapMining(common.HexToAddress(SwapTokenAddress), client)
//	if err != nil {
//		return nil, err
//	}
//	return &Swap{Client: client, SwapContract: swapContract, Farm: farm}, nil
//}
//
//func (c *Swap) SwapBalance(total string) (string, string) {
//
//	totalAmount, _ := new(big.Float).SetString(total)
//	remain := new(big.Float).Quo(totalAmount, big.NewFloat(2))
//	return remain.String(), totalAmount.Sub(totalAmount, remain).String()
//
//}
func (c *Swap) MinExchange(amount, proportion string) (string, error) {

	needExchange, _ := new(big.Float).SetString(amount)
	proportionBig, _ := new(big.Float).SetString(proportion)

	return new(big.Float).Mul(needExchange, proportionBig).String(), nil
}
