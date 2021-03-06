// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhone81TrustedRootCertificate Windows Phone 8.1+ Trusted Root Certificate configuration profile
type WindowsPhone81TrustedRootCertificate struct {
	// DeviceConfiguration is the base model of WindowsPhone81TrustedRootCertificate
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
}
