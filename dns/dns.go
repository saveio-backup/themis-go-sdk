package dns

import (
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/smartcontract/service/native/dns"
)

type Dns struct {
	Client DnsClient
}

var _ DnsClient = (*Dns)(nil)

func (d *Dns) NewClient(dnsClient DnsClient) {
	d.Client = dnsClient
}

func (d *Dns) SetDefaultAccount(acc *account.Account) {
	d.Client.SetDefaultAccount(acc)
}

func (d *Dns) GetDefaultAccount() *account.Account {
	return d.Client.GetDefaultAccount()
}

func (d *Dns) GetClient() *client.ClientMgr {
	return d.Client.GetClient()
}

func (d Dns) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	return d.Client.PreExecInvokeNativeContract(method, params)
}

func (d Dns) InitDNS() (common.Uint256, error) {
	return d.Client.InitDNS()
}

func (d Dns) GetPluginList() (*dns.NameInfoList, error) {
	return d.Client.GetPluginList()
}

func (d Dns) RegisterUrl(url string, rType uint64, name string, desc string, ttl uint64) (common.Uint256, error) {
	return d.Client.RegisterUrl(url, rType, name, desc, ttl)
}

func (d Dns) RegisterHeader(header string, desc string, ttl uint64) (common.Uint256, error) {
	return d.Client.RegisterHeader(header, desc, ttl)
}

func (d Dns) Binding(urlType uint64, url string, name string, desc string, ttl uint64) (common.Uint256, error) {
	return d.Client.Binding(urlType, url, name, desc, ttl)
}

func (d Dns) DeleteUrl(url string) (common.Uint256, error) {
	return d.Client.DeleteUrl(url)
}

func (d Dns) DeleteHeader(header string) (common.Uint256, error) {
	return d.Client.DeleteHeader(header)
}

func (d Dns) TransferUrl(url string, toAdder string) (common.Uint256, error) {
	return d.Client.TransferUrl(url, toAdder)
}

func (d Dns) TransferHeader(header string, toAdder string) (common.Uint256, error) {
	return d.Client.TransferHeader(header, toAdder)
}

func (d Dns) QueryUrl(url string, ownerAddr common.Address) (*dns.NameInfo, error) {
	return d.Client.QueryUrl(url, ownerAddr)
}

func (d Dns) QueryHeader(header string, ownerAddr common.Address) (*dns.HeaderInfo, error) {
	return d.Client.QueryHeader(header, ownerAddr)
}

func (d Dns) UnregisterDNSNode() (common.Uint256, error) {
	return d.Client.UnregisterDNSNode()
}

func (d Dns) DNSNodeReg(ip []byte, port []byte, initPos uint64) (common.Uint256, error) {
	return d.Client.DNSNodeReg(ip, port, initPos)
}

func (d Dns) GetDnsNodeByAddr(wallet common.Address) (*dns.DNSNodeInfo, error) {
	return d.Client.GetDnsNodeByAddr(wallet)
}

func (d Dns) GetAllDnsNodes() (map[string]dns.DNSNodeInfo, error) {
	return d.Client.GetAllDnsNodes()
}

func (d Dns) ApproveCandidate(pubKey string) (common.Uint256, error) {
	return d.Client.ApproveCandidate(pubKey)
}

func (d Dns) RejectDNSCandidate(pubKey string) (common.Uint256, error) {
	return d.Client.RejectDNSCandidate(pubKey)
}

func (d Dns) QuitNode() (common.Uint256, error) {
	return d.Client.QuitNode()
}

func (d Dns) UpdateNode(ip []byte, port []byte) (common.Uint256, error) {
	return d.Client.UpdateNode(ip, port)
}

func (d Dns) AddInitPos(addPosAmount uint64) (common.Uint256, error) {
	return d.Client.AddInitPos(addPosAmount)
}

func (d Dns) ReduceInitPos(changePosAmount uint64) (common.Uint256, error) {
	return d.Client.ReduceInitPos(changePosAmount)
}

func (d Dns) GetPeerPoolMap() (*dns.PeerPoolMap, error) {
	return d.Client.GetPeerPoolMap()
}

func (d Dns) GetPeerPoolItem(pubKey string) (*dns.PeerPoolItem, error) {
	return d.Client.GetPeerPoolItem(pubKey)
}
