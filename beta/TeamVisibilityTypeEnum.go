// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamVisibilityType undocumented
type TeamVisibilityType int

const (
	// TeamVisibilityTypeVPrivate undocumented
	TeamVisibilityTypeVPrivate TeamVisibilityType = 0
	// TeamVisibilityTypeVPublic undocumented
	TeamVisibilityTypeVPublic TeamVisibilityType = 1
	// TeamVisibilityTypeVHiddenMembership undocumented
	TeamVisibilityTypeVHiddenMembership TeamVisibilityType = 2
	// TeamVisibilityTypeVUnknownFutureValue undocumented
	TeamVisibilityTypeVUnknownFutureValue TeamVisibilityType = 3
)

// TeamVisibilityTypePPrivate returns a pointer to TeamVisibilityTypeVPrivate
func TeamVisibilityTypePPrivate() *TeamVisibilityType {
	v := TeamVisibilityTypeVPrivate
	return &v
}

// TeamVisibilityTypePPublic returns a pointer to TeamVisibilityTypeVPublic
func TeamVisibilityTypePPublic() *TeamVisibilityType {
	v := TeamVisibilityTypeVPublic
	return &v
}

// TeamVisibilityTypePHiddenMembership returns a pointer to TeamVisibilityTypeVHiddenMembership
func TeamVisibilityTypePHiddenMembership() *TeamVisibilityType {
	v := TeamVisibilityTypeVHiddenMembership
	return &v
}

// TeamVisibilityTypePUnknownFutureValue returns a pointer to TeamVisibilityTypeVUnknownFutureValue
func TeamVisibilityTypePUnknownFutureValue() *TeamVisibilityType {
	v := TeamVisibilityTypeVUnknownFutureValue
	return &v
}
