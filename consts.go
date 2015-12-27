package common

import (
	"time"
)

const (
	DefaultPort       = 7894
	ProductVersion    = "dev" // protocol version
	ProductName       = "GoshawkDB"
	HeartbeatInterval = 2 * time.Second
)

var (
	VersionZero = MakeTxnId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
)
