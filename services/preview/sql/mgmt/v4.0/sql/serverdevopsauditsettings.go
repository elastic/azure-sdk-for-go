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

// ServerDevOpsAuditSettingsClient is the the Azure SQL Database management API provides a RESTful set of web services
// that interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type ServerDevOpsAuditSettingsClient struct {
	BaseClient
}

// NewServerDevOpsAuditSettingsClient creates an instance of the ServerDevOpsAuditSettingsClient client.
func NewServerDevOpsAuditSettingsClient(subscriptionID string) ServerDevOpsAuditSettingsClient {
	return NewServerDevOpsAuditSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerDevOpsAuditSettingsClientWithBaseURI creates an instance of the ServerDevOpsAuditSettingsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewServerDevOpsAuditSettingsClientWithBaseURI(baseURI string, subscriptionID string) ServerDevOpsAuditSettingsClient {
	return ServerDevOpsAuditSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a server's DevOps audit settings.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// devOpsAuditingSettingsName - the name of the devops audit settings. This should always be 'default'.
// parameters - properties of DevOps audit settings
func (client ServerDevOpsAuditSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings) (result ServerDevOpsAuditSettingsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerDevOpsAuditSettingsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServerDevOpsAuditSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"devOpsAuditingSettingsName": autorest.Encode("path", devOpsAuditingSettingsName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDevOpsAuditSettingsClient) CreateOrUpdateSender(req *http.Request) (future ServerDevOpsAuditSettingsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ServerDevOpsAuditSettingsClient) (sdoas ServerDevOpsAuditingSettings, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("sql.ServerDevOpsAuditSettingsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		sdoas.Response.Response, err = future.GetResult(sender)
		if sdoas.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && sdoas.Response.Response.StatusCode != http.StatusNoContent {
			sdoas, err = client.CreateOrUpdateResponder(sdoas.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsCreateOrUpdateFuture", "Result", sdoas.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServerDevOpsAuditSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result ServerDevOpsAuditingSettings, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a server's DevOps audit settings.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// devOpsAuditingSettingsName - the name of the devops audit settings. This should always be 'default'.
func (client ServerDevOpsAuditSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string) (result ServerDevOpsAuditingSettings, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerDevOpsAuditSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServerDevOpsAuditSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"devOpsAuditingSettingsName": autorest.Encode("path", devOpsAuditingSettingsName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDevOpsAuditSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerDevOpsAuditSettingsClient) GetResponder(resp *http.Response) (result ServerDevOpsAuditingSettings, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer lists DevOps audit settings of a server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ServerDevOpsAuditSettingsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result ServerDevOpsAuditSettingsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerDevOpsAuditSettingsClient.ListByServer")
		defer func() {
			sc := -1
			if result.sdoaslr.Response.Response != nil {
				sc = result.sdoaslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.sdoaslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.sdoaslr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.sdoaslr.hasNextLink() && result.sdoaslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ServerDevOpsAuditSettingsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDevOpsAuditSettingsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ServerDevOpsAuditSettingsClient) ListByServerResponder(resp *http.Response) (result ServerDevOpsAuditSettingsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client ServerDevOpsAuditSettingsClient) listByServerNextResults(ctx context.Context, lastResults ServerDevOpsAuditSettingsListResult) (result ServerDevOpsAuditSettingsListResult, err error) {
	req, err := lastResults.serverDevOpsAuditSettingsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDevOpsAuditSettingsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServerDevOpsAuditSettingsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result ServerDevOpsAuditSettingsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerDevOpsAuditSettingsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}
