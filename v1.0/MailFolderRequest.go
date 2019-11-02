// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// MailFolderRequestBuilder is request builder for MailFolder
type MailFolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns MailFolderRequest
func (b *MailFolderRequestBuilder) Request() *MailFolderRequest {
	return &MailFolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MailFolderRequest is request for MailFolder
type MailFolderRequest struct{ BaseRequest }

// Get performs GET request for MailFolder
func (r *MailFolderRequest) Get(ctx context.Context) (resObj *MailFolder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MailFolder
func (r *MailFolderRequest) Update(ctx context.Context, reqObj *MailFolder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MailFolder
func (r *MailFolderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ChildFolders returns request builder for MailFolder collection
func (b *MailFolderRequestBuilder) ChildFolders() *MailFolderChildFoldersCollectionRequestBuilder {
	bb := &MailFolderChildFoldersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/childFolders"
	return bb
}

// MailFolderChildFoldersCollectionRequestBuilder is request builder for MailFolder collection
type MailFolderChildFoldersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MailFolder collection
func (b *MailFolderChildFoldersCollectionRequestBuilder) Request() *MailFolderChildFoldersCollectionRequest {
	return &MailFolderChildFoldersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MailFolder item
func (b *MailFolderChildFoldersCollectionRequestBuilder) ID(id string) *MailFolderRequestBuilder {
	bb := &MailFolderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderChildFoldersCollectionRequest is request for MailFolder collection
type MailFolderChildFoldersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MailFolder, error) {
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
	var values []MailFolder
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
			value  []MailFolder
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

// Get performs GET request for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Get(ctx context.Context) ([]MailFolder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Add(ctx context.Context, reqObj *MailFolder) (resObj *MailFolder, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MessageRules returns request builder for MessageRule collection
func (b *MailFolderRequestBuilder) MessageRules() *MailFolderMessageRulesCollectionRequestBuilder {
	bb := &MailFolderMessageRulesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messageRules"
	return bb
}

// MailFolderMessageRulesCollectionRequestBuilder is request builder for MessageRule collection
type MailFolderMessageRulesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MessageRule collection
func (b *MailFolderMessageRulesCollectionRequestBuilder) Request() *MailFolderMessageRulesCollectionRequest {
	return &MailFolderMessageRulesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MessageRule item
func (b *MailFolderMessageRulesCollectionRequestBuilder) ID(id string) *MessageRuleRequestBuilder {
	bb := &MessageRuleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderMessageRulesCollectionRequest is request for MessageRule collection
type MailFolderMessageRulesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MessageRule, error) {
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
	var values []MessageRule
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
			value  []MessageRule
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

// Get performs GET request for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Get(ctx context.Context) ([]MessageRule, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Add(ctx context.Context, reqObj *MessageRule) (resObj *MessageRule, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Messages returns request builder for Message collection
func (b *MailFolderRequestBuilder) Messages() *MailFolderMessagesCollectionRequestBuilder {
	bb := &MailFolderMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messages"
	return bb
}

// MailFolderMessagesCollectionRequestBuilder is request builder for Message collection
type MailFolderMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Message collection
func (b *MailFolderMessagesCollectionRequestBuilder) Request() *MailFolderMessagesCollectionRequest {
	return &MailFolderMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Message item
func (b *MailFolderMessagesCollectionRequestBuilder) ID(id string) *MessageRequestBuilder {
	bb := &MessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderMessagesCollectionRequest is request for Message collection
type MailFolderMessagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Message collection
func (r *MailFolderMessagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Message, error) {
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
	var values []Message
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
			value  []Message
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

// Get performs GET request for Message collection
func (r *MailFolderMessagesCollectionRequest) Get(ctx context.Context) ([]Message, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Message collection
func (r *MailFolderMessagesCollectionRequest) Add(ctx context.Context, reqObj *Message) (resObj *Message, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *MailFolderRequestBuilder) MultiValueExtendedProperties() *MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *MailFolderMultiValueExtendedPropertiesCollectionRequest {
	return &MailFolderMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type MailFolderMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
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
	var values []MultiValueLegacyExtendedProperty
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
			value  []MultiValueLegacyExtendedProperty
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

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *MultiValueLegacyExtendedProperty) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *MailFolderRequestBuilder) SingleValueExtendedProperties() *MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *MailFolderSingleValueExtendedPropertiesCollectionRequest {
	return &MailFolderSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type MailFolderSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
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
	var values []SingleValueLegacyExtendedProperty
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
			value  []SingleValueLegacyExtendedProperty
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

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Get(ctx context.Context) ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Add(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
