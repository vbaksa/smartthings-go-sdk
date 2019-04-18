package sdk

import "time"

//Device holds device structure
type Device struct {
	DeviceID               string `json:"deviceId"`
	Name                   string `json:"name"`
	Label                  string `json:"label"`
	DeviceManufacturerCode string `json:"deviceManufacturerCode"`
	LocationID             string `json:"locationId"`
	DeviceTypeID           string `json:"deviceTypeId"`
	DeviceTypeName         string `json:"deviceTypeName"`
	DeviceNetworkType      string `json:"deviceNetworkType"`
	Components             []struct {
		ID           string                    `json:"id"`
		Capabilities []DeviceProfileCapability `json:"capabilities"`
	} `json:"components"`
	App struct {
		InstalledAppID string `json:"installedAppId"`
		ExternalID     string `json:"externalId"`
		Profile        struct {
			ID string `json:"id"`
		} `json:"profile"`
	} `json:"app"`
	Dth struct {
		DeviceTypeID      string `json:"deviceTypeId"`
		DeviceTypeName    string `json:"deviceTypeName"`
		DeviceNetworkType string `json:"deviceNetworkType"`
	} `json:"dth"`
	Type string `json:"type"`
}

//DeviceList holds devices array
type DeviceList struct {
	Devices []Device `json:"items"`
	Links   struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//DeviceInstall device installation details
type DeviceInstall struct {
	Label      string           `json:"label"`
	LocationID string           `json:"locationId"`
	App        DeviceInstallApp `json:"app"`
}

//DeviceInstallApp device application install details
type DeviceInstallApp struct {
	ProfileID      string `json:"profileId"`
	InstalledAppID string `json:"installedAppId"`
	ExternalID     string `json:"externalId"`
}

//DeviceCommand device command request
type DeviceCommand struct {
	Component  string                   `json:"component"`
	Capability string                   `json:"capability"`
	Command    string                   `json:"command"`
	Arguments  []map[string]interface{} `json:"arguments"`
}

//Commands commands array
type Commands struct {
	DeviceCommands []DeviceCommand `json:"commands"`
}

//DeviceEventList structure holds all device events
type DeviceEventList struct {
	DeviceEvents []DeviceEvent `json:"deviceEvents"`
}

//DeviceEvent device event details
type DeviceEvent struct {
	Component  string      `json:"component"`
	Capability string      `json:"capability"`
	Attribute  string      `json:"attribute"`
	Value      interface{} `json:"value"`
	Unit       interface{} `json:"unit"`
}

//DeviceStatus device component status details
type DeviceStatus struct {
	Component map[string]DeviceComponentStatus `json:"components"`
}

//DeviceComponentStatus hold information about various device statuses
type DeviceComponentStatus struct {
	Actuator *struct {
	} `json:"actuator"`
	ColorControl *struct {
		Color struct {
			Value string `json:"value"`
		} `json:"color"`
		Hue struct {
			Value int `json:"value"`
		} `json:"hue"`
		Saturation struct {
			Value int `json:"value"`
		} `json:"saturation"`
	} `json:"colorControl"`
	ColorTemperature *struct {
		ColorTemperature struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"colorTemperature"`
	} `json:"colorTemperature"`
	Battery *struct {
		Battery struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"battery"`
	} `json:"battery"`
	Configuration *struct {
	} `json:"configuration"`
	ContactSensor *struct {
		Contact struct {
			Value string `json:"value"`
		} `json:"contact"`
	} `json:"contactSensor"`
	Button *struct {
		Button struct {
			Value string `json:"value"`
		} `json:"button"`
		NumberOfButtons struct {
			Value int `json:"value"`
		} `json:"numberOfButtons"`
		SupportedButtonValues struct {
			Value interface{} `json:"value"`
		} `json:"supportedButtonValues"`
	} `json:"button"`
	Bridge *struct {
	} `json:"bridge"`
	HealthCheck *struct {
		DeviceWatchDeviceStatus struct {
			Data struct {
			} `json:"data"`
			Value *string `json:"value"`
		} `json:"DeviceWatch-DeviceStatus"`
		CheckInterval struct {
			Data struct {
				HubHardwareID string `json:"hubHardwareId"`
				Protocol      string `json:"protocol"`
			} `json:"data"`
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"checkInterval"`
		HealthStatus struct {
			Data struct {
			} `json:"data"`
			Value string `json:"value"`
		} `json:"healthStatus"`
	} `json:"healthCheck"`
	OccupancySensor *struct {
		Occupancy struct {
			Value string `json:"value"`
		} `json:"occupancy"`
	} `json:"occupancySensor"`
	PresenceSensor *struct {
		Presence struct {
			Value string `json:"value"`
		} `json:"presence"`
	} `json:"presenceSensor"`
	PowerMeter *struct {
		Power struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"power"`
	} `json:"powerMeter"`
	IlluminanceMeasurement *struct {
		Illuminance struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"illuminance"`
	} `json:"illuminanceMeasurement"`
	MotionSensor *struct {
		Motion struct {
			Value string `json:"value"`
		} `json:"motion"`
	} `json:"motionSensor"`
	Light *struct {
		Switch struct {
			Value string `json:"value"`
		} `json:"switch"`
	} `json:"light"`
	ThreeAxis *struct {
		ThreeAxis struct {
			Unit  string `json:"unit"`
			Value []int  `json:"value"`
		} `json:"threeAxis"`
	} `json:"threeAxis"`
	TemperatureMeasurement *struct {
		Temperature struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"temperature"`
	} `json:"temperatureMeasurement"`
	Refresh *struct {
	} `json:"refresh"`
	RelativeHumidityMeasurement *struct {
		Humidity struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"humidity"`
	} `json:"relativeHumidityMeasurement"`
	Sensor *struct {
	} `json:"sensor"`
	Switch *struct {
		Switch struct {
			Value string `json:"value"`
		} `json:"switch"`
	} `json:"switch"`
	SwitchLevel *struct {
		Level struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"level"`
	} `json:"switchLevel"`
	WaterSensor *struct {
		Water struct {
			Value string `json:"value"`
		} `json:"water"`
	} `json:"waterSensor"`
}

//ColorControl color command request details
type ColorControl struct {
	Hue        int    `json:"hue,omitempty"`
	Saturation int    `json:"saturation,omitempty"`
	Hex        string `json:"hex,omitempty"`
	Level      int    `json:"level,omitempty"`
	Switch     string `json:"switch,omitempty"`
}

//DeviceProfileList holds details about all device profiles
type DeviceProfileList struct {
	Items []DeviceProfile `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//DeviceProfileCapability stores device capability name and ID
type DeviceProfileCapability struct {
	ID      string `json:"id"`
	Version int    `json:"version"`
}

//DeviceProfileComponent stores capability array
type DeviceProfileComponent struct {
	ID           string                    `json:"id"`
	Capabilities []DeviceProfileCapability `json:"capabilities"`
}

//DeviceProfile stores device profile details
type DeviceProfile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		OwnerType string `json:"ownerType"`
		OwnerID   string `json:"ownerId"`
	} `json:"owner"`
	Components []DeviceProfileComponent `json:"components"`
	Metadata   DeviceMetadata           `json:"metadata"`
}

//DeviceMetadata stores device metadata details
type DeviceMetadata struct {
	DthName      string `json:"dthName"`
	DthNamespace string `json:"dthNamespace"`
}

//DeviceProfileCreate request to create device profile
type DeviceProfileCreate struct {
	Name       string                   `json:"name"`
	Components []DeviceProfileComponent `json:"components"`
	Metadata   DeviceMetadata           `json:"metadata"`
}

//LocationList location list response structure
type LocationList struct {
	Items []struct {
		LocationID string `json:"locationId"`
		Name       string `json:"name"`
	} `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//Location details structure
type Location struct {
	LocationID       string  `json:"locationId"`
	Name             string  `json:"name"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	RegionRadius     int     `json:"regionRadius"`
	TemperatureScale string  `json:"temperatureScale"`
	TimeZoneID       string  `json:"timeZoneId"`
	Locale           string  `json:"locale"`
}

//LocationUpdate location update request
type LocationUpdate struct {
	Name             string  `json:"name"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	RegionRadius     int     `json:"regionRadius"`
	TemperatureScale string  `json:"temperatureScale"`
	Locale           string  `json:"locale"`
}

//AppCreate application creation request
type AppCreate struct {
	AppName         string         `json:"appName"`
	DisplayName     string         `json:"displayName"`
	Description     string         `json:"description"`
	SingleInstance  bool           `json:"singleInstance"`
	AppType         string         `json:"appType"`
	Classifications string         `json:"classifications,omitempty"`
	LambdaSmartApp  LambdaSmartApp `json:"lambdaSmartApp"`
	WebhookSmartApp struct {
		TargetURL string `json:"targetUrl"`
	} `json:"webhookSmartApp,omitempty"`
}

//LambdaSmartApp smartapp lambda function ARN
type LambdaSmartApp struct {
	Functions []string `json:"functions"`
}

//AppList application list structure
type AppList struct {
	Items []struct {
		AppName     string `json:"appName"`
		AppID       string `json:"appId"`
		AppType     string `json:"appType"`
		DisplayName string `json:"displayName"`
		Description string `json:"description"`
		Owner       struct {
			OwnerType string `json:"ownerType"`
			OwnerID   string `json:"ownerId"`
		} `json:"owner"`
		CreatedDate     time.Time `json:"createdDate"`
		LastUpdatedDate time.Time `json:"lastUpdatedDate"`
	} `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//App holds application details
type App struct {
	AppName     string `json:"appName"`
	AppID       string `json:"appId"`
	AppType     string `json:"appType"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Owner       struct {
		OwnerType string `json:"ownerType"`
		OwnerID   string `json:"ownerId"`
	} `json:"owner"`
	WebhookSmartApp struct {
		TargetURL string `json:"targetUrl"`
		PublicKey string `json:"publicKey"`
	} `json:"webhookSmartApp"`
	CreatedDate     time.Time `json:"createdDate"`
	LastUpdatedDate time.Time `json:"lastUpdatedDate"`
}

//AppSettings holds application configuration settings
type AppSettings struct {
	Settings struct {
		Properties map[string]string
	} `json:"settings"`
}

//AppOauthSettings app oauth settings
type AppOauthSettings struct {
	ClientName string   `json:"clientName"`
	Scope      []string `json:"scope"`
}

//AppOauthSecret secret details
type AppOauthSecret struct {
	OauthClientDetails struct {
		ClientName string   `json:"clientName"`
		Scope      []string `json:"scope"`
	} `json:"oauthClientDetails"`
	OauthClientID     string `json:"oauthClientId"`
	OauthClientSecret string `json:"oauthClientSecret"`
}

//InstalledAppList installed applications list
type InstalledAppList struct {
	Items []struct {
		InstalledAppID     string `json:"installedAppId"`
		InstalledAppType   string `json:"installedAppType"`
		InstalledAppStatus string `json:"installedAppStatus"`
		DisplayName        string `json:"displayName"`
		AppID              string `json:"appId"`
		ReferenceID        string `json:"referenceId"`
		LocationID         string `json:"locationId"`
		Owner              struct {
			OwnerType string `json:"ownerType"`
			OwnerID   string `json:"ownerId"`
		} `json:"owner"`
		CreatedDate     time.Time `json:"createdDate"`
		LastUpdatedDate time.Time `json:"lastUpdatedDate"`
	} `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//InstalledAppConfigList installed application config list
type InstalledAppConfigList struct {
	Items []struct {
		InstalledAppID      string    `json:"installedAppId"`
		ConfigurationID     string    `json:"configurationId"`
		ConfigurationStatus string    `json:"configurationStatus"`
		CreatedDate         time.Time `json:"createdDate"`
		LastUpdatedDate     time.Time `json:"lastUpdatedDate"`
	} `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//InstalledAppConfig application configuration
type InstalledAppConfig struct {
	InstalledAppID      string    `json:"installedAppId"`
	ConfigurationID     string    `json:"configurationId"`
	ConfigurationStatus string    `json:"configurationStatus"`
	CreatedDate         time.Time `json:"createdDate"`
	LastUpdatedDate     time.Time `json:"lastUpdatedDate"`
	Config              struct {
		Switches []struct {
			ValueType    string `json:"valueType"`
			DeviceConfig struct {
				DeviceID    string   `json:"deviceId"`
				ComponentID string   `json:"componentId"`
				Permissions []string `json:"permissions"`
			} `json:"deviceConfig"`
		} `json:"switches"`
	} `json:"config"`
}

//SubscriptionList stores all subscriptions
type SubscriptionList struct {
	Items []struct {
		ID             string `json:"id"`
		InstalledAppID string `json:"installedAppId"`
		SourceType     string `json:"sourceType"`
		Device         struct {
			ComponentID      string `json:"componentId"`
			DeviceID         string `json:"deviceId"`
			Capability       string `json:"capability"`
			Attribute        string `json:"attribute"`
			StateChangeOnly  string `json:"stateChangeOnly"`
			SubscriptionName string `json:"subscriptionName"`
			Value            string `json:"value"`
		} `json:"device"`
	} `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//DeviceSubscription device description details
type DeviceSubscription struct {
	ID             string `json:"id"`
	InstalledAppID string `json:"installedAppId"`
	SourceType     string `json:"sourceType"`
	Device         struct {
		ComponentID      string `json:"componentId"`
		DeviceID         string `json:"deviceId"`
		Capability       string `json:"capability"`
		Attribute        string `json:"attribute"`
		StateChangeOnly  string `json:"stateChangeOnly"`
		SubscriptionName string `json:"subscriptionName"`
		Value            string `json:"value"`
	} `json:"device"`
}

//SceneList stores all scenes
type SceneList struct {
	Items []struct {
		SceneID                      string `json:"sceneId"`
		SceneName                    string `json:"sceneName"`
		SceneIcon                    string `json:"sceneIcon"`
		SceneColor                   string `json:"sceneColor"`
		BehaviorID                   string `json:"behaviorId"`
		LocationID                   string `json:"locationId"`
		CreatedBy                    string `json:"createdBy"`
		CreatedDateMilliseconds      int64  `json:"createdDate"`
		LastUpdatedDateMilliseconds  int64  `json:"lastUpdatedDate"`
		LastExecutedDateMilliseconds int64  `json:"lastExecutedDate"`
		Editable                     bool   `json:"editable"`
		APIVersion                   string `json:"apiVersion"`
	} `json:"items"`
	Links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
}

//ScheduleList stores all schedules
type ScheduleList struct {
	Items []struct {
		ScheduledExecutions []int64 `json:"scheduledExecutions"`
		Name                string  `json:"name"`
		Cron                struct {
			Expression string `json:"expression"`
			Timezone   string `json:"timezone"`
		} `json:"cron"`
		InstalledAppID string `json:"installedAppId"`
		LocationID     string `json:"locationId"`
	} `json:"items"`
	Links struct {
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"_links"`
}

//Schedule schedule details
type Schedule struct {
	InstalledAppID      string  `json:"installedAppId"`
	ScheduledExecutions []int64 `json:"scheduledExecutions"`
	Name                string  `json:"name"`
	Cron                struct {
		Expression string `json:"expression"`
		Timezone   string `json:"timezone"`
	} `json:"cron"`
}

//ScheduleCreate schedule creation request structure
type ScheduleCreate struct {
	Name string       `json:"name"`
	Cron ScheduleCron `json:"cron"`
}

//ScheduleCron cron schedule details
type ScheduleCron struct {
	Expression string `json:"expression"`
	Timezone   string `json:"timezone"`
}
