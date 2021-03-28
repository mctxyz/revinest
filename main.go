package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"math/big"
	"os"
	"reinvest/core/farm"
	"reinvest/printer"
	"reinvest/token"
	"reinvest/utils"
	"runtime"
	"strconv"
	"time"
)

var wallet string
var reinvestInterval int
var TotalReward = big.NewInt(0)
var print = printer.NewPrinter()

func main() {

	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()

	farm, cancel, err := farm.NewFarm()

	if err != nil {
		print.Error(err.Error())
		cancel()
		return
	}
	minutesValidate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("Invalid Reinvest Time Interval")
		}
		return nil
	}

	between := promptui.Prompt{
		Label:    "Reinvest  Interval (minute)",
		Validate: minutesValidate,
	}

	interval, err := between.Run()
	if err != nil {
		print.Error(err.Error())
		pause()
		return
	}

	reinvestInterval, _ = strconv.Atoi(interval)
	if reinvestInterval == 0 {
		reinvestInterval = 10
	}

	if err := farm.Start(); err != nil {
		print.Error(err.Error())
		pause()
		return
	}
	timer := time.NewTimer(time.Minute * time.Duration(reinvestInterval))
	rewardToken := farm.RewardToken()
	Run(farm)
	PrintTotalBalance(rewardToken)
	for {

		select {
		case <-timer.C:
			Run(farm)
			PrintTotalBalance(rewardToken)
			timer.Reset(time.Minute * time.Duration(reinvestInterval))

		default:
			time.Sleep(time.Millisecond * 10)
			_ = struct{}{}
		}

	}
	return

}
func PrintTotalBalance(rewardToken *token.Token) {
	log.Printf(
		"Until %s Total gas used %s  rewards %s %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		utils.ToDecimal(token.GasUsed, 18),
		utils.ToDecimal(TotalReward, int(rewardToken.Decimals)),
		rewardToken.Symbol,
	)
}
func Run(farm farm.Farm) {
	realPendingRewardAmount, err := farm.Harvest()
	if err != nil {
		print.Error(fmt.Sprintf("Get Pending Rewards  Err  %s \n", err.Error()))
	}
	if realPendingRewardAmount.Cmp(big.NewInt(0)) >= 1 {
		TotalReward.Add(TotalReward, realPendingRewardAmount)
		wishA, wishB, tokenAAddres, tokenBAddress, err := farm.SwapRewardToPairWithRetry(realPendingRewardAmount, 10)
		if err != nil {
			print.Error("Swap Error: " + err.Error())
			return
		}
		txHash, err := farm.AddLiquidityWithRetry(wishA, wishB, tokenAAddres, tokenBAddress, 10)
		log.Println(txHash)
		if err != nil {
			print.Error("AddLiquidity Error: " + err.Error() + " Tx: " + txHash)
			return
		}
		print.Success("AddLiquidity Success: " + txHash)
		lpAmount, reinvestTx, err := farm.Reinvest()
		if err != nil {
			print.Error("Reinvest Error: " + err.Error() + " Tx: " + reinvestTx)
			return
		}
		lpToken := farm.LpToken()
		print.Success(fmt.Sprintf("Despoit %s %s Success", utils.ToDecimal(lpAmount, int(lpToken.Decimals)).String(), lpToken.Name))
		fmt.Println("\n")
	}
}
func pause() {
	fmt.Print("Press Any Key to Exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	return
}
