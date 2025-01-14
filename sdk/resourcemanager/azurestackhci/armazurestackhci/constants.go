//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

const (
	moduleName    = "armazurestackhci"
	moduleVersion = "v2.0.0-beta.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CloudInitDataSource - Datasource for the gallery image when provisioning with cloud-init [NoCloud, Azure]
type CloudInitDataSource string

const (
	CloudInitDataSourceAzure   CloudInitDataSource = "Azure"
	CloudInitDataSourceNoCloud CloudInitDataSource = "NoCloud"
)

// PossibleCloudInitDataSourceValues returns the possible values for the CloudInitDataSource const type.
func PossibleCloudInitDataSourceValues() []CloudInitDataSource {
	return []CloudInitDataSource{
		CloudInitDataSourceAzure,
		CloudInitDataSourceNoCloud,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DiskFileFormat - The format of the actual VHD file [vhd, vhdx]
type DiskFileFormat string

const (
	DiskFileFormatVhd  DiskFileFormat = "vhd"
	DiskFileFormatVhdx DiskFileFormat = "vhdx"
)

// PossibleDiskFileFormatValues returns the possible values for the DiskFileFormat const type.
func PossibleDiskFileFormatValues() []DiskFileFormat {
	return []DiskFileFormat{
		DiskFileFormatVhd,
		DiskFileFormatVhdx,
	}
}

// ExtendedLocationTypes - The type of extendedLocation.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation ExtendedLocationTypes = "CustomLocation"
)

// PossibleExtendedLocationTypesValues returns the possible values for the ExtendedLocationTypes const type.
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return []ExtendedLocationTypes{
		ExtendedLocationTypesCustomLocation,
	}
}

// HyperVGeneration - The hypervisor generation of the Virtual Machine [V1, V2]
type HyperVGeneration string

const (
	HyperVGenerationV1 HyperVGeneration = "V1"
	HyperVGenerationV2 HyperVGeneration = "V2"
)

// PossibleHyperVGenerationValues returns the possible values for the HyperVGeneration const type.
func PossibleHyperVGenerationValues() []HyperVGeneration {
	return []HyperVGeneration{
		HyperVGenerationV1,
		HyperVGenerationV2,
	}
}

// IPAllocationMethodEnum - IPAllocationMethod - The IP address allocation method. Possible values include: 'Static', 'Dynamic'
type IPAllocationMethodEnum string

const (
	IPAllocationMethodEnumDynamic IPAllocationMethodEnum = "Dynamic"
	IPAllocationMethodEnumStatic  IPAllocationMethodEnum = "Static"
)

// PossibleIPAllocationMethodEnumValues returns the possible values for the IPAllocationMethodEnum const type.
func PossibleIPAllocationMethodEnumValues() []IPAllocationMethodEnum {
	return []IPAllocationMethodEnum{
		IPAllocationMethodEnumDynamic,
		IPAllocationMethodEnumStatic,
	}
}

// IPPoolTypeEnum - Type of the IP Pool [vm, vippool]
type IPPoolTypeEnum string

const (
	IPPoolTypeEnumVM      IPPoolTypeEnum = "vm"
	IPPoolTypeEnumVippool IPPoolTypeEnum = "vippool"
)

// PossibleIPPoolTypeEnumValues returns the possible values for the IPPoolTypeEnum const type.
func PossibleIPPoolTypeEnumValues() []IPPoolTypeEnum {
	return []IPPoolTypeEnum{
		IPPoolTypeEnumVM,
		IPPoolTypeEnumVippool,
	}
}

// OperatingSystemTypes - Operating system type that the gallery image uses [Windows, Linux]
type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns the possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{
		OperatingSystemTypesLinux,
		OperatingSystemTypesWindows,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PowerStateEnum - The power state of the virtual machine instance
type PowerStateEnum string

const (
	PowerStateEnumDeallocated  PowerStateEnum = "Deallocated"
	PowerStateEnumDeallocating PowerStateEnum = "Deallocating"
	PowerStateEnumRunning      PowerStateEnum = "Running"
	PowerStateEnumStarting     PowerStateEnum = "Starting"
	PowerStateEnumStopped      PowerStateEnum = "Stopped"
	PowerStateEnumStopping     PowerStateEnum = "Stopping"
	PowerStateEnumUnknown      PowerStateEnum = "Unknown"
)

// PossiblePowerStateEnumValues returns the possible values for the PowerStateEnum const type.
func PossiblePowerStateEnumValues() []PowerStateEnum {
	return []PowerStateEnum{
		PowerStateEnumDeallocated,
		PowerStateEnumDeallocating,
		PowerStateEnumRunning,
		PowerStateEnumStarting,
		PowerStateEnumStopped,
		PowerStateEnumStopping,
		PowerStateEnumUnknown,
	}
}

// ProvisioningAction - Defines the different types of operations for guest agent.
type ProvisioningAction string

const (
	ProvisioningActionInstall   ProvisioningAction = "install"
	ProvisioningActionRepair    ProvisioningAction = "repair"
	ProvisioningActionUninstall ProvisioningAction = "uninstall"
)

// PossibleProvisioningActionValues returns the possible values for the ProvisioningAction const type.
func PossibleProvisioningActionValues() []ProvisioningAction {
	return []ProvisioningAction{
		ProvisioningActionInstall,
		ProvisioningActionRepair,
		ProvisioningActionUninstall,
	}
}

// ProvisioningStateEnum - Provisioning state of the gallery image.
type ProvisioningStateEnum string

const (
	ProvisioningStateEnumAccepted   ProvisioningStateEnum = "Accepted"
	ProvisioningStateEnumCanceled   ProvisioningStateEnum = "Canceled"
	ProvisioningStateEnumDeleting   ProvisioningStateEnum = "Deleting"
	ProvisioningStateEnumFailed     ProvisioningStateEnum = "Failed"
	ProvisioningStateEnumInProgress ProvisioningStateEnum = "InProgress"
	ProvisioningStateEnumSucceeded  ProvisioningStateEnum = "Succeeded"
)

// PossibleProvisioningStateEnumValues returns the possible values for the ProvisioningStateEnum const type.
func PossibleProvisioningStateEnumValues() []ProvisioningStateEnum {
	return []ProvisioningStateEnum{
		ProvisioningStateEnumAccepted,
		ProvisioningStateEnumCanceled,
		ProvisioningStateEnumDeleting,
		ProvisioningStateEnumFailed,
		ProvisioningStateEnumInProgress,
		ProvisioningStateEnumSucceeded,
	}
}

// SecurityTypes - Specifies the SecurityType of the virtual machine. EnableTPM and SecureBootEnabled must be set to true
// for SecurityType to function.
type SecurityTypes string

const (
	SecurityTypesConfidentialVM SecurityTypes = "ConfidentialVM"
	SecurityTypesTrustedLaunch  SecurityTypes = "TrustedLaunch"
)

// PossibleSecurityTypesValues returns the possible values for the SecurityTypes const type.
func PossibleSecurityTypesValues() []SecurityTypes {
	return []SecurityTypes{
		SecurityTypesConfidentialVM,
		SecurityTypesTrustedLaunch,
	}
}

// Status - The status of the operation performed on the gallery image [Succeeded, Failed, InProgress]
type Status string

const (
	StatusFailed     Status = "Failed"
	StatusInProgress Status = "InProgress"
	StatusSucceeded  Status = "Succeeded"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusFailed,
		StatusInProgress,
		StatusSucceeded,
	}
}

// StatusLevelTypes - The level code.
type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

// PossibleStatusLevelTypesValues returns the possible values for the StatusLevelTypes const type.
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return []StatusLevelTypes{
		StatusLevelTypesError,
		StatusLevelTypesInfo,
		StatusLevelTypesWarning,
	}
}

// StatusTypes - The installation status of the hybrid machine agent installation.
type StatusTypes string

const (
	StatusTypesFailed     StatusTypes = "Failed"
	StatusTypesInProgress StatusTypes = "InProgress"
	StatusTypesSucceeded  StatusTypes = "Succeeded"
)

// PossibleStatusTypesValues returns the possible values for the StatusTypes const type.
func PossibleStatusTypesValues() []StatusTypes {
	return []StatusTypes{
		StatusTypesFailed,
		StatusTypesInProgress,
		StatusTypesSucceeded,
	}
}

type VMSizeEnum string

const (
	VMSizeEnumCustom         VMSizeEnum = "Custom"
	VMSizeEnumDefault        VMSizeEnum = "Default"
	VMSizeEnumStandardA2V2   VMSizeEnum = "Standard_A2_v2"
	VMSizeEnumStandardA4V2   VMSizeEnum = "Standard_A4_v2"
	VMSizeEnumStandardD16SV3 VMSizeEnum = "Standard_D16s_v3"
	VMSizeEnumStandardD2SV3  VMSizeEnum = "Standard_D2s_v3"
	VMSizeEnumStandardD32SV3 VMSizeEnum = "Standard_D32s_v3"
	VMSizeEnumStandardD4SV3  VMSizeEnum = "Standard_D4s_v3"
	VMSizeEnumStandardD8SV3  VMSizeEnum = "Standard_D8s_v3"
	VMSizeEnumStandardDS13V2 VMSizeEnum = "Standard_DS13_v2"
	VMSizeEnumStandardDS2V2  VMSizeEnum = "Standard_DS2_v2"
	VMSizeEnumStandardDS3V2  VMSizeEnum = "Standard_DS3_v2"
	VMSizeEnumStandardDS4V2  VMSizeEnum = "Standard_DS4_v2"
	VMSizeEnumStandardDS5V2  VMSizeEnum = "Standard_DS5_v2"
	VMSizeEnumStandardK8S2V1 VMSizeEnum = "Standard_K8S2_v1"
	VMSizeEnumStandardK8S3V1 VMSizeEnum = "Standard_K8S3_v1"
	VMSizeEnumStandardK8S4V1 VMSizeEnum = "Standard_K8S4_v1"
	VMSizeEnumStandardK8S5V1 VMSizeEnum = "Standard_K8S5_v1"
	VMSizeEnumStandardK8SV1  VMSizeEnum = "Standard_K8S_v1"
	VMSizeEnumStandardNK12   VMSizeEnum = "Standard_NK12"
	VMSizeEnumStandardNK6    VMSizeEnum = "Standard_NK6"
	VMSizeEnumStandardNV12   VMSizeEnum = "Standard_NV12"
	VMSizeEnumStandardNV6    VMSizeEnum = "Standard_NV6"
)

// PossibleVMSizeEnumValues returns the possible values for the VMSizeEnum const type.
func PossibleVMSizeEnumValues() []VMSizeEnum {
	return []VMSizeEnum{
		VMSizeEnumCustom,
		VMSizeEnumDefault,
		VMSizeEnumStandardA2V2,
		VMSizeEnumStandardA4V2,
		VMSizeEnumStandardD16SV3,
		VMSizeEnumStandardD2SV3,
		VMSizeEnumStandardD32SV3,
		VMSizeEnumStandardD4SV3,
		VMSizeEnumStandardD8SV3,
		VMSizeEnumStandardDS13V2,
		VMSizeEnumStandardDS2V2,
		VMSizeEnumStandardDS3V2,
		VMSizeEnumStandardDS4V2,
		VMSizeEnumStandardDS5V2,
		VMSizeEnumStandardK8S2V1,
		VMSizeEnumStandardK8S3V1,
		VMSizeEnumStandardK8S4V1,
		VMSizeEnumStandardK8S5V1,
		VMSizeEnumStandardK8SV1,
		VMSizeEnumStandardNK12,
		VMSizeEnumStandardNK6,
		VMSizeEnumStandardNV12,
		VMSizeEnumStandardNV6,
	}
}
