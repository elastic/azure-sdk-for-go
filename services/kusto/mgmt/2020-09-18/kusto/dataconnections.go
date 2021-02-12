package kusto

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

// DataConnectionsClient is the the Azure Kusto management API provides a RESTful set of web services that interact
// with Azure Kusto services to manage your clusters and databases. The API enables you to create, update, and delete
// clusters and databases.
type DataConnectionsClient struct {
	BaseClient
}

// NewDataConnectionsClient creates an instance of the DataConnectionsClient client.
func NewDataConnectionsClient(subscriptionID string) DataConnectionsClient {
	return NewDataConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataConnectionsClientWithBaseURI creates an instance of the DataConnectionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataConnectionsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectionsClient {
	return DataConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks that the data connection name is valid and is not already in use.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
func (client DataConnectionsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest) (result CheckNameResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataConnectionName,
			Constraints: []validation.Constraint{{Target: "dataConnectionName.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "dataConnectionName.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kusto.DataConnectionsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client DataConnectionsClient) CheckNameAvailabilityPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/checkNameAvailability", pathParameters),
		autorest.WithJSON(dataConnectionName),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a data connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
// parameters - the data connection parameters supplied to the CreateOrUpdate operation.
func (client DataConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (result DataConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DataConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) CreateOrUpdateSender(req *http.Request) (future DataConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DataConnectionsClient) (dcm DataConnectionModel, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.DataConnectionsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		dcm.Response.Response, err = future.GetResult(sender)
		if dcm.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && dcm.Response.Response.StatusCode != http.StatusNoContent {
			dcm, err = client.CreateOrUpdateResponder(dcm.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "kusto.DataConnectionsCreateOrUpdateFuture", "Result", dcm.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result DataConnectionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DataConnectionValidationMethod checks that the data connection parameters are valid.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// parameters - the data connection parameters supplied to the CreateOrUpdate operation.
func (client DataConnectionsClient) DataConnectionValidationMethod(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation) (result DataConnectionsDataConnectionValidationMethodFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.DataConnectionValidationMethod")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DataConnectionValidationMethodPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "DataConnectionValidationMethod", nil, "Failure preparing request")
		return
	}

	result, err = client.DataConnectionValidationMethodSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "DataConnectionValidationMethod", nil, "Failure sending request")
		return
	}

	return
}

// DataConnectionValidationMethodPreparer prepares the DataConnectionValidationMethod request.
func (client DataConnectionsClient) DataConnectionValidationMethodPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnectionValidation", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DataConnectionValidationMethodSender sends the DataConnectionValidationMethod request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) DataConnectionValidationMethodSender(req *http.Request) (future DataConnectionsDataConnectionValidationMethodFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DataConnectionsClient) (dcvlr DataConnectionValidationListResult, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsDataConnectionValidationMethodFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.DataConnectionsDataConnectionValidationMethodFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		dcvlr.Response.Response, err = future.GetResult(sender)
		if dcvlr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsDataConnectionValidationMethodFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && dcvlr.Response.Response.StatusCode != http.StatusNoContent {
			dcvlr, err = client.DataConnectionValidationMethodResponder(dcvlr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "kusto.DataConnectionsDataConnectionValidationMethodFuture", "Result", dcvlr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// DataConnectionValidationMethodResponder handles the response to the DataConnectionValidationMethod request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) DataConnectionValidationMethodResponder(resp *http.Response) (result DataConnectionValidationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the data connection with the given name.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
func (client DataConnectionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result DataConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DataConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) DeleteSender(req *http.Request) (future DataConnectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DataConnectionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.DataConnectionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a data connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
func (client DataConnectionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result DataConnectionModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) GetResponder(resp *http.Response) (result DataConnectionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase returns the list of data connections of the given Kusto database.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
func (client DataConnectionsClient) ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DataConnectionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client DataConnectionsClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) ListByDatabaseResponder(resp *http.Response) (result DataConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a data connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
// parameters - the data connection parameters supplied to the Update operation.
func (client DataConnectionsClient) Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (result DataConnectionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DataConnectionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) UpdateSender(req *http.Request) (future DataConnectionsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DataConnectionsClient) (dcm DataConnectionModel, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("kusto.DataConnectionsUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		dcm.Response.Response, err = future.GetResult(sender)
		if dcm.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "kusto.DataConnectionsUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && dcm.Response.Response.StatusCode != http.StatusNoContent {
			dcm, err = client.UpdateResponder(dcm.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "kusto.DataConnectionsUpdateFuture", "Result", dcm.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) UpdateResponder(resp *http.Response) (result DataConnectionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
