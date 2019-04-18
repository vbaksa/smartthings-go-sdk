package sdk

import (
	"fmt"
	"net/http"
)

//DeviceProfileDelete delete device profile
func (c *SmartThingsClient) DeviceProfileDelete(profileID string) error {

	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/deviceprofiles/%s", profileID), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//DeviceProfileList list device profiles
func (c *SmartThingsClient) DeviceProfileList() (*DeviceProfileList, error) {

	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/deviceprofiles"), nil)
	if err != nil {
		return nil, err
	}
	var deviceProfiles DeviceProfileList
	_, err = c.do(req, &deviceProfiles)
	return &deviceProfiles, err
}

//DeviceProfileListByPage list device profiles by page URL
func (c *SmartThingsClient) DeviceProfileListByPage(href string) (*DeviceProfileList, error) {

	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var deviceProfileList DeviceProfileList
	_, err = c.do(req, &deviceProfileList)
	return &deviceProfileList, err
}

//DeviceProfileCreate create device profile
func (c *SmartThingsClient) DeviceProfileCreate(profile *DeviceProfileCreate) (*DeviceProfile, error) {

	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/deviceprofiles"), profile)
	if err != nil {
		return nil, err
	}
	var deviceProfile DeviceProfile
	_, err = c.do(req, &deviceProfile)
	return &deviceProfile, err
}
