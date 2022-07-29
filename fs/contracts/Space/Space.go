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

// FileInfo is an auto generated low-level Go binding around an user-defined struct.
type FileInfo struct {
	FileHash       []byte
	FileOwner      common.Address
	FileDesc       []byte
	Privilege      uint64
	FileBlockNum   uint64
	FileBlockSize  uint64
	ProveInterval  uint64
	ProveTimes     uint64
	ExpiredHeight  *big.Int
	CopyNum        uint64
	Deposit        uint64
	FileProveParam []byte
	ProveBlockNum  uint64
	BlockHeight    *big.Int
	ValidFlag      bool
	StorageType    uint8
	RealFileSize   uint64
	PrimaryNodes   []common.Address
	CandidateNodes []common.Address
	BlocksRoot     []byte
	ProveLevel     uint8
	IsPlotFile     bool
	PlotInfo       PlotInfo
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

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

// SpaceAddParams is an auto generated low-level Go binding around an user-defined struct.
type SpaceAddParams struct {
	OldUserSpace  UserSpace
	AddSize       uint64
	AddBlockCount uint64
	CurrentHeight *big.Int
	Setting       Setting
	FileList      [][]byte
}

// SpaceAddReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceAddReturn struct {
	NewUserSpace UserSpace
	UpdatedFiles []FileInfo
	Success      bool
}

// SpaceChangeReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceChangeReturn struct {
	NewUserSpace UserSpace
	State        TransferState
	UpdatedFiles []FileInfo
}

// SpaceProcessParams is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessParams struct {
	UserSpaceParams UserSpaceParams
	OldUserSpace    UserSpace
	Setting         Setting
}

// SpaceProcessReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessReturn struct {
	NewUserSpace UserSpace
	State        TransferState
	UpdatedFiles []FileInfo
	Success      bool
}

// SpaceProcessRevokeParams is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessRevokeParams struct {
	UserSpaceParams UserSpaceParams
	OldUserspace    UserSpace
	Setting         Setting
	FileList        [][]byte
	Ops             uint64
}

// SpaceProcessRevokeReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessRevokeReturn struct {
	UserSpace     UserSpace
	AddedAmount   uint64
	RevokedAmount uint64
	Update        []FileInfo
	Success       bool
}

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// UserSpace is an auto generated low-level Go binding around an user-defined struct.
type UserSpace struct {
	Used         uint64
	Remain       uint64
	Balance      uint64
	ExpireHeight *big.Int
	UpdateHeight *big.Int
}

// UserSpaceOperation is an auto generated low-level Go binding around an user-defined struct.
type UserSpaceOperation struct {
	Type  uint8
	Value uint64
}

// UserSpaceParams is an auto generated low-level Go binding around an user-defined struct.
type UserSpaceParams struct {
	WalletAddr common.Address
	Owner      common.Address
	Size       UserSpaceOperation
	BlockCount UserSpaceOperation
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserSpace\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"addSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"addBlockCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"fileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSpace.AddParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"AddUserSpace\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.AddReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUserSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCost\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUserSpace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ManageUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"_userSpace\",\"type\":\"tuple\"}],\"name\":\"UpdateUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x936f0bd2\",\"type\":\"bytes32\"}],\"name\":\"c_0x936f0bd2\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getUserspaceChange\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"}],\"internalType\":\"structSpace.ChangeReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_fs\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"userSpaceParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserspace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"ops\",\"type\":\"uint64\"}],\"internalType\":\"structSpace.ProcessRevokeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"processForUserSpaceOneAddOneRevoke\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"userSpace\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"addedAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokedAmount\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"update\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.ProcessRevokeReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"userSpaceParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"internalType\":\"structSpace.ProcessParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"processForUserSpaceOperations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.ProcessReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061d45d80620000226000396000f3fe6080604052600436106100bc5760003560e01c80636a4fe52011610074578063b5bc8ee21161004e578063b5bc8ee21461022c578063ed108de914610255578063f1feacbf14610271576100bc565b80636a4fe5201461019c5780636c799c13146101cc578063722b25b9146101fc576100bc565b80633984ff7d116100a55780633984ff7d14610106578063485cc9551461013657806354e0d3c21461015f576100bc565b80630f9fa2eb146100c15780633899831a146100ea575b600080fd5b3480156100cd57600080fd5b506100e860048036038101906100e3919061b996565b6102a1565b005b61010460048036038101906100ff919061bc5b565b6105ff565b005b610120600480360381019061011b919061bc5b565b610e8b565b60405161012d919061cc7a565b60405180910390f35b34801561014257600080fd5b5061015d6004803603810190610158919061bb20565b61101c565b005b34801561016b57600080fd5b506101866004803603810190610181919061b996565b611260565b604051610193919061ccd5565b60405180910390f35b6101b660048036038101906101b1919061bb5c565b6113e6565b6040516101c3919061cbd0565b60405180910390f35b6101e660048036038101906101e1919061bb9d565b61295a565b6040516101f3919061cc36565b60405180910390f35b6102166004803603810190610211919061bc5b565b6147db565b604051610223919061cbf2565b60405180910390f35b34801561023857600080fd5b50610253600480360381019061024e919061baf7565b615263565b005b61026f600480360381019061026a919061b9bf565b615266565b005b61028b6004803603810190610286919061bbc7565b6153d2565b604051610298919061cc58565b60405180910390f35b6102cd7f8bb52b4fd97864729a09aa5073df73f1e5191baf3d7ae35ebf0bf0ee5b7e922560001b615263565b6102f97f4505d0b792326d37084c9843bec76d1311fa85d0705e1311588f1d48878b512b60001b615263565b6103257f277a6a6f1d66187a7793d31ca03765289730204b27c500a99bf5c4195c74de4260001b615263565b600061033082611260565b905061035e7fe815ebd44dabe89ca2f51af0c0dcde3b6ec8a7bffa4058cca9a1ad269e81aa0960001b615263565b61038a7f7d637c7b22816155cf00c4c7f3dd6bb4c7b5702d9dea5fa06a40ca0ac489c31360001b615263565b6000816000015167ffffffffffffffff161480156103b657506000816040015167ffffffffffffffff16115b15610499576103e77ff7041b99835e3ebce206783d1ef8a19efc08755f816d5f158089c838295b40b460001b615263565b6104137f80e96dc30a2f5d57a2f33cd0e4b0d96529ad0efe8cb7f5eeee55711b51143c3c60001b615263565b61043f7fe4372353169058be64cc6ebd20ed390ce1146228411ea359ecd082260fa6929760001b615263565b8173ffffffffffffffffffffffffffffffffffffffff166108fc826040015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015610493573d6000803e3d6000fd5b506104c6565b6104c57fd2b675f5cd718e7bae190ad981107268f99c100ca522403f0df82665990d283660001b615263565b5b6104f27f634439e491daab82d94697cff53e0ab73866a038d47c20b88d65b442b0bf6ce360001b615263565b61051e7fa7425a2b4e30e7e227a8ca7f2201157f6e1c600409973190c4dfb6c80b58dbb460001b615263565b60008160600151118015610536575043816060015111155b156105ce576105677f52d196c9fe640c1a1943594b7eb76eac289588865812ffafbfc96839bc2612bb60001b615263565b6105937ff71d416f71643e11be9d3257a57ec43112490a3fe8ba3d34d7dabef19e1a2dc660001b615263565b6105bf7f3b9f5934bf79c16ddfeb61830626f1f9ec86b6d23bad0ca3d5d22a7220783e3f60001b615263565b6105c981836165e7565b6105fb565b6105fa7f0b0d617378a74e9577df9d89a3f1e955b5a88987bb4d7d315ab9a62910b4fc4660001b615263565b5b5050565b61062b7f9a60ec0fd04f323a277a1b3bf96c575b7e70678364037e16c6f03079ea51050460001b615263565b6106577fd234d7a2e24e1cf7836eba63f8f5a1c2908906b82578881513b16dbb63b5632560001b615263565b6106837ff68f3174180b59179a3fde471f3ab4c980155a66632b211523eb0a749ee19ec160001b615263565b600061068e826147db565b90506106bc7f93af8d75dc621b4fdc05f9efeeb5c917e68544b25eed821adcc78530b703635660001b615263565b6106e87fcd32cf9cb146d39ce1f066e761148f8e7b12a735892b86dbcfe2c3370f6602c760001b615263565b600081602001516040015167ffffffffffffffff161115610ab15761072f7f3a1923645775e86d16051e4f975d9bfaf9e7250cfaa480457068773d12295f9560001b615263565b61075b7f718cdf6cf82fc90d70f91c0bcc425ef708b8dc46220ad630e1054adaae0b52f460001b615263565b6107877f1c80fcc5fd9f1580932d8acf5ac84f4b30cd1481e4d95e8fbb0b68a51a50234460001b615263565b3073ffffffffffffffffffffffffffffffffffffffff1681602001516000015173ffffffffffffffffffffffffffffffffffffffff1614156108ad576107ef7f9b1d2b85f641fd64fef728fb91fa1c6b4651b8539374b2c70fb88abef6540e0960001b615263565b61081b7fba1ac6972e882f3840af0d53debc0556f1b33d9b881482497cae82dd03db02f560001b615263565b6108477f4a9276b330f9dcfcbdb03567ed8dc3428267de843af7f476db1a5a7ebda1ce0d60001b615263565b80602001516020015173ffffffffffffffffffffffffffffffffffffffff166108fc82602001516040015167ffffffffffffffff169081150290604051600060405180830381858888f193505050501580156108a7573d6000803e3d6000fd5b50610aac565b6108d97f10688aad9dbcec0c82a807df1dd0684cf7e156780e9b9a82ff31a5a007abed0260001b615263565b6109057fda4908e7b3dcb5ec9678086cc7370c20e1f3cd58c1a414d576ea2701890f0c8a60001b615263565b6109317f449932c7a37310209d81898ce356a04e973283fa8aabeddc32158963e713cf3960001b615263565b80602001516040015167ffffffffffffffff16341015610a7f576109777fe9a595403b84826d514947b9a63711ae773415d29e7a9c45839a69b9c4239d7060001b615263565b6109a37ff2e384a8b092a72c32634e5b51f928d8c917de853909fd6df5e05312e7fdef9d60001b615263565b6109cf7fb2024282bffdf1c3996697b52e5c8a507ae11c273645e908a0dba674413a9d1d60001b615263565b610a216040518060400160405280601181526020017f496e73756666696369656e7446756e64730000000000000000000000000000008152503483602001516040015167ffffffffffffffff16617318565b610a4d7f88b610feaf79b2c1505e63034586c5533defd051c5ad3e2c25d9590c4624a1ab60001b615263565b6040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aab7f641f5263f3b4d8771b268f398148dda9d11ff745d869438212b91b91420b178c60001b615263565b5b610ade565b610add7f3594ea6fc4108b7ffa717fbde92753692c814a8daf256cb6a220d93d97b69b6c60001b615263565b5b610b0a7ff1db6521e67d8310be67a8a8c14a7d7441e2e8b904bf8a6a36f073358681016f60001b615263565b610b367fbf43d4436a084140bbbb4af1804c05020bfc47a1c176153e171f406f020648cc60001b615263565b60005b816040015151811015610c8257610b727f1bc956a4d7d5561db26c59918e092c093409ca1885eda4008e8a780463166cb860001b615263565b610b9e7f94cc777c0726e00075bc68259b58a21b636d317c8d8776568e54329b36bb56b360001b615263565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663681850d783604001518381518110610c19577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b8152600401610c3d919061cc14565b600060405180830381600087803b158015610c5757600080fd5b505af1158015610c6b573d6000803e3d6000fd5b505050508080610c7a9061d1d0565b915050610b39565b50610caf7fa04af657b6f0423203ade59b86b7be3663c61e7084c6e29ee2bf26cbbde9a3e460001b615263565b610cdb7fbc0422a405e538484b973dc1ebbf7e4389ce5ce200e06e19fbc2768f22561c3160001b615263565b806000015160026000846020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816001015560808201518160020155905050610df37ffffcbbd6f2b7c11521912c214967767e5ea47300a4f5cc4a52278cd94be04b3960001b615263565b610e1f7fd4f77b0edbeb8ed14d78865255f315721af2ffdba70d43daf4c0ee54739f6b8b60001b615263565b7f0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee916003438460000151856040015160000151866040015160200151876060015160000151886060015160200151604051610e7f979695949392919061cacd565b60405180910390a15050565b610e9361a71e565b610ebf7f9b04222d588b0976de17de6fcd5b51bbdba7297d34bf50e82310d41f4825467760001b615263565b610eeb7fa067d7589986e2912fe45f76515ac986ac552311d2bef0d385cd22eb01f7f1ac60001b615263565b610f177f9877b3b92930bf142c42a8984d89b0dfe6cfd3cf1ffb178ab8d79432b1d2410760001b615263565b6000610f22836147db565b9050610f507f9bd050fde2fe32703742b3277a1709edfdf1b8d3f72dc79b2d5ece99e10ea0d060001b615263565b610f7c7f8f08605831ffee06f24aa4a720ee07b8b523b62b385c70574ebe601dc1a5bf9960001b615263565b7f7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e134199270087678160200151604051610faf919061cc7a565b60405180910390a1610fe37fdd2fc421c6b888806810105c9e694283fde0edfefe07c6193220103e8bd8b15b60001b615263565b61100f7fd368f10b9a7652593735085dafe9c8f6bfa5aa758fb628d4364057681047285160001b615263565b8060200151915050919050565b600060019054906101000a900460ff166110445760008054906101000a900460ff161561104d565b61104c6173b7565b5b61108c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110839061cbb0565b60405180910390fd5b60008060019054906101000a900460ff1615905080156110dc576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6111087f13f2d1a75c15f1539ead0415222a269ee63b72017338978714d0a2b75343b20560001b615263565b6111347f88d8bfa543f8baf50810b5f9ad68100bf7d5e3935074ef1e53ac6a82ca033a5960001b615263565b6111607fbfb8612a6005c44a1a9f3d94865993526933b9edd1e205f79dc0f93f43e9232c60001b615263565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111cd7ff0c5e1cfeef07165f48180a33cd9167ae8b42bcb76c6a0c2c0d519bd3b1a49e660001b615263565b6111f97f45e222bd313f9eba9de1bcd67fa2cbab34257ddd517a042afa66837db1960d8460001b615263565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561125b5760008060016101000a81548160ff0219169083151502179055505b505050565b61126861a775565b6112937efc74bc1e763fe03c1558b48dee80e6b2f67585d24037e94ed9ffacbb26e2d960001b615263565b6112bf7fead98bd6c05709604d0ecf7d839d7ae9627495eef8c2472cc6ee47eaf110f8eb60001b615263565b6112eb7fed2c6811bc36ab22ec9fb42aa84996baa3fbae2beab73ce0ae32ab993c635dd060001b615263565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b6113ee61a7c2565b61141a7fd4747aaeaddfb706253b0b418ea7d4926acd09383972c87f64d6540ac172f7b760001b615263565b6114467f033a60ecc01f7e1a938f4fefa683d299c3fb66aa9fca86e52338ea691dd2028260001b615263565b6114727f2c10e5dd4b5728cd2688edcee2be918e9f8d9553f337a73f13668bab4a34c22b60001b615263565b61147a61a7c2565b6114a67f12724a4408602a6ab568f58bc9eaff69df9fa2e212c21ca41b1e0d523cf458b460001b615263565b6114d27f13d9a3291eb32af55ef074126e040b0627a320f4103b0ac2af29a503df90e41960001b615263565b600083600001516060015114156118e45761150f7f951292a230a50ceba2d61f9b2b82b0bc8e4a35dfe751c1282bc26805fe9d460f60001b615263565b61153b7f784a118fdb8ffc036fc10a7f0bf0ffdf8c302fd00f9612b25acac2ed65b617e960001b615263565b6115677f01b9ae254448b15d30aafda74ef84f837e63ccdcca3d61bb9a6855750d5d899a60001b615263565b600081600001516000019067ffffffffffffffff16908167ffffffffffffffff16815250506115b87f8f6cf346e70e85f362f3452ec0babe320316d61e6b051398669323d959db877c60001b615263565b6115e47fb67a7036c14b886932d9d121ab813dc0e78adeea60e50eaa3020a3a68956a5b560001b615263565b826020015181600001516020019067ffffffffffffffff16908167ffffffffffffffff16815250506116387fa6e76cf49de8be2366de1bad2bd9315a43b34590358d45c6008548a84a2913f760001b615263565b6116647f6ef0a409ec57b1a4a1ac902c8423a62b7c3cbe7d289ccb5a44744d6bcdafe8b660001b615263565b826040015167ffffffffffffffff168360600151611682919061cf1f565b816000015160600181815250506116bb7f0102276e48fd40bc01c5fe3a21d5e3f058b3704950f5882d5239886ea038b22560001b615263565b6116e77f4a0df89cdbd6ea8b1095d9abb722c7b4c380f5c6a5f2b8c71fe777a0375fec3060001b615263565b600081600001516040019067ffffffffffffffff16908167ffffffffffffffff16815250506117387ff0930d48acde37aa3e4f8663376100396f2e14887f9828eddbe68421d38c46eb60001b615263565b6117647fb5a00fb370470b812b57022ec1c93ba2c1bebe911251d587099c6b0df3b8eeb060001b615263565b600061177982600001518560800151436173c8565b90506117a77f29713eb32fe92fdffe2c3c7628793b952829679641ab8bd91798026f4fcac21260001b615263565b6117d37fe01d6a8f63de0f00383f126a38341a4dfc7bb6b9dcfd247e4b42d4e95db0a66a60001b615263565b8060400151816020015182600001516117ec919061cf75565b6117f6919061cf75565b82600001516040019067ffffffffffffffff16908167ffffffffffffffff16815250506118457fd6bae6aa3b70e6fb4e2aaf5e6da864bf65ddf0cbf45e1c616fb74b47cc08fc6a60001b615263565b6118717f69ff13835b71e8737d652c90b274225f58a93f69be0d571bc3ab812d41adc79060001b615263565b60018260400190151590811515815250506118ae7f70651119ba5c9ef56097f60fc0cf739174aad1a9ae148fd031485df702b54b8c60001b615263565b6118da7fe31a389003e08d7d32ab9b32bf77466a4dc5ac08c2be177b440d736de25a34ed60001b615263565b8192505050612955565b6119107faefc5af924b2bc143de5c5d99e0c94b7638e96f0d1dc2e43c1567bfe364a5e1e60001b615263565b61193c7f321c2245f08798d3e5a9a1ff8a833870e03b4edf4f4b200864cf4ba56af9b88b60001b615263565b6119687fe4a67e8b333d4f7d3609a1f88fa49ee831b9447d23732aaa4fe836d2aebb4e8660001b615263565b6000836040015167ffffffffffffffff1684600001516060015161198c919061cf1f565b90506119ba7f0c001c1315a72a3392241b149afcc2acec4b82fccd7bdc9fae591ef0b982e34c60001b615263565b6119e67f811b871e1d9cf1325916c6c993dc54690a42c708ab45a4a253b4ad96a20041c960001b615263565b60008460200151856000015160200151611a00919061cf75565b9050611a2e7f9e57d58367567b767ec32bed2c4b7c0f5cc071a50644b9dc4896aff20c94672460001b615263565b611a5a7fa0fa9c2bb19172cf7255053f1597d6a12e30c0c52891f47452575db8b12fcd7660001b615263565b84600001516000015183600001516000019067ffffffffffffffff16908167ffffffffffffffff1681525050611ab27ffb6bc981c795e345487e3a0f585d97cd26f5ddb017ad5da1e782a9ef1cc2ba6260001b615263565b611ade7f654b1752f90f1bcecbf042843736020d19b9051671734462960a33afdb8f48eb60001b615263565b8083600001516020019067ffffffffffffffff16908167ffffffffffffffff1681525050611b2e7fa24a52112398030d793ee669f379db0595815f8f71562c7c1455aacf5a44090f60001b615263565b611b5a7fafaf5b316d0511fa72b4f8eea0c70668465bd0a7396fa6fab63d6afb7d34ee7460001b615263565b8183600001516060018181525050611b947fbc0360e3136a31159418ad27303f443971dccdc5715c45f7a2cea91dc158645e60001b615263565b611bc07f944bfd646a51f4b7f05382e505345f7595e081f8718451e0af9e4ab5c4bd6dfa60001b615263565b84600001516040015183600001516040019067ffffffffffffffff16908167ffffffffffffffff1681525050611c187fbeaa1e0e631a4f9f75d915350ae3423720fc000d08c3b043cdbfe02092bd737560001b615263565b611c447fc52acbb1a13d944add3b1fe22b01e6cace675540561ec2c233ac6f89731123a660001b615263565b6000611c5986600001518760800151436173c8565b9050611c877f7fb4098cdd949b934a410d5f239ff938a7a360a8b12af310a4ca816424bd9e0360001b615263565b611cb37fc42cbd2507ed41633cc1be273a5b1475ec64129d212a4a258a42cb6564d766d060001b615263565b6000611cc885600001518860800151436173c8565b9050611cf67f5eb05c2e1dcb80f9b41a87234cd801771e4b08d1631d7f2ef21609fa1843538160001b615263565b611d227f43d670b9bc6b5457170ad170cbdf7209d4fbf3efd6e6caeba08fb472b8ee0a7b60001b615263565b6000826040015183602001518460000151611d3d919061cf75565b611d47919061cf75565b9050611d757fe52d15698bb257358e79322e838a4f7084427701819bc71d7b406126ae18acdb60001b615263565b611da17fce855786fda36a1a6d62c3aafa5b84a661cd421d0b1a85d65e706fac2064eb0b60001b615263565b6000826040015183602001518460000151611dbc919061cf75565b611dc6919061cf75565b9050611df47fd098edebe5fb3dcab401f5a8414b85fc4d24e76c85f21391c509bc2335363eac60001b615263565b611e207fe92439719f0cbb1953939331e790fab335dcb3daad2e75aa8cd38a1e10273de660001b615263565b8167ffffffffffffffff168167ffffffffffffffff1611611f3757611e677ffe56bb98a0b78db94ab11f039e534be5f436fa9bd98ba9d8231ebaf2f094349a60001b615263565b611e937f76dcdbddd4fc5e12714a57864d33d85f49042c06dbf8706837b54e14790b000360001b615263565b611ebf7f0e598e7702ff1874f92ca0c4a2c4b673a855f9bc52fdc564e903b7afcde3acc260001b615263565b6000876040019015159081151581525050611efc7f4d394f194769213c9ca3062689b3c314f328ca028e74907cae4ea002a42c3f3060001b615263565b611f287fe799b2f012d7c36b9b201b56fc5e190015915f1f882847787616f9f4a93cb67c60001b615263565b86975050505050505050612955565b611f637f6ff606481e59cb6c5a991c2c4ea13ec2fdf50d08a10e4ba3a7718d3b1886860360001b615263565b611f8f7f772085d720af8c8ae29dffb7e9aa8d36ceee9186ddfad2248f0430006150e75460001b615263565b611fbb7fd76c2546a0a47cd048a968b827d81a9a9d9a1b70525358d64d8dfc40df503cf460001b615263565b60008282611fc9919061cfe7565b9050611ff77f30a45de3950b7cb53579e4cfe1348fc26ef3845ebb044ec2cfcf12e6c53bc45b60001b615263565b6120237f5c462cac5f4c6fce7614969240b79b17598417329002d89ef0d6844c0931a9a960001b615263565b60008a600001516000019067ffffffffffffffff16908167ffffffffffffffff16815250506120747f872a8cee1a80042fee580c88a9c1f6cc55d4c6def9c2fab353492fa050fd934260001b615263565b6120a07fc2f0f1e2eade4bac2a5e11df7acc2071f8effcc645b2746a7e06d4d7c01ea31f60001b615263565b60006120b58b600001518c60800151436173c8565b90506120e37fac41b4a64fb7d7a473c82db7e3c83a66fd2c33489cf17079587994669e68f14a60001b615263565b61210f7fb2a315991fe28248dfc4c2785b7dec16d4fd7f738809fc23aa9a77a36a7775c960001b615263565b600081604001518260200151836000015161212a919061cf75565b612134919061cf75565b90506121627f1343b6dd4279a6876e8760773fbc9548900deb89158bb83b095c4fb63375db8760001b615263565b61218e7f9a80bfd1e67a0edc4947d62c6091542af930d56ddd92e516c32880cf437e976b60001b615263565b8067ffffffffffffffff168c600001516040015167ffffffffffffffff16116122b0576121dd7f5d205c204b8a7f11c1ca78bee8cc5825d039b4802515f7fd0d55fb5618c6e19d60001b615263565b6122097ffe5e074a4f2b0344f703b096c9e20c28334f580c74a6aa5ef85e7684ed4feb4c60001b615263565b6122357f6147443af9b101afe5f399da94ce02ee0768fd6dd71630b974ec4796a42c011360001b615263565b60008a60400190151590811515815250506122727ff8ec22d8c0c4d4b310888679142e930c805d1a7e4d98193fab19ab27ca0cf6aa60001b615263565b61229e7f76069a48ac9f65f7af1385104aa161f61fbc0e07d784c35e6df4cc1573605cd560001b615263565b899a5050505050505050505050612955565b6122dc7f4d8d7cb411e95835c18ad239a1ee330007482d41ccc03cb6b1c9aeaf51499c6860001b615263565b6123087f99716b7d38f106c171f73325fc8a195fdc3bb6941db1577fda15e79fdcea058f60001b615263565b6123347f8a74015996043ed35e11c8eb74c06e3008808a1e2aa84ca3d91c8e46f847543a60001b615263565b6000818d600001516040015161234a919061cfe7565b90506123787f0a335feb3118f9eae39f97e04effc2afb011491aab792650ae1e8748bce3c8d060001b615263565b6123a47f7030b650ec564a5d007386f698664078c3397d1d8a330a805c73f669c8c29bc560001b615263565b8367ffffffffffffffff168167ffffffffffffffff16111561244d576123ec7f3bd21fbffa6258dec1f7616d5bcaa24812102fc0e4b8f176020c7f555ed06c6f60001b615263565b6124187fb2e2a0691e8b3da1929b0f144338187c4130f1bc4ceef3326fe549c772997ce460001b615263565b6124447f3746a217d3db99db8fea5024bca252247417a686dd66d7c030b7a01019e9fb5160001b615263565b6000935061256a565b6124797fc5185a58b61f3019aa74da8e96cf013f4c00a1cc587becaad1aeaddaca1bdf4a60001b615263565b6124a57f590a414e8eada1709bb356dbcf1c5f1c504da2f329844244f4651e3cd22e71b160001b615263565b6124d17f378e4fea6ab4fd1d22378d5fb3b0877f206a9af78f6d5d3437af1a20755d98e960001b615263565b80846124dd919061cfe7565b935061250b7fca003b5692d568e491e0a230bb7fbc6370de9ddf07dd5d1f3943571e1a894e4760001b615263565b6125367eb9a2c0eeb48ad4f2bd61f4c27ccb169da890762f9826c6589a8094b58b333760001b615263565b838b6000015160400181815161254c919061cf75565b91509067ffffffffffffffff16908167ffffffffffffffff16815250505b6125967fbd0deaa46517ed8873939383b72fdbb20be8285c5366e3362878df308edd228e60001b615263565b6125c27f45913b6341b4226605260f69055196103d4e947f9c0a4598dff9e6305bf4dc4d60001b615263565b60008d6040015167ffffffffffffffff1614612858576126047feab5d4d6b9e654fa1ae703a011c8d57af4d7077188d9a11817ef33da79dd442c60001b615263565b6126307f6d0d520cf4675842886abb0da77450cfe0fc8fba4c179b52d0fa96ce95dd248060001b615263565b61265c7f73592149cf286772db2b4a06f2beaf1345a2b2fd61a57600e9a2a55ab7e7834660001b615263565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f2cc9a08e60a001518f608001518d6040518463ffffffff1660e01b81526004016126c39392919061ca8d565b600060405180830381600087803b1580156126dd57600080fd5b505af11580156126f1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061271a919061baa3565b8c6020018d6040018215151515815250829052505061275b7f409d03d5c4f1a085c85cf24413f90ca3fd5f87e69fa23bc9722f19314366e93a60001b615263565b6127877f8cfc96a03db25139ed14fe87131a81e3975d4dcef03c126d92c29a5ba3b00c4360001b615263565b8a60400151612827576127bc7f50af1ae06997f8a609c584b13b1964511f1aee1c5533af7047136d960e9ad5a760001b615263565b6127e87f32b0d6e4109e0663257354196a019213ea0363ecd433522edd42eab1c1fe09f160001b615263565b6128147fda0c76a14a38b379bf49ffa453830cdaa7e8c57a39c2cba350812ef66055100e60001b615263565b8a9b505050505050505050505050612955565b6128537f6b53d5e8736c5a15ece3554226843c807c15f53f9e1ccb0a3064de50e3da9d9060001b615263565b612885565b6128847f388ea84ec9b1c616c64f12a1803a229e5a9d145f66615b2cddf1668c6911d1b760001b615263565b5b6128b17fb644d6d079bfe08fa1158d0d872406329cb057d5a8a020975dea4621d18c11d960001b615263565b6128dd7f1d011234fa7f8e1e80ab03ec91f0578bfc828b5468261abb56ad2af9dcddb2da60001b615263565b60018b604001901515908115158152505061291a7f1b5ae6d50f84b67f8ce6afba6d80e5de442c82d557cec1a0e9673263521239c560001b615263565b6129467ff47ad4af06cee3b1a283aa11e4033c5353b4d5f10aa85a35e8070cf7e58d433860001b615263565b8a9b5050505050505050505050505b919050565b61296261a7eb565b61298e7fe3cf0f9585f7392b096c42496b92d5258b92ac29866895cd1a08b7148b196b8b60001b615263565b6129ba7f196047b59960cf54d032cf23fc1613ec503d70e02bf88a13856aa5198c30f93a60001b615263565b6129e67fc0058599697281ba111acac7a77509c3af973175d6b2e72d702d442c96be1b0560001b615263565b6129ee61a7eb565b612a1a7f6750dabd20ffd1d5993ae214fa7756e74f48ced02626f94c8b27b92848a2818a60001b615263565b612a467f7ae3d96ef2f547d23f1e5f66afc278df7afca0f6b86d8ee95cd1f50e779124b960001b615263565b6000612a747fd1f275ee84383c6a767dd9e370d10a4d2f356c491b827b2101290500de33b52260001b615263565b612aa07fca9c80a86c78339c3c5a18540c35f482fd47835102163e7f37f6034185e032a860001b615263565b6000612acd7ec5fc0d3ea69f0a4f3287ba8e1da071a73d151f33bb5558c4f0a319e8331f8c60001b615263565b612af97fd726337b5f42a3d9863085dd790b0812c75eaefaec6eca3a81920471926b648660001b615263565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e638760000151602001516040518263ffffffff1660e01b8152600401612b5e919061c9fd565b60006040518083038186803b158015612b7657600080fd5b505afa158015612b8a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190612bb3919061b9fb565b9050612be17fa707970dce83fc6b45dd61497af268650aa53ff143de25e39df4709dc26f272f60001b615263565b612c0d7feacaf4963e10c2cf4b5b072369330d06476c6a15edf378040bed4705f524ab4060001b615263565b6000612c1c87600001516177b5565b9050612c4a7fbc8a9f50c1b71e24702cbdb38b134a4fe06a3f7feda8edfefc3e604c2b9e532660001b615263565b612c767f031d68bf13afb8869bce9af8e664af47be8ac6c91296af733f1270332de24d0160001b615263565b60016002811115612cb0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460016002811115612cec577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff161480612da7575060006002811115612d4b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460016002811115612d87577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff16145b80612e43575060016002811115612de7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460006002811115612e23577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff16145b1561338657612e747f453ef1968e1dd9c949c4ca33df8af8c937e3e1e76cdcbbfb4fddc2873ee3ae8b60001b615263565b612ea07f5f8d92e8face228f17879aa6c98191a9c9f0b3ec69f9b20927316e899efe192960001b615263565b612ecc7f684ead35571d5f43c0017e721a68f3653890609e10860fd649ba9c9f1e3e885060001b615263565b612ed461a821565b612f007f1227fc76d153e987806667edd2c1c25d8831cac58ac97e4286e6859a9ddaeb5460001b615263565b612f2c7fefb41fb6310dc4dca0757d7a87d16a43cc541985e7e7bf5c93cfdbcd641ebcc760001b615263565b87602001518160000181905250612f657fc3818ca7470bae1aee6b6ca54f8dbe707524f2e726f883929d8ce7c4becb589d60001b615263565b612f917fceb9c2f333f94648256e36307be87e530cb93233c021e2bd10a45ffbbbea604760001b615263565b87600001516040015160200151816020019067ffffffffffffffff16908167ffffffffffffffff1681525050612fe97f5997172babf608ccd31f72e472f401bbe220ab092fca0bb4c8ab684ce58b3b0860001b615263565b6130157f3480d8d327ed7e904bd1033af3f74179beada1a13410f76da81637805a07bf7160001b615263565b87600001516060015160200151816040019067ffffffffffffffff16908167ffffffffffffffff168152505061306d7f3a87eac02b550870b0bd3e7d985b1a8e6412eeeac1568ccf6b5e1c16370f3bcc60001b615263565b6130997f29fa39f509fcc3ab38cf5f19f8134da7b3c32afc4582593d35bae18c032bdda860001b615263565b438160600181815250506130cf7fbcfc6f89925080f0e414c0d3d15c0a69c5cfd664144fa5ce5f644adba50cdee560001b615263565b6130fb7f395e04a2e8eaff9c81246a6db1be15c2435fa03f20e667d49f88bf43a45a3f8e60001b615263565b876040015181608001819052506131347ff36ae94f99ba58c0bb431e1659e8d2556e991c1c2f62129124a16653e393478260001b615263565b6131607f260bfc92d38c5f1244cdc02a5f8bb95f688bebe5a2528702a14ffea4a7ea269060001b615263565b828160a001819052506131957f11ee1b51fbec2be3ce68ca0e1e1c55f4522c75c012f655212acf2039d5710fc460001b615263565b6131c17f3c92d0f6919e8235b609612338070d74cf83df94cfef8b236a1d659312a98dc160001b615263565b60006131cc826113e6565b90506131fa7f32725c6f92b0902fdb78e8ac60c89087bc78167a32693bea9100e952d59c42d360001b615263565b6132267f371754bb02822065b90e63846f8c9e020957055d4df7326846528987f4adebdf60001b615263565b8060000151816000015160400151826020015183604001518a6000018b6040018c6060018315151515815250839052839a508490525050505061328b7f160196be3134bcc9b2e0dd5f6df698d2f5fa6d1f45d168d71d3ee2695529064360001b615263565b6132b77fa1785f792c97e83f192b05277754591b9ed2d6aa452d79a22925367b0007353360001b615263565b8660600151613353576132ec7f064667a1921242fffcc25bef1d97a762810bbd653cec452dfea043bd4a79900e60001b615263565b6133187fc644bd9ab007713bec49d3928238a61fa972409628121ceef057beabd4af6d5160001b615263565b6133447f5d49b9f877edc0123b8493121850eb890ed3b42da6b8ad3deb9f7aede2458adc60001b615263565b869750505050505050506147d6565b61337f7fbc3cae88379e9e21442102f2ce04a1944febef065ce86618d42aa5a28fffd9bd60001b615263565b50506133b3565b6133b27f8ff529e9df233ed39a0f573699e1a0ed55a20e0b7eba11f51c71c5508f75bacd60001b615263565b5b6133df7fa6858e705f48ed73523f46837c1749fe92f3058c6ecc38ba309409b12893a60160001b615263565b61340b7ffe3354d65a9f147fbf74c6d041401f3c3cd53e2df0c134fdcb61645e7a1dd13d60001b615263565b600280811115613444577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460028081111561347f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff16148061353957506002808111156134dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460006002811115613519577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff16145b806135d4575060006002811115613579577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046002808111156135b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff16145b15613aa5576136057f9fb83b469af2a0fb7c1c488df757b019893133382101b3f65b9430199e484ed060001b615263565b6136317f8f8442ae415a43afe9210c3f4ec3a7213bf212f9b5b37bb9909408f6414ee3eb60001b615263565b61365d7f3e05044af65cf380010d9488b84e8ca3bb673ee13cd3a293f51935d1b350210960001b615263565b61366561a877565b6136917f6ddfb2d906473d8bc56e70f63c3052a2e15bdffef5440f19f5bfb7e3e82426e060001b615263565b6136bd7fdab5792ffe0769c87bf753adf400339ace4133767ccc3c774036b313d759929660001b615263565b876020015181600001819052506136f67f9afea905192300fce9f66437becba9633630a81e081b924cc8b60e5a33ae45bd60001b615263565b6137227f5b157c2b1d433636d63f6899bc2a70baa543d19c00f753fabfe8b7976e3ae33a60001b615263565b87600001516040015160200151816020019067ffffffffffffffff16908167ffffffffffffffff168152505061377a7f1aba87e83c8bc6a87929f3deee6d7ce0d1ffacb35b3100d7ae1fd744cbdee94d60001b615263565b6137a67f26b8800f4728686438a1a0ffdae1f3fdbd8fd2d23523461a9c8913feba74c59e60001b615263565b87600001516060015160200151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506137fe7f99ba146f9866fc6807629955617bcd7b6fddb2f5e60bf5364dc83e48483e8d1960001b615263565b61382a7f127ea89e7038355519495ec016047f89efc678212ef94633f33cfe0734f0c33060001b615263565b438160600181815250506138607f9423bd6da49d6d4023e5da76123b27d3ef393e128c4345c68fde213d12c8497960001b615263565b61388c7f4d4a88c265c1b4aef851b6a4d17d374465a103d2b4f62f1c8c4a45f1e4326e1360001b615263565b876040015181608001819052506138c57f25d06c88d4ea8f67ad9538800b272d349d61e6289030d844bbaccab7a911f04960001b615263565b6138f17f8f3867d6d01ad63036b4fa3fd97186074b6b79c120bf4133bc8ae361199950ab60001b615263565b60006138fc82617bb4565b905061392a7fff2dc7d6290baf2824270f802a4ecfd963aa76646eed5faf0f0c0ba8bc02142f60001b615263565b6139567f6ce4d12c726b7bbbaf3380a98c5e88e67b04cd97b4b4fb4a098754295d0f241d60001b615263565b806000015181602001518260400151896000018a60600182151515158152508298508390525050506139aa7fa2e3687551b0395546b07a1c82b5063f0b83826b3ba1a402da48b21035567c4e60001b615263565b6139d67f926de0daeae0d5a787b91bf56a2edffa6cb81cf4209249e5e7a4cf37d604dd5260001b615263565b8660600151613a7257613a0b7f131a2ee002f17c228932d137d6364db0f49853185749d433ebfd56ce6031cc2e60001b615263565b613a377ff5c8019d582f632c15e46ad78c53741bc74b8a7e4488a4297f6c5d990bc71e6160001b615263565b613a637f93548f689180a7e347f1311557bbdd06a7f478e4aa2cb6e94d6475b88f76d1f460001b615263565b869750505050505050506147d6565b613a9e7fe95852f5ab51317a30451c8f31db792767c116a4e3cb3600154f77e973a53a1e60001b615263565b5050613ad2565b613ad17f88536a8bf5ee0265732b3477b564bd414bd9e27a136ba8f205d90cdc59e74f8860001b615263565b5b613afe7f4ce8ad030f3e4b260652460dc78fc89a28fbdc76e6265de967ecbac2dbc2004260001b615263565b613b2a7f9737556e0a04390765dcea32c8694ad4b6bcfa91a07309ce5e681ddeea6a6f5360001b615263565b600280811115613b63577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460016002811115613b9f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff161480613c59575060016002811115613bfe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600280811115613c39577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff16145b1561411457613c8a7fc2d13df0f322a5e99231a3c359d1ed0daf2dbaa2689d367a214935f17d2c02df60001b615263565b613cb67f2177eaa18e6ecdc13eda6af70761af9854f71782fe6003ac1103b1f02df8920a60001b615263565b613ce27f66d774dca648645104b486c1599ebec40598d5ee3da7d63e5c622b211c5bedc760001b615263565b613cea61a8c6565b613d167f02b6831437861de01aa78ae35ea6e99935f986148e6b27556ac9a0e4e7feae6760001b615263565b613d427f47e092c79db5891ebd91a3be0838d96fd2b087cd7ca411faedc7489b51d3aacf60001b615263565b87600001518160000181905250613d7b7f42db01c4dc7d4ded7786cd13b9157a3440cc995b8d68e84be00e654f0afe030560001b615263565b613da77fdbd707f92ce35bdee0786e7d5fcfd0919dba25f20b4022a99f7ce16a9dec224660001b615263565b87602001518160200181905250613de07f0ce1b17c7ed68f6637499472446c1a8584b1b29cc3659cdfd0c38bfb2d208f5060001b615263565b613e0c7f6a3f98b49cc31da6e864d3dba8b36b2841dc470bbf9ccf79c04d84f7ee9258ea60001b615263565b87604001518160400181905250613e457f050847cf3f4ca7cc7dc30574a4d9934b938f4cafd96df3f411d9ecbc4900d8af60001b615263565b613e717f6fe481f26bf9e5e498e800c3f7c69e4312954aab4e4c2546b6a69f6a31d2179260001b615263565b828160600181905250613ea67f90e0de42b0db3a791d4795a58d619341f43b93eb3da6ad93b04b60f0d3edd8e760001b615263565b613ed27f97fa448bc21e6141c0795233b1e16d52d4a547fff91f83cdbda148331cc195c660001b615263565b81816080019067ffffffffffffffff16908167ffffffffffffffff1681525050613f1e7f92bac88735c3bf199e456cb8f7e3296678e6e74adba0408af4b03f8589e26d3360001b615263565b613f4a7f286c930ca1ab87b7030610c58eee03f417f96c4a97e0f9e77cc20bd485d70aee60001b615263565b6000613f55826153d2565b9050613f837fb6935adc75219379e597875e1bc3a0c98a75dbfefcb564827c4122aff3214c9960001b615263565b613faf7f3d614200aaacf2e0021dac10df8b3eeae41a04aa8c1fbd552d63e4eeafe9b21e60001b615263565b806000015181602001518260400151836060015184608001518b6000018c6040018d6060018315151515815250839052839a50849b5085905250505050506140197f1379cb2bf5cdfcc3f3d07121d691ed1496b82abf7a439be225f2dfc85a39e04360001b615263565b6140457ff9d937b4f12d3558f3e77e91fe936c4e807402926970abbb8706fb46268683a560001b615263565b86606001516140e15761407a7f4b4b5128f239c5f865a6b49a00c92f1341df6db215ad584459603b30d2eeecdc60001b615263565b6140a67fe41c419f54bfa564ee8f99e0daeacf7c79af9a957f2cd508f649319f909cf08960001b615263565b6140d27ffa68c660f5b7383efc1716e2fbf5242ed22ca7ad23bd3f9a48117f140e16cf4c60001b615263565b869750505050505050506147d6565b61410d7f83b4bb87c2d166c24292671d5809522db1c73777b3cfdd1ab48781e38eaef8ec60001b615263565b5050614141565b6141407f559f866714fb71bed8fe49bcecf697d394996900e5d1650f67b62d2a1339a88a60001b615263565b5b61416d7fb75b6dd16204d46ee15070d5dd433cb01ad788c4710f86ae93630bacdc61fc5860001b615263565b6141997febdf4a387a08e372921749e9f0f5a3458abda4aa1c8fa731090d1d835b3671d660001b615263565b6000856000015160600151141561423b576141d67f71fc164c5bb667a30dee3422acd3622551cf2c6e7f5adecb1e163f9f1be7ba9160001b615263565b6142027f056f5b08536be14836f2aee449dde376aef12c1e547bba5a9b61b4f9f92603d060001b615263565b61422e7f3d6a39a76826496d7f8542341f6e3df22b40b131ed3a49bb7239de5e9900301360001b615263565b84955050505050506147d6565b6142677f6c089702ec6dcb9b75058e3f7235ba49072306a2d6ab7f248ed62f58497dabd160001b615263565b6142937f3745512734cd12f074dc345747e07d60bf9f0ca25a09d26eedea0e288ae38e3260001b615263565b6142bf7fa67b306a1731b2f40e69770773a55a37fe0a3ba1156718126afdeb5c1ec1f67360001b615263565b43856000015160800181815250506142f97f5fd6ce8703d01c2101a6e648b7956c0b61619b874cc07765681013e1266dec6b60001b615263565b6143257fd816bfdac6dd289fa636c9356a7377a87fbb5e543b3ebbe3ab26bf91db64d9cf60001b615263565b8267ffffffffffffffff168467ffffffffffffffff16106145285761436c7f5a51a3ff99d364079f35b6ebcaaf3906f502c9993ed61a25b23772be080ba6f760001b615263565b6143987fca664c1df81c2df23761c7a1abcc8402962e92854fee710ed755131b8422a1eb60001b615263565b6143c47f698b6df1ba9a5ce454963efae00085dc4e2715d4be7141969bfb0a364f676a6860001b615263565b86600001516000015185602001516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506144347fd806e7ad007ec8a1713f7961e4bf9471f6b188ab044e1fac784b536abdcae6d160001b615263565b6144607fb8236cada4335cd77b3fc6f1dd8e84ee9b8b2ff82d15ea22081039cd705b39ec60001b615263565b3085602001516020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506144c87f6b4a1cde9860e2dd6304759c1df02d251f6a987f13d433f1827be69bdaef87db60001b615263565b6144f47f350eb13aa2792445277a5c69f1d5f049e4695e07dbab53f4e13d726bd22736ec60001b615263565b8284614500919061cfe7565b85602001516040019067ffffffffffffffff16908167ffffffffffffffff168152505061470c565b6145547f5a6e27e7c26c2c169e28af32da0adab0240ef13e9297ae7b31d31a225fbc048a60001b615263565b6145807f99c08f2043e478c8c684c7d7e61793383696f10d5f7ee49b727999063a9f696460001b615263565b6145ac7fada9eb5c75213df15ec4a482100538299802367c72c69dac0888a9fd35dca07460001b615263565b3085602001516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506146147f677ed4205204ed0af2636eda2b2cf98ce299a275cffb53936c09ebbeb3656d0460001b615263565b6146407f27be00ef9ba7f2c2ff317cb057058bcb2b6ff9ef3eb6f30480f551282bc8080f60001b615263565b86600001516000015185602001516020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506146b07f9337375880f044346536cd84e41d8c45c5b36a709b9d9f0599ef014bfaa2320460001b615263565b6146dc7fe079c0a01f3ffd6631066c1b8da9d49ea7a9d55edc20d52b0f6b7a22f2f1343d60001b615263565b83836146e8919061cfe7565b85602001516040019067ffffffffffffffff16908167ffffffffffffffff16815250505b6147387f0303ad19a3e874ade3f33f3d64c9c36f19a4cbc4dfd3b33e7d229ee38ba8859660001b615263565b6147647f8525a28a6c7c7ec76d4aab4c72f4ff861722105dede647106e2ecbd4ca93b2ec60001b615263565b60018560600190151590811515815250506147a17fad61358647b3193952ec4ae76fe090bb1053ccf1cf48700c5d3d01e62845e45e60001b615263565b6147cd7f165156570d0ead0c43cfbd943c309f4d5a0184c1cac312839341d2dc2f5e130660001b615263565b84955050505050505b919050565b6147e361a911565b61480f7ffaba05c5df161e99343da782165e6cee3c1477748d41b0f99f3a3436f44dd78860001b615263565b61483b7f967d98f6fd3c2b18e9037d965c530f4ecbae3c71a2e55adb59772ee363ff67a360001b615263565b6148677f9e1d5262f23e8fb64aa0ba22602ccb1b25b54975afd89e955894f39cf43036f760001b615263565b61486f61a911565b61489b7fdb698f2d94d3874d24883215113837ac33e99385e1e9fac562ea6f1fc926f89160001b615263565b6148c77fe0785b1f849877056febfb4e9385c79cc0e80d553098ed23fb4f5795f3ba244860001b615263565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561493157600080fd5b505afa158015614945573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614969919061bc08565b90506149977f74a5a2f3ef763e7b872aa4c20fc633e934b74714f07bc613bd5a9e0a92243f4160001b615263565b6149c37fa164f0425e4e35a2644364b4b3b1ce38f87911069a9f3a0a6c750096b14dd15c60001b615263565b60006149ce8561886c565b90506149fc7f86eb523554057999aba9337084cc7288399b4b2fa19f64708af4efd9b306a6b060001b615263565b614a287fe756e22c9a38f4f2135b8c116c0a36ded5cfc7f2d4657b12f1917aef894e697560001b615263565b80614ac357614a597fcb97a26397d69b17f4285317425c4515968d11de36fcaeef14b3d542ce63cf8460001b615263565b614a857fd2f9c84b1ea78ad29a5a813a747d14f57ce7bf94a435d86825bd170e40a8243c60001b615263565b60016040517f37159131000000000000000000000000000000000000000000000000000000008152600401614aba919061cb3c565b60405180910390fd5b614aef7f93531c79fce378826520821bc67557742a8a093f1e3907d8415cfa2e46517ec860001b615263565b614b1b7f3345d341564cf064aee4cb178504227dde6c045fee6ed7b1d9cd7ffa48de423960001b615263565b614b477f5a54b810131886c219213ea0f8e6f09be4da2a6f38f7f9ebab7c81190405597c60001b615263565b6000614b568660200151611260565b9050614b847f8ec232cb3f91e9ba3fe701eee2fddc4bcce014ba71896fc1328c8279a0777ab460001b615263565b614bb07fa807e7233deca33da1f3c9c6962a04a3cb1d456c371709da8767ce566513c4fb60001b615263565b60008160600151118015614bc8575043816060015111155b15614c6257614bf97f3ce17805bd3fe0d77f3fd3ab246d262346f4ef9e3b651134228ccc09990273b360001b615263565b614c257ff08ad95e8d38e891d8ea35a795546db11584e2f4db9735c864ff8865bb3fb3a160001b615263565b614c517f1c478c7dcf3a82e005a578588691c18ea4df14a543f66a263df48484b70d024960001b615263565b614c5b8143618fc0565b9050614c8f565b614c8e7fd96cc1a153798a2c728fc515e5b4c708a81ad05bc542d5a1e5817fd061d88e0360001b615263565b5b614cbb7f2518313330ff7567596251814fd25a7dbf0f89aa6e26e05a14988b488065bf2160001b615263565b614ce77f6e6a7a4d878ab6ececf85dc6f14a4e6c93dc4fe5cfab5fba44a213247cfbf15a60001b615263565b600081606001511480614cfd5750438160600151145b15614eb957614d2e7f354c0dd6fab8f00b519f0ba630a051b50219b04acf258cf5cefa6b1c211a202160001b615263565b614d5a7ffd542f18a00f90b748fef00b20082b2d50413920f9c77f06349971cb8a16854760001b615263565b614d867fb661afe5ea6a457539de304c7c94fcb3f0ad256193170dffe031fa202d4e407160001b615263565b6000614d92848861920b565b9050614dc07faa3dfa09784c30f169088642c919f03bc556f2b022c391746bf5192626e3d13060001b615263565b614dec7f6d44b8b1fd1a1c715e534d5d69b19958a0a0f7debf79b7b5951afd89d06fddfe60001b615263565b80614e8757614e1d7f06a02cfd055f027d9326d3a404bc81c1837677625535766d6e9de60e94ad38c160001b615263565b614e497f35f6c940041a3f599f26a8f48f99758c33bed873bde39f0b48ebd4424c34acba60001b615263565b60026040517f37159131000000000000000000000000000000000000000000000000000000008152600401614e7e919061cb57565b60405180910390fd5b614eb37f548f648af75592f9ecacacf210b5b27382284dddd8eaafa232729739725ea80360001b615263565b50614ee6565b614ee57fbdca118167771a95392d87b04f922472cfd60ba09aacaf89d366679803ec893e60001b615263565b5b614f127f9143e447333baf1a6ce75ef1c9cc108e6431a407b102e30078b0415e1c80364a60001b615263565b614f3e7f9648a3a5029ac0fe194309bad01836c5b3636226af74fe9bf042e41b6478d41260001b615263565b614f4661a93e565b614f727f52602c2400bcb384a8467cf0ce1d533861a847a5ec694ee6d9302c9c0861dea560001b615263565b614f9e7fd7892f922c199c64aa531dc28fdee7f7fb084a9887480773b0a784d6bdd928cb60001b615263565b868160000181905250614fd37fe540420e7ec96b7132503ac694297f6a4b8db79cfe0f87abdf5819f3168bdb2460001b615263565b614fff7f7691d3cd7b16a1595b9d4d1253bc7cd0b3e696fd6431bbd94cd6b4d7830d906260001b615263565b8181602001819052506150347fb44e8f3b142a3916b9a06b6effefba5784b4f35db95e205c82178ac3d966c39260001b615263565b6150607ffd80fc91e002ac69c59c1ccc2a880c3921993080430b75c5252f5e5a005b3c6f60001b615263565b8381604001819052506150957fd7afb1924ec778fa21b7fd5f537ca70a9427d030ce9dcb9ced3ee7d33c2bf12b60001b615263565b6150c17fefd41e028075fc5fecf3cd748bf4b7270fe7c3c8c9e9020932b64b7c4329496560001b615263565b60006150cc8261295a565b90506150fa7f4fbb53b496a4db726b515a3938f6dd6e791c79f367eefa399e17eefd5280350860001b615263565b6151267feea9d0bacb23f46a4ba6957beb9a4d52cc5ed0dabf8d3f31e194ccee900ebe4960001b615263565b8060000151866000018190525061515f7f5cc019e8584c12cdfecadcb0d23288eadd1fecc51e3c518c1fcd16d9d90378f160001b615263565b61518b7fba4c95f312a59479ecf689c2e447cc6e65f64032d2d04ff71bda72043a63af8e60001b615263565b806020015186602001819052506151c47f3f9e0bd02415f26553fcbb7ef6296798ec55fcf847d5b49d52928184a9be6b5f60001b615263565b6151f07fe0deaab12a602c67db87e44f308d34cbe0f220b2b8d8353cfea6abed754c209b60001b615263565b806040015186604001819052506152297fc339e28a47dc03426354b6ff1394f590769b36f016ee472048b4637bf0f2066860001b615263565b6152557f267d6948f5d6d26f9aa332bf01197042be4ca2ed1e895200add21ef257a1e73b60001b615263565b859650505050505050919050565b50565b6152927fb6e5ce5cd3fd028a855762cb6d96803e5b499567ace32efa4f3ad979c1f3841e60001b615263565b6152be7fcb90f702271b0de9d7930e82621d47ad7e248a6ab386b22e0179cee866dafd3d60001b615263565b6152ea7f4333e8c69ec6d8a6aeafb50bdde7fe3a1d5a89cb5b8a58ee02fd51863f9b4d9560001b615263565b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010155608082015181600201559050505050565b6153da61a971565b6154067f31ae3f141df72a1b0dad74314e705e211712eaf189529a2619732689877aa2b460001b615263565b6154327fb989112e36291bd3f990665f3fbc19e03e9adac9ea28da6f3b6dc5672d8f93fb60001b615263565b61545e7f44f0e6eb5c69924c17a3ae1b3e0eb442ff5bb50ca40706a0590ab7551589763560001b615263565b61546661a971565b6154927fe476fca19c902c5d38c0ea73f036492a5c7cf17541d1ae50a2d1dc5ea17a811760001b615263565b6154be7f8b09a9fb648a564781e3de9466d22f3b3c3ad02921fd946cc8c1600322d1982a60001b615263565b60006154ec7fbbac7005351c470bfa3f80f7c5f6a060af3984df17f8e54eb391c0aaa20d3cab60001b615263565b6155187f19d013251fde9278f98dc3540018fc48052f3a7b60c7bcf67b068274ca63d1c760001b615263565b60006155467f719f46dfa1c6c38ad83889a207f8bad5c72ffe9458c68557dafa35091da36dc360001b615263565b6155727fb38d7152818cb62bd7fb3dcc4bec72ffd01ee0b4ef82a2d6d3c01c9e09f9beff60001b615263565b60006155a07f1a132b87f6cd1546fa26559281306f731c846c70a4acbaddf5eccc29433d65b960001b615263565b6155cc7f943998d7b90ac5494f562f001fdaab85579b5a362b7f1898d52faab2396d44cb60001b615263565b60006155fa7f4b53dcfe9baedc2d660b7a8a6c1357dd3e5bea19715846bd38929513d29166fd60001b615263565b6156267f03b7df8fd2b485d62ec3c0fb7c57ebb32f9f7e71e95ffe99d9bcd7098bdff9b160001b615263565b60028081111561565f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111561569b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff16876080015167ffffffffffffffff1614156157c2576156ef7fadab3439e2c7ba8dffa40b4af78989e7c286da0596b6787a95c390f7fdd984af60001b615263565b61571b7f1e0effea1669d14bbc3d73994e0a26efcdd69ef1487dabbc5eb3b64871c686f860001b615263565b6157477f32f2a90e0ecc12194b6fe4db5109088f439ce354904ec7e821bf83f0f746ad0160001b615263565b8660000151604001516020015193506157827fc8e8c318ce461d15ed3c66cc8082a15549187c2c366c0d8e7bdf7ee50151382460001b615263565b6157ae7fc671d98c522492df658d240cf674531a83fc62cd36e226628581567eaae612da60001b615263565b8660000151606001516020015190506157ef565b6157ee7fc1bc4424df1dcd5ae20ca6139a2d7fab46afb7dda3ddc68840f2f14e9a51193860001b615263565b5b61581b7f315c7a3ab0b09055592397803d1deaf66863958f94189c99d9096940d8e0979860001b615263565b6158477ffff7bf9b023c7dcf9a7ad8c3bc11bc0500d540b9560252b35fb1b9141f0347dc60001b615263565b60016002811115615881577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046002808111156158bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff16876080015167ffffffffffffffff1614156159e3576159107fc976c2ab07f819eb1ff0d91f1456232f0f8a9d507d55b476653116838e13890560001b615263565b61593c7f86de55610a7d3815bb8fff245ad54cc625fef4bed0fab30f9c508fb4a50cbc4b60001b615263565b6159687f02c0d9f1deb7fd8e68fae1969478ff96268a753cccca01c91e934b19198849d960001b615263565b8660000151604001516020015191506159a37f4e1c48e289262ef64bd11913bd6c12659f0c808c54fbae6481d6d178dae9a7d160001b615263565b6159cf7f4f37dfbc7cd9ae8832850072b856b0330b1e115467fc7b5b03b40d908bf4fbb960001b615263565b866000015160600151602001519250615a10565b615a0f7fa2b85daae4353828797a549d39a8047ac7857cb39da82383e043234afb27c11160001b615263565b5b615a3c7f5e72f4945ad24ebeab4dceece87d8cd75cd0e13947b8bb6d110ed99c35327e7c60001b615263565b615a687f8af625f36d170356c959797afb4dec22f0a646e83bad0a4b1e5a55ee23ad0e8c60001b615263565b615a7061a821565b615a9c7fffa61b8df511fcabdbee88b43d8123560dc979135c0e2e9d9291b7db119eed7560001b615263565b615ac87ff91d5c8ac4abf574dbe6eb9c92e5fea4e8fb99bb5f3941a8a4b202c9c6369aec60001b615263565b87602001518160000181905250615b017f58761027a71cc54978f233403ddd97069fba625761287048a934619d450ceb0560001b615263565b615b2d7f0f17724439fd6a3d37d2f78845b1f252df220e859eef2c6aa1d0fe8a9270043760001b615263565b84816020019067ffffffffffffffff16908167ffffffffffffffff1681525050615b797f8049406728d6ba878bb5f57b64a4d303c62355eaaa2b1faf0aed70fb2703b52460001b615263565b615ba57fb00aa28fcada7baa8e6bdfaef171c5fbf53fc9d5760ca1730cb01adbe872ba4860001b615263565b83816040019067ffffffffffffffff16908167ffffffffffffffff1681525050615bf17f98906e1d6b8badf3db0a4c8f1ff717822d00fb3685de1aefbc874325d1ec1b5d60001b615263565b615c1d7f36d77f8329d19ae9e8a6f9ceac41d02d013357da72c78ee898d099113e4f81ac60001b615263565b43816060018181525050615c537f3c0dca10417a49e1db2d90e935f950ecab1307ab3432496259ff5f44dfc1754560001b615263565b615c7f7f78cefbe59f25a5ac7ec055d6023ae15d9f04bef63772708ff42264537d1c192b60001b615263565b87604001518160800181905250615cb87f8a619254eccc99348b9926eb3423f10f0695d36bd2210024a4cbbb057e8119b560001b615263565b615ce47f6baa99e8f97f13d2e8e13f1e5f89b207574869cc34617a66aad27ca6b5db0e9d60001b615263565b87606001518160a00181905250615d1d7f4b3c1cedb0db6a60436e432d88f74771c2abd11938dd6b1b35e67c12c4f6573d60001b615263565b615d497f06797333a69563137a8a440f2629e04deb23081c1c4706171cda285aabaad74660001b615263565b6000615d54826113e6565b9050615d827fed45a069652ae32a797f53bcbb1f8481cd97212c842b8b43118abaae3ec61bdc60001b615263565b615dae7fdfcfa3a4a632f71cc873d12c105078a0401af92eef1fd39e6aa94a57fe85536160001b615263565b8060400151615eb357615de37f6a30eeb46e823cd1e991e398688fddd38b0f799d13a1b823064a8478be5bed4f60001b615263565b615e0f7f31cbba705e4f4266df137b01f7822ad2456193f365534b0e55126b5c936e5e3960001b615263565b615e3b7f1f6594cb9f51587d8ac62763f6ed834293cbd265bc2f54b1567bc6b10fd325ab60001b615263565b6000876080019015159081151581525050615e787f699d7980168a982d77908be541cad0834ede25ee0a6d833c7560b0fda9bf86d560001b615263565b615ea47f21ed9a55ef2f25d5a3e9298b9c0965e9caffde458c17fe17a5d32789b2b211f460001b615263565b869750505050505050506165e2565b615edf7f0cd70bc77af2a7db1864a410a48a41651f5b02a823fc753479cb0dd7c9a9f1ac60001b615263565b615f0b7f1a49a53c21329180a435fac5409a4930b36b6a4205984c1eba3cd85fe5eb17c660001b615263565b615f377fc03bdd134a349174c292fda2f17856888afb6c66ffc86511f52f0be307c86c8060001b615263565b615f3f61a877565b615f6b7f5fa3816dae3e0dc09728eeb70b2eaaa0e22e95a8f3a0494a9a617de4d8f772da60001b615263565b615f977f0d7309d462bc728536ef1d21f9daa0b11f8be9bb3a4b7d2ae44444dc50318c1d60001b615263565b81600001518160000181905250615fd07fd78dee8f6150c5430f3e3fa2cf3d5cd1faf0320ae84f77164393a67e4e93a46f60001b615263565b615ffc7f8aaeaa27869f33f5f87aef1830b1b2c9d5ef6219345ac20854514aa95a5d191360001b615263565b84816020019067ffffffffffffffff16908167ffffffffffffffff16815250506160487fb459d712c6133188972c54b380d9ad09dd4267acbf29240f4ee23add12b2a9eb60001b615263565b6160747f4b7973f26d13ff88bbe579a0971a457b800079f4fdd25069bf7725c083e0e13660001b615263565b83816040019067ffffffffffffffff16908167ffffffffffffffff16815250506160c07f39dcef896f8f94c8fe385adc487e3e4c6947517c3ac7eae1fcf335046daa251e60001b615263565b6160ec7ff94dce85c1e7817f7cacc96200fbd4bfcc48a8a67f4b4b33a13b39bb3064631460001b615263565b438160600181815250506161227f95425faa1bf69aa5220f00070f396e3212cfb24a618f731c54e272cce0ea15cf60001b615263565b61614e7fa84c044d28042fe1a48167b7a0d174e95c863cc4d981d48ffdb9650b393dffb660001b615263565b896040015181608001819052506161867e1feacad9828d757baabefa1393488922da800f33218e3c3e772fb2faca013760001b615263565b6161b27f433612491a6557f25ee3b416afdf5fb697094df8cb45d8b2ae24c4746825329460001b615263565b60006161bd82617bb4565b90506161eb7f27b27c7ae125297548095624ce1f081669957c60ac5a7064b7619eee54974fc160001b615263565b6162177fe8525e87a2f12dbaf660a9673d9bbc53c3a01de36a694a2d9ee71e9957a91dba60001b615263565b806040015161631e5761624c7f3b5c1e58410e838f183e11dec711ececd4ef976ecea7b0a5219fc62b84d5768360001b615263565b6162787fe20b3946764ec847b74197fa44e711457de1087e35507c5abac200036b64d5d660001b615263565b6162a47f55f8e639471796980bf01a60aab952e2cd7ba5b124d8d7d29b6557c401e391d360001b615263565b60008960800190151590811515815250506162e17faa322728401663eda1ed0d771b018a29f8a8af062508efa5cc2e27a5e4d8901660001b615263565b61630d7fdb0b35a297f18aa2c8d4f84ba25964981ea05540925174e26c68c6e6708dbd6b60001b615263565b8899505050505050505050506165e2565b61634a7f6052654a45ac3dfb455b1b2047c6bab8ba53509561c95b280fbadf421e56f33a60001b615263565b6163767fd64b3b969c84b9c14eb2c5198b93d79d487afb8bbf47d168a121faf3b34ca1ec60001b615263565b6163a27f4bccb0af4547b341a98b3839304532b735db52f64bb6f5479cc531b559776aaa60001b615263565b806000015189600001819052506163db7f7a3dcba130c077e1f10d8e57c1811074d25f0a2e16874341cb15327837ac08f960001b615263565b6164077fd300ea826449ece48df5899c227024d10eb8ce096284fc3ddb712a3c290db3ce60001b615263565b826000015160400151896020019067ffffffffffffffff16908167ffffffffffffffff168152505061645b7fe91ab118dd0ff0c5745f7c62850074507559d2019e7e8803f59a854af57556df60001b615263565b6164877fc5bde7242fcac7bfc18a53226c96c1894e9978cdd17641618ad35f99eb3464cb60001b615263565b806000015160400151896040019067ffffffffffffffff16908167ffffffffffffffff16815250506164db7fa551efe8c3c2f80aa82e989fb4e63a37e23225b18396c482e93449dff090e24a60001b615263565b6165077fea91de0d1e62fdba3127ac23b11b4ca95f3d7a16ea87a6d2cf0c77624e10655760001b615263565b826020015189606001819052506165407f119ba8ff3113a94dd243bc4f551885997f372a053b188852bce0f16f8931791960001b615263565b61656c7f6767caa642b42ec2b851d0a938dd90829c63ef7d0214957f4cc4c7f7d6a1bc6160001b615263565b60018960800190151590811515815250506165a97f973e5523cd42bbb078fc30d415618b734a8003c5e431850a0aa575451c435a3060001b615263565b6165d57f0aeff06b9d99f1390758db8e7a2c12da1bac32a54831d91157f6733c4ad5f87960001b615263565b8899505050505050505050505b919050565b6166137fb8a39981396734afbe7ff4978c383fba4009972088dd62617570fd4bd21b8f8d60001b615263565b61663f7fb19257a208e02863492353a2fb66f6c1d6b792075dcf1262b32193358e7f3b2560001b615263565b61666b7f66353fe9f7f201bd7ff4304bcd5ddead785ea8a7f94c8d88fdc38c91d595ec6d60001b615263565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156166d557600080fd5b505afa1580156166e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061670d919061bc08565b905061673b7f7d0dc31b907921b0ffd1fc303c92ca33823d4412160470a0f7f1cece1d3eb0d160001b615263565b6167677fb87c7e2f283090d7148a00acaacf9fb8d10e78e46f2c90fb51391d9ee37a58ae60001b615263565b4383606001518260c0015167ffffffffffffffff16616786919061cf1f565b1115616816576167b87f31ef80569971d5c4a327db2d2781df87460d9f4a028f2d40752696597b4434d560001b615263565b6167e47fa022d00445d7186335adf84c3e624e17752f696f3e0077bfc1e742c5be51a2cc60001b615263565b6040517f6a7c109700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6168427f4876e3ba627595e693f9ac14bf61d3f90d7490cdfdeb221fb8e0de1d0a7ad6a860001b615263565b61686e7f6ee465a5df2651aedf5add7793353a6a9b70e8f16e1f9af27187ed2f7c45bcd460001b615263565b61689a7ff550644653fb4761b066a0b1d5cfeadff5f83fd6667749c8144abd1c8d582fea60001b615263565b60606168c87faec90c7e0560bd0b4d969fda364d71671127cb6d0205149a7edebd54ebb71a8e60001b615263565b6168f47f70ccce8f3fc8c66cff5856b8b528834c2a5406dbcfce21d1384551fc1189c71560001b615263565b60006169227f9b06af8f1dfe5b400f705b52c6e7d7cc6a0cf11a726bea8e4867c1e17a8607e060001b615263565b61694e7fbc7a606fb6eb97677fc3b17e80e2816e80f73f371034c065fe8d17ee3fd5ad5a60001b615263565b600061697c7f0dbb89c1b065f4dad6553d567bbcb3424fd3251781b6b984098c70075aa8a67a60001b615263565b6169a77e9a2a284a1e5747436983ccc28146b4a21ed00fcd38a678161fbac6467808bc60001b615263565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e63876040518263ffffffff1660e01b8152600401616a04919061c9fd565b60006040518083038186803b158015616a1c57600080fd5b505afa158015616a30573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616a59919061b9fb565b9050616a877f8a5bde4d8a7cafa6b13de304724fe439a3906bd41245a765d8fc7a93812479d560001b615263565b616ab37f4a0fbd8a9bcaab761c19c972c110fed4a89b74374ec9b352ffc133242cab3e3760001b615263565b6000600167ffffffffffffffff811115616af6577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015616b245781602001602082028036833780820191505090505b509050616b537f4c9d5798afd0db15103403cd2e1d2af5f6656b2b43582c9f60d30b1a956db0a460001b615263565b616b7f7f87157521ac4ccf566e59a357aaed64f8fbf83de2fa6afbdf44a2399f9a31f54260001b615263565b600081600081518110616bbb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101906001811115616bfb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b90816001811115616c35577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050616c657faf345a184796e62996f6f044b483caf89b6c447b4a5981142cf507d217499e7460001b615263565b616c917f7aa198d0297f88dadcc1e97983a2f0e4f6ddffecf22a1676752cc777783586ba60001b615263565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad525a98389846040518463ffffffff1660e01b8152600401616cf09392919061ca48565b600060405180830381600087803b158015616d0a57600080fd5b505af1158015616d1e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616d47919061ba3c565b809550819650829750505050616d7f7fa84bb976cb20c82b570c642aa89218b8dab90248201dcbc588e7e29d67da7c8e60001b615263565b616dab7f2f7390cc6d93e159c12d8acaf6bc79be9bbbad3de51a2752cbc72965ecfe4ff660001b615263565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166328a8565c896040518263ffffffff1660e01b8152600401616e08919061c9fd565b60006040518083038186803b158015616e2057600080fd5b505afa158015616e34573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616e5d919061b9fb565b9050616e8b7f6ff922caa227872d85734ebaa84106b47daa3196b905f266ac33e6c0ccc1c7ca60001b615263565b616eb77f1608f79ebd5fa535c1c784a53e5e3b95614d0935cdcbdda73cb161586595e63060001b615263565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad525a9828a856040518463ffffffff1660e01b8152600401616f169392919061ca48565b600060405180830381600087803b158015616f3057600080fd5b505af1158015616f44573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616f6d919061ba3c565b809650819750829850505050616fa57f337e7c2d2b9e721eb830e74906b8d56ffa0ff108062001c132a78f6d0622180e60001b615263565b616fd17f66ecc05af294c8d72562502e2637d074c0c416d4a2af5ee400e89130dc93d5be60001b615263565b60005b8651811015617225576170097fd5011c428c9982272e08aafdcba700007ba1530a9b699f65b00dba940ae72ed860001b615263565b6170357f59e17a3616f6799aa8f880916bb897d0873e721e56277bb6f92d6e243180c96560001b615263565b60005b82518110156172115761706d7faba56ab14940468ab017b54878d67af602c208661c5deb7674de6f0cc847bf8960001b615263565b6170997f3e24e25c3bb1d4323bdc81d0e46303d06161f1f29f2d2c90ebc05b1d515b3c1460001b615263565b8281815181106170d2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518051906020012088838151811061711a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518051906020012014156171d15761715b7fda379f86bf0ab4595104a260da0c7ed535665f6a9201c6dd879f984d10d2352060001b615263565b6171877fb32619df48bb77644cbf30d3f27e901b620087c6b70f72ce77806b9c3befce8160001b615263565b8281815181106171c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010160608152506171fe565b6171fd7ffa4de861231c96435579666ef0837300745e047d542de19de35f7cc81c405a5360001b615263565b5b80806172099061d1d0565b915050617038565b50808061721d9061d1d0565b915050616fd4565b506172527f90abe2811a5ee0bb8420f677537fc74d2ee6600ba5c98486884a1e049ca475bd60001b615263565b61727e7f08a95bacff9d19f180a23ba73c72b3281cb6306c03fd786f7691603a8a380f4b60001b615263565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70e627289836040518363ffffffff1660e01b81526004016172db92919061ca18565b600060405180830381600087803b1580156172f557600080fd5b505af1158015617309573d6000803e3d6000fd5b50505050505050505050505050565b6173b28383836040516024016173309392919061cb72565b6040516020818303038152906040527f969cdd03000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050619731565b505050565b60006173c23061975a565b15905090565b6173d061a9bc565b6173fc7fcd77597e8dab75781cb46341edd4f5e0a9987acca840352c8ac0947e54bc0c9460001b615263565b6174287f4a9c44639e1b17c097f6592d3f6378e1dd31dc79e1074ecdf0d93bf75cf06a9d60001b615263565b6174547fbaddc102603540516c47a65d70c1c4a4d1add2940abbbd5a4a22958b14f43e0260001b615263565b61745c61a9fb565b6174887fa1aa33e19108c6fa2024a252196d8a565c72a7c61a1ce0050774169210f3aca360001b615263565b6174b47fdee8241f2812f1bc03d3012d6661ba32ef1d187fa8c57e19fe0fe3e3585bb1d160001b615263565b846020015185600001516174c8919061cf75565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506175137f2e82e05e1e06f56fe6b90bb25819d7e41ecea46e2aff89134373259391c61b3060001b615263565b61753f7f20a304c7d8fdf440fd97ff7e0870a1e6adca4f15128926f9f4aaa25149c23d0260001b615263565b8360c00151816040019067ffffffffffffffff16908167ffffffffffffffff168152505061758f7f567d87f5ad592a1d9a2ca4aff1da95f66a39f49c8342f41d3b1dcaffb6576c7f60001b615263565b6175bb7ff57f967f952437b7ef51b54c44d70181104e1e49fda55ff7db152d9680f2b63660001b615263565b84606001518160800181815250506175f57f8b37b0909c582d9ccc9b67ef35fffbfaebed7628b774c997fb6965eb6ca1b89260001b615263565b6176217f54a12d6ccfe5a07611d1488d4c1329759a71b4ab4574cc595546083e9953469d60001b615263565b8361010001518160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506176727f7e178e46810e3b6fdcbe6618e3746ae37c563e015767ad68d7b132017581893c60001b615263565b61769e7ffac9af5085858472e6fdd6840e6a5b3b51a89c029af82780a51e065afa8fc6c060001b615263565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663178e4dc08387876040518463ffffffff1660e01b81526004016176ff9392919061cc95565b60606040518083038186803b15801561771757600080fd5b505afa15801561772b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061774f919061bc32565b905061777d7f58af69b2c3a1b45656e71b7bb70e29cfb53208713c33bf119dd0255c3424321360001b615263565b6177a97f43baaab0e358f821b68b7ad0b4d5c9f7d57c8d153102e500c1a4c28d0756161560001b615263565b80925050509392505050565b60006177e37f1a1bb5fd7a3bf76a23df585b48e3df032f4fb0951878ddee2fec92c6581f350460001b615263565b61780f7f76fb9a6e55d104e67a6881382e035a4f63630cd180eb02c93e911350fc9eb94460001b615263565b61783b7f5eb041a50a658eb03552c79dc2daa8419add94ccf3cd6f54599101b904c29f4a60001b615263565b600061784a836040015161976d565b90506178787f63e5c53703e0579b6f844105d867f5c24362fe310cd8928b964f07715db71f6760001b615263565b6178a47f863381359afad8c7f305bf8f4a40342baaa3317304240dc9cbbc4b6355385a4460001b615263565b80617933576178d57f202f7ff32193c58730f11011bfa885e88bd92c6d1e972326047a6d6b9675d05660001b615263565b6179017f6aca6571b1b8b7138cab90808e60486b5e79919c064d5e8f7f0e4cefde27be8560001b615263565b6040517ff75f9e9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61795f7f2bdcf29d2680e5f9febe7b0f9119329ed3823799d1dc3d8181d31ff0cac03d0560001b615263565b61798b7f417cfdf1fa51c3ff1a0d7b65544da51c5a0d723d59ab32f5b26f2578f34c5fdf60001b615263565b6179b77f1347b9b063a611182836dcc9d9faeedd0a4ec043082de6411e58d3512d5cb67060001b615263565b60006179c6846060015161976d565b90506179f47ff2537cfd0a34e92bf52338fe060be78beee8e927c55f72d5aa2579f75e2a706160001b615263565b617a207fa7c6c73ffb52ebbd55f1f69ecb828ff59a808420fde13f15e720e5b69b9c80d160001b615263565b80617aaf57617a517f5530ce426ce9e5555943b0029ff3d32a73fd20f409e1750ddaeec50a4b65160b60001b615263565b617a7d7fb5343537b919d6409f4246854d524d56b292cd119f7469d8f16ce27a82bb811e60001b615263565b6040517ff75f9e9f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b617adb7fba2e828406827562f30f57641e99fd76b461ff555d9dc65fec2a198307f9d22060001b615263565b617b077f92904401604f2d7304fa0b1835b735f405e732493099cf7eda511ac8fd530b6260001b615263565b617b337f45b525dfcbd5ddbbfa0da76a1eeb6f11bd21f2f8c135cdf78d8ce0c82747c72d60001b615263565b6000617b4f85604001516000015186606001516000015161a01d565b9050617b7d7ffceb2ab08d2a23624f0ac61b3db20c883631c9234a67d295cf93228da9853cb660001b615263565b617ba97f89b0d67cc86c3ddaaefe5c096d84133a893e5e7f4e22cf39b918e787642a02a660001b615263565b809350505050919050565b617bbc61aae3565b617be87f253e149e85128803e507db3c938700a7e5878550b31b418f7b56268d5432186160001b615263565b617c147f62cf000c1939ce537efa7ae5a93a3853a2653600a3eded5970661261cdfe9eff60001b615263565b617c407f57f3615f538aa83f52cea4fd871d662be2ed062630ca6c6fa2cbf25c45a9921860001b615263565b617c4861aae3565b617c747f0de28972420fe913f7261f553e9a28f6d5b52c4807f7fe1d093ee343aff3978360001b615263565b617ca07f6b359bb53148eafd718ee9237bbc6b0dbbf72c00a810c27c4d2d8f499e1875df60001b615263565b826020015167ffffffffffffffff1683600001516020015167ffffffffffffffff161015617dbe57617cf47fef14c563fbc95e976a03531b0e4bd5f3850721ee48607dabc0065ae935c9c2de60001b615263565b617d207f884427401ea127595b8cb4befcddd5e5e456e30970d5991bb1f9182195871a9760001b615263565b617d4c7ffb4e07def7f95e91c01db8adb0f43e05910c1b0bd76d716135359d8ad07345af60001b615263565b6000816040019015159081151581525050617d897f3fe29f13a10045abfe89b7dda509b44aeb25fff427d275048033a720abb31ba860001b615263565b617db57f2c00c7b409208de4fc0e767fcd130200f8f853173b2e2690d7eb311b5c59e89260001b615263565b80915050618867565b617dea7fe03db3e6418899f7c2588ab3849e05cc0891b2b0dfaed6d16b0ab4eda5a58dcc60001b615263565b617e167f092d66d97c1e94c3cce45a2039c978ace70437c243f0de643e2864f0c5a9909560001b615263565b617e427f82fcd01c382be2c5212e333c876357e6300d6c0bd9fea340f465ac033fe8e18560001b615263565b8260600151836040015167ffffffffffffffff16846000015160600151617e69919061cfb3565b1015617f6557617e9b7f138742c9d02c2a07f8508d592a471d40078d95aa19ec72b4c2af4c52069777e360001b615263565b617ec77f37ff692c43389fbd1839703b3ef2a02962cd5476ba9a5a11bf4f5f8d83f23fc460001b615263565b617ef37f65b34875464deb09b87ee7e4422552a239c322e9dbc41de7d6af30c80c377a4960001b615263565b6000816040019015159081151581525050617f307f45b048fddb8feed362f87c60626ed35c48ac614f88a94d05e3a3f8b212bf440d60001b615263565b617f5c7f1f7cf5b5212a34caa84c284b9225f23a0ea0b799a35fc4beefd73bf0149e1e4d60001b615263565b80915050618867565b617f917fb3467c8566ed83f0b77cc0c365a6aaefa676b9748d17500b2a45915f001f0ab860001b615263565b617fbd7fab14691358f5b40b65c5435aeac8843c0b487cb4baaa12b46c69af89fd0def8060001b615263565b617fe97f68859bc65271d8ae9b9f236b2444ad0ad8cdac3373f9bc3332e456dd19d1ee6960001b615263565b617ff161a775565b61801d7fb2c3891f6ea96ae0d4bd2fa904d92360b2c70d49e1282d715231faf67902c37d60001b615263565b6180497f06c50c33798c96dc3c1106734782ca7806d8df07ec31e0833ff35b0dd39cc83060001b615263565b836000015160000151816000019067ffffffffffffffff16908167ffffffffffffffff168152505061809d7f6acb838ff72242a8aae1ea1737520cbef8b4f80968263ed1dbdc3b42c26bd57160001b615263565b6180c97f25a71e2dc8a610a0db95ce6b99f379621ab9ad0379d8398a60b994204d77219260001b615263565b83602001518460000151602001516180e1919061cfe7565b816020019067ffffffffffffffff16908167ffffffffffffffff168152505061812c7fa1507a2ef5f9d431de261fe0e380fa4f26f83068f5db84d930d470458492655a60001b615263565b6181587fc6fa7efb98bde6ec18889ec1af692b338ef59b3c097379085f0b107bbb24ce7960001b615263565b836040015167ffffffffffffffff1684600001516060015161817a919061cfb3565b8160600181815250506181af7f1c1373b42cc356319dfd82b7f9648d6a70dbe5ee1623406fa284c62d591d628e60001b615263565b6181db7f2beda2a46fa25f560ac8ca3a37e7b4c50617f9876580d5cfbe0390cac2c8c8f860001b615263565b836000015160400151816040019067ffffffffffffffff16908167ffffffffffffffff168152505061822f7ffad0a22b5ca4c00706e24ccd4ef1850b1a2fc350b2dee8d8735ad731b316252360001b615263565b61825b7fc16e23c174b76bf306769dceb0b149e33a97326d851a78822bda45dba96f2d7a60001b615263565b600061827085600001518660800151436173c8565b905061829e7f817a6949b04a70670da510a2b35458352abc366cc9424bfb2ad90b8a9c5c77b460001b615263565b6182ca7f328af1419381c34228a384944a47fb2080c319d104fd9b73ec8d8770a5427f4160001b615263565b60006182db838760800151436173c8565b90506183097f3a54efe95277753a8f52fe306c62fd3b7de55e009a352f61e0608848a298008e60001b615263565b6183357f6e255bba49aacd0824339a4f6319855f286f6c8d219ff445a33f7ed837f21f8660001b615263565b6000826040015183602001518460000151618350919061cf75565b61835a919061cf75565b90506183887f6f2fc99e616b4f3f3320693e6b1cb234a021020216ef7108d98ced6b236e8c6260001b615263565b6183b47f33932deef29cc55e3bea276dc54af3f902b3a3566da7d5e5d4cebe4daff0eb9960001b615263565b60008260400151836020015184600001516183cf919061cf75565b6183d9919061cf75565b90506184077f92d890cb4088b7ddebc1fb8dac62e5e728d2dfb724277c44f51c856384803dee60001b615263565b6184337ffa34556be74f133edcd769942d95314bb48904bc54fdb1f02cafde1a2694318360001b615263565b8067ffffffffffffffff168267ffffffffffffffff16116185495761847a7ffa481ee48d7310f90a2713d8d32181cdcb7d1643208c22e099b24fdbe846f55e60001b615263565b6184a67fec6ca9f2afb120c968b7465b3136010c777c5874a3c0d587e730f7031d6c292560001b615263565b6184d27f80b68ec14e511a4c52b55c8c7cb3ecfc47cd604a1db9a86c8413e3b7e7cae2c060001b615263565b600086604001901515908115158152505061850f7f561708a4756edb076e00c0b7738fc23e482735ff6af300d2dd0844555a13865660001b615263565b61853b7f3bc5ac30dbe4fea9f822f435071d58ae25ffacb1780a06eecf92e4374aa2f76260001b615263565b859650505050505050618867565b6185757f4dd10983ed14bcf8055df647355183841114f840d83cebdbdea702dc1596ac4760001b615263565b6185a17f71585d9d883ef495903cfb1f569de84f2ba6bc930144fbf8fd9b696c4cdc33d160001b615263565b6185cd7fec82c71ab0eaf7a5a4d8246788878f4021d1c609e4d86edf6479175345c93f9560001b615263565b600081836185db919061cfe7565b90506186097f6974a05dcc9474fb9f105557dd638d69621f7d2a1969d24c2aec30a79e5eeaf060001b615263565b6186357f81e6053b4ba7a9e49ec374108f250efdfae764d9612359274e89b8ee125d6ef960001b615263565b8067ffffffffffffffff16866040015167ffffffffffffffff161015618751576186817f261aed24d570df09f968bd3931c8ecf005435b09a61afa1021a68dc921c9f77f60001b615263565b6186ad7feb9c44fc9ad62d30932888ee37254bcddf8c06c897a3786e1b2cceb6a1e3687f60001b615263565b6186d97ff192731edf34e9044e2c08860a17eb65512f93b3b1484157ce42366088cbcfae60001b615263565b60008760400190151590811515815250506187167f5c101cd1e715f556af660d34c145340cc8cbde5e69b31d0d913811dc2348a1a160001b615263565b6187427f8e86f912b1ecf993616921e69b93aa4bf2d9dede8674a740a9ff37d2985f6f2f60001b615263565b86975050505050505050618867565b61877d7f7ccf9242a70a32e4c43702c1fcc6b390d7c6821e921c6c2fc62f5180744fc5eb60001b615263565b6187a97fbcebf4c35c69d5d0de85bf62ff39f6034136ec6e58b7f0d484ee03e7e096547060001b615263565b6187d57f81933f720cc5fd073ae4b485130579110db81f6e807b51d44ebb865b9827ff4260001b615263565b80866040018181516187e7919061cfe7565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506188307ffe5133e169d5b6d0daeb43ddaf671c7c74e167b9ab9ebf89e282bb6b95bb0ef660001b615263565b61885c7fdb662ce683dad6d8d3814f3c9767a1790f285040cbb1038667274321012180dc60001b615263565b869750505050505050505b919050565b600061889a7fa34bb679eb73a65492301fd8c850ac7b95b4b3d3ca24ecc1fcf67410ab8d7e2060001b615263565b6188c67f3a925a56fb88392812c04ee32daf869298dce53e215d53635e52849efad9246a60001b615263565b6188f27f46573134885690c8d6ae5961c6f022ca1796959e8af9bee2c09d84f827f4549160001b615263565b600082604001516020015167ffffffffffffffff1611618999576189387f8feb54acad75578bf3b11ee9b89210bf9a802f2ab08c2c9a7e65a5d38f17ea0960001b615263565b6189647f3e6cc663bc7360b92e429ccfc493db638bc0201dcff7b45185048e328eab2c6060001b615263565b6189907fb6fa2ceb2b98feba55f1c21c7df148e654d89b2419eedecc5c489ef43ab3115560001b615263565b60009050618fbb565b6189c57f0d95ad5f269bd2ec10841a77bc5c7ebc24406350a5f863d24836dbf95e8205b960001b615263565b6189f17f49a26248ee6c4bac715f42a976bf7a883f54dea7cd73613897cfa39ebb89bcbb60001b615263565b618a1d7fb0f8cedd81167328d920f568bcf54bbe80f51ba2a89ff669f42b176f2240119460001b615263565b600082606001516020015167ffffffffffffffff1611618ac457618a637f6c8fc798b5119a665c1611f8a7f1da8fc7e2bc5c05c908d4bacfd950361acde360001b615263565b618a8f7feb870dafb60bd2622d625f40d1e64445853d891e5e4a8f1225a43ddac544e64660001b615263565b618abb7f40c1c7144cfba103801bc37e5e50bd4bab66726e48285f39b0bb9f880d14522560001b615263565b60009050618fbb565b618af07f160cb0b49cc36d51b7784797baffb35dc91d3f26e48b3d4edf57874cd4beaad960001b615263565b618b1c7f3908c4f5f4ef6929fb6683a9cf367eedc2b0426fcd447abfe022b1868be3bbb860001b615263565b618b487fab5783ae481266a737c60a6175b6e654a95a9b7a38f7fade92783f0209b0556c60001b615263565b6000618b53836177b5565b9050618b817f4d1ca87d3a4718c0960a687c437be6e375bc8114ae771e70318c72856e8be4bf60001b615263565b618bad7f341f2e9f048005d2d54fb3b891ac40c7dc7dffae41d59f9e6ec00494e52596bc60001b615263565b60006002811115618be7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b67ffffffffffffffff168167ffffffffffffffff161415618c9057618c2e7fd399018fb7d3327ff884be70cd31e43ddf28303775bc8033414a1fed8d96961a60001b615263565b618c5a7fe0d529363d7e41eeb6057ac6e64535d00c77d5844a578d6cb53d609fffc161f560001b615263565b618c867f64585f877f1be0454bb570bed9b4fe86576822aac56692aba65a9fab0541310460001b615263565b6000915050618fbb565b618cbc7f6a0e7265968f985cfee1f7490cc8f0a2fe09089ba98382a4031e736ac1d033db60001b615263565b618ce87faff6867647d1792712fb0b0dec35362b3061821ef75a52fdbc93f906fd71bdd360001b615263565b618d147f29440371c22995df0d0bbd009c9830b87ec8ce2b18fa4a797c5017e51b87bc8460001b615263565b6000618d1f8461a18a565b9050618d4d7fafa3e704829f3469f5dae9126de38981da37d9a03c538a77a49dfab2e3e7239960001b615263565b618d797f808fa16dff3d262e062f85f7f5f3526f6eddff9f6d811c61dd210ff88273a49960001b615263565b8015618f2f57618dab7fa6e75f95f3d9194235c9220ff826c6e0c2b88651ff1b97b6d4248b174073332360001b615263565b618dd77f8be5b7a4565770bf93adc6ad1162dc4b68b599550ef2a8b8bfb629d8c20d04dc60001b615263565b618e037f57db46ca992e89f2a112ab17785b103c1ed1e3765cf91b60bf63e0b1b19d0ef760001b615263565b6000618e0e8561a314565b9050618e3c7f9e19ab94a0eb57d4683135641d516914605c82a56fbe727a8181b7013f8cff8160001b615263565b618e687f657a17a2ca23b115323bba00992bb4c75ad1b7f9061a001415dd886ef96d44e060001b615263565b80618efd57618e997f8b8ca6dcb3a4f68702e9adc335938f6a41c0f4ed6761bd5a07a72f65d56be0fe60001b615263565b618ec57fd48998e38086f7f4fa018cd840339cff48442645e17a00682dbbe328a3d8aadb60001b615263565b618ef17fa3655ab87f28ed4d2884ffa3b64f6db06b2a613ea1dbd1f93b121fdd26e13a7e60001b615263565b60009350505050618fbb565b618f297fd786909c04cea7a41143ff74ee0440065065aed5d03ce0022d8fcbd59fdb788760001b615263565b50618f5c565b618f5b7fbc5795aaa6353696af3d9d578edff8fd465891959206c3ac2dbf0e1dd11a421f60001b615263565b5b618f887fc20fe9bbd073a13948e646ece432c764350ef3fac641bdd3be62d418e136c67660001b615263565b618fb47ff56bd32e1b4263816d63443e7c92dcbbe133cf7fb67d4ef53df830c9b201232c60001b615263565b6001925050505b919050565b618fc861a775565b618ff47f7fabca25c47e8a39b4c09bec05cd5d78034523dbdeae34c87e653b243c3295be60001b615263565b6190207f99579c27cc6350afb161c193c008a1ca81dae641c2af3e4360614c9e0eadf24d60001b615263565b61904c7f7601938560528fffd46a11dbdbbad0b60721b5848e4aa7288dfcb5cd68ed967460001b615263565b6000836000019067ffffffffffffffff16908167ffffffffffffffff16815250506190997ffe4a2e3b1d6eae0aaf15b4d6a4183f0ff6037836364e6a73bc54f47237cd0e8b60001b615263565b6190c57f699bd1bca704ba4a402cfe7b4616f8bc60aa601e54c26768683482b95b02288160001b615263565b6000836020019067ffffffffffffffff16908167ffffffffffffffff16815250506191127f55298209fa365d192c52aac03c7bbd6023676425cc04f2a0e3fa12a4293e59f560001b615263565b61913e7f53306b76457f40de0856c06b5632ad77bce333f68ccf624befd445e3f80b15be60001b615263565b818360600181815250506191747f6d9797fff253541768443417a93b60bde17b4189ace6a6120946fcf910d92c1460001b615263565b6191a07f62ef6208e9799e0a6cfbb9b07d111c6ee9a7df165d815b5e2ea5faf023a1bc2e60001b615263565b818360800181815250506191d67f40a30eb3cf58015c672cae6f19f36f60ef964601c136cc2e371c839615109a3760001b615263565b6192027f2b8ce184b14ed7ce929f8d6f7a31f53dab096f8a43e1a23bbcf456352bde9bc660001b615263565b82905092915050565b60006192397fb86598e475625025578a96710d8052af1071e540f0e82e49d81725240cf0302d60001b615263565b6192657f84fbd2cda41c19a017c4e49dc5efe77e8c4a4404d6ce1ef8f9393368b475390f60001b615263565b6192917fc5aa39b9b366e7e92d4ffa88ffe93f78f8553bba3d45435b2db0c551792ece3760001b615263565b600061929c836177b5565b90506192ca7fe3f1a58f26b07ebda561c4e14c7dd19d80e1604f265a03049af5f825086fe61760001b615263565b6192f67f0575db1a5a398b3561ac0636f57b490a25d8d4b47c7fb505ba028125165b008d60001b615263565b60016002811115619330577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111561936c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168167ffffffffffffffff161461941d576193bb7f4ce0ec63b8fd3a9ced23e578cf83aafb107998fc2acac8ac85d67527603b641f60001b615263565b6193e77ffadacfda31b55f76dc097da759d7bec2ec9fe1e31d2a333624086a090b051b9160001b615263565b6194137f8d4d7453ea289d9a6ff395deb1d5158e45ccda7d5e44b65d2d313a371607a60760001b615263565b600091505061972b565b6194497f405e6b6da1c94283d96e6ae0f76d89da537ee390174388f3d1ac6b8ac7f79f1b60001b615263565b6194757f3c44868b93523f03e2dbd995d09e90c3374bbad4c1500526af9ffc7e928b644560001b615263565b6194a17faa5423fa46969bd0a9855acca219dfad5849ad60ea83fb9f226ff040a309166560001b615263565b600083604001516020015167ffffffffffffffff1614806194d45750600083606001516020015167ffffffffffffffff16145b15619567576195057f1e1826bdef793a41ad2a439708b2775260a8a87172826108c57acb50a02ec58d60001b615263565b6195317ffc403f99af60876df68c114d84f9c15632db51ba66173b22fb6f117ff8c7e73a60001b615263565b61955d7fb92d712278763c5d81f757dbf1f2b07a029281a8a35e551cc25dec4de17d852260001b615263565b600091505061972b565b6195937f75b52ed5ef70b8cb51f7b6ad80eed2f7b99c0169fb6544fdf0c04251fc06d6d460001b615263565b6195bf7fcd9f696aa185eda8dcb1b780360ee03042d23ccd87b156fea6f49f8f0653955060001b615263565b6195eb7f74d8607b744f57d0317dfe1215277c765e1366f570c718f3035ef6b5bc525ced60001b615263565b8360c0015167ffffffffffffffff1683606001516020015167ffffffffffffffff1610156196a15761963f7f51730237df2f0f804c4b57cd02103ab8887e375300af395ec2e09dfe4a80978760001b615263565b61966b7f5e51326bcbe07fc05f3d8b510be3c121ac55d6b42d502dae1ee3a9539d266a1460001b615263565b6196977f2d407df40de95a0377843c24e08a23962e097c9900d810a9150ddb8ad9e8bf0160001b615263565b600091505061972b565b6196cd7fda5e5311240b4c8076760c842eea9c922d4287cd6d97b55f19b60513306c724260001b615263565b6196f97f9d4734f9322655ae20b165919e6620ad85acaf49fbddff3e0b490bb0a65228ac60001b615263565b6197257f30292176ed7b43acec2fa67e87128a73b46d8f1a8385189fac0c412dea0cb0f160001b615263565b60019150505b92915050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600080823b905060008111915050919050565b600061979b7f8357822e8bdb92e5a6c0f7d4bbd041facfd4f8718079f95e633a7b6501db1e6060001b615263565b6197c77f058afb760734e140a022c6cbd960ecd9ce0029b68d1e422e4982c5b51b24bd0960001b615263565b6197f37f01acec8ed25a65560bf7526a5129c2c7db3701b645bc3d86ee33f09f7a03f51260001b615263565b60028081111561982c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82600001516002811115619869577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415619a245761989b7fa9d5586eb99953aaf0e29a506ff321a7d3c6e85b540586c2b1c940e7ec434fea60001b615263565b6198c77f1f28a47e198a413a7b890bfa277a9a8ad3df9b08011913de543396a8ff7d83f160001b615263565b6198f37f2e82758de762f0ed598fb8a6f40287fd0f5f415a92d1c35e1fa330cd4a2db9cc60001b615263565b6000826020015167ffffffffffffffff161415619997576199367f10d27ce20d23e8b6cff53ea9d6a58fffafd56ec786ad8e5461c137065cf92d2b60001b615263565b6199627fbf2de73975d148155c1480c613a064fffa9ad57ee41da4ee722fc10bcbfbd23d60001b615263565b61998e7fa64f766e6b52f7da08af2f3cf7809bd6fdcfa3f3a77a78f8d93a65cb5896063b60001b615263565b6000905061a018565b6199c37f2d3d0d6678dae684d01ddb1b51d00955ca185e91153a313843dabde491ae302f60001b615263565b6199ef7f1fd7f7275cbfd14375f322386d222cc60a6fee232080e9563790dd090b76f6bb60001b615263565b619a1b7f832678841268e9e5505dedcb7d24232f01b6e79124f896a9a42af17c12bdd24c60001b615263565b6001905061a018565b619a507f2a5cfe52dfa15db9a2948cf300f0c2367a21315a16f2edd117d1f2b0618e53ba60001b615263565b619a7c7f698c0f07e4c1bf1f38a96e64dbeed6cc185a89de56cf41e5df7478139abc80e860001b615263565b619aa87fbd19ceb7eab4b33bd7aa8258c9e6852d43ac3c6ba554fe0083af752e75d4785f60001b615263565b60016002811115619ae2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82600001516002811115619b1f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415619cda57619b517f214f4d54440fe9ccde5aa1df1992effc36df5726756e3fef119b7a6058b1ca8660001b615263565b619b7d7fd0f140ca4825467963cd22026e51576d8b200e0675d96cb1c3bcf24a01b7638b60001b615263565b619ba97f082d83187e133a14e8dd34d837bfa2c135ffed816d31bd232d86b3f4e7dbebfe60001b615263565b6000826020015167ffffffffffffffff161415619c4d57619bec7f3db80137bb017e7d6294c5542c5ba1212f8baa3bee4c17c0a3cc41e2c12c8e5060001b615263565b619c187f6087e3ee21a69e0d886cd6e9392548411f824aea516444f364f78b59e48b9cf460001b615263565b619c447f9b56b07f5b1e17f6a1a851c0aa3429654fe0d07d0df3c25a2ab30bcf3f2e95f760001b615263565b6000905061a018565b619c797f5e3dc5e813c47b7e5290a54bc8b5acb97a8ad9e5daddecf81ccf14847a9b31ca60001b615263565b619ca57f0f1c10ed41ed42004d6c508ea004fa13b53a2b01c9d3c101a725052c0cd6d73f60001b615263565b619cd17fb5a2b519bd6e35647196388c869314a8160309c4fe1e0fcde2027d95386cdefb60001b615263565b6001905061a018565b619d067ff81e51885fa177107e9f71ad5546d3b8ea0faea83561822755194cc1b71e0a6460001b615263565b619d327f67c4d8898e873578e1482830311474a8d2680252dc77a4e852ae4e5a97f6e20960001b615263565b619d5e7f5fc2027095f63c5e0ca23045ef9cd596fa82af4804ae444343b0b11ee783946760001b615263565b60006002811115619d98577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82600001516002811115619dd5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415619f8f57619e077f2a4c541dd9da0e40e19275133926ec7bcd95786a3e3e2cee2320b3446eff125d60001b615263565b619e337f17a2ba1ef2a22d1fd0708f86755533b58edf4c4e4010d14e528513455b2367ae60001b615263565b619e5f7f227777fc2135e98b6925375bdfa4354cfc854171c117ca2c7030a55d2286bc5f60001b615263565b6000826020015167ffffffffffffffff1614619f0257619ea17ff837c7fb3382f38bb8b6baaf9f6aaf63bc0f324a96f4dc3e9e84997170bc40ab60001b615263565b619ecd7ffa8291eee8197895057bfe57a7a583c64dea7dfeeb16f456430375b0a6dc71de60001b615263565b619ef97fc75d72846d157a8c1be13769fb9361f00b593416fea91898af00e4ee7cbd8e7c60001b615263565b6000905061a018565b619f2e7fa974695399088f6d4b4505ba63a38893f9dc9f14e56109d6c7af193fdcff53a560001b615263565b619f5a7f34da632a92f203ec2746dc630c28ec8f471f09f0abfb3eba741fd9eca488bf7d60001b615263565b619f867f078710240f15d0002258d77fe9439fc2de0be023a4f2a25acbeb51450cccba7160001b615263565b6001905061a018565b619fbb7fd75be6b49f0f6d9380caa66f3c45cf94a2376b0d27c52e9a3aa4353270adf40560001b615263565b619fe77f318add4717c40c47c13613ce6f95f6bcf84fa28cc41e2481bedfb19eb96fc5c760001b615263565b61a0137f5cde360363179e19c1a96a2121956b4dfc983c99ff62c27013823bff8a69e79b60001b615263565b600090505b919050565b600061a04b7f045ebf1823fac1ba60d3fec12a6b806439ae5bd50c2c5eee7fe8dc145160035660001b615263565b61a0777f077c0e134d411c29078dc0a5b0b39cf68f4496ef2e7164353c55331b0a67ed0560001b615263565b61a0a37fb7a2aab74613df5a3e34df0bfc1727538b36a10d33da1a54916cbed1f9f52bbb60001b615263565b600082600281111561a0de577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600485600281111561a119577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b67ffffffffffffffff16901b17905061a1547f4864e1c757f2c840f023ef2bd7a3b448e6b422c63659501310295adc8e75d55160001b615263565b61a1807feea6d4aba5136de03828016afbb70d9db2ae965900eab85fe6b50d65c243721660001b615263565b8091505092915050565b600061a1b87f591480b05bf9c92b173af303986cbdc80f62ffa4f51c18f41ceab37b5953b7e060001b615263565b61a1e47fc146e88cfd3bb0c7f2cbdfa3ce9d3d10b367d74096558cf6f7880e7ff7f6814060001b615263565b61a2107fc9e3c5352a55aea5c9ba2300fa0e9e41f75a80d7b8f12dbe4df37f5f222aee2b60001b615263565b60028081111561a249577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826040015160000151600281111561a28a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b148061a30d575060028081111561a2ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826060015160000151600281111561a30b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b145b9050919050565b600061a3427f256603139ba9722de7ee404f351711a972b4926142fbd82bb877288ca6e0dc9860001b615263565b61a36e7ff47193ca2908e84f681d97ce5189f822c62b87db666f8c431282a68bcf7a1fc760001b615263565b61a39a7f323ea0fa917ca980f3bdc5a750aedec4fee6ad28d123d2ee5e3c6c6660171c8d60001b615263565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e6384602001516040518263ffffffff1660e01b815260040161a3fb919061c9fd565b60006040518083038186803b15801561a41357600080fd5b505afa15801561a427573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061a450919061b9fb565b905061a47e7f28bfc6ba72bfeff75797a7b6edbc36542eb9cd54805e5d6f7d015b5f70c6d4dc60001b615263565b61a4aa7fe3c8833bce819b3ca6d49f7f0e7bb5a5486426f140cfe68407092e8ede234f4560001b615263565b60008151111561a5425761a4e07faaab24b90359f7d92adf395900393e0bc91b31504906a10bc1025e1df9f2d73360001b615263565b61a50c7f42be1dac56ff18cdf845e6e1bc90a24acf7ed0ee75d97dfabe5eba7d37d8a3ac60001b615263565b61a5387fd45c45c014809d07d6d6d56f82106b99a2cb2e3b05931175ebbe9035f9adfd8f60001b615263565b600091505061a719565b61a56e7f91ea8d01af46a920284a7e933cea725e5c85926502b05d1f88876aa153b9000260001b615263565b61a59a7f5120d4c5834d88010d5ad6e9c83ccbd4e18da051a4df2615b85728d6bd7693a960001b615263565b61a5c67fcb4a88d101b649515a745b1a58c010910281e5fc971e83f99895ab90027dde7060001b615263565b826020015173ffffffffffffffffffffffffffffffffffffffff16836000015173ffffffffffffffffffffffffffffffffffffffff161461a68f5761a62d7f630cde4583a12a854c89f32a444125641c2f0168c0eb3ea4f72ca6b66d4e58f160001b615263565b61a6597fb6374c23d95773c3c4f569b054da8c8fdc3c84183b53e4e3dfad10fefed769bd60001b615263565b61a6857f03344db61690b32418fb8051f21f93b99aa6544208b36f8221e863bb2c7393a860001b615263565b600091505061a719565b61a6bb7fc3002452268844c40a4f88b6cf4b556dec21af7bf831560d9c676348d60d135160001b615263565b61a6e77f25645f2e5cdfcc7a70969fc50eca544ca421cf76f6c82392758863796a7da75760001b615263565b61a7137f4cb72c2634d1408aa3e07e0a1c34b5a5283a0bd4e2db0b62520976586176d07c60001b615263565b60019150505b919050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600081525090565b604051806060016040528061a7d561a775565b8152602001606081526020016000151581525090565b604051806080016040528061a7fe61a775565b815260200161a80b61a71e565b8152602001606081526020016000151581525090565b6040518060c0016040528061a83461a775565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000815260200161a86a61ab16565b8152602001606081525090565b6040518060a0016040528061a88a61a775565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000815260200161a8c061ab16565b81525090565b6040518060a0016040528061a8d961abde565b815260200161a8e661a775565b815260200161a8f361ab16565b815260200160608152602001600067ffffffffffffffff1681525090565b604051806060016040528061a92461a775565b815260200161a93161a71e565b8152602001606081525090565b604051806060016040528061a95161abde565b815260200161a95e61a775565b815260200161a96b61ab16565b81525090565b6040518060a0016040528061a98461a775565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016000151581525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b604051806101e0016040528060608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000151581526020016060815260200160001515815260200160001515815260200160608152602001606081526020016000151581526020016000600181111561aadd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090565b604051806060016040528061aaf661a775565b8152602001600067ffffffffffffffff1681526020016000151581525090565b604051806101600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200161ac2b61ac3e565b815260200161ac3861ac3e565b81525090565b60405180604001604052806000600281111561ac83577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff1681525090565b600061acad61aca88461cd15565b61ccf0565b9050808382526020820190508285602086028201111561accc57600080fd5b60005b8581101561acfc578161ace2888261af29565b84526020840193506020830192505060018101905061accf565b5050509392505050565b600061ad1961ad148461cd41565b61ccf0565b9050808382526020820190508285602086028201111561ad3857600080fd5b60005b8581101561ad8257813567ffffffffffffffff81111561ad5a57600080fd5b80860161ad67898261b010565b8552602085019450602084019350505060018101905061ad3b565b5050509392505050565b600061ad9f61ad9a8461cd41565b61ccf0565b9050808382526020820190508285602086028201111561adbe57600080fd5b60005b8581101561ae0857815167ffffffffffffffff81111561ade057600080fd5b80860161aded898261b03a565b8552602085019450602084019350505060018101905061adc1565b5050509392505050565b600061ae2561ae208461cd6d565b61ccf0565b9050808382526020820190508285602086028201111561ae4457600080fd5b60005b8581101561ae8e57815167ffffffffffffffff81111561ae6657600080fd5b80860161ae73898261b184565b8552602085019450602084019350505060018101905061ae47565b5050509392505050565b600061aeab61aea68461cd99565b61ccf0565b90508281526020810184848401111561aec357600080fd5b61aece84828561d15d565b509392505050565b600061aee961aee48461cd99565b61ccf0565b90508281526020810184848401111561af0157600080fd5b61af0c84828561d16c565b509392505050565b60008135905061af238161d356565b92915050565b60008151905061af388161d356565b92915050565b600082601f83011261af4f57600080fd5b815161af5f84826020860161ac9a565b91505092915050565b600082601f83011261af7957600080fd5b813561af8984826020860161ad06565b91505092915050565b600082601f83011261afa357600080fd5b815161afb384826020860161ad8c565b91505092915050565b600082601f83011261afcd57600080fd5b815161afdd84826020860161ae12565b91505092915050565b60008151905061aff58161d36d565b92915050565b60008135905061b00a8161d384565b92915050565b600082601f83011261b02157600080fd5b813561b03184826020860161ae98565b91505092915050565b600082601f83011261b04b57600080fd5b815161b05b84826020860161aed6565b91505092915050565b60008135905061b0738161d39b565b92915050565b60008135905061b0888161d3b2565b92915050565b60008151905061b09d8161d3c9565b92915050565b60008151905061b0b28161d3d9565b92915050565b60008135905061b0c78161d3e9565b92915050565b6000610280828403121561b0e057600080fd5b61b0ea60c061ccf0565b9050600061b0fa8482850161b8ba565b60008301525060a061b10e8482850161b96c565b60208301525060c061b1228482850161b96c565b60408301525060e061b1368482850161b942565b60608301525061010061b14b8482850161b58a565b60808301525061026082013567ffffffffffffffff81111561b16c57600080fd5b61b1788482850161af68565b60a08301525092915050565b6000610320828403121561b19757600080fd5b61b1a26102e061ccf0565b9050600082015167ffffffffffffffff81111561b1be57600080fd5b61b1ca8482850161b03a565b600083015250602061b1de8482850161af29565b602083015250604082015167ffffffffffffffff81111561b1fe57600080fd5b61b20a8482850161b03a565b604083015250606061b21e8482850161b981565b606083015250608061b2328482850161b981565b60808301525060a061b2468482850161b981565b60a08301525060c061b25a8482850161b981565b60c08301525060e061b26e8482850161b981565b60e08301525061010061b2838482850161b957565b6101008301525061012061b2998482850161b981565b6101208301525061014061b2af8482850161b981565b6101408301525061016082015167ffffffffffffffff81111561b2d157600080fd5b61b2dd8482850161b03a565b6101608301525061018061b2f38482850161b981565b610180830152506101a061b3098482850161b957565b6101a0830152506101c061b31f8482850161afe6565b6101c0830152506101e061b3358482850161b0a3565b6101e08301525061020061b34b8482850161b981565b6102008301525061022082015167ffffffffffffffff81111561b36d57600080fd5b61b3798482850161af3e565b6102208301525061024082015167ffffffffffffffff81111561b39b57600080fd5b61b3a78482850161af3e565b6102408301525061026082015167ffffffffffffffff81111561b3c957600080fd5b61b3d58482850161b03a565b6102608301525061028061b3eb8482850161b08e565b610280830152506102a061b4018482850161afe6565b6102a0830152506102c061b4178482850161b424565b6102c08301525092915050565b60006060828403121561b43657600080fd5b61b440606061ccf0565b9050600061b4508482850161b981565b600083015250602061b4648482850161b981565b602083015250604061b4788482850161b981565b60408301525092915050565b60006102c0828403121561b49757600080fd5b61b4a1606061ccf0565b9050600061b4b18482850161b846565b60008301525060c061b4c58482850161b8ba565b60208301525061016061b4da8482850161b58a565b60408301525092915050565b6000610300828403121561b4f957600080fd5b61b50360a061ccf0565b9050600061b5138482850161b846565b60008301525060c061b5278482850161b8ba565b60208301525061016061b53c8482850161b58a565b6040830152506102c082013567ffffffffffffffff81111561b55d57600080fd5b61b5698482850161af68565b6060830152506102e061b57e8482850161b96c565b60808301525092915050565b6000610160828403121561b59d57600080fd5b61b5a861016061ccf0565b9050600061b5b88482850161b96c565b600083015250602061b5cc8482850161b96c565b602083015250604061b5e08482850161b96c565b604083015250606061b5f48482850161b96c565b606083015250608061b6088482850161b96c565b60808301525060a061b61c8482850161b96c565b60a08301525060c061b6308482850161b96c565b60c08301525060e061b6448482850161b96c565b60e08301525061010061b6598482850161b96c565b6101008301525061012061b66f8482850161b96c565b6101208301525061014061b6858482850161b96c565b6101408301525092915050565b6000610160828403121561b6a557600080fd5b61b6b061016061ccf0565b9050600061b6c08482850161b981565b600083015250602061b6d48482850161b981565b602083015250604061b6e88482850161b981565b604083015250606061b6fc8482850161b981565b606083015250608061b7108482850161b981565b60808301525060a061b7248482850161b981565b60a08301525060c061b7388482850161b981565b60c08301525060e061b74c8482850161b981565b60e08301525061010061b7618482850161b981565b6101008301525061012061b7778482850161b981565b6101208301525061014061b78d8482850161b981565b6101408301525092915050565b60006060828403121561b7ac57600080fd5b61b7b6606061ccf0565b9050600061b7c68482850161b981565b600083015250602061b7da8482850161b981565b602083015250604061b7ee8482850161b981565b60408301525092915050565b60006040828403121561b80c57600080fd5b61b816604061ccf0565b9050600061b8268482850161b0b8565b600083015250602061b83a8482850161b96c565b60208301525092915050565b600060c0828403121561b85857600080fd5b61b862608061ccf0565b9050600061b8728482850161af14565b600083015250602061b8868482850161af14565b602083015250604061b89a8482850161b7fa565b604083015250608061b8ae8482850161b7fa565b60608301525092915050565b600060a0828403121561b8cc57600080fd5b61b8d660a061ccf0565b9050600061b8e68482850161b96c565b600083015250602061b8fa8482850161b96c565b602083015250604061b90e8482850161b96c565b604083015250606061b9228482850161b942565b606083015250608061b9368482850161b942565b60808301525092915050565b60008135905061b9518161d3f9565b92915050565b60008151905061b9668161d3f9565b92915050565b60008135905061b97b8161d410565b92915050565b60008151905061b9908161d410565b92915050565b60006020828403121561b9a857600080fd5b600061b9b68482850161af14565b91505092915050565b60008060c0838503121561b9d257600080fd5b600061b9e08582860161af14565b925050602061b9f18582860161b8ba565b9150509250929050565b60006020828403121561ba0d57600080fd5b600082015167ffffffffffffffff81111561ba2757600080fd5b61ba338482850161af92565b91505092915050565b60008060006060848603121561ba5157600080fd5b600084015167ffffffffffffffff81111561ba6b57600080fd5b61ba778682870161af92565b935050602061ba888682870161b981565b925050604061ba998682870161afe6565b9150509250925092565b6000806040838503121561bab657600080fd5b600083015167ffffffffffffffff81111561bad057600080fd5b61badc8582860161afbc565b925050602061baed8582860161afe6565b9150509250929050565b60006020828403121561bb0957600080fd5b600061bb178482850161affb565b91505092915050565b6000806040838503121561bb3357600080fd5b600061bb418582860161b064565b925050602061bb528582860161b079565b9150509250929050565b60006020828403121561bb6e57600080fd5b600082013567ffffffffffffffff81111561bb8857600080fd5b61bb948482850161b0cd565b91505092915050565b60006102c0828403121561bbb057600080fd5b600061bbbe8482850161b484565b91505092915050565b60006020828403121561bbd957600080fd5b600082013567ffffffffffffffff81111561bbf357600080fd5b61bbff8482850161b4e6565b91505092915050565b6000610160828403121561bc1b57600080fd5b600061bc298482850161b692565b91505092915050565b60006060828403121561bc4457600080fd5b600061bc528482850161b79a565b91505092915050565b600060c0828403121561bc6d57600080fd5b600061bc7b8482850161b846565b91505092915050565b600061bc90838361bcf4565b60208301905092915050565b600061bca8838361bf25565b905092915050565b600061bcbc838361bf7c565b60208301905092915050565b600061bcd4838361c0b6565b905092915050565b600061bce8838361c97f565b60608301905092915050565b61bcfd8161d01b565b82525050565b61bd0c8161d01b565b82525050565b600061bd1d8261ce1a565b61bd27818561cea8565b935061bd328361cdca565b8060005b8381101561bd6357815161bd4a888261bc84565b975061bd558361ce67565b92505060018101905061bd36565b5085935050505092915050565b600061bd7b8261ce25565b61bd85818561ceb9565b93508360208202850161bd978561cdda565b8060005b8581101561bdd3578484038952815161bdb4858261bc9c565b945061bdbf8361ce74565b925060208a0199505060018101905061bd9b565b50829750879550505050505092915050565b600061bdf08261ce30565b61bdfa818561ceec565b935061be058361cdea565b8060005b8381101561be3657815161be1d888261bcb0565b975061be288361ce81565b92505060018101905061be09565b5085935050505092915050565b600061be4e8261ce3b565b61be58818561ceca565b93508360208202850161be6a8561cdfa565b8060005b8581101561bea6578484038952815161be87858261bcc8565b945061be928361ce8e565b925060208a0199505060018101905061be6e565b50829750879550505050505092915050565b600061bec38261ce46565b61becd818561cedb565b935061bed88361ce0a565b8060005b8381101561bf0957815161bef0888261bcdc565b975061befb8361ce9b565b92505060018101905061bedc565b5085935050505092915050565b61bf1f8161d02d565b82525050565b600061bf308261ce51565b61bf3a818561cefd565b935061bf4a81856020860161d16c565b61bf538161d2a6565b840191505092915050565b61bf678161d0f1565b82525050565b61bf768161d103565b82525050565b61bf858161d115565b82525050565b61bf948161d127565b82525050565b61bfa38161d139565b82525050565b61bfb28161d14b565b82525050565b600061bfc38261ce5c565b61bfcd818561cf0e565b935061bfdd81856020860161d16c565b61bfe68161d2a6565b840191505092915050565b600061bffe602e8361cf0e565b915061c0098261d2b7565b604082019050919050565b600060e08301600083015161c02c600086018261c8af565b50602083015184820360a086015261c044828261be43565b915050604083015161c05960c086018261bf16565b508091505092915050565b60006101208301600083015161c07d600086018261c8af565b50602083015161c09060a086018261c6d3565b50604083015184820361010086015261c0a9828261be43565b9150508091505092915050565b600061032083016000830151848203600086015261c0d4828261bf25565b915050602083015161c0e9602086018261bcf4565b506040830151848203604086015261c101828261bf25565b915050606083015161c116606086018261c9df565b50608083015161c129608086018261c9df565b5060a083015161c13c60a086018261c9df565b5060c083015161c14f60c086018261c9df565b5060e083015161c16260e086018261c9df565b5061010083015161c17761010086018261c9c1565b5061012083015161c18c61012086018261c9df565b5061014083015161c1a161014086018261c9df565b5061016083015184820361016086015261c1bb828261bf25565b91505061018083015161c1d261018086018261c9df565b506101a083015161c1e76101a086018261c9c1565b506101c083015161c1fc6101c086018261bf16565b506101e083015161c2116101e086018261bf7c565b5061020083015161c22661020086018261c9df565b5061022083015184820361022086015261c240828261bd12565b91505061024083015184820361024086015261c25c828261bd12565b91505061026083015184820361026086015261c278828261bf25565b91505061028083015161c28f61028086018261bf6d565b506102a083015161c2a46102a086018261bf16565b506102c083015161c2b96102c086018261c4d2565b508091505092915050565b600061032083016000830151848203600086015261c2e2828261bf25565b915050602083015161c2f7602086018261bcf4565b506040830151848203604086015261c30f828261bf25565b915050606083015161c324606086018261c9df565b50608083015161c337608086018261c9df565b5060a083015161c34a60a086018261c9df565b5060c083015161c35d60c086018261c9df565b5060e083015161c37060e086018261c9df565b5061010083015161c38561010086018261c9c1565b5061012083015161c39a61012086018261c9df565b5061014083015161c3af61014086018261c9df565b5061016083015184820361016086015261c3c9828261bf25565b91505061018083015161c3e061018086018261c9df565b506101a083015161c3f56101a086018261c9c1565b506101c083015161c40a6101c086018261bf16565b506101e083015161c41f6101e086018261bf7c565b5061020083015161c43461020086018261c9df565b5061022083015184820361022086015261c44e828261bd12565b91505061024083015184820361024086015261c46a828261bd12565b91505061026083015184820361026086015261c486828261bf25565b91505061028083015161c49d61028086018261bf6d565b506102a083015161c4b26102a086018261bf16565b506102c083015161c4c76102c086018261c4d2565b508091505092915050565b60608201600082015161c4e8600085018261c9df565b50602082015161c4fb602085018261c9df565b50604082015161c50e604085018261c9df565b50505050565b60006101408301600083015161c52d600086018261c8af565b50602083015161c54060a086018261c6d3565b50604083015184820361010086015261c559828261be43565b915050606083015161c56f61012086018261bf16565b508091505092915050565b60006101208301600083015161c593600086018261c8af565b50602083015161c5a660a086018261c9df565b50604083015161c5b960c086018261c9df565b50606083015184820360e086015261c5d1828261be43565b915050608083015161c5e761010086018261bf16565b508091505092915050565b6101608201600082015161c609600085018261c9df565b50602082015161c61c602085018261c9df565b50604082015161c62f604085018261c9df565b50606082015161c642606085018261c9df565b50608082015161c655608085018261c9df565b5060a082015161c66860a085018261c9df565b5060c082015161c67b60c085018261c9df565b5060e082015161c68e60e085018261c9df565b5061010082015161c6a361010085018261c9df565b5061012082015161c6b861012085018261c9df565b5061014082015161c6cd61014085018261c9df565b50505050565b60608201600082015161c6e9600085018261bcf4565b50602082015161c6fc602085018261bcf4565b50604082015161c70f604085018261c9df565b50505050565b60608201600082015161c72b600085018261bcf4565b50602082015161c73e602085018261bcf4565b50604082015161c751604085018261c9df565b50505050565b60006101e083016000830151848203600086015261c775828261bf25565b915050602083015161c78a602086018261c9df565b50604083015161c79d604086018261c9df565b50606083015161c7b0606086018261c9df565b50608083015161c7c3608086018261c9c1565b5060a083015161c7d660a086018261c9df565b5060c083015161c7e960c086018261c9df565b5060e083015161c7fc60e086018261bf16565b5061010083015184820361010086015261c816828261bf25565b91505061012083015161c82d61012086018261bf16565b5061014083015161c84261014086018261bf16565b5061016083015184820361016086015261c85c828261bf25565b91505061018083015184820361018086015261c878828261beb8565b9150506101a083015161c88f6101a086018261bf16565b506101c083015161c8a46101c086018261bf7c565b508091505092915050565b60a08201600082015161c8c5600085018261c9df565b50602082015161c8d8602085018261c9df565b50604082015161c8eb604085018261c9df565b50606082015161c8fe606085018261c9c1565b50608082015161c911608085018261c9c1565b50505050565b60a08201600082015161c92d600085018261c9df565b50602082015161c940602085018261c9df565b50604082015161c953604085018261c9df565b50606082015161c966606085018261c9c1565b50608082015161c979608085018261c9c1565b50505050565b60608201600082015161c995600085018261bcf4565b50602082015161c9a8602085018261c9df565b50604082015161c9bb604085018261c9df565b50505050565b61c9ca8161d0d3565b82525050565b61c9d98161d0d3565b82525050565b61c9e88161d0dd565b82525050565b61c9f78161d0dd565b82525050565b600060208201905061ca12600083018461bd03565b92915050565b600060408201905061ca2d600083018561bd03565b818103602083015261ca3f818461bd70565b90509392505050565b6000606082019050818103600083015261ca62818661bd70565b905061ca71602083018561bd03565b818103604083015261ca83818461bde5565b9050949350505050565b60006101a082019050818103600083015261caa8818661bd70565b905061cab7602083018561c5f2565b61cac561018083018461c9d0565b949350505050565b600060e08201905061cae2600083018a61bf5e565b61caef602083018961c9d0565b61cafc604083018861bd03565b61cb09606083018761bf8b565b61cb16608083018661c9ee565b61cb2360a083018561bf8b565b61cb3060c083018461c9ee565b98975050505050505050565b600060208201905061cb51600083018461bf9a565b92915050565b600060208201905061cb6c600083018461bfa9565b92915050565b6000606082019050818103600083015261cb8c818661bfb8565b905061cb9b602083018561c9d0565b61cba8604083018461c9d0565b949350505050565b6000602082019050818103600083015261cbc98161bff1565b9050919050565b6000602082019050818103600083015261cbea818461c014565b905092915050565b6000602082019050818103600083015261cc0c818461c064565b905092915050565b6000602082019050818103600083015261cc2e818461c2c4565b905092915050565b6000602082019050818103600083015261cc50818461c514565b905092915050565b6000602082019050818103600083015261cc72818461c57a565b905092915050565b600060608201905061cc8f600083018461c715565b92915050565b60006101a082019050818103600083015261ccb0818661c757565b905061ccbf602083018561c5f2565b61cccd61018083018461c9d0565b949350505050565b600060a08201905061ccea600083018461c917565b92915050565b600061ccfa61cd0b565b905061cd06828261d19f565b919050565b6000604051905090565b600067ffffffffffffffff82111561cd305761cd2f61d277565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561cd5c5761cd5b61d277565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561cd885761cd8761d277565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561cdb45761cdb361d277565b5b61cdbd8261d2a6565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061cf2a8261d0d3565b915061cf358361d0d3565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561cf6a5761cf6961d219565b5b828201905092915050565b600061cf808261d0dd565b915061cf8b8361d0dd565b92508267ffffffffffffffff0382111561cfa85761cfa761d219565b5b828201905092915050565b600061cfbe8261d0d3565b915061cfc98361d0d3565b92508282101561cfdc5761cfdb61d219565b5b828203905092915050565b600061cff28261d0dd565b915061cffd8361d0dd565b92508282101561d0105761d00f61d219565b5b828203905092915050565b600061d0268261d0b3565b9050919050565b60008115159050919050565b6000819050919050565b600061d04e8261d01b565b9050919050565b600061d0608261d01b565b9050919050565b600081905061d0758261d306565b919050565b600081905061d0888261d31a565b919050565b600081905061d09b8261d32e565b919050565b600081905061d0ae8261d342565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061d0fc8261d067565b9050919050565b600061d10e8261d07a565b9050919050565b600061d1208261d08d565b9050919050565b600061d1328261d0a0565b9050919050565b600061d1448261d0dd565b9050919050565b600061d1568261d0dd565b9050919050565b82818337600083830152505050565b60005b8381101561d18a57808201518184015260208101905061d16f565b8381111561d199576000848401525b50505050565b61d1a88261d2a6565b810181811067ffffffffffffffff8211171561d1c75761d1c661d277565b5b80604052505050565b600061d1db8261d0d3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561d20e5761d20d61d219565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600a811061d3175761d31661d248565b5b50565b6003811061d32b5761d32a61d248565b5b50565b6002811061d33f5761d33e61d248565b5b50565b6003811061d3535761d35261d248565b5b50565b61d35f8161d01b565b811461d36a57600080fd5b50565b61d3768161d02d565b811461d38157600080fd5b50565b61d38d8161d039565b811461d39857600080fd5b50565b61d3a48161d043565b811461d3af57600080fd5b50565b61d3bb8161d055565b811461d3c657600080fd5b50565b6003811061d3d657600080fd5b50565b6002811061d3e657600080fd5b50565b6003811061d3f657600080fd5b50565b61d4028161d0d3565b811461d40d57600080fd5b50565b61d4198161d0dd565b811461d42457600080fd5b5056fea2646970667358221220adb70fa5305efdbb0299ce2314533f2be0f11f7f04ab98fd2cfd418c931bb48264736f6c63430008040033",
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

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreCaller) GetUserSpace(opts *bind.CallOpts, walletAddr common.Address) (UserSpace, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUserSpace", walletAddr)

	if err != nil {
		return *new(UserSpace), err
	}

	out0 := *abi.ConvertType(out[0], new(UserSpace)).(*UserSpace)

	return out0, err

}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreSession) GetUserSpace(walletAddr common.Address) (UserSpace, error) {
	return _Store.Contract.GetUserSpace(&_Store.CallOpts, walletAddr)
}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreCallerSession) GetUserSpace(walletAddr common.Address) (UserSpace, error) {
	return _Store.Contract.GetUserSpace(&_Store.CallOpts, walletAddr)
}

// C0x936f0bd2 is a free data retrieval call binding the contract method 0xb5bc8ee2.
//
// Solidity: function c_0x936f0bd2(bytes32 c__0x936f0bd2) pure returns()
func (_Store *StoreCaller) C0x936f0bd2(opts *bind.CallOpts, c__0x936f0bd2 [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x936f0bd2", c__0x936f0bd2)

	if err != nil {
		return err
	}

	return err

}

// C0x936f0bd2 is a free data retrieval call binding the contract method 0xb5bc8ee2.
//
// Solidity: function c_0x936f0bd2(bytes32 c__0x936f0bd2) pure returns()
func (_Store *StoreSession) C0x936f0bd2(c__0x936f0bd2 [32]byte) error {
	return _Store.Contract.C0x936f0bd2(&_Store.CallOpts, c__0x936f0bd2)
}

// C0x936f0bd2 is a free data retrieval call binding the contract method 0xb5bc8ee2.
//
// Solidity: function c_0x936f0bd2(bytes32 c__0x936f0bd2) pure returns()
func (_Store *StoreCallerSession) C0x936f0bd2(c__0x936f0bd2 [32]byte) error {
	return _Store.Contract.C0x936f0bd2(&_Store.CallOpts, c__0x936f0bd2)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) AddUserSpace(opts *bind.TransactOpts, params SpaceAddParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddUserSpace", params)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) AddUserSpace(params SpaceAddParams) (*types.Transaction, error) {
	return _Store.Contract.AddUserSpace(&_Store.TransactOpts, params)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) AddUserSpace(params SpaceAddParams) (*types.Transaction, error) {
	return _Store.Contract.AddUserSpace(&_Store.TransactOpts, params)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreTransactor) DeleteUserSpace(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteUserSpace", walletAddr)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreSession) DeleteUserSpace(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUserSpace(&_Store.TransactOpts, walletAddr)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreTransactorSession) DeleteUserSpace(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUserSpace(&_Store.TransactOpts, walletAddr)
}

// GetUpdateCost is a paid mutator transaction binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns((address,address,uint64))
func (_Store *StoreTransactor) GetUpdateCost(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "GetUpdateCost", params)
}

// GetUpdateCost is a paid mutator transaction binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns((address,address,uint64))
func (_Store *StoreSession) GetUpdateCost(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUpdateCost(&_Store.TransactOpts, params)
}

// GetUpdateCost is a paid mutator transaction binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns((address,address,uint64))
func (_Store *StoreTransactorSession) GetUpdateCost(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUpdateCost(&_Store.TransactOpts, params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreTransactor) ManageUserSpace(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ManageUserSpace", params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreSession) ManageUserSpace(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.ManageUserSpace(&_Store.TransactOpts, params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreTransactorSession) ManageUserSpace(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.ManageUserSpace(&_Store.TransactOpts, params)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreTransactor) UpdateUserSpace(opts *bind.TransactOpts, walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateUserSpace", walletAddr, _userSpace)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreSession) UpdateUserSpace(walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.Contract.UpdateUserSpace(&_Store.TransactOpts, walletAddr, _userSpace)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreTransactorSession) UpdateUserSpace(walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.Contract.UpdateUserSpace(&_Store.TransactOpts, walletAddr, _userSpace)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]))
func (_Store *StoreTransactor) GetUserspaceChange(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "getUserspaceChange", params)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]))
func (_Store *StoreSession) GetUserspaceChange(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUserspaceChange(&_Store.TransactOpts, params)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]))
func (_Store *StoreTransactorSession) GetUserspaceChange(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUserspaceChange(&_Store.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _fs) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _fs common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _fs)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _fs) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _fs common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _fs) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _fs common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) ProcessForUserSpaceOneAddOneRevoke(opts *bind.TransactOpts, params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "processForUserSpaceOneAddOneRevoke", params)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) ProcessForUserSpaceOneAddOneRevoke(params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOneAddOneRevoke(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) ProcessForUserSpaceOneAddOneRevoke(params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOneAddOneRevoke(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) ProcessForUserSpaceOperations(opts *bind.TransactOpts, params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "processForUserSpaceOperations", params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) ProcessForUserSpaceOperations(params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOperations(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) ProcessForUserSpaceOperations(params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOperations(&_Store.TransactOpts, params)
}

// StoreCreateSectorEventIterator is returned from FilterCreateSectorEvent and is used to iterate over the raw logs and unpacked data for CreateSectorEvent events raised by the Store contract.
type StoreCreateSectorEventIterator struct {
	Event *StoreCreateSectorEvent // Event containing the contract specifics and raw log

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
func (it *StoreCreateSectorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCreateSectorEvent)
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
		it.Event = new(StoreCreateSectorEvent)
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
func (it *StoreCreateSectorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCreateSectorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCreateSectorEvent represents a CreateSectorEvent event raised by the Store contract.
type StoreCreateSectorEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SectorId    uint64
	ProvLevel   uint8
	Size        uint64
	IsPlots     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateSectorEvent is a free log retrieval operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) FilterCreateSectorEvent(opts *bind.FilterOpts) (*StoreCreateSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreCreateSectorEventIterator{contract: _Store.contract, event: "CreateSectorEvent", logs: logs, sub: sub}, nil
}

// WatchCreateSectorEvent is a free log subscription operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) WatchCreateSectorEvent(opts *bind.WatchOpts, sink chan<- *StoreCreateSectorEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCreateSectorEvent)
				if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
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

// ParseCreateSectorEvent is a log parse operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) ParseCreateSectorEvent(log types.Log) (*StoreCreateSectorEvent, error) {
	event := new(StoreCreateSectorEvent)
	if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteFileEventIterator is returned from FilterDeleteFileEvent and is used to iterate over the raw logs and unpacked data for DeleteFileEvent events raised by the Store contract.
type StoreDeleteFileEventIterator struct {
	Event *StoreDeleteFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreDeleteFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteFileEvent)
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
		it.Event = new(StoreDeleteFileEvent)
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
func (it *StoreDeleteFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteFileEvent represents a DeleteFileEvent event raised by the Store contract.
type StoreDeleteFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteFileEvent is a free log retrieval operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) FilterDeleteFileEvent(opts *bind.FilterOpts) (*StoreDeleteFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteFileEventIterator{contract: _Store.contract, event: "DeleteFileEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteFileEvent is a free log subscription operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) WatchDeleteFileEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteFileEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteFileEvent", log); err != nil {
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

// ParseDeleteFileEvent is a log parse operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) ParseDeleteFileEvent(log types.Log) (*StoreDeleteFileEvent, error) {
	event := new(StoreDeleteFileEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteFilesEventIterator is returned from FilterDeleteFilesEvent and is used to iterate over the raw logs and unpacked data for DeleteFilesEvent events raised by the Store contract.
type StoreDeleteFilesEventIterator struct {
	Event *StoreDeleteFilesEvent // Event containing the contract specifics and raw log

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
func (it *StoreDeleteFilesEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteFilesEvent)
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
		it.Event = new(StoreDeleteFilesEvent)
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
func (it *StoreDeleteFilesEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteFilesEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteFilesEvent represents a DeleteFilesEvent event raised by the Store contract.
type StoreDeleteFilesEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHashs   [][]byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteFilesEvent is a free log retrieval operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) FilterDeleteFilesEvent(opts *bind.FilterOpts) (*StoreDeleteFilesEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteFilesEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteFilesEventIterator{contract: _Store.contract, event: "DeleteFilesEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteFilesEvent is a free log subscription operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) WatchDeleteFilesEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteFilesEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteFilesEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteFilesEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteFilesEvent", log); err != nil {
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

// ParseDeleteFilesEvent is a log parse operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) ParseDeleteFilesEvent(log types.Log) (*StoreDeleteFilesEvent, error) {
	event := new(StoreDeleteFilesEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteFilesEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteSectorEventIterator is returned from FilterDeleteSectorEvent and is used to iterate over the raw logs and unpacked data for DeleteSectorEvent events raised by the Store contract.
type StoreDeleteSectorEventIterator struct {
	Event *StoreDeleteSectorEvent // Event containing the contract specifics and raw log

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
func (it *StoreDeleteSectorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteSectorEvent)
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
		it.Event = new(StoreDeleteSectorEvent)
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
func (it *StoreDeleteSectorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteSectorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteSectorEvent represents a DeleteSectorEvent event raised by the Store contract.
type StoreDeleteSectorEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SectorId    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteSectorEvent is a free log retrieval operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) FilterDeleteSectorEvent(opts *bind.FilterOpts) (*StoreDeleteSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteSectorEventIterator{contract: _Store.contract, event: "DeleteSectorEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteSectorEvent is a free log subscription operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) WatchDeleteSectorEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteSectorEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteSectorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteSectorEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteSectorEvent", log); err != nil {
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

// ParseDeleteSectorEvent is a log parse operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) ParseDeleteSectorEvent(log types.Log) (*StoreDeleteSectorEvent, error) {
	event := new(StoreDeleteSectorEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFilePDPSuccessEventIterator is returned from FilterFilePDPSuccessEvent and is used to iterate over the raw logs and unpacked data for FilePDPSuccessEvent events raised by the Store contract.
type StoreFilePDPSuccessEventIterator struct {
	Event *StoreFilePDPSuccessEvent // Event containing the contract specifics and raw log

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
func (it *StoreFilePDPSuccessEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFilePDPSuccessEvent)
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
		it.Event = new(StoreFilePDPSuccessEvent)
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
func (it *StoreFilePDPSuccessEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFilePDPSuccessEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFilePDPSuccessEvent represents a FilePDPSuccessEvent event raised by the Store contract.
type StoreFilePDPSuccessEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFilePDPSuccessEvent is a free log retrieval operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) FilterFilePDPSuccessEvent(opts *bind.FilterOpts) (*StoreFilePDPSuccessEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FilePDPSuccessEvent")
	if err != nil {
		return nil, err
	}
	return &StoreFilePDPSuccessEventIterator{contract: _Store.contract, event: "FilePDPSuccessEvent", logs: logs, sub: sub}, nil
}

// WatchFilePDPSuccessEvent is a free log subscription operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) WatchFilePDPSuccessEvent(opts *bind.WatchOpts, sink chan<- *StoreFilePDPSuccessEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FilePDPSuccessEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFilePDPSuccessEvent)
				if err := _Store.contract.UnpackLog(event, "FilePDPSuccessEvent", log); err != nil {
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

// ParseFilePDPSuccessEvent is a log parse operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) ParseFilePDPSuccessEvent(log types.Log) (*StoreFilePDPSuccessEvent, error) {
	event := new(StoreFilePDPSuccessEvent)
	if err := _Store.contract.UnpackLog(event, "FilePDPSuccessEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreGetUpdateCostEventIterator is returned from FilterGetUpdateCostEvent and is used to iterate over the raw logs and unpacked data for GetUpdateCostEvent events raised by the Store contract.
type StoreGetUpdateCostEventIterator struct {
	Event *StoreGetUpdateCostEvent // Event containing the contract specifics and raw log

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
func (it *StoreGetUpdateCostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreGetUpdateCostEvent)
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
		it.Event = new(StoreGetUpdateCostEvent)
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
func (it *StoreGetUpdateCostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreGetUpdateCostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreGetUpdateCostEvent represents a GetUpdateCostEvent event raised by the Store contract.
type StoreGetUpdateCostEvent struct {
	State TransferState
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGetUpdateCostEvent is a free log retrieval operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) FilterGetUpdateCostEvent(opts *bind.FilterOpts) (*StoreGetUpdateCostEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "GetUpdateCostEvent")
	if err != nil {
		return nil, err
	}
	return &StoreGetUpdateCostEventIterator{contract: _Store.contract, event: "GetUpdateCostEvent", logs: logs, sub: sub}, nil
}

// WatchGetUpdateCostEvent is a free log subscription operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) WatchGetUpdateCostEvent(opts *bind.WatchOpts, sink chan<- *StoreGetUpdateCostEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "GetUpdateCostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreGetUpdateCostEvent)
				if err := _Store.contract.UnpackLog(event, "GetUpdateCostEvent", log); err != nil {
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

// ParseGetUpdateCostEvent is a log parse operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) ParseGetUpdateCostEvent(log types.Log) (*StoreGetUpdateCostEvent, error) {
	event := new(StoreGetUpdateCostEvent)
	if err := _Store.contract.UnpackLog(event, "GetUpdateCostEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreProveFileEventIterator is returned from FilterProveFileEvent and is used to iterate over the raw logs and unpacked data for ProveFileEvent events raised by the Store contract.
type StoreProveFileEventIterator struct {
	Event *StoreProveFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreProveFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreProveFileEvent)
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
		it.Event = new(StoreProveFileEvent)
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
func (it *StoreProveFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreProveFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreProveFileEvent represents a ProveFileEvent event raised by the Store contract.
type StoreProveFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	Profit      uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProveFileEvent is a free log retrieval operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) FilterProveFileEvent(opts *bind.FilterOpts) (*StoreProveFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ProveFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreProveFileEventIterator{contract: _Store.contract, event: "ProveFileEvent", logs: logs, sub: sub}, nil
}

// WatchProveFileEvent is a free log subscription operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) WatchProveFileEvent(opts *bind.WatchOpts, sink chan<- *StoreProveFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ProveFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreProveFileEvent)
				if err := _Store.contract.UnpackLog(event, "ProveFileEvent", log); err != nil {
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

// ParseProveFileEvent is a log parse operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) ParseProveFileEvent(log types.Log) (*StoreProveFileEvent, error) {
	event := new(StoreProveFileEvent)
	if err := _Store.contract.UnpackLog(event, "ProveFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRegisterNodeEventIterator is returned from FilterRegisterNodeEvent and is used to iterate over the raw logs and unpacked data for RegisterNodeEvent events raised by the Store contract.
type StoreRegisterNodeEventIterator struct {
	Event *StoreRegisterNodeEvent // Event containing the contract specifics and raw log

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
func (it *StoreRegisterNodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRegisterNodeEvent)
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
		it.Event = new(StoreRegisterNodeEvent)
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
func (it *StoreRegisterNodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRegisterNodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRegisterNodeEvent represents a RegisterNodeEvent event raised by the Store contract.
type StoreRegisterNodeEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	NodeAddr    common.Address
	Volume      uint64
	ServiceTime uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterNodeEvent is a free log retrieval operation binding the contract event 0x2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, address nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) FilterRegisterNodeEvent(opts *bind.FilterOpts) (*StoreRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreRegisterNodeEventIterator{contract: _Store.contract, event: "RegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchRegisterNodeEvent is a free log subscription operation binding the contract event 0x2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, address nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) WatchRegisterNodeEvent(opts *bind.WatchOpts, sink chan<- *StoreRegisterNodeEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRegisterNodeEvent)
				if err := _Store.contract.UnpackLog(event, "RegisterNodeEvent", log); err != nil {
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

// ParseRegisterNodeEvent is a log parse operation binding the contract event 0x2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, address nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) ParseRegisterNodeEvent(log types.Log) (*StoreRegisterNodeEvent, error) {
	event := new(StoreRegisterNodeEvent)
	if err := _Store.contract.UnpackLog(event, "RegisterNodeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetUserSpaceEventIterator is returned from FilterSetUserSpaceEvent and is used to iterate over the raw logs and unpacked data for SetUserSpaceEvent events raised by the Store contract.
type StoreSetUserSpaceEventIterator struct {
	Event *StoreSetUserSpaceEvent // Event containing the contract specifics and raw log

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
func (it *StoreSetUserSpaceEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetUserSpaceEvent)
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
		it.Event = new(StoreSetUserSpaceEvent)
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
func (it *StoreSetUserSpaceEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetUserSpaceEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetUserSpaceEvent represents a SetUserSpaceEvent event raised by the Store contract.
type StoreSetUserSpaceEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SizeType    uint8
	Size        uint64
	CountType   uint8
	Count       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetUserSpaceEvent is a free log retrieval operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) FilterSetUserSpaceEvent(opts *bind.FilterOpts) (*StoreSetUserSpaceEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetUserSpaceEvent")
	if err != nil {
		return nil, err
	}
	return &StoreSetUserSpaceEventIterator{contract: _Store.contract, event: "SetUserSpaceEvent", logs: logs, sub: sub}, nil
}

// WatchSetUserSpaceEvent is a free log subscription operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) WatchSetUserSpaceEvent(opts *bind.WatchOpts, sink chan<- *StoreSetUserSpaceEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetUserSpaceEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetUserSpaceEvent)
				if err := _Store.contract.UnpackLog(event, "SetUserSpaceEvent", log); err != nil {
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

// ParseSetUserSpaceEvent is a log parse operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) ParseSetUserSpaceEvent(log types.Log) (*StoreSetUserSpaceEvent, error) {
	event := new(StoreSetUserSpaceEvent)
	if err := _Store.contract.UnpackLog(event, "SetUserSpaceEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStoreFileEventIterator is returned from FilterStoreFileEvent and is used to iterate over the raw logs and unpacked data for StoreFileEvent events raised by the Store contract.
type StoreStoreFileEventIterator struct {
	Event *StoreStoreFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreStoreFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStoreFileEvent)
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
		it.Event = new(StoreStoreFileEvent)
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
func (it *StoreStoreFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStoreFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStoreFileEvent represents a StoreFileEvent event raised by the Store contract.
type StoreStoreFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	FileSize    uint64
	WalletAddr  common.Address
	Cost        uint64
	IsPlotFile  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStoreFileEvent is a free log retrieval operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) FilterStoreFileEvent(opts *bind.FilterOpts) (*StoreStoreFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StoreFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreStoreFileEventIterator{contract: _Store.contract, event: "StoreFileEvent", logs: logs, sub: sub}, nil
}

// WatchStoreFileEvent is a free log subscription operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) WatchStoreFileEvent(opts *bind.WatchOpts, sink chan<- *StoreStoreFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StoreFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStoreFileEvent)
				if err := _Store.contract.UnpackLog(event, "StoreFileEvent", log); err != nil {
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

// ParseStoreFileEvent is a log parse operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) ParseStoreFileEvent(log types.Log) (*StoreStoreFileEvent, error) {
	event := new(StoreStoreFileEvent)
	if err := _Store.contract.UnpackLog(event, "StoreFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUnRegisterNodeEventIterator is returned from FilterUnRegisterNodeEvent and is used to iterate over the raw logs and unpacked data for UnRegisterNodeEvent events raised by the Store contract.
type StoreUnRegisterNodeEventIterator struct {
	Event *StoreUnRegisterNodeEvent // Event containing the contract specifics and raw log

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
func (it *StoreUnRegisterNodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUnRegisterNodeEvent)
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
		it.Event = new(StoreUnRegisterNodeEvent)
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
func (it *StoreUnRegisterNodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUnRegisterNodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUnRegisterNodeEvent represents a UnRegisterNodeEvent event raised by the Store contract.
type StoreUnRegisterNodeEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnRegisterNodeEvent is a free log retrieval operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) FilterUnRegisterNodeEvent(opts *bind.FilterOpts) (*StoreUnRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "UnRegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreUnRegisterNodeEventIterator{contract: _Store.contract, event: "UnRegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchUnRegisterNodeEvent is a free log subscription operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) WatchUnRegisterNodeEvent(opts *bind.WatchOpts, sink chan<- *StoreUnRegisterNodeEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "UnRegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUnRegisterNodeEvent)
				if err := _Store.contract.UnpackLog(event, "UnRegisterNodeEvent", log); err != nil {
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

// ParseUnRegisterNodeEvent is a log parse operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) ParseUnRegisterNodeEvent(log types.Log) (*StoreUnRegisterNodeEvent, error) {
	event := new(StoreUnRegisterNodeEvent)
	if err := _Store.contract.UnpackLog(event, "UnRegisterNodeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
