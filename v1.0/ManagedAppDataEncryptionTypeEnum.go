// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppDataEncryptionType undocumented
type ManagedAppDataEncryptionType int

const (
	// ManagedAppDataEncryptionTypeVUseDeviceSettings undocumented
	ManagedAppDataEncryptionTypeVUseDeviceSettings ManagedAppDataEncryptionType = 0
	// ManagedAppDataEncryptionTypeVAfterDeviceRestart undocumented
	ManagedAppDataEncryptionTypeVAfterDeviceRestart ManagedAppDataEncryptionType = 1
	// ManagedAppDataEncryptionTypeVWhenDeviceLockedExceptOpenFiles undocumented
	ManagedAppDataEncryptionTypeVWhenDeviceLockedExceptOpenFiles ManagedAppDataEncryptionType = 2
	// ManagedAppDataEncryptionTypeVWhenDeviceLocked undocumented
	ManagedAppDataEncryptionTypeVWhenDeviceLocked ManagedAppDataEncryptionType = 3
)

// ManagedAppDataEncryptionTypePUseDeviceSettings returns a pointer to ManagedAppDataEncryptionTypeVUseDeviceSettings
func ManagedAppDataEncryptionTypePUseDeviceSettings() *ManagedAppDataEncryptionType {
	v := ManagedAppDataEncryptionTypeVUseDeviceSettings
	return &v
}

// ManagedAppDataEncryptionTypePAfterDeviceRestart returns a pointer to ManagedAppDataEncryptionTypeVAfterDeviceRestart
func ManagedAppDataEncryptionTypePAfterDeviceRestart() *ManagedAppDataEncryptionType {
	v := ManagedAppDataEncryptionTypeVAfterDeviceRestart
	return &v
}

// ManagedAppDataEncryptionTypePWhenDeviceLockedExceptOpenFiles returns a pointer to ManagedAppDataEncryptionTypeVWhenDeviceLockedExceptOpenFiles
func ManagedAppDataEncryptionTypePWhenDeviceLockedExceptOpenFiles() *ManagedAppDataEncryptionType {
	v := ManagedAppDataEncryptionTypeVWhenDeviceLockedExceptOpenFiles
	return &v
}

// ManagedAppDataEncryptionTypePWhenDeviceLocked returns a pointer to ManagedAppDataEncryptionTypeVWhenDeviceLocked
func ManagedAppDataEncryptionTypePWhenDeviceLocked() *ManagedAppDataEncryptionType {
	v := ManagedAppDataEncryptionTypeVWhenDeviceLocked
	return &v
}
