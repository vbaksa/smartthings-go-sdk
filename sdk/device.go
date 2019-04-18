package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//DevicesList list all devices
func (c *SmartThingsClient) DevicesList() (*DeviceList, error) {

	req, err := c.newRequest(http.MethodGet, "/v1/devices", nil)
	if err != nil {
		return nil, err
	}
	var deviceList DeviceList
	_, err = c.do(req, &deviceList)
	return &deviceList, err
}

//DevicesListByPage list all devices by page URL
func (c *SmartThingsClient) DevicesListByPage(href string) (*DeviceList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var deviceList DeviceList
	_, err = c.do(req, &deviceList)
	return &deviceList, err
}

//DeviceDescription get device description
func (c *SmartThingsClient) DeviceDescription(deviceID string) (*Device, error) {

	req, err := c.newRequest(http.MethodGet, "/v1/devices/"+deviceID, nil)
	if err != nil {
		return nil, err
	}
	var device Device
	_, err = c.do(req, &device)
	return &device, err
}

//DeviceInstall install device
func (c *SmartThingsClient) DeviceInstall(deviceInstallReq DeviceInstall) (*Device, error) {

	req, err := c.newRequest(http.MethodPost, "/v1/devices", deviceInstallReq)
	if err != nil {
		return nil, err
	}
	var device Device
	_, err = c.do(req, &device)
	return &device, err
}

//DeviceDelete delete device
func (c *SmartThingsClient) DeviceDelete(deviceID string) error {

	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/devices/%s", deviceID), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//DeviceUpdateLabel update device label
func (c *SmartThingsClient) DeviceUpdateLabel(deviceID string, label string) (*Device, error) {

	lbl := struct {
		label string `json:"label"`
	}{
		label: label,
	}

	req, err := c.newRequest(http.MethodPut, "/v1/devices/"+deviceID, lbl)
	if err != nil {
		return nil, err
	}
	var device Device
	_, err = c.do(req, &device)
	return &device, err
}

//DeviceUpdateProperty update device property
func (c *SmartThingsClient) DeviceUpdateProperty(deviceID string, property string, value interface{}) (*Device, error) {

	res := map[string]interface{}{
		property: value,
	}
	fmt.Printf("lbl: %s\n", res)
	r, err := json.Marshal(res)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("lbl: %s \n", string(r))
	req, err := c.newRequest(http.MethodPut, "/v1/devices/"+deviceID, res)
	if err != nil {
		return nil, err
	}
	var device Device
	_, err = c.do(req, &device)
	return &device, err
}

//DeviceCommand run device command
func (c *SmartThingsClient) DeviceCommand(deviceID string, command ...DeviceCommand) error {
	s := make([]DeviceCommand, 0)
	for _, c := range command {
		s = append(s, c)
	}
	b := &Commands{
		DeviceCommands: s,
	}
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/devices/%s/commands", deviceID), b)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//DeviceCreateEvent create device event
func (c *SmartThingsClient) DeviceCreateEvent(deviceID string, command ...DeviceEvent) error {
	s := make([]DeviceEvent, 0)
	for _, c := range command {
		s = append(s, c)
	}
	b := &DeviceEventList{
		DeviceEvents: s,
	}
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/devices/%s/events", deviceID), b)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//DeviceFullStatus get device full status
func (c *SmartThingsClient) DeviceFullStatus(deviceID string) (*DeviceStatus, error) {

	req, err := c.newRequest(http.MethodGet, "/v1/devices/"+deviceID+"/"+"status", nil)
	if err != nil {
		return nil, err
	}
	var deviceStatus DeviceStatus
	_, err = c.do(req, &deviceStatus)
	return &deviceStatus, err
}

//DeviceComponentStatus get device component status
func (c *SmartThingsClient) DeviceComponentStatus(deviceID string, component string) (*DeviceComponentStatus, error) {

	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/devices/%s/components/%s/status", deviceID, component), nil)
	if err != nil {
		return nil, err
	}
	var deviceStatus DeviceComponentStatus
	_, err = c.do(req, &deviceStatus)
	return &deviceStatus, err
}

//DeviceSupportsCapability check device capability
func DeviceSupportsCapability(capability string, device Device) bool {
	for _, component := range device.Components {
		for _, deviceCapability := range component.Capabilities {
			if deviceCapability.ID == capability {
				return true
			}
		}
	}
	return false
}
