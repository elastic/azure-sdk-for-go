package eventgrid

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

// DomainTopicsClient is the azure EventGrid Management Client
type DomainTopicsClient struct {
	BaseClient
}

// NewDomainTopicsClient creates an instance of the DomainTopicsClient client.
func NewDomainTopicsClient(subscriptionID string) DomainTopicsClient {
	return NewDomainTopicsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDomainTopicsClientWithBaseURI creates an instance of the DomainTopicsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDomainTopicsClientWithBaseURI(baseURI string, subscriptionID string) DomainTopicsClient {
	return DomainTopicsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate asynchronously creates or updates a new domain topic with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// domainName - name of the domain
// domainTopicName - name of the domain topic
func (client DomainTopicsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result DomainTopicsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, domainName, domainTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DomainTopicsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        autorest.Encode("path", domainName),
		"domainTopicName":   autorest.Encode("path", domainTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DomainTopicsClient) CreateOrUpdateSender(req *http.Request) (future DomainTopicsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DomainTopicsClient) (dt DomainTopic, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("eventgrid.DomainTopicsCreateOrUpdateFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		dt.Response.Response, err = future.GetResult(sender)
		if dt.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsCreateOrUpdateFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && dt.Response.Response.StatusCode != http.StatusNoContent {
			dt, err = client.CreateOrUpdateResponder(dt.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsCreateOrUpdateFuture", "Result", dt.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DomainTopicsClient) CreateOrUpdateResponder(resp *http.Response) (result DomainTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete existing domain topic
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// domainName - name of the domain
// domainTopicName - name of the domain topic
func (client DomainTopicsClient) Delete(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result DomainTopicsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, domainName, domainTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DomainTopicsClient) DeletePreparer(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        autorest.Encode("path", domainName),
		"domainTopicName":   autorest.Encode("path", domainTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DomainTopicsClient) DeleteSender(req *http.Request) (future DomainTopicsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client DomainTopicsClient) (ar autorest.Response, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("eventgrid.DomainTopicsDeleteFuture")
			return
		}
		ar.Response = future.Response()
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DomainTopicsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get properties of a domain topic
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// domainName - name of the domain
// domainTopicName - name of the topic
func (client DomainTopicsClient) Get(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result DomainTopic, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, domainName, domainTopicName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DomainTopicsClient) GetPreparer(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        autorest.Encode("path", domainName),
		"domainTopicName":   autorest.Encode("path", domainTopicName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DomainTopicsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DomainTopicsClient) GetResponder(resp *http.Response) (result DomainTopic, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDomain list all the topics in a domain.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// domainName - domain name.
// filter - filter the results using OData syntax.
// top - the number of results to return.
func (client DomainTopicsClient) ListByDomain(ctx context.Context, resourceGroupName string, domainName string, filter string, top *int32) (result DomainTopicsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.ListByDomain")
		defer func() {
			sc := -1
			if result.dtlr.Response.Response != nil {
				sc = result.dtlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDomainNextResults
	req, err := client.ListByDomainPreparer(ctx, resourceGroupName, domainName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "ListByDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDomainSender(req)
	if err != nil {
		result.dtlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "ListByDomain", resp, "Failure sending request")
		return
	}

	result.dtlr, err = client.ListByDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "ListByDomain", resp, "Failure responding to request")
		return
	}
	if result.dtlr.hasNextLink() && result.dtlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDomainPreparer prepares the ListByDomain request.
func (client DomainTopicsClient) ListByDomainPreparer(ctx context.Context, resourceGroupName string, domainName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"domainName":        autorest.Encode("path", domainName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDomainSender sends the ListByDomain request. The method will close the
// http.Response Body if it receives an error.
func (client DomainTopicsClient) ListByDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDomainResponder handles the response to the ListByDomain request. The method always
// closes the http.Response Body.
func (client DomainTopicsClient) ListByDomainResponder(resp *http.Response) (result DomainTopicsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDomainNextResults retrieves the next set of results, if any.
func (client DomainTopicsClient) listByDomainNextResults(ctx context.Context, lastResults DomainTopicsListResult) (result DomainTopicsListResult, err error) {
	req, err := lastResults.domainTopicsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "listByDomainNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "listByDomainNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.DomainTopicsClient", "listByDomainNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDomainComplete enumerates all values, automatically crossing page boundaries as required.
func (client DomainTopicsClient) ListByDomainComplete(ctx context.Context, resourceGroupName string, domainName string, filter string, top *int32) (result DomainTopicsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainTopicsClient.ListByDomain")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDomain(ctx, resourceGroupName, domainName, filter, top)
	return
}
