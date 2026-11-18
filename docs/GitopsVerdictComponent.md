# GitopsVerdictComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimRetirement** | Pointer to **bool** |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**FromDataFormatVersion** | Pointer to **int32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Migrating** | Pointer to **bool** |  | [optional] 
**NewlyStateful** | Pointer to **bool** |  | [optional] 
**ToDataFormatVersion** | Pointer to **int32** |  | [optional] 
**VolumeReview** | Pointer to **bool** |  | [optional] 

## Methods

### NewGitopsVerdictComponent

`func NewGitopsVerdictComponent() *GitopsVerdictComponent`

NewGitopsVerdictComponent instantiates a new GitopsVerdictComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitopsVerdictComponentWithDefaults

`func NewGitopsVerdictComponentWithDefaults() *GitopsVerdictComponent`

NewGitopsVerdictComponentWithDefaults instantiates a new GitopsVerdictComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimRetirement

`func (o *GitopsVerdictComponent) GetClaimRetirement() bool`

GetClaimRetirement returns the ClaimRetirement field if non-nil, zero value otherwise.

### GetClaimRetirementOk

`func (o *GitopsVerdictComponent) GetClaimRetirementOk() (*bool, bool)`

GetClaimRetirementOk returns a tuple with the ClaimRetirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRetirement

`func (o *GitopsVerdictComponent) SetClaimRetirement(v bool)`

SetClaimRetirement sets ClaimRetirement field to given value.

### HasClaimRetirement

`func (o *GitopsVerdictComponent) HasClaimRetirement() bool`

HasClaimRetirement returns a boolean if a field has been set.

### GetComponent

`func (o *GitopsVerdictComponent) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GitopsVerdictComponent) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GitopsVerdictComponent) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *GitopsVerdictComponent) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFromDataFormatVersion

`func (o *GitopsVerdictComponent) GetFromDataFormatVersion() int32`

GetFromDataFormatVersion returns the FromDataFormatVersion field if non-nil, zero value otherwise.

### GetFromDataFormatVersionOk

`func (o *GitopsVerdictComponent) GetFromDataFormatVersionOk() (*int32, bool)`

GetFromDataFormatVersionOk returns a tuple with the FromDataFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDataFormatVersion

`func (o *GitopsVerdictComponent) SetFromDataFormatVersion(v int32)`

SetFromDataFormatVersion sets FromDataFormatVersion field to given value.

### HasFromDataFormatVersion

`func (o *GitopsVerdictComponent) HasFromDataFormatVersion() bool`

HasFromDataFormatVersion returns a boolean if a field has been set.

### GetKind

`func (o *GitopsVerdictComponent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GitopsVerdictComponent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GitopsVerdictComponent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GitopsVerdictComponent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMigrating

`func (o *GitopsVerdictComponent) GetMigrating() bool`

GetMigrating returns the Migrating field if non-nil, zero value otherwise.

### GetMigratingOk

`func (o *GitopsVerdictComponent) GetMigratingOk() (*bool, bool)`

GetMigratingOk returns a tuple with the Migrating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrating

`func (o *GitopsVerdictComponent) SetMigrating(v bool)`

SetMigrating sets Migrating field to given value.

### HasMigrating

`func (o *GitopsVerdictComponent) HasMigrating() bool`

HasMigrating returns a boolean if a field has been set.

### GetNewlyStateful

`func (o *GitopsVerdictComponent) GetNewlyStateful() bool`

GetNewlyStateful returns the NewlyStateful field if non-nil, zero value otherwise.

### GetNewlyStatefulOk

`func (o *GitopsVerdictComponent) GetNewlyStatefulOk() (*bool, bool)`

GetNewlyStatefulOk returns a tuple with the NewlyStateful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewlyStateful

`func (o *GitopsVerdictComponent) SetNewlyStateful(v bool)`

SetNewlyStateful sets NewlyStateful field to given value.

### HasNewlyStateful

`func (o *GitopsVerdictComponent) HasNewlyStateful() bool`

HasNewlyStateful returns a boolean if a field has been set.

### GetToDataFormatVersion

`func (o *GitopsVerdictComponent) GetToDataFormatVersion() int32`

GetToDataFormatVersion returns the ToDataFormatVersion field if non-nil, zero value otherwise.

### GetToDataFormatVersionOk

`func (o *GitopsVerdictComponent) GetToDataFormatVersionOk() (*int32, bool)`

GetToDataFormatVersionOk returns a tuple with the ToDataFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDataFormatVersion

`func (o *GitopsVerdictComponent) SetToDataFormatVersion(v int32)`

SetToDataFormatVersion sets ToDataFormatVersion field to given value.

### HasToDataFormatVersion

`func (o *GitopsVerdictComponent) HasToDataFormatVersion() bool`

HasToDataFormatVersion returns a boolean if a field has been set.

### GetVolumeReview

`func (o *GitopsVerdictComponent) GetVolumeReview() bool`

GetVolumeReview returns the VolumeReview field if non-nil, zero value otherwise.

### GetVolumeReviewOk

`func (o *GitopsVerdictComponent) GetVolumeReviewOk() (*bool, bool)`

GetVolumeReviewOk returns a tuple with the VolumeReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeReview

`func (o *GitopsVerdictComponent) SetVolumeReview(v bool)`

SetVolumeReview sets VolumeReview field to given value.

### HasVolumeReview

`func (o *GitopsVerdictComponent) HasVolumeReview() bool`

HasVolumeReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


