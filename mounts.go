package config

// Mounts stores the (string) mount points
type Mounts struct {
	DMS3FS         string
	DMS3NS         string
	FuseAllowOther bool
}
