package chain

import (
	"encoding/hex"
	"fmt"
	"time"

	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/core/payload"
	"github.com/saveio/themis/core/types"
	httpcom "github.com/saveio/themis/http/base/common"
)

type NeoVMContract struct {
	chain *Chain
}

func newNeoVMContract(chain *Chain) *NeoVMContract {
	return &NeoVMContract{
		chain: chain,
	}
}

func (this *NeoVMContract) NewDeployNeoVMCodeTransaction(gasPrice, gasLimit uint64, contract *sdkcom.SmartContract) *types.MutableTransaction {
	deployPayload := &payload.DeployCode{
		Name:        contract.Name,
		Version:     contract.Version,
		Author:      contract.Author,
		Email:       contract.Email,
		Description: contract.Description,
	}
	tx := &types.MutableTransaction{
		Version:  sdkcom.VERSION_TRANSACTION,
		TxType:   types.Deploy,
		Nonce:    uint32(time.Now().Unix()),
		Payload:  deployPayload,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Sigs:     make([]types.Sig, 0, 0),
	}
	return tx
}

//DeploySmartContract Deploy smart contract to oniChain
func (this *NeoVMContract) DeployNeoVMSmartContract(
	gasPrice,
	gasLimit uint64,
	singer *account.Account,
	needStorage bool,
	code,
	name,
	version,
	author,
	email,
	desc string) (common.Uint256, error) {

	invokeCode, err := hex.DecodeString(code)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("code hex decode error:%s", err)
	}
	// FIXME: setup invokeCode
	fmt.Println(invokeCode)
	tx := this.NewDeployNeoVMCodeTransaction(gasPrice, gasLimit, &sdkcom.SmartContract{
		Name:        name,
		Version:     version,
		Author:      author,
		Email:       email,
		Description: desc,
	})
	err = utils.SignToTransaction(tx, singer)
	if err != nil {
		return common.Uint256{}, err
	}
	txHash, err := this.chain.SendTransaction(tx)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("SendRawTransaction error:%s", err)
	}
	return txHash, nil
}

func (this *NeoVMContract) NewNeoVMInvokeTransaction(
	gasPrice,
	gasLimit uint64,
	contractAddress common.Address,
	params []interface{},
) (*types.MutableTransaction, error) {
	invokeCode, err := httpcom.BuildNeoVMInvokeCode(contractAddress, params)
	if err != nil {
		return nil, err
	}
	return utils.NewInvokeTransaction(gasPrice, gasLimit, invokeCode), nil
}

func (this *NeoVMContract) InvokeNeoVMContract(
	gasPrice,
	gasLimit uint64,
	signer *account.Account,
	contractAddress common.Address,
	params []interface{}) (common.Uint256, error) {
	tx, err := this.NewNeoVMInvokeTransaction(gasPrice, gasLimit, contractAddress, params)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("NewNeoVMInvokeTransaction error:%s", err)
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.chain.SendTransaction(tx)
}

func (this *NeoVMContract) PreExecInvokeNeoVMContract(
	contractAddress common.Address,
	params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := this.NewNeoVMInvokeTransaction(0, 0, contractAddress, params)
	if err != nil {
		return nil, err
	}
	return this.chain.PreExecTransaction(tx)
}
