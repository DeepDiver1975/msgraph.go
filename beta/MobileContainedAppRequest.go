// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MobileContainedAppRequestBuilder is request builder for MobileContainedApp
type MobileContainedAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileContainedAppRequest
func (b *MobileContainedAppRequestBuilder) Request() *MobileContainedAppRequest {
	return &MobileContainedAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileContainedAppRequest is request for MobileContainedApp
type MobileContainedAppRequest struct{ BaseRequest }

// Get performs GET request for MobileContainedApp
func (r *MobileContainedAppRequest) Get(ctx context.Context) (resObj *MobileContainedApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileContainedApp
func (r *MobileContainedAppRequest) Update(ctx context.Context, reqObj *MobileContainedApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileContainedApp
func (r *MobileContainedAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
