// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SettingStateDeviceSummary Device Compilance Policy and Configuration for a Setting State summary
type SettingStateDeviceSummary struct {
	// Entity is the base model of SettingStateDeviceSummary
	Entity
	// SettingName Name of the setting
	SettingName *string `json:"settingName,omitempty"`
	// InstancePath Name of the InstancePath for the setting
	InstancePath *string `json:"instancePath,omitempty"`
	// UnknownDeviceCount Device Unkown count for the setting
	UnknownDeviceCount *int `json:"unknownDeviceCount,omitempty"`
	// NotApplicableDeviceCount Device Not Applicable count for the setting
	NotApplicableDeviceCount *int `json:"notApplicableDeviceCount,omitempty"`
	// CompliantDeviceCount Device Compliant count for the setting
	CompliantDeviceCount *int `json:"compliantDeviceCount,omitempty"`
	// RemediatedDeviceCount Device Compliant count for the setting
	RemediatedDeviceCount *int `json:"remediatedDeviceCount,omitempty"`
	// NonCompliantDeviceCount Device NonCompliant count for the setting
	NonCompliantDeviceCount *int `json:"nonCompliantDeviceCount,omitempty"`
	// ErrorDeviceCount Device error count for the setting
	ErrorDeviceCount *int `json:"errorDeviceCount,omitempty"`
	// ConflictDeviceCount Device conflict error count for the setting
	ConflictDeviceCount *int `json:"conflictDeviceCount,omitempty"`
}
