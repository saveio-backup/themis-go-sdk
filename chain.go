package chain

import (
	"math/rand"
	"time"

	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Chain struct {
	client.ClientMgr
	Native *NativeContract
	NeoVM  *NeoVMContract
	EVM    *EVMContract
}

//NewChain return Chain.
func NewChain() *Chain {
	chain := &Chain{}
	native := newNativeContract(chain.GetClientMgr())
	chain.Native = native
	neoVM := newNeoVMContract(chain)
	chain.NeoVM = neoVM
	evm := newEVMContract(chain.GetClientMgr())
	chain.EVM = evm
	return chain
}

func (this *Chain) SetDefaultAccount(acc *account.Account) {
	this.Native.SetDefaultAccount(acc)
	this.EVM.SetDefaultAccount(acc)
}
func (this *Chain) NewVerifierService() {
	this.EVM.Fs.NewVerifierService()
}

func (this *Chain) InvokeNativeContract(
	gasPrice,
	gasLimit uint64,
	signer *account.Account,
	version byte,
	contractAddress common.Address,
	method string,
	params []interface{},
) (common.Uint256, error) {
	tx, err := utils.NewNativeInvokeTransaction(gasPrice, gasLimit, version, contractAddress, method, params)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.SendTransaction(tx)
}

func (this *Chain) InvokeNativeContractWithCode(
	gasPrice,
	gasLimit uint64,
	signer *account.Account,
	invokeCode []byte,
) (common.Uint256, error) {
	tx := utils.NewInvokeTransaction(gasPrice, gasLimit, invokeCode)
	err := utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.SendTransaction(tx)
}

func (this *Chain) PreExecInvokeNativeContract(
	contractAddress common.Address,
	version byte,
	method string,
	params []interface{},
) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, version, contractAddress, method, params)
	if err != nil {
		return nil, err
	}
	return this.PreExecTransaction(tx)
}
