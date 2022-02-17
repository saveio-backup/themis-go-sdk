package fs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	nodeStore "github.com/saveio/themis-go-sdk/fs/contracts/Node"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"github.com/saveio/themis/smartcontract/service/native/savefs/pdp"
	"github.com/saveio/themis/smartcontract/service/native/usdt"
	"math/big"
	"time"
)

type Ethereum struct {
	Client            *client.ClientMgr
	DefAcc            *account.Account
	PollForTxDuration time.Duration
}

// for Goerli testnet
//var address = ethCommon.HexToAddress("0xa934808a26bd08c5145cf1894d06176d3664f567")
//var privateKey, _ = crypto.HexToECDSA("c9cc08e8564d21be633f55558e9186efcbcd5f2950bee49c9faafb612ccf53fa")
//var ConfigAddress = ethCommon.HexToAddress("0x6b450d2b53Bd6C2d866F5630eDc3bB61e8016A91")
//var NodeAddress = ethCommon.HexToAddress("0x5C373B9C51f65Ec44C315A70c999F7B313Ac90c3")

// for dev mode
var address = ethCommon.HexToAddress("0x792e47e160f4ee67c17714df1c92f678640e0e4c")
var privateKey, _ = crypto.HexToECDSA("97e14a3dc8f8721172090dc5a27681e8eb3e650cb43551b43f0ce4345d4748f7")
var ConfigAddress = ethCommon.HexToAddress("0x6b450d2b53Bd6C2d866F5630eDc3bB61e8016A91")
var NodeAddress = ethCommon.HexToAddress("0x5C373B9C51f65Ec44C315A70c999F7B313Ac90c3")

func (t *Ethereum) GetSigner(value *big.Int) (*bind.TransactOpts, error) {
	ec := t.Client.GetEthClient()

	nonce, err := ec.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.From = address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value       			// in wei
	auth.GasLimit = uint64(5000000) 	// in units
	auth.GasPrice = gasPrice

	return auth, nil
}

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
	// TODO
	return common.Uint256{}, nil
}

func (t *Ethereum) InvokeNativeContractWithGasLimitUserDefine(signer *account.Account, gasLimit uint64, method string, params []interface{}) (common.Uint256, error) {
	// TODO
	return common.Uint256{}, nil
}

func (t *Ethereum) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	// TODO
	return &sdkcom.PreExecResult{}, nil
}

func (t *Ethereum) PreExecInvokeNativeContractWithSigner(signer *account.Account, method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	// TODO
	return &sdkcom.PreExecResult{}, nil
}

func (t *Ethereum) GetSetting() (*fs.FsSetting, error) {
	store, err := configStore.NewStore(ConfigAddress, t.Client.GetEthClient())
	if err != nil {
		return nil, err
	}
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
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient())
	if err != nil {
		return nil, err
	}
	list, err := store.GetNodeList(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	nodes := make([]fs.FsNodeInfo, len(list))
	for k, v := range list {
		info := fs.FsNodeInfo{
			Pledge:      v.Pledge,
			Profit:      v.Profit,
			Volume:      v.Volume,
			RestVol:     v.RestVol,
			ServiceTime: v.ServiceTime,
			WalletAddr:  common.Address(v.WalletAddr),
			NodeAddr:    v.NodeAddr[:],
		}
		nodes[k] = info
	}
	nodeList := &fs.FsNodesInfo{
		NodeNum:  uint64(len(list)),
		NodeInfo: nodes,
	}
	return nodeList, nil
}

// GetNodeListByAddrs get node list by node address
func (t *Ethereum) GetNodeListByAddrs(addrs []common.Address) (*fs.FsNodesInfo, error) {
	if len(addrs) == 0 {
		return nil, errors.New("address cannot empty")
	}
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient())
	if err != nil {
		return nil, err
	}
	info, err := store.GetNodeInfoByNodeAddr(&bind.CallOpts{}, ethCommon.Address(addrs[0]))
	if err != nil {
		return nil, err
	}
	nodes := make([]fs.FsNodeInfo, 1)
	nodeInfo := fs.FsNodeInfo{
		Pledge:      info.Pledge,
		Profit:      info.Profit,
		Volume:      info.Volume,
		RestVol:     info.RestVol,
		ServiceTime: info.ServiceTime,
		WalletAddr:  common.Address(info.WalletAddr),
		NodeAddr:    info.NodeAddr[:],
	}
	nodes[0] = nodeInfo
	nodeList := &fs.FsNodesInfo{
		NodeNum:  1,
		NodeInfo: nodes,
	}
	return nodeList, nil
}

func (t *Ethereum) ProveParamSer(rootHash []byte, fileId pdp.FileID) ([]byte, error) {
    // TODO
	return []byte{}, nil
}

func (t *Ethereum) ProveParamDes(proveParam []byte) (*fs.ProveParam, error) {
	// TODO
	var proveParamSt fs.ProveParam
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
	// TODO timeout
	var hash ethCommon.Hash
	copy(hash[:], txHash)
	receipt, err := t.Client.GetEthClient().TransactionReceipt(context.Background(), hash)
	if err != nil || receipt == nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		return false, nil
	}
	if receipt.Status == 1 {
		return true, nil
	}
	return false, errors.New("unknown status")
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

	ec := t.Client.GetEthClient()
	store, err := nodeStore.NewStore(NodeAddress, ec)
	if err != nil {
		return nil, err
	}

	node := nodeStore.NodeInfo{
		Pledge:      0,
		Profit:      0,
		Volume:      volume,
		RestVol:     0,
		ServiceTime: serviceTime,
		WalletAddr:  t.DefAcc.EthAddress,
		NodeAddr:    ethCommon.HexToAddress(nodeAddr),
	}

	pledge, err := store.CalculateNodePledge(&bind.CallOpts{}, node)
	if err != nil {
		return nil, err
	}

	signer, err := t.GetSigner(big.NewInt(int64(pledge)))
	if err != nil {
		return nil, err
	}

	register, err := store.Register(signer, node)
	if err != nil {
		return nil, err
	}

	hash := register.Hash()
	fmt.Println(hash)
	confirmed, err := t.PollForTxConfirmed(t.PollForTxDuration, hash[:])
	if err != nil {
		return nil, err
	}
	if !confirmed {
		return nil, errors.New("tx not confirmed")
	}
	return hash[:], nil
}

// NodeQuery get node info by wallet address
func (t *Ethereum) NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error) {
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient())
	if err != nil {
		return nil, err
	}
	info, err := store.GetNodeInfoByWalletAddr(&bind.CallOpts{}, ethCommon.Address(nodeWallet))
	if err != nil {
		return nil, err
	}
	nodeInfo := &fs.FsNodeInfo{
		Pledge:      info.Pledge,
		Profit:      info.Profit,
		Volume:      info.Volume,
		RestVol:     info.RestVol,
		ServiceTime: info.ServiceTime,
		WalletAddr:  common.Address(info.WalletAddr),
		NodeAddr:    info.NodeAddr[:],
	}
	return nodeInfo, nil
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
