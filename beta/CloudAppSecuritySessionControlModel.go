// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CloudAppSecuritySessionControl undocumented
type CloudAppSecuritySessionControl struct {
	// ConditionalAccessSessionControl is the base model of CloudAppSecuritySessionControl
	ConditionalAccessSessionControl
	// CloudAppSecurityType undocumented
	CloudAppSecurityType *CloudAppSecuritySessionControlType `json:"cloudAppSecurityType,omitempty"`
}
