package mdex

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"reinvest/token"
	"reinvest/utils"
)

func (c *MdexFarm) RewardToken() *token.Token {
	return c.RewardTokenInfo

}
func (c *MdexFarm) Harvest() (*big.Int, error) {


	c.Printer.Info("Harvest Reward...")
	pendingReward, err := c.Pending(c.FarmConfig.NetWork.FarmAddress, c.FarmConfig.Wallet, int(c.FarmConfig.PooID))
	if err != nil {
		return nil, err
	}
	c.Printer.Info(utils.ToDecimal(pendingReward.Amount, int(c.RewardTokenInfo.Decimals)).String() + " need to harvest")
	if pendingReward.Amount.Cmp(big.NewInt(0)) >= 1 {
		tx, err := c.Deposit(c.FarmConfig.NetWork.FarmAddress, big.NewInt(0), int(c.FarmConfig.PooID))
		if err != nil {
			c.Printer.Error("Harvest error: " + err.Error())
			return big.NewInt(0), err
		}
		txStatus, _tx := c.TokenBasic.WaitForBlockCompletation(tx.Hash().String())
		if txStatus != 1 {
			return big.NewInt(0), errors.New("Harvest Err Tx :" + tx.Hash().String())
		}

		sendRewardAmountToWallet, err := c.TokenBasic.GetTxAmount(c.RewardTokenInfo.Address, "", c.FarmConfig.Wallet, _tx)
		if err != nil {
			return big.NewInt(0), fmt.Errorf("Search Real Reward Amount To Wallet Error: %w", err)
		}
		c.Printer.Success(utils.ToDecimal(sendRewardAmountToWallet, int(c.RewardTokenInfo.Decimals)).String() + " " + c.RewardTokenInfo.Symbol + " -> " + c.FarmConfig.Wallet)

		return sendRewardAmountToWallet, nil
	}
	return big.NewInt(0), nil
}

func (c *MdexFarm) SwapRewardToPairWithRetry(rewardAmount *big.Int, tryCount int) (*big.Int, *big.Int, string, string, error) {
	pairAmount := make(map[string]*big.Int)
	avg := new(big.Int).Div(rewardAmount, big.NewInt(2))
	swapTokenARealAmount, swapTxHash, err := c.Swap(avg, c.RewardTokenInfo.Address, c.TokenAInfo.Address, tryCount)
	if err != nil {
		return nil, nil, c.TokenAInfo.Address, c.TokenBInfo.Address, err
	}
	pairAmount[c.TokenAInfo.Address] = swapTokenARealAmount
	c.Printer.Success(fmt.Sprintf(
		"Swap %s %s -> %s %s Tx: %s",
		utils.ToDecimal(avg, int(c.RewardTokenInfo.Decimals)).String(),
		c.RewardTokenInfo.Symbol,
		utils.ToDecimal(swapTokenARealAmount, int(c.TokenAInfo.Decimals)),
		c.TokenAInfo.Symbol,
		swapTxHash,
	))

	swapTokenBRealAmount, swapTxHash, err := c.Swap(avg, c.RewardTokenInfo.Address, c.TokenBInfo.Address, tryCount)
	if err != nil {
		return nil, nil, c.TokenAInfo.Address, c.TokenBInfo.Address, err
	}
	pairAmount[c.TokenBInfo.Address] = swapTokenBRealAmount
	c.Printer.Success(fmt.Sprintf(
		"Swap %s %s -> %s %s Tx: %s",
		utils.ToDecimal(avg, int(c.RewardTokenInfo.Decimals)).String(),
		c.RewardTokenInfo.Symbol,
		utils.ToDecimal(swapTokenBRealAmount, int(c.TokenBInfo.Decimals)),
		c.TokenBInfo.Symbol,
		swapTxHash,
	))

	return swapTokenARealAmount, swapTokenBRealAmount, c.TokenAInfo.Address, c.TokenBInfo.Address, err
}
func (c *MdexFarm) AddLiquidityWithRetry(wishA *big.Int, wishB *big.Int, tokenAAddress, tokenBAddress string, tryCount int) (string, error) {
	count := 1
	var swapTxHash string
	for {
		if count >= tryCount {
			return swapTxHash, errors.New("Swap  Too Many errors")
		}
		if count > 1 {
			fmt.Printf("Try AddLiquidity %d \n", count)
		}
		addLiquidityTx, err := c.addLiquidity(wishA, wishB, tokenAAddress, tokenBAddress)
		if err != nil {
			fmt.Printf("addLiquidity Error %s \n", err.Error())
			count++
			continue
		}
		swapTxHash = addLiquidityTx.Hash().String()
		c.Printer.Info(fmt.Sprintf("addLiquidity Tx : %s \n", addLiquidityTx.Hash().String()))
		addLiquidityTxStatus, _ := c.TokenBasic.WaitForBlockCompletation(addLiquidityTx.Hash().String())
		if addLiquidityTxStatus == 1 {
			//c.Printer.Success("%s %s + %s %s => %s Send To %s",utils.ToDecimal())
			return swapTxHash, nil
		}
		count++
	}
}
func (c *MdexFarm) Reinvest() (*big.Int, string, error) {
	LpTokenBalance, err := c.TokenBasic.GetMyTokenInfo(c.LpTokenInfo.Address)
	if err != nil {
		c.Printer.Error(fmt.Sprintf("Get My LP Token Info  Err  %s \n", err))

		return nil, "", fmt.Errorf("Get My LP Token Info  Err  %w", err)
	}
	lpTokenApprove, err := c.TokenBasic.Approve(c.LpTokenInfo.Address, c.FarmConfig.NetWork.FarmAddress, LpTokenBalance.Balance)
	if err != nil {

		return nil, "", fmt.Errorf("Approve LP Token Error: %w", err)
	}
	if !lpTokenApprove {
		c.Printer.Error("Approve LP Token  Fail")

		return nil, "", errors.New("Approve LP Token  Fail")
	}
	depTx, err := c.Deposit(c.FarmConfig.NetWork.FarmAddress, LpTokenBalance.Balance, int(c.FarmConfig.PooID))
	if err != nil {

		return nil, "", fmt.Errorf("Deposit Error: %w", err)
	}
	depTxStatus, _ := c.TokenBasic.WaitForBlockCompletation(depTx.Hash().String())
	if depTxStatus != 1 {
		return nil, "", errors.New("Deposit Err Txh :" + depTx.Hash().String())
	}

	return LpTokenBalance.Balance, depTx.Hash().String(), nil
}
func (c *MdexFarm) LpToken() *token.Token {
	return c.LpTokenInfo
}
func (c *MdexFarm) Start() error {

	farmInfo, err := c.GetPoolInfo(c.FarmConfig.NetWork.FarmAddress, int(c.FarmConfig.PooID))
	if err != nil {
		return fmt.Errorf("Get LpToken Info Err %w ", err)
	}
	LpToken, err := NewLpToken(farmInfo.LpToken, c.Client)
	if err != nil {
		return fmt.Errorf("Get LpToken Info Err %w ", err)

	}
	address0, address1, err := LpToken.Pair()
	if err != nil {
		return fmt.Errorf("Get Token Pair Info Err %w ", err)

	}
	LpTokenInfo, err := c.TokenBasic.TokenInfo(farmInfo.LpToken)
	if err != nil {
		return fmt.Errorf("Get LpToken  Info Err %w  ", err)

	}
	swapRouter, err := NewSwapRouter(c.FarmConfig.NetWork.Router, c.Client, c)
	if err != nil {
		return fmt.Errorf("Get Router  Err %w  ", err)
	}
	factory, err := NewSwapFactory(swapRouter.Factory, c.Client, c)

	if err != nil {
		return fmt.Errorf("Get Factory  Err %w   ", err)
	}
	sortRes, err := factory.SwapContract.SortTokens(&bind.CallOpts{}, common.HexToAddress(address0), common.HexToAddress(address1))
	if err != nil {
		log.Printf("Sort Tokens  Err %w  \n", err)
	} else {
		address0 = sortRes.Token0.String()
		address1 = sortRes.Token1.String()
	}
	rewardTokenInfo, err := c.TokenBasic.TokenInfo(c.FarmConfig.NetWork.RewardToken)
	if err != nil {
		return fmt.Errorf("Get Reward Token Info Err %w   ", err)
	}

	tokenAInfo, err := c.TokenBasic.TokenInfo(address0)

	if err != nil {
		return fmt.Errorf("Get My Token A Info  Err %w ", err)
	}
	tokenBInfo, err := c.TokenBasic.TokenInfo(address1)
	if err != nil {
		return fmt.Errorf("Get My Token B Info  Er %w ", err)
	}
	c.TokenAInfo = tokenAInfo
	c.TokenBInfo = tokenBInfo
	c.RewardTokenInfo = rewardTokenInfo
	c.LpTokenInfo = LpTokenInfo
	return nil
}
