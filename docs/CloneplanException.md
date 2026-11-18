# CloneplanException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail is the human-readable fact (source value beside the approved target value for env_local items, the honest replay-empty wording). | [optional] 
**Install** | Pointer to **string** | Install is the item&#39;s install name (a stack unit&#39;s is its parent). | [optional] 
**Kind** | Pointer to **string** | Kind is the exception kind: excluded | env-local | endpoint | replay-empty | undeclared-class. | [optional] 

## Methods

### NewCloneplanException

`func NewCloneplanException() *CloneplanException`

NewCloneplanException instantiates a new CloneplanException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanExceptionWithDefaults

`func NewCloneplanExceptionWithDefaults() *CloneplanException`

NewCloneplanExceptionWithDefaults instantiates a new CloneplanException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *CloneplanException) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CloneplanException) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CloneplanException) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CloneplanException) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstall

`func (o *CloneplanException) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *CloneplanException) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *CloneplanException) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *CloneplanException) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetKind

`func (o *CloneplanException) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloneplanException) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloneplanException) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloneplanException) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


