// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ActivityHistoryItemRequestBuilder is request builder for ActivityHistoryItem
type ActivityHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityHistoryItemRequest
func (b *ActivityHistoryItemRequestBuilder) Request() *ActivityHistoryItemRequest {
	return &ActivityHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityHistoryItemRequest is request for ActivityHistoryItem
type ActivityHistoryItemRequest struct{ BaseRequest }

// Do performs HTTP request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Do(method, path string, reqObj interface{}) (resObj *ActivityHistoryItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Get() (*ActivityHistoryItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Update(reqObj *ActivityHistoryItem) (*ActivityHistoryItem, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Activity is navigation property
func (b *ActivityHistoryItemRequestBuilder) Activity() *UserActivityRequestBuilder {
	bb := &UserActivityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activity"
	return bb
}