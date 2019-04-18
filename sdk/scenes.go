package sdk

import (
	"fmt"
	"net/http"
)

//SceneList list all scenes
func (c *SmartThingsClient) SceneList() (*SceneList, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/scenes"), nil)
	if err != nil {
		return nil, err
	}
	var sceneList SceneList
	_, err = c.do(req, &sceneList)
	return &sceneList, err
}

//SceneListByPage list scenes by page URL
func (c *SmartThingsClient) SceneListByPage(href string) (*SceneList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var sceneList SceneList
	_, err = c.do(req, &sceneList)
	return &sceneList, err
}

//SceneExecute executes scene
func (c *SmartThingsClient) SceneExecute(sceneID string) error {

	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/scenes/%s/execute", sceneID), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}
