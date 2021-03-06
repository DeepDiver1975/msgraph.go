// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookChartDataLabelsRequestBuilder is request builder for WorkbookChartDataLabels
type WorkbookChartDataLabelsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartDataLabelsRequest
func (b *WorkbookChartDataLabelsRequestBuilder) Request() *WorkbookChartDataLabelsRequest {
	return &WorkbookChartDataLabelsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartDataLabelsRequest is request for WorkbookChartDataLabels
type WorkbookChartDataLabelsRequest struct{ BaseRequest }

// Get performs GET request for WorkbookChartDataLabels
func (r *WorkbookChartDataLabelsRequest) Get(ctx context.Context) (resObj *WorkbookChartDataLabels, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChartDataLabels
func (r *WorkbookChartDataLabelsRequest) Update(ctx context.Context, reqObj *WorkbookChartDataLabels) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChartDataLabels
func (r *WorkbookChartDataLabelsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Format is navigation property
func (b *WorkbookChartDataLabelsRequestBuilder) Format() *WorkbookChartDataLabelFormatRequestBuilder {
	bb := &WorkbookChartDataLabelFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}
