package fs

import (
	"fmt"
	ethCom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"reflect"
	"testing"
	"time"
)

var address = ethCom.HexToAddress("0x5406d22c2aDd39Fa7bE8Bc0f18a03D3CEfbBc0E8")

func CreateClientMgr() *client.ClientMgr {
	c := &client.ClientMgr{}
	err := c.NewEthClient().SetAddress([]string{url})
	if err != nil {
		log.Error(err)
	}
	time.Sleep(time.Second * 1)
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
			t := &EVM{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			fmt.Println("===", t.Client.GetEthClient().Client)
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
			t := &EVM{
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
			t := &EVM{
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
			t := &EVM{
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
				volume:      1000 * 1000,
				serviceTime: 0,
				nodeAddr:    "0x792e47e160f4ee67c17714df1c92f678640e0e4c",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &EVM{
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
				volume:      1000 * 1000,
				serviceTime: 0,
				nodeAddr:    "0x792e47e160f4ee67c17714df1c92f678640e0e4c",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &EVM{
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

func TestEVM_CreateSector(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	type args struct {
		sectorId   uint64
		proveLevel uint64
		size       uint64
		isPlots    bool
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
				sectorId:   1,
				proveLevel: 1,
				size:       1024 * 1024,
				isPlots:    false,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &EVM{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.CreateSector(tt.args.sectorId, tt.args.proveLevel, tt.args.size, tt.args.isPlots)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CreateSector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CreateSector() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEVM_UpdateUserSpace(t1 *testing.T) {
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	type args struct {
		walletAddr common.Address
		size       *fs.UserSpaceOperation
		blockCount *fs.UserSpaceOperation
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
				size: &fs.UserSpaceOperation{
					Type:  1,
					Value: 1024000,
				},
				blockCount: &fs.UserSpaceOperation{
					Type:  1,
					Value: 172800,
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &EVM{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.UpdateUserSpace(tt.args.walletAddr, tt.args.size, tt.args.blockCount)
			if (err != nil) != tt.wantErr {
				t1.Errorf("UpdateUserSpace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEVM_NodeRegister(t1 *testing.T) {
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
				volume:      1000 * 1000,
				serviceTime: 0,
				nodeAddr:    "",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &EVM{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.NodeRegister(tt.args.volume, tt.args.serviceTime, tt.args.nodeAddr)
			if err != nil {
				t1.Errorf("NodeRegister() error = %v", err)
				return
			}
			t1.Log(got)
		})
	}
}

func TestEVM_NodeCancel(t1 *testing.T) {
	priKey, _ := crypto.HexToECDSA("qrRvuI+938wW3GPfEquSSzOeQDvDtpWnGx3WcJTv1PUgFmce3VC0a2LNlIXQHBJ9")
	address := ethCom.HexToAddress("F168345D34E76118b2280dBcF905DE98e2905e61")
	type fields struct {
		Client            *client.ClientMgr
		DefAcc            *account.Account
		PollForTxDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "general",
			fields: fields{
				Client: CreateClientMgr(),
				DefAcc: &account.Account{
					PrivateKey: priKey,
					PublicKey:  nil,
					Address:    common.Address{},
					EthAddress: address,
					SigScheme:  0,
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &EVM{
				Client:            tt.fields.Client,
				DefAcc:            tt.fields.DefAcc,
				PollForTxDuration: tt.fields.PollForTxDuration,
			}
			got, err := t.NodeCancel()
			if (err != nil) != tt.wantErr {
				t1.Errorf("NodeCancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NodeCancel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
