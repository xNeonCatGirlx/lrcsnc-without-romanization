package cache

type CacheState byte

var (
	CacheStateActive      CacheState
	CacheStateExpired     CacheState = 1
	CacheStateNonExistant CacheState = 1
	CacheStateDisabled    CacheState = 2
)
