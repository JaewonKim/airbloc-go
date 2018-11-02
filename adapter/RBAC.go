// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RBACABI is the input ABI used to generate the binding from.
const RBACABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RBACBin is the compiled bytecode used for deploying new contracts.
const RBACBin = `0x60806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630988ca8c14610051578063217fe6c6146100da575b600080fd5b34801561005d57600080fd5b506100d8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061017b565b005b3480156100e657600080fd5b50610161600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506101fc565b604051808215151515815260200191505060405180910390f35b6101f8826000836040518082805190602001908083835b6020831015156101b75780518252602082019150602081019050602083039250610192565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902061028390919063ffffffff16565b5050565b600061027b836000846040518082805190602001908083835b60208310151561023a5780518252602082019150602081019050602083039250610215565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902061029c90919063ffffffff16565b905092915050565b61028d828261029c565b151561029857600080fd5b5050565b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050929150505600a165627a7a723058202ec5b7703d986f3b889bffe3e22a045442f734526c17f475e0f411b879a394410029`

// DeployRBAC deploys a new Ethereum contract, binding an instance of RBAC to it.
func DeployRBAC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RBAC, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RBACBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// RBAC is an auto generated Go binding around an Ethereum contract.
type RBAC struct {
	RBACCaller     // Read-only binding to the contract
	RBACTransactor // Write-only binding to the contract
	RBACFilterer   // Log filterer for contract events
}

// RBACCaller is an auto generated read-only Go binding around an Ethereum contract.
type RBACCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RBACTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RBACFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RBACSession struct {
	Contract     *RBAC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RBACCallerSession struct {
	Contract *RBACCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RBACTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RBACTransactorSession struct {
	Contract     *RBACTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACRaw is an auto generated low-level Go binding around an Ethereum contract.
type RBACRaw struct {
	Contract *RBAC // Generic contract binding to access the raw methods on
}

// RBACCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RBACCallerRaw struct {
	Contract *RBACCaller // Generic read-only contract binding to access the raw methods on
}

// RBACTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RBACTransactorRaw struct {
	Contract *RBACTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRBAC creates a new instance of RBAC, bound to a specific deployed contract.
func NewRBAC(address common.Address, backend bind.ContractBackend) (*RBAC, error) {
	contract, err := bindRBAC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// NewRBACCaller creates a new read-only instance of RBAC, bound to a specific deployed contract.
func NewRBACCaller(address common.Address, caller bind.ContractCaller) (*RBACCaller, error) {
	contract, err := bindRBAC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RBACCaller{contract: contract}, nil
}

// NewRBACTransactor creates a new write-only instance of RBAC, bound to a specific deployed contract.
func NewRBACTransactor(address common.Address, transactor bind.ContractTransactor) (*RBACTransactor, error) {
	contract, err := bindRBAC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RBACTransactor{contract: contract}, nil
}

// NewRBACFilterer creates a new log filterer instance of RBAC, bound to a specific deployed contract.
func NewRBACFilterer(address common.Address, filterer bind.ContractFilterer) (*RBACFilterer, error) {
	contract, err := bindRBAC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RBACFilterer{contract: contract}, nil
}

// bindRBAC binds a generic wrapper to an already deployed contract.
func bindRBAC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.RBACCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transact(opts, method, params...)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_RBAC *RBACCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _RBAC.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_RBAC *RBACSession) CheckRole(_operator common.Address, _role string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_RBAC *RBACCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_RBAC *RBACCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBAC.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_RBAC *RBACSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_RBAC *RBACCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, _operator, _role)
}

// RBACRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the RBAC contract.
type RBACRoleAddedIterator struct {
	Event *RBACRoleAdded // Event containing the contract specifics and raw log

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
func (it *RBACRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleAdded)
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
		it.Event = new(RBACRoleAdded)
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
func (it *RBACRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleAdded represents a RoleAdded event raised by the RBAC contract.
type RBACRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_RBAC *RBACFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*RBACRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &RBACRoleAddedIterator{contract: _RBAC.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_RBAC *RBACFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *RBACRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleAdded)
				if err := _RBAC.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// RBACRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the RBAC contract.
type RBACRoleRemovedIterator struct {
	Event *RBACRoleRemoved // Event containing the contract specifics and raw log

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
func (it *RBACRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleRemoved)
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
		it.Event = new(RBACRoleRemoved)
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
func (it *RBACRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleRemoved represents a RoleRemoved event raised by the RBAC contract.
type RBACRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_RBAC *RBACFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*RBACRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &RBACRoleRemovedIterator{contract: _RBAC.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_RBAC *RBACFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *RBACRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleRemoved)
				if err := _RBAC.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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