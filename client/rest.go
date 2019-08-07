package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"sync"

	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/core/types"
	"github.com/saveio/themis/common/log"
	"math/rand"
)

const MonitorBadRestServersInterval = 30

//RestClient for oniChain rest api
type RestClient struct {
	randNum          int
	restServersAddr  []string
	restServerStatus *sync.Map
	httpClient       *http.Client
}

//NewRpcClient return RpcClient instance
func NewRestClient() *RestClient {
	restClient := &RestClient{
		restServerStatus: new(sync.Map),
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false, //enable keepalive
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
			},
			Timeout: time.Second * 300, //timeout for http response
		},
	}
	go restClient.MonitorBadRestServers()
	return restClient
}


func (this *RestClient) MonitorBadRestServers () {
	var err error
	for {
		time.Sleep(MonitorBadRestServersInterval * time.Second)
		this.restServerStatus.Range(func(key, value interface{}) bool {
			restServerAddr := key.(string)
			serverStatus := value.(bool)
			if !serverStatus {
				_, err = this.getVersionBySpecifiedRestServer(restServerAddr)
				if err == nil {
					log.Infof("[MonitorBadRestServers] restServerAddr: %s service resume", restServerAddr)
					this.setServerStatus(restServerAddr, true)
				} else {
					log.Errorf("[MonitorBadRestServers] restServerAddr: %s service is not valid", restServerAddr)
				}
			}
			return true
		})
	}
}

func (this *RestClient) getServerStatus(serverAddr string) (bool, error) {
	if serverStatus, exist := this.restServerStatus.Load(serverAddr); exist {
		status := serverStatus.(bool)
		return status, nil
	} else {
		return false, fmt.Errorf("serverAddr is not exis in restServerStatus")
	}
}

func (this *RestClient) setServerStatus(serverAddr string, status bool) {
	this.restServerStatus.Store(serverAddr, status)
}

func (this *RestClient) getVersionBySpecifiedRestServer(addr string) ([]byte, error) {
	reqPath := GET_VERSION
	return this.sendRestGetRequestToAddr(addr, reqPath)
}


//SetAddress set rest server address. Simple http://localhost:20334
func (this *RestClient) SetAddress(restAddrs []string) *RestClient {
	this.restServersAddr = restAddrs
	for _, addr := range restAddrs {
		this.restServerStatus.Store(addr, true)
	}
	return this
}

//SetHttpClient set rest client to RestClient. In most cases SetHttpClient is not necessary
func (this *RestClient) SetHttpClient(httpClient *http.Client) *RestClient {
	this.httpClient = httpClient
	return this
}

func (this *RestClient) getVersion(qid string) ([]byte, error) {
	reqPath := GET_VERSION
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getNetworkId(qid string) ([]byte, error) {
	reqPath := GET_NETWORK_ID
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getBlockByHash(qid, hash string) ([]byte, error) {
	reqPath := GET_BLK_BY_HASH + hash
	reqValues := &url.Values{}
	reqValues.Add("raw", "1")
	return this.sendRestGetRequest(reqPath, reqValues)
}

func (this *RestClient) getBlockByHeight(qid string, height uint32) ([]byte, error) {
	reqPath := fmt.Sprintf("%s%d", GET_BLK_BY_HEIGHT, height)
	reqValues := &url.Values{}
	reqValues.Add("raw", "1")
	return this.sendRestGetRequest(reqPath, reqValues)
}

func (this *RestClient) getCurrentBlockHeight(qid string) ([]byte, error) {
	reqPath := GET_BLK_HEIGHT
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getCurrentBlockHash(qid string) ([]byte, error) {
	data, err := this.getCurrentBlockHeight(qid)
	if err != nil {
		return nil, err
	}
	height, err := utils.GetUint32(data)
	if err != nil {
		return nil, err
	}
	return this.getBlockHash(qid, height)
}

func (this *RestClient) getBlockHash(qid string, height uint32) ([]byte, error) {
	reqPath := fmt.Sprintf("%s%d", GET_BLK_HASH, height)
	return this.sendRestGetRequest(reqPath)
}

//GetRawTransaction return transaction by transaction hash in hex string code
func (this *RestClient) getRawTransaction(qid, txHash string) ([]byte, error) {
	reqPath := GET_TX + txHash
	reqValues := &url.Values{}
	reqValues.Add("raw", "1")
	return this.sendRestGetRequest(reqPath, reqValues)
}

func (this *RestClient) getStorage(qid, contractAddress string, key []byte) ([]byte, error) {
	reqPath := GET_STORAGE + contractAddress + "/" + hex.EncodeToString(key)
	return this.sendRestGetRequest(reqPath)
}

//GetSmartContractEvent return smart contract event execute by invoke transaction by hex string code
func (this *RestClient) getSmartContractEvent(qid, txHash string) ([]byte, error) {
	reqPath := GET_SMTCOCE_EVTS + txHash
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getSmartContractEventByBlock(qid string, blockHeight uint32) ([]byte, error) {
	reqPath := fmt.Sprintf("%s%d", GET_SMTCOCE_EVT_TXS, blockHeight)
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getSmartContractEventByBlockAndAddress(qid string, blockHeight uint32, contractAddress string) ([]byte, error) {
	reqPath := fmt.Sprintf("%s%d", GET_SMTCOCE_EVT_ADDR, blockHeight) + "/" + contractAddress
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getSmartContractEventByEventId(qid string, contractAddress string, address string, eventId uint32) ([]byte, error) {
	reqPath := fmt.Sprintf("%s%d", GET_SMTCOCE_EVT_ID+"/"+contractAddress+"/"+address+"/", eventId)
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getSmartContract(qid, contractAddress string) ([]byte, error) {
	reqPath := GET_CONTRACT_STATE + contractAddress
	reqValues := &url.Values{}
	reqValues.Add("raw", "1")
	return this.sendRestGetRequest(reqPath, reqValues)
}

func (this RestClient) getMerkleProof(qid, txHash string) ([]byte, error) {
	reqPath := GET_MERKLE_PROOF + txHash
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getMemPoolTxState(qid, txHash string) ([]byte, error) {
	reqPath := GET_MEMPOOL_TXSTATE + txHash
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getMemPoolTxCount(qid string) ([]byte, error) {
	reqPath := GET_MEMPOOL_TXCOUNT
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getBlockHeightByTxHash(qid, txHash string) ([]byte, error) {
	reqPath := GET_BLK_HGT_BY_TXHASH + txHash
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getBlockTxHashesByHeight(qid string, height uint32) ([]byte, error) {
	reqPath := fmt.Sprintf("%s%d", GET_BLK_TXS_BY_HEIGHT, height)
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) getGasPrice(qid string) ([]byte, error) {
	reqPath := GET_GAS_PRICE
	return this.sendRestGetRequest(reqPath)
}

func (this *RestClient) sendRawTransaction(qid string, tx *types.Transaction, isPreExec bool) ([]byte, error) {
	reqPath := POST_RAW_TX
	var buffer bytes.Buffer
	err := tx.Serialize(&buffer)
	if err != nil {
		return nil, fmt.Errorf("Serialize error:%s", err.Error())
	}
	var reqValues *url.Values
	if isPreExec {
		reqValues = &url.Values{}
		reqValues.Add("preExec", "1")
	}
	return this.sendRestPostRequest(buffer.Bytes(), reqPath, reqValues)
}




func (this *RestClient) getRequestUrl(restAddr string, reqPath string, values ...*url.Values) (string, error) {
	if !strings.HasPrefix(restAddr, "http") {
		restAddr = "http://" + restAddr
	}
	reqUrl, err := new(url.URL).Parse(restAddr)
	if err != nil {
		return "", fmt.Errorf("Parse address:%s error:%s", restAddr, err.Error())
	}
	reqUrl.Path = reqPath
	if len(values) > 0 {
		first := values[0]
		if first != nil {
			reqUrl.RawQuery = first.Encode()
		}
	}
	return reqUrl.String(), nil
}

func (this *RestClient) sendRestGetRequest(reqPath string, values ...*url.Values) ([]byte, error) {
	restAddr := this.getNextRestAddress()
	if restAddr == "" {
		return nil, fmt.Errorf("Restful server address is nil")
	}
	reqUrl, err := this.getRequestUrl(restAddr, reqPath, values...)
	if err != nil {
		return nil, err
	}
	resp, err := this.httpClient.Get(reqUrl)
	if err != nil {
		this.setServerStatus(restAddr, false)
		log.Errorf("[sendRestGetRequest] http post request reqPath :%s error:%s", reqPath, err)
		nextRestAddr := this.getNextRestAddress()
		if nextRestAddr == "" {
			return nil, fmt.Errorf("Restful server address is nil")
		}
		newReqUrl, err := this.getRequestUrl(nextRestAddr, reqPath, values...)
		if err != nil {
			return nil, err
		}
		resp, err = this.httpClient.Get(newReqUrl)
		if err != nil {
			this.setServerStatus(nextRestAddr, false)
			log.Errorf("[sendRestGetRequest] http post request reqPath :%s error:%s", reqPath, err)
			return nil, fmt.Errorf("[sendRestGetRequest] http get request reqPath :%s error:%s", reqPath, err)
		}
	}
	defer resp.Body.Close()
	return this.dealRestResponse(resp.Body)
}

func (this *RestClient) sendRestPostRequest(data []byte, reqPath string, values ...*url.Values) ([]byte, error) {
	restAddr := this.getNextRestAddress()
	if restAddr == "" {
		return nil, fmt.Errorf("Restful server address is nil")
	}
	reqUrl, err := this.getRequestUrl(restAddr, reqPath, values...)
	if err != nil {
		return nil, err
	}

	restReq := &RestfulReq{
		Action:  ACTION_SEND_RAW_TRANSACTION,
		Version: REST_VERSION,
		Data:    hex.EncodeToString(data),
	}
	reqData, err := json.Marshal(restReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error:%s", err)
	}

	resp, err := this.httpClient.Post(reqUrl, "application/json", bytes.NewReader(reqData))
	if err != nil {
		this.setServerStatus(restAddr, false)
		log.Errorf("[sendRestPostRequest] http post request reqPath :%s error:%s", reqPath, err)
		nextRestAddr := this.getNextRestAddress()
		if nextRestAddr == "" {
			return nil, fmt.Errorf("Restful server address is nil")
		}
		newReqUrl, err := this.getRequestUrl(nextRestAddr, reqPath, values...)
		if err != nil {
			return nil, err
		}
		resp, err = this.httpClient.Get(newReqUrl)
		if err != nil {
			this.setServerStatus(nextRestAddr, false)
			log.Errorf("[sendRestPostRequest] http post request reqPath :%s error:%s", reqPath, err)
			return nil, fmt.Errorf("[sendRestPostRequest] http post request reqPath :%s error:%s", reqPath, err)
		}
	}
	defer resp.Body.Close()
	return this.dealRestResponse(resp.Body)
}

func (this *RestClient) sendRestGetRequestToAddr(addr string, reqPath string, values ...*url.Values) ([]byte, error) {
	reqUrl, err := this.getRequestUrl(addr, reqPath, values...)
	if err != nil {
		return nil, err
	}
	resp, err := this.httpClient.Get(reqUrl)
	if err != nil {
		log.Errorf("[sendRestGetRequestToAddr] http post request reqUrl :%s error:%s", reqUrl, err.Error())
		return nil, fmt.Errorf("send http get request error:%s", err)
	}
	defer resp.Body.Close()
	return this.dealRestResponse(resp.Body)
}

func (this *RestClient) dealRestResponse(body io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("read http body error:%s", err)
	}
	restRsp := &RestfulResp{}
	err = json.Unmarshal(data, restRsp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal RestfulResp:%s error:%s", body, err)
	}
	if restRsp.Error != 0 {
		return nil, fmt.Errorf("sendRestRequest error code:%d desc:%s result:%s", restRsp.Error, restRsp.Desc, restRsp.Result)
	}
	return restRsp.Result, nil
}

func (this *RestClient) getNextRestAddress() string {
	var restAddr string
	restServersAddrLen := len(this.restServersAddr)
	for i := 0; i < restServersAddrLen; i++ {
		index := this.randNum % restServersAddrLen
		restAddr = this.restServersAddr[index]

		this.randNum++
		ok, exist := this.restServerStatus.Load(restAddr)
		if exist && ok.(bool) {
			log.Info("[getNextRestAddress] Return: ", restAddr)
			return restAddr
		}

		if this.randNum >= 1000000000 {
			this.randNum = 0
		}
	}
	index := rand.Int() % restServersAddrLen
	restAddr = this.restServersAddr[index]

	log.Info("[getNextRestAddress] Return RandRestAddr: ", restAddr)
	return restAddr
}
