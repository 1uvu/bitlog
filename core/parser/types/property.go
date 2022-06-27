package types

type BlockPropertyType int

const (
	IsInvalid BlockPropertyType = iota * 2
	IsMainchain
	IsSidechain
	IsOrphanArrived   // arraved but parent not connected
	IsOrphanDeleted   // arraved but parent not connected, and has been deleted
	IsOrphanProcessed // arraved and parent connected later
	IsStale
	IsFork
	IsSwap
)

type BlockProperty int

func (bp BlockProperty) Set(bt BlockPropertyType, t bool) BlockProperty {
	if t {
		bp |= (1 << bt)
	} else {
		bp &= ^(1 << bt)
	}
	return bp
}

func (bp BlockProperty) Is(bt BlockPropertyType) bool {
	return (bp>>bt)&1 == 1
}
