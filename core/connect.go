package core

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"reinvest/core/farm/config"
)

func Connect(cfg *config.NetInfo) (client *ethclient.Client, cf func(), err error) {
	client, err = ethclient.Dial(cfg.RPC)
	cf = func() { client.Close() }
	return
}
