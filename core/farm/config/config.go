package config

import "reinvest/utils"

type FarmConfig struct {
	NetWork    *NetInfo
	PooID      PoolID
	PrivateKey string
	Wallet     string
}
type NetWorkID int
type PoolID int

func NewFarmConfig(netInfo *NetInfo, poolID PoolID, privateKey string) (*FarmConfig, func(), error) {

	return &FarmConfig{
		NetWork:    netInfo,
		PooID:      poolID,
		PrivateKey: privateKey,
		Wallet:     utils.WalletAddress(privateKey),
	}, Pause, nil
}
