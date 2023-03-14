package film

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/dns"
	"github.com/saveio/themis-go-sdk/fs"
	"github.com/saveio/themis-go-sdk/wallet"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/smartcontract/service/native/film"
)

var testFilm *Film
var walletPath = "./wallet.dat"
var pwd = []byte("pwd")
var rpc_addr = "http://127.0.0.1:20336"
var txt = "QmevhnWdtmz89BMXuuX5pSY2uZtqKLz7frJsrCojT5kmb6"
var current *account.Account
var testDns *dns.Native
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
	testFilm.Client.NewRpcClient().SetAddress([]string{rpc_addr})
	current = acc
	testFilm.DefAcc = acc

	testDns = &dns.Native{}
	testDns.Client = testFilm.Client
	testDns.DefAcc = acc

	testFs = &fs.Fs{}
	//testFs.Client = testFilm.Client
	//testFs.DefAcc = acc

	m.Run()
}

func TestRegisterFilmUrl(t *testing.T) {
	registerCustomHeader := false
	if registerCustomHeader {
		hash, err := testDns.RegisterHeader("oni", "oni", 10000)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("hash: %v\n", hash)
		time.Sleep(time.Duration(3) * time.Second)
	}
	url := "oni://share/123456"
	link := "oni-link://QmUJif4gNXLH5YFgGVUMv86u4K57ET1pED1TLSUaHtUq3p&name=2019-08-23_15.00.32_LOG.log&owner=AY46Kes2ayy8c38hKBqictG9F9ar73mqhD&size=20595&blocknum=82&tr=dWRwOi8vMTY4LjYzLjI1My4yMzE6NjM2OS9hbm5vdW5jZQ==&tr=dWRwOi8vMTY4LjYzLjI1My4yMzE6NjM2OS9hbm5vdW5jZQ== "
	regHash, err := testDns.RegisterUrl(url, 0x02, link, link, 600)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("hash: %v\n", regHash)

}

func TestPublishFilm(t *testing.T) {
	cover := "123"
	url := "oni://share/123456"
	name := "name"
	desc := "desc"
	filmType := 1
	releaseY := 2019
	language := "EN"
	region := "USA"
	price := 0
	params := []interface{}{cover, url, name, desc, filmType, releaseY, language, region, price}
	buf, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("bufs : %v\n", buf)
	hash, err := testFilm.InvokeNativeContract(testFilm.DefAcc, film.FILM_PUBLISH, []interface{}{buf})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("hash: %v\n", hash)
}

func TestGetUserFilms(t *testing.T) {
	ret, err := testFilm.PreExecInvokeNativeContract(film.GET_USER_FILM_LIST, []interface{}{testFilm.DefAcc.Address})
	if err != nil {
		t.Fatal(err)
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		t.Fatal(err)
	}
	films := &film.UserFilmList{}
	buf := bytes.NewReader(data)
	err = films.Deserialize(buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("num : %v\n", films.Num)
	fmt.Printf("owner : %v\n", films.Owner.ToBase58())
	fmt.Printf("films : %s\n", films.FilmHashes)
}

func TestGetFilmLists(t *testing.T) {
	name := "name"
	available := 1
	filmType := 1
	releaseY := 0
	region := "USA"
	offset := 0
	limit := 0
	params := []interface{}{name, filmType, releaseY, region, available, offset, limit}
	buf, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := testFilm.PreExecInvokeNativeContract(film.GET_FILM_LIST, []interface{}{buf})
	if err != nil {
		t.Fatal(err)
	}
	data, err := ret.Result.ToByteArray()
	if err != nil {
		t.Fatal(err)
	}
	str := hex.EncodeToString(data)
	fmt.Printf("strt :%s\n", str)

}

func TestDecodeData(t *testing.T) {
	str := "5b7b2248617368223a2255573155516d35454d56673461465a7164464a6a567a4e4354444e3552476c725a3346435345316a526a464e6331565659334134635755346332746855773d3d222c22436f766572223a22513239325a58493d222c2255726c223a22623235704f69387663326868636d55764e544d344f5755315a54513d222c224e616d65223a224d6a41784f5330774e7930784f4638774e6934784e5334304e31394d543063756247396e222c2244657363223a224d6a41784f5330774e7930784f4638774e6934784e5334304e31394d543063756247396e222c22417661696c61626c65223a747275652c2254797065223a312c2252656c6561736559656172223a323031392c224c616e6775616765223a225255343d222c22526567696f6e223a225130684a546b453d222c225072696365223a302c22437265617465644174223a302c2250616964436f756e74223a302c22546f74616c50726f666974223a307d5d"

	decodedData, err := hex.DecodeString(str)
	if err != nil {
		t.Fatal(err)
	}
	list := make([]*film.FilmInfo, 0)
	err = json.Unmarshal(decodedData, &list)
	if err != nil {
		t.Fatal(err)
	}
	for _, info := range list {
		fmt.Printf("Hash: %s\n", info.Hash)
		fmt.Printf("Cover: %s\n", info.Cover)
		fmt.Printf("Url: %s\n", info.Url)
		fmt.Printf("Name: %s\n", info.Name)
		fmt.Printf("Desc: %s\n", info.Desc)
		fmt.Printf("Available: %v\n", info.Available)
		fmt.Printf("Type: %v\n", info.Type)
		fmt.Printf("ReleaseYear: %v\n", info.ReleaseYear)
		fmt.Printf("Language: %s\n", info.Language)
		fmt.Printf("Region: %s\n", info.Region)
		fmt.Printf("Price: %v\n", info.Price)
		fmt.Printf("CreatedAt: %v\n", info.CreatedAt)
		fmt.Printf("PaidCount: %v\n", info.PaidCount)
		fmt.Printf("TotalProfit: %v\n", info.TotalProfit)
	}
}

func TestConvFilmAddrToBase58(t *testing.T) {
	fmt.Printf("%s\n", FILM_CONTRACT_ADDRESS.ToBase58())
}
