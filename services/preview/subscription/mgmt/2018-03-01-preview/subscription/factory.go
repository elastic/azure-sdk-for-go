package subscription

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

// FactoryClient is the the subscription client
type FactoryClient struct {
	BaseClient
}

// NewFactoryClient creates an instance of the FactoryClient client.
func NewFactoryClient() FactoryClient {
	return NewFactoryClientWithBaseURI(DefaultBaseURI)
}

// NewFactoryClientWithBaseURI creates an instance of the FactoryClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFactoryClientWithBaseURI(baseURI string) FactoryClient {
	return FactoryClient{NewWithBaseURI(baseURI)}
}

// CreateSubscriptionInEnrollmentAccount creates an Azure subscription
// Parameters:
// enrollmentAccountName - the name of the enrollment account to which the subscription will be billed.
// body - the subscription creation parameters.
func (client FactoryClient) CreateSubscriptionInEnrollmentAccount(ctx context.Context, enrollmentAccountName string, body CreationParameters) (result FactoryCreateSubscriptionInEnrollmentAccountFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FactoryClient.CreateSubscriptionInEnrollmentAccount")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateSubscriptionInEnrollmentAccountPreparer(ctx, enrollmentAccountName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.FactoryClient", "CreateSubscriptionInEnrollmentAccount", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSubscriptionInEnrollmentAccountSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.FactoryClient", "CreateSubscriptionInEnrollmentAccount", nil, "Failure sending request")
		return
	}

	return
}

// CreateSubscriptionInEnrollmentAccountPreparer prepares the CreateSubscriptionInEnrollmentAccount request.
func (client FactoryClient) CreateSubscriptionInEnrollmentAccountPreparer(ctx context.Context, enrollmentAccountName string, body CreationParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"enrollmentAccountName": autorest.Encode("path", enrollmentAccountName),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountName}/providers/Microsoft.Subscription/createSubscription", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSubscriptionInEnrollmentAccountSender sends the CreateSubscriptionInEnrollmentAccount request. The method will close the
// http.Response Body if it receives an error.
func (client FactoryClient) CreateSubscriptionInEnrollmentAccountSender(req *http.Request) (future FactoryCreateSubscriptionInEnrollmentAccountFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client FactoryClient) (cr CreationResult, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "subscription.FactoryCreateSubscriptionInEnrollmentAccountFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("subscription.FactoryCreateSubscriptionInEnrollmentAccountFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		cr.Response.Response, err = future.GetResult(sender)
		if cr.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "subscription.FactoryCreateSubscriptionInEnrollmentAccountFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && cr.Response.Response.StatusCode != http.StatusNoContent {
			cr, err = client.CreateSubscriptionInEnrollmentAccountResponder(cr.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "subscription.FactoryCreateSubscriptionInEnrollmentAccountFuture", "Result", cr.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// CreateSubscriptionInEnrollmentAccountResponder handles the response to the CreateSubscriptionInEnrollmentAccount request. The method always
// closes the http.Response Body.
func (client FactoryClient) CreateSubscriptionInEnrollmentAccountResponder(resp *http.Response) (result CreationResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
