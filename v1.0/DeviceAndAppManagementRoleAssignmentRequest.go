// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceAndAppManagementRoleAssignmentRequestBuilder is request builder for DeviceAndAppManagementRoleAssignment
type DeviceAndAppManagementRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceAndAppManagementRoleAssignmentRequest
func (b *DeviceAndAppManagementRoleAssignmentRequestBuilder) Request() *DeviceAndAppManagementRoleAssignmentRequest {
	return &DeviceAndAppManagementRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceAndAppManagementRoleAssignmentRequest is request for DeviceAndAppManagementRoleAssignment
type DeviceAndAppManagementRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for DeviceAndAppManagementRoleAssignment
func (r *DeviceAndAppManagementRoleAssignmentRequest) Get(ctx context.Context) (resObj *DeviceAndAppManagementRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceAndAppManagementRoleAssignment
func (r *DeviceAndAppManagementRoleAssignmentRequest) Update(ctx context.Context, reqObj *DeviceAndAppManagementRoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceAndAppManagementRoleAssignment
func (r *DeviceAndAppManagementRoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
