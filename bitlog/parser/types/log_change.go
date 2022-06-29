package types

import (
	"github.com/1uvu/bitlog/pkg/id"
)

type (
	ChangeLog struct {
		// change detail
		Change RawLog

		// resolver
		ID            id.ID
		PrevChangeLog *ChangeLog // RelevantStatusLogs.Tail.RelevantChangeLog
		NextChangeLog *ChangeLog // create when IsValid=true

		IsValid            bool
		RelevantStatusLogs []*StatusLog
	}
)

// TODO
const (
	ChangeTypeBlock   = RawLogType("change_block")
	ChangeTypeChain   = RawLogType("change_chain")
	ChangeTypeFork    = RawLogType("change_fork")
	ChangeTypeNetwork = RawLogType("change_network")
	ChangeTypeUnknown = RawLogType("change_unknown")
)

func (changeLog *ChangeLog) String() string {
	return ""
}
