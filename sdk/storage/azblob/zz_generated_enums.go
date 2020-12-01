// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

type AccessTier string

const (
	AccessTierArchive AccessTier = "Archive"
	AccessTierCool    AccessTier = "Cool"
	AccessTierHot     AccessTier = "Hot"
	AccessTierP10     AccessTier = "P10"
	AccessTierP15     AccessTier = "P15"
	AccessTierP20     AccessTier = "P20"
	AccessTierP30     AccessTier = "P30"
	AccessTierP4      AccessTier = "P4"
	AccessTierP40     AccessTier = "P40"
	AccessTierP50     AccessTier = "P50"
	AccessTierP6      AccessTier = "P6"
	AccessTierP60     AccessTier = "P60"
	AccessTierP70     AccessTier = "P70"
	AccessTierP80     AccessTier = "P80"
)

func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{
		AccessTierArchive,
		AccessTierCool,
		AccessTierHot,
		AccessTierP10,
		AccessTierP15,
		AccessTierP20,
		AccessTierP30,
		AccessTierP4,
		AccessTierP40,
		AccessTierP50,
		AccessTierP6,
		AccessTierP60,
		AccessTierP70,
		AccessTierP80,
	}
}

func (c AccessTier) ToPtr() *AccessTier {
	return &c
}

type AccountKind string

const (
	AccountKindStorage          AccountKind = "Storage"
	AccountKindBlobStorage      AccountKind = "BlobStorage"
	AccountKindStorageV2        AccountKind = "StorageV2"
	AccountKindFileStorage      AccountKind = "FileStorage"
	AccountKindBlockBlobStorage AccountKind = "BlockBlobStorage"
)

func PossibleAccountKindValues() []AccountKind {
	return []AccountKind{
		AccountKindStorage,
		AccountKindBlobStorage,
		AccountKindStorageV2,
		AccountKindFileStorage,
		AccountKindBlockBlobStorage,
	}
}

func (c AccountKind) ToPtr() *AccountKind {
	return &c
}

type ArchiveStatus string

const (
	ArchiveStatusRehydratePendingToCool ArchiveStatus = "rehydrate-pending-to-cool"
	ArchiveStatusRehydratePendingToHot  ArchiveStatus = "rehydrate-pending-to-hot"
)

func PossibleArchiveStatusValues() []ArchiveStatus {
	return []ArchiveStatus{
		ArchiveStatusRehydratePendingToCool,
		ArchiveStatusRehydratePendingToHot,
	}
}

func (c ArchiveStatus) ToPtr() *ArchiveStatus {
	return &c
}

type BlobExpiryOptions string

const (
	BlobExpiryOptionsAbsolute           BlobExpiryOptions = "Absolute"
	BlobExpiryOptionsNeverExpire        BlobExpiryOptions = "NeverExpire"
	BlobExpiryOptionsRelativeToCreation BlobExpiryOptions = "RelativeToCreation"
	BlobExpiryOptionsRelativeToNow      BlobExpiryOptions = "RelativeToNow"
)

func PossibleBlobExpiryOptionsValues() []BlobExpiryOptions {
	return []BlobExpiryOptions{
		BlobExpiryOptionsAbsolute,
		BlobExpiryOptionsNeverExpire,
		BlobExpiryOptionsRelativeToCreation,
		BlobExpiryOptionsRelativeToNow,
	}
}

func (c BlobExpiryOptions) ToPtr() *BlobExpiryOptions {
	return &c
}

type BlobType string

const (
	BlobTypeBlockBlob  BlobType = "BlockBlob"
	BlobTypePageBlob   BlobType = "PageBlob"
	BlobTypeAppendBlob BlobType = "AppendBlob"
)

func PossibleBlobTypeValues() []BlobType {
	return []BlobType{
		BlobTypeBlockBlob,
		BlobTypePageBlob,
		BlobTypeAppendBlob,
	}
}

func (c BlobType) ToPtr() *BlobType {
	return &c
}

type BlockListType string

const (
	BlockListTypeCommitted   BlockListType = "committed"
	BlockListTypeUncommitted BlockListType = "uncommitted"
	BlockListTypeAll         BlockListType = "all"
)

func PossibleBlockListTypeValues() []BlockListType {
	return []BlockListType{
		BlockListTypeCommitted,
		BlockListTypeUncommitted,
		BlockListTypeAll,
	}
}

func (c BlockListType) ToPtr() *BlockListType {
	return &c
}

type CopyStatusType string

const (
	CopyStatusTypePending CopyStatusType = "pending"
	CopyStatusTypeSuccess CopyStatusType = "success"
	CopyStatusTypeAborted CopyStatusType = "aborted"
	CopyStatusTypeFailed  CopyStatusType = "failed"
)

func PossibleCopyStatusTypeValues() []CopyStatusType {
	return []CopyStatusType{
		CopyStatusTypePending,
		CopyStatusTypeSuccess,
		CopyStatusTypeAborted,
		CopyStatusTypeFailed,
	}
}

func (c CopyStatusType) ToPtr() *CopyStatusType {
	return &c
}

type DeleteSnapshotsOptionType string

const (
	DeleteSnapshotsOptionTypeInclude DeleteSnapshotsOptionType = "include"
	DeleteSnapshotsOptionTypeOnly    DeleteSnapshotsOptionType = "only"
)

func PossibleDeleteSnapshotsOptionTypeValues() []DeleteSnapshotsOptionType {
	return []DeleteSnapshotsOptionType{
		DeleteSnapshotsOptionTypeInclude,
		DeleteSnapshotsOptionTypeOnly,
	}
}

func (c DeleteSnapshotsOptionType) ToPtr() *DeleteSnapshotsOptionType {
	return &c
}

// GeoReplicationStatusType - The status of the secondary location
type GeoReplicationStatusType string

const (
	GeoReplicationStatusTypeBootstrap   GeoReplicationStatusType = "bootstrap"
	GeoReplicationStatusTypeLive        GeoReplicationStatusType = "live"
	GeoReplicationStatusTypeUnavailable GeoReplicationStatusType = "unavailable"
)

func PossibleGeoReplicationStatusTypeValues() []GeoReplicationStatusType {
	return []GeoReplicationStatusType{
		GeoReplicationStatusTypeBootstrap,
		GeoReplicationStatusTypeLive,
		GeoReplicationStatusTypeUnavailable,
	}
}

func (c GeoReplicationStatusType) ToPtr() *GeoReplicationStatusType {
	return &c
}

type LeaseDurationType string

const (
	LeaseDurationTypeInfinite LeaseDurationType = "infinite"
	LeaseDurationTypeFixed    LeaseDurationType = "fixed"
)

func PossibleLeaseDurationTypeValues() []LeaseDurationType {
	return []LeaseDurationType{
		LeaseDurationTypeInfinite,
		LeaseDurationTypeFixed,
	}
}

func (c LeaseDurationType) ToPtr() *LeaseDurationType {
	return &c
}

type LeaseStateType string

const (
	LeaseStateTypeAvailable LeaseStateType = "available"
	LeaseStateTypeLeased    LeaseStateType = "leased"
	LeaseStateTypeExpired   LeaseStateType = "expired"
	LeaseStateTypeBreaking  LeaseStateType = "breaking"
	LeaseStateTypeBroken    LeaseStateType = "broken"
)

func PossibleLeaseStateTypeValues() []LeaseStateType {
	return []LeaseStateType{
		LeaseStateTypeAvailable,
		LeaseStateTypeLeased,
		LeaseStateTypeExpired,
		LeaseStateTypeBreaking,
		LeaseStateTypeBroken,
	}
}

func (c LeaseStateType) ToPtr() *LeaseStateType {
	return &c
}

type LeaseStatusType string

const (
	LeaseStatusTypeLocked   LeaseStatusType = "locked"
	LeaseStatusTypeUnlocked LeaseStatusType = "unlocked"
)

func PossibleLeaseStatusTypeValues() []LeaseStatusType {
	return []LeaseStatusType{
		LeaseStatusTypeLocked,
		LeaseStatusTypeUnlocked,
	}
}

func (c LeaseStatusType) ToPtr() *LeaseStatusType {
	return &c
}

type ListBlobsIncludeItem string

const (
	ListBlobsIncludeItemCopy             ListBlobsIncludeItem = "copy"
	ListBlobsIncludeItemDeleted          ListBlobsIncludeItem = "deleted"
	ListBlobsIncludeItemMetadata         ListBlobsIncludeItem = "metadata"
	ListBlobsIncludeItemSnapshots        ListBlobsIncludeItem = "snapshots"
	ListBlobsIncludeItemUncommittedblobs ListBlobsIncludeItem = "uncommittedblobs"
	ListBlobsIncludeItemVersions         ListBlobsIncludeItem = "versions"
	ListBlobsIncludeItemTags             ListBlobsIncludeItem = "tags"
)

func PossibleListBlobsIncludeItemValues() []ListBlobsIncludeItem {
	return []ListBlobsIncludeItem{
		ListBlobsIncludeItemCopy,
		ListBlobsIncludeItemDeleted,
		ListBlobsIncludeItemMetadata,
		ListBlobsIncludeItemSnapshots,
		ListBlobsIncludeItemUncommittedblobs,
		ListBlobsIncludeItemVersions,
		ListBlobsIncludeItemTags,
	}
}

func (c ListBlobsIncludeItem) ToPtr() *ListBlobsIncludeItem {
	return &c
}

type ListContainersIncludeType string

const (
	ListContainersIncludeTypeMetadata ListContainersIncludeType = "metadata"
	ListContainersIncludeTypeDeleted  ListContainersIncludeType = "deleted"
)

func PossibleListContainersIncludeTypeValues() []ListContainersIncludeType {
	return []ListContainersIncludeType{
		ListContainersIncludeTypeMetadata,
		ListContainersIncludeTypeDeleted,
	}
}

func (c ListContainersIncludeType) ToPtr() *ListContainersIncludeType {
	return &c
}

type PathRenameMode string

const (
	PathRenameModeLegacy PathRenameMode = "legacy"
	PathRenameModePosix  PathRenameMode = "posix"
)

func PossiblePathRenameModeValues() []PathRenameMode {
	return []PathRenameMode{
		PathRenameModeLegacy,
		PathRenameModePosix,
	}
}

func (c PathRenameMode) ToPtr() *PathRenameMode {
	return &c
}

type PremiumPageBlobAccessTier string

const (
	PremiumPageBlobAccessTierP10 PremiumPageBlobAccessTier = "P10"
	PremiumPageBlobAccessTierP15 PremiumPageBlobAccessTier = "P15"
	PremiumPageBlobAccessTierP20 PremiumPageBlobAccessTier = "P20"
	PremiumPageBlobAccessTierP30 PremiumPageBlobAccessTier = "P30"
	PremiumPageBlobAccessTierP4  PremiumPageBlobAccessTier = "P4"
	PremiumPageBlobAccessTierP40 PremiumPageBlobAccessTier = "P40"
	PremiumPageBlobAccessTierP50 PremiumPageBlobAccessTier = "P50"
	PremiumPageBlobAccessTierP6  PremiumPageBlobAccessTier = "P6"
	PremiumPageBlobAccessTierP60 PremiumPageBlobAccessTier = "P60"
	PremiumPageBlobAccessTierP70 PremiumPageBlobAccessTier = "P70"
	PremiumPageBlobAccessTierP80 PremiumPageBlobAccessTier = "P80"
)

func PossiblePremiumPageBlobAccessTierValues() []PremiumPageBlobAccessTier {
	return []PremiumPageBlobAccessTier{
		PremiumPageBlobAccessTierP10,
		PremiumPageBlobAccessTierP15,
		PremiumPageBlobAccessTierP20,
		PremiumPageBlobAccessTierP30,
		PremiumPageBlobAccessTierP4,
		PremiumPageBlobAccessTierP40,
		PremiumPageBlobAccessTierP50,
		PremiumPageBlobAccessTierP6,
		PremiumPageBlobAccessTierP60,
		PremiumPageBlobAccessTierP70,
		PremiumPageBlobAccessTierP80,
	}
}

func (c PremiumPageBlobAccessTier) ToPtr() *PremiumPageBlobAccessTier {
	return &c
}

type PublicAccessType string

const (
	PublicAccessTypeBlob      PublicAccessType = "blob"
	PublicAccessTypeContainer PublicAccessType = "container"
)

func PossiblePublicAccessTypeValues() []PublicAccessType {
	return []PublicAccessType{
		PublicAccessTypeBlob,
		PublicAccessTypeContainer,
	}
}

func (c PublicAccessType) ToPtr() *PublicAccessType {
	return &c
}

// QueryFormatType - The quick query format type.
type QueryFormatType string

const (
	QueryFormatTypeDelimited QueryFormatType = "delimited"
	QueryFormatTypeJSON      QueryFormatType = "json"
)

func PossibleQueryFormatTypeValues() []QueryFormatType {
	return []QueryFormatType{
		QueryFormatTypeDelimited,
		QueryFormatTypeJSON,
	}
}

func (c QueryFormatType) ToPtr() *QueryFormatType {
	return &c
}

// RehydratePriority - If an object is in rehydrate pending state then this header is returned with priority of rehydrate. Valid values are High and Standard.
type RehydratePriority string

const (
	RehydratePriorityHigh     RehydratePriority = "High"
	RehydratePriorityStandard RehydratePriority = "Standard"
)

func PossibleRehydratePriorityValues() []RehydratePriority {
	return []RehydratePriority{
		RehydratePriorityHigh,
		RehydratePriorityStandard,
	}
}

func (c RehydratePriority) ToPtr() *RehydratePriority {
	return &c
}

type SKUName string

const (
	SKUNameStandardLrs   SKUName = "Standard_LRS"
	SKUNameStandardGrs   SKUName = "Standard_GRS"
	SKUNameStandardRagrs SKUName = "Standard_RAGRS"
	SKUNameStandardZrs   SKUName = "Standard_ZRS"
	SKUNamePremiumLrs    SKUName = "Premium_LRS"
)

func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameStandardLrs,
		SKUNameStandardGrs,
		SKUNameStandardRagrs,
		SKUNameStandardZrs,
		SKUNamePremiumLrs,
	}
}

func (c SKUName) ToPtr() *SKUName {
	return &c
}

type SequenceNumberActionType string

const (
	SequenceNumberActionTypeMax       SequenceNumberActionType = "max"
	SequenceNumberActionTypeUpdate    SequenceNumberActionType = "update"
	SequenceNumberActionTypeIncrement SequenceNumberActionType = "increment"
)

func PossibleSequenceNumberActionTypeValues() []SequenceNumberActionType {
	return []SequenceNumberActionType{
		SequenceNumberActionTypeMax,
		SequenceNumberActionTypeUpdate,
		SequenceNumberActionTypeIncrement,
	}
}

func (c SequenceNumberActionType) ToPtr() *SequenceNumberActionType {
	return &c
}

// StorageErrorCode - Error codes returned by the service
type StorageErrorCode string

const (
	StorageErrorCodeAccountAlreadyExists                              StorageErrorCode = "AccountAlreadyExists"
	StorageErrorCodeAccountBeingCreated                               StorageErrorCode = "AccountBeingCreated"
	StorageErrorCodeAccountIsDisabled                                 StorageErrorCode = "AccountIsDisabled"
	StorageErrorCodeAppendPositionConditionNotMet                     StorageErrorCode = "AppendPositionConditionNotMet"
	StorageErrorCodeAuthenticationFailed                              StorageErrorCode = "AuthenticationFailed"
	StorageErrorCodeAuthorizationFailure                              StorageErrorCode = "AuthorizationFailure"
	StorageErrorCodeAuthorizationPermissionMismatch                   StorageErrorCode = "AuthorizationPermissionMismatch"
	StorageErrorCodeAuthorizationProtocolMismatch                     StorageErrorCode = "AuthorizationProtocolMismatch"
	StorageErrorCodeAuthorizationResourceTypeMismatch                 StorageErrorCode = "AuthorizationResourceTypeMismatch"
	StorageErrorCodeAuthorizationServiceMismatch                      StorageErrorCode = "AuthorizationServiceMismatch"
	StorageErrorCodeAuthorizationSourceIPMismatch                     StorageErrorCode = "AuthorizationSourceIPMismatch"
	StorageErrorCodeBlobAlreadyExists                                 StorageErrorCode = "BlobAlreadyExists"
	StorageErrorCodeBlobArchived                                      StorageErrorCode = "BlobArchived"
	StorageErrorCodeBlobBeingRehydrated                               StorageErrorCode = "BlobBeingRehydrated"
	StorageErrorCodeBlobImmutableDueToPolicy                          StorageErrorCode = "BlobImmutableDueToPolicy"
	StorageErrorCodeBlobNotArchived                                   StorageErrorCode = "BlobNotArchived"
	StorageErrorCodeBlobNotFound                                      StorageErrorCode = "BlobNotFound"
	StorageErrorCodeBlobOverwritten                                   StorageErrorCode = "BlobOverwritten"
	StorageErrorCodeBlobTierInadequateForContentLength                StorageErrorCode = "BlobTierInadequateForContentLength"
	StorageErrorCodeBlockCountExceedsLimit                            StorageErrorCode = "BlockCountExceedsLimit"
	StorageErrorCodeBlockListTooLong                                  StorageErrorCode = "BlockListTooLong"
	StorageErrorCodeCannotChangeToLowerTier                           StorageErrorCode = "CannotChangeToLowerTier"
	StorageErrorCodeCannotVerifyCopySource                            StorageErrorCode = "CannotVerifyCopySource"
	StorageErrorCodeConditionHeadersNotSupported                      StorageErrorCode = "ConditionHeadersNotSupported"
	StorageErrorCodeConditionNotMet                                   StorageErrorCode = "ConditionNotMet"
	StorageErrorCodeContainerAlreadyExists                            StorageErrorCode = "ContainerAlreadyExists"
	StorageErrorCodeContainerBeingDeleted                             StorageErrorCode = "ContainerBeingDeleted"
	StorageErrorCodeContainerDisabled                                 StorageErrorCode = "ContainerDisabled"
	StorageErrorCodeContainerNotFound                                 StorageErrorCode = "ContainerNotFound"
	StorageErrorCodeContentLengthLargerThanTierLimit                  StorageErrorCode = "ContentLengthLargerThanTierLimit"
	StorageErrorCodeCopyAcrossAccountsNotSupported                    StorageErrorCode = "CopyAcrossAccountsNotSupported"
	StorageErrorCodeCopyIDMismatch                                    StorageErrorCode = "CopyIdMismatch"
	StorageErrorCodeEmptyMetadataKey                                  StorageErrorCode = "EmptyMetadataKey"
	StorageErrorCodeFeatureVersionMismatch                            StorageErrorCode = "FeatureVersionMismatch"
	StorageErrorCodeIncrementalCopyBlobMismatch                       StorageErrorCode = "IncrementalCopyBlobMismatch"
	StorageErrorCodeIncrementalCopyOfEralierVersionSnapshotNotAllowed StorageErrorCode = "IncrementalCopyOfEralierVersionSnapshotNotAllowed"
	StorageErrorCodeIncrementalCopySourceMustBeSnapshot               StorageErrorCode = "IncrementalCopySourceMustBeSnapshot"
	StorageErrorCodeInfiniteLeaseDurationRequired                     StorageErrorCode = "InfiniteLeaseDurationRequired"
	StorageErrorCodeInsufficientAccountPermissions                    StorageErrorCode = "InsufficientAccountPermissions"
	StorageErrorCodeInternalError                                     StorageErrorCode = "InternalError"
	StorageErrorCodeInvalidAuthenticationInfo                         StorageErrorCode = "InvalidAuthenticationInfo"
	StorageErrorCodeInvalidBlobOrBlock                                StorageErrorCode = "InvalidBlobOrBlock"
	StorageErrorCodeInvalidBlobTier                                   StorageErrorCode = "InvalidBlobTier"
	StorageErrorCodeInvalidBlobType                                   StorageErrorCode = "InvalidBlobType"
	StorageErrorCodeInvalidBlockID                                    StorageErrorCode = "InvalidBlockId"
	StorageErrorCodeInvalidBlockList                                  StorageErrorCode = "InvalidBlockList"
	StorageErrorCodeInvalidHTTPVerb                                   StorageErrorCode = "InvalidHttpVerb"
	StorageErrorCodeInvalidHeaderValue                                StorageErrorCode = "InvalidHeaderValue"
	StorageErrorCodeInvalidInput                                      StorageErrorCode = "InvalidInput"
	StorageErrorCodeInvalidMD5                                        StorageErrorCode = "InvalidMd5"
	StorageErrorCodeInvalidMetadata                                   StorageErrorCode = "InvalidMetadata"
	StorageErrorCodeInvalidOperation                                  StorageErrorCode = "InvalidOperation"
	StorageErrorCodeInvalidPageRange                                  StorageErrorCode = "InvalidPageRange"
	StorageErrorCodeInvalidQueryParameterValue                        StorageErrorCode = "InvalidQueryParameterValue"
	StorageErrorCodeInvalidRange                                      StorageErrorCode = "InvalidRange"
	StorageErrorCodeInvalidResourceName                               StorageErrorCode = "InvalidResourceName"
	StorageErrorCodeInvalidSourceBlobType                             StorageErrorCode = "InvalidSourceBlobType"
	StorageErrorCodeInvalidSourceBlobURL                              StorageErrorCode = "InvalidSourceBlobUrl"
	StorageErrorCodeInvalidURI                                        StorageErrorCode = "InvalidUri"
	StorageErrorCodeInvalidVersionForPageBlobOperation                StorageErrorCode = "InvalidVersionForPageBlobOperation"
	StorageErrorCodeInvalidXMLDocument                                StorageErrorCode = "InvalidXmlDocument"
	StorageErrorCodeInvalidXMLNodeValue                               StorageErrorCode = "InvalidXmlNodeValue"
	StorageErrorCodeLeaseAlreadyBroken                                StorageErrorCode = "LeaseAlreadyBroken"
	StorageErrorCodeLeaseAlreadyPresent                               StorageErrorCode = "LeaseAlreadyPresent"
	StorageErrorCodeLeaseIDMismatchWithBlobOperation                  StorageErrorCode = "LeaseIdMismatchWithBlobOperation"
	StorageErrorCodeLeaseIDMismatchWithContainerOperation             StorageErrorCode = "LeaseIdMismatchWithContainerOperation"
	StorageErrorCodeLeaseIDMismatchWithLeaseOperation                 StorageErrorCode = "LeaseIdMismatchWithLeaseOperation"
	StorageErrorCodeLeaseIDMissing                                    StorageErrorCode = "LeaseIdMissing"
	StorageErrorCodeLeaseIsBreakingAndCannotBeAcquired                StorageErrorCode = "LeaseIsBreakingAndCannotBeAcquired"
	StorageErrorCodeLeaseIsBreakingAndCannotBeChanged                 StorageErrorCode = "LeaseIsBreakingAndCannotBeChanged"
	StorageErrorCodeLeaseIsBrokenAndCannotBeRenewed                   StorageErrorCode = "LeaseIsBrokenAndCannotBeRenewed"
	StorageErrorCodeLeaseLost                                         StorageErrorCode = "LeaseLost"
	StorageErrorCodeLeaseNotPresentWithBlobOperation                  StorageErrorCode = "LeaseNotPresentWithBlobOperation"
	StorageErrorCodeLeaseNotPresentWithContainerOperation             StorageErrorCode = "LeaseNotPresentWithContainerOperation"
	StorageErrorCodeLeaseNotPresentWithLeaseOperation                 StorageErrorCode = "LeaseNotPresentWithLeaseOperation"
	StorageErrorCodeMD5Mismatch                                       StorageErrorCode = "Md5Mismatch"
	StorageErrorCodeMaxBlobSizeConditionNotMet                        StorageErrorCode = "MaxBlobSizeConditionNotMet"
	StorageErrorCodeMetadataTooLarge                                  StorageErrorCode = "MetadataTooLarge"
	StorageErrorCodeMissingContentLengthHeader                        StorageErrorCode = "MissingContentLengthHeader"
	StorageErrorCodeMissingRequiredHeader                             StorageErrorCode = "MissingRequiredHeader"
	StorageErrorCodeMissingRequiredQueryParameter                     StorageErrorCode = "MissingRequiredQueryParameter"
	StorageErrorCodeMissingRequiredXMLNode                            StorageErrorCode = "MissingRequiredXmlNode"
	StorageErrorCodeMultipleConditionHeadersNotSupported              StorageErrorCode = "MultipleConditionHeadersNotSupported"
	StorageErrorCodeNoAuthenticationInformation                       StorageErrorCode = "NoAuthenticationInformation"
	StorageErrorCodeNoPendingCopyOperation                            StorageErrorCode = "NoPendingCopyOperation"
	StorageErrorCodeOperationNotAllowedOnIncrementalCopyBlob          StorageErrorCode = "OperationNotAllowedOnIncrementalCopyBlob"
	StorageErrorCodeOperationTimedOut                                 StorageErrorCode = "OperationTimedOut"
	StorageErrorCodeOutOfRangeInput                                   StorageErrorCode = "OutOfRangeInput"
	StorageErrorCodeOutOfRangeQueryParameterValue                     StorageErrorCode = "OutOfRangeQueryParameterValue"
	StorageErrorCodePendingCopyOperation                              StorageErrorCode = "PendingCopyOperation"
	StorageErrorCodePreviousSnapshotCannotBeNewer                     StorageErrorCode = "PreviousSnapshotCannotBeNewer"
	StorageErrorCodePreviousSnapshotNotFound                          StorageErrorCode = "PreviousSnapshotNotFound"
	StorageErrorCodePreviousSnapshotOperationNotSupported             StorageErrorCode = "PreviousSnapshotOperationNotSupported"
	StorageErrorCodeRequestBodyTooLarge                               StorageErrorCode = "RequestBodyTooLarge"
	StorageErrorCodeRequestURLFailedToParse                           StorageErrorCode = "RequestUrlFailedToParse"
	StorageErrorCodeResourceAlreadyExists                             StorageErrorCode = "ResourceAlreadyExists"
	StorageErrorCodeResourceNotFound                                  StorageErrorCode = "ResourceNotFound"
	StorageErrorCodeResourceTypeMismatch                              StorageErrorCode = "ResourceTypeMismatch"
	StorageErrorCodeSequenceNumberConditionNotMet                     StorageErrorCode = "SequenceNumberConditionNotMet"
	StorageErrorCodeSequenceNumberIncrementTooLarge                   StorageErrorCode = "SequenceNumberIncrementTooLarge"
	StorageErrorCodeServerBusy                                        StorageErrorCode = "ServerBusy"
	StorageErrorCodeSnaphotOperationRateExceeded                      StorageErrorCode = "SnaphotOperationRateExceeded"
	StorageErrorCodeSnapshotCountExceeded                             StorageErrorCode = "SnapshotCountExceeded"
	StorageErrorCodeSnapshotsPresent                                  StorageErrorCode = "SnapshotsPresent"
	StorageErrorCodeSourceConditionNotMet                             StorageErrorCode = "SourceConditionNotMet"
	StorageErrorCodeSystemInUse                                       StorageErrorCode = "SystemInUse"
	StorageErrorCodeTargetConditionNotMet                             StorageErrorCode = "TargetConditionNotMet"
	StorageErrorCodeUnauthorizedBlobOverwrite                         StorageErrorCode = "UnauthorizedBlobOverwrite"
	StorageErrorCodeUnsupportedHTTPVerb                               StorageErrorCode = "UnsupportedHttpVerb"
	StorageErrorCodeUnsupportedHeader                                 StorageErrorCode = "UnsupportedHeader"
	StorageErrorCodeUnsupportedQueryParameter                         StorageErrorCode = "UnsupportedQueryParameter"
	StorageErrorCodeUnsupportedXMLNode                                StorageErrorCode = "UnsupportedXmlNode"
)

func PossibleStorageErrorCodeValues() []StorageErrorCode {
	return []StorageErrorCode{
		StorageErrorCodeAccountAlreadyExists,
		StorageErrorCodeAccountBeingCreated,
		StorageErrorCodeAccountIsDisabled,
		StorageErrorCodeAppendPositionConditionNotMet,
		StorageErrorCodeAuthenticationFailed,
		StorageErrorCodeAuthorizationFailure,
		StorageErrorCodeAuthorizationPermissionMismatch,
		StorageErrorCodeAuthorizationProtocolMismatch,
		StorageErrorCodeAuthorizationResourceTypeMismatch,
		StorageErrorCodeAuthorizationServiceMismatch,
		StorageErrorCodeAuthorizationSourceIPMismatch,
		StorageErrorCodeBlobAlreadyExists,
		StorageErrorCodeBlobArchived,
		StorageErrorCodeBlobBeingRehydrated,
		StorageErrorCodeBlobImmutableDueToPolicy,
		StorageErrorCodeBlobNotArchived,
		StorageErrorCodeBlobNotFound,
		StorageErrorCodeBlobOverwritten,
		StorageErrorCodeBlobTierInadequateForContentLength,
		StorageErrorCodeBlockCountExceedsLimit,
		StorageErrorCodeBlockListTooLong,
		StorageErrorCodeCannotChangeToLowerTier,
		StorageErrorCodeCannotVerifyCopySource,
		StorageErrorCodeConditionHeadersNotSupported,
		StorageErrorCodeConditionNotMet,
		StorageErrorCodeContainerAlreadyExists,
		StorageErrorCodeContainerBeingDeleted,
		StorageErrorCodeContainerDisabled,
		StorageErrorCodeContainerNotFound,
		StorageErrorCodeContentLengthLargerThanTierLimit,
		StorageErrorCodeCopyAcrossAccountsNotSupported,
		StorageErrorCodeCopyIDMismatch,
		StorageErrorCodeEmptyMetadataKey,
		StorageErrorCodeFeatureVersionMismatch,
		StorageErrorCodeIncrementalCopyBlobMismatch,
		StorageErrorCodeIncrementalCopyOfEralierVersionSnapshotNotAllowed,
		StorageErrorCodeIncrementalCopySourceMustBeSnapshot,
		StorageErrorCodeInfiniteLeaseDurationRequired,
		StorageErrorCodeInsufficientAccountPermissions,
		StorageErrorCodeInternalError,
		StorageErrorCodeInvalidAuthenticationInfo,
		StorageErrorCodeInvalidBlobOrBlock,
		StorageErrorCodeInvalidBlobTier,
		StorageErrorCodeInvalidBlobType,
		StorageErrorCodeInvalidBlockID,
		StorageErrorCodeInvalidBlockList,
		StorageErrorCodeInvalidHTTPVerb,
		StorageErrorCodeInvalidHeaderValue,
		StorageErrorCodeInvalidInput,
		StorageErrorCodeInvalidMD5,
		StorageErrorCodeInvalidMetadata,
		StorageErrorCodeInvalidOperation,
		StorageErrorCodeInvalidPageRange,
		StorageErrorCodeInvalidQueryParameterValue,
		StorageErrorCodeInvalidRange,
		StorageErrorCodeInvalidResourceName,
		StorageErrorCodeInvalidSourceBlobType,
		StorageErrorCodeInvalidSourceBlobURL,
		StorageErrorCodeInvalidURI,
		StorageErrorCodeInvalidVersionForPageBlobOperation,
		StorageErrorCodeInvalidXMLDocument,
		StorageErrorCodeInvalidXMLNodeValue,
		StorageErrorCodeLeaseAlreadyBroken,
		StorageErrorCodeLeaseAlreadyPresent,
		StorageErrorCodeLeaseIDMismatchWithBlobOperation,
		StorageErrorCodeLeaseIDMismatchWithContainerOperation,
		StorageErrorCodeLeaseIDMismatchWithLeaseOperation,
		StorageErrorCodeLeaseIDMissing,
		StorageErrorCodeLeaseIsBreakingAndCannotBeAcquired,
		StorageErrorCodeLeaseIsBreakingAndCannotBeChanged,
		StorageErrorCodeLeaseIsBrokenAndCannotBeRenewed,
		StorageErrorCodeLeaseLost,
		StorageErrorCodeLeaseNotPresentWithBlobOperation,
		StorageErrorCodeLeaseNotPresentWithContainerOperation,
		StorageErrorCodeLeaseNotPresentWithLeaseOperation,
		StorageErrorCodeMD5Mismatch,
		StorageErrorCodeMaxBlobSizeConditionNotMet,
		StorageErrorCodeMetadataTooLarge,
		StorageErrorCodeMissingContentLengthHeader,
		StorageErrorCodeMissingRequiredHeader,
		StorageErrorCodeMissingRequiredQueryParameter,
		StorageErrorCodeMissingRequiredXMLNode,
		StorageErrorCodeMultipleConditionHeadersNotSupported,
		StorageErrorCodeNoAuthenticationInformation,
		StorageErrorCodeNoPendingCopyOperation,
		StorageErrorCodeOperationNotAllowedOnIncrementalCopyBlob,
		StorageErrorCodeOperationTimedOut,
		StorageErrorCodeOutOfRangeInput,
		StorageErrorCodeOutOfRangeQueryParameterValue,
		StorageErrorCodePendingCopyOperation,
		StorageErrorCodePreviousSnapshotCannotBeNewer,
		StorageErrorCodePreviousSnapshotNotFound,
		StorageErrorCodePreviousSnapshotOperationNotSupported,
		StorageErrorCodeRequestBodyTooLarge,
		StorageErrorCodeRequestURLFailedToParse,
		StorageErrorCodeResourceAlreadyExists,
		StorageErrorCodeResourceNotFound,
		StorageErrorCodeResourceTypeMismatch,
		StorageErrorCodeSequenceNumberConditionNotMet,
		StorageErrorCodeSequenceNumberIncrementTooLarge,
		StorageErrorCodeServerBusy,
		StorageErrorCodeSnaphotOperationRateExceeded,
		StorageErrorCodeSnapshotCountExceeded,
		StorageErrorCodeSnapshotsPresent,
		StorageErrorCodeSourceConditionNotMet,
		StorageErrorCodeSystemInUse,
		StorageErrorCodeTargetConditionNotMet,
		StorageErrorCodeUnauthorizedBlobOverwrite,
		StorageErrorCodeUnsupportedHTTPVerb,
		StorageErrorCodeUnsupportedHeader,
		StorageErrorCodeUnsupportedQueryParameter,
		StorageErrorCodeUnsupportedXMLNode,
	}
}

func (c StorageErrorCode) ToPtr() *StorageErrorCode {
	return &c
}
