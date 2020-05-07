package client

import (
	"context"
	"encoding/hex"
	watcher "github.com/mapofzones/cosmos-watcher/types"
	"github.com/tendermint/tendermint/rpc/client/http"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	"time"
)

var clientTimeout = 5 * time.Second

// ClientProxy implements a wrapper Tendermint RPC client
type ClientProxy struct {
	clientService *http.HTTP
}

func newRPCClient(serviceRPC string) (*http.HTTP, error) {
	service, err := http.New(serviceRPC, "/websocket")
	if err != nil {
		return nil, err
	}

	err = service.Start()
	if err != nil {
		return nil, err
	}

	return service, err
}

func New(rpcNode string) (ClientProxy, error) {
	rpcClient, err := newRPCClient(rpcNode)
	if err != nil {
		return ClientProxy{} , err
	}

	return ClientProxy{clientService: rpcClient}, nil

}

// LatestHeight returns the latest block height on the active chain. An error
// is returned if the query fails.
func (cp ClientProxy) LatestHeight() (int64, error) {
	status, err := cp.clientService.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}

// GetBlockByHeight queries for a block by height. An error is returned if the query fails.
func (cp ClientProxy) GetBlockByHeight(height int64) (*tmctypes.ResultBlock, error) {
	block, err := cp.clientService.Block(&height)
	if err != nil {
		return nil, err
	}
	return block, err
}

func (cp ClientProxy) GetTxsByBlock(block *tmctypes.ResultBlock) ([]watcher.TxStatus, error) {
	s := []watcher.TxStatus{}
	for _, tx := range block.Block.Txs {
		res, err := cp.clientService.Tx(tx.Hash(), false)
		if err != nil {
			continue
		}
		s = append(s, watcher.TxStatus{
			ResultCode: res.TxResult.Code,
			Hash:       tx.Hash(),
			Height:     res.Height,
		})
	}
	return s, nil
}

// SubscribeNewBlocks subscribes to the new block event handler through the RPC
// client with the given subscriber name. An receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
func (cp ClientProxy) SubscribeNewBlocks(subscriber string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	eventCh, err := cp.clientService.Subscribe(ctx, subscriber, "tm.event = 'NewBlock'")
	return eventCh, cancel, err
}

func (cp ClientProxy) CreateWatcherBlock(block *tmctypes.ResultBlock, txs []watcher.TxStatus) (watcher.Block, error) {
	return watcher.Block{
		ChainID: block.Block.ChainID,
		Height:  block.Block.Height,
		T:       block.Block.Time,
		Txs:     block.Block.Txs,
		Results: txs,
	}, nil
}

// TendermintTx queries for a transaction by hash. An error is returned if the
// query fails.
func (cp ClientProxy) TendermintTx(hash string) (*tmctypes.ResultTx, error) {
	hashRaw, err := hex.DecodeString(hash)
	if err != nil {
		return nil, err
	}

	return cp.clientService.Tx(hashRaw, false)
}

// Stop defers the node stop execution to the RPC client.
func (cp ClientProxy) Stop() error {
	return cp.clientService.Stop()
}
