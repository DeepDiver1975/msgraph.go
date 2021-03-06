// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSImportedPFXCertificateProfile iOS PFX Import certificate profile
type IOSImportedPFXCertificateProfile struct {
	// IOSCertificateProfile is the base model of IOSImportedPFXCertificateProfile
	IOSCertificateProfile
	// IntendedPurpose Intended Purpose of the Certificate Profile - which could be Unassigned, SmimeEncryption, SmimeSigning etc.
	IntendedPurpose *IntendedPurpose `json:"intendedPurpose,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}
