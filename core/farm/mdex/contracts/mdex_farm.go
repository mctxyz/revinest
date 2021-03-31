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

// MdexFarmABI is the input ABI used to generate the binding from.
const MdexFarmABI = "[{\"inputs\":[{\"internalType\":\"contractIMdx\",\"name\":\"_mdx\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mdxPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"LpOfPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addLP\",\"type\":\"address\"}],\"name\":\"addMultLP\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastRewardBlock\",\"type\":\"uint256\"}],\"name\":\"getMdxBlockReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getMultLPAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMultLPLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halvingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_LP\",\"type\":\"address\"}],\"name\":\"isMultLP\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdx\",\"outputs\":[{\"internalType\":\"contractIMdx\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdxPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multLpChef\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multLpToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolCorrespond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMdxPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMultLpPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_multLpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_multLpChef\",\"type\":\"address\"}],\"name\":\"replaceMultLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"setHalvingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPerBlock\",\"type\":\"uint256\"}],\"name\":\"setMdxPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_multLpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_multLpChef\",\"type\":\"address\"}],\"name\":\"setMultLP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sid\",\"type\":\"uint256\"}],\"name\":\"setPoolCorr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multLpRewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MdexFarm is an auto generated Go binding around an Ethereum contract.
type MdexFarm struct {
	MdexFarmCaller     // Read-only binding to the contract
	MdexFarmTransactor // Write-only binding to the contract
	MdexFarmFilterer   // Log filterer for contract events
}

// MdexFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type MdexFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MdexFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MdexFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MdexFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MdexFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MdexFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MdexFarmSession struct {
	Contract     *MdexFarm         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MdexFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MdexFarmCallerSession struct {
	Contract *MdexFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MdexFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MdexFarmTransactorSession struct {
	Contract     *MdexFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MdexFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type MdexFarmRaw struct {
	Contract *MdexFarm // Generic contract binding to access the raw methods on
}

// MdexFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MdexFarmCallerRaw struct {
	Contract *MdexFarmCaller // Generic read-only contract binding to access the raw methods on
}

// MdexFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MdexFarmTransactorRaw struct {
	Contract *MdexFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMdexFarm creates a new instance of MdexFarm, bound to a specific deployed contract.
func NewMdexFarm(address common.Address, backend bind.ContractBackend) (*MdexFarm, error) {
	contract, err := bindMdexFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MdexFarm{MdexFarmCaller: MdexFarmCaller{contract: contract}, MdexFarmTransactor: MdexFarmTransactor{contract: contract}, MdexFarmFilterer: MdexFarmFilterer{contract: contract}}, nil
}

// NewMdexFarmCaller creates a new read-only instance of MdexFarm, bound to a specific deployed contract.
func NewMdexFarmCaller(address common.Address, caller bind.ContractCaller) (*MdexFarmCaller, error) {
	contract, err := bindMdexFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MdexFarmCaller{contract: contract}, nil
}

// NewMdexFarmTransactor creates a new write-only instance of MdexFarm, bound to a specific deployed contract.
func NewMdexFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*MdexFarmTransactor, error) {
	contract, err := bindMdexFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MdexFarmTransactor{contract: contract}, nil
}

// NewMdexFarmFilterer creates a new log filterer instance of MdexFarm, bound to a specific deployed contract.
func NewMdexFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*MdexFarmFilterer, error) {
	contract, err := bindMdexFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MdexFarmFilterer{contract: contract}, nil
}

// bindMdexFarm binds a generic wrapper to an already deployed contract.
func bindMdexFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MdexFarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MdexFarm *MdexFarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MdexFarm.Contract.MdexFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MdexFarm *MdexFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MdexFarm.Contract.MdexFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MdexFarm *MdexFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MdexFarm.Contract.MdexFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MdexFarm *MdexFarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MdexFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MdexFarm *MdexFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MdexFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MdexFarm *MdexFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MdexFarm.Contract.contract.Transact(opts, method, params...)
}

// LpOfPid is a free data retrieval call binding the contract method 0xb0c7044b.
//
// Solidity: function LpOfPid(address ) view returns(uint256)
func (_MdexFarm *MdexFarmCaller) LpOfPid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "LpOfPid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpOfPid is a free data retrieval call binding the contract method 0xb0c7044b.
//
// Solidity: function LpOfPid(address ) view returns(uint256)
func (_MdexFarm *MdexFarmSession) LpOfPid(arg0 common.Address) (*big.Int, error) {
	return _MdexFarm.Contract.LpOfPid(&_MdexFarm.CallOpts, arg0)
}

// LpOfPid is a free data retrieval call binding the contract method 0xb0c7044b.
//
// Solidity: function LpOfPid(address ) view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) LpOfPid(arg0 common.Address) (*big.Int, error) {
	return _MdexFarm.Contract.LpOfPid(&_MdexFarm.CallOpts, arg0)
}

// GetMdxBlockReward is a free data retrieval call binding the contract method 0x1b15b579.
//
// Solidity: function getMdxBlockReward(uint256 _lastRewardBlock) view returns(uint256)
func (_MdexFarm *MdexFarmCaller) GetMdxBlockReward(opts *bind.CallOpts, _lastRewardBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "getMdxBlockReward", _lastRewardBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMdxBlockReward is a free data retrieval call binding the contract method 0x1b15b579.
//
// Solidity: function getMdxBlockReward(uint256 _lastRewardBlock) view returns(uint256)
func (_MdexFarm *MdexFarmSession) GetMdxBlockReward(_lastRewardBlock *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.GetMdxBlockReward(&_MdexFarm.CallOpts, _lastRewardBlock)
}

// GetMdxBlockReward is a free data retrieval call binding the contract method 0x1b15b579.
//
// Solidity: function getMdxBlockReward(uint256 _lastRewardBlock) view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) GetMdxBlockReward(_lastRewardBlock *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.GetMdxBlockReward(&_MdexFarm.CallOpts, _lastRewardBlock)
}

// GetMultLPAddress is a free data retrieval call binding the contract method 0x43154907.
//
// Solidity: function getMultLPAddress(uint256 _pid) view returns(address)
func (_MdexFarm *MdexFarmCaller) GetMultLPAddress(opts *bind.CallOpts, _pid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "getMultLPAddress", _pid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMultLPAddress is a free data retrieval call binding the contract method 0x43154907.
//
// Solidity: function getMultLPAddress(uint256 _pid) view returns(address)
func (_MdexFarm *MdexFarmSession) GetMultLPAddress(_pid *big.Int) (common.Address, error) {
	return _MdexFarm.Contract.GetMultLPAddress(&_MdexFarm.CallOpts, _pid)
}

// GetMultLPAddress is a free data retrieval call binding the contract method 0x43154907.
//
// Solidity: function getMultLPAddress(uint256 _pid) view returns(address)
func (_MdexFarm *MdexFarmCallerSession) GetMultLPAddress(_pid *big.Int) (common.Address, error) {
	return _MdexFarm.Contract.GetMultLPAddress(&_MdexFarm.CallOpts, _pid)
}

// GetMultLPLength is a free data retrieval call binding the contract method 0xe6ab8d48.
//
// Solidity: function getMultLPLength() view returns(uint256)
func (_MdexFarm *MdexFarmCaller) GetMultLPLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "getMultLPLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultLPLength is a free data retrieval call binding the contract method 0xe6ab8d48.
//
// Solidity: function getMultLPLength() view returns(uint256)
func (_MdexFarm *MdexFarmSession) GetMultLPLength() (*big.Int, error) {
	return _MdexFarm.Contract.GetMultLPLength(&_MdexFarm.CallOpts)
}

// GetMultLPLength is a free data retrieval call binding the contract method 0xe6ab8d48.
//
// Solidity: function getMultLPLength() view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) GetMultLPLength() (*big.Int, error) {
	return _MdexFarm.Contract.GetMultLPLength(&_MdexFarm.CallOpts)
}

// HalvingPeriod is a free data retrieval call binding the contract method 0x5a3e251f.
//
// Solidity: function halvingPeriod() view returns(uint256)
func (_MdexFarm *MdexFarmCaller) HalvingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "halvingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalvingPeriod is a free data retrieval call binding the contract method 0x5a3e251f.
//
// Solidity: function halvingPeriod() view returns(uint256)
func (_MdexFarm *MdexFarmSession) HalvingPeriod() (*big.Int, error) {
	return _MdexFarm.Contract.HalvingPeriod(&_MdexFarm.CallOpts)
}

// HalvingPeriod is a free data retrieval call binding the contract method 0x5a3e251f.
//
// Solidity: function halvingPeriod() view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) HalvingPeriod() (*big.Int, error) {
	return _MdexFarm.Contract.HalvingPeriod(&_MdexFarm.CallOpts)
}

// IsMultLP is a free data retrieval call binding the contract method 0x7fe6f5ac.
//
// Solidity: function isMultLP(address _LP) view returns(bool)
func (_MdexFarm *MdexFarmCaller) IsMultLP(opts *bind.CallOpts, _LP common.Address) (bool, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "isMultLP", _LP)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMultLP is a free data retrieval call binding the contract method 0x7fe6f5ac.
//
// Solidity: function isMultLP(address _LP) view returns(bool)
func (_MdexFarm *MdexFarmSession) IsMultLP(_LP common.Address) (bool, error) {
	return _MdexFarm.Contract.IsMultLP(&_MdexFarm.CallOpts, _LP)
}

// IsMultLP is a free data retrieval call binding the contract method 0x7fe6f5ac.
//
// Solidity: function isMultLP(address _LP) view returns(bool)
func (_MdexFarm *MdexFarmCallerSession) IsMultLP(_LP common.Address) (bool, error) {
	return _MdexFarm.Contract.IsMultLP(&_MdexFarm.CallOpts, _LP)
}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_MdexFarm *MdexFarmCaller) Mdx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "mdx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_MdexFarm *MdexFarmSession) Mdx() (common.Address, error) {
	return _MdexFarm.Contract.Mdx(&_MdexFarm.CallOpts)
}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_MdexFarm *MdexFarmCallerSession) Mdx() (common.Address, error) {
	return _MdexFarm.Contract.Mdx(&_MdexFarm.CallOpts)
}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_MdexFarm *MdexFarmCaller) MdxPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "mdxPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_MdexFarm *MdexFarmSession) MdxPerBlock() (*big.Int, error) {
	return _MdexFarm.Contract.MdxPerBlock(&_MdexFarm.CallOpts)
}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) MdxPerBlock() (*big.Int, error) {
	return _MdexFarm.Contract.MdxPerBlock(&_MdexFarm.CallOpts)
}

// MultLpChef is a free data retrieval call binding the contract method 0xdfc7b95b.
//
// Solidity: function multLpChef() view returns(address)
func (_MdexFarm *MdexFarmCaller) MultLpChef(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "multLpChef")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultLpChef is a free data retrieval call binding the contract method 0xdfc7b95b.
//
// Solidity: function multLpChef() view returns(address)
func (_MdexFarm *MdexFarmSession) MultLpChef() (common.Address, error) {
	return _MdexFarm.Contract.MultLpChef(&_MdexFarm.CallOpts)
}

// MultLpChef is a free data retrieval call binding the contract method 0xdfc7b95b.
//
// Solidity: function multLpChef() view returns(address)
func (_MdexFarm *MdexFarmCallerSession) MultLpChef() (common.Address, error) {
	return _MdexFarm.Contract.MultLpChef(&_MdexFarm.CallOpts)
}

// MultLpToken is a free data retrieval call binding the contract method 0x705bbc01.
//
// Solidity: function multLpToken() view returns(address)
func (_MdexFarm *MdexFarmCaller) MultLpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "multLpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultLpToken is a free data retrieval call binding the contract method 0x705bbc01.
//
// Solidity: function multLpToken() view returns(address)
func (_MdexFarm *MdexFarmSession) MultLpToken() (common.Address, error) {
	return _MdexFarm.Contract.MultLpToken(&_MdexFarm.CallOpts)
}

// MultLpToken is a free data retrieval call binding the contract method 0x705bbc01.
//
// Solidity: function multLpToken() view returns(address)
func (_MdexFarm *MdexFarmCallerSession) MultLpToken() (common.Address, error) {
	return _MdexFarm.Contract.MultLpToken(&_MdexFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MdexFarm *MdexFarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MdexFarm *MdexFarmSession) Owner() (common.Address, error) {
	return _MdexFarm.Contract.Owner(&_MdexFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MdexFarm *MdexFarmCallerSession) Owner() (common.Address, error) {
	return _MdexFarm.Contract.Owner(&_MdexFarm.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MdexFarm *MdexFarmCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MdexFarm *MdexFarmSession) Paused() (bool, error) {
	return _MdexFarm.Contract.Paused(&_MdexFarm.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MdexFarm *MdexFarmCallerSession) Paused() (bool, error) {
	return _MdexFarm.Contract.Paused(&_MdexFarm.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256, uint256)
func (_MdexFarm *MdexFarmCaller) Pending(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "pending", _pid, _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256, uint256)
func (_MdexFarm *MdexFarmSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, *big.Int, error) {
	return _MdexFarm.Contract.Pending(&_MdexFarm.CallOpts, _pid, _user)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256, uint256)
func (_MdexFarm *MdexFarmCallerSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, *big.Int, error) {
	return _MdexFarm.Contract.Pending(&_MdexFarm.CallOpts, _pid, _user)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 blockNumber) view returns(uint256)
func (_MdexFarm *MdexFarmCaller) Phase(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "phase", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 blockNumber) view returns(uint256)
func (_MdexFarm *MdexFarmSession) Phase(blockNumber *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.Phase(&_MdexFarm.CallOpts, blockNumber)
}

// Phase is a free data retrieval call binding the contract method 0x135f8aa7.
//
// Solidity: function phase(uint256 blockNumber) view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) Phase(blockNumber *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.Phase(&_MdexFarm.CallOpts, blockNumber)
}

// PoolCorrespond is a free data retrieval call binding the contract method 0xcb4502c4.
//
// Solidity: function poolCorrespond(uint256 ) view returns(uint256)
func (_MdexFarm *MdexFarmCaller) PoolCorrespond(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "poolCorrespond", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCorrespond is a free data retrieval call binding the contract method 0xcb4502c4.
//
// Solidity: function poolCorrespond(uint256 ) view returns(uint256)
func (_MdexFarm *MdexFarmSession) PoolCorrespond(arg0 *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.PoolCorrespond(&_MdexFarm.CallOpts, arg0)
}

// PoolCorrespond is a free data retrieval call binding the contract method 0xcb4502c4.
//
// Solidity: function poolCorrespond(uint256 ) view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) PoolCorrespond(arg0 *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.PoolCorrespond(&_MdexFarm.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accMdxPerShare, uint256 accMultLpPerShare, uint256 totalAmount)
func (_MdexFarm *MdexFarmCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken           common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccMdxPerShare    *big.Int
	AccMultLpPerShare *big.Int
	TotalAmount       *big.Int
}, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken           common.Address
		AllocPoint        *big.Int
		LastRewardBlock   *big.Int
		AccMdxPerShare    *big.Int
		AccMultLpPerShare *big.Int
		TotalAmount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccMdxPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccMultLpPerShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accMdxPerShare, uint256 accMultLpPerShare, uint256 totalAmount)
func (_MdexFarm *MdexFarmSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken           common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccMdxPerShare    *big.Int
	AccMultLpPerShare *big.Int
	TotalAmount       *big.Int
}, error) {
	return _MdexFarm.Contract.PoolInfo(&_MdexFarm.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accMdxPerShare, uint256 accMultLpPerShare, uint256 totalAmount)
func (_MdexFarm *MdexFarmCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken           common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccMdxPerShare    *big.Int
	AccMultLpPerShare *big.Int
	TotalAmount       *big.Int
}, error) {
	return _MdexFarm.Contract.PoolInfo(&_MdexFarm.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MdexFarm *MdexFarmCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MdexFarm *MdexFarmSession) PoolLength() (*big.Int, error) {
	return _MdexFarm.Contract.PoolLength(&_MdexFarm.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) PoolLength() (*big.Int, error) {
	return _MdexFarm.Contract.PoolLength(&_MdexFarm.CallOpts)
}

// Reward is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 blockNumber) view returns(uint256)
func (_MdexFarm *MdexFarmCaller) Reward(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "reward", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reward is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 blockNumber) view returns(uint256)
func (_MdexFarm *MdexFarmSession) Reward(blockNumber *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.Reward(&_MdexFarm.CallOpts, blockNumber)
}

// Reward is a free data retrieval call binding the contract method 0xa9fb763c.
//
// Solidity: function reward(uint256 blockNumber) view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) Reward(blockNumber *big.Int) (*big.Int, error) {
	return _MdexFarm.Contract.Reward(&_MdexFarm.CallOpts, blockNumber)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_MdexFarm *MdexFarmCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_MdexFarm *MdexFarmSession) StartBlock() (*big.Int, error) {
	return _MdexFarm.Contract.StartBlock(&_MdexFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) StartBlock() (*big.Int, error) {
	return _MdexFarm.Contract.StartBlock(&_MdexFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MdexFarm *MdexFarmCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MdexFarm *MdexFarmSession) TotalAllocPoint() (*big.Int, error) {
	return _MdexFarm.Contract.TotalAllocPoint(&_MdexFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_MdexFarm *MdexFarmCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _MdexFarm.Contract.TotalAllocPoint(&_MdexFarm.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 multLpRewardDebt)
func (_MdexFarm *MdexFarmCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	RewardDebt       *big.Int
	MultLpRewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _MdexFarm.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount           *big.Int
		RewardDebt       *big.Int
		MultLpRewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MultLpRewardDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 multLpRewardDebt)
func (_MdexFarm *MdexFarmSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	RewardDebt       *big.Int
	MultLpRewardDebt *big.Int
}, error) {
	return _MdexFarm.Contract.UserInfo(&_MdexFarm.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 multLpRewardDebt)
func (_MdexFarm *MdexFarmCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount           *big.Int
	RewardDebt       *big.Int
	MultLpRewardDebt *big.Int
}, error) {
	return _MdexFarm.Contract.UserInfo(&_MdexFarm.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_MdexFarm *MdexFarmTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_MdexFarm *MdexFarmSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _MdexFarm.Contract.Add(&_MdexFarm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_MdexFarm *MdexFarmTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _MdexFarm.Contract.Add(&_MdexFarm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// AddMultLP is a paid mutator transaction binding the contract method 0x56c5867d.
//
// Solidity: function addMultLP(address _addLP) returns(bool)
func (_MdexFarm *MdexFarmTransactor) AddMultLP(opts *bind.TransactOpts, _addLP common.Address) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "addMultLP", _addLP)
}

// AddMultLP is a paid mutator transaction binding the contract method 0x56c5867d.
//
// Solidity: function addMultLP(address _addLP) returns(bool)
func (_MdexFarm *MdexFarmSession) AddMultLP(_addLP common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.AddMultLP(&_MdexFarm.TransactOpts, _addLP)
}

// AddMultLP is a paid mutator transaction binding the contract method 0x56c5867d.
//
// Solidity: function addMultLP(address _addLP) returns(bool)
func (_MdexFarm *MdexFarmTransactorSession) AddMultLP(_addLP common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.AddMultLP(&_MdexFarm.TransactOpts, _addLP)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MdexFarm *MdexFarmTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MdexFarm *MdexFarmSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.Deposit(&_MdexFarm.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MdexFarm *MdexFarmTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.Deposit(&_MdexFarm.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MdexFarm *MdexFarmTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MdexFarm *MdexFarmSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.EmergencyWithdraw(&_MdexFarm.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MdexFarm *MdexFarmTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.EmergencyWithdraw(&_MdexFarm.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MdexFarm *MdexFarmTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MdexFarm *MdexFarmSession) MassUpdatePools() (*types.Transaction, error) {
	return _MdexFarm.Contract.MassUpdatePools(&_MdexFarm.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MdexFarm *MdexFarmTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _MdexFarm.Contract.MassUpdatePools(&_MdexFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MdexFarm *MdexFarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MdexFarm *MdexFarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _MdexFarm.Contract.RenounceOwnership(&_MdexFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MdexFarm *MdexFarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MdexFarm.Contract.RenounceOwnership(&_MdexFarm.TransactOpts)
}

// ReplaceMultLP is a paid mutator transaction binding the contract method 0xaaae43cc.
//
// Solidity: function replaceMultLP(address _multLpToken, address _multLpChef) returns()
func (_MdexFarm *MdexFarmTransactor) ReplaceMultLP(opts *bind.TransactOpts, _multLpToken common.Address, _multLpChef common.Address) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "replaceMultLP", _multLpToken, _multLpChef)
}

// ReplaceMultLP is a paid mutator transaction binding the contract method 0xaaae43cc.
//
// Solidity: function replaceMultLP(address _multLpToken, address _multLpChef) returns()
func (_MdexFarm *MdexFarmSession) ReplaceMultLP(_multLpToken common.Address, _multLpChef common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.ReplaceMultLP(&_MdexFarm.TransactOpts, _multLpToken, _multLpChef)
}

// ReplaceMultLP is a paid mutator transaction binding the contract method 0xaaae43cc.
//
// Solidity: function replaceMultLP(address _multLpToken, address _multLpChef) returns()
func (_MdexFarm *MdexFarmTransactorSession) ReplaceMultLP(_multLpToken common.Address, _multLpChef common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.ReplaceMultLP(&_MdexFarm.TransactOpts, _multLpToken, _multLpChef)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_MdexFarm *MdexFarmTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_MdexFarm *MdexFarmSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MdexFarm.Contract.Set(&_MdexFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_MdexFarm *MdexFarmTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MdexFarm.Contract.Set(&_MdexFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetHalvingPeriod is a paid mutator transaction binding the contract method 0xb5ec5c99.
//
// Solidity: function setHalvingPeriod(uint256 _block) returns()
func (_MdexFarm *MdexFarmTransactor) SetHalvingPeriod(opts *bind.TransactOpts, _block *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "setHalvingPeriod", _block)
}

// SetHalvingPeriod is a paid mutator transaction binding the contract method 0xb5ec5c99.
//
// Solidity: function setHalvingPeriod(uint256 _block) returns()
func (_MdexFarm *MdexFarmSession) SetHalvingPeriod(_block *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetHalvingPeriod(&_MdexFarm.TransactOpts, _block)
}

// SetHalvingPeriod is a paid mutator transaction binding the contract method 0xb5ec5c99.
//
// Solidity: function setHalvingPeriod(uint256 _block) returns()
func (_MdexFarm *MdexFarmTransactorSession) SetHalvingPeriod(_block *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetHalvingPeriod(&_MdexFarm.TransactOpts, _block)
}

// SetMdxPerBlock is a paid mutator transaction binding the contract method 0x791ba374.
//
// Solidity: function setMdxPerBlock(uint256 _newPerBlock) returns()
func (_MdexFarm *MdexFarmTransactor) SetMdxPerBlock(opts *bind.TransactOpts, _newPerBlock *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "setMdxPerBlock", _newPerBlock)
}

// SetMdxPerBlock is a paid mutator transaction binding the contract method 0x791ba374.
//
// Solidity: function setMdxPerBlock(uint256 _newPerBlock) returns()
func (_MdexFarm *MdexFarmSession) SetMdxPerBlock(_newPerBlock *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetMdxPerBlock(&_MdexFarm.TransactOpts, _newPerBlock)
}

// SetMdxPerBlock is a paid mutator transaction binding the contract method 0x791ba374.
//
// Solidity: function setMdxPerBlock(uint256 _newPerBlock) returns()
func (_MdexFarm *MdexFarmTransactorSession) SetMdxPerBlock(_newPerBlock *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetMdxPerBlock(&_MdexFarm.TransactOpts, _newPerBlock)
}

// SetMultLP is a paid mutator transaction binding the contract method 0xe715e234.
//
// Solidity: function setMultLP(address _multLpToken, address _multLpChef) returns()
func (_MdexFarm *MdexFarmTransactor) SetMultLP(opts *bind.TransactOpts, _multLpToken common.Address, _multLpChef common.Address) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "setMultLP", _multLpToken, _multLpChef)
}

// SetMultLP is a paid mutator transaction binding the contract method 0xe715e234.
//
// Solidity: function setMultLP(address _multLpToken, address _multLpChef) returns()
func (_MdexFarm *MdexFarmSession) SetMultLP(_multLpToken common.Address, _multLpChef common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetMultLP(&_MdexFarm.TransactOpts, _multLpToken, _multLpChef)
}

// SetMultLP is a paid mutator transaction binding the contract method 0xe715e234.
//
// Solidity: function setMultLP(address _multLpToken, address _multLpChef) returns()
func (_MdexFarm *MdexFarmTransactorSession) SetMultLP(_multLpToken common.Address, _multLpChef common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetMultLP(&_MdexFarm.TransactOpts, _multLpToken, _multLpChef)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_MdexFarm *MdexFarmTransactor) SetPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "setPause")
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_MdexFarm *MdexFarmSession) SetPause() (*types.Transaction, error) {
	return _MdexFarm.Contract.SetPause(&_MdexFarm.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_MdexFarm *MdexFarmTransactorSession) SetPause() (*types.Transaction, error) {
	return _MdexFarm.Contract.SetPause(&_MdexFarm.TransactOpts)
}

// SetPoolCorr is a paid mutator transaction binding the contract method 0xb337d32c.
//
// Solidity: function setPoolCorr(uint256 _pid, uint256 _sid) returns()
func (_MdexFarm *MdexFarmTransactor) SetPoolCorr(opts *bind.TransactOpts, _pid *big.Int, _sid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "setPoolCorr", _pid, _sid)
}

// SetPoolCorr is a paid mutator transaction binding the contract method 0xb337d32c.
//
// Solidity: function setPoolCorr(uint256 _pid, uint256 _sid) returns()
func (_MdexFarm *MdexFarmSession) SetPoolCorr(_pid *big.Int, _sid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetPoolCorr(&_MdexFarm.TransactOpts, _pid, _sid)
}

// SetPoolCorr is a paid mutator transaction binding the contract method 0xb337d32c.
//
// Solidity: function setPoolCorr(uint256 _pid, uint256 _sid) returns()
func (_MdexFarm *MdexFarmTransactorSession) SetPoolCorr(_pid *big.Int, _sid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.SetPoolCorr(&_MdexFarm.TransactOpts, _pid, _sid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MdexFarm *MdexFarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MdexFarm *MdexFarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.TransferOwnership(&_MdexFarm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MdexFarm *MdexFarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MdexFarm.Contract.TransferOwnership(&_MdexFarm.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MdexFarm *MdexFarmTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MdexFarm *MdexFarmSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.UpdatePool(&_MdexFarm.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MdexFarm *MdexFarmTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.UpdatePool(&_MdexFarm.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MdexFarm *MdexFarmTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MdexFarm.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MdexFarm *MdexFarmSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.Withdraw(&_MdexFarm.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MdexFarm *MdexFarmTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MdexFarm.Contract.Withdraw(&_MdexFarm.TransactOpts, _pid, _amount)
}

// MdexFarmDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MdexFarm contract.
type MdexFarmDepositIterator struct {
	Event *MdexFarmDeposit // Event containing the contract specifics and raw log

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
func (it *MdexFarmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdexFarmDeposit)
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
		it.Event = new(MdexFarmDeposit)
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
func (it *MdexFarmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdexFarmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdexFarmDeposit represents a Deposit event raised by the MdexFarm contract.
type MdexFarmDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MdexFarm *MdexFarmFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MdexFarmDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MdexFarm.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MdexFarmDepositIterator{contract: _MdexFarm.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MdexFarm *MdexFarmFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MdexFarmDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MdexFarm.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdexFarmDeposit)
				if err := _MdexFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_MdexFarm *MdexFarmFilterer) ParseDeposit(log types.Log) (*MdexFarmDeposit, error) {
	event := new(MdexFarmDeposit)
	if err := _MdexFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdexFarmEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the MdexFarm contract.
type MdexFarmEmergencyWithdrawIterator struct {
	Event *MdexFarmEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MdexFarmEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdexFarmEmergencyWithdraw)
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
		it.Event = new(MdexFarmEmergencyWithdraw)
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
func (it *MdexFarmEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdexFarmEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdexFarmEmergencyWithdraw represents a EmergencyWithdraw event raised by the MdexFarm contract.
type MdexFarmEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MdexFarm *MdexFarmFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MdexFarmEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MdexFarm.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MdexFarmEmergencyWithdrawIterator{contract: _MdexFarm.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MdexFarm *MdexFarmFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MdexFarmEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MdexFarm.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdexFarmEmergencyWithdraw)
				if err := _MdexFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_MdexFarm *MdexFarmFilterer) ParseEmergencyWithdraw(log types.Log) (*MdexFarmEmergencyWithdraw, error) {
	event := new(MdexFarmEmergencyWithdraw)
	if err := _MdexFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdexFarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MdexFarm contract.
type MdexFarmOwnershipTransferredIterator struct {
	Event *MdexFarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MdexFarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdexFarmOwnershipTransferred)
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
		it.Event = new(MdexFarmOwnershipTransferred)
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
func (it *MdexFarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdexFarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdexFarmOwnershipTransferred represents a OwnershipTransferred event raised by the MdexFarm contract.
type MdexFarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MdexFarm *MdexFarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MdexFarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MdexFarm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MdexFarmOwnershipTransferredIterator{contract: _MdexFarm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MdexFarm *MdexFarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MdexFarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MdexFarm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdexFarmOwnershipTransferred)
				if err := _MdexFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MdexFarm *MdexFarmFilterer) ParseOwnershipTransferred(log types.Log) (*MdexFarmOwnershipTransferred, error) {
	event := new(MdexFarmOwnershipTransferred)
	if err := _MdexFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdexFarmWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MdexFarm contract.
type MdexFarmWithdrawIterator struct {
	Event *MdexFarmWithdraw // Event containing the contract specifics and raw log

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
func (it *MdexFarmWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdexFarmWithdraw)
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
		it.Event = new(MdexFarmWithdraw)
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
func (it *MdexFarmWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdexFarmWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdexFarmWithdraw represents a Withdraw event raised by the MdexFarm contract.
type MdexFarmWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MdexFarm *MdexFarmFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MdexFarmWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MdexFarm.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MdexFarmWithdrawIterator{contract: _MdexFarm.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MdexFarm *MdexFarmFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MdexFarmWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MdexFarm.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdexFarmWithdraw)
				if err := _MdexFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_MdexFarm *MdexFarmFilterer) ParseWithdraw(log types.Log) (*MdexFarmWithdraw, error) {
	event := new(MdexFarmWithdraw)
	if err := _MdexFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
