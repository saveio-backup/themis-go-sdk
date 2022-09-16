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

// NameInfo is an auto generated low-level Go binding around an user-defined struct.
type NameInfo struct {
	Type        uint64
	Header      []byte
	URL         []byte
	Name        []byte
	NameOwner   common.Address
	Desc        []byte
	BlockHeight *big.Int
	TTL         uint64
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

// SectorRef is an auto generated low-level Go binding around an user-defined struct.
type SectorRef struct {
	NodeAddr common.Address
	SectorId uint64
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

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
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
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"ref\",\"type\":\"tuple\"}],\"name\":\"AddFileSectorRef\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structOwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structPriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"calcUploadFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"files\",\"type\":\"tuple[]\"}],\"name\":\"deleteFilesInner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractISpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractIProve\",\"name\":\"_prove\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"DEFAULT_BLOCK_INTERVAL\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DEFAULT_PROVE_PERIOD\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"IN_SECTOR_SIZE\",\"type\":\"uint64\"}],\"internalType\":\"structFSConfig\",\"name\":\"fsConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractFileExtra\",\"name\":\"_fileExtra\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615f6f80620000216000396000f3fe6080604052600436106101b75760003560e01c806395b0b54b116100ec578063cc76e80d1161008a578063d70e627211610064578063d70e6272146104ca578063dc1ec30b146104dd578063defce5d4146104fd578063f54cd2951461052a57600080fd5b8063cc76e80d1461046a578063ce9045541461048a578063d49ce874146104aa57600080fd5b8063b6af10fb116100c6578063b6af10fb146103d0578063bc1c783e146103f0578063c43007e514610410578063c6e8392a1461043d57600080fd5b806395b0b54b1461037d5780639a2b5e631461039d5780639cd103e9146103bd57600080fd5b806341bc86cb1161015957806357d943991161013357806357d94399146103245780635a0e748214610337578063681850d71461034a5780638d41f9f81461035d57600080fd5b806341bc86cb146102c4578063432152ce146102e457806352e061461461030457600080fd5b806328a8565c1161019557806328a8565c1461023457806334fddaac146102615780633ad525a9146102745780633f2cc9a0146102a357600080fd5b80630be09702146101bc578063178e4dc0146101de578063207e33be14610214575b600080fd5b3480156101c857600080fd5b506101dc6101d7366004614556565b61053d565b005b3480156101ea57600080fd5b506101fe6101f9366004614c27565b6107c5565b60405161020b9190615b8f565b60405180910390f35b34801561022057600080fd5b506101dc61022f366004614aa0565b610885565b34801561024057600080fd5b5061025461024f366004614556565b610903565b60405161020b91906157c1565b6101dc61026f3660046145c5565b6109a7565b34801561028057600080fd5b5061029461028f366004614674565b610a28565b60405161020b93929190615801565b6102b66102b13660046146ee565b610d2c565b60405161020b92919061583f565b3480156102d057600080fd5b506101fe6102df366004614bf3565b610de1565b3480156102f057600080fd5b506101dc6102ff36600461460c565b610ec3565b34801561031057600080fd5b506101dc61031f36600461485f565b61103a565b6101dc610332366004614a38565b6110e4565b6101dc610345366004614893565b6111e7565b6101dc610358366004614977565b611232565b34801561036957600080fd5b50610254610378366004614556565b61127b565b34801561038957600080fd5b506101dc61039836600461485f565b61132a565b3480156103a957600080fd5b506102546103b8366004614556565b61135a565b6101dc6103cb3660046149df565b6113a4565b3480156103dc57600080fd5b506101dc6103eb366004614ad4565b6114b1565b3480156103fc57600080fd5b506101dc61040b3660046148d9565b6114fa565b34801561041c57600080fd5b5061043061042b366004614749565b6116df565b60405161020b919061585f565b34801561044957600080fd5b5061045d61045836600461460c565b6121bd565b60405161020b919061582e565b34801561047657600080fd5b506101fe610485366004614b5b565b61225b565b34801561049657600080fd5b506101dc6104a536600461485f565b612321565b3480156104b657600080fd5b506101fe6104c5366004614c27565b61240f565b6101dc6104d8366004614574565b612475565b3480156104e957600080fd5b506102546104f8366004614556565b6124c0565b34801561050957600080fd5b5061051d61051836600461485f565b61250f565b60405161020b9190615a5d565b6101dc610538366004614977565b6125da565b600061054882610903565b604080516002808252606082018352929350600092909160208301908036833701905050905060008160008151811061059157634e487b7160e01b600052603260045260246000fd5b602002602001019060018111156105b857634e487b7160e01b600052602160045260246000fd5b908160018111156105d957634e487b7160e01b600052602160045260246000fd5b8152505060018160018151811061060057634e487b7160e01b600052603260045260246000fd5b6020026020010190600181111561062757634e487b7160e01b600052602160045260246000fd5b9081600181111561064857634e487b7160e01b600052602160045260246000fd5b905250606060008061065b858786610a28565b9194509250905060005b83518110156107bc57600061067988610903565b905060005b815181101561072a578583815181106106a757634e487b7160e01b600052603260045260246000fd5b6020026020010151805190602001208282815181106106d657634e487b7160e01b600052603260045260246000fd5b60200260200101518051906020012014156107185781818151811061070b57634e487b7160e01b600052603260045260246000fd5b6020026020010160608152505b8061072281615e1e565b91505061067e565b506005546040517fa0ad3d080000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a0ad3d0890610776908b908590600401615766565b600060405180830381600087803b15801561079057600080fd5b505af11580156107a4573d6000803e3d6000fd5b505050505080806107b490615e1e565b915050610665565b50505050505050565b604080516060810182526000808252602082018190528183015260055491517f178e4dc000000000000000000000000000000000000000000000000000000000815290916001600160a01b03169063178e4dc09061082b90879087908790600401615b9d565b60606040518083038186803b15801561084357600080fd5b505afa158015610857573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087b9190614bd5565b90505b9392505050565b6005546040517f207e33be0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063207e33be906108ce908490600401615a7f565b600060405180830381600087803b1580156108e857600080fd5b505af11580156108fc573d6000803e3d6000fd5b5050505050565b6005546040517f28a8565c0000000000000000000000000000000000000000000000000000000081526060916001600160a01b0316906328a8565c9061094d908590600401615758565b60006040518083038186803b15801561096557600080fd5b505afa158015610979573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109a19190810190614640565b92915050565b6005546040517f34fddaac0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906334fddaac906109f29085908590600401615786565b600060405180830381600087803b158015610a0c57600080fd5b505af1158015610a20573d6000803e3d6000fd5b505050505050565b6060600080600086516001600160401b03811115610a5657634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a8957816020015b6060815260200190600190039081610a745790505b5090506000808060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015610ae057600080fd5b505afa158015610af4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b189190614b3c565b905060005b8b51811015610cca576000610b588d8381518110610b4b57634e487b7160e01b600052603260045260246000fd5b602002602001015161250f565b90506000805b8c51811015610bf4578c8181518110610b8757634e487b7160e01b600052603260045260246000fd5b60200260200101516001811115610bae57634e487b7160e01b600052602160045260246000fd5b836101e001516001811115610bd357634e487b7160e01b600052602160045260246000fd5b1415610be25760019150610bf4565b80610bec81615e1e565b915050610b5e565b5080610c01575050610cb8565b438460c001516001600160401b0316836101000151610c209190615c6e565b1115610c2d575050610cb8565b610140820151610c3d9087615c86565b9550610c4b826001806113a4565b8d8381518110610c6b57634e487b7160e01b600052603260045260246000fd5b602002602001015188886001600160401b031681518110610c9c57634e487b7160e01b600052603260045260246000fd5b60200260200101819052508680610cb290615e39565b97505050505b80610cc281615e1e565b915050610b1d565b506001600160401b03831615610d1a576040516001600160a01b038b16906001600160401b03851680156108fc02916000818181858888f19350505050158015610d18573d6000803e3d6000fd5b505b50929990985091965090945050505050565b6005546040517f3f2cc9a00000000000000000000000000000000000000000000000000000000081526060916000916001600160a01b0390911690633f2cc9a090610d7f908890889088906004016157d2565b600060405180830381600087803b158015610d9957600080fd5b505af1158015610dad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610dd591908101906147b1565b91509150935093915050565b604080516060810182526000808252602080830182905292820152908201516001600160401b0316610e2e5760405162461bcd60e51b8152600401610e2590615a2d565b60405180910390fd5b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015610e7e57600080fd5b505afa158015610e92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb69190614b3c565b905061087e83824361240f565b60008082516001600160401b03811115610eed57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610f2657816020015b610f1361300a565b815260200190600190039081610f0b5790505b50905060005b8351811015610fa3576000610f5a858381518110610b4b57634e487b7160e01b600052603260045260246000fd5b90508060200151935080838381518110610f8457634e487b7160e01b600052603260045260246000fd5b6020026020010181905250508080610f9b90615e1e565b915050610f2c565b506000610faf826116df565b805190915015610ff5577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051610fe891906159bc565b60405180910390a1611034565b7fa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec600143868660405161102b94939291906158b0565b60405180910390a15b50505050565b6005546040516395b0b54b60e01b81526001600160a01b03909116906395b0b54b9061106a90849060040161585f565b600060405180830381600087803b15801561108457600080fd5b505af1158015611098573d6000803e3d6000fd5b5050600480546040517f52e061460000000000000000000000000000000000000000000000000000000081526001600160a01b0390911693506352e0614692506108ce9185910161585f565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561113457600080fd5b505afa158015611148573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116c9190614b3c565b6005546040517f2e37891a0000000000000000000000000000000000000000000000000000000081529192506001600160a01b031690632e37891a9034906111ba9085908790600401615ad4565b6000604051808303818588803b1580156111d357600080fd5b505af11580156107bc573d6000803e3d6000fd5b6005546040517f5a0e74820000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690635a0e7482906109f29085908590600401615890565b6005546040517f681850d70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063681850d7906108ce908490600401615a5d565b6005546040517ff161a50e0000000000000000000000000000000000000000000000000000000081526060916000916001600160a01b039091169063f161a50e906112ca908690600401615758565b60006040518083038186803b1580156112e257600080fd5b505afa1580156112f6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261131e9190810190614640565b905061087e8184612de7565b6005546040516395b0b54b60e01b81526001600160a01b03909116906395b0b54b906108ce90849060040161585f565b6005546040517f9a2b5e630000000000000000000000000000000000000000000000000000000081526060916001600160a01b031690639a2b5e639061094d908590600401615758565b81156114435782516113b59061132a565b82516113c09061103a565b600554602084015184516040517fb64ab1ef0000000000000000000000000000000000000000000000000000000081526001600160a01b039093169263b64ab1ef92611410929091600401615786565b600060405180830381600087803b15801561142a57600080fd5b505af115801561143e573d6000803e3d6000fd5b505050505b80156114ac576005546040517fe457c9d90000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e457c9d990611492908690600401615a5d565b600060405180830381600087803b1580156111d357600080fd5b505050565b6005546040517fb6af10fb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063b6af10fb906108ce908490600401615a90565b600054610100900460ff166115155760005460ff1615611519565b303b155b6115355760405162461bcd60e51b8152600401610e25906159cc565b600054610100900460ff16158015611557576000805461ffff19166101011790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff16620100006001600160a01b038b811691909102919091179091556001805473ffffffffffffffffffffffffffffffffffffffff199081168a841617909155600280548216898416179055600380548216888416179055600480549091169186169190911790558251600580547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000006001600160401b0393841681029190911791829055602086015161164a93919092041690615cb9565b6006805460408601516001600160401b0390811668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909216931692909217919091179055600580546001600160a01b03841673ffffffffffffffffffffffffffffffffffffffff1990911617905580156116d5576000805461ff00191690555b5050505050505050565b606081516000141561172457505060408051808201909152600881527f6e6f2066696c6573000000000000000000000000000000000000000000000000602082015290565b6000808360008151811061174857634e487b7160e01b600052603260045260246000fd5b602002602001015160200151905060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156117a657600080fd5b505afa1580156117ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117de9190614b3c565b905060005b855181101561215557600086828151811061180e57634e487b7160e01b600052603260045260246000fd5b60200260200101519050836001600160a01b031681602001516001600160a01b03161461187457505060408051808201909152601781527f66696c65206f776e657220697320646966666572656e74000000000000000000602082015295945050505050565b60055481516040517f4968e2d60000000000000000000000000000000000000000000000000000000081526000926001600160a01b031691634968e2d6916118bf919060040161585f565b60006040518083038186803b1580156118d757600080fd5b505afa1580156118eb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611913919081019061482b565b905060005b8151811015611a615760035482516000916001600160a01b031690632ba010d79085908590811061195957634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040161197d9190615ac6565b60006040518083038186803b15801561199557600080fd5b505afa1580156119a9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119d19190810190614b08565b6003546040517e47c0030000000000000000000000000000000000000000000000000000000081529192506001600160a01b0316906247c00390611a1b9084908890600401615aa1565b600060405180830381600087803b158015611a3557600080fd5b505af1158015611a49573d6000803e3d6000fd5b50505050508080611a5990615e1e565b915050611918565b506101408201516001600160401b0316611a8857611a81826001806113a4565b5050612143565b6004805483516040517f977fdfe20000000000000000000000000000000000000000000000000000000081526000936001600160a01b039093169263977fdfe292611ad59290910161585f565b60006040518083038186803b158015611aed57600080fd5b505afa158015611b01573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b2991908101906147f7565b90506000836101400151905060008460a001518560800151611b4b9190615cda565b6005546040517fca6142cb0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063ca6142cb90611b9c908b908690600401615af6565b60206040518083038186803b158015611bb457600080fd5b505afa158015611bc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bec9190614c7d565b905060005b8451811015611ec35760015485516000916001600160a01b031690631a65374a90889085908110611c3257634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b8152600401611c5a9190615758565b60006040518083038186803b158015611c7257600080fd5b505afa158015611c86573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611cae9190810190614a6c565b90506000836001888581518110611cd557634e487b7160e01b600052603260045260246000fd5b602002602001015160400151611ceb9190615d31565b611cf59190615cda565b6005546101a08b01519192506000916001600160a01b039091169063df49410a908e908990611d249043615d16565b6040518463ffffffff1660e01b8152600401611d4293929190615b13565b60206040518083038186803b158015611d5a57600080fd5b505afa158015611d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d929190614c7d565b90506000611da08284615c86565b9050876001600160401b0316816001600160401b03161115611e0557505060408051808201909152601181527f70726f66697420697320696e76616c696400000000000000000000000000000060208201529f9e505050505050505050505050505050565b8084602001818151611e179190615c86565b6001600160401b03169052506001546040517ffc2153580000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063fc21535890611e6c908790600401615a6e565b600060405180830381600087803b158015611e8657600080fd5b505af1158015611e9a573d6000803e3d6000fd5b505050508088611eaa9190615d31565b9750505050508080611ebb90615e1e565b915050611bf1565b506000866101e001516001811115611eeb57634e487b7160e01b600052602160045260246000fd5b1415611f0257611efb838b615c86565b9950612130565b6001866101e001516001811115611f2957634e487b7160e01b600052602160045260246000fd5b1415612130576002546020870151604051632a7069e160e11b81526000926001600160a01b0316916354e0d3c291611f649190600401615758565b60a06040518083038186803b158015611f7c57600080fd5b505afa158015611f90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fb49190614c5f565b90508660a001518760800151611fca9190615cda565b6001600160401b031681600001516001600160401b031610612069578381604001818151611ff89190615c86565b6001600160401b031690525060a087015160808801516120189190615cda565b816020018181516120299190615c86565b6001600160401b031690525060a087015160808801516120499190615cda565b81518290612058908390615d31565b6001600160401b03169052506120ae565b505060408051808201909152601581527f7573657220737061636520697320696e76616c6964000000000000000000000060208201529b9a5050505050505050505050565b60025460208801516040517fed108de90000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163ed108de9916120fc9185906004016157a6565b600060405180830381600087803b15801561211657600080fd5b505af115801561212a573d6000803e3d6000fd5b50505050505b61213c866001806113a4565b5050505050505b8061214d81615e1e565b9150506117e3565b506001600160401b038316156121a5576040516001600160a01b038316906001600160401b03851680156108fc02916000818181858888f193505050501580156121a3573d6000803e3d6000fd5b505b50506040805160208101909152600081529392505050565b6005546040517fc6e8392a0000000000000000000000000000000000000000000000000000000081526060916001600160a01b03169063c6e8392a906122079085906004016157c1565b60006040518083038186803b15801561221f57600080fd5b505afa158015612233573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109a1919081019061477d565b604080516060810182526000808252602082018190528183015260055491517fcc76e80d00000000000000000000000000000000000000000000000000000000815290916001600160a01b03169063cc76e80d906122c59089908990899089908990600401615b3e565b60606040518083038186803b1580156122dd57600080fd5b505afa1580156122f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123159190614bd5565b90505b95945050505050565b600061232c8261250f565b60408051600180825281830190925291925060009190816020015b61234f61300a565b815260200190600190039081612347579050509050818160008151811061238657634e487b7160e01b600052603260045260246000fd5b6020026020010181905250600061239c826116df565b8051909150156123d5577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051610fe89190615988565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f4560014386866020015160405161102b94939291906158ec565b604080516060810182526000808252602082018190528183015260055491517fd49ce87400000000000000000000000000000000000000000000000000000000815290916001600160a01b03169063d49ce8749061082b90879087908790600401615b9d565b6005546040517fd70e62720000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063d70e6272906109f29085908590600401615766565b6005546040517f0c18d9cc0000000000000000000000000000000000000000000000000000000081526060916000916001600160a01b0390911690630c18d9cc906112ca908690600401615758565b61251761300a565b6005546040516337bf397560e21b81526000916001600160a01b03169063defce5d49061254890869060040161585f565b60006040518083038186803b15801561256057600080fd5b505afa158015612574573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261259c91908101906149ab565b8051519091506109a157826040517f6c5249c1000000000000000000000000000000000000000000000000000000008152600401610e25919061585f565b60055481516040516337bf397560e21b81526001600160a01b039092169163defce5d49161260a9160040161585f565b60006040518083038186803b15801561262257600080fd5b505afa158015612636573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261265e91908101906149ab565b6101a00151156126805760405162461bcd60e51b8152600401610e2590615a3d565b43816101000151116126a45760405162461bcd60e51b8152600401610e25906159ac565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156126f457600080fd5b505afa158015612708573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061272c9190614b3c565b90506000826102800151600281111561275557634e487b7160e01b600052602160045260246000fd5b141561276d576006546001600160401b031660c08301525b6001826102800151600281111561279457634e487b7160e01b600052602160045260246000fd5b14156127ac576006546001600160401b031660c08301525b600282610280015160028111156127d357634e487b7160e01b600052602160045260246000fd5b14156127eb576006546001600160401b031660c08301525b60016101c083015261286f604080516101e08101825260608082526000602083018190529282018390528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018190526101208201839052610140820183905261016082018190526101808201526101a08101829052906101c082015290565b61010083015160808083019190915260c0808501516001600160401b039081166040850152610120860151169083015283015160a08401516128b19190615cda565b6001600160401b0316602082015260006128cc8284436107c5565b90508060400151816020015182600001516128e79190615c86565b6128f19190615c86565b6001600160401b03166101408501526005546040517fe4bca9730000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e4bca9739061294b9085904390600401615baf565b60206040518083038186803b15801561296357600080fd5b505afa158015612977573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061299b9190614c7d565b6001600160401b031660e08501526000846101e0015160018111156129d057634e487b7160e01b600052602160045260246000fd5b1415612c08576002546020850151604051632a7069e160e11b81526000926001600160a01b0316916354e0d3c291612a0b9190600401615758565b60a06040518083038186803b158015612a2357600080fd5b505afa158015612a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a5b9190614c5f565b90508461014001516001600160401b031681604001516001600160401b03161015612abf5780604001518561014001516040517f62fe63d9000000000000000000000000000000000000000000000000000000008152600401610e25929190615bea565b84608001518560a00151612ad39190615cda565b6001600160401b031681602001516001600160401b03161015612b2e5780602001518560a001516040517fd39dfd03000000000000000000000000000000000000000000000000000000008152600401610e25929190615bea565b84610100015181606001511015612b7e5780606001518561010001516040517f0c803412000000000000000000000000000000000000000000000000000000008152600401610e25929190615bcf565b84610140015181604001818151612b959190615d31565b6001600160401b031690525060a08501516080860151612bb59190615cda565b81602001818151612bc69190615d31565b6001600160401b031690525060a08501516080860151612be69190615cda565b81518290612bf5908390615c86565b6001600160401b0316905250612c3f9050565b8361014001516001600160401b0316341015612c365760405162461bcd60e51b8152600401610e2590615a4d565b60016101e08501525b60808301516001600160401b0316610180850152436101a08501526005546040517f3947b1060000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690633947b10690612ca3908790600401615a5d565b600060405180830381600087803b158015612cbd57600080fd5b505af1158015612cd1573d6000803e3d6000fd5b505060408051808201909152600080825260208201529150612cf09050565b6101208501516001600160401b03168152600060208201526004805486516040517fbb4e4e420000000000000000000000000000000000000000000000000000000081526001600160a01b039092169263bb4e4e4292612d539291869101615870565b600060405180830381600087803b158015612d6d57600080fd5b505af1158015612d81573d6000803e3d6000fd5b505050507fa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18600043876000015188610200015189602001518a61014001518b6102a00151604051612dd89796959493929190615919565b60405180910390a15050505050565b6060600083516001600160401b03811115612e1257634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015612e4557816020015b6060815260200190600190039081612e305790505b5090506000805b85518110156130005760045486516000916001600160a01b03169063977fdfe290899085908110612e8d57634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b8152600401612eb1919061585f565b60006040518083038186803b158015612ec957600080fd5b505afa158015612edd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612f0591908101906147f7565b90506000805b8251811015612f7457876001600160a01b0316838281518110612f3e57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b03161415612f625760019150612f74565b80612f6c81615e1e565b915050612f0b565b5080612f81575050612fee565b878381518110612fa157634e487b7160e01b600052603260045260246000fd5b602002602001015185856001600160401b031681518110612fd257634e487b7160e01b600052603260045260246000fd5b60200260200101819052508380612fe890615e39565b94505050505b80612ff881615e1e565b915050612e4c565b5090949350505050565b604080516102e08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201839052610160820181905261018082018390526101a082018390526101c082018390526101e0820183905261020082018390526102208201819052610240820181905261026082015290610280820190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b60006130ef6130ea84615c21565b615c05565b9050808382526020820190508285602086028201111561310e57600080fd5b60005b8581101561313a57816131248882613530565b8452506020928301929190910190600101613111565b5050509392505050565b60006131526130ea84615c21565b9050808382526020820190508285602086028201111561317157600080fd5b60005b8581101561313a5781613187888261353b565b8452506020928301929190910190600101613174565b60006131ab6130ea84615c21565b905080838252602082019050828560208602820111156131ca57600080fd5b60005b8581101561313a5781356001600160401b038111156131eb57600080fd5b8086016131f889826136ae565b8552505060209283019291909101906001016131cd565b600061321d6130ea84615c21565b9050808382526020820190508285602086028201111561323c57600080fd5b60005b8581101561313a5781516001600160401b0381111561325d57600080fd5b80860161326a89826136cf565b85525050602092830192919091019060010161323f565b600061328f6130ea84615c21565b905080838252602082019050828560208602820111156132ae57600080fd5b60005b8581101561313a57816132c48882613711565b84525060209283019291909101906001016132b1565b60006132e86130ea84615c21565b9050808382526020820190508285602086028201111561330757600080fd5b60005b8581101561313a5781356001600160401b0381111561332857600080fd5b8086016133358982613782565b85525050602092830192919091019060010161330a565b600061335a6130ea84615c21565b9050808382526020820190508285602086028201111561337957600080fd5b60005b8581101561313a5781516001600160401b0381111561339a57600080fd5b8086016133a78982613a16565b85525050602092830192919091019060010161337c565b60006133cc6130ea84615c21565b8381529050602081018260a0850281018610156133e857600080fd5b60005b8581101561313a57816133fe8882613ec2565b84525060209092019160a091909101906001016133eb565b60006134246130ea84615c21565b8381529050602081018260408502810186101561344057600080fd5b60005b8581101561313a5781613456888261409f565b84525060209092019160409190910190600101613443565b600061347c6130ea84615c21565b8381529050602081018260608502810186101561349857600080fd5b60005b8581101561313a57816134ae8882614500565b8452506020909201916060919091019060010161349b565b60006134d46130ea84615c44565b9050828152602081018484840111156134ec57600080fd5b6134f7848285615dba565b509392505050565b600061350d6130ea84615c44565b90508281526020810184848401111561352557600080fd5b6134f7848285615dc6565b80356109a181615ee5565b80516109a181615ee5565b600082601f83011261355757600080fd5b81356135678482602086016130dc565b949350505050565b600082601f83011261358057600080fd5b8151613567848260208601613144565b600082601f8301126135a157600080fd5b813561356784826020860161319d565b600082601f8301126135c257600080fd5b815161356784826020860161320f565b600082601f8301126135e357600080fd5b8135613567848260208601613281565b600082601f83011261360457600080fd5b81356135678482602086016132da565b600082601f83011261362557600080fd5b815161356784826020860161334c565b600082601f83011261364657600080fd5b81516135678482602086016133be565b600082601f83011261366757600080fd5b8151613567848260208601613416565b600082601f83011261368857600080fd5b813561356784826020860161346e565b80356109a181615ef9565b80516109a181615ef9565b600082601f8301126136bf57600080fd5b81356135678482602086016134c6565b600082601f8301126136e057600080fd5b81516135678482602086016134ff565b80356109a181615f01565b80356109a181615f0a565b80516109a181615f0a565b80356109a181615f17565b80516109a181615f17565b60006060828403121561373957600080fd5b6137436060615c05565b905060006137518484614540565b825250602061376284848301614540565b602083015250604061377684828501614540565b60408301525092915050565b6000610320828403121561379557600080fd5b6137a06102e0615c05565b905081356001600160401b038111156137b857600080fd5b6137c4848285016136ae565b82525060206137d584848301613530565b60208301525060408201356001600160401b038111156137f457600080fd5b613800848285016136ae565b604083015250606061381484828501614540565b606083015250608061382884828501614540565b60808301525060a061383c84828501614540565b60a08301525060c061385084828501614540565b60c08301525060e061386484828501614540565b60e0830152506101006138798482850161452a565b6101008301525061012061388f84828501614540565b610120830152506101406138a584828501614540565b610140830152506101608201356001600160401b038111156138c657600080fd5b6138d2848285016136ae565b610160830152506101806138e884828501614540565b610180830152506101a06138fe8482850161452a565b6101a0830152506101c061391484828501613698565b6101c0830152506101e061392a84828501613711565b6101e08301525061020061394084828501614540565b610200830152506102208201356001600160401b0381111561396157600080fd5b61396d84828501613546565b610220830152506102408201356001600160401b0381111561398e57600080fd5b61399a84828501613546565b610240830152506102608201356001600160401b038111156139bb57600080fd5b6139c7848285016136ae565b610260830152506102806139dd848285016136fb565b610280830152506102a06139f384828501613698565b6102a0830152506102c0613a0984828501613727565b6102c08301525092915050565b60006103208284031215613a2957600080fd5b613a346102e0615c05565b82519091506001600160401b03811115613a4d57600080fd5b613a59848285016136cf565b8252506020613a6a8484830161353b565b60208301525060408201516001600160401b03811115613a8957600080fd5b613a95848285016136cf565b6040830152506060613aa98482850161454b565b6060830152506080613abd8482850161454b565b60808301525060a0613ad18482850161454b565b60a08301525060c0613ae58482850161454b565b60c08301525060e0613af98482850161454b565b60e083015250610100613b0e84828501614535565b61010083015250610120613b248482850161454b565b61012083015250610140613b3a8482850161454b565b610140830152506101608201516001600160401b03811115613b5b57600080fd5b613b67848285016136cf565b61016083015250610180613b7d8482850161454b565b610180830152506101a0613b9384828501614535565b6101a0830152506101c0613ba9848285016136a3565b6101c0830152506101e0613bbf8482850161371c565b6101e083015250610200613bd58482850161454b565b610200830152506102208201516001600160401b03811115613bf657600080fd5b613c028482850161356f565b610220830152506102408201516001600160401b03811115613c2357600080fd5b613c2f8482850161356f565b610240830152506102608201516001600160401b03811115613c5057600080fd5b613c5c848285016136cf565b61026083015250610280613c7284828501613706565b610280830152506102a0613c88848285016136a3565b6102a0830152506102c0613a0984828501613e16565b600060608284031215613cb057600080fd5b613cba6060615c05565b905081356001600160401b03811115613cd257600080fd5b613cde848285016136ae565b825250602061376284848301613530565b600060e08284031215613d0157600080fd5b613d0b60e0615c05565b90506000613d19848461454b565b8252506020613d2a8484830161454b565b6020830152506040613d3e8482850161454b565b6040830152506060613d528482850161454b565b6060830152506080613d668482850161454b565b60808301525060a0613d7a8482850161353b565b60a08301525060c08201516001600160401b03811115613d9957600080fd5b613da5848285016136cf565b60c08301525092915050565b600060608284031215613dc357600080fd5b613dcd6060615c05565b905081356001600160401b03811115613de557600080fd5b613df1848285016136ae565b8252506020613e0284848301613530565b602083015250604061377684828501613530565b600060608284031215613e2857600080fd5b613e326060615c05565b90506000613e40848461454b565b8252506020613e518484830161454b565b60208301525060406137768482850161454b565b600060408284031215613e7757600080fd5b613e816040615c05565b905081356001600160401b03811115613e9957600080fd5b613ea5848285016136ae565b8252506020613eb684848301614540565b60208301525092915050565b600060a08284031215613ed457600080fd5b613ede60a0615c05565b90506000613eec848461353b565b8252506020613efd8484830161353b565b6020830152506040613f118482850161454b565b6040830152506060613f2584828501614535565b6060830152506080613f39848285016136a3565b60808301525092915050565b60006101808284031215613f5857600080fd5b613f63610180615c05565b90506000613f71848461353b565b8252506020613f828484830161454b565b6020830152506040613f968482850161454b565b6040830152506060613faa8482850161454b565b6060830152506080613fbe84828501613706565b60808301525060a0613fd284828501614535565b60a08301525060c0613fe684828501614535565b60c08301525060e0613ffa8482850161454b565b60e08301525061010061400f8482850161454b565b610100830152506101206140258482850161454b565b6101208301525061014061403b848285016136a3565b610140830152506101608201516001600160401b0381111561405c57600080fd5b614068848285016135b1565b6101608301525092915050565b60006040828403121561408757600080fd5b6140916040615c05565b90506000613ea58484613530565b6000604082840312156140b157600080fd5b6140bb6040615c05565b905060006140c9848461353b565b8252506020613eb68484830161454b565b600061016082840312156140ed57600080fd5b6140f8610160615c05565b905060006141068484614540565b825250602061411784848301614540565b602083015250604061412b84828501614540565b604083015250606061413f84828501614540565b606083015250608061415384828501614540565b60808301525060a061416784828501614540565b60a08301525060c061417b84828501614540565b60c08301525060e061418f84828501614540565b60e0830152506101006141a484828501614540565b610100830152506101206141ba84828501614540565b610120830152506101406141d084828501614540565b6101408301525092915050565b600061016082840312156141f057600080fd5b6141fb610160615c05565b90506000614209848461454b565b825250602061421a8484830161454b565b602083015250604061422e8482850161454b565b60408301525060606142428482850161454b565b60608301525060806142568482850161454b565b60808301525060a061426a8482850161454b565b60a08301525060c061427e8482850161454b565b60c08301525060e06142928482850161454b565b60e0830152506101006142a78482850161454b565b610100830152506101206142bd8482850161454b565b610120830152506101406141d08482850161454b565b60006101e082840312156142e657600080fd5b6142f16101e0615c05565b905081356001600160401b0381111561430957600080fd5b614315848285016136ae565b825250602061432684848301614540565b602083015250604061433a84828501614540565b604083015250606061434e84828501614540565b60608301525060806143628482850161452a565b60808301525060a061437684828501614540565b60a08301525060c061438a84828501614540565b60c08301525060e061439e84828501613698565b60e0830152506101008201356001600160401b038111156143be57600080fd5b6143ca848285016136ae565b610100830152506101206143e084828501613698565b610120830152506101406143f684828501613698565b610140830152506101608201356001600160401b0381111561441757600080fd5b614423848285016136ae565b610160830152506101808201356001600160401b0381111561444457600080fd5b61445084828501613677565b610180830152506101a061446684828501613698565b6101a0830152506101c061447c84828501613711565b6101c08301525092915050565b600060a0828403121561449b57600080fd5b6144a560a0615c05565b905060006144b3848461454b565b82525060206144c48484830161454b565b60208301525060406144d88482850161454b565b60408301525060606144ec84828501614535565b6060830152506080613f3984828501614535565b60006060828403121561451257600080fd5b61451c6060615c05565b905060006137518484613530565b80356109a181615f24565b80516109a181615f24565b80356109a181615f2a565b80516109a181615f2a565b60006020828403121561456857600080fd5b60006135678484613530565b6000806040838503121561458757600080fd5b60006145938585613530565b92505060208301356001600160401b038111156145af57600080fd5b6145bb85828601613590565b9150509250929050565b600080604083850312156145d857600080fd5b60006145e48585613530565b92505060208301356001600160401b0381111561460057600080fd5b6145bb858286016136ae565b60006020828403121561461e57600080fd5b81356001600160401b0381111561463457600080fd5b61356784828501613590565b60006020828403121561465257600080fd5b81516001600160401b0381111561466857600080fd5b613567848285016135b1565b60008060006060848603121561468957600080fd5b83356001600160401b0381111561469f57600080fd5b6146ab86828701613590565b93505060206146bc86828701613530565b92505060408401356001600160401b038111156146d857600080fd5b6146e4868287016135d2565b9150509250925092565b60008060006101a0848603121561470457600080fd5b83356001600160401b0381111561471a57600080fd5b61472686828701613590565b9350506020614737868287016140da565b9250506101806146e48682870161452a565b60006020828403121561475b57600080fd5b81356001600160401b0381111561477157600080fd5b613567848285016135f3565b60006020828403121561478f57600080fd5b81516001600160401b038111156147a557600080fd5b61356784828501613614565b600080604083850312156147c457600080fd5b82516001600160401b038111156147da57600080fd5b6147e685828601613614565b92505060206145bb858286016136a3565b60006020828403121561480957600080fd5b81516001600160401b0381111561481f57600080fd5b61356784828501613635565b60006020828403121561483d57600080fd5b81516001600160401b0381111561485357600080fd5b61356784828501613656565b60006020828403121561487157600080fd5b81356001600160401b0381111561488757600080fd5b613567848285016136ae565b600080606083850312156148a657600080fd5b82356001600160401b038111156148bc57600080fd5b6148c8858286016136ae565b92505060206145bb85828601614075565b6000806000806000806000610120888a0312156148f557600080fd5b60006149018a8a6136f0565b97505060206149128a828b016136f0565b96505060406149238a828b016136f0565b95505060606149348a828b016136f0565b94505060806149458a828b016136f0565b93505060a06149568a828b01613727565b9250506101006149688a828b016136f0565b91505092959891949750929550565b60006020828403121561498957600080fd5b81356001600160401b0381111561499f57600080fd5b61356784828501613782565b6000602082840312156149bd57600080fd5b81516001600160401b038111156149d357600080fd5b61356784828501613a16565b6000806000606084860312156149f457600080fd5b83356001600160401b03811115614a0a57600080fd5b614a1686828701613782565b9350506020614a2786828701613698565b92505060406146e486828701613698565b600060208284031215614a4a57600080fd5b81356001600160401b03811115614a6057600080fd5b61356784828501613c9e565b600060208284031215614a7e57600080fd5b81516001600160401b03811115614a9457600080fd5b61356784828501613cef565b600060208284031215614ab257600080fd5b81356001600160401b03811115614ac857600080fd5b61356784828501613db1565b600060208284031215614ae657600080fd5b81356001600160401b03811115614afc57600080fd5b61356784828501613e65565b600060208284031215614b1a57600080fd5b81516001600160401b03811115614b3057600080fd5b61356784828501613f45565b60006101608284031215614b4f57600080fd5b600061356784846141dd565b60008060008060006101e08688031215614b7457600080fd5b6000614b8088886140da565b955050610160614b9288828901614540565b945050610180614ba488828901614540565b9350506101a0614bb688828901614540565b9250506101c0614bc888828901614540565b9150509295509295909350565b600060608284031215614be757600080fd5b60006135678484613e16565b600060208284031215614c0557600080fd5b81356001600160401b03811115614c1b57600080fd5b613567848285016142d3565b60008060006101a08486031215614c3d57600080fd5b83356001600160401b03811115614c5357600080fd5b614726868287016142d3565b600060a08284031215614c7157600080fd5b60006135678484614489565b600060208284031215614c8f57600080fd5b6000613567848461454b565b6000614ca78383614cdb565b505060200190565b600061087e8383614e91565b600061087e8383615019565b6000614cd38383615729565b505060600190565b614ce481615d4e565b82525050565b6000614cf4825190565b80845260209384019383018060005b83811015614d28578151614d178882614c9b565b975060208301925050600101614d03565b509495945050505050565b6000614d3d825190565b80845260208401935083602082028501614d578560200190565b8060005b85811015614d8c5784840389528151614d748582614caf565b94506020830160209a909a0199925050600101614d5b565b5091979650505050505050565b6000614da3825190565b80845260208401935083602082028501614dbd8560200190565b8060005b85811015614d8c5784840389528151614dda8582614caf565b94506020830160209a909a0199925050600101614dc1565b6000614dfc825190565b80845260208401935083602082028501614e168560200190565b8060005b85811015614d8c5784840389528151614e338582614cbb565b94506020830160209a909a0199925050600101614e1a565b6000614e55825190565b80845260209384019383018060005b83811015614d28578151614e788882614cc7565b975060208301925050600101614e64565b801515614ce4565b6000614e9b825190565b808452602084019350614eb2818560208601615dc6565b601f01601f19169290920192915050565b614ce481615d88565b614ce481615d93565b614ce481615d9e565b600a81526000602082017f44656c65746546696c6500000000000000000000000000000000000000000000815291505b5060200190565b600c81526000602082017f66696c652065787069726564000000000000000000000000000000000000000081529150614f0e565b600b81526000602082017f44656c65746546696c657300000000000000000000000000000000000000000081529150614f0e565b601f81526000602082017f66696c6553697a65206d7573742062652067726561746572207468616e20300081529150614f0e565b601281526000602082017f66696c6520616c7265616479206578697374000000000000000000000000000081529150614f0e565b601481526000602082017f696e73756666696369656e74206465706f73697400000000000000000000000081529150614f0e565b8051610320808452600091908401906150328282614e91565b91505060208301516150476020860182614cdb565b506040830151848203604086015261505f8282614e91565b91505060608301516150746060860182615749565b5060808301516150876080860182615749565b5060a083015161509a60a0860182615749565b5060c08301516150ad60c0860182615749565b5060e08301516150c060e0860182615749565b506101008301516150d561010086018261573a565b506101208301516150ea610120860182615749565b506101408301516150ff610140860182615749565b506101608301518482036101608601526151198282614e91565b915050610180830151615130610180860182615749565b506101a08301516151456101a086018261573a565b506101c083015161515a6101c0860182614e89565b506101e083015161516f6101e0860182614ed5565b50610200830151615184610200860182615749565b5061022083015184820361022086015261519e8282614cea565b9150506102408301518482036102408601526151ba8282614cea565b9150506102608301518482036102608601526151d68282614e91565b9150506102808301516151ed610280860182614ecc565b506102a08301516152026102a0860182614e89565b506102c08301516134f76102c0860182615322565b805160608084526000919084019061522f8282614e91565b91505060208301516152446020860182614cdb565b5060408301516134f76040860182615749565b805160009060e084019061526b8582615749565b50602083015161527e6020860182615749565b5060408301516152916040860182615749565b5060608301516152a46060860182615749565b5060808301516152b76080860182615749565b5060a08301516152ca60a0860182614cdb565b5060c083015184820360c08601526123188282614e91565b80516060808452600091908401906152fa8282614e91565b915050602083015161530f6020860182614cdb565b5060408301516134f76040860182614cdb565b805160608301906153338482615749565b5060208201516153466020850182615749565b5060408201516110346040850182615749565b80516040808452600091908401906153718282614e91565b91505060208301516134f76020860182615749565b805160408301906153978482615749565b5060208201516110346020850182615749565b80516000906101808401906153bf8582614cdb565b5060208301516153d26020860182615749565b5060408301516153e56040860182615749565b5060608301516153f86060860182615749565b50608083015161540b6080860182614ecc565b5060a083015161541e60a086018261573a565b5060c083015161543160c086018261573a565b5060e083015161544460e0860182615749565b50610100830151615459610100860182615749565b5061012083015161546e610120860182615749565b50610140830151615483610140860182614e89565b506101608301518482036101608601526123188282614d33565b805160408301906153978482614cdb565b80516101608301906154c08482615749565b5060208201516154d36020850182615749565b5060408201516154e66040850182615749565b5060608201516154f96060850182615749565b50608082015161550c6080850182615749565b5060a082015161551f60a0850182615749565b5060c082015161553260c0850182615749565b5060e082015161554560e0850182615749565b5061010082015161555a610100850182615749565b5061012082015161556f610120850182615749565b50610140820151611034610140850182615749565b80516101e08084526000919084019061559d8282614e91565b91505060208301516155b26020860182615749565b5060408301516155c56040860182615749565b5060608301516155d86060860182615749565b5060808301516155eb608086018261573a565b5060a08301516155fe60a0860182615749565b5060c083015161561160c0860182615749565b5060e083015161562460e0860182614e89565b5061010083015184820361010086015261563e8282614e91565b915050610120830151615655610120860182614e89565b5061014083015161566a610140860182614e89565b506101608301518482036101608601526156848282614e91565b9150506101808301518482036101808601526156a08282614e4b565b9150506101a08301516156b76101a0860182614e89565b506101c08301516134f76101c0860182614ed5565b805160a08301906156dd8482615749565b5060208201516156f06020850182615749565b5060408201516157036040850182615749565b506060820151615716606085018261573a565b506080820151611034608085018261573a565b805160608301906153338482614cdb565b80614ce4565b614ce481615da9565b6001600160401b038116614ce4565b602081016109a18284614cdb565b604081016157748285614cdb565b818103602083015261087b8184614d99565b604081016157948285614cdb565b818103602083015261087b8184614e91565b60c081016157b48285614cdb565b61087e60208301846156cc565b6020808252810161087e8184614d99565b6101a080825281016157e48186614d99565b90506157f360208301856154ae565b61356761018083018461573a565b606080825281016158128186614d99565b90506158216020830185615749565b6135676040830184614e89565b6020808252810161087e8184614df2565b604080825281016158508185614df2565b905061087e6020830184614e89565b6020808252810161087e8184614e91565b606080825281016158818185614e91565b905061087e6020830184615386565b606080825281016158a18185614e91565b905061087e602083018461549d565b608081016158be8287614ec3565b6158cb602083018661573a565b81810360408301526158dd8185614d99565b90506123186060830184614cdb565b608081016158fa8287614ec3565b615907602083018661573a565b81810360408301526158dd8185614e91565b60e08101615927828a614ec3565b615934602083018961573a565b81810360408301526159468188614e91565b90506159556060830187615749565b6159626080830186614cdb565b61596f60a0830185615749565b61597c60c0830184614e89565b98975050505050505050565b6040808252810161599881614ede565b9050818103602083015261087e8184614e91565b602080825281016109a181614f15565b6040808252810161599881614f49565b602080825281016109a181602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016109a181614f7d565b602080825281016109a181614fb1565b602080825281016109a181614fe5565b6020808252810161087e8184615019565b6020808252810161087e8184615257565b6020808252810161087e81846152e2565b6020808252810161087e8184615359565b60408082528101615ab281856153aa565b9050818103602083015261087b8184615019565b604081016109a1828461549d565b6101808101615ae382856154ae565b81810361016083015261087b8184615217565b6101808101615b0582856154ae565b61087e610160830184615749565b6101a08101615b2282866154ae565b615b30610160830185615749565b613567610180830184615749565b6101e08101615b4d82886154ae565b615b5b610160830187615749565b615b69610180830186615749565b615b776101a0830185615749565b615b856101c0830184615749565b9695505050505050565b606081016109a18284615322565b6101a080825281016157e48186615584565b60408082528101615bc08185615584565b905061087e602083018461573a565b60408101615bdd828561573a565b61087e602083018461573a565b60408101615bf88285615740565b61087e6020830184615740565b6000615c1060405190565b9050615c1c8282615df2565b919050565b60006001600160401b03821115615c3a57615c3a615e9c565b5060209081020190565b60006001600160401b03821115615c5d57615c5d615e9c565b601f19601f83011660200192915050565b60008219821115615c8157615c81615e5a565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115615c8157615c81615e5a565b6001600160401b039182169116600082615cd557615cd5615e70565b500490565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615615d1157615d11615e5a565b500290565b6000825b925082821015615d2c57615d2c615e5a565b500390565b60006001600160401b03821691506001600160401b038316615d1a565b60006001600160a01b0382166109a1565b60006109a182615d4e565b80615c1c81615eb2565b80615c1c81615ec5565b80615c1c81615ed5565b60006109a182615d6a565b60006109a182615d74565b60006109a182615d7e565b60006001600160401b0382166109a1565b82818337506000910152565b60005b83811015615de1578181015183820152602001615dc9565b838111156110345750506000910152565b601f19601f83011681018181106001600160401b0382111715615e1757615e17615e9c565b6040525050565b6000600019821415615e3257615e32615e5a565b5060010190565b60006001600160401b03821691506001600160401b03821415615e3257615e325b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600a8110615ec257615ec2615e86565b50565b60038110615ec257615ec2615e86565b60028110615ec257615ec2615e86565b615eee81615d4e565b8114615ec257600080fd5b801515615eee565b615eee81615d5f565b60038110615ec257600080fd5b60028110615ec257600080fd5b80615eee565b6001600160401b038116615eee56fea2646970667358221220b05b678aa456ffe74d4b8a16c7e3d43ca939f7bb0b19e3e2f4c1682c6807dffd64736f6c63430008040033",
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
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
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
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) view returns((uint64,uint64,uint64))
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
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
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

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
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
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// AddFileSectorRef is a paid mutator transaction binding the contract method 0x5a0e7482.
//
// Solidity: function AddFileSectorRef(bytes fileHash, (address,uint64) ref) payable returns()
func (_Store *StoreTransactor) AddFileSectorRef(opts *bind.TransactOpts, fileHash []byte, ref SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileSectorRef", fileHash, ref)
}

// AddFileSectorRef is a paid mutator transaction binding the contract method 0x5a0e7482.
//
// Solidity: function AddFileSectorRef(bytes fileHash, (address,uint64) ref) payable returns()
func (_Store *StoreSession) AddFileSectorRef(fileHash []byte, ref SectorRef) (*types.Transaction, error) {
	return _Store.Contract.AddFileSectorRef(&_Store.TransactOpts, fileHash, ref)
}

// AddFileSectorRef is a paid mutator transaction binding the contract method 0x5a0e7482.
//
// Solidity: function AddFileSectorRef(bytes fileHash, (address,uint64) ref) payable returns()
func (_Store *StoreTransactorSession) AddFileSectorRef(fileHash []byte, ref SectorRef) (*types.Transaction, error) {
	return _Store.Contract.AddFileSectorRef(&_Store.TransactOpts, fileHash, ref)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) payable returns()
func (_Store *StoreTransactor) AddFileToUnSettleList(opts *bind.TransactOpts, fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToUnSettleList", fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) payable returns()
func (_Store *StoreSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) payable returns()
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
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) payable returns()
func (_Store *StoreTransactor) CleanupForDeleteFile(opts *bind.TransactOpts, fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CleanupForDeleteFile", fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) payable returns()
func (_Store *StoreSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) payable returns()
func (_Store *StoreTransactorSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
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

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteProveDetails(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteProveDetails", fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) returns()
func (_Store *StoreSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
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
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) payable returns()
func (_Store *StoreTransactor) UpdateFileInfo(opts *bind.TransactOpts, f FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileInfo", f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) payable returns()
func (_Store *StoreSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) payable returns()
func (_Store *StoreTransactorSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) payable returns()
func (_Store *StoreTransactor) UpdateFileList(opts *bind.TransactOpts, walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileList", walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) payable returns()
func (_Store *StoreSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) payable returns()
func (_Store *StoreTransactorSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// UpdateFilesForRenew is a paid mutator transaction binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) payable returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreTransactor) UpdateFilesForRenew(opts *bind.TransactOpts, _fileList [][]byte, setting Setting, newExpireHeight *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFilesForRenew", _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a paid mutator transaction binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) payable returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) (*types.Transaction, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.TransactOpts, _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a paid mutator transaction binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) payable returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreTransactorSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) (*types.Transaction, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.TransactOpts, _fileList, setting, newExpireHeight)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns(string error)
func (_Store *StoreTransactor) DeleteFilesInner(opts *bind.TransactOpts, files []FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteFilesInner", files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns(string error)
func (_Store *StoreSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns(string error)
func (_Store *StoreTransactorSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// Initialize is a paid mutator transaction binding the contract method 0xbc1c783e.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig, address _fileExtra) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig, _fileExtra common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _node, _space, _sector, _prove, fsConfig, _fileExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xbc1c783e.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig, address _fileExtra) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig, _fileExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove, fsConfig, _fileExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xbc1c783e.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig, address _fileExtra) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig, _fileExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove, fsConfig, _fileExtra)
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
	ProveLevel  uint8
	Size        uint64
	IsPlots     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateSectorEvent is a free log retrieval operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) FilterCreateSectorEvent(opts *bind.FilterOpts) (*StoreCreateSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreCreateSectorEventIterator{contract: _Store.contract, event: "CreateSectorEvent", logs: logs, sub: sub}, nil
}

// WatchCreateSectorEvent is a free log subscription operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
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
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) ParseCreateSectorEvent(log types.Log) (*StoreCreateSectorEvent, error) {
	event := new(StoreCreateSectorEvent)
	if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDNSNodeRegisterIterator is returned from FilterDNSNodeRegister and is used to iterate over the raw logs and unpacked data for DNSNodeRegister events raised by the Store contract.
type StoreDNSNodeRegisterIterator struct {
	Event *StoreDNSNodeRegister // Event containing the contract specifics and raw log

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
func (it *StoreDNSNodeRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDNSNodeRegister)
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
		it.Event = new(StoreDNSNodeRegister)
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
func (it *StoreDNSNodeRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDNSNodeRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDNSNodeRegister represents a DNSNodeRegister event raised by the Store contract.
type StoreDNSNodeRegister struct {
	Ip         []byte
	Port       []byte
	WalletAddr common.Address
	Deposit    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDNSNodeRegister is a free log retrieval operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) FilterDNSNodeRegister(opts *bind.FilterOpts) (*StoreDNSNodeRegisterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DNSNodeRegister")
	if err != nil {
		return nil, err
	}
	return &StoreDNSNodeRegisterIterator{contract: _Store.contract, event: "DNSNodeRegister", logs: logs, sub: sub}, nil
}

// WatchDNSNodeRegister is a free log subscription operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) WatchDNSNodeRegister(opts *bind.WatchOpts, sink chan<- *StoreDNSNodeRegister) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DNSNodeRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDNSNodeRegister)
				if err := _Store.contract.UnpackLog(event, "DNSNodeRegister", log); err != nil {
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

// ParseDNSNodeRegister is a log parse operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) ParseDNSNodeRegister(log types.Log) (*StoreDNSNodeRegister, error) {
	event := new(StoreDNSNodeRegister)
	if err := _Store.contract.UnpackLog(event, "DNSNodeRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDNSNodeUnRegIterator is returned from FilterDNSNodeUnReg and is used to iterate over the raw logs and unpacked data for DNSNodeUnReg events raised by the Store contract.
type StoreDNSNodeUnRegIterator struct {
	Event *StoreDNSNodeUnReg // Event containing the contract specifics and raw log

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
func (it *StoreDNSNodeUnRegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDNSNodeUnReg)
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
		it.Event = new(StoreDNSNodeUnReg)
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
func (it *StoreDNSNodeUnRegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDNSNodeUnRegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDNSNodeUnReg represents a DNSNodeUnReg event raised by the Store contract.
type StoreDNSNodeUnReg struct {
	WalletAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDNSNodeUnReg is a free log retrieval operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) FilterDNSNodeUnReg(opts *bind.FilterOpts) (*StoreDNSNodeUnRegIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DNSNodeUnReg")
	if err != nil {
		return nil, err
	}
	return &StoreDNSNodeUnRegIterator{contract: _Store.contract, event: "DNSNodeUnReg", logs: logs, sub: sub}, nil
}

// WatchDNSNodeUnReg is a free log subscription operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) WatchDNSNodeUnReg(opts *bind.WatchOpts, sink chan<- *StoreDNSNodeUnReg) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DNSNodeUnReg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDNSNodeUnReg)
				if err := _Store.contract.UnpackLog(event, "DNSNodeUnReg", log); err != nil {
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

// ParseDNSNodeUnReg is a log parse operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) ParseDNSNodeUnReg(log types.Log) (*StoreDNSNodeUnReg, error) {
	event := new(StoreDNSNodeUnReg)
	if err := _Store.contract.UnpackLog(event, "DNSNodeUnReg", log); err != nil {
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

// StoreFsErrorIterator is returned from FilterFsError and is used to iterate over the raw logs and unpacked data for FsError events raised by the Store contract.
type StoreFsErrorIterator struct {
	Event *StoreFsError // Event containing the contract specifics and raw log

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
func (it *StoreFsErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFsError)
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
		it.Event = new(StoreFsError)
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
func (it *StoreFsErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFsErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFsError represents a FsError event raised by the Store contract.
type StoreFsError struct {
	Method string
	Msg    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFsError is a free log retrieval operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) FilterFsError(opts *bind.FilterOpts) (*StoreFsErrorIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FsError")
	if err != nil {
		return nil, err
	}
	return &StoreFsErrorIterator{contract: _Store.contract, event: "FsError", logs: logs, sub: sub}, nil
}

// WatchFsError is a free log subscription operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) WatchFsError(opts *bind.WatchOpts, sink chan<- *StoreFsError) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FsError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFsError)
				if err := _Store.contract.UnpackLog(event, "FsError", log); err != nil {
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

// ParseFsError is a log parse operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) ParseFsError(log types.Log) (*StoreFsError, error) {
	event := new(StoreFsError)
	if err := _Store.contract.UnpackLog(event, "FsError", log); err != nil {
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

// StoreNotifyHeaderAddIterator is returned from FilterNotifyHeaderAdd and is used to iterate over the raw logs and unpacked data for NotifyHeaderAdd events raised by the Store contract.
type StoreNotifyHeaderAddIterator struct {
	Event *StoreNotifyHeaderAdd // Event containing the contract specifics and raw log

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
func (it *StoreNotifyHeaderAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyHeaderAdd)
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
		it.Event = new(StoreNotifyHeaderAdd)
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
func (it *StoreNotifyHeaderAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyHeaderAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyHeaderAdd represents a NotifyHeaderAdd event raised by the Store contract.
type StoreNotifyHeaderAdd struct {
	Owner  common.Address
	Header []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyHeaderAdd is a free log retrieval operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) FilterNotifyHeaderAdd(opts *bind.FilterOpts) (*StoreNotifyHeaderAddIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyHeaderAdd")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyHeaderAddIterator{contract: _Store.contract, event: "NotifyHeaderAdd", logs: logs, sub: sub}, nil
}

// WatchNotifyHeaderAdd is a free log subscription operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) WatchNotifyHeaderAdd(opts *bind.WatchOpts, sink chan<- *StoreNotifyHeaderAdd) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyHeaderAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyHeaderAdd)
				if err := _Store.contract.UnpackLog(event, "NotifyHeaderAdd", log); err != nil {
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

// ParseNotifyHeaderAdd is a log parse operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) ParseNotifyHeaderAdd(log types.Log) (*StoreNotifyHeaderAdd, error) {
	event := new(StoreNotifyHeaderAdd)
	if err := _Store.contract.UnpackLog(event, "NotifyHeaderAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyHeaderTransferIterator is returned from FilterNotifyHeaderTransfer and is used to iterate over the raw logs and unpacked data for NotifyHeaderTransfer events raised by the Store contract.
type StoreNotifyHeaderTransferIterator struct {
	Event *StoreNotifyHeaderTransfer // Event containing the contract specifics and raw log

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
func (it *StoreNotifyHeaderTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyHeaderTransfer)
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
		it.Event = new(StoreNotifyHeaderTransfer)
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
func (it *StoreNotifyHeaderTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyHeaderTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyHeaderTransfer represents a NotifyHeaderTransfer event raised by the Store contract.
type StoreNotifyHeaderTransfer struct {
	From   common.Address
	To     common.Address
	Header []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyHeaderTransfer is a free log retrieval operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) FilterNotifyHeaderTransfer(opts *bind.FilterOpts) (*StoreNotifyHeaderTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyHeaderTransfer")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyHeaderTransferIterator{contract: _Store.contract, event: "NotifyHeaderTransfer", logs: logs, sub: sub}, nil
}

// WatchNotifyHeaderTransfer is a free log subscription operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) WatchNotifyHeaderTransfer(opts *bind.WatchOpts, sink chan<- *StoreNotifyHeaderTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyHeaderTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyHeaderTransfer)
				if err := _Store.contract.UnpackLog(event, "NotifyHeaderTransfer", log); err != nil {
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

// ParseNotifyHeaderTransfer is a log parse operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) ParseNotifyHeaderTransfer(log types.Log) (*StoreNotifyHeaderTransfer, error) {
	event := new(StoreNotifyHeaderTransfer)
	if err := _Store.contract.UnpackLog(event, "NotifyHeaderTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoAddIterator is returned from FilterNotifyNameInfoAdd and is used to iterate over the raw logs and unpacked data for NotifyNameInfoAdd events raised by the Store contract.
type StoreNotifyNameInfoAddIterator struct {
	Event *StoreNotifyNameInfoAdd // Event containing the contract specifics and raw log

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
func (it *StoreNotifyNameInfoAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoAdd)
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
		it.Event = new(StoreNotifyNameInfoAdd)
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
func (it *StoreNotifyNameInfoAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoAdd represents a NotifyNameInfoAdd event raised by the Store contract.
type StoreNotifyNameInfoAdd struct {
	Owner common.Address
	Url   []byte
	Newer NameInfo
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoAdd is a free log retrieval operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) FilterNotifyNameInfoAdd(opts *bind.FilterOpts) (*StoreNotifyNameInfoAddIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoAdd")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoAddIterator{contract: _Store.contract, event: "NotifyNameInfoAdd", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoAdd is a free log subscription operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) WatchNotifyNameInfoAdd(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoAdd) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoAdd)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoAdd", log); err != nil {
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

// ParseNotifyNameInfoAdd is a log parse operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) ParseNotifyNameInfoAdd(log types.Log) (*StoreNotifyNameInfoAdd, error) {
	event := new(StoreNotifyNameInfoAdd)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoChangeIterator is returned from FilterNotifyNameInfoChange and is used to iterate over the raw logs and unpacked data for NotifyNameInfoChange events raised by the Store contract.
type StoreNotifyNameInfoChangeIterator struct {
	Event *StoreNotifyNameInfoChange // Event containing the contract specifics and raw log

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
func (it *StoreNotifyNameInfoChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoChange)
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
		it.Event = new(StoreNotifyNameInfoChange)
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
func (it *StoreNotifyNameInfoChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoChange represents a NotifyNameInfoChange event raised by the Store contract.
type StoreNotifyNameInfoChange struct {
	Owner common.Address
	Url   []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoChange is a free log retrieval operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) FilterNotifyNameInfoChange(opts *bind.FilterOpts) (*StoreNotifyNameInfoChangeIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoChange")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoChangeIterator{contract: _Store.contract, event: "NotifyNameInfoChange", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoChange is a free log subscription operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) WatchNotifyNameInfoChange(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoChange) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoChange)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoChange", log); err != nil {
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

// ParseNotifyNameInfoChange is a log parse operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) ParseNotifyNameInfoChange(log types.Log) (*StoreNotifyNameInfoChange, error) {
	event := new(StoreNotifyNameInfoChange)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoTransferIterator is returned from FilterNotifyNameInfoTransfer and is used to iterate over the raw logs and unpacked data for NotifyNameInfoTransfer events raised by the Store contract.
type StoreNotifyNameInfoTransferIterator struct {
	Event *StoreNotifyNameInfoTransfer // Event containing the contract specifics and raw log

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
func (it *StoreNotifyNameInfoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoTransfer)
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
		it.Event = new(StoreNotifyNameInfoTransfer)
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
func (it *StoreNotifyNameInfoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoTransfer represents a NotifyNameInfoTransfer event raised by the Store contract.
type StoreNotifyNameInfoTransfer struct {
	From common.Address
	To   common.Address
	Url  []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoTransfer is a free log retrieval operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) FilterNotifyNameInfoTransfer(opts *bind.FilterOpts) (*StoreNotifyNameInfoTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoTransfer")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoTransferIterator{contract: _Store.contract, event: "NotifyNameInfoTransfer", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoTransfer is a free log subscription operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) WatchNotifyNameInfoTransfer(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoTransfer)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoTransfer", log); err != nil {
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

// ParseNotifyNameInfoTransfer is a log parse operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) ParseNotifyNameInfoTransfer(log types.Log) (*StoreNotifyNameInfoTransfer, error) {
	event := new(StoreNotifyNameInfoTransfer)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoTransfer", log); err != nil {
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
	NodeAddr    []byte
	Volume      uint64
	ServiceTime uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterNodeEvent is a free log retrieval operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) FilterRegisterNodeEvent(opts *bind.FilterOpts) (*StoreRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreRegisterNodeEventIterator{contract: _Store.contract, event: "RegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchRegisterNodeEvent is a free log subscription operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
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

// ParseRegisterNodeEvent is a log parse operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
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
