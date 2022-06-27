package types

// ChainStatus records a chain full lifetime
/*
	ID              ChainID        // var: 'c'+height+number
	Timeline        TimelineStatus // only livetime TimestampStatus, the TimelineStatus will be analyzed in analyzer plugin
	PrevChainStatus ChainID        // var: prev id
	BestBlock       BlockID        // b.stateSnapshot
	Info            ChainInfo      // b.stateSnapshot
	// TODO opt zengliang huo jiangliang
	// => zengliang huo jiangliang: index in Block timeline
	OrphanArrivedSet 	map[*TimestampStatus]BlockID // only increasement and decreasement, the set will be analyzed in analyzer plugin
	OrphanDeletedSet 	map[*TimestampStatus]BlockID // only increasement and decreasement, the set will be analyzed in analyzer plugin
	OrphanProcessedSet 	map[*TimestampStatus]BlockID // only increasement and decreasement, the set will be analyzed in analyzer plugin
	StaleSet   			map[*TimestampStatus]BlockID // only increasement and decreasement, the set will be analyzed in analyzer plugin
	ForkSet    			map[*TimestampStatus]BlockID // only increasement and decreasement, the set will be analyzed in analyzer plugin
*/
type ChainStatus struct {
	ID              ChainID        //
	Timeline        TimelineStatus // a timeline that record all event and timestamp about chain types transfer
	PrevChainStatus ChainID        // previous chain types
	BestBlock       BlockID        // current best block about mainchain
	Info            ChainInfo      // chain detail information
	// TODO opt zengliang huo jiangliang
	// => zengliang huo jiangliang: index in Block timeline
	OrphanArrivedSet   map[*TimestampStatus]BlockID // current orphan block in property IsOrphanArrived
	OrphanDeletedSet   map[*TimestampStatus]BlockID // current orphan block in property IsOrphanDeleted
	OrphanProcessedSet map[*TimestampStatus]BlockID // current orphan block in property IsOrphanProcessed
	StaleSet           map[*TimestampStatus]BlockID // current stale block
	ForkSet            map[*TimestampStatus]BlockID // current fork block
}

type ChainID string

type ChainInfo struct {
	Height           int64 // current block height about mainchain
	Worksum          int64 // total worksum until best block in mainchain
	BlockN           int64 // total block, include mainchain and sidechain
	TxN              int64 // total tx, include mainchain and sidechain
	OrphanArrivedN   int64 // total arrived orphan block
	OrphanDeletedN   int64 // total deleted orphan block
	OrphanProcessedN int64 // total processed orphan block
	StaleN           int64 // total stale block
	ForkN            int64 // total fork block
}

func (chainStatus *ChainStatus) Marshal() []byte {
	return nil
}
