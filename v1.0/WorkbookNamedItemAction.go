// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"encoding/json"
)

// WorkbookNamedItemCollectionAddRequestParameter undocumented
type WorkbookNamedItemCollectionAddRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Reference undocumented
	Reference json.RawMessage `json:"reference,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
}

// WorkbookNamedItemCollectionAddFormulaLocalRequestParameter undocumented
type WorkbookNamedItemCollectionAddFormulaLocalRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Formula undocumented
	Formula *string `json:"formula,omitempty"`
	// Comment undocumented
	Comment *string `json:"comment,omitempty"`
}

//
type WorkbookNamedItemCollectionAddRequestBuilder struct{ BaseRequestBuilder }

// Add action undocumented
func (b *WorkbookNamesCollectionRequestBuilder) Add(reqObj *WorkbookNamedItemCollectionAddRequestParameter) *WorkbookNamedItemCollectionAddRequestBuilder {
	bb := &WorkbookNamedItemCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *WorkbookWorksheetNamesCollectionRequestBuilder) Add(reqObj *WorkbookNamedItemCollectionAddRequestParameter) *WorkbookNamedItemCollectionAddRequestBuilder {
	bb := &WorkbookNamedItemCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookNamedItemCollectionAddRequest struct{ BaseRequest }

//
func (b *WorkbookNamedItemCollectionAddRequestBuilder) Request() *WorkbookNamedItemCollectionAddRequest {
	return &WorkbookNamedItemCollectionAddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookNamedItemCollectionAddRequest) Post(ctx context.Context) (resObj *WorkbookNamedItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookNamedItemCollectionAddFormulaLocalRequestBuilder struct{ BaseRequestBuilder }

// AddFormulaLocal action undocumented
func (b *WorkbookNamesCollectionRequestBuilder) AddFormulaLocal(reqObj *WorkbookNamedItemCollectionAddFormulaLocalRequestParameter) *WorkbookNamedItemCollectionAddFormulaLocalRequestBuilder {
	bb := &WorkbookNamedItemCollectionAddFormulaLocalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addFormulaLocal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// AddFormulaLocal action undocumented
func (b *WorkbookWorksheetNamesCollectionRequestBuilder) AddFormulaLocal(reqObj *WorkbookNamedItemCollectionAddFormulaLocalRequestParameter) *WorkbookNamedItemCollectionAddFormulaLocalRequestBuilder {
	bb := &WorkbookNamedItemCollectionAddFormulaLocalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addFormulaLocal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookNamedItemCollectionAddFormulaLocalRequest struct{ BaseRequest }

//
func (b *WorkbookNamedItemCollectionAddFormulaLocalRequestBuilder) Request() *WorkbookNamedItemCollectionAddFormulaLocalRequest {
	return &WorkbookNamedItemCollectionAddFormulaLocalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookNamedItemCollectionAddFormulaLocalRequest) Post(ctx context.Context) (resObj *WorkbookNamedItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
