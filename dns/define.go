package dns

import (
	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/smartcontract/service/native/dns"
)

type DnsClient interface {
	SetDefaultAccount(acc *account.Account)
	GetDefaultAccount() *account.Account
	GetClient() *client.ClientMgr

	PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error)
	InitDNS() (common.Uint256, error)
	GetPluginList() (*dns.NameInfoList, error)
	RegisterUrl(url string, rType uint64, name string, desc string, ttl uint64) (common.Uint256, error)
	RegisterHeader(header string, desc string, ttl uint64) (common.Uint256, error)
	Binding(urlType uint64, url string, name string, desc string, ttl uint64) (common.Uint256, error)
	DeleteUrl(url string) (common.Uint256, error)
	DeleteHeader(header string) (common.Uint256, error)
	TransferUrl(url string, toAdder string) (common.Uint256, error)
	TransferHeader(header string, toAdder string) (common.Uint256, error)
	QueryUrl(url string, ownerAddr common.Address) (*dns.NameInfo, error)
	QueryHeader(header string, ownerAddr common.Address) (*dns.HeaderInfo, error)
	UnregisterDNSNode() (common.Uint256, error)
	DNSNodeReg(ip []byte, port []byte, initPos uint64) (common.Uint256, error)
	GetDnsNodeByAddr(wallet common.Address) (*dns.DNSNodeInfo, error)
	GetAllDnsNodes() (map[string]dns.DNSNodeInfo, error)
	ApproveCandidate(pubKey string) (common.Uint256, error)
	RejectDNSCandidate(pubKey string) (common.Uint256, error)
	QuitNode() (common.Uint256, error)
	UpdateNode(ip []byte, port []byte) (common.Uint256, error)
	AddInitPos(addPosAmount uint64) (common.Uint256, error)
	ReduceInitPos(changePosAmount uint64) (common.Uint256, error)
	GetPeerPoolMap() (*dns.PeerPoolMap, error)
	GetPeerPoolItem(pubKey string) (*dns.PeerPoolItem, error)
}
