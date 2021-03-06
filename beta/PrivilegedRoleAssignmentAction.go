// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PrivilegedRoleAssignmentMakePermanentRequestParameter undocumented
type PrivilegedRoleAssignmentMakePermanentRequestParameter struct {
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// TicketNumber undocumented
	TicketNumber *string `json:"ticketNumber,omitempty"`
	// TicketSystem undocumented
	TicketSystem *string `json:"ticketSystem,omitempty"`
}

// PrivilegedRoleAssignmentMakeEligibleRequestParameter undocumented
type PrivilegedRoleAssignmentMakeEligibleRequestParameter struct {
}

//
type PrivilegedRoleAssignmentMakePermanentRequestBuilder struct{ BaseRequestBuilder }

// MakePermanent action undocumented
func (b *PrivilegedRoleAssignmentRequestBuilder) MakePermanent(reqObj *PrivilegedRoleAssignmentMakePermanentRequestParameter) *PrivilegedRoleAssignmentMakePermanentRequestBuilder {
	bb := &PrivilegedRoleAssignmentMakePermanentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/makePermanent"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrivilegedRoleAssignmentMakePermanentRequest struct{ BaseRequest }

//
func (b *PrivilegedRoleAssignmentMakePermanentRequestBuilder) Request() *PrivilegedRoleAssignmentMakePermanentRequest {
	return &PrivilegedRoleAssignmentMakePermanentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrivilegedRoleAssignmentMakePermanentRequest) Post(ctx context.Context) (resObj *PrivilegedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type PrivilegedRoleAssignmentMakeEligibleRequestBuilder struct{ BaseRequestBuilder }

// MakeEligible action undocumented
func (b *PrivilegedRoleAssignmentRequestBuilder) MakeEligible(reqObj *PrivilegedRoleAssignmentMakeEligibleRequestParameter) *PrivilegedRoleAssignmentMakeEligibleRequestBuilder {
	bb := &PrivilegedRoleAssignmentMakeEligibleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/makeEligible"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrivilegedRoleAssignmentMakeEligibleRequest struct{ BaseRequest }

//
func (b *PrivilegedRoleAssignmentMakeEligibleRequestBuilder) Request() *PrivilegedRoleAssignmentMakeEligibleRequest {
	return &PrivilegedRoleAssignmentMakeEligibleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrivilegedRoleAssignmentMakeEligibleRequest) Post(ctx context.Context) (resObj *PrivilegedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
