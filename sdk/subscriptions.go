package sdk

import (
	"fmt"
	"net/http"
)

//SubscriptionList get all available subscriptions
func (c *SmartThingsClient) SubscriptionList(app string) (*SubscriptionList, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps/%s/subscriptions", app), nil)
	if err != nil {
		return nil, err
	}
	var subscriptions SubscriptionList
	_, err = c.do(req, &subscriptions)
	return &subscriptions, err
}

//SubscriptionListByPage get all available subscriptions by page URL
func (c *SmartThingsClient) SubscriptionListByPage(href string) (*SubscriptionList, error) {
	req, err := http.NewRequest(http.MethodGet, href, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if err != nil {
		return nil, err
	}
	var subscriptions SubscriptionList
	_, err = c.do(req, &subscriptions)
	return &subscriptions, err
}

//SubscriptionGet gets subscriptions details
func (c *SmartThingsClient) SubscriptionGet(app string, subscriptionID string) (*DeviceSubscription, error) {
	req, err := c.newRequest(http.MethodGet, fmt.Sprintf("/v1/installedapps/%s/subscriptions/%s", app, subscriptionID), nil)
	if err != nil {
		return nil, err
	}
	var subscription DeviceSubscription
	_, err = c.do(req, &subscription)
	return &subscription, err
}

//SubscriptionDelete deletes subscription
func (c *SmartThingsClient) SubscriptionDelete(app string, subscriptionID string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/installedapps/%s/subscriptions/%s", app, subscriptionID), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}

//SubscriptionDeleteAll deletes all subscriptions
func (c *SmartThingsClient) SubscriptionDeleteAll(app string) error {
	req, err := c.newRequest(http.MethodDelete, fmt.Sprintf("/v1/installedapps/%s/subscriptions", app), nil)
	if err != nil {
		return err
	}
	_, err = c.do(req, nil)
	return err
}
