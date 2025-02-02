// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package hass

import (
	"errors"
	"fmt"
)

const (
	// RequestTypeEncrypted is a RequestType of type encrypted.
	RequestTypeEncrypted RequestType = "encrypted"
	// RequestTypeGetConfig is a RequestType of type get_config.
	RequestTypeGetConfig RequestType = "get_config"
	// RequestTypeUpdateLocation is a RequestType of type update_location.
	RequestTypeUpdateLocation RequestType = "update_location"
	// RequestTypeRegisterSensor is a RequestType of type register_sensor.
	RequestTypeRegisterSensor RequestType = "register_sensor"
	// RequestTypeUpdateSensorStates is a RequestType of type update_sensor_states.
	RequestTypeUpdateSensorStates RequestType = "update_sensor_states"
)

var ErrInvalidRequestType = errors.New("not a valid RequestType")

// String implements the Stringer interface.
func (x RequestType) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x RequestType) IsValid() bool {
	_, err := ParseRequestType(string(x))
	return err == nil
}

var _RequestTypeValue = map[string]RequestType{
	"encrypted":            RequestTypeEncrypted,
	"get_config":           RequestTypeGetConfig,
	"update_location":      RequestTypeUpdateLocation,
	"register_sensor":      RequestTypeRegisterSensor,
	"update_sensor_states": RequestTypeUpdateSensorStates,
}

// ParseRequestType attempts to convert a string to a RequestType.
func ParseRequestType(name string) (RequestType, error) {
	if x, ok := _RequestTypeValue[name]; ok {
		return x, nil
	}
	return RequestType(""), fmt.Errorf("%s is %w", name, ErrInvalidRequestType)
}

// MarshalText implements the text marshaller method.
func (x RequestType) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *RequestType) UnmarshalText(text []byte) error {
	tmp, err := ParseRequestType(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
