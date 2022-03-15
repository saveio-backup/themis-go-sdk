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

// FileSystemOwnerChange is an auto generated low-level Go binding around an user-defined struct.
type FileSystemOwnerChange struct {
	FileHash []byte
	CurOwner common.Address
	NewOwner common.Address
}

// FileSystemPriChange is an auto generated low-level Go binding around an user-defined struct.
type FileSystemPriChange struct {
	FileHash  []byte
	Privilege uint64
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
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"beginHeight\",\"type\":\"uint256\"}],\"name\":\"CalcProveTimesByUploadInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structFileSystem.OwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structFileSystem.PriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromCandidateList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromPrimaryList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromUnSettledList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfoForRenew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcSingleValidFeeForFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"calcUploadFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"files\",\"type\":\"tuple[]\"}],\"name\":\"deleteFilesInner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"deleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractSpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractProve\",\"name\":\"_prove\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260048054600160a01b600160e01b031916600560a01b17908190556200003f906001600160401b03600160a01b909104166201518062000082565b600580546001600160401b03929092166001600160801b0319909216919091176a0f424000000000000000001790553480156200007b57600080fd5b50620000bc565b6001600160401b039182169116600082620000a157620000a1620000a6565b500490565b634e487b7160e01b600052601260045260246000fd5b617fa080620000cc6000396000f3fe6080604052600436106102855760003560e01c8063793cab8e11610153578063cc76e80d116100cb578063defce5d41161007f578063e4bca97311610064578063e4bca97314610739578063eb7336a014610759578063ebaf02581461078657600080fd5b8063defce5d4146106ec578063df49410a1461071957600080fd5b8063d49ce874116100b0578063d49ce8741461068c578063d70e6272146106ac578063dc1ec30b146106cc57600080fd5b8063cc76e80d1461064c578063ce9045541461066c57600080fd5b80639ffba14911610122578063b64ab1ef11610107578063b64ab1ef146105df578063c6e8392a146105ff578063ca6142cb1461062c57600080fd5b80639ffba149146105ac578063af6a8370146105cc57600080fd5b8063793cab8e1461052c5780638d41f9f81461054c57806395b0b54b1461056c5780639a2b5e631461058c57600080fd5b806334afd8c11161020157806341bc86cb116101b557806357d943991161019a57806357d94399146104d9578063655c6694146104ec57806370b1d8891461050c57600080fd5b806341bc86cb14610499578063432152ce146104b957600080fd5b80633ad525a9116101e65780633ad525a91461041c5780633d4dcb981461044b5780633f2cc9a01461046b57600080fd5b806334afd8c1146103dc57806334fddaac146103fc57600080fd5b80631aa05c5c116102585780631fe07caa1161023d5780631fe07caa1461036f578063207e33be1461038f57806328a8565c146103af57600080fd5b80631aa05c5c1461032f5780631bc558801461034f57600080fd5b80630be097021461028a5780631459457a146102ac578063170bbdf5146102cc578063178e4dc014610302575b600080fd5b34801561029657600080fd5b506102aa6102a53660046168de565b6107a6565b005b3480156102b857600080fd5b506102aa6102c7366004616b39565b610ab6565b3480156102d857600080fd5b506102ec6102e7366004616e1d565b610bb5565b6040516102f99190617be0565b60405180910390f35b34801561030e57600080fd5b5061032261031d366004616f23565b610be2565b6040516102f99190617b9c565b34801561033b57600080fd5b506102aa61034a366004616bae565b610c38565b34801561035b57600080fd5b506102aa61036a36600461694d565b610f5a565b34801561037b57600080fd5b506102aa61038a366004616a9d565b611053565b34801561039b57600080fd5b506102aa6103aa366004616c8d565b6119f0565b3480156103bb57600080fd5b506103cf6103ca3660046168de565b611a5a565b6040516102f99190617916565b3480156103e857600080fd5b506102aa6103f7366004616be2565b611b49565b34801561040857600080fd5b506102aa61041736600461694d565b611c3e565b34801561042857600080fd5b5061043c6104373660046169c8565b611c7a565b6040516102f993929190617927565b34801561045757600080fd5b506102aa610466366004616b05565b611f7e565b34801561047757600080fd5b5061048b610486366004616a42565b612106565b6040516102f9929190617965565b3480156104a557600080fd5b506103226104b4366004616eef565b6121cc565b3480156104c557600080fd5b506102aa6104d4366004616994565b6122a5565b6102aa6104e7366004616c3b565b6123d2565b3480156104f857600080fd5b506102aa61050736600461694d565b612df9565b34801561051857600080fd5b506102ec610527366004616e1d565b612ed7565b34801561053857600080fd5b506102ec610547366004616dd7565b612eef565b34801561055857600080fd5b506103cf6105673660046168de565b612f05565b34801561057857600080fd5b506102aa610587366004616b05565b61320d565b34801561059857600080fd5b506103cf6105a73660046168de565b613322565b3480156105b857600080fd5b506102aa6105c7366004616cc1565b613406565b6102aa6105da366004616bae565b613433565b3480156105eb57600080fd5b506102aa6105fa36600461694d565b613f5d565b34801561060b57600080fd5b5061061f61061a366004616994565b61403b565b6040516102f99190617954565b34801561063857600080fd5b506102ec610647366004616da5565b614a2f565b34801561065857600080fd5b50610322610667366004616e82565b614a58565b34801561067857600080fd5b506102aa610687366004616b05565b614ace565b34801561069857600080fd5b506103226106a7366004616f23565b614b81565b3480156106b857600080fd5b506102aa6106c73660046168fc565b614c6a565b3480156106d857600080fd5b506103cf6106e73660046168de565b614c93565b3480156106f857600080fd5b5061070c610707366004616b05565b614f91565b6040516102f99190617b4a565b34801561072557600080fd5b506102ec610734366004616dd7565b61544d565b34801561074557600080fd5b506102ec610754366004616f5b565b615478565b34801561076557600080fd5b50610779610774366004616d48565b6154a4565b6040516102f99190617985565b34801561079257600080fd5b506102aa6107a136600461694d565b615618565b60006107b182611a5a565b60408051600280825260608201835292935060009290916020830190803683370190505090506000816000815181106107fa57634e487b7160e01b600052603260045260246000fd5b6020026020010190600181111561082157634e487b7160e01b600052602160045260246000fd5b9081600181111561084257634e487b7160e01b600052602160045260246000fd5b8152505060018160018151811061086957634e487b7160e01b600052603260045260246000fd5b6020026020010190600181111561089057634e487b7160e01b600052602160045260246000fd5b908160018111156108b157634e487b7160e01b600052602160045260246000fd5b90525060606000806108c4858786611c7a565b9194509250905060005b8351811015610aad576001600160a01b0387166000908152600b6020908152604080832080548251818502810185019093528083529192909190849084015b828210156109b957838290600052602060002001805461092c90617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461095890617de9565b80156109a55780601f1061097a576101008083540402835291602001916109a5565b820191906000526020600020905b81548152906001019060200180831161098857829003601f168201915b50505050508152602001906001019061090d565b50505050905060005b8151811015610a6e578583815181106109eb57634e487b7160e01b600052603260045260246000fd5b602002602001015180519060200120828281518110610a1a57634e487b7160e01b600052603260045260246000fd5b6020026020010151805190602001201415610a5c57818181518110610a4f57634e487b7160e01b600052603260045260246000fd5b6020026020010160608152505b80610a6681617e3c565b9150506109c2565b506001600160a01b0388166000908152600b602090815260409091208251610a98928401906156f6565b50508080610aa590617e3c565b9150506108ce565b50505050505050565b600054610100900460ff16610ad15760005460ff1615610ad5565b303b155b610afa5760405162461bcd60e51b8152600401610af190617aea565b60405180910390fd5b600054610100900460ff16158015610b1c576000805461ffff19166101011790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff16620100006001600160a01b038981169190910291909117909155600180546001600160a01b031990811688841617909155600280548216878416179055600380548216868416179055600480549091169184169190911790558015610bad576000805461ff00191690555b505050505050565b6000610bc2858584612eef565b610bcd846001617c7d565b610bd79190617cd1565b90505b949350505050565b6040805160608101825260008082526020820181905291810182905290610c098584615478565b90506000610c2c85838860c001518960200151888b608001516106679190617d0d565b925050505b9392505050565b8060068260000151604051610c4d91906178d5565b90815260200160405180910390206000820151816000019080519060200190610c77929190615753565b506020828101516001830180546001600160a01b0319166001600160a01b0390921691909117905560408301518051610cb69260028501920190615753565b506060820151600382018054608085015160a086015160c08701516001600160401b039586166001600160801b031994851617600160401b9387168402176fffffffffffffffffffffffffffffffff16600160801b9287169290920277ffffffffffffffffffffffffffffffffffffffffffffffff1691909117600160c01b918616919091021790925560e085015160048501805467ffffffffffffffff1916918516919091179055610100850151600585015561012085015160068501805461014088015192861693169290921793169091029190911790556101608201518051610dac916007840191602090910190615753565b5061018082015160088201805467ffffffffffffffff19166001600160401b039092169190911790556101a082015160098201556101c0820151600a8201805460ff19811692151592831782556101e08501519261ffff1990911661ff001990911617610100836001811115610e3257634e487b7160e01b600052602160045260246000fd5b0217905550610200820151600a820180546001600160401b03909216620100000269ffffffffffffffff0000199092169190911790556102208201518051610e8491600b8401916020909101906157d3565b506102408201518051610ea191600c8401916020909101906157d3565b50610260820151600d8201805460ff19166001836002811115610ed457634e487b7160e01b600052602160045260246000fd5b0217905550610280820151600d820180549115156101000261ff00199092169190911790556102a0909101518051600e909201805460208301516040909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b03199093169190951617179190911691909117905550565b60005b6001600160a01b0383166000908152600a602052604090205481101561103c578180519060200120600a6000856001600160a01b03166001600160a01b031681526020019081526020016000208281548110610fc957634e487b7160e01b600052603260045260246000fd5b90600052602060002001604051610fe091906178e1565b60405180910390201415611041576001600160a01b0383166000908152600a6020526040902080548290811061102657634e487b7160e01b600052603260045260246000fd5b90600052602060002001600061103c9190615828565b505050565b8061104b81617e3c565b915050610f5d565b805161105c5750565b6000808260008151811061108057634e487b7160e01b600052603260045260246000fd5b602002602001015160200151905060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156110de57600080fd5b505afa1580156110f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111169190616d29565b905060005b84518110156111b257600085828151811061114657634e487b7160e01b600052603260045260246000fd5b60200260200101519050836001600160a01b031681602001516001600160a01b03161461119f576040517f2f7fecf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50806111aa81617e3c565b91505061111b565b5060005b84518110156119965760008582815181106111e157634e487b7160e01b600052603260045260246000fd5b6020026020010151905060006007826000015160405161120191906178d5565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561128f57600084815260209081902060408051808201909152908401546001600160a01b03811682527401000000000000000000000000000000000000000090046001600160401b03168183015282526001909201910161122f565b50505050905060005b81518110156113e35760035482516000916001600160a01b031690632ba010d7908590859081106112d957634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016112fd9190617b8e565b60006040518083038186803b15801561131557600080fd5b505afa158015611329573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113519190810190616cf5565b6003546040517fa3f029180000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063a3f029189061139d9084908890600401617b69565b600060405180830381600087803b1580156113b757600080fd5b505af11580156113cb573d6000803e3d6000fd5b505050505080806113db90617e3c565b915050611298565b506101408201516001600160401b031661140a5761140382600180611b49565b5050611984565b6004805483516040517f977fdfe20000000000000000000000000000000000000000000000000000000081526000936001600160a01b039093169263977fdfe29261145792909101617993565b60006040518083038186803b15801561146f57600080fd5b505afa158015611483573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114ab9190810190616ad1565b90506000836101400151905060008460a0015185608001516114cd9190617cd1565b905060006114db8883614a2f565b905060005b84518110156117145760015485516000916001600160a01b031690631a65374a9088908590811061152157634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b815260040161154991906178ed565b60e06040518083038186803b15801561156157600080fd5b505afa158015611575573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115999190616c6f565b905060008360018885815181106115c057634e487b7160e01b600052603260045260246000fd5b6020026020010151604001516115d69190617d28565b6115e09190617cd1565b905060006115fa8c878c6101a00151436107349190617d0d565b905060006116088284617c7d565b9050876001600160401b0316816001600160401b03161115611656576040517f870303a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80846020018181516116689190617c7d565b6001600160401b03169052506001546040517fc819d86c0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063c819d86c906116bd908790600401617b5b565b600060405180830381600087803b1580156116d757600080fd5b505af11580156116eb573d6000803e3d6000fd5b5050505080886116fb9190617d28565b975050505050808061170c90617e3c565b9150506114e0565b506000866101e00151600181111561173c57634e487b7160e01b600052602160045260246000fd5b14156117535761174c838b617c7d565b9950611971565b6001866101e00151600181111561177a57634e487b7160e01b600052602160045260246000fd5b1415611971576002546020870151604051632a7069e160e11b81526000926001600160a01b0316916354e0d3c2916117b591906004016178ed565b60a06040518083038186803b1580156117cd57600080fd5b505afa1580156117e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118059190616fa1565b90508660a00151876080015161181b9190617cd1565b6001600160401b031681600001516001600160401b0316106118ba5783816040018181516118499190617c7d565b6001600160401b031690525060a087015160808801516118699190617cd1565b8160200181815161187a9190617c7d565b6001600160401b031690525060a0870151608088015161189a9190617cd1565b815182906118a9908390617d28565b6001600160401b03169052506118ef565b60016040517fc8c84b2f000000000000000000000000000000000000000000000000000000008152600401610af19190617a9c565b60025460208801516040517fed108de90000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163ed108de99161193d9185906004016178fb565b600060405180830381600087803b15801561195757600080fd5b505af115801561196b573d6000803e3d6000fd5b50505050505b61197d86600180611b49565b5050505050505b8061198e81617e3c565b9150506111b6565b506001600160401b0383166119ab5750505050565b6040516001600160a01b038316906001600160401b03851680156108fc02916000818181858888f193505050501580156119e9573d6000803e3d6000fd5b5050505050565b60006119ff8260000151614f91565b905081602001516001600160a01b031681602001516001600160a01b031614611a3a5760405162461bcd60e51b8152600401610af190617aca565b60408201516001600160a01b03166020820152611a5681610c38565b5050565b6001600160a01b0381166000908152600b60209081526040808320805482518185028101850190935280835260609492939192909184015b82821015611b3e578382906000526020600020018054611ab190617de9565b80601f0160208091040260200160405190810160405280929190818152602001828054611add90617de9565b8015611b2a5780601f10611aff57610100808354040283529160200191611b2a565b820191906000526020600020905b815481529060010190602001808311611b0d57829003601f168201915b505050505081526020019060010190611a92565b505050509050919050565b82518215611b7157611b5a8161320d565b611b6381611f7e565b611b71846020015182613f5d565b8115611c3857611b85846020015182615618565b60005b84610220015151811015611bde57611bcc8561022001518281518110611bbe57634e487b7160e01b600052603260045260246000fd5b602002602001015183612df9565b80611bd681617e3c565b915050611b88565b5060005b846102400151518110156119e957611c268561024001518281518110611c1857634e487b7160e01b600052603260045260246000fd5b602002602001015183610f5a565b80611c3081617e3c565b915050611be2565b50505050565b6001600160a01b0382166000908152600b602090815260408220805460018101825590835291819020835161103c939190910191840190615753565b6060600080600086516001600160401b03811115611ca857634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611cdb57816020015b6060815260200190600190039081611cc65790505b5090506000808060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015611d3257600080fd5b505afa158015611d46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d6a9190616d29565b905060005b8b51811015611f1c576000611daa8d8381518110611d9d57634e487b7160e01b600052603260045260246000fd5b6020026020010151614f91565b90506000805b8c51811015611e46578c8181518110611dd957634e487b7160e01b600052603260045260246000fd5b60200260200101516001811115611e0057634e487b7160e01b600052602160045260246000fd5b836101e001516001811115611e2557634e487b7160e01b600052602160045260246000fd5b1415611e345760019150611e46565b80611e3e81617e3c565b915050611db0565b5080611e53575050611f0a565b438460c001516001600160401b0316836101000151611e729190617c65565b1115611e7f575050611f0a565b610140820151611e8f9087617c7d565b9550611e9d82600180611b49565b8d8381518110611ebd57634e487b7160e01b600052603260045260246000fd5b602002602001015188886001600160401b031681518110611eee57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508680611f0490617e57565b97505050505b80611f1481617e3c565b915050611d6f565b506001600160401b03831615611f6c576040516001600160a01b038b16906001600160401b03851680156108fc02916000818181858888f19350505050158015611f6a573d6000803e3d6000fd5b505b50929990985091965090945050505050565b600681604051611f8e91906178d5565b9081526040519081900360200190206000611fa98282615828565b6001820180546001600160a01b0319169055611fc9600283016000615828565b60006003830181905560048301805467ffffffffffffffff19169055600583018190556006830180546001600160801b031916905561200c906007840190615828565b60088201805467ffffffffffffffff19169055600060098301819055600a8301805469ffffffffffffffffffff1916905561204b90600b840190615865565b612059600c83016000615865565b50600d8101805461ffff19169055600e0180547fffffffffffffffff000000000000000000000000000000000000000000000000169055600480546040517f52e061460000000000000000000000000000000000000000000000000000000081526001600160a01b03909116916352e06146916120d891859101617993565b600060405180830381600087803b1580156120f257600080fd5b505af11580156119e9573d6000803e3d6000fd5b60606000606060005b86518110156121bc57600061213d888381518110611d9d57634e487b7160e01b600052603260045260246000fd5b90506000816101e00151600181111561216657634e487b7160e01b600052602160045260246000fd5b1461217157506121aa565b806101000151861161218357506121aa565b60006121908888846154a4565b9050806121a75783600095509550505050506121c4565b50505b806121b481617e3c565b91505061210f565b509150600190505b935093915050565b604080516060810182526000808252602080830182905292820152908201516001600160401b03166122105760405162461bcd60e51b8152600401610af190617b0a565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561226057600080fd5b505afa158015612274573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122989190616d29565b9050610c31838243614b81565b60008082516001600160401b038111156122cf57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561230857816020015b6122f5615883565b8152602001906001900390816122ed5790505b50905060005b835181101561238557600061233c858381518110611d9d57634e487b7160e01b600052603260045260246000fd5b9050806020015193508083838151811061236657634e487b7160e01b600052603260045260246000fd5b602002602001018190525050808061237d90617e3c565b91505061230e565b5061238f81611053565b7fa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec60014385856040516123c594939291906179c4565b60405180910390a1505050565b6000600682600001516040516123e891906178d5565b908152602001604051809103902060090154116124175760405162461bcd60e51b8152600401610af190617ada565b6001815160405160069161242a916178d5565b908152604051908190036020019020600a015460ff61010090910416600181111561246557634e487b7160e01b600052602160045260246000fd5b146124825760405162461bcd60e51b8152600401610af190617afa565b436006826000015160405161249791906178d5565b908152602001604051809103902060050154116124c65760405162461bcd60e51b8152600401610af190617aba565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561251657600080fd5b505afa15801561252a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061254e9190616d29565b905060006006836000015160405161256691906178d5565b9081526020016040518091039020604051806102c001604052908160008201805461259090617de9565b80601f01602080910402602001604051908101604052809291908181526020018280546125bc90617de9565b80156126095780601f106125de57610100808354040283529160200191612609565b820191906000526020600020905b8154815290600101906020018083116125ec57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161263990617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461266590617de9565b80156126b25780601f10612687576101008083540402835291602001916126b2565b820191906000526020600020905b81548152906001019060200180831161269557829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e085015291909104166101008201526007820180546101209092019161273f90617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461276b90617de9565b80156127b85780601f1061278d576101008083540402835291602001916127b8565b820191906000526020600020905b81548152906001019060200180831161279b57829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff8082161515606084015260809092019161010090910416600181111561281a57634e487b7160e01b600052602160045260246000fd5b600181111561283957634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b8301805460408051828502810185018252828152940193928301828280156128ab57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161288d575b50505050508152602001600c820180548060200260200160405190810160405280929190818152602001828054801561290d57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116128ef575b5050509183525050600d82015460209091019060ff16600281111561294257634e487b7160e01b600052602160045260246000fd5b600281111561296157634e487b7160e01b600052602160045260246000fd5b8152600d820154610100900460ff16151560208083019190915260408051606081018252600e909401546001600160401b038082168652600160401b8204811693860193909352600160801b9004909116838201529081019190915284015161012082015160a083015160808401519394506000936129fe938793909290916129ea9190617cd1565b8660c0015189604001516106679190617cd1565b9050600081602001518260400151612a169190617c7d565b9050806001600160401b0316341015612a5f5734816040517fb0b78f49000000000000000000000000000000000000000000000000000000008152600401610af1929190617bc5565b84604001518360e001818151612a759190617c7d565b6001600160401b031690525061014083018051829190612a96908390617c7d565b6001600160401b0316905250604085015160c0840151612ab69190617cd1565b6001600160401b03168361010001818151612ad19190617c65565b90525084516040518491600691612ae891906178d5565b90815260200160405180910390206000820151816000019080519060200190612b12929190615753565b506020828101516001830180546001600160a01b0319166001600160a01b0390921691909117905560408301518051612b519260028501920190615753565b506060820151600382018054608085015160a086015160c08701516001600160401b039586166001600160801b031994851617600160401b9387168402176fffffffffffffffffffffffffffffffff16600160801b9287169290920277ffffffffffffffffffffffffffffffffffffffffffffffff1691909117600160c01b918616919091021790925560e085015160048501805467ffffffffffffffff1916918516919091179055610100850151600585015561012085015160068501805461014088015192861693169290921793169091029190911790556101608201518051612c47916007840191602090910190615753565b5061018082015160088201805467ffffffffffffffff19166001600160401b039092169190911790556101a082015160098201556101c0820151600a8201805460ff19811692151592831782556101e08501519261ffff1990911661ff001990911617610100836001811115612ccd57634e487b7160e01b600052602160045260246000fd5b0217905550610200820151600a820180546001600160401b03909216620100000269ffffffffffffffff0000199092169190911790556102208201518051612d1f91600b8401916020909101906157d3565b506102408201518051612d3c91600c8401916020909101906157d3565b50610260820151600d8201805460ff19166001836002811115612d6f57634e487b7160e01b600052602160045260246000fd5b0217905550610280820151600d820180549115156101000261ff00199092169190911790556102a0909101518051600e909201805460208301516040909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909316919095161717919091169190911790555050505050565b60005b6001600160a01b03831660009081526009602052604090205481101561103c57818051906020012060096000856001600160a01b03166001600160a01b031681526020019081526020016000208281548110612e6857634e487b7160e01b600052603260045260246000fd5b90600052602060002001604051612e7f91906178e1565b60405180910390201415612ec5576001600160a01b038316600090815260096020526040902080548290811061102657634e487b7160e01b600052603260045260246000fd5b80612ecf81617e3c565b915050612dfc565b6000612ee485848461544d565b610bcd856001617c7d565b6000612efb8483614a2f565b610bda9084617cd1565b6001600160a01b0381166000908152600960209081526040808320805482518185028101850190935280835260609493849084015b82821015612fe6578382906000526020600020018054612f5990617de9565b80601f0160208091040260200160405190810160405280929190818152602001828054612f8590617de9565b8015612fd25780601f10612fa757610100808354040283529160200191612fd2565b820191906000526020600020905b815481529060010190602001808311612fb557829003601f168201915b505050505081526020019060010190612f3a565b505050509050600081516001600160401b0381111561301557634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561304857816020015b60608152602001906001900390816130335790505b5090506000805b83518110156132035760045484516000916001600160a01b03169063977fdfe29087908590811061309057634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016130b49190617993565b60006040518083038186803b1580156130cc57600080fd5b505afa1580156130e0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526131089190810190616ad1565b90506000805b825181101561317757886001600160a01b031683828151811061314157634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b031614156131655760019150613177565b8061316f81617e3c565b91505061310e565b50806131845750506131f1565b8583815181106131a457634e487b7160e01b600052603260045260246000fd5b602002602001015185856001600160401b0316815181106131d557634e487b7160e01b600052603260045260246000fd5b602002602001018190525083806131eb90617e57565b94505050505b806131fb81617e3c565b91505061304f565b5090949350505050565b60068160405161321d91906178d5565b90815260405190819003602001902060006132388282615828565b6001820180546001600160a01b0319169055613258600283016000615828565b60006003830181905560048301805467ffffffffffffffff19169055600583018190556006830180546001600160801b031916905561329b906007840190615828565b60088201805467ffffffffffffffff19169055600060098301819055600a8301805469ffffffffffffffffffff191690556132da90600b840190615865565b6132e8600c83016000615865565b50600d8101805461ffff19169055600e0180547fffffffffffffffff00000000000000000000000000000000000000000000000016905550565b6001600160a01b0381166000908152600860209081526040808320805482518185028101850190935280835260609492939192909184015b82821015611b3e57838290600052602060002001805461337990617de9565b80601f01602080910402602001604051908101604052809291908181526020018280546133a590617de9565b80156133f25780601f106133c7576101008083540402835291602001916133f2565b820191906000526020600020905b8154815290600101906020018083116133d557829003601f168201915b50505050508152602001906001019061335a565b60006134158260000151614f91565b60208301516001600160401b031660608201529050611a5681610c38565b8051604051600691613444916178d5565b9081526020016040518091039020600901546000146134755760405162461bcd60e51b8152600401610af190617b2a565b43816101000151116134995760405162461bcd60e51b8152600401610af190617aba565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156134e957600080fd5b505afa1580156134fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135219190616d29565b90506000826102600151600281111561354a57634e487b7160e01b600052602160045260246000fd5b1415613562576005546001600160401b031660c08301525b6001826102600151600281111561358957634e487b7160e01b600052602160045260246000fd5b14156135a1576005546001600160401b031660c08301525b600282610260015160028111156135c857634e487b7160e01b600052602160045260246000fd5b14156135e0576005546001600160401b031660c08301525b60016101c0830152613664604080516101e08101825260608082526000602083018190529282018390528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018190526101208201839052610140820183905261016082018190526101808201526101a08101829052906101c082015290565b61010083015160808083019190915260c0808501516001600160401b039081166040850152610120860151169083015283015160a08401516136a69190617cd1565b6001600160401b0316602082015260006136c1828443610be2565b90508060400151816020015182600001516136dc9190617c7d565b6136e69190617c7d565b6001600160401b03166101408501526136ff8243615478565b6001600160401b031660e08501526000846101e00151600181111561373457634e487b7160e01b600052602160045260246000fd5b141561396c576002546020850151604051632a7069e160e11b81526000926001600160a01b0316916354e0d3c29161376f91906004016178ed565b60a06040518083038186803b15801561378757600080fd5b505afa15801561379b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137bf9190616fa1565b90508461014001516001600160401b031681604001516001600160401b031610156138235780604001518561014001516040517f62fe63d9000000000000000000000000000000000000000000000000000000008152600401610af1929190617bee565b84608001518560a001516138379190617cd1565b6001600160401b031681602001516001600160401b031610156138925780602001518560a001516040517fd39dfd03000000000000000000000000000000000000000000000000000000008152600401610af1929190617bee565b846101000151816060015110156138e25780606001518561010001516040517f0c803412000000000000000000000000000000000000000000000000000000008152600401610af1929190617baa565b846101400151816040018181516138f99190617d28565b6001600160401b031690525060a085015160808601516139199190617cd1565b8160200181815161392a9190617d28565b6001600160401b031690525060a0850151608086015161394a9190617cd1565b81518290613959908390617c7d565b6001600160401b03169052506139a39050565b8361014001516001600160401b031634101561399a5760405162461bcd60e51b8152600401610af190617b3a565b60016101e08501525b60808301516001600160401b0316610180850152436101a0850152835160405185916006916139d291906178d5565b908152602001604051809103902060008201518160000190805190602001906139fc929190615753565b506020828101516001830180546001600160a01b0319166001600160a01b0390921691909117905560408301518051613a3b9260028501920190615753565b506060820151600382018054608085015160a086015160c08701516001600160401b039586166001600160801b031994851617600160401b9387168402176fffffffffffffffffffffffffffffffff16600160801b9287169290920277ffffffffffffffffffffffffffffffffffffffffffffffff1691909117600160c01b918616919091021790925560e085015160048501805467ffffffffffffffff1916918516919091179055610100850151600585015561012085015160068501805461014088015192861693169290921793169091029190911790556101608201518051613b31916007840191602090910190615753565b5061018082015160088201805467ffffffffffffffff19166001600160401b039092169190911790556101a082015160098201556101c0820151600a8201805460ff19811692151592831782556101e08501519261ffff1990911661ff001990911617610100836001811115613bb757634e487b7160e01b600052602160045260246000fd5b0217905550610200820151600a820180546001600160401b03909216620100000269ffffffffffffffff0000199092169190911790556102208201518051613c0991600b8401916020909101906157d3565b506102408201518051613c2691600c8401916020909101906157d3565b50610260820151600d8201805460ff19166001836002811115613c5957634e487b7160e01b600052602160045260246000fd5b0217905550610280820151600d820180549115156101000261ff00199092169190911790556102a0909101518051600e90920180546020808401516040948501516001600160401b03908116600160801b0267ffffffffffffffff60801b19928216600160401b026001600160801b031990951691909716179290921791909116939093179055858201516001600160a01b0316600090815260088352908120865181546001810183558284529284902081519294613d1f949190910192910190615753565b5060005b85610220015151811015613dbb576000600860008861022001518481518110613d5c57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0316825281810192909252604001600090812089518154600181018355828452928490208151929550613da6949301920190615753565b50508080613db390617e3c565b915050613d23565b5060005b85610240015151811015613e57576000600860008861024001518481518110613df857634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0316825281810192909252604001600090812089518154600181018355828452928490208151929550613e42949301920190615753565b50508080613e4f90617e3c565b915050613dbf565b50604080518082018252600080825260208201526101208701516001600160401b0316815260048054885193517fbb4e4e4200000000000000000000000000000000000000000000000000000000815292936001600160a01b039091169263bb4e4e4292613ec892918691016179a4565b600060405180830381600087803b158015613ee257600080fd5b505af1158015613ef6573d6000803e3d6000fd5b505050507fa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d1860004388600001518961020001518a602001518b61014001518c6102800151604051613f4d9796959493929190617a2d565b60405180910390a1505050505050565b60005b6001600160a01b0383166000908152600b602052604090205481101561103c578180519060200120600b6000856001600160a01b03166001600160a01b031681526020019081526020016000208281548110613fcc57634e487b7160e01b600052603260045260246000fd5b90600052602060002001604051613fe391906178e1565b60405180910390201415614029576001600160a01b0383166000908152600b6020526040902080548290811061102657634e487b7160e01b600052603260045260246000fd5b8061403381617e3c565b915050613f60565b6060600082511161405e5760405162461bcd60e51b8152600401610af190617aaa565b600082516001600160401b0381111561408757634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156140c057816020015b6140ad615883565b8152602001906001900390816140a55790505b50905060005b8351811015614a285760008482815181106140f157634e487b7160e01b600052603260045260246000fd5b60200260200101519050600060068260405161410d91906178d5565b9081526020016040518091039020604051806102c001604052908160008201805461413790617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461416390617de9565b80156141b05780601f10614185576101008083540402835291602001916141b0565b820191906000526020600020905b81548152906001019060200180831161419357829003601f168201915b505050918352505060018201546001600160a01b031660208201526002820180546040909201916141e090617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461420c90617de9565b80156142595780601f1061422e57610100808354040283529160200191614259565b820191906000526020600020905b81548152906001019060200180831161423c57829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e08501529190910416610100820152600782018054610120909201916142e690617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461431290617de9565b801561435f5780601f106143345761010080835404028352916020019161435f565b820191906000526020600020905b81548152906001019060200180831161434257829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff808216151560608401526080909201916101009091041660018111156143c157634e487b7160e01b600052602160045260246000fd5b60018111156143e057634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b83018054604080518285028101850182528281529401939283018282801561445257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614434575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156144b457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614496575b5050509183525050600d82015460209091019060ff1660028111156144e957634e487b7160e01b600052602160045260246000fd5b600281111561450857634e487b7160e01b600052602160045260246000fd5b8152600d820154610100900460ff16151560208083019190915260408051606081018252600e909401546001600160401b038082168652600160401b8204811693860193909352600160801b90049091168382015201528051519091506145845781604051636c5249c160e01b8152600401610af19190617993565b60068260405161459491906178d5565b9081526020016040518091039020604051806102c00160405290816000820180546145be90617de9565b80601f01602080910402602001604051908101604052809291908181526020018280546145ea90617de9565b80156146375780601f1061460c57610100808354040283529160200191614637565b820191906000526020600020905b81548152906001019060200180831161461a57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161466790617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461469390617de9565b80156146e05780601f106146b5576101008083540402835291602001916146e0565b820191906000526020600020905b8154815290600101906020018083116146c357829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e085015291909104166101008201526007820180546101209092019161476d90617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461479990617de9565b80156147e65780601f106147bb576101008083540402835291602001916147e6565b820191906000526020600020905b8154815290600101906020018083116147c957829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff8082161515606084015260809092019161010090910416600181111561484857634e487b7160e01b600052602160045260246000fd5b600181111561486757634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b8301805460408051828502810185018252828152940193928301828280156148d957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116148bb575b50505050508152602001600c820180548060200260200160405190810160405280929190818152602001828054801561493b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161491d575b5050509183525050600d82015460209091019060ff16600281111561497057634e487b7160e01b600052602160045260246000fd5b600281111561498f57634e487b7160e01b600052602160045260246000fd5b8152600d820154610100900460ff16151560208083019190915260408051606081018252600e909401546001600160401b038082168652600160401b8204811693860193909352600160801b90049091168382015201528451859085908110614a0857634e487b7160e01b600052603260045260246000fd5b602002602001018190525050508080614a2090617e3c565b9150506140c6565b5092915050565b6000620fa000828460600151614a459190617cd1565b614a4f9190617cb0565b90505b92915050565b6040805160608082018352600080835260208084018290528385018290528451928301855281835282018190529281018390529091614a9988888888610bb5565b90506000614aa989888888612ed7565b6001600160401b03928316604085015290911660208301525090505b95945050505050565b6000614ad982614f91565b60408051600180825281830190925291925060009190816020015b614afc615883565b815260200190600190039081614af45790505090508181600081518110614b3357634e487b7160e01b600052603260045260246000fd5b6020026020010181905250614b4781611053565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f456001438585602001516040516123c59493929190617a00565b604080516060810182526000808252602082018190529181019190915261018084015151600090629896809015614bc457614bbd816004617cd1565b9150614bd2565b614bcf816003617cd1565b91505b604080516060810182526000602082018190529181018290526001600160401b03841681526101c08801519091906001811115614c1f57634e487b7160e01b600052602160045260246000fd5b1415614c2f579250610c31915050565b6000614c3c888888610be2565b6040818101516001600160401b03908116918501919091526020918201511690830152509695505050505050565b6001600160a01b0382166000908152600860209081526040909120825161103c928401906156f6565b6001600160a01b0381166000908152600a60209081526040808320805482518185028101850190935280835260609493849084015b82821015614d74578382906000526020600020018054614ce790617de9565b80601f0160208091040260200160405190810160405280929190818152602001828054614d1390617de9565b8015614d605780601f10614d3557610100808354040283529160200191614d60565b820191906000526020600020905b815481529060010190602001808311614d4357829003601f168201915b505050505081526020019060010190614cc8565b505050509050600081516001600160401b03811115614da357634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015614dd657816020015b6060815260200190600190039081614dc15790505b5090506000805b83518110156132035760045484516000916001600160a01b03169063977fdfe290879085908110614e1e57634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b8152600401614e429190617993565b60006040518083038186803b158015614e5a57600080fd5b505afa158015614e6e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614e969190810190616ad1565b90506000805b8251811015614f0557886001600160a01b0316838281518110614ecf57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b03161415614ef35760019150614f05565b80614efd81617e3c565b915050614e9c565b5080614f12575050614f7f565b858381518110614f3257634e487b7160e01b600052603260045260246000fd5b602002602001015185856001600160401b031681518110614f6357634e487b7160e01b600052603260045260246000fd5b60200260200101819052508380614f7990617e57565b94505050505b80614f8981617e3c565b915050614ddd565b614f99615883565b816000815111614fbb5760405162461bcd60e51b8152600401610af190617b1a565b6000600684604051614fcd91906178d5565b9081526020016040518091039020604051806102c0016040529081600082018054614ff790617de9565b80601f016020809104026020016040519081016040528092919081815260200182805461502390617de9565b80156150705780601f1061504557610100808354040283529160200191615070565b820191906000526020600020905b81548152906001019060200180831161505357829003601f168201915b505050918352505060018201546001600160a01b031660208201526002820180546040909201916150a090617de9565b80601f01602080910402602001604051908101604052809291908181526020018280546150cc90617de9565b80156151195780601f106150ee57610100808354040283529160200191615119565b820191906000526020600020905b8154815290600101906020018083116150fc57829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e08501529190910416610100820152600782018054610120909201916151a690617de9565b80601f01602080910402602001604051908101604052809291908181526020018280546151d290617de9565b801561521f5780601f106151f45761010080835404028352916020019161521f565b820191906000526020600020905b81548152906001019060200180831161520257829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff8082161515606084015260809092019161010090910416600181111561528157634e487b7160e01b600052602160045260246000fd5b60018111156152a057634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b83018054604080518285028101850182528281529401939283018282801561531257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116152f4575b50505050508152602001600c820180548060200260200160405190810160405280929190818152602001828054801561537457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311615356575b5050509183525050600d82015460209091019060ff1660028111156153a957634e487b7160e01b600052602160045260246000fd5b60028111156153c857634e487b7160e01b600052602160045260246000fd5b8152600d820154610100900460ff16151560208083019190915260408051606081018252600e909401546001600160401b038082168652600160401b8204811693860193909352600160801b90049091168382015201528051519091506154445783604051636c5249c160e01b8152600401610af19190617993565b91505b50919050565b6000620fa000828486602001516154649190617cd1565b61546e9190617cd1565b610bda9190617cb0565b6000826040015182846080015161548f9190617d0d565b6154999190617cb0565b614a4f906001617c7d565b6101008101829052600061552a604080516101e08101825260608082526000602083018190529282018390528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018190526101208201839052610140820183905261016082018190526101808201526101a08101829052906101c082015290565b61010083015160808083019190915260c0808501516001600160401b039081166040850152610120860151169083015283015160a084015161556c9190617cd1565b6001600160401b031660208201526101a0830151600061558d838884610be2565b905060008160400151826020015183600001516155aa9190617c7d565b6155b49190617c7d565b90508561014001516001600160401b0316816001600160401b0316116155e1576000945050505050610c31565b6001600160401b0381166101408701526155fb8484615478565b6001600160401b031660e087015250600193505050509392505050565b60005b6001600160a01b03831660009081526008602052604090205481101561103c57818051906020012060086000856001600160a01b03166001600160a01b03168152602001908152602001600020828154811061568757634e487b7160e01b600052603260045260246000fd5b9060005260206000200160405161569e91906178e1565b604051809103902014156156e4576001600160a01b038316600090815260086020526040902080548290811061102657634e487b7160e01b600052603260045260246000fd5b806156ee81617e3c565b91505061561b565b828054828255906000526020600020908101928215615743579160200282015b828111156157435782518051615733918491602090910190615753565b5091602001919060010190615716565b5061574f92915061594b565b5090565b82805461575f90617de9565b90600052602060002090601f01602090048101928261578157600085556157c7565b82601f1061579a57805160ff19168380011785556157c7565b828001600101855582156157c7579182015b828111156157c75782518255916020019190600101906157ac565b5061574f929150615968565b8280548282559060005260206000209081019282156157c7579160200282015b828111156157c757825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906157f3565b50805461583490617de9565b6000825580601f10615844575050565b601f0160209004906000526020600020908101906158629190615968565b50565b50805460008255906000526020600020908101906158629190615968565b604080516102c08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018390526101208201839052610140820183905261016082015261018081018290526101a081018290526101c08101829052906101e0820190815260006020820181905260606040830181905280830152608090910190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b8082111561574f57600061595f8282615828565b5060010161594b565b5b8082111561574f5760008155600101615969565b600061599061598b84617c18565b617bfc565b905080838252602082019050828560208602820111156159af57600080fd5b60005b858110156159db57816159c58882615cae565b84525060209283019291909101906001016159b2565b5050509392505050565b60006159f361598b84617c18565b90508083825260208201905082856020860282011115615a1257600080fd5b60005b858110156159db5781356001600160401b03811115615a3357600080fd5b808601615a408982615dc1565b855250506020928301929190910190600101615a15565b6000615a6561598b84617c18565b90508083825260208201905082856020860282011115615a8457600080fd5b60005b858110156159db5781516001600160401b03811115615aa557600080fd5b808601615ab28982615de2565b855250506020928301929190910190600101615a87565b6000615ad761598b84617c18565b90508083825260208201905082856020860282011115615af657600080fd5b60005b858110156159db5781615b0c8882615e24565b8452506020928301929190910190600101615af9565b6000615b3061598b84617c18565b90508083825260208201905082856020860282011115615b4f57600080fd5b60005b858110156159db5781356001600160401b03811115615b7057600080fd5b808601615b7d8982615e2f565b855250506020928301929190910190600101615b52565b6000615ba261598b84617c18565b8381529050602081018260a085028101861015615bbe57600080fd5b60005b858110156159db5781615bd488826162af565b84525060209092019160a09190910190600101615bc1565b6000615bfa61598b84617c18565b83815290506020810182606085028101861015615c1657600080fd5b60005b858110156159db5781615c2c8882616888565b84525060209092019160609190910190600101615c19565b6000615c5261598b84617c3b565b905082815260208101848484011115615c6a57600080fd5b615c75848285617db1565b509392505050565b6000615c8b61598b84617c3b565b905082815260208101848484011115615ca357600080fd5b615c75848285617dbd565b8035614a5281617f16565b8051614a5281617f16565b600082601f830112615cd557600080fd5b8135610bda84826020860161597d565b600082601f830112615cf657600080fd5b8135610bda8482602086016159e5565b600082601f830112615d1757600080fd5b8151610bda848260208601615a57565b600082601f830112615d3857600080fd5b8135610bda848260208601615ac9565b600082601f830112615d5957600080fd5b8135610bda848260208601615b22565b600082601f830112615d7a57600080fd5b8151610bda848260208601615b94565b600082601f830112615d9b57600080fd5b8135610bda848260208601615bec565b8035614a5281617f2a565b8051614a5281617f2a565b600082601f830112615dd257600080fd5b8135610bda848260208601615c44565b600082601f830112615df357600080fd5b8151610bda848260208601615c7d565b8035614a5281617f32565b8035614a5281617f3b565b8051614a5281617f3b565b8035614a5281617f48565b60006103008284031215615e4257600080fd5b615e4d6102c0617bfc565b905081356001600160401b03811115615e6557600080fd5b615e7184828501615dc1565b8252506020615e8284848301615cae565b60208301525060408201356001600160401b03811115615ea157600080fd5b615ead84828501615dc1565b6040830152506060615ec1848285016168c8565b6060830152506080615ed5848285016168c8565b60808301525060a0615ee9848285016168c8565b60a08301525060c0615efd848285016168c8565b60c08301525060e0615f11848285016168c8565b60e083015250610100615f26848285016168b2565b61010083015250610120615f3c848285016168c8565b61012083015250610140615f52848285016168c8565b610140830152506101608201356001600160401b03811115615f7357600080fd5b615f7f84828501615dc1565b61016083015250610180615f95848285016168c8565b610180830152506101a0615fab848285016168b2565b6101a0830152506101c0615fc184828501615dab565b6101c0830152506101e0615fd784828501615e24565b6101e083015250610200615fed848285016168c8565b610200830152506102208201356001600160401b0381111561600e57600080fd5b61601a84828501615cc4565b610220830152506102408201356001600160401b0381111561603b57600080fd5b61604784828501615cc4565b6102408301525061026061605d84828501615e0e565b6102608301525061028061607384828501615dab565b610280830152506102a061608984828501616217565b6102a08301525092915050565b6000606082840312156160a857600080fd5b6160b26060617bfc565b905081356001600160401b038111156160ca57600080fd5b6160d684828501615dc1565b82525060206160e784848301615cae565b60208301525060406160fb848285016168c8565b60408301525092915050565b600060e0828403121561611957600080fd5b61612360e0617bfc565b9050600061613184846168d3565b8252506020616142848483016168d3565b6020830152506040616156848285016168d3565b604083015250606061616a848285016168d3565b606083015250608061617e848285016168d3565b60808301525060a061619284828501615cb9565b60a08301525060c06161a684828501615cb9565b60c08301525092915050565b6000606082840312156161c457600080fd5b6161ce6060617bfc565b905081356001600160401b038111156161e657600080fd5b6161f284828501615dc1565b825250602061620384848301615cae565b60208301525060406160fb84828501615cae565b60006060828403121561622957600080fd5b6162336060617bfc565b9050600061624184846168c8565b82525060206160e7848483016168c8565b60006040828403121561626457600080fd5b61626e6040617bfc565b905081356001600160401b0381111561628657600080fd5b61629284828501615dc1565b82525060206162a3848483016168c8565b60208301525092915050565b600060a082840312156162c157600080fd5b6162cb60a0617bfc565b905060006162d98484615cb9565b82525060206162ea84848301615cb9565b60208301525060406162fe848285016168d3565b6040830152506060616312848285016168bd565b606083015250608061632684828501615db6565b60808301525092915050565b6000610180828403121561634557600080fd5b616350610180617bfc565b9050600061635e8484615cb9565b825250602061636f848483016168d3565b6020830152506040616383848285016168d3565b6040830152506060616397848285016168d3565b60608301525060806163ab84828501615e19565b60808301525060a06163bf848285016168bd565b60a08301525060c06163d3848285016168bd565b60c08301525060e06163e7848285016168d3565b60e0830152506101006163fc848285016168d3565b61010083015250610120616412848285016168d3565b6101208301525061014061642884828501615db6565b610140830152506101608201516001600160401b0381111561644957600080fd5b61645584828501615d06565b6101608301525092915050565b6000610160828403121561647557600080fd5b616480610160617bfc565b9050600061648e84846168c8565b825250602061649f848483016168c8565b60208301525060406164b3848285016168c8565b60408301525060606164c7848285016168c8565b60608301525060806164db848285016168c8565b60808301525060a06164ef848285016168c8565b60a08301525060c0616503848285016168c8565b60c08301525060e0616517848285016168c8565b60e08301525061010061652c848285016168c8565b61010083015250610120616542848285016168c8565b61012083015250610140616558848285016168c8565b6101408301525092915050565b6000610160828403121561657857600080fd5b616583610160617bfc565b9050600061659184846168d3565b82525060206165a2848483016168d3565b60208301525060406165b6848285016168d3565b60408301525060606165ca848285016168d3565b60608301525060806165de848285016168d3565b60808301525060a06165f2848285016168d3565b60a08301525060c0616606848285016168d3565b60c08301525060e061661a848285016168d3565b60e08301525061010061662f848285016168d3565b61010083015250610120616645848285016168d3565b61012083015250610140616558848285016168d3565b60006101e0828403121561666e57600080fd5b6166796101e0617bfc565b905081356001600160401b0381111561669157600080fd5b61669d84828501615dc1565b82525060206166ae848483016168c8565b60208301525060406166c2848285016168c8565b60408301525060606166d6848285016168c8565b60608301525060806166ea848285016168b2565b60808301525060a06166fe848285016168c8565b60a08301525060c0616712848285016168c8565b60c08301525060e061672684828501615dab565b60e0830152506101008201356001600160401b0381111561674657600080fd5b61675284828501615dc1565b6101008301525061012061676884828501615dab565b6101208301525061014061677e84828501615dab565b610140830152506101608201356001600160401b0381111561679f57600080fd5b6167ab84828501615dc1565b610160830152506101808201356001600160401b038111156167cc57600080fd5b6167d884828501615d8a565b610180830152506101a06167ee84828501615dab565b6101a0830152506101c061680484828501615e24565b6101c08301525092915050565b600060a0828403121561682357600080fd5b61682d60a0617bfc565b9050600061683b84846168d3565b825250602061684c848483016168d3565b6020830152506040616860848285016168d3565b6040830152506060616874848285016168bd565b6060830152506080616326848285016168bd565b60006060828403121561689a57600080fd5b6168a46060617bfc565b905060006162418484615cae565b8035614a5281617f55565b8051614a5281617f55565b8035614a5281617f5b565b8051614a5281617f5b565b6000602082840312156168f057600080fd5b6000610bda8484615cae565b6000806040838503121561690f57600080fd5b600061691b8585615cae565b92505060208301356001600160401b0381111561693757600080fd5b61694385828601615ce5565b9150509250929050565b6000806040838503121561696057600080fd5b600061696c8585615cae565b92505060208301356001600160401b0381111561698857600080fd5b61694385828601615dc1565b6000602082840312156169a657600080fd5b81356001600160401b038111156169bc57600080fd5b610bda84828501615ce5565b6000806000606084860312156169dd57600080fd5b83356001600160401b038111156169f357600080fd5b6169ff86828701615ce5565b9350506020616a1086828701615cae565b92505060408401356001600160401b03811115616a2c57600080fd5b616a3886828701615d27565b9150509250925092565b60008060006101a08486031215616a5857600080fd5b83356001600160401b03811115616a6e57600080fd5b616a7a86828701615ce5565b9350506020616a8b86828701616462565b925050610180616a38868287016168b2565b600060208284031215616aaf57600080fd5b81356001600160401b03811115616ac557600080fd5b610bda84828501615d48565b600060208284031215616ae357600080fd5b81516001600160401b03811115616af957600080fd5b610bda84828501615d69565b600060208284031215616b1757600080fd5b81356001600160401b03811115616b2d57600080fd5b610bda84828501615dc1565b600080600080600060a08688031215616b5157600080fd5b6000616b5d8888615e03565b9550506020616b6e88828901615e03565b9450506040616b7f88828901615e03565b9350506060616b9088828901615e03565b9250506080616ba188828901615e03565b9150509295509295909350565b600060208284031215616bc057600080fd5b81356001600160401b03811115616bd657600080fd5b610bda84828501615e2f565b600080600060608486031215616bf757600080fd5b83356001600160401b03811115616c0d57600080fd5b616c1986828701615e2f565b9350506020616c2a86828701615dab565b9250506040616a3886828701615dab565b600060208284031215616c4d57600080fd5b81356001600160401b03811115616c6357600080fd5b610bda84828501616096565b600060e08284031215616c8157600080fd5b6000610bda8484616107565b600060208284031215616c9f57600080fd5b81356001600160401b03811115616cb557600080fd5b610bda848285016161b2565b600060208284031215616cd357600080fd5b81356001600160401b03811115616ce957600080fd5b610bda84828501616252565b600060208284031215616d0757600080fd5b81516001600160401b03811115616d1d57600080fd5b610bda84828501616332565b60006101608284031215616d3c57600080fd5b6000610bda8484616565565b60008060006101a08486031215616d5e57600080fd5b6000616d6a8686616462565b935050610160616d7c868287016168b2565b9250506101808401356001600160401b03811115616d9957600080fd5b616a3886828701615e2f565b6000806101808385031215616db957600080fd5b6000616dc58585616462565b925050610160616943858286016168c8565b60008060006101a08486031215616ded57600080fd5b6000616df98686616462565b935050610160616e0b868287016168c8565b925050610180616a38868287016168c8565b6000806000806101c08587031215616e3457600080fd5b6000616e408787616462565b945050610160616e52878288016168c8565b935050610180616e64878288016168c8565b9250506101a0616e76878288016168c8565b91505092959194509250565b60008060008060006101e08688031215616e9b57600080fd5b6000616ea78888616462565b955050610160616eb9888289016168c8565b945050610180616ecb888289016168c8565b9350506101a0616edd888289016168c8565b9250506101c0616ba1888289016168c8565b600060208284031215616f0157600080fd5b81356001600160401b03811115616f1757600080fd5b610bda8482850161665b565b60008060006101a08486031215616f3957600080fd5b83356001600160401b03811115616f4f57600080fd5b616a7a8682870161665b565b60008060408385031215616f6e57600080fd5b82356001600160401b03811115616f8457600080fd5b616f908582860161665b565b9250506020616943858286016168b2565b600060a08284031215616fb357600080fd5b6000610bda8484616811565b6000616fcb8383616feb565b505060200190565b6000614a4f8383617163565b6000614a4f83836174a3565b616ff481617d45565b82525050565b6000617004825190565b80845260209384019383018060005b838110156170385781516170278882616fbf565b975060208301925050600101617013565b509495945050505050565b600061704d825190565b808452602084019350836020820285016170678560200190565b8060005b8581101561709c57848403895281516170848582616fd3565b94506020830160209a909a019992505060010161706b565b5091979650505050505050565b60006170b3825190565b808452602084019350836020820285016170cd8560200190565b8060005b8581101561709c57848403895281516170ea8582616fd3565b94506020830160209a909a01999250506001016170d1565b600061710c825190565b808452602084019350836020820285016171268560200190565b8060005b8581101561709c57848403895281516171438582616fdf565b94506020830160209a909a019992505060010161712a565b801515616ff4565b600061716d825190565b808452602084019350617184818560208601617dbd565b601f01601f19169290920192915050565b600061719f825190565b6171ad818560208601617dbd565b9290920192915050565b600081546171c481617de9565b6001821680156171db57600181146171ec5761721c565b60ff1983168652818601935061721c565b60008581526020902060005b83811015617214578154888201526001909101906020016171f8565b838801955050505b50505092915050565b616ff481617d7f565b616ff481617d8a565b616ff481617d95565b616ff481617da0565b601181526000602082017f66696c654c69737420697320656d707479000000000000000000000000000000815291505b5060200190565b600c81526000602082017f66696c652065787069726564000000000000000000000000000000000000000081529150617279565b602b81526000602082017f43757272656e74206f776e6572206d75737420626520746865206f776e65722081527f6f66207468652066696c65000000000000000000000000000000000000000000602082015291505b5060400190565b600e81526000602082017f66696c65206e6f7420657869737400000000000000000000000000000000000081529150617279565b602e81526000602082017f496e697469616c697a61626c653a20636f6e747261637420697320616c72656181527f647920696e697469616c697a65640000000000000000000000000000000000006020820152915061730a565b600f81526000602082017f66696c652074797065206572726f72000000000000000000000000000000000081529150617279565b601f81526000602082017f66696c6553697a65206d7573742062652067726561746572207468616e20300081529150617279565b601a81526000602082017f66696c6548617368206d757374206265206e6f7420656d70747900000000000081529150617279565b601281526000602082017f66696c6520616c7265616479206578697374000000000000000000000000000081529150617279565b601481526000602082017f696e73756666696369656e74206465706f73697400000000000000000000000081529150617279565b8051610300808452600091908401906174bc8282617163565b91505060208301516174d16020860182616feb565b50604083015184820360408601526174e98282617163565b91505060608301516174fe60608601826178c6565b50608083015161751160808601826178c6565b5060a083015161752460a08601826178c6565b5060c083015161753760c08601826178c6565b5060e083015161754a60e08601826178c6565b5061010083015161755f6101008601826178c0565b506101208301516175746101208601826178c6565b506101408301516175896101408601826178c6565b506101608301518482036101608601526175a38282617163565b9150506101808301516175ba6101808601826178c6565b506101a08301516175cf6101a08601826178c0565b506101c08301516175e46101c086018261715b565b506101e08301516175f96101e0860182617237565b5061020083015161760e6102008601826178c6565b506102208301518482036102208601526176288282616ffa565b9150506102408301518482036102408601526176448282616ffa565b91505061026083015161765b61026086018261722e565b5061028083015161767061028086018261715b565b506102a0830151615c756102a0860182617708565b805160e083019061769684826178c6565b5060208201516176a960208501826178c6565b5060408201516176bc60408501826178c6565b5060608201516176cf60608501826178c6565b5060808201516176e260808501826178c6565b5060a08201516176f560a0850182616feb565b5060c0820151611c3860c0850182616feb565b8051606083019061771984826178c6565b50602082015161772c60208501826178c6565b506040820151611c3860408501826178c6565b8051604083019061775084826178c6565b506020820151611c3860208501826178c6565b80516000906101808401906177788582616feb565b50602083015161778b60208601826178c6565b50604083015161779e60408601826178c6565b5060608301516177b160608601826178c6565b5060808301516177c4608086018261722e565b5060a08301516177d760a08601826178c0565b5060c08301516177ea60c08601826178c0565b5060e08301516177fd60e08601826178c6565b506101008301516178126101008601826178c6565b506101208301516178276101208601826178c6565b5061014083015161783c61014086018261715b565b50610160830151848203610160860152614ac58282617043565b805160408301906177508482616feb565b805160a083019061787884826178c6565b50602082015161788b60208501826178c6565b50604082015161789e60408501826178c6565b5060608201516178b160608501826178c0565b506080820151611c3860808501825b80616ff4565b6001600160401b038116616ff4565b6000610c318284617195565b6000610c3182846171b7565b60208101614a528284616feb565b60c081016179098285616feb565b610c316020830184617867565b60208082528101614a4f81846170a9565b6060808252810161793881866170a9565b905061794760208301856178c6565b610bda604083018461715b565b60208082528101614a4f8184617102565b604080825281016179768185617102565b9050610c31602083018461715b565b60208101614a52828461715b565b60208082528101614a4f8184617163565b606080825281016179b58185617163565b9050610c31602083018461773f565b608081016179d28287617225565b6179df60208301866178c0565b81810360408301526179f181856170a9565b9050614ac56060830184616feb565b60808101617a0e8287617225565b617a1b60208301866178c0565b81810360408301526179f18185617163565b60e08101617a3b828a617225565b617a4860208301896178c0565b8181036040830152617a5a8188617163565b9050617a6960608301876178c6565b617a766080830186616feb565b617a8360a08301856178c6565b617a9060c083018461715b565b98975050505050505050565b60208101614a528284617240565b60208082528101614a5281617249565b60208082528101614a5281617280565b60208082528101614a52816172b4565b60208082528101614a5281617311565b60208082528101614a5281617345565b60208082528101614a528161739f565b60208082528101614a52816173d3565b60208082528101614a5281617407565b60208082528101614a528161743b565b60208082528101614a528161746f565b60208082528101614a4f81846174a3565b60e08101614a528284617685565b60408082528101617b7a8185617763565b90508181036020830152610bda81846174a3565b60408101614a528284617856565b60608101614a528284617708565b60408101617bb882856178c0565b610c3160208301846178c0565b60408101617bd382856178c0565b610c316020830184617240565b60208101614a5282846178c6565b60408101617bd38285617240565b6000617c0760405190565b9050617c138282617e10565b919050565b60006001600160401b03821115617c3157617c31617ed0565b5060209081020190565b60006001600160401b03821115617c5457617c54617ed0565b601f19601f83011660200192915050565b60008219821115617c7857617c78617e78565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115617c7857617c78617e78565b6001600160401b039182169116600082617ccc57617ccc617e8e565b500490565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615617d0857617d08617e78565b500290565b6000825b925082821015617d2357617d23617e78565b500390565b60006001600160401b03821691506001600160401b038316617d11565b60006001600160a01b038216614a52565b6000614a5282617d45565b80617c1381617ee6565b80617c1381617ef6565b80617c1381617f06565b6000614a5282617d61565b6000614a5282617d6b565b6000614a5282617d75565b60006001600160401b038216614a52565b82818337506000910152565b60005b83811015617dd8578181015183820152602001617dc0565b83811115611c385750506000910152565b600281046001821680617dfd57607f821691505b6020821081141561544757615447617eba565b601f19601f83011681018181106001600160401b0382111715617e3557617e35617ed0565b6040525050565b6000600019821415617e5057617e50617e78565b5060010190565b60006001600160401b03821691506001600160401b03821415617e5057617e505b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600a811061586257615862617ea4565b6003811061586257615862617ea4565b6002811061586257615862617ea4565b617f1f81617d45565b811461586257600080fd5b801515617f1f565b617f1f81617d56565b6003811061586257600080fd5b6002811061586257600080fd5b80617f1f565b6001600160401b038116617f1f56fea2646970667358221220a450342e4401c31025962e14f358794455d1ffa9868133824ecdcb28f89215e264736f6c63430008040033",
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
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)))
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
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreCallerSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[])
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
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreSession) GetFileInfos(_fileList [][]byte) ([]FileInfo, error) {
	return _Store.Contract.GetFileInfos(&_Store.CallOpts, _fileList)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[])
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

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0xeb7336a0.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCaller) UpdateFileInfoForRenew(opts *bind.CallOpts, setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "UpdateFileInfoForRenew", setting, newExpireHeight, fileInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0xeb7336a0.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0xeb7336a0.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCallerSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[], bool)
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
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreCallerSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
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
func (_Store *StoreTransactor) ChangeFileOwner(opts *bind.TransactOpts, ownerChange FileSystemOwnerChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFileOwner", ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreSession) ChangeFileOwner(ownerChange FileSystemOwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactorSession) ChangeFileOwner(ownerChange FileSystemOwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFilePivilege is a paid mutator transaction binding the contract method 0x9ffba149.
//
// Solidity: function ChangeFilePivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactor) ChangeFilePivilege(opts *bind.TransactOpts, priChange FileSystemPriChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFilePivilege", priChange)
}

// ChangeFilePivilege is a paid mutator transaction binding the contract method 0x9ffba149.
//
// Solidity: function ChangeFilePivilege((bytes,uint64) priChange) returns()
func (_Store *StoreSession) ChangeFilePivilege(priChange FileSystemPriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePivilege(&_Store.TransactOpts, priChange)
}

// ChangeFilePivilege is a paid mutator transaction binding the contract method 0x9ffba149.
//
// Solidity: function ChangeFilePivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactorSession) ChangeFilePivilege(priChange FileSystemPriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePivilege(&_Store.TransactOpts, priChange)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x34afd8c1.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreTransactor) CleanupForDeleteFile(opts *bind.TransactOpts, fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CleanupForDeleteFile", fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x34afd8c1.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x34afd8c1.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
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

// StoreFile is a paid mutator transaction binding the contract method 0xaf6a8370.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) StoreFile(opts *bind.TransactOpts, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "StoreFile", fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xaf6a8370.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xaf6a8370.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x1aa05c5c.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreTransactor) UpdateFileInfo(opts *bind.TransactOpts, f FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileInfo", f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x1aa05c5c.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x1aa05c5c.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) f) returns()
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

// DeleteFilesInner is a paid mutator transaction binding the contract method 0x1fe07caa.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreTransactor) DeleteFilesInner(opts *bind.TransactOpts, files []FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteFilesInner", files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0x1fe07caa.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0x1fe07caa.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[] files) returns()
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

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _node, _space, _sector, _prove)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove)
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
