package types

type LyricsState byte

const (
	LyricsStateSynced       LyricsState = 0
	LyricsStatePlain        LyricsState = 1
	LyricsStateInstrumental LyricsState = 2
	LyricsStateNotFound     LyricsState = 3
	LyricsStateLoading      LyricsState = 4
	LyricsStateUnknown      LyricsState = 5
)

func (l LyricsState) String() string {
	switch l {
	case LyricsStateSynced:
		return "synced"
	case LyricsStatePlain:
		return "plain"
	case LyricsStateInstrumental:
		return "instrumental"
	case LyricsStateNotFound:
		return "not-found"
	case LyricsStateLoading:
		return "loading"
	default:
		return "unknown"
	}
}

func (l LyricsState) ToCacheStoreCondition() CacheStoreConditionType {
	switch l {
	case LyricsStateSynced:
		return CacheStoreConditionSynced
	case LyricsStatePlain:
		return CacheStoreConditionPlain
	case LyricsStateInstrumental:
		return CacheStoreConditionInstrumental
	default:
		return CacheStoreConditionNone
	}
}
