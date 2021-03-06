// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOSWiredNetworkConfigurationRequestBuilder is request builder for MacOSWiredNetworkConfiguration
type MacOSWiredNetworkConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSWiredNetworkConfigurationRequest
func (b *MacOSWiredNetworkConfigurationRequestBuilder) Request() *MacOSWiredNetworkConfigurationRequest {
	return &MacOSWiredNetworkConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSWiredNetworkConfigurationRequest is request for MacOSWiredNetworkConfiguration
type MacOSWiredNetworkConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Get(ctx context.Context) (resObj *MacOSWiredNetworkConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Update(ctx context.Context, reqObj *MacOSWiredNetworkConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSWiredNetworkConfiguration
func (r *MacOSWiredNetworkConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *MacOSWiredNetworkConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *MacOSWiredNetworkConfigurationRequestBuilder) RootCertificateForServerValidation() *MacOSTrustedRootCertificateRequestBuilder {
	bb := &MacOSTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}
