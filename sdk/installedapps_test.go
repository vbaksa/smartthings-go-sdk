package sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInstalledAppList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/installedapps")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "installedAppId": "8885b9c3-a50d-427e-ad07-d7a2f94421af",
      "installedAppType": "WEBHOOK_SMART_APP",
      "installedAppStatus": "PENDING",
      "displayName": "Color App",
      "appId": "fdd21b3e-5880-44c8-9f61-ce474814f587",
      "referenceId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
      "locationId": "4ae86cfc-63be-4bcc-8514-836c873b4d7a",
      "owner": {
        "ownerType": "USER",
        "ownerId": "3f178632-79ed-4358-94de-280dbd7c206c"
      },
      "createdDate": "2019-03-30T14:20:41.732Z",
      "lastUpdatedDate": "2019-03-30T14:20:41.732Z"
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

	appList, err := sc.InstalledAppList()
	assert.Nil(t, err)
	assert.NotNil(t, appList.Items)
}

func TestInstalledAppListByPage(t *testing.T) {
	link := "https://api.smartthings.com/v1/installedapps?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/installedapps?page=2")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "installedAppId": "8885b9c3-a50d-427e-ad07-d7a2f94421af",
      "installedAppType": "WEBHOOK_SMART_APP",
      "installedAppStatus": "PENDING",
      "displayName": "Color App",
      "appId": "fdd21b3e-5880-44c8-9f61-ce474814f587",
      "referenceId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
      "locationId": "4ae86cfc-63be-4bcc-8514-836c873b4d7a",
      "owner": {
        "ownerType": "USER",
        "ownerId": "3f178632-79ed-4358-94de-280dbd7c206c"
      },
      "createdDate": "2019-03-30T14:20:41.732Z",
      "lastUpdatedDate": "2019-03-30T14:20:41.732Z"
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

	appList, err := sc.InstalledAppListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, appList.Items)
}
func TestInstalledAppDelete(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id", req.URL.String())
		assert.Equal(t, http.MethodDelete, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	err := sc.InstalledAppDelete("app-id")
	assert.Nil(t, err)
}

func TestInstalledAppConfigList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/configs", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "installedAppId": "8885b9c3-a50d-427e-ad07-d7a2f94421af",
      "configurationId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
      "configurationStatus": "DONE",
      "createdDate": "2019-03-30T14:20:41.732Z",
      "lastUpdatedDate": "2019-03-30T14:20:41.732Z"
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

	appList, err := sc.InstalledConfigList("app-id")
	assert.Nil(t, err)
	assert.NotNil(t, appList.Items)
}

func TestInstalledAppConfigListByPage(t *testing.T) {
	link := "https://api.smartthings.com/v1/installedapps/app-id/configs?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/configs?page=2", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "installedAppId": "8885b9c3-a50d-427e-ad07-d7a2f94421af",
      "configurationId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
      "configurationStatus": "DONE",
      "createdDate": "2019-03-30T14:20:41.732Z",
      "lastUpdatedDate": "2019-03-30T14:20:41.732Z"
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

	appList, err := sc.InstalledAppConfigListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, appList.Items)
}

func TestInstalledAppConfig(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/installedapps/app-id/configs/config-id")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "installedAppId": "8885b9c3-a50d-427e-ad07-d7a2f94421af",
  "configurationId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
  "configurationStatus": "AUTHORIZED",
  "createdDate": "2019-03-30T14:20:41.732Z",
  "lastUpdatedDate": "2019-03-30T14:20:41.732Z",
  "config": {
    "switches": [
      {
        "valueType": "DEVICE",
        "deviceConfig": {
          "deviceId": "5ccbf4ba-7b5f-4eb3-85b8-ab8705ff8d61",
          "componentId": "main",
          "permissions": [
            "r:devices"
          ]
        }
      }
    ]
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	config, err := sc.InstalledConfig("app-id", "config-id")
	assert.Nil(t, err)
	assert.NotNil(t, config.ConfigurationID)
}
