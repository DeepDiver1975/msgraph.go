// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EdgeSearchEngineType undocumented
type EdgeSearchEngineType int

const (
	// EdgeSearchEngineTypeVDefault undocumented
	EdgeSearchEngineTypeVDefault EdgeSearchEngineType = 0
	// EdgeSearchEngineTypeVBing undocumented
	EdgeSearchEngineTypeVBing EdgeSearchEngineType = 1
)

// EdgeSearchEngineTypePDefault returns a pointer to EdgeSearchEngineTypeVDefault
func EdgeSearchEngineTypePDefault() *EdgeSearchEngineType {
	v := EdgeSearchEngineTypeVDefault
	return &v
}

// EdgeSearchEngineTypePBing returns a pointer to EdgeSearchEngineTypeVBing
func EdgeSearchEngineTypePBing() *EdgeSearchEngineType {
	v := EdgeSearchEngineTypeVBing
	return &v
}
