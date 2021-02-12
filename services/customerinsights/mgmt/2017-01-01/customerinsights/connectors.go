package customerinsights

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

// ConnectorsClient is the the Azure Customer Insights management API provides a RESTful set of web services that
// interact with Azure Customer Insights service to manage your resources. The API has entities that capture the
// relationship between an end user and the Azure Customer Insights service.
type ConnectorsClient struct {
	BaseClient
}

// NewConnectorsClient creates an instance of the ConnectorsClient client.
func NewConnectorsClient(subscriptionID string) ConnectorsClient {
	return NewConnectorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConnectorsClientWithBaseURI creates an instance of the ConnectorsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewConnectorsClientWithBaseURI(baseURI string, subscriptionID string) ConnectorsClient {
	return ConnectorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a connector or updates an existing connector in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// connectorName - the name of the connector.
// parameters - parameters supplied to the CreateOrUpdate Connector operation.
func (client ConnectorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, connectorName string, parameters ConnectorResourceFormat) (result ConnectorsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: connectorName,
			Constraints: []validation.Constraint{{Target: "connectorName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "connectorName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "connectorName", Name: validation.Pattern, Rule: `^[a-zA-Z][a-zA-Z0-9_]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Connector", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Connector.ConnectorProperties", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("customerinsights.ConnectorsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, hubName, connectorName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ConnectorsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, hubName string, connectorName string, parameters ConnectorResourceFormat) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorsClient) CreateOrUpdateSender(req *http.Request) (future ConnectorsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConnectorsClient) (crf ConnectorResourceFormat, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("customerinsights.ConnectorsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		crf.Response.Response, err = future.GetResult(sender)
		if crf.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && crf.Response.Response.StatusCode != http.StatusNoContent {
			crf, err = client.CreateOrUpdateResponder(crf.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsCreateOrUpdateFuture", "Result", crf.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ConnectorsClient) CreateOrUpdateResponder(resp *http.Response) (result ConnectorResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a connector in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// connectorName - the name of the connector.
func (client ConnectorsClient) Delete(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (result ConnectorsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, hubName, connectorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ConnectorsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorsClient) DeleteSender(req *http.Request) (future ConnectorsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ConnectorsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("customerinsights.ConnectorsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ConnectorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a connector in the hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
// connectorName - the name of the connector.
func (client ConnectorsClient) Get(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (result ConnectorResourceFormat, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hubName, connectorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ConnectorsClient) GetPreparer(ctx context.Context, resourceGroupName string, hubName string, connectorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectorName":     autorest.Encode("path", connectorName),
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ConnectorsClient) GetResponder(resp *http.Response) (result ConnectorResourceFormat, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByHub gets all the connectors in the specified hub.
// Parameters:
// resourceGroupName - the name of the resource group.
// hubName - the name of the hub.
func (client ConnectorsClient) ListByHub(ctx context.Context, resourceGroupName string, hubName string) (result ConnectorListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorsClient.ListByHub")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByHubNextResults
	req, err := client.ListByHubPreparer(ctx, resourceGroupName, hubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "ListByHub", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "ListByHub", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "ListByHub", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByHubPreparer prepares the ListByHub request.
func (client ConnectorsClient) ListByHubPreparer(ctx context.Context, resourceGroupName string, hubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hubName":           autorest.Encode("path", hubName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHubSender sends the ListByHub request. The method will close the
// http.Response Body if it receives an error.
func (client ConnectorsClient) ListByHubSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHubResponder handles the response to the ListByHub request. The method always
// closes the http.Response Body.
func (client ConnectorsClient) ListByHubResponder(resp *http.Response) (result ConnectorListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHubNextResults retrieves the next set of results, if any.
func (client ConnectorsClient) listByHubNextResults(ctx context.Context, lastResults ConnectorListResult) (result ConnectorListResult, err error) {
	req, err := lastResults.connectorListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "listByHubNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHubSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "listByHubNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHubResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customerinsights.ConnectorsClient", "listByHubNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByHubComplete enumerates all values, automatically crossing page boundaries as required.
func (client ConnectorsClient) ListByHubComplete(ctx context.Context, resourceGroupName string, hubName string) (result ConnectorListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectorsClient.ListByHub")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHub(ctx, resourceGroupName, hubName)
	return
}
