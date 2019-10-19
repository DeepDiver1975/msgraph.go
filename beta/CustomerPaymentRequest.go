// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CustomerPaymentRequestBuilder is request builder for CustomerPayment
type CustomerPaymentRequestBuilder struct{ BaseRequestBuilder }

// Request returns CustomerPaymentRequest
func (b *CustomerPaymentRequestBuilder) Request() *CustomerPaymentRequest {
	return &CustomerPaymentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CustomerPaymentRequest is request for CustomerPayment
type CustomerPaymentRequest struct{ BaseRequest }

// Do performs HTTP request for CustomerPayment
func (r *CustomerPaymentRequest) Do(method, path string, reqObj interface{}) (resObj *CustomerPayment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for CustomerPayment
func (r *CustomerPaymentRequest) Get() (*CustomerPayment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for CustomerPayment
func (r *CustomerPaymentRequest) Update(reqObj *CustomerPayment) (*CustomerPayment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for CustomerPayment
func (r *CustomerPaymentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Customer is navigation property
func (b *CustomerPaymentRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}