package sdk

import "github.com/vbaksa/smartthings-go-sdk/sdk/capabilities"

//NewSwitchOn turn on switch
func NewSwitchOn() DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.Switch,
		Command:    "on",
		Arguments:  []map[string]interface{}{},
	}
}

//NewSwitchOff turn off switch
func NewSwitchOff() DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.Switch,
		Command:    "off",
		Arguments:  []map[string]interface{}{},
	}
}

//NewSwitchLevel change switch level
func NewSwitchLevel(level int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.SwitchLevel,
		Command:    "setLevel",
		Arguments:  []map[string]interface{}{},
	}
}

//NewColorControlLevel change light level
func NewColorControlLevel(level int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorControl,
		Command:    "setColor",
		Arguments: []map[string]interface{}{
			{
				"level": int(level),
			},
		},
	}
}

//NewColorControlLevelRate change light level and change rate
func NewColorControlLevelRate(level int, rate int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorControl,
		Command:    "setColor",
		Arguments: []map[string]interface{}{
			{
				"level": int(level),
				"rate":  int(rate),
			},
		},
	}
}

//NewColorControl control light settings
func NewColorControl(hue int, saturation int, hex string, level int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorControl,
		Command:    "setColor",
		Arguments: []map[string]interface{}{
			{
				"hue":        int(hue),
				"saturation": int(saturation),
				"hex":        hex,
				"level":      int(level),
			},
		},
	}
}

//NewColorControlHex change light color
func NewColorControlHex(hex string) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorControl,
		Command:    "setColor",
		Arguments: []map[string]interface{}{
			{
				"hex": hex,
			},
		},
	}
}

//NewColorControlHue change light hue setting
func NewColorControlHue(hue int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorControl,
		Command:    "setHue",
		Arguments: []map[string]interface{}{
			{
				"hex": int(hue),
			},
		},
	}
}

//NewColorControlSaturation change light saturation
func NewColorControlSaturation(saturation int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorControl,
		Command:    "setSaturation",
		Arguments: []map[string]interface{}{
			{
				"saturation": int(saturation),
			},
		},
	}
}

//NewColorTemperature change light temperature
func NewColorTemperature(temperature int) DeviceCommand {
	return DeviceCommand{
		Component:  "main",
		Capability: capabilities.ColorTemperature,
		Command:    "setColorTemperature",
		Arguments: []map[string]interface{}{
			{
				"temperature": int(temperature),
			},
		},
	}
}
