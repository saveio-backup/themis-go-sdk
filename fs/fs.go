package fs

import (
	"github.com/saveio/themis-go-sdk/client"
	"math/big"
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
	Client ContractClient
}

var _ ContractClient = (*Fs)(nil)

func (f *Fs) SetDefaultAccount(acc *account.Account) {
	f.Client.SetDefaultAccount(acc)
}

func (f *Fs) GetDefaultAccount() *account.Account {
	return f.Client.GetDefaultAccount()
}

func (f *Fs) GetClient() *client.ClientMgr {
	return f.Client.GetClient()
}

func (f *Fs) NewClient(contractClient ContractClient) {
	f.Client = contractClient
}

func (f *Fs) InvokeNativeContract(signer *account.Account, method string,
	params []interface{}) (common.Uint256, error) {
	return f.Client.InvokeNativeContract(signer, method, params)
}

func (f *Fs) InvokeNativeContractWithGasLimitUserDefine(signer *account.Account, gasLimit uint64, method string,
	params []interface{}) (common.Uint256, error) {
	return f.Client.InvokeNativeContractWithGasLimitUserDefine(signer, gasLimit, method, params)
}

func (f *Fs) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	return f.Client.PreExecInvokeNativeContract(method, params)
}

func (f *Fs) PreExecInvokeNativeContractWithSigner(signer *account.Account, method string,
	params []interface{}) (*sdkcom.PreExecResult, error) {
	return f.Client.PreExecInvokeNativeContractWithSigner(signer, method, params)
}

func (f *Fs) GetSetting() (*fs.FsSetting, error) {
	return f.Client.GetSetting()
}

func (f *Fs) GetNodeList() (*fs.FsNodesInfo, error) {
	return f.Client.GetNodeList()
}

func (f *Fs) GetNodeListByAddrs(addrs []common.Address) (*fs.FsNodesInfo, error) {
	return f.Client.GetNodeListByAddrs(addrs)
}

func (f *Fs) ProveParamSer(rootHash []byte, fileId pdp.FileID) ([]byte, error) {
	return f.Client.ProveParamSer(rootHash, fileId)
}

func (f *Fs) ProveParamDes(proveParam []byte) (*fs.ProveParam, error) {
	return f.Client.ProveParamDes(proveParam)
}

func (f *Fs) StoreFile(fileHashStr, blocksRoot string, blockNum uint64,
	blockSize uint64, proveLevel uint64, expiredHeight uint64, copyNum uint64,
	fileDesc []byte, privilege uint64, proveParam []byte, storageType uint64, realFileSize uint64,
	primaryNodes, candidateNodes []common.Address, plotInfo *fs.PlotInfo, url string) ([]byte, error) {
	return f.Client.StoreFile(fileHashStr, blocksRoot, blockNum, blockSize,
		proveLevel, expiredHeight, copyNum, fileDesc, privilege, proveParam, storageType,
		realFileSize, primaryNodes, candidateNodes, plotInfo, url)
}

func (f *Fs) FileRenew(fileHashStr string, renewTimes uint64) ([]byte, error) {
	return f.Client.FileRenew(fileHashStr, renewTimes)
}

func (f *Fs) GetFileInfo(fileHashStr string) (*fs.FileInfo, error) {
	return f.Client.GetFileInfo(fileHashStr)
}

func (f *Fs) GetFileInfos(fileHashStrs []string) (*fs.FileInfoList, error) {
	return f.Client.GetFileInfos(fileHashStrs)
}

func (f *Fs) ChangeFileOwner(fileHashStr string, newOwner common.Address) ([]byte, error) {
	return f.Client.ChangeFileOwner(fileHashStr, newOwner)
}

func (f *Fs) ChangeFilePrivilege(fileHashStr string, newPrivilege uint64) ([]byte, error) {
	return f.Client.ChangeFilePrivilege(fileHashStr, newPrivilege)
}

func (f *Fs) GetFileList(addr common.Address) (*fs.FileList, error) {
	return f.Client.GetFileList(addr)
}

func (f *Fs) GetUnprovePrimaryFileList(addr common.Address) (*fs.FileList, error) {
	return f.Client.GetUnprovePrimaryFileList(addr)
}

func (f *Fs) GetUnProveCandidateFileList(addr common.Address) (*fs.FileList, error) {
	return f.Client.GetUnProveCandidateFileList(addr)
}

func (f *Fs) GetFileProveDetails(fileHashStr string) (*fs.FsProveDetails, error) {
	return f.Client.GetFileProveDetails(fileHashStr)
}

func (f *Fs) AddWhiteLists(fileHashStr string, whitelists []fs.Rule) ([]byte, error) {
	return f.Client.AddWhiteLists(fileHashStr, whitelists)
}

func (f *Fs) WhiteListOp(fileHashStr string, op uint64, whiteList fs.WhiteList) ([]byte, error) {
	return f.Client.WhiteListOp(fileHashStr, op, whiteList)
}

func (f *Fs) GetWhiteList(fileHashStr string) (*fs.WhiteList, error) {
	return f.Client.GetWhiteList(fileHashStr)
}

func (f *Fs) DeleteFile(fileHashStr string) ([]byte, error) {
	return f.Client.DeleteFile(fileHashStr)
}

func (f *Fs) DeleteFiles(fileHashStrs []string, gasLimit uint64) ([]byte, error) {
	return f.Client.DeleteFiles(fileHashStrs, gasLimit)
}

func (f *Fs) PollForTxConfirmed(timeout time.Duration, txHash []byte) (bool, error) {
	return f.Client.PollForTxConfirmed(timeout, txHash)
}

func (f *Fs) savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
	maxProveBlockNum, defProveLevel uint64, minVolume uint64) ([]byte, error) {
	return f.Client.savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
		maxProveBlockNum, defProveLevel, minVolume)
}

func (f *Fs) NewVerifierService() {
	f.Client.NewVerifierService()
}

func (f *Fs) NodeRegister(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	return f.Client.NodeRegister(volume, serviceTime, nodeAddr)
}

func (f *Fs) NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error) {
	return f.Client.NodeQuery(nodeWallet)
}

func (f *Fs) NodeUpdate(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	return f.Client.NodeUpdate(volume, serviceTime, nodeAddr)
}

func (f *Fs) NodeCancel() ([]byte, error) {
	return f.Client.NodeCancel()
}

func (f *Fs) NodeWithDrawProfit() ([]byte, error) {
	return f.Client.NodeWithDrawProfit()
}

func (f *Fs) FileProve(fileHashStr string, proveData []byte, blockHeight uint64, sectorId uint64) ([]byte, error) {
	return f.Client.FileProve(fileHashStr, proveData, blockHeight, sectorId)
}

func (f *Fs) GenChallenge(walletAddr common.Address, hash common.Uint256, fileBlockNum, proveNum uint64) []pdp.Challenge {
	return fs.GenChallenge(walletAddr, hash, uint32(fileBlockNum), uint32(proveNum))
}

// UpdateUserSpace. user space operation for space owner.
func (f *Fs) UpdateUserSpace(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) ([]byte, error) {
	return f.Client.UpdateUserSpace(walletAddr, size, blockCount)
}

// UpdateUserSpace. user space operation for space owner.
func (f *Fs) CashUserSpace(walletAddr common.Address) ([]byte, error) {
	return f.Client.CashUserSpace(walletAddr)
}

// GetUserSpace. get user space with wallet address
func (f *Fs) GetUserSpace(walletAddr common.Address) (*fs.UserSpace, error) {
	return f.Client.GetUserSpace(walletAddr)
}

func (f *Fs) GetUpdateSpaceCost(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) (*usdt.State, error) {
	return f.Client.GetUpdateSpaceCost(walletAddr, size, blockCount)
}

func (f *Fs) DeleteUserSpace() ([]byte, error) {
	return f.Client.DeleteUserSpace()
}

func (f *Fs) GetUploadStorageFee(opt *fs.UploadOption) (*fs.StorageFee, error) {
	return f.Client.GetUploadStorageFee(opt)
}

func (f *Fs) GetDeleteFilesStorageFee(fileHashStrs []string) (uint64, error) {
	return f.Client.GetDeleteFilesStorageFee(fileHashStrs)
}

func (f *Fs) CreateSector(sectorId uint64, proveLevel uint64, size uint64, isPlots bool) ([]byte, error) {
	return f.Client.CreateSector(sectorId, proveLevel, size, isPlots)
}

func (f *Fs) DeleteSector(sectorId uint64) ([]byte, error) {
	return f.Client.DeleteSector(sectorId)
}

func (f *Fs) GetSectorInfo(sectorId uint64) (*fs.SectorInfo, error) {
	return f.Client.GetSectorInfo(sectorId)
}

func (f *Fs) DeleteFileInSector(sectorId uint64, fileHashStr string) ([]byte, error) {
	return f.Client.DeleteFileInSector(sectorId, fileHashStr)
}

func (f *Fs) GetSectorInfosForNode(addr common.Address) (*fs.SectorInfos, error) {
	return f.Client.GetSectorInfosForNode(addr)
}

func (f *Fs) SectorProve(sectorId uint64, challengeHeight uint64, proveData []byte) ([]byte, error) {
	return f.Client.SectorProve(sectorId, challengeHeight, proveData)
}

func (f *Fs) GetUnSettledFiles(addr common.Address) (*fs.FileList, error) {
	return f.Client.GetUnSettledFiles(addr)
}

func (f *Fs) DeleteUnSettledFiles() ([]byte, error) {
	return f.Client.DeleteUnSettledFiles()
}

func (f *Fs) CheckNodeSectorProveInTime(addr common.Address, sectorId uint64) ([]byte, error) {
	return f.Client.CheckNodeSectorProveInTime(addr, sectorId)
}

func (t *Fs) GetEventsByBlockHeight(blockHeight *big.Int) ([]map[string]interface{}, error) {
	return t.Client.GetEventsByBlockHeight(blockHeight)
}
