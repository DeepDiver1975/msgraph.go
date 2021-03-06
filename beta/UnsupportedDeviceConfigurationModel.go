// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UnsupportedDeviceConfiguration UnsupportedDeviceConfiguration is used when an entity cannot be mapped to another model-compliant subtype of deviceConfiguration.
type UnsupportedDeviceConfiguration struct {
	// DeviceConfiguration is the base model of UnsupportedDeviceConfiguration
	DeviceConfiguration
	// OriginalEntityTypeName The type of entity that would be returned otherwise.
	OriginalEntityTypeName *string `json:"originalEntityTypeName,omitempty"`
	// Details Details describing why the entity is unsupported. This collection can contain a maximum of 1000 elements.
	Details []UnsupportedDeviceConfigurationDetail `json:"details,omitempty"`
}
