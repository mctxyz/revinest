package core

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"reinvest/core/farm/config"
)

func Connect(rpcURL *config.NetInfo) (client *ethclient.Client, cf func(), err error) {
	client, err = ethclient.Dial(rpcURL.RPC)
	cf = func() { client.Close() }
	if err != nil {
		return
	}
	return
}
