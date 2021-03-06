// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SalesInvoiceLineRequestBuilder is request builder for SalesInvoiceLine
type SalesInvoiceLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesInvoiceLineRequest
func (b *SalesInvoiceLineRequestBuilder) Request() *SalesInvoiceLineRequest {
	return &SalesInvoiceLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesInvoiceLineRequest is request for SalesInvoiceLine
type SalesInvoiceLineRequest struct{ BaseRequest }

// Get performs GET request for SalesInvoiceLine
func (r *SalesInvoiceLineRequest) Get(ctx context.Context) (resObj *SalesInvoiceLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesInvoiceLine
func (r *SalesInvoiceLineRequest) Update(ctx context.Context, reqObj *SalesInvoiceLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesInvoiceLine
func (r *SalesInvoiceLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Account is navigation property
func (b *SalesInvoiceLineRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// Item is navigation property
func (b *SalesInvoiceLineRequestBuilder) Item() *ItemRequestBuilder {
	bb := &ItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
