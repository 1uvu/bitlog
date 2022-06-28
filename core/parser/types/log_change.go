package types

import (
	"github.com/1uvu/bitlog/core/pkg/id"
)

type (
	ChangeLog struct {
		// inject
		Type      ChangeType
		Timestamp Timestamp

		// parsing with inject
		Raw []byte

		// resolver
		ID                 id.ID
		IsValid            bool
		RelevantStatusLogs []*StatusLog
		PrevChangeLog      *ChangeLog // RelevantStatusLogs.Tail.RelevantChangeLog
		NextChangeLog      *ChangeLog // create when IsValid=true
	}
	ChangeType string
)

// TODO
const (
	ChangeTypeBlock   = ChangeType("block")
	ChangeTypeChain   = ChangeType("chain")
	ChangeTypeFork    = ChangeType("fork")
	ChangeTypeNetwork = ChangeType("network")
	ChangeTypeUnknown = ChangeType("unknown")
)

func (changeLog *ChangeLog) String() string {
	return ""
}
