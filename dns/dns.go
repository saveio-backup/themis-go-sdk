package dns

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"encoding/json"

	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/crypto/keypair"
	"github.com/saveio/themis/smartcontract/service/native/dns"
)

var (
	DNS_CONTRACT_ADDRESS, _ = utils.AddressFromHexString("1200000000000000000000000000000000000000")
	DNS_CONTRACT_VERSION    = byte(0)
)

type Dns struct {
	Client *client.ClientMgr
	DefAcc *account.Account
}

func (this *Dns) InvokeNativeContract(signer *account.Account, method string, params []interface{}) (common.Uint256, error) {
	if signer == nil {
		return common.UINT256_EMPTY, errors.New("signer is nil")
	}
	tx, err := utils.NewNativeInvokeTransaction(sdkcom.GAS_PRICE, sdkcom.GAS_LIMIT, DNS_CONTRACT_VERSION, DNS_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Dns) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, DNS_CONTRACT_VERSION, DNS_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return nil, err
	}
	return this.Client.PreExecTransaction(tx)
}

// InitDNS. init dns native contract.
func (this *Dns) InitDNS() (common.Uint256, error) {
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.INIT_NAME, []interface{}{})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// RegisterUrl. register url
func (this *Dns) RegisterUrl(url string, rType uint64, name, desc string, ttl uint64) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	var req dns.RequestName
	if rType == dns.SYSTEM {
		req = dns.RequestName{
			Type:      rType,
			Header:    []byte{0},
			URL:       []byte{0},
			Name:      []byte(name),
			NameOwner: this.DefAcc.Address,
			Desc:      []byte(desc),
			DesireTTL: ttl,
		}
	} else {
		strs := strings.Split(url, "://")
		if len(strs) != 2 {
			return common.UINT256_EMPTY, errors.New("QueryUrl input url format invalid")
		}
		req = dns.RequestName{
			Type:      rType,
			Header:    []byte(strs[0]),
			URL:       []byte(strs[1]),
			Name:      []byte(name),
			NameOwner: this.DefAcc.Address,
			Desc:      []byte(desc),
			DesireTTL: ttl,
		}
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.REGISTER_NAME, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// RegisterHeader. register url header prefix
func (this *Dns) RegisterHeader(header, desc string, ttl uint64) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	req := dns.RequestHeader{
		Header:    []byte(header),
		NameOwner: this.DefAcc.Address,
		Desc:      []byte(desc),
		DesireTTL: ttl,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.REGISTER_HEADER, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// Binding. bind url with name, desc, ttl
func (this *Dns) Binding(url string, name, desc string, ttl uint64) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return common.UINT256_EMPTY, errors.New("QueryUrl input url format valid")
	}
	req := dns.RequestName{
		Header:    []byte(strs[0]),
		URL:       []byte(strs[1]),
		Name:      []byte(name),
		NameOwner: this.DefAcc.Address,
		Desc:      []byte(desc),
		DesireTTL: ttl,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.UPDATE_DNS_NAME, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// DeleteUrl. delete url
func (this *Dns) DeleteUrl(url string) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return common.UINT256_EMPTY, errors.New("QueryUrl input url format valid")
	}
	req := dns.ReqInfo{
		Header: []byte(strs[0]),
		URL:    []byte(strs[1]),
		Owner:  this.DefAcc.Address,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.DEL_DNS, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// DeleteHeader. delete header
func (this *Dns) DeleteHeader(header string) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	req := dns.ReqInfo{
		Header: []byte(header),
		Owner:  this.DefAcc.Address,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.DEL_DNS_HEADER, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// TransferUrl. transfer the ownership of url to other user.
func (this *Dns) TransferUrl(url string, toAdder string) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return common.UINT256_EMPTY, errors.New("QueryUrl input url format valid")
	}
	to, err := common.AddressFromHexString(toAdder)
	req := dns.TranferInfo{
		Header: []byte(strs[0]),
		URL:    []byte(strs[1]),
		From:   this.DefAcc.Address,
		To:     to,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.TRANSFER_NAME, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// TransferHeader. transfer the ownership of header to other user.
func (this *Dns) TransferHeader(header, toAdder string) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	to, err := common.AddressFromBase58(toAdder)
	req := dns.TranferInfo{
		Header: []byte(header),
		From:   this.DefAcc.Address,
		To:     to,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.TRANSFER_HEADER_NAME, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, err
}

// QueryUrl. query a url to find it's exist or not
func (this *Dns) QueryUrl(url string, ownerAddr common.Address) (*dns.NameInfo, error) {
	strs := strings.Split(url, "://")
	if len(strs) != 2 {
		return nil, errors.New("QueryUrl input url format valid")
	}
	req := dns.ReqInfo{
		Header: []byte(strs[0]),
		URL:    []byte(strs[1]),
		Owner:  ownerAddr,
	}
	ret, err := this.PreExecInvokeNativeContract(
		dns.GET_DNS_NAME, []interface{}{req},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, errors.New("QueryUrl result toByteArray err")
	}

	var name dns.NameInfo
	infoReader := bytes.NewReader(data)
	err = name.Deserialize(infoReader)
	if err != nil {
		return nil, errors.New("NameInfo deserialize error")
	}
	return &name, nil
}

// QueryHeader. query a header to find it's exist or not
func (this *Dns) QueryHeader(header string, ownerAddr common.Address) (*dns.HeaderInfo, error) {
	req := dns.ReqInfo{
		Header: []byte(header),
		Owner:  ownerAddr,
	}

	ret, err := this.PreExecInvokeNativeContract(
		dns.GET_HEADER_NAME, []interface{}{req},
	)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, errors.New("QueryHeader result toByteArray err")
	}

	var headerInfo dns.HeaderInfo
	infoReader := bytes.NewReader(data)
	err = headerInfo.Deserialize(infoReader)
	if err != nil {
		return nil, err
	}
	return &headerInfo, nil

}

// UnregisterDNSNode. unregister dns node to quit
func (this *Dns) UnregisterDNSNode() (common.Uint256, error) {
	pk := keypair.SerializePublicKey(this.DefAcc.PublicKey)
	param := dns.UnRegisterCandidateParam{
		PeerPubkey: hex.EncodeToString(pk),
		Address:    this.DefAcc.Address,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.UN_DNS_NODE_REG, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// DNSNodeReg. register dns node to dns node govern contract
func (this *Dns) DNSNodeReg(ip, port []byte, initPos uint64) (common.Uint256, error) {
	pk := keypair.SerializePublicKey(this.DefAcc.PublicKey)
	peerPubKey := hex.EncodeToString(pk)
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	req := dns.DNSNodeInfo{
		WalletAddr:  this.DefAcc.Address,
		IP:          ip,
		Port:        port,
		InitDeposit: initPos,
		PeerPubKey:  peerPubKey,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.DNS_NODE_REG, []interface{}{req})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

func (this *Dns) GetDnsNodeByAddr(wallet common.Address) (*dns.DNSNodeInfo, error) {
	var req common.Address
	if wallet == common.ADDRESS_EMPTY {
		req = this.DefAcc.Address
	} else {
		req = wallet
	}
	ret, err := this.PreExecInvokeNativeContract(dns.GET_DNSNODE_BYADDR, []interface{}{req})
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, errors.New("QueryDnsNode result toByteArray error")
	}
	var dn dns.DNSNodeInfo
	reader := bytes.NewReader(data)
	err = dn.Deserialize(reader)
	if err != nil {
		return nil, errors.New("QueryDnsNode dnsNodeInfo deserialize error")
	}
	return &dn, nil
}

func (this *Dns) GetAllDnsNodes() (map[string]dns.DNSNodeInfo, error) {
	dn := make(map[string]dns.DNSNodeInfo)
	ret, err := this.PreExecInvokeNativeContract(dns.GET_ALL_DNSNODES, nil)

	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, errors.New("GetDnsNodes result toByteArray error")
	}
	err = json.Unmarshal(data, &dn)
	if err != nil {
		return nil, errors.New("QueryDnsNode dnsNodeInfo deserialize error")
	}
	return dn, nil
}

func (this *Dns) GetPluginList() (*dns.NameInfoList, error) {
	ret, err := this.PreExecInvokeNativeContract(dns.GET_PLUGIN_LIST, nil)
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, fmt.Errorf("GetPluginList result toByteArray error: %v", err.Error())
	}

	dnList := &dns.NameInfoList{}
	retInfo := dns.DecRet(data)
	if retInfo.Ret {
		reader := bytes.NewReader(retInfo.Info)
		err = dnList.Deserialize(reader)
		if err != nil {
			return nil, fmt.Errorf("GetPluginList dns.NameInfoList deserialize error: %v", err)
		}
		return dnList, nil
	} else {
		return nil, errors.New(string(retInfo.Info))
	}
}

// ApproveCandidate. approve a dns candidate node.
func (this *Dns) ApproveCandidate(pubKey string) (common.Uint256, error) {
	param := dns.ApproveCandidateParam{
		PeerPubkey: pubKey,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.APPROVE_CANDIDATE, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// RejectDNSCandidate. reject a dns candidate node.
func (this *Dns) RejectDNSCandidate(pubKey string) (common.Uint256, error) {
	param := dns.PubKeyParam{
		PeerPubkey: pubKey,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.REJECT_CANDIDATE, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// QuitNode. quit node.
func (this *Dns) QuitNode() (common.Uint256, error) {
	pk := keypair.SerializePublicKey(this.DefAcc.PublicKey)
	param := dns.QuitNodeParam{
		PeerPubkey: hex.EncodeToString(pk),
		Address:    this.DefAcc.Address,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.QUIT_NODE, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// UpdateNode. update dns node host info.
func (this *Dns) UpdateNode(ip, port []byte) (common.Uint256, error) {
	if this.DefAcc == nil {
		return common.UINT256_EMPTY, errors.New("account is nil")
	}
	param := dns.UpdateNodeParam{
		WalletAddr: this.DefAcc.Address,
		IP:         ip,
		Port:       port,
	}
	log.Infof("UpdateNodeParam: %v", param)
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.UPDATE_DNSNODE, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// AddInitPos. add init pos.
func (this *Dns) AddInitPos(addPosAmount uint64) (common.Uint256, error) {
	pk := keypair.SerializePublicKey(this.DefAcc.PublicKey)
	param := dns.ChangeInitPosParam{
		PeerPubkey: hex.EncodeToString(pk),
		Address:    this.DefAcc.Address,
		Pos:        addPosAmount,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.ADD_INIT_POS, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// ReduceInitPos. reduce init pos.
func (this *Dns) ReduceInitPos(changePosAmount uint64) (common.Uint256, error) {
	pk := keypair.SerializePublicKey(this.DefAcc.PublicKey)
	param := dns.ChangeInitPosParam{
		PeerPubkey: hex.EncodeToString(pk),
		Address:    this.DefAcc.Address,
		Pos:        changePosAmount,
	}
	ret, err := this.InvokeNativeContract(this.DefAcc, dns.REDUCE_INIT_POS, []interface{}{param})
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return ret, nil
}

// GetPeerPoolMap. get peer pool map
func (this *Dns) GetPeerPoolMap() (*dns.PeerPoolMap, error) {
	ret, err := this.PreExecInvokeNativeContract(dns.GET_PEER_POOLMAP, []interface{}{})
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, errors.New("GetPeerPoolMap result toByteArray err")
	}
	if len(data) == 0 {
		return nil, errors.New("data is nil")
	}
	var peerPoolMap dns.PeerPoolMap
	err = peerPoolMap.Deserialize(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("peerPoolMap deserialize error")
	}
	return &peerPoolMap, nil
}

// GetPeerPoolItem. get peer pool item
func (this *Dns) GetPeerPoolItem(pubKey string) (*dns.PeerPoolItem, error) {
	if pubKey == "" {
		pk := keypair.SerializePublicKey(this.DefAcc.PublicKey)
		pubKey = hex.EncodeToString(pk)
	}
	param := dns.PubKeyParam{
		PeerPubkey: pubKey,
	}
	ret, err := this.PreExecInvokeNativeContract(dns.GET_PEER_POOLITEM, []interface{}{param})
	if err != nil {
		return nil, err
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, errors.New("GetPeerPoolItem result toByteArray err")
	}
	if len(data) == 0 {
		return nil, errors.New("data is nil")
	}
	var peerPoolItem dns.PeerPoolItem
	err = peerPoolItem.Deserialize(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("peerPoolItem deserialize error")
	}
	return &peerPoolItem, nil
}
