package managedapplications

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

// ApplianceDefinitionsClient is the ARM managed applications (appliances)
type ApplianceDefinitionsClient struct {
	BaseClient
}

// NewApplianceDefinitionsClient creates an instance of the ApplianceDefinitionsClient client.
func NewApplianceDefinitionsClient(subscriptionID string) ApplianceDefinitionsClient {
	return NewApplianceDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplianceDefinitionsClientWithBaseURI creates an instance of the ApplianceDefinitionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewApplianceDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ApplianceDefinitionsClient {
	return ApplianceDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new appliance definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applianceDefinitionName - the name of the appliance definition.
// parameters - parameters supplied to the create or update an appliance definition.
func (client ApplianceDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applianceDefinitionName string, parameters ApplianceDefinition) (result ApplianceDefinitionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: applianceDefinitionName,
			Constraints: []validation.Constraint{{Target: "applianceDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applianceDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties.Authorizations", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ApplianceDefinitionProperties.PackageFileURI", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, applianceDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplianceDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, applianceDefinitionName string, parameters ApplianceDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionName": autorest.Encode("path", applianceDefinitionName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) CreateOrUpdateSender(req *http.Request) (future ApplianceDefinitionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplianceDefinitionsClient) (ad ApplianceDefinition, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplianceDefinitionsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ad.Response.Response, err = future.GetResult(sender)
		if ad.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ad.Response.Response.StatusCode != http.StatusNoContent {
			ad, err = client.CreateOrUpdateResponder(ad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsCreateOrUpdateFuture", "Result", ad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByID creates a new appliance definition.
// Parameters:
// applianceDefinitionID - the fully qualified ID of the appliance definition, including the appliance name and
// the appliance definition resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applianceDefinitions/{applianceDefinition-name}
// parameters - parameters supplied to the create or update an appliance definition.
func (client ApplianceDefinitionsClient) CreateOrUpdateByID(ctx context.Context, applianceDefinitionID string, parameters ApplianceDefinition) (result ApplianceDefinitionsCreateOrUpdateByIDFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.CreateOrUpdateByID")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ApplianceDefinitionProperties.Authorizations", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ApplianceDefinitionProperties.PackageFileURI", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID", err.Error())
	}

	req, err := client.CreateOrUpdateByIDPreparer(ctx, applianceDefinitionID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateByIDSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "CreateOrUpdateByID", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateByIDPreparer prepares the CreateOrUpdateByID request.
func (client ApplianceDefinitionsClient) CreateOrUpdateByIDPreparer(ctx context.Context, applianceDefinitionID string, parameters ApplianceDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionId": applianceDefinitionID,
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applianceDefinitionId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateByIDSender sends the CreateOrUpdateByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) CreateOrUpdateByIDSender(req *http.Request) (future ApplianceDefinitionsCreateOrUpdateByIDFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplianceDefinitionsClient) (ad ApplianceDefinition, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsCreateOrUpdateByIDFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplianceDefinitionsCreateOrUpdateByIDFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		ad.Response.Response, err = future.GetResult(sender)
		if ad.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsCreateOrUpdateByIDFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && ad.Response.Response.StatusCode != http.StatusNoContent {
			ad, err = client.CreateOrUpdateByIDResponder(ad.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsCreateOrUpdateByIDFuture", "Result", ad.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateByIDResponder handles the response to the CreateOrUpdateByID request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) CreateOrUpdateByIDResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the appliance definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applianceDefinitionName - the name of the appliance definition to delete.
func (client ApplianceDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, applianceDefinitionName string) (result ApplianceDefinitionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: applianceDefinitionName,
			Constraints: []validation.Constraint{{Target: "applianceDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applianceDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplianceDefinitionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, applianceDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplianceDefinitionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, applianceDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionName": autorest.Encode("path", applianceDefinitionName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) DeleteSender(req *http.Request) (future ApplianceDefinitionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplianceDefinitionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplianceDefinitionsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByID deletes the appliance definition.
// Parameters:
// applianceDefinitionID - the fully qualified ID of the appliance definition, including the appliance name and
// the appliance definition resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applianceDefinitions/{applianceDefinition-name}
func (client ApplianceDefinitionsClient) DeleteByID(ctx context.Context, applianceDefinitionID string) (result ApplianceDefinitionsDeleteByIDFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.DeleteByID")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByIDPreparer(ctx, applianceDefinitionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "DeleteByID", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteByIDSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "DeleteByID", nil, "Failure sending request")
		return
	}

	return
}

// DeleteByIDPreparer prepares the DeleteByID request.
func (client ApplianceDefinitionsClient) DeleteByIDPreparer(ctx context.Context, applianceDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionId": applianceDefinitionID,
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applianceDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByIDSender sends the DeleteByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) DeleteByIDSender(req *http.Request) (future ApplianceDefinitionsDeleteByIDFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ApplianceDefinitionsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsDeleteByIDFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("managedapplications.ApplianceDefinitionsDeleteByIDFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteByIDResponder handles the response to the DeleteByID request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) DeleteByIDResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the appliance definition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applianceDefinitionName - the name of the appliance definition.
func (client ApplianceDefinitionsClient) Get(ctx context.Context, resourceGroupName string, applianceDefinitionName string) (result ApplianceDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}},
		{TargetValue: applianceDefinitionName,
			Constraints: []validation.Constraint{{Target: "applianceDefinitionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applianceDefinitionName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplianceDefinitionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, applianceDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplianceDefinitionsClient) GetPreparer(ctx context.Context, resourceGroupName string, applianceDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionName": autorest.Encode("path", applianceDefinitionName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) GetResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets the appliance definition.
// Parameters:
// applianceDefinitionID - the fully qualified ID of the appliance definition, including the appliance name and
// the appliance definition resource type. Use the format,
// /subscriptions/{guid}/resourceGroups/{resource-group-name}/Microsoft.Solutions/applianceDefinitions/{applianceDefinition-name}
func (client ApplianceDefinitionsClient) GetByID(ctx context.Context, applianceDefinitionID string) (result ApplianceDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, applianceDefinitionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "GetByID", resp, "Failure responding to request")
		return
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client ApplianceDefinitionsClient) GetByIDPreparer(ctx context.Context, applianceDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applianceDefinitionId": applianceDefinitionID,
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{applianceDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) GetByIDResponder(resp *http.Response) (result ApplianceDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists the appliance definitions in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ApplianceDefinitionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ApplianceDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.adlr.Response.Response != nil {
				sc = result.adlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\p{L}\._\(\)\w]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.adlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.adlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.adlr.hasNextLink() && result.adlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ApplianceDefinitionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ApplianceDefinitionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ApplianceDefinitionsClient) ListByResourceGroupResponder(resp *http.Response) (result ApplianceDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ApplianceDefinitionsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ApplianceDefinitionListResult) (result ApplianceDefinitionListResult, err error) {
	req, err := lastResults.applianceDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedapplications.ApplianceDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplianceDefinitionsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ApplianceDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplianceDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}
