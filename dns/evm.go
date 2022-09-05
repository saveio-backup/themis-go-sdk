package dns

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/smartcontract/service/native/dns"
	"math/big"
)

type EVM struct {
	Client *client.ClientMgr
	DefAcc *account.Account
}

var _ DnsClient = (*EVM)(nil)

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

func (E EVM) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) InitDNS() (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) GetPluginList() (*dns.NameInfoList, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) RegisterUrl(url string, rType uint64, name string, desc string, ttl uint64) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) RegisterHeader(header string, desc string, ttl uint64) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) Binding(urlType uint64, url string, name string, desc string, ttl uint64) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) DeleteUrl(url string) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) DeleteHeader(header string) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) TransferUrl(url string, toAdder string) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) TransferHeader(header string, toAdder string) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) QueryUrl(url string, ownerAddr common.Address) (*dns.NameInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) QueryHeader(header string, ownerAddr common.Address) (*dns.HeaderInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) UnregisterDNSNode() (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) DNSNodeReg(ip []byte, port []byte, initPos uint64) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) GetDnsNodeByAddr(wallet common.Address) (*dns.DNSNodeInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) GetAllDnsNodes() (map[string]dns.DNSNodeInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) ApproveCandidate(pubKey string) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) RejectDNSCandidate(pubKey string) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) QuitNode() (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) UpdateNode(ip []byte, port []byte) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) AddInitPos(addPosAmount uint64) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) ReduceInitPos(changePosAmount uint64) (common.Uint256, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) GetPeerPoolMap() (*dns.PeerPoolMap, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVM) GetPeerPoolItem(pubKey string) (*dns.PeerPoolItem, error) {
	//TODO implement me
	panic("implement me")
}
