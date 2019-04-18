package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDeviceProfileList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/deviceprofiles")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "id": "a362ddb6-349b-4650-9911-681b51069a57",
      "name": "thermostat1.model1",
      "owner": {
        "ownerType": "USER",
        "ownerId": "72fd3b54-d243-4bf7-a845-92eb7956b982"
      },
      "components": [
        {
          "id": "main",
          "capabilities": []
        }
      ],
      "metadata": null
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

	devicesProfiles, err := sc.DeviceProfileList()
	assert.Nil(t, err)
	for _, deviceProfile := range devicesProfiles.Items {
		assert.NotNil(t, deviceProfile)
		assert.NotEmpty(t, deviceProfile.Name)
	}
}

func TestDeviceProfileListByPage(t *testing.T) {

	link := "https://api.smartthings.com/v1/deviceprofiles?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/deviceprofiles?page=2")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "id": "a362ddb6-349b-4650-9911-681b51069a57",
      "name": "thermostat1.model1",
      "owner": {
        "ownerType": "USER",
        "ownerId": "72fd3b54-d243-4bf7-a845-92eb7956b982"
      },
      "components": [
        {
          "id": "main",
          "capabilities": []
        }
      ],
      "metadata": null
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

	devicesProfiles, err := sc.DeviceProfileListByPage(link)
	assert.Nil(t, err)
	for _, deviceProfile := range devicesProfiles.Items {
		assert.NotNil(t, deviceProfile)
		assert.NotEmpty(t, deviceProfile.Name)
	}
}
func TestDeviceProfileCreate(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {

		var createReq DeviceProfileCreate
		decoder := json.NewDecoder(req.Body)
		err2 := decoder.Decode(&createReq)
		assert.Nil(t, err2)

		expectedReq := []byte(`{
  "name": "thermostat1.model1",
  "components": [],
  "metadata": null
}`)
		var expectedBody DeviceProfileCreate
		json.Unmarshal(expectedReq, &expectedBody)
		assert.Equal(t, expectedBody, createReq)
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/deviceprofiles")
		assert.Equal(t, req.Method, http.MethodPost)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": "a362ddb6-349b-4650-9911-681b51069a57",
  "name": "thermostat1.model1",
  "owner": {
    "ownerType": "USER",
    "ownerId": "72fd3b54-d243-4bf7-a845-92eb7956b982"
  },
  "components": [
    {
      "id": "main",
      "capabilities": []
    }
  ],
  "metadata": null
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	sc := NewClientCustomHTTP(client)
	profileCreate := &DeviceProfileCreate{
		Name:       "thermostat1.model1",
		Components: []DeviceProfileComponent{},
	}
	profile, err := sc.DeviceProfileCreate(profileCreate)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.NotEmpty(t, profile.ID)
}
func TestDeviceProfileDelete(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/deviceprofiles/profile-id", req.URL.String())
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
	err := sc.DeviceProfileDelete("profile-id")
	assert.Nil(t, err)
}
