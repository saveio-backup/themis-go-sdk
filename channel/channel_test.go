package channel

import (
	"encoding/hex"
	"fmt"
	"github.com/saveio/themis/crypto/signature"
	"github.com/saveio/themis/smartcontract/service/native/micropayment"
	"github.com/saveio/themis/smartcontract/service/native/utils"
	"strconv"
	"testing"
	"time"

	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis-go-sdk/wallet"
	"github.com/saveio/themis/common"
	"github.com/stretchr/testify/assert"
)

var walletPath = "/Users/smallyu/work/gogs/scan-deploy/node2/wallet.dat"
var pwd = []byte("pwd")
var rpc_addr = "http://127.0.0.1:20336"
var PARTICIPANT2_WALLETADDR = "AWcxaqe5chY3gSyUcBgdEgQrm7Ntr15qrZ" // Channel participant2 address base58 format
var testChannel *Channel

func init() {
	var err error
	w, err := wallet.OpenWallet(walletPath)
	if err != nil {
		fmt.Printf("Account.Open error:%s\n", err)
	}
	acc, err := w.GetDefaultAccount(pwd)
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
	}
	testChannel := &Channel{}
	testChannel.Client = &client.ClientMgr{}
	testChannel.Client.NewRpcClient().SetAddress([]string{rpc_addr})
	testChannel.DefAcc = acc
}

func Init() *Channel {
	var err error
	w, err := wallet.OpenWallet(walletPath)
	if err != nil {
		fmt.Printf("Account.Open error:%s\n", err)
	}
	acc, err := w.GetDefaultAccount(pwd)
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
	}
	testChannel := &Channel{}
	testChannel.Client = &client.ClientMgr{}
	testChannel.Client.NewRpcClient().SetAddress([]string{rpc_addr})
	testChannel.DefAcc = acc
	return testChannel
}

func TestSetTotalDeposit(t *testing.T) {

	wallet2Addr, err := common.AddressFromBase58(PARTICIPANT2_WALLETADDR)
	h, err := testChannel.Client.GetCurrentBlockHeight()
	assert.Nil(t, err)
	height, err := strconv.ParseUint(fmt.Sprint(h), 10, 64)
	fmt.Printf("height: %d\n", height)
	assert.Nil(t, err)
	tx, err := testChannel.SetTotalDeposit(104, testChannel.DefAcc.Address, wallet2Addr, 25)
	assert.Nil(t, err)
	fmt.Printf("tx :%s\n", hex.EncodeToString(tx))
	confirmed, err := testChannel.Client.PollForTxConfirmed(time.Duration(60)*time.Second, tx)
	assert.Nil(t, err)
	assert.True(t, confirmed)
}

func TestGetChannelId(t *testing.T) {

	wallet2Addr, err := common.AddressFromBase58(PARTICIPANT2_WALLETADDR)
	assert.Nil(t, err)
	id, err := testChannel.GetChannelIdentifier(wallet2Addr, testChannel.DefAcc.Address)
	assert.Nil(t, err)
	fmt.Printf("id:%d\n", id)
}

func TestGetChannelInfo(t *testing.T) {
	wallet2Addr, err := common.AddressFromBase58(PARTICIPANT2_WALLETADDR)
	assert.Nil(t, err)
	info, err := testChannel.GetChannelInfo(104, testChannel.DefAcc.Address, wallet2Addr)
	assert.Nil(t, err)
	fmt.Printf("SettleBlockHeight :%v\n", info.SettleBlockHeight)
	fmt.Printf("ChannelState :%v\n", info.ChannelState)
}

func TestGetChannelParticipantInfo(t *testing.T) {
	wallet2Addr, err := common.AddressFromBase58(PARTICIPANT2_WALLETADDR)
	assert.Nil(t, err)
	p, err := testChannel.GetChannelParticipantInfo(104, testChannel.DefAcc.Address, wallet2Addr)
	assert.Nil(t, err)
	fmt.Printf("WalletAddr :%v\n", p.WalletAddr.ToBase58())
	fmt.Printf("Deposit :%v\n", p.Deposit)
	fmt.Printf("WithDrawAmount :%v\n", p.WithDrawAmount)
	fmt.Printf("IP :%v\n", p.IP)
	fmt.Printf("Port :%v\n", p.Port)
	fmt.Printf("Balance :%v\n", p.Balance)
	fmt.Printf("BalanceHash :%v\n", p.BalanceHash)
	fmt.Printf("Nonce :%v\n", p.Nonce)
	fmt.Printf("IsCloser :%v\n", p.IsCloser)
}

func TestRegisterSecret(t *testing.T) {
	secret := "hello world117"
	tx, err := testChannel.RegisterSecret([]byte(secret))
	assert.Nil(t, err)
	fmt.Printf("tx :%s\n", hex.EncodeToString(tx))
	confirmed, err := testChannel.Client.PollForTxConfirmed(time.Duration(60)*time.Second, tx)
	fmt.Printf("confirmed:%t\n", confirmed)
	assert.Nil(t, err)
	assert.True(t, confirmed)
}

func TestGetSecretRevealBlockHeight(t *testing.T) {
	secretHash, err := hex.DecodeString("6e13cc69a7ec5fbd77f98f6e2f2c017a58e48aec55ae93a3428b884ab7068c29")
	assert.Nil(t, err)
	height, err := testChannel.GetSecretRevealBlockHeight(secretHash)
	assert.Nil(t, err)
	fmt.Printf("height :%d\n", height)
}

func TestGetCurrentHeight(t *testing.T) {
	h, _ := testChannel.Client.GetCurrentBlockHeight()
	height, _ := strconv.ParseUint(fmt.Sprint(h), 10, 64)
	fmt.Printf("h:%d\n", height)
}

func TestFilter(t *testing.T) {
	// client := NewRpcClient("http://127.0.0.1:20336")
	// ret, _ := client.GetFilterArgsForAllEventsFromChannel(0, 7373, 7375)
	// fmt.Printf("ret:%v\n", ret)
}

func TestChannel_GetNodePubKey(t *testing.T) {
	channel := Init()
	var addr string
	addr = "ANQKbEjoGFu7Qg9wF4y5i748umAWmrRhnB"
	base58, err := common.AddressFromBase58(addr)
	if err != nil {
		t.Error(err)
	}
	key, err := channel.GetNodePubKey(base58)
	if err != nil {
		t.Error(err)
	}
	t.Log("get key:", key)
}

func TestChannel_SetNodePubKey(t *testing.T) {
	channel := Init()
	toHex := common.PubKeyToHex(channel.DefAcc.PublicKey)
	bytes, err := common.HexToBytes(toHex)
	if err != nil {
		t.Error(err)
	}
	tx, err := channel.SetNodePubKey(channel.DefAcc.Address, bytes)
	if err != nil {
		t.Error(err)
	}
	t.Log("tx:", tx)
}

func TestChannel_GetFeeInfo(t *testing.T) {
	channel := Init()
	var addr string
	addr = "AUwpxrg3Bc8StyDn5tB1eWdHtv5rSmTFzR"
	base58, err := common.AddressFromBase58(addr)
	if err != nil {
		t.Error(err)
	}
	key, err := channel.GetFeeInfo(base58, utils.UsdtContractAddress)
	if err != nil {
		t.Error(err)
	}
	t.Log("get key:", key)
}

func TestChannel_SetFeeInfo(t *testing.T) {
	channel := Init()
	addr := "AUwpxrg3Bc8StyDn5tB1eWdHtv5rSmTFzR"
	walletAddr, err := common.AddressFromBase58(addr)
	msgHash := micropayment.FeeInfoMessageBundleHash(walletAddr, utils.UsdtContractAddress)
	sign, err := signature.Sign(channel.DefAcc.SigScheme, channel.DefAcc.PrivateKey, msgHash[:], nil)
	serialize, err := signature.Serialize(sign)
	pkHex := common.PubKeyToHex(channel.DefAcc.PubKey())
	bytes, err := common.HexToBytes(pkHex)
	feeInfo := micropayment.FeeInfo{
		WalletAddr: walletAddr,
		TokenAddr:  utils.UsdtContractAddress,
		Flat:       20,
		PublicKey:  bytes,
		Signature:  serialize,
	}
	verify := signature.Verify(channel.DefAcc.PublicKey, msgHash[:], sign)
	if !verify {
		t.Error(err)
	}
	tx, err := channel.SetFeeInfo(feeInfo)
	if err != nil {
		t.Error(err)
	}
	t.Log("tx:", tx)
}
