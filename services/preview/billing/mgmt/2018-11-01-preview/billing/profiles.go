package billing

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

// ProfilesClient is the billing client provides access to billing resources for Azure subscriptions.
type ProfilesClient struct {
	BaseClient
}

// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return ProfilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create the operation to create a BillingProfile.
// Parameters:
// billingAccountName - billing Account Id.
// parameters - parameters supplied to the Create BillingProfile operation.
func (client ProfilesClient) Create(ctx context.Context, billingAccountName string, parameters ProfileCreationParameters) (result ProfilesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, billingAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Create", nil, "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ProfilesClient) CreatePreparer(ctx context.Context, billingAccountName string, parameters ProfileCreationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) CreateSender(req *http.Request) (future ProfilesCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ProfilesClient) (p Profile, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "billing.ProfilesCreateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("billing.ProfilesCreateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "billing.ProfilesCreateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.CreateResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "billing.ProfilesCreateFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CreateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the billing profile by id.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
// expand - may be used to expand the invoiceSections.
func (client ProfilesClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, expand string) (result Profile, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, billingProfileName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProfilesClient) GetPreparer(ctx context.Context, billingAccountName string, billingProfileName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccountName lists all billing profiles for a user which that user has access to.
// Parameters:
// billingAccountName - billing Account Id.
// expand - may be used to expand the invoiceSections.
func (client ProfilesClient) ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string) (result ProfileListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNameNextResults
	req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "ListByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "ListByBillingAccountName", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "ListByBillingAccountName", resp, "Failure responding to request")
		return
	}
	if result.plr.hasNextLink() && result.plr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
func (client ProfilesClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListByBillingAccountNameResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNameNextResults retrieves the next set of results, if any.
func (client ProfilesClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults ProfileListResult) (result ProfileListResult, err error) {
	req, err := lastResults.profileListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProfilesClient", "listByBillingAccountNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProfilesClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProfilesClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string, expand string) (result ProfileListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccountName(ctx, billingAccountName, expand)
	return
}

// Update the operation to update a billing profile.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
// parameters - parameters supplied to the update billing profile operation.
func (client ProfilesClient) Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile) (result ProfilesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProfilesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, billingAccountName, billingProfileName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProfilesClient", "Update", nil, "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ProfilesClient) UpdatePreparer(ctx context.Context, billingAccountName string, billingProfileName string, parameters Profile) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) UpdateSender(req *http.Request) (future ProfilesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client ProfilesClient) (p Profile, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "billing.ProfilesUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("billing.ProfilesUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		p.Response.Response, err = future.GetResult(sender)
		if p.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "billing.ProfilesUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && p.Response.Response.StatusCode != http.StatusNoContent {
			p, err = client.UpdateResponder(p.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "billing.ProfilesUpdateFuture", "Result", p.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProfilesClient) UpdateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
