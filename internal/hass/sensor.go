// Copyright (c) 2023 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package hass

import "strings"

//go:generate stringer -type=SensorType,SensorDeviceClass,SensorStateClass -output sensor_strings.go -trimprefix Sensor

const (
	TypeSensor SensorType = iota + 1
	TypeBinary

	Apparent_power SensorDeviceClass = iota + 1
	Aqi
	Atmospheric_pressure
	SensorBattery
	Carbon_dioxide
	Carbon_monoxide
	Current
	Data_rate
	Data_size
	Date
	Distance
	Duration
	Energy
	Enum
	Frequency
	Gas
	Humidity
	Illuminance
	Irradiance
	Moisture
	Monetary
	Nitrogen_dioxide
	Nitrogen_monoxide
	Nitrous_oxide
	Ozone
	Pm1
	Pm25
	Pm10
	Power_factor
	SensorPower
	Precipitation
	Precipitation_intensity
	Pressure
	Reactive_power
	Signal_strength
	Sound_pressure
	Speed
	Sulphur_dioxide
	SensorTemperature
	Timestamp
	Volatile_organic_compounds
	Voltage
	Volume
	Water
	Weight
	Wind_speed

	Measurement SensorStateClass = iota
	Total
	Total_increasing
)

// SensorType reflects the type of sensor, sensor or binary_sensor.
type SensorType int

// SensorDeviceClass reflects the HA device class of the sensor.
type SensorDeviceClass int

// SensorStateClass reflects the HA state class of the sensor.
type SensorStateClass int

// Sensor represents a sensor in HA. As an interface, it leaves it up to the
// underlying struct to provide the appropriate data for this representation.
type Sensor interface {
	Attributes() interface{}
	DeviceClass() string
	Icon() string
	Name() string
	State() interface{}
	Type() string
	UniqueID() string
	UnitOfMeasurement() string
	StateClass() string
	EntityCategory() string
	Disabled() bool
	Registered() bool
}

// SensorUpdate represents an update for a sensor. It reflects the current state
// of the sensor at the point in time it is used. It provides a bridge between
// platform/device and HA implementations of what a sensor is.
type SensorUpdate interface {
	Name() string
	ID() string
	Icon() string
	SensorType() SensorType
	DeviceClass() SensorDeviceClass
	StateClass() SensorStateClass
	State() interface{}
	Units() string
	Category() string
	Attributes() interface{}
}

// sensorRegistrationInfo is the JSON structure required to register a sensor
// with HA.
type sensorRegistrationInfo struct {
	StateAttributes   interface{} `json:"attributes,omitempty"`
	DeviceClass       string      `json:"device_class,omitempty"`
	Icon              string      `json:"icon,omitempty"`
	Name              string      `json:"name"`
	State             interface{} `json:"state"`
	Type              string      `json:"type"`
	UniqueID          string      `json:"unique_id"`
	UnitOfMeasurement string      `json:"unit_of_measurement,omitempty"`
	StateClass        string      `json:"state_class,omitempty"`
	EntityCategory    string      `json:"entity_category,omitempty"`
	Disabled          bool        `json:"disabled,omitempty"`
}

// sensorUpdateInfo is the JSON structure required to update HA with the current
// sensor state.
type sensorUpdateInfo struct {
	StateAttributes interface{} `json:"attributes,omitempty"`
	Icon            string      `json:"icon,omitempty"`
	State           interface{} `json:"state"`
	Type            string      `json:"type"`
	UniqueID        string      `json:"unique_id"`
}

// MarshalSensorDate takes the sensor data and returns the appropriate JSON
// structure to either register or update the sensor in HA.
func MarshalSensorData(s Sensor) interface{} {
	if s.Registered() {
		return []sensorUpdateInfo{{
			StateAttributes: s.Attributes(),
			Icon:            s.Icon(),
			State:           s.State(),
			Type:            s.Type(),
			UniqueID:        s.UniqueID(),
		},
		}
	} else {
		return sensorRegistrationInfo{
			StateAttributes:   s.Attributes(),
			DeviceClass:       s.DeviceClass(),
			Icon:              s.Icon(),
			Name:              s.Name(),
			State:             s.State(),
			Type:              s.Type(),
			UniqueID:          s.UniqueID(),
			UnitOfMeasurement: s.UnitOfMeasurement(),
			StateClass:        strings.ToLower(s.StateClass()),
			EntityCategory:    s.EntityCategory(),
			Disabled:          s.Disabled(),
		}
	}
}
