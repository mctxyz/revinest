package alpaca

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"reinvest/core/farm/alpaca/contracts"
)

type LpToken struct {
	Client   *ethclient.Client
	Contract *contracts.AplpacaLpToken
	Address  string
}

func NewLpToken(address string, client *ethclient.Client) (*LpToken, error) {
	swapLpTokenContract, err := contracts.NewAplpacaLpToken(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}
	return &LpToken{
		Client:   client,
		Contract: swapLpTokenContract,
		Address:  address,
	}, nil

}

func (c *LpToken) Pair() (string, string, error) {
	address0, err := c.Contract.Token0(&bind.CallOpts{})
	if err != nil {
		return "", "", err
	}
	address1, err := c.Contract.Token1(&bind.CallOpts{})
	if err != nil {
		return "", "", err
	}
	return address0.String(), address1.String(), nil
}

func (c *LpToken) GetReserves() (*big.Int, *big.Int, error) {
	reserves, err := c.Contract.GetReserves(&bind.CallOpts{})
	if err != nil {
		return nil, nil, err
	}
	return reserves.Reserve0, reserves.Reserve1, nil
}
