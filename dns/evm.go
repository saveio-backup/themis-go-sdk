package dns

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/saveio/themis/crypto/keypair"
	"math/big"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	store "github.com/saveio/themis-go-sdk/dns/contracts/Dns"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/smartcontract/service/native/dns"
)

type EVM struct {
	Client *client.ClientMgr
	DefAcc *account.Account
}

var _ DnsClient = (*EVM)(nil)

var DnsAddress = ethCommon.HexToAddress("0xE85aA093b9bB9E41D3c31f98AE5220D74DA9573B")

func (E *EVM) GetSigner(value *big.Int) (*bind.TransactOpts, error) {
	ec := E.Client.GetEthClient().Client
	nonce, err := ec.PendingNonceAt(context.TODO(), E.DefAcc.EthAddress)
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
	keyBytes := E.DefAcc.GetEthPrivateKey()
	keyStr := hexutil.Encode(keyBytes)
	key, err := crypto.HexToECDSA(keyStr[2:])
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	auth.From = E.DefAcc.EthAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value                 // in wei
	auth.GasLimit = uint64(10_000_000) // in units
	auth.GasPrice = gasPrice

	gas := new(big.Int).Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit)))
	gas = new(big.Int).Add(auth.Value, gas)
	log.Debugf("get signer with gas: %v, %v, %v, %v", auth.GasPrice, auth.GasLimit, auth.Value, gas)
	return auth, nil
}

func GetStore(E *EVM) (*ethclient.Client, *store.Store, error) {
	ec := E.Client.GetEthClient().Client
	if ec == nil {
		return nil, nil, errors.New("eth client is nil")
	}
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return nil, nil, err
	}
	return ec, newStore, nil
}

func (E *EVM) SetDefaultAccount(acc *account.Account) {
	E.DefAcc = acc
}

func (E *EVM) GetDefaultAccount() *account.Account {
	return E.DefAcc
}

func (E *EVM) GetClient() *client.ClientMgr {
	return E.Client
}

func (E *EVM) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	//TODO implement me
	panic("implement me")
}

func (E *EVM) InitDNS() (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E *EVM) GetPluginList() (*dns.NameInfoList, error) {
	//TODO implement me
	panic("implement me")
}

func (E *EVM) RegisterUrl(url string, rType uint64, name string, desc string, ttl uint64) (common.Uint256, error) {
	ec := E.Client.GetEthClient().Client
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return common.Uint256{}, err
	}
	auth, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	var req store.RequestName
	if rType == dns.SYSTEM {
		req = store.RequestName{
			Type:      rType,
			Header:    []byte{0},
			URL:       []byte{0},
			Name:      []byte(name),
			NameOwner: E.DefAcc.EthAddress,
			Desc:      []byte(desc),
			DesireTTL: ttl,
		}
	} else {
		strs := strings.Split(url, "://")
		if len(strs) != 2 {
			return common.UINT256_EMPTY, errors.New("QueryUrl input url format invalid")
		}
		req = store.RequestName{
			Type:      rType,
			Header:    []byte(strs[0]),
			URL:       []byte(strs[1]),
			Name:      []byte(name),
			NameOwner: E.DefAcc.EthAddress,
			Desc:      []byte(desc),
			DesireTTL: ttl,
		}
	}
	tx, err := newStore.RegisterName(auth, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) RegisterHeader(header string, desc string, ttl uint64) (common.Uint256, error) {
	ec := E.Client.GetEthClient().Client
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return common.Uint256{}, err
	}
	auth, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	addr := ethCommon.BytesToAddress(E.DefAcc.Address[:])
	req := store.RequestHeader{
		Header:    []byte(header),
		NameOwner: addr,
		Desc:      []byte(desc),
		DesireTTL: ttl,
	}
	tx, err := newStore.RegisterHeader(auth, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResult(ec, tx)
}

func (E *EVM) Binding(urlType uint64, url string, name string, desc string, ttl uint64) (common.Uint256, error) {
	ec := E.Client.GetEthClient().Client
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return common.UINT256_EMPTY, errors.New("QueryUrl input url format valid")
	}
	req := store.RequestName{
		Type:      urlType,
		Header:    []byte(strs[0]),
		URL:       []byte(strs[1]),
		Name:      []byte(name),
		NameOwner: E.DefAcc.EthAddress,
		Desc:      []byte(desc),
		DesireTTL: ttl,
	}
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return common.Uint256{}, err
	}
	auth, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	tx, err := newStore.UpdateName(auth, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) DeleteUrl(url string) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return common.UINT256_EMPTY, errors.New("QueryUrl input url format valid")
	}
	req := store.ReqInfo{
		Header: []byte(strs[0]),
		URL:    []byte(strs[1]),
		Owner:  E.DefAcc.EthAddress,
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	tx, err := getStore.DelDNS(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) DeleteHeader(header string) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	req := store.ReqInfo{
		Header: []byte(header),
		Owner:  E.DefAcc.EthAddress,
	}
	tx, err := getStore.DelHeader(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) TransferUrl(url string, toAdder string) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return common.UINT256_EMPTY, errors.New("QueryUrl input url format valid")
	}
	to, err := common.AddressFromHexString(toAdder)
	if err != nil {
		return common.Uint256{}, err
	}
	req := store.TransferInfo{
		Header: []byte(strs[0]),
		URL:    []byte(strs[1]),
		From:   E.DefAcc.EthAddress,
		To:     ethCommon.BytesToAddress(to[:]),
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	tx, err := getStore.TransferName(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) TransferHeader(header string, toAdder string) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	to, err := common.AddressFromHexString(toAdder)
	if err != nil {
		return common.Uint256{}, err
	}
	req := store.TransferInfo{
		Header: []byte(header),
		From:   E.DefAcc.EthAddress,
		To:     ethCommon.BytesToAddress(to[:]),
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	tx, err := getStore.TransferHeader(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) QueryUrl(url string, ownerAddr common.Address) (*dns.NameInfo, error) {
	ec := E.Client.GetEthClient().Client
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return nil, err
	}
	addr := ethCommon.BytesToAddress(ownerAddr[:])
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return nil, err
	}
	req := store.ReqInfo{
		Header: []byte(strs[0]),
		URL:    []byte(strs[1]),
		Owner:  addr,
	}
	res, err := newStore.GetName(nil, req)
	if err != nil {
		return nil, err
	}
	if res.BlockHeight.Uint64() == 0 {
		return nil, errors.New("not found")
	}
	return &dns.NameInfo{
		Type:        res.Type,
		Header:      res.Header,
		URL:         res.URL,
		Name:        res.Name,
		NameOwner:   common.Address(res.NameOwner),
		Desc:        res.Desc,
		BlockHeight: res.BlockHeight.Uint64(),
		TTL:         res.TTL,
	}, nil
}

func (E *EVM) QueryHeader(header string, ownerAddr common.Address) (*dns.HeaderInfo, error) {
	ec := E.Client.GetEthClient().Client
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return nil, err
	}
	addr := ethCommon.BytesToAddress(ownerAddr[:])
	req := store.ReqInfo{
		Header: []byte(header),
		Owner:  addr,
	}
	res, err := newStore.GetHeader(nil, req)
	if err != nil {
		return nil, err
	}
	if res.BlockHeight.Uint64() == 0 {
		return nil, errors.New("not found")
	}
	return &dns.HeaderInfo{
		Header:      res.Header,
		HeaderOwner: common.Address(res.HeaderOwner),
		BlockHeight: res.BlockHeight.Uint64(),
	}, nil
}

func (E *EVM) UnregisterDNSNode() (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	pk := keypair.SerializePublicKey(E.DefAcc.PublicKey)
	req := store.UnRegisterCandidateParam{
		PeerPubKey: hex.EncodeToString(pk),
		Address:    E.DefAcc.EthAddress,
	}
	tx, err := getStore.UnRegDNSNode(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) DNSNodeReg(ip []byte, port []byte, initPos uint64) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	pk := keypair.SerializePublicKey(E.DefAcc.PublicKey)
	req := store.DNSNodeInfo{
		PeerPubKey:  hex.EncodeToString(pk),
		WalletAddr:  E.DefAcc.EthAddress,
		IP:          ip,
		Port:        port,
		InitDeposit: initPos,
	}
	tx, err := getStore.DNSNodeReg(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) GetDnsNodeByAddr(wallet common.Address) (*dns.DNSNodeInfo, error) {
	ec := E.Client.GetEthClient().Client
	if ec == nil {
		return nil, errors.New("eth client is nil")
	}
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return nil, err
	}
	info, err := newStore.GetDNSNodeByAddress(nil, ethCommon.BytesToAddress(wallet[:]))
	res := &dns.DNSNodeInfo{
		WalletAddr:  common.Address(info.WalletAddr),
		IP:          info.IP,
		Port:        info.Port,
		InitDeposit: info.InitDeposit,
		PeerPubKey:  info.PeerPubKey,
	}
	return res, nil
}

func (E *EVM) GetAllDnsNodes() (map[string]dns.DNSNodeInfo, error) {
	ec := E.Client.GetEthClient().Client
	if ec == nil {
		return nil, errors.New("eth client is nil")
	}
	newStore, err := store.NewStore(DnsAddress, ec)
	if err != nil {
		return nil, err
	}
	nodes, err := newStore.GetAllDnsNodes(nil)
	if err != nil {
		return nil, err
	}
	res := make(map[string]dns.DNSNodeInfo)
	for _, v := range nodes {
		res[v.PeerPubKey] = dns.DNSNodeInfo{
			WalletAddr:  common.Address(v.WalletAddr),
			PeerPubKey:  v.PeerPubKey,
			IP:          v.IP,
			Port:        v.Port,
			InitDeposit: v.InitDeposit,
		}
	}
	return res, nil
}

func (E *EVM) ApproveCandidate(pubKey string) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	tx, err := getStore.ApproveDNSCandidate(signer, pubKey)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) RejectDNSCandidate(pubKey string) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}
	tx, err := getStore.RejectDNSCandidate(signer, pubKey)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) QuitNode() (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}

	pk := keypair.SerializePublicKey(E.DefAcc.PublicKey)
	req := store.QuitNodeParam{
		PeerPubKey: hex.EncodeToString(pk),
		Address:    E.DefAcc.EthAddress,
	}
	tx, err := getStore.QuitNode(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) UpdateNode(ip []byte, port []byte) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}

	req := store.UpdateNodeParam{
		WalletAddr: E.DefAcc.EthAddress,
		IP:         ip,
		Port:       port,
	}
	tx, err := getStore.UpdateDNSNodesInfo(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) AddInitPos(addPosAmount uint64) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}

	pk := keypair.SerializePublicKey(E.DefAcc.PublicKey)
	req := store.ChangeInitPosParam{
		PeerPubKey: hex.EncodeToString(pk),
		Address:    E.DefAcc.EthAddress,
		Pos:        addPosAmount,
	}
	tx, err := getStore.AddInitPos(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) ReduceInitPos(changePosAmount uint64) (common.Uint256, error) {
	if E.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	ec, getStore, err := GetStore(E)
	if err != nil {
		return common.Uint256{}, err
	}
	signer, err := E.GetSigner(big.NewInt(0))
	if err != nil {
		return common.Uint256{}, err
	}

	pk := keypair.SerializePublicKey(E.DefAcc.PublicKey)
	req := store.ChangeInitPosParam{
		PeerPubKey: hex.EncodeToString(pk),
		Address:    E.DefAcc.EthAddress,
		Pos:        changePosAmount,
	}
	tx, err := getStore.ReduceInitPos(signer, req)
	if err != nil {
		return common.Uint256{}, err
	}
	return TxResultWithError(ec, tx)
}

func (E *EVM) GetPeerPoolMap() (*dns.PeerPoolMap, error) {
	_, getStore, err := GetStore(E)
	if err != nil {
		return nil, err
	}
	poolMap, err := getStore.GetPeerPoolMap(nil)
	if err != nil {
		return nil, err
	}
	m := make(map[string]*dns.PeerPoolItem)
	for _, v := range poolMap {
		m[v.PeerPubKey] = &dns.PeerPoolItem{
			PeerPubkey:    v.PeerPubKey,
			WalletAddress: common.Address(v.WalletAddress),
			TotalInitPos:  v.TotalInitPos,
			Status:        dns.Status(v.Status),
		}
	}
	res := &dns.PeerPoolMap{
		PeerPoolMap: m,
	}
	return res, nil
}

func (E *EVM) GetPeerPoolItem(pubKey string) (*dns.PeerPoolItem, error) {
	_, getStore, err := GetStore(E)
	if err != nil {
		return nil, err
	}
	item, err := getStore.GetPeerPoolItem(nil, pubKey)
	if err != nil {
		return nil, err
	}
	return &dns.PeerPoolItem{
		PeerPubkey:    item.PeerPubKey,
		WalletAddress: common.Address(item.WalletAddress),
		TotalInitPos:  item.TotalInitPos,
		Status:        dns.Status(item.Status),
	}, nil
}

func TxResult(ec *ethclient.Client, tx *types.Transaction) (common.Uint256, error) {
	mined, err := bind.WaitMined(context.Background(), ec, tx)
	if err != nil {
		return common.Uint256{}, err
	}
	if mined.Status != types.ReceiptStatusSuccessful {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		var funcName string
		if ok && details != nil {
			funcName = details.Name()
		}
		log.Errorf("tx failed: %s, func: %s, logs: %d", tx.Hash().String(), funcName, len(mined.Logs))
		return common.Uint256{}, errors.New("tx failed")
	}
	return common.Uint256(tx.Hash()), nil
}

func TxResultWithError(ec *ethclient.Client, tx *types.Transaction) (common.Uint256, error) {
	mined, err := bind.WaitMined(context.Background(), ec, tx)
	if err != nil {
		return common.Uint256{}, err
	}
	if mined.Status != types.ReceiptStatusSuccessful {
		return common.Uint256{}, errors.New("tx mined failed")
	}
	var errMsg string
	for _, v := range mined.Logs {
		contractAbi, err := abi.JSON(strings.NewReader(store.StoreMetaData.ABI))
		if err != nil {
			return common.Uint256{}, err
		}
		data, err := contractAbi.Unpack("DnsError", v.Data)
		for _, v := range data {
			errMsg += v.(string)
		}
	}
	if errMsg != "" {
		return common.Uint256(tx.Hash()), errors.New(errMsg)
	}
	return common.Uint256(tx.Hash()), nil
}
