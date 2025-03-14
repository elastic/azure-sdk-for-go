//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

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

// InternetGatewayRulesClient contains the methods for the InternetGatewayRules group.
// Don't use this type directly, use NewInternetGatewayRulesClient() instead.
type InternetGatewayRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInternetGatewayRulesClient creates a new instance of InternetGatewayRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInternetGatewayRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InternetGatewayRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".InternetGatewayRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InternetGatewayRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates an Internet Gateway rule resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayRuleName - Name of the Internet Gateway rule.
//   - body - Request payload.
//   - options - InternetGatewayRulesClientBeginCreateOptions contains the optional parameters for the InternetGatewayRulesClient.BeginCreate
//     method.
func (client *InternetGatewayRulesClient) BeginCreate(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, body InternetGatewayRule, options *InternetGatewayRulesClientBeginCreateOptions) (*runtime.Poller[InternetGatewayRulesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, internetGatewayRuleName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InternetGatewayRulesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InternetGatewayRulesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates an Internet Gateway rule resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *InternetGatewayRulesClient) create(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, body InternetGatewayRule, options *InternetGatewayRulesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, internetGatewayRuleName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *InternetGatewayRulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, body InternetGatewayRule, options *InternetGatewayRulesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/{internetGatewayRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayRuleName == "" {
		return nil, errors.New("parameter internetGatewayRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayRuleName}", url.PathEscape(internetGatewayRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements Internet Gateway Rules DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayRuleName - Name of the Internet Gateway rule.
//   - options - InternetGatewayRulesClientBeginDeleteOptions contains the optional parameters for the InternetGatewayRulesClient.BeginDelete
//     method.
func (client *InternetGatewayRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, options *InternetGatewayRulesClientBeginDeleteOptions) (*runtime.Poller[InternetGatewayRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, internetGatewayRuleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InternetGatewayRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InternetGatewayRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Implements Internet Gateway Rules DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *InternetGatewayRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, options *InternetGatewayRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, internetGatewayRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InternetGatewayRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, options *InternetGatewayRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/{internetGatewayRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayRuleName == "" {
		return nil, errors.New("parameter internetGatewayRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayRuleName}", url.PathEscape(internetGatewayRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an Internet Gateway Rule resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayRuleName - Name of the Internet Gateway rule.
//   - options - InternetGatewayRulesClientGetOptions contains the optional parameters for the InternetGatewayRulesClient.Get
//     method.
func (client *InternetGatewayRulesClient) Get(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, options *InternetGatewayRulesClientGetOptions) (InternetGatewayRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, internetGatewayRuleName, options)
	if err != nil {
		return InternetGatewayRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InternetGatewayRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InternetGatewayRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InternetGatewayRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, options *InternetGatewayRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/{internetGatewayRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayRuleName == "" {
		return nil, errors.New("parameter internetGatewayRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayRuleName}", url.PathEscape(internetGatewayRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InternetGatewayRulesClient) getHandleResponse(resp *http.Response) (InternetGatewayRulesClientGetResponse, error) {
	result := InternetGatewayRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternetGatewayRule); err != nil {
		return InternetGatewayRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Implements Internet Gateway Rules list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - InternetGatewayRulesClientListByResourceGroupOptions contains the optional parameters for the InternetGatewayRulesClient.NewListByResourceGroupPager
//     method.
func (client *InternetGatewayRulesClient) NewListByResourceGroupPager(resourceGroupName string, options *InternetGatewayRulesClientListByResourceGroupOptions) *runtime.Pager[InternetGatewayRulesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[InternetGatewayRulesClientListByResourceGroupResponse]{
		More: func(page InternetGatewayRulesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InternetGatewayRulesClientListByResourceGroupResponse) (InternetGatewayRulesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InternetGatewayRulesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InternetGatewayRulesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InternetGatewayRulesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *InternetGatewayRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *InternetGatewayRulesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *InternetGatewayRulesClient) listByResourceGroupHandleResponse(resp *http.Response) (InternetGatewayRulesClientListByResourceGroupResponse, error) {
	result := InternetGatewayRulesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternetGatewayRulesListResult); err != nil {
		return InternetGatewayRulesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all Internet Gateway rules in the given subscription.
//
// Generated from API version 2023-06-15
//   - options - InternetGatewayRulesClientListBySubscriptionOptions contains the optional parameters for the InternetGatewayRulesClient.NewListBySubscriptionPager
//     method.
func (client *InternetGatewayRulesClient) NewListBySubscriptionPager(options *InternetGatewayRulesClientListBySubscriptionOptions) *runtime.Pager[InternetGatewayRulesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[InternetGatewayRulesClientListBySubscriptionResponse]{
		More: func(page InternetGatewayRulesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InternetGatewayRulesClientListBySubscriptionResponse) (InternetGatewayRulesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InternetGatewayRulesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InternetGatewayRulesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InternetGatewayRulesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *InternetGatewayRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *InternetGatewayRulesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *InternetGatewayRulesClient) listBySubscriptionHandleResponse(resp *http.Response) (InternetGatewayRulesClientListBySubscriptionResponse, error) {
	result := InternetGatewayRulesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternetGatewayRulesListResult); err != nil {
		return InternetGatewayRulesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the Internet Gateway Rule resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayRuleName - Name of the Internet Gateway rule.
//   - body - Internet Gateway Rule properties to update.
//   - options - InternetGatewayRulesClientBeginUpdateOptions contains the optional parameters for the InternetGatewayRulesClient.BeginUpdate
//     method.
func (client *InternetGatewayRulesClient) BeginUpdate(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, body InternetGatewayRulePatch, options *InternetGatewayRulesClientBeginUpdateOptions) (*runtime.Poller[InternetGatewayRulesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, internetGatewayRuleName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InternetGatewayRulesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InternetGatewayRulesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - API to update certain properties of the Internet Gateway Rule resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *InternetGatewayRulesClient) update(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, body InternetGatewayRulePatch, options *InternetGatewayRulesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, internetGatewayRuleName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *InternetGatewayRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayRuleName string, body InternetGatewayRulePatch, options *InternetGatewayRulesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/{internetGatewayRuleName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayRuleName == "" {
		return nil, errors.New("parameter internetGatewayRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayRuleName}", url.PathEscape(internetGatewayRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
