// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EventMessage undocumented
type EventMessage struct {
	Message
	// MeetingMessageType undocumented
	MeetingMessageType *MeetingMessageType `json:"meetingMessageType,omitempty"`
	// Event undocumented
	Event *Event `json:"event,omitempty"`
}