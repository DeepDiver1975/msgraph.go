// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PrivilegedApprovalRequestBuilder is request builder for PrivilegedApproval
type PrivilegedApprovalRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedApprovalRequest
func (b *PrivilegedApprovalRequestBuilder) Request() *PrivilegedApprovalRequest {
	return &PrivilegedApprovalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedApprovalRequest is request for PrivilegedApproval
type PrivilegedApprovalRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedApproval
func (r *PrivilegedApprovalRequest) Get(ctx context.Context) (resObj *PrivilegedApproval, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedApproval
func (r *PrivilegedApprovalRequest) Update(ctx context.Context, reqObj *PrivilegedApproval) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedApproval
func (r *PrivilegedApprovalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RequestNavigation is navigation property
func (b *PrivilegedApprovalRequestBuilder) RequestNavigation() *PrivilegedRoleAssignmentRequestObjectRequestBuilder {
	bb := &PrivilegedRoleAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/request"
	return bb
}

// RoleInfo is navigation property
func (b *PrivilegedApprovalRequestBuilder) RoleInfo() *PrivilegedRoleRequestBuilder {
	bb := &PrivilegedRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleInfo"
	return bb
}
