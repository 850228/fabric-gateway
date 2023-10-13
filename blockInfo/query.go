package blockinfo

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-protos-go-apiv2/common"
	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

/*
查询指定的peer节点加入了哪些通道
输入：*client.Network
输出：字符串切片
*/
func QueryChannels(network *client.Network) ([]string, error) {
	chaincodeName := "cscc"
	contract := network.GetContract(chaincodeName)

	evaluateResult, err := contract.EvaluateTransaction("GetChannels")
	if err != nil {
		return nil, errors.WithMessage(err, "failed to query channels: %s")
	}

	cqr := peer.ChannelQueryResponse{}
	err = proto.Unmarshal(evaluateResult, &cqr)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal: %s")
	}

	channels := make([]string, len(cqr.Channels))
	for i, _ := range cqr.Channels {
		channels[i] = cqr.Channels[i].ChannelId
	}
	return channels, nil
}

/*
查询指定通道的区块高度
输入：*client.Network , 通道名
输出：int
*/
func QueryBlockHeight(network *client.Network, channelName string) (uint64, error) {
	BCI, err := QueryBlockInfo(network, channelName)
	if err != nil {
		return ^uint64(0), err
	}
	return BCI.Height, nil
}

/*
查询指定通道的当前区块哈希
输入：*client.Network , 通道名
输出：[]byte
*/
func QueryCurrentBlockHash(network *client.Network, channelName string) ([]byte, error) {
	BCI, err := QueryBlockInfo(network, channelName)
	if err != nil {
		return nil, err
	}
	return BCI.CurrentBlockHash, nil
}

/*
查询指定通道的上一区块哈希
输入：*client.Network , 通道名
输出：[]byte
*/
func QueryPreviousBlockHash(network *client.Network, channelName string) ([]byte, error) {
	BCI, err := QueryBlockInfo(network, channelName)
	if err != nil {
		return nil, err
	}
	return BCI.PreviousBlockHash, nil
}

/*
查询指定通道的区块信息，包括高度、当前区块哈希、前一个区块哈希
输入：*client.Network , 通道名
输出：*common.BlockchainInfo
*/
func QueryBlockInfo(network *client.Network, channelName string) (*common.BlockchainInfo, error) {
	chaincodeName := "qscc"
	contract := network.GetContract(chaincodeName)

	evaluateResult, err := contract.EvaluateTransaction("GetChainInfo", channelName)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get chainInfo: %s")
	}

	bci := common.BlockchainInfo{}
	err = proto.Unmarshal(evaluateResult, &bci)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal: %s")
	}
	return &bci, nil
}

/*
查询指定通道和区块高度的区块信息
输入：network *client.Network , 通道名 , 高度
输出：*common.Block
*/
func QueryBlockByIndex(network *client.Network, channelName string, height string) (*common.Block, error) {
	chaincodeName := "qscc"
	contract := network.GetContract(chaincodeName)

	evaluateResult, err := contract.EvaluateTransaction("GetBlockByNumber", channelName, height)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to query block by index: %s")
	}

	block := common.Block{}
	err = proto.Unmarshal(evaluateResult, &block)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal: %s")
	}
	return &block, nil
}

/*
查询指定通道和区块哈希的区块信息
输入：network *client.Network , 通道名 , 区块哈希
输出：*common.Block
*/
func QueryBlockByHash(network *client.Network, channelName string, hash []byte) (*common.Block, error) {
	chaincodeName := "qscc"
	contract := network.GetContract(chaincodeName)

	evaluateResult, err := contract.EvaluateTransaction("GetBlockByHash", channelName, string(hash))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to query block by hash: %s")
	}
	block := common.Block{}
	err = proto.Unmarshal(evaluateResult, &block)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal: %s")
	}
	return &block, nil
}

/*
查询指定通道和交易id的区块信息
输入：network *client.Network , 通道名 , 交易id
输出：*common.Block
*/
func QueryBlockByTxID(network *client.Network, channelName string, TxID []byte) (*common.Block, error) {
	chaincodeName := "qscc"
	contract := network.GetContract(chaincodeName)

	evaluateResult, err := contract.EvaluateTransaction("GetBlockByTxID", channelName, string(TxID))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to query block by TxID: %s")
	}
	block := common.Block{}
	err = proto.Unmarshal(evaluateResult, &block)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal: %s")
	}
	return &block, nil
}
