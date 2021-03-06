// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ListItemVersionRestoreVersionRequestParameter undocumented
type ListItemVersionRestoreVersionRequestParameter struct {
}

//
type ListItemVersionRestoreVersionRequestBuilder struct{ BaseRequestBuilder }

// RestoreVersion action undocumented
func (b *ListItemVersionRequestBuilder) RestoreVersion(reqObj *ListItemVersionRestoreVersionRequestParameter) *ListItemVersionRestoreVersionRequestBuilder {
	bb := &ListItemVersionRestoreVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/restoreVersion"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ListItemVersionRestoreVersionRequest struct{ BaseRequest }

//
func (b *ListItemVersionRestoreVersionRequestBuilder) Request() *ListItemVersionRestoreVersionRequest {
	return &ListItemVersionRestoreVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ListItemVersionRestoreVersionRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
