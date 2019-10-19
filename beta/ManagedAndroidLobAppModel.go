// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAndroidLobApp Contains properties and inherited properties for Managed Android Line Of Business apps.
type ManagedAndroidLobApp struct {
	ManagedMobileLobApp
	// PackageID The package identifier.
	PackageID *string `json:"packageId,omitempty"`
	// IdentityName The Identity Name.
	IdentityName *string `json:"identityName,omitempty"`
	// MinimumSupportedOperatingSystem The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *AndroidMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// VersionName The version name of managed Android Line of Business (LoB) app.
	VersionName *string `json:"versionName,omitempty"`
	// VersionCode The version code of managed Android Line of Business (LoB) app.
	VersionCode *string `json:"versionCode,omitempty"`
	// IdentityVersion The identity version.
	IdentityVersion *string `json:"identityVersion,omitempty"`
}