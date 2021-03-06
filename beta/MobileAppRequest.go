// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// MobileAppRequestBuilder is request builder for MobileApp
type MobileAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppRequest
func (b *MobileAppRequestBuilder) Request() *MobileAppRequest {
	return &MobileAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppRequest is request for MobileApp
type MobileAppRequest struct{ BaseRequest }

// Get performs GET request for MobileApp
func (r *MobileAppRequest) Get(ctx context.Context) (resObj *MobileApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileApp
func (r *MobileAppRequest) Update(ctx context.Context, reqObj *MobileApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileApp
func (r *MobileAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for MobileAppAssignment collection
func (b *MobileAppRequestBuilder) Assignments() *MobileAppAssignmentsCollectionRequestBuilder {
	bb := &MobileAppAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// MobileAppAssignmentsCollectionRequestBuilder is request builder for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppAssignment collection
func (b *MobileAppAssignmentsCollectionRequestBuilder) Request() *MobileAppAssignmentsCollectionRequest {
	return &MobileAppAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppAssignment item
func (b *MobileAppAssignmentsCollectionRequestBuilder) ID(id string) *MobileAppAssignmentRequestBuilder {
	bb := &MobileAppAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppAssignmentsCollectionRequest is request for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppAssignment, error) {
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
	var values []MobileAppAssignment
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
			value  []MobileAppAssignment
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

// Get performs GET request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Get(ctx context.Context) ([]MobileAppAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *MobileAppAssignment) (resObj *MobileAppAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Categories returns request builder for MobileAppCategory collection
func (b *MobileAppRequestBuilder) Categories() *MobileAppCategoriesCollectionRequestBuilder {
	bb := &MobileAppCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/categories"
	return bb
}

// MobileAppCategoriesCollectionRequestBuilder is request builder for MobileAppCategory collection
type MobileAppCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppCategory collection
func (b *MobileAppCategoriesCollectionRequestBuilder) Request() *MobileAppCategoriesCollectionRequest {
	return &MobileAppCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppCategory item
func (b *MobileAppCategoriesCollectionRequestBuilder) ID(id string) *MobileAppCategoryRequestBuilder {
	bb := &MobileAppCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppCategoriesCollectionRequest is request for MobileAppCategory collection
type MobileAppCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppCategory, error) {
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
	var values []MobileAppCategory
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
			value  []MobileAppCategory
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

// Get performs GET request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Get(ctx context.Context) ([]MobileAppCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Add(ctx context.Context, reqObj *MobileAppCategory) (resObj *MobileAppCategory, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceStatuses returns request builder for MobileAppInstallStatus collection
func (b *MobileAppRequestBuilder) DeviceStatuses() *MobileAppDeviceStatusesCollectionRequestBuilder {
	bb := &MobileAppDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// MobileAppDeviceStatusesCollectionRequestBuilder is request builder for MobileAppInstallStatus collection
type MobileAppDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppInstallStatus collection
func (b *MobileAppDeviceStatusesCollectionRequestBuilder) Request() *MobileAppDeviceStatusesCollectionRequest {
	return &MobileAppDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppInstallStatus item
func (b *MobileAppDeviceStatusesCollectionRequestBuilder) ID(id string) *MobileAppInstallStatusRequestBuilder {
	bb := &MobileAppInstallStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppDeviceStatusesCollectionRequest is request for MobileAppInstallStatus collection
type MobileAppDeviceStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppInstallStatus, error) {
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
	var values []MobileAppInstallStatus
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
			value  []MobileAppInstallStatus
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

// Get performs GET request for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Get(ctx context.Context) ([]MobileAppInstallStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppInstallStatus collection
func (r *MobileAppDeviceStatusesCollectionRequest) Add(ctx context.Context, reqObj *MobileAppInstallStatus) (resObj *MobileAppInstallStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// InstallSummary is navigation property
func (b *MobileAppRequestBuilder) InstallSummary() *MobileAppInstallSummaryRequestBuilder {
	bb := &MobileAppInstallSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installSummary"
	return bb
}

// Relationships returns request builder for MobileAppRelationship collection
func (b *MobileAppRequestBuilder) Relationships() *MobileAppRelationshipsCollectionRequestBuilder {
	bb := &MobileAppRelationshipsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/relationships"
	return bb
}

// MobileAppRelationshipsCollectionRequestBuilder is request builder for MobileAppRelationship collection
type MobileAppRelationshipsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppRelationship collection
func (b *MobileAppRelationshipsCollectionRequestBuilder) Request() *MobileAppRelationshipsCollectionRequest {
	return &MobileAppRelationshipsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppRelationship item
func (b *MobileAppRelationshipsCollectionRequestBuilder) ID(id string) *MobileAppRelationshipRequestBuilder {
	bb := &MobileAppRelationshipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppRelationshipsCollectionRequest is request for MobileAppRelationship collection
type MobileAppRelationshipsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppRelationship, error) {
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
	var values []MobileAppRelationship
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
			value  []MobileAppRelationship
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

// Get performs GET request for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Get(ctx context.Context) ([]MobileAppRelationship, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppRelationship collection
func (r *MobileAppRelationshipsCollectionRequest) Add(ctx context.Context, reqObj *MobileAppRelationship) (resObj *MobileAppRelationship, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserStatuses returns request builder for UserAppInstallStatus collection
func (b *MobileAppRequestBuilder) UserStatuses() *MobileAppUserStatusesCollectionRequestBuilder {
	bb := &MobileAppUserStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatuses"
	return bb
}

// MobileAppUserStatusesCollectionRequestBuilder is request builder for UserAppInstallStatus collection
type MobileAppUserStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserAppInstallStatus collection
func (b *MobileAppUserStatusesCollectionRequestBuilder) Request() *MobileAppUserStatusesCollectionRequest {
	return &MobileAppUserStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserAppInstallStatus item
func (b *MobileAppUserStatusesCollectionRequestBuilder) ID(id string) *UserAppInstallStatusRequestBuilder {
	bb := &UserAppInstallStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppUserStatusesCollectionRequest is request for UserAppInstallStatus collection
type MobileAppUserStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]UserAppInstallStatus, error) {
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
	var values []UserAppInstallStatus
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
			value  []UserAppInstallStatus
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

// Get performs GET request for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Get(ctx context.Context) ([]UserAppInstallStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for UserAppInstallStatus collection
func (r *MobileAppUserStatusesCollectionRequest) Add(ctx context.Context, reqObj *UserAppInstallStatus) (resObj *UserAppInstallStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
