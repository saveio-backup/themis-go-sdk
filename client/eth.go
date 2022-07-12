package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/saveio/themis/core/types"
)

type EthClient struct {
	Client *ethclient.Client
}

func (e EthClient) getCurrentBlockHeight(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getCurrentBlockHash(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getVersion(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getNetworkId(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getBlockByHash(qid, hash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getBlockByHeight(qid string, height uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getBlockHash(qid string, height uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getBlockHeightByTxHash(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getBlockTxHashesByHeight(qid string, height uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getRawTransaction(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getSmartContract(qid, contractAddress string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getSmartContractEvent(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getSmartContractEventByBlock(qid string, blockHeight uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getSmartContractEventByBlockAndAddress(qid string, blockHeight uint32, contractAddress string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getSmartContractEventByEventId(qid string, contractAddress string, address string, eventId uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getSmartContractEventByEventIdAndHeights(qid string, contractAddress string, address string, eventId, startHeight, endHeight uint32) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getStorage(qid, contractAddress string, key []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getMerkleProof(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getMemPoolTxState(qid, txHash string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getMemPoolTxCount(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) sendRawTransaction(qid string, tx *types.Transaction, isPreExec bool) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) getGasPrice(qid string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (e EthClient) SetAddress(rpcAddrs []string) *EthClient {
	//TODO implement me
	panic("implement me")
}
