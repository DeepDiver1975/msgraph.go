// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsFeatureUpdateProfileAssignmentRequestBuilder is request builder for WindowsFeatureUpdateProfileAssignment
type WindowsFeatureUpdateProfileAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsFeatureUpdateProfileAssignmentRequest
func (b *WindowsFeatureUpdateProfileAssignmentRequestBuilder) Request() *WindowsFeatureUpdateProfileAssignmentRequest {
	return &WindowsFeatureUpdateProfileAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsFeatureUpdateProfileAssignmentRequest is request for WindowsFeatureUpdateProfileAssignment
type WindowsFeatureUpdateProfileAssignmentRequest struct{ BaseRequest }

// Get performs GET request for WindowsFeatureUpdateProfileAssignment
func (r *WindowsFeatureUpdateProfileAssignmentRequest) Get(ctx context.Context) (resObj *WindowsFeatureUpdateProfileAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsFeatureUpdateProfileAssignment
func (r *WindowsFeatureUpdateProfileAssignmentRequest) Update(ctx context.Context, reqObj *WindowsFeatureUpdateProfileAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsFeatureUpdateProfileAssignment
func (r *WindowsFeatureUpdateProfileAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
