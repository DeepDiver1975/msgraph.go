// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ThreatAssessmentRequestPivotProperty undocumented
type ThreatAssessmentRequestPivotProperty int

const (
	// ThreatAssessmentRequestPivotPropertyVThreatCategory undocumented
	ThreatAssessmentRequestPivotPropertyVThreatCategory ThreatAssessmentRequestPivotProperty = 1
	// ThreatAssessmentRequestPivotPropertyVMailDestinationRoutingReason undocumented
	ThreatAssessmentRequestPivotPropertyVMailDestinationRoutingReason ThreatAssessmentRequestPivotProperty = 2
)

// ThreatAssessmentRequestPivotPropertyPThreatCategory returns a pointer to ThreatAssessmentRequestPivotPropertyVThreatCategory
func ThreatAssessmentRequestPivotPropertyPThreatCategory() *ThreatAssessmentRequestPivotProperty {
	v := ThreatAssessmentRequestPivotPropertyVThreatCategory
	return &v
}

// ThreatAssessmentRequestPivotPropertyPMailDestinationRoutingReason returns a pointer to ThreatAssessmentRequestPivotPropertyVMailDestinationRoutingReason
func ThreatAssessmentRequestPivotPropertyPMailDestinationRoutingReason() *ThreatAssessmentRequestPivotProperty {
	v := ThreatAssessmentRequestPivotPropertyVMailDestinationRoutingReason
	return &v
}
