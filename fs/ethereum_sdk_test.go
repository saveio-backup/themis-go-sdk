package fs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	nodeStore "github.com/saveio/themis-go-sdk/fs/contracts/Node"
	"log"
	"math/big"
	"testing"
)

var url = "http://104.225.158.242:8545"

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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice.Div(gasPrice, big.NewInt(100000))

	t.Log("gasPrice", auth.GasPrice)
	t.Log("gas * price", gasPrice.Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))

	address, tx, instance, err := configStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x6b450d2b53Bd6C2d866F5630eDc3bB61e8016A91
	fmt.Println(tx.Hash().Hex()) // 0x1a3f4e7ec26b67255efbf696d48484b56760d76668ef312deb3cee30773a2d86

	_ = instance
}

func Test_Deploy_Node(t *testing.T) {
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = gasPrice.Div(gasPrice, big.NewInt(100000))

	t.Log("gasPrice", auth.GasPrice)
	t.Log("gas * price", gasPrice.Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))

	address, tx, instance, err := nodeStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x5C373B9C51f65Ec44C315A70c999F7B313Ac90c3
	fmt.Println(tx.Hash().Hex()) // 0xfb7030e6e1e0b082e90da1e3a6d63335cc88f3f8073d55cb3d298aa5b76bb2eb

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

func Test_GetSetting(t *testing.T) {
	Test_Init(t)

	store, err := configStore.NewStore(ConfigAddress, ec)
	if err != nil {
		log.Fatal(err)
	}
	setting, err := store.GetSetting(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	t.Log(setting)
}

func Test_GetTx(t *testing.T) {
	Test_Init(t)

	receipt, err := ec.TransactionReceipt(context.Background(), ethCommon.HexToHash("0xfb7030e6e1e0b082e90da1e3a6d63335cc88f3f8073d55cb3d298aa5b76bb2eb"))
	if err != nil {
		log.Fatal(err)
	}
	t.Log(receipt.Status)
}

