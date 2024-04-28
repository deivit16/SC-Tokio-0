// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TwitterTokioTweet is an auto generated low-level Go binding around an user-defined struct.
type TwitterTokioTweet struct {
	Username  string
	Tweet     string
	Timestamp *big.Int
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tweetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tweet\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TweetPublished\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"}],\"name\":\"getTweet\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tweet\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTwitterTokio.Tweet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tweetContent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"}],\"name\":\"setTweet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTweets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unreadMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tweet\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structTwitterTokio.Tweet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610a8a8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063448027021461004e5780634847e35a1461006a578063c3ad5ecb1461007f578063de440e0e14610092575b5f80fd5b61005760015481565b6040519081526020015b60405180910390f35b6100726100a7565b60405161006191906106e1565b61007261008d366004610772565b6100d2565b6100a56100a0366004610826565b610332565b005b335f9081526002602052604090208054600154918290556060916100cb8282610460565b9250505090565b606060015482111561012b5760405162461bcd60e51b815260206004820152601860248201527f537461727420696e646578206f7574206f662072616e6765000000000000000060448201526064015b60405180910390fd5b5f8260015461013a919061089a565b90505f8167ffffffffffffffff81111561015657610156610789565b6040519080825280602002602001820160405280156101aa57816020015b610197604051806060016040528060608152602001606081526020015f81525090565b8152602001906001900390816101745790505b509050835b60015481101561032a575f81815260208190526040908190208151606081019092528054829082906101e0906108ad565b80601f016020809104026020016040519081016040528092919081815260200182805461020c906108ad565b80156102575780601f1061022e57610100808354040283529160200191610257565b820191905f5260205f20905b81548152906001019060200180831161023a57829003601f168201915b50505050508152602001600182018054610270906108ad565b80601f016020809104026020016040519081016040528092919081815260200182805461029c906108ad565b80156102e75780601f106102be576101008083540402835291602001916102e7565b820191905f5260205f20905b8154815290600101906020018083116102ca57829003601f168201915b50505050508152602001600282015481525050828683610307919061089a565b81518110610317576103176108e5565b60209081029190910101526001016101af565b509392505050565b61012c825111156103915760405162461bcd60e51b8152602060048201526024808201527f547765657420636f6e74656e742065786365656473206d6178696d756d206c656044820152630dccee8d60e31b6064820152608401610122565b60408051606081018252828152602080820185905242828401526001545f9081529081905291909120815181906103c89082610945565b50602082015160018201906103dd9082610945565b506040919091015160029091015560018054905f6103fa83610a05565b91905055506001805461040d919061089a565b8160405161041b9190610a1d565b60405180910390207f8eed20624ab7dd7268daa6258d774554c67a144b8dd33c14b38e9da7ba0e54898442604051610454929190610a33565b60405180910390a35050565b6060818310801561047357506001548211155b6104af5760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642072616e676560981b6044820152606401610122565b5f6104ba848461089a565b90505f8167ffffffffffffffff8111156104d6576104d6610789565b60405190808252806020026020018201604052801561052a57816020015b610517604051806060016040528060608152602001606081526020015f81525090565b8152602001906001900390816104f45790505b509050845b848110156106a8575f818152602081905260409081902081516060810190925280548290829061055e906108ad565b80601f016020809104026020016040519081016040528092919081815260200182805461058a906108ad565b80156105d55780601f106105ac576101008083540402835291602001916105d5565b820191905f5260205f20905b8154815290600101906020018083116105b857829003601f168201915b505050505081526020016001820180546105ee906108ad565b80601f016020809104026020016040519081016040528092919081815260200182805461061a906108ad565b80156106655780601f1061063c57610100808354040283529160200191610665565b820191905f5260205f20905b81548152906001019060200180831161064857829003601f168201915b50505050508152602001600282015481525050828783610685919061089a565b81518110610695576106956108e5565b602090810291909101015260010161052f565b509150505b92915050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b8381101561076457603f1989840301855281516060815181865261072e828701826106b3565b915050888201518582038a87015261074682826106b3565b92890151958901959095525094870194925090860190600101610708565b509098975050505050505050565b5f60208284031215610782575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126107ac575f80fd5b813567ffffffffffffffff808211156107c7576107c7610789565b604051601f8301601f19908116603f011681019082821181831017156107ef576107ef610789565b81604052838152866020858801011115610807575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f8060408385031215610837575f80fd5b823567ffffffffffffffff8082111561084e575f80fd5b61085a8683870161079d565b9350602085013591508082111561086f575f80fd5b5061087c8582860161079d565b9150509250929050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156106ad576106ad610886565b600181811c908216806108c157607f821691505b6020821081036108df57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b601f82111561094057805f5260205f20601f840160051c8101602085101561091e5750805b601f840160051c820191505b8181101561093d575f815560010161092a565b50505b505050565b815167ffffffffffffffff81111561095f5761095f610789565b6109738161096d84546108ad565b846108f9565b602080601f8311600181146109a6575f841561098f5750858301515b5f19600386901b1c1916600185901b1785556109fd565b5f85815260208120601f198616915b828110156109d4578886015182559484019460019091019084016109b5565b50858210156109f157878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60018201610a1657610a16610886565b5060010190565b5f82518060208501845e5f920191825250919050565b604081525f610a4560408301856106b3565b9050826020830152939250505056fea2646970667358221220ff161e7fafbfcbecda8e65ce460b5afc677e6f532ce7efc6dc4b120ad8e4944b64736f6c63430008190033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// GetTweet is a free data retrieval call binding the contract method 0xc3ad5ecb.
//
// Solidity: function getTweet(uint256 _startIndex) view returns((string,string,uint256)[])
func (_Main *MainCaller) GetTweet(opts *bind.CallOpts, _startIndex *big.Int) ([]TwitterTokioTweet, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getTweet", _startIndex)

	if err != nil {
		return *new([]TwitterTokioTweet), err
	}

	out0 := *abi.ConvertType(out[0], new([]TwitterTokioTweet)).(*[]TwitterTokioTweet)

	return out0, err

}

// GetTweet is a free data retrieval call binding the contract method 0xc3ad5ecb.
//
// Solidity: function getTweet(uint256 _startIndex) view returns((string,string,uint256)[])
func (_Main *MainSession) GetTweet(_startIndex *big.Int) ([]TwitterTokioTweet, error) {
	return _Main.Contract.GetTweet(&_Main.CallOpts, _startIndex)
}

// GetTweet is a free data retrieval call binding the contract method 0xc3ad5ecb.
//
// Solidity: function getTweet(uint256 _startIndex) view returns((string,string,uint256)[])
func (_Main *MainCallerSession) GetTweet(_startIndex *big.Int) ([]TwitterTokioTweet, error) {
	return _Main.Contract.GetTweet(&_Main.CallOpts, _startIndex)
}

// TotalTweets is a free data retrieval call binding the contract method 0x44802702.
//
// Solidity: function totalTweets() view returns(uint256)
func (_Main *MainCaller) TotalTweets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "totalTweets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTweets is a free data retrieval call binding the contract method 0x44802702.
//
// Solidity: function totalTweets() view returns(uint256)
func (_Main *MainSession) TotalTweets() (*big.Int, error) {
	return _Main.Contract.TotalTweets(&_Main.CallOpts)
}

// TotalTweets is a free data retrieval call binding the contract method 0x44802702.
//
// Solidity: function totalTweets() view returns(uint256)
func (_Main *MainCallerSession) TotalTweets() (*big.Int, error) {
	return _Main.Contract.TotalTweets(&_Main.CallOpts)
}

// SetTweet is a paid mutator transaction binding the contract method 0xde440e0e.
//
// Solidity: function setTweet(string _tweetContent, string _username) returns()
func (_Main *MainTransactor) SetTweet(opts *bind.TransactOpts, _tweetContent string, _username string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setTweet", _tweetContent, _username)
}

// SetTweet is a paid mutator transaction binding the contract method 0xde440e0e.
//
// Solidity: function setTweet(string _tweetContent, string _username) returns()
func (_Main *MainSession) SetTweet(_tweetContent string, _username string) (*types.Transaction, error) {
	return _Main.Contract.SetTweet(&_Main.TransactOpts, _tweetContent, _username)
}

// SetTweet is a paid mutator transaction binding the contract method 0xde440e0e.
//
// Solidity: function setTweet(string _tweetContent, string _username) returns()
func (_Main *MainTransactorSession) SetTweet(_tweetContent string, _username string) (*types.Transaction, error) {
	return _Main.Contract.SetTweet(&_Main.TransactOpts, _tweetContent, _username)
}

// UnreadMessages is a paid mutator transaction binding the contract method 0x4847e35a.
//
// Solidity: function unreadMessages() returns((string,string,uint256)[])
func (_Main *MainTransactor) UnreadMessages(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "unreadMessages")
}

// UnreadMessages is a paid mutator transaction binding the contract method 0x4847e35a.
//
// Solidity: function unreadMessages() returns((string,string,uint256)[])
func (_Main *MainSession) UnreadMessages() (*types.Transaction, error) {
	return _Main.Contract.UnreadMessages(&_Main.TransactOpts)
}

// UnreadMessages is a paid mutator transaction binding the contract method 0x4847e35a.
//
// Solidity: function unreadMessages() returns((string,string,uint256)[])
func (_Main *MainTransactorSession) UnreadMessages() (*types.Transaction, error) {
	return _Main.Contract.UnreadMessages(&_Main.TransactOpts)
}

// MainTweetPublishedIterator is returned from FilterTweetPublished and is used to iterate over the raw logs and unpacked data for TweetPublished events raised by the Main contract.
type MainTweetPublishedIterator struct {
	Event *MainTweetPublished // Event containing the contract specifics and raw log

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
func (it *MainTweetPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTweetPublished)
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
		it.Event = new(MainTweetPublished)
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
func (it *MainTweetPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTweetPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTweetPublished represents a TweetPublished event raised by the Main contract.
type MainTweetPublished struct {
	Username  common.Hash
	TweetId   *big.Int
	Tweet     string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTweetPublished is a free log retrieval operation binding the contract event 0x8eed20624ab7dd7268daa6258d774554c67a144b8dd33c14b38e9da7ba0e5489.
//
// Solidity: event TweetPublished(string indexed username, uint256 indexed tweetId, string tweet, uint256 timestamp)
func (_Main *MainFilterer) FilterTweetPublished(opts *bind.FilterOpts, username []string, tweetId []*big.Int) (*MainTweetPublishedIterator, error) {

	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}
	var tweetIdRule []interface{}
	for _, tweetIdItem := range tweetId {
		tweetIdRule = append(tweetIdRule, tweetIdItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "TweetPublished", usernameRule, tweetIdRule)
	if err != nil {
		return nil, err
	}
	return &MainTweetPublishedIterator{contract: _Main.contract, event: "TweetPublished", logs: logs, sub: sub}, nil
}

// WatchTweetPublished is a free log subscription operation binding the contract event 0x8eed20624ab7dd7268daa6258d774554c67a144b8dd33c14b38e9da7ba0e5489.
//
// Solidity: event TweetPublished(string indexed username, uint256 indexed tweetId, string tweet, uint256 timestamp)
func (_Main *MainFilterer) WatchTweetPublished(opts *bind.WatchOpts, sink chan<- *MainTweetPublished, username []string, tweetId []*big.Int) (event.Subscription, error) {

	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}
	var tweetIdRule []interface{}
	for _, tweetIdItem := range tweetId {
		tweetIdRule = append(tweetIdRule, tweetIdItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "TweetPublished", usernameRule, tweetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTweetPublished)
				if err := _Main.contract.UnpackLog(event, "TweetPublished", log); err != nil {
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

// ParseTweetPublished is a log parse operation binding the contract event 0x8eed20624ab7dd7268daa6258d774554c67a144b8dd33c14b38e9da7ba0e5489.
//
// Solidity: event TweetPublished(string indexed username, uint256 indexed tweetId, string tweet, uint256 timestamp)
func (_Main *MainFilterer) ParseTweetPublished(log types.Log) (*MainTweetPublished, error) {
	event := new(MainTweetPublished)
	if err := _Main.contract.UnpackLog(event, "TweetPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
