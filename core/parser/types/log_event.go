package types

import (
	"github.com/1uvu/bitlog/pkg/id"
)

type (
	EventLogLinkedList struct {
		Timeline   Timeline
		Head, Tail *EventLog
	}
	EventLog struct {
		// inject
		Type      EventType
		Timestamp Timestamp

		// parsing with inject
		Raw []byte

		// resolver
		ID                id.ID
		RelevantStatusLog *StatusLog
		PrevEventLog      *EventLog // RelevantStatusLog.RelevantEventLogs.Tail
		NextEventLog      *EventLog // create when happen a new event
	}
	EventType string
)

// TODO
const (
	// EventType for ChangeTypeBlock
	EventTypeBlockArrival       = EventType("block_arrival")
	EventTypeBlockVerifySuccess = EventType("block_verify_success")
	EventTypeBlockVerifyFailed  = EventType("block_verify_failed")
	EventTypeBlockConnect       = EventType("block_connect")
	EventTypeBlockDisconnect    = EventType("block_disconnect")
	// others block event...

	// EventType for ChangeTypeChain
	EventTypeChainIncrease   = EventType("chain_increase")
	EventTypeChainDecrease   = EventType("chain_decrease")
	EventTypeChainReorganize = EventType("chain_reorganize")
	// others chain event...

	EventTypeFork    = EventType("fork")
	EventTypeNetwork = EventType("network")
	EventTypeUnknown = EventType("unknown")
)

func (eventLog *EventLog) String() string {
	return ""
}
