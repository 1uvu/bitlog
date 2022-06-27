package types

// ForkStatus records a fork full lifetime
/*
	ID              ForkID         // var: 'f'+height+number
	Timeline        TimelineStatus // only livetime TimestampStatus, the TimelineStatus will be analyzed in analyzer plugin
	PrevForkStatus  ForkID         // var: prev id
	ForkBlock       BlockID        // Connectbestchain FindFork()
	BestBlock       BlockID        // b.stateSnapshot
	CurrentBlock    BlockID        // Connectbestchain FindFork()
	CheckPoint 		int64          // CurrentBlock.Height+6
	Reorganized     bool           // Connectbestchain Reorganizechain()
*/
type ForkStatus struct {
	ID             ForkID         //
	Timeline       TimelineStatus // a timeline that record all event and timestamp about fork types transfer
	PrevForkStatus ForkID         // previous fork types
	ForkBlock      BlockID        // fork block about a fork
	BestBlock      BlockID        // best block of mainchain
	CurrentBlock   BlockID        // current block in sidechain, is best block in sidechain currently
	CheckPoint     int64          // checkpoint block in sidechain, fork Height + 6
	Reorganized    bool           // current fork cause a Reorganize
}

type ForkID string

func (forkStatus *ForkStatus) Marshal() []byte {
	return nil
}
