package channel

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/saveio/themis-go-sdk/client"
	sdkcom "github.com/saveio/themis-go-sdk/common"
	"github.com/saveio/themis-go-sdk/utils"
	"github.com/saveio/themis/account"
	"github.com/saveio/themis/common"
	"github.com/saveio/themis/smartcontract/service/native/micropayment"
	sutils "github.com/saveio/themis/smartcontract/service/native/utils"
	"github.com/saveio/themis/vm/neovm/types"
)

var (
	CHANNEL_CONTRACT_ADDRESS, _ = utils.AddressFromHexString("0900000000000000000000000000000000000000")
	CHANNEL_CONTRACT_VERSION    = byte(0)
)

type Channel struct {
	Client *client.ClientMgr
	DefAcc *account.Account
}

func (this *Channel) InvokeNativeContract(signer *account.Account, method string, params []interface{}) (common.Uint256, error) {
	if signer == nil {
		return common.UINT256_EMPTY, errors.New("signer is nil")
	}
	tx, err := utils.NewNativeInvokeTransaction(sdkcom.GAS_PRICE, sdkcom.GAS_LIMIT, CHANNEL_CONTRACT_VERSION, CHANNEL_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	err = utils.SignToTransaction(tx, signer)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return this.Client.SendTransaction(tx)
}

func (this *Channel) PreExecInvokeNativeContract(method string, params []interface{}) (*sdkcom.PreExecResult, error) {
	tx, err := utils.NewNativeInvokeTransaction(0, 0, CHANNEL_CONTRACT_VERSION, CHANNEL_CONTRACT_ADDRESS, method, params)
	if err != nil {
		return nil, err
	}
	return this.Client.PreExecTransaction(tx)
}

func (this *Channel) RegisterPaymentEndPoint(protocol, ip, port []byte, regAccount common.Address) ([]byte, error) {
	params := &micropayment.NodeInfo{
		WalletAddr: regAccount,
		Protocol:   protocol,
		IP:         ip,
		Port:       port,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_ENDPOINT_REGISTRY,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}

func (this *Channel) OpenChannel(wallet1Addr common.Address, wallet1PubKey []byte, wallet2Addr common.Address,
	blockHeight uint64) ([]byte, error) {

	params := &micropayment.OpenChannelInfo{
		Participant1WalletAddr: wallet1Addr,
		Participant1PubKey: wallet1PubKey,
		Participant2WalletAddr: wallet2Addr,
		SettleBlockHeight:      blockHeight,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_OPEN_CHANNEL,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}

func (this *Channel) SetTotalDeposit(channelId uint64, participantWalletAddr common.Address, partnerWalletAddr common.Address, deposit uint64) ([]byte, error) {
	params := &micropayment.SetTotalDepositInfo{
		ChannelID:             channelId,
		ParticipantWalletAddr: participantWalletAddr,
		PartnerWalletAddr:     partnerWalletAddr,
		SetTotalDeposit:       deposit,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_SET_TOTALDEPOSIT,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}

func (this *Channel) GetNodePubKey(walletAddr common.Address) ([]byte, error) {
	params := &micropayment.NodePubKey{
		Participant: walletAddr,
		PublicKey:   []byte{},
	}
	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_NODE_PUBKEY,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}

	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, err
	}
	nodePubKey := &micropayment.NodePubKey{}
	source := common.NewZeroCopySource(buf)
	if err := nodePubKey.Deserialization(source); err != nil {
		return nil, fmt.Errorf("[GetNodePubKey] NodePubKey deserialization error!")
	}
	return nodePubKey.PublicKey, nil
}

func (this *Channel) SetTotalWithdraw(channelID uint64, participant, partner common.Address, totalWithdraw uint64, participantSig, participantPubKey, partnerSig, partnerPubKey []byte) ([]byte, error) {
	params := &micropayment.WithDraw{
		ChannelID:         channelID,
		Participant:       participant,
		Partner:           partner,
		TotalWithdraw:     totalWithdraw,
		ParticipantSig:    participantSig,
		ParticipantPubKey: participantPubKey,
		PartnerSig:        partnerSig,
		PartnerPubKey:     partnerPubKey,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_SET_TOTALWITHDRAW,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}
func (this *Channel) CooperativeSettle(channelID uint64, participant1Address common.Address, participant1Balance uint64, participant2Address common.Address, participant2Balance uint64, participant1Signature, participant1PubKey, participant2Signature, participant2PubKey []byte) ([]byte, error) {
	params := &micropayment.CooperativeSettleInfo{
		ChannelID:             channelID,
		Participant1Address:   participant1Address,
		Participant1Balance:   participant1Balance,
		Participant2Address:   participant2Address,
		Participant2Balance:   participant2Balance,
		Participant1Signature: participant1Signature,
		Participant1PubKey:    participant1PubKey,
		Participant2Signature: participant2Signature,
		Participant2PubKey:    participant2PubKey,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_COOPERATIVESETTLE,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}
func (this *Channel) CloseChannel(channelID uint64, participantAddress, partnerAddress common.Address, balanceHash []byte, nonce uint64, additionalHash, partnerSignature, partnerPubKey []byte) ([]byte, error) {
	params := &micropayment.CloseChannelInfo{
		ChannelID:          channelID,
		ParticipantAddress: participantAddress,
		PartnerAddress:     partnerAddress,
		BalanceHash:        balanceHash,
		Nonce:              nonce,
		AdditionalHash:     additionalHash,
		PartnerSignature:   partnerSignature,
		PartnerPubKey:      partnerPubKey,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_CLOSE_CHANNEL,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}
func (this *Channel) RegisterSecret(secret []byte) ([]byte, error) {
	tx, err := this.InvokeNativeContract(

		this.DefAcc,

		micropayment.MP_SECRET_REG,
		[]interface{}{secret},
	)
	if err != nil {
		return nil, err
	}

	return tx[:], nil
}
func (this *Channel) RegisterSecretBatch(secrets []byte) ([]byte, error) {
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_SECRET_REG_BATCH,
		[]interface{}{secrets},
	)
	if err != nil {
		return nil, err
	}
	return tx[:], nil
}

func (this *Channel) GetSecretRevealBlockHeight(secretHash []byte) (uint64, error) {
	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_SECRET_REVEAL_BLOCKHEIGHT,
		[]interface{}{secretHash},
	)
	if err != nil {
		return 0, err
	}
	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return 0, err
	}
	height, err := strconv.ParseUint(hex.EncodeToString(common.ToArrayReverse(buf)), 16, 64)
	if err != nil {
		return 0, err
	}
	return height, nil
}
func (this *Channel) UpdateNonClosingBalanceProof(chanID uint64, closeParticipant, nonCloseParticipant common.Address, balanceHash []byte, nonce uint64, additionalHash, closeSignature, nonCloseSignature, closePubKey, nonClosePubKey []byte) ([]byte, error) {
	params := &micropayment.UpdateNonCloseBalanceProof{
		ChanID:              chanID,
		CloseParticipant:    closeParticipant,
		NonCloseParticipant: nonCloseParticipant,
		BalanceHash:         balanceHash,
		Nonce:               nonce,
		AdditionalHash:      additionalHash,
		CloseSignature:      closeSignature,
		NonCloseSignature:   nonCloseSignature,
		ClosePubKey:         closePubKey,
		NonClosePubKey:      nonClosePubKey,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_UPDATE_NONCLOSING_BPF,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}

	return tx[:], nil
}

func (this *Channel) SettleChannel(chanID uint64, participant1 common.Address, p1TransferredAmount, p1LockedAmount uint64, p1LocksRoot []byte, participant2 common.Address, p2TransferredAmount, p2LockedAmount uint64, p2LocksRoot []byte) ([]byte, error) {
	params := &micropayment.SettleChannelInfo{
		ChanID:              chanID,
		Participant1:        participant1,
		P1TransferredAmount: p1TransferredAmount,
		P1LockedAmount:      p1LockedAmount,
		P1LocksRoot:         p1LocksRoot,
		Participant2:        participant2,
		P2TransferredAmount: p2TransferredAmount,
		P2LockedAmount:      p2LockedAmount,
		P2LocksRoot:         p2LocksRoot,
	}
	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_SETTLE_CHANNEL,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}

	return tx[:], nil
}

func (this *Channel) Unlock(channelID uint64, participant, partner common.Address, merkleTreeLeaves []byte) ([]byte, error) {
	params := &micropayment.UnlockInfo{
		ChannelID:          channelID,
		ParticipantAddress: participant,
		PartnerAddress:     partner,
		MerkleTreeLeaves:   merkleTreeLeaves,
	}

	tx, err := this.InvokeNativeContract(
		this.DefAcc,
		micropayment.MP_UNLOCK,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}

	return tx[:], nil
}

func (this *Channel) GetChannelInfo(channelID uint64, participant1, participant2 common.Address) (*micropayment.ChannelInfo, error) {
	params := &micropayment.GetChanInfo{
		ChannelID:    channelID,
		Participant1: participant1,
		Participant2: participant2,
	}

	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_CHANNELINFO,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}

	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, err
	}
	channelInfo := &micropayment.ChannelInfo{}
	source := common.NewZeroCopySource(buf)
	channelInfo.ChannelID, err = sutils.DecodeVarUint(source)
	if err != nil {
		return nil, err
	}
	channelInfo.SettleBlockHeight, err = sutils.DecodeVarUint(source)
	if err != nil {
		return nil, err
	}
	channelInfo.ChannelState, err = sutils.DecodeVarUint(source)
	if err != nil {
		return nil, err
	}
	channelInfo.Participant1.WalletAddr, err = sutils.DecodeAddress(source)
	if err != nil {
		return nil, err
	}
	channelInfo.Participant1.IsCloser, err = sutils.DecodeBool(source)
	if err != nil {
		return nil, err
	}
	channelInfo.Participant2.WalletAddr, err = sutils.DecodeAddress(source)
	if err != nil {
		return nil, err
	}
	channelInfo.Participant2.IsCloser, err = sutils.DecodeBool(source)
	if err != nil {
		return nil, err
	}
	return channelInfo, nil
}

func (this *Channel) GetAllOpenChannels () (*micropayment.AllChannels, error) {
	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_ALL_OPEN_CHANNELS,
		[]interface{}{},
	)
	if err != nil {
		return nil, err
	}

	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, err
	}
	allOpenChannels := &micropayment.AllChannels{}
	reader := bytes.NewReader(buf)
	err = allOpenChannels.Deserialize(reader)
	return allOpenChannels, err
}

func (this *Channel) GetChannelParticipantInfo(channelID uint64, participant1, participant2 common.Address) (*micropayment.Participant, error) {
	params := &micropayment.GetChanInfo{
		ChannelID:    channelID,
		Participant1: participant1,
		Participant2: participant2,
	}

	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_CHANNEL_PARTICIPANTINFO,
		[]interface{}{params},
	)
	if err != nil {
		return nil, err
	}

	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, err
	}
	source := common.NewZeroCopySource(buf)
	participant := &micropayment.Participant{}
	err = participant.Deserialization(source)
	if err != nil {
		return nil, err
	}
	return participant, nil
}

// GetEndpointByAddress get endpoint by user wallet address
func (this *Channel) GetEndpointByAddress(nodeAddress common.Address) (*micropayment.NodeInfo, error) {
	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_FIND_ENDPOINT,
		[]interface{}{nodeAddress},
	)
	if err != nil {
		return nil, err
	}

	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return nil, err
	}
	//haven`t register
	if len(buf) == 0 {
		return nil, nil
	}
	nodeInfo := &micropayment.NodeInfo{}
	err = nodeInfo.Deserialize(bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	return nodeInfo, nil
}

func (this *Channel) GetChannelIdentifier(participant1WalletAddr, participant2WalletAddr common.Address) (uint64, error) {
	params := &micropayment.GetChannelId{
		Participant1WalletAddr: participant1WalletAddr,
		Participant2WalletAddr: participant2WalletAddr,
	}
	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_CHANNELID,
		[]interface{}{params},
	)
	if err != nil {
		return 0, err
	}
	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return 0, err
	}
	valStr := fmt.Sprintf("%s", types.BigIntFromBytes(buf))
	return strconv.ParseUint(valStr, 10, 64)
}

func (this *Channel) GetChannelCounter() (uint64, error) {
	ret, err := this.PreExecInvokeNativeContract(
		micropayment.MP_GET_CHANNELCOUNTER,
		[]interface{}{},
	)
	if err != nil {
		return 0, err
	}
	buf, err := ret.Result.ToByteArray()
	if err != nil {
		return 0, err
	}
	valStr := fmt.Sprintf("%s", types.BigIntFromBytes(buf))
	return strconv.ParseUint(valStr, 10, 64)
}

func (this *Channel) GetFilterArgsForAllEventsFromChannelByEventId(contractAddress common.Address, address common.Address,
	eventId uint32, fromBlock uint32, toBlock uint32) ([]map[string]interface{}, map[uint32][2]uint32, error) {
	toBlockUint := uint32(toBlock)
	currentH, _ := this.Client.GetCurrentBlockHeight()
	if toBlockUint > currentH {
		return nil, nil, fmt.Errorf("toBlock bigger than currentBlockHeight:%d", currentH)
	} else if toBlockUint == 0 {
		toBlockUint = currentH
	} else if uint32(fromBlock) > toBlockUint {
		return nil, nil, errors.New("fromBlock bigger than toBlock")
	}

	events, err := this.Client.GetSmartContractEventByEventId(contractAddress.ToBase58(), address.ToBase58(), eventId)
	if err != nil {
		return nil, nil, err
	}

	var eventRe = make([]map[string]interface{}, 0)
	// map from block height to start and end index of events in the same block
	var posMap = make(map[uint32][2]uint32, 0)
	var index uint32

	for _, event := range events {
		buf, err := json.Marshal(event)
		if err != nil {
			return nil, nil, err
		}
		result, err := utils.GetSmartContractEvent(buf)
		if err != nil {
			return nil, nil, err
		}
		if result == nil {
			continue
		}

		height, err := this.Client.GetBlockHeightByTxHash(result.TxHash)
		if err != nil {
			return nil, nil, err
		}

		if height < fromBlock || height > toBlock {
			continue
		}

		for _, notify := range result.Notify {
			if _, ok := notify.States.(map[string]interface{}); !ok {
				continue
			}
			eventRe = append(eventRe, notify.States.(map[string]interface{}))

			pos, ok := posMap[height]
			if !ok {
				posMap[height] = [2]uint32{index, 0}
			} else {
				pos[1] = index
				posMap[height] = pos
			}
			index++
		}
	}
	return eventRe, posMap, nil
}

func (this *Channel) GetFilterArgsForAllEventsFromChannel(chanID int, fromBlock, toBlock uint32) ([]map[string]interface{}, error) {
	toBlockUint := uint32(toBlock)
	currentH, _ := this.Client.GetCurrentBlockHeight()
	if toBlockUint > currentH {
		return nil, fmt.Errorf("toBlock bigger than currentBlockHeight:%d", currentH)
	} else if toBlockUint == 0 {
		toBlockUint = currentH
	} else if uint32(fromBlock) > toBlockUint {
		return nil, errors.New("fromBlock bigger than toBlock")
	}
	var eventRe = make([]map[string]interface{}, 0)
	for bc := uint32(fromBlock); bc <= toBlockUint; bc++ {
		raws, err := this.Client.GetSmartContractEventsByBlock(bc)
		if err != nil {
			return nil, err
		}
		if len(raws) == 0 {
			continue
		}
		for _, r := range raws {
			buf, err := json.Marshal(r)
			if err != nil {
				return nil, err
			}
			result, err := utils.GetSmartContractEvent(buf)
			if err != nil {
				return nil, err
			}
			if result == nil {
				continue
			}
			for _, notify := range result.Notify {
				if _, ok := notify.States.(map[string]interface{}); !ok {
					continue
				}
				eventRe = append(eventRe, notify.States.(map[string]interface{}))
			}
		}
	}
	return eventRe, nil
}

// GetAllFilterArgsForAllEventsFromChannel get all events from fromBlock to current block height
// return a slice of map[string]interface{}
func (this *Channel) GetAllFilterArgsForAllEventsFromChannel(chanID int, fromBlock uint32) ([]map[string]interface{}, error) {
	height, err := this.Client.GetCurrentBlockHeight()
	if err != nil {
		return nil, err
	}
	return this.GetFilterArgsForAllEventsFromChannel(chanID, fromBlock, height)
}
