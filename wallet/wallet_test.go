package wallet

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestOpenWalletWithData(t *testing.T) {
	dataStr := []byte(`{"name":"MyWallet","version":"1.1","scrypt":{"p":8,"n":16384,"r":8,"dkLen":64},"accounts":[{"address":"AYMnqA65pJFKAbbpD8hi5gdNDBmeFBy5hS","enc-alg":"aes-256-gcm","key":"9tRqfmW3H3cTOGQgKjwZWlOz0dKY5xfbiQ2Hw+eiNShhCcDqXIk73a8lvh/uNwve","algorithm":"ECDSA","salt":"feUbTaXb83Vy7QrGxJRQWw==","parameters":{"curve":"P-256"},"label":"123","publicKey":"02ead0eeca8355e2a5c444fa193a390ecdb8c671f26120de6bc329ab72586979c0","signatureScheme":"SHA256withECDSA","isDefault":true,"lock":false}]}`)
	ioutil.WriteFile("./wallet.dat", dataStr, 0666)
	data, err := ioutil.ReadFile("./wallet.dat")
	if err != nil {
		t.Fatal(err)
	}

	wal, err := OpenWithWalletData(data)
	if err != nil {
		t.Fatal(err)
	}
	if wal == nil {
		t.Fatal("wallet is nil")
	}
	acc, err := wal.GetDefaultAccount([]byte("pwd"))
	if err != nil {
		t.Fatal(err)
	}
	if acc == nil {
		t.Fatal("acc is nil")
	}
	accData, err := wal.GetDefaultAccountData()
	if err != nil {
		t.Fatal(err)
	}
	if accData == nil {
		t.Fatal("acc is nil")
	}
	fmt.Printf("label :%s", accData.Label)
	fmt.Printf("acc :%s", acc.Address.ToBase58())
}
