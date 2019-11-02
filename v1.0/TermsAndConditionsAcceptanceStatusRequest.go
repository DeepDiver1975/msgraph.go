// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TermsAndConditionsAcceptanceStatusRequestBuilder is request builder for TermsAndConditionsAcceptanceStatus
type TermsAndConditionsAcceptanceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermsAndConditionsAcceptanceStatusRequest
func (b *TermsAndConditionsAcceptanceStatusRequestBuilder) Request() *TermsAndConditionsAcceptanceStatusRequest {
	return &TermsAndConditionsAcceptanceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermsAndConditionsAcceptanceStatusRequest is request for TermsAndConditionsAcceptanceStatus
type TermsAndConditionsAcceptanceStatusRequest struct{ BaseRequest }

// Get performs GET request for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) Get(ctx context.Context) (resObj *TermsAndConditionsAcceptanceStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) Update(ctx context.Context, reqObj *TermsAndConditionsAcceptanceStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TermsAndConditionsAcceptanceStatus
func (r *TermsAndConditionsAcceptanceStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TermsAndConditions is navigation property
func (b *TermsAndConditionsAcceptanceStatusRequestBuilder) TermsAndConditions() *TermsAndConditionsRequestBuilder {
	bb := &TermsAndConditionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/termsAndConditions"
	return bb
}
