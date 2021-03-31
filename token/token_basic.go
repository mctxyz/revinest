package token

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"reinvest/core/farm/config"
	"reinvest/utils"
	"strings"
	"time"
)

type MyTokenInfo struct {
	Balance  *big.Int
	Name     string
	Decimals uint8
	Symbol   string
}
type Token struct {
	TokenContract *Hrc20
	Name          string
	Decimals      uint8
	Symbol        string
	Address       string
}

type TokenBasic struct {
	Client     *ethclient.Client
	Wallet     string
	PrivateKey string
}

func NewTokenBasic(farmConfig *config.FarmConfig, client *ethclient.Client) (*TokenBasic, func(), error) {
	return &TokenBasic{
		Client:     client,
		Wallet:     farmConfig.Wallet,
		PrivateKey: farmConfig.PrivateKey,
	}, func() {}, nil
}

var GasUsed = big.NewInt(0)

func (c *TokenBasic) GetMyTokenInfo(tokenAddress string) (*MyTokenInfo, error) {
	walletAddress := c.Wallet
	if !utils.IsValidAddress(tokenAddress) {
		return nil, errors.New("Token Address Is InValid!")
	}
	if !utils.IsValidAddress(walletAddress) {
		return nil, errors.New("Wallet Address Is InValid!")
	}
	token, err := c.TokenInfo(tokenAddress)
	if err != nil {

		return nil, errors.New("Get Token Error")
	}
	balance, err := token.TokenContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(walletAddress))
	if err != nil {
		return nil, errors.New("Get Balance Error")
	}

	return &MyTokenInfo{
		Balance:  balance,
		Name:     token.Name,
		Decimals: token.Decimals,
		Symbol:   token.Symbol,
	}, nil
}

func (c TokenBasic) TokenInfo(tokenAddress string) (*Token, error) {
	if !utils.IsValidAddress(tokenAddress) {
		return nil, errors.New("Token Address Is InValid!")
	}
	tokenContract, err := NewHrc20(common.HexToAddress(tokenAddress), c.Client)
	if err != nil {

		return nil, errors.New("Get Token Error")

	}
	decimals, err := tokenContract.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.New("Get Decimals Error")
	}
	tokenName, err := tokenContract.Name(&bind.CallOpts{})
	if err != nil {
		return nil, errors.New("Get Token Name Error")
	}
	symbol, err := tokenContract.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, errors.New("Get Token Name Error")
	}
	return &Token{
		TokenContract: tokenContract,
		Name:          tokenName,
		Decimals:      decimals,
		Symbol:        symbol,
		Address:       tokenAddress,
	}, nil
}
func (c *TokenBasic) CreateTx() (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, err
	}

	chainID, err := c.Client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(privateKey, chainID)
}
func (c *TokenBasic) Approve(tokenAddress string, contractAddress string, amount *big.Int) (bool, error) {
	wallet := c.Wallet
	hrc20, err := NewHrc20(common.HexToAddress(tokenAddress), c.Client)
	if err != nil {
		return false, err
	}
	allowance, err := hrc20.Allowance(&bind.CallOpts{}, common.HexToAddress(wallet), common.HexToAddress(contractAddress))
	if err != nil {
		return false, err
	}
	if allowance.Cmp(amount) >= 1 {

		return true, nil
	}
	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	ABI, err := abi.JSON(strings.NewReader(Hrc20ABI))
	if err != nil {
		return false, err
	}
	unlimit := new(big.Int).Exp(big.NewInt(2), big.NewInt(250), nil)
	txData, err := ABI.Pack("approve", common.HexToAddress(contractAddress), unlimit.Sub(unlimit, big.NewInt(1)))
	if err != nil {
		return false, err
	}
	toContract := common.HexToAddress(tokenAddress)
	gas, err := c.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     common.HexToAddress(c.Wallet),
		To:       &toContract,
		GasPrice: gasPrice.Mul(gasPrice, big.NewInt(2)),
		Value:    utils.ToWei(big.NewFloat(0.00), 18),
		Data:     txData,
	})
	nonce, err := c.Client.PendingNonceAt(context.Background(), common.HexToAddress(wallet))
	if err != nil {
		return false, fmt.Errorf("Get Nonce Error %w", err)
	}
	auth, err := c.CreateTx()
	if err != nil {
		return false, err
	}
	auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2))
	auth.From = common.HexToAddress(c.Wallet)
	auth.GasLimit = gas * 2
	auth.Context = context.Background()
	auth.Nonce = big.NewInt(int64(nonce))

	tx, err := hrc20.Approve(auth, common.HexToAddress(contractAddress), unlimit.Sub(unlimit, big.NewInt(1)))
	if err != nil {
		return false, err
	}
	fmt.Println("Approve Tx: " + tx.Hash().String())
	txStatus, _ := c.WaitForBlockCompletation(tx.Hash().String())
	if txStatus == 1 {
		return true, nil
	}
	return false, errors.New("Approve Faild")
}

func (c *TokenBasic) WaitForBlockCompletation(hash string) (int, *types.Receipt) {
	log.Println("Search Tx: " + hash)
	ctx, chancel := context.WithTimeout(context.Background(), time.Second*60)
	defer chancel()
	transaction := make(chan *types.Receipt)
	go func(context context.Context, client *ethclient.Client) {
		for {
			//log.Println("wait tx")
			statusCode := -1
			txHash := common.HexToHash(hash)
			tx, err := client.TransactionReceipt(ctx, txHash)
			if err == nil {
				statusCode = int(tx.Status)
				transaction <- tx
				return
			} else {
				statusCode = -1
			}
			select {
			case <-ctx.Done():
				if statusCode == -1 {
					transaction <- nil
				} else {
					transaction <- tx
				}
				break
			default:
				_ = struct{}{}
			}
			time.Sleep(time.Second * 1)
		}
	}(ctx, c.Client)
	select {
	case tx := <-transaction:
		if tx != nil {
			txd, _, _ := c.Client.TransactionByHash(context.Background(), tx.TxHash)
			total := new(big.Int).Mul(txd.GasPrice(), new(big.Int).SetUint64(tx.GasUsed))
			GasUsed.Add(GasUsed, total)
			return int(tx.Status), tx
		}
		return -1, nil
	}
}

func (c *TokenBasic) GetTxAmount(tokenAddress, from, to string, tx *types.Receipt) (*big.Int, error) {
	hrc20, err := NewHrc20(common.HexToAddress(tokenAddress), c.Client)
	if err != nil {
		return nil, err
	}
	end := uint64(tx.BlockNumber.Int64())
	var fromCommon []common.Address
	if from != "" {
		fromCommon = []common.Address{common.HexToAddress(from)}
	}

	filter, err := hrc20.FilterTransfer(&bind.FilterOpts{Start: uint64(tx.BlockNumber.Int64()), End: &end}, fromCommon, []common.Address{common.HexToAddress(to)})
	if err != nil {
		return nil, err
	}
	defer filter.Close()
	for {
		if filter.Next() {
			if tx.TxHash.String() == filter.Event.Raw.TxHash.String() {
				return filter.Event.Value, nil

			}
		} else {
			return nil, errors.New("not found")
		}
	}
	return nil, errors.New("not found")
}

func (c *TokenBasic) CheckTransactionReceipt(_txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := c.Client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return -1
	}
	return (int(tx.Status))
}
