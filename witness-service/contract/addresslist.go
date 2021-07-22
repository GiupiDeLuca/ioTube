// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// AddressListABI is the input ABI used to generate the binding from.
const AddressListABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfActive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_item\",\"type\":\"address\"}],\"name\":\"isExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_item\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"limit\",\"type\":\"uint8\"}],\"name\":\"getActiveItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"items_\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AddressListBin is the compiled bytecode used for deploying new contracts.
var AddressListBin = "0x6080604052600080546001600160a01b03191633179055610404806100256000396000f3fe608060405234801561001057600080fd5b506004361061007c5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100dd5780639f8a13d714610101578063f2fde38b14610127578063f7cb13121461014f5761007c565b806213eb4b1461008157806306661abd146100bb578063593f6969146100d5575b600080fd5b6100a76004803603602081101561009757600080fd5b50356001600160a01b03166101d0565b604080519115158252519081900360200190f35b6100c36101ee565b60408051918252519081900360200190f35b6100c36101f4565b6100e56101fa565b604080516001600160a01b039092168252519081900360200190f35b6100a76004803603602081101561011757600080fd5b50356001600160a01b0316610209565b61014d6004803603602081101561013d57600080fd5b50356001600160a01b031661022c565b005b6101756004803603604081101561016557600080fd5b508035906020013560ff166102b1565b6040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156101bb5781810151838201526020016101a3565b50505050905001935050505060405180910390f35b6001600160a01b031660009081526003602052604090205460ff1690565b60025490565b60015490565b6000546001600160a01b031681565b6001600160a01b0316600090815260036020526040902054610100900460ff1690565b6000546001600160a01b0316331461024357600080fd5b6001600160a01b03811661025657600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600254600090606090841080156102ca575060ff831615155b6102d357600080fd5b8260ff16604051908082528060200260200182016040528015610300578160200160208202803883390190505b50905060005b8360ff168110156103c75760025485820110610321576103c7565b6003600060028388018154811061033457fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff61010090910416156103bf5760028186018154811061037657fe5b9060005260206000200160009054906101000a90046001600160a01b03168284815181106103a057fe5b6001600160a01b03909216602092830291909101909101526001909201915b600101610306565b50925092905056fea265627a7a723158207106cbee636d71fa92d11f5bea1b7e98875ab2464c410cfa68b873b1fff774eb64736f6c63430005110032"

// DeployAddressList deploys a new Ethereum contract, binding an instance of AddressList to it.
func DeployAddressList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressList, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressList{AddressListCaller: AddressListCaller{contract: contract}, AddressListTransactor: AddressListTransactor{contract: contract}, AddressListFilterer: AddressListFilterer{contract: contract}}, nil
}

// AddressList is an auto generated Go binding around an Ethereum contract.
type AddressList struct {
	AddressListCaller     // Read-only binding to the contract
	AddressListTransactor // Write-only binding to the contract
	AddressListFilterer   // Log filterer for contract events
}

// AddressListCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressListSession struct {
	Contract     *AddressList      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressListCallerSession struct {
	Contract *AddressListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AddressListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressListTransactorSession struct {
	Contract     *AddressListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AddressListRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressListRaw struct {
	Contract *AddressList // Generic contract binding to access the raw methods on
}

// AddressListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressListCallerRaw struct {
	Contract *AddressListCaller // Generic read-only contract binding to access the raw methods on
}

// AddressListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressListTransactorRaw struct {
	Contract *AddressListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressList creates a new instance of AddressList, bound to a specific deployed contract.
func NewAddressList(address common.Address, backend bind.ContractBackend) (*AddressList, error) {
	contract, err := bindAddressList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressList{AddressListCaller: AddressListCaller{contract: contract}, AddressListTransactor: AddressListTransactor{contract: contract}, AddressListFilterer: AddressListFilterer{contract: contract}}, nil
}

// NewAddressListCaller creates a new read-only instance of AddressList, bound to a specific deployed contract.
func NewAddressListCaller(address common.Address, caller bind.ContractCaller) (*AddressListCaller, error) {
	contract, err := bindAddressList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressListCaller{contract: contract}, nil
}

// NewAddressListTransactor creates a new write-only instance of AddressList, bound to a specific deployed contract.
func NewAddressListTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressListTransactor, error) {
	contract, err := bindAddressList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressListTransactor{contract: contract}, nil
}

// NewAddressListFilterer creates a new log filterer instance of AddressList, bound to a specific deployed contract.
func NewAddressListFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressListFilterer, error) {
	contract, err := bindAddressList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressListFilterer{contract: contract}, nil
}

// bindAddressList binds a generic wrapper to an already deployed contract.
func bindAddressList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressList *AddressListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressList.Contract.AddressListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressList *AddressListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressList.Contract.AddressListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressList *AddressListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressList.Contract.AddressListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressList *AddressListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressList *AddressListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressList *AddressListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressList.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_AddressList *AddressListCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressList.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_AddressList *AddressListSession) Count() (*big.Int, error) {
	return _AddressList.Contract.Count(&_AddressList.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_AddressList *AddressListCallerSession) Count() (*big.Int, error) {
	return _AddressList.Contract.Count(&_AddressList.CallOpts)
}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_AddressList *AddressListCaller) GetActiveItems(opts *bind.CallOpts, offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	var out []interface{}
	err := _AddressList.contract.Call(opts, &out, "getActiveItems", offset, limit)

	outstruct := new(struct {
		Count *big.Int
		Items []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Items = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_AddressList *AddressListSession) GetActiveItems(offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	return _AddressList.Contract.GetActiveItems(&_AddressList.CallOpts, offset, limit)
}

// GetActiveItems is a free data retrieval call binding the contract method 0xf7cb1312.
//
// Solidity: function getActiveItems(uint256 offset, uint8 limit) view returns(uint256 count_, address[] items_)
func (_AddressList *AddressListCallerSession) GetActiveItems(offset *big.Int, limit uint8) (struct {
	Count *big.Int
	Items []common.Address
}, error) {
	return _AddressList.Contract.GetActiveItems(&_AddressList.CallOpts, offset, limit)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_AddressList *AddressListCaller) IsActive(opts *bind.CallOpts, _item common.Address) (bool, error) {
	var out []interface{}
	err := _AddressList.contract.Call(opts, &out, "isActive", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_AddressList *AddressListSession) IsActive(_item common.Address) (bool, error) {
	return _AddressList.Contract.IsActive(&_AddressList.CallOpts, _item)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _item) view returns(bool)
func (_AddressList *AddressListCallerSession) IsActive(_item common.Address) (bool, error) {
	return _AddressList.Contract.IsActive(&_AddressList.CallOpts, _item)
}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_AddressList *AddressListCaller) IsExist(opts *bind.CallOpts, _item common.Address) (bool, error) {
	var out []interface{}
	err := _AddressList.contract.Call(opts, &out, "isExist", _item)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_AddressList *AddressListSession) IsExist(_item common.Address) (bool, error) {
	return _AddressList.Contract.IsExist(&_AddressList.CallOpts, _item)
}

// IsExist is a free data retrieval call binding the contract method 0x0013eb4b.
//
// Solidity: function isExist(address _item) view returns(bool)
func (_AddressList *AddressListCallerSession) IsExist(_item common.Address) (bool, error) {
	return _AddressList.Contract.IsExist(&_AddressList.CallOpts, _item)
}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_AddressList *AddressListCaller) NumOfActive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressList.contract.Call(opts, &out, "numOfActive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_AddressList *AddressListSession) NumOfActive() (*big.Int, error) {
	return _AddressList.Contract.NumOfActive(&_AddressList.CallOpts)
}

// NumOfActive is a free data retrieval call binding the contract method 0x593f6969.
//
// Solidity: function numOfActive() view returns(uint256)
func (_AddressList *AddressListCallerSession) NumOfActive() (*big.Int, error) {
	return _AddressList.Contract.NumOfActive(&_AddressList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressList *AddressListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressList *AddressListSession) Owner() (common.Address, error) {
	return _AddressList.Contract.Owner(&_AddressList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressList *AddressListCallerSession) Owner() (common.Address, error) {
	return _AddressList.Contract.Owner(&_AddressList.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressList *AddressListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressList *AddressListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressList.Contract.TransferOwnership(&_AddressList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressList *AddressListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressList.Contract.TransferOwnership(&_AddressList.TransactOpts, newOwner)
}

// AddressListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressList contract.
type AddressListOwnershipTransferredIterator struct {
	Event *AddressListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressListOwnershipTransferred)
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
		it.Event = new(AddressListOwnershipTransferred)
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
func (it *AddressListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressListOwnershipTransferred represents a OwnershipTransferred event raised by the AddressList contract.
type AddressListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressList *AddressListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressListOwnershipTransferredIterator{contract: _AddressList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressList *AddressListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressListOwnershipTransferred)
				if err := _AddressList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressList *AddressListFilterer) ParseOwnershipTransferred(log types.Log) (*AddressListOwnershipTransferred, error) {
	event := new(AddressListOwnershipTransferred)
	if err := _AddressList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
