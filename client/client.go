package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/core/types"
)

type ClientMgr struct {
	rpc       *RpcClient  //Rpc client used the rpc api of oniChain
	rest      *RestClient //Rest client used the rest api of oniChain
	ws        *WSClient   //Web socket client used the web socket api of oniChain
	defClient ChainClient
	qid       uint64
	ethClient *EthClient // Rpc client used the ethclient of Ethereum
}

func (this *ClientMgr) NewRpcClient() *RpcClient {
	this.rpc = NewRpcClient()
	return this.rpc
}

func (this *ClientMgr) NewRpcClientWithAddrs(addrs []string) *RpcClient {
	this.rpc = &RpcClient{
		callMode:        ServerScore,
		rpcServerStatus: new(sync.Map),
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxConnsPerHost:       2,
				MaxIdleConnsPerHost:   12,
				DisableKeepAlives:     false, //enable keepalive
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 100,
				ExpectContinueTimeout: 20,
			},
			Timeout: time.Second * 50, //timeout for http response
		},
	}
	this.rpc.SetAddress(addrs)
	return this.rpc
}

func (this *ClientMgr) GetRpcClient() *RpcClient {
	return this.rpc
}

func (this *ClientMgr) NewRestClient() *RestClient {
	this.rest = NewRestClient()
	return this.rest
}

func (this *ClientMgr) GetRestClient() *RestClient {
	return this.rest
}

func (this *ClientMgr) NewWebSocketClient() *WSClient {
	wsClient := NewWSClient()
	this.ws = wsClient
	return wsClient
}

func (this *ClientMgr) GetWebSocketClient() *WSClient {
	return this.ws
}

func (this *ClientMgr) NewEthClient() *EthClient {
	this.ethClient = NewEthClient()
	return this.ethClient
}

func (this *ClientMgr) GetEthClient() *EthClient {
	return this.ethClient
}

func (this *ClientMgr) SetDefaultClient(client ChainClient) {
	this.defClient = client
}

func (this *ClientMgr) GetCurrentBlockHeight() (uint32, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getCurrentBlockHeight(this.getNextQid())
	if err != nil {
		return 0, err
	}
	return utils.GetUint32(data)
}

func (this *ClientMgr) GetCurrentBlockHash() (common.Uint256, error) {
	client := this.getClient()
	if client == nil {
		return common.UINT256_EMPTY, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getCurrentBlockHash(this.getNextQid())
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return utils.GetUint256(data)
}

func (this *ClientMgr) GetBlockByHeight(height uint32) (*types.Block, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getBlockByHeight(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetBlock(data)
}

func (this *ClientMgr) GetBlockByHash(blockHash string) (*types.Block, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getBlockByHash(this.getNextQid(), blockHash)
	if err != nil {
		return nil, err
	}
	return utils.GetBlock(data)
}

func (this *ClientMgr) GetRawTransaction(txHash string) ([]byte, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	return client.getRawTransaction(this.getNextQid(), txHash)
}

func (this *ClientMgr) GetTransaction(txHash string) (*types.Transaction, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getRawTransaction(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetTransaction(data)
}

func (this *ClientMgr) GetBlockHash(height uint32) (common.Uint256, error) {
	client := this.getClient()
	if client == nil {
		return common.UINT256_EMPTY, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getBlockHash(this.getNextQid(), height)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return utils.GetUint256(data)
}

func (this *ClientMgr) GetBlockHeightByTxHash(txHash string) (uint32, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getBlockHeightByTxHash(this.getNextQid(), txHash)
	if err != nil {
		return 0, err
	}
	return utils.GetUint32(data)
}

func (this *ClientMgr) GetBlockTxHashesByHeight(height uint32) (*sdkcom.BlockTxHashes, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getBlockTxHashesByHeight(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetBlockTxHashes(data)
}

func (this *ClientMgr) GetStorage(contractAddress string, key []byte) ([]byte, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getStorage(this.getNextQid(), contractAddress, key)
	if err != nil {
		return nil, err
	}
	return utils.GetStorage(data)
}

func (this *ClientMgr) GetSmartContract(contractAddress string) (*sdkcom.SmartContract, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContract(this.getNextQid(), contractAddress)
	if err != nil {
		return nil, err
	}
	deployCode, err := utils.GetSmartContract(data)
	if err != nil {
		return nil, err
	}
	sm := sdkcom.SmartContract(*deployCode)
	return &sm, nil
}

func (this *ClientMgr) GetSmartContractEvent(txHash string) (*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContractEvent(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContractEvent(data)
}

func (this *ClientMgr) GetSmartContractEventsByBlock(height uint32) ([]*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContractEventByBlock(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContractEvents(data)
}

func (this *ClientMgr) GetSmartContractEventByBlock(height uint32) (*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContractEventByBlock(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContractEvent(data)
}

func (this *ClientMgr) GetSmartContractEventsByBlockAndAddress(height uint32, contractAddress string) ([]interface{}, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContractEventByBlockAndAddress(this.getNextQid(), height, contractAddress)
	if err != nil {
		return nil, err
	}

	events := make([]interface{}, 0)
	err = json.Unmarshal(data, &events)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal SmartContactEvent:%s error:%s", data, err)
	}
	return events, nil

}

func (this *ClientMgr) GetSmartContractEventByEventId(contractAddress string, address string, eventId uint32) ([]*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContractEventByEventId(this.getNextQid(), contractAddress, address, eventId)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContractEvents(data)

}

func (this *ClientMgr) GetSmartContractEventByEventIdAndHeights(contractAddress string, address string, eventId, startHeight, endHeight uint32) ([]*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getSmartContractEventByEventIdAndHeights(this.getNextQid(), contractAddress, address, eventId, startHeight, endHeight)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContractEvents(data)
}

func (this *ClientMgr) GetMerkleProof(txHash string) (*sdkcom.MerkleProof, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getMerkleProof(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetMerkleProof(data)
}

func (this *ClientMgr) GetMemPoolTxState(txHash string) (*sdkcom.MemPoolTxState, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getMemPoolTxState(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetMemPoolTxState(data)
}

func (this *ClientMgr) GetMemPoolTxCount() (*sdkcom.MemPoolTxCount, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getMemPoolTxCount(this.getNextQid())
	if err != nil {
		return nil, err
	}
	return utils.GetMemPoolTxCount(data)
}

func (this *ClientMgr) GetVersion() (string, error) {
	client := this.getClient()
	if client == nil {
		return "", fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getVersion(this.getNextQid())
	if err != nil {
		return "", err
	}
	return utils.GetVersion(data)
}

func (this *ClientMgr) GetNetworkId() (uint32, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getNetworkId(this.getNextQid())
	if err != nil {
		return 0, err
	}
	return utils.GetUint32(data)
}

func (this *ClientMgr) GetGasPrice() (uint64, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of oniChain")
	}
	data, err := client.getGasPrice(this.getNextQid())
	if err != nil {
		return 0, err
	}
	return utils.GetUint64(data)
}

func (this *ClientMgr) SendTransaction(mutTx *types.MutableTransaction) (common.Uint256, error) {
	client := this.getClient()
	if client == nil {
		return common.UINT256_EMPTY, fmt.Errorf("don't have available client of oniChain")
	}
	tx, err := mutTx.IntoImmutable()
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	data, err := client.sendRawTransaction(this.getNextQid(), tx, false)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return utils.GetUint256(data)
}

func (this *ClientMgr) PreExecTransaction(mutTx *types.MutableTransaction) (*sdkcom.PreExecResult, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of oniChain")
	}
	tx, err := mutTx.IntoImmutable()
	if err != nil {
		return nil, err
	}
	data, err := client.sendRawTransaction(this.getNextQid(), tx, true)
	if err != nil {
		return nil, err
	}
	preResult := &sdkcom.PreExecResult{}
	err = json.Unmarshal(data, &preResult)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal PreExecResult:%s error:%s", data, err)
	}
	return preResult, nil
}

//WaitForGenerateBlock Wait oniChain generate block. Default wait 2 blocks.
//return timeout error when there is no block generate in some time.
func (this *ClientMgr) WaitForGenerateBlock(timeout time.Duration, blockCount ...uint32) (bool, error) {
	count := uint32(2)
	if len(blockCount) > 0 && blockCount[0] > 0 {
		count = blockCount[0]
	}
	blockHeight, err := this.GetCurrentBlockHeight()
	if err != nil {
		return false, fmt.Errorf("GetCurrentBlockHeight error:%s", err)
	}
	secs := int(timeout / time.Second)
	if secs <= 0 {
		secs = 1
	}
	for i := 0; i < secs; i++ {
		time.Sleep(time.Second)
		curBlockHeigh, err := this.GetCurrentBlockHeight()
		if err != nil {
			continue
		}
		if curBlockHeigh-blockHeight >= count {
			return true, nil
		}
	}
	return false, fmt.Errorf("timeout after %d (s)", secs)
}

// PollForTxConfirmed Poll tx for confirmation. This is a thread-block method
func (this *ClientMgr) PollForTxConfirmed(timeout time.Duration, txHash []byte) (bool, error) {
	if len(txHash) == 0 {
		return false, fmt.Errorf("txHash is empty")
	}
	txHashStr := hex.EncodeToString(common.ToArrayReverse(txHash))
	interval := time.Duration(sdkcom.POLL_TX_INTERVAL) * time.Second
	secs := int(timeout / interval)
	if secs <= 0 {
		secs = 1
	}
	log.Debugf("txHashStr: %s", txHashStr)
	for i := 0; i < secs; i++ {
		time.Sleep(interval)
		ret, err := this.GetBlockHeightByTxHash(txHashStr)
		if err != nil || ret == 0 {
			continue
		}
		return true, nil
	}
	return false, fmt.Errorf("timeout after %d (s)", secs)
}

// PollForTxConfirmedHeight Poll tx for confirmation. This is a thread-block method
func (this *ClientMgr) PollForTxConfirmedHeight(timeout time.Duration, txHash []byte) (uint32, error) {
	if len(txHash) == 0 {
		return 0, fmt.Errorf("txHash is empty")
	}
	txHashStr := hex.EncodeToString(common.ToArrayReverse(txHash))
	interval := time.Duration(sdkcom.POLL_TX_INTERVAL) * time.Second
	secs := int(timeout / interval)
	if secs <= 0 {
		secs = 1
	}
	log.Debugf("txHashStr: %s", txHashStr)
	for i := 0; i < secs; i++ {
		time.Sleep(interval)
		ret, err := this.GetBlockHeightByTxHash(txHashStr)
		if err != nil || ret == 0 {
			continue
		}
		return ret, nil
	}
	return 0, fmt.Errorf("timeout after %d (s)", secs)
}

func (this *ClientMgr) getClient() ChainClient {
	if this.defClient != nil {
		return this.defClient
	}
	if this.rpc != nil {
		return this.rpc
	}
	if this.rest != nil {
		return this.rest
	}
	if this.ws != nil {
		return this.ws
	}
	if this.ethClient != nil {
		return this.ethClient
	}
	return nil
}

func (this *ClientMgr) GetClientMgr() *ClientMgr {
	return this
}

func (this *ClientMgr) getNextQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&this.qid, 1))
}
