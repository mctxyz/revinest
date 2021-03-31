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

// PancakeFarmABI is the input ABI used to generate the binding from.
const PancakeFarmABI = "[{\"inputs\":[{\"internalType\":\"contractCakeToken\",\"name\":\"_cake\",\"type\":\"address\"},{\"internalType\":\"contractSyrupBar\",\"name\":\"_syrup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cakePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cake\",\"outputs\":[{\"internalType\":\"contractCakeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"enterStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"leaveStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingCake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syrup\",\"outputs\":[{\"internalType\":\"contractSyrupBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplierNumber\",\"type\":\"uint256\"}],\"name\":\"updateMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PancakeFarm is an auto generated Go binding around an Ethereum contract.
type PancakeFarm struct {
	PancakeFarmCaller     // Read-only binding to the contract
	PancakeFarmTransactor // Write-only binding to the contract
	PancakeFarmFilterer   // Log filterer for contract events
}

// PancakeFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeFarmSession struct {
	Contract     *PancakeFarm      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeFarmCallerSession struct {
	Contract *PancakeFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PancakeFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeFarmTransactorSession struct {
	Contract     *PancakeFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PancakeFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeFarmRaw struct {
	Contract *PancakeFarm // Generic contract binding to access the raw methods on
}

// PancakeFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeFarmCallerRaw struct {
	Contract *PancakeFarmCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeFarmTransactorRaw struct {
	Contract *PancakeFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeFarm creates a new instance of PancakeFarm, bound to a specific deployed contract.
func NewPancakeFarm(address common.Address, backend bind.ContractBackend) (*PancakeFarm, error) {
	contract, err := bindPancakeFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeFarm{PancakeFarmCaller: PancakeFarmCaller{contract: contract}, PancakeFarmTransactor: PancakeFarmTransactor{contract: contract}, PancakeFarmFilterer: PancakeFarmFilterer{contract: contract}}, nil
}

// NewPancakeFarmCaller creates a new read-only instance of PancakeFarm, bound to a specific deployed contract.
func NewPancakeFarmCaller(address common.Address, caller bind.ContractCaller) (*PancakeFarmCaller, error) {
	contract, err := bindPancakeFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmCaller{contract: contract}, nil
}

// NewPancakeFarmTransactor creates a new write-only instance of PancakeFarm, bound to a specific deployed contract.
func NewPancakeFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeFarmTransactor, error) {
	contract, err := bindPancakeFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmTransactor{contract: contract}, nil
}

// NewPancakeFarmFilterer creates a new log filterer instance of PancakeFarm, bound to a specific deployed contract.
func NewPancakeFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeFarmFilterer, error) {
	contract, err := bindPancakeFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmFilterer{contract: contract}, nil
}

// bindPancakeFarm binds a generic wrapper to an already deployed contract.
func bindPancakeFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeFarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeFarm *PancakeFarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeFarm.Contract.PancakeFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeFarm *PancakeFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeFarm.Contract.PancakeFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeFarm *PancakeFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeFarm.Contract.PancakeFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeFarm *PancakeFarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeFarm *PancakeFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeFarm *PancakeFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeFarm.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _PancakeFarm.Contract.BONUSMULTIPLIER(&_PancakeFarm.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _PancakeFarm.Contract.BONUSMULTIPLIER(&_PancakeFarm.CallOpts)
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() view returns(address)
func (_PancakeFarm *PancakeFarmCaller) Cake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "cake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() view returns(address)
func (_PancakeFarm *PancakeFarmSession) Cake() (common.Address, error) {
	return _PancakeFarm.Contract.Cake(&_PancakeFarm.CallOpts)
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() view returns(address)
func (_PancakeFarm *PancakeFarmCallerSession) Cake() (common.Address, error) {
	return _PancakeFarm.Contract.Cake(&_PancakeFarm.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) CakePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "cakePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) CakePerBlock() (*big.Int, error) {
	return _PancakeFarm.Contract.CakePerBlock(&_PancakeFarm.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) CakePerBlock() (*big.Int, error) {
	return _PancakeFarm.Contract.CakePerBlock(&_PancakeFarm.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_PancakeFarm *PancakeFarmCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_PancakeFarm *PancakeFarmSession) Devaddr() (common.Address, error) {
	return _PancakeFarm.Contract.Devaddr(&_PancakeFarm.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_PancakeFarm *PancakeFarmCallerSession) Devaddr() (common.Address, error) {
	return _PancakeFarm.Contract.Devaddr(&_PancakeFarm.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _PancakeFarm.Contract.GetMultiplier(&_PancakeFarm.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _PancakeFarm.Contract.GetMultiplier(&_PancakeFarm.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_PancakeFarm *PancakeFarmCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_PancakeFarm *PancakeFarmSession) Migrator() (common.Address, error) {
	return _PancakeFarm.Contract.Migrator(&_PancakeFarm.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_PancakeFarm *PancakeFarmCallerSession) Migrator() (common.Address, error) {
	return _PancakeFarm.Contract.Migrator(&_PancakeFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeFarm *PancakeFarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeFarm *PancakeFarmSession) Owner() (common.Address, error) {
	return _PancakeFarm.Contract.Owner(&_PancakeFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeFarm *PancakeFarmCallerSession) Owner() (common.Address, error) {
	return _PancakeFarm.Contract.Owner(&_PancakeFarm.CallOpts)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) PendingCake(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "pendingCake", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _PancakeFarm.Contract.PendingCake(&_PancakeFarm.CallOpts, _pid, _user)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _PancakeFarm.Contract.PendingCake(&_PancakeFarm.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_PancakeFarm *PancakeFarmCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccCakePerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccCakePerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_PancakeFarm *PancakeFarmSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _PancakeFarm.Contract.PoolInfo(&_PancakeFarm.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_PancakeFarm *PancakeFarmCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _PancakeFarm.Contract.PoolInfo(&_PancakeFarm.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) PoolLength() (*big.Int, error) {
	return _PancakeFarm.Contract.PoolLength(&_PancakeFarm.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) PoolLength() (*big.Int, error) {
	return _PancakeFarm.Contract.PoolLength(&_PancakeFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) StartBlock() (*big.Int, error) {
	return _PancakeFarm.Contract.StartBlock(&_PancakeFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) StartBlock() (*big.Int, error) {
	return _PancakeFarm.Contract.StartBlock(&_PancakeFarm.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_PancakeFarm *PancakeFarmCaller) Syrup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "syrup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_PancakeFarm *PancakeFarmSession) Syrup() (common.Address, error) {
	return _PancakeFarm.Contract.Syrup(&_PancakeFarm.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() view returns(address)
func (_PancakeFarm *PancakeFarmCallerSession) Syrup() (common.Address, error) {
	return _PancakeFarm.Contract.Syrup(&_PancakeFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_PancakeFarm *PancakeFarmCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_PancakeFarm *PancakeFarmSession) TotalAllocPoint() (*big.Int, error) {
	return _PancakeFarm.Contract.TotalAllocPoint(&_PancakeFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_PancakeFarm *PancakeFarmCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _PancakeFarm.Contract.TotalAllocPoint(&_PancakeFarm.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_PancakeFarm *PancakeFarmCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _PancakeFarm.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_PancakeFarm *PancakeFarmSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _PancakeFarm.Contract.UserInfo(&_PancakeFarm.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_PancakeFarm *PancakeFarmCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _PancakeFarm.Contract.UserInfo(&_PancakeFarm.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_PancakeFarm *PancakeFarmTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_PancakeFarm *PancakeFarmSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Add(&_PancakeFarm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Add(&_PancakeFarm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Deposit(&_PancakeFarm.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Deposit(&_PancakeFarm.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_PancakeFarm *PancakeFarmTransactor) Dev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "dev", _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_PancakeFarm *PancakeFarmSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Dev(&_PancakeFarm.TransactOpts, _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Dev(&_PancakeFarm.TransactOpts, _devaddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.EmergencyWithdraw(&_PancakeFarm.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.EmergencyWithdraw(&_PancakeFarm.TransactOpts, _pid)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactor) EnterStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "enterStaking", _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.EnterStaking(&_PancakeFarm.TransactOpts, _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.EnterStaking(&_PancakeFarm.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactor) LeaveStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "leaveStaking", _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.LeaveStaking(&_PancakeFarm.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.LeaveStaking(&_PancakeFarm.TransactOpts, _amount)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PancakeFarm *PancakeFarmTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PancakeFarm *PancakeFarmSession) MassUpdatePools() (*types.Transaction, error) {
	return _PancakeFarm.Contract.MassUpdatePools(&_PancakeFarm.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PancakeFarm *PancakeFarmTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _PancakeFarm.Contract.MassUpdatePools(&_PancakeFarm.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Migrate(&_PancakeFarm.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Migrate(&_PancakeFarm.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeFarm *PancakeFarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeFarm *PancakeFarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeFarm.Contract.RenounceOwnership(&_PancakeFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeFarm *PancakeFarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeFarm.Contract.RenounceOwnership(&_PancakeFarm.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_PancakeFarm *PancakeFarmTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_PancakeFarm *PancakeFarmSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Set(&_PancakeFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Set(&_PancakeFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_PancakeFarm *PancakeFarmTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_PancakeFarm *PancakeFarmSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _PancakeFarm.Contract.SetMigrator(&_PancakeFarm.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _PancakeFarm.Contract.SetMigrator(&_PancakeFarm.TransactOpts, _migrator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeFarm *PancakeFarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeFarm *PancakeFarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeFarm.Contract.TransferOwnership(&_PancakeFarm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeFarm.Contract.TransferOwnership(&_PancakeFarm.TransactOpts, newOwner)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_PancakeFarm *PancakeFarmTransactor) UpdateMultiplier(opts *bind.TransactOpts, multiplierNumber *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "updateMultiplier", multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_PancakeFarm *PancakeFarmSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.UpdateMultiplier(&_PancakeFarm.TransactOpts, multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.UpdateMultiplier(&_PancakeFarm.TransactOpts, multiplierNumber)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.UpdatePool(&_PancakeFarm.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.UpdatePool(&_PancakeFarm.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Withdraw(&_PancakeFarm.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_PancakeFarm *PancakeFarmTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeFarm.Contract.Withdraw(&_PancakeFarm.TransactOpts, _pid, _amount)
}

// PancakeFarmDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PancakeFarm contract.
type PancakeFarmDepositIterator struct {
	Event *PancakeFarmDeposit // Event containing the contract specifics and raw log

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
func (it *PancakeFarmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeFarmDeposit)
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
		it.Event = new(PancakeFarmDeposit)
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
func (it *PancakeFarmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeFarmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeFarmDeposit represents a Deposit event raised by the PancakeFarm contract.
type PancakeFarmDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeFarm *PancakeFarmFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*PancakeFarmDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeFarm.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmDepositIterator{contract: _PancakeFarm.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeFarm *PancakeFarmFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PancakeFarmDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeFarm.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeFarmDeposit)
				if err := _PancakeFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_PancakeFarm *PancakeFarmFilterer) ParseDeposit(log types.Log) (*PancakeFarmDeposit, error) {
	event := new(PancakeFarmDeposit)
	if err := _PancakeFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeFarmEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the PancakeFarm contract.
type PancakeFarmEmergencyWithdrawIterator struct {
	Event *PancakeFarmEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeFarmEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeFarmEmergencyWithdraw)
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
		it.Event = new(PancakeFarmEmergencyWithdraw)
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
func (it *PancakeFarmEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeFarmEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeFarmEmergencyWithdraw represents a EmergencyWithdraw event raised by the PancakeFarm contract.
type PancakeFarmEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeFarm *PancakeFarmFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*PancakeFarmEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeFarm.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmEmergencyWithdrawIterator{contract: _PancakeFarm.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeFarm *PancakeFarmFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeFarmEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeFarm.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeFarmEmergencyWithdraw)
				if err := _PancakeFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_PancakeFarm *PancakeFarmFilterer) ParseEmergencyWithdraw(log types.Log) (*PancakeFarmEmergencyWithdraw, error) {
	event := new(PancakeFarmEmergencyWithdraw)
	if err := _PancakeFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeFarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PancakeFarm contract.
type PancakeFarmOwnershipTransferredIterator struct {
	Event *PancakeFarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeFarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeFarmOwnershipTransferred)
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
		it.Event = new(PancakeFarmOwnershipTransferred)
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
func (it *PancakeFarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeFarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeFarmOwnershipTransferred represents a OwnershipTransferred event raised by the PancakeFarm contract.
type PancakeFarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeFarm *PancakeFarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakeFarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeFarm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmOwnershipTransferredIterator{contract: _PancakeFarm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeFarm *PancakeFarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeFarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeFarm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeFarmOwnershipTransferred)
				if err := _PancakeFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PancakeFarm *PancakeFarmFilterer) ParseOwnershipTransferred(log types.Log) (*PancakeFarmOwnershipTransferred, error) {
	event := new(PancakeFarmOwnershipTransferred)
	if err := _PancakeFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeFarmWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PancakeFarm contract.
type PancakeFarmWithdrawIterator struct {
	Event *PancakeFarmWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeFarmWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeFarmWithdraw)
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
		it.Event = new(PancakeFarmWithdraw)
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
func (it *PancakeFarmWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeFarmWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeFarmWithdraw represents a Withdraw event raised by the PancakeFarm contract.
type PancakeFarmWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeFarm *PancakeFarmFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*PancakeFarmWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeFarm.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeFarmWithdrawIterator{contract: _PancakeFarm.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeFarm *PancakeFarmFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeFarmWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeFarm.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeFarmWithdraw)
				if err := _PancakeFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_PancakeFarm *PancakeFarmFilterer) ParseWithdraw(log types.Log) (*PancakeFarmWithdraw, error) {
	event := new(PancakeFarmWithdraw)
	if err := _PancakeFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
