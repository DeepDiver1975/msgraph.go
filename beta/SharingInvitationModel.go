// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SharingInvitation undocumented
type SharingInvitation struct {
	// Object is the base model of SharingInvitation
	Object
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// InvitedBy undocumented
	InvitedBy *IdentitySet `json:"invitedBy,omitempty"`
	// RedeemedBy undocumented
	RedeemedBy *string `json:"redeemedBy,omitempty"`
	// SignInRequired undocumented
	SignInRequired *bool `json:"signInRequired,omitempty"`
}
