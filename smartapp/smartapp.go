package smartapp

import "time"

//PingRequest ping request triggered by API
type PingRequest struct {
	Lifecycle   string    `json:"lifecycle"`
	ExecutionID string    `json:"executionId"`
	Locale      string    `json:"locale"`
	Version     string    `json:"version"`
	PingData    *PingData `json:"pingData"`
}

//PingResponse sent back by smart app
type PingResponse struct {
	PingData *PingData `json:"pingData"`
}

//PingData ping details
type PingData struct {
	Challenge string `json:"challenge"`
}

//ConfigurationRequest stores configuration request details
type ConfigurationRequest struct {
	Lifecycle   string `json:"lifecycle"`
	ExecutionID string `json:"executionId"`
	Locale      string `json:"locale"`
	Version     string `json:"version"`
	Client      struct {
		Os                 string        `json:"os"`
		Version            string        `json:"version"`
		Language           string        `json:"language"`
		SupportedTemplates []interface{} `json:"supportedTemplates"`
	} `json:"sdk"`
	ConfigurationData struct {
		InstalledAppID string                      `json:"installedAppId"`
		Phase          string                      `json:"phase"`
		PageID         string                      `json:"pageId"`
		PreviousPageID string                      `json:"previousPageId"`
		Config         map[string][]ConfigProperty `json:"config"`
	} `json:"configurationData"`
	Settings struct {
	} `json:"settings"`
}

//ConfigProperty configuration property details
type ConfigProperty struct {
	ValueType    string `json:"valueType"`
	DeviceConfig *struct {
		DeviceID    string `json:"deviceId"`
		ComponentID string `json:"componentId"`
	} `json:"deviceConfig,omitempty"`
	StringConfig *struct {
		Value string `json:"value"`
	} `json:"stringConfig,omitempty"`
}

//ConfigResponse configuration response
type ConfigResponse struct {
	ConfigurationData ConfigurationData `json:"configurationData,omitempty"`
}

//ConfigurationData configuration details
type ConfigurationData struct {
	Initialize *Initialize `json:"initialize,omitempty"`
	Page       *Page       `json:"page,omitempty"`
}

//Initialize payload details
type Initialize struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ID          string   `json:"id"`
	Permissions []string `json:"permissions"`
	FirstPageID string   `json:"firstPageId"`
}

//Page configuration page details
type Page struct {
	PageID         string      `json:"pageId"`
	Name           string      `json:"name"`
	NextPageID     interface{} `json:"nextPageId"`
	PreviousPageID interface{} `json:"previousPageId"`
	Complete       bool        `json:"complete"`
	Sections       []Sections  `json:"sections"`
}

//Sections store page sections
type Sections struct {
	Name     string            `json:"name"`
	Settings []SectionSettings `json:"settings"`
}

//SectionSettings stores sections settings details
type SectionSettings struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	Required     bool     `json:"required,omitempty"`
	Multiple     bool     `json:"multiple,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
	Permissions  []string `json:"permissions,omitempty"`
	DefaultValue string   `json:"defaultValue,omitempty"`
}

//InstalledApp installed application details
type InstalledApp struct {
	InstalledAppID      string                      `json:"installedAppId"`
	LocationID          string                      `json:"locationId"`
	Config              map[string][]ConfigProperty `json:"config"`
	Permissions         []string                    `json:"permissions"`
	PreviousConfig      map[string][]ConfigProperty `json:"previousConfig"`
	PreviousPermissions []string                    `json:"previousPermissions"`
}

//InstallRequest installation request
type InstallRequest struct {
	Lifecycle   string `json:"lifecycle"`
	ExecutionID string `json:"executionId"`
	Locale      string `json:"locale"`
	Version     string `json:"version"`
	InstallData struct {
		AuthToken    string       `json:"authToken"`
		RefreshToken string       `json:"refreshToken"`
		InstalledApp InstalledApp `json:"installedApp"`
	} `json:"installData"`
	Settings struct {
	} `json:"settings"`
}

//UpdateRequest update payload
type UpdateRequest struct {
	Lifecycle   string `json:"lifecycle"`
	ExecutionID string `json:"executionId"`
	Locale      string `json:"locale"`
	Version     string `json:"version"`
	UpdateData  struct {
		AuthToken    string       `json:"authToken"`
		RefreshToken string       `json:"refreshToken"`
		InstalledApp InstalledApp `json:"installedApp"`
	} `json:"updateData"`
	Settings struct {
	} `json:"settings"`
}

//UninstallRequest uninstall payload
type UninstallRequest struct {
	Lifecycle     string `json:"lifecycle"`
	ExecutionID   string `json:"executionId"`
	Locale        string `json:"locale"`
	Version       string `json:"version"`
	UninstallData struct {
		InstalledApp InstalledApp `json:"installedApp"`
	} `json:"uninstallData"`
	Settings struct {
	} `json:"settings"`
}

//Event smart app event
type Event struct {
	Lifecycle   string `json:"lifecycle"`
	ExecutionID string `json:"executionId"`
	Locale      string `json:"locale"`
	Version     string `json:"version"`
	EventData   struct {
		AuthToken    string       `json:"authToken"`
		InstalledApp InstalledApp `json:"installedApp"`
		Events       []struct {
			EventTime           time.Time `json:"eventTime"`
			EventType           string    `json:"eventType"`
			DeviceCommandsEvent struct {
				EventID    string `json:"eventId"`
				DeviceID   string `json:"deviceId"`
				ProfileID  string `json:"profileId"`
				ExternalID string `json:"externalId"`
				Commands   []struct {
					ComponentID string                   `json:"componentId"`
					Capability  string                   `json:"capability"`
					Command     string                   `json:"command"`
					Arguments   []map[string]interface{} `json:"arguments"`
				} `json:"commands"`
			} `json:"deviceCommandsEvent"`
		} `json:"events"`
	} `json:"eventData"`
	Settings struct {
	} `json:"settings"`
}
