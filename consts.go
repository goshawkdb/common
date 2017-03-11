package common

import (
	"time"
)

const (
	DefaultPort           = 7894
	DefaultPrometheusPort = 8080
	DefaultWSSPort        = 7895
	ProductVersion        = "dev" // protocol version!
	ProductName           = "GoshawkDB"
	HeartbeatInterval     = 2 * time.Second
	ConnectionBufferSize  = 131072
)

var (
	VersionZero = MakeTxnId([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
)
