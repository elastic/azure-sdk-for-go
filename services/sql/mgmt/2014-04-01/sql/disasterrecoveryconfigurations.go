package sql

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DisasterRecoveryConfigurationsClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type DisasterRecoveryConfigurationsClient struct {
	BaseClient
}

// NewDisasterRecoveryConfigurationsClient creates an instance of the DisasterRecoveryConfigurationsClient client.
func NewDisasterRecoveryConfigurationsClient(subscriptionID string) DisasterRecoveryConfigurationsClient {
	return NewDisasterRecoveryConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDisasterRecoveryConfigurationsClientWithBaseURI creates an instance of the DisasterRecoveryConfigurationsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewDisasterRecoveryConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) DisasterRecoveryConfigurationsClient {
	return DisasterRecoveryConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a disaster recovery configuration.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// disasterRecoveryConfigurationName - the name of the disaster recovery configuration to be created/updated.
func (client DisasterRecoveryConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result DisasterRecoveryConfigurationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisasterRecoveryConfigurationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, disasterRecoveryConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DisasterRecoveryConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"disasterRecoveryConfigurationName": autorest.Encode("path", disasterRecoveryConfigurationName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigurationsClient) CreateOrUpdateSender(req *http.Request) (future DisasterRecoveryConfigurationsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DisasterRecoveryConfigurationsClient) (drc DisasterRecoveryConfiguration, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("sql.DisasterRecoveryConfigurationsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		drc.Response.Response, err = future.GetResult(sender)
		if drc.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && drc.Response.Response.StatusCode != http.StatusNoContent {
			drc, err = client.CreateOrUpdateResponder(drc.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsCreateOrUpdateFuture", "Result", drc.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result DisasterRecoveryConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a disaster recovery configuration.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// disasterRecoveryConfigurationName - the name of the disaster recovery configuration to be deleted.
func (client DisasterRecoveryConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result DisasterRecoveryConfigurationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisasterRecoveryConfigurationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, disasterRecoveryConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DisasterRecoveryConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"disasterRecoveryConfigurationName": autorest.Encode("path", disasterRecoveryConfigurationName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigurationsClient) DeleteSender(req *http.Request) (future DisasterRecoveryConfigurationsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DisasterRecoveryConfigurationsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("sql.DisasterRecoveryConfigurationsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Failover fails over from the current primary server to this server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// disasterRecoveryConfigurationName - the name of the disaster recovery configuration to failover.
func (client DisasterRecoveryConfigurationsClient) Failover(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result DisasterRecoveryConfigurationsFailoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisasterRecoveryConfigurationsClient.Failover")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.FailoverPreparer(ctx, resourceGroupName, serverName, disasterRecoveryConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Failover", nil, "Failure preparing request")
		return
	}

	result, err = client.FailoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Failover", nil, "Failure sending request")
		return
	}

	return
}

// FailoverPreparer prepares the Failover request.
func (client DisasterRecoveryConfigurationsClient) FailoverPreparer(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"disasterRecoveryConfigurationName": autorest.Encode("path", disasterRecoveryConfigurationName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}/failover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FailoverSender sends the Failover request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigurationsClient) FailoverSender(req *http.Request) (future DisasterRecoveryConfigurationsFailoverFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DisasterRecoveryConfigurationsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsFailoverFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("sql.DisasterRecoveryConfigurationsFailoverFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// FailoverResponder handles the response to the Failover request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigurationsClient) FailoverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// FailoverAllowDataLoss fails over from the current primary server to this server. This operation might result in data
// loss.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// disasterRecoveryConfigurationName - the name of the disaster recovery configuration to failover forcefully.
func (client DisasterRecoveryConfigurationsClient) FailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result DisasterRecoveryConfigurationsFailoverAllowDataLossFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisasterRecoveryConfigurationsClient.FailoverAllowDataLoss")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.FailoverAllowDataLossPreparer(ctx, resourceGroupName, serverName, disasterRecoveryConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "FailoverAllowDataLoss", nil, "Failure preparing request")
		return
	}

	result, err = client.FailoverAllowDataLossSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "FailoverAllowDataLoss", nil, "Failure sending request")
		return
	}

	return
}

// FailoverAllowDataLossPreparer prepares the FailoverAllowDataLoss request.
func (client DisasterRecoveryConfigurationsClient) FailoverAllowDataLossPreparer(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"disasterRecoveryConfigurationName": autorest.Encode("path", disasterRecoveryConfigurationName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}/forceFailoverAllowDataLoss", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FailoverAllowDataLossSender sends the FailoverAllowDataLoss request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigurationsClient) FailoverAllowDataLossSender(req *http.Request) (future DisasterRecoveryConfigurationsFailoverAllowDataLossFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DisasterRecoveryConfigurationsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsFailoverAllowDataLossFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("sql.DisasterRecoveryConfigurationsFailoverAllowDataLossFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// FailoverAllowDataLossResponder handles the response to the FailoverAllowDataLoss request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigurationsClient) FailoverAllowDataLossResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a disaster recovery configuration.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// disasterRecoveryConfigurationName - the name of the disaster recovery configuration.
func (client DisasterRecoveryConfigurationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result DisasterRecoveryConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisasterRecoveryConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, disasterRecoveryConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DisasterRecoveryConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"disasterRecoveryConfigurationName": autorest.Encode("path", disasterRecoveryConfigurationName),
		"resourceGroupName":                 autorest.Encode("path", resourceGroupName),
		"serverName":                        autorest.Encode("path", serverName),
		"subscriptionId":                    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigurationsClient) GetResponder(resp *http.Response) (result DisasterRecoveryConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists a server's disaster recovery configuration.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client DisasterRecoveryConfigurationsClient) List(ctx context.Context, resourceGroupName string, serverName string) (result DisasterRecoveryConfigurationListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisasterRecoveryConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DisasterRecoveryConfigurationsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DisasterRecoveryConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DisasterRecoveryConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DisasterRecoveryConfigurationsClient) ListResponder(resp *http.Response) (result DisasterRecoveryConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
