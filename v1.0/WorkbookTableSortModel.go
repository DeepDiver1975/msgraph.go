// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookTableSort undocumented
type WorkbookTableSort struct {
	Entity
	// Fields undocumented
	Fields []WorkbookSortField `json:"fields,omitempty"`
	// MatchCase undocumented
	MatchCase *bool `json:"matchCase,omitempty"`
	// Method undocumented
	Method *string `json:"method,omitempty"`
}