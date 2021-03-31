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

// BoardroomFarmABI is the input ABI used to generate the binding from.
const BoardroomFarmABI = "[{\"inputs\":[{\"internalType\":\"contractIWHT\",\"name\":\"_wht\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cycle\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_whtAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"name\":\"newAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accWhtPerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCycle\",\"type\":\"uint256\"}],\"name\":\"setCycle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wht\",\"outputs\":[{\"internalType\":\"contractIWHT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whtPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BoardroomFarm is an auto generated Go binding around an Ethereum contract.
type BoardroomFarm struct {
	BoardroomFarmCaller     // Read-only binding to the contract
	BoardroomFarmTransactor // Write-only binding to the contract
	BoardroomFarmFilterer   // Log filterer for contract events
}

// BoardroomFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoardroomFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardroomFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoardroomFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardroomFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoardroomFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardroomFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoardroomFarmSession struct {
	Contract     *BoardroomFarm    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoardroomFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoardroomFarmCallerSession struct {
	Contract *BoardroomFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BoardroomFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoardroomFarmTransactorSession struct {
	Contract     *BoardroomFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BoardroomFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoardroomFarmRaw struct {
	Contract *BoardroomFarm // Generic contract binding to access the raw methods on
}

// BoardroomFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoardroomFarmCallerRaw struct {
	Contract *BoardroomFarmCaller // Generic read-only contract binding to access the raw methods on
}

// BoardroomFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoardroomFarmTransactorRaw struct {
	Contract *BoardroomFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoardroomFarm creates a new instance of BoardroomFarm, bound to a specific deployed contract.
func NewBoardroomFarm(address common.Address, backend bind.ContractBackend) (*BoardroomFarm, error) {
	contract, err := bindBoardroomFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarm{BoardroomFarmCaller: BoardroomFarmCaller{contract: contract}, BoardroomFarmTransactor: BoardroomFarmTransactor{contract: contract}, BoardroomFarmFilterer: BoardroomFarmFilterer{contract: contract}}, nil
}

// NewBoardroomFarmCaller creates a new read-only instance of BoardroomFarm, bound to a specific deployed contract.
func NewBoardroomFarmCaller(address common.Address, caller bind.ContractCaller) (*BoardroomFarmCaller, error) {
	contract, err := bindBoardroomFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmCaller{contract: contract}, nil
}

// NewBoardroomFarmTransactor creates a new write-only instance of BoardroomFarm, bound to a specific deployed contract.
func NewBoardroomFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*BoardroomFarmTransactor, error) {
	contract, err := bindBoardroomFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmTransactor{contract: contract}, nil
}

// NewBoardroomFarmFilterer creates a new log filterer instance of BoardroomFarm, bound to a specific deployed contract.
func NewBoardroomFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*BoardroomFarmFilterer, error) {
	contract, err := bindBoardroomFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmFilterer{contract: contract}, nil
}

// bindBoardroomFarm binds a generic wrapper to an already deployed contract.
func bindBoardroomFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoardroomFarmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoardroomFarm *BoardroomFarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoardroomFarm.Contract.BoardroomFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoardroomFarm *BoardroomFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.BoardroomFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoardroomFarm *BoardroomFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.BoardroomFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoardroomFarm *BoardroomFarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoardroomFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoardroomFarm *BoardroomFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoardroomFarm *BoardroomFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.contract.Transact(opts, method, params...)
}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) Cycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "cycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) Cycle() (*big.Int, error) {
	return _BoardroomFarm.Contract.Cycle(&_BoardroomFarm.CallOpts)
}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) Cycle() (*big.Int, error) {
	return _BoardroomFarm.Contract.Cycle(&_BoardroomFarm.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) EndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "endBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) EndBlock() (*big.Int, error) {
	return _BoardroomFarm.Contract.EndBlock(&_BoardroomFarm.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) EndBlock() (*big.Int, error) {
	return _BoardroomFarm.Contract.EndBlock(&_BoardroomFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoardroomFarm *BoardroomFarmCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoardroomFarm *BoardroomFarmSession) Owner() (common.Address, error) {
	return _BoardroomFarm.Contract.Owner(&_BoardroomFarm.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoardroomFarm *BoardroomFarmCallerSession) Owner() (common.Address, error) {
	return _BoardroomFarm.Contract.Owner(&_BoardroomFarm.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) Pending(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "pending", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _BoardroomFarm.Contract.Pending(&_BoardroomFarm.CallOpts, _pid, _user)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _BoardroomFarm.Contract.Pending(&_BoardroomFarm.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accWhtPerShare)
func (_BoardroomFarm *BoardroomFarmCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccWhtPerShare  *big.Int
}, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccWhtPerShare  *big.Int
	})

	outstruct.LpToken = out[0].(common.Address)
	outstruct.AllocPoint = out[1].(*big.Int)
	outstruct.LastRewardBlock = out[2].(*big.Int)
	outstruct.AccWhtPerShare = out[3].(*big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accWhtPerShare)
func (_BoardroomFarm *BoardroomFarmSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccWhtPerShare  *big.Int
}, error) {
	return _BoardroomFarm.Contract.PoolInfo(&_BoardroomFarm.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accWhtPerShare)
func (_BoardroomFarm *BoardroomFarmCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccWhtPerShare  *big.Int
}, error) {
	return _BoardroomFarm.Contract.PoolInfo(&_BoardroomFarm.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) PoolLength() (*big.Int, error) {
	return _BoardroomFarm.Contract.PoolLength(&_BoardroomFarm.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) PoolLength() (*big.Int, error) {
	return _BoardroomFarm.Contract.PoolLength(&_BoardroomFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) StartBlock() (*big.Int, error) {
	return _BoardroomFarm.Contract.StartBlock(&_BoardroomFarm.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) StartBlock() (*big.Int, error) {
	return _BoardroomFarm.Contract.StartBlock(&_BoardroomFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) TotalAllocPoint() (*big.Int, error) {
	return _BoardroomFarm.Contract.TotalAllocPoint(&_BoardroomFarm.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _BoardroomFarm.Contract.TotalAllocPoint(&_BoardroomFarm.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BoardroomFarm *BoardroomFarmCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.RewardDebt = out[1].(*big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BoardroomFarm *BoardroomFarmSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _BoardroomFarm.Contract.UserInfo(&_BoardroomFarm.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BoardroomFarm *BoardroomFarmCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _BoardroomFarm.Contract.UserInfo(&_BoardroomFarm.CallOpts, arg0, arg1)
}

// Wht is a free data retrieval call binding the contract method 0x785f3499.
//
// Solidity: function wht() view returns(address)
func (_BoardroomFarm *BoardroomFarmCaller) Wht(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "wht")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wht is a free data retrieval call binding the contract method 0x785f3499.
//
// Solidity: function wht() view returns(address)
func (_BoardroomFarm *BoardroomFarmSession) Wht() (common.Address, error) {
	return _BoardroomFarm.Contract.Wht(&_BoardroomFarm.CallOpts)
}

// Wht is a free data retrieval call binding the contract method 0x785f3499.
//
// Solidity: function wht() view returns(address)
func (_BoardroomFarm *BoardroomFarmCallerSession) Wht() (common.Address, error) {
	return _BoardroomFarm.Contract.Wht(&_BoardroomFarm.CallOpts)
}

// WhtPerBlock is a free data retrieval call binding the contract method 0x0cef9396.
//
// Solidity: function whtPerBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCaller) WhtPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarm.contract.Call(opts, &out, "whtPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhtPerBlock is a free data retrieval call binding the contract method 0x0cef9396.
//
// Solidity: function whtPerBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmSession) WhtPerBlock() (*big.Int, error) {
	return _BoardroomFarm.Contract.WhtPerBlock(&_BoardroomFarm.CallOpts)
}

// WhtPerBlock is a free data retrieval call binding the contract method 0x0cef9396.
//
// Solidity: function whtPerBlock() view returns(uint256)
func (_BoardroomFarm *BoardroomFarmCallerSession) WhtPerBlock() (*big.Int, error) {
	return _BoardroomFarm.Contract.WhtPerBlock(&_BoardroomFarm.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_BoardroomFarm *BoardroomFarmSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Add(&_BoardroomFarm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Add(&_BoardroomFarm.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarm *BoardroomFarmSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Deposit(&_BoardroomFarm.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Deposit(&_BoardroomFarm.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BoardroomFarm *BoardroomFarmSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.EmergencyWithdraw(&_BoardroomFarm.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.EmergencyWithdraw(&_BoardroomFarm.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BoardroomFarm *BoardroomFarmTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BoardroomFarm *BoardroomFarmSession) MassUpdatePools() (*types.Transaction, error) {
	return _BoardroomFarm.Contract.MassUpdatePools(&_BoardroomFarm.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _BoardroomFarm.Contract.MassUpdatePools(&_BoardroomFarm.TransactOpts)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0x4ebe495c.
//
// Solidity: function newAirdrop(uint256 _whtAmount, uint256 _newPerBlock, uint256 _startBlock) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) NewAirdrop(opts *bind.TransactOpts, _whtAmount *big.Int, _newPerBlock *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "newAirdrop", _whtAmount, _newPerBlock, _startBlock)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0x4ebe495c.
//
// Solidity: function newAirdrop(uint256 _whtAmount, uint256 _newPerBlock, uint256 _startBlock) returns()
func (_BoardroomFarm *BoardroomFarmSession) NewAirdrop(_whtAmount *big.Int, _newPerBlock *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.NewAirdrop(&_BoardroomFarm.TransactOpts, _whtAmount, _newPerBlock, _startBlock)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0x4ebe495c.
//
// Solidity: function newAirdrop(uint256 _whtAmount, uint256 _newPerBlock, uint256 _startBlock) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) NewAirdrop(_whtAmount *big.Int, _newPerBlock *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.NewAirdrop(&_BoardroomFarm.TransactOpts, _whtAmount, _newPerBlock, _startBlock)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoardroomFarm *BoardroomFarmTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoardroomFarm *BoardroomFarmSession) RenounceOwnership() (*types.Transaction, error) {
	return _BoardroomFarm.Contract.RenounceOwnership(&_BoardroomFarm.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BoardroomFarm.Contract.RenounceOwnership(&_BoardroomFarm.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_BoardroomFarm *BoardroomFarmSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Set(&_BoardroomFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Set(&_BoardroomFarm.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetCycle is a paid mutator transaction binding the contract method 0x6d93c3ba.
//
// Solidity: function setCycle(uint256 _newCycle) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) SetCycle(opts *bind.TransactOpts, _newCycle *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "setCycle", _newCycle)
}

// SetCycle is a paid mutator transaction binding the contract method 0x6d93c3ba.
//
// Solidity: function setCycle(uint256 _newCycle) returns()
func (_BoardroomFarm *BoardroomFarmSession) SetCycle(_newCycle *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.SetCycle(&_BoardroomFarm.TransactOpts, _newCycle)
}

// SetCycle is a paid mutator transaction binding the contract method 0x6d93c3ba.
//
// Solidity: function setCycle(uint256 _newCycle) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) SetCycle(_newCycle *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.SetCycle(&_BoardroomFarm.TransactOpts, _newCycle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoardroomFarm *BoardroomFarmSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.TransferOwnership(&_BoardroomFarm.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.TransferOwnership(&_BoardroomFarm.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BoardroomFarm *BoardroomFarmSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.UpdatePool(&_BoardroomFarm.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.UpdatePool(&_BoardroomFarm.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarm *BoardroomFarmTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarm *BoardroomFarmSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Withdraw(&_BoardroomFarm.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarm *BoardroomFarmTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarm.Contract.Withdraw(&_BoardroomFarm.TransactOpts, _pid, _amount)
}

// BoardroomFarmDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BoardroomFarm contract.
type BoardroomFarmDepositIterator struct {
	Event *BoardroomFarmDeposit // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmDeposit)
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
		it.Event = new(BoardroomFarmDeposit)
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
func (it *BoardroomFarmDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmDeposit represents a Deposit event raised by the BoardroomFarm contract.
type BoardroomFarmDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarm *BoardroomFarmFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BoardroomFarmDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarm.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmDepositIterator{contract: _BoardroomFarm.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarm *BoardroomFarmFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BoardroomFarmDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarm.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmDeposit)
				if err := _BoardroomFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BoardroomFarm *BoardroomFarmFilterer) ParseDeposit(log types.Log) (*BoardroomFarmDeposit, error) {
	event := new(BoardroomFarmDeposit)
	if err := _BoardroomFarm.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoardroomFarmEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the BoardroomFarm contract.
type BoardroomFarmEmergencyWithdrawIterator struct {
	Event *BoardroomFarmEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmEmergencyWithdraw)
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
		it.Event = new(BoardroomFarmEmergencyWithdraw)
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
func (it *BoardroomFarmEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmEmergencyWithdraw represents a EmergencyWithdraw event raised by the BoardroomFarm contract.
type BoardroomFarmEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarm *BoardroomFarmFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BoardroomFarmEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarm.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmEmergencyWithdrawIterator{contract: _BoardroomFarm.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarm *BoardroomFarmFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *BoardroomFarmEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarm.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmEmergencyWithdraw)
				if err := _BoardroomFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_BoardroomFarm *BoardroomFarmFilterer) ParseEmergencyWithdraw(log types.Log) (*BoardroomFarmEmergencyWithdraw, error) {
	event := new(BoardroomFarmEmergencyWithdraw)
	if err := _BoardroomFarm.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoardroomFarmOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BoardroomFarm contract.
type BoardroomFarmOwnershipTransferredIterator struct {
	Event *BoardroomFarmOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmOwnershipTransferred)
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
		it.Event = new(BoardroomFarmOwnershipTransferred)
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
func (it *BoardroomFarmOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmOwnershipTransferred represents a OwnershipTransferred event raised by the BoardroomFarm contract.
type BoardroomFarmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BoardroomFarm *BoardroomFarmFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BoardroomFarmOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BoardroomFarm.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmOwnershipTransferredIterator{contract: _BoardroomFarm.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BoardroomFarm *BoardroomFarmFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BoardroomFarmOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BoardroomFarm.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmOwnershipTransferred)
				if err := _BoardroomFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BoardroomFarm *BoardroomFarmFilterer) ParseOwnershipTransferred(log types.Log) (*BoardroomFarmOwnershipTransferred, error) {
	event := new(BoardroomFarmOwnershipTransferred)
	if err := _BoardroomFarm.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoardroomFarmWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BoardroomFarm contract.
type BoardroomFarmWithdrawIterator struct {
	Event *BoardroomFarmWithdraw // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmWithdraw)
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
		it.Event = new(BoardroomFarmWithdraw)
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
func (it *BoardroomFarmWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmWithdraw represents a Withdraw event raised by the BoardroomFarm contract.
type BoardroomFarmWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarm *BoardroomFarmFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BoardroomFarmWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarm.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmWithdrawIterator{contract: _BoardroomFarm.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarm *BoardroomFarmFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BoardroomFarmWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarm.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmWithdraw)
				if err := _BoardroomFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BoardroomFarm *BoardroomFarmFilterer) ParseWithdraw(log types.Log) (*BoardroomFarmWithdraw, error) {
	event := new(BoardroomFarmWithdraw)
	if err := _BoardroomFarm.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
