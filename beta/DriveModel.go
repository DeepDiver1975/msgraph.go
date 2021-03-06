// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Drive undocumented
type Drive struct {
	// BaseItem is the base model of Drive
	BaseItem
	// DriveType undocumented
	DriveType *string `json:"driveType,omitempty"`
	// Owner undocumented
	Owner *IdentitySet `json:"owner,omitempty"`
	// Quota undocumented
	Quota *Quota `json:"quota,omitempty"`
	// SharePointIDs undocumented
	SharePointIDs *SharepointIDs `json:"sharePointIds,omitempty"`
	// System undocumented
	System *SystemFacet `json:"system,omitempty"`
	// Activities undocumented
	Activities []ItemActivityOLD `json:"activities,omitempty"`
	// Bundles undocumented
	Bundles []DriveItem `json:"bundles,omitempty"`
	// Following undocumented
	Following []DriveItem `json:"following,omitempty"`
	// Items undocumented
	Items []DriveItem `json:"items,omitempty"`
	// List undocumented
	List *List `json:"list,omitempty"`
	// Root undocumented
	Root *DriveItem `json:"root,omitempty"`
	// Special undocumented
	Special []DriveItem `json:"special,omitempty"`
}
