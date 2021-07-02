package governance

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/wallet"
	"github.com/saveio/themis/common"
)

var testGov *Governance
var walletPath = "./wallet.dat"
var walletAddr = "AdCRkcsgAux2bZe2feK3FwgPSDd9qedpPR"
var peerPubKey = "0264a6f59d821c0b4bded42804ccb749da1006ed2dc234adbc990ffe3dba565245"
var pwd = []byte("passw0rd")
var rpc_addr = "http://127.0.0.1:20336"

func TestMain(m *testing.M) {
	var err error
	w, err := wallet.OpenWallet(walletPath)
	if err != nil {
		fmt.Printf("Account.Open error:%s\n", err)
	}
	acc, err := w.GetDefaultAccount(pwd)
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
	}
	testGov = &Governance{}
	testGov.Client = &client.ClientMgr{}
	testGov.Client.NewRpcClient().SetAddress([]string{rpc_addr})
	testGov.DefAcc = acc
	m.Run()
}

func TestGov_RegisterCandidate(t *testing.T) {
	tx, err := testGov.RegisterCandidate(peerPubKey, 10000)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("RegisterCandidate Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}

func TestGov_UnRegisterCandidate(t *testing.T) {
	tx, err := testGov.UnRegisterCandidate(peerPubKey)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("UnRegisterCandidate Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}

func TestGov_Withdraw(t *testing.T) {
	peerPubkeyList := []string{peerPubKey}
	withdrawList := []uint32{100}
	tx, err := testGov.Withdraw(peerPubkeyList, withdrawList)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("Withdraw Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}

func TestGov_QuitNode(t *testing.T) {
	tx, err := testGov.QuitNode(peerPubKey)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("QuitNode Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}

func TestGov_WithdrawFee(t *testing.T) {
	tx, err := testGov.WithdrawFee(peerPubKey)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("WithdrawFee Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}

func TestGov_AddInitPos(t *testing.T) {
	tx, err := testGov.AddInitPos(peerPubKey, 12345)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("AddInitPos Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}

func TestGov_ReduceInitPos(t *testing.T) {
	tx, err := testGov.ReduceInitPos(peerPubKey, 10000)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println("ReduceInitPos Success")
	fmt.Printf("tx :%s\n", hex.EncodeToString(common.ToArrayReverse(tx)))
}
