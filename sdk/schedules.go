package sdk

import (
	"fmt"
	"net/http"
)

//ScheduleList list all schedules
func (c *SmartThingsClient) ScheduleList(app string) (*ScheduleList, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps/%s/schedules", app), nil)
	if err != nil {
		return nil, err
	}
	var scheduleList ScheduleList
	_, err = c.do(req, &scheduleList)
	return &scheduleList, err
}

//ScheduleListByPage list schedules by page URL
func (c *SmartThingsClient) ScheduleListByPage(href string) (*ScheduleList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var scheduleList ScheduleList
	_, err = c.do(req, &scheduleList)
	return &scheduleList, err
}

//ScheduleGet get schedule by schedule ID
func (c *SmartThingsClient) ScheduleGet(app string, scheduleID string) (*Schedule, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps/%s/schedules/%s", app, scheduleID), nil)
	if err != nil {
		return nil, err
	}
	var schedule Schedule
	_, err = c.do(req, &schedule)
	return &schedule, err
}

//ScheduleDelete deletes schedule
func (c *SmartThingsClient) ScheduleDelete(app string, scheduleID string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/installedapps/%s/schedules/%s", app, scheduleID), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//ScheduleDeleteAll deletes all application schedules
func (c *SmartThingsClient) ScheduleDeleteAll(app string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/installedapps/%s/schedules", app), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//ScheduleCreate creates schedule
func (c *SmartThingsClient) ScheduleCreate(app string, scheduleCreate ScheduleCreate) (*Schedule, error) {
	req, err := c.newRequest(http.MethodPost, fmt.Sprintf("/v1/installedapps/%s/schedules", app), scheduleCreate)
	if err != nil {
		return nil, err
	}
	var schedule Schedule
	_, err = c.do(req, &schedule)
	return &schedule, err
}
