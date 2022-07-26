package usdt

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/saveio/themis-go-sdk/client"
)

type ERC20 struct {
	Client *client.ClientMgr
}

func (e *ERC20) BalanceOf(address common.Address) (uint64, error) {
	at, err := e.Client.GetEthClient().Client.BalanceAt(context.TODO(), address, nil)
	if err != nil {
		return 0, err
	}
	return at.Uint64(), nil
}
