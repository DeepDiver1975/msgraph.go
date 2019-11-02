// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// BusinessFlowTemplateRequestBuilder is request builder for BusinessFlowTemplate
type BusinessFlowTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns BusinessFlowTemplateRequest
func (b *BusinessFlowTemplateRequestBuilder) Request() *BusinessFlowTemplateRequest {
	return &BusinessFlowTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BusinessFlowTemplateRequest is request for BusinessFlowTemplate
type BusinessFlowTemplateRequest struct{ BaseRequest }

// Get performs GET request for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) Get(ctx context.Context) (resObj *BusinessFlowTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) Update(ctx context.Context, reqObj *BusinessFlowTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BusinessFlowTemplate
func (r *BusinessFlowTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
