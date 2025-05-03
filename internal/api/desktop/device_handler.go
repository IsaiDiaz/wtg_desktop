package desktop

import "wtg_desktop/internal/domain/device"

type DeviceHandler struct {
	service device.Service
}

func NewDeviceHandler(service device.Service) *DeviceHandler {
	return &DeviceHandler{service: service}
}

func (d *DeviceHandler) CreateDevice(dev *device.Device) error {
	return d.service.CreateDevice(dev)
}

func (d *DeviceHandler) GetDeviceByID(id int) (*device.Device, error) {
	return d.service.GetDeviceByID(id)
}

func (d *DeviceHandler) GetAllDevices() ([]*device.Device, error) {
	return d.service.GetAllDevices()
}

func (d *DeviceHandler) UpdateDevice(dev *device.Device) error {
	return d.service.UpdateDevice(dev)
}

func (d *DeviceHandler) DeleteDevice(id int) error {
	return d.service.DeleteDevice(id)
}

func (d *DeviceHandler) GetDevicesByEmployeeID(employeeID int) ([]*device.Device, error) {
	return d.service.GetDevicesByEmployeeID(employeeID)
}
