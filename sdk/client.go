package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

//SmartThingsClient hold sdk configuration structures
type SmartThingsClient struct {
	BaseURL   *url.URL
	UserAgent string

	HTTPClient *http.Client
}

//NewClient create new sdk instance
func NewClient(token string) *SmartThingsClient {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	c := &SmartThingsClient{HTTPClient: tc,
		UserAgent: "Smart House",
		BaseURL: &url.URL{Host: "api.smartthings.com",
			Scheme: "https"}}
	return c
}

//NewClientCustomHTTP create new custom sdk instance
func NewClientCustomHTTP(client *http.Client) *SmartThingsClient {

	c := &SmartThingsClient{HTTPClient: client,
		UserAgent: "Smart House",
		BaseURL: &url.URL{Host: "api.smartthings.com",
			Scheme: "https"}}
	return c
}

func (c *SmartThingsClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *SmartThingsClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if 200 != resp.StatusCode {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("expected status code 200, but API responded with %d . Response Body: %s", resp.StatusCode, bodyString)
	}
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
