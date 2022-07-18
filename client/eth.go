package client

import (
	"context"
	"encoding/json"
	"errors"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/core/types"
	"math/big"
	"sync"
	"time"
)

type EthClient struct {
	Client *ethclient.Client
	Urls   *sync.Map // string => int
}

func NewEthClient() *EthClient {
	client := &EthClient{
		Urls:   new(sync.Map),
		Client: nil,
	}
	go client.MonitorBasEthServer()
	return client
}

func (e *EthClient) MonitorBasEthServer() {
	for {
		e.Urls.Range(func(key, value interface{}) bool {
			url := key.(string)
			dial, err := ethclient.Dial(url)
			if err != nil {
				return true
			}
			id, err := dial.ChainID(context.TODO())
			if err != nil {
				return true
			}
			e.Urls.Store(url, id)
			e.Client = dial
			return false
		})
		log.Debug("monitor eth client:", e.Client)
		time.Sleep(time.Second * MonitorBadRpcServersInterval)
	}
}

func (e *EthClient) SetAddress(urls []string) *ethclient.Client {
	for _, url := range urls {
		e.Urls.Store(url, 0)
	}
	return e.Client
}

func (e *EthClient) getCurrentBlockHeight(qid string) ([]byte, error) {
	if e.Client == nil {
		return nil, errors.New("no eth client")
	}
	number, err := e.Client.BlockNumber(context.TODO())
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(number - 1)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (e *EthClient) getCurrentBlockHash(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getVersion(qid string) ([]byte, error) {
	return json.Marshal("1.0.0")
}

func (e *EthClient) getNetworkId(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getBlockByHash(qid, hash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getBlockByHeight(qid string, height uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getBlockHash(qid string, height uint32) ([]byte, error) {
	block, err := e.Client.BlockByNumber(context.TODO(), big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	hex := block.Hash().Hex()
	hex = hex[2:]
	marshal, err := json.Marshal(hex)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (e *EthClient) getBlockHeightByTxHash(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getBlockTxHashesByHeight(qid string, height uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getRawTransaction(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContract(qid, contractAddress string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContractEvent(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContractEventByBlock(qid string, blockHeight uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContractEventByBlockAndAddress(qid string, blockHeight uint32, contractAddress string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContractEventByEventId(qid string, contractAddress string, address string, eventId uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContractEventByEventIdAndHeights(qid string, contractAddress string, address string,
	eventId, startHeight, endHeight uint32) ([]byte, error) {
	// TODO wangyu : implement me
	return nil, nil
}

func (e *EthClient) getStorage(qid, contractAddress string, key []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getMerkleProof(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getMemPoolTxState(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getMemPoolTxCount(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) sendRawTransaction(qid string, tx *types.Transaction, isPreExec bool) ([]byte, error) {
	// TODO wangyu: implement me
	txData := &ethTypes.DynamicFeeTx{
		ChainID:    nil,
		Nonce:      0,
		GasTipCap:  nil,
		GasFeeCap:  nil,
		Gas:        0,
		To:         nil,
		Value:      nil,
		Data:       nil,
		AccessList: nil,
		V:          nil,
		R:          nil,
		S:          nil,
	}
	txRes := ethTypes.NewTx(txData)
	return txRes.MarshalJSON()
}

func (e *EthClient) getGasPrice(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
