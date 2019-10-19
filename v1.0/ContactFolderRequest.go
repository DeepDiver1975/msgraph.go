// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ContactFolderRequestBuilder is request builder for ContactFolder
type ContactFolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContactFolderRequest
func (b *ContactFolderRequestBuilder) Request() *ContactFolderRequest {
	return &ContactFolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContactFolderRequest is request for ContactFolder
type ContactFolderRequest struct{ BaseRequest }

// Do performs HTTP request for ContactFolder
func (r *ContactFolderRequest) Do(method, path string, reqObj interface{}) (resObj *ContactFolder, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ContactFolder
func (r *ContactFolderRequest) Get() (*ContactFolder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ContactFolder
func (r *ContactFolderRequest) Update(reqObj *ContactFolder) (*ContactFolder, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ContactFolder
func (r *ContactFolderRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ChildFolders returns request builder for ContactFolder collection
func (b *ContactFolderRequestBuilder) ChildFolders() *ContactFolderChildFoldersCollectionRequestBuilder {
	bb := &ContactFolderChildFoldersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/childFolders"
	return bb
}

// ContactFolderChildFoldersCollectionRequestBuilder is request builder for ContactFolder collection
type ContactFolderChildFoldersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContactFolder collection
func (b *ContactFolderChildFoldersCollectionRequestBuilder) Request() *ContactFolderChildFoldersCollectionRequest {
	return &ContactFolderChildFoldersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContactFolder item
func (b *ContactFolderChildFoldersCollectionRequestBuilder) ID(id string) *ContactFolderRequestBuilder {
	bb := &ContactFolderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContactFolderChildFoldersCollectionRequest is request for ContactFolder collection
type ContactFolderChildFoldersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ContactFolder collection
func (r *ContactFolderChildFoldersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ContactFolder, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ContactFolder collection
func (r *ContactFolderChildFoldersCollectionRequest) Paging(method, path string, obj interface{}) ([]ContactFolder, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ContactFolder
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ContactFolder
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ContactFolder collection
func (r *ContactFolderChildFoldersCollectionRequest) Get() ([]ContactFolder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ContactFolder collection
func (r *ContactFolderChildFoldersCollectionRequest) Add(reqObj *ContactFolder) (*ContactFolder, error) {
	return r.Do("POST", "", reqObj)
}

// Contacts returns request builder for Contact collection
func (b *ContactFolderRequestBuilder) Contacts() *ContactFolderContactsCollectionRequestBuilder {
	bb := &ContactFolderContactsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contacts"
	return bb
}

// ContactFolderContactsCollectionRequestBuilder is request builder for Contact collection
type ContactFolderContactsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Contact collection
func (b *ContactFolderContactsCollectionRequestBuilder) Request() *ContactFolderContactsCollectionRequest {
	return &ContactFolderContactsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Contact item
func (b *ContactFolderContactsCollectionRequestBuilder) ID(id string) *ContactRequestBuilder {
	bb := &ContactRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContactFolderContactsCollectionRequest is request for Contact collection
type ContactFolderContactsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Contact collection
func (r *ContactFolderContactsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Contact, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Contact collection
func (r *ContactFolderContactsCollectionRequest) Paging(method, path string, obj interface{}) ([]Contact, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Contact
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Contact
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for Contact collection
func (r *ContactFolderContactsCollectionRequest) Get() ([]Contact, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Contact collection
func (r *ContactFolderContactsCollectionRequest) Add(reqObj *Contact) (*Contact, error) {
	return r.Do("POST", "", reqObj)
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *ContactFolderRequestBuilder) MultiValueExtendedProperties() *ContactFolderMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &ContactFolderMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// ContactFolderMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type ContactFolderMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *ContactFolderMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *ContactFolderMultiValueExtendedPropertiesCollectionRequest {
	return &ContactFolderMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *ContactFolderMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContactFolderMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type ContactFolderMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MultiValueLegacyExtendedProperty collection
func (r *ContactFolderMultiValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *ContactFolderMultiValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MultiValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MultiValueLegacyExtendedProperty
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *ContactFolderMultiValueExtendedPropertiesCollectionRequest) Get() ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *ContactFolderMultiValueExtendedPropertiesCollectionRequest) Add(reqObj *MultiValueLegacyExtendedProperty) (*MultiValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *ContactFolderRequestBuilder) SingleValueExtendedProperties() *ContactFolderSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &ContactFolderSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// ContactFolderSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type ContactFolderSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *ContactFolderSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *ContactFolderSingleValueExtendedPropertiesCollectionRequest {
	return &ContactFolderSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *ContactFolderSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContactFolderSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type ContactFolderSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SingleValueLegacyExtendedProperty collection
func (r *ContactFolderSingleValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *ContactFolderSingleValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SingleValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SingleValueLegacyExtendedProperty
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *ContactFolderSingleValueExtendedPropertiesCollectionRequest) Get() ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *ContactFolderSingleValueExtendedPropertiesCollectionRequest) Add(reqObj *SingleValueLegacyExtendedProperty) (*SingleValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}