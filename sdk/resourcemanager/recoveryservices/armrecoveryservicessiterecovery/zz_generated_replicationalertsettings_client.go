//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationAlertSettingsClient contains the methods for the ReplicationAlertSettings group.
// Don't use this type directly, use NewReplicationAlertSettingsClient() instead.
type ReplicationAlertSettingsClient struct {
	ep                string
	pl                runtime.Pipeline
	resourceName      string
	resourceGroupName string
	subscriptionID    string
}

// NewReplicationAlertSettingsClient creates a new instance of ReplicationAlertSettingsClient with the specified values.
func NewReplicationAlertSettingsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationAlertSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReplicationAlertSettingsClient{resourceName: resourceName, resourceGroupName: resourceGroupName, subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Create or update an email notification(alert) configuration.
// If the operation fails it returns a generic error.
func (client *ReplicationAlertSettingsClient) Create(ctx context.Context, alertSettingName string, request ConfigureAlertRequest, options *ReplicationAlertSettingsCreateOptions) (ReplicationAlertSettingsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, alertSettingName, request, options)
	if err != nil {
		return ReplicationAlertSettingsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationAlertSettingsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationAlertSettingsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ReplicationAlertSettingsClient) createCreateRequest(ctx context.Context, alertSettingName string, request ConfigureAlertRequest, options *ReplicationAlertSettingsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertSettingName == "" {
		return nil, errors.New("parameter alertSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertSettingName}", url.PathEscape(alertSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// createHandleResponse handles the Create response.
func (client *ReplicationAlertSettingsClient) createHandleResponse(resp *http.Response) (ReplicationAlertSettingsCreateResponse, error) {
	result := ReplicationAlertSettingsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return ReplicationAlertSettingsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *ReplicationAlertSettingsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the details of the specified email notification(alert) configuration.
// If the operation fails it returns a generic error.
func (client *ReplicationAlertSettingsClient) Get(ctx context.Context, alertSettingName string, options *ReplicationAlertSettingsGetOptions) (ReplicationAlertSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, alertSettingName, options)
	if err != nil {
		return ReplicationAlertSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationAlertSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationAlertSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationAlertSettingsClient) getCreateRequest(ctx context.Context, alertSettingName string, options *ReplicationAlertSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertSettingName == "" {
		return nil, errors.New("parameter alertSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertSettingName}", url.PathEscape(alertSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationAlertSettingsClient) getHandleResponse(resp *http.Response) (ReplicationAlertSettingsGetResponse, error) {
	result := ReplicationAlertSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return ReplicationAlertSettingsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ReplicationAlertSettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// List - Gets the list of email notification(alert) configurations for the vault.
// If the operation fails it returns a generic error.
func (client *ReplicationAlertSettingsClient) List(options *ReplicationAlertSettingsListOptions) *ReplicationAlertSettingsListPager {
	return &ReplicationAlertSettingsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ReplicationAlertSettingsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AlertCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ReplicationAlertSettingsClient) listCreateRequest(ctx context.Context, options *ReplicationAlertSettingsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationAlertSettingsClient) listHandleResponse(resp *http.Response) (ReplicationAlertSettingsListResponse, error) {
	result := ReplicationAlertSettingsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertCollection); err != nil {
		return ReplicationAlertSettingsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReplicationAlertSettingsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}