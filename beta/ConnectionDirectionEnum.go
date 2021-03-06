// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConnectionDirection undocumented
type ConnectionDirection int

const (
	// ConnectionDirectionVUnknown undocumented
	ConnectionDirectionVUnknown ConnectionDirection = 0
	// ConnectionDirectionVInbound undocumented
	ConnectionDirectionVInbound ConnectionDirection = 1
	// ConnectionDirectionVOutbound undocumented
	ConnectionDirectionVOutbound ConnectionDirection = 2
	// ConnectionDirectionVUnknownFutureValue undocumented
	ConnectionDirectionVUnknownFutureValue ConnectionDirection = 127
)

// ConnectionDirectionPUnknown returns a pointer to ConnectionDirectionVUnknown
func ConnectionDirectionPUnknown() *ConnectionDirection {
	v := ConnectionDirectionVUnknown
	return &v
}

// ConnectionDirectionPInbound returns a pointer to ConnectionDirectionVInbound
func ConnectionDirectionPInbound() *ConnectionDirection {
	v := ConnectionDirectionVInbound
	return &v
}

// ConnectionDirectionPOutbound returns a pointer to ConnectionDirectionVOutbound
func ConnectionDirectionPOutbound() *ConnectionDirection {
	v := ConnectionDirectionVOutbound
	return &v
}

// ConnectionDirectionPUnknownFutureValue returns a pointer to ConnectionDirectionVUnknownFutureValue
func ConnectionDirectionPUnknownFutureValue() *ConnectionDirection {
	v := ConnectionDirectionVUnknownFutureValue
	return &v
}
