// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsInformationProtectionAppLearningSummaryRequestBuilder is request builder for WindowsInformationProtectionAppLearningSummary
type WindowsInformationProtectionAppLearningSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionAppLearningSummaryRequest
func (b *WindowsInformationProtectionAppLearningSummaryRequestBuilder) Request() *WindowsInformationProtectionAppLearningSummaryRequest {
	return &WindowsInformationProtectionAppLearningSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionAppLearningSummaryRequest is request for WindowsInformationProtectionAppLearningSummary
type WindowsInformationProtectionAppLearningSummaryRequest struct{ BaseRequest }

// Get performs GET request for WindowsInformationProtectionAppLearningSummary
func (r *WindowsInformationProtectionAppLearningSummaryRequest) Get(ctx context.Context) (resObj *WindowsInformationProtectionAppLearningSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsInformationProtectionAppLearningSummary
func (r *WindowsInformationProtectionAppLearningSummaryRequest) Update(ctx context.Context, reqObj *WindowsInformationProtectionAppLearningSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsInformationProtectionAppLearningSummary
func (r *WindowsInformationProtectionAppLearningSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
