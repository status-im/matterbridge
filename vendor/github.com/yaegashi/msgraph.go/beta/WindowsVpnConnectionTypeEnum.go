// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsVpnConnectionType undocumented
type WindowsVpnConnectionType int

const (
	// WindowsVpnConnectionTypeVPulseSecure undocumented
	WindowsVpnConnectionTypeVPulseSecure WindowsVpnConnectionType = 0
	// WindowsVpnConnectionTypeVF5EdgeClient undocumented
	WindowsVpnConnectionTypeVF5EdgeClient WindowsVpnConnectionType = 1
	// WindowsVpnConnectionTypeVDellSonicWallMobileConnect undocumented
	WindowsVpnConnectionTypeVDellSonicWallMobileConnect WindowsVpnConnectionType = 2
	// WindowsVpnConnectionTypeVCheckPointCapsuleVpn undocumented
	WindowsVpnConnectionTypeVCheckPointCapsuleVpn WindowsVpnConnectionType = 3
)

// WindowsVpnConnectionTypePPulseSecure returns a pointer to WindowsVpnConnectionTypeVPulseSecure
func WindowsVpnConnectionTypePPulseSecure() *WindowsVpnConnectionType {
	v := WindowsVpnConnectionTypeVPulseSecure
	return &v
}

// WindowsVpnConnectionTypePF5EdgeClient returns a pointer to WindowsVpnConnectionTypeVF5EdgeClient
func WindowsVpnConnectionTypePF5EdgeClient() *WindowsVpnConnectionType {
	v := WindowsVpnConnectionTypeVF5EdgeClient
	return &v
}

// WindowsVpnConnectionTypePDellSonicWallMobileConnect returns a pointer to WindowsVpnConnectionTypeVDellSonicWallMobileConnect
func WindowsVpnConnectionTypePDellSonicWallMobileConnect() *WindowsVpnConnectionType {
	v := WindowsVpnConnectionTypeVDellSonicWallMobileConnect
	return &v
}

// WindowsVpnConnectionTypePCheckPointCapsuleVpn returns a pointer to WindowsVpnConnectionTypeVCheckPointCapsuleVpn
func WindowsVpnConnectionTypePCheckPointCapsuleVpn() *WindowsVpnConnectionType {
	v := WindowsVpnConnectionTypeVCheckPointCapsuleVpn
	return &v
}