package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/core/types"
)

const MonitorBadRpcServersInterval = 30
const RoundRobin = 0
const MasterSlave = 1
const ServerScore = 2

//RpcClient for themis rpc api
type RpcClient struct {
	callMode        int
	randNum         int
	masterIndex     int
	rpcServersAddr  []string
	rpcServerStatus *sync.Map
	httpClient      *http.Client
}

//NewRpcClient return RpcClient instance
func NewRpcClient() *RpcClient {
	rpcClient := &RpcClient{
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
	go rpcClient.MonitorBadRpcServers()
	return rpcClient
}

func (this *RpcClient) MonitorBadRpcServers() {
	for {
		this.rpcServerStatus.Range(func(key, value interface{}) bool {
			rpcServerAddr := key.(string)
			score, err, timeout := this.getSysStatusScore(rpcServerAddr)
			if timeout {
				log.Errorf("[MonitorBadRpcServers] rpcServerAddr: %s service is not valid", rpcServerAddr)
				this.setServerScore(rpcServerAddr, 0)
			} else if err != nil {
				log.Errorf("[MonitorBadRpcServers] rpcServerAddr: %s getScore error: %s", rpcServerAddr, err.Error())
				this.setServerScore(rpcServerAddr, 0)
			} else {
				//log.Infof("[MonitorBadRpcServers] rpcServerAddr: %s score %d", rpcServerAddr, score)
				this.setServerScore(rpcServerAddr, score)
			}
			return true
		})
		time.Sleep(MonitorBadRpcServersInterval * time.Second)
	}
}

func (this *RpcClient) getServerScore(serverAddr string) (uint32, error) {
	if serverStatus, exist := this.rpcServerStatus.Load(serverAddr); exist {
		status := serverStatus.(uint32)
		return status, nil
	} else {
		return 0, fmt.Errorf("serverAddr is not exis in rpcServerStatus")
	}
}

func (this *RpcClient) setServerScore(serverAddr string, score uint32) {
	this.rpcServerStatus.Store(serverAddr, score)
}

//SetAddress set rpc server address. Simple http://localhost:20336
func (this *RpcClient) SetAddress(rpcAddrs []string) *RpcClient {
	this.rpcServersAddr = rpcAddrs
	for _, addr := range rpcAddrs {
		this.rpcServerStatus.Store(addr, uint32(1))
	}
	this.masterIndex = rand.Int() % len(this.rpcServersAddr)
	return this
}

func (this *RpcClient) GetAddress() []string {
	return this.rpcServersAddr
}

func (this *RpcClient) SetCallMode(callMode int) error {
	if 0 == len(this.rpcServersAddr) {
		return fmt.Errorf("[SetCallMode] error: RpcServerAddress is not set")
	}
	switch callMode {
	case MasterSlave:
		log.Info("[SetCallMode] Call mode is: MasterSlave")
	case RoundRobin:
		log.Info("[SetCallMode] Call mode is: RoundRobin")
	case ServerScore:
		log.Info("[SetCallMode] Call mode is: ServerCore")
	default:
		log.Error("[SetCallMode] Unknown callMode")
		return fmt.Errorf("[SetCallMode] Unknown callMode")
	}
	this.callMode = callMode
	return nil
}

//SetHttpClient set http client to RpcClient. In most cases SetHttpClient is not necessary
func (this *RpcClient) SetHttpClient(httpClient *http.Client) *RpcClient {
	this.httpClient = httpClient
	return this
}

//GetSysStatusScore return the status score of themis
func (this *RpcClient) getSysStatusScore(rpcServerAddr string) (uint32, error, bool) {
	data, err, b := this.sendRpcRequestToAddr(rpcServerAddr, "0", RPC_GET_SYSTEM_STATUS_SCORE, []interface{}{})
	if err != nil {
		return 0, err, b
	}
	if b {
		return 0, fmt.Errorf("server refused or timeout"), b
	}
	count, err := utils.GetUint32(data)
	if err != nil {
		return 0, err, false
	}
	return count, nil, false
}

//GetVersion return the version of themis
func (this *RpcClient) getVersionBySpecifiedRpcServer(rpcServerAddr string) ([]byte, error, bool) {
	return this.sendRpcRequestToAddr(rpcServerAddr, "0", RPC_GET_VERSION, []interface{}{})
}

//GetVersion return the version of themis
func (this *RpcClient) getVersion(qid string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_VERSION, []interface{}{})
}

func (this *RpcClient) getNetworkId(qid string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_NETWORK_ID, []interface{}{})
}

//GetBlockByHash return block with specified block hash in hex string code
func (this *RpcClient) getBlockByHash(qid, hash string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_BLOCK, []interface{}{hash})
}

//GetBlockByHeight return block by specified block height
func (this *RpcClient) getBlockByHeight(qid string, height uint32) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_BLOCK, []interface{}{height})
}

//GetBlockCount return the total block count of oniChain
func (this *RpcClient) getBlockCount(qid string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_BLOCK_COUNT, []interface{}{})
}

//GetCurrentBlockHeight return the current block height of oniChain
func (this *RpcClient) getCurrentBlockHeight(qid string) ([]byte, error) {
	data, err := this.getBlockCount(qid)
	if err != nil {
		return nil, err
	}
	count, err := utils.GetUint32(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(count - 1)
}

//GetCurrentBlockHash return the current block hash of oniChain
func (this *RpcClient) getCurrentBlockHash(qid string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_CURRENT_BLOCK_HASH, []interface{}{})
}

//GetBlockHash return block hash by block height
func (this *RpcClient) getBlockHash(qid string, height uint32) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_BLOCK_HASH, []interface{}{height})
}

//GetStorage return smart contract storage item.
//addr is smart contact address
//key is the key of value in smart contract
func (this *RpcClient) getStorage(qid, contractAddress string, key []byte) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_STORAGE, []interface{}{contractAddress, hex.EncodeToString(key)})
}

//GetSmartContractEvent return smart contract event execute by invoke transaction by hex string code
func (this *RpcClient) getSmartContractEvent(qid, txHash string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_SMART_CONTRACT_EVENT, []interface{}{txHash})
}

func (this *RpcClient) getSmartContractEventByBlock(qid string, blockHeight uint32) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_SMART_CONTRACT_EVENT, []interface{}{blockHeight})
}

func (this *RpcClient) getSmartContractEventByBlockAndAddress(qid string, blockHeight uint32, contractAddress string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_SMART_CONTRACT_EVENT, []interface{}{blockHeight, contractAddress})
}

func (this *RpcClient) getSmartContractEventByEventId(qid string, contractAddress string, address string, eventId uint32) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_SMART_CONTRACT_EVENT_BY_EVENT_ID, []interface{}{contractAddress, address, eventId})
}

func (this *RpcClient) getSmartContractEventByEventIdAndHeights(qid string, contractAddress string, address string, eventId, startHeight, endHeight uint32) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_SMART_CONTRACT_EVENT_BY_EVENT_ID_HEIGHTS, []interface{}{contractAddress, eventId, startHeight, endHeight, address})
}

//GetRawTransaction return transaction by transaction hash
func (this *RpcClient) getRawTransaction(qid, txHash string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_TRANSACTION, []interface{}{txHash})
}

//GetSmartContract return smart contract deployed in oniChain by specified smart contract address
func (this *RpcClient) getSmartContract(qid, contractAddress string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_SMART_CONTRACT, []interface{}{contractAddress})
}

//GetMerkleProof return the merkle proof whether tx is exist in ledger. Param txHash is in hex string code
func (this *RpcClient) getMerkleProof(qid, txHash string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_MERKLE_PROOF, []interface{}{txHash})
}

func (this *RpcClient) getMemPoolTxState(qid, txHash string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_MEM_POOL_TX_STATE, []interface{}{txHash})
}

func (this *RpcClient) getMemPoolTxCount(qid string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_MEM_POOL_TX_COUNT, []interface{}{})
}

func (this *RpcClient) getBlockHeightByTxHash(qid, txHash string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_BLOCK_HEIGHT_BY_TX_HASH, []interface{}{txHash})
}

func (this *RpcClient) getBlockTxHashesByHeight(qid string, height uint32) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_BLOCK_TX_HASH_BY_HEIGHT, []interface{}{height})
}

func (this *RpcClient) getGasPrice(qid string) ([]byte, error) {
	return this.sendRpcRequest(qid, RPC_GET_GAS_PRICE, []interface{}{})
}

func (this *RpcClient) sendRawTransaction(qid string, tx *types.Transaction, isPreExec bool) ([]byte, error) {
	sink := &common.ZeroCopySink{}
	tx.Serialization(sink)
	txData := hex.EncodeToString(sink.Bytes())
	params := []interface{}{txData}
	if isPreExec {
		params = append(params, 1)
	}
	return this.sendRpcRequest(qid, RPC_SEND_TRANSACTION, params)
}

//sendRpcRequest send Rpc request to oniChain
func (this *RpcClient) sendRpcRequest(qid, method string, params []interface{}) ([]byte, error) {
	rpcAddr := this.getNextRpcAddress()
	if rpcAddr == "" {
		return nil, fmt.Errorf("[sendRpcRequest] no usable rpc server")
	}
	resp, err, badServer := this.sendRpcRequestToAddr(rpcAddr, qid, method, params)
	if err != nil {
		if badServer {
			this.setServerScore(rpcAddr, 0)
			nextRpcAddr := this.getNextRpcAddress()
			if nextRpcAddr == "" {
				return nil, fmt.Errorf("[sendRpcRequest] no usable rpc server")
			}
			log.Warnf("[sendRpcRequest] http post request rpcAddr: %s method: %s error: %s, getNextRpcAddress %s Retry",
				rpcAddr, method, err, nextRpcAddr)

			resp, err, badServer = this.sendRpcRequestToAddr(nextRpcAddr, qid, method, params)
			if err != nil {
				if badServer {
					this.setServerScore(nextRpcAddr, 0)
				}
				err2 := fmt.Errorf("[sendRpcRequest] http post request rpcAddr: %s method: %s error: %s", nextRpcAddr,
					method, err.Error())
				log.Error(err2.Error())
				return nil, err2
			}
		}
	}
	return resp, err
}

//sendRpcRequest send Rpc request to oniChain
func (this *RpcClient) sendRpcRequestToAddr(rpcAddr string, qid, method string, params []interface{}) ([]byte, error, bool) {
	rpcReq := &JsonRpcRequest{
		Version: JSON_RPC_VERSION,
		Id:      qid,
		Method:  method,
		Params:  params,
	}
	data, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, fmt.Errorf("JsonRpcRequest json.Marsha error:%s", err), false
	}

	done := make(chan bool)
	var resp *http.Response
	go func() {
		resp, err = this.httpClient.Post(rpcAddr, "application/json", bytes.NewReader(data))
		done <- true
	}()

	select {
	case <-done:
		defer close(done)
		if err != nil {
			errInfo := err.Error()
			if strings.Contains(errInfo, "connect: connection refused") {
				return nil, err, true
			} else {
				return nil, err, false
			}
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read rpc response body error:%s", err.Error()), false
		}

		rpcRsp := &JsonRpcResponse{}
		err = json.Unmarshal(body, rpcRsp)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal JsonRpcResponse:%s error:%s", body, err.Error()), false
		}
		if rpcRsp.Error != 0 {
			return nil, fmt.Errorf("JsonRpcResponse error code:%d desc:%s result:%s", rpcRsp.Error, rpcRsp.Desc,
				rpcRsp.Result), false
		}
		return rpcRsp.Result, nil, false
	case <-time.After(20 * time.Second):
		err = fmt.Errorf("[sendRpcRequestToAddr] RpcAddr(%s) timeout", rpcAddr)
		return nil, err, true
	}
}

func (this *RpcClient) getNextRpcAddress() string {
	var rpcAddr string
	rpcSrvAddrsLen := len(this.rpcServersAddr)
	if this.callMode == MasterSlave {
		masterIndex := this.masterIndex
		for {
			rpcAddr = this.rpcServersAddr[masterIndex]
			score, err := this.getServerScore(rpcAddr)
			if err == nil && score != 0 {
				this.masterIndex = masterIndex
				//fmt.Printf("Server: %s, score: %d\n", rpcAddr, score)
				return rpcAddr
			} else {
				masterIndex = (masterIndex + 1) % rpcSrvAddrsLen
				if masterIndex == this.masterIndex {
					log.Errorf("[getNextRpcAddress] no usable server address")
					return rpcAddr
				}
				log.Warnf("[getNextRpcAddress] %s is not valid, try %s", this.rpcServersAddr[this.masterIndex],
					this.rpcServersAddr[masterIndex])
			}
		}
	} else if this.callMode == RoundRobin {
		for i := 0; i < rpcSrvAddrsLen; i++ {
			index := this.randNum % rpcSrvAddrsLen
			rpcAddr = this.rpcServersAddr[index]

			this.randNum++
			score, err := this.getServerScore(rpcAddr)
			if err == nil && score != 0 {
				//fmt.Printf("Server: %s, score: %d\n", rpcAddr, score)
				return rpcAddr
			}

			if this.randNum >= 999999999 {
				this.randNum = 0
			}
		}
		index := rand.Int() % rpcSrvAddrsLen
		rpcAddr = this.rpcServersAddr[index]
	} else if this.callMode == ServerScore {
		var bestSrv string
		var maxScore uint32

		for i := 0; i < rpcSrvAddrsLen; i++ {
			tmpAddr := this.rpcServersAddr[i]
			score, err := this.getServerScore(tmpAddr)
			//fmt.Printf("Server: %s, score: %d\n", tmpAddr, score)
			if err == nil && score > maxScore {
				maxScore = score
				bestSrv = tmpAddr
			}
		}
		if len(bestSrv) == 0 && len(this.rpcServersAddr) > 0 {
			bestSrv = this.rpcServersAddr[0]
		}
		//fmt.Printf("best server addr: %s\n", bestSrv)
		rpcAddr = bestSrv
	}
	return rpcAddr
}
