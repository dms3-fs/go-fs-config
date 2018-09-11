package config

type SwarmConfig struct {
	AddrFilters             []string
	DisableBandwidthMetrics bool
	DisableNatPortMap       bool
	DisableRelay            bool
	EnableRelayHop          bool

	ConnMgr ConnMgr
}

// ConnMgr defines configuration options for the dms3-p2p connection manager
type ConnMgr struct {
	Type        string
	LowWater    int
	HighWater   int
	GracePeriod string
}
