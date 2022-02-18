package fs

import (
	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"testing"
	"time"
)

func CreateClientMgr() *client.ClientMgr {
	c := &client.ClientMgr{}
	err := c.NewEthClient(url)
	if err != nil {
		log.Error(err)
	}
	return c
}

func TestEthereum_GetSetting(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    *fs.FsSetting
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ethereum{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.GetSetting()
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEthereum_GetNodeList(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    *fs.FsNodesInfo
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ethereum{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.GetNodeList()
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetNodeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEthereum_GetNodeListByAddrs(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	type args struct {
		addrs []common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *fs.FsNodesInfo
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
			},
			args: args{addrs: make([]common.Address, 1)},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ethereum{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.GetNodeListByAddrs(tt.args.addrs)
			if err != nil {
				t1.Errorf("GetNodeListByAddrs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEthereum_NodeQuery(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	type args struct {
		nodeWallet common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *fs.FsNodeInfo
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
			},
			args: args{nodeWallet: common.Address{}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ethereum{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.NodeQuery(tt.args.nodeWallet)
			if err != nil {
				t1.Errorf("NodeQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEthereum_NodeRegister(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	type args struct {
		volume      uint64
		serviceTime uint64
		nodeAddr    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
				DefAcc: &account.Account{
					PrivateKey: nil,
					PublicKey:  nil,
					Address:    common.Address{},
					EthAddress: address,
					SigScheme:  0,
				},
			},
			args: args{
				volume: 1000 * 1000,
				serviceTime: 0,
				nodeAddr: "0x792e47e160f4ee67c17714df1c92f678640e0e4c",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ethereum{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.NodeRegister(tt.args.volume, tt.args.serviceTime, tt.args.nodeAddr)
			if err != nil {
				t1.Errorf("NodeRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEthereum_NodeUpdate(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	type args struct {
		volume      uint64
		serviceTime uint64
		nodeAddr    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
				DefAcc: &account.Account{
					PrivateKey: nil,
					PublicKey:  nil,
					Address:    common.Address{},
					EthAddress: address,
					SigScheme:  0,
				},
			},
			args: args{
				volume: 1000 * 1000,
				serviceTime: 0,
				nodeAddr: "0x792e47e160f4ee67c17714df1c92f678640e0e4c",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Ethereum{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.NodeUpdate(tt.args.volume, tt.args.serviceTime, tt.args.nodeAddr)
			if err != nil {
				t1.Errorf("NodeUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}