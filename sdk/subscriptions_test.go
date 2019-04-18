package sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSubscriptionList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/installedapps/app-id/subscriptions")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "id": "5e1b134b-bd85-4125-9c25-4a8291e754aa",
      "installedAppId": "fb05c874-cf1d-406a-930c-69a081e0eaee",
      "sourceType": "DEVICE",
      "device": {
        "componentId": "main",
        "deviceId": "e457978e-5e37-43e6-979d-18112e12c961,",
        "capability": "contactSensor,",
        "attribute": "contact,",
        "stateChangeOnly": "true,",
        "subscriptionName": "contact_subscription',",
        "value": "*"
      }
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

	subscriptions, err := sc.SubscriptionList("app-id")
	assert.Nil(t, err)
	assert.NotNil(t, subscriptions.Items)
}

func TestSubscriptionListByPage(t *testing.T) {
	link := "https://api.smartthings.com/v1/installedapps/app-id/subscriptions?page=1"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/installedapps/app-id/subscriptions?page=1")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "id": "5e1b134b-bd85-4125-9c25-4a8291e754aa",
      "installedAppId": "fb05c874-cf1d-406a-930c-69a081e0eaee",
      "sourceType": "DEVICE",
      "device": {
        "componentId": "main",
        "deviceId": "e457978e-5e37-43e6-979d-18112e12c961,",
        "capability": "contactSensor,",
        "attribute": "contact,",
        "stateChangeOnly": "true,",
        "subscriptionName": "contact_subscription',",
        "value": "*"
      }
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

	subscriptions, err := sc.SubscriptionListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, subscriptions.Items)
}

func TestSubscriptionDelete(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/subscriptions/id", req.URL.String())
		assert.Equal(t, http.MethodDelete, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "count": 0
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	err := sc.SubscriptionDelete("app-id", "id")
	assert.Nil(t, err)
}
func TestSubscriptionDeleteAll(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/subscriptions", req.URL.String())
		assert.Equal(t, http.MethodDelete, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "count": 0
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	err := sc.SubscriptionDeleteAll("app-id")
	assert.Nil(t, err)
}

func TestSubscriptionGet(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "https://api.smartthings.com/v1/installedapps/app-id/subscriptions/id")
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "id": "5e1b134b-bd85-4125-9c25-4a8291e754aa",
  "installedAppId": "fb05c874-cf1d-406a-930c-69a081e0eaee",
  "sourceType": "DEVICE",
  "device": {
    "componentId": "main",
    "deviceId": "e457978e-5e37-43e6-979d-18112e12c961,",
    "capability": "contactSensor,",
    "attribute": "contact,",
    "stateChangeOnly": "true,",
    "subscriptionName": "contact_subscription',",
    "value": "*"
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	subscription, err := sc.SubscriptionGet("app-id", "id")
	assert.Nil(t, err)
	assert.NotNil(t, subscription.ID)
}
