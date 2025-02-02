// Code generated by "stringer -type=networkProp -output network_props_linux.go"; DO NOT EDIT.

package device

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ConnectionState-7]
	_ = x[ConnectionID-8]
	_ = x[ConnectionDevices-9]
	_ = x[ConnectionType-10]
	_ = x[ConnectionIPv4-11]
	_ = x[ConnectionIPv6-12]
	_ = x[AddressIPv4-13]
	_ = x[AddressIPv6-14]
	_ = x[WifiSSID-15]
	_ = x[WifiFrequency-16]
	_ = x[WifiSpeed-17]
	_ = x[WifiStrength-18]
	_ = x[WifiHWAddress-19]
}

const _networkProp_name = "ConnectionStateConnectionIDConnectionDevicesConnectionTypeConnectionIPv4ConnectionIPv6AddressIPv4AddressIPv6WifiSSIDWifiFrequencyWifiSpeedWifiStrengthWifiHWAddress"

var _networkProp_index = [...]uint8{0, 15, 27, 44, 58, 72, 86, 97, 108, 116, 129, 138, 150, 163}

func (i networkProp) String() string {
	i -= 7
	if i < 0 || i >= networkProp(len(_networkProp_index)-1) {
		return "networkProp(" + strconv.FormatInt(int64(i+7), 10) + ")"
	}
	return _networkProp_name[_networkProp_index[i]:_networkProp_index[i+1]]
}
