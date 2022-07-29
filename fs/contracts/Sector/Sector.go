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

// SectorConfig is an auto generated low-level Go binding around an user-defined struct.
type SectorConfig struct {
	SECTORFILEINFOGROUPMAXLEN uint64
}

// SectorInfo is an auto generated low-level Go binding around an user-defined struct.
type SectorInfo struct {
	NodeAddr         common.Address
	SectorID         uint64
	Size             uint64
	Used             uint64
	ProveLevel       uint8
	FirstProveHeight *big.Int
	NextProveHeight  *big.Int
	TotalBlockNum    uint64
	FileNum          uint64
	GroupNum         uint64
	IsPlots          bool
	FileList         [][]byte
}

// SectorRef is an auto generated low-level Go binding around an user-defined struct.
type SectorRef struct {
	NodeAddr common.Address
	SectorId uint64
}

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"AddFileToSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"AddSectorRefForFileInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"CreateSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"DeleteFileFromSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"DeleteSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"GetSectorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetSectorsForNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sector\",\"type\":\"tuple\"}],\"name\":\"UpdateSectorInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x29d6276d\",\"type\":\"bytes32\"}],\"name\":\"c_0x29d6276d\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_FILE_INFO_GROUP_MAX_LEN\",\"type\":\"uint64\"}],\"internalType\":\"structSectorConfig\",\"name\":\"sectorConfig\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061872180620000226000396000f3fe6080604052600436106100b05760003560e01c8063931bb19a11610069578063ba9210041161004e578063ba921004146101c1578063dcf74960146101ea578063e3168f9e14610206576100b0565b8063931bb19a1461017c578063955f98b7146101a5576100b0565b8063112346c21161009a578063112346c2146100fa5780632384a6aa146101235780632ba010d71461013f576100b0565b806247c003146100b55780630daf9945146100d1575b600080fd5b6100cf60048036038101906100ca91906176f6565b610243565b005b3480156100dd57600080fd5b506100f860048036038101906100f39190617627565b610708565b005b34801561010657600080fd5b50610121600480360381019061011c9190617650565b61070b565b005b61013d600480360381019061013891906176b5565b61093b565b005b34801561014b57600080fd5b5061016660048036038101906101619190617762565b610f5c565b6040516101739190617f1b565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e9190617762565b6117a7565b005b6101bf60048036038101906101ba91906176f6565b611a79565b005b3480156101cd57600080fd5b506101e860048036038101906101e391906176b5565b611f7f565b005b61020460048036038101906101ff91906176b5565b612ae8565b005b34801561021257600080fd5b5061022d600480360381019061022891906175d5565b612d0b565b60405161023a9190617d41565b60405180910390f35b61026f7f1919bca4ebae8d5fd402a1f6ae9c2bbff240b0d0dba56b72b3dc0cfb493de49e60001b610708565b61029b7f5c530f06ed17aa39adf8ec8643cd28552d30842a3a132d8e3f263cfe6885851f60001b610708565b6102c77f785030e7927d3075bff2f60fb3ac8c8b1478021c4a7c8a49cd47f04c3728506960001b610708565b60006102e083600001518460200151846000015161313f565b905061030e7f163ba845eb2f7dfbb7e7c65ccf6e321214ec289e87dac56493b7d3cc32b7e84860001b610708565b82610100018051809190610321906182d9565b67ffffffffffffffff1667ffffffffffffffff16815250506103657f37e0db7561b5458e105ab4ed2c2a9866c85f21e6631e5dbb9e4bc81d6509f46d60001b610708565b6103917ffa4eb79160d83b3505f88b9ce64963415ef9e9dada6c7204808fd9a2b7bf9a8160001b610708565b81608001518360e0018181516103a7919061817d565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506103f07f25c69ea2f875822b905948192875227b86814c1e8a60f4dc909a0e99f156894a60001b610708565b61041c7f1023127b5b2bc1ae6d51179d06c47d271b4779fbf7594d186ff2cddce5cf6dcf60001b610708565b8160a0015182608001516104309190618107565b83606001818151610441919061817d565b91509067ffffffffffffffff16908167ffffffffffffffff168152505061048a7f126ce44f4b0c5439c83b1bf16c0aa376bf9784bf220d8a1a3d349f17a5949ef960001b610708565b6104b67f7d5121356f6aee93ed8f0b567329599bd801abbfb865a8987c42cce31732e99d60001b610708565b8015610544576104e87f69a6ff89bc3c5f3292d244b8879ecf9d7674e5e40d45c68afef0eccc3e956e3760001b610708565b6105147f4e610f3eb313b2153cd3295876818cc1b8dc781ab65012e1133e26ac589fe98360001b610708565b82610120018051809190610527906182d9565b67ffffffffffffffff1667ffffffffffffffff1681525050610571565b6105707f590cb79bbc7705eea4334205859a24ad9c2ac5239e9948a184ed2f44dc065ed660001b610708565b5b61059d7f510d7abdee056cc770e4fc07e7f9ddb9721c1b0f8e5b107cf0a7af5518b3057860001b610708565b6105c97f9c6359425d4f44dfb7a47c51b58f9c955ed250acff1cf0f04998b1d2c752c51e60001b610708565b600083610100015167ffffffffffffffff1614156106755761060d7f57354aa8cc124e150e6ecb2785d8fb73ebcb8db649cd9d493e5618de0342d14660001b610708565b6106397f35b807f7272db07e484b6e496535872b710653c404a101b90cbe0414779d91c060001b610708565b6106657fbde6e7a054afcc21cbd32ac62579a5d09ba88f9c87218c98fb21b131a37391b260001b610708565b60008360c00181815250506106a2565b6106a17fe85520bb40c8d01fc80dd4ecce392376f17cd229b860b379f4993f89ae82a62660001b610708565b5b6106ce7f51e57f662c2b25ddcff64e0fd9f4df037bb07b901d8572385b40b64205a1bf4c60001b610708565b6106fa7f7bdc3e8e9982178f5ac88c637c0f7fe1a25b0dcfda1275e6a241875bfa8f6c7d60001b610708565b6107038361093b565b505050565b50565b600060019054906101000a900460ff166107335760008054906101000a900460ff161561073c565b61073b6138d5565b5b61077b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077290617ebb565b60405180910390fd5b60008060019054906101000a900460ff1615905080156107cb576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6107f77f94df7151c99d18fe5b39bf5694e4f0ad881a47be2108ac09c94e18fe704ca64b60001b610708565b6108237f68b2f4f8496436daa2a159137cff788170e1438c557a00fe5dd4dde2ee3b369560001b610708565b61084f7f063e2c51dd767ad26c998da4faf5bac68436ff4dd29e431849457614e0f3342060001b610708565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506108bc7fd805d5d9c79338b1f7c74b301df5c3baee97b86179bb4298fc6575bf1a810cdc60001b610708565b6108e87fe6856d2e9024c9ce4ab20b7e0d2cc8de7c990727f43d56241c1e57553813616760001b610708565b8160000151600060166101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080156109365760008060016101000a81548160ff0219169083151502179055505b505050565b6109677fd297d909672a9a2727fd26d0d5e4f3c3dd17c6c4f54f91735379e003ad8b3faa60001b610708565b6109937f564f4d2fbcbba648a76ab9bf1dd0ecf7f1eecc23d035051d9b6d30b875a74b3f60001b610708565b6109bf7f87e045822cada65242b4e2b5e422d327218f143bd816dd664168876aea10ab7a60001b610708565b600060016000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050610a327f1ec55d904c2c8f3edc0aa0a97fe99f333ccf42d3413f9b4dc2a5325a857cffc660001b610708565b610a5e7f3082e6e4a4632d143607cc14f9e4884f4a26ec934e9dd90f6790a4d0ed0c00cc60001b610708565b60005b81805490508167ffffffffffffffff161015610eac57610aa37f4d0edc8169ef7845c7bd26419933e9b4759f413cf4a03be6111c525e7609dc9060001b610708565b610acf7f6bfab9f8c963d0f1deca1b2be6afc4e750205731aba454d10151d607025bf02d60001b610708565b826020015167ffffffffffffffff16828267ffffffffffffffff1681548110610b21577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff161415610e6d57610b837fe543d18fc27955debe4e43583f6a3bdc06a0ce392d5557d92ff6b303cad9196e60001b610708565b610baf7f06ab8e1272587f8b6010c111daf764fe07c34416359e2c61bcef909d7fab9d4a60001b610708565b610bdb7ffc4392297080a5def23f259dc95e7b8255ca66bb91f09cff1cfae560c10d3afe60001b610708565b82828267ffffffffffffffff1681548110610c1f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115610d52577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190610e389291906165a7565b50905050610e687f391806f95b44a93ce4c209bda3373f2810879f34e3a96529966d5fa3795247c160001b610708565b610eac565b610e997fdc7b558ffb8a487d1445ceba4ad66ed73d1215feb82bd0e2c23652530951cb5660001b610708565b8080610ea4906183af565b915050610a61565b50610ed97fdb1172d0b4e11c82c81ea5ef955316cacdc964839ef3e57dc4185a6cbbbca34a60001b610708565b610f057fa923f46bb383dc7a698d1f66066a2ad4f37a315f6889e72c7d902ca87417ea2760001b610708565b8060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908054610f57929190616607565b505050565b610f64616913565b610f907f641b38a3521f55ae4d729f204f83cdd7b93fddce5a6059fe563aab7256fa93f260001b610708565b610fbc7f68263802c04f466a4e4900c266a39b44b30b16d12a634347c22538d93a35892d60001b610708565b610fe87f7669b90bc12843f4542330e89e1c359e518e50d4b67893314b78a7ec96d5f3d260001b610708565b6110147f1491df281584694724aabe378947fb4ea7ec0e748cccc3f846014579ddea54d960001b610708565b6000826020015167ffffffffffffffff1611611065576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161105c90617efb565b60405180910390fd5b6110917fd6c2a49f8cfde3699f02b96ea4e2b88c782adbdb3f73e007e0968f2e4151294060001b610708565b6110bd7f54abf930c4a14b608868747693c8f98b597749f2d18abe0d8c59b93eb4cb9f8a60001b610708565b6110e97f05e15157583fd1ec87049aff436287f15b0dff1fa21f029ca05822aae071650960001b610708565b600060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156114925783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff1660028111156112a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156112e1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b8282101561147b5783829060005260206000200180546113ee90618303565b80601f016020809104026020016040519081016040528092919081815260200182805461141a90618303565b80156114675780601f1061143c57610100808354040283529160200191611467565b820191906000526020600020905b81548152906001019060200180831161144a57829003601f168201915b5050505050815260200190600101906113cf565b50505050815250508152602001906001019061114e565b5050505090506114c47fc722cd062822a2b731951b4558ee226877f4b1afde336927c058c3fd25d5fb5d60001b610708565b6114f07f5f0c594302665b781cead17a2cb04cebe9a3a0e337d94cc7c58233a40d8eb75960001b610708565b60005b81518167ffffffffffffffff1610156116e3576115327f3a61e4a89832ccfff3af94b6949a2a19c80d8cdc2de466e70c0f9b6d508090e060001b610708565b61155e7f47cda66b16cc71c048a497ce8e7eb898dcf46c9d3dd262a835e3e4a3374d2f9160001b610708565b836020015167ffffffffffffffff16828267ffffffffffffffff16815181106115b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff1614156116a4576115f87f1748b449385775d8e4abd7fdca18cab2dc05f431b162a538047dea696783df5260001b610708565b6116247fed6a8808fa7b14b7cd9021db7a4d5329c2fb1c2bf7f617f17e2fbef78e38b85260001b610708565b6116507fad3c80c8eacbd7ec60b0632d1ddf0bfc09ff8cad25b657d43e91f86a91ac32c960001b610708565b818167ffffffffffffffff1681518110611693577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151925050506117a2565b6116d07f334b3bb33c061d41bf67c5878563475f66c69013b84de5ad739fbb23208b05de60001b610708565b80806116db906183af565b9150506114f3565b506117107f494336d6fd19594181463b84a8009d2dac56369fad8e482712a72640f15156e560001b610708565b61173c7f4cf8b20b7a43ef77f3b2ae8740abc6442cd1cbafc463c9527c233c9e68ac14c860001b610708565b611744616913565b6117707fee7242e2b7a06f41360e9e1fb6ae232b14979876423f5f5d46cf00eb0df9a32860001b610708565b61179c7f6ef08a460089223497415e6a1caab8f72684857192188ff3b8dbaadcfd3b8cb860001b610708565b80925050505b919050565b6117d37fec361340e10c6c9034b819f881aa8b2e4343cce346bf143fcd4fdd977fa33de460001b610708565b6117ff7f08880e75794ed1fef6e2237346cbb6995ea581d700b5ae314ac6a2a04472aa6b60001b610708565b61182b7f53bb85a76895147542fd05807464cadceeb1aca73f33b55b7b75e9966527b9b860001b610708565b600061183682610f5c565b90506118647fee8119464a84a9358ac363019bc804d77ea4f08a9a364d7612af3989962b40e860001b610708565b6118907f031cfaa4f196a5b60b52a52ec244e9fb1d75b0b9ee0ff8e5df1d16346b058bb460001b610708565b600081610100015167ffffffffffffffff161115611945576118d47f6d97e234ff39343dc396abb7afe4ff323b74876404e32158d4b75997c14a95a060001b610708565b6119007f284104df9a0f536819d8befd8f35f6e24d2f96dae0e650162e1a3958a0a4a70160001b610708565b60008161010001516040517f25a56d9300000000000000000000000000000000000000000000000000000000815260040161193c929190617e17565b60405180910390fd5b6119717fb82375164b8052ac667c07dc5ac0516385b0d7064609a62dae625af03697d37d60001b610708565b61199d7fd2eae4cc0ad3ac0e9ecf5495eea13c28b8cc2c134d3c6f6d484ed6e091968bc460001b610708565b6119c97f946665e5e4029752557f5b15847ac1b1f458deeed1909a25cb4cad8b3819a6db60001b610708565b6119db826000015183602001516138e6565b611a077f391cafdd153706d2543201367a276d67d6a49dfb1c65774e028a650520706e3860001b610708565b611a337f7cbf4f3dc2995bd887dbc0a05a4674f0435b95973e7e1b2119ab1627ddadce9e60001b610708565b7f4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c600943338560200151604051611a6d9493929190617d63565b60405180910390a15050565b611aa57ffa06ebdb614d17dd832e394d3f9bcde75d265f6fddac1f70592fa66eee7670a360001b610708565b611ad17fdb2fa5ffecf75a543e389a088c9943bbff6b5bec4d8531bae3836fc9111a29ae60001b610708565b611afd7fb948d0ba24a2c4df0b73a05f009dcb456711dd8b0cdf30c8b079e696b529a05a60001b610708565b816040015167ffffffffffffffff168160a001518260800151611b209190618107565b8360600151611b2f91906180c9565b67ffffffffffffffff161115611bc957611b6b7f74e608459a2112da89f19d0e46abcb5696f78c5e04cd865e3d9317d35ce6f19660001b610708565b611b977f1f20f6324d272416600a1579c08ccb2f9f97cf24624ec1c44549b1a0c5fa0e5c60001b610708565b6040517f6073072a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bf57fa213c17241a72d319fdf1b620c1d4f2735eeb91e29ab71424e940949596ffbab60001b610708565b611c217ff37f7ae7ea5cb13b63be8efbe4d38fd352a0c272ae7e432a60190e081d1999fe60001b610708565b611c4d7fd504763f93a854e3b3acf1dd3d439230ad1ac28cf503bd463b4eda788e6c100c60001b610708565b6000611c8883600001518460200151604051806040016040528086600001518152602001866080015167ffffffffffffffff168152506139da565b9050611cb67f3c2045dd8d906c971b8ea4edef96a4e24cffd6a6ccdb2c7ff325d29abacc4c7360001b610708565b82610100018051809190611cc9906183af565b67ffffffffffffffff1667ffffffffffffffff1681525050611d0d7f6e1eddf16f77ad9be7671617f6abeadca254d36fb7c58d05cefc70b6e12ffa8a60001b610708565b611d397f3acf5e0503bc9445a44d0b183b710dbaa050ab2b23e28d75d60350f59b0f791d60001b610708565b8160a001518260800151611d4d9190618107565b83606001818151611d5e91906180c9565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611da77f3c0b7578c7c28ab28c47f38439fb264a96dcbf9e471be307853638c0faf3df4f60001b610708565b611dd37f63c8ff23704830631276afe0c5b24d7d017b4042a51cdce664b5c848217b54cd60001b610708565b81608001518360e001818151611de991906180c9565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611e327f123e697b7637dee140e3961f0e89497c7c57dbb56d1e4b8b38985b30821ebc1d60001b610708565b611e5e7f4f7bb092bee5a4d39f365dd8d633e97836359ed473adf42f8b1b980682d6fb1960001b610708565b8015611eec57611e907f0356932f04ba2cf659aef767288e5c5e131836d4b85345269ff126cacf9e73b060001b610708565b611ebc7f4234215dd452c4b20a3f7bf9728760e0d3c1aff73dd6e49d807b259da17b846460001b610708565b82610120018051809190611ecf906183af565b67ffffffffffffffff1667ffffffffffffffff1681525050611f19565b611f187f58d1f5dccdfa035d70e8c62893f9eb7e036852283bd28fe7f9085b391aa8a49d60001b610708565b5b611f457fd58f810a3790d58ad24805e15b7d2abb6ee39a270c469722815ec7184ac61f1960001b610708565b611f717f40b3b656367a9ab079367839ce9ee1a6bd37640adf55d47507b713e5abde47c560001b610708565b611f7a8361093b565b505050565b611fab7fdd4918620906bd18a8dbfe593f0797ec5e831bcebcf3b4c86837b25c6afaacfd60001b610708565b611fd77f5383785f617cf8d11dd39a7dac9954debf1b2aa7f86dac9becc30170b1602c1460001b610708565b6120037f744d51c97dc8800a1702c00f87f7494921f41760da4d40985c4fe8fea307ca2760001b610708565b61202f7fc2a5542a87a66452b910a83aab05f8bd5ee487976eec78fb39a57429aa3725f560001b610708565b6000816020015167ffffffffffffffff1611612080576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161207790617edb565b60405180910390fd5b6120ac7fe5b41227978c3572a6a9f296278853475280aa88a8339ebaa518e6068e30f21260001b610708565b6120d87fafd911f07c818c562218e665a20915d820b2d11e50e565e50fc788660597aabd60001b610708565b6121047fa7a121658f83864466ade33c52b27a2c36f77a116d8c080cc69ce09346cd087b60001b610708565b6121307f8b5d8ef462b2f6225fc628c7e72b4b692eed42933c8c5cec84eb1b6a977452d260001b610708565b6000816040015167ffffffffffffffff1611612181576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161217890617e9b565b60405180910390fd5b6121ad7f14d90f0df29effba0b61ce90648846191f5465d3435d80b1e5e33ed63b8c65d360001b610708565b6121d97f3ac600e0bc78016e9d25379c795976ceb96aaf263635745b1d197e803289c60760001b610708565b6122057f74850e734a3555ee34b7f57a8685f07576ae619985c41db65e0135e1c0a19f4860001b610708565b6122317f3fd1f3511d0e6a17f1bfea3090645fa6fa03ea15e33401b9e674b17b455181a660001b610708565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636683899482600001516040518263ffffffff1660e01b81526004016122909190617d26565b60206040518083038186803b1580156122a857600080fd5b505afa1580156122bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122e091906175fe565b61231f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161231690617e7b565b60405180910390fd5b61234b7f6edb7ce36c4344d8aeb9e017379198064314f879006c5b930deed32091010ba660001b610708565b6123777fe757ccad20c1e422f629ffefa7106594379a2e61ea125c5f8e5b93bc1a47ad5c60001b610708565b6123a37fb08f9e0952d9cfa016ffdf7dfac04bca3af5cff72f4b84722daa3442ccac7eaf60001b610708565b6123cf7f6f165d9c6a18a5403762059fb1dd37e9d75d3ddd5cc6af6a0550041861cedec960001b610708565b60006124166040518060400160405280846000015173ffffffffffffffffffffffffffffffffffffffff168152602001846020015167ffffffffffffffff16815250610f5c565b6040015167ffffffffffffffff1614612464576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161245b90617e5b565b60405180910390fd5b6124907f753a2e9c35ecee68d571fd78eda15867785909ef62082a0ab4f881feb4d2a66160001b610708565b6124bc7fe11e3712a8038ffec887e81a3deb04e2ca6ed833eea03a3cd82d39cf0eccc66660001b610708565b6124e87ffe92da217a2486ca2ca9c265debe8c1bc75919f8fad6dd39ea523212e66ba35e60001b610708565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe83600001516040518263ffffffff1660e01b81526004016125489190617d26565b60e06040518083038186803b15801561256057600080fd5b505afa158015612574573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612598919061768c565b90506125c67fd526401d864e7e54b788f897d7f801679394a124889c0cb36ad077bdc6346ea360001b610708565b6125f27f86901653a79a1e07a93e7ead9f51ea0a7a97cc8a4c520baa7ec1edfc4cc10d7160001b610708565b60006126018360000151614335565b905061262f7ff1d35bcd3706f6ca940f0a7a632c876354eada821806b17da86a1c2abcd7fa8460001b610708565b61265b7fa299207df252ab374781bd84d6ba4e00881e1bb3a45f0faac8054115a17ec9d860001b610708565b816040015167ffffffffffffffff1681846040015161267a91906180c9565b67ffffffffffffffff161115612734576126b67fb770a05813e7e9e5a7c5838c65b2b003e06afbab666e99b9d24d01418c1acf0c60001b610708565b6126e27fdcfa3cb1c9de24582a87b8a99ecc16fce9fc119bce7289e35d1ededf7349913660001b610708565b81604001518184604001516126f791906180c9565b6040517f95c91e8e00000000000000000000000000000000000000000000000000000000815260040161272b929190617f3d565b60405180910390fd5b6127607f19ead433b6820abf04255ee9b682cdb7786e268c4427c566ae30834bde74c20f60001b610708565b61278c7f5d884843e62e6bcf82e0956a44e1ac3dfeae98a77697d8a6f4cfbeeff4bc954860001b610708565b6127b87f1a06d2bec1f25293687c16837ddaf9186187e8fd2527b67a4b6789ad763625ab60001b610708565b60016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff0219169083600281111561294a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190612a309291906165a7565b505050612a5f7f69118391ee3d378b0157cf479dfee3ff506016155ee16cb194b0c840fc7d18c260001b610708565b612a8a7e70985d4b5a69c9d7ac27fe5ab8d35320213b6556c4ead45c2c2a11835dbcee60001b610708565b7fc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb6008438560000151866020015187608001518860400151896101400151604051612adb9796959493929190617da8565b60405180910390a1505050565b612b147f2fac9d52d6df0b2e3e2540e42da31ed1c072c3061f898d12e20e5498b629da3f60001b610708565b612b407ff8d197e7fbe12293adf37db3947c2989e18cc274dda25ea5175389856f40017060001b610708565b612b6c7f60bbc05049ddf9506ddcd585fa363cfe054a1a900ae6c57cf25c584885dd3d1960001b610708565b6000612b8082600001518360200151614944565b9050612bae7f7667e547b2b162dff3df9e7a9738359bc3b9381ac76075ea7e6ef80c799445e060001b610708565b612bda7f23bea6f3a3f48eb661bdfbc870a8fdf29b78cbf95fbb587116d6425c76d8361760001b610708565b80612c7557612c0b7f77352e890805fb5c32b8bc3b4c2973f2381eef9a6c931d9cc3e8ee01dcc4e37e60001b610708565b612c377f5c0725ffef4664bb84fa03e1629b447c3260ffd26eca0b3216b4221470915dba60001b610708565b60036040517feca42856000000000000000000000000000000000000000000000000000000008152600401612c6c9190617e40565b60405180910390fd5b612ca17fb417300de2ef4924aee66889fc6db4c9db983b7e410d8d8e2ec854fbb4fa908760001b610708565b612ccd7f73f4341f5794095b7aa9d0bd297a0712aca86a19c0385f57e8b3d77ca5c75c2360001b610708565b612cf97f2284266bbfcc7589b68f9983ec060c7aa2c444750fdb988dfc593af175b5407b60001b610708565b612d07826000015183614ca1565b5050565b6060612d397f69e3afeccb4d559459dcc3996200dab360b2fdb3c8e8aa49dc40d669c8bead5860001b610708565b612d657f058c43fb00c9a64896c3448beee16e193d83fb55bb57c5181b2a783daec1bfbf60001b610708565b612d917f52eb86ea507b50e0c43458431b293ce378b9e4a7a51258781c4a73a0176b703f60001b610708565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156131345783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff166002811115612f4b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115612f83577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b8282101561311d57838290600052602060002001805461309090618303565b80601f01602080910402602001604051908101604052809291908181526020018280546130bc90618303565b80156131095780601f106130de57610100808354040283529160200191613109565b820191906000526020600020905b8154815290600101906020018083116130ec57829003601f168201915b505050505081526020019060010190613071565b505050508152505081526020019060010190612df0565b505050509050919050565b600061316d7fe870bd9bcdf60905126a5b1aca7a95e6979ff877b52a09959208cbac049c2d7060001b610708565b6131997f159e77117ba1fd850073d314e44cbe1511187c826ca50a34718921242064a4ba60001b610708565b6131c57fd71b040656e0d03fd531efa91d9ec2194b733eeee4a967ebec5e1402f0fa1ac860001b610708565b60006131d18585614fa0565b90506131ff7f27b08b775e100644752b6bac2834789025a45d8dc7cda45a1f5067b5ec67ebe660001b610708565b61322b7f0a570e804cce371cb0c1c7005384fb04582247c542440a0ea300bc40bb4712b060001b610708565b60005b8167ffffffffffffffff168167ffffffffffffffff16101561386f576132767f57477047b774705b86767f676777402006e63f4c41d022ffe77ce8e0b105426f60001b610708565b6132a27fd8559a89e9658ab23780349ed71d35da8aab72693a59653dc021e30c87267fbf60001b610708565b60006132af8787846150ce565b90506132dd7f8cb7ec3857cbde6cd22a72796edcfd63ba861ae4280d93d0e874a63d5ba056b060001b610708565b6133097fb17f94f80ca9b36f3beb705671136ca610251eda85a3137386acb2bb93782a4a60001b610708565b6000806133168388615363565b915091506133467f39e556b846dfd7f0036336855ba5ca779c408d748fb74c28c15029950f06802960001b610708565b6133727f219b75189bb77a39d3eb650d6038b16c8e1d071edb988f102d513e216971a3ca60001b610708565b801561382d576133a47f20b113ea6f803ebc8d476356cb1f047c14090da4ad2e1ed6c79188dc0bd4d68560001b610708565b6133d07f546a10e7f911a1436b6a2e2de15bcccd88aa901e7d545fd3f735f1b9f4f68b3a60001b610708565b6133fc7f6f186ae4e621f201e15c5130a515fae5ebe40545273d85c9c2f5a5880024722860001b610708565b600060036000856000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002080549050905061345b7fbb75de0725d74fbc2d5ed57e863a85d22ba4bb66f045586a30e98bcd5b1a378d60001b610708565b60036000856000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208367ffffffffffffffff16815481106134c8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600080820160006134e79190616a00565b6001820160006101000a81549067ffffffffffffffff021916905550506135307fe1f07eff54869d7a56d8518209aec8aa2bf8ae9d2c3c6c4744f5605469945f6660001b610708565b61355c7ffb55f47076886b68be2e3b441c33067ed0ce9b0d6ee846e8fa0ab83754dcb5ea60001b610708565b60008367ffffffffffffffff1614156136015761359b7f0a846942129d18645466e17d9b14817e1fc0529e2c2ac59568aa11f0f298016c60001b610708565b6135c77fa3401bedcebe2762e6af8d3c9712569a8e7ef62aa67913b454c36fd55270c15260001b610708565b6135f37f5a2d14051cf50816196ec869838694f41b9e669a75a28b0551385c309b14e26160001b610708565b87846020018190525061362e565b61362d7fdfcb353d1b1a466bc4e79e002388e01138b43d56d54a421f8f9c7524b0fc39ad60001b610708565b5b61365a7fb8b97a6aa7c7dae79841ffc4b1ca1c99dc15b4cdbd07051eaaf379794841524860001b610708565b6136867f7427c26ad991217d85f3ebe2a1cfcbf89f7ad4bc1ab67929d3f7b08bd6b690b760001b610708565b6001816136939190618149565b8367ffffffffffffffff161415613736576136d07fdc5bf24c1855956212c6ea41a95a5a640befe5a1a52c5d8060fbbdcd9a83209960001b610708565b6136fc7f0fcd8756f3a59152d7c91f0926f059232a986d4d86f201ece26c0c1733acc54660001b610708565b6137287f95658c079e66fc3f583a64ccb1e516053fd614bf6426985132a9f008907bf5a260001b610708565b878460400181905250613763565b6137627f1ac1572b7db2079a161a4ed05ba839e87da7c20cf5556e9dba283100463273e960001b610708565b5b61378f7f54989847dc03b6eb238a2781af84179b47cad6c2e54848b7e34f1e1bd4eedc3560001b610708565b6137bb7f16bfe8417b31e895188a60169fe8fe28477f0561a683491ed488d7f42e02f40d60001b610708565b6137c68a8a866159e3565b6137f27fab196df1a7861911bc73eb5275e2b2b8838d61f4d7d192bb483a305b88e7822960001b610708565b61381e7f942a7d88cdb05ecb711daa4a8d3afb748713a118023f983b6034120c4997474660001b610708565b600196505050505050506138ce565b6138597fb98a990a78fd06b2156947baf4b25c5aefe9ec35bd563410f5ff441472ac68e560001b610708565b5050508080613867906183af565b91505061322e565b5061389c7fb286e1a974abb858fd4bd6b0cbee2a95f935b1c471e95650cb46029c2e80577060001b610708565b6138c87fc375d4a9f01e6da1e676bf8e7cb938b14dfa5d185f7f82f44503e073ae31770460001b610708565b60009150505b9392505050565b60006138e030615b7c565b15905090565b6139127f698947cc5c4a7fd78e529bbcc993d4b5ea5d7e56caf8bf3af77c9eebcfe4871160001b610708565b61393e7f69cbfb5b05f7c7a0b63d16d4064c3e555ba5dce7377775c9dc9f2bdcc5dce77260001b610708565b61396a7f097d528a4a2b35a610e2e393301def8f44e17bc1eedcf05c752e5f913931b54a60001b610708565b6139748282615b8f565b6139a07f1c1cfd23abd0d5e09f7a72dc325a617c80e493afe136cdfce399439f0c7fc0fe60001b610708565b6139cc7f98386665762483197cdc068a908f4acd74e38f6ae1136c80ce04f050a91c190060001b610708565b6139d68282615d14565b5050565b6000613a087fd3ad23d4cda7cf17b4f7334a7fbf9bf35ce564c19cf5eda4c1c19ab49b79ab7060001b610708565b613a347f3e46baf3a1ae6f71ccd99daa20ca4ad7f96ad7e1a8c95b4b131ad62629234d3960001b610708565b613a607f3f65c84b03457e3ee1b8dc7b5969754f6a319e2179f7ab845c072fea305460de60001b610708565b613a68616a40565b613a947f970226277e5ad7ef90bf36ca69a65f599ff0b87476b2dfc4cba29e320a88813760001b610708565b613ac07f1f3483dd07c2df7235ce7ecd7e5b4205ee84768a6a4e43c80f68447c06695d7a60001b610708565b6000613aee7f6cda95a4c4f4d59f4e50e64b06724328d8affaf6e59f047abcc94c7209f0b0e560001b610708565b613b1a7f2088cb317fe5d85b09ecf939c847b73b95f343455a3409bda7acff0cc6dbf6b460001b610708565b6000613b487f32bd0f72df208fe527b9fb07b8e5ff5e5e189dbd14e6e93e8df0c466bc9ab9b960001b610708565b613b747f3bd744b4f0cb9f34cfa56aede6d36cb7ccd73d6b8c4303265c16e3d1b286ceb760001b610708565b6000613bb360405180604001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018967ffffffffffffffff16815250610f5c565b9050613be17fe0775fd8d8f2430bbf141665d8bdaaf4e2282ca8526e45620de11517655276ec60001b610708565b613c0d7f16af40ccdacb55322dca2dfc7537677836823e85a5a78b634e17859175ecbd0160001b610708565b6060613c3b7feacc2a0f3b1bfdb218dbb55930b9dc2227bad9a604dac8761439659ede4ee4e960001b610708565b613c677f84931eac100ad095d3515b0a80e8ec111c4b87c84bc0b39f16b8fe10880f0e7360001b610708565b8161012001519350613c9b7f012889bbce1ad80536f58d262d9fbd3197f08942d5d27e747c483d9134156dcb60001b610708565b613cc77fdecc667777e173a614110ddde1a5ea32de1bda362cf603b28b4e8f0b5f59bb6e60001b610708565b60008467ffffffffffffffff161415613e4257613d067fc12256b0530319e6b3fd6b261a0ee830fd4475a355b853d707ca37d8588d0bb960001b610708565b613d327fa7ee6454538e2edb5bce610ab1bbfb09c60d36b1968c7d0e1acf34c4b82a621260001b610708565b613d5e7f89fb0b169176bcf50d158e9fe0a442418fd29e68740803fdbf1df27a20a3a0e660001b610708565b60019350613d8e7f1ae08fc1ec360e382dcaa2dde8e822a19af148080ff09e6306d30f6f952e461b60001b610708565b613dba7fcc8eb0623751ec5f6e95f7bd3dd85551d6efdc1d3f556fdb3e1595827968b3b960001b610708565b60019250613dea7f8a15f61858ac420399bdb255083147d126ae45b7f29c167d2a50d7ea9f32779360001b610708565b613e167f149ba3cec578af3856f595695512ed2d7881924cb6aba748f0f2e31d2cf7544560001b610708565b60405180606001604052808567ffffffffffffffff16815260200182815260200182815250945061426a565b613e6e7f1d17502b1a8e7db7980e0ce9f81dbe31cf7232d861ed11e658227052a22e2bf660001b610708565b613e9a7f2ecbbe89e21fa4116c5b324a39baedc78ed286d500201e6f2e022a7b0fbd7d3260001b610708565b613ec67fb56014fc991d774b2cc23e58fd447758850c3929af6c8480dbc1af34d8d1b37860001b610708565b613ed18989866150ce565b9450613eff7ffc6ef0bd0a104f1c184c325d5dda5fce31c3516710e10c5bb71f63fc228e707260001b610708565b613f2b7fba0414a5c4e0ea6a14ff9265a7ebada39bebe8bcdfe9f186f1a77afddcb9b01960001b610708565b600060036000876000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561406e5783829060005260206000209060020201604051806040016040529081600082018054613fab90618303565b80601f0160208091040260200160405190810160405280929190818152602001828054613fd790618303565b80156140245780601f10613ff957610100808354040283529160200191614024565b820191906000526020600020905b81548152906001019060200180831161400757829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190613f78565b5050505090506140a07fc13739bf4dd9b9b36af6c777b4cf51326ffc2f7e9ad39b5a561b8a7ac2ad83f460001b610708565b6140cc7fcb1eff8dede3e0033c17523fe739b0930827bb116127905dda931ab0b3672fc760001b610708565b600060169054906101000a900467ffffffffffffffff1667ffffffffffffffff168151141561423b576141217fe0db4d86d4f13ea560949911caaa6d9449fc19a1d27d4f0842f9839a28c196b460001b610708565b61414d7f8afc9266de811521ab58cc1ee7be13c88cdcf47f7947611e1e043b34d0b2fe0c60001b610708565b8480614158906183af565b9550506141877f233a373e493c8bbff8c4b852481ec86d4eebac44dbce5395a9f0a685a42f51b560001b610708565b6141b37fed075eb98a1cdadf955c3e96d4be67fdc9854ce85b585c2854d3ff7d934e101760001b610708565b600193506141e37f9ab838b6a611c18b5ffa6dbe763f199bd2a3b50e93b5c133c815046d7e446b8e60001b610708565b61420f7f403ada0f1b9922170593c52135bdadf9682b0b26d54cca69aa9a228da1c8292560001b610708565b60405180606001604052808667ffffffffffffffff168152602001838152602001838152509550614268565b6142677f0b224ee8641c55ed20be4fcd221366546700502750a93494f38bc1ec3ed7d57d60001b610708565b5b505b6142967fa099831781b28930643d93eddec0e9545d2e564016aa4d9c90617c98f609c11c60001b610708565b6142c27f3516db01dc1927b78206ef086546ecd86cb7ddf44a2f683a27032aa74a100c4260001b610708565b6142ce8989878a616179565b6142fa7f18d71158e4bed03f6767a2d50d4c410e215171650fdb18f55106dd5dd8fd177860001b610708565b6143267fc347d3e28813d8ced8f06470153ba5f580c36fef40fec9a2b801bcb97488f90d60001b610708565b82955050505050509392505050565b60006143637f727776808cdf7c0386c4d04de7a487f1e4efb857d537025fc000f1ec0cccb17660001b610708565b61438f7f2db9093f38cce8130e63a2dafef628d24d93175aebe51402a4ede776b65f1ce760001b610708565b6143bb7f8003cc263ea6d1bdbfda8990266bdda47a2855ca9d49f514be287e160512d31660001b610708565b60006143e97f9136ce5995c45e12e9e53a96500f98aaa357943a20f748b76ec49eb548d9ca5d60001b610708565b6144157f22f4fe9575a1b65c9e62e79b27479315cbb2894edf37af541aa817b597a296c760001b610708565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156147ba5783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff1660028111156145d1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115614609577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156147a357838290600052602060002001805461471690618303565b80601f016020809104026020016040519081016040528092919081815260200182805461474290618303565b801561478f5780601f106147645761010080835404028352916020019161478f565b820191906000526020600020905b81548152906001019060200180831161477257829003601f168201915b5050505050815260200190600101906146f7565b505050508152505081526020019060010190614476565b5050505090506147ec7fc72ae74c6b6a3c935f6a132d1739d6b2d6d4ac2eb4c73e538fa733f09db33a5560001b610708565b6148187f01d6145fae03b849125293edcfc82f5767945620736f06cbca6798ef54425df760001b610708565b60005b81518110156148e1576148507ff8e1a7a75b5a33fa7d2dd11c520fa193f74aa73831165da87b5ca01c5d2a707260001b610708565b61487c7f130555d9d3343b7f209bd27ae80b3142b567b12623873f347a12c41f0cc5298660001b610708565b8181815181106148b5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160400151836148cc91906180c9565b925080806148d990618366565b91505061481b565b5061490e7f8b8338f32183a3f7f2dee2fb6b57c51e48a63c1df55921c91658bba19efd38b660001b610708565b61493a7fac2966ae12fbebe1273801315e2afe83b7c11c53ca14b047b3c91481390d860b60001b610708565b8192505050919050565b60006149727f769e816c7ed60337ba352d71a5687c83cfb1c6e57395bf8d2743bcfa6d51aac660001b610708565b61499e7fb76c9edc5325d2190f8b415bd94be79f3f567ee95cb84260e46b611e6716e7d760001b610708565b6149ca7fe952608570efbd273b78ef5bb2da16e4b6f32f453564fe98a0c8b09ff7e724cd60001b610708565b60006149d584612d0b565b9050614a037f6b30fed03acd5fc7882a50c740716fa1c5b3d2a1a221414a2ff1045f16f4de5260001b610708565b614a2f7f5346d600527ad7dfbf0c8cf4941ab1b02e05b852c8843aca86e9f11db9af635c60001b610708565b60005b8151811015614c3c57614a677fe423cb92a85981f3b7008291bf62a326050f24cca06633dbb404ddb566f5944c60001b610708565b614a937fba7a59354f295e9cfa29b0c61ba5b57ad255aefda16543c531b2e97b07b184d760001b610708565b8473ffffffffffffffffffffffffffffffffffffffff16828281518110614ae3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff16148015614b6957508367ffffffffffffffff16828281518110614b51577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff16145b15614bfd57614b9a7f76eba042b8cc59bbf45380851c350b1b02b972f5bd022be8d2d8ac0b2711663260001b610708565b614bc67f2bffca3adb168293b23513ff9b15cf7278722e6f871e7271d645d6ab79827f5760001b610708565b614bf27ffd15fbd409606a0e1426f1a4e35ef9918c7d64e1d15d42239b4c12500b92458b60001b610708565b600192505050614c9b565b614c297f45b11e98e533070af462a3c15704ff945089d4ea678a722397bdfe3ff50e95a060001b610708565b8080614c3490618366565b915050614a32565b50614c697f7d28183d6f64ec80db9b3e600be71e64adf9ee58cce72cc45960158ce6e3145160001b610708565b614c957fc345614eae689e5a50172d2a2e1b691d70be2c5b8fefb9136671dd325a11e5e560001b610708565b60009150505b92915050565b614ccd7fdf63988290c6710a1cd554785556e3d51f4fd7e74f82c3f06837558c1d14ea8260001b610708565b614cf97ff0abaf79dd5c09a8840c7b9d4c6bb1dc09ad34fee7de6abf83c923c3b118ba3260001b610708565b614d257f62c78c20e861d4240c4bed6274293cb7d78f4f16f1d1438f30771a23f0d42a8d60001b610708565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115614eb3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190614f999291906165a7565b5050505050565b6000614fce7fb01281155baa7b88feecd0930d5cf4c6266f9d522a85c0621d3644932a3b594760001b610708565b614ffa7facebfca90c258e8aa5f4ab833c955b598a3bf6b31d678e742583bb20672e1de260001b610708565b6150267faa0eb19ac85ea38d1f6fb181c8ed4de81b9e1b4f55af20003a34d2a59b790e9e60001b610708565b600061506560405180604001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff16815250610f5c565b90506150937fc6ef2bc50a885eadf8c0c29cb86cfb06d4213329934e7a435b22e7e82fbe639560001b610708565b6150bf7f4befa1a9931c63025d3c955158ba7b904b96329a414fc6feea57f16a043c1ae260001b610708565b80610120015191505092915050565b6150d6616a40565b6151027fb9a82c772c3aaafab1ae4bffc3ace84eb73eaa94136176969aa3cc4199432b1f60001b610708565b61512e7fee08fc3de5c29399704e7e6aaa93e2b60bd5cc0cbde613f349f5fbcbde822f9860001b610708565b61515a7f4b4b7119685992083cbcb14f763b316479271f791391d488b842a6b6f885d88460001b610708565b600084848460405160200161517193929190617cd2565b60405160208183030381529060405290506151ae7f07ef6845f56900d94c72efe775178c5026dec3d0c83e4f0fbca0f71c6d71a1af60001b610708565b6151da7fafeaea48299d6db0659fcbbc77a6aac1a8fcacd75c6b1de1a14d0e147a65049b60001b610708565b6002816040516151ea9190617d0f565b90815260200160405180910390206040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461524590618303565b80601f016020809104026020016040519081016040528092919081815260200182805461527190618303565b80156152be5780601f10615293576101008083540402835291602001916152be565b820191906000526020600020905b8154815290600101906020018083116152a157829003601f168201915b505050505081526020016002820180546152d790618303565b80601f016020809104026020016040519081016040528092919081815260200182805461530390618303565b80156153505780601f1061532557610100808354040283529160200191615350565b820191906000526020600020905b81548152906001019060200180831161533357829003601f168201915b5050505050815250509150509392505050565b6000806153927ffc472cad1b8f070e3d7aa05e69cf133c9dce52fa3f37913c8032ef699231129f60001b610708565b6153be7f314810b000b85644c6c451f948d0a98e99f5395f2afbbcab81bb47903a73f9af60001b610708565b6153ea7f5656a04d13d7ba2af338362a439c2084309c758b6a076c151525adb2b71a0d6660001b610708565b600060036000866000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561552d578382906000526020600020906002020160405180604001604052908160008201805461546a90618303565b80601f016020809104026020016040519081016040528092919081815260200182805461549690618303565b80156154e35780601f106154b8576101008083540402835291602001916154e3565b820191906000526020600020905b8154815290600101906020018083116154c657829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190615437565b50505050905061555f7f91c908db2366b23cc8dbaa5f61b5d61c276f043e2b47b07d6a55ba384207163760001b610708565b61558b7f6add403f5596368dd8845e16e962aee7833a7de1ef4a2c17ef2c0212f1eaafde60001b610708565b600081511415615625576155c17fb0fbfcdf0d7c9ecceaa0d3c55d517c9b4291b87bb3d4b4ebd8892b843d6c580d60001b610708565b6155ed7f33ac0c7c9819c6e3d0a3b89bb3e85a3eb050da15bd2bcd09068742ed49a1986c60001b610708565b6156187e8d180ee8d236494a86eb79318aa6350b9104351b6cfd8852c998ff1684090560001b610708565b60008092509250506159dc565b6156517f0943b81a3a885523673274dc1d1117c93c782a4bbc0bb82935848aef618e200260001b610708565b61567d7fb032e2095cbc28d0f235a2e160ffd1663cdb408b07da7abf28e8705ccd19781a60001b610708565b6156a97f3e572d93b794231a5458aafdb92dcb1b4f65275fb49e2687f0d4ad4faa944e0f60001b610708565b838051906020012085602001518051906020012014615753576156ee7fdd6285bdab49765913189d2591c9cd826d6f3d11285880b0bf84fc7314b975d960001b610708565b61571a7f5ecc780fa85463ac070e8837e6a4db9fc2cfc9bbfc3b21c3b67a6fbbc6301d4060001b610708565b6157467f3b2a1f7cafd6d43902d0786ec0bd3bda25ed5e1588a5b9ebdddaaff07ee8c62460001b610708565b60008092509250506159dc565b61577f7f5a94db2011d9fe5f0c012cfcedcae984db69fb1032527aae707a34e27832f09e60001b610708565b6157ab7f97b83973de14e0b49b59da0d55a7a0c1c648b782e2304300f0d3e1641d73c33660001b610708565b6157d77f4243a46fc39fb778460fc80f40bd695c16d52cba9d56b5fbc50e4188deabdcd960001b610708565b60005b81518167ffffffffffffffff16101561597a576158197f43c8229e4e5cccc637468470508129e6eb806725c74a0d0f55e7029158e70fa160001b610708565b6158457f4da08d767b1789a0c4c4660f0de8139fc272d9bd7796d3bc543e63e90de744d160001b610708565b8480519060200120828267ffffffffffffffff1681518110615890577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015180519060200120141561593b576158d57f80ceed0912e79f43a9b3efcc9bc75484dddf5524f8adc0d47010f5eea644427360001b610708565b6159017f8d284bb2237d8acb7757279d76baacb2bcd574ab4bf9cebc88c0372b6a13538d60001b610708565b61592d7f3cd0f6175d75479ff04f2a14ab7fb2fb9fa9bc69d243495cb07cc2402f6b5c3060001b610708565b8060019350935050506159dc565b6159677f4233ddd007c76fd2e7c8da18b2a61666a7700da11929510859e9fd5b79a1ee1760001b610708565b8080615972906183af565b9150506157da565b506159a77fa0dbef2e65b1bfc50b116a6e248ff521858bbc99c8ceceef5f0af11c812e0cdb60001b610708565b6159d37f568dbce0ee71b0470c70a6cd48fdbea18736945aecd6227faeac9e0e8066cf4f60001b610708565b60008092509250505b9250929050565b615a0f7fa9b32b63360a0ea16f25995ef904699b071a7a5af893d0fe410792ebe980c8bd60001b610708565b615a3b7fa091b05eed7f65636ca2cdbf50db950beb8be4f178047e7b23a211e32f03fddf60001b610708565b615a677f45491887f3a4ac8f9f1a716a8b97eedcb4e19e9ac7bc0bbc58f3b0a88135c4ca60001b610708565b600083838360000151604051602001615a8293929190617cd2565b6040516020818303038152906040529050615abf7fb1ba2fba7a788206300a9ed48d2311a7358279a16bb540074afda3b37276c6c360001b610708565b615aeb7f43ffcbe6a729df46768767367a80c4b908fcd4157cb1b83d66423efad201315960001b610708565b81600282604051615afc9190617d0f565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151816001019080519060200190615b55929190616a6b565b506040820151816002019080519060200190615b72929190616a6b565b5090505050505050565b600080823b905060008111915050919050565b615bbb7f30ca766cfe97e4b7b3bb81700a641ea3b32825709b3c351667213a10c627d4fe60001b610708565b615be77ff0a37467ddbb3829c9f1dc6c3faa42480a0722af0c99d7c6ac2af9a2d6bfbc0a60001b610708565b615c137f0dcb6b408e3b3e883dcbd215f0782196a4283692b3d38309fb70b178b59f9ad960001b610708565b6000615c1f8383614fa0565b9050615c4d7f410b61d7b79d9eb9371bb9f3aaa3a1a1151e3c4bc6a63c015d407b8a116b366660001b610708565b615c797fe7cb7738e5ebb9973071aac5b57dc4adc6c9c783995608a43a190b17ab4adb8c60001b610708565b60005b8167ffffffffffffffff168167ffffffffffffffff161015615d0e57615cc47f70290edf72b863994382da5bc09e738640df1e16815df09326eb6eb45dceb0ea60001b610708565b615cf07fa0a6466d9faa50161c316a2e644fa739b23bad20be1517c2b03ed5001d178fa460001b610708565b615cfb84848361640e565b8080615d06906183af565b915050615c7c565b50505050565b615d407fe95eec86f1a043f4a812724bf97fab5e06394b9434688c3eb31a4e664576173560001b610708565b615d6c7f280c358ee84a037f8a65453039f307c9051ac2bf1e9584be4acf0410c7a6420c60001b610708565b615d987ff6f02c9b3f370303ab30cc08dec3cb1fd2af2425041dcdb18c79cd15d48e5f7a60001b610708565b60005b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508167ffffffffffffffff16101561617457615e1b7feb6ab60f659c6b76ad666a94d5c5b58b143ab9b8278380c9f29037751b837bc060001b610708565b615e477ffca22b4c4999db08ae9e0812cd63683a4d7aa291e1b174a201d591446a415e7960001b610708565b8167ffffffffffffffff16600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208267ffffffffffffffff1681548110615ed3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16141561613557615f357faf623cf28fcd2130b6d9f334729d4319aa26375e28c06a8b4d5f910688ae982560001b610708565b615f617f20b0843d8e48a731417b60e65da298fdc31fb8dc7ffc51e59d0af08f7177b20d60001b610708565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208167ffffffffffffffff1681548110615fe2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060060201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff02191690556005820160006161029190616af1565b50506161307fcccee5c8e6b6ecb5ccd6a7eed0021686708b7f2dda961754d1cd16a771d29cd360001b610708565b616174565b6161617f55bc55036f1410647a205886bbd20fc60c471df3347be494e27dd24f9679193c60001b610708565b808061616c906183af565b915050615d9b565b505050565b6161a57f93b0a643e9fbf7d81bffe63554f55846866565eadfd8288ae4f197ad43b3548b60001b610708565b6161d17fe28872f838ff3721ff78fb46da5053417db786bc4e9fe37369e702fa1cab972460001b610708565b6161fd7f86e8198551edc7c0f7ee70d937331f587241a2008129cae564cd40eb5125f12060001b610708565b60008484846000015160405160200161621893929190617cd2565b60405160208183030381529060405290506162557f5c3a21513294b8b6adc616be2794ef0f18e2501e4572a5066e0340d57da7c53160001b610708565b6162817fb3b96e0c6c140fa286814d6983a47c8660543084c2a95b215d9078c211c3b89360001b610708565b826002826040516162929190617d0f565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160010190805190602001906162eb929190616a6b565b506040820151816002019080519060200190616308929190616a6b565b509050506163387fb1b938147c593e8cbea165cad2cfed128f79fb59823a7d9b05a581472cedffd560001b610708565b6163647f26d288c13b61475b0bc42638f2e5095d509caf4da2b30373965164c1c656b01b60001b610708565b60036000846000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002082908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906163d5929190616a6b565b5060208201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b61643a7f54f198330c96315d9e50278bc21d317232fd6d02ff2426c905da3aa98c6ae48c60001b610708565b6164667f9a611bba6aa85df2c63a02ef84c7c91867ef8f1ed1b6073edbcfa185ed3303fa60001b610708565b6164927f80278e146f18bb776a8ec8db30fc5dbef5310fd89b47de3ba50de7d6c6e9823a60001b610708565b60008383836040516020016164a993929190617cd2565b60405160208183030381529060405290506164e67f8d2daa5c42dd47c202f87ed942d98b83d07f0e3b6b739d9a46141edbc523455960001b610708565b6002816040516164f69190617d0f565b9081526020016040518091039020600080820160006101000a81549067ffffffffffffffff02191690556001820160006165309190616a00565b6002820160006165409190616a00565b505061656e7f9d241acda3cb91ae9aa77b6fd9724997960bf2da713992d5476de0f23a25184760001b610708565b600360008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006165a19190616b12565b50505050565b8280548282559060005260206000209081019282156165f6579160200282015b828111156165f55782518290805190602001906165e5929190616a6b565b50916020019190600101906165c7565b5b5090506166039190616b36565b5090565b8280548282559060005260206000209060060281019282156169025760005260206000209160060282015b828111156169015782826000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900467ffffffffffffffff168160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160089054906101000a900467ffffffffffffffff168160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160109054906101000a900460ff168160010160106101000a81548160ff021916908360028111156167ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060028201548160020155600382015481600301556004820160009054906101000a900467ffffffffffffffff168160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160089054906101000a900467ffffffffffffffff168160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160109054906101000a900467ffffffffffffffff168160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160189054906101000a900460ff168160040160186101000a81548160ff02191690831515021790555060058201816005019080546168ef929190616b5a565b50505091600601919060060190616632565b5b50905061690f9190616bc2565b5090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600060028111156169a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b508054616a0c90618303565b6000825580601f10616a1e5750616a3d565b601f016020900490600052602060002090810190616a3c9190616cea565b5b50565b6040518060600160405280600067ffffffffffffffff16815260200160608152602001606081525090565b828054616a7790618303565b90600052602060002090601f016020900481019282616a995760008555616ae0565b82601f10616ab257805160ff1916838001178555616ae0565b82800160010185558215616ae0579182015b82811115616adf578251825591602001919060010190616ac4565b5b509050616aed9190616cea565b5090565b5080546000825590600052602060002090810190616b0f9190616b36565b50565b5080546000825560020290600052602060002090810190616b339190616d07565b50565b5b80821115616b565760008181616b4d9190616a00565b50600101616b37565b5090565b828054828255906000526020600020908101928215616bb15760005260206000209182015b82811115616bb0578282908054616b9590618303565b616ba0929190616d49565b5091600101919060010190616b7f565b5b509050616bbe9190616b36565b5090565b5b80821115616ce657600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff0219169055600582016000616cdd9190616af1565b50600601616bc3565b5090565b5b80821115616d03576000816000905550600101616ceb565b5090565b5b80821115616d455760008082016000616d219190616a00565b6001820160006101000a81549067ffffffffffffffff021916905550600201616d08565b5090565b828054616d5590618303565b90600052602060002090601f016020900481019282616d775760008555616dc5565b82601f10616d885780548555616dc5565b82800160010185558215616dc557600052602060002091601f016020900482015b82811115616dc4578254825591600101919060010190616da9565b5b509050616dd29190616cea565b5090565b6000616de9616de484617f8b565b617f66565b90508083825260208201905082856020860282011115616e0857600080fd5b60005b85811015616e385781616e1e8882616f06565b845260208401935060208301925050600181019050616e0b565b5050509392505050565b6000616e55616e5084617fb7565b617f66565b90508083825260208201905082856020860282011115616e7457600080fd5b60005b85811015616ebe57813567ffffffffffffffff811115616e9657600080fd5b808601616ea38982616fc3565b85526020850194506020840193505050600181019050616e77565b5050509392505050565b6000616edb616ed684617fe3565b617f66565b905082815260208101848484011115616ef357600080fd5b616efe848285618297565b509392505050565b600081359050616f1581618641565b92915050565b600081519050616f2a81618641565b92915050565b600082601f830112616f4157600080fd5b8135616f51848260208601616dd6565b91505092915050565b600082601f830112616f6b57600080fd5b8135616f7b848260208601616e42565b91505092915050565b600081359050616f9381618658565b92915050565b600081519050616fa881618658565b92915050565b600081359050616fbd8161866f565b92915050565b600082601f830112616fd457600080fd5b8135616fe4848260208601616ec8565b91505092915050565b600081359050616ffc81618686565b92915050565b6000813590506170118161869d565b92915050565b600081359050617026816186ad565b92915050565b6000610320828403121561703f57600080fd5b61704a6102e0617f66565b9050600082013567ffffffffffffffff81111561706657600080fd5b61707284828501616fc3565b600083015250602061708684828501616f06565b602083015250604082013567ffffffffffffffff8111156170a657600080fd5b6170b284828501616fc3565b60408301525060606170c6848285016175ab565b60608301525060806170da848285016175ab565b60808301525060a06170ee848285016175ab565b60a08301525060c0617102848285016175ab565b60c08301525060e0617116848285016175ab565b60e08301525061010061712b84828501617596565b61010083015250610120617141848285016175ab565b61012083015250610140617157848285016175ab565b6101408301525061016082013567ffffffffffffffff81111561717957600080fd5b61718584828501616fc3565b6101608301525061018061719b848285016175ab565b610180830152506101a06171b184828501617596565b6101a0830152506101c06171c784828501616f84565b6101c0830152506101e06171dd84828501617017565b6101e0830152506102006171f3848285016175ab565b6102008301525061022082013567ffffffffffffffff81111561721557600080fd5b61722184828501616f30565b6102208301525061024082013567ffffffffffffffff81111561724357600080fd5b61724f84828501616f30565b6102408301525061026082013567ffffffffffffffff81111561727157600080fd5b61727d84828501616fc3565b6102608301525061028061729384828501617002565b610280830152506102a06172a984828501616f84565b6102a0830152506102c06172bf8482850161737c565b6102c08301525092915050565b600060e082840312156172de57600080fd5b6172e860e0617f66565b905060006172f8848285016175c0565b600083015250602061730c848285016175c0565b6020830152506040617320848285016175c0565b6040830152506060617334848285016175c0565b6060830152506080617348848285016175c0565b60808301525060a061735c84828501616f1b565b60a08301525060c061737084828501616f1b565b60c08301525092915050565b60006060828403121561738e57600080fd5b6173986060617f66565b905060006173a8848285016175ab565b60008301525060206173bc848285016175ab565b60208301525060406173d0848285016175ab565b60408301525092915050565b6000602082840312156173ee57600080fd5b6173f86020617f66565b90506000617408848285016175ab565b60008301525092915050565b6000610180828403121561742757600080fd5b617432610180617f66565b9050600061744284828501616f06565b6000830152506020617456848285016175ab565b602083015250604061746a848285016175ab565b604083015250606061747e848285016175ab565b606083015250608061749284828501617002565b60808301525060a06174a684828501617596565b60a08301525060c06174ba84828501617596565b60c08301525060e06174ce848285016175ab565b60e0830152506101006174e3848285016175ab565b610100830152506101206174f9848285016175ab565b6101208301525061014061750f84828501616f84565b6101408301525061016082013567ffffffffffffffff81111561753157600080fd5b61753d84828501616f5a565b6101608301525092915050565b60006040828403121561755c57600080fd5b6175666040617f66565b9050600061757684828501616f06565b600083015250602061758a848285016175ab565b60208301525092915050565b6000813590506175a5816186bd565b92915050565b6000813590506175ba816186d4565b92915050565b6000815190506175cf816186d4565b92915050565b6000602082840312156175e757600080fd5b60006175f584828501616f06565b91505092915050565b60006020828403121561761057600080fd5b600061761e84828501616f99565b91505092915050565b60006020828403121561763957600080fd5b600061764784828501616fae565b91505092915050565b6000806040838503121561766357600080fd5b600061767185828601616fed565b9250506020617682858286016173dc565b9150509250929050565b600060e0828403121561769e57600080fd5b60006176ac848285016172cc565b91505092915050565b6000602082840312156176c757600080fd5b600082013567ffffffffffffffff8111156176e157600080fd5b6176ed84828501617414565b91505092915050565b6000806040838503121561770957600080fd5b600083013567ffffffffffffffff81111561772357600080fd5b61772f85828601617414565b925050602083013567ffffffffffffffff81111561774c57600080fd5b6177588582860161702c565b9150509250929050565b60006040828403121561777457600080fd5b60006177828482850161754a565b91505092915050565b600061779783836178f0565b905092915050565b60006177ab8383617a77565b905092915050565b6177bc816181b1565b82525050565b6177cb816181b1565b82525050565b6177e26177dd826181b1565b6183e0565b82525050565b60006177f382618034565b6177fd818561807a565b93508360208202850161780f85618014565b8060005b8581101561784b578484038952815161782c858261778b565b945061783783618060565b925060208a01995050600181019050617813565b50829750879550505050505092915050565b60006178688261803f565b617872818561808b565b93508360208202850161788485618024565b8060005b858110156178c057848403895281516178a1858261779f565b94506178ac8361806d565b925060208a01995050600181019050617888565b50829750879550505050505092915050565b6178db816181c3565b82525050565b6178ea816181c3565b82525050565b60006178fb8261804a565b617905818561809c565b93506179158185602086016182a6565b61791e816184d2565b840191505092915050565b6179328161824f565b82525050565b61794181618261565b82525050565b61795081618261565b82525050565b61795f81618273565b82525050565b61796e81618285565b82525050565b600061797f82618055565b61798981856180be565b93506179998185602086016182a6565b80840191505092915050565b60006179b26015836180ad565b91506179bd826184fd565b602082019050919050565b60006179d56013836180ad565b91506179e082618526565b602082019050919050565b60006179f86014836180ad565b9150617a038261854f565b602082019050919050565b6000617a1b602e836180ad565b9150617a2682618578565b604082019050919050565b6000617a3e6011836180ad565b9150617a49826185c7565b602082019050919050565b6000617a61601f836180ad565b9150617a6c826185f0565b602082019050919050565b600061018083016000830151617a9060008601826177b3565b506020830151617aa36020860182617c9d565b506040830151617ab66040860182617c9d565b506060830151617ac96060860182617c9d565b506080830151617adc6080860182617938565b5060a0830151617aef60a0860182617c7f565b5060c0830151617b0260c0860182617c7f565b5060e0830151617b1560e0860182617c9d565b50610100830151617b2a610100860182617c9d565b50610120830151617b3f610120860182617c9d565b50610140830151617b546101408601826178d2565b50610160830151848203610160860152617b6e82826177e8565b9150508091505092915050565b600061018083016000830151617b9460008601826177b3565b506020830151617ba76020860182617c9d565b506040830151617bba6040860182617c9d565b506060830151617bcd6060860182617c9d565b506080830151617be06080860182617938565b5060a0830151617bf360a0860182617c7f565b5060c0830151617c0660c0860182617c7f565b5060e0830151617c1960e0860182617c9d565b50610100830151617c2e610100860182617c9d565b50610120830151617c43610120860182617c9d565b50610140830151617c586101408601826178d2565b50610160830151848203610160860152617c7282826177e8565b9150508091505092915050565b617c8881618231565b82525050565b617c9781618231565b82525050565b617ca68161823b565b82525050565b617cb58161823b565b82525050565b617ccc617cc78261823b565b618404565b82525050565b6000617cde82866177d1565b601482019150617cee8285617cbb565b600882019150617cfe8284617cbb565b600882019150819050949350505050565b6000617d1b8284617974565b915081905092915050565b6000602082019050617d3b60008301846177c2565b92915050565b60006020820190508181036000830152617d5b818461785d565b905092915050565b6000608082019050617d786000830187617929565b617d856020830186617c8e565b617d9260408301856177c2565b617d9f6060830184617cac565b95945050505050565b600060e082019050617dbd600083018a617929565b617dca6020830189617c8e565b617dd760408301886177c2565b617de46060830187617cac565b617df16080830186617947565b617dfe60a0830185617cac565b617e0b60c08301846178e1565b98975050505050505050565b6000604082019050617e2c6000830185617956565b617e396020830184617cac565b9392505050565b6000602082019050617e556000830184617965565b92915050565b60006020820190508181036000830152617e74816179a5565b9050919050565b60006020820190508181036000830152617e94816179c8565b9050919050565b60006020820190508181036000830152617eb4816179eb565b9050919050565b60006020820190508181036000830152617ed481617a0e565b9050919050565b60006020820190508181036000830152617ef481617a31565b9050919050565b60006020820190508181036000830152617f1481617a54565b9050919050565b60006020820190508181036000830152617f358184617b7b565b905092915050565b6000604082019050617f526000830185617cac565b617f5f6020830184617cac565b9392505050565b6000617f70617f81565b9050617f7c8282618335565b919050565b6000604051905090565b600067ffffffffffffffff821115617fa657617fa56184a3565b5b602082029050602081019050919050565b600067ffffffffffffffff821115617fd257617fd16184a3565b5b602082029050602081019050919050565b600067ffffffffffffffff821115617ffe57617ffd6184a3565b5b618007826184d2565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006180d48261823b565b91506180df8361823b565b92508267ffffffffffffffff038211156180fc576180fb618416565b5b828201905092915050565b60006181128261823b565b915061811d8361823b565b92508167ffffffffffffffff048311821515161561813e5761813d618416565b5b828202905092915050565b600061815482618231565b915061815f83618231565b92508282101561817257618171618416565b5b828203905092915050565b60006181888261823b565b91506181938361823b565b9250828210156181a6576181a5618416565b5b828203905092915050565b60006181bc82618211565b9050919050565b60008115159050919050565b6000819050919050565b60006181e4826181b1565b9050919050565b60008190506181f982618619565b919050565b600081905061820c8261862d565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061825a826181eb565b9050919050565b600061826c826181fe565b9050919050565b600061827e8261823b565b9050919050565b60006182908261823b565b9050919050565b82818337600083830152505050565b60005b838110156182c45780820151818401526020810190506182a9565b838111156182d3576000848401525b50505050565b60006182e48261823b565b915060008214156182f8576182f7618416565b5b600182039050919050565b6000600282049050600182168061831b57607f821691505b6020821081141561832f5761832e618474565b5b50919050565b61833e826184d2565b810181811067ffffffffffffffff8211171561835d5761835c6184a3565b5b80604052505050565b600061837182618231565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156183a4576183a3618416565b5b600182019050919050565b60006183ba8261823b565b915067ffffffffffffffff8214156183d5576183d4618416565b5b600182019050919050565b60006183eb826183f2565b9050919050565b60006183fd826184f0565b9050919050565b600061840f826184e3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f736563746f7220616c7265616479206578697374730000000000000000000000600082015250565b7f6e6f6465206e6f74207265676973746572656400000000000000000000000000600082015250565b7f736563746f722073697a652069732077726f6e67000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f736563746f7249642069732077726f6e67000000000000000000000000000000600082015250565b7f736563746f724964206d7573742062652067726561746572207468616e203000600082015250565b600a811061862a57618629618445565b5b50565b6003811061863e5761863d618445565b5b50565b61864a816181b1565b811461865557600080fd5b50565b618661816181c3565b811461866c57600080fd5b50565b618678816181cf565b811461868357600080fd5b50565b61868f816181d9565b811461869a57600080fd5b50565b600381106186aa57600080fd5b50565b600281106186ba57600080fd5b50565b6186c681618231565b81146186d157600080fd5b50565b6186dd8161823b565b81146186e857600080fd5b5056fea2646970667358221220f30c7d7ceaaf4c23614020b29962bba466e604a3541150e66ff927068ff612e664736f6c63430008040033",
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

// GetSectorInfo is a free data retrieval call binding the contract method 0x2ba010d7.
//
// Solidity: function GetSectorInfo((address,uint64) sectorRef) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]))
func (_Store *StoreCaller) GetSectorInfo(opts *bind.CallOpts, sectorRef SectorRef) (SectorInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSectorInfo", sectorRef)

	if err != nil {
		return *new(SectorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SectorInfo)).(*SectorInfo)

	return out0, err

}

// GetSectorInfo is a free data retrieval call binding the contract method 0x2ba010d7.
//
// Solidity: function GetSectorInfo((address,uint64) sectorRef) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]))
func (_Store *StoreSession) GetSectorInfo(sectorRef SectorRef) (SectorInfo, error) {
	return _Store.Contract.GetSectorInfo(&_Store.CallOpts, sectorRef)
}

// GetSectorInfo is a free data retrieval call binding the contract method 0x2ba010d7.
//
// Solidity: function GetSectorInfo((address,uint64) sectorRef) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]))
func (_Store *StoreCallerSession) GetSectorInfo(sectorRef SectorRef) (SectorInfo, error) {
	return _Store.Contract.GetSectorInfo(&_Store.CallOpts, sectorRef)
}

// GetSectorsForNode is a free data retrieval call binding the contract method 0xe3168f9e.
//
// Solidity: function GetSectorsForNode(address nodeAddr) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[])[])
func (_Store *StoreCaller) GetSectorsForNode(opts *bind.CallOpts, nodeAddr common.Address) ([]SectorInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSectorsForNode", nodeAddr)

	if err != nil {
		return *new([]SectorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]SectorInfo)).(*[]SectorInfo)

	return out0, err

}

// GetSectorsForNode is a free data retrieval call binding the contract method 0xe3168f9e.
//
// Solidity: function GetSectorsForNode(address nodeAddr) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[])[])
func (_Store *StoreSession) GetSectorsForNode(nodeAddr common.Address) ([]SectorInfo, error) {
	return _Store.Contract.GetSectorsForNode(&_Store.CallOpts, nodeAddr)
}

// GetSectorsForNode is a free data retrieval call binding the contract method 0xe3168f9e.
//
// Solidity: function GetSectorsForNode(address nodeAddr) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[])[])
func (_Store *StoreCallerSession) GetSectorsForNode(nodeAddr common.Address) ([]SectorInfo, error) {
	return _Store.Contract.GetSectorsForNode(&_Store.CallOpts, nodeAddr)
}

// C0x29d6276d is a free data retrieval call binding the contract method 0x0daf9945.
//
// Solidity: function c_0x29d6276d(bytes32 c__0x29d6276d) pure returns()
func (_Store *StoreCaller) C0x29d6276d(opts *bind.CallOpts, c__0x29d6276d [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x29d6276d", c__0x29d6276d)

	if err != nil {
		return err
	}

	return err

}

// C0x29d6276d is a free data retrieval call binding the contract method 0x0daf9945.
//
// Solidity: function c_0x29d6276d(bytes32 c__0x29d6276d) pure returns()
func (_Store *StoreSession) C0x29d6276d(c__0x29d6276d [32]byte) error {
	return _Store.Contract.C0x29d6276d(&_Store.CallOpts, c__0x29d6276d)
}

// C0x29d6276d is a free data retrieval call binding the contract method 0x0daf9945.
//
// Solidity: function c_0x29d6276d(bytes32 c__0x29d6276d) pure returns()
func (_Store *StoreCallerSession) C0x29d6276d(c__0x29d6276d [32]byte) error {
	return _Store.Contract.C0x29d6276d(&_Store.CallOpts, c__0x29d6276d)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) AddFileToSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToSector", sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) payable returns()
func (_Store *StoreTransactor) AddSectorRefForFileInfo(opts *bind.TransactOpts, sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddSectorRefForFileInfo", sectorInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) payable returns()
func (_Store *StoreSession) AddSectorRefForFileInfo(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.AddSectorRefForFileInfo(&_Store.TransactOpts, sectorInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) payable returns()
func (_Store *StoreTransactorSession) AddSectorRefForFileInfo(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.AddSectorRefForFileInfo(&_Store.TransactOpts, sectorInfo)
}

// CreateSector is a paid mutator transaction binding the contract method 0xba921004.
//
// Solidity: function CreateSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreTransactor) CreateSector(opts *bind.TransactOpts, sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CreateSector", sectorInfo)
}

// CreateSector is a paid mutator transaction binding the contract method 0xba921004.
//
// Solidity: function CreateSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreSession) CreateSector(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.CreateSector(&_Store.TransactOpts, sectorInfo)
}

// CreateSector is a paid mutator transaction binding the contract method 0xba921004.
//
// Solidity: function CreateSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreTransactorSession) CreateSector(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.CreateSector(&_Store.TransactOpts, sectorInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) DeleteFileFromSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFileFromSector", sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteSector is a paid mutator transaction binding the contract method 0x931bb19a.
//
// Solidity: function DeleteSector((address,uint64) sectorRef) returns()
func (_Store *StoreTransactor) DeleteSector(opts *bind.TransactOpts, sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteSector", sectorRef)
}

// DeleteSector is a paid mutator transaction binding the contract method 0x931bb19a.
//
// Solidity: function DeleteSector((address,uint64) sectorRef) returns()
func (_Store *StoreSession) DeleteSector(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.DeleteSector(&_Store.TransactOpts, sectorRef)
}

// DeleteSector is a paid mutator transaction binding the contract method 0x931bb19a.
//
// Solidity: function DeleteSector((address,uint64) sectorRef) returns()
func (_Store *StoreTransactorSession) DeleteSector(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.DeleteSector(&_Store.TransactOpts, sectorRef)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) payable returns()
func (_Store *StoreTransactor) UpdateSectorInfo(opts *bind.TransactOpts, sector SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateSectorInfo", sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) payable returns()
func (_Store *StoreSession) UpdateSectorInfo(sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) payable returns()
func (_Store *StoreTransactorSession) UpdateSectorInfo(sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, sector)
}

// Initialize is a paid mutator transaction binding the contract method 0x112346c2.
//
// Solidity: function initialize(address _node, (uint64) sectorConfig) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _node common.Address, sectorConfig SectorConfig) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _node, sectorConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x112346c2.
//
// Solidity: function initialize(address _node, (uint64) sectorConfig) returns()
func (_Store *StoreSession) Initialize(_node common.Address, sectorConfig SectorConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _node, sectorConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x112346c2.
//
// Solidity: function initialize(address _node, (uint64) sectorConfig) returns()
func (_Store *StoreTransactorSession) Initialize(_node common.Address, sectorConfig SectorConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _node, sectorConfig)
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
