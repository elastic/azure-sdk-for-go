//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticsPackagesClient contains the methods for the DiagnosticsPackages group.
// Don't use this type directly, use NewDiagnosticsPackagesClient() instead.
type DiagnosticsPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDiagnosticsPackagesClient creates a new instance of DiagnosticsPackagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDiagnosticsPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticsPackagesClient, error) {
	cl, err := arm.NewClient(moduleName+".DiagnosticsPackagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DiagnosticsPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a diagnostics package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - diagnosticsPackageName - The name of the diagnostics package.
//   - options - DiagnosticsPackagesClientBeginCreateOrUpdateOptions contains the optional parameters for the DiagnosticsPackagesClient.BeginCreateOrUpdate
//     method.
func (client *DiagnosticsPackagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DiagnosticsPackagesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, packetCoreControlPlaneName, diagnosticsPackageName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiagnosticsPackagesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DiagnosticsPackagesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a diagnostics package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *DiagnosticsPackagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, diagnosticsPackageName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticsPackagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/diagnosticsPackages/{diagnosticsPackageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if diagnosticsPackageName == "" {
		return nil, errors.New("parameter diagnosticsPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticsPackageName}", url.PathEscape(diagnosticsPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Deletes the specified diagnostics package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - diagnosticsPackageName - The name of the diagnostics package.
//   - options - DiagnosticsPackagesClientBeginDeleteOptions contains the optional parameters for the DiagnosticsPackagesClient.BeginDelete
//     method.
func (client *DiagnosticsPackagesClient) BeginDelete(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientBeginDeleteOptions) (*runtime.Poller[DiagnosticsPackagesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, packetCoreControlPlaneName, diagnosticsPackageName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiagnosticsPackagesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DiagnosticsPackagesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified diagnostics package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *DiagnosticsPackagesClient) deleteOperation(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, diagnosticsPackageName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticsPackagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/diagnosticsPackages/{diagnosticsPackageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if diagnosticsPackageName == "" {
		return nil, errors.New("parameter diagnosticsPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticsPackageName}", url.PathEscape(diagnosticsPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified diagnostics package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - diagnosticsPackageName - The name of the diagnostics package.
//   - options - DiagnosticsPackagesClientGetOptions contains the optional parameters for the DiagnosticsPackagesClient.Get method.
func (client *DiagnosticsPackagesClient) Get(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientGetOptions) (DiagnosticsPackagesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, diagnosticsPackageName, options)
	if err != nil {
		return DiagnosticsPackagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticsPackagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DiagnosticsPackagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DiagnosticsPackagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, diagnosticsPackageName string, options *DiagnosticsPackagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/diagnosticsPackages/{diagnosticsPackageName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	if diagnosticsPackageName == "" {
		return nil, errors.New("parameter diagnosticsPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diagnosticsPackageName}", url.PathEscape(diagnosticsPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticsPackagesClient) getHandleResponse(resp *http.Response) (DiagnosticsPackagesClientGetResponse, error) {
	result := DiagnosticsPackagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticsPackage); err != nil {
		return DiagnosticsPackagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPacketCoreControlPlanePager - Lists all the diagnostics packages under a packet core control plane.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - packetCoreControlPlaneName - The name of the packet core control plane.
//   - options - DiagnosticsPackagesClientListByPacketCoreControlPlaneOptions contains the optional parameters for the DiagnosticsPackagesClient.NewListByPacketCoreControlPlanePager
//     method.
func (client *DiagnosticsPackagesClient) NewListByPacketCoreControlPlanePager(resourceGroupName string, packetCoreControlPlaneName string, options *DiagnosticsPackagesClientListByPacketCoreControlPlaneOptions) *runtime.Pager[DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse] {
	return runtime.NewPager(runtime.PagingHandler[DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse]{
		More: func(page DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse) (DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPacketCoreControlPlaneCreateRequest(ctx, resourceGroupName, packetCoreControlPlaneName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPacketCoreControlPlaneHandleResponse(resp)
		},
	})
}

// listByPacketCoreControlPlaneCreateRequest creates the ListByPacketCoreControlPlane request.
func (client *DiagnosticsPackagesClient) listByPacketCoreControlPlaneCreateRequest(ctx context.Context, resourceGroupName string, packetCoreControlPlaneName string, options *DiagnosticsPackagesClientListByPacketCoreControlPlaneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/{packetCoreControlPlaneName}/diagnosticsPackages"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if packetCoreControlPlaneName == "" {
		return nil, errors.New("parameter packetCoreControlPlaneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCoreControlPlaneName}", url.PathEscape(packetCoreControlPlaneName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPacketCoreControlPlaneHandleResponse handles the ListByPacketCoreControlPlane response.
func (client *DiagnosticsPackagesClient) listByPacketCoreControlPlaneHandleResponse(resp *http.Response) (DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse, error) {
	result := DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticsPackageListResult); err != nil {
		return DiagnosticsPackagesClientListByPacketCoreControlPlaneResponse{}, err
	}
	return result, nil
}
