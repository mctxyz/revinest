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

// BoardroomFarmMdxABI is the input ABI used to generate the binding from.
const BoardroomFarmMdxABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mdx\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cycle\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdxPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mdxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"name\":\"newAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accMDXPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mdxAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCycle\",\"type\":\"uint256\"}],\"name\":\"setCycle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BoardroomFarmMdx is an auto generated Go binding around an Ethereum contract.
type BoardroomFarmMdx struct {
	BoardroomFarmMdxCaller     // Read-only binding to the contract
	BoardroomFarmMdxTransactor // Write-only binding to the contract
	BoardroomFarmMdxFilterer   // Log filterer for contract events
}

// BoardroomFarmMdxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoardroomFarmMdxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardroomFarmMdxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoardroomFarmMdxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardroomFarmMdxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoardroomFarmMdxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoardroomFarmMdxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoardroomFarmMdxSession struct {
	Contract     *BoardroomFarmMdx // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoardroomFarmMdxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoardroomFarmMdxCallerSession struct {
	Contract *BoardroomFarmMdxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BoardroomFarmMdxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoardroomFarmMdxTransactorSession struct {
	Contract     *BoardroomFarmMdxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BoardroomFarmMdxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoardroomFarmMdxRaw struct {
	Contract *BoardroomFarmMdx // Generic contract binding to access the raw methods on
}

// BoardroomFarmMdxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoardroomFarmMdxCallerRaw struct {
	Contract *BoardroomFarmMdxCaller // Generic read-only contract binding to access the raw methods on
}

// BoardroomFarmMdxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoardroomFarmMdxTransactorRaw struct {
	Contract *BoardroomFarmMdxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoardroomFarmMdx creates a new instance of BoardroomFarmMdx, bound to a specific deployed contract.
func NewBoardroomFarmMdx(address common.Address, backend bind.ContractBackend) (*BoardroomFarmMdx, error) {
	contract, err := bindBoardroomFarmMdx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdx{BoardroomFarmMdxCaller: BoardroomFarmMdxCaller{contract: contract}, BoardroomFarmMdxTransactor: BoardroomFarmMdxTransactor{contract: contract}, BoardroomFarmMdxFilterer: BoardroomFarmMdxFilterer{contract: contract}}, nil
}

// NewBoardroomFarmMdxCaller creates a new read-only instance of BoardroomFarmMdx, bound to a specific deployed contract.
func NewBoardroomFarmMdxCaller(address common.Address, caller bind.ContractCaller) (*BoardroomFarmMdxCaller, error) {
	contract, err := bindBoardroomFarmMdx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxCaller{contract: contract}, nil
}

// NewBoardroomFarmMdxTransactor creates a new write-only instance of BoardroomFarmMdx, bound to a specific deployed contract.
func NewBoardroomFarmMdxTransactor(address common.Address, transactor bind.ContractTransactor) (*BoardroomFarmMdxTransactor, error) {
	contract, err := bindBoardroomFarmMdx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxTransactor{contract: contract}, nil
}

// NewBoardroomFarmMdxFilterer creates a new log filterer instance of BoardroomFarmMdx, bound to a specific deployed contract.
func NewBoardroomFarmMdxFilterer(address common.Address, filterer bind.ContractFilterer) (*BoardroomFarmMdxFilterer, error) {
	contract, err := bindBoardroomFarmMdx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxFilterer{contract: contract}, nil
}

// bindBoardroomFarmMdx binds a generic wrapper to an already deployed contract.
func bindBoardroomFarmMdx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoardroomFarmMdxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoardroomFarmMdx *BoardroomFarmMdxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoardroomFarmMdx.Contract.BoardroomFarmMdxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoardroomFarmMdx *BoardroomFarmMdxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.BoardroomFarmMdxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoardroomFarmMdx *BoardroomFarmMdxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.BoardroomFarmMdxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BoardroomFarmMdx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.contract.Transact(opts, method, params...)
}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) Cycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "cycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Cycle() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.Cycle(&_BoardroomFarmMdx.CallOpts)
}

// Cycle is a free data retrieval call binding the contract method 0x6190c9d5.
//
// Solidity: function cycle() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) Cycle() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.Cycle(&_BoardroomFarmMdx.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) EndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "endBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) EndBlock() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.EndBlock(&_BoardroomFarmMdx.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) EndBlock() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.EndBlock(&_BoardroomFarmMdx.CallOpts)
}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) Mdx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "mdx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Mdx() (common.Address, error) {
	return _BoardroomFarmMdx.Contract.Mdx(&_BoardroomFarmMdx.CallOpts)
}

// Mdx is a free data retrieval call binding the contract method 0x5672ab55.
//
// Solidity: function mdx() view returns(address)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) Mdx() (common.Address, error) {
	return _BoardroomFarmMdx.Contract.Mdx(&_BoardroomFarmMdx.CallOpts)
}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) MdxPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "mdxPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) MdxPerBlock() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.MdxPerBlock(&_BoardroomFarmMdx.CallOpts)
}

// MdxPerBlock is a free data retrieval call binding the contract method 0x8fa39a88.
//
// Solidity: function mdxPerBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) MdxPerBlock() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.MdxPerBlock(&_BoardroomFarmMdx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Owner() (common.Address, error) {
	return _BoardroomFarmMdx.Contract.Owner(&_BoardroomFarmMdx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) Owner() (common.Address, error) {
	return _BoardroomFarmMdx.Contract.Owner(&_BoardroomFarmMdx.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) Pending(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "pending", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.Pending(&_BoardroomFarmMdx.CallOpts, _pid, _user)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.Pending(&_BoardroomFarmMdx.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accMDXPerShare, uint256 mdxAmount)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccMDXPerShare  *big.Int
	MdxAmount       *big.Int
}, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccMDXPerShare  *big.Int
		MdxAmount       *big.Int
	})

	outstruct.LpToken = out[0].(common.Address)
	outstruct.AllocPoint = out[1].(*big.Int)
	outstruct.LastRewardBlock = out[2].(*big.Int)
	outstruct.AccMDXPerShare = out[3].(*big.Int)
	outstruct.MdxAmount = out[4].(*big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accMDXPerShare, uint256 mdxAmount)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccMDXPerShare  *big.Int
	MdxAmount       *big.Int
}, error) {
	return _BoardroomFarmMdx.Contract.PoolInfo(&_BoardroomFarmMdx.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accMDXPerShare, uint256 mdxAmount)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccMDXPerShare  *big.Int
	MdxAmount       *big.Int
}, error) {
	return _BoardroomFarmMdx.Contract.PoolInfo(&_BoardroomFarmMdx.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) PoolLength() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.PoolLength(&_BoardroomFarmMdx.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) PoolLength() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.PoolLength(&_BoardroomFarmMdx.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) StartBlock() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.StartBlock(&_BoardroomFarmMdx.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) StartBlock() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.StartBlock(&_BoardroomFarmMdx.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) TotalAllocPoint() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.TotalAllocPoint(&_BoardroomFarmMdx.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _BoardroomFarmMdx.Contract.TotalAllocPoint(&_BoardroomFarmMdx.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BoardroomFarmMdx *BoardroomFarmMdxCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _BoardroomFarmMdx.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _BoardroomFarmMdx.Contract.UserInfo(&_BoardroomFarmMdx.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_BoardroomFarmMdx *BoardroomFarmMdxCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _BoardroomFarmMdx.Contract.UserInfo(&_BoardroomFarmMdx.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Add(&_BoardroomFarmMdx.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Add(&_BoardroomFarmMdx.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Deposit(&_BoardroomFarmMdx.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Deposit(&_BoardroomFarmMdx.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.EmergencyWithdraw(&_BoardroomFarmMdx.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.EmergencyWithdraw(&_BoardroomFarmMdx.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) MassUpdatePools() (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.MassUpdatePools(&_BoardroomFarmMdx.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.MassUpdatePools(&_BoardroomFarmMdx.TransactOpts)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0x4ebe495c.
//
// Solidity: function newAirdrop(uint256 _mdxAmount, uint256 _newPerBlock, uint256 _startBlock) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) NewAirdrop(opts *bind.TransactOpts, _mdxAmount *big.Int, _newPerBlock *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "newAirdrop", _mdxAmount, _newPerBlock, _startBlock)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0x4ebe495c.
//
// Solidity: function newAirdrop(uint256 _mdxAmount, uint256 _newPerBlock, uint256 _startBlock) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) NewAirdrop(_mdxAmount *big.Int, _newPerBlock *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.NewAirdrop(&_BoardroomFarmMdx.TransactOpts, _mdxAmount, _newPerBlock, _startBlock)
}

// NewAirdrop is a paid mutator transaction binding the contract method 0x4ebe495c.
//
// Solidity: function newAirdrop(uint256 _mdxAmount, uint256 _newPerBlock, uint256 _startBlock) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) NewAirdrop(_mdxAmount *big.Int, _newPerBlock *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.NewAirdrop(&_BoardroomFarmMdx.TransactOpts, _mdxAmount, _newPerBlock, _startBlock)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) RenounceOwnership() (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.RenounceOwnership(&_BoardroomFarmMdx.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.RenounceOwnership(&_BoardroomFarmMdx.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Set(&_BoardroomFarmMdx.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Set(&_BoardroomFarmMdx.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetCycle is a paid mutator transaction binding the contract method 0x6d93c3ba.
//
// Solidity: function setCycle(uint256 _newCycle) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) SetCycle(opts *bind.TransactOpts, _newCycle *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "setCycle", _newCycle)
}

// SetCycle is a paid mutator transaction binding the contract method 0x6d93c3ba.
//
// Solidity: function setCycle(uint256 _newCycle) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) SetCycle(_newCycle *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.SetCycle(&_BoardroomFarmMdx.TransactOpts, _newCycle)
}

// SetCycle is a paid mutator transaction binding the contract method 0x6d93c3ba.
//
// Solidity: function setCycle(uint256 _newCycle) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) SetCycle(_newCycle *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.SetCycle(&_BoardroomFarmMdx.TransactOpts, _newCycle)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.TransferOwnership(&_BoardroomFarmMdx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.TransferOwnership(&_BoardroomFarmMdx.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.UpdatePool(&_BoardroomFarmMdx.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.UpdatePool(&_BoardroomFarmMdx.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Withdraw(&_BoardroomFarmMdx.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_BoardroomFarmMdx *BoardroomFarmMdxTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _BoardroomFarmMdx.Contract.Withdraw(&_BoardroomFarmMdx.TransactOpts, _pid, _amount)
}

// BoardroomFarmMdxDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxDepositIterator struct {
	Event *BoardroomFarmMdxDeposit // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmMdxDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmMdxDeposit)
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
		it.Event = new(BoardroomFarmMdxDeposit)
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
func (it *BoardroomFarmMdxDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmMdxDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmMdxDeposit represents a Deposit event raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BoardroomFarmMdxDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxDepositIterator{contract: _BoardroomFarmMdx.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BoardroomFarmMdxDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmMdxDeposit)
				if err := _BoardroomFarmMdx.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) ParseDeposit(log types.Log) (*BoardroomFarmMdxDeposit, error) {
	event := new(BoardroomFarmMdxDeposit)
	if err := _BoardroomFarmMdx.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoardroomFarmMdxEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxEmergencyWithdrawIterator struct {
	Event *BoardroomFarmMdxEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmMdxEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmMdxEmergencyWithdraw)
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
		it.Event = new(BoardroomFarmMdxEmergencyWithdraw)
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
func (it *BoardroomFarmMdxEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmMdxEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmMdxEmergencyWithdraw represents a EmergencyWithdraw event raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BoardroomFarmMdxEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxEmergencyWithdrawIterator{contract: _BoardroomFarmMdx.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *BoardroomFarmMdxEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmMdxEmergencyWithdraw)
				if err := _BoardroomFarmMdx.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) ParseEmergencyWithdraw(log types.Log) (*BoardroomFarmMdxEmergencyWithdraw, error) {
	event := new(BoardroomFarmMdxEmergencyWithdraw)
	if err := _BoardroomFarmMdx.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoardroomFarmMdxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxOwnershipTransferredIterator struct {
	Event *BoardroomFarmMdxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmMdxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmMdxOwnershipTransferred)
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
		it.Event = new(BoardroomFarmMdxOwnershipTransferred)
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
func (it *BoardroomFarmMdxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmMdxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmMdxOwnershipTransferred represents a OwnershipTransferred event raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BoardroomFarmMdxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxOwnershipTransferredIterator{contract: _BoardroomFarmMdx.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BoardroomFarmMdxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmMdxOwnershipTransferred)
				if err := _BoardroomFarmMdx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) ParseOwnershipTransferred(log types.Log) (*BoardroomFarmMdxOwnershipTransferred, error) {
	event := new(BoardroomFarmMdxOwnershipTransferred)
	if err := _BoardroomFarmMdx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoardroomFarmMdxWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxWithdrawIterator struct {
	Event *BoardroomFarmMdxWithdraw // Event containing the contract specifics and raw log

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
func (it *BoardroomFarmMdxWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoardroomFarmMdxWithdraw)
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
		it.Event = new(BoardroomFarmMdxWithdraw)
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
func (it *BoardroomFarmMdxWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoardroomFarmMdxWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoardroomFarmMdxWithdraw represents a Withdraw event raised by the BoardroomFarmMdx contract.
type BoardroomFarmMdxWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BoardroomFarmMdxWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BoardroomFarmMdxWithdrawIterator{contract: _BoardroomFarmMdx.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BoardroomFarmMdxWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _BoardroomFarmMdx.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoardroomFarmMdxWithdraw)
				if err := _BoardroomFarmMdx.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BoardroomFarmMdx *BoardroomFarmMdxFilterer) ParseWithdraw(log types.Log) (*BoardroomFarmMdxWithdraw, error) {
	event := new(BoardroomFarmMdxWithdraw)
	if err := _BoardroomFarmMdx.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
