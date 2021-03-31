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

// SwapMiningABI is the input ABI used to generate the binding from.
const SwapMiningABI = "[{\"inputs\":[{\"internalType\":\"contractIMdx\",\"name\":\"_mdx\",\"type\":\"address\"},{\"internalType\":\"contractIMdexFactory\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"contractIOracle\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mdxPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"addPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addToken\",\"type\":\"address\"}],\"name\":\"addWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delToken\",\"type\":\"address\"}],\"name\":\"delWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIMdexFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastRewardBlock\",\"type\":\"uint256\"}],\"name\":\"getMdxReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"anchorToken\",\"type\":\"address\"}],\"name\":\"getQuantity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getUserReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getWhitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halvingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massMintPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdx\",\"outputs\":[{\"internalType\":\"contractIMdx\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdxPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairOfPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocMdxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"setHalvingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPerBlock\",\"type\":\"uint256\"}],\"name\":\"setMdxPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"setPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takerWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwapMining is an auto generated Go binding around an Ethereum contract.
type SwapMining struct {
	SwapMiningCaller     // Read-only binding to the contract
	SwapMiningTransactor // Write-only binding to the contract
	SwapMiningFilterer   // Log filterer for contract events
}

// SwapMiningCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapMiningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapMiningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapMiningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapMiningFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapMiningFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapMiningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapMiningSession struct {
	Contract     *SwapMining       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapMiningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapMiningCallerSession struct {
	Contract *SwapMiningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SwapMiningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapMiningTransactorSession struct {
	Contract     *SwapMiningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SwapMiningRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapMiningRaw struct {
	Contract *SwapMining // Generic contract binding to access the raw methods on
}

// SwapMiningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapMiningCallerRaw struct {
	Contract *SwapMiningCaller // Generic read-only contract binding to access the raw methods on
}

// SwapMiningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapMiningTransactorRaw struct {
	Contract *SwapMiningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapMining creates a new instance of SwapMining, bound to a specific deployed contract.
func NewSwapMining(address common.Address, backend bind.ContractBackend) (*SwapMining, error) {
	contract, err := bindSwapMining(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapMining{SwapMiningCaller: SwapMiningCaller{contract: contract}, SwapMiningTransactor: SwapMiningTransactor{contract: contract}, SwapMiningFilterer: SwapMiningFilterer{contract: contract}}, nil
}

// NewSwapMiningCaller creates a new read-only instance of SwapMining, bound to a specific deployed contract.
func NewSwapMiningCaller(address common.Address, caller bind.ContractCaller) (*SwapMiningCaller, error) {
	contract, err := bindSwapMining(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapMiningCaller{contract: contract}, nil
}

// NewSwapMiningTransactor creates a new write-only instance of SwapMining, bound to a specific deployed contract.
func NewSwapMiningTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapMiningTransactor, error) {
	contract, err := bindSwapMining(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapMiningTransactor{contract: contract}, nil
}

// NewSwapMiningFilterer creates a new log filterer instance of SwapMining, bound to a specific deployed contract.
func NewSwapMiningFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapMiningFilterer, error) {
	contract, err := bindSwapMining(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapMiningFilterer{contract: contract}, nil
}

// bindSwapMining binds a generic wrapper to an already deployed contract.
func bindSwapMining(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapMiningABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapMining *SwapMiningRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapMining.Contract.SwapMiningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapMining *SwapMiningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMining.Contract.SwapMiningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapMining *SwapMiningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapMining.Contract.SwapMiningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapMining *SwapMiningCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapMining.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapMining *SwapMiningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMining.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapMining *SwapMiningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapMining.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapMining *SwapMiningCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapMining *SwapMiningSession) Factory() (common.Address, error) {
	return _SwapMining.Contract.Factory(&_SwapMining.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SwapMining *SwapMiningCallerSession) Factory() (common.Address, error) {
	return _SwapMining.Contract.Factory(&_SwapMining.CallOpts)
}

// GetMdxReward is a free data retrieval call binding the contract method 0xe377fe35.
//
// Solidity: function getMdxReward(uint256 _lastRewardBlock) view returns(uint256)
func (_SwapMining *SwapMiningCaller) GetMdxReward(opts *bind.CallOpts, _lastRewardBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "getMdxReward", _lastRewardBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMdxReward is a free data retrieval call binding the contract method 0xe377fe35.
//
// Solidity: function getMdxReward(uint256 _lastRewardBlock) view returns(uint256)
func (_SwapMining *SwapMiningSession) GetMdxReward(_lastRewardBlock *big.Int) (*big.Int, error) {
	return _SwapMining.Contract.GetMdxReward(&_SwapMining.CallOpts, _lastRewardBlock)
}

// GetMdxReward is a free data retrieval call binding the contract method 0xe377fe35.
//
// Solidity: function getMdxReward(uint256 _lastRewardBlock) view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) GetMdxReward(_lastRewardBlock *big.Int) (*big.Int, error) {
	return _SwapMining.Contract.GetMdxReward(&_SwapMining.CallOpts, _lastRewardBlock)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 _pid) view returns(address, address, uint256, uint256, uint256, uint256)
func (_SwapMining *SwapMiningCaller) GetPoolInfo(opts *bind.CallOpts, _pid *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "getPoolInfo", _pid)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 _pid) view returns(address, address, uint256, uint256, uint256, uint256)
func (_SwapMining *SwapMiningSession) GetPoolInfo(_pid *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SwapMining.Contract.GetPoolInfo(&_SwapMining.CallOpts, _pid)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 _pid) view returns(address, address, uint256, uint256, uint256, uint256)
func (_SwapMining *SwapMiningCallerSession) GetPoolInfo(_pid *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SwapMining.Contract.GetPoolInfo(&_SwapMining.CallOpts, _pid)
}

// GetQuantity is a free data retrieval call binding the contract method 0xa7ef67f6.
//
// Solidity: function getQuantity(address outputToken, uint256 outputAmount, address anchorToken) view returns(uint256)
func (_SwapMining *SwapMiningCaller) GetQuantity(opts *bind.CallOpts, outputToken common.Address, outputAmount *big.Int, anchorToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "getQuantity", outputToken, outputAmount, anchorToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuantity is a free data retrieval call binding the contract method 0xa7ef67f6.
//
// Solidity: function getQuantity(address outputToken, uint256 outputAmount, address anchorToken) view returns(uint256)
func (_SwapMining *SwapMiningSession) GetQuantity(outputToken common.Address, outputAmount *big.Int, anchorToken common.Address) (*big.Int, error) {
	return _SwapMining.Contract.GetQuantity(&_SwapMining.CallOpts, outputToken, outputAmount, anchorToken)
}

// GetQuantity is a free data retrieval call binding the contract method 0xa7ef67f6.
//
// Solidity: function getQuantity(address outputToken, uint256 outputAmount, address anchorToken) view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) GetQuantity(outputToken common.Address, outputAmount *big.Int, anchorToken common.Address) (*big.Int, error) {
	return _SwapMining.Contract.GetQuantity(&_SwapMining.CallOpts, outputToken, outputAmount, anchorToken)
}

// GetUserReward is a free data retrieval call binding the contract method 0x9fb08cb7.
//
// Solidity: function getUserReward(uint256 _pid) view returns(uint256, uint256)
func (_SwapMining *SwapMiningCaller) GetUserReward(opts *bind.CallOpts, _pid *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "getUserReward", _pid)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserReward is a free data retrieval call binding the contract method 0x9fb08cb7.
//
// Solidity: function getUserReward(uint256 _pid) view returns(uint256, uint256)
func (_SwapMining *SwapMiningSession) GetUserReward(_pid *big.Int) (*big.Int, *big.Int, error) {
	return _SwapMining.Contract.GetUserReward(&_SwapMining.CallOpts, _pid)
}

// GetUserReward is a free data retrieval call binding the contract method 0x9fb08cb7.
//
// Solidity: function getUserReward(uint256 _pid) view returns(uint256, uint256)
func (_SwapMining *SwapMiningCallerSession) GetUserReward(_pid *big.Int) (*big.Int, *big.Int, error) {
	return _SwapMining.Contract.GetUserReward(&_SwapMining.CallOpts, _pid)
}

// GetWhitelist is a free data retrieval call binding the contract method 0x96db752a.
//
// Solidity: function getWhitelist(uint256 _index) view returns(address)
func (_SwapMining *SwapMiningCaller) GetWhitelist(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "getWhitelist", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWhitelist is a free data retrieval call binding the contract method 0x96db752a.
//
// Solidity: function getWhitelist(uint256 _index) view returns(address)
func (_SwapMining *SwapMiningSession) GetWhitelist(_index *big.Int) (common.Address, error) {
	return _SwapMining.Contract.GetWhitelist(&_SwapMining.CallOpts, _index)
}

// GetWhitelist is a free data retrieval call binding the contract method 0x96db752a.
//
// Solidity: function getWhitelist(uint256 _index) view returns(address)
func (_SwapMining *SwapMiningCallerSession) GetWhitelist(_index *big.Int) (common.Address, error) {
	return _SwapMining.Contract.GetWhitelist(&_SwapMining.CallOpts, _index)
}

// GetWhitelistLength is a free data retrieval call binding the contract method 0xa13202e9.
//
// Solidity: function getWhitelistLength() view returns(uint256)
func (_SwapMining *SwapMiningCaller) GetWhitelistLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "getWhitelistLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWhitelistLength is a free data retrieval call binding the contract method 0xa13202e9.
//
// Solidity: function getWhitelistLength() view returns(uint256)
func (_SwapMining *SwapMiningSession) GetWhitelistLength() (*big.Int, error) {
	return _SwapMining.Contract.GetWhitelistLength(&_SwapMining.CallOpts)
}

// GetWhitelistLength is a free data retrieval call binding the contract method 0xa13202e9.
//
// Solidity: function getWhitelistLength() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) GetWhitelistLength() (*big.Int, error) {
	return _SwapMining.Contract.GetWhitelistLength(&_SwapMining.CallOpts)
}

// HalvingPeriod is a free data retrieval call binding the contract method 0x5a3e251f.
//
// Solidity: function halvingPeriod() view returns(uint256)
func (_SwapMining *SwapMiningCaller) HalvingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "halvingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalvingPeriod is a free data retrieval call binding the contract method 0x5a3e251f.
//
// Solidity: function halvingPeriod() view returns(uint256)
func (_SwapMining *SwapMiningSession) HalvingPeriod() (*big.Int, error) {
	return _SwapMining.Contract.HalvingPeriod(&_SwapMining.CallOpts)
}

// HalvingPeriod is a free data retrieval call binding the contract method 0x5a3e251f.
//
// Solidity: function halvingPeriod() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) HalvingPeriod() (*big.Int, error) {
	return _SwapMining.Contract.HalvingPeriod(&_SwapMining.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_SwapMining *SwapMiningCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_SwapMining *SwapMiningSession) IsOwner(account common.Address) (bool, error) {
	return _SwapMining.Contract.IsOwner(&_SwapMining.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_SwapMining *SwapMiningCallerSession) IsOwner(account common.Address) (bool, error) {
	return _SwapMining.Contract.IsOwner(&_SwapMining.CallOpts, account)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address _token) view returns(bool)
func (_SwapMining *SwapMiningCaller) IsWhitelist(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "isWhitelist", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address _token) view returns(bool)
func (_SwapMining *SwapMiningSession) IsWhitelist(_token common.Address) (bool, error) {
	return _SwapMining.Contract.IsWhitelist(&_SwapMining.CallOpts, _token)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address _token) view returns(bool)
func (_SwapMining *SwapMiningCallerSession) IsWhitelist(_token common.Address) (bool, error) {
	return _SwapMining.Contract.IsWhitelist(&_SwapMining.CallOpts, _token)
}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_SwapMining *SwapMiningCaller) Mdx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "mdx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_SwapMining *SwapMiningSession) Mdx() (common.Address, error) {
	return _SwapMining.Contract.Mdx(&_SwapMining.CallOpts)
}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_SwapMining *SwapMiningCallerSession) Mdx() (common.Address, error) {
	return _SwapMining.Contract.Mdx(&_SwapMining.CallOpts)
}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_SwapMining *SwapMiningCaller) MdxPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "mdxPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_SwapMining *SwapMiningSession) MdxPerBlock() (*big.Int, error) {
	return _SwapMining.Contract.MdxPerBlock(&_SwapMining.CallOpts)
}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) MdxPerBlock() (*big.Int, error) {
	return _SwapMining.Contract.MdxPerBlock(&_SwapMining.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_SwapMining *SwapMiningCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_SwapMining *SwapMiningSession) Oracle() (common.Address, error) {
	return _SwapMining.Contract.Oracle(&_SwapMining.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_SwapMining *SwapMiningCallerSession) Oracle() (common.Address, error) {
	return _SwapMining.Contract.Oracle(&_SwapMining.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapMining *SwapMiningCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapMining *SwapMiningSession) Owner() (common.Address, error) {
	return _SwapMining.Contract.Owner(&_SwapMining.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapMining *SwapMiningCallerSession) Owner() (common.Address, error) {
	return _SwapMining.Contract.Owner(&_SwapMining.CallOpts)
}

// PairOfPid is a free data retrieval call binding the contract method 0x0495de50.
//
// Solidity: function pairOfPid(address ) view returns(uint256)
func (_SwapMining *SwapMiningCaller) PairOfPid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "pairOfPid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairOfPid is a free data retrieval call binding the contract method 0x0495de50.
//
// Solidity: function pairOfPid(address ) view returns(uint256)
func (_SwapMining *SwapMiningSession) PairOfPid(arg0 common.Address) (*big.Int, error) {
	return _SwapMining.Contract.PairOfPid(&_SwapMining.CallOpts, arg0)
}

// PairOfPid is a free data retrieval call binding the contract method 0x0495de50.
//
// Solidity: function pairOfPid(address ) view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) PairOfPid(arg0 common.Address) (*big.Int, error) {
	return _SwapMining.Contract.PairOfPid(&_SwapMining.CallOpts, arg0)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 blockNumber) view returns(uint256)
func (_SwapMining *SwapMiningCaller) Phase(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "phase", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 blockNumber) view returns(uint256)
func (_SwapMining *SwapMiningSession) Phase(blockNumber *big.Int) (*big.Int, error) {
	return _SwapMining.Contract.Phase(&_SwapMining.CallOpts, blockNumber)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 blockNumber) view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) Phase(blockNumber *big.Int) (*big.Int, error) {
	return _SwapMining.Contract.Phase(&_SwapMining.CallOpts, blockNumber)
}

// Phase0 is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint256)
func (_SwapMining *SwapMiningCaller) Phase0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "phase0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Phase0 is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint256)
func (_SwapMining *SwapMiningSession) Phase0() (*big.Int, error) {
	return _SwapMining.Contract.Phase0(&_SwapMining.CallOpts)
}

// Phase0 is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) Phase0() (*big.Int, error) {
	return _SwapMining.Contract.Phase0(&_SwapMining.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address pair, uint256 quantity, uint256 totalQuantity, uint256 allocPoint, uint256 allocMdxAmount, uint256 lastRewardBlock)
func (_SwapMining *SwapMiningCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Pair            common.Address
	Quantity        *big.Int
	TotalQuantity   *big.Int
	AllocPoint      *big.Int
	AllocMdxAmount  *big.Int
	LastRewardBlock *big.Int
}, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Pair            common.Address
		Quantity        *big.Int
		TotalQuantity   *big.Int
		AllocPoint      *big.Int
		AllocMdxAmount  *big.Int
		LastRewardBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pair = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalQuantity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AllocPoint = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AllocMdxAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address pair, uint256 quantity, uint256 totalQuantity, uint256 allocPoint, uint256 allocMdxAmount, uint256 lastRewardBlock)
func (_SwapMining *SwapMiningSession) PoolInfo(arg0 *big.Int) (struct {
	Pair            common.Address
	Quantity        *big.Int
	TotalQuantity   *big.Int
	AllocPoint      *big.Int
	AllocMdxAmount  *big.Int
	LastRewardBlock *big.Int
}, error) {
	return _SwapMining.Contract.PoolInfo(&_SwapMining.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address pair, uint256 quantity, uint256 totalQuantity, uint256 allocPoint, uint256 allocMdxAmount, uint256 lastRewardBlock)
func (_SwapMining *SwapMiningCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Pair            common.Address
	Quantity        *big.Int
	TotalQuantity   *big.Int
	AllocPoint      *big.Int
	AllocMdxAmount  *big.Int
	LastRewardBlock *big.Int
}, error) {
	return _SwapMining.Contract.PoolInfo(&_SwapMining.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SwapMining *SwapMiningCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SwapMining *SwapMiningSession) PoolLength() (*big.Int, error) {
	return _SwapMining.Contract.PoolLength(&_SwapMining.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) PoolLength() (*big.Int, error) {
	return _SwapMining.Contract.PoolLength(&_SwapMining.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() view returns(uint256)
func (_SwapMining *SwapMiningCaller) Reward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "reward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() view returns(uint256)
func (_SwapMining *SwapMiningSession) Reward() (*big.Int, error) {
	return _SwapMining.Contract.Reward(&_SwapMining.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0x228cb733.
//
// Solidity: function reward() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) Reward() (*big.Int, error) {
	return _SwapMining.Contract.Reward(&_SwapMining.CallOpts)
}

// Reward0 is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 blockNumber) view returns(uint256)
func (_SwapMining *SwapMiningCaller) Reward0(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "reward0", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reward0 is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 blockNumber) view returns(uint256)
func (_SwapMining *SwapMiningSession) Reward0(blockNumber *big.Int) (*big.Int, error) {
	return _SwapMining.Contract.Reward0(&_SwapMining.CallOpts, blockNumber)
}

// Reward0 is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 blockNumber) view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) Reward0(blockNumber *big.Int) (*big.Int, error) {
	return _SwapMining.Contract.Reward0(&_SwapMining.CallOpts, blockNumber)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_SwapMining *SwapMiningCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_SwapMining *SwapMiningSession) Router() (common.Address, error) {
	return _SwapMining.Contract.Router(&_SwapMining.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_SwapMining *SwapMiningCallerSession) Router() (common.Address, error) {
	return _SwapMining.Contract.Router(&_SwapMining.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SwapMining *SwapMiningCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SwapMining *SwapMiningSession) StartBlock() (*big.Int, error) {
	return _SwapMining.Contract.StartBlock(&_SwapMining.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) StartBlock() (*big.Int, error) {
	return _SwapMining.Contract.StartBlock(&_SwapMining.CallOpts)
}

// TargetToken is a free data retrieval call binding the contract method 0x327107f7.
//
// Solidity: function targetToken() view returns(address)
func (_SwapMining *SwapMiningCaller) TargetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "targetToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetToken is a free data retrieval call binding the contract method 0x327107f7.
//
// Solidity: function targetToken() view returns(address)
func (_SwapMining *SwapMiningSession) TargetToken() (common.Address, error) {
	return _SwapMining.Contract.TargetToken(&_SwapMining.CallOpts)
}

// TargetToken is a free data retrieval call binding the contract method 0x327107f7.
//
// Solidity: function targetToken() view returns(address)
func (_SwapMining *SwapMiningCallerSession) TargetToken() (common.Address, error) {
	return _SwapMining.Contract.TargetToken(&_SwapMining.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SwapMining *SwapMiningCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SwapMining *SwapMiningSession) TotalAllocPoint() (*big.Int, error) {
	return _SwapMining.Contract.TotalAllocPoint(&_SwapMining.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SwapMining *SwapMiningCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _SwapMining.Contract.TotalAllocPoint(&_SwapMining.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 quantity, uint256 blockNumber)
func (_SwapMining *SwapMiningCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Quantity    *big.Int
	BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _SwapMining.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Quantity    *big.Int
		BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quantity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 quantity, uint256 blockNumber)
func (_SwapMining *SwapMiningSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Quantity    *big.Int
	BlockNumber *big.Int
}, error) {
	return _SwapMining.Contract.UserInfo(&_SwapMining.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 quantity, uint256 blockNumber)
func (_SwapMining *SwapMiningCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Quantity    *big.Int
	BlockNumber *big.Int
}, error) {
	return _SwapMining.Contract.UserInfo(&_SwapMining.CallOpts, arg0, arg1)
}

// AddPair is a paid mutator transaction binding the contract method 0xf9d181f9.
//
// Solidity: function addPair(uint256 _allocPoint, address _pair, bool _withUpdate) returns()
func (_SwapMining *SwapMiningTransactor) AddPair(opts *bind.TransactOpts, _allocPoint *big.Int, _pair common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "addPair", _allocPoint, _pair, _withUpdate)
}

// AddPair is a paid mutator transaction binding the contract method 0xf9d181f9.
//
// Solidity: function addPair(uint256 _allocPoint, address _pair, bool _withUpdate) returns()
func (_SwapMining *SwapMiningSession) AddPair(_allocPoint *big.Int, _pair common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _SwapMining.Contract.AddPair(&_SwapMining.TransactOpts, _allocPoint, _pair, _withUpdate)
}

// AddPair is a paid mutator transaction binding the contract method 0xf9d181f9.
//
// Solidity: function addPair(uint256 _allocPoint, address _pair, bool _withUpdate) returns()
func (_SwapMining *SwapMiningTransactorSession) AddPair(_allocPoint *big.Int, _pair common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _SwapMining.Contract.AddPair(&_SwapMining.TransactOpts, _allocPoint, _pair, _withUpdate)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address _addToken) returns(bool)
func (_SwapMining *SwapMiningTransactor) AddWhitelist(opts *bind.TransactOpts, _addToken common.Address) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "addWhitelist", _addToken)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address _addToken) returns(bool)
func (_SwapMining *SwapMiningSession) AddWhitelist(_addToken common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.AddWhitelist(&_SwapMining.TransactOpts, _addToken)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address _addToken) returns(bool)
func (_SwapMining *SwapMiningTransactorSession) AddWhitelist(_addToken common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.AddWhitelist(&_SwapMining.TransactOpts, _addToken)
}

// DelWhitelist is a paid mutator transaction binding the contract method 0x97ecfaab.
//
// Solidity: function delWhitelist(address _delToken) returns(bool)
func (_SwapMining *SwapMiningTransactor) DelWhitelist(opts *bind.TransactOpts, _delToken common.Address) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "delWhitelist", _delToken)
}

// DelWhitelist is a paid mutator transaction binding the contract method 0x97ecfaab.
//
// Solidity: function delWhitelist(address _delToken) returns(bool)
func (_SwapMining *SwapMiningSession) DelWhitelist(_delToken common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.DelWhitelist(&_SwapMining.TransactOpts, _delToken)
}

// DelWhitelist is a paid mutator transaction binding the contract method 0x97ecfaab.
//
// Solidity: function delWhitelist(address _delToken) returns(bool)
func (_SwapMining *SwapMiningTransactorSession) DelWhitelist(_delToken common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.DelWhitelist(&_SwapMining.TransactOpts, _delToken)
}

// MassMintPools is a paid mutator transaction binding the contract method 0x9f4ece2e.
//
// Solidity: function massMintPools() returns()
func (_SwapMining *SwapMiningTransactor) MassMintPools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "massMintPools")
}

// MassMintPools is a paid mutator transaction binding the contract method 0x9f4ece2e.
//
// Solidity: function massMintPools() returns()
func (_SwapMining *SwapMiningSession) MassMintPools() (*types.Transaction, error) {
	return _SwapMining.Contract.MassMintPools(&_SwapMining.TransactOpts)
}

// MassMintPools is a paid mutator transaction binding the contract method 0x9f4ece2e.
//
// Solidity: function massMintPools() returns()
func (_SwapMining *SwapMiningTransactorSession) MassMintPools() (*types.Transaction, error) {
	return _SwapMining.Contract.MassMintPools(&_SwapMining.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _pid) returns(bool)
func (_SwapMining *SwapMiningTransactor) Mint(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "mint", _pid)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _pid) returns(bool)
func (_SwapMining *SwapMiningSession) Mint(_pid *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.Mint(&_SwapMining.TransactOpts, _pid)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _pid) returns(bool)
func (_SwapMining *SwapMiningTransactorSession) Mint(_pid *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.Mint(&_SwapMining.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapMining *SwapMiningTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapMining *SwapMiningSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapMining.Contract.RenounceOwnership(&_SwapMining.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapMining *SwapMiningTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapMining.Contract.RenounceOwnership(&_SwapMining.TransactOpts)
}

// SetHalvingPeriod is a paid mutator transaction binding the contract method 0xb5ec5c99.
//
// Solidity: function setHalvingPeriod(uint256 _block) returns()
func (_SwapMining *SwapMiningTransactor) SetHalvingPeriod(opts *bind.TransactOpts, _block *big.Int) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "setHalvingPeriod", _block)
}

// SetHalvingPeriod is a paid mutator transaction binding the contract method 0xb5ec5c99.
//
// Solidity: function setHalvingPeriod(uint256 _block) returns()
func (_SwapMining *SwapMiningSession) SetHalvingPeriod(_block *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.SetHalvingPeriod(&_SwapMining.TransactOpts, _block)
}

// SetHalvingPeriod is a paid mutator transaction binding the contract method 0xb5ec5c99.
//
// Solidity: function setHalvingPeriod(uint256 _block) returns()
func (_SwapMining *SwapMiningTransactorSession) SetHalvingPeriod(_block *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.SetHalvingPeriod(&_SwapMining.TransactOpts, _block)
}

// SetMdxPerBlock is a paid mutator transaction binding the contract method 0x791ba374.
//
// Solidity: function setMdxPerBlock(uint256 _newPerBlock) returns()
func (_SwapMining *SwapMiningTransactor) SetMdxPerBlock(opts *bind.TransactOpts, _newPerBlock *big.Int) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "setMdxPerBlock", _newPerBlock)
}

// SetMdxPerBlock is a paid mutator transaction binding the contract method 0x791ba374.
//
// Solidity: function setMdxPerBlock(uint256 _newPerBlock) returns()
func (_SwapMining *SwapMiningSession) SetMdxPerBlock(_newPerBlock *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.SetMdxPerBlock(&_SwapMining.TransactOpts, _newPerBlock)
}

// SetMdxPerBlock is a paid mutator transaction binding the contract method 0x791ba374.
//
// Solidity: function setMdxPerBlock(uint256 _newPerBlock) returns()
func (_SwapMining *SwapMiningTransactorSession) SetMdxPerBlock(_newPerBlock *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.SetMdxPerBlock(&_SwapMining.TransactOpts, _newPerBlock)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_SwapMining *SwapMiningTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_SwapMining *SwapMiningSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.SetOracle(&_SwapMining.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_SwapMining *SwapMiningTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.SetOracle(&_SwapMining.TransactOpts, _oracle)
}

// SetPair is a paid mutator transaction binding the contract method 0x190cf6c9.
//
// Solidity: function setPair(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_SwapMining *SwapMiningTransactor) SetPair(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "setPair", _pid, _allocPoint, _withUpdate)
}

// SetPair is a paid mutator transaction binding the contract method 0x190cf6c9.
//
// Solidity: function setPair(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_SwapMining *SwapMiningSession) SetPair(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _SwapMining.Contract.SetPair(&_SwapMining.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetPair is a paid mutator transaction binding the contract method 0x190cf6c9.
//
// Solidity: function setPair(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_SwapMining *SwapMiningTransactorSession) SetPair(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _SwapMining.Contract.SetPair(&_SwapMining.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_SwapMining *SwapMiningTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "setRouter", newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_SwapMining *SwapMiningSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.SetRouter(&_SwapMining.TransactOpts, newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_SwapMining *SwapMiningTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.SetRouter(&_SwapMining.TransactOpts, newRouter)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address account, address input, address output, uint256 amount) returns(bool)
func (_SwapMining *SwapMiningTransactor) Swap(opts *bind.TransactOpts, account common.Address, input common.Address, output common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "swap", account, input, output, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address account, address input, address output, uint256 amount) returns(bool)
func (_SwapMining *SwapMiningSession) Swap(account common.Address, input common.Address, output common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.Swap(&_SwapMining.TransactOpts, account, input, output, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xa9678a18.
//
// Solidity: function swap(address account, address input, address output, uint256 amount) returns(bool)
func (_SwapMining *SwapMiningTransactorSession) Swap(account common.Address, input common.Address, output common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SwapMining.Contract.Swap(&_SwapMining.TransactOpts, account, input, output, amount)
}

// TakerWithdraw is a paid mutator transaction binding the contract method 0xb872dd0e.
//
// Solidity: function takerWithdraw() returns()
func (_SwapMining *SwapMiningTransactor) TakerWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "takerWithdraw")
}

// TakerWithdraw is a paid mutator transaction binding the contract method 0xb872dd0e.
//
// Solidity: function takerWithdraw() returns()
func (_SwapMining *SwapMiningSession) TakerWithdraw() (*types.Transaction, error) {
	return _SwapMining.Contract.TakerWithdraw(&_SwapMining.TransactOpts)
}

// TakerWithdraw is a paid mutator transaction binding the contract method 0xb872dd0e.
//
// Solidity: function takerWithdraw() returns()
func (_SwapMining *SwapMiningTransactorSession) TakerWithdraw() (*types.Transaction, error) {
	return _SwapMining.Contract.TakerWithdraw(&_SwapMining.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapMining *SwapMiningTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapMining.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapMining *SwapMiningSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.TransferOwnership(&_SwapMining.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapMining *SwapMiningTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapMining.Contract.TransferOwnership(&_SwapMining.TransactOpts, newOwner)
}

// SwapMiningOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapMining contract.
type SwapMiningOwnershipTransferredIterator struct {
	Event *SwapMiningOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapMiningOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapMiningOwnershipTransferred)
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
		it.Event = new(SwapMiningOwnershipTransferred)
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
func (it *SwapMiningOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapMiningOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapMiningOwnershipTransferred represents a OwnershipTransferred event raised by the SwapMining contract.
type SwapMiningOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapMining *SwapMiningFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapMiningOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapMining.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapMiningOwnershipTransferredIterator{contract: _SwapMining.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapMining *SwapMiningFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapMiningOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapMining.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapMiningOwnershipTransferred)
				if err := _SwapMining.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapMining *SwapMiningFilterer) ParseOwnershipTransferred(log types.Log) (*SwapMiningOwnershipTransferred, error) {
	event := new(SwapMiningOwnershipTransferred)
	if err := _SwapMining.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
