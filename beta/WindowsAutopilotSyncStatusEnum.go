// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsAutopilotSyncStatus undocumented
type WindowsAutopilotSyncStatus int

const (
	// WindowsAutopilotSyncStatusVUnknown undocumented
	WindowsAutopilotSyncStatusVUnknown WindowsAutopilotSyncStatus = 0
	// WindowsAutopilotSyncStatusVInProgress undocumented
	WindowsAutopilotSyncStatusVInProgress WindowsAutopilotSyncStatus = 1
	// WindowsAutopilotSyncStatusVCompleted undocumented
	WindowsAutopilotSyncStatusVCompleted WindowsAutopilotSyncStatus = 2
	// WindowsAutopilotSyncStatusVFailed undocumented
	WindowsAutopilotSyncStatusVFailed WindowsAutopilotSyncStatus = 3
)

// WindowsAutopilotSyncStatusPUnknown returns a pointer to WindowsAutopilotSyncStatusVUnknown
func WindowsAutopilotSyncStatusPUnknown() *WindowsAutopilotSyncStatus {
	v := WindowsAutopilotSyncStatusVUnknown
	return &v
}

// WindowsAutopilotSyncStatusPInProgress returns a pointer to WindowsAutopilotSyncStatusVInProgress
func WindowsAutopilotSyncStatusPInProgress() *WindowsAutopilotSyncStatus {
	v := WindowsAutopilotSyncStatusVInProgress
	return &v
}

// WindowsAutopilotSyncStatusPCompleted returns a pointer to WindowsAutopilotSyncStatusVCompleted
func WindowsAutopilotSyncStatusPCompleted() *WindowsAutopilotSyncStatus {
	v := WindowsAutopilotSyncStatusVCompleted
	return &v
}

// WindowsAutopilotSyncStatusPFailed returns a pointer to WindowsAutopilotSyncStatusVFailed
func WindowsAutopilotSyncStatusPFailed() *WindowsAutopilotSyncStatus {
	v := WindowsAutopilotSyncStatusVFailed
	return &v
}
