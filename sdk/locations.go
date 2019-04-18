package sdk

import (
	"fmt"
	"net/http"
)

//LocationsList get all locations
func (c *SmartThingsClient) LocationsList() (*LocationList, error) {

	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/locations"), nil)
	if err != nil {
		return nil, err
	}
	var locationList LocationList
	_, err = c.do(req, &locationList)
	return &locationList, err
}

//LocationsListByPage get all locations by page
func (c *SmartThingsClient) LocationsListByPage(href string) (*LocationList, error) {

	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var locationList LocationList
	_, err = c.do(req, &locationList)
	return &locationList, err
}

//LocationGet get location by ID
func (c *SmartThingsClient) LocationGet(locationID string) (*Location, error) {

	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/locations/%s", locationID), nil)
	if err != nil {
		return nil, err
	}
	var location Location
	_, err = c.do(req, &location)
	return &location, err
}

//LocationUpdate update location
func (c *SmartThingsClient) LocationUpdate(locationID string, locationUpdate LocationUpdate) (*Location, error) {

	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/v1/locations/%s", locationID), locationUpdate)
	if err != nil {
		return nil, err
	}
	var location Location
	_, err = c.do(req, &location)
	return &location, err
}

//LocationDelete delete location
func (c *SmartThingsClient) LocationDelete(locationID string) error {

	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/locations/%s", locationID), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}
