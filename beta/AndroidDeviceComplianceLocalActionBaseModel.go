// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidDeviceComplianceLocalActionBase Local Action Configuration
type AndroidDeviceComplianceLocalActionBase struct {
	// Entity is the base model of AndroidDeviceComplianceLocalActionBase
	Entity
	// GracePeriodInMinutes Number of minutes to wait till a local action is enforced. Valid values 0 to 2147483647
	GracePeriodInMinutes *int `json:"gracePeriodInMinutes,omitempty"`
}
