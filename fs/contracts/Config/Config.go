// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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
)

// Setting is an auto generated low-level Go binding around an user-defined struct.
type Setting struct {
	GasPrice             uint64
	GasPerGBPerBlock     uint64
	GasPerKBPerBlock     uint64
	GasForChallenge      uint64
	MaxProveBlockNum     uint64
	MinVolume            uint64
	DefaultProvePeriod   uint64
	DefaultProveLevel    uint64
	DefaultCopyNum       uint64
	DefaultBlockInterval uint64
	MinSectorSize        uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"}],\"name\":\"GetProveIntervalByProveLevel\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSetting\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"}],\"name\":\"GetSettingWithProveLevel\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xea7eb5a7\",\"type\":\"bytes32\"}],\"name\":\"c_0xea7eb5a7\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506111b6806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806322cf12cf146100515780636c5ab29f1461006f578063a86593981461008b578063c5c81b20146100bb575b600080fd5b6100596100eb565b6040516100669190611093565b60405180910390f35b61008960048036038101906100849190610f42565b61071b565b005b6100a560048036038101906100a09190610f6b565b61071e565b6040516100b291906110af565b60405180910390f35b6100d560048036038101906100d09190610f6b565b610cd7565b6040516100e29190611093565b60405180910390f35b6100f3610e50565b61011f7f25d761c9f975c5dcebb1527bd6c91c83fc5e16d5a604c09f21248986d83e33b160001b61071b565b61014b7fa80353c277b53781550e8004b1183c216bd93b917d47010877fff9bb52de772e60001b61071b565b6101777ff05d7d6f1ba86c204ce69188f35d2c944728d8559f9f05bc3b0b77400ebe5a4d60001b61071b565b61017f610e50565b6101ab7f37e84a7390b556cf030ee8c20b97ed7d71e8f6135be5db59c68f3562f30873ef60001b61071b565b6101d77fa8c6e9662e351b8c75cce8fb32853a5499a472a41d2029fa271be02ff1bd8ee560001b61071b565b6001816000019067ffffffffffffffff16908167ffffffffffffffff16815250506102247fef666733d531e627a1919556160cc7c161761406a77b5c3de0bcb1da574aa76a60001b61071b565b6102507f12c87d395ab915dd9042a1819ceb3ea5f0ffcf4236420920b69055c3b235c83a60001b61071b565b6001816020019067ffffffffffffffff16908167ffffffffffffffff168152505061029d7f2ccf26ed9c534b634feab9391b63c8cbd26b236f99115300e0e2dfd19bebd9e960001b61071b565b6102c97f3be63c117ebfa011c25170601a6c3bd43cc6a108fee6227571c64dcf891e95b760001b61071b565b6001816040019067ffffffffffffffff16908167ffffffffffffffff16815250506103167fcd3537bff4e09eb61f4824541bf653e4852ebedd49a7be17210ccf11a038443360001b61071b565b6103427f53f421ebc62b0de495bb3db322ac1e998c5683a67d233f3e0858ff65191f6be360001b61071b565b62030d40816060019067ffffffffffffffff16908167ffffffffffffffff16815250506103917f53573f5987f90cbd0cd6921943b89dd0c40ce907c97195c2ba01846f0529868c60001b61071b565b6103bd7f53a2464869bc6d494e715d29a41c0f7fe3c19af2addff420ca32b4b878c7a2f660001b61071b565b6020816080019067ffffffffffffffff16908167ffffffffffffffff168152505061040a7f6a872a919f2f332a8c486e49d9f9fb203c5a0e38369e74901b7198e75bb8c76060001b61071b565b6104367f5e00c08b59e97462d174caf1e2587d125dd222a4dd00ce8ee1cd68b466af6de360001b61071b565b620f42408160a0019067ffffffffffffffff16908167ffffffffffffffff16815250506104857f796c17eacfab6188ca7ed4319a6c13e2c30a7021f434110cf62565c3fb5c4ab860001b61071b565b6104b17f74b5e848b29d1b6c3abc26050b14af5bf073ea45e7327245d59670213551b6df60001b61071b565b6143808160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506104ff7fc9aa0981ac513ffeb618117c2e000e79b436285fd3a3079dfd6fe6354c3d8cd460001b61071b565b61052b7f9f97aa5cb3acc197a8171a9d7bd2edce17b19830724d2d416c6d5f98f0ee0a0360001b61071b565b60018160e0019067ffffffffffffffff16908167ffffffffffffffff16815250506105787f2847aa11c20d9c577abf2394c8ddbb95012a12166c8ef3ccc9b3eb1a43e45ff460001b61071b565b6105a47f5e4bb4bd07dad6c14da1c76335806a1432024f74477b8530e1b32eb3a1b87be460001b61071b565b600281610100019067ffffffffffffffff16908167ffffffffffffffff16815250506105f27f4c3c41793b9fa035a29dfc6e70035d9c21f267ca4e820589cf4da3f606b5696960001b61071b565b61061e7fa39e8a7fdf63f8371cc2bfaa7d3d1b0e5cefd9af8a31902b5e8bbacae06e42de60001b61071b565b600581610120019067ffffffffffffffff16908167ffffffffffffffff168152505061066c7f3850faac504ab2a8611b052165c03b054a970d9a60bce8d540ad78850286b79460001b61071b565b6106987ff0f748406f19a1a4d98e2e36d9d2a905b4016f6935faa847ceb6a9c66d6b2bce60001b61071b565b620f424081610140019067ffffffffffffffff16908167ffffffffffffffff16815250506106e87fe26160b8f051f556343f31e1d250b10b0219c9124c75a16460e67b03508eb68360001b61071b565b6107147fc678ba4c35663e79443fb1e214039601fc842fe5b35dba01f567fca6ccd48de060001b61071b565b8091505090565b50565b600061074c7f614df1bd4327fb0c869ecf3e236a4ee2649289e5c02064dd7411186a1125d24960001b61071b565b6107787f26e89168660b6ccdf25691fcc110d08a39bdf4eceedc969d0c6c5f45594340c760001b61071b565b6107a47f64e445cc373687869264b5e180eb3a6ec2e4f2a7060469ed5e000c019da4772660001b61071b565b60006107ae6100eb565b90506107dc7f9b03e1c9338d4c5e189fb17c56fc1072f0d300adbdba403f8166a52c55e05f5e60001b61071b565b6108087f178b4b8cbbfcac236b7eee87a2db07e39a756a91e7877670da062377915855bf60001b61071b565b60006002811115610842577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b83600281111561087b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610912576108ad7f168f6a8da2e4e7364cd8d88a7a58fde70ca9a0d58bf8e45baa6e1ea0d4c3d5c260001b61071b565b6108d97f5cda96c5cd21dceadeaa8c83bbf06872914dfd160e61114af4e70090f950eb4560001b61071b565b6109057f9a39b5ee9852edb6c54d46a8c0df4a9629702b9e9296d483ea1a433dc0ddab1960001b61071b565b8060c00151915050610cd2565b61093e7fa5f7c1d5d50cb03d7acfb5f1d3aeaab17764001ed69bcbec24ea66205691ecea60001b61071b565b61096a7f4bad6ed4794dbcfe94cea0a6ae415e2f5b0957b7c07a16d8d96ebf035137153f60001b61071b565b6109967f96a16c4b58bd18911a06835b7d96e5a2b134a9205492a09b2e84888405376c2e60001b61071b565b600160028111156109d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836002811115610a09577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610aac57610a3b7f5628b65c4a34916ce8c367a693953bdaf86fd43928fe227089b849bdbc5847f560001b61071b565b610a677f31b9fecfe2d8f6c4e4b5217112397513d85a47e21f120271de45b01b9c26006460001b61071b565b610a937f40884f6f1ebb0d843b2ea7abcac5247e050bdf64783ff787aea3f8f8d428ebcf60001b61071b565b60028160c00151610aa491906110ca565b915050610cd2565b610ad87fd55b0d6139f352e34894955aba3c7926281def55b9550ffa93c042321832a22a60001b61071b565b610b047f4e4ee34ed87474c85792ae3604b75094c383eec58cfa1150bf6a803e1663049b60001b61071b565b610b307ff4b9c6d6804f4e52e7dae60a45ecd834c597a063203dc734f6930ec7758a094c60001b61071b565b600280811115610b69577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836002811115610ba2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610c4557610bd47f6d7c9451c27f2ebbdcddf13a1e0dbe7a84ce0ee333768f0b90d39acf4cdb864e60001b61071b565b610c007f9d6340a64a50ca1ea8e056197b630013f07ef3d5a169e853f0b687d728ba18d160001b61071b565b610c2c7fe666bfb1f53aa51da17df7e6ed561101132f79f5e588ffc7c49339ba62d3292960001b61071b565b60088160c00151610c3d91906110ca565b915050610cd2565b610c717f25b8ff10f37ff975574e3f4de5220dda500ecce328f728d26ffc429af6187cf860001b61071b565b610c9d7fc4f443a7e180c0c4f030453d2101c13f39034589679f5e7f09707b75029c322560001b61071b565b610cc97f031e36f0a055e7aba370ef121b5eba7a68ed4336247eb347cfbcc87c3f67b72a60001b61071b565b8060c001519150505b919050565b610cdf610e50565b610d0b7f2f051ee7d764748f04c000da93c0175a11df665135278c88d5efe5c82da83b8560001b61071b565b610d377fa6797a0969ee35120d3ede94ab84b3b912d59725401c954655cf6d7dd3355b0f60001b61071b565b610d637fb503a6c37b579b8e8005cc7e699fcb9c4efc9ade251c9f481a1d7afd329776d560001b61071b565b6000610d6d6100eb565b9050610d9b7f456352522c85cc1113bb1c513c70af80c0bc8807d1608f523b34a60817a7b40360001b61071b565b610dc77ff7168de228a32c4aa5217eb424bf3150d427c6fc6fff194ba2a163154511133f60001b61071b565b610dd08361071e565b8160c0019067ffffffffffffffff16908167ffffffffffffffff1681525050610e1b7f328190f7d2335fb4817fe4f0a4385abb59145a2cd556d97c613d1bd50afd64f960001b61071b565b610e477f221dcf64f94dfa75f5e3387343b72e3b183c42c3f409826888d316b7e487369560001b61071b565b80915050919050565b604051806101600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600081359050610f2781611159565b92915050565b600081359050610f3c81611170565b92915050565b600060208284031215610f5457600080fd5b6000610f6284828501610f18565b91505092915050565b600060208284031215610f7d57600080fd5b6000610f8b84828501610f2d565b91505092915050565b61016082016000820151610fab6000850182611075565b506020820151610fbe6020850182611075565b506040820151610fd16040850182611075565b506060820151610fe46060850182611075565b506080820151610ff76080850182611075565b5060a082015161100a60a0850182611075565b5060c082015161101d60c0850182611075565b5060e082015161103060e0850182611075565b50610100820151611045610100850182611075565b5061012082015161105a610120850182611075565b5061014082015161106f610140850182611075565b50505050565b61107e81611116565b82525050565b61108d81611116565b82525050565b6000610160820190506110a96000830184610f94565b92915050565b60006020820190506110c46000830184611084565b92915050565b60006110d582611116565b91506110e083611116565b92508167ffffffffffffffff04831182151516156111015761110061112a565b5b828202905092915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6111628161110c565b811461116d57600080fd5b50565b6003811061117d57600080fd5b5056fea2646970667358221220260d486497c4c57da24341b435a0f2536fe9f07d56cb6ccf476cb3052eb511f664736f6c63430008040033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// GetProveIntervalByProveLevel is a free data retrieval call binding the contract method 0xa8659398.
//
// Solidity: function GetProveIntervalByProveLevel(uint8 proveLevel) pure returns(uint64)
func (_Store *StoreCaller) GetProveIntervalByProveLevel(opts *bind.CallOpts, proveLevel uint8) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetProveIntervalByProveLevel", proveLevel)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetProveIntervalByProveLevel is a free data retrieval call binding the contract method 0xa8659398.
//
// Solidity: function GetProveIntervalByProveLevel(uint8 proveLevel) pure returns(uint64)
func (_Store *StoreSession) GetProveIntervalByProveLevel(proveLevel uint8) (uint64, error) {
	return _Store.Contract.GetProveIntervalByProveLevel(&_Store.CallOpts, proveLevel)
}

// GetProveIntervalByProveLevel is a free data retrieval call binding the contract method 0xa8659398.
//
// Solidity: function GetProveIntervalByProveLevel(uint8 proveLevel) pure returns(uint64)
func (_Store *StoreCallerSession) GetProveIntervalByProveLevel(proveLevel uint8) (uint64, error) {
	return _Store.Contract.GetProveIntervalByProveLevel(&_Store.CallOpts, proveLevel)
}

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCaller) GetSetting(opts *bind.CallOpts) (Setting, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSetting")

	if err != nil {
		return *new(Setting), err
	}

	out0 := *abi.ConvertType(out[0], new(Setting)).(*Setting)

	return out0, err

}

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreSession) GetSetting() (Setting, error) {
	return _Store.Contract.GetSetting(&_Store.CallOpts)
}

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetSetting() (Setting, error) {
	return _Store.Contract.GetSetting(&_Store.CallOpts)
}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCaller) GetSettingWithProveLevel(opts *bind.CallOpts, proveLevel uint8) (Setting, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSettingWithProveLevel", proveLevel)

	if err != nil {
		return *new(Setting), err
	}

	out0 := *abi.ConvertType(out[0], new(Setting)).(*Setting)

	return out0, err

}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreSession) GetSettingWithProveLevel(proveLevel uint8) (Setting, error) {
	return _Store.Contract.GetSettingWithProveLevel(&_Store.CallOpts, proveLevel)
}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetSettingWithProveLevel(proveLevel uint8) (Setting, error) {
	return _Store.Contract.GetSettingWithProveLevel(&_Store.CallOpts, proveLevel)
}

// C0xea7eb5a7 is a free data retrieval call binding the contract method 0x6c5ab29f.
//
// Solidity: function c_0xea7eb5a7(bytes32 c__0xea7eb5a7) pure returns()
func (_Store *StoreCaller) C0xea7eb5a7(opts *bind.CallOpts, c__0xea7eb5a7 [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0xea7eb5a7", c__0xea7eb5a7)

	if err != nil {
		return err
	}

	return err

}

// C0xea7eb5a7 is a free data retrieval call binding the contract method 0x6c5ab29f.
//
// Solidity: function c_0xea7eb5a7(bytes32 c__0xea7eb5a7) pure returns()
func (_Store *StoreSession) C0xea7eb5a7(c__0xea7eb5a7 [32]byte) error {
	return _Store.Contract.C0xea7eb5a7(&_Store.CallOpts, c__0xea7eb5a7)
}

// C0xea7eb5a7 is a free data retrieval call binding the contract method 0x6c5ab29f.
//
// Solidity: function c_0xea7eb5a7(bytes32 c__0xea7eb5a7) pure returns()
func (_Store *StoreCallerSession) C0xea7eb5a7(c__0xea7eb5a7 [32]byte) error {
	return _Store.Contract.C0xea7eb5a7(&_Store.CallOpts, c__0xea7eb5a7)
}
