package common

import (
	"time"
)

const (
	DefaultPort       = 7894
	ProductVersion    = "1"
	ProductName       = "goshawkdb"
	HeartbeatInterval = 2 * time.Second
)

var (
	VersionZero = MakeTxnId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
)
