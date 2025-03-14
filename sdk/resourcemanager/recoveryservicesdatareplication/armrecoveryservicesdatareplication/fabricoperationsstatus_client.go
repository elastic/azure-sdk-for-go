//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

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

// FabricOperationsStatusClient contains the methods for the FabricOperationsStatus group.
// Don't use this type directly, use NewFabricOperationsStatusClient() instead.
type FabricOperationsStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFabricOperationsStatusClient creates a new instance of FabricOperationsStatusClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFabricOperationsStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FabricOperationsStatusClient, error) {
	cl, err := arm.NewClient(moduleName+".FabricOperationsStatusClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FabricOperationsStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Tracks the results of an asynchronous operation on the fabric.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fabricName - The fabric name.
//   - operationID - The ID of an ongoing async operation.
//   - options - FabricOperationsStatusClientGetOptions contains the optional parameters for the FabricOperationsStatusClient.Get
//     method.
func (client *FabricOperationsStatusClient) Get(ctx context.Context, resourceGroupName string, fabricName string, operationID string, options *FabricOperationsStatusClientGetOptions) (FabricOperationsStatusClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, fabricName, operationID, options)
	if err != nil {
		return FabricOperationsStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FabricOperationsStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FabricOperationsStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FabricOperationsStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, fabricName string, operationID string, options *FabricOperationsStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationFabrics/{fabricName}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FabricOperationsStatusClient) getHandleResponse(resp *http.Response) (FabricOperationsStatusClientGetResponse, error) {
	result := FabricOperationsStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return FabricOperationsStatusClientGetResponse{}, err
	}
	return result, nil
}
