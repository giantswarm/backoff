package backoff

import (
	"time"
)

const (
	LongMaxWait   = 40 * time.Minute
	MediumMaxWait = 10 * time.Minute
	TinyMaxWait   = 5 * time.Second
	ShortMaxWait  = 2 * time.Minute
)

const (
	LongMaxInterval  = 60 * time.Second
	TinyMaxInterval  = 1 * time.Second
	ShortMaxInterval = 5 * time.Second
)
