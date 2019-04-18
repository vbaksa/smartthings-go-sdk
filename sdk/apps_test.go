package sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAppList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "appName": "security-home-monitor",
      "appId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
      "appType": "WEBHOOK_SMART_APP",
      "displayName": "Security Home Monitor",
      "description": "An integration that leverages all security devices in your home - cameras, motion sensors, etc to provide\nyou with peace of mind and #security.\n",
      "owner": {
        "ownerType": "USER",
        "ownerId": "3f178632-79ed-4358-94de-280dbd7c206c"
      },
      "createdDate": "2019-03-30T14:20:41.733Z",
      "lastUpdatedDate": "2019-03-30T14:20:41.733Z"
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

	appList, err := sc.AppList()
	assert.Nil(t, err)
	assert.NotNil(t, appList.Items)
}

func TestAppListByPage(t *testing.T) {
	link := "https://api.smartthings.com/v1/apps?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps?page=2")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "appName": "security-home-monitor",
      "appId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
      "appType": "WEBHOOK_SMART_APP",
      "displayName": "Security Home Monitor",
      "description": "An integration that leverages all security devices in your home - cameras, motion sensors, etc to provide\nyou with peace of mind and #security.\n",
      "owner": {
        "ownerType": "USER",
        "ownerId": "3f178632-79ed-4358-94de-280dbd7c206c"
      },
      "createdDate": "2019-03-30T14:20:41.733Z",
      "lastUpdatedDate": "2019-03-30T14:20:41.733Z"
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

	appList, err := sc.AppListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, appList.Items)
}

func TestAppGet(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps/app-id")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "appName": "security-home-monitor",
  "appId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
  "appType": "WEBHOOK_SMART_APP",
  "displayName": "Security Home Monitor",
  "description": "An intergration that leverages all security devices in your home - cameras, motion sensors, etc to provide\nyou with peace of mind and #security.\n",
  "owner": {
    "ownerType": "USER",
    "ownerId": "3f178632-79ed-4358-94de-280dbd7c206c"
  },
  "webhookSmartApp": {
    "targetUrl": "https://securityhome.com/smartthings",
    "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlz00iTNPXIo+RM8n0Eudnh84XDWBFAkR\nL7mkOtXvZiKViCR1qs5LT3SaNy62hkUroexCoii4vjRTSHaQ9FwyBRo98EY9U3qzWmr9Row9ONcQ\nVKFVX6qiwb9SZG7QvHSfqVAuWhMGb1fFGUWl94L7kY0CihF+WJS5fdhXc1h/f/tY1439hruL3+BR\n/idfYjAKDa1+euW4b0HKMi4BdJutL/G5FXjDP7MP5fTRuM52erKpON1pwPD8ow/03MhK4wRikHqS\n/Kvd3UnQ8Q7OSrfJevqR0CqBf83Am1HSHh/bqOfsG0epl+2I1RKLBzB2tSTvMfk0ftiRSq/fRrhl\nxsJmFQIDAQAB\n-----END PUBLIC KEY-----\n"
  },
  "createdDate": "2019-03-30T14:20:41.733Z",
  "lastUpdatedDate": "2019-03-30T14:20:41.733Z"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	app, err := sc.AppGet("app-id")
	assert.Nil(t, err)
	assert.NotNil(t, app.AppName)
}

func TestAppCreate(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps")
		assert.Equal(t, req.Method, http.MethodPost)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "app": {
    "appName": "security-home-monitor",
    "appId": "c71b0a5b-7dab-41fe-b75f-7cf30246468e",
    "appType": "WEBHOOK_SMART_APP",
    "displayName": "Security Home Monitor",
    "description": "An intergration that leverages all security devices in your home - cameras, motion sensors, etc to provide\nyou with peace of mind and #security.\n",
    "owner": {
      "ownerType": "USER",
      "ownerId": "3f178632-79ed-4358-94de-280dbd7c206c"
    },
    "webhookSmartApp": {
      "targetUrl": "https://securityhome.com/smartthings",
      "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlz00iTNPXIo+RM8n0Eudnh84XDWBFAkR\nL7mkOtXvZiKViCR1qs5LT3SaNy62hkUroexCoii4vjRTSHaQ9FwyBRo98EY9U3qzWmr9Row9ONcQ\nVKFVX6qiwb9SZG7QvHSfqVAuWhMGb1fFGUWl94L7kY0CihF+WJS5fdhXc1h/f/tY1439hruL3+BR\n/idfYjAKDa1+euW4b0HKMi4BdJutL/G5FXjDP7MP5fTRuM52erKpON1pwPD8ow/03MhK4wRikHqS\n/Kvd3UnQ8Q7OSrfJevqR0CqBf83Am1HSHh/bqOfsG0epl+2I1RKLBzB2tSTvMfk0ftiRSq/fRrhl\nxsJmFQIDAQAB\n-----END PUBLIC KEY-----\n"
    },
    "createdDate": "2019-03-30T14:20:41.733Z",
    "lastUpdatedDate": "2019-03-30T14:20:41.733Z"
  },
  "oauthClientId": "7cd4d474-7b36-4e03-bbdb-4cd4ae45a2be",
  "oauthClientSecret": "9b3fd445-42d6-441b-b386-99ea51e13cb0"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	appCreate := AppCreate{
		AppName: "Test",
	}
	app, err := sc.AppCreate(appCreate)
	assert.Nil(t, err)
	assert.NotNil(t, app.AppName)
}
func TestAppUpdate(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/apps", req.URL.String())
		assert.Equal(t, http.MethodPut, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "appName": "string",
  "displayName": "string",
  "description": "string",
  "singleInstance": false,
  "appType": "LAMBDA_SMART_APP",
  "classifications": "AUTOMATION",
  "lambdaSmartApp": {
    "functions": [
      "arn:aws:lambda:eu-central-1:account-id:function:function-name:alias-name",
      "arn:aws:lambda:ap-southeast-1:account-id:function:function-name:alias-name"
    ]
  },
  "webhookSmartApp": {
    "targetUrl": "https://securityhome.com/smartthings"
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	appCreate := AppCreate{
		AppName: "Test",
	}
	app, err := sc.AppUpdate(appCreate)
	assert.Nil(t, err)
	assert.NotNil(t, app.AppName)
}

func TestAppDelete(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/apps/app-id", req.URL.String())
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
	err := sc.AppDelete("app-id")
	assert.Nil(t, err)
}

func TestAppGetSettings(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps/app-id/settings")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "settings": {
    "mySecretData": "6c21f9ee-7634-413f-89f7-bdfdce88bf0e"
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	app, err := sc.AppGetSettings("app-id")
	assert.Nil(t, err)
	assert.NotNil(t, app.Settings)
	for k, v := range app.Settings.Properties {
		assert.NotEmpty(t, k)
		assert.NotEmpty(t, v)
	}
}

func TestAppUpdateSettings(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps/app-id/settings")
		assert.Equal(t, req.Method, http.MethodPut)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "settings": {
    "mySecretData": "6c21f9ee-7634-413f-89f7-bdfdce88bf0e"
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	settings := AppSettings{
		Settings: struct{ Properties map[string]string }{Properties: map[string]string{}},
	}
	app, err := sc.AppUpdateSettings("app-id", settings)
	assert.Nil(t, err)
	assert.NotNil(t, app.Settings)
	for k, v := range app.Settings.Properties {
		assert.NotEmpty(t, k)
		assert.NotEmpty(t, v)
	}
}
func TestAppGetOAuthSettings(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps/app-id/oauth")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "clientName": "My SmartThings Integration",
  "scope": [
    "r:installedapps",
    "w:installedapps",
    "x:installedapps"
  ]
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	settings, err := sc.AppGetOauthSettings("app-id")
	assert.Nil(t, err)
	assert.NotNil(t, settings.ClientName)
	assert.NotEmpty(t, settings.Scope)
}

func TestAppUpdateOAuthSettings(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/apps/app-id/oauth")
		assert.Equal(t, req.Method, http.MethodPut)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "clientName": "My SmartThings Integration",
  "scope": [
    "r:installedapps",
    "w:installedapps",
    "x:installedapps"
  ]
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	oauthSettings := AppOauthSettings{
		ClientName: "Client name",
		Scope: []string{
			"r:installedapps",
		},
	}
	settings, err := sc.AppUpdateOauthSettings("app-id", oauthSettings)
	assert.Nil(t, err)
	assert.NotNil(t, settings.ClientName)
	assert.NotEmpty(t, settings.Scope)
}
func TestAppGenerateOAuthSettings(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/apps/app-id/oauth/generate", req.URL.String())
		assert.Equal(t, http.MethodPost, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "oauthClientDetails": {
    "clientName": "My SmartThings Integration",
    "scope": [
      "r:installedapps",
      "w:installedapps",
      "x:installedapps"
    ]
  },
  "oauthClientId": "7cd4d474-7b36-4e03-bbdb-4cd4ae45a2be",
  "oauthClientSecret": "9b3fd445-42d6-441b-b386-99ea51e13cb0"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	oauthSettings := AppOauthSettings{
		ClientName: "Client name",
		Scope: []string{
			"r:installedapps",
		},
	}
	settings, err := sc.AppGenerateOauthSettings("app-id", oauthSettings)
	assert.Nil(t, err)
	assert.NotNil(t, settings.OauthClientSecret)
	assert.NotEmpty(t, settings.OauthClientID)
}
