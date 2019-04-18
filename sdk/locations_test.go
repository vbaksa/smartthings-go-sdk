package sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLocationList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/locations")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "locationId": "6b3d1909-1e1c-43ec-adc2-5f941de4fbf9",
      "name": "Home"
    }
  ],
  "_links": {
    "next": {
      "href": "https://..."
    },
    "previous": {
      "href": "https://..."
    }
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	locations, err := sc.LocationsList()
	assert.Nil(t, err)
	assert.NotNil(t, locations.Items)
}
func TestLocationsListByPage(t *testing.T) {

	link := "https://api.smartthings.com/v1/locations?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/locations?page=2")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "locationId": "6b3d1909-1e1c-43ec-adc2-5f941de4fbf9",
      "name": "Home"
    }
  ],
  "_links": {
    "next": {
      "href": "https://..."
    },
    "previous": {
      "href": "https://..."
    }
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	locations, err := sc.LocationsListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, locations.Items)
}

func TestLocation(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/locations/location-id", req.URL.String())
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "locationId": "6b3d1909-1e1c-43ec-adc2-5f941de4fbf9",
  "name": "Home",
  "latitude": 45.00708112,
  "longitude": -93.11223629,
  "regionRadius": 150,
  "temperatureScale": "F",
  "timeZoneId": "America/Chicago",
  "locale": "en"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	location, err := sc.LocationGet("location-id")
	assert.Nil(t, err)
	assert.NotNil(t, location.Name)
	assert.NotNil(t, location.LocationID)
}
func TestLocationUpdate(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/locations/location-id", req.URL.String())
		assert.Equal(t, req.Method, http.MethodPut)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "locationId": "6b3d1909-1e1c-43ec-adc2-5f941de4fbf9",
  "name": "Home",
  "latitude": 45.00708112,
  "longitude": -93.11223629,
  "regionRadius": 150,
  "temperatureScale": "F",
  "timeZoneId": "America/Chicago",
  "locale": "en"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	locationUpdate := &LocationUpdate{
		Name: "Home",
	}
	location, err := sc.LocationUpdate("location-id", *locationUpdate)
	assert.Nil(t, err)
	assert.NotNil(t, location.Name)
	assert.NotNil(t, location.LocationID)
}

func TestLocationDelete(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/locations/location-id", req.URL.String())
		assert.Equal(t, req.Method, http.MethodDelete)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	err := sc.LocationDelete("location-id")
	assert.Nil(t, err)
}
