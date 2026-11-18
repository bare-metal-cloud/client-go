# GitopsVolumeDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | Pointer to **[]string** | Added / Removed / Remounted name the deltas: a remount is the D1 same-name mount-path change (the claim survives, the app would boot reading an empty directory). | [optional] 
**Current** | Pointer to [**[]GitopsParkVolume**](GitopsParkVolume.md) | Current is the deployed claim set (what the app holds now). | [optional] 
**Proposed** | Pointer to [**[]GitopsParkVolume**](GitopsParkVolume.md) | Proposed is the target&#39;s claim set (this push). | [optional] 
**Remounted** | Pointer to **[]string** |  | [optional] 
**Removed** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGitopsVolumeDelta

`func NewGitopsVolumeDelta() *GitopsVolumeDelta`

NewGitopsVolumeDelta instantiates a new GitopsVolumeDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitopsVolumeDeltaWithDefaults

`func NewGitopsVolumeDeltaWithDefaults() *GitopsVolumeDelta`

NewGitopsVolumeDeltaWithDefaults instantiates a new GitopsVolumeDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *GitopsVolumeDelta) GetAdded() []string`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *GitopsVolumeDelta) GetAddedOk() (*[]string, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *GitopsVolumeDelta) SetAdded(v []string)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *GitopsVolumeDelta) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetCurrent

`func (o *GitopsVolumeDelta) GetCurrent() []GitopsParkVolume`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GitopsVolumeDelta) GetCurrentOk() (*[]GitopsParkVolume, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GitopsVolumeDelta) SetCurrent(v []GitopsParkVolume)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *GitopsVolumeDelta) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetProposed

`func (o *GitopsVolumeDelta) GetProposed() []GitopsParkVolume`

GetProposed returns the Proposed field if non-nil, zero value otherwise.

### GetProposedOk

`func (o *GitopsVolumeDelta) GetProposedOk() (*[]GitopsParkVolume, bool)`

GetProposedOk returns a tuple with the Proposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposed

`func (o *GitopsVolumeDelta) SetProposed(v []GitopsParkVolume)`

SetProposed sets Proposed field to given value.

### HasProposed

`func (o *GitopsVolumeDelta) HasProposed() bool`

HasProposed returns a boolean if a field has been set.

### GetRemounted

`func (o *GitopsVolumeDelta) GetRemounted() []string`

GetRemounted returns the Remounted field if non-nil, zero value otherwise.

### GetRemountedOk

`func (o *GitopsVolumeDelta) GetRemountedOk() (*[]string, bool)`

GetRemountedOk returns a tuple with the Remounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemounted

`func (o *GitopsVolumeDelta) SetRemounted(v []string)`

SetRemounted sets Remounted field to given value.

### HasRemounted

`func (o *GitopsVolumeDelta) HasRemounted() bool`

HasRemounted returns a boolean if a field has been set.

### GetRemoved

`func (o *GitopsVolumeDelta) GetRemoved() []string`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *GitopsVolumeDelta) GetRemovedOk() (*[]string, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *GitopsVolumeDelta) SetRemoved(v []string)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *GitopsVolumeDelta) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


