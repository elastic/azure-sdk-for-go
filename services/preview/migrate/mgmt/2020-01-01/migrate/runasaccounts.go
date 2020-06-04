package migrate

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

// RunAsAccountsClient is the discover your workloads for Azure.
type RunAsAccountsClient struct {
	BaseClient
}

// NewRunAsAccountsClient creates an instance of the RunAsAccountsClient client.
func NewRunAsAccountsClient() RunAsAccountsClient {
	return NewRunAsAccountsClientWithBaseURI(DefaultBaseURI)
}

// NewRunAsAccountsClientWithBaseURI creates an instance of the RunAsAccountsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRunAsAccountsClientWithBaseURI(baseURI string) RunAsAccountsClient {
	return RunAsAccountsClient{NewWithBaseURI(baseURI)}
}

// GetAllRunAsAccountsInSite sends the get all run as accounts in site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client RunAsAccountsClient) GetAllRunAsAccountsInSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result VMwareRunAsAccountCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunAsAccountsClient.GetAllRunAsAccountsInSite")
		defer func() {
			sc := -1
			if result.vmraac.Response.Response != nil {
				sc = result.vmraac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAllRunAsAccountsInSiteNextResults
	req, err := client.GetAllRunAsAccountsInSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "GetAllRunAsAccountsInSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAllRunAsAccountsInSiteSender(req)
	if err != nil {
		result.vmraac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "GetAllRunAsAccountsInSite", resp, "Failure sending request")
		return
	}

	result.vmraac, err = client.GetAllRunAsAccountsInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "GetAllRunAsAccountsInSite", resp, "Failure responding to request")
	}

	return
}

// GetAllRunAsAccountsInSitePreparer prepares the GetAllRunAsAccountsInSite request.
func (client RunAsAccountsClient) GetAllRunAsAccountsInSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/runAsAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAllRunAsAccountsInSiteSender sends the GetAllRunAsAccountsInSite request. The method will close the
// http.Response Body if it receives an error.
func (client RunAsAccountsClient) GetAllRunAsAccountsInSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAllRunAsAccountsInSiteResponder handles the response to the GetAllRunAsAccountsInSite request. The method always
// closes the http.Response Body.
func (client RunAsAccountsClient) GetAllRunAsAccountsInSiteResponder(resp *http.Response) (result VMwareRunAsAccountCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAllRunAsAccountsInSiteNextResults retrieves the next set of results, if any.
func (client RunAsAccountsClient) getAllRunAsAccountsInSiteNextResults(ctx context.Context, lastResults VMwareRunAsAccountCollection) (result VMwareRunAsAccountCollection, err error) {
	req, err := lastResults.vMwareRunAsAccountCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "getAllRunAsAccountsInSiteNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAllRunAsAccountsInSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "getAllRunAsAccountsInSiteNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAllRunAsAccountsInSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "getAllRunAsAccountsInSiteNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAllRunAsAccountsInSiteComplete enumerates all values, automatically crossing page boundaries as required.
func (client RunAsAccountsClient) GetAllRunAsAccountsInSiteComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result VMwareRunAsAccountCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunAsAccountsClient.GetAllRunAsAccountsInSite")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAllRunAsAccountsInSite(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	return
}

// GetRunAsAccount sends the get run as account request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// accountName - run as account ARM name.
// APIVersion - the API version to use for this operation.
func (client RunAsAccountsClient) GetRunAsAccount(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, accountName string, APIVersion string) (result VMwareRunAsAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RunAsAccountsClient.GetRunAsAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRunAsAccountPreparer(ctx, subscriptionID, resourceGroupName, siteName, accountName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "GetRunAsAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRunAsAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "GetRunAsAccount", resp, "Failure sending request")
		return
	}

	result, err = client.GetRunAsAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.RunAsAccountsClient", "GetRunAsAccount", resp, "Failure responding to request")
	}

	return
}

// GetRunAsAccountPreparer prepares the GetRunAsAccount request.
func (client RunAsAccountsClient) GetRunAsAccountPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, accountName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/runAsAccounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRunAsAccountSender sends the GetRunAsAccount request. The method will close the
// http.Response Body if it receives an error.
func (client RunAsAccountsClient) GetRunAsAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetRunAsAccountResponder handles the response to the GetRunAsAccount request. The method always
// closes the http.Response Body.
func (client RunAsAccountsClient) GetRunAsAccountResponder(resp *http.Response) (result VMwareRunAsAccount, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
