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

// FSConfig is an auto generated low-level Go binding around an user-defined struct.
type FSConfig struct {
	DEFAULTBLOCKINTERVAL uint64
	DEFAULTPROVEPERIOD   uint64
	INSECTORSIZE         uint64
}

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

// FileReNewInfo is an auto generated low-level Go binding around an user-defined struct.
type FileReNewInfo struct {
	FileHash   []byte
	FromAddr   common.Address
	ReNewTimes uint64
}

// OwnerChange is an auto generated low-level Go binding around an user-defined struct.
type OwnerChange struct {
	FileHash []byte
	CurOwner common.Address
	NewOwner common.Address
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

// PriChange is an auto generated low-level Go binding around an user-defined struct.
type PriChange struct {
	FileHash  []byte
	Privilege uint64
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

// StorageFee is an auto generated low-level Go binding around an user-defined struct.
type StorageFee struct {
	TxnFee        uint64
	SpaceFee      uint64
	ValidationFee uint64
}

// UploadOption is an auto generated low-level Go binding around an user-defined struct.
type UploadOption struct {
	FileDesc        []byte
	FileSize        uint64
	ProveInterval   uint64
	ProveLevel      uint64
	ExpiredHeight   *big.Int
	Privilege       uint64
	CopyNum         uint64
	Encrypt         bool
	EncryptPassword []byte
	RegisterDNS     bool
	BindDNS         bool
	DnsURL          []byte
	WhiteList       []WhiteList
	Share           bool
	StorageType     uint8
}

// WhiteList is an auto generated low-level Go binding around an user-defined struct.
type WhiteList struct {
	Addr         common.Address
	BaseHeight   uint64
	ExpireHeight uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"beginHeight\",\"type\":\"uint256\"}],\"name\":\"CalcProveTimesByUploadInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structOwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structPriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromCandidateList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromPrimaryList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromUnSettledList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfoForRenew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xf932268c\",\"type\":\"bytes32\"}],\"name\":\"c_0xf932268c\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcSingleValidFeeForFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"calcUploadFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"files\",\"type\":\"tuple[]\"}],\"name\":\"deleteFilesInner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"deleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractSpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractProve\",\"name\":\"_prove\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"DEFAULT_BLOCK_INTERVAL\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DEFAULT_PROVE_PERIOD\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"IN_SECTOR_SIZE\",\"type\":\"uint64\"}],\"internalType\":\"structFSConfig\",\"name\":\"fsConfig\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50620149f480620000236000396000f3fe608060405260043610620002e45760003560e01c80638d41f9f81162000193578063cc76e80d11620000df578063defce5d4116200008b578063e4bca973116200006d578063e4bca9731462000ae2578063ebaf02581462000b26578063f54cd2951462000b5457620002e4565b8063defce5d41462000a5a578063df49410a1462000a9e57620002e4565b8063d49ce87411620000c1578063d49ce87414620009a4578063d70e627214620009e8578063dc1ec30b1462000a1657620002e4565b8063cc76e80d1462000932578063ce904554146200097657620002e4565b8063b64ab1ef116200013f578063c43007e51162000121578063c43007e5146200087c578063c6e8392a14620008aa578063ca6142cb14620008ee57620002e4565b8063b64ab1ef1462000820578063b6af10fb146200084e57620002e4565b80639a2b5e6311620001755780639a2b5e6314620007805780639cd103e914620007c4578063a1276a5414620007f257620002e4565b80638d41f9f8146200070e57806395b0b54b146200075257620002e4565b80633d4dcb981162000253578063655c669411620001ff57806370b1d88911620001e157806370b1d8891462000642578063740f69131462000686578063793cab8e14620006ca57620002e4565b8063655c669414620005e6578063681850d7146200061457620002e4565b806341bc86cb116200023557806341bc86cb1462000554578063432152ce146200059857806357d9439914620005c657620002e4565b80633d4dcb9814620004e15780633f2cc9a0146200050f57620002e4565b80631bc5588011620002b357806328a8565c116200029557806328a8565c146200042957806334fddaac146200046d5780633ad525a9146200049b57620002e4565b80631bc5588014620003cd578063207e33be14620003fb57620002e4565b80630be0970214620002e95780630df7e99e1462000317578063170bbdf51462000345578063178e4dc01462000389575b600080fd5b348015620002f657600080fd5b506200031560048036038101906200030f919062012275565b62000b74565b005b3480156200032457600080fd5b506200034360048036038101906200033d91906201251d565b620015a0565b005b3480156200035257600080fd5b506200037160048036038101906200036b919062012954565b620015a3565b60405162000380919062013e57565b60405180910390f35b3480156200039657600080fd5b50620003b56004803603810190620003af919062012a90565b62001661565b604051620003c4919062013de0565b60405180910390f35b348015620003da57600080fd5b50620003f96004803603810190620003f39190620122fb565b620017f5565b005b3480156200040857600080fd5b506200042760048036038101906200042191906201274a565b62001b3a565b005b3480156200043657600080fd5b506200045560048036038101906200044f919062012275565b62001e0b565b60405162000464919062013975565b60405180910390f35b3480156200047a57600080fd5b50620004996004803603810190620004939190620122fb565b62001fb7565b005b348015620004a857600080fd5b50620004c76004803603810190620004c191906201239a565b620020c0565b604051620004d89392919062013999565b60405180910390f35b348015620004ee57600080fd5b506200050d600480360381019062000507919062012549565b62002e08565b005b3480156200051c57600080fd5b506200053b600480360381019062000535919062012422565b620031b5565b6040516200054b92919062013a01565b60405180910390f35b3480156200056157600080fd5b506200058060048036038101906200057a919062012a4b565b620037e0565b6040516200058f919062013de0565b60405180910390f35b348015620005a557600080fd5b50620005c46004803603810190620005be919062012355565b62003a99565b005b620005e46004803603810190620005de9190620126d9565b62003f3f565b005b348015620005f357600080fd5b506200061260048036038101906200060c9190620122fb565b62005577565b005b3480156200062157600080fd5b506200064060048036038101906200063a919062012625565b620058bc565b005b3480156200064f57600080fd5b506200066e600480360381019062000668919062012954565b62005dc2565b6040516200067d919062013e57565b60405180910390f35b3480156200069357600080fd5b50620006b26004803603810190620006ac919062012846565b62005e80565b604051620006c1919062013a35565b60405180910390f35b348015620006d757600080fd5b50620006f66004803603810190620006f09190620128fb565b62006587565b60405162000705919062013e57565b60405180910390f35b3480156200071b57600080fd5b506200073a600480360381019062000734919062012275565b62006635565b60405162000749919062013975565b60405180910390f35b3480156200075f57600080fd5b506200077e600480360381019062000778919062012549565b62007018565b005b3480156200078d57600080fd5b50620007ac6004803603810190620007a6919062012275565b620072d8565b604051620007bb919062013975565b60405180910390f35b348015620007d157600080fd5b50620007f06004803603810190620007ea91906201266a565b62007484565b005b348015620007ff57600080fd5b506200081e60048036038101906200081891906201258e565b62007ab4565b005b3480156200082d57600080fd5b506200084c6004803603810190620008469190620122fb565b620080a4565b005b3480156200085b57600080fd5b506200087a60048036038101906200087491906201278f565b620083e8565b005b3480156200088957600080fd5b50620008a86004803603810190620008a2919062012493565b62008570565b005b348015620008b757600080fd5b50620008d66004803603810190620008d0919062012355565b6200a54c565b604051620008e59190620139dd565b60405180910390f35b348015620008fb57600080fd5b506200091a6004803603810190620009149190620128b8565b6200babd565b60405162000929919062013e57565b60405180910390f35b3480156200093f57600080fd5b506200095e6004803603810190620009589190620129c4565b6200bb73565b6040516200096d919062013de0565b60405180910390f35b3480156200098357600080fd5b50620009a260048036038101906200099c919062012549565b6200be50565b005b348015620009b157600080fd5b50620009d06004803603810190620009ca919062012a90565b6200c178565b604051620009df919062013de0565b60405180910390f35b348015620009f557600080fd5b5062000a14600480360381019062000a0e9190620122a1565b6200c876565b005b34801562000a2357600080fd5b5062000a42600480360381019062000a3c919062012275565b6200c95a565b60405162000a51919062013975565b60405180910390f35b34801562000a6757600080fd5b5062000a86600480360381019062000a80919062012549565b6200d33d565b60405162000a95919062013d47565b60405180910390f35b34801562000aab57600080fd5b5062000aca600480360381019062000ac49190620128fb565b6200de9d565b60405162000ad9919062013e57565b60405180910390f35b34801562000aef57600080fd5b5062000b0e600480360381019062000b08919062012b01565b6200df61565b60405162000b1d919062013e57565b60405180910390f35b34801562000b3357600080fd5b5062000b52600480360381019062000b4c9190620122fb565b6200e026565b005b62000b72600480360381019062000b6c919062012625565b6200e36b565b005b62000ba27fbfe1a6994139c5c3920edf29161b82f022ff62ad1ab4a250834c06835a8c344660001b620015a0565b62000bd07f7bcc5fdd9ec0580c1722f61e75e2a62b10d0d71770cc7324436aca6efc22fcf960001b620015a0565b62000bfe7f9105a3a8553720748ef3a0cfbdebd7109dcce51d78bf89ea9c327684fc2ed7de60001b620015a0565b600062000c0b8262001e0b565b905062000c3b7ff646fd7c0d38b20681bf0586134cb86672a277fc328eeefc556e6efbe95f3d4660001b620015a0565b62000c697f0cc778ea07ee202fd18b2705852842da3dcb7cbf439e4e0a821f9611801c962760001b620015a0565b6000600267ffffffffffffffff81111562000cad577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562000cdc5781602001602082028036833780820191505090505b50905062000d0d7fb64fcc500acc72e76b1848a0ad077381b819865ffa83369ffd161f552e2fd9ed60001b620015a0565b62000d3b7fe4a62461237b44e16869ffc69797af14a272a13baf3eb2b412dc2fbfbd66000160001b620015a0565b60008160008151811062000d78577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010190600181111562000db9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9081600181111562000df4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505062000e267fca4edae81da8eeb08c4484c02b1e6151deda998a354857fe9f91aaf8c5e6af4760001b620015a0565b62000e547f870c243e274a3db17179c525c3b7ab115dc1bfc3fdb42fccb21becc1a1a0b24260001b620015a0565b60018160018151811062000e91577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010190600181111562000ed2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9081600181111562000f0d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505062000f3f7fe6a3c0fbfe2be967f37f499c0f2a74de4f231e11420b3f5f89edf9c4a509ed4160001b620015a0565b62000f6d7faf522b3197196cec4d0743bc6045edcd673f31563b41276ae1621a9fd4709c3660001b620015a0565b606062000f9d7fc166f1f38d83c8c12c65ee27a04d3fc0eac5da2dbc69a041994f27300ac4982160001b620015a0565b62000fcb7f7cf81d927fcb489089907f5bdeac3fd9368422e60270e67b1d2264dfb82b036960001b620015a0565b600062000ffb7fab403076d8474c4a1221c1f940f84b0242d8b058bc2377442ae9685e1fd15e7660001b620015a0565b620010297f113b3da8e79fe1b8a48119013ef02b1a60b6d2fac8b2e72110150f91f01d5a6a60001b620015a0565b6000620010597f6eecab00b91fc42f7567b5e0307da3bcdc5050b8c965d23e0835a7468a82049860001b620015a0565b620010877fa031341cdb61fe16eda0b3bebe42c7df9608ce0dea049b26776b05ef262131fd60001b620015a0565b62001094858786620020c0565b809350819450829550505050620010ce7f704981f986bd5173be4afa36ce2a703df7c8a2d4604ec9189ca46c87ad89f8f360001b620015a0565b620010fc7f3c339fa14726e804890f4753bfa51d6a5270383cab3195999a7e6734696d792a60001b620015a0565b60005b83518110156200159757620011377fb65e33cda92155c40a62d362b87bb240c3c13ebdb57c436fae96ee5c9f05e58860001b620015a0565b620011657f538bd8a7ec958bc898413e0a64adcf2e4f29ca801d67134bf01edfe2bc1b21b860001b620015a0565b6000600b60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156200127c578382906000526020600020018054620011e89062014481565b80601f0160208091040260200160405190810160405280929190818152602001828054620012169062014481565b8015620012675780601f106200123b5761010080835404028352916020019162001267565b820191906000526020600020905b8154815290600101906020018083116200124957829003601f168201915b505050505081526020019060010190620011c6565b505050509050620012b07fa74cd022319f44cef0d19be54826a544d4e719ac292f6dbf6d094c893a8ae8a160001b620015a0565b620012de7f406b247f893c76c384514430b2940b5d2419f279946863b5e9a608aae14ded0460001b620015a0565b60005b8151811015620014cd57620013197f290c43f8188491d9b7e9453926d8bd30aeefab5fc523e2ab141fda0b84535cf960001b620015a0565b620013477fafeb82dcf7ab32416f024cc84323c42a611caa1dba81732c52233ef9d9f53e4260001b620015a0565b85838151811062001381577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015180519060200120828281518110620013ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015180519060200120141562001488576200140e7f8cc5a2bac7d4765f8118076b9184c3b7dbc258e8641ec0cc8da292147306276960001b620015a0565b6200143c7f53e7b591012f2b40706d30807bf14750a5c09698ba4a9eee97f34cc88c6342c660001b620015a0565b81818151811062001476577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001016060815250620014b7565b620014b67f7f36442dc390b58f719b50e43f6e33bb9b69535074d08077f75e7b1db8ba4eb460001b620015a0565b5b8080620014c490620144ed565b915050620012e1565b50620014fc7f77b74c4251ee86c46c6e36a82c45754d953e260a7106636b6f284faf49cd0aa260001b620015a0565b6200152a7fadd58cd4ed0661c490feb138b9ae2842f2a03017e9b4535bf309ab7876aaf50e60001b620015a0565b80600b60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906200157f929190620108fd565b505080806200158e90620144ed565b915050620010ff565b50505050505050565b50565b6000620015d37f880a161c3ba002f98117f1ef5f63d73872c6609c1f2cfd9c883e74b35826496c60001b620015a0565b620016017fa255701f88816329d852a4ea9749c203c3db798d1ce2879aa8e8b22406a819d260001b620015a0565b6200162f7fb313296457da02f8adaa7bfc56b2a35fef7516fc5616fa54e4bf991244c6336560001b620015a0565b6200163c85858462006587565b6001846200164b919062014191565b6200165791906201420e565b9050949350505050565b6200166b62010964565b620016997f4e91a6681efa3db9ad639967a2c82936ee06769fd7e9ab2b90cc2ae33c1f8a2960001b620015a0565b620016c77f9f208eafe9c9a21956a1236aff1641893c24f7725cd4984fd51ef87c96c30e6f60001b620015a0565b620016f57f37b8ec15a0ad47b0ecc19089c7489787c113552898ad84af2e7444579f78986160001b620015a0565b60006200170385846200df61565b9050620017337fbbb16efbcee538bab3f0d880f6510d16da0d619eb024dcf2a775e64b115038e060001b620015a0565b620017617f017530f9276699c0934f17fe346e81e82c694264ccaaedf81802cd3842ed5d0560001b620015a0565b60006200178b85838860c001518960200151888b6080015162001785919062014257565b6200bb73565b9050620017bb7fa7b766ec203c491458a4d371f22ded7db42f460eccc6b776230b5688be25cf9e60001b620015a0565b620017e97f280190ba76b3f309f97c7c32c6dabd1cd81be7db942c38d0020fa4f92e533c7860001b620015a0565b80925050509392505050565b620018237f969a90bf79c018a79a82dbc0ac5a7f992106179602efb699e16f46e212fd437360001b620015a0565b620018517f264b9c954684728d32674a7eb6775add718ae7e5e4a4d20b89606f63e298bf9e60001b620015a0565b6200187f7f2c5c929dfdac5df60c00ca60e13514c29bfa78beb04120499c8bddba12124d7760001b620015a0565b60005b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905081101562001b3557620018fb7f68de6f91f5ee1a8079c782d29c33d90ccd9523231b758f1bc1ecb8e80480ed4f60001b620015a0565b620019297f9e56d2be8696791257ffed1c089b4a8a0e127e009b61868467ef3fdf1f64319260001b620015a0565b8180519060200120600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110620019a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001604051620019c2919062013912565b6040518091039020141562001af157620019ff7fa1781f850664cac5b11a6b9e481dfb02edcae89096e6e0508a9208112410c33b60001b620015a0565b62001a2d7fe1dfee9a86c62f3285b21d6c5ad6f99cc02bc819ddfcc5cd9c21a3dfb2bce81160001b620015a0565b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020818154811062001aa5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001600062001abd9190620109a3565b62001aeb7f5b3457f4e6bf0c5d630cce97a025036cdc2c0d4ce500b54d52859f1d6dc997a560001b620015a0565b62001b35565b62001b1f7f50a091d1db902d839c2b19bf3770bda484c7b93f0b12c36a3285a8fff4437ac760001b620015a0565b808062001b2c90620144ed565b91505062001882565b505050565b62001b687f0586e30b57be65e3ec9f86dc4cdf927f9886d8f5e73c596f6218c020ad717ebf60001b620015a0565b62001b967f8d77dc23123b50fe58d4d181a26c2d360f4d8b44ecee2ebb4e73c8ae3a18c1e660001b620015a0565b62001bc47f9a5653a2eaa6d6cdf3fb7187f885c0dcf3b3ffd70bcc64cc3014fb8febae87fd60001b620015a0565b600062001bd582600001516200d33d565b905062001c057f8893787abbeb55c1e961ba93eb8d9c85f8ce04bd170e6a38dc3f1485e1b9b42060001b620015a0565b62001c337f3f13b7f759a50557a7c26e9b9f6ee9990ce85966d48780e06b1248d8b59a403f60001b620015a0565b62001c617fd2b19eb65fae328aa751afeed0f67d241dc4fee6d2fe97fc65d931a0d0c7967060001b620015a0565b816020015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161462001cda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162001cd19062013c37565b60405180910390fd5b62001d087fd8c707514eac9dc587d13bf51e560c29a5dbef3c0b4263d785ab1be578794a4260001b620015a0565b62001d367fe654e10850e2fcb6da9922b85f3f108bd37e12fefceff0b8738b6fd0f993037160001b620015a0565b62001d647f478c582e21d9516412c65a8574c762cce218b8e79100fbc3bc22047484a8d79d60001b620015a0565b8160400151816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505062001dce7fa0c83b9f239e712c72e7d53608b37be13d49a33da87536cd27bfed1fbc91c2d260001b620015a0565b62001dfc7fd24646d62f6bc931cb4e0f2dd0588bd65201134c2a2cfc248d42e1da68c43b4e60001b620015a0565b62001e0781620058bc565b5050565b606062001e3b7f1429b295cf906a2d14788dd8ea97a649275d99383540b7bf96d73106cd23076560001b620015a0565b62001e697f23c5a9f785df02655797d51d91f4f35352a1e3cb255214770895860bd24bdd5a60001b620015a0565b62001e977fd43e529f56e455ac2f4fb45259e3b83f3e2aa2f872983dc4c36529a53d10149960001b620015a0565b600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101562001fac57838290600052602060002001805462001f189062014481565b80601f016020809104026020016040519081016040528092919081815260200182805462001f469062014481565b801562001f975780601f1062001f6b5761010080835404028352916020019162001f97565b820191906000526020600020905b81548152906001019060200180831162001f7957829003601f168201915b50505050508152602001906001019062001ef6565b505050509050919050565b62001fe57fe648fc8c65a2ebedeafae6ec7519fd4397a1e2b4011c85238dfcdd5d3cef469d60001b620015a0565b620020137fadcffe0225e69819f6e7e26c9dd6d9530ae911b021cd434fe70c7a42312ca0ea60001b620015a0565b620020417fe5e5a57148cd6e37d886f38d7ef20525f6290d446b85e6f781747a32880319ef60001b620015a0565b600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190620020bb929190620109e9565b505050565b6060600080620020f37f1b1b0212a93dfb3b904a5fdd751d7b502f6dc88a19fc0bbb752f7fe49f9af24460001b620015a0565b620021217f58494430edebd0bfc25dacfe8e9c116a568776ac7286d2a926b527b9d7c0069460001b620015a0565b6200214f7ff6828bf3114da7f08750dae437fcf16c48972b15cd9f330c6cbfcb0e73a5d7ea60001b620015a0565b6000865167ffffffffffffffff81111562002193577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015620021c857816020015b6060815260200190600190039081620021b25790505b509050620021f97f2080fe531c704284ca9356f9302cd072d7a39d48ab8341227501d6e4ceef99e160001b620015a0565b620022277fea67f9b66be064445b1748b5f096828c679f2de93343a02076cff35fde7ed55960001b620015a0565b6000620022577f1add6da8ee3f6ba578d4e447e109d644317c4022ba910bd61c45227f0393268060001b620015a0565b620022857f120d3e6be5131e1413abf41f1fccb91d6cc4d5e5975e0ed3035fd8b7a5080c8960001b620015a0565b6000620022b57f8f85ea4682355a67f38023a766f7501af204ae69530ea809a22ee3eb3507363660001b620015a0565b620022e37fafd419c1af0ff347baef3621ba4e45236cf645d10468edbce4533409c1429f0f60001b620015a0565b6000620023137f74b08d2e3558ef981c442a3b6f217925bc7caae9d7723a26aa7d28ccd8b3dec960001b620015a0565b620023417f86d041f35e25c5f6f2e7791f0113620939cc879649e92cd77e0e1dc8271c9c4360001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015620023ac57600080fd5b505afa158015620023c1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620023e7919062012819565b9050620024177f405dd08ecaffd0f74468cc8551d16c96c7c6bc3b0e382a82e6128c25100a4a9e60001b620015a0565b620024457fa805888c8acbb6ca934192a87ac98cd34f9779df5551be637ff49ecd8fde23f460001b620015a0565b60005b8b5181101562002c1357620024807f174b30a2fc5e050f8268f1355b0ed19e5dd85ec739f3e1b3624331c3fde9cb1d60001b620015a0565b620024ae7f1cd874e1ee896136f87dba1cb043958a016cbd92452d2ccee2a0073ad7398bf460001b620015a0565b6000620024fc8d8381518110620024ee577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516200d33d565b90506200252c7f506db94eb2080c8150dc396fd554c0c911906f22aad4267835d347c11a2cb03060001b620015a0565b6200255a7fd4bb94c1916f2b5d8096b0d607692e10247e648222e19c045b813504bf824a2860001b620015a0565b60006200258a7f8220d22fec50a13c1077b57020d05c2d17d619cbfd6ee62d121cfb1ef489c63660001b620015a0565b620025b87f555811514dee7d461e40903b051c65e8b6060e252142c0cdf0c4e9c292c6b87d60001b620015a0565b60005b8c51811015620027e857620025f37f6a416f336b7f35827e0d7e174577f20406569be6daa4e7b3773a0a5d8b75c6dc60001b620015a0565b620026217f3855e23598b247d76ca0fa210cbb9aa86d0c263a1710c64194ca8ea08b08020760001b620015a0565b8c81815181106200265b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160018111156200269c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836101e001516001811115620026db577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415620027a457620027107fa943e4a93fee2ab99bf1ca9ed36f4d052df1f86be70336acc5fd89e40394422060001b620015a0565b6200273e7ff02c8956f11b0335c7e2eae97034a6e64091dd767989a9719b02c5a78f284a8360001b620015a0565b6200276c7f1cb610e47574e5666dd3bd891e2f445445ebd07a3c18844d3b41fb6b8d2a37fc60001b620015a0565b600191506200279e7f9f10e5b9647235fa3306ee2b62dd9bb3c335e3b077713037789703104944f2fb60001b620015a0565b620027e8565b620027d27f1d1c75644c2971330c0b84fe3c988168a543487eab9a89612feec8bf649444ba60001b620015a0565b8080620027df90620144ed565b915050620025bb565b50620028177f5f75ab263a06300b4372c2bb83f5dda60988655bec508216def027de7827711060001b620015a0565b620028457f66589873f30c5f2b42addfdad2ec2d4fb2c08e7add63cbeac0c9cd67a11a0ce060001b620015a0565b80620028af57620028797f4190543ce769b9973d293c75d872d10ac222300cfc26ddba3e019d097d70b59260001b620015a0565b620028a77f67abbd33dfc55f75ecc348ca69152ba7285c5392748b74df2ce7f0c7fdf6321760001b620015a0565b505062002bfd565b620028dd7fb4b6a28a49a7092351d025e37a9e0c18e8e2e0b818d2ab7caf0e8f67ead364d360001b620015a0565b6200290b7f5944910d27bd38b2ac9cb780fe9e7937407a6cb75683a1b0cf8796a1044fa61060001b620015a0565b620029397f165944f80b87f76e325962a6b249087d508c16881b11e470403855580a006f3a60001b620015a0565b438460c0015167ffffffffffffffff168361010001516200295b919062014134565b1115620029c657620029907fe13da2b1b58484afe5e3677da2e17300d01024d06769b69b5b75f24e094129f660001b620015a0565b620029be7f2fa34e256e0cda332c9fd48aa33b0c84ed6453df96ace899eac89f497082f32660001b620015a0565b505062002bfd565b620029f47fb9de3f8d94fc35de895453f5230018b4c6f43c569d54ba27aab1a693258e43e460001b620015a0565b62002a227f66e0871a83953cd29d41d2e4f8a9e4c6cb9ec0a727b2020cc8cb0c7748fa37f760001b620015a0565b62002a507fc735706bd6112444995bce76f422654ff0d79ff7994747ad50d87aabb850676e60001b620015a0565b8161014001518662002a63919062014191565b955062002a937f26311b37e4bfd4755ffafe382e1d71071623e5e2a07cddd67daceb520de2c64060001b620015a0565b62002ac17f9e8683e1b0b5aa14a71406c61d4184d91ececebe77af04edeb6a20ee474aa36260001b620015a0565b62002acf8260018062007484565b62002afd7f1e74b958b6c76e3d0d4e7f3f937b38d9a8854a3632b64f0674c51ce2f1e23e1f60001b620015a0565b62002b2b7fea4bdad1d85a9ab02103c89a505c7a96f6539c654d59d16a4af43288ae5b2a9b60001b620015a0565b8d838151811062002b65577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151888867ffffffffffffffff168151811062002bb1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525062002bea7f439242bc37ca4a4c87f069e153edb6fd6a8c67dd1b3480d6ca7c0a06f8106d3f60001b620015a0565b868062002bf7906201453b565b97505050505b808062002c0a90620144ed565b91505062002448565b5062002c427f3b1331e7f62e037876d65d6c88c662a459747d88d88dd89e2761174a8a1ae5c660001b620015a0565b62002c707f0c22055de75bd34d8b9103d179cffaea3e66e550b482d864e78158ab1a4554ad60001b620015a0565b60008367ffffffffffffffff16111562002d665762002cb27f354630a43ff387e74a2fd2d8f3083c2e62a1e2d3c4af9d6dde6100160d9696ee60001b620015a0565b62002ce07fc8f418fa287040deb2d2097e74f3b9bc872088318b47b468b5cff84470270f1760001b620015a0565b62002d0e7ff14f7e455a52bc8a661d6f61fc30432c3e29157dd5ca7002e229620558851ce460001b620015a0565b8973ffffffffffffffffffffffffffffffffffffffff166108fc8467ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801562002d5f573d6000803e3d6000fd5b5062002d95565b62002d947fb14e8ae0fb2f14b08c41fdededc353b50721ef93564eabeb5fb668765a96041860001b620015a0565b5b62002dc37f1454479f0c8650b4809ffba427d909ae52c97123c24f000e444041f78da0175660001b620015a0565b62002df17f6806dbdaa4e5b4063411a681076c3b69d21ba4660663c76b667409d899ab1f2360001b620015a0565b848383975097509750505050505093509350939050565b62002e367f639c1b25dc58bacd876d45ec5d597116c0c2cef6667a1ddc86875217813c15a360001b620015a0565b62002e647f92c1de1403fbda8fae9718b088715cfb8f3787c4da1dbf44164e8b9e4a6bcb6360001b620015a0565b60068160405162002e769190620138f9565b90815260200160405180910390206000808201600062002e979190620109a3565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560028201600062002ed09190620109a3565b6003820160006101000a81549067ffffffffffffffff02191690556003820160086101000a81549067ffffffffffffffff02191690556003820160106101000a81549067ffffffffffffffff02191690556003820160186101000a81549067ffffffffffffffff02191690556004820160006101000a81549067ffffffffffffffff021916905560058201600090556006820160006101000a81549067ffffffffffffffff02191690556006820160086101000a81549067ffffffffffffffff021916905560078201600062002fa79190620109a3565b6008820160006101000a81549067ffffffffffffffff02191690556009820160009055600a820160006101000a81549060ff0219169055600a820160016101000a81549060ff0219169055600a820160026101000a81549067ffffffffffffffff0219169055600b820160006200301f919062010a7a565b600c8201600062003031919062010a7a565b600d82016000620030439190620109a3565b600e820160006101000a81549060ff0219169055600e820160016101000a81549060ff0219169055600f8201600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556000820160106101000a81549067ffffffffffffffff021916905550505050620030f37f46df501b828f66bb3584107d7d6bba29236465293566e299cfecaf4ade57723e60001b620015a0565b620031217fdb703d80a5d70bb51644e2a3efe9d313984bc908f8a3478e3a76eccc08f45c9460001b620015a0565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166352e06146826040518263ffffffff1660e01b81526004016200317e919062013a52565b600060405180830381600087803b1580156200319957600080fd5b505af1158015620031ae573d6000803e3d6000fd5b5050505050565b60606000620031e77fcd5cfe51a4c6469b3d1e24f084efd9b4b658068c35f2641bf704b18c5fa5f08660001b620015a0565b620032157ffd83ca1eeb55d9e798ef6a1e5dda8484f9766f4036a06ac12761b480d31f87c460001b620015a0565b620032437f0239c4dd6dc1065463e0a22351d7a7b5508ee169c2c793e0a063e54a7ce564f560001b620015a0565b6060620032737ff66051e17f76627f2d53ed53656cf358851d8db646b8897a2e222771544aa56760001b620015a0565b620032a17f26d2ac6fd0a897c30bb5bceb8173852a7d5caf31fe24647708eaa6fa126f64eb60001b620015a0565b60005b86518110156200377257620032dc7f60065f963c9d71e293f25e58671f200bf9c513b521353ff94b85d9c64363bf7c60001b620015a0565b6200330a7fb75ad8e05510fe1ed6e7e20dc0a386002b1cd28947be51e1fc38667ee9d949a160001b620015a0565b6000620033588883815181106200334a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516200d33d565b9050620033887fcd2002e3dec244b5e4322e10bcb6257fde29577ea868e4dd6245449f9cdf27d360001b620015a0565b620033b67f84efe295e61d25dc20dabe94318bff15bda9b00690a4ebff604745cd06db81a560001b620015a0565b60006001811115620033f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816101e00151600181111562003430577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b146200349957620034647fbca3d29fa5b2281a996c894fe154e36f6ecba9491dd5bb46e9e915872d29696160001b620015a0565b620034927fa784c98d9bb91970f40fd7e3e55ec1da4a73af358c43af7465bfaefe077379ce60001b620015a0565b506200375c565b620034c77f5f049a6d20ee8954720ed052d48d10fdaeebb9be4ce662522fe2e273c68625da60001b620015a0565b620034f57fc328f9f4aadb4dfcfa0c933dcde1900deb4951fa9c98c89ce86eab49f1dca15860001b620015a0565b620035237f0693eeca19bea5f4cf8939e06670b9790c00644db5c2ac89a875bdd1f6dc3dd760001b620015a0565b806101000151861162003593576200355e7f56c2a32830c0bec428b1f3e93116bd115b6a10a18b37381f30865c12969dca4260001b620015a0565b6200358c7f721421755122ba597fe3e094786facd8dafa166f425d5711822029108862247860001b620015a0565b506200375c565b620035c17f7cb5379989cf75f03acf505a4e88c9ec793e698c9486805acbc3f3452c4e4ac960001b620015a0565b620035ef7f7b26c618f5595a902926dce4eba153c186e7299d75e3a66158b669ff0ace33c960001b620015a0565b6200361d7f217e0bbf81979bc8174959e73b12e34858f32bd7130f7c2c9c3f1f3d61081cc260001b620015a0565b60006200362c88888462005e80565b90506200365c7f169ad7ee9ba9c6d102fc4f9e85adac96f05101adc28ff38a5198834fdf51610d60001b620015a0565b6200368a7f084f4205d7c8296cb9d766978521980e1bd49132a0a76755f25a49f7b8318fea60001b620015a0565b806200372b57620036be7f0adfa40c30c7870865b0657f8dd46c9a734012328c182263552d2e1be584a55960001b620015a0565b620036ec7f052e70b433acdb52f3e35dc9babc521180c17cde8039644c72266f70b8512b6360001b620015a0565b6200371a7f1f0f5b6c23a32892d8d203b8dac8cb283bfe2410378642cdcf962100f351de3a60001b620015a0565b8360009550955050505050620037d8565b620037597fd04a05a7cef13603c2c665d1dc0bc684a446826863efc95b6b46acb8efafa76e60001b620015a0565b50505b80806200376990620144ed565b915050620032a4565b50620037a17f2a7f931bccf0c4219b8fb3268df5c8ea5f8cfcabe1725b7fb82fc67bf4f4463e60001b620015a0565b620037cf7f7e12c3dae6e9b1186c1d66c29c9ccb7a8dbfe77407e49d87a4084c6ba7607f0860001b620015a0565b80600192509250505b935093915050565b620037ea62010964565b620038187f447e082935b60124cf7d67a9e39b717c25b8941c25b81201e109122845824c0260001b620015a0565b620038467fbf63ee34c621f3a9272e67cd11cd0b6768c19765d9548255c8a9ccdd78440b6460001b620015a0565b620038747f63c9253e2607d9058069a3c09f2e1091e81594f4ae5eacb6de828e744da0425860001b620015a0565b620038a27f867a02572bd89c1dfd96049f4a189f54d37f73a254a81f9f491b3b8465e3be8b60001b620015a0565b6000826020015167ffffffffffffffff1611620038f6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620038ed9062013cbf565b60405180910390fd5b620039247fbaa62176caf78a3cba6c909e27dda9360cb993374322f1412476dd001f5ad70d60001b620015a0565b620039527f57a8ff50a61d952950fad07c2466d071a00684c355d3ec578c8f75414d46ca3b60001b620015a0565b620039807fd0b1a533f9b5ab69bc5e3f2d8b097a71e01cf8a819ae45ab5bb0d166b5f852e960001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015620039eb57600080fd5b505afa15801562003a00573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062003a26919062012819565b905062003a567f283c882f63869cff2b4c5701680a09d34b9776baaf3eea492db71b41c3d27be560001b620015a0565b62003a847feae3ca41b54342974c3cf68f821e34dc39f4c6d3358928868e6fc64cbd987bea60001b620015a0565b62003a918382436200c178565b915050919050565b62003ac77f0c6ae0a4e89937a95680b764edbdd4b873d70af1f8507d1e6f2b992f1296c2e360001b620015a0565b62003af57f28c5aced2023321b4ccfdee6c001705366f67b98ae0f90f96266205d80b4793d60001b620015a0565b62003b237f5d02a89c7e25afc1a50c75bae83ced379c1c52bd1c9040129708c29a01227c8e60001b620015a0565b600062003b537f95be1fe609c3d716f162de967a2470a9d39fcaab833d5b8f03a78d9fcf853a7e60001b620015a0565b62003b817f7c94a48499c9fe2708fbcbf5259341769177eb5ec12d75324b532b0fee70043360001b620015a0565b6000825167ffffffffffffffff81111562003bc5577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562003c0257816020015b62003bee62010a9d565b81526020019060019003908162003be45790505b50905062003c337fdd0206a010ef963d409282a886808b7b9381e7454be786adfb12fcd58a27f2b360001b620015a0565b62003c617f12c5ce26c452a7b83292a1cf9e94f651047444af3a2c471652bcdc370ea426a360001b620015a0565b60005b835181101562003e365762003c9c7ff8e0c00ea63238e947730512e66a045c88980cfc2a5697b52f76caccfd20201a60001b620015a0565b62003cca7f6b3db2a2a89b712e39dc8cd29bdd6018a1c67f75275d4331ee63c21a8da75c9060001b620015a0565b600062003d1885838151811062003d0a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516200d33d565b905062003d487f4ec6c0b28c10db44e5c2a084bf45824f83c55e87350afd7033d16140038ece6960001b620015a0565b62003d767fcf0e8737dc93a82d5257532020ee7a703e1e3ad256ad7968c440a2fc02aaaa7e60001b620015a0565b8060200151935062003dab7f3a15c5c0c506433295c2dafc65e14f508ab36d895cb94f48a4694d093273863d60001b620015a0565b62003dd97ff85fb9f7dd8a2bebdbe2e8c4cb5cd62aa685ab1164d7f6a3711fe088579d275c60001b620015a0565b8083838151811062003e14577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525050808062003e2d90620144ed565b91505062003c64565b5062003e657f057e2a5b0b9fd7dd1d9bac5efd407dabcb9fc0a7d5575ba9019419808e6e81ca60001b620015a0565b62003e937f9f1bc3830680a2f300bd0a956aa5afccc3c783b9519f5574d28be70fa29a7e1d60001b620015a0565b62003e9e8162008570565b62003ecc7f107b92c1cd97edad0fd90de45b6a8ff55ff2dcdad83c5527b493d30ffa2a9db660001b620015a0565b62003efa7f4304d515de4d365cec1c9add6533ad68533aaaffc20c543e0568226c7ab1f1ed60001b620015a0565b7fa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec600143858560405162003f32949392919062013aaa565b60405180910390a1505050565b62003f6d7f3e9991811ccf4429588f79283ed27119012f8983ad2ea50e40c9815621d7b1c360001b620015a0565b62003f9b7fae21f9f6f684141ff5d0e00271fdcf59775d4e662a63122b170169f8ea82304660001b620015a0565b62003fc97ff6d73a5c93d5376b2a0669dbc907714570bc06af6d6e2d7a659b73833e0e81bc60001b620015a0565b62003ff77f5b1e2d0fd003c0a913297c15e867afea02b03711635ec908206edd08decae66760001b620015a0565b6000600682600001516040516200400f9190620138f9565b9081526020016040518091039020600901541162004064576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200405b9062013c59565b60405180910390fd5b620040927fc232047193d89dd107fb5857b0f6ae860b7c043fa13331f950e16d8413ca01cc60001b620015a0565b620040c07f31b02e90378b75fe911716448a115407f943ca56d98b29d63f4ca1a53be7c7fb60001b620015a0565b620040ee7f1a27cff00c88e431b78b0d26982a3c41ddd60fd7eea6fb0acac22033e827c91c60001b620015a0565b6200411c7f981b9648ec972d2e47ea401e022795346b2128779d7384752120816e21e7665160001b620015a0565b60018081111562004156577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600682600001516040516200416c9190620138f9565b9081526020016040518091039020600a0160019054906101000a900460ff166001811115620041c4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1462004207576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620041fe9062013c9d565b60405180910390fd5b620042357fe499b64fe35656dc0a4c4e77a8983871300fbfee55c4945b3dfc583c06ced0fa60001b620015a0565b620042637f81d6e1ec87d8279ad412de3dccaa256078557bafb8a11a458f572bd480e4ac6360001b620015a0565b620042917f9e329ed053749ac56adb3c618eeed7dd79fb666e3fca83643d237a4ce3b33b2560001b620015a0565b620042bf7ff2bc6fc60153386fd8ce0fc6133387c918ec7d796b57257f82e8903235746c9f60001b620015a0565b4360068260000151604051620042d69190620138f9565b908152602001604051809103902060050154116200432b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620043229062013c15565b60405180910390fd5b620043597ffbbf85826141628e49de79f7f404deaee0dc7a767b333eae663b55a71c19ea3f60001b620015a0565b620043877f961b7992dcc5345cc732c81d7e18b2f0ddcba1af02b03585d06d2f19d75c3ed060001b620015a0565b620043b57fbb385af377f8540ddb2be261478f1f59fef205dc5b8e6f332e1f8e088954922260001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200442057600080fd5b505afa15801562004435573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200445b919062012819565b90506200448b7ffdb5bb71d145d5b3fc9d0e09ef10b30290a45e5c19b4e3f492b94f35d08b4c1860001b620015a0565b620044b97f921d46e74b139a78908d3ba9228a22dc320bd3db2846540ac802e86033f1ebc460001b620015a0565b600060068360000151604051620044d19190620138f9565b9081526020016040518091039020604051806102e0016040529081600082018054620044fd9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200452b9062014481565b80156200457c5780601f1062004550576101008083540402835291602001916200457c565b820191906000526020600020905b8154815290600101906020018083116200455e57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054620045ed9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200461b9062014481565b80156200466c5780601f1062004640576101008083540402835291602001916200466c565b820191906000526020600020905b8154815290600101906020018083116200464e57829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600782018054620047ef9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200481d9062014481565b80156200486e5780601f1062004842576101008083540402835291602001916200486e565b820191906000526020600020905b8154815290600101906020018083116200485057829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200491a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600181111562004953577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b820180548060200260200160405190810160405280929190818152602001828054801562004a1057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620049c5575b50505050508152602001600c820180548060200260200160405190810160405280929190818152602001828054801562004aa057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162004a55575b50505050508152602001600d8201805462004abb9062014481565b80601f016020809104026020016040519081016040528092919081815260200182805462004ae99062014481565b801562004b3a5780601f1062004b0e5761010080835404028352916020019162004b3a565b820191906000526020600020905b81548152906001019060200180831162004b1c57829003601f168201915b50505050508152602001600e820160009054906101000a900460ff16600281111562004b8f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281111562004bc8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600e820160019054906101000a900460ff16151515158152602001600f82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081525050905062004cc17ff95b5cc806e1fe96d109a790d392fe21d65b6471ea42fa98d2ec90996aaa702160001b620015a0565b62004cef7f61c3ec6b41aa5d892cc158332e078622b7675d7b28e5ec52d38631d3b490c8a560001b620015a0565b600062004d338385604001518461012001518560a00151866080015162004d1791906201420e565b8660c00151896040015162004d2d91906201420e565b6200bb73565b905062004d637f1b43559dd947cb5e13cc583585127d502f7dbeef1083fe15c04a902c60ac7f3660001b620015a0565b62004d917f44c50de07efcdc0fffbc404d9eb5aeb2ec073688d7d98dd78767305c1f4676d160001b620015a0565b60008160200151826040015162004da9919062014191565b905062004dd97f8d64a5094b58eb77946f12f86d32a680fc0026fac51a7ec4412539f426861c6f60001b620015a0565b62004e077f9bcadc2a0e6ceb48581e1001a638e8cc24c6274c639c91bc341f173338cf071760001b620015a0565b8067ffffffffffffffff1634101562004eb75762004e487f40ce558a527bb9faf205d2546431885ce64f6e7917e55575a3c61a38071abc5c60001b620015a0565b62004e767fb6be8962502d55a9414f78d2d0b138d4d5d13bd187a08304b744e7c5049510a960001b620015a0565b34816040517fb0b78f4900000000000000000000000000000000000000000000000000000000815260040162004eae92919062013e2a565b60405180910390fd5b62004ee57fcb2004a05b6ae580e073350a258ab5e7655298e0517afc044fd98a72dda2712360001b620015a0565b62004f137f44d5e14b343b985dcaa77ac740c1d4e494392de3d8e179b73fcbff884e0d89d160001b620015a0565b62004f417f5eb9fe692af152eeaeae6cbcffe67da91d08ae8572acdf60b88869dcf5a19a1c60001b620015a0565b84604001518360e00181815162004f59919062014191565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062004fa47fe66b23ecc532fe479a1fa2d9d96b13f9ba1d0374b9a8b19e1d4f73369a895a8e60001b620015a0565b62004fd27f7c795242005044e557eceb72bbe6d71a4861e15b7ae779e347f62e7cd31109bd60001b620015a0565b80836101400181815162004fe7919062014191565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050620050327f3f1d625862e8616925d4bc4a13bcf40dd24712f7b4fc4d3b3d9c315152cd5dce60001b620015a0565b620050607f42e0b5c64f3aa706d678eb71d6e13564fab06e1bec72830f91450c133eedaa5660001b620015a0565b84604001518360c001516200507691906201420e565b67ffffffffffffffff16836101000181815162005094919062014134565b91508181525050620050c97fa2999d9ddd7087cb98b36a781b16a450a2ba2d008477701bd368fe8d1fe547fe60001b620015a0565b620050f77f5ad31e9a4cd2ffb5f7e0db63b08f5ced4005b01145383a781bae521849f402f060001b620015a0565b82600686600001516040516200510e9190620138f9565b908152602001604051809103902060008201518160000190805190602001906200513a929190620109e9565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190620051a0929190620109e9565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061010082015181600501556101208201518160060160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160060160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061016082015181600701908051906020019062005316929190620109e9565b506101808201518160080160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600901556101c082015181600a0160006101000a81548160ff0219169083151502179055506101e082015181600a0160016101000a81548160ff02191690836001811115620053c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061020082015181600a0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061022082015181600b0190805190602001906200541a92919062010c39565b5061024082015181600c0190805190602001906200543a92919062010c39565b5061026082015181600d0190805190602001906200545a929190620109e9565b5061028082015181600e0160006101000a81548160ff02191690836002811115620054ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055506102a082015181600e0160016101000a81548160ff0219169083151502179055506102c082015181600f0160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050509050505050505050565b620055a57f1f94138e9a33b18d4f1458abef85f3863f456e5b1aa7e26077ca85085bbc06a160001b620015a0565b620055d37f772e38d3099bf19a8dfb845f612c923be2eabb9a7f2dad03ba2d1e4be13dcfd760001b620015a0565b620056017fb860fb88daf96cf03b146758004ee725ec265e0dd7001c60be8b99d7f4c5f18a60001b620015a0565b60005b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050811015620058b7576200567d7f4de2d54671678c15be153bb8d329c0fd442b46834bfe03fe56cc62b0d034aeef60001b620015a0565b620056ab7ff4002df7aae4444c68150d646ab3ff0f9a2e3b3d804623f8cff010e6538e62ca60001b620015a0565b8180519060200120600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481106200572b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160405162005744919062013912565b604051809103902014156200587357620057817f80e0526bfd7cfda7d7022153ac33fdb4dff324cefc7644e68e175be57becf66a60001b620015a0565b620057af7f41225f43f2382c9f83c85a90f86600e9857936fbf9727ac048c70bf399dad05460001b620015a0565b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020818154811062005827577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006200583f9190620109a3565b6200586d7f619b9e0390f34a875722793ad5ed0917e4db3a5ff2b1f7e5ded979393feb53ba60001b620015a0565b620058b7565b620058a17f1dbab3312bb21e744593709286852d193ffebea4ad8265edfa2852903cafcb2360001b620015a0565b8080620058ae90620144ed565b91505062005604565b505050565b620058ea7f8c60b50f8cb941e11ef012eac6e238c8449dcdee55ee332b37549c53db1c2f5f60001b620015a0565b620059187fda3bc62f49fc366348d1100ba32e8b2e88684b40624f7332de0565f35eb614ca60001b620015a0565b620059467f3fd428942c87989e6a3472b7764e3ddc97e2afa1d79c9a744f36a16245c3de1b60001b620015a0565b80600682600001516040516200595d9190620138f9565b9081526020016040518091039020600082015181600001908051906020019062005989929190620109e9565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190620059ef929190620109e9565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061010082015181600501556101208201518160060160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160060160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061016082015181600701908051906020019062005b65929190620109e9565b506101808201518160080160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600901556101c082015181600a0160006101000a81548160ff0219169083151502179055506101e082015181600a0160016101000a81548160ff0219169083600181111562005c15577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061020082015181600a0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061022082015181600b01908051906020019062005c6992919062010c39565b5061024082015181600c01908051906020019062005c8992919062010c39565b5061026082015181600d01908051906020019062005ca9929190620109e9565b5061028082015181600e0160006101000a81548160ff0219169083600281111562005cfd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055506102a082015181600e0160016101000a81548160ff0219169083151502179055506102c082015181600f0160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505090505050565b600062005df27fb195c33d923490965474ed61da2086b01cea0bff8ec52609939c3283785965e560001b620015a0565b62005e207f3d36a9f405bc0ce9ea63a7a293f5191fb1fd905b40caa99ca7c960d6cec54c4360001b620015a0565b62005e4e7f4c93587fe3303c8f4b9ce7cb659791580e0bb9c4087f3120911a1965381ba62d60001b620015a0565b62005e5b8584846200de9d565b60018562005e6a919062014191565b62005e7691906201420e565b9050949350505050565b600062005eb07ff6b7b5fae9dedda0f487a6d9b50a912b42669f8ebde0237969a23c7dc148a40d60001b620015a0565b62005ede7fcd0f23b0fdd0c6c6f86891931674f7e5253a4cc6f22d237e348124e16b82963060001b620015a0565b62005f0c7f73a73f34f582d6c37316641eb60b79129fa0301a8de5329d20bcd22c48e57e5860001b620015a0565b828261010001818152505062005f457f1502a0f78e02d97146ad524a6ad70324b9c1cf46c49132b1ccff31981ee7322b60001b620015a0565b62005f737f27998aa87b098b76afe9d84f1f1f84f6ec28f705683a78064aa9d97467fd6bd560001b620015a0565b62005f7d62010cc8565b62005fab7f4391215a194a1e339ed3d899f0237e4e25e5567cd7267d60bcaeb9480c44d94460001b620015a0565b62005fd97f902ed3e3ffe158a6577b5a678ce1f61d7cc7aee32a7a04f190f5e1926f1de72b60001b620015a0565b826101000151816080018181525050620060167f9402a71f36be7735c925ab19c4b222d6cf4c5362136f215dbe81585132b4a10260001b620015a0565b620060447f6cca6dbb597602c989f9385928265941dd0fdeb19901a94984091d1027d132ee60001b620015a0565b8260c00151816040019067ffffffffffffffff16908167ffffffffffffffff1681525050620060967f9d61adcb7d8de38326080451d6b6062c3e9883bd703eef940123b6328178508960001b620015a0565b620060c47fb85fa0c2ba1f3dfc45f9ac728996b9b5115769cc119e59d06cbb7448bde4705560001b620015a0565b8261012001518160c0019067ffffffffffffffff16908167ffffffffffffffff1681525050620061177fa8f23140eeff226f2fa651cecb4bf71f4bd4eeb04377777ee4f07b23206efcec60001b620015a0565b620061457f17d4ce84ab38db9751999ef5804c7d9880a7b4f8f6844a0678dac15ec8ad5d8660001b620015a0565b82608001518360a001516200615b91906201420e565b816020019067ffffffffffffffff16908167ffffffffffffffff1681525050620061a87f2fa0fa6bdd49e88c619efa12bac96ffea381dc91c102493b972439716953cbb760001b620015a0565b620061d67fd56a2cd8bcddae1bf3be685feca29b6994c2f94751d90c5a45fa58a1258c906860001b620015a0565b6000836101a0015190506200620e7f019c0f8331752afe2515b71d3e6cd862f51d400c32358d8729baa76c9c858c7a60001b620015a0565b6200623c7f4def059c2fd8442a85b5a8206516cb80bf8c283fa15f36166ccaaa1431c312e660001b620015a0565b60006200624b83888462001661565b90506200627b7f5026f785df6e533659c03b246091f1feb2ebd386d2415d220829115bc16b1d1660001b620015a0565b620062a97f0c1d72ae9f3a74452e2c8d3f4f17594377e401afb2fe7ba1f6a5c738bce8ab8160001b620015a0565b6000816040015182602001518360000151620062c6919062014191565b620062d2919062014191565b9050620063027f76b71838af0234fcda6c36fb8da83e4b05c436bcc25e4141ae23db9e5e49f70f60001b620015a0565b620063307fc12e84a892cbb1acff726b76b238bc62b3a1f80a1f84289b44854662b835b3f660001b620015a0565b85610140015167ffffffffffffffff168167ffffffffffffffff1611620063e9576200637f7ff7f52db577dfcc483aa2ca8ce9f28ebec6bb73fbef4a477c06460d3a3265dd8760001b620015a0565b620063ad7f9a88d05f70bf6088172eabb953cfb548175ec95a760357fae58ee72638477ad360001b620015a0565b620063db7f64268433e8261a1e559e661e622d7f9e985588f27c64be78866cf8571392a56560001b620015a0565b600094505050505062006580565b620064177f2f15d3deb42c3213a46690100d5f689ea3c412ed686c960e24d714f9e3e1f50960001b620015a0565b620064457fc44bb44de2474117bb8f10ef00a8254d19c5a4b8e5bd96ade6054c543b7820d460001b620015a0565b620064737f0ed23c6776c4fc8914ad94a8384c85fdb666f2e39e7f2b7f874e1b7b9e59ffd660001b620015a0565b8086610140019067ffffffffffffffff16908167ffffffffffffffff1681525050620064c27f09b906236dba6b1858c675081a6b2ed2c160b4d711fc31930382ea9dbe9004e860001b620015a0565b620064f07f5fa72ba24b0f066106b91e31054526a2e6a2fa8f8814932f497573d90de6f16060001b620015a0565b620064fc84846200df61565b8660e0019067ffffffffffffffff16908167ffffffffffffffff1681525050620065497f6cd79e0e068846b37280ec647c6654c5c31c05dfe22638c1886841e3f9433da260001b620015a0565b620065777f8927dbfa57261b5e38bcc7244c769af0e5a5a1d7827cfcdf3cbb857a5275a28a60001b620015a0565b60019450505050505b9392505050565b6000620065b77fc3f82c4d6c9ab658616bf090fa5f016979e0dc644a6c91614be4a5f4fdfb021460001b620015a0565b620065e57f11f58a6b53e0398335c024a1e1b668ef5f10bc496bbbebd5cb5cdd5af3234b3560001b620015a0565b620066137fe74735d3ccb835b9096ea2f5e4307ae72fd5de874483b9a1b1b306713ae7563760001b620015a0565b6200661f84836200babd565b836200662c91906201420e565b90509392505050565b6060620066657f388baf0212ca8889cebfc57d28ea7c190d1b5b0b04ac62b3adb10417f30d2dc660001b620015a0565b620066937f53f0fb8a16094b39ea00fd8ace536e378e685671fbd9dfd5b0b8eb5687912a6960001b620015a0565b620066c17f7746d3145f2d20f9e2c98fdb18b0827008f5771f206d363ea6dc0bee54142af560001b620015a0565b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015620067d8578382906000526020600020018054620067449062014481565b80601f0160208091040260200160405190810160405280929190818152602001828054620067729062014481565b8015620067c35780601f106200679757610100808354040283529160200191620067c3565b820191906000526020600020905b815481529060010190602001808311620067a557829003601f168201915b50505050508152602001906001019062006722565b5050505090506200680c7f79b8ae175be0a6c5474e1da01b6c33146f8316ad20e6d4629cc21e65eb4dfc5b60001b620015a0565b6200683a7fd5ea27528f4bcab085eb2c96b80c4e9a57be31cbaa88865713be71bd9da6b02660001b620015a0565b6000815167ffffffffffffffff8111156200687e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015620068b357816020015b60608152602001906001900390816200689d5790505b509050620068e47f5661b06e75adb6526474e1b3ff2b51492410037aaa8e5be9d62592e2c1e70e8060001b620015a0565b620069127f811ca069aee492ceebd10e4bcddf4cb095aa3f7b05827a5ad920702b1e249ab760001b620015a0565b6000620069427f7b23d086bdb40b89524c7ebbf6efdcbeb1ad51138ef78ff843bf3cebe044b4f260001b620015a0565b620069707f282a2e203bf20e4b12d64adac9b63ec9eef99d7e7013eb7a117eb7384fbfac2360001b620015a0565b60005b835181101562006fb057620069ab7f7ded9006d32f67ec194c72a67bdc3091afefe3352479e2332220f68d140bddf860001b620015a0565b620069d97fc0ad66a5be81cb1f141ab5aeb52b146f7163c726c4929b8a766e14335e4c9a3f60001b620015a0565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe286848151811062006a53577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040162006a79919062013a52565b60006040518083038186803b15801562006a9257600080fd5b505afa15801562006aa7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062006ad29190620124d8565b905062006b027feffd430b4ee804f8cd2c2340b0c4e9ab454ad200d9eb1cc1fc941dc200ccd21e60001b620015a0565b62006b307f507cc4bb977412c31099b28cad2199c8537b4ba2041befb2caad5bf1cf92f0b360001b620015a0565b600062006b607f4ed64843a3ca1c5d48a23d6a2cc8437a307ab09e430ec152b7eb806ca26c34b960001b620015a0565b62006b8e7fe259954f4165993c53f66e49a4c5658bb56b75858681bd750b6e8fbfabf007e360001b620015a0565b60005b825181101562006d775762006bc97f504625989b9ec2c61960b1b1f38cec27b159d6813d95835e2cdc773142a5692d60001b620015a0565b62006bf77feaee30452ea629889d22a584c8c806c8ea340ce9c9182975e48c1e2dd98d29f660001b620015a0565b8873ffffffffffffffffffffffffffffffffffffffff1683828151811062006c48577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141562006d335762006c9f7fbdfcf26de3561f1e8b68867cf2760b2760aaff774478a7d39796d163480a8e6460001b620015a0565b62006ccd7f9f972284d511d7c3ee430b88dc4d5e4a45185740880b7a0504940527fd8b407a60001b620015a0565b62006cfb7fe2c77ed87cedf8dddffd03196a9f7a22f4ee58367060e9c2158b55a794fe0c8a60001b620015a0565b6001915062006d2d7f460a25f2acad0469abaad3d90531af060c3730ede7fd8581d157ef4e4c7e394060001b620015a0565b62006d77565b62006d617f10189c223c7807a0b5ca34fd36833fd7f5a799a00731b4a1157e966326a81e9560001b620015a0565b808062006d6e90620144ed565b91505062006b91565b5062006da67f8c378dd10f86bbd1ccc62258965490dba24c8f9120ae381082c54bfc025e7a5960001b620015a0565b62006dd47f4a9751caf88e39c9ad5ed65d96db81398384f9a37bc6024301698dde0ff6a4f660001b620015a0565b8062006e3e5762006e087f42ef728a33b1349b163f6a139ee1a3b6e550d25b64da49e05f25f4fb36e6b2cb60001b620015a0565b62006e367fab338a37a9a4ed52f8a714dd7943b321c5bd474d1a46624cd99f8cd00d0648cc60001b620015a0565b505062006f9a565b62006e6c7fb9da73a04beb3ff88caebedd89002391d5935a0ac75b16810bebc819957e5b1b60001b620015a0565b62006e9a7f240e9da687309d02d789b76bab78e360733d6de668348f937744fffcc3393c1d60001b620015a0565b62006ec87f702bb253c15ca8c6e28e245e9635668b33ff78d35698bd73c3c34197f46a480b60001b620015a0565b85838151811062006f02577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858567ffffffffffffffff168151811062006f4e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525062006f877ff531e3c8787290c8678a3e2f0f8bbd86077fa896fe4bb5bd750189c3ec4f527e60001b620015a0565b838062006f94906201453b565b94505050505b808062006fa790620144ed565b91505062006973565b5062006fdf7f056756053e5182fd9bfaa62c53a52f52fbec57565d531b067d34de72a6be260f60001b620015a0565b6200700d7fdb01789ab99ab24e5787b38f325d35c5e01c49fd8c2ebf7a7cb4fb4d535bb3f060001b620015a0565b819350505050919050565b620070467f87b4148840511a558d9520aa269ad194c3befdd158a8313e31f245d5d654391d60001b620015a0565b620070747f2c0dc343a7de0fd692ee77940655c221a4f426e744f119b36e5ee7300ebb5db160001b620015a0565b600681604051620070869190620138f9565b908152602001604051809103902060008082016000620070a79190620109a3565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600282016000620070e09190620109a3565b6003820160006101000a81549067ffffffffffffffff02191690556003820160086101000a81549067ffffffffffffffff02191690556003820160106101000a81549067ffffffffffffffff02191690556003820160186101000a81549067ffffffffffffffff02191690556004820160006101000a81549067ffffffffffffffff021916905560058201600090556006820160006101000a81549067ffffffffffffffff02191690556006820160086101000a81549067ffffffffffffffff0219169055600782016000620071b79190620109a3565b6008820160006101000a81549067ffffffffffffffff02191690556009820160009055600a820160006101000a81549060ff0219169055600a820160016101000a81549060ff0219169055600a820160026101000a81549067ffffffffffffffff0219169055600b820160006200722f919062010a7a565b600c8201600062007241919062010a7a565b600d82016000620072539190620109a3565b600e820160006101000a81549060ff0219169055600e820160016101000a81549060ff0219169055600f8201600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556000820160106101000a81549067ffffffffffffffff02191690555050505050565b6060620073087f62a02aba0a01aff589b3934d3ff3b584bf6585f0eae0e29ce27520a14d44df2a60001b620015a0565b620073367f390e2f304bae752d8cc90c59ecea01be9e709d948682ffea2a7a20fd069cb8e960001b620015a0565b620073647fa561962b7a6dabdccf63b8a3ce8f86b536e8dd05d5ac95390281f77a2974d40460001b620015a0565b600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101562007479578382906000526020600020018054620073e59062014481565b80601f0160208091040260200160405190810160405280929190818152602001828054620074139062014481565b8015620074645780601f10620074385761010080835404028352916020019162007464565b820191906000526020600020905b8154815290600101906020018083116200744657829003601f168201915b505050505081526020019060010190620073c3565b505050509050919050565b620074b27f958eed56fdafd1dee1feb6a2adce329ed83a818ca82d4f3e9951dfdb19a26a6560001b620015a0565b620074e07fdb124247ff4c649c0c8e2b6341cf83d9c2d5addd8d00486e04962b0020e3bdb660001b620015a0565b6200750e7fceb0d9c67ce42d504cfe5bf5217887a65b6b815d8a0acac7b22ff3626d74343260001b620015a0565b600083600001519050620075457f3b10d4392a724f64d01ab62981bdbfebdaa6a138cabfb454527414f120c2a95360001b620015a0565b620075737f17df30f62f2c7c62e98a332cba855584657331a50921289a498bfe0bf42543e460001b620015a0565b8215620076e757620075a87f9d71ca31f8c70a1030a4a2562436b86b36a622113b961f401bd71adbafe8120e60001b620015a0565b620075d67faa9c2eb2b7201248cc0c653e0a511cda314395fd9d5e868599b9f5505f24463560001b620015a0565b620076037e8265f59f47a982857c668d7b4a65efdd282296c2bd2367f9740500f71087d060001b620015a0565b6200760e8162007018565b6200763c7f5d8428453f6e92ce2fad737e28c81e6d08e388ec07ace3fae69694d869221d4360001b620015a0565b6200766a7f5eb30bafd704c7b6a22c401822cddc09f8209b2306c1d2104694f81fc6a07fec60001b620015a0565b620076758162002e08565b620076a37fb8df03942981319cbb18aed20e2981ff1fa0ca501144cd96e461b892ffa5844260001b620015a0565b620076d17fce544cae4e90aaef5b2d9488288fc12c5e5982f87db30c702e6620702291e04660001b620015a0565b620076e1846020015182620080a4565b62007716565b620077157fec78cf7e8e91585096452345627bb71b1045af1833a57f2dea84eb6d2023051c60001b620015a0565b5b620077447fe2a0196d1f505023ee1ef74394e4f21479eaa79624eaf3db739dcf816d672ef360001b620015a0565b620077727f9d8b1d6cd7b1a382665c171f128ddb755628372a260a3d6b5ac018eab5c1d4a860001b620015a0565b811562007a7f57620077a77fe8a61fe8d78f357872bfb46d8eaf2e075f221c61c8e53f73ee930fe745defea060001b620015a0565b620077d57f28e115199828f8b23b7093ef0f6a139054118a674b70085f61833f1fb777242060001b620015a0565b620078037fb6fb671f0fd736857a0ea8f531bd27f125f7fcfadda6675becfa9878a202779160001b620015a0565b620078138460200151826200e026565b620078417f62a554a5e8c99d561be3d36a05e1dd5451322f31fbfb1ba62304cf6f3962777d60001b620015a0565b6200786f7f76f219ddc649b089e4ea142cbddff8e18a04c5710858b9cfa435cfe818426e7e60001b620015a0565b60005b846102200151518110156200794557620078af7f2d4a186ef5dcd1948396f16a804c8d6057042635b1e076bec95154062962b48360001b620015a0565b620078dd7f46abce34920238dc89a48186b9f06e80272b0a6b9347b2edcf09be302aeed18560001b620015a0565b6200792f856102200151828151811062007920577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518362005577565b80806200793c90620144ed565b91505062007872565b50620079747f1eb67b138d643afe5d08525cfd5ac3f8844274ddc973c2b9547bbfa7657acf3660001b620015a0565b620079a27f816b5edcb116f5cf4b7f89b2a5d7a082708ad20ec82eedfdd49f81380e1eeb8760001b620015a0565b60005b8461024001515181101562007a7857620079e27f33b9fa0f7a541df7837fad43c95bb632fe7df21fe07b9ffd62f9a039cdbdc58d60001b620015a0565b62007a107f41f5c625df63eb1fd41cee4e28869378bfb881add1bd86e3aced521570514f0160001b620015a0565b62007a62856102400151828151811062007a53577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015183620017f5565b808062007a6f90620144ed565b915050620079a5565b5062007aae565b62007aad7fec59598a07ea9f6504111408edacad04477a537e83fce736aad95f2d8e8dd17660001b620015a0565b5b50505050565b600060019054906101000a900460ff1662007ade5760008054906101000a900460ff161562007ae9565b62007ae8620108d7565b5b62007b2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162007b229062013c7b565b60405180910390fd5b60008060019054906101000a900460ff16159050801562007b7c576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b62007baa7f57879634bd2dfa3ddbf06b3b735c29ad8e51016bb0c4f5d407ba7a411c0ef36660001b620015a0565b62007bd87faef3e50a0b73da2aafc5f30a7b84e5a51881148024382c3586930fa1611d05dd60001b620015a0565b62007c067fed3311f93163521dac2206de7278dfb3d29d54811c042eacbe492691aa58815760001b620015a0565b86600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062007c757f8cda38560b56ef3acd37f866324212bd50f30886763063555be10aa0659f8cac60001b620015a0565b62007ca37fcaf13f8d7d3adcfa35498a36d395538f7986f622c21f900dc9f6f4e3cd2da8a060001b620015a0565b85600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062007d127f9f485013c6bd038643674c2e83960d78f4624f205ad8fcca46880105f03e8fc360001b620015a0565b62007d407f24c038856a123a351ec5c98bb117cb09149daa55e640031b1d8b1358e9f4886a60001b620015a0565b84600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062007daf7fa97789f41db57f357f3488ebfc68d20349da283e89335fc1dae994fcecaaabbd60001b620015a0565b62007ddd7f167acaa778b52b0c023c82d706a1695195011dc22ddefda802d6bf2cb33212b660001b620015a0565b83600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062007e4c7f5bc409bf112e50c0b8106dcfb85c8768859033ffca97c18d78ca6d0966273a8b60001b620015a0565b62007e7a7fcd544d73c3ff60a21e9c799bf6a160189eab5aadd68094969cf2f9e33dbe729460001b620015a0565b82600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062007ee97fd921e2e99ea30ead8738c6fd252327de7856ba808e8226704f367a512d439b2560001b620015a0565b62007f177fc52f52cda96c6738011ea1e110849683642e743dfc9c061054569bbd3326459e60001b620015a0565b8160000151600460146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555062007f727fc1d97d74fe6595ec7ba319d5afcdeb58beefd65fc1591ddc83ef18e054516a8960001b620015a0565b62007fa07f22cd66f2d2611ad157d487a561d1e87ab1d61f7f6b81d8cedd896c27ddd4f9cd60001b620015a0565b600460149054906101000a900467ffffffffffffffff16826020015162007fc89190620141d6565b600560006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506200801e7f14c0a98e3670b109cff339d59fb3b486fced2394053e64ae3b25fe9c247ba72f60001b620015a0565b6200804c7f890a7f6f3711fdc59755177d27724d66706a2c2412a841d3d4a1d1a1a8313ded60001b620015a0565b8160400151600560086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080156200809b5760008060016101000a81548160ff0219169083151502179055505b50505050505050565b620080d27f60f4ad7a32a98583407569415f6bb676c8e5600b3263cec10bdecf3b66394a8560001b620015a0565b620081007fc8319b01485dac020e8b5b51bb2d5db722b7cd3b93ddebb01d1ca09f852af6e260001b620015a0565b6200812e7fb7a256da74bfb28af42e2f39baff2c27316e7d840a84f93d210ca225980d829660001b620015a0565b60005b600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050811015620083e357620081a97e6e027161f2f321aafb51a5025906fec65d3f71c010dcc7a7fec027a0ec5ce360001b620015a0565b620081d77f85f4a5bd4455ed052654f8e0324e0ec109507a45d44dff13dce52183f2a34bfd60001b620015a0565b8180519060200120600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020828154811062008257577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160405162008270919062013912565b604051809103902014156200839f57620082ad7fb3da975fd0934fc57a83f19c98f177026fd935cad5aeb530a5be48fafcbf50f360001b620015a0565b620082db7fb2835011e520bff0614d05b41f22229df6c8c61385e5d36f53d1ccdd6e7367f560001b620015a0565b600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020818154811062008353577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006200836b9190620109a3565b620083997fd27e87e69fd5b31e928274138ea3dea87dac3c5e1b0069a957b45dd9148d7f3560001b620015a0565b620083e3565b620083cd7fef72770d7d84b23281a09f97cb36d2ddf24ed749080be126f8fbe4c28ffafbf860001b620015a0565b8080620083da90620144ed565b91505062008131565b505050565b620084167f99aeff814146c54ea04bbe440a73f62f7199f4c3cf83831eb3a649e223b280a760001b620015a0565b620084447fdbe9861e8e5762841e618df500e933e87ee1212151579015c41674ec784d5ec260001b620015a0565b620084727f122c6ff6b646163db8db3492449180ae21dea711655f77250de30bed8d03817260001b620015a0565b60006200848382600001516200d33d565b9050620084b37f4298f47ad90cf278fc0374a16b4330b66bc9212a79aff3673af3946f3b116d4a60001b620015a0565b620084e17f7079fc557e49f9573c55f0c78a14cb3a3a05e63227cc2dfad9e267c01a570c1c60001b620015a0565b8160200151816060019067ffffffffffffffff16908167ffffffffffffffff1681525050620085337fb51d7e6f39d45df623a1f53c6b914bdee6665563ec43712d7e49e4ddd6990b1e60001b620015a0565b620085617faedba580f299a7c1244118fed3a9c7387a4ccd12a95d3d5dedc3a24ac462048260001b620015a0565b6200856c81620058bc565b5050565b6200859e7ffec1e0c92f64247167617933cf4a94ca73cbf5b53834d04a45702275e5b0871360001b620015a0565b620085cc7f5c6f83f77a69331b91f32995d0bf3d5251f62018371d1ba82f515d8d675939a460001b620015a0565b620085fa7f7edafe165a55d8bf1b374f81e56317170cedfd8a504706dbad1c13306b7ef8e360001b620015a0565b6000815114156200869557620086337f5b742bf846744189dcb20662dbe034cf3a8d728e1bebc1aa4635b3818d86cacd60001b620015a0565b620086617fb2dc7b267e1f7052516ddd142ad721c4243b7d5f1de11b4630048c403649b20860001b620015a0565b6200868f7f0d5b81b5145e3b5cf0435caa8a7baca595268cf341f6e948ce5d4f9d206b662e60001b620015a0565b6200a549565b620086c37fa07972223aac44602f2ede0469abb68f6ead68e45749438ca80ab54634e2ab1660001b620015a0565b620086f17f1536f27a76046b679e62b09682fd42873a57b3ec725392b790d0aad981b9a6fe60001b620015a0565b6200871f7f8abf294a065c517380985b36111f22fe635ace0e880f4f661c3e212861c4fae160001b620015a0565b60006200874f7f9f2ef5f0877c3583bcac870c298bd2f0e145eea4cb257b67f75bc6da6678006b60001b620015a0565b6200877d7f77f7c945626170bcad0de9afe11f86c83f8f3c6a2fee463d001f7fcab134fda360001b620015a0565b600082600081518110620087ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001519050620087f67fcaf1dec1d4ab98ea29ce77c4ae043c39e777fda52a8c3bf0d1a57e827788766260001b620015a0565b620088247f02500afe9f3ad8d5c3c4b6209a4d49c4100aa07cdccac6fbfe5d1a2e6596a2e860001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200888f57600080fd5b505afa158015620088a4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620088ca919062012819565b9050620088fa7fbe491b7a7084baf0e8f3fd530436ec5f73c6aa86e0d755e67ee6d9f44dd3dfd360001b620015a0565b620089287feb0fdb6bfacd32e721fe398f7b50939d1bc9d7bc0b39509df95a44edf061d64160001b620015a0565b60005b845181101562008b3e57620089637f833e2ba4c35cc0d040958d790b057665c8ecfd7531557b1e0ff14c4c107a177c60001b620015a0565b620089917fc8b18c248d4b62c6dd2bfb55daafb0f55d643a91181c89abfc4aa054ddccd26060001b620015a0565b6000858281518110620089cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151905062008a057f596fe214bc38b0aaa0c3efd2f9597f9d4c029b7e56047111d9d9ce4b0274983b60001b620015a0565b62008a337fd798db5a967634c6fc6ab9ee01f8352a103d773df45b89a93bc2858ba70af5b560001b620015a0565b8373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161462008af95762008a997f722dfd020b96eeda4d628c96c4b94069620005784e5417b5851ef0eb53e00d7460001b620015a0565b62008ac77f3bd61223627eb34fae9620dfdd0947aeb721f184dd7347b4a4b9b7605291e07a60001b620015a0565b6040517f2f7fecf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62008b277fe4323601c35961f2934c5742906856954f7f1b292fad95fe5e7b35ec06c2539760001b620015a0565b50808062008b3590620144ed565b9150506200892b565b5062008b6d7f5368e6403ccbf59553954750121332161cd900abc20ebc68bf7f282756701d8d60001b620015a0565b62008b9b7f8d1fcacfdfee8d3dad1f0988fb749df24eeeae47125dc2450a00258268da98f960001b620015a0565b60005b84518110156200a3655762008bd67f7582c1f7d7cb8ef97095497fa459be80917aa3027590111019662f86a16d387660001b620015a0565b62008c047f2a9e8e7dc8d63d7c23a030724c07bbdf2e1e85c49b070d15dccd15266b5b3ac460001b620015a0565b600085828151811062008c40577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151905062008c787f42743ef157b72ec28f1c725160ce66abf49d43b22267404dd90227a1d7c76fb460001b620015a0565b62008ca67f26d8c049d8338f867bf194bd1530f18e59807f6a63444dc832b0130a6ebf4b0760001b620015a0565b60006007826000015160405162008cbe9190620138f9565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101562008da4578382906000526020600020016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001906001019062008cec565b50505050905062008dd87ff494f922483f335cfaecd0760b6928a97501002ac99c80f756704e0e7f3095a460001b620015a0565b62008e067f50496bc33f510b3fb9d2750123a6fcf7550fb0cf37bd2c044a139f55f8778cfe60001b620015a0565b60005b81518110156200906f5762008e417f63c930a7cf17a38b073c96d64be7e3a7d49514103c9506208f2bb5137d49afc760001b620015a0565b62008e6f7f6e38fe8aca7e8bcb1d5536358ed018563172e75ecb575372a4fd0a5cd4954c2e60001b620015a0565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d784848151811062008ee9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040162008f0f919062013dc3565b60006040518083038186803b15801562008f2857600080fd5b505afa15801562008f3d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062008f689190620127d4565b905062008f987fc9ff4220484d3f9458a0d2e3b462a02f90036b0fa557b8e0bd54f58cbcc925b560001b620015a0565b62008fc67fa568c33573a18ff02b069e40d62f69c830bdb31120b07120b751ef32826c3ebb60001b620015a0565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c00382866040518363ffffffff1660e01b81526004016200902492919062013d88565b600060405180830381600087803b1580156200903f57600080fd5b505af115801562009054573d6000803e3d6000fd5b505050505080806200906690620144ed565b91505062008e09565b506200909e7f46a5877324f3c7c7bb5f268a888d9421406e799fd8e8c9c94ff6da06e53245dc60001b620015a0565b620090cc7f2051d4eb957f2a75ad63acba42ae5da2926d1c05e7895ab94db6358bbf65883460001b620015a0565b600082610140015167ffffffffffffffff161415620091b357620091137f9161c85b0f093c4bd1c157308e36dc2d38b5484542091b21fc63c27fa9295f7b60001b620015a0565b620091417f3e3a6be7ba1cb650e05c00bd4d4b10d76de373c5131abe671f56e6bbe2fa5cc060001b620015a0565b6200916f7fbadb00f25ee0bb872631c63a38cfb86c8b1aa13ce8104dc1a45c90dcac98718660001b620015a0565b6200917d8260018062007484565b620091ab7f5a1122b1559a58cc549474032921bd01c7f8fb6555d11e12a6ef5b78d90b958560001b620015a0565b50506200a34f565b620091e17fc10acb0d43a791642286c564333f99bea4a316c6be59468fad2b97f1858add7f60001b620015a0565b6200920f7f5b91100d35730d278baea1728df98ad9e365875f6da418c8f7ab092a6334a96760001b620015a0565b6200923d7f4cbf12a30e116d25102eac7fddbb8653c968ba354501a37147c593fbdc7246d760001b620015a0565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe284600001516040518263ffffffff1660e01b8152600401620092a0919062013a52565b60006040518083038186803b158015620092b957600080fd5b505afa158015620092ce573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620092f99190620124d8565b9050620093297ffff9218c5caa959e92bf145df8d38d706bd35cd3fb9d2585604445e60164923260001b620015a0565b620093577f33d375f9d42d7c701b3bf432bd06f131c29e994a68b158f2834d631a0ac458dd60001b620015a0565b600083610140015190506200938f7f040afd990879fcdc42ee17775af451ddfdce4af252de00bba93191b9be85db9960001b620015a0565b620093bd7fe6b1a7f1e1c3c164dab0fafec13071d7e6645a8704751cff4a9e392cef4e483860001b620015a0565b60008460a001518560800151620093d591906201420e565b9050620094057ffe31150f992673f9d3d8379febc9cda01f9d5419f1d5bd911a2aae867d6dd5af60001b620015a0565b620094337f8d6e19b5c22b49b24cdc6f2dbffaa98566f654c03a7227e523d3e6ca4ba3257d60001b620015a0565b60006200944188836200babd565b9050620094717f3fda9983dcc516dc5352523d698059269843831a8d3f69867085ea802b850e4860001b620015a0565b6200949f7f4e971878e934e38074b23682d9a1de4183b86e8eee0ab9295a67f66349ac1dc760001b620015a0565b60005b845181101562009ae557620094da7f7c1959527d9fc9c7e0255520d47c16d9585faa3b91d61deed26eeb00090a46d460001b620015a0565b620095087f6cf24c1ef949a958d9c8493bd221d5233402e6b30be41f6619f84941892b8efb60001b620015a0565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a87848151811062009582577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b8152600401620095ac91906201392b565b60e06040518083038186803b158015620095c557600080fd5b505afa158015620095da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200960091906201271e565b9050620096307ff2d4470e3730069059ff1fbe4f49d8f5edd0a81d25911434817b67ea873ec9fa60001b620015a0565b6200965e7fb6430f90b728ec97641f542a39251f29e419ea8fa8c823bdddeceac1b416caa960001b620015a0565b60008360018885815181106200969d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160400151620096b5919062014292565b620096c191906201420e565b9050620096f17f3ea4dfe0e0ee29c668b945e46d6ee0420bb6e25d82001cbe93ca5d20876d0e3d60001b620015a0565b6200971f7fbb67ed0bf8a63594c5d0ee3e5a5ac7745a4f693faae1e266bd66048bc03aa57860001b620015a0565b6000620097408c878c6101a00151436200973a919062014257565b6200de9d565b9050620097707fdf28b9eb13a43e9f17ea81aa346906479ea69d844f9c870d7176c963fcc98a2960001b620015a0565b6200979e7f97cf79d2a257e2bc91a3900cc9b35124ab0d9384a1131ff3b8ba68afaed93e8460001b620015a0565b60008183620097ae919062014191565b9050620097de7faeb2e81a509bbc1eaeaf899d7a1bad6e389fbb49da7745493ef49e4c1c32219d60001b620015a0565b6200980c7fae97696aa73793f3c8d819b6ebf5109eea208256dd7fa3ddb496bcab661c19d460001b620015a0565b8767ffffffffffffffff168167ffffffffffffffff161115620098b757620098577f0bb1ded3dc58a436ec95237250838c3d7a70bda698712f2275c642a3c65e51ab60001b620015a0565b620098857fb1b55a6c00587969576bc73cc96df668bd3505e9ad6075eed6d7b3fb4208e6b260001b620015a0565b6040517f870303a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620098e57ff4ed5bf7270322275c120c565e8560a4df1b77963b12d5e892767a9410aaf0bd60001b620015a0565b620099137f337f5aa028d412380996bae7990084095129a3908b924255d6472470ee04299d60001b620015a0565b620099417f395c7dd84494560cf7074d5b39fcb620177637ac6d1c4bff9503d4209ccd97fd60001b620015a0565b808460200181815162009955919062014191565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050620099a07f2a1f3ac5aeedbae936eeecb997b3135e9f7f66e04ba0e700a62d8872f44f913660001b620015a0565b620099ce7f1aeeb9a1b11480db21e3a54789d9a1c8164f6449bc197e5f3bf6bccdd439328c60001b620015a0565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c856040518263ffffffff1660e01b815260040162009a2b919062013d6b565b600060405180830381600087803b15801562009a4657600080fd5b505af115801562009a5b573d6000803e3d6000fd5b5050505062009a8d7fbb08638ef09e7e6b5afe90409c5d1042ba32ed2447eab02e76b73b0382e914a660001b620015a0565b62009abb7f766475e6d5edfa4ad466cd1f2ea945d4175e3cfc3683acae4a2f654b8cf000c060001b620015a0565b808862009ac9919062014292565b975050505050808062009adc90620144ed565b915050620094a2565b5062009b147fe7e679d8de5580b377c52cbb98f9b612be10b6e3e49d1f7d7ccff30a53ce323060001b620015a0565b62009b427fdc77753cb50f0510cd4df0ca43b4c3667e7c3aaefaeff0bf42decf3468eec63a60001b620015a0565b6000600181111562009b7d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b866101e00151600181111562009bbc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141562009c635762009bf17fa710508cc8fd98abba4e1e3a3d43eae2fb4b463eb4960af11e4abff42be86b7460001b620015a0565b62009c1f7f7498183a7af821a54bfdb8a0047c3d10e1fc80d68a3e8d4242e35192dfe387b160001b620015a0565b62009c4d7ffc3936561907adabb910231832e0d57f298ad160444e092724e614f4e698210e60001b620015a0565b828a62009c5b919062014191565b99506200a2de565b62009c917f593c183963387acfdd09f0afbf837bb83e20b4d00fd0d072d3c7d12552952ecc60001b620015a0565b62009cbf7f227032a4d1df4bbd0868f6ab7e5fd77efc850eb45f43c8f1fc4d67ec84341d6860001b620015a0565b60018081111562009cf9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b866101e00151600181111562009d38577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200a2ae5762009d6d7f1448b95e8691d4d6076281a0c8b43d3b63cb6e278f91d9d19a99a7185b14485560001b620015a0565b62009d9b7f97a7415bde6afff03e4dce77fbbd01ba90c27116097d00fe081a8b58e1652c3560001b620015a0565b62009dc97fd2d92f5c985868675c9c052e8c1cb29692eb01a00922196338a937592bad03b360001b620015a0565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354e0d3c288602001516040518263ffffffff1660e01b815260040162009e2c91906201392b565b60a06040518083038186803b15801562009e4557600080fd5b505afa15801562009e5a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062009e80919062012b5b565b905062009eb07f6635b6c18b61ad05ff6085e05bdeb72c85f5617bfe0b5e39ecbd1f12d18d25f560001b620015a0565b62009ede7ff6fb0adeb720b9b7adb754dc2bd3cf5e394ea270ce25102c2c29628c6acf167b60001b620015a0565b8660a00151876080015162009ef491906201420e565b67ffffffffffffffff16816000015167ffffffffffffffff16106200a1185762009f417ffeadf1187eba19b4dedd01e784e492c0c20da5f2bb3a3692a67d41f5df24ee9c60001b620015a0565b62009f6f7fc71ccbf38068984707bb8d283618a998098b7cf6dfb2b27a736049139e917dff60001b620015a0565b62009f9d7f9e65e97b6e5a85e9acee41cf5afa3bd955dbdd51ff742835e46a24359c8e333f60001b620015a0565b838160400181815162009fb1919062014191565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062009ffc7f558c6dddd8699dbbed23751de97d8359c6314f354aace96272f842c6bc694a1b60001b620015a0565b6200a02a7f87bcd5b2e5230e7e1ae3cd6021063ac3274ffd36a7d5cc10c14d589af1b4044560001b620015a0565b8660a0015187608001516200a04091906201420e565b816020018181516200a053919062014191565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200a09e7f921cb1d70bcaf7d903fb866472d9952263c59de050bda85b0b36bafd2b382be560001b620015a0565b6200a0cc7f2cf463ddbcb1745e05483e5aff5c2ddf39662b2529cdd54554002fbab84015e060001b620015a0565b8660a0015187608001516200a0e291906201420e565b816000018181516200a0f5919062014292565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200a1b4565b6200a1467fedc538cf02256fed4ce0eb662276466efedf12421ca36c9f8c7c7469b476bfbe60001b620015a0565b6200a1747f595a05b8db2d056bb1902a6f6cbd5e7c83ac179740bc04445a06d2718d265b1d60001b620015a0565b60016040517fc8c84b2f0000000000000000000000000000000000000000000000000000000081526004016200a1ab919062013bd6565b60405180910390fd5b6200a1e27fc48475f37b0d45795ffaf6a5db31b820886bd10e00b04ea6c9970b50ad901e2e60001b620015a0565b6200a2107f474d5811acd5ea3ee50823d08cd42550faf6b0ac439acca4873d19bb57c8764060001b620015a0565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed108de98860200151836040518363ffffffff1660e01b81526004016200a27392919062013948565b600060405180830381600087803b1580156200a28e57600080fd5b505af11580156200a2a3573d6000803e3d6000fd5b50505050506200a2dd565b6200a2dc7fa5529085b72aa4c0aa6fa65ba3b0d5fca2a4b26673f18abc2030ad75ee4602e960001b620015a0565b5b5b6200a30c7f676c6b045ee52448a2839cfdeafadd1a579b1e6fe4214b6c3452d557c7f2c70860001b620015a0565b6200a33a7fdf2f288b96b7c5b1712b47265cde57abbbc3fa5611d8d827c9db7b60724e3a7460001b620015a0565b6200a3488660018062007484565b5050505050505b80806200a35c90620144ed565b91505062008b9e565b506200a3947f77457e8a0f34ae5998ddda0f18b716941b9a0a82b490404441b3c6e2be5a3dec60001b620015a0565b6200a3c27fb5c9253e2b19e7d9cad31d6a08a8063fe6f4e767ba3a6c95635b7f4d23494c5c60001b620015a0565b60008367ffffffffffffffff1614156200a469576200a4047feb0df40120ea2ee7285d6d8f9f912539375b874b094134f0b49b86c55cb22c2f60001b620015a0565b6200a4327ff9d75f8d49ef5a80ced7f72a7e94d1e45eb315f01eabfe1708b940e356c7056760001b620015a0565b6200a4607f54d9a71d4c2ccce5312953d7e1ee31270836f4d0abe9b7490f9f61863b890a5860001b620015a0565b5050506200a549565b6200a4977f2e0aa2bfdce8de0f057a2f81c2c3682b85ebc4fee4dd05a80535ba2c877b03cc60001b620015a0565b6200a4c57fce3e204081da479054d8dca9f9ac6d9dd7440620797cb88da4141b046d44461460001b620015a0565b6200a4f37f18b204f7426e4f90fae55fd63d354d515a315218c697fb49115b638c3824eb3260001b620015a0565b8173ffffffffffffffffffffffffffffffffffffffff166108fc8467ffffffffffffffff169081150290604051600060405180830381858888f193505050501580156200a544573d6000803e3d6000fd5b505050505b50565b60606200a57c7fdb6458cc05e08b336e0790c1451ea98da8c55f399b2a316045d0549c396e9e3360001b620015a0565b6200a5aa7fc65e1f8e6f042b7c7aa93b55d0fc464984a780f68ff21cd64a56b899352f7cea60001b620015a0565b6200a5d87f9a304036fb6719d5f8d0fffd6370a620c2d130c7c28d326e6f5f84460b25efb960001b620015a0565b6200a6067f641159e61bdf2c900ee6a738405853a032863463826fd94d402662f2d8240f5060001b620015a0565b60008251116200a64d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200a6449062013bf3565b60405180910390fd5b6200a67b7f6277ba5a34555d5469f97c1f931d34a369e81cb48854a651be60df154855812860001b620015a0565b6200a6a97ffa16a54cf63b8cd89c7675386925c636f2c47c5c746221bc7c152e8354fe069160001b620015a0565b6200a6d77f5fb8e95eb88d48f32ef01617185defc615c22c6c828c0419a074d8582cd0620a60001b620015a0565b6000825167ffffffffffffffff8111156200a71b577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200a75857816020015b6200a74462010a9d565b8152602001906001900390816200a73a5790505b5090506200a7897f51ef10fbad4408679bbbf1d905136fdb6618b47ecff0aca9649216229ca580b460001b620015a0565b6200a7b77f5ab2dc6e80adc23371dac0b7a9dad1ef7d36e49d92922bc361e29328f47e641060001b620015a0565b60005b83518110156200ba57576200a7f27f8dd1976522cf158a4277cb8ae8a6399490fac084fffc5ba86a51f33069f1853960001b620015a0565b6200a8207f0a4499cf04fb742c3d3042d95bae68c89ebb31b5a0a4745b16bc7500a49c820c60001b620015a0565b60008482815181106200a85c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190506200a8947f9d0b34ee6cfe5f201aa5540590e09a118d2fc7b5a32143010983702ad120ef2360001b620015a0565b6200a8c27f4f699d550ff91e67e4483bb58a04563e94e1c0d6500ae2fa3f41aa785117b93960001b620015a0565b60006006826040516200a8d69190620138f9565b9081526020016040518091039020604051806102e00160405290816000820180546200a9029062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200a9309062014481565b80156200a9815780601f106200a955576101008083540402835291602001916200a981565b820191906000526020600020905b8154815290600101906020018083116200a96357829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546200a9f29062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200aa209062014481565b80156200aa715780601f106200aa45576101008083540402835291602001916200aa71565b820191906000526020600020905b8154815290600101906020018083116200aa5357829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016007820180546200abf49062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200ac229062014481565b80156200ac735780601f106200ac47576101008083540402835291602001916200ac73565b820191906000526020600020905b8154815290600101906020018083116200ac5557829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200ad1f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60018111156200ad58577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200ae1557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200adca575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156200aea557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200ae5a575b50505050508152602001600d820180546200aec09062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200aeee9062014481565b80156200af3f5780601f106200af13576101008083540402835291602001916200af3f565b820191906000526020600020905b8154815290600101906020018083116200af2157829003601f168201915b50505050508152602001600e820160009054906101000a900460ff1660028111156200af94577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156200afcd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600e820160019054906101000a900460ff16151515158152602001600f82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152505090506200b0c67f4bb39638ccc1a9f38d690908c14e1296d5a7c53e49e1afc4c6e7b5f8de71ed5360001b620015a0565b6200b0f47f21b0024845eabfff4136cae851052a4550ee32d964bf9521863a54a959d189eb60001b620015a0565b600081600001515114156200b19e576200b1317f8bfb80018024877eecf7365b3f2bf626355fafd59f0514835f7d44600a0f377860001b620015a0565b6200b15f7feb48c84303c06de35c0a22c1fb5eeefa35531efd3897a76346caeed5a171e17460001b620015a0565b816040517f6c5249c10000000000000000000000000000000000000000000000000000000081526004016200b195919062013a52565b60405180910390fd5b6200b1cc7f2f15e999569c7e08b0701a041b321e49d8ba0b7e2c2bf79056227601e913a7d260001b620015a0565b6200b1fa7fba98946e89854548220573adb3d7e1d3847c705c16e671c28499dd8b6780ef1060001b620015a0565b6200b2287f2da02bc1ef8fc815c8e1b07922ea78f8f53c46e7e2a4c0041888a666e38521e760001b620015a0565b6006826040516200b23a9190620138f9565b9081526020016040518091039020604051806102e00160405290816000820180546200b2669062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200b2949062014481565b80156200b2e55780601f106200b2b9576101008083540402835291602001916200b2e5565b820191906000526020600020905b8154815290600101906020018083116200b2c757829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546200b3569062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200b3849062014481565b80156200b3d55780601f106200b3a9576101008083540402835291602001916200b3d5565b820191906000526020600020905b8154815290600101906020018083116200b3b757829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016007820180546200b5589062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200b5869062014481565b80156200b5d75780601f106200b5ab576101008083540402835291602001916200b5d7565b820191906000526020600020905b8154815290600101906020018083116200b5b957829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200b683577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60018111156200b6bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200b77957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200b72e575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156200b80957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200b7be575b50505050508152602001600d820180546200b8249062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200b8529062014481565b80156200b8a35780601f106200b877576101008083540402835291602001916200b8a3565b820191906000526020600020905b8154815290600101906020018083116200b88557829003601f168201915b50505050508152602001600e820160009054906101000a900460ff1660028111156200b8f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156200b931577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600e820160019054906101000a900460ff16151515158152602001600f82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250508484815181106200ba34577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250505080806200ba4e90620144ed565b9150506200a7ba565b506200ba867f436d834e84d366d7ed3fdab4b2731aea1fcb52cb1a25922f0107770427dff21a60001b620015a0565b6200bab47fd38dc8ee092fe161bb932e0e1b507bae35797ae4513841ac7488c29e170e1a6d60001b620015a0565b80915050919050565b60006200baed7f438b7c63a5809859924121b74371bffdb86715a85d11f4142fb04730466f503460001b620015a0565b6200bb1b7fa257f053309ee13448581c267e478ec419ed24ebb37c69dc325c0ee4d9f23a4e60001b620015a0565b6200bb497f44a6cdc69e1afba364d68d3c73a49210c19debc0dc4ae6ecac73a44eff72e0eb60001b620015a0565b620fa0008284606001516200bb5f91906201420e565b6200bb6b9190620141d6565b905092915050565b6200bb7d62010964565b6200bbab7fe53554de5b2c6dcbb32fb5ddc14a029494fa514fdf498d2d40865dcda916b45f60001b620015a0565b6200bbd97f9c13f4ff734a6f7eee3bd47e6f495b0056c856ddbe4a0cbe4577a966fa7b346860001b620015a0565b6200bc077f99daf98c674c78c79c2c0af2516ffd710cc5f198e098930c864aafd9adb920a360001b620015a0565b6200bc1162010964565b6200bc3f7fa062ea5685fee4aab19e0e7ad8be4acfde38de148ccf5adca7b9be7a93c56e0b60001b620015a0565b6200bc6d7f9efaaf89811b62ce32fa7a3db9890652fca00f1827de21afe479768a74b7703760001b620015a0565b60006200bc7d88888888620015a3565b90506200bcad7f217985ad4a7266d851008a41b4e34e40d43570ab7c5f8caac3721c7731a09c4060001b620015a0565b6200bcdb7f7061159c8da20a70dd4a7e7602c2486f7f9284afaa07fd8c66c70da0aea843cb60001b620015a0565b60006200bceb8988888862005dc2565b90506200bd1b7ffc760423793427029ac6e6caa5c91221b38bfe2ddc75098cc682e5e7a4e59e8560001b620015a0565b6200bd497f4700e78bbb56c6c3cb3949c24c4a48600d4f0c5f990b0769b30e5a0bfee4d40860001b620015a0565b81836040019067ffffffffffffffff16908167ffffffffffffffff16815250506200bd977f79cb8ddeb2a6319b9a71e06be79c27d2ce5f5c4a256aa0e6e051ebafbe71bc1460001b620015a0565b6200bdc57ffd56fcbfeae8b7f81d992296d272259147969496a822f0c75311a11c9508a19f60001b620015a0565b80836020019067ffffffffffffffff16908167ffffffffffffffff16815250506200be137faeeab0f86657b027adf6135d4b646cd8d1997aea8b834ad8f1e9820161e978a260001b620015a0565b6200be417fe66dd2e4b236471e68dfa46856642dc70c0591ebe4d54a3bf6b66d1957f6638860001b620015a0565b82935050505095945050505050565b6200be7e7f944bf0004173e26ccacff736a32e72e30731cc1357af87c6a64b068ba0aac50260001b620015a0565b6200beac7fe20adb1a3dc2e3ad849b255a8e79fb7e371bf8c976a3f2c1c3e156f617f9e96060001b620015a0565b6200beda7f931025eaa38d4100f9536d325d95da9bf06f411f760ceea4409a79f889a2ea0060001b620015a0565b60006200bee7826200d33d565b90506200bf177f40e7ba2dcc65eba1ed28c934e69e2787a99b2e2b6866ba0787afe8c6fd0a144f60001b620015a0565b6200bf457fe04029617dec975d0c5fa51940a0213540119e2d4d26a6a1f75e459ea5339f8260001b620015a0565b6000600167ffffffffffffffff8111156200bf89577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200bfc657816020015b6200bfb262010a9d565b8152602001906001900390816200bfa85790505b5090506200bff77fff7ef79382f5411b6b61e1e0f19eb62a8f4238076ef00b9b5bceb015b4e8db2760001b620015a0565b6200c0257f8d385dc8fc6b11adc66c09365ddeaa5a54152b265b49cc2418b8d79e967f443560001b620015a0565b81816000815181106200c061577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506200c09a7fd24c944d706b4662785cc657c7efcce90012d00964f546832cc5c30973e261c560001b620015a0565b6200c0c87f48950c13e89965881ae3d403cc658eab4310d04b567409fc9f7dab142bc56cb860001b620015a0565b6200c0d38162008570565b6200c1017ff4f33d55631a58ada81a0821224271f35eecf482a489cf712d448c7dcb4b341760001b620015a0565b6200c12f7fd4684b5b53cc0618e1c7bd5b83df1555772d0700830a4cd421438a7bc284512d60001b620015a0565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f456001438585602001516040516200c16b949392919062013afe565b60405180910390a1505050565b6200c18262010964565b6200c1b07fb2465e4b5c6ee98bb70addaa69d1bbf2b2cf1f471a2cb19d079fd9e4162574aa60001b620015a0565b6200c1de7fddc24047eb9ac7f2c488986ef60d8f5c8110a4bb81102134217573da52dbb84960001b620015a0565b6200c20c7f16528a43ed8102cacbf9d96a60d43372852b65e42499dc2ea560606046ca737560001b620015a0565b60006200c23c7fe582a221657b948230177c6476bc47a665d7d805eb4f2b9ac77278b7d201922e60001b620015a0565b6200c26a7fc7ffd4f66bae2be4ac58ad6a49398278b0005c85dee7162a0c19f645d2ae2be060001b620015a0565b60006298968090506200c2a07f0d60a796501a28fc32c78678cf4ffa493da5c07cd50f630b1b64f2e49f51118b60001b620015a0565b6200c2ce7ff7a9c29e816451d3dd3f705ac9c22f9a282f17dbc7703ba04f3a89db001a11f260001b620015a0565b60008661018001515111156200c37f576200c30c7f9b5f4500b980ab024a912e1192232e1d34875790084998618cf410eae20b2f5160001b620015a0565b6200c33a7fb06787a1d86be30928c28df7c7bc75d6f3d8238424ae538d0278bd41abc9063460001b620015a0565b6200c3687f863a0eabc2f3a141356a9f08e21aecbd4eba8d0ae5d7f5c3414fa687e451bb5260001b620015a0565b6004816200c37791906201420e565b91506200c41b565b6200c3ad7fc6f44ff32a0bcf7ba8a4716a3e0abd272b463afd2c9d02ac9fab5d71758d855d60001b620015a0565b6200c3db7f9fd52ccc2f41147a23410e80ebe5727d77a796ee95e390d52f61b13c530f758a60001b620015a0565b6200c4097f1f640d5d0ba9262d20dd82045622636cf1debcc40e73c68b5b561235f034c38d60001b620015a0565b6003816200c41891906201420e565b91505b6200c4497fce8c73828107763f824cc7a2fdbcecd6653e252e8dfbc6d1750e3a6f0b11632a60001b620015a0565b6200c4777fbc43b330e0c207bd19776067478d6493561f06b7f915a15d7dc56deb818aac5360001b620015a0565b6200c48162010964565b6200c4af7fbe192d68aab4df86ab4b98a7a9335eccc0bf203d6f2d5aa8a24d2c3b654d023260001b620015a0565b6200c4dd7fc030bd9adf8a5bef1541ef5faaeff6ee3f884feb14a8a1aa840cac16b7ea3c5f60001b620015a0565b82816000019067ffffffffffffffff16908167ffffffffffffffff16815250506200c52b7fe5d0de1b3c8958b76ad4ec361afcf3515314d533038b7c2e21ebd0a161a94cde60001b620015a0565b6200c5597f1952e3ea2ae72b328ac1d20e4581218c2541a1fa56e39850bd51ec7597daa21160001b620015a0565b600060018111156200c594577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b876101c0015160018111156200c5d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200c670576200c6087f0ebb73d09e14ff731e20626c777eab77e923a0c40c848fb2fae3bf5ef692ead760001b620015a0565b6200c6367f0b2f01f68a1e0fbe450eae0aa512e394a8204bfd805f4fa6ec4d13863b518c7860001b620015a0565b6200c6647f74e6ecd310e70b198143f17816b79dbe720368983000ecfd4f2d56f88ec2b40060001b620015a0565b8093505050506200c86f565b6200c69e7fbd0472ec4cc9554e7069b20ea2f0873d9a60da0621ce9a6beb0cc37b82522e3760001b620015a0565b6200c6cc7fed9399dd85d4438fdeb2365229bdd5d7a427a130cc1492e9ce92a6487648c34960001b620015a0565b6200c6fa7f8d6cda09b6ac598965523e960cf0efdde48898dad6bda53b86a403ee0ea0ff8360001b620015a0565b60006200c70988888862001661565b90506200c7397fd54804b2ee6d59a2aa938e306af5a218ae12e0ef286fd1e4f7a45ad87c09ce6160001b620015a0565b6200c7677f11a1fa1ab968cb207dac44ba78e6a32101fcc6f7e556578e73614df995b36ff160001b620015a0565b8060400151826040019067ffffffffffffffff16908167ffffffffffffffff16815250506200c7b97fd62d8bc4a654258f1479e5b6bfbfe30d65376a5b7382613032f4c306d57be34f60001b620015a0565b6200c7e77f1769590e259297417b0378a986ef120be903c8c356df373b2c857a491bf8597460001b620015a0565b8060200151826020019067ffffffffffffffff16908167ffffffffffffffff16815250506200c8397fb219c0b9d4050add9615ce47cd042f593150936f4b0a36de6db6b4487ff22b1160001b620015a0565b6200c8677fb122c90c16be5246fa626aac8e1e18ec51dae63b99f4735df8540e400957216060001b620015a0565b819450505050505b9392505050565b6200c8a47f5898ff48267bc362603a052d3f8cbcbc76704e71b12ad683cab294a0aab838d260001b620015a0565b6200c8d27f90bab23fc9a94161cbb3ad896fa9587e398713cb77d4a16c3c03710bb7d8e8e960001b620015a0565b6200c9007fac31b44f3ca53a31ce9b1469dd4c82a6e995303d8021db80c646b6d38b821e3060001b620015a0565b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906200c955929190620108fd565b505050565b60606200c98a7f387fc187174dc5bf3f189e4b690132eff002c22d10b933ab9f8db4142efabd8e60001b620015a0565b6200c9b87f9ae13d2bb882495c21f1b892685a4983901d54a519db1054bfc7434dee2e0c6660001b620015a0565b6200c9e67f90f6755c20e8557b6498a716e740a280e3060e0378ae5a008c12d4c2882a5bb660001b620015a0565b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156200cafd5783829060005260206000200180546200ca699062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200ca979062014481565b80156200cae85780601f106200cabc576101008083540402835291602001916200cae8565b820191906000526020600020905b8154815290600101906020018083116200caca57829003601f168201915b5050505050815260200190600101906200ca47565b5050505090506200cb317f4b3f6eafc480a7379146cca4fa23335282a7215c27f64ead52865de3b9e1d13060001b620015a0565b6200cb5f7f66d109be6d09d9150ece19d0312ae36e25e7f8bed259900be86106b87048644860001b620015a0565b6000815167ffffffffffffffff8111156200cba3577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200cbd857816020015b60608152602001906001900390816200cbc25790505b5090506200cc097fd89e51d83ac654caeab43c3e7fb336cbd2347c365cfe6f2a7814dd5e528ed60a60001b620015a0565b6200cc377f068088e1e0a3164cbec40d00bb33366f9028fe61a5e8c43c4d0d1e71fca5232160001b620015a0565b60006200cc677f57f546cfeb2ecdb41382d9a94561376beacbf13961fae380ee3fdfe517e8462660001b620015a0565b6200cc957f202378abc95088401786fd8d5f9f94b6b3eabd3abca05b67e24acc89e12704a060001b620015a0565b60005b83518110156200d2d5576200ccd07fedceecd1c56994214bb7345439cbdb617d7bbea4a3a68c505d550254038406af60001b620015a0565b6200ccfe7fa390cb4e8bb8bdd8717b87ba4a4044cb89611814d4b246654618625afddd2dbd60001b620015a0565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe28684815181106200cd78577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016200cd9e919062013a52565b60006040518083038186803b1580156200cdb757600080fd5b505afa1580156200cdcc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200cdf79190620124d8565b90506200ce277f0e9a73e50c0a93382778abeb582b3cca3089c4cd1ff82036f4f16db53a87325960001b620015a0565b6200ce557f679f353d2641982e84b49a838e61c4b9a83eb82320777752dd0583f6e8e7308460001b620015a0565b60006200ce857feb6291909339b89d5da54b588b5ff89d166078a702e0e4adf72532971fa7c4a860001b620015a0565b6200ceb37f939f69ec31cd9e529374f946d88b4041c2228e2338fcd88c25e4fcf8d4fa343260001b620015a0565b60005b82518110156200d09c576200ceee7f91005e23bb9c69a1b427e5670bd2721e990fb07187aa595ff7f14ec603b362c160001b620015a0565b6200cf1c7f2c84a07270952b2ca79bdd33d571914fd8fa6b4434acd41bc200bad73b32a1da60001b620015a0565b8873ffffffffffffffffffffffffffffffffffffffff168382815181106200cf6d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff1614156200d058576200cfc47f0c01464459099109af6ddb252e27b60276a1f77b4ba0ea40a511d08950bc266260001b620015a0565b6200cff27f4a16a3888408d130267800974cb3f5e6a5e9b6d79655204ed3b1cad92ebe378a60001b620015a0565b6200d0207f7912575e61593fe96f1d13272db936b7df89b855d08af2740fb192877759ab4a60001b620015a0565b600191506200d0527f61c2f39b6dfbebf6aa4b08795f43da369833f08168f16d2bbea18e98037a48d360001b620015a0565b6200d09c565b6200d0867f1ae1b38b7b7495121937c55fe8d91aa8e3d6383075bc0ae87d1ba60c7e7482f860001b620015a0565b80806200d09390620144ed565b9150506200ceb6565b506200d0cb7fb53f9b6348cabac9b3a8b7be41201dba84c1c50549ac3c9e361701bb37539dd560001b620015a0565b6200d0f97fecbc1df80962fef1814fcc10acd9307f01ca060a681b78ce0be586e07f71410a60001b620015a0565b806200d163576200d12d7fbe3619b2c3d340abd8740482d1d5bf3ecc7a3217e1cf7bbf928f12aa0707e7f660001b620015a0565b6200d15b7f77b0c12f7bfdd306b6f1bc623eb1784c042d33b3caa93bea46fc291388d2763960001b620015a0565b50506200d2bf565b6200d1917ffa626994d9dad6c2993be845812fff6b38adc278e284f8cfcdda4bf1d28bab1f60001b620015a0565b6200d1bf7fb021bc26b7f58e22c635847dbb9505287fd72a27a30358a09a415e6d46a4ff3460001b620015a0565b6200d1ed7f46f9a942d369ce077ce569a431899e1e51eb97ecf28fadda98bd81f4c91ce3de60001b620015a0565b8583815181106200d227577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858567ffffffffffffffff16815181106200d273577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506200d2ac7f7917f4dc92fe5bc718f1a0c0b3e3e64658e4d5c1ffaff204d6b9941947def0c860001b620015a0565b83806200d2b9906201453b565b94505050505b80806200d2cc90620144ed565b9150506200cc98565b506200d3047fc5484101124372f011979c6677f1a3f14cd94c01457fb7b80441a411101b851a60001b620015a0565b6200d3327f4725ddfbc6d7d212581bb45f78d106d3dabb63258454133b4376ccdaecb2984160001b620015a0565b819350505050919050565b6200d34762010a9d565b816200d3767f1bf68ce7aedc9874685902dbac8645673c058afb04cc4647ac7e75b26c510f5360001b620015a0565b6200d3a47feb31a4ba25c36671c3f5cb88ffcb8511a1ac39e6e6ebf73c241566266d9c9ca160001b620015a0565b6200d3d27fa4fb360503ed57bee61d9d1c9d9da8dded5a43c82844162d715f2099b56c989660001b620015a0565b6200d4007ff1484d6eff33dcd8113e762310d7500b14808fa5c4923a7421c8a0420a90cc7160001b620015a0565b60008151116200d447576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200d43e9062013ce1565b60405180910390fd5b6200d4757f9141be990549866bcb5dc4e4211093aa53353e33774e2160ecf1643ba6c9f1b260001b620015a0565b6200d4a37f8d010b557c8db11436730cd7e786539722e1a23fedf99dc43e32ed2da04b86ce60001b620015a0565b6200d4d17fc044375fadb95d358cee357feb34383d9cd5405865bfc6ab5f36bef1efc31b3460001b620015a0565b6200d4ff7fcf4a4be628adacfb14428def7410ad07ff6fc3ad8578eecbc4b7f3597fac52f160001b620015a0565b6200d52d7fb10642d10e46a8f2b6cdde9473b70fc6091b1997274ce55d46926208737522ac60001b620015a0565b60006006846040516200d5419190620138f9565b9081526020016040518091039020604051806102e00160405290816000820180546200d56d9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200d59b9062014481565b80156200d5ec5780601f106200d5c0576101008083540402835291602001916200d5ec565b820191906000526020600020905b8154815290600101906020018083116200d5ce57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546200d65d9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200d68b9062014481565b80156200d6dc5780601f106200d6b0576101008083540402835291602001916200d6dc565b820191906000526020600020905b8154815290600101906020018083116200d6be57829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016007820180546200d85f9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200d88d9062014481565b80156200d8de5780601f106200d8b2576101008083540402835291602001916200d8de565b820191906000526020600020905b8154815290600101906020018083116200d8c057829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200d98a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60018111156200d9c3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200da8057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200da35575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156200db1057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200dac5575b50505050508152602001600d820180546200db2b9062014481565b80601f01602080910402602001604051908101604052809291908181526020018280546200db599062014481565b80156200dbaa5780601f106200db7e576101008083540402835291602001916200dbaa565b820191906000526020600020905b8154815290600101906020018083116200db8c57829003601f168201915b50505050508152602001600e820160009054906101000a900460ff1660028111156200dbff577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156200dc38577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600e820160019054906101000a900460ff16151515158152602001600f82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152505090506200dd317f6a91e296095e9591810a3946baffc11ba323b2de00c1b0db24e4bb1bd927c57160001b620015a0565b6200dd5f7f925a21be4c480dcb9911569499d901357d0f1d6e69e5712a161402e33dea108460001b620015a0565b600081600001515114156200de09576200dd9c7f75259cb50e4defc4f3ab59cceda4487f4f84be9296792b7cf5ccd438dd74a2a860001b620015a0565b6200ddca7fd9a98aa7f5271ed6c5339c9f92e613610cc8a42f138b163e90950e425cf3303660001b620015a0565b836040517f6c5249c10000000000000000000000000000000000000000000000000000000081526004016200de00919062013a52565b60405180910390fd5b6200de377f52133f711aa408c13cc50008f42e1bfc4185b0e8917aad8d80b2a97607390c9d60001b620015a0565b6200de657f5f105ba679e5fc64a3480e95b5b824e38d53d3212f88691274338a38c56ad18b60001b620015a0565b6200de937f5ae2833e3301c0662eb171913d4fa8296ef4a90b05ba2ee257cc34876a5e128960001b620015a0565b8092505050919050565b60006200decd7f1fe9fc1472a69e0466fd47bfa10832b522e056154c6e360255eae6696802263560001b620015a0565b6200defb7fe0925de2bee54b459f5437ea8911810de76fd7826191d4d4e30e358406a5d08960001b620015a0565b6200df297f1f139fb914682594f7d735af39cffe774340f9bcf8e119b1fbf6ffa116380d5660001b620015a0565b620fa000828486602001516200df4091906201420e565b6200df4c91906201420e565b6200df589190620141d6565b90509392505050565b60006200df917f8c6d1ddd2276911fff5706bdd5e97f5a0a3be042836862980245cccff080280860001b620015a0565b6200dfbf7f7393aa05998eb8d0019286df986114f0c27a37c68238cf531fcb29d468e9c0b160001b620015a0565b6200dfed7f9e7875b337bc99fa44106d38412015164105b797494c6bb2ab24b3b6a9b831ff60001b620015a0565b600183604001518385608001516200e006919062014257565b6200e0129190620141d6565b6200e01e919062014191565b905092915050565b6200e0547f203dc81212c80726daf4194adc73ded580b9b650705f22d5fb48495128b5e40660001b620015a0565b6200e0827f3c9f51c76a2c5c0d3fb2964cdac85e0a239461872035fc298fa5ac048805be6860001b620015a0565b6200e0b07f8d5061506bb3c941d1368ab7b0d7c96bc3d122c40e37124dbde7cad089b3b62360001b620015a0565b60005b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156200e366576200e12c7fcaba45db971d21f9d4a03d9e43d29098d41f747fd71bc38b21828e6e70ce484460001b620015a0565b6200e15a7f9dc08cd7c408ed1de7bb5137fab198aa099646990a36b887ab6878e76708903c60001b620015a0565b8180519060200120600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481106200e1da577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020016040516200e1f3919062013912565b604051809103902014156200e322576200e2307f6290722ae13fd47962163472d09b14446e7895d79982f46f4ef19798b7b1753260001b620015a0565b6200e25e7f195845903c31ef7d6e353dfee81d50838b743bde79239c60634880d279a7a89360001b620015a0565b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081815481106200e2d6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006200e2ee9190620109a3565b6200e31c7f5c30b6d9dcdcbbd07ff74f956c74d210f08139274540453dcdf7489c60cf331760001b620015a0565b6200e366565b6200e3507f28b476c4a37d663b13f6004a1ecd9ea2f97ca3bc40690c820921a70347ee85f460001b620015a0565b80806200e35d90620144ed565b9150506200e0b3565b505050565b6200e3997f67019cf9fccb4f80a4caee87009ece260606c01f458095d5ff841b3b4ac1c80560001b620015a0565b6200e3c77f129579356cca7a0a41602dce565cddc4d9aefee3d2da3ba6654fa3863ab8651860001b620015a0565b6200e3f57f2fc6a7e4a7e99902408144a526d72240ed118c5532b26e3de1bcd113cb9993d160001b620015a0565b6200e4237faa1a1b7cf3cb3eadbb2ec979cee1b256ce63ba462af24fe34132f5fbfbf6258060001b620015a0565b6000600682600001516040516200e43b9190620138f9565b908152602001604051809103902060090154146200e490576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200e4879062013d03565b60405180910390fd5b6200e4be7fe8140b4abe7593d0578e30c4e3e3969f8718eafb61539843ed8a2ddc5e835f2860001b620015a0565b6200e4ec7f34b11acbd5fa3a61151082edb2b8b7d99af998f5ed49753604c1b1653aa9c5cc60001b620015a0565b6200e51a7f61f794a4e9b4d35baf4a0cfc5eb8b347bd3eceebfa177a817c6dcc21f5b5cccb60001b620015a0565b6200e5487f0769876d7522279fb30115f9fb9ce7133c4328dc1eda15cf1dddc2493e56c12660001b620015a0565b43816101000151116200e592576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200e5899062013c15565b60405180910390fd5b6200e5c07fc1b826d6bc33134e676fd18c11ebf1e38979453839b38b6a2e5461dc13f47cc860001b620015a0565b6200e5ee7f15410b331db777f9e3124cab1c29961c540e596de47829aaf50bdf1526ab1fa460001b620015a0565b6200e61c7f4d423ee8d384a2e9421a6d127a1f5989671402b3841959e160d50a1639ac070d60001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200e68757600080fd5b505afa1580156200e69c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200e6c2919062012819565b90506200e6f27f03bfda97ce2f32529924ce17aa6b6679bb2504bd6b32441b1f10e4993c8760ba60001b620015a0565b6200e7207f2f443119e878cd8c96fff0abed6b0e95bb4a64de85744a05785604855841d3aa60001b620015a0565b600060028111156200e75b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82610280015160028111156200e79a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200e867576200e7cf7f1221c6390881ff415a4bb3ebe9fe07e6e4ba5dc8751179787a9bb9ced6733d6f60001b620015a0565b6200e7fd7f11b63e4db0ff33f4ec928284eca1145cbf6a664bbbd4e2fec712e46f7ced3f5060001b620015a0565b6200e82b7fb80fd6cb192a74545fb0006d28744b6a09cb317cf265bf450e816d77fbc352af60001b620015a0565b600560009054906101000a900467ffffffffffffffff168260c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200e896565b6200e8957fb80b728a5f386cca4a81807588afdf58b2d5c880ef266d9ca9a1625e36c2f7cd60001b620015a0565b5b6200e8c47f884cdeeac7e0b4b64a24cb2e435ce790c31886fa8bb34da0218248c579169de560001b620015a0565b6200e8f27f7676a18e0ed939a80bf384e96137add3fa77742cbf9d5ad50c11e65a37ddba1a60001b620015a0565b600160028111156200e92d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82610280015160028111156200e96c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200ea39576200e9a17f73e6f8b6eac0720a019ba5e7651649459f3bcfc805b0bf270a919201dca5ce1160001b620015a0565b6200e9cf7fce9fc85a7213753f5203eb0573b3cf0756877b019544478864efe34c1b3b127e60001b620015a0565b6200e9fd7f3a6ddb196f47d6038cd44f9f2aa91e652bc56c4adf81919244e1160cf314112860001b620015a0565b600560009054906101000a900467ffffffffffffffff168260c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200ea68565b6200ea677fc1938dafece2f497da4878e6f86f390519a14f71badd7018d05bbd771b3a51a060001b620015a0565b5b6200ea967fd74a703b65743152804fde57b9477936709bac93f200266a0c874f00f68f5b8860001b620015a0565b6200eac47f4ade62425b730cc3d9934307e6dd2d4bdd0566016ad0b3bddb74ecdd7c61e41b60001b620015a0565b6002808111156200eafe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82610280015160028111156200eb3d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200ec0a576200eb727fbf31ddd5b543927bd0ef9b67abc63c493abcf9ba604a3b6d0811659329b2452f60001b620015a0565b6200eba07f2c3fbb47390a63e8480dc10a1f812e30dc9ff35e6da1f5f938a8e5f99a0f2a3360001b620015a0565b6200ebce7fd437b53f04f2eae768159aa9ef9297b8aa9da0bfda3f64794cf2955933f21bb460001b620015a0565b600560009054906101000a900467ffffffffffffffff168260c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200ec39565b6200ec387f970a09f32507fc20ce31ba01dd673be592ce026f938c277e8d00fe1a3717ab4d60001b620015a0565b5b6200ec677f4b3bb5ce59c0fabf6c04027f862efb42b9829d804997b3e02248bdca7f95077d60001b620015a0565b6200ec957fd761a8ae2fa58b5a2f18c798212fc679a730428feead73bd5d0ce68b7571bb7060001b620015a0565b6001826101c00190151590811515815250506200ecd57f5dd4a42b222c75d1e428eb6c4b35f04f2bd113ecd8411f5819e7fc766a70485c60001b620015a0565b6200ed037f3d9a19adfed7b393178e734f9811f5d81ebb7bcb69fe6494a824c67d08a1ffff60001b620015a0565b6200ed0d62010cc8565b6200ed3b7f9ea16bcb96e0a9de9b060c175dc7c2fe8b5f5e61525f365b9c6473f28534cb9f60001b620015a0565b6200ed697fd38e261fd40727a7c42fe1bf95b3833d66ca47a1ad4dc4cabe2bae116d60f63760001b620015a0565b8261010001518160800181815250506200eda67ffa297a33092e9ed4e9f9108420625c4df6aaf792ccb430ed6bfae9ed56df465160001b620015a0565b6200edd47f20758964a68f56709ca6dae4735572a8b0409348a438e71217269be5aff7ff2b60001b620015a0565b8260c00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200ee267f781c85f49e865980d5ccde5443f8d67c88b173fccd41fcfd88327c8deeebd34e60001b620015a0565b6200ee547f13e970289b93971335134c53f6d1d2070d03bf8d758f0d2f47106aa7f66cb94660001b620015a0565b8261012001518160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200eea77f7531ad35f7b2d4c0b5258794b46ce84a80bbff6b6e218cf53c2c9f8fda484f5f60001b620015a0565b6200eed57f59975163f464eeeb625479e1ac69fefdbe60dc069fdf3f3668a7fdbcf224af6360001b620015a0565b82608001518360a001516200eeeb91906201420e565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200ef387ff79dc69b33ef26c71a0f0e4c22e3b32b0d48fec555ef342ce164a30156fb696460001b620015a0565b6200ef667fc4f37235380951ca9ca054719429dc8637b9ef6839bacc749f6012b53cb075f460001b620015a0565b60006200ef7582844362001661565b90506200efa57f3ba3bf987e52f312e64c7d8a38330e66a7452fe9df5f2cd5c68453fc0267fbe760001b620015a0565b6200efd37f04abcfc39202ee7f8e925a1decf58889b009d87d313ebc037ccada9eaa0b1aac60001b620015a0565b8060400151816020015182600001516200efee919062014191565b6200effa919062014191565b84610140019067ffffffffffffffff16908167ffffffffffffffff16815250506200f0487f5f50d342d134a0372e97a3d506f49381b0fc757ba586dbdf0a34c9a517e2984c60001b620015a0565b6200f0767f0b30cae3e0aba407ab3885836c0a30271cac40019b4477beb744255af12d7bc060001b620015a0565b6200f08282436200df61565b8460e0019067ffffffffffffffff16908167ffffffffffffffff16815250506200f0cf7fbd42828f125948f244949752bb7bd81f13c342004ae6d0395abbc686fd8f2c9a60001b620015a0565b6200f0fd7f409796b52537f7e5955a4243eca2e28f04a6b9b173d1f8dc0cff3e6fa5ee59cd60001b620015a0565b600060018111156200f138577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b846101e0015160018111156200f177577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200f89a576200f1ac7f1479adb8afe4794d69ad8c4aba29f82c9434d5cb49325c446cdb8914e534a2e060001b620015a0565b6200f1da7f5e785cb86e6cf030931e0abae28ac5a9197d1d74a1d471389c6132c8a1426b3f60001b620015a0565b6200f2087f9cb9f632848548094a484a02abacccc92ae4b7137d873b10abc1dec928dd777460001b620015a0565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354e0d3c286602001516040518263ffffffff1660e01b81526004016200f26b91906201392b565b60a06040518083038186803b1580156200f28457600080fd5b505afa1580156200f299573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200f2bf919062012b5b565b90506200f2ef7f4bc14e30120e64e30cba7bca50d3e896091707ea147a7b5661992cd8bb2a8e3d60001b620015a0565b6200f31d7f8c90efca346a7e7e6c062e8b8b0ce22989b5fafa6f4b2931b8ba1de9acad914260001b620015a0565b84610140015167ffffffffffffffff16816040015167ffffffffffffffff1610156200f3e8576200f3707e952324f4814e140cabcf35efa7990248b38511894b9744165b960e5fbfc35360001b620015a0565b6200f39e7ff08c3fe625d5030617c1099a1f7340f398fa7724aa671d26693ebf3214b2670b60001b620015a0565b80604001518561014001516040517f62fe63d90000000000000000000000000000000000000000000000000000000081526004016200f3df92919062013e74565b60405180910390fd5b6200f4167fdf8ed198e73d2731e4e4cba189fd6e739f96689a5bd7a83aaab4d97f41bcbb8560001b620015a0565b6200f4447f1d78b2f786405b24946f50eec23fdb02f744334747fc70aa1c3df77a733e30f460001b620015a0565b6200f4727fd7ff92196792e1d01f7b5abc5d90ab6eb312ad1557d7e9ed898cdada23b3443a60001b620015a0565b84608001518560a001516200f48891906201420e565b67ffffffffffffffff16816020015167ffffffffffffffff1610156200f54d576200f4d67f9b44bd83e8eca3ad20a65eda17edc84d6d78fa7197d71e9191ae9a715fecf1e260001b620015a0565b6200f5047f4924eb45e77693ede686422ac1276125c280861908344a59187e5133c210bcc760001b620015a0565b80602001518560a001516040517fd39dfd030000000000000000000000000000000000000000000000000000000081526004016200f54492919062013e74565b60405180910390fd5b6200f57b7feff9fb634b58e60e575d86e6fcde63b4eebc1e40021987af76cd8215458c05f260001b620015a0565b6200f5a97f540d560d4f33c9ebc3a1413252c62c200392d5ca524d5cc0b90ac62088c48dde60001b620015a0565b6200f5d77ff130a7a8f5f2ac96c46df65c77cd7320c6730a94417018deae3e00c2ac89f46060001b620015a0565b846101000151816060015110156200f68f576200f6177ff2cf3a1e53fbc5f992de757ac4db78204164ec34e63cd02b83163da8ad58495060001b620015a0565b6200f6457fffee23fd612c326f9e7ffb1eba576dbb63442700c0bfbe7305d49f5567ff901160001b620015a0565b80606001518561010001516040517f0c8034120000000000000000000000000000000000000000000000000000000081526004016200f68692919062013dfd565b60405180910390fd5b6200f6bd7fc305c6179e209930eeca61dc9244f9f8a8a146d49c42cf024cbce3ad46e770cd60001b620015a0565b6200f6eb7f2b32de74fe4815d0fbc38e765c24b7e4dde562265af48d7d4dc64f8e1f94655760001b620015a0565b6200f7197fb93a08e573fe5aa186015a0662db7f07457c013e15be8b95b2fb0180ee84b33960001b620015a0565b846101400151816040018181516200f732919062014292565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200f77d7fba009a7cfabb9654f6a672058ee982660885421484fa3af692100e78e11c58ae60001b620015a0565b6200f7ab7f4c158c26e138b54b4cca747b05493d75cb465280964877455f7377751e2e121760001b620015a0565b8460a0015185608001516200f7c191906201420e565b816020018181516200f7d4919062014292565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200f81f7f8a0dbd432c7164af56d3d52de3c6442ab9b7f19ab0425bf8752048414532323f60001b620015a0565b6200f84d7f62cc8b6ed4686abb7418ed14c24ee360a06c7480f9987312914b82c76fd404e260001b620015a0565b8460a0015185608001516200f86391906201420e565b816000018181516200f876919062014191565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050506200fab2565b6200f8c87f200272346cd8bf268e5ac1ade077dee3a786c8cdc1649fd770a1f4162a4b212660001b620015a0565b6200f8f67ff53027ba5569a61a7b0dd189760f02eb5375a43e5c54dd3875002106443fdb7760001b620015a0565b6200f9247f1aa350aeb8445d62dc5c88058de45a7b7342894e086d6c2065aef5fa5781a06760001b620015a0565b6200f9527f2506346691b230a587fcd4c86a7084f357a741648936eafe972d58ab317d28a860001b620015a0565b83610140015167ffffffffffffffff163410156200f9a7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200f99e9062013d25565b60405180910390fd5b6200f9d57f70aa5b90757d4acc23e59215cfe919ce03f38707614047805b6c6d0329c05a3a60001b620015a0565b6200fa037f7ec3882214732c029850267d55f6e9888ce66ddc329c5426504a8546904d3b6760001b620015a0565b6200fa317f5073e3fc43b87f60e13e9ad0141b08f32c8bdf1024c04044b30b17430ed3f05e60001b620015a0565b6001846101e0019060018111156200fa72577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160018111156200faad577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250505b6200fae07f47dc3c728d9aad722386d6ec2973fb608d952a9be7d9e080e6e8b0a80338bc3660001b620015a0565b6200fb0e7fc3df9c86daab4b2736d603ea8137f4d4610c210355438f1f2baea58acb1db35260001b620015a0565b826080015184610180019067ffffffffffffffff16908167ffffffffffffffff16815250506200fb617f3c6d1240d08716de2cd588ced6a63df66c8e327f1a6f3b7a9f08951e82a3dd4560001b620015a0565b6200fb8f7f87a3cb63d3e6cf2e60f4e1c8950c05c246319424f8dfc460d004fbafa74caa9c60001b620015a0565b43846101a00181815250506200fbc87fcac91423f201d19cb2bb455fa3c41359a907f94ea1e7add26bc2177ad9e86aaa60001b620015a0565b6200fbf67f3c6e2e4161b61d98b40030ce95b6fe5758fe384e85d378a59ebb767df9f157ee60001b620015a0565b83600685600001516040516200fc0d9190620138f9565b908152602001604051809103902060008201518160000190805190602001906200fc39929190620109e9565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906200fc9f929190620109e9565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061010082015181600501556101208201518160060160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160060160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101608201518160070190805190602001906200fe15929190620109e9565b506101808201518160080160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600901556101c082015181600a0160006101000a81548160ff0219169083151502179055506101e082015181600a0160016101000a81548160ff021916908360018111156200fec5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061020082015181600a0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061022082015181600b0190805190602001906200ff1992919062010c39565b5061024082015181600c0190805190602001906200ff3992919062010c39565b5061026082015181600d0190805190602001906200ff59929190620109e9565b5061028082015181600e0160006101000a81548160ff021916908360028111156200ffad577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055506102a082015181600e0160016101000a81548160ff0219169083151502179055506102c082015181600f0160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050509050506201009d7fb83d52fa82fd630d69314173c7a974e5dac556af46cfd8c85f26d3e07e98362a60001b620015a0565b620100cb7f57455c5fb2b987f5651ad56e4cbb938381b29f19c89939aff5e7be3115c6b62c60001b620015a0565b600060086000866020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050620101407fc43aa74a0d97c9e21159797ae853b80e81addf72f69c572ff7d793d05721cb1060001b620015a0565b6201016e7f6136151bc8bf0b4879cf3c3cd98f8d64d00fb2c65c28972d86a511d58f538c6960001b620015a0565b808560000151908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190620101ae929190620109e9565b50620101dd7f017e4a80834f76ec8277f87e0d93447b45fd5c8f80684f3e2c2af9c32375befc60001b620015a0565b6201020b7fbd54cfd3a9146eee0afb53d2a3f93d70566c869b4b05917d7c64b52af7863ac160001b620015a0565b60005b85610220015151811015620103b6576201024b7f35e7bdafa21c98c3c025a9afb78d6402f75ceeed14ff3bde2746704013105dd060001b620015a0565b620102797f4d492e388c813287c07aa5b10b2794f6f2d839339ba881d9945163324ed5261560001b620015a0565b6000600860008861022001518481518110620102be577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050620103307f3aefa0f9b9aea7a8d0b1db248d032b2c3e80d088583db7730fe7714ea77cbf0260001b620015a0565b6201035e7f9a16d308a298ce4317b1d35918ba2677ac974857f1f739996ec6c3327de01c2e60001b620015a0565b8087600001519080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906201039e929190620109e9565b50508080620103ad90620144ed565b9150506201020e565b50620103e57f106ea2301b67d3da118c5314e4651798ac4273dbd3ec00bd792618b97653be8e60001b620015a0565b620104137ffd68b719fce5c674a6005a2f4291e039f2bb8fcfe46314aac96420d9a1c8d23160001b620015a0565b60005b85610240015151811015620105be57620104537fc63b5d6e797490c7b610cf05c6c5329548b54f954483839814174f5531ee32a460001b620015a0565b620104817f6a12cb375ef74befdac946b195a1907cfeacd392aca6ed522d26b3c636fb7cae60001b620015a0565b6000600860008861024001518481518110620104c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050620105387f38b264ea93f0b51eec711e14ed38d493733c7771527277519a76d0137f95da5660001b620015a0565b620105667fab578f961e96f7e0d9e980b428e1caeb76293eb2be36024bebda0027ecc3857260001b620015a0565b808760000151908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190620105a6929190620109e9565b50508080620105b590620144ed565b91505062010416565b50620105ed7f7bb3abe63085651b900427d704548b9f6f92dd598300edfabe476ad6da0a682c60001b620015a0565b6201061b7fb9a2bcd169436f695dc40128957614ef38c1cb883d0bd14fed8cd4244aed30b260001b620015a0565b6201062562010db1565b620106537fd93403149fe66c5fa56f3896de7e3c89509ab34a5b86528d0f104eec1fad540e60001b620015a0565b620106817f452678ef4f605b5bab621803e787064ce3668351cc6a6b7f25ed6d28664d060660001b620015a0565b856101200151816000019067ffffffffffffffff16908167ffffffffffffffff1681525050620106d47fbcdff6b64b4dc066d651baa27bdd3dbcd28f299895efbb44c022a28ed28a801360001b620015a0565b620107027f450ce935f17fe5cb13f764fdba8098a418b4dc096fe52b15171ea5e5130604cf60001b620015a0565b6000816020019067ffffffffffffffff16908167ffffffffffffffff1681525050620107517f843b7d9b779e805acbf276a29187d459bc1320f792004023e95bc74f0773216460001b620015a0565b6201077f7f68dd9b25246f7382bf0c9c678f0404b0f1f04695784718dfae534f4be41629ef60001b620015a0565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb4e4e428760000151836040518363ffffffff1660e01b8152600401620107e292919062013a76565b600060405180830381600087803b158015620107fd57600080fd5b505af115801562010812573d6000803e3d6000fd5b50505050620108447f1c017d94f61fad97b863f4014cf61e27747882b50232355319af1e507309e03e60001b620015a0565b620108727fa8b25ac94132cb5de833164098b3f8ab60595b902d380942ca5cf92c35ec00fb60001b620015a0565b7fa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d1860004388600001518961020001518a602001518b61014001518c6102a00151604051620108c7979695949392919062013b52565b60405180910390a1505050505050565b6000620108e430620108ea565b15905090565b600080823b905060008111915050919050565b82805482825590600052602060002090810192821562010951579160200282015b82811115620109505782518290805190602001906201093f929190620109e9565b50916020019190600101906201091e565b5b50905062010960919062010ddf565b5090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b508054620109b19062014481565b6000825580601f10620109c55750620109e6565b601f016020900490600052602060002090810190620109e5919062010e07565b5b50565b828054620109f79062014481565b90600052602060002090601f01602090048101928262010a1b576000855562010a67565b82601f1062010a3657805160ff191683800117855562010a67565b8280016001018555821562010a67579182015b8281111562010a6657825182559160200191906001019062010a49565b5b50905062010a76919062010e07565b5090565b508054600082559060005260206000209081019062010a9a919062010e07565b50565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600081526020016000151581526020016000600181111562010bb5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff1681526020016060815260200160608152602001606081526020016000600281111562010c1b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160001515815260200162010c3362010e26565b81525090565b82805482825590600052602060002090810192821562010cb5579160200282015b8281111562010cb45782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019062010c5a565b5b50905062010cc4919062010e07565b5090565b604051806101e0016040528060608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000151581526020016060815260200160001515815260200160001515815260200160608152602001606081526020016000151581526020016000600181111562010dab577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090565b6040518060400160405280600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b5b8082111562010e03576000818162010df99190620109a3565b5060010162010de0565b5090565b5b8082111562010e2257600081600090555060010162010e08565b5090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600062010e7c62010e768462013eca565b62013ea1565b9050808382526020820190508285602086028201111562010e9c57600080fd5b60005b8581101562010ed0578162010eb5888262011273565b84526020840193506020830192505060018101905062010e9f565b5050509392505050565b600062010ef162010eeb8462013ef9565b62013ea1565b9050808382526020820190508285602086028201111562010f1157600080fd5b60005b8581101562010f6057813567ffffffffffffffff81111562010f3557600080fd5b80860162010f44898262011421565b8552602085019450602084019350505060018101905062010f14565b5050509392505050565b600062010f8162010f7b8462013ef9565b62013ea1565b9050808382526020820190508285602086028201111562010fa157600080fd5b60005b8581101562010ff057815167ffffffffffffffff81111562010fc557600080fd5b80860162010fd489826201144e565b8552602085019450602084019350505060018101905062010fa4565b5050509392505050565b6000620110116201100b8462013f28565b62013ea1565b905080838252602082019050828560208602820111156201103157600080fd5b60005b858110156201106557816201104a88826201151c565b84526020840193506020830192505060018101905062011034565b5050509392505050565b600062011086620110808462013f57565b62013ea1565b90508083825260208201905082856020860282011115620110a657600080fd5b60005b85811015620110f557813567ffffffffffffffff811115620110ca57600080fd5b808601620110d989826201159c565b85526020850194506020840193505050600181019050620110a9565b5050509392505050565b600062011116620111108462013f86565b62013ea1565b905080838252602082019050828560a08602820111156201113657600080fd5b60005b858110156201116a57816201114f888262011b0d565b845260208401935060a0830192505060018101905062011139565b5050509392505050565b60006201118b620111858462013fb5565b62013ea1565b90508083825260208201905082856060860282011115620111ab57600080fd5b60005b85811015620111df5781620111c48882620121b0565b845260208401935060608301925050600181019050620111ae565b5050509392505050565b600062011200620111fa8462013fe4565b62013ea1565b9050828152602081018484840111156201121957600080fd5b620112268482856201443c565b509392505050565b6000620112456201123f8462013fe4565b62013ea1565b9050828152602081018484840111156201125e57600080fd5b6201126b8482856201444b565b509392505050565b600081359050620112848162014898565b92915050565b6000815190506201129b8162014898565b92915050565b600082601f830112620112b357600080fd5b8135620112c584826020860162010e65565b91505092915050565b600082601f830112620112e057600080fd5b8135620112f284826020860162010eda565b91505092915050565b600082601f8301126201130d57600080fd5b81516201131f84826020860162010f6a565b91505092915050565b600082601f8301126201133a57600080fd5b81356201134c84826020860162010ffa565b91505092915050565b600082601f8301126201136757600080fd5b8135620113798482602086016201106f565b91505092915050565b600082601f8301126201139457600080fd5b8151620113a6848260208601620110ff565b91505092915050565b600082601f830112620113c157600080fd5b8135620113d384826020860162011174565b91505092915050565b600081359050620113ed81620148b2565b92915050565b6000815190506201140481620148b2565b92915050565b6000813590506201141b81620148cc565b92915050565b600082601f8301126201143357600080fd5b813562011445848260208601620111e9565b91505092915050565b600082601f8301126201146057600080fd5b8151620114728482602086016201122e565b91505092915050565b6000813590506201148c81620148e6565b92915050565b600081359050620114a38162014900565b92915050565b600081359050620114ba816201491a565b92915050565b600081359050620114d18162014934565b92915050565b600081359050620114e8816201494e565b92915050565b600081359050620114ff8162014968565b92915050565b600081519050620115168162014968565b92915050565b6000813590506201152d8162014979565b92915050565b6000606082840312156201154657600080fd5b62011552606062013ea1565b90506000620115648482850162012247565b60008301525060206201157a8482850162012247565b6020830152506040620115908482850162012247565b60408301525092915050565b60006103208284031215620115b057600080fd5b620115bd6102e062013ea1565b9050600082013567ffffffffffffffff811115620115da57600080fd5b620115e88482850162011421565b6000830152506020620115fe8482850162011273565b602083015250604082013567ffffffffffffffff8111156201161f57600080fd5b6201162d8482850162011421565b6040830152506060620116438482850162012247565b6060830152506080620116598482850162012247565b60808301525060a06201166f8482850162012247565b60a08301525060c0620116858482850162012247565b60c08301525060e06201169b8482850162012247565b60e083015250610100620116b28482850162012219565b61010083015250610120620116ca8482850162012247565b61012083015250610140620116e28482850162012247565b6101408301525061016082013567ffffffffffffffff8111156201170557600080fd5b620117138482850162011421565b610160830152506101806201172b8482850162012247565b610180830152506101a0620117438482850162012219565b6101a0830152506101c06201175b84828501620113dc565b6101c0830152506101e062011773848285016201151c565b6101e0830152506102006201178b8482850162012247565b6102008301525061022082013567ffffffffffffffff811115620117ae57600080fd5b620117bc84828501620112a1565b6102208301525061024082013567ffffffffffffffff811115620117df57600080fd5b620117ed84828501620112a1565b6102408301525061026082013567ffffffffffffffff8111156201181057600080fd5b6201181e8482850162011421565b610260830152506102806201183684828501620114ee565b610280830152506102a06201184e84828501620113dc565b6102a0830152506102c0620118668482850162011a38565b6102c08301525092915050565b6000606082840312156201188657600080fd5b62011892606062013ea1565b9050600082013567ffffffffffffffff811115620118af57600080fd5b620118bd8482850162011421565b6000830152506020620118d38482850162011273565b6020830152506040620118e98482850162012247565b60408301525092915050565b600060e082840312156201190857600080fd5b6201191460e062013ea1565b9050600062011926848285016201225e565b60008301525060206201193c848285016201225e565b602083015250604062011952848285016201225e565b604083015250606062011968848285016201225e565b60608301525060806201197e848285016201225e565b60808301525060a062011994848285016201128a565b60a08301525060c0620119aa848285016201128a565b60c08301525092915050565b600060608284031215620119c957600080fd5b620119d5606062013ea1565b9050600082013567ffffffffffffffff811115620119f257600080fd5b62011a008482850162011421565b600083015250602062011a168482850162011273565b602083015250604062011a2c8482850162011273565b60408301525092915050565b60006060828403121562011a4b57600080fd5b62011a57606062013ea1565b9050600062011a698482850162012247565b600083015250602062011a7f8482850162012247565b602083015250604062011a958482850162012247565b60408301525092915050565b60006040828403121562011ab457600080fd5b62011ac0604062013ea1565b9050600082013567ffffffffffffffff81111562011add57600080fd5b62011aeb8482850162011421565b600083015250602062011b018482850162012247565b60208301525092915050565b600060a0828403121562011b2057600080fd5b62011b2c60a062013ea1565b9050600062011b3e848285016201128a565b600083015250602062011b54848285016201128a565b602083015250604062011b6a848285016201225e565b604083015250606062011b808482850162012230565b606083015250608062011b9684828501620113f3565b60808301525092915050565b6000610180828403121562011bb657600080fd5b62011bc361018062013ea1565b9050600062011bd5848285016201128a565b600083015250602062011beb848285016201225e565b602083015250604062011c01848285016201225e565b604083015250606062011c17848285016201225e565b606083015250608062011c2d8482850162011505565b60808301525060a062011c438482850162012230565b60a08301525060c062011c598482850162012230565b60c08301525060e062011c6f848285016201225e565b60e08301525061010062011c86848285016201225e565b6101008301525061012062011c9e848285016201225e565b6101208301525061014062011cb684828501620113f3565b6101408301525061016082015167ffffffffffffffff81111562011cd957600080fd5b62011ce784828501620112fb565b6101608301525092915050565b6000610160828403121562011d0857600080fd5b62011d1561016062013ea1565b9050600062011d278482850162012247565b600083015250602062011d3d8482850162012247565b602083015250604062011d538482850162012247565b604083015250606062011d698482850162012247565b606083015250608062011d7f8482850162012247565b60808301525060a062011d958482850162012247565b60a08301525060c062011dab8482850162012247565b60c08301525060e062011dc18482850162012247565b60e08301525061010062011dd88482850162012247565b6101008301525061012062011df08482850162012247565b6101208301525061014062011e088482850162012247565b6101408301525092915050565b6000610160828403121562011e2957600080fd5b62011e3661016062013ea1565b9050600062011e48848285016201225e565b600083015250602062011e5e848285016201225e565b602083015250604062011e74848285016201225e565b604083015250606062011e8a848285016201225e565b606083015250608062011ea0848285016201225e565b60808301525060a062011eb6848285016201225e565b60a08301525060c062011ecc848285016201225e565b60c08301525060e062011ee2848285016201225e565b60e08301525061010062011ef9848285016201225e565b6101008301525061012062011f11848285016201225e565b6101208301525061014062011f29848285016201225e565b6101408301525092915050565b60006101e0828403121562011f4a57600080fd5b62011f576101e062013ea1565b9050600082013567ffffffffffffffff81111562011f7457600080fd5b62011f828482850162011421565b600083015250602062011f988482850162012247565b602083015250604062011fae8482850162012247565b604083015250606062011fc48482850162012247565b606083015250608062011fda8482850162012219565b60808301525060a062011ff08482850162012247565b60a08301525060c0620120068482850162012247565b60c08301525060e06201201c84828501620113dc565b60e08301525061010082013567ffffffffffffffff8111156201203e57600080fd5b6201204c8482850162011421565b610100830152506101206201206484828501620113dc565b610120830152506101406201207c84828501620113dc565b6101408301525061016082013567ffffffffffffffff8111156201209f57600080fd5b620120ad8482850162011421565b6101608301525061018082013567ffffffffffffffff811115620120d057600080fd5b620120de84828501620113af565b610180830152506101a0620120f684828501620113dc565b6101a0830152506101c06201210e848285016201151c565b6101c08301525092915050565b600060a082840312156201212e57600080fd5b6201213a60a062013ea1565b905060006201214c848285016201225e565b600083015250602062012162848285016201225e565b602083015250604062012178848285016201225e565b60408301525060606201218e8482850162012230565b6060830152506080620121a48482850162012230565b60808301525092915050565b600060608284031215620121c357600080fd5b620121cf606062013ea1565b90506000620121e18482850162011273565b6000830152506020620121f78482850162012247565b60208301525060406201220d8482850162012247565b60408301525092915050565b6000813590506201222a816201498a565b92915050565b60008151905062012241816201498a565b92915050565b6000813590506201225881620149a4565b92915050565b6000815190506201226f81620149a4565b92915050565b6000602082840312156201228857600080fd5b6000620122988482850162011273565b91505092915050565b60008060408385031215620122b557600080fd5b6000620122c58582860162011273565b925050602083013567ffffffffffffffff811115620122e357600080fd5b620122f185828601620112ce565b9150509250929050565b600080604083850312156201230f57600080fd5b60006201231f8582860162011273565b925050602083013567ffffffffffffffff8111156201233d57600080fd5b6201234b8582860162011421565b9150509250929050565b6000602082840312156201236857600080fd5b600082013567ffffffffffffffff8111156201238357600080fd5b6201239184828501620112ce565b91505092915050565b600080600060608486031215620123b057600080fd5b600084013567ffffffffffffffff811115620123cb57600080fd5b620123d986828701620112ce565b9350506020620123ec8682870162011273565b925050604084013567ffffffffffffffff8111156201240a57600080fd5b620124188682870162011328565b9150509250925092565b60008060006101a084860312156201243957600080fd5b600084013567ffffffffffffffff8111156201245457600080fd5b6201246286828701620112ce565b9350506020620124758682870162011cf4565b925050610180620124898682870162012219565b9150509250925092565b600060208284031215620124a657600080fd5b600082013567ffffffffffffffff811115620124c157600080fd5b620124cf8482850162011355565b91505092915050565b600060208284031215620124eb57600080fd5b600082015167ffffffffffffffff8111156201250657600080fd5b620125148482850162011382565b91505092915050565b6000602082840312156201253057600080fd5b600062012540848285016201140a565b91505092915050565b6000602082840312156201255c57600080fd5b600082013567ffffffffffffffff8111156201257757600080fd5b620125858482850162011421565b91505092915050565b6000806000806000806101008789031215620125a957600080fd5b6000620125b989828a016201147b565b9650506020620125cc89828a0162011492565b9550506040620125df89828a01620114d7565b9450506060620125f289828a01620114c0565b93505060806201260589828a01620114a9565b92505060a06201261889828a0162011533565b9150509295509295509295565b6000602082840312156201263857600080fd5b600082013567ffffffffffffffff8111156201265357600080fd5b62012661848285016201159c565b91505092915050565b6000806000606084860312156201268057600080fd5b600084013567ffffffffffffffff8111156201269b57600080fd5b620126a9868287016201159c565b9350506020620126bc86828701620113dc565b9250506040620126cf86828701620113dc565b9150509250925092565b600060208284031215620126ec57600080fd5b600082013567ffffffffffffffff8111156201270757600080fd5b620127158482850162011873565b91505092915050565b600060e082840312156201273157600080fd5b60006201274184828501620118f5565b91505092915050565b6000602082840312156201275d57600080fd5b600082013567ffffffffffffffff8111156201277857600080fd5b6201278684828501620119b6565b91505092915050565b600060208284031215620127a257600080fd5b600082013567ffffffffffffffff811115620127bd57600080fd5b620127cb8482850162011aa1565b91505092915050565b600060208284031215620127e757600080fd5b600082015167ffffffffffffffff8111156201280257600080fd5b620128108482850162011ba2565b91505092915050565b600061016082840312156201282d57600080fd5b60006201283d8482850162011e15565b91505092915050565b60008060006101a084860312156201285d57600080fd5b60006201286d8682870162011cf4565b935050610160620128818682870162012219565b92505061018084013567ffffffffffffffff811115620128a057600080fd5b620128ae868287016201159c565b9150509250925092565b6000806101808385031215620128cd57600080fd5b6000620128dd8582860162011cf4565b925050610160620128f18582860162012247565b9150509250929050565b60008060006101a084860312156201291257600080fd5b6000620129228682870162011cf4565b935050610160620129368682870162012247565b9250506101806201294a8682870162012247565b9150509250925092565b6000806000806101c085870312156201296c57600080fd5b60006201297c8782880162011cf4565b945050610160620129908782880162012247565b935050610180620129a48782880162012247565b9250506101a0620129b88782880162012247565b91505092959194509250565b60008060008060006101e08688031215620129de57600080fd5b6000620129ee8882890162011cf4565b95505061016062012a028882890162012247565b94505061018062012a168882890162012247565b9350506101a062012a2a8882890162012247565b9250506101c062012a3e8882890162012247565b9150509295509295909350565b60006020828403121562012a5e57600080fd5b600082013567ffffffffffffffff81111562012a7957600080fd5b62012a878482850162011f36565b91505092915050565b60008060006101a0848603121562012aa757600080fd5b600084013567ffffffffffffffff81111562012ac257600080fd5b62012ad08682870162011f36565b935050602062012ae38682870162011cf4565b92505061018062012af78682870162012219565b9150509250925092565b6000806040838503121562012b1557600080fd5b600083013567ffffffffffffffff81111562012b3057600080fd5b62012b3e8582860162011f36565b925050602062012b518582860162012219565b9150509250929050565b600060a0828403121562012b6e57600080fd5b600062012b7e848285016201211b565b91505092915050565b600062012b95838362012bcd565b60208301905092915050565b600062012baf838362012dfe565b905092915050565b600062012bc583836201310c565b905092915050565b62012bd881620142cd565b82525050565b62012be981620142cd565b82525050565b600062012bfc826201405f565b62012c088185620140b2565b935062012c15836201401a565b8060005b8381101562012c4c57815162012c30888262012b87565b975062012c3d836201408b565b92505060018101905062012c19565b5085935050505092915050565b600062012c66826201406a565b62012c728185620140c3565b93508360208202850162012c86856201402a565b8060005b8581101562012cc8578484038952815162012ca6858262012ba1565b945062012cb38362014098565b925060208a0199505060018101905062012c8a565b50829750879550505050505092915050565b600062012ce7826201406a565b62012cf38185620140d4565b93508360208202850162012d07856201402a565b8060005b8581101562012d49578484038952815162012d27858262012ba1565b945062012d348362014098565b925060208a0199505060018101905062012d0b565b50829750879550505050505092915050565b600062012d688262014075565b62012d748185620140e5565b93508360208202850162012d88856201403a565b8060005b8581101562012dca578484038952815162012da8858262012bb7565b945062012db583620140a5565b925060208a0199505060018101905062012d8c565b50829750879550505050505092915050565b62012de781620142e1565b82525050565b62012df881620142e1565b82525050565b600062012e0b8262014080565b62012e178185620140f6565b935062012e298185602086016201444b565b62012e34816201465c565b840191505092915050565b600062012e4c8262014080565b62012e58818562014107565b935062012e6a8185602086016201444b565b62012e75816201465c565b840191505092915050565b600062012e8d8262014080565b62012e99818562014118565b935062012eab8185602086016201444b565b80840191505092915050565b6000815462012ec68162014481565b62012ed2818662014118565b9450600182166000811462012ef0576001811462012f025762012f39565b60ff1983168652818601935062012f39565b62012f0d856201404a565b60005b8381101562012f315781548189015260018201915060208101905062012f10565b838801955050505b50505092915050565b62012f4d81620143d8565b82525050565b62012f5e81620143ec565b82525050565b62012f6f8162014400565b82525050565b62012f808162014414565b82525050565b600062012f9560118362014123565b915062012fa2826201466d565b602082019050919050565b600062012fbc600c8362014123565b915062012fc98262014696565b602082019050919050565b600062012fe3602b8362014123565b915062012ff082620146bf565b604082019050919050565b60006201300a600e8362014123565b915062013017826201470e565b602082019050919050565b600062013031602e8362014123565b91506201303e8262014737565b604082019050919050565b600062013058600f8362014123565b9150620130658262014786565b602082019050919050565b60006201307f601f8362014123565b91506201308c82620147af565b602082019050919050565b6000620130a6601a8362014123565b9150620130b382620147d8565b602082019050919050565b6000620130cd60128362014123565b9150620130da8262014801565b602082019050919050565b6000620130f460148362014123565b915062013101826201482a565b602082019050919050565b60006103208301600083015184820360008601526201312c828262012dfe565b915050602083015162013143602086018262012bcd565b50604083015184820360408601526201315d828262012dfe565b9150506060830151620131746060860182620138d7565b506080830151620131896080860182620138d7565b5060a08301516201319e60a0860182620138d7565b5060c0830151620131b360c0860182620138d7565b5060e0830151620131c860e0860182620138d7565b50610100830151620131df610100860182620138a4565b50610120830151620131f6610120860182620138d7565b506101408301516201320d610140860182620138d7565b5061016083015184820361016086015262013229828262012dfe565b91505061018083015162013242610180860182620138d7565b506101a0830151620132596101a0860182620138a4565b506101c0830151620132706101c086018262012ddc565b506101e0830151620132876101e086018262012f64565b506102008301516201329e610200860182620138d7565b50610220830151848203610220860152620132ba828262012bef565b915050610240830151848203610240860152620132d8828262012bef565b915050610260830151848203610260860152620132f6828262012dfe565b9150506102808301516201330f61028086018262012f53565b506102a0830151620133266102a086018262012ddc565b506102c08301516201333d6102c086018262013620565b508091505092915050565b600061032083016000830151848203600086015262013368828262012dfe565b91505060208301516201337f602086018262012bcd565b506040830151848203604086015262013399828262012dfe565b9150506060830151620133b06060860182620138d7565b506080830151620133c56080860182620138d7565b5060a0830151620133da60a0860182620138d7565b5060c0830151620133ef60c0860182620138d7565b5060e08301516201340460e0860182620138d7565b506101008301516201341b610100860182620138a4565b5061012083015162013432610120860182620138d7565b5061014083015162013449610140860182620138d7565b5061016083015184820361016086015262013465828262012dfe565b9150506101808301516201347e610180860182620138d7565b506101a0830151620134956101a0860182620138a4565b506101c0830151620134ac6101c086018262012ddc565b506101e0830151620134c36101e086018262012f64565b50610200830151620134da610200860182620138d7565b50610220830151848203610220860152620134f6828262012bef565b91505061024083015184820361024086015262013514828262012bef565b91505061026083015184820361026086015262013532828262012dfe565b9150506102808301516201354b61028086018262012f53565b506102a0830151620135626102a086018262012ddc565b506102c0830151620135796102c086018262013620565b508091505092915050565b60e0820160008201516201359c6000850182620138d7565b506020820151620135b16020850182620138d7565b506040820151620135c66040850182620138d7565b506060820151620135db6060850182620138d7565b506080820151620135f06080850182620138d7565b5060a08201516201360560a085018262012bcd565b5060c08201516201361a60c085018262012bcd565b50505050565b606082016000820151620136386000850182620138d7565b5060208201516201364d6020850182620138d7565b506040820151620136626040850182620138d7565b50505050565b604082016000820151620136806000850182620138d7565b506020820151620136956020850182620138d7565b50505050565b600061018083016000830151620136b6600086018262012bcd565b506020830151620136cb6020860182620138d7565b506040830151620136e06040860182620138d7565b506060830151620136f56060860182620138d7565b5060808301516201370a608086018262012f53565b5060a08301516201371f60a0860182620138a4565b5060c08301516201373460c0860182620138a4565b5060e08301516201374960e0860182620138d7565b5061010083015162013760610100860182620138d7565b5061012083015162013777610120860182620138d7565b506101408301516201378e61014086018262012ddc565b50610160830151848203610160860152620137aa828262012c59565b9150508091505092915050565b604082016000820151620137cf600085018262012bcd565b506020820151620137e46020850182620138d7565b50505050565b606082016000820151620138026000850182620138d7565b506020820151620138176020850182620138d7565b5060408201516201382c6040850182620138d7565b50505050565b60a0820160008201516201384a6000850182620138d7565b5060208201516201385f6020850182620138d7565b506040820151620138746040850182620138d7565b506060820151620138896060850182620138a4565b5060808201516201389e6080850182620138a4565b50505050565b620138af81620143ba565b82525050565b620138c081620143ba565b82525050565b620138d18162014428565b82525050565b620138e281620143c4565b82525050565b620138f381620143c4565b82525050565b600062013907828462012e80565b915081905092915050565b600062013920828462012eb7565b915081905092915050565b600060208201905062013942600083018462012bde565b92915050565b600060c0820190506201395f600083018562012bde565b6201396e602083018462013832565b9392505050565b6000602082019050818103600083015262013991818462012cda565b905092915050565b60006060820190508181036000830152620139b5818662012cda565b9050620139c66020830185620138e8565b620139d5604083018462012ded565b949350505050565b60006020820190508181036000830152620139f9818462012d5b565b905092915050565b6000604082019050818103600083015262013a1d818562012d5b565b905062013a2e602083018462012ded565b9392505050565b600060208201905062013a4c600083018462012ded565b92915050565b6000602082019050818103600083015262013a6e818462012e3f565b905092915050565b6000606082019050818103600083015262013a92818562012e3f565b905062013aa3602083018462013668565b9392505050565b600060808201905062013ac1600083018762012f42565b62013ad06020830186620138b5565b818103604083015262013ae4818562012cda565b905062013af5606083018462012bde565b95945050505050565b600060808201905062013b15600083018762012f42565b62013b246020830186620138b5565b818103604083015262013b38818562012e3f565b905062013b49606083018462012bde565b95945050505050565b600060e08201905062013b69600083018a62012f42565b62013b786020830189620138b5565b818103604083015262013b8c818862012e3f565b905062013b9d6060830187620138e8565b62013bac608083018662012bde565b62013bbb60a0830185620138e8565b62013bca60c083018462012ded565b98975050505050505050565b600060208201905062013bed600083018462012f75565b92915050565b6000602082019050818103600083015262013c0e8162012f86565b9050919050565b6000602082019050818103600083015262013c308162012fad565b9050919050565b6000602082019050818103600083015262013c528162012fd4565b9050919050565b6000602082019050818103600083015262013c748162012ffb565b9050919050565b6000602082019050818103600083015262013c968162013022565b9050919050565b6000602082019050818103600083015262013cb88162013049565b9050919050565b6000602082019050818103600083015262013cda8162013070565b9050919050565b6000602082019050818103600083015262013cfc8162013097565b9050919050565b6000602082019050818103600083015262013d1e81620130be565b9050919050565b6000602082019050818103600083015262013d4081620130e5565b9050919050565b6000602082019050818103600083015262013d63818462013348565b905092915050565b600060e08201905062013d82600083018462013584565b92915050565b6000604082019050818103600083015262013da481856201369b565b9050818103602083015262013dba818462013348565b90509392505050565b600060408201905062013dda6000830184620137b7565b92915050565b600060608201905062013df76000830184620137ea565b92915050565b600060408201905062013e146000830185620138b5565b62013e236020830184620138b5565b9392505050565b600060408201905062013e416000830185620138b5565b62013e506020830184620138c6565b9392505050565b600060208201905062013e6e6000830184620138e8565b92915050565b600060408201905062013e8b6000830185620138c6565b62013e9a6020830184620138c6565b9392505050565b600062013ead62013ec0565b905062013ebb8282620144b7565b919050565b6000604051905090565b600067ffffffffffffffff82111562013ee85762013ee76201462d565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013f175762013f166201462d565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013f465762013f456201462d565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013f755762013f746201462d565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013fa45762013fa36201462d565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013fd35762013fd26201462d565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156201400257620140016201462d565b5b6201400d826201465c565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b60006201414182620143ba565b91506201414e83620143ba565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562014186576201418562014571565b5b828201905092915050565b60006201419e82620143c4565b9150620141ab83620143c4565b92508267ffffffffffffffff03821115620141cb57620141ca62014571565b5b828201905092915050565b6000620141e382620143c4565b9150620141f083620143c4565b925082620142035762014202620145a0565b5b828204905092915050565b60006201421b82620143c4565b91506201422883620143c4565b92508167ffffffffffffffff04831182151516156201424c576201424b62014571565b5b828202905092915050565b60006201426482620143ba565b91506201427183620143ba565b92508282101562014287576201428662014571565b5b828203905092915050565b60006201429f82620143c4565b9150620142ac83620143c4565b925082821015620142c257620142c162014571565b5b828203905092915050565b6000620142da826201439a565b9050919050565b60008115159050919050565b6000819050919050565b60006201430482620142cd565b9050919050565b60006201431882620142cd565b9050919050565b60006201432c82620142cd565b9050919050565b60006201434082620142cd565b9050919050565b60006201435482620142cd565b9050919050565b60008190506201436b8262014853565b919050565b600081905062014380826201486a565b919050565b6000819050620143958262014881565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000620143e5826201435b565b9050919050565b6000620143f98262014370565b9050919050565b60006201440d8262014385565b9050919050565b60006201442182620143c4565b9050919050565b60006201443582620143c4565b9050919050565b82818337600083830152505050565b60005b838110156201446b5780820151818401526020810190506201444e565b838111156201447b576000848401525b50505050565b600060028204905060018216806201449a57607f821691505b60208210811415620144b157620144b0620145fe565b5b50919050565b620144c2826201465c565b810181811067ffffffffffffffff82111715620144e457620144e36201462d565b5b80604052505050565b6000620144fa82620143ba565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562014530576201452f62014571565b5b600182019050919050565b60006201454882620143c4565b915067ffffffffffffffff82141562014566576201456562014571565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f66696c654c69737420697320656d707479000000000000000000000000000000600082015250565b7f66696c6520657870697265640000000000000000000000000000000000000000600082015250565b7f43757272656e74206f776e6572206d75737420626520746865206f776e65722060008201527f6f66207468652066696c65000000000000000000000000000000000000000000602082015250565b7f66696c65206e6f74206578697374000000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f66696c652074797065206572726f720000000000000000000000000000000000600082015250565b7f66696c6553697a65206d7573742062652067726561746572207468616e203000600082015250565b7f66696c6548617368206d757374206265206e6f7420656d707479000000000000600082015250565b7f66696c6520616c72656164792065786973740000000000000000000000000000600082015250565b7f696e73756666696369656e74206465706f736974000000000000000000000000600082015250565b600a8110620148675762014866620145cf565b5b50565b600381106201487e576201487d620145cf565b5b50565b60028110620148955762014894620145cf565b5b50565b620148a381620142cd565b8114620148af57600080fd5b50565b620148bd81620142e1565b8114620148c957600080fd5b50565b620148d781620142ed565b8114620148e357600080fd5b50565b620148f181620142f7565b8114620148fd57600080fd5b50565b6201490b816201430b565b81146201491757600080fd5b50565b62014925816201431f565b81146201493157600080fd5b50565b6201493f8162014333565b81146201494b57600080fd5b50565b620149598162014347565b81146201496557600080fd5b50565b600381106201497657600080fd5b50565b600281106201498757600080fd5b50565b6201499581620143ba565b8114620149a157600080fd5b50565b620149af81620143c4565b8114620149bb57600080fd5b5056fea2646970667358221220fd836012301de72d776eb535f20ffc7e145b5fb331f21b746eb04d130f9c210464736f6c63430008040033",
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

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCaller) CalcDepositFee(opts *bind.CallOpts, uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalcDepositFee", uploadOption, setting, currentHeight)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns((uint64,uint64,uint64))
func (_Store *StoreCaller) CalcFee(opts *bind.CallOpts, setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalcFee", setting, proveTime, copyNum, fileSize, duration)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
}

// CalcProveTimesByUploadInfo is a free data retrieval call binding the contract method 0xe4bca973.
//
// Solidity: function CalcProveTimesByUploadInfo((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, uint256 beginHeight) pure returns(uint64)
func (_Store *StoreCaller) CalcProveTimesByUploadInfo(opts *bind.CallOpts, uploadOption UploadOption, beginHeight *big.Int) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalcProveTimesByUploadInfo", uploadOption, beginHeight)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcProveTimesByUploadInfo is a free data retrieval call binding the contract method 0xe4bca973.
//
// Solidity: function CalcProveTimesByUploadInfo((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, uint256 beginHeight) pure returns(uint64)
func (_Store *StoreSession) CalcProveTimesByUploadInfo(uploadOption UploadOption, beginHeight *big.Int) (uint64, error) {
	return _Store.Contract.CalcProveTimesByUploadInfo(&_Store.CallOpts, uploadOption, beginHeight)
}

// CalcProveTimesByUploadInfo is a free data retrieval call binding the contract method 0xe4bca973.
//
// Solidity: function CalcProveTimesByUploadInfo((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, uint256 beginHeight) pure returns(uint64)
func (_Store *StoreCallerSession) CalcProveTimesByUploadInfo(uploadOption UploadOption, beginHeight *big.Int) (uint64, error) {
	return _Store.Contract.CalcProveTimesByUploadInfo(&_Store.CallOpts, uploadOption, beginHeight)
}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreCaller) GetFileInfo(opts *bind.CallOpts, fileHash []byte) (FileInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetFileInfo", fileHash)

	if err != nil {
		return *new(FileInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FileInfo)).(*FileInfo)

	return out0, err

}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreCallerSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreCaller) GetFileInfos(opts *bind.CallOpts, _fileList [][]byte) ([]FileInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetFileInfos", _fileList)

	if err != nil {
		return *new([]FileInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]FileInfo)).(*[]FileInfo)

	return out0, err

}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreSession) GetFileInfos(_fileList [][]byte) ([]FileInfo, error) {
	return _Store.Contract.GetFileInfos(&_Store.CallOpts, _fileList)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreCallerSession) GetFileInfos(_fileList [][]byte) ([]FileInfo, error) {
	return _Store.Contract.GetFileInfos(&_Store.CallOpts, _fileList)
}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetFileList(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetFileList", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetFileList(&_Store.CallOpts, walletAddr)
}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetFileList(&_Store.CallOpts, walletAddr)
}

// GetUnProveCandidateFiles is a free data retrieval call binding the contract method 0xdc1ec30b.
//
// Solidity: function GetUnProveCandidateFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetUnProveCandidateFiles(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnProveCandidateFiles", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetUnProveCandidateFiles is a free data retrieval call binding the contract method 0xdc1ec30b.
//
// Solidity: function GetUnProveCandidateFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetUnProveCandidateFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProveCandidateFiles(&_Store.CallOpts, walletAddr)
}

// GetUnProveCandidateFiles is a free data retrieval call binding the contract method 0xdc1ec30b.
//
// Solidity: function GetUnProveCandidateFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetUnProveCandidateFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProveCandidateFiles(&_Store.CallOpts, walletAddr)
}

// GetUnProvePrimaryFiles is a free data retrieval call binding the contract method 0x8d41f9f8.
//
// Solidity: function GetUnProvePrimaryFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetUnProvePrimaryFiles(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnProvePrimaryFiles", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetUnProvePrimaryFiles is a free data retrieval call binding the contract method 0x8d41f9f8.
//
// Solidity: function GetUnProvePrimaryFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetUnProvePrimaryFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProvePrimaryFiles(&_Store.CallOpts, walletAddr)
}

// GetUnProvePrimaryFiles is a free data retrieval call binding the contract method 0x8d41f9f8.
//
// Solidity: function GetUnProvePrimaryFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetUnProvePrimaryFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProvePrimaryFiles(&_Store.CallOpts, walletAddr)
}

// GetUnSettledFileList is a free data retrieval call binding the contract method 0x28a8565c.
//
// Solidity: function GetUnSettledFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetUnSettledFileList(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnSettledFileList", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetUnSettledFileList is a free data retrieval call binding the contract method 0x28a8565c.
//
// Solidity: function GetUnSettledFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetUnSettledFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnSettledFileList(&_Store.CallOpts, walletAddr)
}

// GetUnSettledFileList is a free data retrieval call binding the contract method 0x28a8565c.
//
// Solidity: function GetUnSettledFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetUnSettledFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnSettledFileList(&_Store.CallOpts, walletAddr)
}

// GetUploadStorageFee is a free data retrieval call binding the contract method 0x41bc86cb.
//
// Solidity: function GetUploadStorageFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption) view returns((uint64,uint64,uint64))
func (_Store *StoreCaller) GetUploadStorageFee(opts *bind.CallOpts, uploadOption UploadOption) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUploadStorageFee", uploadOption)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// GetUploadStorageFee is a free data retrieval call binding the contract method 0x41bc86cb.
//
// Solidity: function GetUploadStorageFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) GetUploadStorageFee(uploadOption UploadOption) (StorageFee, error) {
	return _Store.Contract.GetUploadStorageFee(&_Store.CallOpts, uploadOption)
}

// GetUploadStorageFee is a free data retrieval call binding the contract method 0x41bc86cb.
//
// Solidity: function GetUploadStorageFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetUploadStorageFee(uploadOption UploadOption) (StorageFee, error) {
	return _Store.Contract.GetUploadStorageFee(&_Store.CallOpts, uploadOption)
}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0x740f6913.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCaller) UpdateFileInfoForRenew(opts *bind.CallOpts, setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "UpdateFileInfoForRenew", setting, newExpireHeight, fileInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0x740f6913.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0x740f6913.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCallerSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreCaller) UpdateFilesForRenew(opts *bind.CallOpts, _fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "UpdateFilesForRenew", _fileList, setting, newExpireHeight)

	if err != nil {
		return *new([]FileInfo), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]FileInfo)).(*[]FileInfo)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreCallerSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
}

// C0xf932268c is a free data retrieval call binding the contract method 0x0df7e99e.
//
// Solidity: function c_0xf932268c(bytes32 c__0xf932268c) pure returns()
func (_Store *StoreCaller) C0xf932268c(opts *bind.CallOpts, c__0xf932268c [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0xf932268c", c__0xf932268c)

	if err != nil {
		return err
	}

	return err

}

// C0xf932268c is a free data retrieval call binding the contract method 0x0df7e99e.
//
// Solidity: function c_0xf932268c(bytes32 c__0xf932268c) pure returns()
func (_Store *StoreSession) C0xf932268c(c__0xf932268c [32]byte) error {
	return _Store.Contract.C0xf932268c(&_Store.CallOpts, c__0xf932268c)
}

// C0xf932268c is a free data retrieval call binding the contract method 0x0df7e99e.
//
// Solidity: function c_0xf932268c(bytes32 c__0xf932268c) pure returns()
func (_Store *StoreCallerSession) C0xf932268c(c__0xf932268c [32]byte) error {
	return _Store.Contract.C0xf932268c(&_Store.CallOpts, c__0xf932268c)
}

// CalcSingleValidFeeForFile is a free data retrieval call binding the contract method 0xca6142cb.
//
// Solidity: function calcSingleValidFeeForFile((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcSingleValidFeeForFile(opts *bind.CallOpts, setting Setting, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcSingleValidFeeForFile", setting, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcSingleValidFeeForFile is a free data retrieval call binding the contract method 0xca6142cb.
//
// Solidity: function calcSingleValidFeeForFile((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcSingleValidFeeForFile(setting Setting, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcSingleValidFeeForFile(&_Store.CallOpts, setting, fileSize)
}

// CalcSingleValidFeeForFile is a free data retrieval call binding the contract method 0xca6142cb.
//
// Solidity: function calcSingleValidFeeForFile((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcSingleValidFeeForFile(setting Setting, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcSingleValidFeeForFile(&_Store.CallOpts, setting, fileSize)
}

// CalcStorageFee is a free data retrieval call binding the contract method 0x70b1d889.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCaller) CalcStorageFee(opts *bind.CallOpts, setting Setting, copyNum uint64, fileSize uint64, duration uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcStorageFee", setting, copyNum, fileSize, duration)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcStorageFee is a free data retrieval call binding the contract method 0x70b1d889.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreSession) CalcStorageFee(setting Setting, copyNum uint64, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFee(&_Store.CallOpts, setting, copyNum, fileSize, duration)
}

// CalcStorageFee is a free data retrieval call binding the contract method 0x70b1d889.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCallerSession) CalcStorageFee(setting Setting, copyNum uint64, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFee(&_Store.CallOpts, setting, copyNum, fileSize, duration)
}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0xdf49410a.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCaller) CalcStorageFeeForOneNode(opts *bind.CallOpts, setting Setting, fileSize uint64, duration uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcStorageFeeForOneNode", setting, fileSize, duration)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0xdf49410a.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreSession) CalcStorageFeeForOneNode(setting Setting, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFeeForOneNode(&_Store.CallOpts, setting, fileSize, duration)
}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0xdf49410a.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCallerSession) CalcStorageFeeForOneNode(setting Setting, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFeeForOneNode(&_Store.CallOpts, setting, fileSize, duration)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCaller) CalcUploadFee(opts *bind.CallOpts, uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcUploadFee", uploadOption, setting, currentHeight)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcValidFee is a free data retrieval call binding the contract method 0x170bbdf5.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcValidFee(opts *bind.CallOpts, setting Setting, proveTime uint64, copyNum uint64, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcValidFee", setting, proveTime, copyNum, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcValidFee is a free data retrieval call binding the contract method 0x170bbdf5.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcValidFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize)
}

// CalcValidFee is a free data retrieval call binding the contract method 0x170bbdf5.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcValidFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize)
}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x793cab8e.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcValidFeeForOneNode(opts *bind.CallOpts, setting Setting, proveTime uint64, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcValidFeeForOneNode", setting, proveTime, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x793cab8e.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcValidFeeForOneNode(setting Setting, proveTime uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFeeForOneNode(&_Store.CallOpts, setting, proveTime, fileSize)
}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x793cab8e.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcValidFeeForOneNode(setting Setting, proveTime uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFeeForOneNode(&_Store.CallOpts, setting, proveTime, fileSize)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) returns()
func (_Store *StoreTransactor) AddFileToUnSettleList(opts *bind.TransactOpts, fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToUnSettleList", fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) returns()
func (_Store *StoreSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) returns()
func (_Store *StoreTransactorSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactor) ChangeFileOwner(opts *bind.TransactOpts, ownerChange OwnerChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFileOwner", ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreSession) ChangeFileOwner(ownerChange OwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactorSession) ChangeFileOwner(ownerChange OwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactor) ChangeFilePrivilege(opts *bind.TransactOpts, priChange PriChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFilePrivilege", priChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreSession) ChangeFilePrivilege(priChange PriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePrivilege(&_Store.TransactOpts, priChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactorSession) ChangeFilePrivilege(priChange PriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePrivilege(&_Store.TransactOpts, priChange)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreTransactor) CleanupForDeleteFile(opts *bind.TransactOpts, fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CleanupForDeleteFile", fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreTransactorSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// DelFileFromCandidateList is a paid mutator transaction binding the contract method 0x1bc55880.
//
// Solidity: function DelFileFromCandidateList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromCandidateList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromCandidateList", walletAddr, fileHash)
}

// DelFileFromCandidateList is a paid mutator transaction binding the contract method 0x1bc55880.
//
// Solidity: function DelFileFromCandidateList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromCandidateList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromCandidateList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromCandidateList is a paid mutator transaction binding the contract method 0x1bc55880.
//
// Solidity: function DelFileFromCandidateList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromCandidateList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromCandidateList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromList is a paid mutator transaction binding the contract method 0xebaf0258.
//
// Solidity: function DelFileFromList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromList", walletAddr, fileHash)
}

// DelFileFromList is a paid mutator transaction binding the contract method 0xebaf0258.
//
// Solidity: function DelFileFromList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromList is a paid mutator transaction binding the contract method 0xebaf0258.
//
// Solidity: function DelFileFromList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromPrimaryList is a paid mutator transaction binding the contract method 0x655c6694.
//
// Solidity: function DelFileFromPrimaryList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromPrimaryList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromPrimaryList", walletAddr, fileHash)
}

// DelFileFromPrimaryList is a paid mutator transaction binding the contract method 0x655c6694.
//
// Solidity: function DelFileFromPrimaryList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromPrimaryList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromPrimaryList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromPrimaryList is a paid mutator transaction binding the contract method 0x655c6694.
//
// Solidity: function DelFileFromPrimaryList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromPrimaryList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromPrimaryList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromUnSettledList is a paid mutator transaction binding the contract method 0xb64ab1ef.
//
// Solidity: function DelFileFromUnSettledList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromUnSettledList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromUnSettledList", walletAddr, fileHash)
}

// DelFileFromUnSettledList is a paid mutator transaction binding the contract method 0xb64ab1ef.
//
// Solidity: function DelFileFromUnSettledList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromUnSettledList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromUnSettledList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromUnSettledList is a paid mutator transaction binding the contract method 0xb64ab1ef.
//
// Solidity: function DelFileFromUnSettledList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromUnSettledList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromUnSettledList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DeleteExpiredFilesFromList is a paid mutator transaction binding the contract method 0x3ad525a9.
//
// Solidity: function DeleteExpiredFilesFromList(bytes[] _fileList, address walletAddr, uint8[] storageType) returns(bytes[], uint64, bool)
func (_Store *StoreTransactor) DeleteExpiredFilesFromList(opts *bind.TransactOpts, _fileList [][]byte, walletAddr common.Address, storageType []uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteExpiredFilesFromList", _fileList, walletAddr, storageType)
}

// DeleteExpiredFilesFromList is a paid mutator transaction binding the contract method 0x3ad525a9.
//
// Solidity: function DeleteExpiredFilesFromList(bytes[] _fileList, address walletAddr, uint8[] storageType) returns(bytes[], uint64, bool)
func (_Store *StoreSession) DeleteExpiredFilesFromList(_fileList [][]byte, walletAddr common.Address, storageType []uint8) (*types.Transaction, error) {
	return _Store.Contract.DeleteExpiredFilesFromList(&_Store.TransactOpts, _fileList, walletAddr, storageType)
}

// DeleteExpiredFilesFromList is a paid mutator transaction binding the contract method 0x3ad525a9.
//
// Solidity: function DeleteExpiredFilesFromList(bytes[] _fileList, address walletAddr, uint8[] storageType) returns(bytes[], uint64, bool)
func (_Store *StoreTransactorSession) DeleteExpiredFilesFromList(_fileList [][]byte, walletAddr common.Address, storageType []uint8) (*types.Transaction, error) {
	return _Store.Contract.DeleteExpiredFilesFromList(&_Store.TransactOpts, _fileList, walletAddr, storageType)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteFile(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFile", fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) returns()
func (_Store *StoreSession) DeleteFile(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFile(&_Store.TransactOpts, fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteFile(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFile(&_Store.TransactOpts, fileHash)
}

// DeleteFileInfo is a paid mutator transaction binding the contract method 0x95b0b54b.
//
// Solidity: function DeleteFileInfo(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteFileInfo(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFileInfo", fileHash)
}

// DeleteFileInfo is a paid mutator transaction binding the contract method 0x95b0b54b.
//
// Solidity: function DeleteFileInfo(bytes fileHash) returns()
func (_Store *StoreSession) DeleteFileInfo(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileInfo(&_Store.TransactOpts, fileHash)
}

// DeleteFileInfo is a paid mutator transaction binding the contract method 0x95b0b54b.
//
// Solidity: function DeleteFileInfo(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteFileInfo(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileInfo(&_Store.TransactOpts, fileHash)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) returns()
func (_Store *StoreTransactor) DeleteFiles(opts *bind.TransactOpts, fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFiles", fileHashs)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) returns()
func (_Store *StoreSession) DeleteFiles(fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFiles(&_Store.TransactOpts, fileHashs)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) returns()
func (_Store *StoreTransactorSession) DeleteFiles(fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFiles(&_Store.TransactOpts, fileHashs)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) returns()
func (_Store *StoreTransactor) DeleteUnsettledFiles(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteUnsettledFiles", walletAddr)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) returns()
func (_Store *StoreSession) DeleteUnsettledFiles(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUnsettledFiles(&_Store.TransactOpts, walletAddr)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) returns()
func (_Store *StoreTransactorSession) DeleteUnsettledFiles(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUnsettledFiles(&_Store.TransactOpts, walletAddr)
}

// FileReNew is a paid mutator transaction binding the contract method 0x57d94399.
//
// Solidity: function FileReNew((bytes,address,uint64) fileReNewInfo) payable returns()
func (_Store *StoreTransactor) FileReNew(opts *bind.TransactOpts, fileReNewInfo FileReNewInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileReNew", fileReNewInfo)
}

// FileReNew is a paid mutator transaction binding the contract method 0x57d94399.
//
// Solidity: function FileReNew((bytes,address,uint64) fileReNewInfo) payable returns()
func (_Store *StoreSession) FileReNew(fileReNewInfo FileReNewInfo) (*types.Transaction, error) {
	return _Store.Contract.FileReNew(&_Store.TransactOpts, fileReNewInfo)
}

// FileReNew is a paid mutator transaction binding the contract method 0x57d94399.
//
// Solidity: function FileReNew((bytes,address,uint64) fileReNewInfo) payable returns()
func (_Store *StoreTransactorSession) FileReNew(fileReNewInfo FileReNewInfo) (*types.Transaction, error) {
	return _Store.Contract.FileReNew(&_Store.TransactOpts, fileReNewInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xf54cd295.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) StoreFile(opts *bind.TransactOpts, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "StoreFile", fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xf54cd295.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xf54cd295.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreTransactor) UpdateFileInfo(opts *bind.TransactOpts, f FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileInfo", f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreTransactorSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) returns()
func (_Store *StoreTransactor) UpdateFileList(opts *bind.TransactOpts, walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileList", walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) returns()
func (_Store *StoreSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) returns()
func (_Store *StoreTransactorSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreTransactor) DeleteFilesInner(opts *bind.TransactOpts, files []FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteFilesInner", files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreTransactorSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x3d4dcb98.
//
// Solidity: function deleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteProveDetails(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteProveDetails", fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x3d4dcb98.
//
// Solidity: function deleteProveDetails(bytes fileHash) returns()
func (_Store *StoreSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x3d4dcb98.
//
// Solidity: function deleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xa1276a54.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _node, _space, _sector, _prove, fsConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xa1276a54.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove, fsConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xa1276a54.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove, fsConfig)
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
