// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceOwnerDefaultAppPermissionPolicyType undocumented
type AndroidDeviceOwnerDefaultAppPermissionPolicyType int

const (
	// AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVDeviceDefault undocumented
	AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVDeviceDefault AndroidDeviceOwnerDefaultAppPermissionPolicyType = 0
	// AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVPrompt undocumented
	AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVPrompt AndroidDeviceOwnerDefaultAppPermissionPolicyType = 1
	// AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoGrant undocumented
	AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoGrant AndroidDeviceOwnerDefaultAppPermissionPolicyType = 2
	// AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoDeny undocumented
	AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoDeny AndroidDeviceOwnerDefaultAppPermissionPolicyType = 3
)

// AndroidDeviceOwnerDefaultAppPermissionPolicyTypePDeviceDefault returns a pointer to AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVDeviceDefault
func AndroidDeviceOwnerDefaultAppPermissionPolicyTypePDeviceDefault() *AndroidDeviceOwnerDefaultAppPermissionPolicyType {
	v := AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVDeviceDefault
	return &v
}

// AndroidDeviceOwnerDefaultAppPermissionPolicyTypePPrompt returns a pointer to AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVPrompt
func AndroidDeviceOwnerDefaultAppPermissionPolicyTypePPrompt() *AndroidDeviceOwnerDefaultAppPermissionPolicyType {
	v := AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVPrompt
	return &v
}

// AndroidDeviceOwnerDefaultAppPermissionPolicyTypePAutoGrant returns a pointer to AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoGrant
func AndroidDeviceOwnerDefaultAppPermissionPolicyTypePAutoGrant() *AndroidDeviceOwnerDefaultAppPermissionPolicyType {
	v := AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoGrant
	return &v
}

// AndroidDeviceOwnerDefaultAppPermissionPolicyTypePAutoDeny returns a pointer to AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoDeny
func AndroidDeviceOwnerDefaultAppPermissionPolicyTypePAutoDeny() *AndroidDeviceOwnerDefaultAppPermissionPolicyType {
	v := AndroidDeviceOwnerDefaultAppPermissionPolicyTypeVAutoDeny
	return &v
}