package fs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	nodeStore "github.com/saveio/themis-go-sdk/fs/contracts/Node"
	"log"
	"math/big"
	"testing"
)

var url = "http://104.225.158.242:8545"

var ec *ethclient.Client

func Test_Init(t *testing.T) {
	var err error
	// eth client
	ec, err = ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Deploy_Config(t *testing.T) {
	Test_Init(t)

	eth := &Ethereum{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

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
	auth.GasPrice = gasPrice.Div(gasPrice, big.NewInt(1000000))

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

func Test_Init_Node(t *testing.T) {
	Test_Init(t)

	store, err := nodeStore.NewStore(NodeAddress, ec)
	if err != nil {
		t.Fatal(err)
	}

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
	auth.GasPrice = gasPrice.Div(gasPrice, big.NewInt(10000000))

	t.Log("gasPrice", auth.GasPrice)
	t.Log("gas * price", gasPrice.Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))

	initialize, err := store.Initialize(auth, ConfigAddress, ethCommon.Address{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(initialize.Hash()) // 0x64cf635bbc8e591daca19bac6ed8033ed38c11895e79eadd8b5322c6ceda26b3
}

// ----------------------------------------------------

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

	receipt, err := ec.TransactionReceipt(context.Background(), ethCommon.HexToHash("0xe7fe4a760f3b16c07dd92c2813cf635413595b1bb13e425ba34a75f68ec419b3"))
	if err != nil {
		log.Fatal(err)
	}
	t.Log(receipt.Status)
}


