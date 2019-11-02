// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookChartAxisTitleRequestBuilder is request builder for WorkbookChartAxisTitle
type WorkbookChartAxisTitleRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartAxisTitleRequest
func (b *WorkbookChartAxisTitleRequestBuilder) Request() *WorkbookChartAxisTitleRequest {
	return &WorkbookChartAxisTitleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartAxisTitleRequest is request for WorkbookChartAxisTitle
type WorkbookChartAxisTitleRequest struct{ BaseRequest }

// Get performs GET request for WorkbookChartAxisTitle
func (r *WorkbookChartAxisTitleRequest) Get(ctx context.Context) (resObj *WorkbookChartAxisTitle, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChartAxisTitle
func (r *WorkbookChartAxisTitleRequest) Update(ctx context.Context, reqObj *WorkbookChartAxisTitle) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChartAxisTitle
func (r *WorkbookChartAxisTitleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Format is navigation property
func (b *WorkbookChartAxisTitleRequestBuilder) Format() *WorkbookChartAxisTitleFormatRequestBuilder {
	bb := &WorkbookChartAxisTitleFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}
