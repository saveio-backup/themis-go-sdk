package fs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis/account"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/saveio/themis-go-sdk/fs/contracts/Config"
)

var url = "http://104.225.158.242:8545"
var address = common.HexToAddress("0xa934808a26bd08c5145cf1894d06176d3664f567")
var ConfigAddress = common.HexToAddress("0x6b450d2b53Bd6C2d866F5630eDc3bB61e8016A91")

var ec *ethclient.Client
var privateKey *ecdsa.PrivateKey

func Test_Init(t *testing.T) {
	var err error
	// eth client
	ec, err = ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// private key
	privateKey, err = crypto.HexToECDSA("c9cc08e8564d21be633f55558e9186efcbcd5f2950bee49c9faafb612ccf53fa")
	if err != nil {
		log.Fatal(err)
	}
}

// -----------------------------------------------------

func Test_Deploy_Config(t *testing.T) {
	Test_Init(t)

	nonce, err := ec.PendingNonceAt(context.Background(), address)
	if err != nil {
		t.Fatal(err)
	}

	gasPrice, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}
	auth.From = address
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(1000000)    // in units
	auth.GasPrice = gasPrice.Div(gasPrice, big.NewInt(10000))

	t.Log("gasPrice", auth.GasPrice)
	t.Log("gas * price", gasPrice.Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))

	address, tx, instance, err := store.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x6b450d2b53Bd6C2d866F5630eDc3bB61e8016A91
	fmt.Println(tx.Hash().Hex()) // 0x1a3f4e7ec26b67255efbf696d48484b56760d76668ef312deb3cee30773a2d86

	_ = instance
}

func Test_Balance(t *testing.T) {
	Test_Init(t)

	balance, err := ec.BalanceAt(context.Background(), address, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(balance)
}

func Test_ChainID(t *testing.T) {
	Test_Init(t)

	id, err := ec.ChainID(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Log(id)
}

// -----------------------------------------------------

func CreateClientMgr () *client.ClientMgr {
	c := &client.ClientMgr{}

	ec, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewStore(ConfigAddress, ec)
	if err != nil {
		log.Fatal(err)
	}

	c.NewEthClient("Config", instance)
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
			name:    "general",
			fields:  fields{
				Client:            CreateClientMgr(),
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