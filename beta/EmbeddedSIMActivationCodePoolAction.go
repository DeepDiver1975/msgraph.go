// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EmbeddedSIMActivationCodePoolAssignRequestParameter undocumented
type EmbeddedSIMActivationCodePoolAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
}

//
type EmbeddedSIMActivationCodePoolAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Assign(reqObj *EmbeddedSIMActivationCodePoolAssignRequestParameter) *EmbeddedSIMActivationCodePoolAssignRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EmbeddedSIMActivationCodePoolAssignRequest struct{ BaseRequest }

//
func (b *EmbeddedSIMActivationCodePoolAssignRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignRequest {
	return &EmbeddedSIMActivationCodePoolAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EmbeddedSIMActivationCodePoolAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]EmbeddedSIMActivationCodePoolAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]EmbeddedSIMActivationCodePoolAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  [][]EmbeddedSIMActivationCodePoolAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *EmbeddedSIMActivationCodePoolAssignRequest) Get(ctx context.Context) ([][]EmbeddedSIMActivationCodePoolAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}
