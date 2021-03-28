package config

import (
	"github.com/manifoldco/promptui"
)
type RPCURL string
type NetInfo struct {
	Router      string
	RewardToken string
	FarmAddress string
	Name        string
	RPC         string
}

var NET_MAPPING = map[int]*NetInfo{
	0: &NetInfo{
		Router:      "0xED7d5F38C79115ca12fe6C0041abb22F0A06C300",
		RewardToken: "0x25d2e80cb6b86881fd7e07dd263fb79f4abe033c",
		FarmAddress: "0xFB03e11D93632D97a8981158A632Dd5986F5E909",
		Name:        "Mdex",
		RPC:         "https://http-mainnet-node.huobichain.com",
	},
	1: &NetInfo{
		Router:      "0x05ff2b0db69458a0750badebc4f9e13add608c7f",
		RewardToken: "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
		FarmAddress: "0x73feaa1eE314F8c655E354234017bE2193C9E24E",
		Name:        "Pancake",
		RPC:         "https://bsc-dataseed1.binance.org",
	},
}

func NewNetWork() (*NetInfo, func(), error) {

	netWorks := make([]string, len(NET_MAPPING), len(NET_MAPPING))
	for id, netInfo := range NET_MAPPING {
		netWorks[id] = netInfo.Name
	}
	prompt := promptui.Select{
		Label: "Select Network",
		Items: netWorks,
	}
	number, _, err := prompt.Run()
	if err != nil {
		return nil, func() {

		}, err
	}
	if _, ok := NET_MAPPING[number]; ok {
		return NET_MAPPING[number], func() {
		}, nil
	}
	return nil, func() {
	}, nil
}
