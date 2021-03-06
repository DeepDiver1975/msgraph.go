// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceManagementTroubleshootingEventRequestBuilder is request builder for DeviceManagementTroubleshootingEvent
type DeviceManagementTroubleshootingEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementTroubleshootingEventRequest
func (b *DeviceManagementTroubleshootingEventRequestBuilder) Request() *DeviceManagementTroubleshootingEventRequest {
	return &DeviceManagementTroubleshootingEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementTroubleshootingEventRequest is request for DeviceManagementTroubleshootingEvent
type DeviceManagementTroubleshootingEventRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Get(ctx context.Context) (resObj *DeviceManagementTroubleshootingEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Update(ctx context.Context, reqObj *DeviceManagementTroubleshootingEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
