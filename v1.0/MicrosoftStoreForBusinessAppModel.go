// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MicrosoftStoreForBusinessApp Microsoft Store for Business Apps. This class does not support Create, Delete, or Update.
type MicrosoftStoreForBusinessApp struct {
	MobileApp
	// UsedLicenseCount The number of Microsoft Store for Business licenses in use.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// TotalLicenseCount The total number of Microsoft Store for Business licenses.
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// ProductKey The app product key
	ProductKey *string `json:"productKey,omitempty"`
	// LicenseType The app license type
	LicenseType *MicrosoftStoreForBusinessLicenseType `json:"licenseType,omitempty"`
	// PackageIdentityName The app package identifier
	PackageIdentityName *string `json:"packageIdentityName,omitempty"`
}