// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IPv4Range undocumented
type IPv4Range struct {
	// IPRange is the base model of IPv4Range
	IPRange
	// LowerAddress Lower address.
	LowerAddress *string `json:"lowerAddress,omitempty"`
	// UpperAddress Upper address.
	UpperAddress *string `json:"upperAddress,omitempty"`
}
