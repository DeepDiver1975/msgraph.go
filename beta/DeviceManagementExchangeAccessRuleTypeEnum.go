// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementExchangeAccessRuleType undocumented
type DeviceManagementExchangeAccessRuleType int

const (
	// DeviceManagementExchangeAccessRuleTypeVFamily undocumented
	DeviceManagementExchangeAccessRuleTypeVFamily DeviceManagementExchangeAccessRuleType = 0
	// DeviceManagementExchangeAccessRuleTypeVModel undocumented
	DeviceManagementExchangeAccessRuleTypeVModel DeviceManagementExchangeAccessRuleType = 1
)

// DeviceManagementExchangeAccessRuleTypePFamily returns a pointer to DeviceManagementExchangeAccessRuleTypeVFamily
func DeviceManagementExchangeAccessRuleTypePFamily() *DeviceManagementExchangeAccessRuleType {
	v := DeviceManagementExchangeAccessRuleTypeVFamily
	return &v
}

// DeviceManagementExchangeAccessRuleTypePModel returns a pointer to DeviceManagementExchangeAccessRuleTypeVModel
func DeviceManagementExchangeAccessRuleTypePModel() *DeviceManagementExchangeAccessRuleType {
	v := DeviceManagementExchangeAccessRuleTypeVModel
	return &v
}
