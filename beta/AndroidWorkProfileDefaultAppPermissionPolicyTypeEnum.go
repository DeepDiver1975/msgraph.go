// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidWorkProfileDefaultAppPermissionPolicyType undocumented
type AndroidWorkProfileDefaultAppPermissionPolicyType int

const (
	// AndroidWorkProfileDefaultAppPermissionPolicyTypeVDeviceDefault undocumented
	AndroidWorkProfileDefaultAppPermissionPolicyTypeVDeviceDefault AndroidWorkProfileDefaultAppPermissionPolicyType = 0
	// AndroidWorkProfileDefaultAppPermissionPolicyTypeVPrompt undocumented
	AndroidWorkProfileDefaultAppPermissionPolicyTypeVPrompt AndroidWorkProfileDefaultAppPermissionPolicyType = 1
	// AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoGrant undocumented
	AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoGrant AndroidWorkProfileDefaultAppPermissionPolicyType = 2
	// AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoDeny undocumented
	AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoDeny AndroidWorkProfileDefaultAppPermissionPolicyType = 3
)

// AndroidWorkProfileDefaultAppPermissionPolicyTypePDeviceDefault returns a pointer to AndroidWorkProfileDefaultAppPermissionPolicyTypeVDeviceDefault
func AndroidWorkProfileDefaultAppPermissionPolicyTypePDeviceDefault() *AndroidWorkProfileDefaultAppPermissionPolicyType {
	v := AndroidWorkProfileDefaultAppPermissionPolicyTypeVDeviceDefault
	return &v
}

// AndroidWorkProfileDefaultAppPermissionPolicyTypePPrompt returns a pointer to AndroidWorkProfileDefaultAppPermissionPolicyTypeVPrompt
func AndroidWorkProfileDefaultAppPermissionPolicyTypePPrompt() *AndroidWorkProfileDefaultAppPermissionPolicyType {
	v := AndroidWorkProfileDefaultAppPermissionPolicyTypeVPrompt
	return &v
}

// AndroidWorkProfileDefaultAppPermissionPolicyTypePAutoGrant returns a pointer to AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoGrant
func AndroidWorkProfileDefaultAppPermissionPolicyTypePAutoGrant() *AndroidWorkProfileDefaultAppPermissionPolicyType {
	v := AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoGrant
	return &v
}

// AndroidWorkProfileDefaultAppPermissionPolicyTypePAutoDeny returns a pointer to AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoDeny
func AndroidWorkProfileDefaultAppPermissionPolicyTypePAutoDeny() *AndroidWorkProfileDefaultAppPermissionPolicyType {
	v := AndroidWorkProfileDefaultAppPermissionPolicyTypeVAutoDeny
	return &v
}
