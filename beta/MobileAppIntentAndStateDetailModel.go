// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppIntentAndStateDetail undocumented
type MobileAppIntentAndStateDetail struct {
	// Object is the base model of MobileAppIntentAndStateDetail
	Object
	// ApplicationID MobieApp identifier.
	ApplicationID *string `json:"applicationId,omitempty"`
	// DisplayName The admin provided or imported title of the app.
	DisplayName *string `json:"displayName,omitempty"`
	// MobileAppIntent Mobile App Intent.
	MobileAppIntent *MobileAppIntent `json:"mobileAppIntent,omitempty"`
	// DisplayVersion Human readable version of the application
	DisplayVersion *string `json:"displayVersion,omitempty"`
	// InstallState The install state of the app.
	InstallState *ResultantAppState `json:"installState,omitempty"`
	// SupportedDeviceTypes The supported platforms for the app.
	SupportedDeviceTypes []MobileAppSupportedDeviceType `json:"supportedDeviceTypes,omitempty"`
}
