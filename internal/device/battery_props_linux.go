// Code generated by "stringer -type=BatteryProp -output battery_props_linux.go -trimprefix batt"; DO NOT EDIT.

package device

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[battType-3]
	_ = x[Percentage-4]
	_ = x[Temperature-5]
	_ = x[Voltage-6]
	_ = x[Energy-7]
	_ = x[EnergyRate-8]
	_ = x[battState-9]
	_ = x[NativePath-10]
	_ = x[BatteryLevel-11]
}

const _BatteryProp_name = "TypePercentageTemperatureVoltageEnergyEnergyRateStateNativePathBatteryLevel"

var _BatteryProp_index = [...]uint8{0, 4, 14, 25, 32, 38, 48, 53, 63, 75}

func (i BatteryProp) String() string {
	i -= 3
	if i < 0 || i >= BatteryProp(len(_BatteryProp_index)-1) {
		return "BatteryProp(" + strconv.FormatInt(int64(i+3), 10) + ")"
	}
	return _BatteryProp_name[_BatteryProp_index[i]:_BatteryProp_index[i+1]]
}