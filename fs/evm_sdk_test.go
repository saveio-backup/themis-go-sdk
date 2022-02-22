package fs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	configStore "github.com/saveio/themis-go-sdk/fs/contracts/Config"
	fsStore "github.com/saveio/themis-go-sdk/fs/contracts/FileSystem"
	listStore "github.com/saveio/themis-go-sdk/fs/contracts/List"
	nodeStore "github.com/saveio/themis-go-sdk/fs/contracts/Node"
	pdpStore "github.com/saveio/themis-go-sdk/fs/contracts/PDP"
	proveStore "github.com/saveio/themis-go-sdk/fs/contracts/Prove"
	sectorStore "github.com/saveio/themis-go-sdk/fs/contracts/Sector"
	spaceStore "github.com/saveio/themis-go-sdk/fs/contracts/Space"
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

// ====================================================

func Test_Deploy_Config(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := configStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_Node(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := nodeStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_Sector(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := sectorStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_Space(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := spaceStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_FS(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := fsStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_List(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := listStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_Prove(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := proveStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func Test_Deploy_PDP(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	address, tx, instance, err := pdpStore.DeployStore(auth, ec)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

// ====================================================

func Test_Init_Node(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	store, err := nodeStore.NewStore(NodeAddress, ec)
	if err != nil {
		t.Fatal(err)
	}

	initialize, err := store.Initialize(auth, ConfigAddress, SectorAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(initialize.Hash())
}

func Test_Init_Sector(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	store, err := sectorStore.NewStore(NodeAddress, ec)
	if err != nil {
		t.Fatal(err)
	}

	initialize, err := store.Initialize(auth, NodeAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(initialize.Hash())
}

func Test_Init_FS(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	store, err := fsStore.NewStore(NodeAddress, ec)
	if err != nil {
		t.Fatal(err)
	}

	initialize, err := store.Initialize(auth, ConfigAddress, NodeAddress, SpaceAddress, SectorAddress, ProveAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(initialize.Hash())
}

func Test_Init_Space(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	store, err := spaceStore.NewStore(NodeAddress, ec)
	if err != nil {
		t.Fatal(err)
	}

	initialize, err := store.Initialize(auth, ConfigAddress, FSAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(initialize.Hash())
}

func Test_Init_Prove(t *testing.T) {
	Test_Init(t)

	eth := &EVM{
		Client: CreateClientMgr(),
	}
	auth, _ := eth.GetSigner(big.NewInt(0))

	store, err := proveStore.NewStore(NodeAddress, ec)
	if err != nil {
		t.Fatal(err)
	}

	initialize, err := store.Initialize(auth, ConfigAddress, FSAddress, NodeAddress, PDPAddress, SectorAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(initialize.Hash())
}

// ====================================================
// ====================================================
// ====================================================

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
