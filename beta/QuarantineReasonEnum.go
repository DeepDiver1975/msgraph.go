// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// QuarantineReason undocumented
type QuarantineReason int

const (
	// QuarantineReasonVEncounteredBaseEscrowThreshold undocumented
	QuarantineReasonVEncounteredBaseEscrowThreshold QuarantineReason = 0
	// QuarantineReasonVEncounteredTotalEscrowThreshold undocumented
	QuarantineReasonVEncounteredTotalEscrowThreshold QuarantineReason = 1
	// QuarantineReasonVEncounteredEscrowProportionThreshold undocumented
	QuarantineReasonVEncounteredEscrowProportionThreshold QuarantineReason = 2
	// QuarantineReasonVEncounteredQuarantineException undocumented
	QuarantineReasonVEncounteredQuarantineException QuarantineReason = 4
	// QuarantineReasonVUnknown undocumented
	QuarantineReasonVUnknown QuarantineReason = 8
	// QuarantineReasonVQuarantinedOnDemand undocumented
	QuarantineReasonVQuarantinedOnDemand QuarantineReason = 16
	// QuarantineReasonVTooManyDeletes undocumented
	QuarantineReasonVTooManyDeletes QuarantineReason = 32
)

// QuarantineReasonPEncounteredBaseEscrowThreshold returns a pointer to QuarantineReasonVEncounteredBaseEscrowThreshold
func QuarantineReasonPEncounteredBaseEscrowThreshold() *QuarantineReason {
	v := QuarantineReasonVEncounteredBaseEscrowThreshold
	return &v
}

// QuarantineReasonPEncounteredTotalEscrowThreshold returns a pointer to QuarantineReasonVEncounteredTotalEscrowThreshold
func QuarantineReasonPEncounteredTotalEscrowThreshold() *QuarantineReason {
	v := QuarantineReasonVEncounteredTotalEscrowThreshold
	return &v
}

// QuarantineReasonPEncounteredEscrowProportionThreshold returns a pointer to QuarantineReasonVEncounteredEscrowProportionThreshold
func QuarantineReasonPEncounteredEscrowProportionThreshold() *QuarantineReason {
	v := QuarantineReasonVEncounteredEscrowProportionThreshold
	return &v
}

// QuarantineReasonPEncounteredQuarantineException returns a pointer to QuarantineReasonVEncounteredQuarantineException
func QuarantineReasonPEncounteredQuarantineException() *QuarantineReason {
	v := QuarantineReasonVEncounteredQuarantineException
	return &v
}

// QuarantineReasonPUnknown returns a pointer to QuarantineReasonVUnknown
func QuarantineReasonPUnknown() *QuarantineReason {
	v := QuarantineReasonVUnknown
	return &v
}

// QuarantineReasonPQuarantinedOnDemand returns a pointer to QuarantineReasonVQuarantinedOnDemand
func QuarantineReasonPQuarantinedOnDemand() *QuarantineReason {
	v := QuarantineReasonVQuarantinedOnDemand
	return &v
}

// QuarantineReasonPTooManyDeletes returns a pointer to QuarantineReasonVTooManyDeletes
func QuarantineReasonPTooManyDeletes() *QuarantineReason {
	v := QuarantineReasonVTooManyDeletes
	return &v
}
