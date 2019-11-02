// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ChatRequestBuilder is request builder for Chat
type ChatRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatRequest
func (b *ChatRequestBuilder) Request() *ChatRequest {
	return &ChatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatRequest is request for Chat
type ChatRequest struct{ BaseRequest }

// Get performs GET request for Chat
func (r *ChatRequest) Get(ctx context.Context) (resObj *Chat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Chat
func (r *ChatRequest) Update(ctx context.Context, reqObj *Chat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Chat
func (r *ChatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *ChatRequestBuilder) InstalledApps() *ChatInstalledAppsCollectionRequestBuilder {
	bb := &ChatInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// ChatInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection
type ChatInstalledAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppInstallation collection
func (b *ChatInstalledAppsCollectionRequestBuilder) Request() *ChatInstalledAppsCollectionRequest {
	return &ChatInstalledAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppInstallation item
func (b *ChatInstalledAppsCollectionRequestBuilder) ID(id string) *TeamsAppInstallationRequestBuilder {
	bb := &TeamsAppInstallationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatInstalledAppsCollectionRequest is request for TeamsAppInstallation collection
type ChatInstalledAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TeamsAppInstallation, error) {
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
	var values []TeamsAppInstallation
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
			value  []TeamsAppInstallation
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

// Get performs GET request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Get(ctx context.Context) ([]TeamsAppInstallation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAppInstallation) (resObj *TeamsAppInstallation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Members returns request builder for ConversationMember collection
func (b *ChatRequestBuilder) Members() *ChatMembersCollectionRequestBuilder {
	bb := &ChatMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// ChatMembersCollectionRequestBuilder is request builder for ConversationMember collection
type ChatMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConversationMember collection
func (b *ChatMembersCollectionRequestBuilder) Request() *ChatMembersCollectionRequest {
	return &ChatMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConversationMember item
func (b *ChatMembersCollectionRequestBuilder) ID(id string) *ConversationMemberRequestBuilder {
	bb := &ConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMembersCollectionRequest is request for ConversationMember collection
type ChatMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConversationMember collection
func (r *ChatMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ConversationMember, error) {
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
	var values []ConversationMember
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
			value  []ConversationMember
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

// Get performs GET request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Get(ctx context.Context) ([]ConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Add(ctx context.Context, reqObj *ConversationMember) (resObj *ConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Messages returns request builder for ChatMessage collection
func (b *ChatRequestBuilder) Messages() *ChatMessagesCollectionRequestBuilder {
	bb := &ChatMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messages"
	return bb
}

// ChatMessagesCollectionRequestBuilder is request builder for ChatMessage collection
type ChatMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChatMessage collection
func (b *ChatMessagesCollectionRequestBuilder) Request() *ChatMessagesCollectionRequest {
	return &ChatMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChatMessage item
func (b *ChatMessagesCollectionRequestBuilder) ID(id string) *ChatMessageRequestBuilder {
	bb := &ChatMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMessagesCollectionRequest is request for ChatMessage collection
type ChatMessagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ChatMessage, error) {
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
	var values []ChatMessage
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
			value  []ChatMessage
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

// Get performs GET request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Get(ctx context.Context) ([]ChatMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Add(ctx context.Context, reqObj *ChatMessage) (resObj *ChatMessage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
