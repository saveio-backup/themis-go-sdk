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
	_ = abi.ConvertType
)

// FileProveParams is an auto generated low-level Go binding around an user-defined struct.
type FileProveParams struct {
	FileHash    []byte
	ProveData   []byte
	BlockHeight *big.Int
	NodeWallet  common.Address
	Profit      uint64
	SectorID    uint64
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

// NodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeInfo struct {
	Pledge      uint64
	Profit      uint64
	Volume      uint64
	RestVol     uint64
	ServiceTime uint64
	WalletAddr  common.Address
	NodeAddr    []byte
}

// ProveConfig is an auto generated low-level Go binding around an user-defined struct.
type ProveConfig struct {
	SECTORPROVEBLOCKNUM uint64
}

// ProveDetail is an auto generated low-level Go binding around an user-defined struct.
type ProveDetail struct {
	NodeAddr    []byte
	WalletAddr  common.Address
	ProveTimes  uint64
	BlockHeight *big.Int
	Finished    bool
}

// ProveDetailMeta is an auto generated low-level Go binding around an user-defined struct.
type ProveDetailMeta struct {
	CopyNum        uint64
	ProveDetailNum uint64
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

// SectorProveParams is an auto generated low-level Go binding around an user-defined struct.
type SectorProveParams struct {
	NodeAddr        common.Address
	SectorID        uint64
	ChallengeHeight uint64
	ProveData       []byte
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

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structFileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"checkSectorProve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractIPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_PROVE_BLOCK_NUM\",\"type\":\"uint64\"}],\"internalType\":\"structProveConfig\",\"name\":\"proveConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractProveExtra\",\"name\":\"_proveExtra\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"profitSplitForSector\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615cb480620000216000396000f3fe6080604052600436106100c75760003560e01c8063648d225d11610074578063a00044121161004e578063a0004412146101e1578063bb4e4e42146101f4578063d2561e0a1461020757600080fd5b8063648d225d1461018e5780638e270530146101a1578063977fdfe2146101b457600080fd5b8063181197f7116100a5578063181197f71461012a5780633ec0f5df1461013d57806352e061461461017b57600080fd5b806309cbe391146100cc5780630e3459fd146100e15780630fece8691461010a575b600080fd5b6100df6100da3660046142f5565b610227565b005b6100f46100ef3660046141f4565b61070c565b6040516101019190615420565b60405180910390f35b34801561011657600080fd5b506100f4610125366004614329565b610ae6565b6100df610138366004613f55565b610f90565b6100df61014b366004613e4e565b6001600160a01b0390921660009081526006602090815260408083206001600160401b0390941683529290522055565b6100df610189366004613f21565b611011565b6100df61019c366004614386565b61108f565b6100df6101af366004614106565b6113b9565b3480156101c057600080fd5b506101d46101cf366004613f21565b612415565b604051610101919061540f565b6100df6101ef366004614265565b6124b3565b6100df610202366004613fbc565b6125f1565b34801561021357600080fd5b506100df610222366004614002565b61263c565b6002548151604051630d329ba560e11b81526000926001600160a01b031691631a65374a9161025991906004016153e6565b60006040518083038186803b15801561027157600080fd5b505afa158015610285573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102ad919081019061413a565b6004805460408051808201825286516001600160a01b0390811682526020808901516001600160401b0316908301529151632ba010d760e01b81529495506000949190921692632ba010d79261030592909101615902565b60006040518083038186803b15801561031d57600080fd5b505afa158015610331573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261035991908101906141c0565b905080602001516001600160401b03166000141561039957600080516020615c5f83398151915260405161038c9061568f565b60405180910390a1505050565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156103e957600080fd5b505afa1580156103fd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042191906143a4565b90508160c0015143101561045857600080516020615c5f83398151915260405161044a906156f8565b60405180910390a150505050565b8160c0015184604001516001600160401b03161461048b57600080516020615c5f83398151915260405161044a9061566c565b60006104978584610ae6565b9050806104d5576104ab83858460016124b3565b600080516020615c5f8339815191526040516104c6906156b2565b60405180910390a15050505050565b60006104e284868561070c565b90508061051457600080516020615c5f83398151915260405161050490615649565b60405180910390a1505050505050565b60a0840151610524574360a08501525b60c083015161053c906001600160401b031643615998565b60c0850152600480546040516311c2535560e11b81526001600160a01b0390911691632384a6aa91610570918891016158cc565b600060405180830381600087803b15801561058a57600080fd5b505af115801561059e573d6000803e3d6000fd5b505050508361014001516105c757600080516020615c5f833981519152604051610504906156d5565b60055484516040517f9908f2bf0000000000000000000000000000000000000000000000000000000081526000926001600160a01b031691639908f2bf91610614919043906004016153f4565b60606040518083038186803b15801561062c57600080fd5b505afa158015610640573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066491906141a2565b43602082015285516001600160a01b03908116825260608701516001600160401b031660408084019190915260055490517f1581d5450000000000000000000000000000000000000000000000000000000081529293501690631581d545906106d190849060040161589c565b600060405180830381600087803b1580156106eb57600080fd5b505af11580156106ff573d6000803e3d6000fd5b5050505050505050505050565b6000805b84610160015151811015610ad9576000856101600151828151811061074557634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001546040516337bf397560e21b81529192506000916001600160a01b039091169063defce5d49061078690859060040161542e565b60006040518083038186803b15801561079e57600080fd5b505afa1580156107b2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107da91908101906140d2565b905060006107eb8260000151612415565b6101008301516040805160a081018252606080825260006020830181905292820183905281018290526080810182905292935091829060005b8551811015610916578c600001516001600160a01b0316868a8151811061085b57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b03161415610904576001925085898151811061089a57634e487b7160e01b600052603260045260246000fd5b602002602001015191506000826040015190508760e001516001600160401b0316816001600160401b031614806108d057508443115b156108e15760016080840181905295505b604083018051906108f182615b52565b6001600160401b03169052506109169050565b8061090e81615b37565b915050610824565b508161092d57600098505050505050505050610adf565b85518c9087908d908d90610941908a610f90565b8715610abb57600560009054906101000a90046001600160a01b03166001600160a01b031663d63877b9600260009054906101000a90046001600160a01b0316600160009054906101000a90046001600160a01b031686868a8f886040518863ffffffff1660e01b81526004016109be97969594939291906154a4565b600060405180830381600087803b1580156109d857600080fd5b505af11580156109ec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a14919081019061409e565b50600480546040516247c00360e01b81526001600160a01b03909116916247c00391610a449188918891016158dd565b600060405180830381600087803b158015610a5e57600080fd5b505af1158015610a72573d6000803e3d6000fd5b5050845160a08501516040517fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f459450610ab2935060019243929091615599565b60405180910390a15b50505050505050505050508080610ad190615b37565b915050610710565b50600190505b9392505050565b60408051608081018252600080825260606020830181905282840182815281840183815287516001600160a01b03908116865260e08801516001600160401b03908116909352600554600160a01b9004909216905260035494517ffd7c9808000000000000000000000000000000000000000000000000000000008152929491939285929091169063fd7c980890610b8290859060040161587a565b60006040518083038186803b158015610b9a57600080fd5b505afa158015610bae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610bd69190810190613e9b565b9050610c236040518060c0016040528060006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b610c2b6128a8565b86815260408101829052610c3d61296f565b6003546040517f57d60d6e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906357d60d6e90610c869085906004016158aa565b60006040518083038186803b158015610c9e57600080fd5b505afa158015610cb2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cda919081019061416e565b90508060c00151610cf45760009650505050505050610f8a565b610d2f6040518060a0016040528060006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b600080825260408086015160208085019190915284518285015284015160608401526080808501519084015260035490517f507af8050000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063507af80590610da39085906004016158bb565b60206040518083038186803b158015610dbb57600080fd5b505afa158015610dcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df39190613f03565b905080610e0b57600098505050505050505050610f8a565b89610140015115610f7d578260a001516102a001511580610e40575060a08301516102c00151604001516001600160401b0316155b506040805160c081018252600060608083018281526080840183905260a08401839052835260208301529181019190915260a0808501516102c0015182528601516020820152865115610ec65786600081518110610eae57634e487b7160e01b600052603260045260246000fd5b60209081029190910101515163ffffffff1660408201525b6003546040517f2e19aeff0000000000000000000000000000000000000000000000000000000081526000916001600160a01b031690632e19aeff90610f10908590600401615910565b60206040518083038186803b158015610f2857600080fd5b505afa158015610f3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f609190613f03565b905080610f7a5760009a5050505050505050505050610f8a565b50505b6001985050505050505050505b92915050565b6005546040517f181197f70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063181197f790610fdb908590859060040161543f565b600060405180830381600087803b158015610ff557600080fd5b505af1158015611009573d6000803e3d6000fd5b505050505050565b6005546040517f52e061460000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906352e061469061105a90849060040161542e565b600060405180830381600087803b15801561107457600080fd5b505af1158015611088573d6000803e3d6000fd5b5050505050565b80516020820151600254604051630d329ba560e11b81526000916001600160a01b031690631a65374a906110c79086906004016153e6565b60006040518083038186803b1580156110df57600080fd5b505afa1580156110f3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261111b919081019061413a565b90504281608001516001600160401b0316101561114d57600080516020615c5f83398151915260405161044a906155e3565b6001600160401b03821661117657600080516020615c5f83398151915260405161044a90615606565b60048054604051632ba010d760e01b81526000926001600160a01b0390921691632ba010d7916111a891899101615902565b60006040518083038186803b1580156111c057600080fd5b505afa1580156111d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111fc91908101906141c0565b90508061010001516001600160401b03166000141561123057600080516020615c5f8339815191526040516104c690615606565b6000805460808301516040517fc5c81b20000000000000000000000000000000000000000000000000000000008152620100009092046001600160a01b03169163c5c81b2091611282916004016155d5565b6101606040518083038186803b15801561129b57600080fd5b505afa1580156112af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d391906143a4565b90506000439050808260c001516001600160401b03168460c001516112f89190615998565b101561132a57600080516020615c5f83398151915260405161131990615606565b60405180910390a150505050505050565b6001600160a01b03861660009081526006602090815260408083206001600160401b03891684529091528120549061136485858486612791565b90506001600160401b0381166113a257600080516020615c5f83398151915260405161138f90615606565b60405180910390a1505050505050505050565b6113ae858786846124b3565b505050505050505050565b60015481516040516337bf397560e21b81526000926001600160a01b03169163defce5d4916113eb919060040161542e565b60006040518083038186803b15801561140357600080fd5b505afa158015611417573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261143f91908101906140d2565b9050806102a00151156114995780602001516001600160a01b031682606001516001600160a01b03161461149457600080516020615c5f8339815191526040516114889061573f565b60405180910390a15050565b6115ab565b610220810151600090815b815181101561150d5784606001516001600160a01b03168282815181106114db57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156114fb576001925061150d565b8061150581615b37565b9150506114a4565b50816115885761024083015160005b81518110156115855785606001516001600160a01b031682828151811061155357634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156115735760019350611585565b8061157d81615b37565b91505061151c565b50505b816115a857600080516020615c5f83398151915260405161044a90615857565b50505b6002546060830151604051630d329ba560e11b81526000926001600160a01b031691631a65374a916115e091906004016153e6565b60006040518083038186803b1580156115f857600080fd5b505afa15801561160c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611634919081019061413a565b60055460025484516040517f23fdc52c0000000000000000000000000000000000000000000000000000000081529394506000936001600160a01b03938416936323fdc52c9361168993911691600401615484565b60006040518083038186803b1580156116a157600080fd5b505afa1580156116b5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116dd9190810190613ecf565b90508360a001516001600160401b0316600014158015611701575082610100015143105b1561178a5760005b81518110156117885784606001516001600160a01b031682828151811061174057634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b0316141561177657600080516020615c5f8339815191526040516104c690615811565b8061178081615b37565b915050611709565b505b6005546003546040517f483377d60000000000000000000000000000000000000000000000000000000081526000926001600160a01b039081169263483377d6926117df929091169089908990600401615529565b60206040518083038186803b1580156117f757600080fd5b505afa15801561180b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061182f9190613f03565b90508061185157600080516020615c5f8339815191526040516104c6906157cb565b6040805160a08101825260608082526000602083018190529282018390528101829052608081018290528190819061010088015160005b8751811015611a74578a606001516001600160a01b03168882815181106118bf57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b03161415611a6257826040015193508960e001516001600160401b0316846001600160401b0316148061190457508143115b156119155760016080840181905294505b8960e001516001600160401b0316846001600160401b0316111561196357600080516020615c5f83398151915260405161194e906157a8565b60405180910390a15050505050505050505050565b6005546101008b01516040517f624e1dd00000000000000000000000000000000000000000000000000000000081526000926001600160a01b03169163624e1dd0916119b29190600401615921565b60206040518083038186803b1580156119ca57600080fd5b505afa1580156119de573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a029190613f03565b90508015611a3b57600080516020615c5f833981519152604051611a2590615785565b60405180910390a1505050505050505050505050565b60408401805190611a4b82615b52565b6001600160401b031690525060019650611a749050565b80611a6c81615b37565b915050611888565b5084611d1e57610120890151611a8b9060016159b0565b6001600160401b031687511415611acb57600080516020615c5f833981519152604051611ab7906157ee565b60405180910390a150505050505050505050565b8860a001518960800151611adf9190615a04565b6001600160401b031688606001516001600160401b03161015611b1757600080516020615c5f833981519152604051611ab790615834565b8860a001518960800151611b2b9190615a04565b88606001818151611b3c9190615a5b565b6001600160401b0316905250600254604051631f842a6b60e31b81526001600160a01b039091169063fc21535890611b78908b9060040161588b565b600060405180830381600087803b158015611b9257600080fd5b505af1158015611ba6573d6000803e3d6000fd5b50505060c08901518352506060808b01516001600160a01b03166020840152600160408401819052439184019190915260006080840181905288519091611bed9190615998565b6001600160401b03811115611c1257634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611c6c57816020015b6040805160a08101825260608082526000602080840182905293830181905290820181905260808201528252600019909201910181611c305790505b50905060005b8851811015611ce157888181518110611c9b57634e487b7160e01b600052603260045260246000fd5b6020026020010151828281518110611cc357634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080611cd990615b37565b915050611c72565b50828160018351611cf29190615a40565b81518110611d1057634e487b7160e01b600052603260045260246000fd5b602090810291909101015296505b8851611d2a9088610f90565b60a0808b01516004546040808e01518151808301909252938c01516001600160a01b0390811682526001600160401b0384166020830152929392909116918c91896120de57604051632ba010d760e01b81526000906001600160a01b03861690632ba010d790611d9e908590600401615902565b60006040518083038186803b158015611db657600080fd5b505afa158015611dca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611df291908101906141c0565b9050836102a001511515816101400151151514611e3e57600080516020615c5f833981519152604051611e2490615762565b60405180910390a150505050505050505050505050505050565b6040517f955f98b70000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063955f98b790611e8590849088906004016158dd565b600060405180830381600087803b158015611e9f57600080fd5b505af1158015611eb3573d6000803e3d6000fd5b5050604051632ba010d760e01b81526001600160a01b0388169250632ba010d79150611ee3908590600401615902565b60006040518083038186803b158015611efb57600080fd5b505afa158015611f0f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f3791908101906141c0565b9050836102a001511515816101400151151514611f6957600080516020615c5f833981519152604051611e2490615762565b6005546001546040517fde4d268b0000000000000000000000000000000000000000000000000000000081526000926001600160a01b039081169263de4d268b92611fbe928b9216908a90899060040161555d565b600060405180830381600087803b158015611fd857600080fd5b505af1158015611fec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612014919081019061409e565b80519091501561205657600080516020615c5f8339815191528160405161203b919061571b565b60405180910390a15050505050505050505050505050505050565b60c082015161207d5760c0850151612077906001600160401b031685615998565b60c08301525b6040516311c2535560e11b81526001600160a01b03871690632384a6aa906120a99085906004016158cc565b600060405180830381600087803b1580156120c357600080fd5b505af11580156120d7573d6000803e3d6000fd5b5050505050505b60008d9050600088905060008e905060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561213d57600080fd5b505afa158015612151573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061217591906143a4565b90508c156123bc576001600160401b038916156122765760048054604051632ba010d760e01b81526000926001600160a01b0390921691632ba010d7916121be918a9101615902565b60006040518083038186803b1580156121d657600080fd5b505afa1580156121ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261221291908101906141c0565b600480546040516247c00360e01b81529293506001600160a01b0316916247c003916122429185918d91016158dd565b600060405180830381600087803b15801561225c57600080fd5b505af1158015612270573d6000803e3d6000fd5b50505050505b6005546002546001546040517fd63877b90000000000000000000000000000000000000000000000000000000081526000936001600160a01b039081169363d63877b9936122d6939183169216908d908b908b908b908b906004016154a4565b600060405180830381600087803b1580156122f057600080fd5b505af1158015612304573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261232c919081019061409e565b905060405160200161233d906153de565b604051602081830303815290604052805190602001208160405160200161236491906153d2565b60405160208183030381529060405280519060200120146123ba57600080516020615c5f8339815191528160405161239c919061571b565b60405180910390a15050505050505050505050505050505050505050565b505b865160a08501516040517fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6926123f89260079243929190615599565b60405180910390a150505050505050505050505050505050505050565b6005546040517f977fdfe20000000000000000000000000000000000000000000000000000000081526060916001600160a01b03169063977fdfe29061245f90859060040161542e565b60006040518083038186803b15801561247757600080fd5b505afa15801561248b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f8a9190810190613ecf565b60006124bf838661285e565b6124c99083615a04565b9050806001600160401b031684600001516001600160401b03161061250b5780846000018181516124fa9190615a5b565b6001600160401b0316905250612512565b5060008084525b6001600160401b038116156125b657806001600160401b03163410156125535760405162461bcd60e51b815260040161254a90615629565b60405180910390fd5b600254604051631f842a6b60e31b81526001600160a01b039091169063fc2153589061258390879060040161588b565b600060405180830381600087803b15801561259d57600080fd5b505af11580156125b1573d6000803e3d6000fd5b505050505b84516020808701516001600160a01b0390921660009081526006825260408082206001600160401b0390941682529290915220439055611088565b6005546040517fbb4e4e420000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063bb4e4e4290610fdb9085908590600401615464565b600054610100900460ff166126575760005460ff161561265b565b303b155b6126775760405162461bcd60e51b815260040161254a90615639565b600054610100900460ff16158015612699576000805461ffff19166101011790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff16620100006001600160a01b038b81169190910291909117909155600180547fffffffffffffffffffffffff00000000000000000000000000000000000000009081168a8416179091556002805482168984161790556003805482168884161790556004805482168784161790558451600580547fffffffff0000000000000000000000000000000000000000000000000000000016600160a01b6001600160401b0390931692909202909216179184169190911790558015612787576000805461ff00191690555b5050505050505050565b60c0808401519085015160009190836127b36001600160401b03841683615998565b106127c357600092505050612856565b6000826127d08387615a40565b6127da91906159e3565b90506000861561281e576127f76001600160401b03851684615998565b87111561281a57836128098489615a40565b61281391906159e3565b905061281e565b5060005b806001600160401b0316826001600160401b03161015612845576000945050505050612856565b61284f8183615a5b565b9450505050505b949350505050565b600080600290506000612875858560600151612888565b61287f9083615a04565b95945050505050565b6000620fa00082846060015161289e9190615a04565b610adf91906159e3565b604080516101e08101909152600060608083018281526080840183905260a0840183905260c0840183905260e08401839052610100840183905261012084018390526101408401839052610160840183905261018084018390526101a08401929092526101c0830152819081526020016060815260200161296a6040518060c0016040528060006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b905290565b6040518060e001604052806060815260200160608152602001606081526020016060815260200160608152602001612a72604080516102e08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201839052610160820181905261018082018390526101a082018390526101c082018390526101e0820183905261020082018390526102208201819052610240820181905261026082015290610280820190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b8152600060209091015290565b6000612a92612a8d8461594b565b61592f565b90508083825260208201905082856020860282011115612ab157600080fd5b60005b85811015612add5781612ac78882612e60565b8452506020928301929190910190600101612ab4565b5050509392505050565b6000612af5612a8d8461594b565b90508083825260208201905082856020860282011115612b1457600080fd5b60005b85811015612add5781356001600160401b03811115612b3557600080fd5b808601612b428982612f89565b855250506020928301929190910190600101612b17565b6000612b67612a8d8461594b565b90508083825260208201905082856020860282011115612b8657600080fd5b60005b85811015612add5781516001600160401b03811115612ba757600080fd5b808601612bb48982612faa565b855250506020928301929190910190600101612b89565b6000612bd9612a8d8461594b565b83815290506020810182604085028101861015612bf557600080fd5b60005b85811015612add5781612c0b8882612ff7565b84525060209092019160409190910190600101612bf8565b6000612c31612a8d8461594b565b90508083825260208201905082856020860282011115612c5057600080fd5b60005b85811015612add5781516001600160401b03811115612c7157600080fd5b808601612c7e8982613397565b855250506020928301929190910190600101612c53565b6000612ca3612a8d8461594b565b90508083825260208201905082856020860282011115612cc257600080fd5b60005b85811015612add5781516001600160401b03811115612ce357600080fd5b808601612cf08982613409565b855250506020928301929190910190600101612cc5565b6000612d15612a8d8461594b565b90508083825260208201905082856020860282011115612d3457600080fd5b60005b85811015612add5781356001600160401b03811115612d5557600080fd5b808601612d6289826137f4565b855250506020928301929190910190600101612d37565b6000612d87612a8d8461594b565b90508083825260208201905082856020860282011115612da657600080fd5b60005b85811015612add5781516001600160401b03811115612dc757600080fd5b808601612dd4898261388d565b855250506020928301929190910190600101612da9565b6000612df9612a8d8461596e565b905082815260208101848484011115612e1157600080fd5b612e1c848285615ad3565b509392505050565b6000612e32612a8d8461596e565b905082815260208101848484011115612e4a57600080fd5b612e1c848285615adf565b8035610f8a81615bfe565b8051610f8a81615bfe565b600082601f830112612e7c57600080fd5b8151612856848260208601612a7f565b600082601f830112612e9d57600080fd5b8135612856848260208601612ae7565b600082601f830112612ebe57600080fd5b8151612856848260208601612b59565b600082601f830112612edf57600080fd5b8151612856848260208601612bcb565b600082601f830112612f0057600080fd5b8151612856848260208601612c23565b600082601f830112612f2157600080fd5b8151612856848260208601612c95565b600082601f830112612f4257600080fd5b8135612856848260208601612d07565b600082601f830112612f6357600080fd5b8151612856848260208601612d79565b8035610f8a81615c12565b8051610f8a81615c12565b600082601f830112612f9a57600080fd5b8135612856848260208601612deb565b600082601f830112612fbb57600080fd5b8151612856848260208601612e24565b8035610f8a81615c1a565b8035610f8a81615c23565b8051610f8a81615c23565b8051610f8a81615c30565b60006040828403121561300957600080fd5b613013604061592f565b905060006130218484613e2d565b825250602061303284848301613e2d565b60208301525092915050565b6000610320828403121561305157600080fd5b61305c6102e061592f565b82519091506001600160401b0381111561307557600080fd5b61308184828501612faa565b825250602061309284848301612e60565b60208301525060408201516001600160401b038111156130b157600080fd5b6130bd84828501612faa565b60408301525060606130d184828501613e43565b60608301525060806130e584828501613e43565b60808301525060a06130f984828501613e43565b60a08301525060c061310d84828501613e43565b60c08301525060e061312184828501613e43565b60e08301525061010061313684828501613e22565b6101008301525061012061314c84828501613e43565b6101208301525061014061316284828501613e43565b610140830152506101608201516001600160401b0381111561318357600080fd5b61318f84828501612faa565b610160830152506101806131a584828501613e43565b610180830152506101a06131bb84828501613e22565b6101a0830152506101c06131d184828501612f7e565b6101c0830152506101e06131e784828501612fec565b6101e0830152506102006131fd84828501613e43565b610200830152506102208201516001600160401b0381111561321e57600080fd5b61322a84828501612e6b565b610220830152506102408201516001600160401b0381111561324b57600080fd5b61325784828501612e6b565b610240830152506102608201516001600160401b0381111561327857600080fd5b61328484828501612faa565b6102608301525061028061329a84828501612fe1565b610280830152506102a06132b084828501612f7e565b6102a0830152506102c06132c6848285016136fc565b6102c08301525092915050565b600060c082840312156132e557600080fd5b6132ef60c061592f565b905081356001600160401b0381111561330757600080fd5b61331384828501612f89565b82525060208201356001600160401b0381111561332f57600080fd5b61333b84828501612f89565b602083015250604061334f84828501613e17565b604083015250606061336384828501612e55565b606083015250608061337784828501613e38565b60808301525060a061338b84828501613e38565b60a08301525092915050565b6000606082840312156133a957600080fd5b6133b3606061592f565b905060006133c18484613e43565b82525060206133d284848301613e43565b60208301525060408201516001600160401b038111156133f157600080fd5b6133fd84828501612faa565b60408301525092915050565b60006040828403121561341b57600080fd5b613425604061592f565b905060006134338484613e43565b82525060208201516001600160401b0381111561344f57600080fd5b61303284828501612eef565b600060e0828403121561346d57600080fd5b61347760e061592f565b905060006134858484613e38565b825250602061349684848301613e38565b60208301525060406134aa84828501613e38565b60408301525060606134be84828501613e38565b60608301525060806134d284828501613e38565b60808301525060a06134e684828501612e55565b60a08301525060c08201356001600160401b0381111561350557600080fd5b61351184828501612f89565b60c08301525092915050565b600060e0828403121561352f57600080fd5b61353960e061592f565b905060006135478484613e43565b825250602061355884848301613e43565b602083015250604061356c84828501613e43565b604083015250606061358084828501613e43565b606083015250608061359484828501613e43565b60808301525060a06135a884828501612e60565b60a08301525060c08201516001600160401b038111156135c757600080fd5b61351184828501612faa565b600060e082840312156135e557600080fd5b6135ef60e061592f565b82519091506001600160401b0381111561360857600080fd5b61361484828501612faa565b82525060208201516001600160401b0381111561363057600080fd5b61363c84828501612ead565b60208301525060408201516001600160401b0381111561365b57600080fd5b61366784828501612ece565b60408301525060608201516001600160401b0381111561368657600080fd5b61369284828501612f10565b60608301525060808201516001600160401b038111156136b157600080fd5b6136bd84828501612faa565b60808301525060a08201516001600160401b038111156136dc57600080fd5b6136e88482850161303e565b60a08301525060c061351184828501612f7e565b60006060828403121561370e57600080fd5b613718606061592f565b905060006137268484613e43565b825250602061373784848301613e43565b60208301525060406133fd84828501613e43565b60006060828403121561375d57600080fd5b613767606061592f565b905060006137758484612e60565b825250602061373784848301613e22565b60006020828403121561379857600080fd5b6137a2602061592f565b905060006137b08484613e38565b82525092915050565b6000604082840312156137cb57600080fd5b6137d5604061592f565b905060006137e38484613e38565b825250602061303284848301613e38565b600060a0828403121561380657600080fd5b61381060a061592f565b905081356001600160401b0381111561382857600080fd5b61383484828501612f89565b825250602061384584848301612e55565b602083015250604061385984828501613e38565b604083015250606061386d84828501613e17565b606083015250608061388184828501612f73565b60808301525092915050565b600060a0828403121561389f57600080fd5b6138a960a061592f565b82519091506001600160401b038111156138c257600080fd5b6138ce84828501612faa565b82525060206138df84848301612e60565b60208301525060406138f384828501613e43565b604083015250606061390784828501613e22565b606083015250608061388184828501612f7e565b6000610180828403121561392e57600080fd5b61393961018061592f565b905060006139478484612e55565b825250602061395884848301613e38565b602083015250604061396c84828501613e38565b604083015250606061398084828501613e38565b606083015250608061399484828501612fd6565b60808301525060a06139a884828501613e17565b60a08301525060c06139bc84828501613e17565b60c08301525060e06139d084828501613e38565b60e0830152506101006139e584828501613e38565b610100830152506101206139fb84828501613e38565b61012083015250610140613a1184828501612f73565b610140830152506101608201356001600160401b03811115613a3257600080fd5b613a3e84828501612e8c565b6101608301525092915050565b60006101808284031215613a5e57600080fd5b613a6961018061592f565b90506000613a778484612e60565b8252506020613a8884848301613e43565b6020830152506040613a9c84828501613e43565b6040830152506060613ab084828501613e43565b6060830152506080613ac484828501612fe1565b60808301525060a0613ad884828501613e22565b60a08301525060c0613aec84828501613e22565b60c08301525060e0613b0084828501613e43565b60e083015250610100613b1584828501613e43565b61010083015250610120613b2b84828501613e43565b61012083015250610140613b4184828501612f7e565b610140830152506101608201516001600160401b03811115613b6257600080fd5b613a3e84828501612ead565b600060808284031215613b8057600080fd5b613b8a608061592f565b90506000613b988484612e55565b8252506020613ba984848301613e38565b6020830152506040613bbd84828501613e38565b60408301525060608201356001600160401b03811115613bdc57600080fd5b613be884828501612f89565b60608301525092915050565b600060408284031215613c0657600080fd5b613c10604061592f565b905060006137e38484612e55565b60006101608284031215613c3157600080fd5b613c3c61016061592f565b90506000613c4a8484613e38565b8252506020613c5b84848301613e38565b6020830152506040613c6f84828501613e38565b6040830152506060613c8384828501613e38565b6060830152506080613c9784828501613e38565b60808301525060a0613cab84828501613e38565b60a08301525060c0613cbf84828501613e38565b60c08301525060e0613cd384828501613e38565b60e083015250610100613ce884828501613e38565b61010083015250610120613cfe84828501613e38565b61012083015250610140613d1484828501613e38565b6101408301525092915050565b60006101608284031215613d3457600080fd5b613d3f61016061592f565b90506000613d4d8484613e43565b8252506020613d5e84848301613e43565b6020830152506040613d7284828501613e43565b6040830152506060613d8684828501613e43565b6060830152506080613d9a84828501613e43565b60808301525060a0613dae84828501613e43565b60a08301525060c0613dc284828501613e43565b60c08301525060e0613dd684828501613e43565b60e083015250610100613deb84828501613e43565b61010083015250610120613e0184828501613e43565b61012083015250610140613d1484828501613e43565b8035610f8a81615c3d565b8051610f8a81615c3d565b8051610f8a81615c43565b8035610f8a81615c4f565b8051610f8a81615c4f565b600080600060608486031215613e6357600080fd5b6000613e6f8686612e55565b9350506020613e8086828701613e38565b9250506040613e9186828701613e17565b9150509250925092565b600060208284031215613ead57600080fd5b81516001600160401b03811115613ec357600080fd5b61285684828501612ece565b600060208284031215613ee157600080fd5b81516001600160401b03811115613ef757600080fd5b61285684828501612f52565b600060208284031215613f1557600080fd5b60006128568484612f7e565b600060208284031215613f3357600080fd5b81356001600160401b03811115613f4957600080fd5b61285684828501612f89565b60008060408385031215613f6857600080fd5b82356001600160401b03811115613f7e57600080fd5b613f8a85828601612f89565b92505060208301356001600160401b03811115613fa657600080fd5b613fb285828601612f31565b9150509250929050565b60008060608385031215613fcf57600080fd5b82356001600160401b03811115613fe557600080fd5b613ff185828601612f89565b9250506020613fb2858286016137b9565b600080600080600080600060e0888a03121561401d57600080fd5b60006140298a8a612fcb565b975050602061403a8a828b01612fcb565b965050604061404b8a828b01612fcb565b955050606061405c8a828b01612fcb565b945050608061406d8a828b01612fcb565b93505060a061407e8a828b01613786565b92505060c061408f8a828b01612fcb565b91505092959891949750929550565b6000602082840312156140b057600080fd5b81516001600160401b038111156140c657600080fd5b61285684828501612faa565b6000602082840312156140e457600080fd5b81516001600160401b038111156140fa57600080fd5b6128568482850161303e565b60006020828403121561411857600080fd5b81356001600160401b0381111561412e57600080fd5b612856848285016132d3565b60006020828403121561414c57600080fd5b81516001600160401b0381111561416257600080fd5b6128568482850161351d565b60006020828403121561418057600080fd5b81516001600160401b0381111561419657600080fd5b612856848285016135d3565b6000606082840312156141b457600080fd5b6000612856848461374b565b6000602082840312156141d257600080fd5b81516001600160401b038111156141e857600080fd5b61285684828501613a4b565b60008060006101a0848603121561420a57600080fd5b83356001600160401b0381111561422057600080fd5b61422c8682870161391b565b93505060208401356001600160401b0381111561424857600080fd5b6142548682870161345b565b9250506040613e9186828701613c1e565b6000806000806101c0858703121561427c57600080fd5b84356001600160401b0381111561429257600080fd5b61429e8782880161391b565b94505060208501356001600160401b038111156142ba57600080fd5b6142c68782880161345b565b93505060406142d787828801613c1e565b9250506101a06142e987828801613e38565b91505092959194509250565b60006020828403121561430757600080fd5b81356001600160401b0381111561431d57600080fd5b61285684828501613b6e565b6000806040838503121561433c57600080fd5b82356001600160401b0381111561435257600080fd5b61435e85828601613b6e565b92505060208301356001600160401b0381111561437a57600080fd5b613fb28582860161391b565b60006040828403121561439857600080fd5b60006128568484613bf4565b600061016082840312156143b757600080fd5b60006128568484613d21565b60006143cf838361441b565b505060200190565b6000610adf838361462a565b60006143ef8383614b6c565b505060400190565b6000610adf8383614e68565b6000610adf8383614ea7565b6000610adf83836150a3565b61442481615a78565b82525050565b6000614434825190565b80845260209384019383018060005b8381101561446857815161445788826143c3565b975060208301925050600101614443565b509495945050505050565b600061447d825190565b808452602084019350836020820285016144978560200190565b8060005b858110156144cc57848403895281516144b485826143d7565b94506020830160209a909a019992505060010161449b565b5091979650505050505050565b60006144e3825190565b80845260209384019383018060005b8381101561446857815161450688826143e3565b9750602083019250506001016144f2565b6000614521825190565b8084526020840193508360208202850161453b8560200190565b8060005b858110156144cc578484038952815161455885826143f7565b94506020830160209a909a019992505060010161453f565b600061457a825190565b808452602084019350836020820285016145948560200190565b8060005b858110156144cc57848403895281516145b18582614403565b94506020830160209a909a0199925050600101614598565b60006145d3825190565b808452602084019350836020820285016145ed8560200190565b8060005b858110156144cc578484038952815161460a858261440f565b94506020830160209a909a01999250506001016145f1565b801515614424565b6000614634825190565b80845260208401935061464b818560208601615adf565b601f01601f19169290920192915050565b61442481615a89565b61442481615ab2565b61442481615abd565b61442481615ac8565b600061468a825190565b614698818560208601615adf565b9290920192915050565b601781526000602082017f4e6f64655365727669636554696d654e6f744d61746368000000000000000000815291505b5060200190565b601581526000602082017f46696c6550726f76654e6f7446696c654f776e65720000000000000000000000815291506146d2565b601c81526000602082017f536563746f7250726f766550726f66697453706c69744661696c656400000000815291506146d2565b602281526000602082017f536563746f7250726f76654368616c6c656e67654865696768744e6f744d617481527f6368000000000000000000000000000000000000000000000000000000000000602082015291505b5060400190565b601b81526000602082017f46696c6550726f7665536563746f72547970654e6f744d617463680000000000815291506146d2565b601081526000602082017f46696c6550726f76654578706972656400000000000000000000000000000000815291506146d2565b601981526000602082017f536563746f7250726f7665536563746f724e6f74457869737400000000000000815291506146d2565b601b81526000602082017f4e6f6465536563746f7250726f766564496e54696d654572726f720000000000815291506146d2565b601681526000602082017f536563746f7250726f7665436865636b4661696c656400000000000000000000815291506146d2565b601b81526000602082017f436865636b4e6f6465536563746f7250726f766564496e54696d650000000000815291506146d2565b601681526000602082017f50756e697368466f72536563746f72206661696c656400000000000000000000815291506146d2565b602e81526000602082017f496e697469616c697a61626c653a20636f6e747261637420697320616c72656181527f647920696e697469616c697a656400000000000000000000000000000000000060208201529150614797565b601281526000602082017f536563746f7250726f76654e6f74506c6f740000000000000000000000000000815291506146d2565b601481526000602082017f46696c6550726f766554696d6573457863656564000000000000000000000000815291506146d2565b601481526000602082017f46696c6550726f7665436865636b4661696c6564000000000000000000000000815291506146d2565b601681526000602082017f46696c6550726f7665436f70794e756d45786365656400000000000000000000815291506146d2565b600b81526000602082017f536563746f7250726f7665000000000000000000000000000000000000000000815291506146d2565b601681526000602082017f46696c6550726f7665416c726561647950726f76656400000000000000000000815291506146d2565b601981526000602082017f46696c6550726f76654e6f6465566f6c4e6f74456e6f75676800000000000000815291506146d2565b601581526000602082017f536563746f7250726f76654e6f74457870697265640000000000000000000000815291506146d2565b600981526000602082017f46696c6550726f76650000000000000000000000000000000000000000000000815291506146d2565b601081526000602082017f46696c6550726f76654e6f744e6f646500000000000000000000000000000000815291506146d2565b80516040830190614b7d84826153b7565b506020820151614b9060208501826153b7565b50505050565b805161032080845260009190840190614baf828261462a565b9150506020830151614bc4602086018261441b565b5060408301518482036040860152614bdc828261462a565b9150506060830151614bf160608601826153c3565b506080830151614c0460808601826153c3565b5060a0830151614c1760a08601826153c3565b5060c0830151614c2a60c08601826153c3565b5060e0830151614c3d60e08601826153c3565b50610100830151614c526101008601826153b1565b50610120830151614c676101208601826153c3565b50610140830151614c7c6101408601826153c3565b50610160830151848203610160860152614c96828261462a565b915050610180830151614cad6101808601826153c3565b506101a0830151614cc26101a08601826153b1565b506101c0830151614cd76101c0860182614622565b506101e0830151614cec6101e0860182614677565b50610200830151614d016102008601826153c3565b50610220830151848203610220860152614d1b828261442a565b915050610240830151848203610240860152614d37828261442a565b915050610260830151848203610260860152614d53828261462a565b915050610280830151614d6a61028086018261466e565b506102a0830151614d7f6102a0860182614622565b506102c0830151612e1c6102c0860182614f5e565b805160c080845260009190840190614dac828261462a565b91505060208301518482036020860152614dc6828261462a565b9150506040830151614ddb60408601826153b1565b506060830151614dee606086018261441b565b506080830151614e0160808601826153c3565b5060a0830151612e1c60a08601826153c3565b80516000906080840190614e28858261441b565b5060208301518482036020860152614e40828261462a565b9150506040830151614e5560408601826153c3565b506060830151612e1c60608601826153c3565b80516000906060840190614e7c85826153c3565b506020830151614e8f60208601826153c3565b506040830151848203604086015261287f828261462a565b80516000906040840190614ebb85826153c3565b506020830151848203602086015261287f8282614517565b805160009060e0840190614ee785826153c3565b506020830151614efa60208601826153c3565b506040830151614f0d60408601826153c3565b506060830151614f2060608601826153c3565b506080830151614f3360808601826153c3565b5060a0830151614f4660a086018261441b565b5060c083015184820360c086015261287f828261462a565b80516060830190614f6f84826153c3565b506020820151614f8260208501826153c3565b506040820151614b9060408501826153c3565b80516060830190614fa6848261441b565b506020820151614f8260208501826153b1565b8051606080845260009190840190614fd18282615109565b91505060208301518482036020860152614feb82826144d9565b9150506040830151848203604086015261287f82826151fc565b805160009060a084019061501985826153c3565b5060208301518482036020860152615031828261462a565b9150506040830151848203604086015261504b828261462a565b915050606083015184820360608601526150658282614473565b9150506080830151848203608086015261287f828261462a565b8051604083019061509084826153c3565b506020820151614b9060208501826153c3565b805160a0808452600091908401906150bb828261462a565b91505060208301516150d0602086018261441b565b5060408301516150e360408601826153c3565b5060608301516150f660608601826153b1565b506080830151612e1c6080860182614622565b805160009061018084019061511e858261441b565b50602083015161513160208601826153c3565b50604083015161514460408601826153c3565b50606083015161515760608601826153c3565b50608083015161516a608086018261466e565b5060a083015161517d60a08601826153b1565b5060c083015161519060c08601826153b1565b5060e08301516151a360e08601826153c3565b506101008301516151b86101008601826153c3565b506101208301516151cd6101208601826153c3565b506101408301516151e2610140860182614622565b5061016083015184820361016086015261287f8282614473565b805160009060c084019061521085826153c3565b50602083015161522360208601826153c3565b506040830151848203604086015261523b828261462a565b915050606083015184820360608601526152558282614473565b9150506080830151848203608086015261526f8282614570565b91505060a083015184820360a086015261287f828261462a565b80516040830190615090848261441b565b80516101608301906152ac84826153c3565b5060208201516152bf60208501826153c3565b5060408201516152d260408501826153c3565b5060608201516152e560608501826153c3565b5060808201516152f860808501826153c3565b5060a082015161530b60a08501826153c3565b5060c082015161531e60c08501826153c3565b5060e082015161533160e08501826153c3565b506101008201516153466101008501826153c3565b5061012082015161535b6101208501826153c3565b50610140820151614b906101408501826153c3565b805160009060a08401906153848582614f5e565b506020830151848203606086015261539c828261462a565b9150506040830151612e1c60808601826153c3565b80614424565b63ffffffff8116614424565b6001600160401b038116614424565b6000610adf8284614680565b600081610f8a565b60208101610f8a828461441b565b60408101615402828561441b565b610adf60208301846153b1565b60208082528101610adf81846145c9565b60208101610f8a8284614622565b60208082528101610adf818461462a565b60408082528101615450818561462a565b9050818103602083015261285681846145c9565b60608082528101615475818561462a565b9050610adf602083018461507f565b60408101615492828561465c565b8181036020830152612856818461462a565b61022081016154b3828a61465c565b6154c0602083018961465c565b81810360408301526154d28188614b96565b905081810360608301526154e68187614ed3565b905081810360808301526154fa81866150a3565b905081810360a083015261550e81856145c9565b905061551d60c083018461529a565b98975050505050505050565b60608101615537828661465c565b81810360208301526155498185614d94565b9050818103604083015261287f8184614b96565b60a0810161556b828761465c565b615578602083018661465c565b818103604083015261558a8185614b96565b905061287f6060830184615289565b608081016155a78287614665565b6155b460208301866153b1565b81810360408301526155c6818561462a565b905061287f606083018461441b565b60208101610f8a828461466e565b604080825281016155f3816148a2565b90508181036020830152610f8a816146a2565b60408082528101615616816148a2565b90508181036020830152610f8a8161483a565b60208082528101610f8a816148d6565b60208082528101610f8a8161490a565b6040808252810161565981614a34565b90508181036020830152610f8a8161470d565b6040808252810161567c81614a34565b90508181036020830152610f8a81614741565b6040808252810161569f81614a34565b90508181036020830152610f8a81614806565b604080825281016156c281614a34565b90508181036020830152610f8a8161486e565b604080825281016156e581614a34565b90508181036020830152610f8a81614964565b6040808252810161570881614a34565b90508181036020830152610f8a81614ad0565b6040808252810161572b81614b04565b90508181036020830152610adf818461462a565b6040808252810161574f81614b04565b90508181036020830152610f8a816146d9565b6040808252810161577281614b04565b90508181036020830152610f8a8161479e565b6040808252810161579581614b04565b90508181036020830152610f8a816147d2565b604080825281016157b881614b04565b90508181036020830152610f8a81614998565b604080825281016157db81614b04565b90508181036020830152610f8a816149cc565b604080825281016157fe81614b04565b90508181036020830152610f8a81614a00565b6040808252810161582181614b04565b90508181036020830152610f8a81614a68565b6040808252810161584481614b04565b90508181036020830152610f8a81614a9c565b6040808252810161586781614b04565b90508181036020830152610f8a81614b38565b60208082528101610adf8184614e14565b60208082528101610adf8184614ed3565b60608101610f8a8284614f95565b60208082528101610adf8184614fb9565b60208082528101610adf8184615005565b60208082528101610adf8184615109565b604080825281016158ee8185615109565b905081810360208301526128568184614b96565b60408101610f8a8284615289565b60208082528101610adf8184615370565b60208101610f8a82846153b1565b600061593a60405190565b90506159468282615b0b565b919050565b60006001600160401b0382111561596457615964615bb5565b5060209081020190565b60006001600160401b0382111561598757615987615bb5565b601f19601f83011660200192915050565b600082198211156159ab576159ab615b73565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b03038211156159ab576159ab615b73565b6001600160401b0391821691166000826159ff576159ff615b89565b500490565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615615a3b57615a3b615b73565b500290565b6000825b925082821015615a5657615a56615b73565b500390565b60006001600160401b03821691506001600160401b038316615a44565b60006001600160a01b038216610f8a565b6000610f8a82615a78565b8061594681615bce565b8061594681615bde565b8061594681615bee565b6000610f8a82615a94565b6000610f8a82615a9e565b6000610f8a82615aa8565b82818337506000910152565b60005b83811015615afa578181015183820152602001615ae2565b83811115614b905750506000910152565b601f19601f83011681018181106001600160401b0382111715615b3057615b30615bb5565b6040525050565b6000600019821415615b4b57615b4b615b73565b5060010190565b60006001600160401b03821691506001600160401b03821415615b4b57615b4b5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b50565b600b8110615bcb57615bcb615b9f565b60038110615bcb57615bcb615b9f565b60028110615bcb57615bcb615b9f565b615c0781615a78565b8114615bcb57600080fd5b801515615c07565b615c0781615a89565b60038110615bcb57600080fd5b60028110615bcb57600080fd5b80615c07565b63ffffffff8116615c07565b6001600160401b038116615c0756fe61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5a264697066735822122083e88f8bae526395c109fbe479234261ad4dc520dfb19a664ebf24fe2ffda87464736f6c63430008040033",
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
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((bytes,address,uint64,uint256,bool)[])
func (_Store *StoreCaller) GetProveDetailList(opts *bind.CallOpts, fileHash []byte) ([]ProveDetail, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetProveDetailList", fileHash)

	if err != nil {
		return *new([]ProveDetail), err
	}

	out0 := *abi.ConvertType(out[0], new([]ProveDetail)).(*[]ProveDetail)

	return out0, err

}

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((bytes,address,uint64,uint256,bool)[])
func (_Store *StoreSession) GetProveDetailList(fileHash []byte) ([]ProveDetail, error) {
	return _Store.Contract.GetProveDetailList(&_Store.CallOpts, fileHash)
}

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((bytes,address,uint64,uint256,bool)[])
func (_Store *StoreCallerSession) GetProveDetailList(fileHash []byte) ([]ProveDetail, error) {
	return _Store.Contract.GetProveDetailList(&_Store.CallOpts, fileHash)
}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreCaller) CheckSectorProve(opts *bind.CallOpts, sectorProve SectorProveParams, sectorInfo SectorInfo) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkSectorProve", sectorProve, sectorInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreSession) CheckSectorProve(sectorProve SectorProveParams, sectorInfo SectorInfo) (bool, error) {
	return _Store.Contract.CheckSectorProve(&_Store.CallOpts, sectorProve, sectorInfo)
}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreCallerSession) CheckSectorProve(sectorProve SectorProveParams, sectorInfo SectorInfo) (bool, error) {
	return _Store.Contract.CheckSectorProve(&_Store.CallOpts, sectorProve, sectorInfo)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) payable returns()
func (_Store *StoreTransactor) CheckNodeSectorProvedInTime(opts *bind.TransactOpts, sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CheckNodeSectorProvedInTime", sectorRef)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) payable returns()
func (_Store *StoreSession) CheckNodeSectorProvedInTime(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.CheckNodeSectorProvedInTime(&_Store.TransactOpts, sectorRef)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) payable returns()
func (_Store *StoreTransactorSession) CheckNodeSectorProvedInTime(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.CheckNodeSectorProvedInTime(&_Store.TransactOpts, sectorRef)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) payable returns()
func (_Store *StoreTransactor) DeleteProveDetails(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteProveDetails", fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) payable returns()
func (_Store *StoreSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) payable returns()
func (_Store *StoreTransactorSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreTransactor) FileProve(opts *bind.TransactOpts, fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileProve", fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreTransactorSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) payable returns()
func (_Store *StoreTransactor) SectorProve(opts *bind.TransactOpts, sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SectorProve", sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) payable returns()
func (_Store *StoreSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) payable returns()
func (_Store *StoreTransactorSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) payable returns()
func (_Store *StoreTransactor) SetLastPunishmentHeightForNode(opts *bind.TransactOpts, nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetLastPunishmentHeightForNode", nodeAddr, sectorId, height)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) payable returns()
func (_Store *StoreSession) SetLastPunishmentHeightForNode(nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastPunishmentHeightForNode(&_Store.TransactOpts, nodeAddr, sectorId, height)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) payable returns()
func (_Store *StoreTransactorSession) SetLastPunishmentHeightForNode(nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastPunishmentHeightForNode(&_Store.TransactOpts, nodeAddr, sectorId, height)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x181197f7.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (bytes,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreTransactor) UpdateProveDetailList(opts *bind.TransactOpts, fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailList", fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x181197f7.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (bytes,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x181197f7.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (bytes,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreTransactorSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) payable returns()
func (_Store *StoreTransactor) UpdateProveDetailMeta(opts *bind.TransactOpts, fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailMeta", fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) payable returns()
func (_Store *StoreSession) UpdateProveDetailMeta(fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailMeta(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) payable returns()
func (_Store *StoreTransactorSession) UpdateProveDetailMeta(fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailMeta(&_Store.TransactOpts, fileHash, details)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2561e0a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig, address _proveExtra) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig, _proveExtra common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _fs, _node, _pdp, _sector, proveConfig, _proveExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2561e0a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig, address _proveExtra) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig, _proveExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector, proveConfig, _proveExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2561e0a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig, address _proveExtra) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig, _proveExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector, proveConfig, _proveExtra)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x0e3459fd.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreTransactor) ProfitSplitForSector(opts *bind.TransactOpts, sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "profitSplitForSector", sectorInfo, nodeInfo, setting)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x0e3459fd.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreSession) ProfitSplitForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.ProfitSplitForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x0e3459fd.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreTransactorSession) ProfitSplitForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.ProfitSplitForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xa0004412.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns()
func (_Store *StoreTransactor) PunishForSector(opts *bind.TransactOpts, sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "punishForSector", sectorInfo, nodeInfo, setting, times)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xa0004412.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns()
func (_Store *StoreSession) PunishForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.Contract.PunishForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting, times)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xa0004412.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns()
func (_Store *StoreTransactorSession) PunishForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.Contract.PunishForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting, times)
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

// StoreDnsErrorIterator is returned from FilterDnsError and is used to iterate over the raw logs and unpacked data for DnsError events raised by the Store contract.
type StoreDnsErrorIterator struct {
	Event *StoreDnsError // Event containing the contract specifics and raw log

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
func (it *StoreDnsErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDnsError)
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
		it.Event = new(StoreDnsError)
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
func (it *StoreDnsErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDnsErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDnsError represents a DnsError event raised by the Store contract.
type StoreDnsError struct {
	Method string
	Msg    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDnsError is a free log retrieval operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) FilterDnsError(opts *bind.FilterOpts) (*StoreDnsErrorIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DnsError")
	if err != nil {
		return nil, err
	}
	return &StoreDnsErrorIterator{contract: _Store.contract, event: "DnsError", logs: logs, sub: sub}, nil
}

// WatchDnsError is a free log subscription operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) WatchDnsError(opts *bind.WatchOpts, sink chan<- *StoreDnsError) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DnsError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDnsError)
				if err := _Store.contract.UnpackLog(event, "DnsError", log); err != nil {
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

// ParseDnsError is a log parse operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) ParseDnsError(log types.Log) (*StoreDnsError, error) {
	event := new(StoreDnsError)
	if err := _Store.contract.UnpackLog(event, "DnsError", log); err != nil {
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

// StorePDPVerifyEventIterator is returned from FilterPDPVerifyEvent and is used to iterate over the raw logs and unpacked data for PDPVerifyEvent events raised by the Store contract.
type StorePDPVerifyEventIterator struct {
	Event *StorePDPVerifyEvent // Event containing the contract specifics and raw log

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
func (it *StorePDPVerifyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePDPVerifyEvent)
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
		it.Event = new(StorePDPVerifyEvent)
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
func (it *StorePDPVerifyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePDPVerifyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePDPVerifyEvent represents a PDPVerifyEvent event raised by the Store contract.
type StorePDPVerifyEvent struct {
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPDPVerifyEvent is a free log retrieval operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) FilterPDPVerifyEvent(opts *bind.FilterOpts) (*StorePDPVerifyEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "PDPVerifyEvent")
	if err != nil {
		return nil, err
	}
	return &StorePDPVerifyEventIterator{contract: _Store.contract, event: "PDPVerifyEvent", logs: logs, sub: sub}, nil
}

// WatchPDPVerifyEvent is a free log subscription operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) WatchPDPVerifyEvent(opts *bind.WatchOpts, sink chan<- *StorePDPVerifyEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "PDPVerifyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePDPVerifyEvent)
				if err := _Store.contract.UnpackLog(event, "PDPVerifyEvent", log); err != nil {
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

// ParsePDPVerifyEvent is a log parse operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) ParsePDPVerifyEvent(log types.Log) (*StorePDPVerifyEvent, error) {
	event := new(StorePDPVerifyEvent)
	if err := _Store.contract.UnpackLog(event, "PDPVerifyEvent", log); err != nil {
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
