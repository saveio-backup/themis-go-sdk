package chain

import (
	"github.com/saveio/themis-go-sdk/channel"
	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/dns"
	"github.com/saveio/themis-go-sdk/fs"
	"github.com/saveio/themis-go-sdk/governance"
	"github.com/saveio/themis/account"
)

type EVMContract struct {
	Fs           *fs.Fs
	Dns          *dns.Dns
	Channel      *channel.Channel
	Governance   *governance.Governance
}

func newEthereumContract(client *client.ClientMgr) *EVMContract {
	native := &EVMContract{}
	native.Fs = &fs.Fs{}
	native.Fs.NewClient(&fs.Ethereum{
		Client:            client,
		DefAcc:            nil,
		PollForTxDuration: DEFAULT_POLL_FOR_CONFIRM_DURATION,
	})
	native.Channel = &channel.Channel{Client: client}
	native.Governance = &governance.Governance{Client: client}
	return native
}

func (this *EVMContract) SetDefaultAccount(acc *account.Account) {
	this.Channel.DefAcc = acc
	this.Fs.Client.SetDefaultAccount(acc)
	this.Dns.DefAcc = acc
	this.Governance.DefAcc = acc
}
