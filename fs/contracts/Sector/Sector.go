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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"AddFileToSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"AddSectorRefForFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"CreateSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"DeleteFileFromSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"DeleteSecotr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"GetSectorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetSectorsForNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sector\",\"type\":\"tuple\"}],\"name\":\"UpdateSectorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x29d6276d\",\"type\":\"bytes32\"}],\"name\":\"c_0x29d6276d\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_FILE_INFO_GROUP_MAX_LEN\",\"type\":\"uint64\"}],\"internalType\":\"structSectorConfig\",\"name\":\"sectorConfig\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506186e280620000226000396000f3fe608060405234801561001057600080fd5b50600436106100bd5760003560e01c80632ba010d711610076578063ba9210041161005b578063ba9210041461019a578063dcf74960146101b6578063e3168f9e146101d2576100bd565b80632ba010d71461014e578063955f98b71461017e576100bd565b8063112346c2116100a7578063112346c2146100fa5780632384a6aa1461011657806323de5b9814610132576100bd565b806247c003146100c25780630daf9945146100de575b600080fd5b6100dc60048036038101906100d791906176b7565b610202565b005b6100f860048036038101906100f391906175e8565b6106c7565b005b610114600480360381019061010f9190617611565b6106ca565b005b610130600480360381019061012b9190617676565b6108fa565b005b61014c60048036038101906101479190617723565b610f1b565b005b61016860048036038101906101639190617723565b6111ed565b6040516101759190617edc565b60405180910390f35b610198600480360381019061019391906176b7565b611a38565b005b6101b460048036038101906101af9190617676565b611f3e565b005b6101d060048036038101906101cb9190617676565b612aa8565b005b6101ec60048036038101906101e79190617596565b612ccb565b6040516101f99190617d02565b60405180910390f35b61022e7f6946773d589d5aa7cca324353bfe08904598f9bf404adab26232f8b5cf21005d60001b6106c7565b61025a7fbd7447020a27aaa6a74daf5cc19e46e79b45a399121a83937226d9a57866d0ab60001b6106c7565b6102867fdc604c12c1261a0de72ae24e49a6c9faab1b4a289bc7e67f6a7810b386e4c65160001b6106c7565b600061029f8360000151846020015184600001516130ff565b90506102cd7ff29e3d9ca8bfba6ddd1ae5d1618ab99fd0d77f3cd0b69c5c5df87ec61ba4844360001b6106c7565b826101000180518091906102e09061829a565b67ffffffffffffffff1667ffffffffffffffff16815250506103247f3c15cea3d76d44a626f8e25d095890a230a4e52587f1346063f200db9186eadb60001b6106c7565b6103507fadce9fcba52bc14487ce82b6a50f02a59543c5ed78da5791d7d22a268a99c4ba60001b6106c7565b81608001518360e001818151610366919061813e565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506103af7f7eecdbd0a4369e120c0e1e5af3b381580c1e892fe353bfe3cf979c72ad5c87d660001b6106c7565b6103db7fb92fa28e211d3aee5046d8a76d1db546278d00f28abf198994077fedae88edc360001b6106c7565b8160a0015182608001516103ef91906180c8565b83606001818151610400919061813e565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506104497f35e5d832fb326b1c73cdff9fbbf3cac0aa4b2f3455378c524bfdad1cbd48da9960001b6106c7565b6104757fa86f16bcb3c15d10543a74ab7211410c3da7e2d3bbeba8a051e076fc1c85953260001b6106c7565b8015610503576104a77f85746e64c7ecf3041cbff4e0572b3d42195be5cfeba9547cb0e3c4cbf97ac46a60001b6106c7565b6104d37f1591fa4847409025f3a857f5b86285e99331efa817ef14549ee6ec9f75791b7b60001b6106c7565b826101200180518091906104e69061829a565b67ffffffffffffffff1667ffffffffffffffff1681525050610530565b61052f7fb4da21644e4a7fc90068443205ac764cdb35c12805d0f686614a9e0a69d52f0c60001b6106c7565b5b61055c7fdc2ad75a81dab5ec57ede91a9f94c16aafebe920cd5e5a2133de6493aa80cc2060001b6106c7565b6105887f6e4803494ed94c785344d3c421414130c5f641b6ed361f7212b9b3e9b658e23e60001b6106c7565b600083610100015167ffffffffffffffff161415610634576105cc7f2b83974a4c538480f01ec56a74b005a4c417f25f988c43e3169ab586ca3ad53260001b6106c7565b6105f87fde2148f3fde82993d48bd2da757673252b00ab21760ef165950452197b7aa9ce60001b6106c7565b6106247f16debf4c3c03075887844a037302a98adb003be1a8c34864aaf48f663656fbcd60001b6106c7565b60008360c0018181525050610661565b6106607fcd327c750cfc21eb7cfdafb1e430ca04c369bba3986aa0ce707e8f5f77a520cf60001b6106c7565b5b61068d7f1131f4db2427d7176c491d7f5238319e4d4cfe6c3e97bc68a0c3aa2c2162f71760001b6106c7565b6106b97f767670dfc055717e0e815cf3af7f2e03b74bc72f6bec0659520175fa10bc3a1b60001b6106c7565b6106c2836108fa565b505050565b50565b600060019054906101000a900460ff166106f25760008054906101000a900460ff16156106fb565b6106fa613895565b5b61073a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073190617e7c565b60405180910390fd5b60008060019054906101000a900460ff16159050801561078a576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6107b67f26d288c13b61475b0bc42638f2e5095d509caf4da2b30373965164c1c656b01b60001b6106c7565b6107e27fa9b32b63360a0ea16f25995ef904699b071a7a5af893d0fe410792ebe980c8bd60001b6106c7565b61080e7fa091b05eed7f65636ca2cdbf50db950beb8be4f178047e7b23a211e32f03fddf60001b6106c7565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061087b7f45491887f3a4ac8f9f1a716a8b97eedcb4e19e9ac7bc0bbc58f3b0a88135c4ca60001b6106c7565b6108a77fb1ba2fba7a788206300a9ed48d2311a7358279a16bb540074afda3b37276c6c360001b6106c7565b8160000151600060166101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080156108f55760008060016101000a81548160ff0219169083151502179055505b505050565b6109267fbab33c829c9996924ede78ce61c6eff32465be5989392bda441ff134786c74fe60001b6106c7565b6109527f3cf02417cc53ee01c167db8b94f9b63822e8a29a93f29b37a4c1a8774f8c32c760001b6106c7565b61097e7f8a907d1912286be9ce0a50144367f9cd82078da8356afd97c6d1fc17dcb7a47c60001b6106c7565b600060016000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506109f17ff5e36d18e978c23ab4709c319958d0ea692a04ee2a36be6413b999490ccec4f360001b6106c7565b610a1d7fb250b34256aad002f500055491df26039a71a69d7eb891182f6e47bd21ef795960001b6106c7565b60005b81805490508167ffffffffffffffff161015610e6b57610a627f44631d57c04619ffdc78464ae9410f4633b14817b2955503787654c50be04de760001b6106c7565b610a8e7f839dbb50d8d37a5cacb3575f76247cec8339cf683a40d8e7b9be3536acf18a8e60001b6106c7565b826020015167ffffffffffffffff16828267ffffffffffffffff1681548110610ae0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff161415610e2c57610b427f0b20b4459e1e710bdb6ccdf7bddfe080d20d29fb466e929c552bc3be884883c260001b6106c7565b610b6e7fbc666fce9df3eb8e9b6cb2e77bbed4c996a8251b175c78a68b8a4eac0dc8ae0f60001b6106c7565b610b9a7feac775a9cf16c38d8b9c629eae11acdcc5bf10312e56348a1b8d6632480f719b60001b6106c7565b82828267ffffffffffffffff1681548110610bde577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115610d11577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190610df7929190616568565b50905050610e277fcac248084862bf4c8841d11ef4dbbb8cbdaa54b53bf3593eb84b192ee90939c760001b6106c7565b610e6b565b610e587f70da512a885a081f1e609180c719bad9a2b3484c95a934258318203d49b296df60001b6106c7565b8080610e6390618370565b915050610a20565b50610e987fd89b77b4e91411fc228e9065a26a61da092b561f169d340e325e63fdac7779ee60001b6106c7565b610ec47f96bb39363f6c41b477fc2ffacfb9652fab0266dafaf62a323c02508f13b5024860001b6106c7565b8060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908054610f169291906165c8565b505050565b610f477fba7a59354f295e9cfa29b0c61ba5b57ad255aefda16543c531b2e97b07b184d760001b6106c7565b610f737f76eba042b8cc59bbf45380851c350b1b02b972f5bd022be8d2d8ac0b2711663260001b6106c7565b610f9f7f2bffca3adb168293b23513ff9b15cf7278722e6f871e7271d645d6ab79827f5760001b6106c7565b6000610faa826111ed565b9050610fd87ffd15fbd409606a0e1426f1a4e35ef9918c7d64e1d15d42239b4c12500b92458b60001b6106c7565b6110047f45b11e98e533070af462a3c15704ff945089d4ea678a722397bdfe3ff50e95a060001b6106c7565b600081610100015167ffffffffffffffff1611156110b9576110487f7d28183d6f64ec80db9b3e600be71e64adf9ee58cce72cc45960158ce6e3145160001b6106c7565b6110747fc345614eae689e5a50172d2a2e1b691d70be2c5b8fefb9136671dd325a11e5e560001b6106c7565b60008161010001516040517f25a56d930000000000000000000000000000000000000000000000000000000081526004016110b0929190617dd8565b60405180910390fd5b6110e57f7d7d8b950be65c0724bee103a4e03a57fdb35b112631bde51e202d3bf41dedd260001b6106c7565b6111117fb4f0b5d974531722ff35261c80b14503c988029d42bb0ec4775c6b9cbc810aab60001b6106c7565b61113d7fa82c541da1ad66baabcd968f8cf906b932aacd7497cd44a83b33d59b1d49f0a660001b6106c7565b61114f826000015183602001516138a6565b61117b7fdd64fb8c4bb9f7e7da161a472f4cbc6038801e52448233ffafc903b1dddc3f6660001b6106c7565b6111a77f8a27c9efbc617dd2395610a1be5a73ba4b96d5367ae05f93fcbfdbfff13d6e4660001b6106c7565b7f4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c6009433385602001516040516111e19493929190617d24565b60405180910390a15050565b6111f56168d4565b6112217fb56014fc991d774b2cc23e58fd447758850c3929af6c8480dbc1af34d8d1b37860001b6106c7565b61124d7ffc6ef0bd0a104f1c184c325d5dda5fce31c3516710e10c5bb71f63fc228e707260001b6106c7565b6112797fba0414a5c4e0ea6a14ff9265a7ebada39bebe8bcdfe9f186f1a77afddcb9b01960001b6106c7565b6112a57fc13739bf4dd9b9b36af6c777b4cf51326ffc2f7e9ad39b5a561b8a7ac2ad83f460001b6106c7565b6000826020015167ffffffffffffffff16116112f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ed90617ebc565b60405180910390fd5b6113227fcb1eff8dede3e0033c17523fe739b0930827bb116127905dda931ab0b3672fc760001b6106c7565b61134e7fe0db4d86d4f13ea560949911caaa6d9449fc19a1d27d4f0842f9839a28c196b460001b6106c7565b61137a7f8afc9266de811521ab58cc1ee7be13c88cdcf47f7947611e1e043b34d0b2fe0c60001b6106c7565b600060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156117235783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff16600281111561153a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611572577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b8282101561170c57838290600052602060002001805461167f906182c4565b80601f01602080910402602001604051908101604052809291908181526020018280546116ab906182c4565b80156116f85780601f106116cd576101008083540402835291602001916116f8565b820191906000526020600020905b8154815290600101906020018083116116db57829003601f168201915b505050505081526020019060010190611660565b5050505081525050815260200190600101906113df565b5050505090506117557f233a373e493c8bbff8c4b852481ec86d4eebac44dbce5395a9f0a685a42f51b560001b6106c7565b6117817fed075eb98a1cdadf955c3e96d4be67fdc9854ce85b585c2854d3ff7d934e101760001b6106c7565b60005b81518167ffffffffffffffff161015611974576117c37f9ab838b6a611c18b5ffa6dbe763f199bd2a3b50e93b5c133c815046d7e446b8e60001b6106c7565b6117ef7f403ada0f1b9922170593c52135bdadf9682b0b26d54cca69aa9a228da1c8292560001b6106c7565b836020015167ffffffffffffffff16828267ffffffffffffffff1681518110611841577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff161415611935576118897f0b224ee8641c55ed20be4fcd221366546700502750a93494f38bc1ec3ed7d57d60001b6106c7565b6118b57fa099831781b28930643d93eddec0e9545d2e564016aa4d9c90617c98f609c11c60001b6106c7565b6118e17f3516db01dc1927b78206ef086546ecd86cb7ddf44a2f683a27032aa74a100c4260001b6106c7565b818167ffffffffffffffff1681518110611924577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015192505050611a33565b6119617f18d71158e4bed03f6767a2d50d4c410e215171650fdb18f55106dd5dd8fd177860001b6106c7565b808061196c90618370565b915050611784565b506119a17fc347d3e28813d8ced8f06470153ba5f580c36fef40fec9a2b801bcb97488f90d60001b6106c7565b6119cd7f769e816c7ed60337ba352d71a5687c83cfb1c6e57395bf8d2743bcfa6d51aac660001b6106c7565b6119d56168d4565b611a017fb76c9edc5325d2190f8b415bd94be79f3f567ee95cb84260e46b611e6716e7d760001b6106c7565b611a2d7fe952608570efbd273b78ef5bb2da16e4b6f32f453564fe98a0c8b09ff7e724cd60001b6106c7565b80925050505b919050565b611a647f768656c0cdd52722593e2115406d54117e9bd6111cb52089f5b59c7fcf122d8c60001b6106c7565b611a907f1e9ac4df7ae1fcd83ebfd2f052742b296ac1ab9867bf03397118a92714eff0ac60001b6106c7565b611abc7f181ff6d3dbb564b62c5a10cb9f6c8d01addcf825a82c46c6dedb4efb2a3285fd60001b6106c7565b816040015167ffffffffffffffff168160a001518260800151611adf91906180c8565b8360600151611aee919061808a565b67ffffffffffffffff161115611b8857611b2a7f47ecb464575db4ad742bf4bcb665850697e92f6111a7c68a23dafd8899ef356760001b6106c7565b611b567fb9ffaf145d63153401605995fc5aa7ae5ed9e3f4de53bdb96e85835b2a2ed9f160001b6106c7565b6040517f6073072a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bb47f807ad8a78f0dfa57aae627e207cef03254721efb201a332afcdfb7227eea6a1460001b6106c7565b611be07ff543cf3b05be29618a0632c7ee129e0ef38370ce135832acf25b5205b981f86060001b6106c7565b611c0c7fd638fa618cf0ed1094306a24de3aa59f34f5e5109233ca8b8bb4b0cecb8e674b60001b6106c7565b6000611c4783600001518460200151604051806040016040528086600001518152602001866080015167ffffffffffffffff1681525061399a565b9050611c757fb723d1ddc452a736b1018e9a5177876d92ab5888472b31d37d9032f0651b5e7c60001b6106c7565b82610100018051809190611c8890618370565b67ffffffffffffffff1667ffffffffffffffff1681525050611ccc7f2625bcc04197a9eec35125c3d604c539936c5777fd7d3b6d546aecfc21314bb860001b6106c7565b611cf87fb12168ace3033c5c71c6c284837441a07034250f27324108d98fc815e60ce6c060001b6106c7565b8160a001518260800151611d0c91906180c8565b83606001818151611d1d919061808a565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611d667f6fc6e948d56023414ed5ee082fc755bc9b267318aa22a4dc9a2df0a2a548267060001b6106c7565b611d927f4e5aa2f42aad718572b8e312ef3cc8d501cc37fae1ec6ed8e58136f3d7972b5a60001b6106c7565b81608001518360e001818151611da8919061808a565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611df17f81040a56e8e7fba8b618c700914350ce80252f33721d471405c0c16ebf9a9b9760001b6106c7565b611e1d7fb67f616f7083384b537f72c851b990da40e2cdcd06a99d4b82c6f3776f5b578f60001b6106c7565b8015611eab57611e4f7fc942034943bbd5a425fd0240e27b83bfe331c5bcf2c2d77d23fba6b2c6ebbd0060001b6106c7565b611e7b7fc61f44af1f16d42dba6080b97cd9006c9b38435a832384f77c1900e4e74d315360001b6106c7565b82610120018051809190611e8e90618370565b67ffffffffffffffff1667ffffffffffffffff1681525050611ed8565b611ed77f115e7c0c4d3f6c2f36cc31e01e2d562b29eb22c2c23e76712eea934b3bc5ee2a60001b6106c7565b5b611f047f122630e1e79f3f1a744d6538736e44ca8399d3244b807bfe3d252b0fd74f4c8560001b6106c7565b611f307f0f19a9022ff6695c306985e4add75dbc8c230adcd36a662929533008bcd953ff60001b6106c7565b611f39836108fa565b505050565b611f6a7f43ffcbe6a729df46768767367a80c4b908fcd4157cb1b83d66423efad201315960001b6106c7565b611f967f54f198330c96315d9e50278bc21d317232fd6d02ff2426c905da3aa98c6ae48c60001b6106c7565b611fc27f9a611bba6aa85df2c63a02ef84c7c91867ef8f1ed1b6073edbcfa185ed3303fa60001b6106c7565b611fee7f80278e146f18bb776a8ec8db30fc5dbef5310fd89b47de3ba50de7d6c6e9823a60001b6106c7565b6000816020015167ffffffffffffffff161161203f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161203690617e9c565b60405180910390fd5b61206b7f8d2daa5c42dd47c202f87ed942d98b83d07f0e3b6b739d9a46141edbc523455960001b6106c7565b6120977f9d241acda3cb91ae9aa77b6fd9724997960bf2da713992d5476de0f23a25184760001b6106c7565b6120c37fd3ad23d4cda7cf17b4f7334a7fbf9bf35ce564c19cf5eda4c1c19ab49b79ab7060001b6106c7565b6120ef7f3e46baf3a1ae6f71ccd99daa20ca4ad7f96ad7e1a8c95b4b131ad62629234d3960001b6106c7565b6000816040015167ffffffffffffffff1611612140576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161213790617e5c565b60405180910390fd5b61216c7f3f65c84b03457e3ee1b8dc7b5969754f6a319e2179f7ab845c072fea305460de60001b6106c7565b6121987f970226277e5ad7ef90bf36ca69a65f599ff0b87476b2dfc4cba29e320a88813760001b6106c7565b6121c47f1f3483dd07c2df7235ce7ecd7e5b4205ee84768a6a4e43c80f68447c06695d7a60001b6106c7565b6121f07f6cda95a4c4f4d59f4e50e64b06724328d8affaf6e59f047abcc94c7209f0b0e560001b6106c7565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636683899482600001516040518263ffffffff1660e01b815260040161224f9190617ce7565b60206040518083038186803b15801561226757600080fd5b505afa15801561227b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061229f91906175bf565b6122de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122d590617e3c565b60405180910390fd5b61230a7f2088cb317fe5d85b09ecf939c847b73b95f343455a3409bda7acff0cc6dbf6b460001b6106c7565b6123367f32bd0f72df208fe527b9fb07b8e5ff5e5e189dbd14e6e93e8df0c466bc9ab9b960001b6106c7565b6123627f3bd744b4f0cb9f34cfa56aede6d36cb7ccd73d6b8c4303265c16e3d1b286ceb760001b6106c7565b61238e7fe0775fd8d8f2430bbf141665d8bdaaf4e2282ca8526e45620de11517655276ec60001b6106c7565b60006123d56040518060400160405280846000015173ffffffffffffffffffffffffffffffffffffffff168152602001846020015167ffffffffffffffff168152506111ed565b6040015167ffffffffffffffff1614612423576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161241a90617e1c565b60405180910390fd5b61244f7f16af40ccdacb55322dca2dfc7537677836823e85a5a78b634e17859175ecbd0160001b6106c7565b61247b7feacc2a0f3b1bfdb218dbb55930b9dc2227bad9a604dac8761439659ede4ee4e960001b6106c7565b6124a77f84931eac100ad095d3515b0a80e8ec111c4b87c84bc0b39f16b8fe10880f0e7360001b6106c7565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe83600001516040518263ffffffff1660e01b81526004016125079190617ce7565b60e06040518083038186803b15801561251f57600080fd5b505afa158015612533573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612557919061764d565b90506125857f012889bbce1ad80536f58d262d9fbd3197f08942d5d27e747c483d9134156dcb60001b6106c7565b6125b17fdecc667777e173a614110ddde1a5ea32de1bda362cf603b28b4e8f0b5f59bb6e60001b6106c7565b60006125c083600001516142f5565b90506125ee7fc12256b0530319e6b3fd6b261a0ee830fd4475a355b853d707ca37d8588d0bb960001b6106c7565b61261a7fa7ee6454538e2edb5bce610ab1bbfb09c60d36b1968c7d0e1acf34c4b82a621260001b6106c7565b816040015167ffffffffffffffff16818460400151612639919061808a565b67ffffffffffffffff1611156126f3576126757f89fb0b169176bcf50d158e9fe0a442418fd29e68740803fdbf1df27a20a3a0e660001b6106c7565b6126a17f1ae08fc1ec360e382dcaa2dde8e822a19af148080ff09e6306d30f6f952e461b60001b6106c7565b81604001518184604001516126b6919061808a565b6040517f95c91e8e0000000000000000000000000000000000000000000000000000000081526004016126ea929190617efe565b60405180910390fd5b61271f7fcc8eb0623751ec5f6e95f7bd3dd85551d6efdc1d3f556fdb3e1595827968b3b960001b6106c7565b61274b7f8a15f61858ac420399bdb255083147d126ae45b7f29c167d2a50d7ea9f32779360001b6106c7565b6127777f149ba3cec578af3856f595695512ed2d7881924cb6aba748f0f2e31d2cf7544560001b6106c7565b60016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115612909577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff0219169083151502179055506101608201518160050190805190602001906129ef929190616568565b505050612a1e7f1d17502b1a8e7db7980e0ce9f81dbe31cf7232d861ed11e658227052a22e2bf660001b6106c7565b612a4a7f2ecbbe89e21fa4116c5b324a39baedc78ed286d500201e6f2e022a7b0fbd7d3260001b6106c7565b7fc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb6008438560000151866020015187608001518860400151896101400151604051612a9b9796959493929190617d69565b60405180910390a1505050565b612ad47f225f8600a660f39f2fdf5d7a5be49ea24132cea3eb60722a4f9405ed9d42e72d60001b6106c7565b612b007f246c11e19d66495175f6594cfa67fbbed20ba1426f444c75c35dd9d9d171100860001b6106c7565b612b2c7f623d477333391236a147d2a22cbe1504ac9f0865e097f4a5a1a6241dd29207a060001b6106c7565b6000612b4082600001518360200151614904565b9050612b6e7fb3c0d0b50c9a48e866b9b9adecbe67693f6e105686f6ab410015f7eeee4b717d60001b6106c7565b612b9a7f27749044d2a0cfe7e6099c7adf57f43e349db5dd844c126c892396e3a4cd539160001b6106c7565b80612c3557612bcb7fee0e521f14a2b1771fa341230360cc39d92fde6eecdf467106a80e8507d8e79960001b6106c7565b612bf77f8611527167e4ed1332de539005826d6793e986f27a0227006249d2bf30cfdf3460001b6106c7565b60036040517fc8c84b2f000000000000000000000000000000000000000000000000000000008152600401612c2c9190617e01565b60405180910390fd5b612c617f9deefab10afb6c63bcec0656a339a91552ed41fc33b7a7d79da08fdf7eeff84860001b6106c7565b612c8d7f0ec86240bfcd64c89bd79f242a6f4e5218bce65717193de26ed518b9c9808a3360001b6106c7565b612cb97fafbd1daa00f23bf67920d1e31f2d6252a4095bf5e3517ebaf3781850d5d5007860001b6106c7565b612cc7826000015183614c61565b5050565b6060612cf97f6b30fed03acd5fc7882a50c740716fa1c5b3d2a1a221414a2ff1045f16f4de5260001b6106c7565b612d257f5346d600527ad7dfbf0c8cf4941ab1b02e05b852c8843aca86e9f11db9af635c60001b6106c7565b612d517fe423cb92a85981f3b7008291bf62a326050f24cca06633dbb404ddb566f5944c60001b6106c7565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156130f45783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff166002811115612f0b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115612f43577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156130dd578382906000526020600020018054613050906182c4565b80601f016020809104026020016040519081016040528092919081815260200182805461307c906182c4565b80156130c95780601f1061309e576101008083540402835291602001916130c9565b820191906000526020600020905b8154815290600101906020018083116130ac57829003601f168201915b505050505081526020019060010190613031565b505050508152505081526020019060010190612db0565b505050509050919050565b600061312d7fdac88d5ce47849788b6c5e32a9726d7bbb1bb2bdd7b02ffce7acbcb87c58273660001b6106c7565b6131597f07c61050d6945fdd7819b6b33377d6f2f48759bb97dc0b6939dc46afeff2731b60001b6106c7565b6131857fc8f78c84b097ad198920893d4f53a0509a165b625f1de787383029a24a336dd760001b6106c7565b60006131918585614f60565b90506131bf7fa66dcf269ff77c63e90566a75c57469f4aa9290518e1367f377b7accdd816b3e60001b6106c7565b6131eb7fb6792e31b3fdb9c1414e86aed3de6b1fed6320922e5883bee4c5ccbb29931c0760001b6106c7565b60005b8167ffffffffffffffff168167ffffffffffffffff16101561382f576132367fbd39394156306901aa9af11890ee930e0b7af308e791712bab7040f116e82a1460001b6106c7565b6132627ff7e99ea508feb050a2d450ad40327b58de7a6736e52bc1efd9b047691baf7c9460001b6106c7565b600061326f87878461508e565b905061329d7ff59bb7618b53e585c1431d1503ad6a9f90c32a6f7b95d8c0c6e2bf8e2eaee0af60001b6106c7565b6132c97fbbd546b4ca8e58c5c7a3fbe8dd8c8a0cc6b6087073449211b8db38f267cb903a60001b6106c7565b6000806132d68388615323565b915091506133067f07bba7a72522c90b12fdf465cbeb5fc960c2a77a154f8da56bc58124c3872c6d60001b6106c7565b6133327f9c41522d6e0997cf823bb65b5068cad62ad2a196b5f2f460e03c0341bf7650c360001b6106c7565b80156137ed576133647f87d5ec1a5a29c49325e018de65f6ab86b5cc57cbe2d7f5095c17b869e55ae28e60001b6106c7565b6133907fca85af5125cb6ecbbb5e9f654dd6f914cc8b4675d90cffa38644c1fc5378c0ae60001b6106c7565b6133bc7fbb2bd7f05fd2640aaf73d5c233471bf2ea0a4dcb12e2c9600f045aef9f14b88760001b6106c7565b600060036000856000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002080549050905061341b7f50ab706e4617d5eb8de08acac13511755a0a4d423c30d93dc34f0d79a984bed360001b6106c7565b60036000856000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208367ffffffffffffffff1681548110613488577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600080820160006134a791906169c1565b6001820160006101000a81549067ffffffffffffffff021916905550506134f07f78e1f923dbe0b11b5658b4e63c8144af4dd8cf751e7db8bed0135fddd9a2e03b60001b6106c7565b61351c7febc00f45981d674a3743bcc5ff5ee7829eadc2f8a8c7359553b1cb7c1dfd2b1360001b6106c7565b60008367ffffffffffffffff1614156135c15761355b7ffd1a095e5351f4e3a841003648de95ec2d1eefb52b9cf0c6a83eb3610beba66560001b6106c7565b6135877fa691e72f136973af49f4fa51c0a0d977e8f13f20520a051c8b1716d670d274cd60001b6106c7565b6135b37ffb04c4e0fb6cfd659e95c7e4f22042013a5edd611bd39d00e4a9643c57b1eb5b60001b6106c7565b8784602001819052506135ee565b6135ed7ffbbe3920bca96bdaba2e9b4340c8b073658422616e91f1c3280d61119763645560001b6106c7565b5b61361a7f353f728d9f480136a7edc1ebc5ea94a94e798ed7901b33bb287288478b61a34260001b6106c7565b6136467fd3f76268d24c345e12a1a6b663eb884ca24bb09a3575bafb82006844689d5bd860001b6106c7565b600181613653919061810a565b8367ffffffffffffffff1614156136f6576136907f488ac80d0fb99a19fc961d64080000ef3c14efef2a4954dc1d720abec7fe8f2560001b6106c7565b6136bc7fb166b42de903af6b6be0fd46151b2cb6aab7987eb207f914876c76b8a0c5cec660001b6106c7565b6136e87f3753de9ea28d2a57d29b8ff40c266203361496f6656f4cd6d359bbc094e068b860001b6106c7565b878460400181905250613723565b6137227f9dc0f72e827d527e04de4e3ac16fcd4b5b72e84be469493ec1b5aa829f19dfb860001b6106c7565b5b61374f7f706b559ce5426ea5ae75383382122ba5ef34f644ba6193b7a109cd99685c007d60001b6106c7565b61377b7ff1cde91f775dc9e5a1bb46ee8ebc06139f53f481c404c24a92ea760599f547b660001b6106c7565b6137868a8a866159a4565b6137b27f130e29db6e85ef26d23c5b8526c9744cfe3bb7632005d33ac534bac9c1e64f4c60001b6106c7565b6137de7ff69a8352b6aa4db37bb1fe79c04a1d9e064a88c6eeb7c37f5da33d8a797640b860001b6106c7565b6001965050505050505061388e565b6138197f50ef4904e6cad02dc79e14eddba9e551fdd97615f54b7cc4f80cbafaddd3c08c60001b6106c7565b505050808061382790618370565b9150506131ee565b5061385c7f99b7c52368369cd7325b91636fd7b6a6603f5c8e22948511835a1c08e43cb70960001b6106c7565b6138887f70ecbcb3545d17c3807009eb8d4015025da94cfa7c3edee1cbedde8c5b7ed9d460001b6106c7565b60009150505b9392505050565b60006138a030615b3d565b15905090565b6138d27f7d35758dd7f5de269b84df640b5d49193676dbbe59480dfe2dfdc3deedafd84060001b6106c7565b6138fe7fc0fee63bafa13d2cec230f0cee4242af64f498a635a5d8e49c18f6d97f470f8760001b6106c7565b61392a7f76758cf24ba199c7b6c6b3ff9d9e27bd6065fe96e8acdadda6c7bf33c007114860001b6106c7565b6139348282615b50565b6139607f442ccbb45816882a67baea0189ad130d085424e98d65d89f1df41af5f8c8552c60001b6106c7565b61398c7fe67d3e4d917d418322c088123a49f7857874a09d1c5a6afb5c47d9b27060525660001b6106c7565b6139968282615cd5565b5050565b60006139c87f1e153624cad8a2723628b56bdc82f8daef69e1a60e687d985213cc7c9520fb9260001b6106c7565b6139f47f97b2abc29a4bcfcb4d242ec6e1d5a036cf36eaebc26080a7e06006c23dfab34e60001b6106c7565b613a207f8a57d9b8646c215eb9082de2f87984a5295b3186e8e3adda4180f1fb6785713e60001b6106c7565b613a28616a01565b613a547f8dba24be634a7a004c37a5e64c061e492b231517f3b3d66bc95bf4e2279ef9a360001b6106c7565b613a807f89c72acccf5209b7891d0356a88cf55247158eeb57edb95f208bbedd2692a71d60001b6106c7565b6000613aae7f41b6857935051e34bfe024152599187f1da412f53bb89b5219bdaf496f14d4c060001b6106c7565b613ada7fdd22ffc7f39a1291d79cb9383d49658dfaaf31faa33d2836fed39b2527fe470160001b6106c7565b6000613b087fb5bf9e4c61f2238fe5023a64ee095a2be066294584723c63e2c5081f027ef27e60001b6106c7565b613b347ffe15be7941cafa0a8b2fca5eff8b458a47fe7e846679b17cc9ef85fb52fcffdd60001b6106c7565b6000613b7360405180604001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018967ffffffffffffffff168152506111ed565b9050613ba17ff004f46c68cfe8cd630d55e335465d1ea108fff843ac5f80bf32f982abf15e4460001b6106c7565b613bcd7faeb85a1c3eae9fe199db5d1e4fc2a222b85f33d0f3cd85619e9aa8b6cb0ce90960001b6106c7565b6060613bfb7f4b5fc591f43c8d1066a1d609fd8918b95fc1314bd2985eb1137019cd7e8d0f8e60001b6106c7565b613c277f81ea11037364be1288001d7a16094eb5f6dbfafe813ab8e3bc4a904b7114bdbe60001b6106c7565b8161012001519350613c5b7f04dc0b91aec1d09f76491ed6e6e7d6e2d5d2cd2a6b5f38407dd33b1e7103da3560001b6106c7565b613c877fa1052197ae9f3df966568858e6b136e0fa835708beb068ac22d14459a0c6b7b260001b6106c7565b60008467ffffffffffffffff161415613e0257613cc67fce1b94a0004782776a7f260590d5e9d6b1f31032de0b73c63d543879df6dd3b460001b6106c7565b613cf27f6a1370d1d138502e35e6301884817e8c04a36db390a91e24e93f1fa5bdc16f0760001b6106c7565b613d1e7f8c3d233ff1dd2a456cd874374d43ccc7850db1d10769a75ee88547f62f3a7a2060001b6106c7565b60019350613d4e7ffd54be18c0512f6f0e383f20f2dbfdda83c7d31f18daf86a2dbb0db98cda127e60001b6106c7565b613d7a7f0bb186568fef3c3b035f1f42fcd85dde265ca1a02cbfad15bc56a4430d1f742060001b6106c7565b60019250613daa7f6801affa0d531a3ed14b1d33964ae5ffe0a55a5857f7da61999c0d2eca29e4e260001b6106c7565b613dd67f47fda20fce7b19a657e14f1e34187d27469b7bab997fd2d9a8bbf10dba60cc8160001b6106c7565b60405180606001604052808567ffffffffffffffff16815260200182815260200182815250945061422a565b613e2e7ff7771f6a63727796d36780efa2ca0f83cf1bb91145bd73733f85a94f9995ee2860001b6106c7565b613e5a7f95798bf728399f3b5c61bc3ec0733d875f98c93313fd37f3b83fb6360739a01e60001b6106c7565b613e867f2ba9846f9e1c382634b9d28d0328d45a559c8e7a50345a2f0765710589c3748260001b6106c7565b613e9189898661508e565b9450613ebf7f4b57fb6d3fbcd009dbc662f7ee28c18ce902f103e717dd57c46e36547302c09960001b6106c7565b613eeb7f1fc9b12785bcf8c4455a7dcb03fa55e8366124c470c0bf7650dfbd1e8b4a444b60001b6106c7565b600060036000876000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561402e5783829060005260206000209060020201604051806040016040529081600082018054613f6b906182c4565b80601f0160208091040260200160405190810160405280929190818152602001828054613f97906182c4565b8015613fe45780601f10613fb957610100808354040283529160200191613fe4565b820191906000526020600020905b815481529060010190602001808311613fc757829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190613f38565b5050505090506140607f9489a5ba767a56a08ab818d3716866bbf0a3f3a7bbd561b6963711d7289d316b60001b6106c7565b61408c7f63885d494b06d40fdaab29072a8e2389b3ac36482994221aa674b138e711e42e60001b6106c7565b600060169054906101000a900467ffffffffffffffff1667ffffffffffffffff16815114156141fb576140e17f036270ea840fb0833bd9bc13a6fb0fbf4e6ee8af23fd67827768298a564b494060001b6106c7565b61410d7fbc2a2cbc25169ffc9f7b24b8c16df0d9a0d087bb4ac6c49cc28f8680690a70c960001b6106c7565b848061411890618370565b9550506141477f9b9911afbeb5bf3cc48f022f2cdce7f8b7e5086aaeb390423aaf22d1b11d380060001b6106c7565b6141737faca8e5d4df0c77ce9bf2c37efbd7dcdbc270a03a5752820197d6882dfdd8bb1f60001b6106c7565b600193506141a37f13318de186d1d43942b6f7647a035ade4b9f3c5bf767a31cf76d3d0cc58ea98c60001b6106c7565b6141cf7f13388afff1a17fc7509847ad2c619975756f43624630f79995d548dcce39038f60001b6106c7565b60405180606001604052808667ffffffffffffffff168152602001838152602001838152509550614228565b6142277f27d4288e3a25f0c908e2cc7992b9b50fead82eb4f3e718ffde62e3daf992862560001b6106c7565b5b505b6142567f01bece67bceca80702f8fcff5fa98116d496919e3b26376e8c107ad7e1a59fb460001b6106c7565b6142827f6c63f7e1823190802ecc0867993c42e1fac591fedb459005351d3ca685dc179a60001b6106c7565b61428e8989878a61613a565b6142ba7f7cf8135c47bac0fae4815286e05fe2f0b671d019be33ec6a05d0f45de392a84a60001b6106c7565b6142e67f1c067c484456eafc36c92a28f1dfc668af794a9bd41908c83752b5c3e10fc9ae60001b6106c7565b82955050505050509392505050565b60006143237f88954ea6c643a364b6368e36c9f0ef039824e65ebe790790bdc9dbbce792b67f60001b6106c7565b61434f7f23cbe6edc71733af978599a3e038078cb03f991bc29d7981b471119230b3901760001b6106c7565b61437b7f7aa4226e08770231c8fe569d00e719f9ca76d8573627c7d449aad1fc88a5c0f960001b6106c7565b60006143a97fc17bae04e15157d03c2089f176c76c5c5cb19a8ff23148cae01cc2d87addd1b160001b6106c7565b6143d57fc87dd4a2396fc5879de3f8ecbd9451a60148b715d68bc89f2502be4305e7537460001b6106c7565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561477a5783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff166002811115614591577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156145c9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156147635783829060005260206000200180546146d6906182c4565b80601f0160208091040260200160405190810160405280929190818152602001828054614702906182c4565b801561474f5780601f106147245761010080835404028352916020019161474f565b820191906000526020600020905b81548152906001019060200180831161473257829003601f168201915b5050505050815260200190600101906146b7565b505050508152505081526020019060010190614436565b5050505090506147ac7f8a87b94961f8ede1121dea7439799293e2435baadae5db5749563ae05614f36d60001b6106c7565b6147d87f300412e48bedd2ee422d86aaeca06b2fa0a982faf358dea823ee8255c986783860001b6106c7565b60005b81518110156148a1576148107f535908c7e33095dc6ec38daf846163093eb119323985bc1958e7e6b480946b0160001b6106c7565b61483c7fae6e8e32b20d696ebe3dc5e9d179608fe563e65cc66421584d4ab24f76f5a6b660001b6106c7565b818181518110614875577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151604001518361488c919061808a565b9250808061489990618327565b9150506147db565b506148ce7f4d33fe4d0c5dc1b6b90e005d9a75b885f4d64147001d55e79c26971d7f9f61d760001b6106c7565b6148fa7f6bc29a05f6d0f227310b9d16fc6b371b1e3aaa27f3711ae2d2942c88585def0e60001b6106c7565b8192505050919050565b60006149327fe6246d44d7ec74722c93ab922d30ab08921af1369c86f4ef33f9293294d9f76060001b6106c7565b61495e7f0520949701a9afbd50ad72c71b829f166a84e807564612846851d3144b57413e60001b6106c7565b61498a7fc062cca1232ef3c74f8e221d7ba32ff4f3f060549986e2feef259a5f9ec5c81f60001b6106c7565b600061499584612ccb565b90506149c37f7704da39fda0192d6ed64c91f96d0332cce7e68e073920dc06e3957a1f9881dd60001b6106c7565b6149ef7fe85f0e3b5b574800811a6d7a3eacb82c01d16576eea1a90667d5de62c4a4d68260001b6106c7565b60005b8151811015614bfc57614a277f268651888a32456a957a41c0f974f5f790653cd0517734db2f526d4b9ca4b74c60001b6106c7565b614a537f9176715d3817467dc52c4964006806dc3788c4bc0933995b696d520981db7b2060001b6106c7565b8473ffffffffffffffffffffffffffffffffffffffff16828281518110614aa3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff16148015614b2957508367ffffffffffffffff16828281518110614b11577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff16145b15614bbd57614b5a7f04645f0a5c5593ef1ea01702d558a84050ccf769a2970056c9133ec4126c8aed60001b6106c7565b614b867f7b9ec74a6d5bae05b733d0cf32d5b06704f5162a1b39861527f567246489c74860001b6106c7565b614bb27f45b5f260c01924b817d24c8fa6c4164207a32a5f07b7e022717afd537308dc4460001b6106c7565b600192505050614c5b565b614be97fe6ab4fd87ce299c46fd8c092c0eee85333e8189b1cf2e4db412f34c54e475d0260001b6106c7565b8080614bf490618327565b9150506149f2565b50614c297fd95a29d5b248a0cc338ee13a9c2ddd27dfcf2bc8fb7b77aab7cbde6be1a3678560001b6106c7565b614c557feb6ff4957f49ed0a62be45652e831d37f40f211806fc65a16a138f1095206d4960001b6106c7565b60009150505b92915050565b614c8d7f6f4e2ac6f88d0352cc5ed946913f1f1efe1d2f1caa028309f3b5766daf57f21d60001b6106c7565b614cb97f5d46446705b389fcb9794236660e8125b8d5735988d49b9ef31be07f839b890060001b6106c7565b614ce57f100595cda781c30d289a21fdfaad4a6c8da80fae35332b97f56d97051856c8c460001b6106c7565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115614e73577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190614f59929190616568565b5050505050565b6000614f8e7f98a2ba0c4f3f8f5c1ad77e1ca6d302cae8044505ba4a08e5496b52ca741bf5e260001b6106c7565b614fba7fafa2c992801a26149525d135ae4d5cad412d81465fe5460e9c25a8423e5d808760001b6106c7565b614fe67f3addb1a0993289ff4e2a99b00633787233b715d361bb32ec682badc68340391e60001b6106c7565b600061502560405180604001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff168152506111ed565b90506150537f1f99e6fbde8df0e716b94d50c369ad61afb96c7f37627ef39a418ea22b72a02660001b6106c7565b61507f7fe83baa9c930b0cf8f351c1c9edda5f4fdbae6626b95a6e73bdc7af3702b40e1060001b6106c7565b80610120015191505092915050565b615096616a01565b6150c27f9faac8b9e0720d8800490364229facc0315645d3261a2825ceccce218d595b8060001b6106c7565b6150ee7ffacdde5a900b43a79f1325b5fc658032bb5a5882e99f1ab58c684fd6ed8740bf60001b6106c7565b61511a7f9855253135f06ae48caa7bdb53e22239f8bd4ef4d7241ae6e7efbb09c4b5136b60001b6106c7565b600084848460405160200161513193929190617c93565b604051602081830303815290604052905061516e7f3ba3c0bae7dc1432dd1e175b42a031a1b959de568c1145dbedb7306f253f87e960001b6106c7565b61519a7f28fe145b37402e181e4bd3c1b2d441baf1467b8bcf10b0aa549f5f6535dd028060001b6106c7565b6002816040516151aa9190617cd0565b90815260200160405180910390206040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615205906182c4565b80601f0160208091040260200160405190810160405280929190818152602001828054615231906182c4565b801561527e5780601f106152535761010080835404028352916020019161527e565b820191906000526020600020905b81548152906001019060200180831161526157829003601f168201915b50505050508152602001600282018054615297906182c4565b80601f01602080910402602001604051908101604052809291908181526020018280546152c3906182c4565b80156153105780601f106152e557610100808354040283529160200191615310565b820191906000526020600020905b8154815290600101906020018083116152f357829003601f168201915b5050505050815250509150509392505050565b6000806153527f9c682b2292b6b73653e0024d31742a34f7d362eb1ec7801fb2dd7375dcfbbd3f60001b6106c7565b61537e7f44ab4aaf9c9e937a9ff5813e7a3f070ecb0143d8293dd11f24bb1fb6adabb68660001b6106c7565b6153aa7ffc7a93ef14375293879373fbde5ae2f3beb9671fd8ae8371147985e3ec44b6b560001b6106c7565b600060036000866000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156154ed578382906000526020600020906002020160405180604001604052908160008201805461542a906182c4565b80601f0160208091040260200160405190810160405280929190818152602001828054615456906182c4565b80156154a35780601f10615478576101008083540402835291602001916154a3565b820191906000526020600020905b81548152906001019060200180831161548657829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815260200190600101906153f7565b50505050905061551f7f7219c30c781337520f55c752a5b445db599be9721bba75ca4a9749c69f905d8e60001b6106c7565b61554b7f74e0d2cdade352af02189ccc61fa04bbaf62875889de59fd620dbdb8e2cd7c4760001b6106c7565b6000815114156155e6576155817fbf83b2ae8399ac2e5750453260bbd031e58e2abb9c9630ecb2597494ae63b05b60001b6106c7565b6155ad7f103f6228748b4dd04ee7818be5e5b3ae51bcf6b8f2576f29b98b94487e80660060001b6106c7565b6155d97fecfb868b6ced0f454cc278053a96e4f81cf586152118413d9017a39f76cc570f60001b6106c7565b600080925092505061599d565b6156127f7c796f4fd653179b6e7ce8c6901cc4a49c36409270450bd442115916bdc5214d60001b6106c7565b61563e7f1a6ccfb704e932f121b461e4f9917013f84e4e967661ef07fa4c3ed1238b051b60001b6106c7565b61566a7fed9d50eefe5fce550359e2042ef6fa0ab0055a4290bfa0558497faf944fdb65060001b6106c7565b838051906020012085602001518051906020012014615714576156af7fe8e9d22bf1a7d512a0913fef4c182892d38a21ce6635d192705264ea5d3b3fda60001b6106c7565b6156db7f7357d17652acbb4889a4900e11454ba1b1224d3aae4a21a28389f6f8b4e120aa60001b6106c7565b6157077fb8f4fe738f9b11733d37e00eb2a9fae416f2c7b3888533a917425d2680469b6560001b6106c7565b600080925092505061599d565b6157407f1f213faface3c19439795092c5dcc13455c6eceef95fb04bac9cef5a5614880b60001b6106c7565b61576c7f5713f6375ab83abf367c9f340528da7428bc3fc82dc7817bd4077ecd4de9d71660001b6106c7565b6157987fc0a2c8027d1be3707a4b7ecb4ac43d4fda271daeff1f4ec35dae44ebf55db09760001b6106c7565b60005b81518167ffffffffffffffff16101561593b576157da7f495e56b3c4f2b09c1c48505e53c4d9132b8b1584948740c0c2f5730c078d0e8d60001b6106c7565b6158067f2b73c931aa25c770283babb76eadc0659bea8425d20c17611cb4259e3cbfae8260001b6106c7565b8480519060200120828267ffffffffffffffff1681518110615851577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600001518051906020012014156158fc576158967f43d9bb361795ae26f72509b6037d1d6f30163201d94ffb9b3ca3d65488639e2e60001b6106c7565b6158c27f3239d6be6dda14ac9fd7cc2ab71c1410287bb1018fd51f39c4449d95fb3af88d60001b6106c7565b6158ee7f447961f1f80b0a7090fd9d7be22a9a618d397efa2e3039c96b99d610cc40ecc360001b6106c7565b80600193509350505061599d565b6159287f1c28798eca951252a7d13a8344361b362fd9b8771e962df4afdf1e15655e065b60001b6106c7565b808061593390618370565b91505061579b565b506159687f056e9410f1b889d3557d1653ebf626594cd8fc75880f9270cc6c5c4ad2b3f8ee60001b6106c7565b6159947f2284870b8b1affffbbd4ded57ca189d44134321ec0094e40c1256302c94e14f860001b6106c7565b60008092509250505b9250929050565b6159d07fad73b3102144104c71f8412733a243b462d5923c459b2265af0915c083c0935560001b6106c7565b6159fc7f9b2a2cd9a3e2b0f8ad5273f24efc86152bd8bdae117d558e2650eebe67ad4c2460001b6106c7565b615a287ffaf37c014b854fe778486fb4e46d3965cda25887a8559f9e839ca2680690e66160001b6106c7565b600083838360000151604051602001615a4393929190617c93565b6040516020818303038152906040529050615a807f85de4f30ba1049eab69f0afce5d8f2b4936dd8facab8a9c9833e8027932c9d4460001b6106c7565b615aac7f070083e9614eecb253ff9784d9be5b4b53e05bca773e070d9f96454991c7040f60001b6106c7565b81600282604051615abd9190617cd0565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151816001019080519060200190615b16929190616a2c565b506040820151816002019080519060200190615b33929190616a2c565b5090505050505050565b600080823b905060008111915050919050565b615b7c7f0f58d24cdc473993c3760e934f14269c30cc8608236a4b7815f4a9dfae63b12860001b6106c7565b615ba87f73c509a1dd49c36f3ad01cf45043bce236af918fb54736371209b97b38f953bc60001b6106c7565b615bd47fa6b73cbb67995daee3241670e9754a41e6b167ed0d0dfdd63ab955315ae2291560001b6106c7565b6000615be08383614f60565b9050615c0e7ffc90b2a44a0a3ca8f33581dca89bf6aac58c545837134eb3d3e317fa700e726f60001b6106c7565b615c3a7fc827d34dd8d658eeee0115a8862bd8f8673e88faf2fdab031259bd55be46396160001b6106c7565b60005b8167ffffffffffffffff168167ffffffffffffffff161015615ccf57615c857f3e42eee18224f16272817a5a06388b0582111a41ea5da0bfaf1630b586f28e9260001b6106c7565b615cb17f840eaff1b93d07d7b684dd69c1af2f791c7a90cd4854c6507646dfd4fb554a2960001b6106c7565b615cbc8484836163cf565b8080615cc790618370565b915050615c3d565b50505050565b615d017f645eadf21c8ca050bc83ffcd3e887ed1739c8da58172bb7d7087299fc198998160001b6106c7565b615d2d7fb09cd63be178679f2a2eb19d4569facb58db5706bb8b64b4ccf203316e33b67c60001b6106c7565b615d597fa9dc3235459080e1fc665e42700c646798f73c047df2acf7770dcafbc58d809c60001b6106c7565b60005b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508167ffffffffffffffff16101561613557615ddc7fa57325d810ce13eaad3e0238796c92b38d4caa139dbab291fe292babf307c8a060001b6106c7565b615e087fc1d6ef36c034d9c82ec5246e2e22e1458bbf8b155d5867a3135aaa1bb646b7d060001b6106c7565b8167ffffffffffffffff16600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208267ffffffffffffffff1681548110615e94577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1614156160f657615ef67fae5a6958c3ed7cf1fbabc62d99a55df18689a6aabf7f3a6f07562bab7f16a83060001b6106c7565b615f227f51d3ad352dffbb219284a862da131f3941f8b5c8946981189a2b672d1836d13160001b6106c7565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208167ffffffffffffffff1681548110615fa3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060060201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff02191690556005820160006160c39190616ab2565b50506160f17f38c5a4a4bca5bb899c56027e15fcb7dec61b629afd815ff7d1e2156c2c86fd6860001b6106c7565b616135565b6161227ff12767c815a2799a54f5dcd3d1744757535cf5ab46a554a9def058c4f2297c5560001b6106c7565b808061612d90618370565b915050615d5c565b505050565b6161667fc5f9790451f67a50dd162f174b75050b4aa4006e1418c918fc721da2d803175860001b6106c7565b6161927fcf44cc969a3aece8cdf31ceea2360446b0e0fde9462dd87b6b9e52f2e72f166860001b6106c7565b6161be7ffd86ba8091dc033eefbe760e285bbb70ee6718c8a909000ff9cc9402c71b9af260001b6106c7565b6000848484600001516040516020016161d993929190617c93565b60405160208183030381529060405290506162167f9e6a03c000a81d9b92a5dc2045dcef7eb64381708572dffe96072c7c51ff07b560001b6106c7565b6162427fc46485066475e0d446d7cd4c915759bb8faed1a8bf4e347a4ede19b26656d51b60001b6106c7565b826002826040516162539190617cd0565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160010190805190602001906162ac929190616a2c565b5060408201518160020190805190602001906162c9929190616a2c565b509050506162f97fb24f94b22e81afaba1cc93e3759ff237ed9beb4e2c31c7adcf476c5e0e47008f60001b6106c7565b6163257f1699877131bcec2ccc42e6c733c4d6b2e47d6cd1484dcec23e8d6d0ce323f40c60001b6106c7565b60036000846000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208290806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000019080519060200190616396929190616a2c565b5060208201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6163fb7f39633b1cedefaa2c00234699518df64e3a0af47098f8bf80db8bc7b9fa9aa9b760001b6106c7565b6164277fd82166d586970dbb2b6382d5cf5312acd0eb3122786d28bce66b321ae24db2bb60001b6106c7565b6164537f5f512d484bc0d8f3de349b8e3c03dd2d65445caeacab0da4e1d2fba8c95263dc60001b6106c7565b600083838360405160200161646a93929190617c93565b60405160208183030381529060405290506164a77fa3621aae50e00786bf415e2e2d28823d0eb36e0c905c4a7c3b4106542a3a4c5760001b6106c7565b6002816040516164b79190617cd0565b9081526020016040518091039020600080820160006101000a81549067ffffffffffffffff02191690556001820160006164f191906169c1565b60028201600061650191906169c1565b505061652f7ff034141d625caf774e26d9fb6795a1dc9c060caeb93d25bcbc2e2acd7ebc75b660001b6106c7565b600360008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006165629190616ad3565b50505050565b8280548282559060005260206000209081019282156165b7579160200282015b828111156165b65782518290805190602001906165a6929190616a2c565b5091602001919060010190616588565b5b5090506165c49190616af7565b5090565b8280548282559060005260206000209060060281019282156168c35760005260206000209160060282015b828111156168c25782826000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900467ffffffffffffffff168160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160089054906101000a900467ffffffffffffffff168160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160109054906101000a900460ff168160010160106101000a81548160ff0219169083600281111561678b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060028201548160020155600382015481600301556004820160009054906101000a900467ffffffffffffffff168160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160089054906101000a900467ffffffffffffffff168160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160109054906101000a900467ffffffffffffffff168160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160189054906101000a900460ff168160040160186101000a81548160ff02191690831515021790555060058201816005019080546168b0929190616b1b565b505050916006019190600601906165f3565b5b5090506168d09190616b83565b5090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000600281111561696a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b5080546169cd906182c4565b6000825580601f106169df57506169fe565b601f0160209004906000526020600020908101906169fd9190616cab565b5b50565b6040518060600160405280600067ffffffffffffffff16815260200160608152602001606081525090565b828054616a38906182c4565b90600052602060002090601f016020900481019282616a5a5760008555616aa1565b82601f10616a7357805160ff1916838001178555616aa1565b82800160010185558215616aa1579182015b82811115616aa0578251825591602001919060010190616a85565b5b509050616aae9190616cab565b5090565b5080546000825590600052602060002090810190616ad09190616af7565b50565b5080546000825560020290600052602060002090810190616af49190616cc8565b50565b5b80821115616b175760008181616b0e91906169c1565b50600101616af8565b5090565b828054828255906000526020600020908101928215616b725760005260206000209182015b82811115616b71578282908054616b56906182c4565b616b61929190616d0a565b5091600101919060010190616b40565b5b509050616b7f9190616af7565b5090565b5b80821115616ca757600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff0219169055600582016000616c9e9190616ab2565b50600601616b84565b5090565b5b80821115616cc4576000816000905550600101616cac565b5090565b5b80821115616d065760008082016000616ce291906169c1565b6001820160006101000a81549067ffffffffffffffff021916905550600201616cc9565b5090565b828054616d16906182c4565b90600052602060002090601f016020900481019282616d385760008555616d86565b82601f10616d495780548555616d86565b82800160010185558215616d8657600052602060002091601f016020900482015b82811115616d85578254825591600101919060010190616d6a565b5b509050616d939190616cab565b5090565b6000616daa616da584617f4c565b617f27565b90508083825260208201905082856020860282011115616dc957600080fd5b60005b85811015616df95781616ddf8882616ec7565b845260208401935060208301925050600181019050616dcc565b5050509392505050565b6000616e16616e1184617f78565b617f27565b90508083825260208201905082856020860282011115616e3557600080fd5b60005b85811015616e7f57813567ffffffffffffffff811115616e5757600080fd5b808601616e648982616f84565b85526020850194506020840193505050600181019050616e38565b5050509392505050565b6000616e9c616e9784617fa4565b617f27565b905082815260208101848484011115616eb457600080fd5b616ebf848285618258565b509392505050565b600081359050616ed681618602565b92915050565b600081519050616eeb81618602565b92915050565b600082601f830112616f0257600080fd5b8135616f12848260208601616d97565b91505092915050565b600082601f830112616f2c57600080fd5b8135616f3c848260208601616e03565b91505092915050565b600081359050616f5481618619565b92915050565b600081519050616f6981618619565b92915050565b600081359050616f7e81618630565b92915050565b600082601f830112616f9557600080fd5b8135616fa5848260208601616e89565b91505092915050565b600081359050616fbd81618647565b92915050565b600081359050616fd28161865e565b92915050565b600081359050616fe78161866e565b92915050565b6000610320828403121561700057600080fd5b61700b6102e0617f27565b9050600082013567ffffffffffffffff81111561702757600080fd5b61703384828501616f84565b600083015250602061704784828501616ec7565b602083015250604082013567ffffffffffffffff81111561706757600080fd5b61707384828501616f84565b60408301525060606170878482850161756c565b606083015250608061709b8482850161756c565b60808301525060a06170af8482850161756c565b60a08301525060c06170c38482850161756c565b60c08301525060e06170d78482850161756c565b60e0830152506101006170ec84828501617557565b610100830152506101206171028482850161756c565b610120830152506101406171188482850161756c565b6101408301525061016082013567ffffffffffffffff81111561713a57600080fd5b61714684828501616f84565b6101608301525061018061715c8482850161756c565b610180830152506101a061717284828501617557565b6101a0830152506101c061718884828501616f45565b6101c0830152506101e061719e84828501616fd8565b6101e0830152506102006171b48482850161756c565b6102008301525061022082013567ffffffffffffffff8111156171d657600080fd5b6171e284828501616ef1565b6102208301525061024082013567ffffffffffffffff81111561720457600080fd5b61721084828501616ef1565b6102408301525061026082013567ffffffffffffffff81111561723257600080fd5b61723e84828501616f84565b6102608301525061028061725484828501616fc3565b610280830152506102a061726a84828501616f45565b6102a0830152506102c06172808482850161733d565b6102c08301525092915050565b600060e0828403121561729f57600080fd5b6172a960e0617f27565b905060006172b984828501617581565b60008301525060206172cd84828501617581565b60208301525060406172e184828501617581565b60408301525060606172f584828501617581565b606083015250608061730984828501617581565b60808301525060a061731d84828501616edc565b60a08301525060c061733184828501616edc565b60c08301525092915050565b60006060828403121561734f57600080fd5b6173596060617f27565b905060006173698482850161756c565b600083015250602061737d8482850161756c565b60208301525060406173918482850161756c565b60408301525092915050565b6000602082840312156173af57600080fd5b6173b96020617f27565b905060006173c98482850161756c565b60008301525092915050565b600061018082840312156173e857600080fd5b6173f3610180617f27565b9050600061740384828501616ec7565b60008301525060206174178482850161756c565b602083015250604061742b8482850161756c565b604083015250606061743f8482850161756c565b606083015250608061745384828501616fc3565b60808301525060a061746784828501617557565b60a08301525060c061747b84828501617557565b60c08301525060e061748f8482850161756c565b60e0830152506101006174a48482850161756c565b610100830152506101206174ba8482850161756c565b610120830152506101406174d084828501616f45565b6101408301525061016082013567ffffffffffffffff8111156174f257600080fd5b6174fe84828501616f1b565b6101608301525092915050565b60006040828403121561751d57600080fd5b6175276040617f27565b9050600061753784828501616ec7565b600083015250602061754b8482850161756c565b60208301525092915050565b6000813590506175668161867e565b92915050565b60008135905061757b81618695565b92915050565b60008151905061759081618695565b92915050565b6000602082840312156175a857600080fd5b60006175b684828501616ec7565b91505092915050565b6000602082840312156175d157600080fd5b60006175df84828501616f5a565b91505092915050565b6000602082840312156175fa57600080fd5b600061760884828501616f6f565b91505092915050565b6000806040838503121561762457600080fd5b600061763285828601616fae565b92505060206176438582860161739d565b9150509250929050565b600060e0828403121561765f57600080fd5b600061766d8482850161728d565b91505092915050565b60006020828403121561768857600080fd5b600082013567ffffffffffffffff8111156176a257600080fd5b6176ae848285016173d5565b91505092915050565b600080604083850312156176ca57600080fd5b600083013567ffffffffffffffff8111156176e457600080fd5b6176f0858286016173d5565b925050602083013567ffffffffffffffff81111561770d57600080fd5b61771985828601616fed565b9150509250929050565b60006040828403121561773557600080fd5b60006177438482850161750b565b91505092915050565b600061775883836178b1565b905092915050565b600061776c8383617a38565b905092915050565b61777d81618172565b82525050565b61778c81618172565b82525050565b6177a361779e82618172565b6183a1565b82525050565b60006177b482617ff5565b6177be818561803b565b9350836020820285016177d085617fd5565b8060005b8581101561780c57848403895281516177ed858261774c565b94506177f883618021565b925060208a019950506001810190506177d4565b50829750879550505050505092915050565b600061782982618000565b617833818561804c565b93508360208202850161784585617fe5565b8060005b8581101561788157848403895281516178628582617760565b945061786d8361802e565b925060208a01995050600181019050617849565b50829750879550505050505092915050565b61789c81618184565b82525050565b6178ab81618184565b82525050565b60006178bc8261800b565b6178c6818561805d565b93506178d6818560208601618267565b6178df81618493565b840191505092915050565b6178f381618210565b82525050565b61790281618222565b82525050565b61791181618222565b82525050565b61792081618234565b82525050565b61792f81618246565b82525050565b600061794082618016565b61794a818561807f565b935061795a818560208601618267565b80840191505092915050565b600061797360158361806e565b915061797e826184be565b602082019050919050565b600061799660138361806e565b91506179a1826184e7565b602082019050919050565b60006179b960148361806e565b91506179c482618510565b602082019050919050565b60006179dc602e8361806e565b91506179e782618539565b604082019050919050565b60006179ff60118361806e565b9150617a0a82618588565b602082019050919050565b6000617a22601f8361806e565b9150617a2d826185b1565b602082019050919050565b600061018083016000830151617a516000860182617774565b506020830151617a646020860182617c5e565b506040830151617a776040860182617c5e565b506060830151617a8a6060860182617c5e565b506080830151617a9d60808601826178f9565b5060a0830151617ab060a0860182617c40565b5060c0830151617ac360c0860182617c40565b5060e0830151617ad660e0860182617c5e565b50610100830151617aeb610100860182617c5e565b50610120830151617b00610120860182617c5e565b50610140830151617b15610140860182617893565b50610160830151848203610160860152617b2f82826177a9565b9150508091505092915050565b600061018083016000830151617b556000860182617774565b506020830151617b686020860182617c5e565b506040830151617b7b6040860182617c5e565b506060830151617b8e6060860182617c5e565b506080830151617ba160808601826178f9565b5060a0830151617bb460a0860182617c40565b5060c0830151617bc760c0860182617c40565b5060e0830151617bda60e0860182617c5e565b50610100830151617bef610100860182617c5e565b50610120830151617c04610120860182617c5e565b50610140830151617c19610140860182617893565b50610160830151848203610160860152617c3382826177a9565b9150508091505092915050565b617c49816181f2565b82525050565b617c58816181f2565b82525050565b617c67816181fc565b82525050565b617c76816181fc565b82525050565b617c8d617c88826181fc565b6183c5565b82525050565b6000617c9f8286617792565b601482019150617caf8285617c7c565b600882019150617cbf8284617c7c565b600882019150819050949350505050565b6000617cdc8284617935565b915081905092915050565b6000602082019050617cfc6000830184617783565b92915050565b60006020820190508181036000830152617d1c818461781e565b905092915050565b6000608082019050617d3960008301876178ea565b617d466020830186617c4f565b617d536040830185617783565b617d606060830184617c6d565b95945050505050565b600060e082019050617d7e600083018a6178ea565b617d8b6020830189617c4f565b617d986040830188617783565b617da56060830187617c6d565b617db26080830186617908565b617dbf60a0830185617c6d565b617dcc60c08301846178a2565b98975050505050505050565b6000604082019050617ded6000830185617917565b617dfa6020830184617c6d565b9392505050565b6000602082019050617e166000830184617926565b92915050565b60006020820190508181036000830152617e3581617966565b9050919050565b60006020820190508181036000830152617e5581617989565b9050919050565b60006020820190508181036000830152617e75816179ac565b9050919050565b60006020820190508181036000830152617e95816179cf565b9050919050565b60006020820190508181036000830152617eb5816179f2565b9050919050565b60006020820190508181036000830152617ed581617a15565b9050919050565b60006020820190508181036000830152617ef68184617b3c565b905092915050565b6000604082019050617f136000830185617c6d565b617f206020830184617c6d565b9392505050565b6000617f31617f42565b9050617f3d82826182f6565b919050565b6000604051905090565b600067ffffffffffffffff821115617f6757617f66618464565b5b602082029050602081019050919050565b600067ffffffffffffffff821115617f9357617f92618464565b5b602082029050602081019050919050565b600067ffffffffffffffff821115617fbf57617fbe618464565b5b617fc882618493565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000618095826181fc565b91506180a0836181fc565b92508267ffffffffffffffff038211156180bd576180bc6183d7565b5b828201905092915050565b60006180d3826181fc565b91506180de836181fc565b92508167ffffffffffffffff04831182151516156180ff576180fe6183d7565b5b828202905092915050565b6000618115826181f2565b9150618120836181f2565b925082821015618133576181326183d7565b5b828203905092915050565b6000618149826181fc565b9150618154836181fc565b925082821015618167576181666183d7565b5b828203905092915050565b600061817d826181d2565b9050919050565b60008115159050919050565b6000819050919050565b60006181a582618172565b9050919050565b60008190506181ba826185da565b919050565b60008190506181cd826185ee565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061821b826181ac565b9050919050565b600061822d826181bf565b9050919050565b600061823f826181fc565b9050919050565b6000618251826181fc565b9050919050565b82818337600083830152505050565b60005b8381101561828557808201518184015260208101905061826a565b83811115618294576000848401525b50505050565b60006182a5826181fc565b915060008214156182b9576182b86183d7565b5b600182039050919050565b600060028204905060018216806182dc57607f821691505b602082108114156182f0576182ef618435565b5b50919050565b6182ff82618493565b810181811067ffffffffffffffff8211171561831e5761831d618464565b5b80604052505050565b6000618332826181f2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415618365576183646183d7565b5b600182019050919050565b600061837b826181fc565b915067ffffffffffffffff821415618396576183956183d7565b5b600182019050919050565b60006183ac826183b3565b9050919050565b60006183be826184b1565b9050919050565b60006183d0826184a4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f736563746f7220616c7265616479206578697374730000000000000000000000600082015250565b7f6e6f6465206e6f74207265676973746572656400000000000000000000000000600082015250565b7f736563746f722073697a652069732077726f6e67000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f736563746f7249642069732077726f6e67000000000000000000000000000000600082015250565b7f736563746f724964206d7573742062652067726561746572207468616e203000600082015250565b600a81106185eb576185ea618406565b5b50565b600381106185ff576185fe618406565b5b50565b61860b81618172565b811461861657600080fd5b50565b61862281618184565b811461862d57600080fd5b50565b61863981618190565b811461864457600080fd5b50565b6186508161819a565b811461865b57600080fd5b50565b6003811061866b57600080fd5b50565b6002811061867b57600080fd5b50565b618687816181f2565b811461869257600080fd5b50565b61869e816181fc565b81146186a957600080fd5b5056fea2646970667358221220d44306237ead0cf02b5a9414b7807378ab8b6c76b1b9c09b29cb54f4dd88e3a264736f6c63430008040033",
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
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) returns()
func (_Store *StoreTransactor) AddFileToSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToSector", sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) returns()
func (_Store *StoreSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) returns()
func (_Store *StoreTransactorSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreTransactor) AddSectorRefForFileInfo(opts *bind.TransactOpts, sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddSectorRefForFileInfo", sectorInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreSession) AddSectorRefForFileInfo(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.AddSectorRefForFileInfo(&_Store.TransactOpts, sectorInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
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
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) returns()
func (_Store *StoreTransactor) DeleteFileFromSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFileFromSector", sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) returns()
func (_Store *StoreSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) returns()
func (_Store *StoreTransactorSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteSecotr is a paid mutator transaction binding the contract method 0x23de5b98.
//
// Solidity: function DeleteSecotr((address,uint64) sectorRef) returns()
func (_Store *StoreTransactor) DeleteSecotr(opts *bind.TransactOpts, sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteSecotr", sectorRef)
}

// DeleteSecotr is a paid mutator transaction binding the contract method 0x23de5b98.
//
// Solidity: function DeleteSecotr((address,uint64) sectorRef) returns()
func (_Store *StoreSession) DeleteSecotr(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.DeleteSecotr(&_Store.TransactOpts, sectorRef)
}

// DeleteSecotr is a paid mutator transaction binding the contract method 0x23de5b98.
//
// Solidity: function DeleteSecotr((address,uint64) sectorRef) returns()
func (_Store *StoreTransactorSession) DeleteSecotr(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.DeleteSecotr(&_Store.TransactOpts, sectorRef)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) returns()
func (_Store *StoreTransactor) UpdateSectorInfo(opts *bind.TransactOpts, sector SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateSectorInfo", sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) returns()
func (_Store *StoreSession) UpdateSectorInfo(sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) returns()
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
