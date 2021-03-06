// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessPackageResourceScopeRequestBuilder is request builder for AccessPackageResourceScope
type AccessPackageResourceScopeRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageResourceScopeRequest
func (b *AccessPackageResourceScopeRequestBuilder) Request() *AccessPackageResourceScopeRequest {
	return &AccessPackageResourceScopeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageResourceScopeRequest is request for AccessPackageResourceScope
type AccessPackageResourceScopeRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageResourceScope
func (r *AccessPackageResourceScopeRequest) Get(ctx context.Context) (resObj *AccessPackageResourceScope, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageResourceScope
func (r *AccessPackageResourceScopeRequest) Update(ctx context.Context, reqObj *AccessPackageResourceScope) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageResourceScope
func (r *AccessPackageResourceScopeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackageResource is navigation property
func (b *AccessPackageResourceScopeRequestBuilder) AccessPackageResource() *AccessPackageResourceRequestBuilder {
	bb := &AccessPackageResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResource"
	return bb
}
