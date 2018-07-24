package util


type CongestionLevel int

const (
	FREE_FLOW CongestionLevel = iota+1
	REASONABLY_FREE_FLOW
	STABLE_FLOW
	APPROACHING_UNSTABLE_FLOW
	UNSTABLE_FLOW
	BREAKDOWN_FLOW
	UNKNOWN_STATE
)

func TranslateCongestionLevel(duration int)(CongestionLevel)  {

	return FREE_FLOW

}



