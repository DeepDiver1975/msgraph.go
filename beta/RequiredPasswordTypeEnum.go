// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RequiredPasswordType undocumented
type RequiredPasswordType int

const (
	// RequiredPasswordTypeVDeviceDefault undocumented
	RequiredPasswordTypeVDeviceDefault RequiredPasswordType = 0
	// RequiredPasswordTypeVAlphanumeric undocumented
	RequiredPasswordTypeVAlphanumeric RequiredPasswordType = 1
	// RequiredPasswordTypeVNumeric undocumented
	RequiredPasswordTypeVNumeric RequiredPasswordType = 2
)

// RequiredPasswordTypePDeviceDefault returns a pointer to RequiredPasswordTypeVDeviceDefault
func RequiredPasswordTypePDeviceDefault() *RequiredPasswordType {
	v := RequiredPasswordTypeVDeviceDefault
	return &v
}

// RequiredPasswordTypePAlphanumeric returns a pointer to RequiredPasswordTypeVAlphanumeric
func RequiredPasswordTypePAlphanumeric() *RequiredPasswordType {
	v := RequiredPasswordTypeVAlphanumeric
	return &v
}

// RequiredPasswordTypePNumeric returns a pointer to RequiredPasswordTypeVNumeric
func RequiredPasswordTypePNumeric() *RequiredPasswordType {
	v := RequiredPasswordTypeVNumeric
	return &v
}
