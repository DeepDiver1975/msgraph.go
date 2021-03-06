// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MediaStream undocumented
type MediaStream struct {
	// Object is the base model of MediaStream
	Object
	// MediaType undocumented
	MediaType *Modality `json:"mediaType,omitempty"`
	// Label undocumented
	Label *string `json:"label,omitempty"`
	// SourceID undocumented
	SourceID *string `json:"sourceId,omitempty"`
	// Direction undocumented
	Direction *MediaDirection `json:"direction,omitempty"`
	// ServerMuted undocumented
	ServerMuted *bool `json:"serverMuted,omitempty"`
}
