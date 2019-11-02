// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ChatMessageHostedContentRequestBuilder is request builder for ChatMessageHostedContent
type ChatMessageHostedContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatMessageHostedContentRequest
func (b *ChatMessageHostedContentRequestBuilder) Request() *ChatMessageHostedContentRequest {
	return &ChatMessageHostedContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatMessageHostedContentRequest is request for ChatMessageHostedContent
type ChatMessageHostedContentRequest struct{ BaseRequest }

// Get performs GET request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Get(ctx context.Context) (resObj *ChatMessageHostedContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Update(ctx context.Context, reqObj *ChatMessageHostedContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ChatMessageHostedContent
func (r *ChatMessageHostedContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
