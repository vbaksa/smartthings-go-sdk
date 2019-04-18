package sdk

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSceneList(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/scenes", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "sceneId": "string",
      "sceneName": "string",
      "sceneIcon": "string",
      "sceneColor": "string",
      "behaviorId": "string",
      "locationId": "string",
      "createdBy": "string",
      "createdDate": 1553985049,
      "lastUpdatedDate": 1553985049,
      "lastExecutedDate": 1553985049,
      "editable": true,
      "apiVersion": "string"
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

	sceneList, err := sc.SceneList()
	assert.Nil(t, err)
	assert.NotNil(t, sceneList.Items)
}
func TestSceneListByPage(t *testing.T) {
	link := "https://api.smartthings.com/v1/scenes?page=2"
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/scenes?page=2", req.URL.String())
		assert.Equal(t, http.MethodGet, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "items": [
    {
      "sceneId": "string",
      "sceneName": "string",
      "sceneIcon": "string",
      "sceneColor": "string",
      "behaviorId": "string",
      "locationId": "string",
      "createdBy": "string",
      "createdDate": 1553985049,
      "lastUpdatedDate": 1553985049,
      "lastExecutedDate": 1553985049,
      "editable": true,
      "apiVersion": "string"
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

	sceneList, err := sc.SceneListByPage(link)
	assert.Nil(t, err)
	assert.NotNil(t, sceneList.Items)
}
func TestSceneExecute(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "https://api.smartthings.com/v1/scenes/scene-id/execute", req.URL.String())
		assert.Equal(t, http.MethodPost, req.Method)
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{
  "status": "success"
}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	sc := NewClientCustomHTTP(client)

	err := sc.SceneExecute("scene-id")
	assert.Nil(t, err)
}
