// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementPartnerAppType undocumented
type DeviceManagementPartnerAppType int

const (
	// DeviceManagementPartnerAppTypeVUnknown undocumented
	DeviceManagementPartnerAppTypeVUnknown DeviceManagementPartnerAppType = 0
	// DeviceManagementPartnerAppTypeVSingleTenantApp undocumented
	DeviceManagementPartnerAppTypeVSingleTenantApp DeviceManagementPartnerAppType = 1
	// DeviceManagementPartnerAppTypeVMultiTenantApp undocumented
	DeviceManagementPartnerAppTypeVMultiTenantApp DeviceManagementPartnerAppType = 2
)

// DeviceManagementPartnerAppTypePUnknown returns a pointer to DeviceManagementPartnerAppTypeVUnknown
func DeviceManagementPartnerAppTypePUnknown() *DeviceManagementPartnerAppType {
	v := DeviceManagementPartnerAppTypeVUnknown
	return &v
}

// DeviceManagementPartnerAppTypePSingleTenantApp returns a pointer to DeviceManagementPartnerAppTypeVSingleTenantApp
func DeviceManagementPartnerAppTypePSingleTenantApp() *DeviceManagementPartnerAppType {
	v := DeviceManagementPartnerAppTypeVSingleTenantApp
	return &v
}

// DeviceManagementPartnerAppTypePMultiTenantApp returns a pointer to DeviceManagementPartnerAppTypeVMultiTenantApp
func DeviceManagementPartnerAppTypePMultiTenantApp() *DeviceManagementPartnerAppType {
	v := DeviceManagementPartnerAppTypeVMultiTenantApp
	return &v
}
