package usdt

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/saveio/themis-go-sdk/client"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common/log"
	"math/big"
	"math/rand"
)

type ERC20 struct {
	Client *client.ClientMgr
	DefAcc *account.Account
}

func (t *ERC20) GetSigner(value *big.Int) (*bind.TransactOpts, error) {
	ec := t.Client.GetEthClient().Client
	nonce, err := ec.PendingNonceAt(context.TODO(), t.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}
	chainID, err := ec.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}
	keyBytes := t.DefAcc.GetEthPrivateKey()
	keyStr := hexutil.Encode(keyBytes)
	key, err := crypto.HexToECDSA(keyStr[2:])
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	auth.From = t.DefAcc.EthAddress
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value                 // in wei
	auth.GasLimit = uint64(10_000_000) // in units
	auth.GasPrice = gasPrice.Add(gasPrice, big.NewInt(int64(rand.Intn(1))))

	gas := new(big.Int).Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit)))
	gas = new(big.Int).Add(auth.Value, gas)
	log.Debugf("get signer with gas: %v, %v, %v, %v", auth.GasPrice, auth.GasLimit, auth.Value, gas)
	return auth, nil
}

func (e *ERC20) BalanceOf(address common.Address) (uint64, error) {
	at, err := e.Client.GetEthClient().Client.BalanceAt(context.TODO(), address, nil)
	if err != nil {
		return 0, err
	}
	return at.Uint64(), nil
}

func (e *ERC20) Transfer(from, to common.Address, amount uint64) ([]byte, error) {
	ec := e.Client.GetEthClient().Client
	nonce, err := ec.PendingNonceAt(context.TODO(), e.DefAcc.EthAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := ec.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}

	var data []byte
	tx := types.NewTransaction(nonce, to, big.NewInt(int64(amount)), uint64(10_000_000), gasPrice, data)
	chainID, err := ec.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}
	keyBytes := e.DefAcc.GetEthPrivateKey()
	keyStr := hexutil.Encode(keyBytes)
	privateKey, err := crypto.HexToECDSA(keyStr[2:])
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	err = ec.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx.Hash().Bytes(), nil
}
