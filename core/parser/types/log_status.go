package types

import "github.com/1uvu/bitlog/pkg/id"

type (
	StatusLog struct {
		// inject
		Type      StatusType
		Timestamp Timestamp

		// parsing with inject
		Raw []byte

		// resolver
		ID                id.ID
		RelevantEventLogs *EventLogLinkedList
		RelevantChangeLog *ChangeLog
		PrevStatusLog     *StatusLog // last status
		NextStatusLog     *StatusLog // create when status changed
	}
	StatusType string
)

// TODO
const (
	StatueTypeChain   = StatusType("chain")
	StatueTypeNetwork = StatusType("network")
	StatusTypeUnknown = StatusType("unknown")
)

func (statusLog *StatusLog) String() string {
	return ""
}
