package types

// BlockStatus records a block (any type) full lifetime
/*
	ID              BlockID        // var: 'b'+height+number
	Timeline        TimelineStatus // only livetime TimestampStatus, the TimelineStatus will be analyzed in analyzer plugin
	ParentBlock     BlockID        // block.Parent
	PrevBlockStatus BlockID        // var: prev id
	BestBlock       BlockID        // b.stateSnapshot
	Info            BlockInfo      // block + block header
	Property        BlockProperty  // releaed to timeline
*/
type BlockStatus struct {
	ID              BlockID        // for logging, use block id to replace the pointer can find a block types by search logging text
	Timeline        TimelineStatus // a timeline that record all event and timestamp about block types transfer
	ParentBlock     BlockID        // parent block types
	PrevBlockStatus BlockID        // previous block types
	BestBlock       BlockID        // current best block types about mainchain
	Info            BlockInfo      // block detail information
	Property        BlockProperty  // block detail property
}

type BlockID string

type BlockInfo struct {
	Height     int64  // current block height about mainchain
	TxN        int64  // tx number in block
	Size       int64  // block size in bytes
	Difficulty int64  // target block difficulty in bytes
	Worksum    int64  // block worksum since genesis in bytes
	Hash       string // block hash
	Nonce      string // block nonce for mining
	Miner      string // miner address
}

func (blockStatus *BlockStatus) Marshal() []byte {
	return nil
}
