// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementAutopilotEventRequestBuilder is request builder for DeviceManagementAutopilotEvent
type DeviceManagementAutopilotEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementAutopilotEventRequest
func (b *DeviceManagementAutopilotEventRequestBuilder) Request() *DeviceManagementAutopilotEventRequest {
	return &DeviceManagementAutopilotEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementAutopilotEventRequest is request for DeviceManagementAutopilotEvent
type DeviceManagementAutopilotEventRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementAutopilotEvent, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Get() (*DeviceManagementAutopilotEvent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Update(reqObj *DeviceManagementAutopilotEvent) (*DeviceManagementAutopilotEvent, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}