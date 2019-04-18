package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDeviceList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "deviceId": "6f5ea629-4c05-4a90-a244-cc129b0a80c3",
      "name": "color.light.100x",
      "label": "color.light.100x",
      "deviceManufacturerCode": "010F-0B01-2002",
      "locationId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
      "deviceTypeId": "Deprecated please look under \"dth\".",
      "deviceTypeName": "Deprecated please look under \"dth\".",
      "deviceNetworkType": "Deprecated please look under \"dth\".",
      "components": [
        {
          "id": "main",
          "capabilities": []
        }
      ],
      "app": {
        "installedAppId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
        "externalId": "Th13390",
        "profile": {
          "id": "user@example.com/thermostat.model1"
        }
      },
      "dth": {
        "deviceTypeId": "7b8514e6-230d-41cc-b3c2-512bca15abf0",
        "deviceTypeName": "x.com.samsung.da.fridge",
        "deviceNetworkType": "ZWAVE"
      },
      "type": "DTH"
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

	devices, err := sc.DevicesList()
	assert.Nil(t, err)
	for _, device := range devices.Devices {
		assert.NotNil(t, device)
		assert.NotEmpty(t, device.DeviceID)
	}
}
func TestDeviceListByPage(t *testing.T) {

	link := "https://api.smartthings.com/v1/devices?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/devices?page=2", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "deviceId": "6f5ea629-4c05-4a90-a244-cc129b0a80c3",
      "name": "color.light.100x",
      "label": "color.light.100x",
      "deviceManufacturerCode": "010F-0B01-2002",
      "locationId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
      "deviceTypeId": "Deprecated please look under \"dth\".",
      "deviceTypeName": "Deprecated please look under \"dth\".",
      "deviceNetworkType": "Deprecated please look under \"dth\".",
      "components": [
        {
          "id": "main",
          "capabilities": []
        }
      ],
      "app": {
        "installedAppId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
        "externalId": "Th13390",
        "profile": {
          "id": "user@example.com/thermostat.model1"
        }
      },
      "dth": {
        "deviceTypeId": "7b8514e6-230d-41cc-b3c2-512bca15abf0",
        "deviceTypeName": "x.com.samsung.da.fridge",
        "deviceNetworkType": "ZWAVE"
      },
      "type": "DTH"
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

	devices, err := sc.DevicesListByPage(link)
	assert.Nil(t, err)
	for _, device := range devices.Devices {
		assert.NotNil(t, device)
		assert.NotEmpty(t, device.DeviceID)
	}
}
func TestDeviceDescription(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "deviceId": "6f5ea629-4c05-4a90-a244-cc129b0a80c3",
  "name": "color.light.100x",
  "label": "color.light.100x",
  "deviceManufacturerCode": "010F-0B01-2002",
  "locationId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
  "deviceTypeId": "Deprecated please look under \"dth\".",
  "deviceTypeName": "Deprecated please look under \"dth\".",
  "deviceNetworkType": "Deprecated please look under \"dth\".",
  "components": [
    {
      "id": "main",
      "capabilities": []
    }
  ],
  "app": {
    "installedAppId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
    "externalId": "Th13390",
    "profile": {
      "id": "user@example.com/thermostat.model1"
    }
  },
  "dth": {
    "deviceTypeId": "7b8514e6-230d-41cc-b3c2-512bca15abf0",
    "deviceTypeName": "x.com.samsung.da.fridge",
    "deviceNetworkType": "ZWAVE"
  },
  "type": "DTH"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	device, err := sc.DeviceDescription("device-id")
	assert.Nil(t, err)
	assert.NotNil(t, device)
}

func TestDeviceInstall(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices")
		assert.Equal(t, req.Method, http.MethodPost)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "deviceId": "6f5ea629-4c05-4a90-a244-cc129b0a80c3",
  "name": "color.light.100x",
  "label": "color.light.100x",
  "deviceManufacturerCode": "010F-0B01-2002",
  "locationId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
  "deviceTypeId": "Deprecated please look under \"dth\".",
  "deviceTypeName": "Deprecated please look under \"dth\".",
  "deviceNetworkType": "Deprecated please look under \"dth\".",
  "components": [
    {
      "id": "main",
      "capabilities": []
    }
  ],
  "app": {
    "installedAppId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
    "externalId": "Th13390",
    "profile": {
      "id": "user@example.com/thermostat.model1"
    }
  },
  "dth": {
    "deviceTypeId": "7b8514e6-230d-41cc-b3c2-512bca15abf0",
    "deviceTypeName": "x.com.samsung.da.fridge",
    "deviceNetworkType": "ZWAVE"
  },
  "type": "DTH"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	deviceInstallReq := &DeviceInstall{
		Label:      "Test Label",
		LocationID: "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
		App: struct {
			ProfileID      string `json:"profileId"`
			InstalledAppID string `json:"installedAppId"`
			ExternalID     string `json:"externalId"`
		}{
			ProfileID:      "",
			InstalledAppID: "",
			ExternalID:     "",
		},
	}
	device, err := sc.DeviceInstall(*deviceInstallReq)
	assert.Nil(t, err)
	assert.NotNil(t, device)
}
func TestDeviceDelete(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id")
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
	err := sc.DeviceDelete("device-id")
	assert.Nil(t, err)
}

func TestDeviceUpdateLabel(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id")
		assert.Equal(t, req.Method, http.MethodPut)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "deviceId": "6f5ea629-4c05-4a90-a244-cc129b0a80c3",
  "name": "color.light.100x",
  "label": "color.light.100x",
  "deviceManufacturerCode": "010F-0B01-2002",
  "locationId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
  "deviceTypeId": "Deprecated please look under \"dth\".",
  "deviceTypeName": "Deprecated please look under \"dth\".",
  "deviceNetworkType": "Deprecated please look under \"dth\".",
  "components": [
    {
      "id": "main",
      "capabilities": []
    }
  ],
  "app": {
    "installedAppId": "0c0b935d-0616-4441-a0bf-da7aeec3dc0a",
    "externalId": "Th13390",
    "profile": {
      "id": "user@example.com/thermostat.model1"
    }
  },
  "dth": {
    "deviceTypeId": "7b8514e6-230d-41cc-b3c2-512bca15abf0",
    "deviceTypeName": "x.com.samsung.da.fridge",
    "deviceNetworkType": "ZWAVE"
  },
  "type": "DTH"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	device, err := sc.DeviceUpdateLabel("device-id", "test")
	assert.Nil(t, err)
	assert.NotNil(t, device)
}
func TestDeviceExecuteCommand(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id/commands")
		assert.Equal(t, req.Method, http.MethodPost)

		//	bodyBytes, _ := ioutil.ReadAll(req.Body)
		var hh Commands
		decoder := json.NewDecoder(req.Body)
		err2 := decoder.Decode(&hh)
		assert.Nil(t, err2)

		expectedResponse := []byte(`{
			"commands": [
		{
			"component": "main",
			"capability": "temperatureSetPoint",
			"command": "setTemperature",
			"arguments": [
			{
				"value": 5,
				"unit": "°C"
			}
		]
		}
	]
		}`)
		var expectedCommand Commands
		json.Unmarshal(expectedResponse, &expectedCommand)

		fmt.Println(expectedCommand.DeviceCommands[0].Arguments)
		assert.Equal(t, expectedCommand, hh)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	command := &DeviceCommand{
		Component:  "main",
		Capability: "temperatureSetPoint",
		Command:    "setTemperature",
		Arguments: []map[string]interface{}{
			{
				"value": int(5),
				"unit":  "°C",
			},
		},
	}
	err := sc.DeviceCommand("device-id", *command)
	assert.Nil(t, err)
}

func TestDeviceCreateEvent(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id/events")
		assert.Equal(t, req.Method, http.MethodPost)

		var hh DeviceEventList
		decoder := json.NewDecoder(req.Body)
		err2 := decoder.Decode(&hh)
		assert.Nil(t, err2)

		expectedResponse := []byte(`{
  "deviceEvents": [
    {
      "component": "main",
      "capability": "switchLevel",
      "attribute": "level",
      "value": 0,
      "unit": null
    }
  ]
}`)
		var expectedCommand DeviceEventList
		json.Unmarshal(expectedResponse, &expectedCommand)
		assert.Equal(t, expectedCommand, hh)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	event := &DeviceEvent{
		Component:  "main",
		Capability: "switchLevel",
		Attribute:  "level",
		Value:      0,
		Unit:       nil,
	}
	err := sc.DeviceCreateEvent("device-id", *event)
	assert.Nil(t, err)
}

func TestDeviceStatus(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id/status")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "components": {
    "main": {
      "switch": {
        "switch": {
          "value": "on"
        }
      },
      "switchLevel": {
        "level": {
          "value": 90
        }
      }
    }
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	deviceStatus, err := sc.DeviceFullStatus("device-id")
	assert.Nil(t, err)
	assert.NotNil(t, deviceStatus.Component)
	for k, v := range deviceStatus.Component {
		assert.NotNil(t, k)
		assert.NotEmpty(t, v)
	}
}

func TestDeviceComponentStatus(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/devices/device-id/components/main/status")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "switch": {
    "switch": {
      "value": "on"
    }
  },
  "switchLevel": {
    "level": {
      "value": 90
    }
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	deviceStatus, err := sc.DeviceComponentStatus("device-id", "main")
	assert.Nil(t, err)
	assert.NotNil(t, deviceStatus.Switch)
	assert.NotNil(t, deviceStatus.SwitchLevel)
	assert.Nil(t, deviceStatus.TemperatureMeasurement)

}
