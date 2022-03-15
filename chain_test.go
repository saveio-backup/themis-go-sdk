package chain

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/saveio/themis-go-sdk/wallet"
	"github.com/saveio/themis/account"
)

var (
	testChain    *Chain
	testWallet   *wallet.Wallet
	testPasswd   = []byte("pwd")
	testDefAcc   *account.Account
	testGasPrice = uint64(0)
	testGasLimit = uint64(20000)
)

func TestMain(m *testing.M) {
	testChain = NewChain()
	testChain.NewRpcClient().SetAddress([]string{"http://localhost:20336"})

	var err error
	testWallet, err = wallet.OpenWallet("./wallet.dat")
	if err != nil {
		fmt.Printf("account.Open error:%s\n", err)
		m.Run()
		return
	}
	testDefAcc, err = testWallet.GetDefaultAccount(testPasswd)
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
		return
	}

	ws := testChain.NewWebSocketClient()
	err = ws.Connect("ws://localhost:20335")
	if err != nil {
		fmt.Printf("Connect ws error:%s", err)
		return
	}
	m.Run()
}

func TestOnt_Transfer(t *testing.T) {
	txHash, err := testChain.Native.Usdt.Transfer(testGasPrice, testGasLimit, testDefAcc, testDefAcc.Address, 1)
	if err != nil {
		t.Errorf("NewTransferTransaction error:%s", err)
		return
	}
	testChain.WaitForGenerateBlock(30*time.Second, 1)
	evts, err := testChain.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		t.Errorf("GetSmartContractEvent error:%s", err)
		return
	}
	fmt.Printf("TxHash:%s\n", txHash.ToHexString())
	fmt.Printf("State:%d\n", evts.State)
	fmt.Printf("GasConsume:%d\n", evts.GasConsumed)
	for _, notify := range evts.Notify {
		fmt.Printf("ContractAddress:%s\n", notify.ContractAddress)
		fmt.Printf("States:%+v\n", notify.States)
	}
}

func TestGlobalParam_GetGlobalParams(t *testing.T) {
	gasPrice := "gasPrice"
	params := []string{gasPrice}
	results, err := testChain.Native.GlobalParams.GetGlobalParams(params)
	if err != nil {
		t.Errorf("GetGlobalParams:%+v error:%s", params, err)
		return
	}
	fmt.Printf("Params:%s Value:%v\n", gasPrice, results[gasPrice])
}

func TestGlobalParam_SetGlobalParams(t *testing.T) {
	gasPrice := "gasPrice"
	globalParams, err := testChain.Native.GlobalParams.GetGlobalParams([]string{gasPrice})
	if err != nil {
		t.Errorf("GetGlobalParams error:%s", err)
		return
	}
	gasPriceValue, err := strconv.Atoi(globalParams[gasPrice])
	if err != nil {
		t.Errorf("Get prama value error:%s", err)
		return
	}
	_, err = testChain.Native.GlobalParams.SetGlobalParams(testGasPrice, testGasLimit, testDefAcc, map[string]string{gasPrice: strconv.Itoa(gasPriceValue + 1)})
	if err != nil {
		t.Errorf("SetGlobalParams error:%s", err)
		return
	}
	testChain.WaitForGenerateBlock(30*time.Second, 1)
	globalParams, err = testChain.Native.GlobalParams.GetGlobalParams([]string{gasPrice})
	if err != nil {
		t.Errorf("GetGlobalParams error:%s", err)
		return
	}
	gasPriceValueAfter, err := strconv.Atoi(globalParams[gasPrice])
	if err != nil {
		t.Errorf("Get prama value error:%s", err)
		return
	}
	fmt.Printf("After set params gasPrice:%d\n", gasPriceValueAfter)
}

func TestWsScribeEvent(t *testing.T) {
	wsClient := testChain.ClientMgr.GetWebSocketClient()
	err := wsClient.SubscribeEvent()
	if err != nil {
		t.Errorf("SubscribeTxHash error:%s", err)
		return
	}
	defer wsClient.UnsubscribeTxHash()

	actionCh := wsClient.GetActionCh()
	timer := time.NewTimer(time.Minute * 3)
	for {
		select {
		case <-timer.C:
			return
		case action := <-actionCh:
			fmt.Printf("Action:%s\n", action.Action)
			fmt.Printf("Result:%s\n", action.Result)
		}
	}
}

func TestWsTransfer(t *testing.T) {
	wsClient := testChain.ClientMgr.GetWebSocketClient()
	testChain.ClientMgr.SetDefaultClient(wsClient)
	txHash, err := testChain.Native.Usdt.Transfer(testGasPrice, testGasLimit, testDefAcc, testDefAcc.Address, 1)
	if err != nil {
		t.Errorf("NewTransferTransaction error:%s", err)
		return
	}
	testChain.WaitForGenerateBlock(30*time.Second, 1)
	evts, err := testChain.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		t.Errorf("GetSmartContractEvent error:%s", err)
		return
	}
	fmt.Printf("TxHash:%s\n", txHash.ToHexString())
	fmt.Printf("State:%d\n", evts.State)
	fmt.Printf("GasConsume:%d\n", evts.GasConsumed)
	for _, notify := range evts.Notify {
		fmt.Printf("ContractAddress:%s\n", notify.ContractAddress)
		fmt.Printf("States:%+v\n", notify.States)
	}
}

func TestsavefsInit(t *testing.T) {
	testChain.Native.Fs.Client.SetDefaultAccount(testDefAcc)
	setting, _ := testChain.Native.Fs.GetSetting()
	fmt.Printf("setting: %v\n", setting)
}

func TestGetCurrentHeight(t *testing.T) {
	height, err := testChain.GetCurrentBlockHeight()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("height:%d\n", height)
}
