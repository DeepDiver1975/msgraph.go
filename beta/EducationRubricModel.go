// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationRubric undocumented
type EducationRubric struct {
	// Entity is the base model of EducationRubric
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// Qualities undocumented
	Qualities []RubricQuality `json:"qualities,omitempty"`
	// Levels undocumented
	Levels []RubricLevel `json:"levels,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
}
