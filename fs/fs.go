package fs

import (
	"time"

	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"github.com/saveio/themis/smartcontract/service/native/savefs/pdp"
	"github.com/saveio/themis/smartcontract/service/native/usdt"
	nutils "github.com/saveio/themis/smartcontract/service/native/utils"
)

var (
	FS_CONTRACT_ADDRESS = nutils.OntFSContractAddress
	FS_CONTRACT_VERSION = byte(0)
)

type Fs struct {
	ChainClient ChainClient
}

func (f *Fs) NewChainClient(chainClient ChainClient) {
	f.ChainClient = chainClient
}

func (f *Fs) InvokeNativeContract(signer *account.Account, method string,
	params []interface{}) (common.Uint256, error) {
	return f.ChainClient.InvokeNativeContract(signer, method, params)
}

func (f *Fs) InvokeNativeContractWithGasLimitUserDefine(signer *account.Account, gasLimit uint64, method string,
	params []interface{}) (common.Uint256, error) {
	return f.ChainClient.InvokeNativeContractWithGasLimitUserDefine(signer, gasLimit, method, params)
}

func (f *Fs) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	return f.ChainClient.PreExecInvokeNativeContract(method, params)
}

func (f *Fs) PreExecInvokeNativeContractWithSigner(signer *account.Account, method string,
	params []interface{}) (*sdkcom.PreExecResult, error) {
	return f.ChainClient.PreExecInvokeNativeContractWithSigner(signer, method, params)
}

func (f *Fs) GetSetting() (*fs.FsSetting, error) {
	return f.ChainClient.GetSetting()
}

func (f *Fs) GetNodeList() (*fs.FsNodesInfo, error) {
	return f.ChainClient.GetNodeList()
}

func (f *Fs) GetNodeListByAddrs(addrs []common.Address) (*fs.FsNodesInfo, error) {
	return f.ChainClient.GetNodeListByAddrs(addrs)
}

func (f *Fs) ProveParamSer(rootHash []byte, fileId pdp.FileID) ([]byte, error) {
	return f.ChainClient.ProveParamSer(rootHash, fileId)
}

func (f *Fs) ProveParamDes(proveParam []byte) (*fs.ProveParam, error) {
	return f.ChainClient.ProveParamDes(proveParam)
}

func (f *Fs) StoreFile(fileHashStr, blocksRoot string, blockNum uint64,
	blockSize uint64, proveLevel uint64, expiredHeight uint64, copyNum uint64,
	fileDesc []byte, privilege uint64, proveParam []byte, storageType uint64, realFileSize uint64,
	primaryNodes, candidateNodes []common.Address, plotInfo *fs.PlotInfo) ([]byte, error) {
	return f.ChainClient.StoreFile(fileHashStr, blocksRoot, blockNum, blockSize,
		proveLevel, expiredHeight, copyNum, fileDesc, privilege, proveParam, storageType,
		realFileSize, primaryNodes, candidateNodes, plotInfo)
}

func (f *Fs) FileRenew(fileHashStr string, renewTimes uint64) ([]byte, error) {
	return f.ChainClient.FileRenew(fileHashStr, renewTimes)
}

func (f *Fs) GetFileInfo(fileHashStr string) (*fs.FileInfo, error) {
	return f.ChainClient.GetFileInfo(fileHashStr)
}

func (f *Fs) GetFileInfos(fileHashStrs []string) (*fs.FileInfoList, error) {
	return f.ChainClient.GetFileInfos(fileHashStrs)
}

func (f *Fs) ChangeFileOwner(fileHashStr string, newOwner common.Address) ([]byte, error) {
	return f.ChainClient.ChangeFileOwner(fileHashStr, newOwner)
}

func (f *Fs) ChangeFilePrivilege(fileHashStr string, newPrivilege uint64) ([]byte, error) {
	return f.ChainClient.ChangeFilePrivilege(fileHashStr, newPrivilege)
}

func (f *Fs) GetFileList(addr common.Address) (*fs.FileList, error) {
	return f.ChainClient.GetFileList(addr)
}

func (f *Fs) GetUnprovePrimaryFileList(addr common.Address) (*fs.FileList, error) {
	return f.ChainClient.GetUnprovePrimaryFileList(addr)
}

func (f *Fs) GetUnProveCandidateFileList(addr common.Address) (*fs.FileList, error) {
	return f.ChainClient.GetUnProveCandidateFileList(addr)
}

func (f *Fs) GetFileProveDetails(fileHashStr string) (*fs.FsProveDetails, error) {
	return f.ChainClient.GetFileProveDetails(fileHashStr)
}

func (f *Fs) AddWhiteLists(fileHashStr string, whitelists []fs.Rule) ([]byte, error) {
	return f.ChainClient.AddWhiteLists(fileHashStr, whitelists)
}

func (f *Fs) WhiteListOp(fileHashStr string, op uint64, whiteList fs.WhiteList) ([]byte, error) {
	return f.ChainClient.WhiteListOp(fileHashStr, op, whiteList)
}

func (f *Fs) GetWhiteList(fileHashStr string) (*fs.WhiteList, error) {
	return f.ChainClient.GetWhiteList(fileHashStr)
}

func (f *Fs) DeleteFile(fileHashStr string) ([]byte, error) {
	return f.ChainClient.DeleteFile(fileHashStr)
}

func (f *Fs) DeleteFiles(fileHashStrs []string, gasLimit uint64) ([]byte, error) {
	return f.ChainClient.DeleteFiles(fileHashStrs, gasLimit)
}

func (f *Fs) PollForTxConfirmed(timeout time.Duration, txHash []byte) (bool, error) {
	return f.ChainClient.PollForTxConfirmed(timeout, txHash)
}

func (f *Fs) savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
	maxProveBlockNum, defProveLevel uint64, minVolume uint64) ([]byte, error) {
	return f.ChainClient.savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
		maxProveBlockNum, defProveLevel, minVolume)
}

func (f *Fs) NodeRegister(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	return f.ChainClient.NodeRegister(volume, serviceTime, nodeAddr)
}

func (f *Fs) NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error) {
	return f.ChainClient.NodeQuery(nodeWallet)
}

func (f *Fs) NodeUpdate(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	return f.ChainClient.NodeUpdate(volume, serviceTime, nodeAddr)
}

func (f *Fs) NodeCancel() ([]byte, error) {
	return f.ChainClient.NodeCancel()
}

func (f *Fs) NodeWithDrawProfit() ([]byte, error) {
	return f.ChainClient.NodeWithDrawProfit()
}

func (f *Fs) FileProve(fileHashStr string, proveData []byte, blockHeight uint64, sectorId uint64) ([]byte, error) {
	return f.ChainClient.FileProve(fileHashStr, proveData, blockHeight, sectorId)
}

func (f *Fs) GenChallenge(walletAddr common.Address, hash common.Uint256, fileBlockNum, proveNum uint64) []pdp.Challenge {
	return fs.GenChallenge(walletAddr, hash, uint32(fileBlockNum), uint32(proveNum))
}

// UpdateUserSpace. user space operation for space owner.
func (f *Fs) UpdateUserSpace(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) ([]byte, error) {
	return f.ChainClient.UpdateUserSpace(walletAddr, size, blockCount)
}

// GetUserSpace. get user space with wallet address
func (f *Fs) GetUserSpace(walletAddr common.Address) (*fs.UserSpace, error) {
	return f.ChainClient.GetUserSpace(walletAddr)
}

func (f *Fs) GetUpdateSpaceCost(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) (*usdt.State, error) {
	return f.ChainClient.GetUpdateSpaceCost(walletAddr, size, blockCount)
}

func (f *Fs) DeleteUserSpace() ([]byte, error) {
	return f.ChainClient.DeleteUserSpace()
}

func (f *Fs) GetUploadStorageFee(opt *fs.UploadOption) (*fs.StorageFee, error) {
	return f.ChainClient.GetUploadStorageFee(opt)
}

func (f *Fs) GetDeleteFilesStorageFee(fileHashStrs []string) (uint64, error) {
	return f.ChainClient.GetDeleteFilesStorageFee(fileHashStrs)
}

func (f *Fs) CreateSector(sectorId uint64, proveLevel uint64, size uint64, isPlots bool) ([]byte, error) {
	return f.ChainClient.CreateSector(sectorId, proveLevel, size, isPlots)
}

func (f *Fs) DeleteSector(sectorId uint64) ([]byte, error) {
	return f.ChainClient.DeleteSector(sectorId)
}

func (f *Fs) GetSectorInfo(sectorId uint64) (*fs.SectorInfo, error) {
	return f.ChainClient.GetSectorInfo(sectorId)
}

func (f *Fs) DeleteFileInSector(sectorId uint64, fileHashStr string) ([]byte, error) {
	return f.ChainClient.DeleteFileInSector(sectorId, fileHashStr)
}

func (f *Fs) GetSectorInfosForNode(addr common.Address) (*fs.SectorInfos, error) {
	return f.ChainClient.GetSectorInfosForNode(addr)
}

func (f *Fs) SectorProve(sectorId uint64, challengeHeight uint64, proveData []byte) ([]byte, error) {
	return f.ChainClient.SectorProve(sectorId, challengeHeight, proveData)
}

func (f *Fs) GetUnSettledFiles(addr common.Address) (*fs.FileList, error) {
	return f.ChainClient.GetUnSettledFiles(addr)
}

func (f *Fs) DeleteUnSettledFiles() ([]byte, error) {
	return f.ChainClient.DeleteUnSettledFiles()
}

func (f *Fs) CheckNodeSectorProveInTime(addr common.Address, sectorId uint64) ([]byte, error) {
	return f.ChainClient.CheckNodeSectorProveInTime(addr, sectorId)
}
