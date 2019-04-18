package sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestScheduleList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/schedules", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "scheduledExecutions": [
        1505752823000
      ],
      "name": "schedule-name",
      "cron": {
        "expression": "0/10 * * * ? *",
        "timezone": "America/Chicago"
      },
      "installedAppId": "937e11d5-317d-445f-bec7-3055fdb827a3",
      "locationId": "8418eebd-8d5f-48dd-a028-054744j8secb"
    }
  ],
  "_links": {
    "next": null,
    "previous": null
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	scheduleList, err := sc.ScheduleList("app-id")
	assert.Nil(t, err)
	assert.NotNil(t, scheduleList.Items)
}
func TestScheduleListByPage(t *testing.T) {
	link := "https://api.smartthings.com/v1/installedapps/app-id/schedules?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/schedules?page=2", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "scheduledExecutions": [
        1505752823000
      ],
      "name": "schedule-name",
      "cron": {
        "expression": "0/10 * * * ? *",
        "timezone": "America/Chicago"
      },
      "installedAppId": "937e11d5-317d-445f-bec7-3055fdb827a3",
      "locationId": "8418eebd-8d5f-48dd-a028-054744j8secb"
    }
  ],
  "_links": {
    "next": null,
    "previous": null
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	scheduleList, err := sc.ScheduleListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, scheduleList.Items)
}

func TestScheduleDeleteAll(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/schedules", req.URL.String())
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

	err := sc.ScheduleDeleteAll("app-id")
	assert.Nil(t, err)
}
func TestScheduleDelete(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/schedules/schedule-id", req.URL.String())
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

	err := sc.ScheduleDelete("app-id", "schedule-id")
	assert.Nil(t, err)
}

func TestScheduleGet(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/schedules/schedule-id", req.URL.String())
		assert.Equal(t, req.Method, http.MethodGet)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "installedAppId": "736e3903-001c-4d40-b408-ff40d162a06b",
  "scheduledExecutions": [
    1490892856362
  ],
  "name": "on_schedule-1",
  "cron": {
    "expression": "15 10 * * ? *",
    "timezone": "GMT"
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	schedule, err := sc.ScheduleGet("app-id", "schedule-id")
	assert.Nil(t, err)
	assert.NotNil(t, schedule.Name)
}

func TestScheduleCreate(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/installedapps/app-id/schedules", req.URL.String())
		assert.Equal(t, req.Method, http.MethodPost)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "installedAppId": "736e3903-001c-4d40-b408-ff40d162a06b",
  "scheduledExecutions": [
    1490892856362
  ],
  "name": "on_schedule-1",
  "cron": {
    "expression": "15 10 * * ? *",
    "timezone": "GMT"
  }
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)
	create := ScheduleCreate{
		Name: "on_schedule_1",
		Cron: ScheduleCron{
			Expression: "15 10 * * ? *",
			Timezone:   "GMT",
		},
	}
	schedule, err := sc.ScheduleCreate("app-id", create)
	assert.Nil(t, err)
	assert.NotNil(t, schedule.Name)
}
