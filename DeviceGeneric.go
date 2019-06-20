package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceGenericInterface = DeviceInterface + ".Generic"

	// Properties
	DeviceGenericPropertyHwAddress       = DeviceGenericInterface + ".HwAddress"       // readable   s
	DeviceGenericPropertyTypeDescription = DeviceGenericInterface + ".TypeDescription" // readable   s
)

type DeviceGeneric interface {
	Device

	// Active hardware address of the device.
	GetPropertyHwAddress() (string, error)

	// A (non-localized) description of the interface type, if known.
	GetPropertyTypeDescription() (string, error)
}

func NewDeviceGeneric(objectPath dbus.ObjectPath) (DeviceGeneric, error) {
	var d deviceGeneric
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceGeneric struct {
	device
}

func (d *deviceGeneric) GetPropertyHwAddress() (string, error) {
	return d.getStringProperty(DeviceGenericPropertyHwAddress)
}

func (d *deviceGeneric) GetPropertyTypeDescription() (string, error) {
	return d.getStringProperty(DeviceGenericPropertyTypeDescription)
}

func (d *deviceGeneric) MarshalJSON() ([]byte, error) {
	m := d.device.marshalMap()
	m["HwAddress"], _ = d.GetPropertyHwAddress()
	m["TypeDescription"], _ = d.GetPropertyTypeDescription()
	return json.Marshal(m)
}
