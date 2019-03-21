package chain

import (
	"github.com/oniio/oniChain-go-sdk/auth"
	"github.com/oniio/oniChain-go-sdk/channel"
	"github.com/oniio/oniChain-go-sdk/client"
	"github.com/oniio/oniChain-go-sdk/dns"
	"github.com/oniio/oniChain-go-sdk/fs"
	cgp "github.com/oniio/oniChain-go-sdk/globalparam"
	"github.com/oniio/oniChain-go-sdk/ong"
	"github.com/oniio/oniChain-go-sdk/ont"
	"github.com/oniio/oniChain-go-sdk/ontid"
	"github.com/oniio/oniChain/account"
)

type NativeContract struct {
	Ont          *ont.Ont
	Ong          *ong.Ong
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
	native.Ont = &ont.Ont{Client: client}
	native.Ong = &ong.Ong{Client: client}
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
