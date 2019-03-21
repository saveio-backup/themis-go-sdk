package governance

import (
	"errors"
	"fmt"

	"github.com/oniio/oniChain-go-sdk/client"
	sdkcom "github.com/oniio/oniChain-go-sdk/common"
	"github.com/oniio/oniChain-go-sdk/utils"
	"github.com/oniio/oniChain/account"
	"github.com/oniio/oniChain/common"
	gov "github.com/oniio/oniChain/smartcontract/service/native/governance"
)

var (
	GOV_CONTRACT_ADDRESS, _ = utils.AddressFromHexString("0700000000000000000000000000000000000000")
	GOV_CONTRACT_VERSION    = byte(0)
)

type Governance struct {
	Client *client.ClientMgr
	DefAcc *account.Account
}

func (this *Governance) InvokeNativeContract(signer *account.Account, method string, params []interface{}) (common.Uint256, error) {
	if signer == nil {
		return common.UINT256_EMPTY, errors.New("signer is nil")
	}
	tx, err := utils.NewNativeInvokeTransaction(sdkcom.GAS_PRICE, sdkcom.GAS_LIMIT, GOV_CONTRACT_VERSION, GOV_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Governance) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, GOV_CONTRACT_VERSION, GOV_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return nil, err
	}
	return this.Client.PreExecTransaction(tx)
}

func (this *Governance) RegisterCandidate(peerPubkey string, initPos uint64) ([]byte, error) {
	if this.DefAcc == nil {
		fmt.Println("DefAcc is nil")
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.REGISTER_CANDIDATE,
		[]interface{}{&gov.RegisterCandidateParam{PeerPubkey: peerPubkey, Address: this.DefAcc.Address, InitPos: initPos}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (this *Governance) UnRegisterCandidate(peerPubkey string) ([]byte, error) {
	if this.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.UNREGISTER_CANDIDATE,
		[]interface{}{&gov.UnRegisterCandidateParam{PeerPubkey: peerPubkey, Address: this.DefAcc.Address}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (this *Governance) Withdraw(peerPubkeyList []string, withdrawList []uint64) ([]byte, error) {
	if this.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}

	if len(peerPubkeyList) != len(withdrawList) {
		return nil, errors.New("peerPubkeyList and withdrawList not match")
	}

	withdrawParam := &gov.WithdrawParam{}
	withdrawParam.Address = this.DefAcc.Address
	withdrawParam.PeerPubkeyList = peerPubkeyList
	withdrawParam.WithdrawList = withdrawList

	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.WITHDRAW,
		[]interface{}{withdrawParam},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (this *Governance) QuitNode(peerPubkey string) ([]byte, error) {
	if this.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.QUIT_NODE,
		[]interface{}{&gov.QuitNodeParam{PeerPubkey: peerPubkey, Address: this.DefAcc.Address}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (this *Governance) WithdrawFee(peerPubkey string) ([]byte, error) {
	if this.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.WITHDRAW_FEE,
		[]interface{}{&gov.WithdrawFeeParam{Address: this.DefAcc.Address}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (this *Governance) AddInitPos(peerPubkey string, pos uint64) ([]byte, error) {
	if this.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.ADD_INIT_POS,
		[]interface{}{&gov.ChangeInitPosParam{PeerPubkey: peerPubKey, Address: this.DefAcc.Address, Pos: pos}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}

func (this *Governance) ReduceInitPos(peerPubkey string, pos uint64) ([]byte, error) {
	if this.DefAcc == nil {
		return nil, errors.New("DefAcc is nil")
	}
	ret, err := this.InvokeNativeContract(this.DefAcc,
		gov.REDUCE_INIT_POS,
		[]interface{}{&gov.ChangeInitPosParam{PeerPubkey: peerPubKey, Address: this.DefAcc.Address, Pos: pos}},
	)
	if err != nil {
		return nil, err
	}
	return ret.ToArray(), err
}
