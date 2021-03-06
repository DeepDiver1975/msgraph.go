// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MeetingMessageType undocumented
type MeetingMessageType int

const (
	// MeetingMessageTypeVNone undocumented
	MeetingMessageTypeVNone MeetingMessageType = 0
	// MeetingMessageTypeVMeetingRequest undocumented
	MeetingMessageTypeVMeetingRequest MeetingMessageType = 1
	// MeetingMessageTypeVMeetingCancelled undocumented
	MeetingMessageTypeVMeetingCancelled MeetingMessageType = 2
	// MeetingMessageTypeVMeetingAccepted undocumented
	MeetingMessageTypeVMeetingAccepted MeetingMessageType = 3
	// MeetingMessageTypeVMeetingTenativelyAccepted undocumented
	MeetingMessageTypeVMeetingTenativelyAccepted MeetingMessageType = 4
	// MeetingMessageTypeVMeetingDeclined undocumented
	MeetingMessageTypeVMeetingDeclined MeetingMessageType = 5
)

// MeetingMessageTypePNone returns a pointer to MeetingMessageTypeVNone
func MeetingMessageTypePNone() *MeetingMessageType {
	v := MeetingMessageTypeVNone
	return &v
}

// MeetingMessageTypePMeetingRequest returns a pointer to MeetingMessageTypeVMeetingRequest
func MeetingMessageTypePMeetingRequest() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingRequest
	return &v
}

// MeetingMessageTypePMeetingCancelled returns a pointer to MeetingMessageTypeVMeetingCancelled
func MeetingMessageTypePMeetingCancelled() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingCancelled
	return &v
}

// MeetingMessageTypePMeetingAccepted returns a pointer to MeetingMessageTypeVMeetingAccepted
func MeetingMessageTypePMeetingAccepted() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingAccepted
	return &v
}

// MeetingMessageTypePMeetingTenativelyAccepted returns a pointer to MeetingMessageTypeVMeetingTenativelyAccepted
func MeetingMessageTypePMeetingTenativelyAccepted() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingTenativelyAccepted
	return &v
}

// MeetingMessageTypePMeetingDeclined returns a pointer to MeetingMessageTypeVMeetingDeclined
func MeetingMessageTypePMeetingDeclined() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingDeclined
	return &v
}
