// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UsedInsightRequestBuilder is request builder for UsedInsight
type UsedInsightRequestBuilder struct{ BaseRequestBuilder }

// Request returns UsedInsightRequest
func (b *UsedInsightRequestBuilder) Request() *UsedInsightRequest {
	return &UsedInsightRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UsedInsightRequest is request for UsedInsight
type UsedInsightRequest struct{ BaseRequest }

// Do performs HTTP request for UsedInsight
func (r *UsedInsightRequest) Do(method, path string, reqObj interface{}) (resObj *UsedInsight, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for UsedInsight
func (r *UsedInsightRequest) Get() (*UsedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for UsedInsight
func (r *UsedInsightRequest) Update(reqObj *UsedInsight) (*UsedInsight, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for UsedInsight
func (r *UsedInsightRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Resource is navigation property
func (b *UsedInsightRequestBuilder) Resource() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}