package sdk

import (
	"fmt"
	"net/http"
)

//InstalledAppList list all installed applications
func (c *SmartThingsClient) InstalledAppList() (*InstalledAppList, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps"), nil)
	if err != nil {
		return nil, err
	}
	var installedAppList InstalledAppList
	_, err = c.do(req, &installedAppList)
	return &installedAppList, err
}

//InstalledAppListByPage list installed applications by page URl
func (c *SmartThingsClient) InstalledAppListByPage(href string) (*InstalledAppList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var installedAppList InstalledAppList
	_, err = c.do(req, &installedAppList)
	return &installedAppList, err
}

//InstalledAppDelete delete installed appplication
func (c *SmartThingsClient) InstalledAppDelete(app string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/installedapps/%s", app), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//InstalledConfigList get installed applications config
func (c *SmartThingsClient) InstalledConfigList(app string) (*InstalledAppConfigList, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps/%s/configs", app), nil)
	if err != nil {
		return nil, err
	}
	var installedConfig InstalledAppConfigList
	_, err = c.do(req, &installedConfig)
	return &installedConfig, err
}

//InstalledAppConfigListByPage get installed application config by page URL
func (c *SmartThingsClient) InstalledAppConfigListByPage(href string) (*InstalledAppConfigList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var installedAppConfig InstalledAppConfigList
	_, err = c.do(req, &installedAppConfig)
	return &installedAppConfig, err
}

//InstalledConfig get install application config by config ID
func (c *SmartThingsClient) InstalledConfig(app string, configID string) (*InstalledAppConfig, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps/%s/configs/%s", app, configID), nil)
	if err != nil {
		return nil, err
	}
	var installedConfig InstalledAppConfig
	_, err = c.do(req, &installedConfig)
	return &installedConfig, err
}
