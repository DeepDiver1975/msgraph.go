// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsDeviceHealthState undocumented
type WindowsDeviceHealthState int

const (
	// WindowsDeviceHealthStateVClean undocumented
	WindowsDeviceHealthStateVClean WindowsDeviceHealthState = 0
	// WindowsDeviceHealthStateVFullScanPending undocumented
	WindowsDeviceHealthStateVFullScanPending WindowsDeviceHealthState = 1
	// WindowsDeviceHealthStateVRebootPending undocumented
	WindowsDeviceHealthStateVRebootPending WindowsDeviceHealthState = 2
	// WindowsDeviceHealthStateVManualStepsPending undocumented
	WindowsDeviceHealthStateVManualStepsPending WindowsDeviceHealthState = 4
	// WindowsDeviceHealthStateVOfflineScanPending undocumented
	WindowsDeviceHealthStateVOfflineScanPending WindowsDeviceHealthState = 8
	// WindowsDeviceHealthStateVCritical undocumented
	WindowsDeviceHealthStateVCritical WindowsDeviceHealthState = 16
)

// WindowsDeviceHealthStatePClean returns a pointer to WindowsDeviceHealthStateVClean
func WindowsDeviceHealthStatePClean() *WindowsDeviceHealthState {
	v := WindowsDeviceHealthStateVClean
	return &v
}

// WindowsDeviceHealthStatePFullScanPending returns a pointer to WindowsDeviceHealthStateVFullScanPending
func WindowsDeviceHealthStatePFullScanPending() *WindowsDeviceHealthState {
	v := WindowsDeviceHealthStateVFullScanPending
	return &v
}

// WindowsDeviceHealthStatePRebootPending returns a pointer to WindowsDeviceHealthStateVRebootPending
func WindowsDeviceHealthStatePRebootPending() *WindowsDeviceHealthState {
	v := WindowsDeviceHealthStateVRebootPending
	return &v
}

// WindowsDeviceHealthStatePManualStepsPending returns a pointer to WindowsDeviceHealthStateVManualStepsPending
func WindowsDeviceHealthStatePManualStepsPending() *WindowsDeviceHealthState {
	v := WindowsDeviceHealthStateVManualStepsPending
	return &v
}

// WindowsDeviceHealthStatePOfflineScanPending returns a pointer to WindowsDeviceHealthStateVOfflineScanPending
func WindowsDeviceHealthStatePOfflineScanPending() *WindowsDeviceHealthState {
	v := WindowsDeviceHealthStateVOfflineScanPending
	return &v
}

// WindowsDeviceHealthStatePCritical returns a pointer to WindowsDeviceHealthStateVCritical
func WindowsDeviceHealthStatePCritical() *WindowsDeviceHealthState {
	v := WindowsDeviceHealthStateVCritical
	return &v
}
