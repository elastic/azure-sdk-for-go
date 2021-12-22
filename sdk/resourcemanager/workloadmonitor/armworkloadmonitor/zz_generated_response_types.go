//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadmonitor

import "net/http"

// HealthMonitorsGetResponse contains the response from method HealthMonitors.Get.
type HealthMonitorsGetResponse struct {
	HealthMonitorsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HealthMonitorsGetResult contains the result from method HealthMonitors.Get.
type HealthMonitorsGetResult struct {
	HealthMonitor
}

// HealthMonitorsGetStateChangeResponse contains the response from method HealthMonitors.GetStateChange.
type HealthMonitorsGetStateChangeResponse struct {
	HealthMonitorsGetStateChangeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HealthMonitorsGetStateChangeResult contains the result from method HealthMonitors.GetStateChange.
type HealthMonitorsGetStateChangeResult struct {
	HealthMonitorStateChange
}

// HealthMonitorsListResponse contains the response from method HealthMonitors.List.
type HealthMonitorsListResponse struct {
	HealthMonitorsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HealthMonitorsListResult contains the result from method HealthMonitors.List.
type HealthMonitorsListResult struct {
	HealthMonitorList
}

// HealthMonitorsListStateChangesResponse contains the response from method HealthMonitors.ListStateChanges.
type HealthMonitorsListStateChangesResponse struct {
	HealthMonitorsListStateChangesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HealthMonitorsListStateChangesResult contains the result from method HealthMonitors.ListStateChanges.
type HealthMonitorsListStateChangesResult struct {
	HealthMonitorStateChangeList
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationList
}