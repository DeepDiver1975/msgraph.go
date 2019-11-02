// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ApplicationAddPasswordRequestParameter undocumented
type ApplicationAddPasswordRequestParameter struct {
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
}

//
type ApplicationAddPasswordRequestBuilder struct{ BaseRequestBuilder }

// AddPassword action undocumented
func (b *ApplicationRequestBuilder) AddPassword(reqObj *ApplicationAddPasswordRequestParameter) *ApplicationAddPasswordRequestBuilder {
	bb := &ApplicationAddPasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addPassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ApplicationAddPasswordRequest struct{ BaseRequest }

//
func (b *ApplicationAddPasswordRequestBuilder) Request() *ApplicationAddPasswordRequest {
	return &ApplicationAddPasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ApplicationAddPasswordRequest) Post(ctx context.Context) (resObj *PasswordCredential, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
