package fs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"runtime"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	listStore "github.com/saveio/themis-go-sdk/fs/contracts/List"
	sectorStore "github.com/saveio/themis-go-sdk/fs/contracts/Sector"
	spaceStore "github.com/saveio/themis-go-sdk/fs/contracts/Space"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	fileStore "github.com/saveio/themis-go-sdk/fs/contracts/File"
	nodeStore "github.com/saveio/themis-go-sdk/fs/contracts/Node"
	pdpStore "github.com/saveio/themis-go-sdk/fs/contracts/PDP"
	proveStore "github.com/saveio/themis-go-sdk/fs/contracts/Prove"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"github.com/saveio/themis/smartcontract/service/native/savefs/pdp"
	"github.com/saveio/themis/smartcontract/service/native/usdt"
)

type EVM struct {
	Client            *client.ClientMgr
	DefAcc            *account.Account
	PollForTxDuration time.Duration
}

var _ ContractClient = (*EVM)(nil)

var ConfigAddress = ethCommon.HexToAddress("0xad973E6356f34254035300624eCacde4985f5928")

var NodeAddress = ethCommon.HexToAddress("0xa6579eA478BC818157BfEf4dA1236Db4513b1414")

var SectorAddress = ethCommon.HexToAddress("0x8baD872E3623f6b92357474640975e7fDF4F9BcC")

var SpaceAddress = ethCommon.HexToAddress("0xA8C748e48418b8eBDfBD4826a0c0f2e519A68765")

var FileAddress = ethCommon.HexToAddress("0xAe7651927963e62b2F01fe51d8C95b0D00D11812")

var FileExtraAddress = ethCommon.HexToAddress("0xB2f4Bce0069d35aa3fAA1cf052Bf93aA1d01fdC3")

var ListAddress = ethCommon.HexToAddress("0xFFaFf9145c42e04E6654211713d84D2368d947E1")

var ProveAddress = ethCommon.HexToAddress("0xB658B307F3d7B9fcFb6697cf3c746Cb21427d053")

var ProveExtraAddress = ethCommon.HexToAddress("0xAd151b6c3821F07A359A3928b45Be2018a704e8b")

var PDPAddress = ethCommon.HexToAddress("0x546869055C8B668Cc21E6683A78Ed231A16b7669")

func (t *EVM) GetSigner(value *big.Int) (*bind.TransactOpts, error) {
	ec := t.Client.GetEthClient().Client
	nonce, err := ec.PendingNonceAt(context.TODO(), t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}
	chainID, err := ec.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}
	keyBytes := t.DefAcc.GetEthPrivateKey()
	keyStr := hexutil.Encode(keyBytes)
	key, err := crypto.HexToECDSA(keyStr[2:])
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	auth.From = t.DefAcc.EthAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value                 // in wei
	auth.GasLimit = uint64(10_000_000) // in units
	auth.GasPrice = gasPrice

	gas := new(big.Int).Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit)))
	gas = new(big.Int).Add(auth.Value, gas)
	log.Debugf("get signer with gas: %v, %v, %v, %v", auth.GasPrice, auth.GasLimit, auth.Value, gas)
	return auth, nil
}

func (t *EVM) SetDefaultAccount(acc *account.Account) {
	t.DefAcc = acc
}

func (t *EVM) GetDefaultAccount() *account.Account {
	return t.DefAcc
}

func (t *EVM) GetClient() *client.ClientMgr {
	return t.Client
}

func (t *EVM) InvokeNativeContract(signer *account.Account, method string, params []interface{}) (common.Uint256, error) {
	// TODO
	return common.Uint256{}, nil
}

func (t *EVM) InvokeNativeContractWithGasLimitUserDefine(signer *account.Account, gasLimit uint64, method string, params []interface{}) (common.Uint256, error) {
	// TODO
	return common.Uint256{}, nil
}

func (t *EVM) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	// TODO
	return &sdkcom.PreExecResult{}, nil
}

func (t *EVM) PreExecInvokeNativeContractWithSigner(signer *account.Account, method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	// TODO
	return &sdkcom.PreExecResult{}, nil
}

func (t *EVM) PollForTxConfirmed(timeout time.Duration, txHash []byte) (bool, error) {
	var hash ethCommon.Hash
	copy(hash[:], txHash)
	interval := time.Duration(sdkcom.POLL_TX_INTERVAL) * time.Second
	secs := int(timeout / interval)
	if secs <= 0 {
		secs = 1
	}
	for i := 0; i < secs; i++ {
		time.Sleep(interval)
		receipt, err := t.Client.GetEthClient().Client.TransactionReceipt(context.Background(), hash)
		if err != nil || receipt == nil {
			return false, err
		}
		if receipt.Status == 1 {
			return true, nil
		}
	}
	return false, nil
}

func (t *EVM) GetSetting() (*fs.FsSetting, error) {
	store, err := configStore.NewStore(ConfigAddress, t.Client.GetEthClient().Client)
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

func (t *EVM) savefsInit(fsGasPrice, gasPerGBPerBlock, gasPerKBForRead, gasForChallenge,
	maxProveBlockNum, defProveLevel uint64, minVolume uint64) ([]byte, error) {
	// TODO not be implemented in contract
	return []byte{}, nil
}

func (t *EVM) GetNodeList() (*fs.FsNodesInfo, error) {
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient().Client)
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
func (t *EVM) GetNodeListByAddrs(addrs []common.Address) (*fs.FsNodesInfo, error) {
	if len(addrs) == 0 {
		return nil, errors.New("address cannot empty")
	}
	nodes := make([]fs.FsNodeInfo, 0)
	for _, addr := range addrs {
		address := ethCommon.BytesToAddress(addr[:])
		store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient().Client)
		if err != nil {
			return nil, err
		}
		info, err := store.GetNodeInfoByWalletAddr(&bind.CallOpts{}, address)
		if err != nil {
			return nil, err
		}
		nodeInfo := fs.FsNodeInfo{
			Pledge:      info.Pledge,
			Profit:      info.Profit,
			Volume:      info.Volume,
			RestVol:     info.RestVol,
			ServiceTime: info.ServiceTime,
			WalletAddr:  common.Address(info.WalletAddr),
			NodeAddr:    info.NodeAddr[:],
		}
		nodes = append(nodes, nodeInfo)
	}
	nodeList := &fs.FsNodesInfo{
		NodeNum:  uint64(len(nodes)),
		NodeInfo: nodes,
	}
	return nodeList, nil
}

// NodeQuery get node info by wallet address
func (t *EVM) NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error) {
	address := ethCommon.BytesToAddress(nodeWallet[:])
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient().Client)
	if err != nil {
		return nil, err
	}
	info, err := store.GetNodeInfoByWalletAddr(&bind.CallOpts{}, address)
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

func (t *EVM) NodeRegister(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
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
		NodeAddr:    []byte(nodeAddr),
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
	return TxResult(ec, register)
}

func (t *EVM) NodeUpdate(volume uint64, serviceTime uint64, nodeAddr string) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := nodeStore.NewStore(NodeAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
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
		NodeAddr:    []byte(nodeAddr),
	}
	pledge, err := store.GetPledgeUpdate(&bind.CallOpts{}, node)
	if err != nil {
		return nil, err
	}
	if pledge < 0 {
		pledge = 0
	}
	signer, err = t.GetSigner(big.NewInt(pledge))
	if err != nil {
		return nil, err
	}
	tx, err := store.NodeUpdate(signer, node)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) NodeCancel() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := nodeStore.NewStore(NodeAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	tx, err := store.Cancel(signer, t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) NodeWithDrawProfit() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := nodeStore.NewStore(NodeAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	cancel, err := store.WithDrawProfit(signer, t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, cancel)
}

func copyUploadOption(opt *fs.UploadOption) fileStore.UploadOption {
	list := make([]fileStore.WhiteList, opt.WhiteList.Num)
	for k, v := range opt.WhiteList.List {
		addr := ethCommon.BytesToAddress(v.Addr[:])
		list[k] = fileStore.WhiteList{
			Addr:         addr,
			BaseHeight:   v.BaseHeight,
			ExpireHeight: v.ExpireHeight,
		}
	}
	o := fileStore.UploadOption{
		FileDesc:        opt.FileDesc,
		FileSize:        opt.FileSize,
		ProveInterval:   opt.ProveInterval,
		ProveLevel:      opt.ProveLevel,
		ExpiredHeight:   big.NewInt(int64(opt.ExpiredHeight)),
		Privilege:       opt.Privilege,
		CopyNum:         opt.CopyNum,
		Encrypt:         opt.Encrypt,
		EncryptPassword: opt.EncryptPassword,
		RegisterDNS:     opt.RegisterDNS,
		BindDNS:         opt.BindDNS,
		DnsURL:          opt.DnsURL,
		WhiteList:       list,
		Share:           opt.Share,
		StorageType:     uint8(opt.StorageType),
	}
	return o
}

func (t *EVM) GetUploadStorageFee(opt *fs.UploadOption) (*fs.StorageFee, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	option := copyUploadOption(opt)
	fee, err := store.GetUploadStorageFee(&bind.CallOpts{}, option)
	if err != nil {
		return nil, err
	}
	storageFee := &fs.StorageFee{
		TxnFee:        fee.TxnFee,
		SpaceFee:      fee.SpaceFee,
		ValidationFee: fee.ValidationFee,
	}
	return storageFee, nil
}

func (t *EVM) GetDeleteFilesStorageFee(fileHashStrs []string) (uint64, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return 0, err
	}
	log.Info(store)
	// TODO how get the fees
	return 0, nil
}

func (t *EVM) FileRenew(fileHashStr string, renewTimes uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	fileHash := []byte(fileHashStr)
	opt := fileStore.FileReNewInfo{
		FileHash:   fileHash,
		FromAddr:   t.DefAcc.EthAddress,
		ReNewTimes: renewTimes,
	}
	reNew, err := store.FileReNew(signer, opt)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, reNew)
}

func (t *EVM) ProveParamSer(rootHash []byte, fileId pdp.FileID) ([]byte, error) {
	var proveParam fs.ProveParam
	proveParam.RootHash = rootHash
	proveParam.FileID = fileId
	bf := new(bytes.Buffer)
	if err := proveParam.Serialize(bf); err != nil {
		return nil, fmt.Errorf("ProveParam Serialize error: %s", err.Error())
	}
	return bf.Bytes(), nil
}

func (t *EVM) ProveParamDes(proveParam []byte) (*fs.ProveParam, error) {
	var proveParamSt fs.ProveParam
	reader := bytes.NewReader(proveParam)
	if err := proveParamSt.Deserialize(reader); err != nil {
		return nil, fmt.Errorf("ProveParamDes Deserialize error: %s", err.Error())
	}
	return &proveParamSt, nil
}

func (t *EVM) StoreFile(fileHashStr, blocksRoot string, blockNum uint64,
	blockSize uint64, proveLevel uint64, expiredHeight uint64, copyNum uint64,
	fileDesc []byte, privilege uint64, proveParam []byte, storageType uint64, realFileSize uint64,
	primaryNodes, candidateNodes []common.Address, plotInfo *fs.PlotInfo, url string) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	// TODO wangyu
	signer, err := t.GetSigner(big.NewInt(10000))
	if err != nil {
		return nil, err
	}

	primary := make([]ethCommon.Address, len(primaryNodes))
	for k, v := range primaryNodes {
		primary[k] = ethCommon.Address(v)
	}
	candidates := make([]ethCommon.Address, len(candidateNodes))
	for k, v := range candidateNodes {
		candidates[k] = ethCommon.BytesToAddress(v[:])
	}
	fileHash := []byte(fileHashStr)
	var evmPlotInfo fileStore.PlotInfo
	if plotInfo == nil {
		evmPlotInfo = fileStore.PlotInfo{
			NumberID:   0,
			StartNonce: 0,
			Nonces:     0,
		}
	} else {
		evmPlotInfo = fileStore.PlotInfo{
			NumberID:   plotInfo.NumericID,
			StartNonce: plotInfo.StartNonce,
			Nonces:     plotInfo.Nonces,
		}
	}
	f := fileStore.FileInfo{
		FileHash:       fileHash,
		BlocksRoot:     []byte(blocksRoot),
		FileOwner:      t.DefAcc.EthAddress,
		FileDesc:       fileDesc,
		Privilege:      privilege,
		FileBlockNum:   blockNum,
		FileBlockSize:  blockSize,
		ProveLevel:     uint8(proveLevel),
		ProveTimes:     0,
		ExpiredHeight:  big.NewInt(int64(expiredHeight)),
		CopyNum:        copyNum,
		Deposit:        0,
		FileProveParam: proveParam,
		ProveBlockNum:  0,
		BlockHeight:    big.NewInt(0),
		ValidFlag:      true,
		StorageType:    uint8(storageType),
		RealFileSize:   realFileSize,
		PrimaryNodes:   primary,
		CandidateNodes: candidates,
		IsPlotFile:     plotInfo != nil,
		PlotInfo:       evmPlotInfo,
	}
	tx, err := store.StoreFile(signer, f)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func copyFileInfo(info fileStore.FileInfo) fs.FileInfo {
	pa := make([]common.Address, len(info.PrimaryNodes))
	for k, v := range info.PrimaryNodes {
		pa[k] = common.Address(v)
	}
	pnode := fs.NodeList{
		AddrNum:  0,
		AddrList: pa,
	}
	ca := make([]common.Address, len(info.CandidateNodes))
	for k, v := range info.PrimaryNodes {
		ca[k] = common.Address(v)
	}
	cnode := fs.NodeList{
		AddrNum:  0,
		AddrList: ca,
	}
	p := &fs.PlotInfo{
		NumericID:  info.PlotInfo.NumberID,
		StartNonce: info.PlotInfo.StartNonce,
		Nonces:     info.PlotInfo.Nonces,
	}
	fileInfo := fs.FileInfo{
		FileHash:       info.FileHash,
		FileOwner:      common.Address(info.FileOwner),
		FileDesc:       info.FileDesc,
		Privilege:      info.Privilege,
		FileBlockNum:   info.FileBlockNum,
		FileBlockSize:  info.FileBlockSize,
		ProveInterval:  info.ProveInterval,
		ProveTimes:     info.ProveTimes,
		ExpiredHeight:  info.ExpiredHeight.Uint64(),
		CopyNum:        info.CopyNum,
		Deposit:        info.Deposit,
		FileProveParam: info.FileProveParam,
		ProveBlockNum:  info.ProveBlockNum,
		BlockHeight:    info.BlockHeight.Uint64(),
		ValidFlag:      info.ValidFlag,
		StorageType:    uint64(info.StorageType),
		RealFileSize:   info.RealFileSize,
		PrimaryNodes:   pnode,
		CandidateNodes: cnode,
		BlocksRoot:     info.BlocksRoot,
		ProveLevel:     uint64(info.ProveLevel),
		SectorRefs:     nil, // TODO get from sector info?
		IsPlotFile:     info.IsPlotFile,
		PlotInfo:       p,
	}
	return fileInfo
}

func (t *EVM) GetFileInfo(fileHashStr string) (*fs.FileInfo, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	info, err := store.GetFileInfo(&bind.CallOpts{}, []byte(fileHashStr))
	if err != nil {
		return nil, err
	}
	if info.ExpiredHeight.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("file not exist")
	}
	fileInfo := copyFileInfo(info)
	return &fileInfo, nil
}

func (t *EVM) GetFileInfos(fileHashStrs []string) (*fs.FileInfoList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	list := make([][]byte, len(fileHashStrs))
	for k, v := range fileHashStrs {
		list[k] = []byte(v)
	}
	info, err := store.GetFileInfos(&bind.CallOpts{}, list)
	if err != nil {
		return nil, err
	}
	rlist := make([]fs.FileInfo, len(info))
	for k, v := range info {
		rlist[k] = copyFileInfo(v)
	}
	ret := &fs.FileInfoList{
		FileNum: uint64(len(info)),
		List:    rlist,
	}
	return ret, nil
}

func (t *EVM) ChangeFileOwner(fileHashStr string, newOwner common.Address) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	ownerChange := fileStore.OwnerChange{
		FileHash: []byte(fileHashStr),
		CurOwner: t.DefAcc.EthAddress,
		NewOwner: ethCommon.Address(newOwner),
	}
	tx, err := store.ChangeFileOwner(signer, ownerChange)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) ChangeFilePrivilege(fileHashStr string, newPrivilege uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	pri := fileStore.PriChange{
		FileHash:  []byte(fileHashStr),
		Privilege: newPrivilege,
	}
	tx, err := store.ChangeFilePrivilege(signer, pri)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) GetFileList(addr common.Address) (*fs.FileList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetFileList(&bind.CallOpts{}, ethCommon.BytesToAddress(addr[:]))
	if err != nil {
		return nil, err
	}
	flist := make([]fs.FileHash, len(list))
	for k, v := range list {
		flist[k] = fs.FileHash{Hash: v}
	}
	rlist := &fs.FileList{
		FileNum: uint64(len(list)),
		List:    flist,
	}
	return rlist, nil
}

func (t *EVM) GetUnprovePrimaryFileList(addr common.Address) (*fs.FileList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	files, err := store.GetUnProvePrimaryFiles(&bind.CallOpts{}, ethCommon.BytesToAddress(addr[:]))
	if err != nil {
		return nil, err
	}
	flist := make([]fs.FileHash, len(files))
	for k, v := range files {
		flist[k] = fs.FileHash{Hash: v}
	}
	list := &fs.FileList{
		FileNum: uint64(len(files)),
		List:    nil,
	}
	return list, nil
}

func (t *EVM) GetUnProveCandidateFileList(addr common.Address) (*fs.FileList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	files, err := store.GetUnProveCandidateFiles(&bind.CallOpts{}, ethCommon.BytesToAddress(addr[:]))
	if err != nil {
		return nil, err
	}
	flist := make([]fs.FileHash, len(files))
	for k, v := range files {
		flist[k] = fs.FileHash{Hash: v}
	}
	list := &fs.FileList{
		FileNum: uint64(len(files)),
		List:    nil,
	}
	return list, nil
}

func (t *EVM) GetFileProveDetails(fileHashStr string) (*fs.FsProveDetails, error) {
	ec := t.Client.GetEthClient().Client
	store, err := proveStore.NewStore(ProveAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetProveDetailList(&bind.CallOpts{}, []byte(fileHashStr))
	if err != nil {
		return nil, err
	}
	rlist := make([]fs.ProveDetail, len(list))
	for k, v := range list {
		rlist[k] = fs.ProveDetail{
			NodeAddr:    v.NodeAddr[:],
			WalletAddr:  common.Address(v.WalletAddr),
			ProveTimes:  v.ProveTimes,
			BlockHeight: v.BlockHeight.Uint64(),
			Finished:    v.Finished,
		}
	}
	ret := &fs.FsProveDetails{
		CopyNum:        uint64(len(list)), // TODO ?
		ProveDetailNum: uint64(len(list)),
		ProveDetails:   rlist,
	}
	return ret, nil
}

func (t *EVM) AddWhiteLists(fileHashStr string, whitelists []fs.Rule) ([]byte, error) {
	if len(fileHashStr) == 0 || len(whitelists) == 0 {
		return nil, errors.New("invalid params")
	}
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := listStore.NewStore(ListAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}

	list := make([]listStore.WhiteList, len(whitelists))
	for k, v := range whitelists {
		list[k] = listStore.WhiteList{
			Addr:         ethCommon.Address(v.Addr),
			BaseHeight:   v.BaseHeight,
			ExpireHeight: v.ExpireHeight,
		}
	}
	param := listStore.WhiteListParams{
		FileHash: []byte(fileHashStr),
		Op:       fs.ADD,
		List:     list,
	}
	tx, err := store.WhiteListOperate(signer, param)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) WhiteListOp(fileHashStr string, op uint64, whiteList fs.WhiteList) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	if op != fs.ADD && op != fs.ADD_COV && op != fs.DEL && op != fs.DEL_ALL {
		return nil, errors.New("Param [op] error")
	}
	ec := t.Client.GetEthClient().Client
	store, err := listStore.NewStore(ListAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	list := make([]listStore.WhiteList, whiteList.Num)
	for k, v := range whiteList.List {
		list[k] = listStore.WhiteList{
			Addr:         ethCommon.Address(v.Addr),
			BaseHeight:   v.BaseHeight,
			ExpireHeight: v.ExpireHeight,
		}
	}
	param := listStore.WhiteListParams{
		FileHash: []byte(fileHashStr),
		Op:       fs.ADD,
		List:     list,
	}
	tx, err := store.WhiteListOperate(signer, param)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) GetWhiteList(fileHashStr string) (*fs.WhiteList, error) {
	fileHash := []byte(fileHashStr)
	ec := t.Client.GetEthClient().Client
	store, err := listStore.NewStore(ListAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetWhiteList(&bind.CallOpts{}, fileHash)
	if err != nil {
		return nil, err
	}
	rules := make([]fs.Rule, len(list))
	for k, v := range list {
		rules[k] = fs.Rule{
			Addr:         common.Address(v.Addr),
			BaseHeight:   v.BaseHeight,
			ExpireHeight: v.ExpireHeight,
		}
	}
	ret := &fs.WhiteList{
		Num:  uint64(len(list)),
		List: rules,
	}
	return ret, nil
}

func (t *EVM) DeleteFile(fileHashStr string) ([]byte, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	tx, err := store.DeleteFile(signer, []byte(fileHashStr))
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) DeleteFiles(fileHashStrs []string, gasLimit uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	p := make([][]byte, len(fileHashStrs))
	for k, v := range fileHashStrs {
		p[k] = []byte(v)
	}
	tx, err := store.DeleteFiles(signer, p)
	if err != nil {
		return nil, err
	}
	return TxResultWithError(ec, tx, fileStore.StoreMetaData.ABI)
}

func (t *EVM) FileProve(fileHashStr string, proveData []byte, blockHeight uint64, sectorId uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	fileHash := []byte(fileHashStr)
	ec := t.Client.GetEthClient().Client
	store, err := proveStore.NewStore(ProveAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	param := proveStore.FileProveParams{
		FileHash:    fileHash,
		ProveData:   proveData,
		BlockHeight: big.NewInt(int64(blockHeight)),
		NodeWallet:  t.DefAcc.EthAddress,
		Profit:      0,
		SectorID:    sectorId,
	}
	tx, err := store.FileProve(signer, param)
	if err != nil {
		return nil, err
	}
	return TxResultWithError(ec, tx, proveStore.StoreMetaData.ABI)
}

func (t *EVM) GenChallenge(walletAddr common.Address, hash common.Uint256, fileBlockNum, proveNum uint64) []pdp.Challenge {
	ec := t.Client.GetEthClient().Client
	store, err := pdpStore.NewStore(PDPAddress, ec)
	if err != nil {
		return nil
	}
	param := pdpStore.GenChallengeParams{
		WalletAddr:   ethCommon.BytesToAddress(walletAddr[:]),
		HashValue:    hash[:],
		FileBlockNum: fileBlockNum,
		ProveNum:     proveNum,
	}
	challenge, err := store.GenChallenge(&bind.CallOpts{}, param)
	if err != nil {
		return nil
	}
	ret := make([]pdp.Challenge, proveNum)
	for k, v := range challenge {
		ret[k] = pdp.Challenge{
			Index: v.Index,
			Rand:  v.Rand,
		}
	}
	return ret
}

// UpdateUserSpace. user space operation for space owner.
func (t *EVM) UpdateUserSpace(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	address := ethCommon.BytesToAddress(walletAddr[:])
	ec := t.Client.GetEthClient().Client
	store, err := spaceStore.NewStore(SpaceAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	param := spaceStore.UserSpaceParams{
		WalletAddr: t.DefAcc.EthAddress,
		Owner:      address,
		Size: spaceStore.UserSpaceOperation{
			Type:  uint8(size.Type),
			Value: size.Value,
		},
		BlockCount: spaceStore.UserSpaceOperation{
			Type:  uint8(blockCount.Type),
			Value: blockCount.Value,
		},
	}
	tx, err := store.GetUpdateCost(signer, param)
	if err != nil {
		return nil, err
	}
	mined, err := bind.WaitMined(context.Background(), ec, tx)
	var res spaceStore.TransferState
	for _, vLog := range mined.Logs {
		contractAbi, err := abi.JSON(strings.NewReader(spaceStore.StoreMetaData.ABI))
		if err != nil {
			return nil, err
		}
		data, err := contractAbi.Unpack("GetUpdateCostEvent", vLog.Data)
		if err != nil {
			return nil, err
		}
		for _, v := range data {
			res = spaceStore.TransferState(v.(struct {
				From  ethCommon.Address `json:"From"`
				To    ethCommon.Address `json:"To"`
				Value uint64            `json:"Value"`
			}))
		}
	}
	signer, err = t.GetSigner(big.NewInt(int64(res.Value)))
	if err != nil {
		return nil, err
	}
	tx, err = store.ManageUserSpace(signer, param)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

// GetUserSpace. get user space with wallet address
func (t *EVM) GetUserSpace(walletAddr common.Address) (*fs.UserSpace, error) {
	ec := t.Client.GetEthClient().Client
	store, err := spaceStore.NewStore(SpaceAddress, ec)
	if err != nil {
		return nil, err
	}
	address := ethCommon.BytesToAddress(walletAddr[:])
	space, err := store.GetUserSpace(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	ret := &fs.UserSpace{
		Used:         space.Used,
		Remain:       space.Remain,
		ExpireHeight: space.ExpireHeight.Uint64(),
		Balance:      space.Balance,
		UpdateHeight: space.UpdateHeight.Uint64(),
	}
	return ret, nil
}

func (t *EVM) GetUpdateSpaceCost(walletAddr common.Address, size, blockCount *fs.UserSpaceOperation) (*usdt.State, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	address := ethCommon.BytesToAddress(walletAddr[:])
	ec := t.Client.GetEthClient().Client
	store, err := spaceStore.NewStore(SpaceAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	param := spaceStore.UserSpaceParams{
		WalletAddr: address,
		Owner:      t.DefAcc.EthAddress,
		Size: spaceStore.UserSpaceOperation{
			Type:  uint8(size.Type),
			Value: size.Value,
		},
		BlockCount: spaceStore.UserSpaceOperation{
			Type:  uint8(blockCount.Type),
			Value: blockCount.Value,
		},
	}
	cost, err := store.GetUpdateCost(signer, param)
	if err != nil {
		return nil, err
	}
	log.Debugf("GetUpdateSpaceCost: %+v", cost)
	// TODO wangyu get res
	state := &usdt.State{
		From:  common.Address{},
		To:    common.Address{},
		Value: 0,
	}
	return state, nil
}

func (t *EVM) DeleteUserSpace() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := spaceStore.NewStore(SpaceAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	tx, err := store.DeleteUserSpace(signer, t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) CreateSector(sectorId uint64, proveLevel uint64, size uint64, isPlots bool) ([]byte, error) {
	ec := t.Client.GetEthClient().Client
	store, err := sectorStore.NewStore(SectorAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	sector := sectorStore.SectorInfo{
		NodeAddr:         t.DefAcc.EthAddress,
		SectorID:         sectorId,
		Size:             size,
		ProveLevel:       uint8(proveLevel),
		IsPlots:          isPlots,
		Used:             0,
		FirstProveHeight: big.NewInt(0),
		NextProveHeight:  big.NewInt(0),
		TotalBlockNum:    0,
		FileNum:          0,
		GroupNum:         0,
		FileList:         nil,
	}
	tx, err := store.CreateSector(signer, sector)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) DeleteSector(sectorId uint64) ([]byte, error) {
	ec := t.Client.GetEthClient().Client
	store, err := sectorStore.NewStore(SectorAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	sector := sectorStore.SectorRef{
		NodeAddr: t.DefAcc.EthAddress,
		SectorId: sectorId,
	}
	tx, err := store.DeleteSector(signer, sector)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) GetSectorInfo(sectorId uint64) (*fs.SectorInfo, error) {
	ec := t.Client.GetEthClient().Client
	store, err := sectorStore.NewStore(SectorAddress, ec)
	if err != nil {
		return nil, err
	}
	sector := sectorStore.SectorRef{
		NodeAddr: t.DefAcc.EthAddress,
		SectorId: sectorId,
	}
	info, err := store.GetSectorInfo(&bind.CallOpts{}, sector)
	if err != nil {
		return nil, err
	}
	hlist := make([]fs.FileHash, len(info.FileList))
	for k, v := range info.FileList {
		hlist[k] = fs.FileHash{Hash: v}
	}
	flist := fs.FileList{
		FileNum: uint64(len(info.FileList)),
		List:    hlist,
	}
	ret := &fs.SectorInfo{
		NodeAddr:         common.Address(info.NodeAddr),
		SectorID:         info.SectorID,
		Size:             info.Size,
		Used:             info.Used,
		ProveLevel:       uint64(info.ProveLevel),
		FirstProveHeight: info.FirstProveHeight.Uint64(),
		NextProveHeight:  info.NextProveHeight.Uint64(),
		TotalBlockNum:    info.TotalBlockNum,
		FileNum:          info.FileNum,
		GroupNum:         info.GroupNum,
		IsPlots:          info.IsPlots,
		FileList:         flist,
	}
	return ret, err
}

func (t *EVM) DeleteFileInSector(sectorId uint64, fileHashStr string) ([]byte, error) {
	ec := t.Client.GetEthClient().Client
	store, err := sectorStore.NewStore(SectorAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	s := sectorStore.SectorInfo{
		NodeAddr: t.DefAcc.EthAddress,
		SectorID: sectorId,
	}
	f := sectorStore.FileInfo{
		FileHash:  []byte(fileHashStr),
		FileOwner: t.DefAcc.EthAddress,
	}
	tx, err := store.DeleteFileFromSector(signer, s, f) // TODO why todo?
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func copeSectorInfo(info sectorStore.SectorInfo) *fs.SectorInfo {
	hlist := make([]fs.FileHash, len(info.FileList))
	for k, v := range info.FileList {
		hlist[k] = fs.FileHash{Hash: v}
	}
	flist := fs.FileList{
		FileNum: uint64(len(info.FileList)),
		List:    hlist,
	}
	ret := &fs.SectorInfo{
		NodeAddr:         common.Address(info.NodeAddr),
		SectorID:         info.SectorID,
		Size:             info.Size,
		Used:             info.Used,
		ProveLevel:       uint64(info.ProveLevel),
		FirstProveHeight: info.FirstProveHeight.Uint64(),
		NextProveHeight:  info.NextProveHeight.Uint64(),
		TotalBlockNum:    info.TotalBlockNum,
		FileNum:          info.FileNum,
		GroupNum:         info.GroupNum,
		IsPlots:          info.IsPlots,
		FileList:         flist,
	}
	return ret
}

func (t *EVM) GetSectorInfosForNode(addr common.Address) (*fs.SectorInfos, error) {
	address := ethCommon.BytesToAddress(addr[:])
	ec := t.Client.GetEthClient().Client
	store, err := sectorStore.NewStore(SectorAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetSectorsForNode(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	infos := make([]*fs.SectorInfo, len(list))
	for k, v := range list {
		infos[k] = copeSectorInfo(v)
	}
	ret := &fs.SectorInfos{
		SectorCount: uint64(len(infos)),
		Sectors:     infos,
	}
	return ret, err
}

func (t *EVM) SectorProve(sectorId uint64, challengeHeight uint64, proveData []byte) ([]byte, error) {
	ec := t.Client.GetEthClient().Client
	store, err := proveStore.NewStore(ProveAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	param := proveStore.SectorProveParams{
		NodeAddr:        t.DefAcc.EthAddress,
		SectorID:        sectorId,
		ChallengeHeight: challengeHeight,
		ProveData:       proveData,
	}
	tx, err := store.SectorProve(signer, param)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) GetUnSettledFiles(addr common.Address) (*fs.FileList, error) {
	address := ethCommon.BytesToAddress(addr[:])
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetUnSettledFileList(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}
	rlist := make([]fs.FileHash, len(list))
	for k, v := range list {
		rlist[k] = fs.FileHash{Hash: v}
	}
	ret := &fs.FileList{
		FileNum: uint64(len(list)),
		List:    rlist,
	}
	return ret, nil
}

func (t *EVM) DeleteUnSettledFiles() ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fileStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	tx, err := store.DeleteUnsettledFiles(signer, t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) CheckNodeSectorProveInTime(addr common.Address, sectorId uint64) ([]byte, error) {
	address := ethCommon.BytesToAddress(addr[:])
	ec := t.Client.GetEthClient().Client
	store, err := proveStore.NewStore(ProveAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	param := proveStore.SectorRef{
		NodeAddr: address,
		SectorId: sectorId,
	}
	tx, err := store.CheckNodeSectorProvedInTime(signer, param)
	if err != nil {
		return nil, err
	}
	return TxResult(ec, tx)
}

func (t *EVM) GetEventsByBlockHeight(blockHeight *big.Int) ([]map[string]interface{}, error) {
	res := make([]map[string]interface{}, 0)
	query := ethereum.FilterQuery{
		FromBlock: blockHeight,
		ToBlock:   blockHeight,
		Addresses: []ethCommon.Address{
			FileAddress,
			SectorAddress,
		},
	}
	ec := t.Client.GetEthClient().Client
	logs, err := ec.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	fsAbi, err := abi.JSON(strings.NewReader(fileStore.StoreMetaData.ABI))
	if err != nil {
		return nil, err
	}
	sectorAbi, err := abi.JSON(strings.NewReader(sectorStore.StoreMetaData.ABI))
	if err != nil {
		return nil, err
	}
	for _, vLog := range logs {
		event := make(map[string]interface{})
		err := sectorAbi.UnpackIntoMap(event, "CreateSectorEvent", vLog.Data)
		if err != nil {
			continue
		}
		event["eventName"] = GetFsEventNameByType(event["eventType"].(uint8))
		if event["eventName"] != "createSector" {
			continue
		}
		res = append(res, event)

		event = make(map[string]interface{})
		err = sectorAbi.UnpackIntoMap(event, "DeleteSectorEvent", vLog.Data)
		if err != nil {
			continue
		}
		event["eventName"] = GetFsEventNameByType(event["eventType"].(uint8))
		if event["eventName"] != "deleteSector" {
			continue
		}
		res = append(res, event)

		event = make(map[string]interface{})
		err = fsAbi.UnpackIntoMap(event, "DeleteFileEvent", vLog.Data)
		if err != nil {
			continue
		}
		event["eventName"] = GetFsEventNameByType(event["eventType"].(uint8))
		if event["eventName"] != "deleteFile" {
			continue
		}
		res = append(res, event)

		event = make(map[string]interface{})
		err = fsAbi.UnpackIntoMap(event, "DeleteFilesEvent", vLog.Data)
		if err != nil {
			continue
		}
		event["eventName"] = GetFsEventNameByType(event["eventType"].(uint8))
		if event["eventName"] != "deleteFiles" {
			continue
		}
		res = append(res, event)
	}
	return res, nil
}

func TxResult(ec *ethclient.Client, tx *types.Transaction) ([]byte, error) {
	mined, err := bind.WaitMined(context.Background(), ec, tx)
	if err != nil {
		return nil, err
	}
	if mined.Status != types.ReceiptStatusSuccessful {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		var funcName string
		if ok && details != nil {
			funcName = details.Name()
		}
		log.Errorf("tx failed: %s, func: %s, logs: %d", tx.Hash().String(), funcName, len(mined.Logs))
		return nil, errors.New("tx failed")
	}
	hash := tx.Hash()
	return hash[:], nil
}

func TxResultWithError(ec *ethclient.Client, tx *types.Transaction, abiData string) ([]byte, error) {
	mined, err := bind.WaitMined(context.Background(), ec, tx)
	if err != nil {
		return nil, err
	}
	if mined.Status != types.ReceiptStatusSuccessful {
		return nil, errors.New("tx mined failed")
	}
	var errMsg string
	for _, v := range mined.Logs {
		contractAbi, err := abi.JSON(strings.NewReader(abiData))
		if err != nil {
			return nil, err
		}
		data, err := contractAbi.Unpack("FsError", v.Data)
		for _, v := range data {
			errMsg += v.(string)
		}
	}
	hash := tx.Hash()
	if errMsg != "" {
		return hash[:], errors.New(errMsg)
	}
	return hash[:], nil
}

func GetFsEventNameByType(t uint8) string {
	t += 1
	switch t {
	case fs.EVENT_FS_STORE_FILE:
		return "uploadFile"
	case fs.EVENT_FS_DELETE_FILE:
		return "deleteFile"
	case fs.EVENT_FS_DELETE_FILES:
		return "deleteFiles"
	case fs.EVENT_FS_SET_USER_SPACE:
		return "setUserSpace"
	case fs.EVENT_FS_REG_NODE:
		return "registerNode"
	case fs.EVENT_FS_UN_REG_NODE:
		return "unregisterNode"
	case fs.EVENT_FS_PROVE_FILE:
		return "proveFile"
	case fs.EVENT_FS_FILE_PDP_SUCCESS:
		return "filePdpSuccess"
	case fs.EVENT_FS_CREATE_SECTOR:
		return "createSector"
	case fs.EVENT_FS_DELETE_SECTOR:
		return "deleteSector"
	}
	return ""
}
