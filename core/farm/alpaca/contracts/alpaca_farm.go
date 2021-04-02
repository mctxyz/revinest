// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AplpacaFarmABI is the input ABI used to generate the binding from.
const AplpacaFarmABI = "[{\"inputs\":[{\"internalType\":\"contractAlpacaToken\",\"name\":\"_alpaca\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_alpacaPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusLockupBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alpaca\",\"outputs\":[{\"internalType\":\"contractAlpacaToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alpacaPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusLockUpBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentBlock\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"name\":\"isDuplicatedPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"manualMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingAlpaca\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accAlpacaPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accAlpacaPerShareTilBonusEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_alpacaPerBlock\",\"type\":\"uint256\"}],\"name\":\"setAlpacaPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusLockUpBps\",\"type\":\"uint256\"}],\"name\":\"setBonus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bonusDebt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fundedBy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AplpacaFarm is an auto generated Go binding around an Ethereum contract.
type AplpacaFarm struct {
	AplpacaFarmCaller     // Read-only binding to the contract
	AplpacaFarmTransactor // Write-only binding to the contract
	AplpacaFarmFilterer   // Log filterer for contract events
}

// AplpacaFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type AplpacaFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AplpacaFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AplpacaFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AplpacaFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AplpacaFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AplpacaFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AplpacaFarmSession struct {
	Contract     *AplpacaFarm      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AplpacaFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AplpacaFarmCallerSession struct {
	Contract *AplpacaFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AplpacaFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AplpacaFarmTransactorSession struct {
	Contract     *AplpacaFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AplpacaFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type AplpacaFarmRaw struct {
	Contract *AplpacaFarm // Generic contract binding to access the raw methods on
}

// AplpacaFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AplpacaFarmCallerRaw struct {
	Contract *AplpacaFarmCaller // Generic read-only contract binding to access the raw methods on
}

// AplpacaFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AplpacaFarmTransactorRaw struct {
	Contract *AplpacaFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAplpacaFarm creates a new instance of AplpacaFarm, bound to a specific deployed contract.
func NewAplpacaFarm(address common.Address, backend bind.ContractBackend) (*AplpacaFarm, error) {
	contract, err := bindAplpacaFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarm{AplpacaFarmCaller: AplpacaFarmCaller{contract: contract}, AplpacaFarmTransactor: AplpacaFarmTransactor{contract: contract}, AplpacaFarmFilterer: AplpacaFarmFilterer{contract: contract}}, nil
}

// NewAplpacaFarmCaller creates a new read-only instance of AplpacaFarm, bound to a specific deployed contract.
func NewAplpacaFarmCaller(address common.Address, caller bind.ContractCaller) (*AplpacaFarmCaller, error) {
	contract, err := bindAplpacaFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmCaller{contract: contract}, nil
}

// NewAplpacaFarmTransactor creates a new write-only instance of AplpacaFarm, bound to a specific deployed contract.
func NewAplpacaFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*AplpacaFarmTransactor, error) {
	contract, err := bindAplpacaFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmTransactor{contract: contract}, nil
}

// NewAplpacaFarmFilterer creates a new log filterer instance of AplpacaFarm, bound to a specific deployed contract.
func NewAplpacaFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*AplpacaFarmFilterer, error) {
	contract, err := bindAplpacaFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmFilterer{contract: contract}, nil
}

// bindAplpacaFarm binds a generic wrapper to an already deployed contract.
func bindAplpacaFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AplpacaFarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AplpacaFarm *AplpacaFarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AplpacaFarm.Contract.AplpacaFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AplpacaFarm *AplpacaFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.AplpacaFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AplpacaFarm *AplpacaFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.AplpacaFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AplpacaFarm *AplpacaFarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AplpacaFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AplpacaFarm *AplpacaFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AplpacaFarm *AplpacaFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.contract.Transact(opts, method, params...)
}

// Alpaca is a free data retrieval call binding the contract method 0x94faab23.
//
// Solidity: function alpaca() view returns(address)
func (_AplpacaFarm *AplpacaFarmCaller) Alpaca(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "alpaca")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Alpaca is a free data retrieval call binding the contract method 0x94faab23.
//
// Solidity: function alpaca() view returns(address)
func (_AplpacaFarm *AplpacaFarmSession) Alpaca() (common.Address, error) {
	return _AplpacaFarm.Contract.Alpaca(&_AplpacaFarm.CallOpts)
}

// Alpaca is a free data retrieval call binding the contract method 0x94faab23.
//
// Solidity: function alpaca() view returns(address)
func (_AplpacaFarm *AplpacaFarmCallerSession) Alpaca() (common.Address, error) {
	return _AplpacaFarm.Contract.Alpaca(&_AplpacaFarm.CallOpts)
}

// AlpacaPerBlock is a free data retrieval call binding the contract method 0x20f33d59.
//
// Solidity: function alpacaPerBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) AlpacaPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "alpacaPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlpacaPerBlock is a free data retrieval call binding the contract method 0x20f33d59.
//
// Solidity: function alpacaPerBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) AlpacaPerBlock() (*big.Int, error) {
	return _AplpacaFarm.Contract.AlpacaPerBlock(&_AplpacaFarm.CallOpts)
}

// AlpacaPerBlock is a free data retrieval call binding the contract method 0x20f33d59.
//
// Solidity: function alpacaPerBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) AlpacaPerBlock() (*big.Int, error) {
	return _AplpacaFarm.Contract.AlpacaPerBlock(&_AplpacaFarm.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) BonusEndBlock() (*big.Int, error) {
	return _AplpacaFarm.Contract.BonusEndBlock(&_AplpacaFarm.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) BonusEndBlock() (*big.Int, error) {
	return _AplpacaFarm.Contract.BonusEndBlock(&_AplpacaFarm.CallOpts)
}

// BonusLockUpBps is a free data retrieval call binding the contract method 0xcbdec1fc.
//
// Solidity: function bonusLockUpBps() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) BonusLockUpBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "bonusLockUpBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusLockUpBps is a free data retrieval call binding the contract method 0xcbdec1fc.
//
// Solidity: function bonusLockUpBps() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) BonusLockUpBps() (*big.Int, error) {
	return _AplpacaFarm.Contract.BonusLockUpBps(&_AplpacaFarm.CallOpts)
}

// BonusLockUpBps is a free data retrieval call binding the contract method 0xcbdec1fc.
//
// Solidity: function bonusLockUpBps() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) BonusLockUpBps() (*big.Int, error) {
	return _AplpacaFarm.Contract.BonusLockUpBps(&_AplpacaFarm.CallOpts)
}

// BonusMultiplier is a free data retrieval call binding the contract method 0xa8b973a1.
//
// Solidity: function bonusMultiplier() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) BonusMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "bonusMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusMultiplier is a free data retrieval call binding the contract method 0xa8b973a1.
//
// Solidity: function bonusMultiplier() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) BonusMultiplier() (*big.Int, error) {
	return _AplpacaFarm.Contract.BonusMultiplier(&_AplpacaFarm.CallOpts)
}

// BonusMultiplier is a free data retrieval call binding the contract method 0xa8b973a1.
//
// Solidity: function bonusMultiplier() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) BonusMultiplier() (*big.Int, error) {
	return _AplpacaFarm.Contract.BonusMultiplier(&_AplpacaFarm.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_AplpacaFarm *AplpacaFarmCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_AplpacaFarm *AplpacaFarmSession) Devaddr() (common.Address, error) {
	return _AplpacaFarm.Contract.Devaddr(&_AplpacaFarm.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_AplpacaFarm *AplpacaFarmCallerSession) Devaddr() (common.Address, error) {
	return _AplpacaFarm.Contract.Devaddr(&_AplpacaFarm.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _lastRewardBlock, uint256 _currentBlock) view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) GetMultiplier(opts *bind.CallOpts, _lastRewardBlock *big.Int, _currentBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "getMultiplier", _lastRewardBlock, _currentBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _lastRewardBlock, uint256 _currentBlock) view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) GetMultiplier(_lastRewardBlock *big.Int, _currentBlock *big.Int) (*big.Int, error) {
	return _AplpacaFarm.Contract.GetMultiplier(&_AplpacaFarm.CallOpts, _lastRewardBlock, _currentBlock)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _lastRewardBlock, uint256 _currentBlock) view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) GetMultiplier(_lastRewardBlock *big.Int, _currentBlock *big.Int) (*big.Int, error) {
	return _AplpacaFarm.Contract.GetMultiplier(&_AplpacaFarm.CallOpts, _lastRewardBlock, _currentBlock)
}

// IsDuplicatedPool is a free data retrieval call binding the contract method 0xb03df69d.
//
// Solidity: function isDuplicatedPool(address _stakeToken) view returns(bool)
func (_AplpacaFarm *AplpacaFarmCaller) IsDuplicatedPool(opts *bind.CallOpts, _stakeToken common.Address) (bool, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "isDuplicatedPool", _stakeToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDuplicatedPool is a free data retrieval call binding the contract method 0xb03df69d.
//
// Solidity: function isDuplicatedPool(address _stakeToken) view returns(bool)
func (_AplpacaFarm *AplpacaFarmSession) IsDuplicatedPool(_stakeToken common.Address) (bool, error) {
	return _AplpacaFarm.Contract.IsDuplicatedPool(&_AplpacaFarm.CallOpts, _stakeToken)
}

// IsDuplicatedPool is a free data retrieval call binding the contract method 0xb03df69d.
//
// Solidity: function isDuplicatedPool(address _stakeToken) view returns(bool)
func (_AplpacaFarm *AplpacaFarmCallerSession) IsDuplicatedPool(_stakeToken common.Address) (bool, error) {
	return _AplpacaFarm.Contract.IsDuplicatedPool(&_AplpacaFarm.CallOpts, _stakeToken)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AplpacaFarm *AplpacaFarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AplpacaFarm *AplpacaFarmSession) Owner() (common.Address, error) {
	return _AplpacaFarm.Contract.Owner(&_AplpacaFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AplpacaFarm *AplpacaFarmCallerSession) Owner() (common.Address, error) {
	return _AplpacaFarm.Contract.Owner(&_AplpacaFarm.CallOpts)
}

// PendingAlpaca is a free data retrieval call binding the contract method 0x94443b73.
//
// Solidity: function pendingAlpaca(uint256 _pid, address _user) view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) PendingAlpaca(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "pendingAlpaca", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingAlpaca is a free data retrieval call binding the contract method 0x94443b73.
//
// Solidity: function pendingAlpaca(uint256 _pid, address _user) view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) PendingAlpaca(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _AplpacaFarm.Contract.PendingAlpaca(&_AplpacaFarm.CallOpts, _pid, _user)
}

// PendingAlpaca is a free data retrieval call binding the contract method 0x94443b73.
//
// Solidity: function pendingAlpaca(uint256 _pid, address _user) view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) PendingAlpaca(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _AplpacaFarm.Contract.PendingAlpaca(&_AplpacaFarm.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accAlpacaPerShare, uint256 accAlpacaPerShareTilBonusEnd)
func (_AplpacaFarm *AplpacaFarmCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakeToken                   common.Address
	AllocPoint                   *big.Int
	LastRewardBlock              *big.Int
	AccAlpacaPerShare            *big.Int
	AccAlpacaPerShareTilBonusEnd *big.Int
}, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		StakeToken                   common.Address
		AllocPoint                   *big.Int
		LastRewardBlock              *big.Int
		AccAlpacaPerShare            *big.Int
		AccAlpacaPerShareTilBonusEnd *big.Int
	})

	outstruct.StakeToken = out[0].(common.Address)
	outstruct.AllocPoint = out[1].(*big.Int)
	outstruct.LastRewardBlock = out[2].(*big.Int)
	outstruct.AccAlpacaPerShare = out[3].(*big.Int)
	outstruct.AccAlpacaPerShareTilBonusEnd = out[4].(*big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accAlpacaPerShare, uint256 accAlpacaPerShareTilBonusEnd)
func (_AplpacaFarm *AplpacaFarmSession) PoolInfo(arg0 *big.Int) (struct {
	StakeToken                   common.Address
	AllocPoint                   *big.Int
	LastRewardBlock              *big.Int
	AccAlpacaPerShare            *big.Int
	AccAlpacaPerShareTilBonusEnd *big.Int
}, error) {
	return _AplpacaFarm.Contract.PoolInfo(&_AplpacaFarm.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accAlpacaPerShare, uint256 accAlpacaPerShareTilBonusEnd)
func (_AplpacaFarm *AplpacaFarmCallerSession) PoolInfo(arg0 *big.Int) (struct {
	StakeToken                   common.Address
	AllocPoint                   *big.Int
	LastRewardBlock              *big.Int
	AccAlpacaPerShare            *big.Int
	AccAlpacaPerShareTilBonusEnd *big.Int
}, error) {
	return _AplpacaFarm.Contract.PoolInfo(&_AplpacaFarm.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) PoolLength() (*big.Int, error) {
	return _AplpacaFarm.Contract.PoolLength(&_AplpacaFarm.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) PoolLength() (*big.Int, error) {
	return _AplpacaFarm.Contract.PoolLength(&_AplpacaFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) StartBlock() (*big.Int, error) {
	return _AplpacaFarm.Contract.StartBlock(&_AplpacaFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) StartBlock() (*big.Int, error) {
	return _AplpacaFarm.Contract.StartBlock(&_AplpacaFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmSession) TotalAllocPoint() (*big.Int, error) {
	return _AplpacaFarm.Contract.TotalAllocPoint(&_AplpacaFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_AplpacaFarm *AplpacaFarmCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _AplpacaFarm.Contract.TotalAllocPoint(&_AplpacaFarm.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 bonusDebt, address fundedBy)
func (_AplpacaFarm *AplpacaFarmCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	BonusDebt  *big.Int
	FundedBy   common.Address
}, error) {
	var out []interface{}
	err := _AplpacaFarm.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
		BonusDebt  *big.Int
		FundedBy   common.Address
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.RewardDebt = out[1].(*big.Int)
	outstruct.BonusDebt = out[2].(*big.Int)
	outstruct.FundedBy = out[3].(common.Address)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 bonusDebt, address fundedBy)
func (_AplpacaFarm *AplpacaFarmSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	BonusDebt  *big.Int
	FundedBy   common.Address
}, error) {
	return _AplpacaFarm.Contract.UserInfo(&_AplpacaFarm.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 bonusDebt, address fundedBy)
func (_AplpacaFarm *AplpacaFarmCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	BonusDebt  *big.Int
	FundedBy   common.Address
}, error) {
	return _AplpacaFarm.Contract.UserInfo(&_AplpacaFarm.CallOpts, arg0, arg1)
}

// AddPool is a paid mutator transaction binding the contract method 0x7abceffd.
//
// Solidity: function addPool(uint256 _allocPoint, address _stakeToken, bool _withUpdate) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) AddPool(opts *bind.TransactOpts, _allocPoint *big.Int, _stakeToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "addPool", _allocPoint, _stakeToken, _withUpdate)
}

// AddPool is a paid mutator transaction binding the contract method 0x7abceffd.
//
// Solidity: function addPool(uint256 _allocPoint, address _stakeToken, bool _withUpdate) returns()
func (_AplpacaFarm *AplpacaFarmSession) AddPool(_allocPoint *big.Int, _stakeToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.AddPool(&_AplpacaFarm.TransactOpts, _allocPoint, _stakeToken, _withUpdate)
}

// AddPool is a paid mutator transaction binding the contract method 0x7abceffd.
//
// Solidity: function addPool(uint256 _allocPoint, address _stakeToken, bool _withUpdate) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) AddPool(_allocPoint *big.Int, _stakeToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.AddPool(&_AplpacaFarm.TransactOpts, _allocPoint, _stakeToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _for, uint256 _pid, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) Deposit(opts *bind.TransactOpts, _for common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "deposit", _for, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _for, uint256 _pid, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmSession) Deposit(_for common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.Deposit(&_AplpacaFarm.TransactOpts, _for, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _for, uint256 _pid, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) Deposit(_for common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.Deposit(&_AplpacaFarm.TransactOpts, _for, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.EmergencyWithdraw(&_AplpacaFarm.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.EmergencyWithdraw(&_AplpacaFarm.TransactOpts, _pid)
}

// Harvest is a paid mutator transaction binding the contract method 0xddc63262.
//
// Solidity: function harvest(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) Harvest(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "harvest", _pid)
}

// Harvest is a paid mutator transaction binding the contract method 0xddc63262.
//
// Solidity: function harvest(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmSession) Harvest(_pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.Harvest(&_AplpacaFarm.TransactOpts, _pid)
}

// Harvest is a paid mutator transaction binding the contract method 0xddc63262.
//
// Solidity: function harvest(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) Harvest(_pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.Harvest(&_AplpacaFarm.TransactOpts, _pid)
}

// ManualMint is a paid mutator transaction binding the contract method 0xe4c5ff46.
//
// Solidity: function manualMint(address _to, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) ManualMint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "manualMint", _to, _amount)
}

// ManualMint is a paid mutator transaction binding the contract method 0xe4c5ff46.
//
// Solidity: function manualMint(address _to, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmSession) ManualMint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.ManualMint(&_AplpacaFarm.TransactOpts, _to, _amount)
}

// ManualMint is a paid mutator transaction binding the contract method 0xe4c5ff46.
//
// Solidity: function manualMint(address _to, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) ManualMint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.ManualMint(&_AplpacaFarm.TransactOpts, _to, _amount)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_AplpacaFarm *AplpacaFarmTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_AplpacaFarm *AplpacaFarmSession) MassUpdatePools() (*types.Transaction, error) {
	return _AplpacaFarm.Contract.MassUpdatePools(&_AplpacaFarm.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _AplpacaFarm.Contract.MassUpdatePools(&_AplpacaFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AplpacaFarm *AplpacaFarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AplpacaFarm *AplpacaFarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _AplpacaFarm.Contract.RenounceOwnership(&_AplpacaFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AplpacaFarm.Contract.RenounceOwnership(&_AplpacaFarm.TransactOpts)
}

// SetAlpacaPerBlock is a paid mutator transaction binding the contract method 0x4c8f70fd.
//
// Solidity: function setAlpacaPerBlock(uint256 _alpacaPerBlock) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) SetAlpacaPerBlock(opts *bind.TransactOpts, _alpacaPerBlock *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "setAlpacaPerBlock", _alpacaPerBlock)
}

// SetAlpacaPerBlock is a paid mutator transaction binding the contract method 0x4c8f70fd.
//
// Solidity: function setAlpacaPerBlock(uint256 _alpacaPerBlock) returns()
func (_AplpacaFarm *AplpacaFarmSession) SetAlpacaPerBlock(_alpacaPerBlock *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetAlpacaPerBlock(&_AplpacaFarm.TransactOpts, _alpacaPerBlock)
}

// SetAlpacaPerBlock is a paid mutator transaction binding the contract method 0x4c8f70fd.
//
// Solidity: function setAlpacaPerBlock(uint256 _alpacaPerBlock) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) SetAlpacaPerBlock(_alpacaPerBlock *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetAlpacaPerBlock(&_AplpacaFarm.TransactOpts, _alpacaPerBlock)
}

// SetBonus is a paid mutator transaction binding the contract method 0xfa5be8f8.
//
// Solidity: function setBonus(uint256 _bonusMultiplier, uint256 _bonusEndBlock, uint256 _bonusLockUpBps) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) SetBonus(opts *bind.TransactOpts, _bonusMultiplier *big.Int, _bonusEndBlock *big.Int, _bonusLockUpBps *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "setBonus", _bonusMultiplier, _bonusEndBlock, _bonusLockUpBps)
}

// SetBonus is a paid mutator transaction binding the contract method 0xfa5be8f8.
//
// Solidity: function setBonus(uint256 _bonusMultiplier, uint256 _bonusEndBlock, uint256 _bonusLockUpBps) returns()
func (_AplpacaFarm *AplpacaFarmSession) SetBonus(_bonusMultiplier *big.Int, _bonusEndBlock *big.Int, _bonusLockUpBps *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetBonus(&_AplpacaFarm.TransactOpts, _bonusMultiplier, _bonusEndBlock, _bonusLockUpBps)
}

// SetBonus is a paid mutator transaction binding the contract method 0xfa5be8f8.
//
// Solidity: function setBonus(uint256 _bonusMultiplier, uint256 _bonusEndBlock, uint256 _bonusLockUpBps) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) SetBonus(_bonusMultiplier *big.Int, _bonusEndBlock *big.Int, _bonusLockUpBps *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetBonus(&_AplpacaFarm.TransactOpts, _bonusMultiplier, _bonusEndBlock, _bonusLockUpBps)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _devaddr) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) SetDev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "setDev", _devaddr)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _devaddr) returns()
func (_AplpacaFarm *AplpacaFarmSession) SetDev(_devaddr common.Address) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetDev(&_AplpacaFarm.TransactOpts, _devaddr)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _devaddr) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) SetDev(_devaddr common.Address) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetDev(&_AplpacaFarm.TransactOpts, _devaddr)
}

// SetPool is a paid mutator transaction binding the contract method 0x46ca6bea.
//
// Solidity: function setPool(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) SetPool(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "setPool", _pid, _allocPoint, _withUpdate)
}

// SetPool is a paid mutator transaction binding the contract method 0x46ca6bea.
//
// Solidity: function setPool(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_AplpacaFarm *AplpacaFarmSession) SetPool(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetPool(&_AplpacaFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetPool is a paid mutator transaction binding the contract method 0x46ca6bea.
//
// Solidity: function setPool(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) SetPool(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.SetPool(&_AplpacaFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AplpacaFarm *AplpacaFarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.TransferOwnership(&_AplpacaFarm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.TransferOwnership(&_AplpacaFarm.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.UpdatePool(&_AplpacaFarm.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.UpdatePool(&_AplpacaFarm.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address _for, uint256 _pid, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) Withdraw(opts *bind.TransactOpts, _for common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "withdraw", _for, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address _for, uint256 _pid, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmSession) Withdraw(_for common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.Withdraw(&_AplpacaFarm.TransactOpts, _for, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb5c5f672.
//
// Solidity: function withdraw(address _for, uint256 _pid, uint256 _amount) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) Withdraw(_for common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.Withdraw(&_AplpacaFarm.TransactOpts, _for, _pid, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xcc6dbc27.
//
// Solidity: function withdrawAll(address _for, uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactor) WithdrawAll(opts *bind.TransactOpts, _for common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.contract.Transact(opts, "withdrawAll", _for, _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xcc6dbc27.
//
// Solidity: function withdrawAll(address _for, uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmSession) WithdrawAll(_for common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.WithdrawAll(&_AplpacaFarm.TransactOpts, _for, _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0xcc6dbc27.
//
// Solidity: function withdrawAll(address _for, uint256 _pid) returns()
func (_AplpacaFarm *AplpacaFarmTransactorSession) WithdrawAll(_for common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _AplpacaFarm.Contract.WithdrawAll(&_AplpacaFarm.TransactOpts, _for, _pid)
}

// AplpacaFarmDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AplpacaFarm contract.
type AplpacaFarmDepositIterator struct {
	Event *AplpacaFarmDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AplpacaFarmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AplpacaFarmDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AplpacaFarmDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AplpacaFarmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AplpacaFarmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AplpacaFarmDeposit represents a Deposit event raised by the AplpacaFarm contract.
type AplpacaFarmDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*AplpacaFarmDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AplpacaFarm.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmDepositIterator{contract: _AplpacaFarm.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AplpacaFarmDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AplpacaFarm.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AplpacaFarmDeposit)
				if err := _AplpacaFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) ParseDeposit(log types.Log) (*AplpacaFarmDeposit, error) {
	event := new(AplpacaFarmDeposit)
	if err := _AplpacaFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AplpacaFarmEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the AplpacaFarm contract.
type AplpacaFarmEmergencyWithdrawIterator struct {
	Event *AplpacaFarmEmergencyWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AplpacaFarmEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AplpacaFarmEmergencyWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AplpacaFarmEmergencyWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AplpacaFarmEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AplpacaFarmEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AplpacaFarmEmergencyWithdraw represents a EmergencyWithdraw event raised by the AplpacaFarm contract.
type AplpacaFarmEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*AplpacaFarmEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AplpacaFarm.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmEmergencyWithdrawIterator{contract: _AplpacaFarm.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *AplpacaFarmEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AplpacaFarm.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AplpacaFarmEmergencyWithdraw)
				if err := _AplpacaFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) ParseEmergencyWithdraw(log types.Log) (*AplpacaFarmEmergencyWithdraw, error) {
	event := new(AplpacaFarmEmergencyWithdraw)
	if err := _AplpacaFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AplpacaFarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AplpacaFarm contract.
type AplpacaFarmOwnershipTransferredIterator struct {
	Event *AplpacaFarmOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AplpacaFarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AplpacaFarmOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AplpacaFarmOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AplpacaFarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AplpacaFarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AplpacaFarmOwnershipTransferred represents a OwnershipTransferred event raised by the AplpacaFarm contract.
type AplpacaFarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AplpacaFarm *AplpacaFarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AplpacaFarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AplpacaFarm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmOwnershipTransferredIterator{contract: _AplpacaFarm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AplpacaFarm *AplpacaFarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AplpacaFarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AplpacaFarm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AplpacaFarmOwnershipTransferred)
				if err := _AplpacaFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AplpacaFarm *AplpacaFarmFilterer) ParseOwnershipTransferred(log types.Log) (*AplpacaFarmOwnershipTransferred, error) {
	event := new(AplpacaFarmOwnershipTransferred)
	if err := _AplpacaFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AplpacaFarmWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AplpacaFarm contract.
type AplpacaFarmWithdrawIterator struct {
	Event *AplpacaFarmWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AplpacaFarmWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AplpacaFarmWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AplpacaFarmWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AplpacaFarmWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AplpacaFarmWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AplpacaFarmWithdraw represents a Withdraw event raised by the AplpacaFarm contract.
type AplpacaFarmWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*AplpacaFarmWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AplpacaFarm.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &AplpacaFarmWithdrawIterator{contract: _AplpacaFarm.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AplpacaFarmWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _AplpacaFarm.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AplpacaFarmWithdraw)
				if err := _AplpacaFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_AplpacaFarm *AplpacaFarmFilterer) ParseWithdraw(log types.Log) (*AplpacaFarmWithdraw, error) {
	event := new(AplpacaFarmWithdraw)
	if err := _AplpacaFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
