// +build linux

package defaults

// GetDefaults returns sane defaults for the Linux platform.
// The "default" options may be may be replaced by the running configuration.
func GetDefaults() PlatformDefaultParameters {
	return PlatformDefaultParameters{
		// Admin
		DefaultAdminListen: "unix:///var/run/yggdrasil.sock",

		// Configuration (used for yggdrasilctl)
		DefaultConfigFile: "/etc/yggdrasil.conf",

		// Multicast interfaces
		DefaultMulticastInterfaces: []string{
			".*",
		},

		// TUN/TAP
		MaximumIfMTU:  65535,
		DefaultIfMTU:  65535,
		DefaultIfName: "auto",
	}
}
