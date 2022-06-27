package types

import (
	"sort"
	"time"
)

type TimelineStatus struct {
	Timeline []*TimestampStatus // a time queue
	Since    Timestamp          // first timestamp types
}

func NewTimelineStatus(status *TimestampStatus) TimelineStatus {
	tl := TimelineStatus{make([]*TimestampStatus, 0), FromInt64(0)}
	tl.Timeline = append(tl.Timeline, status)
	tl.Since = status.T
	return tl
}

func (tl *TimelineStatus) PushStatus(statuses ...*TimestampStatus) {
	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].T < statuses[j].T
	})
	if len(tl.Timeline) == 0 {
		tl.Since = statuses[0].T
	}
	tl.Timeline = append(tl.Timeline, statuses...)
}

// find range TimestampStatus between start and end
func (tl *TimelineStatus) GetStatusByRange(start, end Timestamp) TimelineStatus {
	n := len(tl.Timeline)
	l := sort.Search(n, func(i int) bool {
		return tl.Timeline[i].T <= start
	})
	r := sort.Search(n, func(i int) bool {
		return tl.Timeline[i].T >= end
	})
	if l >= n || l > r {
		return TimelineStatus{}
	}
	return TimelineStatus{tl.Timeline[l : r+1], tl.Timeline[l].T}
}

// find left and right TimestampStatus about the pivot
func (tl *TimelineStatus) GetStatusByTime(pivot Timestamp) TimelineStatus {
	n := len(tl.Timeline)
	r := sort.Search(n, func(i int) bool {
		return tl.Timeline[i].T >= pivot
	})
	if r >= n || r <= 0 {
		return TimelineStatus{}
	}
	return TimelineStatus{tl.Timeline[r-1 : r+1], tl.Timeline[r-1].T}
}

// find the TimestampStatus that Type in ts
func (tl *TimelineStatus) GetStatusByEvents(es ...TimestampEvent) TimelineStatus {
	m := make(map[TimestampEvent]struct{}, len(es))
	for i := range es {
		m[es[i]] = struct{}{}
	}
	_tl := []*TimestampStatus{}
	for i := range tl.Timeline {
		es := tl.Timeline[i].Es
		for _, e := range es {
			if _, ok := m[e]; ok {
				_tl = append(_tl, tl.Timeline[i])
			}
		}
	}
	if len(_tl) == 0 {
		return TimelineStatus{}
	}
	return TimelineStatus{_tl, _tl[0].T}
}

// find the TimestampStatus Since the time point
func (tl *TimelineStatus) GetStatusSince(point Timestamp) TimelineStatus {
	n := len(tl.Timeline)
	r := sort.Search(n, func(i int) bool {
		return tl.Timeline[i].T >= point
	})
	if r >= n || r <= 0 {
		return TimelineStatus{}
	}
	return TimelineStatus{tl.Timeline[r:], tl.Timeline[r].T}
}

// find the TimestampStatus After the time point
func (tl *TimelineStatus) GetStatusAfter(point Timestamp) TimelineStatus {
	n := len(tl.Timeline)
	r := sort.Search(n, func(i int) bool {
		return tl.Timeline[i].T >= point
	})
	if r >= n || r <= 0 {
		return TimelineStatus{}
	}
	return TimelineStatus{tl.Timeline[:r], tl.Since}
}

type TimestampStatus struct {
	T  Timestamp
	Es []TimestampEvent
}

func NewTimestampStatus(t Timestamp, es ...TimestampEvent) *TimestampStatus {
	return &TimestampStatus{t, es}
}

func (ts *TimestampStatus) PushEvent(es ...TimestampEvent) {
	ts.Es = append(ts.Es, es...)
}

type TimestampEvent int8

const (
	Present TimestampEvent = iota
	// for block types
	Mined
	Arrived
	Verified
	Connected
	Disconnected
	// for chain types
	Genesis
	// for chain and fork types
	Fork
	Reorganize
	// for network types
	UpdateDifficulty
)

type Timestamp int64

func FromTime(t time.Time) Timestamp {
	return Timestamp(t.UnixNano())
}

func FromNow() Timestamp {
	return FromTime(time.Now())
}

func FromInt64(t int64) Timestamp {
	return Timestamp(t)
}

// TODO add more method about Timestamp
func (t Timestamp) FormatString() string {
	return ""
}

// TODO add more method about Timestamp
func (t Timestamp) FormatTime() time.Time {
	return time.Now()
}
