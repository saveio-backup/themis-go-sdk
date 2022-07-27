package fs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	listStore "github.com/saveio/themis-go-sdk/fs/contracts/List"
	sectorStore "github.com/saveio/themis-go-sdk/fs/contracts/Sector"
	spaceStore "github.com/saveio/themis-go-sdk/fs/contracts/Space"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	fsStore "github.com/saveio/themis-go-sdk/fs/contracts/FileSystem"
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

// for dev mode
var address = ethCommon.HexToAddress("0xB22f5E275a30826B5718A3a06612567701DD9126")
var privateKey, _ = crypto.HexToECDSA("58d923527aac8f3e792f341e7e186f5aa9a8555ec7426bb639415710015460d2")

var ConfigAddress = ethCommon.HexToAddress("0xA009F5dFA3D2c89b14E7a64ef1b3A954C1a8AACb")
var NodeAddress = ethCommon.HexToAddress("0xE3a76fCa97bdFCa0274388b91901028Cc90635C8")
var SectorAddress = ethCommon.HexToAddress("0xFD6c336668Df20A9505A7e072A6c6827caF8124f")
var SpaceAddress = ethCommon.HexToAddress("0xDB430EA3d781E6746D6963123483F5E5E9bA009A")
var FileAddress = ethCommon.HexToAddress("0x9E300194A11da9Be466368d7af3470B72B3D2F2a")
var ListAddress = ethCommon.HexToAddress("0xfecd702e246469aaCf94a0c94Cc3561Db527Bd67")
var ProveAddress = ethCommon.HexToAddress("0x644D1bEFF706B5161E5D208B1e870519117940f7")
var PDPAddress = ethCommon.HexToAddress("0x3518D099E55512dEf1827c643c528a8ba8234aD4")

func (t *EVM) GetSigner(value *big.Int) (*bind.TransactOpts, error) {
	ec := t.Client.GetEthClient().Client

	nonce, err := ec.PendingNonceAt(context.TODO(), address)
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.From = address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value                 // in wei
	auth.GasLimit = uint64(10_000_000) // in units
	auth.GasPrice = gasPrice

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
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient().Client)
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

// NodeQuery get node info by wallet address
func (t *EVM) NodeQuery(nodeWallet common.Address) (*fs.FsNodeInfo, error) {
	store, err := nodeStore.NewStore(NodeAddress, t.Client.GetEthClient().Client)
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
	return hash[:], nil
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
		NodeAddr:    ethCommon.HexToAddress(nodeAddr),
	}
	update, err := store.NodeUpdate(signer, node)
	if err != nil {
		return nil, err
	}
	hash := update.Hash()
	return hash[:], nil
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
	cancel, err := store.Cancel(signer, t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	hash := cancel.Hash()
	return hash[:], err
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
	hash := cancel.Hash()
	return hash[:], nil
}

func copyUploadOption(opt *fs.UploadOption) fsStore.UploadOption {
	list := make([]fsStore.WhiteList, opt.WhiteList.Num)
	for k, v := range opt.WhiteList.List {
		list[k] = fsStore.WhiteList{
			Addr:         ethCommon.Address(v.Addr),
			BaseHeight:   v.BaseHeight,
			ExpireHeight: v.ExpireHeight,
		}
	}
	o := fsStore.UploadOption{
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
	store, err := fsStore.NewStore(FileAddress, ec)
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
	store, err := fsStore.NewStore(FileAddress, ec)
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
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	fileHash := []byte(fileHashStr)
	opt := fsStore.FileReNewInfo{
		FileHash:   fileHash,
		FromAddr:   t.DefAcc.EthAddress,
		ReNewTimes: renewTimes,
	}
	reNew, err := store.FileReNew(signer, opt)
	if err != nil {
		return nil, err
	}
	hash := reNew.Hash()
	return hash[:], err
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
	primaryNodes, candidateNodes []common.Address, plotInfo *fs.PlotInfo) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}

	primary := make([]ethCommon.Address, len(primaryNodes))
	for k, v := range primaryNodes {
		primary[k] = ethCommon.Address(v)
	}
	candidates := make([]ethCommon.Address, len(candidateNodes))
	for k, v := range candidateNodes {
		candidates[k] = ethCommon.Address(v)
	}
	fileHash := []byte(fileHashStr)
	f := fsStore.FileInfo{
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
		PlotInfo: fsStore.PlotInfo{
			NumberID:   plotInfo.NumericID,
			StartNonce: plotInfo.StartNonce,
			Nonces:     plotInfo.Nonces,
		},
	}

	file, err := store.StoreFile(signer, f)
	if err != nil {
		return nil, err
	}
	hash := file.Hash()
	return hash[:], err
}

func copyFileInfo(info fsStore.FileInfo) fs.FileInfo {
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
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	info, err := store.GetFileInfo(&bind.CallOpts{}, []byte(fileHashStr))
	if err != nil {
		return nil, err
	}
	fileInfo := copyFileInfo(info)
	return &fileInfo, nil
}

func (t *EVM) GetFileInfos(fileHashStrs []string) (*fs.FileInfoList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fsStore.NewStore(FileAddress, ec)
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
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	ownerChange := fsStore.OwnerChange{
		FileHash: []byte(fileHashStr),
		CurOwner: t.DefAcc.EthAddress,
		NewOwner: ethCommon.Address(newOwner),
	}
	owner, err := store.ChangeFileOwner(signer, ownerChange)
	if err != nil {
		return nil, err
	}
	hash := owner.Hash()
	return hash[:], err
}

func (t *EVM) ChangeFilePrivilege(fileHashStr string, newPrivilege uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ec := t.Client.GetEthClient().Client
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	pri := fsStore.PriChange{
		FileHash:  []byte(fileHashStr),
		Privilege: newPrivilege,
	}
	privilege, err := store.ChangeFilePrivilege(signer, pri)
	if err != nil {
		return nil, err
	}
	hash := privilege.Hash()
	return hash[:], nil
}

func (t *EVM) GetFileList(addr common.Address) (*fs.FileList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetFileList(&bind.CallOpts{}, ethCommon.Address(addr))
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
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	files, err := store.GetUnProvePrimaryFiles(&bind.CallOpts{}, ethCommon.Address(addr))
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
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	files, err := store.GetUnProveCandidateFiles(&bind.CallOpts{}, ethCommon.Address(addr))
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
	operate, err := store.WhiteListOperate(signer, param)
	if err != nil {
		return nil, err
	}
	hash := operate.Hash()
	return hash[:], nil
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
	operate, err := store.WhiteListOperate(signer, param)
	if err != nil {
		return nil, err
	}
	hash := operate.Hash()
	return hash[:], err
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
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	signer, err := t.GetSigner(big.NewInt(0))
	if err != nil {
		return nil, err
	}
	file, err := store.DeleteFile(signer, []byte(fileHashStr))
	if err != nil {
		return nil, err
	}
	hash := file.Hash()
	return hash[:], nil
}

func (t *EVM) DeleteFiles(fileHashStrs []string, gasLimit uint64) ([]byte, error) {
	if t.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}

	if gasLimit == 0 {
		gasLimit = sdkcom.GAS_LIMIT
	}

	ec := t.Client.GetEthClient().Client
	store, err := fsStore.NewStore(FileAddress, ec)
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
	file, err := store.DeleteFiles(signer, p)
	if err != nil {
		return nil, err
	}
	hash := file.Hash()
	return hash[:], nil
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
	prove, err := store.FileProve(signer, param)
	if err != nil {
		return nil, err
	}
	hash := prove.Hash()
	confirmed, err := t.PollForTxConfirmed(t.PollForTxDuration, hash[:])
	if err != nil {
		return nil, err
	}
	if !confirmed {
		return nil, errors.New("tx failed")
	}
	return hash[:], nil
}

func (t *EVM) GenChallenge(walletAddr common.Address, hash common.Uint256, fileBlockNum, proveNum uint64) []pdp.Challenge {
	ec := t.Client.GetEthClient().Client
	store, err := pdpStore.NewStore(PDPAddress, ec)
	if err != nil {
		return nil
	}
	param := pdpStore.GenChallengeParams{
		WalletAddr:   ethCommon.Address(walletAddr),
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
		WalletAddr: ethCommon.Address(walletAddr),
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
	space, err := store.ManageUserSpace(signer, param)
	if err != nil {
		return nil, err
	}
	hash := space.Hash()
	return hash[:], err
}

// GetUserSpace. get user space with wallet address
func (t *EVM) GetUserSpace(walletAddr common.Address) (*fs.UserSpace, error) {
	ec := t.Client.GetEthClient().Client
	store, err := spaceStore.NewStore(SpaceAddress, ec)
	if err != nil {
		return nil, err
	}
	space, err := store.GetUserSpace(&bind.CallOpts{}, ethCommon.Address(walletAddr))
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
		WalletAddr: ethCommon.Address(walletAddr),
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
	space, err := store.DeleteUserSpace(signer, t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	hash := space.Hash()
	return hash[:], err
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
		NodeAddr:   t.DefAcc.EthAddress,
		SectorID:   sectorId,
		Size:       size,
		ProveLevel: uint8(proveLevel),
		IsPlots:    isPlots,
	}
	createSector, err := store.CreateSector(signer, sector)
	if err != nil {
		return nil, err
	}
	hash := createSector.Hash()
	return hash[:], err
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
	createSector, err := store.DeleteSector(signer, sector)
	if err != nil {
		return nil, err
	}
	hash := createSector.Hash()
	return hash[:], err
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
	createSector, err := store.DeleteFileFromSector(signer, s, f) // TODO why todo?
	if err != nil {
		return nil, err
	}
	hash := createSector.Hash()
	return hash[:], err
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
	ec := t.Client.GetEthClient().Client
	store, err := sectorStore.NewStore(SectorAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetSectorsForNode(&bind.CallOpts{}, ethCommon.Address(addr))
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
	prove, err := store.SectorProve(signer, param)
	if err != nil {
		return nil, err
	}
	hash := prove.Hash()
	return hash[:], err
}

func (t *EVM) GetUnSettledFiles(addr common.Address) (*fs.FileList, error) {
	ec := t.Client.GetEthClient().Client
	store, err := fsStore.NewStore(FileAddress, ec)
	if err != nil {
		return nil, err
	}
	list, err := store.GetUnSettledFileList(&bind.CallOpts{}, ethCommon.Address(addr))
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
	store, err := fsStore.NewStore(FileAddress, ec)
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
	hash := tx.Hash()
	return hash[:], nil
}

func (t *EVM) CheckNodeSectorProveInTime(addr common.Address, sectorId uint64) ([]byte, error) {
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
		NodeAddr: ethCommon.Address(addr),
		SectorId: sectorId,
	}
	inTime, err := store.CheckNodeSectorProvedInTime(signer, param)
	if err != nil {
		return nil, err
	}
	hash := inTime.Hash()
	return hash[:], err
}
