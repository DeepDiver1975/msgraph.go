// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SharedPCAllowedAccountType undocumented
type SharedPCAllowedAccountType int

const (
	// SharedPCAllowedAccountTypeVGuest undocumented
	SharedPCAllowedAccountTypeVGuest SharedPCAllowedAccountType = 1
	// SharedPCAllowedAccountTypeVDomain undocumented
	SharedPCAllowedAccountTypeVDomain SharedPCAllowedAccountType = 2
)

// SharedPCAllowedAccountTypePGuest returns a pointer to SharedPCAllowedAccountTypeVGuest
func SharedPCAllowedAccountTypePGuest() *SharedPCAllowedAccountType {
	v := SharedPCAllowedAccountTypeVGuest
	return &v
}

// SharedPCAllowedAccountTypePDomain returns a pointer to SharedPCAllowedAccountTypeVDomain
func SharedPCAllowedAccountTypePDomain() *SharedPCAllowedAccountType {
	v := SharedPCAllowedAccountTypeVDomain
	return &v
}
