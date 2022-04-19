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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"AddFileToSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"AddSectorRefForFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"CreateSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"DeleteFileFromSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"DeleteSecotr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"GetSectorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetSectorsForNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sector\",\"type\":\"tuple\"}],\"name\":\"UpdateSectorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x29d6276d\",\"type\":\"bytes32\"}],\"name\":\"c_0x29d6276d\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052611388600060026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503480156200003c57600080fd5b50618611806200004d6000396000f3fe608060405234801561001057600080fd5b50600436106100bd5760003560e01c8063955f98b711610076578063c4d66de81161005b578063c4d66de81461019a578063dcf74960146101b6578063e3168f9e146101d2576100bd565b8063955f98b714610162578063ba9210041461017e576100bd565b80632384a6aa116100a75780632384a6aa146100fa57806323de5b98146101165780632ba010d714610132576100bd565b806247c003146100c25780630daf9945146100de575b600080fd5b6100dc60048036038101906100d791906175e6565b610202565b005b6100f860048036038101906100f3919061752a565b6106c7565b005b610114600480360381019061010f91906175a5565b6106ca565b005b610130600480360381019061012b9190617652565b610ceb565b005b61014c60048036038101906101479190617652565b610fbd565b6040516101599190617e0b565b60405180910390f35b61017c600480360381019061017791906175e6565b611808565b005b610198600480360381019061019391906175a5565b611d0e565b005b6101b460048036038101906101af9190617553565b612878565b005b6101d060048036038101906101cb91906175a5565b612a22565b005b6101ec60048036038101906101e791906174d8565b612c45565b6040516101f99190617c31565b60405180910390f35b61022e7f35e5d832fb326b1c73cdff9fbbf3cac0aa4b2f3455378c524bfdad1cbd48da9960001b6106c7565b61025a7fa86f16bcb3c15d10543a74ab7211410c3da7e2d3bbeba8a051e076fc1c85953260001b6106c7565b6102867f85746e64c7ecf3041cbff4e0572b3d42195be5cfeba9547cb0e3c4cbf97ac46a60001b6106c7565b600061029f836000015184602001518460000151613079565b90506102cd7f1591fa4847409025f3a857f5b86285e99331efa817ef14549ee6ec9f75791b7b60001b6106c7565b826101000180518091906102e0906181c9565b67ffffffffffffffff1667ffffffffffffffff16815250506103247fb4da21644e4a7fc90068443205ac764cdb35c12805d0f686614a9e0a69d52f0c60001b6106c7565b6103507fdc2ad75a81dab5ec57ede91a9f94c16aafebe920cd5e5a2133de6493aa80cc2060001b6106c7565b81608001518360e001818151610366919061806d565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506103af7f6e4803494ed94c785344d3c421414130c5f641b6ed361f7212b9b3e9b658e23e60001b6106c7565b6103db7f2b83974a4c538480f01ec56a74b005a4c417f25f988c43e3169ab586ca3ad53260001b6106c7565b8160a0015182608001516103ef9190617ff7565b83606001818151610400919061806d565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506104497fde2148f3fde82993d48bd2da757673252b00ab21760ef165950452197b7aa9ce60001b6106c7565b6104757f16debf4c3c03075887844a037302a98adb003be1a8c34864aaf48f663656fbcd60001b6106c7565b8015610503576104a77fcd327c750cfc21eb7cfdafb1e430ca04c369bba3986aa0ce707e8f5f77a520cf60001b6106c7565b6104d37f1131f4db2427d7176c491d7f5238319e4d4cfe6c3e97bc68a0c3aa2c2162f71760001b6106c7565b826101200180518091906104e6906181c9565b67ffffffffffffffff1667ffffffffffffffff1681525050610530565b61052f7f767670dfc055717e0e815cf3af7f2e03b74bc72f6bec0659520175fa10bc3a1b60001b6106c7565b5b61055c7f768656c0cdd52722593e2115406d54117e9bd6111cb52089f5b59c7fcf122d8c60001b6106c7565b6105887f1e9ac4df7ae1fcd83ebfd2f052742b296ac1ab9867bf03397118a92714eff0ac60001b6106c7565b600083610100015167ffffffffffffffff161415610634576105cc7f181ff6d3dbb564b62c5a10cb9f6c8d01addcf825a82c46c6dedb4efb2a3285fd60001b6106c7565b6105f87f47ecb464575db4ad742bf4bcb665850697e92f6111a7c68a23dafd8899ef356760001b6106c7565b6106247fb9ffaf145d63153401605995fc5aa7ae5ed9e3f4de53bdb96e85835b2a2ed9f160001b6106c7565b60008360c0018181525050610661565b6106607f807ad8a78f0dfa57aae627e207cef03254721efb201a332afcdfb7227eea6a1460001b6106c7565b5b61068d7ff543cf3b05be29618a0632c7ee129e0ef38370ce135832acf25b5205b981f86060001b6106c7565b6106b97fd638fa618cf0ed1094306a24de3aa59f34f5e5109233ca8b8bb4b0cecb8e674b60001b6106c7565b6106c2836106ca565b505050565b50565b6106f67fbc666fce9df3eb8e9b6cb2e77bbed4c996a8251b175c78a68b8a4eac0dc8ae0f60001b6106c7565b6107227feac775a9cf16c38d8b9c629eae11acdcc5bf10312e56348a1b8d6632480f719b60001b6106c7565b61074e7fcac248084862bf4c8841d11ef4dbbb8cbdaa54b53bf3593eb84b192ee90939c760001b6106c7565b600060016000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506107c17f70da512a885a081f1e609180c719bad9a2b3484c95a934258318203d49b296df60001b6106c7565b6107ed7fd89b77b4e91411fc228e9065a26a61da092b561f169d340e325e63fdac7779ee60001b6106c7565b60005b81805490508167ffffffffffffffff161015610c3b576108327f96bb39363f6c41b477fc2ffacfb9652fab0266dafaf62a323c02508f13b5024860001b6106c7565b61085e7f6946773d589d5aa7cca324353bfe08904598f9bf404adab26232f8b5cf21005d60001b6106c7565b826020015167ffffffffffffffff16828267ffffffffffffffff16815481106108b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff161415610bfc576109127fbd7447020a27aaa6a74daf5cc19e46e79b45a399121a83937226d9a57866d0ab60001b6106c7565b61093e7fdc604c12c1261a0de72ae24e49a6c9faab1b4a289bc7e67f6a7810b386e4c65160001b6106c7565b61096a7ff29e3d9ca8bfba6ddd1ae5d1618ab99fd0d77f3cd0b69c5c5df87ec61ba4844360001b6106c7565b82828267ffffffffffffffff16815481106109ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115610ae1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190610bc79291906164e2565b50905050610bf77f3c15cea3d76d44a626f8e25d095890a230a4e52587f1346063f200db9186eadb60001b6106c7565b610c3b565b610c287fadce9fcba52bc14487ce82b6a50f02a59543c5ed78da5791d7d22a268a99c4ba60001b6106c7565b8080610c339061829f565b9150506107f0565b50610c687f7eecdbd0a4369e120c0e1e5af3b381580c1e892fe353bfe3cf979c72ad5c87d660001b6106c7565b610c947fb92fa28e211d3aee5046d8a76d1db546278d00f28abf198994077fedae88edc360001b6106c7565b8060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908054610ce6929190616542565b505050565b610d177fb4f0b5d974531722ff35261c80b14503c988029d42bb0ec4775c6b9cbc810aab60001b6106c7565b610d437fa82c541da1ad66baabcd968f8cf906b932aacd7497cd44a83b33d59b1d49f0a660001b6106c7565b610d6f7fdd64fb8c4bb9f7e7da161a472f4cbc6038801e52448233ffafc903b1dddc3f6660001b6106c7565b6000610d7a82610fbd565b9050610da87f8a27c9efbc617dd2395610a1be5a73ba4b96d5367ae05f93fcbfdbfff13d6e4660001b6106c7565b610dd47fbab33c829c9996924ede78ce61c6eff32465be5989392bda441ff134786c74fe60001b6106c7565b600081610100015167ffffffffffffffff161115610e8957610e187f3cf02417cc53ee01c167db8b94f9b63822e8a29a93f29b37a4c1a8774f8c32c760001b6106c7565b610e447f8a907d1912286be9ce0a50144367f9cd82078da8356afd97c6d1fc17dcb7a47c60001b6106c7565b60008161010001516040517f25a56d93000000000000000000000000000000000000000000000000000000008152600401610e80929190617d07565b60405180910390fd5b610eb57ff5e36d18e978c23ab4709c319958d0ea692a04ee2a36be6413b999490ccec4f360001b6106c7565b610ee17fb250b34256aad002f500055491df26039a71a69d7eb891182f6e47bd21ef795960001b6106c7565b610f0d7f44631d57c04619ffdc78464ae9410f4633b14817b2955503787654c50be04de760001b6106c7565b610f1f8260000151836020015161380f565b610f4b7f839dbb50d8d37a5cacb3575f76247cec8339cf683a40d8e7b9be3536acf18a8e60001b6106c7565b610f777f0b20b4459e1e710bdb6ccdf7bddfe080d20d29fb466e929c552bc3be884883c260001b6106c7565b7f4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c600943338560200151604051610fb19493929190617c53565b60405180910390a15050565b610fc561684e565b610ff17fed075eb98a1cdadf955c3e96d4be67fdc9854ce85b585c2854d3ff7d934e101760001b6106c7565b61101d7f9ab838b6a611c18b5ffa6dbe763f199bd2a3b50e93b5c133c815046d7e446b8e60001b6106c7565b6110497f403ada0f1b9922170593c52135bdadf9682b0b26d54cca69aa9a228da1c8292560001b6106c7565b6110757f0b224ee8641c55ed20be4fcd221366546700502750a93494f38bc1ec3ed7d57d60001b6106c7565b6000826020015167ffffffffffffffff16116110c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110bd90617deb565b60405180910390fd5b6110f27fa099831781b28930643d93eddec0e9545d2e564016aa4d9c90617c98f609c11c60001b6106c7565b61111e7f3516db01dc1927b78206ef086546ecd86cb7ddf44a2f683a27032aa74a100c4260001b6106c7565b61114a7f18d71158e4bed03f6767a2d50d4c410e215171650fdb18f55106dd5dd8fd177860001b6106c7565b600060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156114f35783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff16600281111561130a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611342577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156114dc57838290600052602060002001805461144f906181f3565b80601f016020809104026020016040519081016040528092919081815260200182805461147b906181f3565b80156114c85780601f1061149d576101008083540402835291602001916114c8565b820191906000526020600020905b8154815290600101906020018083116114ab57829003601f168201915b505050505081526020019060010190611430565b5050505081525050815260200190600101906111af565b5050505090506115257fc347d3e28813d8ced8f06470153ba5f580c36fef40fec9a2b801bcb97488f90d60001b6106c7565b6115517f769e816c7ed60337ba352d71a5687c83cfb1c6e57395bf8d2743bcfa6d51aac660001b6106c7565b60005b81518167ffffffffffffffff161015611744576115937fb76c9edc5325d2190f8b415bd94be79f3f567ee95cb84260e46b611e6716e7d760001b6106c7565b6115bf7fe952608570efbd273b78ef5bb2da16e4b6f32f453564fe98a0c8b09ff7e724cd60001b6106c7565b836020015167ffffffffffffffff16828267ffffffffffffffff1681518110611611577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff161415611705576116597f6b30fed03acd5fc7882a50c740716fa1c5b3d2a1a221414a2ff1045f16f4de5260001b6106c7565b6116857f5346d600527ad7dfbf0c8cf4941ab1b02e05b852c8843aca86e9f11db9af635c60001b6106c7565b6116b17fe423cb92a85981f3b7008291bf62a326050f24cca06633dbb404ddb566f5944c60001b6106c7565b818167ffffffffffffffff16815181106116f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015192505050611803565b6117317fba7a59354f295e9cfa29b0c61ba5b57ad255aefda16543c531b2e97b07b184d760001b6106c7565b808061173c9061829f565b915050611554565b506117717f76eba042b8cc59bbf45380851c350b1b02b972f5bd022be8d2d8ac0b2711663260001b6106c7565b61179d7f2bffca3adb168293b23513ff9b15cf7278722e6f871e7271d645d6ab79827f5760001b6106c7565b6117a561684e565b6117d17ffd15fbd409606a0e1426f1a4e35ef9918c7d64e1d15d42239b4c12500b92458b60001b6106c7565b6117fd7f45b11e98e533070af462a3c15704ff945089d4ea678a722397bdfe3ff50e95a060001b6106c7565b80925050505b919050565b6118347fb723d1ddc452a736b1018e9a5177876d92ab5888472b31d37d9032f0651b5e7c60001b6106c7565b6118607f2625bcc04197a9eec35125c3d604c539936c5777fd7d3b6d546aecfc21314bb860001b6106c7565b61188c7fb12168ace3033c5c71c6c284837441a07034250f27324108d98fc815e60ce6c060001b6106c7565b816040015167ffffffffffffffff168160a0015182608001516118af9190617ff7565b83606001516118be9190617fb9565b67ffffffffffffffff161115611958576118fa7f6fc6e948d56023414ed5ee082fc755bc9b267318aa22a4dc9a2df0a2a548267060001b6106c7565b6119267f4e5aa2f42aad718572b8e312ef3cc8d501cc37fae1ec6ed8e58136f3d7972b5a60001b6106c7565b6040517f6073072a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119847f81040a56e8e7fba8b618c700914350ce80252f33721d471405c0c16ebf9a9b9760001b6106c7565b6119b07fb67f616f7083384b537f72c851b990da40e2cdcd06a99d4b82c6f3776f5b578f60001b6106c7565b6119dc7fc942034943bbd5a425fd0240e27b83bfe331c5bcf2c2d77d23fba6b2c6ebbd0060001b6106c7565b6000611a1783600001518460200151604051806040016040528086600001518152602001866080015167ffffffffffffffff16815250613903565b9050611a457fc61f44af1f16d42dba6080b97cd9006c9b38435a832384f77c1900e4e74d315360001b6106c7565b82610100018051809190611a589061829f565b67ffffffffffffffff1667ffffffffffffffff1681525050611a9c7f115e7c0c4d3f6c2f36cc31e01e2d562b29eb22c2c23e76712eea934b3bc5ee2a60001b6106c7565b611ac87f122630e1e79f3f1a744d6538736e44ca8399d3244b807bfe3d252b0fd74f4c8560001b6106c7565b8160a001518260800151611adc9190617ff7565b83606001818151611aed9190617fb9565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611b367f0f19a9022ff6695c306985e4add75dbc8c230adcd36a662929533008bcd953ff60001b6106c7565b611b627f225f8600a660f39f2fdf5d7a5be49ea24132cea3eb60722a4f9405ed9d42e72d60001b6106c7565b81608001518360e001818151611b789190617fb9565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611bc17f246c11e19d66495175f6594cfa67fbbed20ba1426f444c75c35dd9d9d171100860001b6106c7565b611bed7f623d477333391236a147d2a22cbe1504ac9f0865e097f4a5a1a6241dd29207a060001b6106c7565b8015611c7b57611c1f7fb3c0d0b50c9a48e866b9b9adecbe67693f6e105686f6ab410015f7eeee4b717d60001b6106c7565b611c4b7f27749044d2a0cfe7e6099c7adf57f43e349db5dd844c126c892396e3a4cd539160001b6106c7565b82610120018051809190611c5e9061829f565b67ffffffffffffffff1667ffffffffffffffff1681525050611ca8565b611ca77fee0e521f14a2b1771fa341230360cc39d92fde6eecdf467106a80e8507d8e79960001b6106c7565b5b611cd47f8611527167e4ed1332de539005826d6793e986f27a0227006249d2bf30cfdf3460001b6106c7565b611d007f9deefab10afb6c63bcec0656a339a91552ed41fc33b7a7d79da08fdf7eeff84860001b6106c7565b611d09836106ca565b505050565b611d3a7f3f65c84b03457e3ee1b8dc7b5969754f6a319e2179f7ab845c072fea305460de60001b6106c7565b611d667f970226277e5ad7ef90bf36ca69a65f599ff0b87476b2dfc4cba29e320a88813760001b6106c7565b611d927f1f3483dd07c2df7235ce7ecd7e5b4205ee84768a6a4e43c80f68447c06695d7a60001b6106c7565b611dbe7f6cda95a4c4f4d59f4e50e64b06724328d8affaf6e59f047abcc94c7209f0b0e560001b6106c7565b6000816020015167ffffffffffffffff1611611e0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e0690617dcb565b60405180910390fd5b611e3b7f2088cb317fe5d85b09ecf939c847b73b95f343455a3409bda7acff0cc6dbf6b460001b6106c7565b611e677f32bd0f72df208fe527b9fb07b8e5ff5e5e189dbd14e6e93e8df0c466bc9ab9b960001b6106c7565b611e937f3bd744b4f0cb9f34cfa56aede6d36cb7ccd73d6b8c4303265c16e3d1b286ceb760001b6106c7565b611ebf7fe0775fd8d8f2430bbf141665d8bdaaf4e2282ca8526e45620de11517655276ec60001b6106c7565b6000816040015167ffffffffffffffff1611611f10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f0790617d8b565b60405180910390fd5b611f3c7f16af40ccdacb55322dca2dfc7537677836823e85a5a78b634e17859175ecbd0160001b6106c7565b611f687feacc2a0f3b1bfdb218dbb55930b9dc2227bad9a604dac8761439659ede4ee4e960001b6106c7565b611f947f84931eac100ad095d3515b0a80e8ec111c4b87c84bc0b39f16b8fe10880f0e7360001b6106c7565b611fc07f012889bbce1ad80536f58d262d9fbd3197f08942d5d27e747c483d9134156dcb60001b6106c7565b6000600a9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636683899482600001516040518263ffffffff1660e01b815260040161201f9190617c16565b60206040518083038186803b15801561203757600080fd5b505afa15801561204b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061206f9190617501565b6120ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120a590617d6b565b60405180910390fd5b6120da7fdecc667777e173a614110ddde1a5ea32de1bda362cf603b28b4e8f0b5f59bb6e60001b6106c7565b6121067fc12256b0530319e6b3fd6b261a0ee830fd4475a355b853d707ca37d8588d0bb960001b6106c7565b6121327fa7ee6454538e2edb5bce610ab1bbfb09c60d36b1968c7d0e1acf34c4b82a621260001b6106c7565b61215e7f89fb0b169176bcf50d158e9fe0a442418fd29e68740803fdbf1df27a20a3a0e660001b6106c7565b60006121a56040518060400160405280846000015173ffffffffffffffffffffffffffffffffffffffff168152602001846020015167ffffffffffffffff16815250610fbd565b6040015167ffffffffffffffff16146121f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121ea90617d4b565b60405180910390fd5b61221f7f1ae08fc1ec360e382dcaa2dde8e822a19af148080ff09e6306d30f6f952e461b60001b6106c7565b61224b7fcc8eb0623751ec5f6e95f7bd3dd85551d6efdc1d3f556fdb3e1595827968b3b960001b6106c7565b6122777f8a15f61858ac420399bdb255083147d126ae45b7f29c167d2a50d7ea9f32779360001b6106c7565b600080600a9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe83600001516040518263ffffffff1660e01b81526004016122d79190617c16565b60e06040518083038186803b1580156122ef57600080fd5b505afa158015612303573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612327919061757c565b90506123557f149ba3cec578af3856f595695512ed2d7881924cb6aba748f0f2e31d2cf7544560001b6106c7565b6123817f1d17502b1a8e7db7980e0ce9f81dbe31cf7232d861ed11e658227052a22e2bf660001b6106c7565b6000612390836000015161425e565b90506123be7f2ecbbe89e21fa4116c5b324a39baedc78ed286d500201e6f2e022a7b0fbd7d3260001b6106c7565b6123ea7fb56014fc991d774b2cc23e58fd447758850c3929af6c8480dbc1af34d8d1b37860001b6106c7565b816040015167ffffffffffffffff168184604001516124099190617fb9565b67ffffffffffffffff1611156124c3576124457ffc6ef0bd0a104f1c184c325d5dda5fce31c3516710e10c5bb71f63fc228e707260001b6106c7565b6124717fba0414a5c4e0ea6a14ff9265a7ebada39bebe8bcdfe9f186f1a77afddcb9b01960001b6106c7565b81604001518184604001516124869190617fb9565b6040517f95c91e8e0000000000000000000000000000000000000000000000000000000081526004016124ba929190617e2d565b60405180910390fd5b6124ef7fc13739bf4dd9b9b36af6c777b4cf51326ffc2f7e9ad39b5a561b8a7ac2ad83f460001b6106c7565b61251b7fcb1eff8dede3e0033c17523fe739b0930827bb116127905dda931ab0b3672fc760001b6106c7565b6125477fe0db4d86d4f13ea560949911caaa6d9449fc19a1d27d4f0842f9839a28c196b460001b6106c7565b60016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff021916908360028111156126d9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff0219169083151502179055506101608201518160050190805190602001906127bf9291906164e2565b5050506127ee7f8afc9266de811521ab58cc1ee7be13c88cdcf47f7947611e1e043b34d0b2fe0c60001b6106c7565b61281a7f233a373e493c8bbff8c4b852481ec86d4eebac44dbce5395a9f0a685a42f51b560001b6106c7565b7fc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb600843856000015186602001518760800151886040015189610140015160405161286b9796959493929190617c98565b60405180910390a1505050565b600060019054906101000a900460ff166128a05760008054906101000a900460ff16156128a9565b6128a861486d565b5b6128e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128df90617dab565b60405180910390fd5b60008060019054906101000a900460ff161590508015612938576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6129647f9d241acda3cb91ae9aa77b6fd9724997960bf2da713992d5476de0f23a25184760001b6106c7565b6129907fd3ad23d4cda7cf17b4f7334a7fbf9bf35ce564c19cf5eda4c1c19ab49b79ab7060001b6106c7565b6129bc7f3e46baf3a1ae6f71ccd99daa20ca4ad7f96ad7e1a8c95b4b131ad62629234d3960001b6106c7565b816000600a6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015612a1e5760008060016101000a81548160ff0219169083151502179055505b5050565b612a4e7f0ec86240bfcd64c89bd79f242a6f4e5218bce65717193de26ed518b9c9808a3360001b6106c7565b612a7a7fafbd1daa00f23bf67920d1e31f2d6252a4095bf5e3517ebaf3781850d5d5007860001b6106c7565b612aa67f88954ea6c643a364b6368e36c9f0ef039824e65ebe790790bdc9dbbce792b67f60001b6106c7565b6000612aba8260000151836020015161487e565b9050612ae87f23cbe6edc71733af978599a3e038078cb03f991bc29d7981b471119230b3901760001b6106c7565b612b147f7aa4226e08770231c8fe569d00e719f9ca76d8573627c7d449aad1fc88a5c0f960001b6106c7565b80612baf57612b457fc17bae04e15157d03c2089f176c76c5c5cb19a8ff23148cae01cc2d87addd1b160001b6106c7565b612b717fc87dd4a2396fc5879de3f8ecbd9451a60148b715d68bc89f2502be4305e7537460001b6106c7565b60036040517fc8c84b2f000000000000000000000000000000000000000000000000000000008152600401612ba69190617d30565b60405180910390fd5b612bdb7f8a87b94961f8ede1121dea7439799293e2435baadae5db5749563ae05614f36d60001b6106c7565b612c077f300412e48bedd2ee422d86aaeca06b2fa0a982faf358dea823ee8255c986783860001b6106c7565b612c337f535908c7e33095dc6ec38daf846163093eb119323985bc1958e7e6b480946b0160001b6106c7565b612c41826000015183614bdb565b5050565b6060612c737f7d28183d6f64ec80db9b3e600be71e64adf9ee58cce72cc45960158ce6e3145160001b6106c7565b612c9f7fc345614eae689e5a50172d2a2e1b691d70be2c5b8fefb9136671dd325a11e5e560001b6106c7565b612ccb7f7d7d8b950be65c0724bee103a4e03a57fdb35b112631bde51e202d3bf41dedd260001b6106c7565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561306e5783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff166002811115612e85577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115612ebd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b82821015613057578382906000526020600020018054612fca906181f3565b80601f0160208091040260200160405190810160405280929190818152602001828054612ff6906181f3565b80156130435780601f1061301857610100808354040283529160200191613043565b820191906000526020600020905b81548152906001019060200180831161302657829003601f168201915b505050505081526020019060010190612fab565b505050508152505081526020019060010190612d2a565b505050509050919050565b60006130a77fbbd546b4ca8e58c5c7a3fbe8dd8c8a0cc6b6087073449211b8db38f267cb903a60001b6106c7565b6130d37f07bba7a72522c90b12fdf465cbeb5fc960c2a77a154f8da56bc58124c3872c6d60001b6106c7565b6130ff7f9c41522d6e0997cf823bb65b5068cad62ad2a196b5f2f460e03c0341bf7650c360001b6106c7565b600061310b8585614eda565b90506131397f87d5ec1a5a29c49325e018de65f6ab86b5cc57cbe2d7f5095c17b869e55ae28e60001b6106c7565b6131657fca85af5125cb6ecbbb5e9f654dd6f914cc8b4675d90cffa38644c1fc5378c0ae60001b6106c7565b60005b8167ffffffffffffffff168167ffffffffffffffff1610156137a9576131b07fbb2bd7f05fd2640aaf73d5c233471bf2ea0a4dcb12e2c9600f045aef9f14b88760001b6106c7565b6131dc7f50ab706e4617d5eb8de08acac13511755a0a4d423c30d93dc34f0d79a984bed360001b6106c7565b60006131e9878784615008565b90506132177f78e1f923dbe0b11b5658b4e63c8144af4dd8cf751e7db8bed0135fddd9a2e03b60001b6106c7565b6132437febc00f45981d674a3743bcc5ff5ee7829eadc2f8a8c7359553b1cb7c1dfd2b1360001b6106c7565b600080613250838861529d565b915091506132807ffd1a095e5351f4e3a841003648de95ec2d1eefb52b9cf0c6a83eb3610beba66560001b6106c7565b6132ac7fa691e72f136973af49f4fa51c0a0d977e8f13f20520a051c8b1716d670d274cd60001b6106c7565b8015613767576132de7ffb04c4e0fb6cfd659e95c7e4f22042013a5edd611bd39d00e4a9643c57b1eb5b60001b6106c7565b61330a7ffbbe3920bca96bdaba2e9b4340c8b073658422616e91f1c3280d61119763645560001b6106c7565b6133367f353f728d9f480136a7edc1ebc5ea94a94e798ed7901b33bb287288478b61a34260001b6106c7565b600060036000856000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208054905090506133957fd3f76268d24c345e12a1a6b663eb884ca24bb09a3575bafb82006844689d5bd860001b6106c7565b60036000856000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208367ffffffffffffffff1681548110613402577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160008082016000613421919061693b565b6001820160006101000a81549067ffffffffffffffff0219169055505061346a7f488ac80d0fb99a19fc961d64080000ef3c14efef2a4954dc1d720abec7fe8f2560001b6106c7565b6134967fb166b42de903af6b6be0fd46151b2cb6aab7987eb207f914876c76b8a0c5cec660001b6106c7565b60008367ffffffffffffffff16141561353b576134d57f3753de9ea28d2a57d29b8ff40c266203361496f6656f4cd6d359bbc094e068b860001b6106c7565b6135017f9dc0f72e827d527e04de4e3ac16fcd4b5b72e84be469493ec1b5aa829f19dfb860001b6106c7565b61352d7f706b559ce5426ea5ae75383382122ba5ef34f644ba6193b7a109cd99685c007d60001b6106c7565b878460200181905250613568565b6135677ff1cde91f775dc9e5a1bb46ee8ebc06139f53f481c404c24a92ea760599f547b660001b6106c7565b5b6135947f130e29db6e85ef26d23c5b8526c9744cfe3bb7632005d33ac534bac9c1e64f4c60001b6106c7565b6135c07ff69a8352b6aa4db37bb1fe79c04a1d9e064a88c6eeb7c37f5da33d8a797640b860001b6106c7565b6001816135cd9190618039565b8367ffffffffffffffff1614156136705761360a7f50ef4904e6cad02dc79e14eddba9e551fdd97615f54b7cc4f80cbafaddd3c08c60001b6106c7565b6136367f99b7c52368369cd7325b91636fd7b6a6603f5c8e22948511835a1c08e43cb70960001b6106c7565b6136627f70ecbcb3545d17c3807009eb8d4015025da94cfa7c3edee1cbedde8c5b7ed9d460001b6106c7565b87846040018190525061369d565b61369c7f6f4e2ac6f88d0352cc5ed946913f1f1efe1d2f1caa028309f3b5766daf57f21d60001b6106c7565b5b6136c97f5d46446705b389fcb9794236660e8125b8d5735988d49b9ef31be07f839b890060001b6106c7565b6136f57f100595cda781c30d289a21fdfaad4a6c8da80fae35332b97f56d97051856c8c460001b6106c7565b6137008a8a8661591e565b61372c7f9faac8b9e0720d8800490364229facc0315645d3261a2825ceccce218d595b8060001b6106c7565b6137587ffacdde5a900b43a79f1325b5fc658032bb5a5882e99f1ab58c684fd6ed8740bf60001b6106c7565b60019650505050505050613808565b6137937f9855253135f06ae48caa7bdb53e22239f8bd4ef4d7241ae6e7efbb09c4b5136b60001b6106c7565b50505080806137a19061829f565b915050613168565b506137d67f3ba3c0bae7dc1432dd1e175b42a031a1b959de568c1145dbedb7306f253f87e960001b6106c7565b6138027f28fe145b37402e181e4bd3c1b2d441baf1467b8bcf10b0aa549f5f6535dd028060001b6106c7565b60009150505b9392505050565b61383b7f7219c30c781337520f55c752a5b445db599be9721bba75ca4a9749c69f905d8e60001b6106c7565b6138677f74e0d2cdade352af02189ccc61fa04bbaf62875889de59fd620dbdb8e2cd7c4760001b6106c7565b6138937fbf83b2ae8399ac2e5750453260bbd031e58e2abb9c9630ecb2597494ae63b05b60001b6106c7565b61389d8282615ab7565b6138c97f103f6228748b4dd04ee7818be5e5b3ae51bcf6b8f2576f29b98b94487e80660060001b6106c7565b6138f57fecfb868b6ced0f454cc278053a96e4f81cf586152118413d9017a39f76cc570f60001b6106c7565b6138ff8282615c3c565b5050565b60006139317ffe15be7941cafa0a8b2fca5eff8b458a47fe7e846679b17cc9ef85fb52fcffdd60001b6106c7565b61395d7ff004f46c68cfe8cd630d55e335465d1ea108fff843ac5f80bf32f982abf15e4460001b6106c7565b6139897faeb85a1c3eae9fe199db5d1e4fc2a222b85f33d0f3cd85619e9aa8b6cb0ce90960001b6106c7565b61399161697b565b6139bd7f4b5fc591f43c8d1066a1d609fd8918b95fc1314bd2985eb1137019cd7e8d0f8e60001b6106c7565b6139e97f81ea11037364be1288001d7a16094eb5f6dbfafe813ab8e3bc4a904b7114bdbe60001b6106c7565b6000613a177f04dc0b91aec1d09f76491ed6e6e7d6e2d5d2cd2a6b5f38407dd33b1e7103da3560001b6106c7565b613a437fa1052197ae9f3df966568858e6b136e0fa835708beb068ac22d14459a0c6b7b260001b6106c7565b6000613a717fce1b94a0004782776a7f260590d5e9d6b1f31032de0b73c63d543879df6dd3b460001b6106c7565b613a9d7f6a1370d1d138502e35e6301884817e8c04a36db390a91e24e93f1fa5bdc16f0760001b6106c7565b6000613adc60405180604001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018967ffffffffffffffff16815250610fbd565b9050613b0a7f8c3d233ff1dd2a456cd874374d43ccc7850db1d10769a75ee88547f62f3a7a2060001b6106c7565b613b367ffd54be18c0512f6f0e383f20f2dbfdda83c7d31f18daf86a2dbb0db98cda127e60001b6106c7565b6060613b647f0bb186568fef3c3b035f1f42fcd85dde265ca1a02cbfad15bc56a4430d1f742060001b6106c7565b613b907f6801affa0d531a3ed14b1d33964ae5ffe0a55a5857f7da61999c0d2eca29e4e260001b6106c7565b8161012001519350613bc47f47fda20fce7b19a657e14f1e34187d27469b7bab997fd2d9a8bbf10dba60cc8160001b6106c7565b613bf07ff7771f6a63727796d36780efa2ca0f83cf1bb91145bd73733f85a94f9995ee2860001b6106c7565b60008467ffffffffffffffff161415613d6b57613c2f7f95798bf728399f3b5c61bc3ec0733d875f98c93313fd37f3b83fb6360739a01e60001b6106c7565b613c5b7f2ba9846f9e1c382634b9d28d0328d45a559c8e7a50345a2f0765710589c3748260001b6106c7565b613c877f4b57fb6d3fbcd009dbc662f7ee28c18ce902f103e717dd57c46e36547302c09960001b6106c7565b60019350613cb77f1fc9b12785bcf8c4455a7dcb03fa55e8366124c470c0bf7650dfbd1e8b4a444b60001b6106c7565b613ce37f9489a5ba767a56a08ab818d3716866bbf0a3f3a7bbd561b6963711d7289d316b60001b6106c7565b60019250613d137f63885d494b06d40fdaab29072a8e2389b3ac36482994221aa674b138e711e42e60001b6106c7565b613d3f7f036270ea840fb0833bd9bc13a6fb0fbf4e6ee8af23fd67827768298a564b494060001b6106c7565b60405180606001604052808567ffffffffffffffff168152602001828152602001828152509450614193565b613d977fbc2a2cbc25169ffc9f7b24b8c16df0d9a0d087bb4ac6c49cc28f8680690a70c960001b6106c7565b613dc37f9b9911afbeb5bf3cc48f022f2cdce7f8b7e5086aaeb390423aaf22d1b11d380060001b6106c7565b613def7faca8e5d4df0c77ce9bf2c37efbd7dcdbc270a03a5752820197d6882dfdd8bb1f60001b6106c7565b613dfa898986615008565b9450613e287f13318de186d1d43942b6f7647a035ade4b9f3c5bf767a31cf76d3d0cc58ea98c60001b6106c7565b613e547f13388afff1a17fc7509847ad2c619975756f43624630f79995d548dcce39038f60001b6106c7565b600060036000876000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015613f975783829060005260206000209060020201604051806040016040529081600082018054613ed4906181f3565b80601f0160208091040260200160405190810160405280929190818152602001828054613f00906181f3565b8015613f4d5780601f10613f2257610100808354040283529160200191613f4d565b820191906000526020600020905b815481529060010190602001808311613f3057829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190613ea1565b505050509050613fc97f27d4288e3a25f0c908e2cc7992b9b50fead82eb4f3e718ffde62e3daf992862560001b6106c7565b613ff57f01bece67bceca80702f8fcff5fa98116d496919e3b26376e8c107ad7e1a59fb460001b6106c7565b600060029054906101000a900467ffffffffffffffff1667ffffffffffffffff16815114156141645761404a7f6c63f7e1823190802ecc0867993c42e1fac591fedb459005351d3ca685dc179a60001b6106c7565b6140767f7cf8135c47bac0fae4815286e05fe2f0b671d019be33ec6a05d0f45de392a84a60001b6106c7565b84806140819061829f565b9550506140b07f1c067c484456eafc36c92a28f1dfc668af794a9bd41908c83752b5c3e10fc9ae60001b6106c7565b6140dc7fe6246d44d7ec74722c93ab922d30ab08921af1369c86f4ef33f9293294d9f76060001b6106c7565b6001935061410c7f0520949701a9afbd50ad72c71b829f166a84e807564612846851d3144b57413e60001b6106c7565b6141387fc062cca1232ef3c74f8e221d7ba32ff4f3f060549986e2feef259a5f9ec5c81f60001b6106c7565b60405180606001604052808667ffffffffffffffff168152602001838152602001838152509550614191565b6141907f7704da39fda0192d6ed64c91f96d0332cce7e68e073920dc06e3957a1f9881dd60001b6106c7565b5b505b6141bf7fe85f0e3b5b574800811a6d7a3eacb82c01d16576eea1a90667d5de62c4a4d68260001b6106c7565b6141eb7f268651888a32456a957a41c0f974f5f790653cd0517734db2f526d4b9ca4b74c60001b6106c7565b6141f78989878a6160a1565b6142237f9176715d3817467dc52c4964006806dc3788c4bc0933995b696d520981db7b2060001b6106c7565b61424f7f04645f0a5c5593ef1ea01702d558a84050ccf769a2970056c9133ec4126c8aed60001b6106c7565b82955050505050509392505050565b600061428c7fae6e8e32b20d696ebe3dc5e9d179608fe563e65cc66421584d4ab24f76f5a6b660001b6106c7565b6142b87f4d33fe4d0c5dc1b6b90e005d9a75b885f4d64147001d55e79c26971d7f9f61d760001b6106c7565b6142e47f6bc29a05f6d0f227310b9d16fc6b371b1e3aaa27f3711ae2d2942c88585def0e60001b6106c7565b60006143127f98a2ba0c4f3f8f5c1ad77e1ca6d302cae8044505ba4a08e5496b52ca741bf5e260001b6106c7565b61433e7fafa2c992801a26149525d135ae4d5cad412d81465fe5460e9c25a8423e5d808760001b6106c7565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156146e35783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff1660028111156144fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115614532577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156146cc57838290600052602060002001805461463f906181f3565b80601f016020809104026020016040519081016040528092919081815260200182805461466b906181f3565b80156146b85780601f1061468d576101008083540402835291602001916146b8565b820191906000526020600020905b81548152906001019060200180831161469b57829003601f168201915b505050505081526020019060010190614620565b50505050815250508152602001906001019061439f565b5050505090506147157f3addb1a0993289ff4e2a99b00633787233b715d361bb32ec682badc68340391e60001b6106c7565b6147417f1f99e6fbde8df0e716b94d50c369ad61afb96c7f37627ef39a418ea22b72a02660001b6106c7565b60005b815181101561480a576147797fe83baa9c930b0cf8f351c1c9edda5f4fdbae6626b95a6e73bdc7af3702b40e1060001b6106c7565b6147a57f0f58d24cdc473993c3760e934f14269c30cc8608236a4b7815f4a9dfae63b12860001b6106c7565b8181815181106147de577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160400151836147f59190617fb9565b9250808061480290618256565b915050614744565b506148377f73c509a1dd49c36f3ad01cf45043bce236af918fb54736371209b97b38f953bc60001b6106c7565b6148637fa6b73cbb67995daee3241670e9754a41e6b167ed0d0dfdd63ab955315ae2291560001b6106c7565b8192505050919050565b600061487830616336565b15905090565b60006148ac7f7b9ec74a6d5bae05b733d0cf32d5b06704f5162a1b39861527f567246489c74860001b6106c7565b6148d87f45b5f260c01924b817d24c8fa6c4164207a32a5f07b7e022717afd537308dc4460001b6106c7565b6149047fe6ab4fd87ce299c46fd8c092c0eee85333e8189b1cf2e4db412f34c54e475d0260001b6106c7565b600061490f84612c45565b905061493d7fd95a29d5b248a0cc338ee13a9c2ddd27dfcf2bc8fb7b77aab7cbde6be1a3678560001b6106c7565b6149697feb6ff4957f49ed0a62be45652e831d37f40f211806fc65a16a138f1095206d4960001b6106c7565b60005b8151811015614b76576149a17f647e3d2f29200b4a6e9f1eb39862439a27418d79c2f77261391f55ed3d555f4860001b6106c7565b6149cd7ff378f63cfa867f3664db76e2db3d2c21f29b80cdb572e3a9d6c17c40a46478a160001b6106c7565b8473ffffffffffffffffffffffffffffffffffffffff16828281518110614a1d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff16148015614aa357508367ffffffffffffffff16828281518110614a8b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff16145b15614b3757614ad47fdd481cd666cef0d7b6b9a742577a1740f5ceab85dec307f79214f72572a482ee60001b6106c7565b614b007fa4566dd704fa2666802f708b15152e4b13d62a37273c70dd9a7136fc7d1dc93060001b6106c7565b614b2c7f69b0b71d9c6e75198962e471db3f34573b3d2389d9955af7e1a832680dc555a560001b6106c7565b600192505050614bd5565b614b637f781cae78c8e517a6513e12c5472315d4d5e9bac1ba1b47c81639fc257174555e60001b6106c7565b8080614b6e90618256565b91505061496c565b50614ba37f0631f214fa9dd0c80b52704432b4e1148274ece7205fd78365873b022f186b0a60001b6106c7565b614bcf7f093fc8452f0edd17d53dc169dfd4882f0080817ddc9bf54f36e4796f4905ed7960001b6106c7565b60009150505b92915050565b614c077fc5f9790451f67a50dd162f174b75050b4aa4006e1418c918fc721da2d803175860001b6106c7565b614c337fcf44cc969a3aece8cdf31ceea2360446b0e0fde9462dd87b6b9e52f2e72f166860001b6106c7565b614c5f7ffd86ba8091dc033eefbe760e285bbb70ee6718c8a909000ff9cc9402c71b9af260001b6106c7565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115614ded577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190614ed39291906164e2565b5050505050565b6000614f087ffc90b2a44a0a3ca8f33581dca89bf6aac58c545837134eb3d3e317fa700e726f60001b6106c7565b614f347fc827d34dd8d658eeee0115a8862bd8f8673e88faf2fdab031259bd55be46396160001b6106c7565b614f607f3e42eee18224f16272817a5a06388b0582111a41ea5da0bfaf1630b586f28e9260001b6106c7565b6000614f9f60405180604001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff16815250610fbd565b9050614fcd7f840eaff1b93d07d7b684dd69c1af2f791c7a90cd4854c6507646dfd4fb554a2960001b6106c7565b614ff97f645eadf21c8ca050bc83ffcd3e887ed1739c8da58172bb7d7087299fc198998160001b6106c7565b80610120015191505092915050565b61501061697b565b61503c7f9e6a03c000a81d9b92a5dc2045dcef7eb64381708572dffe96072c7c51ff07b560001b6106c7565b6150687fc46485066475e0d446d7cd4c915759bb8faed1a8bf4e347a4ede19b26656d51b60001b6106c7565b6150947fb24f94b22e81afaba1cc93e3759ff237ed9beb4e2c31c7adcf476c5e0e47008f60001b6106c7565b60008484846040516020016150ab93929190617bc2565b60405160208183030381529060405290506150e87f1699877131bcec2ccc42e6c733c4d6b2e47d6cd1484dcec23e8d6d0ce323f40c60001b6106c7565b6151147fad73b3102144104c71f8412733a243b462d5923c459b2265af0915c083c0935560001b6106c7565b6002816040516151249190617bff565b90815260200160405180910390206040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461517f906181f3565b80601f01602080910402602001604051908101604052809291908181526020018280546151ab906181f3565b80156151f85780601f106151cd576101008083540402835291602001916151f8565b820191906000526020600020905b8154815290600101906020018083116151db57829003601f168201915b50505050508152602001600282018054615211906181f3565b80601f016020809104026020016040519081016040528092919081815260200182805461523d906181f3565b801561528a5780601f1061525f5761010080835404028352916020019161528a565b820191906000526020600020905b81548152906001019060200180831161526d57829003601f168201915b5050505050815250509150509392505050565b6000806152cc7f7c796f4fd653179b6e7ce8c6901cc4a49c36409270450bd442115916bdc5214d60001b6106c7565b6152f87f1a6ccfb704e932f121b461e4f9917013f84e4e967661ef07fa4c3ed1238b051b60001b6106c7565b6153247fed9d50eefe5fce550359e2042ef6fa0ab0055a4290bfa0558497faf944fdb65060001b6106c7565b600060036000866000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561546757838290600052602060002090600202016040518060400160405290816000820180546153a4906181f3565b80601f01602080910402602001604051908101604052809291908181526020018280546153d0906181f3565b801561541d5780601f106153f25761010080835404028352916020019161541d565b820191906000526020600020905b81548152906001019060200180831161540057829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190615371565b5050505090506154997fe8e9d22bf1a7d512a0913fef4c182892d38a21ce6635d192705264ea5d3b3fda60001b6106c7565b6154c57f7357d17652acbb4889a4900e11454ba1b1224d3aae4a21a28389f6f8b4e120aa60001b6106c7565b600081511415615560576154fb7fb8f4fe738f9b11733d37e00eb2a9fae416f2c7b3888533a917425d2680469b6560001b6106c7565b6155277f1f213faface3c19439795092c5dcc13455c6eceef95fb04bac9cef5a5614880b60001b6106c7565b6155537f5713f6375ab83abf367c9f340528da7428bc3fc82dc7817bd4077ecd4de9d71660001b6106c7565b6000809250925050615917565b61558c7fc0a2c8027d1be3707a4b7ecb4ac43d4fda271daeff1f4ec35dae44ebf55db09760001b6106c7565b6155b87f495e56b3c4f2b09c1c48505e53c4d9132b8b1584948740c0c2f5730c078d0e8d60001b6106c7565b6155e47f2b73c931aa25c770283babb76eadc0659bea8425d20c17611cb4259e3cbfae8260001b6106c7565b83805190602001208560200151805190602001201461568e576156297f43d9bb361795ae26f72509b6037d1d6f30163201d94ffb9b3ca3d65488639e2e60001b6106c7565b6156557f3239d6be6dda14ac9fd7cc2ab71c1410287bb1018fd51f39c4449d95fb3af88d60001b6106c7565b6156817f447961f1f80b0a7090fd9d7be22a9a618d397efa2e3039c96b99d610cc40ecc360001b6106c7565b6000809250925050615917565b6156ba7f1c28798eca951252a7d13a8344361b362fd9b8771e962df4afdf1e15655e065b60001b6106c7565b6156e67f056e9410f1b889d3557d1653ebf626594cd8fc75880f9270cc6c5c4ad2b3f8ee60001b6106c7565b6157127f2284870b8b1affffbbd4ded57ca189d44134321ec0094e40c1256302c94e14f860001b6106c7565b60005b81518167ffffffffffffffff1610156158b5576157547fdac88d5ce47849788b6c5e32a9726d7bbb1bb2bdd7b02ffce7acbcb87c58273660001b6106c7565b6157807f07c61050d6945fdd7819b6b33377d6f2f48759bb97dc0b6939dc46afeff2731b60001b6106c7565b8480519060200120828267ffffffffffffffff16815181106157cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160000151805190602001201415615876576158107fc8f78c84b097ad198920893d4f53a0509a165b625f1de787383029a24a336dd760001b6106c7565b61583c7fa66dcf269ff77c63e90566a75c57469f4aa9290518e1367f377b7accdd816b3e60001b6106c7565b6158687fb6792e31b3fdb9c1414e86aed3de6b1fed6320922e5883bee4c5ccbb29931c0760001b6106c7565b806001935093505050615917565b6158a27fbd39394156306901aa9af11890ee930e0b7af308e791712bab7040f116e82a1460001b6106c7565b80806158ad9061829f565b915050615715565b506158e27ff7e99ea508feb050a2d450ad40327b58de7a6736e52bc1efd9b047691baf7c9460001b6106c7565b61590e7ff59bb7618b53e585c1431d1503ad6a9f90c32a6f7b95d8c0c6e2bf8e2eaee0af60001b6106c7565b60008092509250505b9250929050565b61594a7fa3621aae50e00786bf415e2e2d28823d0eb36e0c905c4a7c3b4106542a3a4c5760001b6106c7565b6159767ff034141d625caf774e26d9fb6795a1dc9c060caeb93d25bcbc2e2acd7ebc75b660001b6106c7565b6159a27f1e153624cad8a2723628b56bdc82f8daef69e1a60e687d985213cc7c9520fb9260001b6106c7565b6000838383600001516040516020016159bd93929190617bc2565b60405160208183030381529060405290506159fa7f97b2abc29a4bcfcb4d242ec6e1d5a036cf36eaebc26080a7e06006c23dfab34e60001b6106c7565b615a267f8a57d9b8646c215eb9082de2f87984a5295b3186e8e3adda4180f1fb6785713e60001b6106c7565b81600282604051615a379190617bff565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151816001019080519060200190615a909291906169a6565b506040820151816002019080519060200190615aad9291906169a6565b5090505050505050565b615ae37fb09cd63be178679f2a2eb19d4569facb58db5706bb8b64b4ccf203316e33b67c60001b6106c7565b615b0f7fa9dc3235459080e1fc665e42700c646798f73c047df2acf7770dcafbc58d809c60001b6106c7565b615b3b7fa57325d810ce13eaad3e0238796c92b38d4caa139dbab291fe292babf307c8a060001b6106c7565b6000615b478383614eda565b9050615b757fc1d6ef36c034d9c82ec5246e2e22e1458bbf8b155d5867a3135aaa1bb646b7d060001b6106c7565b615ba17fae5a6958c3ed7cf1fbabc62d99a55df18689a6aabf7f3a6f07562bab7f16a83060001b6106c7565b60005b8167ffffffffffffffff168167ffffffffffffffff161015615c3657615bec7f51d3ad352dffbb219284a862da131f3941f8b5c8946981189a2b672d1836d13160001b6106c7565b615c187f38c5a4a4bca5bb899c56027e15fcb7dec61b629afd815ff7d1e2156c2c86fd6860001b6106c7565b615c23848483616349565b8080615c2e9061829f565b915050615ba4565b50505050565b615c687ff12767c815a2799a54f5dcd3d1744757535cf5ab46a554a9def058c4f2297c5560001b6106c7565b615c947f7d35758dd7f5de269b84df640b5d49193676dbbe59480dfe2dfdc3deedafd84060001b6106c7565b615cc07fc0fee63bafa13d2cec230f0cee4242af64f498a635a5d8e49c18f6d97f470f8760001b6106c7565b60005b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508167ffffffffffffffff16101561609c57615d437f76758cf24ba199c7b6c6b3ff9d9e27bd6065fe96e8acdadda6c7bf33c007114860001b6106c7565b615d6f7f442ccbb45816882a67baea0189ad130d085424e98d65d89f1df41af5f8c8552c60001b6106c7565b8167ffffffffffffffff16600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208267ffffffffffffffff1681548110615dfb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16141561605d57615e5d7fe67d3e4d917d418322c088123a49f7857874a09d1c5a6afb5c47d9b27060525660001b6106c7565b615e897f9c682b2292b6b73653e0024d31742a34f7d362eb1ec7801fb2dd7375dcfbbd3f60001b6106c7565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208167ffffffffffffffff1681548110615f0a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060060201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff021916905560058201600061602a9190616a2c565b50506160587f44ab4aaf9c9e937a9ff5813e7a3f070ecb0143d8293dd11f24bb1fb6adabb68660001b6106c7565b61609c565b6160897ffc7a93ef14375293879373fbde5ae2f3beb9671fd8ae8371147985e3ec44b6b560001b6106c7565b80806160949061829f565b915050615cc3565b505050565b6160cd7f9b2a2cd9a3e2b0f8ad5273f24efc86152bd8bdae117d558e2650eebe67ad4c2460001b6106c7565b6160f97ffaf37c014b854fe778486fb4e46d3965cda25887a8559f9e839ca2680690e66160001b6106c7565b6161257f85de4f30ba1049eab69f0afce5d8f2b4936dd8facab8a9c9833e8027932c9d4460001b6106c7565b60008484846000015160405160200161614093929190617bc2565b604051602081830303815290604052905061617d7f070083e9614eecb253ff9784d9be5b4b53e05bca773e070d9f96454991c7040f60001b6106c7565b6161a97f39633b1cedefaa2c00234699518df64e3a0af47098f8bf80db8bc7b9fa9aa9b760001b6106c7565b826002826040516161ba9190617bff565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160010190805190602001906162139291906169a6565b5060408201518160020190805190602001906162309291906169a6565b509050506162607fd82166d586970dbb2b6382d5cf5312acd0eb3122786d28bce66b321ae24db2bb60001b6106c7565b61628c7f5f512d484bc0d8f3de349b8e3c03dd2d65445caeacab0da4e1d2fba8c95263dc60001b6106c7565b60036000846000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002082908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000190805190602001906162fd9291906169a6565b5060208201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b600080823b905060008111915050919050565b6163757f8dba24be634a7a004c37a5e64c061e492b231517f3b3d66bc95bf4e2279ef9a360001b6106c7565b6163a17f89c72acccf5209b7891d0356a88cf55247158eeb57edb95f208bbedd2692a71d60001b6106c7565b6163cd7f41b6857935051e34bfe024152599187f1da412f53bb89b5219bdaf496f14d4c060001b6106c7565b60008383836040516020016163e493929190617bc2565b60405160208183030381529060405290506164217fdd22ffc7f39a1291d79cb9383d49658dfaaf31faa33d2836fed39b2527fe470160001b6106c7565b6002816040516164319190617bff565b9081526020016040518091039020600080820160006101000a81549067ffffffffffffffff021916905560018201600061646b919061693b565b60028201600061647b919061693b565b50506164a97fb5bf9e4c61f2238fe5023a64ee095a2be066294584723c63e2c5081f027ef27e60001b6106c7565b600360008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006164dc9190616a4d565b50505050565b828054828255906000526020600020908101928215616531579160200282015b828111156165305782518290805190602001906165209291906169a6565b5091602001919060010190616502565b5b50905061653e9190616a71565b5090565b82805482825590600052602060002090600602810192821561683d5760005260206000209160060282015b8281111561683c5782826000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900467ffffffffffffffff168160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160089054906101000a900467ffffffffffffffff168160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160109054906101000a900460ff168160010160106101000a81548160ff02191690836002811115616705577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060028201548160020155600382015481600301556004820160009054906101000a900467ffffffffffffffff168160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160089054906101000a900467ffffffffffffffff168160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160109054906101000a900467ffffffffffffffff168160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160189054906101000a900460ff168160040160186101000a81548160ff021916908315150217905550600582018160050190805461682a929190616a95565b5050509160060191906006019061656d565b5b50905061684a9190616afd565b5090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600060028111156168e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b508054616947906181f3565b6000825580601f106169595750616978565b601f0160209004906000526020600020908101906169779190616c25565b5b50565b6040518060600160405280600067ffffffffffffffff16815260200160608152602001606081525090565b8280546169b2906181f3565b90600052602060002090601f0160209004810192826169d45760008555616a1b565b82601f106169ed57805160ff1916838001178555616a1b565b82800160010185558215616a1b579182015b82811115616a1a5782518255916020019190600101906169ff565b5b509050616a289190616c25565b5090565b5080546000825590600052602060002090810190616a4a9190616a71565b50565b5080546000825560020290600052602060002090810190616a6e9190616c42565b50565b5b80821115616a915760008181616a88919061693b565b50600101616a72565b5090565b828054828255906000526020600020908101928215616aec5760005260206000209182015b82811115616aeb578282908054616ad0906181f3565b616adb929190616c84565b5091600101919060010190616aba565b5b509050616af99190616a71565b5090565b5b80821115616c2157600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff0219169055600582016000616c189190616a2c565b50600601616afe565b5090565b5b80821115616c3e576000816000905550600101616c26565b5090565b5b80821115616c805760008082016000616c5c919061693b565b6001820160006101000a81549067ffffffffffffffff021916905550600201616c43565b5090565b828054616c90906181f3565b90600052602060002090601f016020900481019282616cb25760008555616d00565b82601f10616cc35780548555616d00565b82800160010185558215616d0057600052602060002091601f016020900482015b82811115616cff578254825591600101919060010190616ce4565b5b509050616d0d9190616c25565b5090565b6000616d24616d1f84617e7b565b617e56565b90508083825260208201905082856020860282011115616d4357600080fd5b60005b85811015616d735781616d598882616e41565b845260208401935060208301925050600181019050616d46565b5050509392505050565b6000616d90616d8b84617ea7565b617e56565b90508083825260208201905082856020860282011115616daf57600080fd5b60005b85811015616df957813567ffffffffffffffff811115616dd157600080fd5b808601616dde8982616efe565b85526020850194506020840193505050600181019050616db2565b5050509392505050565b6000616e16616e1184617ed3565b617e56565b905082815260208101848484011115616e2e57600080fd5b616e39848285618187565b509392505050565b600081359050616e5081618531565b92915050565b600081519050616e6581618531565b92915050565b600082601f830112616e7c57600080fd5b8135616e8c848260208601616d11565b91505092915050565b600082601f830112616ea657600080fd5b8135616eb6848260208601616d7d565b91505092915050565b600081359050616ece81618548565b92915050565b600081519050616ee381618548565b92915050565b600081359050616ef88161855f565b92915050565b600082601f830112616f0f57600080fd5b8135616f1f848260208601616e03565b91505092915050565b600081359050616f3781618576565b92915050565b600081359050616f4c8161858d565b92915050565b600081359050616f618161859d565b92915050565b60006103208284031215616f7a57600080fd5b616f856102e0617e56565b9050600082013567ffffffffffffffff811115616fa157600080fd5b616fad84828501616efe565b6000830152506020616fc184828501616e41565b602083015250604082013567ffffffffffffffff811115616fe157600080fd5b616fed84828501616efe565b6040830152506060617001848285016174ae565b6060830152506080617015848285016174ae565b60808301525060a0617029848285016174ae565b60a08301525060c061703d848285016174ae565b60c08301525060e0617051848285016174ae565b60e08301525061010061706684828501617499565b6101008301525061012061707c848285016174ae565b61012083015250610140617092848285016174ae565b6101408301525061016082013567ffffffffffffffff8111156170b457600080fd5b6170c084828501616efe565b610160830152506101806170d6848285016174ae565b610180830152506101a06170ec84828501617499565b6101a0830152506101c061710284828501616ebf565b6101c0830152506101e061711884828501616f52565b6101e08301525061020061712e848285016174ae565b6102008301525061022082013567ffffffffffffffff81111561715057600080fd5b61715c84828501616e6b565b6102208301525061024082013567ffffffffffffffff81111561717e57600080fd5b61718a84828501616e6b565b6102408301525061026082013567ffffffffffffffff8111156171ac57600080fd5b6171b884828501616efe565b610260830152506102806171ce84828501616f3d565b610280830152506102a06171e484828501616ebf565b6102a0830152506102c06171fa848285016172b7565b6102c08301525092915050565b600060e0828403121561721957600080fd5b61722360e0617e56565b90506000617233848285016174c3565b6000830152506020617247848285016174c3565b602083015250604061725b848285016174c3565b604083015250606061726f848285016174c3565b6060830152506080617283848285016174c3565b60808301525060a061729784828501616e56565b60a08301525060c06172ab84828501616e56565b60c08301525092915050565b6000606082840312156172c957600080fd5b6172d36060617e56565b905060006172e3848285016174ae565b60008301525060206172f7848285016174ae565b602083015250604061730b848285016174ae565b60408301525092915050565b6000610180828403121561732a57600080fd5b617335610180617e56565b9050600061734584828501616e41565b6000830152506020617359848285016174ae565b602083015250604061736d848285016174ae565b6040830152506060617381848285016174ae565b606083015250608061739584828501616f3d565b60808301525060a06173a984828501617499565b60a08301525060c06173bd84828501617499565b60c08301525060e06173d1848285016174ae565b60e0830152506101006173e6848285016174ae565b610100830152506101206173fc848285016174ae565b6101208301525061014061741284828501616ebf565b6101408301525061016082013567ffffffffffffffff81111561743457600080fd5b61744084828501616e95565b6101608301525092915050565b60006040828403121561745f57600080fd5b6174696040617e56565b9050600061747984828501616e41565b600083015250602061748d848285016174ae565b60208301525092915050565b6000813590506174a8816185ad565b92915050565b6000813590506174bd816185c4565b92915050565b6000815190506174d2816185c4565b92915050565b6000602082840312156174ea57600080fd5b60006174f884828501616e41565b91505092915050565b60006020828403121561751357600080fd5b600061752184828501616ed4565b91505092915050565b60006020828403121561753c57600080fd5b600061754a84828501616ee9565b91505092915050565b60006020828403121561756557600080fd5b600061757384828501616f28565b91505092915050565b600060e0828403121561758e57600080fd5b600061759c84828501617207565b91505092915050565b6000602082840312156175b757600080fd5b600082013567ffffffffffffffff8111156175d157600080fd5b6175dd84828501617317565b91505092915050565b600080604083850312156175f957600080fd5b600083013567ffffffffffffffff81111561761357600080fd5b61761f85828601617317565b925050602083013567ffffffffffffffff81111561763c57600080fd5b61764885828601616f67565b9150509250929050565b60006040828403121561766457600080fd5b60006176728482850161744d565b91505092915050565b600061768783836177e0565b905092915050565b600061769b8383617967565b905092915050565b6176ac816180a1565b82525050565b6176bb816180a1565b82525050565b6176d26176cd826180a1565b6182d0565b82525050565b60006176e382617f24565b6176ed8185617f6a565b9350836020820285016176ff85617f04565b8060005b8581101561773b578484038952815161771c858261767b565b945061772783617f50565b925060208a01995050600181019050617703565b50829750879550505050505092915050565b600061775882617f2f565b6177628185617f7b565b93508360208202850161777485617f14565b8060005b858110156177b05784840389528151617791858261768f565b945061779c83617f5d565b925060208a01995050600181019050617778565b50829750879550505050505092915050565b6177cb816180b3565b82525050565b6177da816180b3565b82525050565b60006177eb82617f3a565b6177f58185617f8c565b9350617805818560208601618196565b61780e816183c2565b840191505092915050565b6178228161813f565b82525050565b61783181618151565b82525050565b61784081618151565b82525050565b61784f81618163565b82525050565b61785e81618175565b82525050565b600061786f82617f45565b6178798185617fae565b9350617889818560208601618196565b80840191505092915050565b60006178a2601583617f9d565b91506178ad826183ed565b602082019050919050565b60006178c5601383617f9d565b91506178d082618416565b602082019050919050565b60006178e8601483617f9d565b91506178f38261843f565b602082019050919050565b600061790b602e83617f9d565b915061791682618468565b604082019050919050565b600061792e601183617f9d565b9150617939826184b7565b602082019050919050565b6000617951601f83617f9d565b915061795c826184e0565b602082019050919050565b60006101808301600083015161798060008601826176a3565b5060208301516179936020860182617b8d565b5060408301516179a66040860182617b8d565b5060608301516179b96060860182617b8d565b5060808301516179cc6080860182617828565b5060a08301516179df60a0860182617b6f565b5060c08301516179f260c0860182617b6f565b5060e0830151617a0560e0860182617b8d565b50610100830151617a1a610100860182617b8d565b50610120830151617a2f610120860182617b8d565b50610140830151617a446101408601826177c2565b50610160830151848203610160860152617a5e82826176d8565b9150508091505092915050565b600061018083016000830151617a8460008601826176a3565b506020830151617a976020860182617b8d565b506040830151617aaa6040860182617b8d565b506060830151617abd6060860182617b8d565b506080830151617ad06080860182617828565b5060a0830151617ae360a0860182617b6f565b5060c0830151617af660c0860182617b6f565b5060e0830151617b0960e0860182617b8d565b50610100830151617b1e610100860182617b8d565b50610120830151617b33610120860182617b8d565b50610140830151617b486101408601826177c2565b50610160830151848203610160860152617b6282826176d8565b9150508091505092915050565b617b7881618121565b82525050565b617b8781618121565b82525050565b617b968161812b565b82525050565b617ba58161812b565b82525050565b617bbc617bb78261812b565b6182f4565b82525050565b6000617bce82866176c1565b601482019150617bde8285617bab565b600882019150617bee8284617bab565b600882019150819050949350505050565b6000617c0b8284617864565b915081905092915050565b6000602082019050617c2b60008301846176b2565b92915050565b60006020820190508181036000830152617c4b818461774d565b905092915050565b6000608082019050617c686000830187617819565b617c756020830186617b7e565b617c8260408301856176b2565b617c8f6060830184617b9c565b95945050505050565b600060e082019050617cad600083018a617819565b617cba6020830189617b7e565b617cc760408301886176b2565b617cd46060830187617b9c565b617ce16080830186617837565b617cee60a0830185617b9c565b617cfb60c08301846177d1565b98975050505050505050565b6000604082019050617d1c6000830185617846565b617d296020830184617b9c565b9392505050565b6000602082019050617d456000830184617855565b92915050565b60006020820190508181036000830152617d6481617895565b9050919050565b60006020820190508181036000830152617d84816178b8565b9050919050565b60006020820190508181036000830152617da4816178db565b9050919050565b60006020820190508181036000830152617dc4816178fe565b9050919050565b60006020820190508181036000830152617de481617921565b9050919050565b60006020820190508181036000830152617e0481617944565b9050919050565b60006020820190508181036000830152617e258184617a6b565b905092915050565b6000604082019050617e426000830185617b9c565b617e4f6020830184617b9c565b9392505050565b6000617e60617e71565b9050617e6c8282618225565b919050565b6000604051905090565b600067ffffffffffffffff821115617e9657617e95618393565b5b602082029050602081019050919050565b600067ffffffffffffffff821115617ec257617ec1618393565b5b602082029050602081019050919050565b600067ffffffffffffffff821115617eee57617eed618393565b5b617ef7826183c2565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000617fc48261812b565b9150617fcf8361812b565b92508267ffffffffffffffff03821115617fec57617feb618306565b5b828201905092915050565b60006180028261812b565b915061800d8361812b565b92508167ffffffffffffffff048311821515161561802e5761802d618306565b5b828202905092915050565b600061804482618121565b915061804f83618121565b92508282101561806257618061618306565b5b828203905092915050565b60006180788261812b565b91506180838361812b565b92508282101561809657618095618306565b5b828203905092915050565b60006180ac82618101565b9050919050565b60008115159050919050565b6000819050919050565b60006180d4826180a1565b9050919050565b60008190506180e982618509565b919050565b60008190506180fc8261851d565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061814a826180db565b9050919050565b600061815c826180ee565b9050919050565b600061816e8261812b565b9050919050565b60006181808261812b565b9050919050565b82818337600083830152505050565b60005b838110156181b4578082015181840152602081019050618199565b838111156181c3576000848401525b50505050565b60006181d48261812b565b915060008214156181e8576181e7618306565b5b600182039050919050565b6000600282049050600182168061820b57607f821691505b6020821081141561821f5761821e618364565b5b50919050565b61822e826183c2565b810181811067ffffffffffffffff8211171561824d5761824c618393565b5b80604052505050565b600061826182618121565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561829457618293618306565b5b600182019050919050565b60006182aa8261812b565b915067ffffffffffffffff8214156182c5576182c4618306565b5b600182019050919050565b60006182db826182e2565b9050919050565b60006182ed826183e0565b9050919050565b60006182ff826183d3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f736563746f7220616c7265616479206578697374730000000000000000000000600082015250565b7f6e6f6465206e6f74207265676973746572656400000000000000000000000000600082015250565b7f736563746f722073697a652069732077726f6e67000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f736563746f7249642069732077726f6e67000000000000000000000000000000600082015250565b7f736563746f724964206d7573742062652067726561746572207468616e203000600082015250565b600a811061851a57618519618335565b5b50565b6003811061852e5761852d618335565b5b50565b61853a816180a1565b811461854557600080fd5b50565b618551816180b3565b811461855c57600080fd5b50565b618568816180bf565b811461857357600080fd5b50565b61857f816180c9565b811461858a57600080fd5b50565b6003811061859a57600080fd5b50565b600281106185aa57600080fd5b50565b6185b681618121565b81146185c157600080fd5b50565b6185cd8161812b565b81146185d857600080fd5b5056fea2646970667358221220abd42d4e4709ff7114e51a7c62e113bc202bdeb1dda170840d74394a2f33a84a64736f6c63430008040033",
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _node) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _node)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _node) returns()
func (_Store *StoreSession) Initialize(_node common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _node)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _node) returns()
func (_Store *StoreTransactorSession) Initialize(_node common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _node)
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
