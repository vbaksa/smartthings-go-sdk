package sdk

import (
	"fmt"
	"net/http"
)

//AppList list all installed applications
func (c *SmartThingsClient) AppList() (*AppList, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/apps"), nil)
	if err != nil {
		return nil, err
	}
	var appList AppList
	_, err = c.do(req, &appList)
	return &appList, err
}

//AppListByPage list applications by page URL
func (c *SmartThingsClient) AppListByPage(href string) (*AppList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var appList AppList
	_, err = c.do(req, &appList)
	return &appList, err
}

//AppGet get application details
func (c *SmartThingsClient) AppGet(app string) (*App, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/apps/%s", app), nil)
	if err != nil {
		return nil, err
	}
	var appResult App
	_, err = c.do(req, &appResult)
	return &appResult, err
}

//AppCreate create new application
func (c *SmartThingsClient) AppCreate(appCreate AppCreate) (*App, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/apps"), appCreate)
	if err != nil {
		return nil, err
	}
	var appResult App
	_, err = c.do(req, &appResult)
	return &appResult, err
}

//AppUpdate update existing application details
func (c *SmartThingsClient) AppUpdate(appCreate AppCreate) (*App, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/v1/apps"), appCreate)
	if err != nil {
		return nil, err
	}
	var appResult App
	_, err = c.do(req, &appResult)
	return &appResult, err
}

//AppDelete delete installed application details
func (c *SmartThingsClient) AppDelete(app string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/apps/%s", app), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//AppGetSettings get application settings
func (c *SmartThingsClient) AppGetSettings(app string) (*AppSettings, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/apps/%s/settings", app), nil)
	if err != nil {
		return nil, err
	}
	var settings AppSettings
	_, err = c.do(req, &settings)
	return &settings, err
}

//AppUpdateSettings update application settings
func (c *SmartThingsClient) AppUpdateSettings(appID string, settings AppSettings) (*AppSettings, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/v1/apps/%s/settings", appID), settings)
	if err != nil {
		return nil, err
	}
	var appSettings AppSettings
	_, err = c.do(req, &appSettings)
	return &appSettings, err
}

//AppGetOauthSettings get oauth configuration
func (c *SmartThingsClient) AppGetOauthSettings(app string) (*AppOauthSettings, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/apps/%s/oauth", app), nil)
	if err != nil {
		return nil, err
	}
	var oauthSettings AppOauthSettings
	_, err = c.do(req, &oauthSettings)
	return &oauthSettings, err
}

//AppUpdateOauthSettings update oauth settings
func (c *SmartThingsClient) AppUpdateOauthSettings(app string, settings AppOauthSettings) (*AppOauthSettings, error) {
	req, err := c.newRequest(http.MethodPut, fmt.Sprintf("/v1/apps/%s/oauth", app), settings)
	if err != nil {
		return nil, err
	}
	var oauthSettings AppOauthSettings
	_, err = c.do(req, &oauthSettings)
	return &oauthSettings, err
}

//AppGenerateOauthSettings generate oauth settings
func (c *SmartThingsClient) AppGenerateOauthSettings(app string, settings AppOauthSettings) (*AppOauthSecret, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/apps/%s/oauth/generate", app), settings)
	if err != nil {
		return nil, err
	}
	var oauthSecret AppOauthSecret
	_, err = c.do(req, &oauthSecret)
	return &oauthSecret, err
}
