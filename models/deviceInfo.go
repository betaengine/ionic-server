package models

type DeviceInfo struct {
	Device_app_version string `json:"device_app_version"`
	Device_deploy_uuid string `json:"device_deploy_uuid"`
	Device_platform    string `json:"device_platform"`
	Channel_tag        string `json:"channel_tag"`
}