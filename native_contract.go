package chain

import (
	"github.com/saveio/themis-go-sdk/auth"
	"github.com/saveio/themis-go-sdk/channel"
	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/dns"
	"github.com/saveio/themis-go-sdk/fs"
	cgp "github.com/saveio/themis-go-sdk/globalparam"
	"github.com/saveio/themis-go-sdk/governance"
	"github.com/saveio/themis-go-sdk/ontid"
	"github.com/saveio/themis-go-sdk/usdt"
	"github.com/saveio/themis/account"
)

type NativeContract struct {
	Usdt         *usdt.Usdt
	OntId        *ontid.OntId
	GlobalParams *cgp.GlobalParam
	Auth         *auth.Auth
	Dns          *dns.Dns
	Fs           *fs.Fs
	Channel      *channel.Channel
	Governance   *governance.Governance
}

func newNativeContract(client *client.ClientMgr) *NativeContract {
	native := &NativeContract{}
	native.Usdt = &usdt.Usdt{Client: client}
	native.OntId = &ontid.OntId{Client: client}
	native.GlobalParams = &cgp.GlobalParam{Client: client}
	native.Auth = &auth.Auth{Client: client}
	native.Dns = &dns.Dns{Client: client}
	native.Fs = &fs.Fs{Client: client}
	native.Channel = &channel.Channel{Client: client}
	native.Governance = &governance.Governance{Client: client}
	return native
}

func (this *NativeContract) SetDefaultAccount(acc *account.Account) {
	this.Channel.DefAcc = acc
	this.Fs.DefAcc = acc
	this.Dns.DefAcc = acc
	this.Governance.DefAcc = acc
}
