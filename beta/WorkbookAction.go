// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookCreateSessionRequestParameter undocumented
type WorkbookCreateSessionRequestParameter struct {
	// PersistChanges undocumented
	PersistChanges *bool `json:"persistChanges,omitempty"`
	// CommitExplicitly undocumented
	CommitExplicitly *bool `json:"commitExplicitly,omitempty"`
}

// WorkbookCloseSessionRequestParameter undocumented
type WorkbookCloseSessionRequestParameter struct {
}

// WorkbookRefreshSessionRequestParameter undocumented
type WorkbookRefreshSessionRequestParameter struct {
}

//
type WorkbookCreateSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateSession action undocumented
func (b *WorkbookRequestBuilder) CreateSession(reqObj *WorkbookCreateSessionRequestParameter) *WorkbookCreateSessionRequestBuilder {
	bb := &WorkbookCreateSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookCreateSessionRequest struct{ BaseRequest }

//
func (b *WorkbookCreateSessionRequestBuilder) Request() *WorkbookCreateSessionRequest {
	return &WorkbookCreateSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookCreateSessionRequest) Post(ctx context.Context) (resObj *WorkbookSessionInfo, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookCloseSessionRequestBuilder struct{ BaseRequestBuilder }

// CloseSession action undocumented
func (b *WorkbookRequestBuilder) CloseSession(reqObj *WorkbookCloseSessionRequestParameter) *WorkbookCloseSessionRequestBuilder {
	bb := &WorkbookCloseSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/closeSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookCloseSessionRequest struct{ BaseRequest }

//
func (b *WorkbookCloseSessionRequestBuilder) Request() *WorkbookCloseSessionRequest {
	return &WorkbookCloseSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookCloseSessionRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type WorkbookRefreshSessionRequestBuilder struct{ BaseRequestBuilder }

// RefreshSession action undocumented
func (b *WorkbookRequestBuilder) RefreshSession(reqObj *WorkbookRefreshSessionRequestParameter) *WorkbookRefreshSessionRequestBuilder {
	bb := &WorkbookRefreshSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/refreshSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookRefreshSessionRequest struct{ BaseRequest }

//
func (b *WorkbookRefreshSessionRequestBuilder) Request() *WorkbookRefreshSessionRequest {
	return &WorkbookRefreshSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookRefreshSessionRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
