// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ResourceSpecificPermissionGrant undocumented
type ResourceSpecificPermissionGrant struct {
	// DirectoryObject is the base model of ResourceSpecificPermissionGrant
	DirectoryObject
	// ClientID undocumented
	ClientID *string `json:"clientId,omitempty"`
	// ClientAppID undocumented
	ClientAppID *string `json:"clientAppId,omitempty"`
	// ResourceAppID undocumented
	ResourceAppID *string `json:"resourceAppId,omitempty"`
	// PermissionType undocumented
	PermissionType *string `json:"permissionType,omitempty"`
	// Permission undocumented
	Permission *string `json:"permission,omitempty"`
}
