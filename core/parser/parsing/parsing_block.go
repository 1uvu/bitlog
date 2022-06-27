package parsing

import (
	"fmt"
	"time"

	"github.com/1uvu/bitlog/pkg/utils"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type Block struct {
	MinerAddress btcutil.Address
	BlockHash    chainhash.Hash
	Header       wire.BlockHeader
	Txn          int64
	ConnectTime  time.Time
}

func NewBlock(b *wire.MsgBlock, connectTime time.Time) *Block {
	var genesisTX *wire.MsgTx
	for _, tx := range b.Transactions {
		if blockchain.IsCoinBaseTx(tx) {
			genesisTX = tx
			break
		}
	}
	minerAddress := func(tx *wire.MsgTx) btcutil.Address {
		var minerAddress btcutil.Address
		for _, out := range tx.TxOut {
			_, minerAddresses, _, _ := txscript.ExtractPkScriptAddrs(out.PkScript, &chaincfg.TestNet3Params)
			if len(minerAddresses) > 0 {
				minerAddress = minerAddresses[0]
			}
		}
		return minerAddress
	}(genesisTX)
	return &Block{
		MinerAddress: minerAddress,
		BlockHash:    b.BlockHash(),
		Header:       b.Header,
		Txn:          int64(len(b.Transactions)),
		ConnectTime:  connectTime,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"{blockMiner=%s blockHash=%s blockHeader=%+v transactionNumber=%d timestampConnected=%s}",
		b.MinerAddress.EncodeAddress(),
		b.BlockHash.String(),
		b.Header,
		b.Txn,
		utils.TimeLocal(b.ConnectTime),
	)
}
