package chain

import (
	"time"

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

const DEFAULT_POLL_FOR_CONFIRM_DURATION = time.Duration(10) * time.Second

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
	native.Fs = &fs.Fs{}
	native.Fs.NewClient(&fs.Themis{
		Client:            client,
		DefAcc:            nil,
		PollForTxDuration: DEFAULT_POLL_FOR_CONFIRM_DURATION,
	})
	native.Channel = &channel.Channel{Client: client}
	native.Governance = &governance.Governance{Client: client}
	return native
}

func newEthereumContract(client *client.ClientMgr) *NativeContract {
	native := &NativeContract{}
	native.Usdt = &usdt.Usdt{Client: client}
	native.OntId = &ontid.OntId{Client: client}
	native.GlobalParams = &cgp.GlobalParam{Client: client}
	native.Auth = &auth.Auth{Client: client}
	native.Dns = &dns.Dns{Client: client}
	native.Fs = &fs.Fs{}
	native.Fs.NewClient(&fs.Ethereum{
		Client: 		   client,
		DefAcc:            nil,
		PollForTxDuration: DEFAULT_POLL_FOR_CONFIRM_DURATION,
	})
	native.Channel = &channel.Channel{Client: client}
	native.Governance = &governance.Governance{Client: client}
	return native
}

func (this *NativeContract) SetDefaultAccount(acc *account.Account) {
	this.Channel.DefAcc = acc
	this.Fs.Client.SetDefaultAccount(acc)
	this.Dns.DefAcc = acc
	this.Governance.DefAcc = acc
}
