//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// VPNConnectionsServer is a fake server for instances of the armnetwork.VPNConnectionsClient type.
type VPNConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method VPNConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters armnetwork.VPNConnection, options *armnetwork.VPNConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VPNConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VPNConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *armnetwork.VPNConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VPNConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VPNConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *armnetwork.VPNConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.VPNConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVPNGatewayPager is the fake for method VPNConnectionsClient.NewListByVPNGatewayPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVPNGatewayPager func(resourceGroupName string, gatewayName string, options *armnetwork.VPNConnectionsClientListByVPNGatewayOptions) (resp azfake.PagerResponder[armnetwork.VPNConnectionsClientListByVPNGatewayResponse])

	// BeginStartPacketCapture is the fake for method VPNConnectionsClient.BeginStartPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartPacketCapture func(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *armnetwork.VPNConnectionsClientBeginStartPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VPNConnectionsClientStartPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginStopPacketCapture is the fake for method VPNConnectionsClient.BeginStopPacketCapture
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStopPacketCapture func(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *armnetwork.VPNConnectionsClientBeginStopPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.VPNConnectionsClientStopPacketCaptureResponse], errResp azfake.ErrorResponder)
}

// NewVPNConnectionsServerTransport creates a new instance of VPNConnectionsServerTransport with the provided implementation.
// The returned VPNConnectionsServerTransport instance is connected to an instance of armnetwork.VPNConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVPNConnectionsServerTransport(srv *VPNConnectionsServer) *VPNConnectionsServerTransport {
	return &VPNConnectionsServerTransport{
		srv:                      srv,
		beginCreateOrUpdate:      newTracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientDeleteResponse]](),
		newListByVPNGatewayPager: newTracker[azfake.PagerResponder[armnetwork.VPNConnectionsClientListByVPNGatewayResponse]](),
		beginStartPacketCapture:  newTracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientStartPacketCaptureResponse]](),
		beginStopPacketCapture:   newTracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientStopPacketCaptureResponse]](),
	}
}

// VPNConnectionsServerTransport connects instances of armnetwork.VPNConnectionsClient to instances of VPNConnectionsServer.
// Don't use this type directly, use NewVPNConnectionsServerTransport instead.
type VPNConnectionsServerTransport struct {
	srv                      *VPNConnectionsServer
	beginCreateOrUpdate      *tracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientCreateOrUpdateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientDeleteResponse]]
	newListByVPNGatewayPager *tracker[azfake.PagerResponder[armnetwork.VPNConnectionsClientListByVPNGatewayResponse]]
	beginStartPacketCapture  *tracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientStartPacketCaptureResponse]]
	beginStopPacketCapture   *tracker[azfake.PollerResponder[armnetwork.VPNConnectionsClientStopPacketCaptureResponse]]
}

// Do implements the policy.Transporter interface for VPNConnectionsServerTransport.
func (v *VPNConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VPNConnectionsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VPNConnectionsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VPNConnectionsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VPNConnectionsClient.NewListByVPNGatewayPager":
		resp, err = v.dispatchNewListByVPNGatewayPager(req)
	case "VPNConnectionsClient.BeginStartPacketCapture":
		resp, err = v.dispatchBeginStartPacketCapture(req)
	case "VPNConnectionsClient.BeginStopPacketCapture":
		resp, err = v.dispatchBeginStopPacketCapture(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VPNConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, connectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VPNConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, connectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VPNConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, connectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VPNConnectionsServerTransport) dispatchNewListByVPNGatewayPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVPNGatewayPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVPNGatewayPager not implemented")}
	}
	newListByVPNGatewayPager := v.newListByVPNGatewayPager.get(req)
	if newListByVPNGatewayPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByVPNGatewayPager(resourceGroupNameUnescaped, gatewayNameUnescaped, nil)
		newListByVPNGatewayPager = &resp
		v.newListByVPNGatewayPager.add(req, newListByVPNGatewayPager)
		server.PagerResponderInjectNextLinks(newListByVPNGatewayPager, req, func(page *armnetwork.VPNConnectionsClientListByVPNGatewayResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVPNGatewayPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByVPNGatewayPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVPNGatewayPager) {
		v.newListByVPNGatewayPager.remove(req)
	}
	return resp, nil
}

func (v *VPNConnectionsServerTransport) dispatchBeginStartPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStartPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartPacketCapture not implemented")}
	}
	beginStartPacketCapture := v.beginStartPacketCapture.get(req)
	if beginStartPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections/(?P<vpnConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startpacketcapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNConnectionPacketCaptureStartParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		vpnConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vpnConnectionName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VPNConnectionsClientBeginStartPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VPNConnectionsClientBeginStartPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStartPacketCapture(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, vpnConnectionNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStartPacketCapture = &respr
		v.beginStartPacketCapture.add(req, beginStartPacketCapture)
	}

	resp, err := server.PollerResponderNext(beginStartPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginStartPacketCapture.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStartPacketCapture) {
		v.beginStartPacketCapture.remove(req)
	}

	return resp, nil
}

func (v *VPNConnectionsServerTransport) dispatchBeginStopPacketCapture(req *http.Request) (*http.Response, error) {
	if v.srv.BeginStopPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopPacketCapture not implemented")}
	}
	beginStopPacketCapture := v.beginStopPacketCapture.get(req)
	if beginStopPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConnections/(?P<vpnConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stoppacketcapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNConnectionPacketCaptureStopParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		vpnConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("vpnConnectionName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VPNConnectionsClientBeginStopPacketCaptureOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VPNConnectionsClientBeginStopPacketCaptureOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := v.srv.BeginStopPacketCapture(req.Context(), resourceGroupNameUnescaped, gatewayNameUnescaped, vpnConnectionNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStopPacketCapture = &respr
		v.beginStopPacketCapture.add(req, beginStopPacketCapture)
	}

	resp, err := server.PollerResponderNext(beginStopPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginStopPacketCapture.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStopPacketCapture) {
		v.beginStopPacketCapture.remove(req)
	}

	return resp, nil
}
