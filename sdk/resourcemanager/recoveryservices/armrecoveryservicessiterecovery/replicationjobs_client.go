//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ReplicationJobsClient contains the methods for the ReplicationJobs group.
// Don't use this type directly, use NewReplicationJobsClient() instead.
type ReplicationJobsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationJobsClient creates a new instance of ReplicationJobsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationJobsClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicationJobsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationJobsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - The operation to cancel an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - jobName - Job identifier.
//   - options - ReplicationJobsClientBeginCancelOptions contains the optional parameters for the ReplicationJobsClient.BeginCancel
//     method.
func (client *ReplicationJobsClient) BeginCancel(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientBeginCancelOptions) (*runtime.Poller[ReplicationJobsClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceName, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ReplicationJobsClientCancelResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationJobsClientCancelResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Cancel - The operation to cancel an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ReplicationJobsClient) cancel(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientBeginCancelOptions) (*http.Response, error) {
	var err error
	req, err := client.cancelCreateRequest(ctx, resourceName, resourceGroupName, jobName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ReplicationJobsClient) cancelCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}/cancel"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExport - The operation to export the details of the Azure Site Recovery jobs of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - jobQueryParameter - job query filter.
//   - options - ReplicationJobsClientBeginExportOptions contains the optional parameters for the ReplicationJobsClient.BeginExport
//     method.
func (client *ReplicationJobsClient) BeginExport(ctx context.Context, resourceName string, resourceGroupName string, jobQueryParameter JobQueryParameter, options *ReplicationJobsClientBeginExportOptions) (*runtime.Poller[ReplicationJobsClientExportResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.export(ctx, resourceName, resourceGroupName, jobQueryParameter, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ReplicationJobsClientExportResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationJobsClientExportResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Export - The operation to export the details of the Azure Site Recovery jobs of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ReplicationJobsClient) export(ctx context.Context, resourceName string, resourceGroupName string, jobQueryParameter JobQueryParameter, options *ReplicationJobsClientBeginExportOptions) (*http.Response, error) {
	var err error
	req, err := client.exportCreateRequest(ctx, resourceName, resourceGroupName, jobQueryParameter, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// exportCreateRequest creates the Export request.
func (client *ReplicationJobsClient) exportCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, jobQueryParameter JobQueryParameter, options *ReplicationJobsClientBeginExportOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/export"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, jobQueryParameter); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get the details of an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - jobName - Job identifier.
//   - options - ReplicationJobsClientGetOptions contains the optional parameters for the ReplicationJobsClient.Get method.
func (client *ReplicationJobsClient) Get(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientGetOptions) (ReplicationJobsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, jobName, options)
	if err != nil {
		return ReplicationJobsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationJobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationJobsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationJobsClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
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
func (client *ReplicationJobsClient) getHandleResponse(resp *http.Response) (ReplicationJobsClientGetResponse, error) {
	result := ReplicationJobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return ReplicationJobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of Azure Site Recovery Jobs for the vault.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationJobsClientListOptions contains the optional parameters for the ReplicationJobsClient.NewListPager
//     method.
func (client *ReplicationJobsClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationJobsClientListOptions) *runtime.Pager[ReplicationJobsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationJobsClientListResponse]{
		More: func(page ReplicationJobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationJobsClientListResponse) (ReplicationJobsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationJobsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationJobsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationJobsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationJobsClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationJobsClient) listHandleResponse(resp *http.Response) (ReplicationJobsClientListResponse, error) {
	result := ReplicationJobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollection); err != nil {
		return ReplicationJobsClientListResponse{}, err
	}
	return result, nil
}

// BeginRestart - The operation to restart an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - jobName - Job identifier.
//   - options - ReplicationJobsClientBeginRestartOptions contains the optional parameters for the ReplicationJobsClient.BeginRestart
//     method.
func (client *ReplicationJobsClient) BeginRestart(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientBeginRestartOptions) (*runtime.Poller[ReplicationJobsClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceName, resourceGroupName, jobName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ReplicationJobsClientRestartResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationJobsClientRestartResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Restart - The operation to restart an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ReplicationJobsClient) restart(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientBeginRestartOptions) (*http.Response, error) {
	var err error
	req, err := client.restartCreateRequest(ctx, resourceName, resourceGroupName, jobName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ReplicationJobsClient) restartCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, jobName string, options *ReplicationJobsClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}/restart"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginResume - The operation to resume an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - jobName - Job identifier.
//   - resumeJobParams - Resume rob comments.
//   - options - ReplicationJobsClientBeginResumeOptions contains the optional parameters for the ReplicationJobsClient.BeginResume
//     method.
func (client *ReplicationJobsClient) BeginResume(ctx context.Context, resourceName string, resourceGroupName string, jobName string, resumeJobParams ResumeJobParams, options *ReplicationJobsClientBeginResumeOptions) (*runtime.Poller[ReplicationJobsClientResumeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resume(ctx, resourceName, resourceGroupName, jobName, resumeJobParams, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ReplicationJobsClientResumeResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationJobsClientResumeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Resume - The operation to resume an Azure Site Recovery job.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ReplicationJobsClient) resume(ctx context.Context, resourceName string, resourceGroupName string, jobName string, resumeJobParams ResumeJobParams, options *ReplicationJobsClientBeginResumeOptions) (*http.Response, error) {
	var err error
	req, err := client.resumeCreateRequest(ctx, resourceName, resourceGroupName, jobName, resumeJobParams, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *ReplicationJobsClient) resumeCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, jobName string, resumeJobParams ResumeJobParams, options *ReplicationJobsClientBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}/resume"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resumeJobParams); err != nil {
		return nil, err
	}
	return req, nil
}
