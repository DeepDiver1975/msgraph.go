// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ProgramRequestBuilder is request builder for Program
type ProgramRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProgramRequest
func (b *ProgramRequestBuilder) Request() *ProgramRequest {
	return &ProgramRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProgramRequest is request for Program
type ProgramRequest struct{ BaseRequest }

// Get performs GET request for Program
func (r *ProgramRequest) Get(ctx context.Context) (resObj *Program, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Program
func (r *ProgramRequest) Update(ctx context.Context, reqObj *Program) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Program
func (r *ProgramRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Controls returns request builder for ProgramControl collection
func (b *ProgramRequestBuilder) Controls() *ProgramControlsCollectionRequestBuilder {
	bb := &ProgramControlsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/controls"
	return bb
}

// ProgramControlsCollectionRequestBuilder is request builder for ProgramControl collection
type ProgramControlsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ProgramControl collection
func (b *ProgramControlsCollectionRequestBuilder) Request() *ProgramControlsCollectionRequest {
	return &ProgramControlsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ProgramControl item
func (b *ProgramControlsCollectionRequestBuilder) ID(id string) *ProgramControlRequestBuilder {
	bb := &ProgramControlRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ProgramControlsCollectionRequest is request for ProgramControl collection
type ProgramControlsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ProgramControl, error) {
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
	var values []ProgramControl
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
			value  []ProgramControl
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

// Get performs GET request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Get(ctx context.Context) ([]ProgramControl, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ProgramControl collection
func (r *ProgramControlsCollectionRequest) Add(ctx context.Context, reqObj *ProgramControl) (resObj *ProgramControl, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
