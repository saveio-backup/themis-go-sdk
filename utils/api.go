package utils

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"

	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/core/payload"
	"github.com/saveio/themis/core/types"
)

func GetVersion(data []byte) (string, error) {
	version := ""
	err := json.Unmarshal(data, &version)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return version, nil
}

func GetBlock(data []byte) (*types.Block, error) {
	hexStr := ""
	err := json.Unmarshal(data, &hexStr)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	blockData, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	return types.BlockFromRawBytes(blockData)
}

func GetUint32(data []byte) (uint32, error) {
	count := uint32(0)
	err := json.Unmarshal(data, &count)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return count, nil
}

func GetUint64(data []byte) (uint64, error) {
	count := uint64(0)
	err := json.Unmarshal(data, &count)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return count, nil
}

func GetInt(data []byte) (int, error) {
	integer := 0
	err := json.Unmarshal(data, &integer)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return integer, nil
}

func GetUint256(data []byte) (common.Uint256, error) {
	hexHash := ""
	err := json.Unmarshal(data, &hexHash)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("json.Unmarshal hash:%s error:%s", data, err)
	}
	hx, err := common.HexToBytes(hexHash)
	if err != nil {
		return common.Uint256{}, err
	}
	return Uint256ParseFromBytes(hx)
}

func GetUint256Reversed(data []byte) (common.Uint256, error) {
	hexHash := ""
	err := json.Unmarshal(data, &hexHash)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("json.Unmarshal hash:%s error:%s", data, err)
	}
	hash, err := common.Uint256FromHexString(hexHash)
	if err != nil {
		return common.Uint256{}, fmt.Errorf("ParseUint256FromHexString:%s error:%s", data, err)
	}
	return hash, nil
}

func GetTransaction(data []byte) (*types.Transaction, error) {
	hexStr := ""
	err := json.Unmarshal(data, &hexStr)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	txData, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	return types.TransactionFromRawBytes(txData)
}

func GetStorage(data []byte) ([]byte, error) {
	hexData := ""
	err := json.Unmarshal(data, &hexData)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	value, err := hex.DecodeString(hexData)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	return value, nil
}

func GetSmartContractEvent(data []byte) (*sdkcom.SmartContactEvent, error) {
	event := &sdkcom.SmartContactEvent{}
	err := json.Unmarshal(data, &event)
	if err == nil {
		return event, nil
	}
	events, err := GetSmartContractEvents(data)
	if err != nil {
		return nil, err
	}
	if len(events) > 0 {
		return events[0], nil
	}
	return nil, nil
}

func GetSmartContractEventLog(data []byte) (*sdkcom.SmartContractEventLog, error) {
	log := &sdkcom.SmartContractEventLog{}
	err := json.Unmarshal(data, &log)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal SmartContractEventLog:%s error:%s", data, err)
	}
	return log, nil
}

func GetSmartContractEvents(data []byte) ([]*sdkcom.SmartContactEvent, error) {
	events := make([]*sdkcom.SmartContactEvent, 0)
	err := json.Unmarshal(data, &events)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal SmartContactEvent:%s error:%s", data, err)
	}
	return events, nil
}

func GetSmartContract(data []byte) (*payload.DeployCode, error) {
	hexStr := ""
	err := json.Unmarshal(data, &hexStr)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	if hexStr == "" {
		return nil, nil
	}
	hexData, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("hex.DecodeString error:%s", err)
	}
	buf := bytes.NewReader(hexData)
	deploy := &payload.DeployCode{}
	err = deploy.Deserialize(buf)
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func GetMerkleProof(data []byte) (*sdkcom.MerkleProof, error) {
	proof := &sdkcom.MerkleProof{}
	err := json.Unmarshal(data, proof)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	return proof, nil
}

func GetBlockTxHashes(data []byte) (*sdkcom.BlockTxHashes, error) {
	blockTxHashesStr := &sdkcom.BlockTxHashesStr{}
	err := json.Unmarshal(data, &blockTxHashesStr)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal")
	}
	blockTxHashes := &sdkcom.BlockTxHashes{}

	blockHash, err := common.Uint256FromHexString(blockTxHashesStr.Hash)
	if err != nil {
		return nil, err
	}
	txHashes := make([]common.Uint256, 0, len(blockTxHashesStr.Transactions))
	for _, txHashStr := range blockTxHashesStr.Transactions {
		txHash, err := common.Uint256FromHexString(txHashStr)
		if err != nil {
			return nil, err
		}
		txHashes = append(txHashes, txHash)
	}
	blockTxHashes.Hash = blockHash
	blockTxHashes.Height = blockTxHashesStr.Height
	blockTxHashes.Transactions = txHashes
	return blockTxHashes, nil
}

func GetMemPoolTxState(data []byte) (*sdkcom.MemPoolTxState, error) {
	txState := &sdkcom.MemPoolTxState{}
	err := json.Unmarshal(data, txState)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	return txState, nil
}

func GetMemPoolTxCount(data []byte) (*sdkcom.MemPoolTxCount, error) {
	count := make([]uint32, 0, 2)
	err := json.Unmarshal(data, &count)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error:%s", err)
	}
	if len(count) != 2 {
		return nil, fmt.Errorf("count len != 2")
	}
	return &sdkcom.MemPoolTxCount{
		Verified: count[0],
		Verifing: count[1],
	}, nil
}
