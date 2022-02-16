package fs

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"github.com/saveio/themis/smartcontract/service/native/savefs/pdp"
	"github.com/saveio/themis/smartcontract/service/native/usdt"
	"time"
)

type Ethereum struct {
	Client            *client.ClientMgr
	DefAcc            *account.Account
	PollForTxDuration time.Duration
}

const (
	Config string = "Config"
)

func (t *Ethereum) SetDefaultAccount(acc *account.Account) {
	t.DefAcc = acc
}

func (t *Ethereum) GetDefaultAccount() *account.Account {
	return t.DefAcc
}

// GetClient only oni used
func (t *Ethereum) GetClient() *client.ClientMgr {
	return t.Client;
}

func (t *Ethereum) InvokeNativeContract(signer *account.Account, method string, params []interface{}) (common.Uint256, error) {
	if signer == nil {
		return common.UINT256_EMPTY, errors.New("signer is nil")
	}
	tx, err := utils.NewNativeInvokeTransaction(sdkcom.GAS_PRICE, sdkcom.GAS_LIMIT, FS_CONTRACT_VERSION, FS_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return t.Client.SendTransaction(tx)
}

func (t *Ethereum) InvokeNativeContractWithGasLimitUserDefine(signer *account.Account, gasLimit uint64, method string, params []interface{}) (common.Uint256, error) {
	if signer == nil {
		return common.UINT256_EMPTY, errors.New("signer is nil")
	}
	tx, err := utils.NewNativeInvokeTransaction(sdkcom.GAS_PRICE, gasLimit, FS_CONTRACT_VERSION, FS_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return t.Client.SendTransaction(tx)
}

func (t *Ethereum) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, FS_CONTRACT_VERSION, FS_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return nil, err
	}
	return t.Client.PreExecTransaction(tx)
}

func (t *Ethereum) PreExecInvokeNativeContractWithSigner(signer *account.Account, method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, FS_CONTRACT_VERSION, FS_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return nil, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return nil, err
	}
	return t.Client.PreExecTransaction(tx)
}

func (t *Ethereum) GetSetting() (*fs.FsSetting, error) {
	ethClient := t.Client.GetEthClient()
	store := ethClient[Config].(*configStore.Store)
	setting, err := store.GetSetting(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	fsSetting := &fs.FsSetting{
		FsGasPrice:         setting.GasPrice,
		GasPerGBPerBlock:   setting.GasPerGBPerBlock,
		GasPerKBForRead:    setting.GasPerKBPerBlock,
		GasForChallenge:    setting.GasForChallenge,
		MaxProveBlockNum:   setting.MaxProveBlockNum,
		MinVolume:          setting.MinVolume,
		DefaultProvePeriod: setting.DefaultProvePeriod,
		DefaultProveLevel:  setting.DefaultProveLevel,
		DefaultCopyNum:     setting.DefaultCopyNum,
	}
	return fsSetting, nil
}

func (t *Ethereum) GetNodeList() (*fs.FsNodesInfo, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_NODE_LIST, []interface{}{},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetNodeList result toByteArray: %s", err.Error())
	}
	var nodesInfo fs.FsNodesInfo
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		reader := bytes.NewReader(retInfo.Info)
		if err = nodesInfo.Deserialize(reader); err != nil {
			return nil, fmt.Errorf("GetNodeList json Unmarshal: %s", err.Error())
		}
		return &nodesInfo, nil
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetNodeListByAddrs(addrs []common.Address) (*fs.FsNodesInfo, error) {
	nodeList := &fs.NodeList{
		AddrNum:  uint64(len(addrs)),
		AddrList: addrs,
	}
	buf := new(bytes.Buffer)
	if err := nodeList.Serialize(buf); err != nil {
		return nil, fmt.Errorf("nodeList serialize error: %s", err.Error())
	}
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_NODE_LIST_BY_ADDRS, []interface{}{buf.Bytes()},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetNodeList result toByteArray: %s", err.Error())
	}
	var nodesInfo fs.FsNodesInfo
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		reader := bytes.NewReader(retInfo.Info)
		if err = nodesInfo.Deserialize(reader); err != nil {
			return nil, fmt.Errorf("GetNodeList json Unmarshal: %s", err.Error())
		}
		return &nodesInfo, nil
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) ProveParamSer(rootHash []byte, fileId pdp.FileID) ([]byte, error) {
	var proveParam fs.ProveParam
	proveParam.RootHash = rootHash
	proveParam.FileID = fileId
	bf := new(bytes.Buffer)
	if err := proveParam.Serialize(bf); err != nil {
		return nil, fmt.Errorf("ProveParam Serialize error: %s", err.Error())
	}
	return bf.Bytes(), nil
}

func (t *Ethereum) ProveParamDes(proveParam []byte) (*fs.ProveParam, error) {
	var proveParamSt fs.ProveParam
	reader := bytes.NewReader(proveParam)
	if err := proveParamSt.Deserialize(reader); err != nil {
		return nil, fmt.Errorf("ProveParamDes Deserialize error: %s", err.Error())
	}
	return &proveParamSt, nil
}

func (t *Ethereum) StoreFile(fileHashStr, blocksRoot string, blockNum uint64,
	blockSize uint64, proveLevel uint64, expiredHeight uint64, copyNum uint64,
	fileDesc []byte, privilege uint64, proveParam []byte, storageType uint64, realFileSize uint64,
	primaryNodes, candidateNodes []common.Address, plotInfo *fs.PlotInfo) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}

	primary := fs.NodeList{
		AddrNum:  uint64(len(primaryNodes)),
		AddrList: primaryNodes,
	}
	candidates := fs.NodeList{
		AddrNum:  uint64(len(candidateNodes)),
		AddrList: candidateNodes,
	}

	fileHash := []byte(fileHashStr)
	fileInfo := &fs.FileInfo{
		FileHash:       fileHash,
		BlocksRoot:     []byte(blocksRoot),
		FileOwner:      t.DefAcc.Address,
		FileDesc:       fileDesc,
		Privilege:      privilege,
		FileBlockNum:   blockNum,
		FileBlockSize:  blockSize,
		ProveLevel:     proveLevel,
		ProveTimes:     0,
		ExpiredHeight:  expiredHeight,
		CopyNum:        copyNum,
		Deposit:        0,
		FileProveParam: proveParam,
		ProveBlockNum:  0,
		BlockHeight:    0,
		ValidFlag:      true,
		StorageType:    storageType,
		RealFileSize:   realFileSize,
		PrimaryNodes:   primary,
		CandidateNodes: candidates,
		IsPlotFile:     plotInfo != nil,
		PlotInfo:       plotInfo,
	}
	buf := new(bytes.Buffer)
	if err := fileInfo.Serialize(buf); err != nil {
		return nil, fmt.Errorf("fileInfo serialize error: %s", err.Error())
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_STORE_FILE, []interface{}{buf.Bytes()},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) FileRenew(fileHashStr string, renewTimes uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	fileHash := []byte(fileHashStr)
	fileRenew := &fs.FileReNew{
		FileHash:   fileHash,
		FromAddr:   t.DefAcc.Address,
		ReNewTimes: renewTimes,
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_FILE_RENEW, []interface{}{fileRenew},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetFileInfo(fileHashStr string) (*fs.FileInfo, error) {
	fileHash := []byte(fileHashStr)
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_FILE_INFO, []interface{}{fileHash},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetFileInfo result toByteArray: %s", err.Error())
	}

	var fileInfo fs.FileInfo
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsFileInfoReader := bytes.NewReader(retInfo.Info)
		err = fileInfo.Deserialize(fsFileInfoReader)
		if err != nil {
			return nil, fmt.Errorf("GetFileInfo error: %s", err.Error())
		}
		return &fileInfo, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetFileInfos(fileHashStrs []string) (*fs.FileInfoList, error) {
	fileHashes := make([]fs.FileHash, 0, len(fileHashStrs))
	for _, fileHashStr := range fileHashStrs {
		fileHashes = append(fileHashes, fs.FileHash{
			Hash: []byte(fileHashStr),
		})
	}
	fileList := fs.FileList{
		FileNum: uint64(len(fileHashStrs)),
		List:    fileHashes,
	}
	buf := new(bytes.Buffer)
	if err := fileList.Serialize(buf); err != nil {
		return nil, fmt.Errorf("fileList serialize error: %s", err.Error())
	}
	ret, err := t.PreExecInvokeNativeContract(fs.FS_GET_FILE_INFOS, []interface{}{buf.Bytes()})
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetFileInfo result toByteArray: %s", err.Error())
	}
	fileInfos := &fs.FileInfoList{}
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsFileInfoReader := bytes.NewReader(retInfo.Info)
		err = fileInfos.Deserialize(fsFileInfoReader)
		if err != nil {
			return nil, fmt.Errorf("GetFileInfo error: %s", err.Error())
		}
		return fileInfos, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) ChangeFileOwner(fileHashStr string, newOwner common.Address) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	fileHash := []byte(fileHashStr)
	ownerChange := &fs.OwnerChange{
		FileHash: fileHash,
		CurOwner: t.DefAcc.Address,
		NewOwner: newOwner,
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_CHANGE_FILE_OWNER, []interface{}{ownerChange},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) ChangeFilePrivilege(fileHashStr string, newPrivilege uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	fileHash := []byte(fileHashStr)
	priChange := &fs.PriChange{
		FileHash:  fileHash,
		Privilege: newPrivilege,
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_CHANGE_FILE_PRIVILEGE, []interface{}{priChange},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetFileList(addr common.Address) (*fs.FileList, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_FILE_LIST, []interface{}{addr},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetFileList result toByteArray: %s", err.Error())
	}

	var fileList fs.FileList
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsFileListReader := bytes.NewReader(retInfo.Info)
		err = fileList.Deserialize(fsFileListReader)
		if err != nil {
			return nil, fmt.Errorf("GetFileList error: %s", err.Error())
		}
		return &fileList, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetUnprovePrimaryFileList(addr common.Address) (*fs.FileList, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_UNPROVE_PRIMARY_FILES, []interface{}{addr},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetFileList result toByteArray: %s", err.Error())
	}

	var fileList fs.FileList
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsFileListReader := bytes.NewReader(retInfo.Info)
		err = fileList.Deserialize(fsFileListReader)
		if err != nil {
			return nil, fmt.Errorf("GetFileList error: %s", err.Error())
		}
		return &fileList, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetUnProveCandidateFileList(addr common.Address) (*fs.FileList, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_UNPROVE_PRIMARY_FILES, []interface{}{addr},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetFileList result toByteArray: %s", err.Error())
	}

	var fileList fs.FileList
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsFileListReader := bytes.NewReader(retInfo.Info)
		err = fileList.Deserialize(fsFileListReader)
		if err != nil {
			return nil, fmt.Errorf("GetFileList error: %s", err.Error())
		}
		return &fileList, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetFileProveDetails(fileHashStr string) (*fs.FsProveDetails, error) {
	fileHash := []byte(fileHashStr)
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_FILE_PROVE_DETAILS, []interface{}{fileHash},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetProveDetails result toByteArray: %s", err.Error())
	}
	var fileProveDetails fs.FsProveDetails
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fileProveDetailReader := bytes.NewReader(retInfo.Info)
		err = fileProveDetails.Deserialize(fileProveDetailReader)
		if err != nil {
			return nil, fmt.Errorf("GetProveDetails deserialize error: %s", err.Error())
		}
		return &fileProveDetails, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) AddWhiteLists(fileHashStr string, whitelists []fs.Rule) ([]byte, error) {
	if len(fileHashStr) == 0 || len(whitelists) == 0 {
		return nil, errors.New("invalid params")
	}
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	return t.WhiteListOp(fileHashStr, fs.ADD, fs.WhiteList{
		Num:  uint64(len(whitelists)),
		List: whitelists,
	})
}

func (t *Ethereum) WhiteListOp(fileHashStr string, op uint64, whiteList fs.WhiteList) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	if op != fs.ADD && op != fs.ADD_COV && op != fs.DEL && op != fs.DEL_ALL {
		return nil, errors.New("Param [op] error")
	}
	fileHash := []byte(fileHashStr)
	whiteListOp := fs.WhiteListOp{FileHash: fileHash, Op: op, List: whiteList}
	buf := new(bytes.Buffer)
	if err := whiteListOp.Serialize(buf); err != nil {
		return nil, fmt.Errorf("WhiteListOp serialize error: %s", err.Error())
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_WHITE_LIST_OP, []interface{}{buf.Bytes()},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetWhiteList(fileHashStr string) (*fs.WhiteList, error) {
	fileHash := []byte(fileHashStr)
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_WHITE_LIST, []interface{}{fileHash},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetProveDetails result toByteArray: %s", err.Error())
	}
	var whiteList fs.WhiteList
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		whiteListReader := bytes.NewReader(retInfo.Info)
		err = whiteList.Deserialize(whiteListReader)
		if err != nil {
			return nil, fmt.Errorf("GetWhiteList deserialize error: %s", err.Error())
		}
		return &whiteList, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) DeleteFile(fileHashStr string) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	fileHash := []byte(fileHashStr)
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_DELETE_FILE, []interface{}{fileHash},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) DeleteFiles(fileHashStrs []string, gasLimit uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}

	if gasLimit == 0 {
		gasLimit = sdkcom.GAS_LIMIT
	}

	fileHashes := make([]fs.FileHash, 0, len(fileHashStrs))
	for _, fileHashStr := range fileHashStrs {
		fileHashes = append(fileHashes, fs.FileHash{
			Hash: []byte(fileHashStr),
		})
	}
	fileList := fs.FileList{
		FileNum: uint64(len(fileHashStrs)),
		List:    fileHashes,
	}
	fmt.Printf("file list: %v\n", fileList)
	buf := new(bytes.Buffer)
	if err := fileList.Serialize(buf); err != nil {
		return nil, fmt.Errorf("fileList serialize error: %s", err.Error())
	}
	ret, err := t.InvokeNativeContractWithGasLimitUserDefine(t.DefAcc,
		gasLimit, fs.FS_DELETE_FILES, []interface{}{buf.Bytes()},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) PollForTxConfirmed(timeout time.Duration, txHash []byte) (bool, error) {
	return t.Client.PollForTxConfirmed(timeout, txHash)
}

func (t *Ethereum) savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
	maxProveBlockNum, defProveLevel uint64, minVolume uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_INIT,
		[]interface{}{&fs.FsSetting{FsGasPrice: fsGasPrice,
			GasPerGBPerBlock:  gasPerGBPerBlock,
			GasPerKBForRead:   gasPerKBForRead,
			GasForChallenge:   gasForChallenge,
			MaxProveBlockNum:  maxProveBlockNum,
			DefaultProveLevel: defProveLevel,
			MinVolume:         minVolume}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) NodeRegister(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_NODE_REGISTER,
		[]interface{}{&fs.FsNodeInfo{Pledge: 0, Profit: 0, Volume: volume, RestVol: 0,
			ServiceTime: serviceTime, WalletAddr: t.DefAcc.Address, NodeAddr: []byte(nodeAddr)}},
	)
	if err != nil {
		return nil, err
	}

	if t.PollForTxDuration != 0 {
		confirmed, _ := t.PollForTxConfirmed(t.PollForTxDuration, ret.ToArray())
		if !confirmed {
			return nil, fmt.Errorf("tx %x not confirmed", ret.ToArray())
		}

	}
	return ret.ToArray(), err
}

func (t *Ethereum) NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_NODE_QUERY, []interface{}{nodeWallet},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetNodeList result toByteArray: %s", err.Error())
	}

	var fsNodeInfo fs.FsNodeInfo
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsNodeInfoReader := bytes.NewReader(retInfo.Info)
		err = fsNodeInfo.Deserialize(fsNodeInfoReader)
		if err != nil {
			return nil, fmt.Errorf("FsNodeQuery error: %s", err.Error())
		}
		return &fsNodeInfo, nil
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) NodeUpdate(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_NODE_UPDATE,
		[]interface{}{&fs.FsNodeInfo{Pledge: 0, Profit: 0, Volume: volume, RestVol: 0,
			ServiceTime: serviceTime, WalletAddr: t.DefAcc.Address, NodeAddr: []byte(nodeAddr)}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) NodeCancel() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_NODE_CANCEL,
		[]interface{}{t.DefAcc.Address},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) NodeWithDrawProfit() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_NODE_WITH_DRAW_PROFIT,
		[]interface{}{t.DefAcc.Address},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) FileProve(fileHashStr string, proveData []byte, blockHeight uint64, sectorId uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	fileHash := []byte(fileHashStr)
	ret, err := t.InvokeNativeContract(
		t.DefAcc, fs.FS_FILE_PROVE,
		[]interface{}{&fs.FileProve{FileHash: fileHash,
			ProveData:   proveData,
			BlockHeight: blockHeight,
			NodeWallet:  t.DefAcc.Address,
			Profit:      0,
			SectorID:    sectorId}},
	)
	if err != nil {
		return nil, err
	}
	if t.PollForTxDuration != 0 {
		confirmed, _ := t.PollForTxConfirmed(t.PollForTxDuration, ret.ToArray())
		if !confirmed {
			return nil, fmt.Errorf("tx %x not confirmed", ret.ToArray())
		}

	}
	return ret.ToArray(), err
}

func (t *Ethereum) GenChallenge(walletAddr common.Address, hash common.Uint256, fileBlockNum, proveNum uint64) []pdp.Challenge {
	return fs.GenChallenge(walletAddr, hash, uint32(fileBlockNum), uint32(proveNum))
}

// UpdateUserSpace. user space operation for space owner.
func (t *Ethereum) UpdateUserSpace(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	params := &fs.UserSpaceParams{
		WalletAddr: t.DefAcc.Address,
		Owner:      walletAddr,
		Size:       size,
		BlockCount: blockCount,
	}
	ret, err := t.InvokeNativeContract(
		t.DefAcc, fs.FS_MANAGE_USER_SPACE,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

// GetUserSpace. get user space with wallet address
func (t *Ethereum) GetUserSpace(walletAddr common.Address) (*fs.UserSpace, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_USER_SPACE, []interface{}{walletAddr})
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetUserSpace result toByteArray: %s", err.Error())
	}

	var userspace fs.UserSpace
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		r := bytes.NewReader(retInfo.Info)
		err = userspace.Deserialize(r)
		if err != nil {
			return nil, fmt.Errorf("GetUserSpace error: %s", err.Error())
		}
		return &userspace, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetUpdateSpaceCost(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) (*usdt.State, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	params := &fs.UserSpaceParams{
		WalletAddr: t.DefAcc.Address,
		Owner:      walletAddr,
		Size:       size,
		BlockCount: blockCount,
	}
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_USER_SPACE_COST, []interface{}{params})
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetUserSpace result toByteArray: %s", err.Error())
	}
	var state usdt.State
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		r := bytes.NewReader(retInfo.Info)
		err = state.Deserialize(r)
		if err != nil {
			return nil, fmt.Errorf("GetUserSpace error: %s", err.Error())
		}
		return &state, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) DeleteUserSpace() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}

	ret, err := t.InvokeNativeContract(
		t.DefAcc, fs.FS_DELETE_USER_SPACE,
		[]interface{}{t.DefAcc.Address},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetUploadStorageFee(opt *fs.UploadOption) (*fs.StorageFee, error) {
	log.Debugf("opt :%v", opt.StorageType)
	buf := new(bytes.Buffer)
	if err := opt.Serialize(buf); err != nil {
		return nil, fmt.Errorf("UploadOption serialize error: %s", err.Error())
	}
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GETSTORAGEFEE, []interface{}{buf.Bytes()},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetUploadStorageFee result toByteArray: %s", err.Error())
	}

	var storageFee fs.StorageFee
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		reader := bytes.NewReader(retInfo.Info)
		if err = storageFee.Deserialize(reader); err != nil {
			return nil, fmt.Errorf("GetUploadStorageFee StorageFee Deserialize: %s", err.Error())
		}
		return &storageFee, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) GetDeleteFilesStorageFee(fileHashStrs []string) (uint64, error) {
	fileHashes := make([]fs.FileHash, 0, len(fileHashStrs))
	for _, fileHashStr := range fileHashStrs {
		fileHashes = append(fileHashes, fs.FileHash{
			Hash: []byte(fileHashStr),
		})
	}
	fileList := fs.FileList{
		FileNum: uint64(len(fileHashStrs)),
		List:    fileHashes,
	}
	buf := new(bytes.Buffer)
	if err := fileList.Serialize(buf); err != nil {
		return 0, fmt.Errorf("fileList serialize error: %s", err.Error())
	}
	ret, err := t.PreExecInvokeNativeContractWithSigner(t.DefAcc,
		fs.FS_DELETE_FILES, []interface{}{buf.Bytes()},
	)
	if err != nil {
		return 0, err
	}
	return ret.Gas, err
}

func (t *Ethereum) CreateSector(sectorId uint64, proveLevel uint64, size uint64, isPlots bool) ([]byte, error) {
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_CREATE_SECTOR, []interface{}{
			&fs.SectorInfo{
				NodeAddr:   t.DefAcc.Address,
				SectorID:   sectorId,
				ProveLevel: proveLevel,
				Size:       size,
				IsPlots:    isPlots,
			}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) DeleteSector(sectorId uint64) ([]byte, error) {
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_DELETE_SECTOR_INFO, []interface{}{
			&fs.SectorRef{
				NodeAddr: t.DefAcc.Address,
				SectorID: sectorId,
			}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetSectorInfo(sectorId uint64) (*fs.SectorInfo, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_SECTOR_INFO, []interface{}{
			&fs.SectorRef{
				NodeAddr: t.DefAcc.Address,
				SectorID: sectorId,
			}},
	)
	if err != nil {
		return nil, err
	}

	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetSectorInfo result toByteArray: %s", err.Error())
	}

	var sectorInfo fs.SectorInfo
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		sectorInfoReader := bytes.NewReader(retInfo.Info)
		err = sectorInfo.Deserialize(sectorInfoReader)
		if err != nil {
			return nil, fmt.Errorf("GetSectorInfo error: %s", err.Error())
		}
		return &sectorInfo, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) DeleteFileInSector(sectorId uint64, fileHashStr string) ([]byte, error) {
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_DELETE_FILE_IN_SECTOR, []interface{}{
			&fs.SectorFileRef{
				NodeAddr: t.DefAcc.Address,
				SectorID: sectorId,
				FileHash: ([]byte)(fileHashStr),
			}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetSectorInfosForNode(addr common.Address) (*fs.SectorInfos, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_SECTORS_FOR_NODE, []interface{}{addr[:]},
	)
	if err != nil {
		return nil, err
	}

	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetSectorInfo result toByteArray: %s", err.Error())
	}

	var sectorInfos fs.SectorInfos
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		sectorInfosReader := bytes.NewReader(retInfo.Info)
		err = sectorInfos.Deserialize(sectorInfosReader)
		if err != nil {
			return nil, fmt.Errorf("GetSectorInfo error: %s", err.Error())
		}
		return &sectorInfos, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) SectorProve(sectorId uint64, challengeHeight uint64, proveData []byte) ([]byte, error) {
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_SECTOR_PROVE, []interface{}{
			&fs.SectorProve{
				NodeAddr:        t.DefAcc.Address,
				SectorID:        sectorId,
				ChallengeHeight: challengeHeight,
				ProveData:       proveData,
			}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) GetUnSettledFiles(addr common.Address) (*fs.FileList, error) {
	ret, err := t.PreExecInvokeNativeContract(
		fs.FS_GET_USER_UNSETTLED_FILES, []interface{}{addr},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetUnSettledFiles result toByteArray: %s", err.Error())
	}

	var fileList fs.FileList
	retInfo := fs.DecRet(data)
	if retInfo.Ret {
		fsFileListReader := bytes.NewReader(retInfo.Info)
		err = fileList.Deserialize(fsFileListReader)
		if err != nil {
			return nil, fmt.Errorf("GetUnSettledFiles error: %s", err.Error())
		}
		return &fileList, err
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

func (t *Ethereum) DeleteUnSettledFiles() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}

	ret, err := t.InvokeNativeContract(
		t.DefAcc, fs.FS_DELETE_UNSETTLED_FILES, []interface{}{t.DefAcc.Address},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (t *Ethereum) CheckNodeSectorProveInTime(addr common.Address, sectorId uint64) ([]byte, error) {
	ret, err := t.InvokeNativeContract(t.DefAcc,
		fs.FS_CHECK_NODE_SECTOR_PROVED_INTIME, []interface{}{
			&fs.SectorRef{
				NodeAddr: addr,
				SectorID: sectorId,
			}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}
