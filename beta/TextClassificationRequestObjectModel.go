// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TextClassificationRequestObject undocumented
type TextClassificationRequestObject struct {
	// Entity is the base model of TextClassificationRequestObject
	Entity
	// Text undocumented
	Text *string `json:"text,omitempty"`
	// SensitiveTypeIDs undocumented
	SensitiveTypeIDs []string `json:"sensitiveTypeIds,omitempty"`
}
