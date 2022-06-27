package types

// NetworkStatus records a network full lifetime
/*
	Timeline   TimelineStatus // only livetime TimestampStatus, the TimelineStatus will be analyzed in analyzer plugin
	Difficulty int64 		  // bits
*/
type NetworkStatus struct {
	Timeline   TimelineStatus // a timeline that record all event and timestamp about network types transfer
	Difficulty int64
}

type NetworkLocation string

func (networkStatus *NetworkStatus) Marshal() []byte {
	return nil
}
