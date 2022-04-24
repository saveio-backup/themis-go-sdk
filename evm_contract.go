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
	Fs         *fs.Fs
	Dns        *dns.Dns
	Channel    *channel.Channel
	Governance *governance.Governance
}

func newEVMContract(client *client.ClientMgr) *EVMContract {
	evm := &EVMContract{}
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
	this.Fs.Client.SetDefaultAccount(acc)
	this.Dns.DefAcc = acc
	this.Governance.DefAcc = acc
}
