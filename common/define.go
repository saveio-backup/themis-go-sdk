package common

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/saveio/themis/common"
	"github.com/saveio/themis/core/payload"
)

var (
	VERSION_TRANSACTION = byte(0)
	GAS_PRICE           = uint64(1)
	GAS_LIMIT = uint64(1000000)
)

type SmartContract payload.DeployCode

type PreExecResult struct {
	State  byte
	Gas    uint64
	Result *ResultItem
}

func (this *PreExecResult) UnmarshalJSON(data []byte) (err error) {
	var state byte
	var gas uint64
	var resultItem *ResultItem
	defer func() {
		if err == nil {
			this.State = state
			this.Gas = gas
			this.Result = resultItem
		}
	}()

	objects := make(map[string]interface{})
	err = json.Unmarshal(data, &objects)
	if err != nil {
		return err
	}
	stateField, ok := objects["State"].(float64)
	if !ok {
		err = fmt.Errorf("Parse State field failed, type error")
		return
	}
	state = byte(stateField)

	gasField, ok := objects["Gas"].(float64)
	if !ok {
		err = fmt.Errorf("Parse Gas field failed, type error")
		return
	}
	gas = uint64(gasField)
	resultField, ok := objects["Result"]
	if !ok {
		return nil
	}
	resultItem = &ResultItem{}
	value, ok := resultField.(string)
	if ok {
		resultItem.value = value
		return nil
	}
	values, ok := resultField.([]interface{})
	if !ok {
		err = fmt.Errorf("Parse Result field, type error")
		return
	}
	resultItem.values = values
	return nil
}

type ResultItem struct {
	value  string
	values []interface{}
}

func (this *ResultItem) ToArray() ([]*ResultItem, error) {
	if this.values == nil {
		return nil, fmt.Errorf("type error")
	}
	items := make([]*ResultItem, 0)
	for _, res := range this.values {
		item := &ResultItem{}
		value, ok := res.(string)
		if ok {
			item.value = value
			items = append(items, item)
			continue
		}
		values, ok := res.([]interface{})
		if !ok {
			return nil, fmt.Errorf("parse items:%v failed, type error", res)
		}
		item.values = values
		items = append(items, item)
	}
	return items, nil
}

func (this ResultItem) ToBool() (bool, error) {
	if this.values != nil {
		return false, fmt.Errorf("type error")
	}
	return this.value == "01", nil
}

func (this ResultItem) ToInteger() (*big.Int, error) {
	data, err := this.ToByteArray()
	if err != nil {
		return nil, err
	}
	return common.BigIntFromNeoBytes(data), nil
}

func (this ResultItem) ToByteArray() ([]byte, error) {
	if this.values != nil {
		return nil, fmt.Errorf("type error")
	}
	return hex.DecodeString(this.value)
}

func (this ResultItem) ToString() (string, error) {
	data, err := this.ToByteArray()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//SmartContactEvent object for event of transaction
type SmartContactEvent struct {
	TxHash      string
	State       byte
	GasConsumed uint64
	Notify      []*NotifyEventInfo
}

type NotifyEventInfo struct {
	ContractAddress string
	EventIdentifier uint32
	Addresses       []common.Address // addresses involved in the notify
	States          interface{}
}

func (this *NotifyEventInfo) UnmarshalJSON(data []byte) error {
	type evtInfo struct {
		ContractAddress string
		States          json.RawMessage
	}
	info := &evtInfo{}
	err := json.Unmarshal(data, info)
	if err != nil {
		return err
	}
	this.ContractAddress = info.ContractAddress

	dec := json.NewDecoder(bytes.NewReader(info.States))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if delim, ok := token.(json.Delim); !ok || delim.String() != "[" {
		return this.originUnmarshal(info.States)
	}
	notifyMethod, err := dec.Token()
	if err != nil {
		return this.originUnmarshal(info.States)
	}
	if notifyMethod != "transfer" {
		return this.originUnmarshal(info.States)
	}
	transferFrom, err := dec.Token()
	if err != nil {
		return this.originUnmarshal(info.States)
	}
	transferTo, err := dec.Token()
	if err != nil {
		return this.originUnmarshal(info.States)
	}
	//using uint64 to decode, avoid precision lost decode by float64
	transferAmount := uint64(0)
	err = dec.Decode(&transferAmount)
	if err != nil {
		return this.originUnmarshal(info.States)
	}
	blockHeight := uint64(0)
	err = dec.Decode(&blockHeight)
	if err != nil {
		return this.originUnmarshal(info.States)
	}
	this.States = []interface{}{
		notifyMethod,
		transferFrom,
		transferTo,
		transferAmount,
		blockHeight,
	}
	return nil
}

func (this *NotifyEventInfo) originUnmarshal(data []byte) error {
	return json.Unmarshal(data, &this.States)
}

type SmartContractEventLog struct {
	TxHash          string
	ContractAddress string
	Message         string
}

//MerkleProof return struct
type MerkleProof struct {
	Type             string
	TransactionsRoot string
	BlockHeight      uint32
	CurBlockRoot     string
	CurBlockHeight   uint32
	TargetHashes     []string
}

type BlockTxHashes struct {
	Hash         common.Uint256
	Height       uint32
	Transactions []common.Uint256
}

type BlockTxHashesStr struct {
	Hash         string
	Height       uint32
	Transactions []string
}

type MemPoolTxState struct {
	State []*MemPoolTxStateItem
}

type MemPoolTxStateItem struct {
	Height  uint32 // The height in which tx was verified
	Type    int    // The validator flag: stateless/stateful
	ErrCode int    // Verified result
}

type MemPoolTxCount struct {
	Verified uint32 //Tx count of verified
	Verifing uint32 //Tx count of verifing
}

type GlobalParam struct {
	Key   string
	Value string
}
