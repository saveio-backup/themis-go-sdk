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

type EVMContract struct {
	ERC20        *usdt.ERC20
	OntId        *ontid.OntId
	GlobalParams *cgp.GlobalParam
	Auth         *auth.Auth
	Fs           *fs.Fs
	Dns          *dns.Dns
	Channel      *channel.Channel
	Governance   *governance.Governance
}

func newEVMContract(client *client.ClientMgr) *EVMContract {
	evm := &EVMContract{}
	evm.ERC20 = &usdt.ERC20{Client: client}
	evm.OntId = &ontid.OntId{Client: client}
	evm.GlobalParams = &cgp.GlobalParam{Client: client}
	evm.Auth = &auth.Auth{Client: client}
	evm.Dns = &dns.Dns{}
	evm.Dns.NewClient(&dns.EVM{
		Client: client,
		DefAcc: nil,
	})
	evm.Fs = &fs.Fs{}
	evm.Fs.NewClient(&fs.EVM{
		Client:            client,
		DefAcc:            nil,
		PollForTxDuration: DEFAULT_POLL_FOR_CONFIRM_DURATION,
	})
	evm.Channel = &channel.Channel{Client: client}
	evm.Governance = &governance.Governance{Client: client}
	return evm
}

func (this *EVMContract) SetDefaultAccount(acc *account.Account) {
	this.Channel.DefAcc = acc
	this.Fs.SetDefaultAccount(acc)
	this.Dns.SetDefaultAccount(acc)
	this.Governance.DefAcc = acc
}
