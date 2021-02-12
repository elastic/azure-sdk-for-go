package batch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PoolClient is the client for the Pool methods of the Batch service.
type PoolClient struct {
	BaseClient
}

// NewPoolClient creates an instance of the PoolClient client.
func NewPoolClient(subscriptionID string) PoolClient {
	return NewPoolClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPoolClientWithBaseURI creates an instance of the PoolClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPoolClientWithBaseURI(baseURI string, subscriptionID string) PoolClient {
	return PoolClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new pool inside the specified account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// poolName - the pool name. This must be unique within the account.
// parameters - additional parameters for pool creation.
// ifMatch - the entity state (ETag) version of the pool to update. A value of "*" can be used to apply the
// operation only if the pool already exists. If omitted, this operation will always be applied.
// ifNoneMatch - set to '*' to allow a new pool to be created, but to prevent updating an existing pool. Other
// values will be ignored.
func (client PoolClient) Create(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string, ifNoneMatch string) (result PoolCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PoolProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.CloudServiceConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.CloudServiceConfiguration.OsFamily", Name: validation.Null, Rule: true, Chain: nil}}},
						{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration.ImageReference", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration.NodeAgentSkuID", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration.ContainerConfiguration", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.PoolProperties.DeploymentConfiguration.VirtualMachineConfiguration.ContainerConfiguration.Type", Name: validation.Null, Rule: true, Chain: nil}}},
							}},
					}},
					{Target: "parameters.PoolProperties.ScaleSettings", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.ScaleSettings.AutoScale", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.ScaleSettings.AutoScale.Formula", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "parameters.PoolProperties.AutoScaleRun", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.AutoScaleRun.EvaluationTime", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.PoolProperties.AutoScaleRun.Error", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.PoolProperties.AutoScaleRun.Error.Code", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.PoolProperties.AutoScaleRun.Error.Message", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
					{Target: "parameters.PoolProperties.NetworkConfiguration", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.NetworkConfiguration.EndpointConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.NetworkConfiguration.EndpointConfiguration.InboundNatPools", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "parameters.PoolProperties.StartTask", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PoolProperties.StartTask.ContainerSettings", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.PoolProperties.StartTask.ContainerSettings.ImageName", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.PoolProperties.StartTask.ContainerSettings.Registry", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.PoolProperties.StartTask.ContainerSettings.Registry.UserName", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.PoolProperties.StartTask.ContainerSettings.Registry.Password", Name: validation.Null, Rule: true, Chain: nil},
									}},
							}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, poolName, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PoolClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) CreateSender(req *http.Request) (future PoolCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PoolClient) (p Pool, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.PoolCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("batch.PoolCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "batch.PoolCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.CreateResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "batch.PoolCreateFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PoolClient) CreateResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// poolName - the pool name. This must be unique within the account.
func (client PoolClient) Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result PoolDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PoolClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) DeleteSender(req *http.Request) (future PoolDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PoolClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.PoolDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("batch.PoolDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PoolClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableAutoScale disables automatic scaling for a pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// poolName - the pool name. This must be unique within the account.
func (client PoolClient) DisableAutoScale(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result Pool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.DisableAutoScale")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "DisableAutoScale", err.Error())
	}

	req, err := client.DisableAutoScalePreparer(ctx, resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "DisableAutoScale", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableAutoScaleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "DisableAutoScale", resp, "Failure sending request")
		return
	}

	result, err = client.DisableAutoScaleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "DisableAutoScale", resp, "Failure responding to request")
		return
	}

	return
}

// DisableAutoScalePreparer prepares the DisableAutoScale request.
func (client PoolClient) DisableAutoScalePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}/disableAutoScale", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableAutoScaleSender sends the DisableAutoScale request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) DisableAutoScaleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DisableAutoScaleResponder handles the response to the DisableAutoScale request. The method always
// closes the http.Response Body.
func (client PoolClient) DisableAutoScaleResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets information about the specified pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// poolName - the pool name. This must be unique within the account.
func (client PoolClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result Pool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PoolClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PoolClient) GetResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBatchAccount lists all of the pools in the specified account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// maxresults - the maximum number of items to return in the response.
// selectParameter - comma separated list of properties that should be returned. e.g.
// "properties/provisioningState". Only top level properties under properties/ are valid for selection.
// filter - oData filter expression. Valid properties for filtering are:
//
// name
// properties/allocationState
// properties/allocationStateTransitionTime
// properties/creationTime
// properties/provisioningState
// properties/provisioningStateTransitionTime
// properties/lastModified
// properties/vmSize
// properties/interNodeCommunication
// properties/scaleSettings/autoScale
// properties/scaleSettings/fixedScale
func (client PoolClient) ListByBatchAccount(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result ListPoolsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.lpr.Response.Response != nil {
				sc = result.lpr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "ListByBatchAccount", err.Error())
	}

	result.fn = client.listByBatchAccountNextResults
	req, err := client.ListByBatchAccountPreparer(ctx, resourceGroupName, accountName, maxresults, selectParameter, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.lpr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", resp, "Failure sending request")
		return
	}

	result.lpr, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "ListByBatchAccount", resp, "Failure responding to request")
		return
	}
	if result.lpr.hasNextLink() && result.lpr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBatchAccountPreparer prepares the ListByBatchAccount request.
func (client PoolClient) ListByBatchAccountPreparer(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBatchAccountSender sends the ListByBatchAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) ListByBatchAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByBatchAccountResponder handles the response to the ListByBatchAccount request. The method always
// closes the http.Response Body.
func (client PoolClient) ListByBatchAccountResponder(resp *http.Response) (result ListPoolsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBatchAccountNextResults retrieves the next set of results, if any.
func (client PoolClient) listByBatchAccountNextResults(ctx context.Context, lastResults ListPoolsResult) (result ListPoolsResult, err error) {
	req, err := lastResults.listPoolsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.PoolClient", "listByBatchAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.PoolClient", "listByBatchAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "listByBatchAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBatchAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client PoolClient) ListByBatchAccountComplete(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32, selectParameter string, filter string) (result ListPoolsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBatchAccount(ctx, resourceGroupName, accountName, maxresults, selectParameter, filter)
	return
}

// StopResize this does not restore the pool to its previous state before the resize operation: it only stops any
// further changes being made, and the pool maintains its current state. After stopping, the pool stabilizes at the
// number of nodes it was at when the stop operation was done. During the stop operation, the pool allocation state
// changes first to stopping and then to steady. A resize operation need not be an explicit resize pool request; this
// API can also be used to halt the initial sizing of the pool when it is created.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// poolName - the pool name. This must be unique within the account.
func (client PoolClient) StopResize(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result Pool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.StopResize")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "StopResize", err.Error())
	}

	req, err := client.StopResizePreparer(ctx, resourceGroupName, accountName, poolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "StopResize", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopResizeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "StopResize", resp, "Failure sending request")
		return
	}

	result, err = client.StopResizeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "StopResize", resp, "Failure responding to request")
		return
	}

	return
}

// StopResizePreparer prepares the StopResize request.
func (client PoolClient) StopResizePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}/stopResize", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopResizeSender sends the StopResize request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) StopResizeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// StopResizeResponder handles the response to the StopResize request. The method always
// closes the http.Response Body.
func (client PoolClient) StopResizeResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the properties of an existing pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// poolName - the pool name. This must be unique within the account.
// parameters - pool properties that should be updated. Properties that are supplied will be updated, any
// property not supplied will be unchanged.
// ifMatch - the entity state (ETag) version of the pool to update. This value can be omitted or set to "*" to
// apply the operation unconditionally.
func (client PoolClient) Update(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string) (result Pool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PoolClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: poolName,
			Constraints: []validation.Constraint{{Target: "poolName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "poolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "poolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PoolClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, poolName, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PoolClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client PoolClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, parameters Pool, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client PoolClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PoolClient) UpdateResponder(resp *http.Response) (result Pool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
