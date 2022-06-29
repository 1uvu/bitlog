package types

import (
	"github.com/1uvu/bitlog/pkg/id"
)

type (
	EventLogLinkedList struct {
		Timeline   Timeline // TODO timestamp + event = timeline
		Head, Tail *EventLog
	}
	EventLog struct {
		// event detail
		EventRaw RawLog

		// resolver
		ID           id.ID
		PrevEventLog *EventLog // RelevantStatusLog.RelevantEventLogs.Tail
		NextEventLog *EventLog // create when happen a new event

		RelevantStatusLog *StatusLog
	}
)

// TODO
const (
	// EventType for ChangeTypeBlock
	EventTypeBlockArrival       = RawLogType("event_block_arrival")
	EventTypeBlockVerifySuccess = RawLogType("event_block_verify_success")
	EventTypeBlockVerifyFailed  = RawLogType("event_block_verify_failed")
	EventTypeBlockConnect       = RawLogType("event_block_connect")
	EventTypeBlockDisconnect    = RawLogType("event_block_disconnect")
	// others block event...

	// EventType for ChangeTypeChain
	EventTypeChainIncrease   = RawLogType("event_chain_increase")
	EventTypeChainDecrease   = RawLogType("event_chain_decrease")
	EventTypeChainReorganize = RawLogType("event_chain_reorganize")
	// others chain event...

	EventTypeFork    = RawLogType("event_fork")
	EventTypeNetwork = RawLogType("event_network")
	EventTypeUnknown = RawLogType("event_unknown")
)

func (eventLog *EventLog) String() string {
	return ""
}
