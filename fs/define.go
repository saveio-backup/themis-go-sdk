package fs

import (
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"github.com/saveio/themis/smartcontract/service/native/savefs/pdp"
	"github.com/saveio/themis/smartcontract/service/native/usdt"
	"math/big"
	"time"
)

type ContractClient interface {
	SetDefaultAccount(acc *account.Account)
	GetDefaultAccount() *account.Account
	GetClient() *client.ClientMgr

	InvokeNativeContract(signer *account.Account, method string, params []interface{}) (common.Uint256, error)
	InvokeNativeContractWithGasLimitUserDefine(signer *account.Account, gasLimit uint64, method string,
		params []interface{}) (common.Uint256, error)
	PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error)
	PreExecInvokeNativeContractWithSigner(signer *account.Account, method string,
		params []interface{}) (*sdkcom.PreExecResult, error)
	PollForTxConfirmed(timeout time.Duration, txHash []byte) (bool, error)

	GetSetting() (*fs.FsSetting, error)
	savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
		maxProveBlockNum, defProveLevel uint64, minVolume uint64) ([]byte, error)
	NewVerifierService()

	GetNodeList() (*fs.FsNodesInfo, error)
	GetNodeListByAddrs(addrs []common.Address) (*fs.FsNodesInfo, error)
	NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error)
	NodeRegister(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error)
	NodeUpdate(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error)
	NodeCancel() ([]byte, error)
	NodeWithDrawProfit() ([]byte, error)

	GetUploadStorageFee(opt *fs.UploadOption) (*fs.StorageFee, error)
	GetDeleteFilesStorageFee(fileHashStrs []string) (uint64, error)
	StoreFile(fileHashStr, blocksRoot string, blockNum uint64,
		blockSize uint64, proveLevel uint64, expiredHeight uint64, copyNum uint64,
		fileDesc []byte, privilege uint64, proveParam []byte, storageType uint64, realFileSize uint64,
		primaryNodes, candidateNodes []common.Address, plotInfo *fs.PlotInfo, url string) ([]byte, error)
	FileRenew(fileHashStr string, renewTimes uint64) ([]byte, error)
	GetFileInfo(fileHashStr string) (*fs.FileInfo, error)
	GetFileInfos(fileHashStrs []string) (*fs.FileInfoList, error)
	ChangeFileOwner(fileHashStr string, newOwner common.Address) ([]byte, error)
	ChangeFilePrivilege(fileHashStr string, newPrivilege uint64) ([]byte, error)
	GetFileList(addr common.Address) (*fs.FileList, error)
	GetUnprovePrimaryFileList(addr common.Address) (*fs.FileList, error)
	GetUnProveCandidateFileList(addr common.Address) (*fs.FileList, error)
	DeleteFile(fileHashStr string) ([]byte, error)
	DeleteFiles(fileHashStrs []string, gasLimit uint64) ([]byte, error)
	DeleteUnSettledFiles() ([]byte, error)
	GetUnSettledFiles(addr common.Address) (*fs.FileList, error)

	AddWhiteLists(fileHashStr string, whitelists []fs.Rule) ([]byte, error)
	WhiteListOp(fileHashStr string, op uint64, whiteList fs.WhiteList) ([]byte, error)
	GetWhiteList(fileHashStr string) (*fs.WhiteList, error)

	ProveParamSer(rootHash []byte, fileId pdp.FileID) ([]byte, error)
	ProveParamDes(proveParam []byte) (*fs.ProveParam, error)
	FileProve(fileHashStr string, proveData []byte, blockHeight uint64, sectorId uint64) ([]byte, error)
	SectorProve(sectorId uint64, challengeHeight uint64, proveData []byte) ([]byte, error)
	GenChallenge(walletAddr common.Address, hash common.Uint256, fileBlockNum, proveNum uint64) []pdp.Challenge
	CheckNodeSectorProveInTime(addr common.Address, sectorId uint64) ([]byte, error)
	GetFileProveDetails(fileHashStr string) (*fs.FsProveDetails, error)

	UpdateUserSpace(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) ([]byte, error)
	CashUserSpace(walletAddr common.Address) ([]byte, error)
	GetUserSpace(walletAddr common.Address) (*fs.UserSpace, error)
	GetUpdateSpaceCost(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) (*usdt.State, error)
	DeleteUserSpace() ([]byte, error)

	CreateSector(sectorId uint64, proveLevel uint64, size uint64, isPlots bool) ([]byte, error)
	DeleteSector(sectorId uint64) ([]byte, error)
	GetSectorInfo(sectorId uint64) (*fs.SectorInfo, error)
	DeleteFileInSector(sectorId uint64, fileHashStr string) ([]byte, error)
	GetSectorInfosForNode(addr common.Address) (*fs.SectorInfos, error)

	GetEventsByBlockHeight(blockHeight *big.Int) ([]map[string]interface{}, error)
}
