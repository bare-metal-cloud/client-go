# HandlersVaultRotationHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobsOnOldVersion** | Pointer to **int32** |  | [optional] 
**CurrentActiveVersion** | Pointer to **int32** |  | [optional] 
**DualKeyWindowActive** | Pointer to **bool** |  | [optional] 
**LastRotationTimestamp** | Pointer to **string** |  | [optional] 
**RegisteredVersions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewHandlersVaultRotationHealthResponse

`func NewHandlersVaultRotationHealthResponse() *HandlersVaultRotationHealthResponse`

NewHandlersVaultRotationHealthResponse instantiates a new HandlersVaultRotationHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVaultRotationHealthResponseWithDefaults

`func NewHandlersVaultRotationHealthResponseWithDefaults() *HandlersVaultRotationHealthResponse`

NewHandlersVaultRotationHealthResponseWithDefaults instantiates a new HandlersVaultRotationHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobsOnOldVersion

`func (o *HandlersVaultRotationHealthResponse) GetBlobsOnOldVersion() int32`

GetBlobsOnOldVersion returns the BlobsOnOldVersion field if non-nil, zero value otherwise.

### GetBlobsOnOldVersionOk

`func (o *HandlersVaultRotationHealthResponse) GetBlobsOnOldVersionOk() (*int32, bool)`

GetBlobsOnOldVersionOk returns a tuple with the BlobsOnOldVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsOnOldVersion

`func (o *HandlersVaultRotationHealthResponse) SetBlobsOnOldVersion(v int32)`

SetBlobsOnOldVersion sets BlobsOnOldVersion field to given value.

### HasBlobsOnOldVersion

`func (o *HandlersVaultRotationHealthResponse) HasBlobsOnOldVersion() bool`

HasBlobsOnOldVersion returns a boolean if a field has been set.

### GetCurrentActiveVersion

`func (o *HandlersVaultRotationHealthResponse) GetCurrentActiveVersion() int32`

GetCurrentActiveVersion returns the CurrentActiveVersion field if non-nil, zero value otherwise.

### GetCurrentActiveVersionOk

`func (o *HandlersVaultRotationHealthResponse) GetCurrentActiveVersionOk() (*int32, bool)`

GetCurrentActiveVersionOk returns a tuple with the CurrentActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveVersion

`func (o *HandlersVaultRotationHealthResponse) SetCurrentActiveVersion(v int32)`

SetCurrentActiveVersion sets CurrentActiveVersion field to given value.

### HasCurrentActiveVersion

`func (o *HandlersVaultRotationHealthResponse) HasCurrentActiveVersion() bool`

HasCurrentActiveVersion returns a boolean if a field has been set.

### GetDualKeyWindowActive

`func (o *HandlersVaultRotationHealthResponse) GetDualKeyWindowActive() bool`

GetDualKeyWindowActive returns the DualKeyWindowActive field if non-nil, zero value otherwise.

### GetDualKeyWindowActiveOk

`func (o *HandlersVaultRotationHealthResponse) GetDualKeyWindowActiveOk() (*bool, bool)`

GetDualKeyWindowActiveOk returns a tuple with the DualKeyWindowActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualKeyWindowActive

`func (o *HandlersVaultRotationHealthResponse) SetDualKeyWindowActive(v bool)`

SetDualKeyWindowActive sets DualKeyWindowActive field to given value.

### HasDualKeyWindowActive

`func (o *HandlersVaultRotationHealthResponse) HasDualKeyWindowActive() bool`

HasDualKeyWindowActive returns a boolean if a field has been set.

### GetLastRotationTimestamp

`func (o *HandlersVaultRotationHealthResponse) GetLastRotationTimestamp() string`

GetLastRotationTimestamp returns the LastRotationTimestamp field if non-nil, zero value otherwise.

### GetLastRotationTimestampOk

`func (o *HandlersVaultRotationHealthResponse) GetLastRotationTimestampOk() (*string, bool)`

GetLastRotationTimestampOk returns a tuple with the LastRotationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotationTimestamp

`func (o *HandlersVaultRotationHealthResponse) SetLastRotationTimestamp(v string)`

SetLastRotationTimestamp sets LastRotationTimestamp field to given value.

### HasLastRotationTimestamp

`func (o *HandlersVaultRotationHealthResponse) HasLastRotationTimestamp() bool`

HasLastRotationTimestamp returns a boolean if a field has been set.

### GetRegisteredVersions

`func (o *HandlersVaultRotationHealthResponse) GetRegisteredVersions() []int32`

GetRegisteredVersions returns the RegisteredVersions field if non-nil, zero value otherwise.

### GetRegisteredVersionsOk

`func (o *HandlersVaultRotationHealthResponse) GetRegisteredVersionsOk() (*[]int32, bool)`

GetRegisteredVersionsOk returns a tuple with the RegisteredVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredVersions

`func (o *HandlersVaultRotationHealthResponse) SetRegisteredVersions(v []int32)`

SetRegisteredVersions sets RegisteredVersions field to given value.

### HasRegisteredVersions

`func (o *HandlersVaultRotationHealthResponse) HasRegisteredVersions() bool`

HasRegisteredVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


