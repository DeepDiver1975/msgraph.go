// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserAccountInformation undocumented
type UserAccountInformation struct {
	// ItemFacet is the base model of UserAccountInformation
	ItemFacet
	// AgeGroup undocumented
	AgeGroup *string `json:"ageGroup,omitempty"`
	// CountryCode undocumented
	CountryCode *string `json:"countryCode,omitempty"`
	// PreferredLanguageTag undocumented
	PreferredLanguageTag *LocaleInfo `json:"preferredLanguageTag,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
