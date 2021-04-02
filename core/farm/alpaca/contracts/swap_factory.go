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

// AplpacaFactoryABI is the input ABI used to generate the binding from.
const AplpacaFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_CODE_PAIR_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AplpacaFactory is an auto generated Go binding around an Ethereum contract.
type AplpacaFactory struct {
	AplpacaFactoryCaller     // Read-only binding to the contract
	AplpacaFactoryTransactor // Write-only binding to the contract
	AplpacaFactoryFilterer   // Log filterer for contract events
}

// AplpacaFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AplpacaFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AplpacaFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AplpacaFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AplpacaFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AplpacaFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AplpacaFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AplpacaFactorySession struct {
	Contract     *AplpacaFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AplpacaFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AplpacaFactoryCallerSession struct {
	Contract *AplpacaFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AplpacaFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AplpacaFactoryTransactorSession struct {
	Contract     *AplpacaFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AplpacaFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AplpacaFactoryRaw struct {
	Contract *AplpacaFactory // Generic contract binding to access the raw methods on
}

// AplpacaFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AplpacaFactoryCallerRaw struct {
	Contract *AplpacaFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AplpacaFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AplpacaFactoryTransactorRaw struct {
	Contract *AplpacaFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAplpacaFactory creates a new instance of AplpacaFactory, bound to a specific deployed contract.
func NewAplpacaFactory(address common.Address, backend bind.ContractBackend) (*AplpacaFactory, error) {
	contract, err := bindAplpacaFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AplpacaFactory{AplpacaFactoryCaller: AplpacaFactoryCaller{contract: contract}, AplpacaFactoryTransactor: AplpacaFactoryTransactor{contract: contract}, AplpacaFactoryFilterer: AplpacaFactoryFilterer{contract: contract}}, nil
}

// NewAplpacaFactoryCaller creates a new read-only instance of AplpacaFactory, bound to a specific deployed contract.
func NewAplpacaFactoryCaller(address common.Address, caller bind.ContractCaller) (*AplpacaFactoryCaller, error) {
	contract, err := bindAplpacaFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AplpacaFactoryCaller{contract: contract}, nil
}

// NewAplpacaFactoryTransactor creates a new write-only instance of AplpacaFactory, bound to a specific deployed contract.
func NewAplpacaFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AplpacaFactoryTransactor, error) {
	contract, err := bindAplpacaFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AplpacaFactoryTransactor{contract: contract}, nil
}

// NewAplpacaFactoryFilterer creates a new log filterer instance of AplpacaFactory, bound to a specific deployed contract.
func NewAplpacaFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AplpacaFactoryFilterer, error) {
	contract, err := bindAplpacaFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AplpacaFactoryFilterer{contract: contract}, nil
}

// bindAplpacaFactory binds a generic wrapper to an already deployed contract.
func bindAplpacaFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AplpacaFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AplpacaFactory *AplpacaFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AplpacaFactory.Contract.AplpacaFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AplpacaFactory *AplpacaFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.AplpacaFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AplpacaFactory *AplpacaFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.AplpacaFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AplpacaFactory *AplpacaFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AplpacaFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AplpacaFactory *AplpacaFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AplpacaFactory *AplpacaFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.contract.Transact(opts, method, params...)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_AplpacaFactory *AplpacaFactoryCaller) INITCODEPAIRHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AplpacaFactory.contract.Call(opts, &out, "INIT_CODE_PAIR_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_AplpacaFactory *AplpacaFactorySession) INITCODEPAIRHASH() ([32]byte, error) {
	return _AplpacaFactory.Contract.INITCODEPAIRHASH(&_AplpacaFactory.CallOpts)
}

// INITCODEPAIRHASH is a free data retrieval call binding the contract method 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (_AplpacaFactory *AplpacaFactoryCallerSession) INITCODEPAIRHASH() ([32]byte, error) {
	return _AplpacaFactory.Contract.INITCODEPAIRHASH(&_AplpacaFactory.CallOpts)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_AplpacaFactory *AplpacaFactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_AplpacaFactory *AplpacaFactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _AplpacaFactory.Contract.AllPairs(&_AplpacaFactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_AplpacaFactory *AplpacaFactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _AplpacaFactory.Contract.AllPairs(&_AplpacaFactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_AplpacaFactory *AplpacaFactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AplpacaFactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_AplpacaFactory *AplpacaFactorySession) AllPairsLength() (*big.Int, error) {
	return _AplpacaFactory.Contract.AllPairsLength(&_AplpacaFactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_AplpacaFactory *AplpacaFactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _AplpacaFactory.Contract.AllPairsLength(&_AplpacaFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_AplpacaFactory *AplpacaFactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_AplpacaFactory *AplpacaFactorySession) FeeTo() (common.Address, error) {
	return _AplpacaFactory.Contract.FeeTo(&_AplpacaFactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_AplpacaFactory *AplpacaFactoryCallerSession) FeeTo() (common.Address, error) {
	return _AplpacaFactory.Contract.FeeTo(&_AplpacaFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_AplpacaFactory *AplpacaFactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_AplpacaFactory *AplpacaFactorySession) FeeToSetter() (common.Address, error) {
	return _AplpacaFactory.Contract.FeeToSetter(&_AplpacaFactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_AplpacaFactory *AplpacaFactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _AplpacaFactory.Contract.FeeToSetter(&_AplpacaFactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_AplpacaFactory *AplpacaFactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AplpacaFactory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_AplpacaFactory *AplpacaFactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _AplpacaFactory.Contract.GetPair(&_AplpacaFactory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_AplpacaFactory *AplpacaFactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _AplpacaFactory.Contract.GetPair(&_AplpacaFactory.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_AplpacaFactory *AplpacaFactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_AplpacaFactory *AplpacaFactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.CreatePair(&_AplpacaFactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_AplpacaFactory *AplpacaFactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.CreatePair(&_AplpacaFactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_AplpacaFactory *AplpacaFactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_AplpacaFactory *AplpacaFactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.SetFeeTo(&_AplpacaFactory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_AplpacaFactory *AplpacaFactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.SetFeeTo(&_AplpacaFactory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_AplpacaFactory *AplpacaFactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_AplpacaFactory *AplpacaFactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.SetFeeToSetter(&_AplpacaFactory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_AplpacaFactory *AplpacaFactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _AplpacaFactory.Contract.SetFeeToSetter(&_AplpacaFactory.TransactOpts, _feeToSetter)
}

// AplpacaFactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the AplpacaFactory contract.
type AplpacaFactoryPairCreatedIterator struct {
	Event *AplpacaFactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *AplpacaFactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AplpacaFactoryPairCreated)
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
		it.Event = new(AplpacaFactoryPairCreated)
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
func (it *AplpacaFactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AplpacaFactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AplpacaFactoryPairCreated represents a PairCreated event raised by the AplpacaFactory contract.
type AplpacaFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_AplpacaFactory *AplpacaFactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*AplpacaFactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _AplpacaFactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &AplpacaFactoryPairCreatedIterator{contract: _AplpacaFactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_AplpacaFactory *AplpacaFactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *AplpacaFactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _AplpacaFactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AplpacaFactoryPairCreated)
				if err := _AplpacaFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_AplpacaFactory *AplpacaFactoryFilterer) ParsePairCreated(log types.Log) (*AplpacaFactoryPairCreated, error) {
	event := new(AplpacaFactoryPairCreated)
	if err := _AplpacaFactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
