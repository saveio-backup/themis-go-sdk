package usdt

import (
	"github.com/oniio/oniChain-go-sdk/client"
	sdkcom "github.com/oniio/oniChain-go-sdk/common"
	"github.com/oniio/oniChain-go-sdk/utils"
	"github.com/oniio/oniChain/account"
	"github.com/oniio/oniChain/common"
	"github.com/oniio/oniChain/core/types"
	"github.com/oniio/oniChain/smartcontract/service/native/usdt"
)

var (
	USDT_CONTRACT_ADDRESS, _ = utils.AddressFromHexString("0100000000000000000000000000000000000000")
	USDT_CONTRACT_VERSION    = byte(0)
)

type Usdt struct {
	Client *client.ClientMgr
}

func (this *Usdt) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, USDT_CONTRACT_VERSION, USDT_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return nil, err
	}
	return this.Client.PreExecTransaction(tx)
}

func (this *Usdt) NewTransferTransaction(gasPrice, gasLimit uint64, from, to common.Address, amount uint64) (*types.MutableTransaction, error) {
	state := &usdt.State{
		From:  from,
		To:    to,
		Value: amount,
	}
	return this.NewMultiTransferTransaction(gasPrice, gasLimit, []*usdt.State{state})
}

func (this *Usdt) Transfer(gasPrice, gasLimit uint64, from *account.Account, to common.Address, amount uint64) (common.Uint256, error) {
	tx, err := this.NewTransferTransaction(gasPrice, gasLimit, from.Address, to, amount)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, from)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Usdt) NewMultiTransferTransaction(gasPrice, gasLimit uint64, states []*usdt.State) (*types.MutableTransaction, error) {
	return utils.NewNativeInvokeTransaction(
		gasPrice,
		gasLimit,
		USDT_CONTRACT_VERSION,
		USDT_CONTRACT_ADDRESS,
		usdt.TRANSFER_NAME,
		[]interface{}{states})
}

func (this *Usdt) MultiTransfer(gasPrice, gasLimit uint64, states []*usdt.State, signer *account.Account) (common.Uint256, error) {
	tx, err := this.NewMultiTransferTransaction(gasPrice, gasLimit, states)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Usdt) NewTransferFromTransaction(gasPrice, gasLimit uint64, sender, from, to common.Address, amount uint64) (*types.MutableTransaction, error) {
	state := &usdt.TransferFrom{
		Sender: sender,
		From:   from,
		To:     to,
		Value:  amount,
	}
	return utils.NewNativeInvokeTransaction(
		gasPrice,
		gasLimit,
		USDT_CONTRACT_VERSION,
		USDT_CONTRACT_ADDRESS,
		usdt.TRANSFERFROM_NAME,
		[]interface{}{state},
	)
}

func (this *Usdt) TransferFrom(gasPrice, gasLimit uint64, sender *account.Account, from, to common.Address, amount uint64) (common.Uint256, error) {
	tx, err := this.NewTransferFromTransaction(gasPrice, gasLimit, sender.Address, from, to, amount)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, sender)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Usdt) NewApproveTransaction(gasPrice, gasLimit uint64, from, to common.Address, amount uint64) (*types.MutableTransaction, error) {
	state := &usdt.State{
		From:  from,
		To:    to,
		Value: amount,
	}
	return utils.NewNativeInvokeTransaction(
		gasPrice,
		gasLimit,
		USDT_CONTRACT_VERSION,
		USDT_CONTRACT_ADDRESS,
		usdt.APPROVE_NAME,
		[]interface{}{state},
	)
}

func (this *Usdt) Approve(gasPrice, gasLimit uint64, from *account.Account, to common.Address, amount uint64) (common.Uint256, error) {
	tx, err := this.NewApproveTransaction(gasPrice, gasLimit, from.Address, to, amount)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, from)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Usdt) Allowance(from, to common.Address) (uint64, error) {
	type allowanceStruct struct {
		From common.Address
		To   common.Address
	}

	preResult, err := this.PreExecInvokeNativeContract(
		usdt.ALLOWANCE_NAME,
		[]interface{}{&allowanceStruct{From: from, To: to}},
	)
	if err != nil {
		return 0, err
	}
	balance, err := preResult.Result.ToInteger()
	if err != nil {
		return 0, err
	}
	return balance.Uint64(), nil
}

func (this *Usdt) Symbol() (string, error) {
	preResult, err := this.PreExecInvokeNativeContract(
		usdt.SYMBOL_NAME,
		[]interface{}{},
	)
	if err != nil {
		return "", err
	}
	return preResult.Result.ToString()
}

func (this *Usdt) BalanceOf(address common.Address) (uint64, error) {
	preResult, err := this.PreExecInvokeNativeContract(
		usdt.BALANCEOF_NAME,
		[]interface{}{address[:]},
	)
	if err != nil {
		return 0, err
	}
	balance, err := preResult.Result.ToInteger()
	if err != nil {
		return 0, err
	}
	return balance.Uint64(), nil
}

func (this *Usdt) Name() (string, error) {
	preResult, err := this.PreExecInvokeNativeContract(
		usdt.NAME_NAME,
		[]interface{}{},
	)
	if err != nil {
		return "", err
	}
	return preResult.Result.ToString()
}

func (this *Usdt) Decimals() (byte, error) {
	preResult, err := this.PreExecInvokeNativeContract(
		usdt.DECIMALS_NAME,
		[]interface{}{},
	)
	if err != nil {
		return 0, err
	}
	decimals, err := preResult.Result.ToInteger()
	if err != nil {
		return 0, err
	}
	return byte(decimals.Uint64()), nil
}

func (this *Usdt) TotalSupply() (uint64, error) {
	preResult, err := this.PreExecInvokeNativeContract(
		usdt.TOTAL_SUPPLY_NAME,
		[]interface{}{},
	)
	if err != nil {
		return 0, err
	}
	balance, err := preResult.Result.ToInteger()
	if err != nil {
		return 0, err
	}
	return balance.Uint64(), nil
}
