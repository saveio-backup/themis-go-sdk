package fs

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/wallet"
	"github.com/saveio/themis/common"
	fs "github.com/saveio/themis/smartcontract/service/native/savefs"
)

var testFs *Fs
var walletAddr = "AHjjdbVLhfTyiNFEq2X8mFnnirZY1yK8Rq"
var walletPath = "./wallet.dat"
var pwd = []byte("pwd")
var rpc_addr = "http://127.0.0.1:20336"
var txt = "QmevhnWdtmz89BMXuuX5pSY2uZtqKLz7frJsrCojT5kmb6"

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	var err error
	w, err := wallet.OpenWallet(walletPath)
	if err != nil {
		fmt.Printf("Account.Open error:%s\n", err)
	}
	acc, err := w.GetDefaultAccount(pwd)
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
	}
	testFs = &Fs{}
	testFs.Client = &client.ClientMgr{}
	testFs.Client.NewRpcClient().SetAddress([]string{rpc_addr})
	testFs.DefAcc = acc
	m.Run()
}

func TestCalculateFee(t *testing.T) {
	// addr, _ := common.AddressFromBase58("ARy6wVzHWsrzmKzSorrTRzf8B5oYGvdUmK")
	// rule := fs.Rule{
	// 	Addr:         addr,
	// 	BaseHeight:   0,
	// 	ExpireHeight: 0,
	// }
	rules := make([]fs.Rule, 0)
	// rules = append(rules, rule)
	wh := fs.WhiteList{
		Num:  0,
		List: rules,
	}

	opt := &fs.UploadOption{
		FileDesc:        []byte(""),
		FileSize:        uint64(166974),
		ProveInterval:   uint64(60),
		ExpiredHeight:   1563560,
		Privilege:       1,
		CopyNum:         uint64(2),
		Encrypt:         false,
		EncryptPassword: []byte{},
		RegisterDNS:     true,
		BindDNS:         true,
		DnsURL:          []byte{},
		WhiteList:       wh,
		Share:           true,
		StorageType:     uint64(0),
	}
	fee, err := testFs.GetUploadStorageFee(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("fee :%v\n", fee)
}

func TestOntFsClient_GetNodeList(t *testing.T) {
	ret, err := testFs.GetNodeList()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	nodesInfo := (*fs.FsNodesInfo)(ret)

	for _, nodeInfo := range nodesInfo.NodeInfo {
		fmt.Println("FsNodeQuery Success")
		fmt.Println("Pledge: ", nodeInfo.Pledge)
		fmt.Println("Profit: ", nodeInfo.Profit)
		fmt.Println("Volume: ", nodeInfo.Volume)
		fmt.Println("RestVol: ", nodeInfo.RestVol)
		fmt.Println("NodeAddr: ", string(nodeInfo.NodeAddr))
		fmt.Println("WalletAddr: ", nodeInfo.WalletAddr)
		fmt.Println("ServiceTime:", nodeInfo.ServiceTime)
	}
}

func TestOntFsClient_GetFileInfo(t *testing.T) {
	fileInfo, err := testFs.GetFileInfo(txt)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("FileHash:", fileInfo.FileHash)
	fmt.Println("FileOwner:", fileInfo.FileOwner)
	fmt.Println("CopyNum:", fileInfo.CopyNum)
	fmt.Println("ProveInterval:", fileInfo.ProveInterval)
	fmt.Println("ExpiredHeight:", fileInfo.ExpiredHeight)
	fmt.Println("FileBlockNum:", fileInfo.FileBlockNum)
	fmt.Println("FIleBlockSize:", fileInfo.FileBlockSize)
	fmt.Println("Deposit:", fileInfo.Deposit)
	fmt.Println("FileProveParam:", string(fileInfo.FileProveParam))
	fmt.Println("ProveBlockNum:", fileInfo.ProveBlockNum)
}

func TestOntFsClient_GetFileProveDetails(t *testing.T) {
	fileProveDetails, err := testFs.GetFileProveDetails(txt)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("CopyNum:", fileProveDetails.CopyNum)
	fmt.Println("ProveDetailNum:", fileProveDetails.ProveDetailNum)
	var i uint64
	for i = 0; i < fileProveDetails.ProveDetailNum; i++ {
		fmt.Println("NodeAddr: ", fileProveDetails.ProveDetails[i].NodeAddr)
		fmt.Println("WalletAddr:", fileProveDetails.ProveDetails[i].WalletAddr)
		fmt.Println("ProveTimes:", fileProveDetails.ProveDetails[i].ProveTimes)
	}
}

func TestOntFs_OntFsInit(t *testing.T) {
	testFs.savefsInit(2000, 2, 1, 1,
		32, fs.PROVE_LEVEL_HIGH, 1024*1024)
}

func TestOntFs_GetSetting(t *testing.T) {
	fsSet, err := testFs.GetSetting()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("FsGetSetting Success")
	fmt.Println("FsGasPrice:", fsSet.FsGasPrice,
		"GasPerKBForStore:", fsSet.GasPerGBPerBlock,
		"GasPerKBForRead:", fsSet.GasPerKBForRead,
		"GasForChallenge:", fsSet.GasForChallenge,
		"MaxProveBlockNum:", fsSet.MaxProveBlockNum)
}

func TestOntFs_NodeRegister(t *testing.T) {
	_, err := testFs.NodeRegister(107374182400, 10000, "https://127.0.0.1:5002")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("NodeRegister Success")
}

func TestOntFs_NodeQuery(t *testing.T) {
	fsNodeInfo, err := testFs.NodeQuery(testFs.DefAcc.Address)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("NodeQuery Success")
	fmt.Println("Pledge: ", fsNodeInfo.Pledge)
	fmt.Println("Profit: ", fsNodeInfo.Profit)
	fmt.Println("Volume: ", fsNodeInfo.Volume)
	fmt.Println("RestVol: ", fsNodeInfo.RestVol)
	fmt.Println("NodeAddr: ", string(fsNodeInfo.NodeAddr))
	fmt.Println("WalletAddr: ", fsNodeInfo.WalletAddr)
	fmt.Println("ServiceTime:", fsNodeInfo.ServiceTime)
}

func TestOntFs_NodeUpdate(t *testing.T) {
	_, err := testFs.NodeUpdate(100000, 100000, "https://127.0.0.1:5001")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("NodeUpdate Success")
}

func TestOntFs_NodeCancel(t *testing.T) {
	_, err := testFs.NodeCancel()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("NodeCancel Success")
}

func TestOntFs_NodeWithDrawProfit(t *testing.T) {
	_, err := testFs.NodeWithDrawProfit()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("NodeWithDrawProfit Success")
}

func Testsavefs_AddUserSpace(t *testing.T) {
	if testFs == nil {
		t.Fatal("testFs is nil")
	}
	ret, err := testFs.UpdateUserSpace(testFs.DefAcc.Address, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("add user space ret %v\n", ret)
}

func TestOntFs_GetUserSpace(t *testing.T) {
	userspace, err := testFs.GetUserSpace(testFs.DefAcc.Address)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("Used : %d\n", userspace.Used)
	fmt.Printf("Remain : %d\n", userspace.Remain)
	fmt.Printf("ExpiredHeight : %d\n", userspace.ExpireHeight)
	fmt.Printf("Blanace : %d\n", userspace.Balance)
}

func TestGetNodeListByAddrs(t *testing.T) {
	addr, err := common.AddressFromBase58("ALQ6RWJENsELE7ATuzHz4zgHrq573xJsnM")
	if err != nil {
		t.Fatal(err)
	}
	addrs := make([]common.Address, 0)
	addrs = append(addrs, addr)
	info, err := testFs.GetNodeListByAddrs(addrs)
	if err != nil {
		t.Fatal(err)
	}
	if info.NodeNum != uint64(len(addrs)) {
		t.Fatalf("info.NodeNum %d != len(addrs) %d", info.NodeNum, len(addrs))
	}
	for _, info := range info.NodeInfo {
		if info.WalletAddr.ToBase58() != addr.ToBase58() {
			t.Fatalf("wallet not equal %s %s", info.WalletAddr.ToBase58(), addr.ToBase58())
		}
	}
	fmt.Printf("%v\n", info)
}
