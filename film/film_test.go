package film

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/dns"
	"github.com/saveio/themis-go-sdk/fs"
	"github.com/saveio/themis-go-sdk/wallet"
	"github.com/saveio/themis/account"
)

var testFilm *Film
var walletPath = "./wallet.dat"
var pwd = []byte("pwd")
var rpc_addr = "http://127.0.0.1:20336"
var txt = "QmevhnWdtmz89BMXuuX5pSY2uZtqKLz7frJsrCojT5kmb6"
var current *account.Account
var testDns *dns.Dns
var testFs *fs.Fs

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
	testFilm = &Film{}

	testFilm.Client = &client.ClientMgr{}
	testFilm.Client.NewRpcClient().SetAddress(rpc_addr)
	current = acc
	testFilm.DefAcc = acc

	testDns = &dns.Dns{}
	testDns.Client = testFilm.Client
	testDns.DefAcc = acc

	testFs = &fs.Fs{}
	testFs.Client = testFilm.Client
	testFs.DefAcc = acc

	m.Run()
}

func TestPublishFilm(t *testing.T) {
	hash := "QmevhnWdtmz89BMXuuX5pSY2uZtqKLz7frJsrCojT5kmb6"
	proveParam := []byte("ProveProveProveProveProveProveProveProveProveProveProveProve")
	ret, err := testFs.StoreFile(hash, 12, 32,
		320, 32, 0, []byte("1.jpg"), 1, proveParam)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(ret)
	time.Sleep(time.Duration(10) * time.Second)
	link := "save-link://" + hash
	tx, err := testDns.RegisterUrl("save://share/123123", 0x02, link, link, 600)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tx: %v", tx)
}
