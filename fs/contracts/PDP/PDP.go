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

// Challenge is an auto generated low-level Go binding around an user-defined struct.
type Challenge struct {
	Index uint32
	Rand  uint32
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

// GenChallengeParams is an auto generated low-level Go binding around an user-defined struct.
type GenChallengeParams struct {
	WalletAddr   common.Address
	HashValue    []byte
	FileBlockNum uint64
	ProveNum     uint64
}

// MerkleNode is an auto generated low-level Go binding around an user-defined struct.
type MerkleNode struct {
	Layer uint64
	Index uint64
	Hash  []byte
}

// MerklePath is an auto generated low-level Go binding around an user-defined struct.
type MerklePath struct {
	PathLen uint64
	Path    []MerkleNode
}

// PdpVerificationReturns is an auto generated low-level Go binding around an user-defined struct.
type PdpVerificationReturns struct {
	FileIDs     [][]byte
	Tags        [][]byte
	UpdatedChal []Challenge
	Path        []MerklePath
	RootHashes  []byte
	FileInfo    FileInfo
	Success     bool
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

// PrepareForPdpVerificationParams is an auto generated low-level Go binding around an user-defined struct.
type PrepareForPdpVerificationParams struct {
	SectorInfo SectorInfo
	Challenges []Challenge
	ProveData  SectorProveData
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

// SectorProveData is an auto generated low-level Go binding around an user-defined struct.
type SectorProveData struct {
	ProveFileNum uint64
	BlockNum     uint64
	Proofs       []byte
	Tags         []byte
	MerklePath   []MerklePath
	PlotData     []byte
}

// VerifyPlotDataParams is an auto generated low-level Go binding around an user-defined struct.
type VerifyPlotDataParams struct {
	PlotInfo PlotInfo
	PlotData []byte
	Index    uint64
}

// VerifyProofWithMerklePathForFileParams is an auto generated low-level Go binding around an user-defined struct.
type VerifyProofWithMerklePathForFileParams struct {
	Version    uint64
	Proofs     []byte
	FileIds    [][]byte
	Tags       [][]byte
	Challenges []Challenge
	MerklePath []MerklePath
	RootHashes []byte
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Tags\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structVerifyProofWithMerklePathForFileParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x7b63f42e\",\"type\":\"bytes32\"}],\"name\":\"c_0x7b63f42e\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506120d6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063b6ddeca011610050578063b6ddeca0146100e1578063cca9fb6e14610111578063fd7c98081461012d57610072565b80632e19aeff146100775780638129fc1c146100a75780639f9ca644146100b1575b600080fd5b610091600480360381019061008c9190611307565b61015d565b60405161009e9190611b2f565b60405180910390f35b6100af6101ec565b005b6100cb60048036038101906100c691906112c6565b6102fc565b6040516100d89190611b6a565b60405180910390f35b6100fb60048036038101906100f69190611348565b61045a565b6040516101089190611b2f565b60405180910390f35b61012b6004803603810190610126919061125c565b6104e9565b005b61014760048036038101906101429190611285565b6104ec565b6040516101549190611b0d565b60405180910390f35b600061018b7fa597ff8b0a702f59c85cd3ed85d8395813ae9e9bb46cdcf6beb4f0e928e5678e60001b6104e9565b6101b77f55ebb1eae64711c74d3cb2d2035bd58e59bc1ec78c50d366bac24c88fd9f458e60001b6104e9565b6101e37f71196d46a8deed898dd1e340afa788e478ba85c78873bb4ee477dad93cd1f8f960001b6104e9565b60019050919050565b600060019054906101000a900460ff166102145760008054906101000a900460ff161561021d565b61021c6105d5565b5b61025c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025390611b4a565b60405180910390fd5b60008060019054906101000a900460ff1615905080156102ac576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6102d87fbeacf65b2e9fb70c5e3266b40b3eb6ce13c3d20a2b20923c19e6a1030d005bd960001b6104e9565b80156102f95760008060016101000a81548160ff0219169083151502179055505b50565b6103046105f9565b6103307f8f435c4c33f9896d6942b8ae68c6ca26ec0f4c8cd2d2a83a4ecc4f8cea364cec60001b6104e9565b61035c7fa0cfe98c15ec7ac68842d309bc20cd6b2c030c0fb42ff100ccd5b6f509ee0d6960001b6104e9565b6103887fbff3dd3df5c77d7f11f5ad21bc2691c474d0e54db9362f9bf664ee2b83c1e43e60001b6104e9565b6103906105f9565b6103bc7f8f0c382108085c0e04d6742fcccc306de91487f471dba5391d3cc2d5fba8fe3260001b6104e9565b6103e87f0ca89a4fa54fea040c720511cff7162bf49d3bbb6c04c7023ac7b1a8987feee360001b6104e9565b60018160c00190151590811515815250506104257f0d21943d8586fce18331fcef08bbad6175e9141d2915965320b191f41067aff860001b6104e9565b6104517f2ab16454be301a116bb5bee94c6685116b44139da4552225e42a11f023693bf760001b6104e9565b80915050919050565b60006104887f5d81fbf8acdcb7b0be288d85e9f8ac258a22672c85193b8c515b6412a3eca52660001b6104e9565b6104b47f172c645e79cc8413c7fd23127fab5a6f52ce48afd485f63ef880a40e7396e2e260001b6104e9565b6104e07fc2cf5e96723b26f7d4e6ae1f76c0d00cfec94352f511e947ad5443ebad6443e360001b6104e9565b60019050919050565b50565b606061051a7f6aa9bd5258eaddc20d1e2c2c3264a8acc8dada31b2bba2f05a8ce3b9f9f0b9a360001b6104e9565b6105467f71224835ce6d3d2143732a263c201d3b63b8607e74b329bbdd3ef3e6dbab5b2b60001b6104e9565b6105727fa1376d9420f338781b1b376d75b8c12fbda3dec06751399c652ac0702db5254360001b6104e9565b60606105a07f380227e718becf37426a4fcbf5bff46adb3e4e737306737f6810b18f00fb2c5c60001b6104e9565b6105cc7f585e1981fa6067c33e835818234011f05b98d8704a225596b89b62da6975daf860001b6104e9565b80915050919050565b60006105e0306105e6565b15905090565b600080823b905060008111915050919050565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200161062f61063e565b81526020016000151581525090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff1681526020016000815260200160001515815260200160006001811115610755577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001600060028111156107ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581526020016107d06107d6565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600061082861082384611bb1565b611b8c565b9050808382526020820190508285602086028201111561084757600080fd5b60005b8581101561089157813567ffffffffffffffff81111561086957600080fd5b8086016108768982610b38565b8552602085019450602084019350505060018101905061084a565b5050509392505050565b60006108ae6108a984611bdd565b611b8c565b905080838252602082019050828560408602820111156108cd57600080fd5b60005b858110156108fd57816108e38882610b77565b8452602084019350604083019250506001810190506108d0565b5050509392505050565b600061091a61091584611c09565b611b8c565b9050808382526020820190508285602086028201111561093957600080fd5b60005b8581101561098357813567ffffffffffffffff81111561095b57600080fd5b8086016109688982610c4f565b8552602085019450602084019350505060018101905061093c565b5050509392505050565b60006109a061099b84611c35565b611b8c565b905080838252602082019050828560208602820111156109bf57600080fd5b60005b85811015610a0957813567ffffffffffffffff8111156109e157600080fd5b8086016109ee8982610cc7565b855260208501945060208401935050506001810190506109c2565b5050509392505050565b6000610a26610a2184611c61565b611b8c565b905082815260208101848484011115610a3e57600080fd5b610a49848285611ead565b509392505050565b600081359050610a6081612006565b92915050565b600082601f830112610a7757600080fd5b8135610a87848260208601610815565b91505092915050565b600082601f830112610aa157600080fd5b8135610ab184826020860161089b565b91505092915050565b600082601f830112610acb57600080fd5b8135610adb848260208601610907565b91505092915050565b600082601f830112610af557600080fd5b8135610b0584826020860161098d565b91505092915050565b600081359050610b1d8161201d565b92915050565b600081359050610b3281612034565b92915050565b600082601f830112610b4957600080fd5b8135610b59848260208601610a13565b91505092915050565b600081359050610b718161204b565b92915050565b600060408284031215610b8957600080fd5b610b936040611b8c565b90506000610ba384828501611232565b6000830152506020610bb784828501611232565b60208301525092915050565b600060808284031215610bd557600080fd5b610bdf6080611b8c565b90506000610bef84828501610a51565b600083015250602082013567ffffffffffffffff811115610c0f57600080fd5b610c1b84828501610b38565b6020830152506040610c2f84828501611247565b6040830152506060610c4384828501611247565b60608301525092915050565b600060608284031215610c6157600080fd5b610c6b6060611b8c565b90506000610c7b84828501611247565b6000830152506020610c8f84828501611247565b602083015250604082013567ffffffffffffffff811115610caf57600080fd5b610cbb84828501610b38565b60408301525092915050565b600060408284031215610cd957600080fd5b610ce36040611b8c565b90506000610cf384828501611247565b600083015250602082013567ffffffffffffffff811115610d1357600080fd5b610d1f84828501610aba565b60208301525092915050565b600060608284031215610d3d57600080fd5b610d476060611b8c565b90506000610d5784828501611247565b6000830152506020610d6b84828501611247565b6020830152506040610d7f84828501611247565b60408301525092915050565b600060608284031215610d9d57600080fd5b610da76060611b8c565b9050600082013567ffffffffffffffff811115610dc357600080fd5b610dcf84828501610e33565b600083015250602082013567ffffffffffffffff811115610def57600080fd5b610dfb84828501610a90565b602083015250604082013567ffffffffffffffff811115610e1b57600080fd5b610e2784828501610f69565b60408301525092915050565b60006101808284031215610e4657600080fd5b610e51610180611b8c565b90506000610e6184828501610a51565b6000830152506020610e7584828501611247565b6020830152506040610e8984828501611247565b6040830152506060610e9d84828501611247565b6060830152506080610eb184828501610b62565b60808301525060a0610ec58482850161121d565b60a08301525060c0610ed98482850161121d565b60c08301525060e0610eed84828501611247565b60e083015250610100610f0284828501611247565b61010083015250610120610f1884828501611247565b61012083015250610140610f2e84828501610b0e565b6101408301525061016082013567ffffffffffffffff811115610f5057600080fd5b610f5c84828501610a66565b6101608301525092915050565b600060c08284031215610f7b57600080fd5b610f8560c0611b8c565b90506000610f9584828501611247565b6000830152506020610fa984828501611247565b602083015250604082013567ffffffffffffffff811115610fc957600080fd5b610fd584828501610b38565b604083015250606082013567ffffffffffffffff811115610ff557600080fd5b61100184828501610b38565b606083015250608082013567ffffffffffffffff81111561102157600080fd5b61102d84828501610ae4565b60808301525060a082013567ffffffffffffffff81111561104d57600080fd5b61105984828501610b38565b60a08301525092915050565b600060a0828403121561107757600080fd5b6110816060611b8c565b9050600061109184828501610d2b565b600083015250606082013567ffffffffffffffff8111156110b157600080fd5b6110bd84828501610b38565b60208301525060806110d184828501611247565b60408301525092915050565b600060e082840312156110ef57600080fd5b6110f960e0611b8c565b9050600061110984828501611247565b600083015250602082013567ffffffffffffffff81111561112957600080fd5b61113584828501610b38565b602083015250604082013567ffffffffffffffff81111561115557600080fd5b61116184828501610a66565b604083015250606082013567ffffffffffffffff81111561118157600080fd5b61118d84828501610a66565b606083015250608082013567ffffffffffffffff8111156111ad57600080fd5b6111b984828501610a90565b60808301525060a082013567ffffffffffffffff8111156111d957600080fd5b6111e584828501610ae4565b60a08301525060c082013567ffffffffffffffff81111561120557600080fd5b61121184828501610b38565b60c08301525092915050565b60008135905061122c8161205b565b92915050565b60008135905061124181612072565b92915050565b60008135905061125681612089565b92915050565b60006020828403121561126e57600080fd5b600061127c84828501610b23565b91505092915050565b60006020828403121561129757600080fd5b600082013567ffffffffffffffff8111156112b157600080fd5b6112bd84828501610bc3565b91505092915050565b6000602082840312156112d857600080fd5b600082013567ffffffffffffffff8111156112f257600080fd5b6112fe84828501610d8b565b91505092915050565b60006020828403121561131957600080fd5b600082013567ffffffffffffffff81111561133357600080fd5b61133f84828501611065565b91505092915050565b60006020828403121561135a57600080fd5b600082013567ffffffffffffffff81111561137457600080fd5b611380848285016110dd565b91505092915050565b600061139583836113f5565b60208301905092915050565b60006113ad838361169b565b905092915050565b60006113c18383611715565b60408301905092915050565b60006113d98383611952565b905092915050565b60006113ed83836119a2565b905092915050565b6113fe81611ded565b82525050565b600061140f82611ce2565b6114198185611d65565b935061142483611c92565b8060005b8381101561145557815161143c8882611389565b975061144783611d24565b925050600181019050611428565b5085935050505092915050565b600061146d82611ced565b6114778185611d76565b93508360208202850161148985611ca2565b8060005b858110156114c557848403895281516114a685826113a1565b94506114b183611d31565b925060208a0199505060018101905061148d565b50829750879550505050505092915050565b60006114e282611cf8565b6114ec8185611d87565b93506114f783611cb2565b8060005b8381101561152857815161150f88826113b5565b975061151a83611d3e565b9250506001810190506114fb565b5085935050505092915050565b600061154082611cf8565b61154a8185611d98565b935061155583611cb2565b8060005b8381101561158657815161156d88826113b5565b975061157883611d3e565b925050600181019050611559565b5085935050505092915050565b600061159e82611d03565b6115a88185611da9565b9350836020820285016115ba85611cc2565b8060005b858110156115f657848403895281516115d785826113cd565b94506115e283611d4b565b925060208a019950506001810190506115be565b50829750879550505050505092915050565b600061161382611d0e565b61161d8185611dba565b93508360208202850161162f85611cd2565b8060005b8581101561166b578484038952815161164c85826113e1565b945061165783611d58565b925060208a01995050600181019050611633565b50829750879550505050505092915050565b61168681611dff565b82525050565b61169581611dff565b82525050565b60006116a682611d19565b6116b08185611dcb565b93506116c0818560208601611ebc565b6116c981611f7e565b840191505092915050565b6116dd81611e89565b82525050565b6116ec81611e9b565b82525050565b60006116ff602e83611ddc565b915061170a82611f8f565b604082019050919050565b60408201600082015161172b6000850182611aef565b50602082015161173e6020850182611aef565b50505050565b6000610320830160008301518482036000860152611762828261169b565b915050602083015161177760208601826113f5565b506040830151848203604086015261178f828261169b565b91505060608301516117a46060860182611afe565b5060808301516117b76080860182611afe565b5060a08301516117ca60a0860182611afe565b5060c08301516117dd60c0860182611afe565b5060e08301516117f060e0860182611afe565b50610100830151611805610100860182611ae0565b5061012083015161181a610120860182611afe565b5061014083015161182f610140860182611afe565b50610160830151848203610160860152611849828261169b565b915050610180830151611860610180860182611afe565b506101a08301516118756101a0860182611ae0565b506101c083015161188a6101c086018261167d565b506101e083015161189f6101e08601826116e3565b506102008301516118b4610200860182611afe565b506102208301518482036102208601526118ce8282611404565b9150506102408301518482036102408601526118ea8282611404565b915050610260830151848203610260860152611906828261169b565b91505061028083015161191d6102808601826116d4565b506102a08301516119326102a086018261167d565b506102c08301516119476102c0860182611a9e565b508091505092915050565b600060608301600083015161196a6000860182611afe565b50602083015161197d6020860182611afe565b5060408301518482036040860152611995828261169b565b9150508091505092915050565b60006040830160008301516119ba6000860182611afe565b50602083015184820360208601526119d28282611593565b9150508091505092915050565b600060e08301600083015184820360008601526119fc8282611462565b91505060208301518482036020860152611a168282611462565b91505060408301518482036040860152611a3082826114d7565b91505060608301518482036060860152611a4a8282611608565b91505060808301518482036080860152611a64828261169b565b91505060a083015184820360a0860152611a7e8282611744565b91505060c0830151611a9360c086018261167d565b508091505092915050565b606082016000820151611ab46000850182611afe565b506020820151611ac76020850182611afe565b506040820151611ada6040850182611afe565b50505050565b611ae981611e5b565b82525050565b611af881611e65565b82525050565b611b0781611e75565b82525050565b60006020820190508181036000830152611b278184611535565b905092915050565b6000602082019050611b44600083018461168c565b92915050565b60006020820190508181036000830152611b63816116f2565b9050919050565b60006020820190508181036000830152611b8481846119df565b905092915050565b6000611b96611ba7565b9050611ba28282611eef565b919050565b6000604051905090565b600067ffffffffffffffff821115611bcc57611bcb611f4f565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611bf857611bf7611f4f565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c2457611c23611f4f565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c5057611c4f611f4f565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c7c57611c7b611f4f565b5b611c8582611f7e565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611df882611e3b565b9050919050565b60008115159050919050565b6000819050919050565b6000819050611e2382611fde565b919050565b6000819050611e3682611ff2565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000611e9482611e15565b9050919050565b6000611ea682611e28565b9050919050565b82818337600083830152505050565b60005b83811015611eda578082015181840152602081019050611ebf565b83811115611ee9576000848401525b50505050565b611ef882611f7e565b810181811067ffffffffffffffff82111715611f1757611f16611f4f565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60038110611fef57611fee611f20565b5b50565b6002811061200357612002611f20565b5b50565b61200f81611ded565b811461201a57600080fd5b50565b61202681611dff565b811461203157600080fd5b50565b61203d81611e0b565b811461204857600080fd5b50565b6003811061205857600080fd5b50565b61206481611e5b565b811461206f57600080fd5b50565b61207b81611e65565b811461208657600080fd5b50565b61209281611e75565b811461209d57600080fd5b5056fea26469706673582212202a004df68dccce74893792a04000236e6528fca7f12b4fdca8affd160fa12bf464736f6c63430008040033",
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

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) pure returns((uint32,uint32)[])
func (_Store *StoreCaller) GenChallenge(opts *bind.CallOpts, gParams GenChallengeParams) ([]Challenge, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GenChallenge", gParams)

	if err != nil {
		return *new([]Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new([]Challenge)).(*[]Challenge)

	return out0, err

}

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) pure returns((uint32,uint32)[])
func (_Store *StoreSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) pure returns((uint32,uint32)[])
func (_Store *StoreCallerSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreCaller) PrepareForPdpVerification(opts *bind.CallOpts, pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "PrepareForPdpVerification", pParams)

	if err != nil {
		return *new(PdpVerificationReturns), err
	}

	out0 := *abi.ConvertType(out[0], new(PdpVerificationReturns)).(*PdpVerificationReturns)

	return out0, err

}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreCallerSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) pure returns(bool)
func (_Store *StoreCaller) VerifyPlotData(opts *bind.CallOpts, vParams VerifyPlotDataParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyPlotData", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) pure returns(bool)
func (_Store *StoreSession) VerifyPlotData(vParams VerifyPlotDataParams) (bool, error) {
	return _Store.Contract.VerifyPlotData(&_Store.CallOpts, vParams)
}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) pure returns(bool)
func (_Store *StoreCallerSession) VerifyPlotData(vParams VerifyPlotDataParams) (bool, error) {
	return _Store.Contract.VerifyPlotData(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) pure returns(bool)
func (_Store *StoreCaller) VerifyProofWithMerklePathForFile(opts *bind.CallOpts, vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyProofWithMerklePathForFile", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) pure returns(bool)
func (_Store *StoreSession) VerifyProofWithMerklePathForFile(vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) pure returns(bool)
func (_Store *StoreCallerSession) VerifyProofWithMerklePathForFile(vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// C0x7b63f42e is a free data retrieval call binding the contract method 0xcca9fb6e.
//
// Solidity: function c_0x7b63f42e(bytes32 c__0x7b63f42e) pure returns()
func (_Store *StoreCaller) C0x7b63f42e(opts *bind.CallOpts, c__0x7b63f42e [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x7b63f42e", c__0x7b63f42e)

	if err != nil {
		return err
	}

	return err

}

// C0x7b63f42e is a free data retrieval call binding the contract method 0xcca9fb6e.
//
// Solidity: function c_0x7b63f42e(bytes32 c__0x7b63f42e) pure returns()
func (_Store *StoreSession) C0x7b63f42e(c__0x7b63f42e [32]byte) error {
	return _Store.Contract.C0x7b63f42e(&_Store.CallOpts, c__0x7b63f42e)
}

// C0x7b63f42e is a free data retrieval call binding the contract method 0xcca9fb6e.
//
// Solidity: function c_0x7b63f42e(bytes32 c__0x7b63f42e) pure returns()
func (_Store *StoreCallerSession) C0x7b63f42e(c__0x7b63f42e [32]byte) error {
	return _Store.Contract.C0x7b63f42e(&_Store.CallOpts, c__0x7b63f42e)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactorSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}
