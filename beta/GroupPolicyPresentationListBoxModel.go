// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyPresentationListBox Represents an ADMX listBox element and an ADMX list element.
type GroupPolicyPresentationListBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationListBox
	GroupPolicyPresentation
	// ExplicitValue If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
	ExplicitValue *bool `json:"explicitValue,omitempty"`
	// ValuePrefix undocumented
	ValuePrefix *string `json:"valuePrefix,omitempty"`
}
