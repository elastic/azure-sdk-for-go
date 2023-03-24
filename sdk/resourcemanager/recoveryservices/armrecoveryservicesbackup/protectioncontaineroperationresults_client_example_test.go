//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureStorage/ProtectionContainers_Inquire_Result.json
func ExampleProtectionContainerOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectionContainerOperationResultsClient().Get(ctx, "testvault", "test-rg", "Azure", "VMAppContainer;Compute;testRG;testSQL", "00000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectionContainerResource = armrecoveryservicesbackup.ProtectionContainerResource{
	// 	Name: to.Ptr("VMAppContainer;Compute;testRG;testSQL"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
	// 	ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.RecoveryServices/vaults/testVault/backupFabrics/Azure/protectionContainers/VMAppContainer;Compute;testRG;testSQL"),
	// 	Properties: &armrecoveryservicesbackup.AzureVMAppContainerProtectionContainer{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureWorkload),
	// 		ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeVMAppContainer),
	// 		FriendlyName: to.Ptr("testSQL"),
	// 		ExtendedInfo: &armrecoveryservicesbackup.AzureWorkloadContainerExtendedInfo{
	// 			HostServerName: to.Ptr("testsql"),
	// 			InquiryInfo: &armrecoveryservicesbackup.InquiryInfo{
	// 				ErrorDetail: &armrecoveryservicesbackup.ErrorDetail{
	// 					Code: to.Ptr("Success"),
	// 					Message: to.Ptr("Not Available"),
	// 					Recommendations: []*string{
	// 						to.Ptr("Not Available")},
	// 					},
	// 					InquiryDetails: []*armrecoveryservicesbackup.WorkloadInquiryDetails{
	// 						{
	// 							Type: to.Ptr("sql"),
	// 							InquiryValidation: &armrecoveryservicesbackup.InquiryValidation{
	// 								ErrorDetail: &armrecoveryservicesbackup.ErrorDetail{
	// 									Code: to.Ptr("Success"),
	// 									Message: to.Ptr("Not Available"),
	// 									Recommendations: []*string{
	// 										to.Ptr("Not Available")},
	// 									},
	// 									Status: to.Ptr("Success"),
	// 								},
	// 								ItemCount: to.Ptr[int64](14),
	// 						}},
	// 						Status: to.Ptr("Success"),
	// 					},
	// 					NodesList: []*armrecoveryservicesbackup.DistributedNodesInfo{
	// 					},
	// 				},
	// 				SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/testSQL"),
	// 			},
	// 		}
}
