// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// EducationOutcomeRequestBuilder is request builder for EducationOutcome
type EducationOutcomeRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationOutcomeRequest
func (b *EducationOutcomeRequestBuilder) Request() *EducationOutcomeRequest {
	return &EducationOutcomeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationOutcomeRequest is request for EducationOutcome
type EducationOutcomeRequest struct{ BaseRequest }

// Get performs GET request for EducationOutcome
func (r *EducationOutcomeRequest) Get(ctx context.Context) (resObj *EducationOutcome, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EducationOutcome
func (r *EducationOutcomeRequest) Update(ctx context.Context, reqObj *EducationOutcome) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EducationOutcome
func (r *EducationOutcomeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
