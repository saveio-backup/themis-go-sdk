package client

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"sync"
	"time"

	ethCom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/core/types"
)

type EthClient struct {
	Client *ethclient.Client
	Urls   *sync.Map // string => int
	dial   chan bool
}

func NewEthClient() *EthClient {
	client := &EthClient{
		Urls:   new(sync.Map),
		Client: nil,
		dial:   make(chan bool),
	}
	go client.MonitorBasEthServer()
	return client
}

func (e *EthClient) DialClient() {
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
	log.Debugf("eth client dialed")
}

func (e *EthClient) MonitorBasEthServer() {
	ticker := time.NewTicker(time.Second * MonitorBadRpcServersInterval)
	for {
		select {
		case <-ticker.C:
			e.DialClient()
		case <-e.dial:
			e.DialClient()
		}
	}
}

func (e *EthClient) SetAddress(urls []string) *ethclient.Client {
	for _, url := range urls {
		e.Urls.Store(url, 0)
	}
	e.dial <- true
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
	marshal, err := json.Marshal(number)
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
	header, err := e.Client.HeaderByNumber(context.Background(), big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return json.Marshal(header)
}

func (e *EthClient) getBlockHash(qid string, height uint32) ([]byte, error) {
	header, err := e.Client.HeaderByNumber(context.Background(), big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	hex := header.Hash().String()

	marshal, err := json.Marshal(hex[2:])
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (e *EthClient) getBlockHeightByTxHash(qid, txHash string) ([]byte, error) {
	receipt, err := e.Client.TransactionReceipt(context.Background(), ethCom.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	return json.Marshal(uint32(receipt.BlockNumber.Uint64()))
}

func (e *EthClient) getBlockTxHashesByHeight(qid string, height uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getRawTransaction(qid, txHash string) ([]byte, error) {
	tx, _, err := e.Client.TransactionByHash(context.Background(), ethCom.HexToHash(txHash))
	if err != nil {
		return nil, err
	}
	return tx.MarshalBinary()
}

func (e *EthClient) getSmartContract(qid, contractAddress string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e *EthClient) getSmartContractEvent(qid, txHash string) ([]byte, error) {
	log.Errorf("getSmartContractEvent not implemented")
	events := make([]*sdkcom.SmartContactEvent, 0)
	return json.Marshal(events)
}

func (e *EthClient) getSmartContractEventByBlock(qid string, blockHeight uint32) ([]byte, error) {
	log.Errorf("getSmartContractEventByBlock not implemented")
	events := make([]*sdkcom.SmartContactEvent, 0)
	return json.Marshal(events)
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
	log.Errorf("sendRawTransaction not implemented")
	if isPreExec {
		res := &sdkcom.PreExecResult{
			Result: &sdkcom.ResultItem{},
		}
		marshal, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	} else {
		res := common.Uint256{}
		marshal, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
}

func (e *EthClient) getGasPrice(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
