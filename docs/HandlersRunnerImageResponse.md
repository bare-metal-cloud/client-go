# HandlersRunnerImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**Preinstalled** | Pointer to **[]string** | Preinstalled is the comma-joined column split onto the wire array (the profile labels&#39; transport shape) — last: the pinned fieldalignment analyzer wants the pointer word trailing here (review round 1 verified with the standalone analyzer). | [optional] 
**ReleaseState** | Pointer to **string** | ReleaseState is the closed available|archived vocabulary, derived from the row&#39;s Archived bool (the D1 decision: one fact, one column, the state rendered on the wire). | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerImageResponse

`func NewHandlersRunnerImageResponse() *HandlersRunnerImageResponse`

NewHandlersRunnerImageResponse instantiates a new HandlersRunnerImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerImageResponseWithDefaults

`func NewHandlersRunnerImageResponseWithDefaults() *HandlersRunnerImageResponse`

NewHandlersRunnerImageResponseWithDefaults instantiates a new HandlersRunnerImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *HandlersRunnerImageResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *HandlersRunnerImageResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *HandlersRunnerImageResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *HandlersRunnerImageResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetDigest

`func (o *HandlersRunnerImageResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *HandlersRunnerImageResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *HandlersRunnerImageResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *HandlersRunnerImageResponse) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetImageId

`func (o *HandlersRunnerImageResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *HandlersRunnerImageResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *HandlersRunnerImageResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *HandlersRunnerImageResponse) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetPreinstalled

`func (o *HandlersRunnerImageResponse) GetPreinstalled() []string`

GetPreinstalled returns the Preinstalled field if non-nil, zero value otherwise.

### GetPreinstalledOk

`func (o *HandlersRunnerImageResponse) GetPreinstalledOk() (*[]string, bool)`

GetPreinstalledOk returns a tuple with the Preinstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreinstalled

`func (o *HandlersRunnerImageResponse) SetPreinstalled(v []string)`

SetPreinstalled sets Preinstalled field to given value.

### HasPreinstalled

`func (o *HandlersRunnerImageResponse) HasPreinstalled() bool`

HasPreinstalled returns a boolean if a field has been set.

### GetReleaseState

`func (o *HandlersRunnerImageResponse) GetReleaseState() string`

GetReleaseState returns the ReleaseState field if non-nil, zero value otherwise.

### GetReleaseStateOk

`func (o *HandlersRunnerImageResponse) GetReleaseStateOk() (*string, bool)`

GetReleaseStateOk returns a tuple with the ReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseState

`func (o *HandlersRunnerImageResponse) SetReleaseState(v string)`

SetReleaseState sets ReleaseState field to given value.

### HasReleaseState

`func (o *HandlersRunnerImageResponse) HasReleaseState() bool`

HasReleaseState returns a boolean if a field has been set.

### GetVersion

`func (o *HandlersRunnerImageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersRunnerImageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersRunnerImageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HandlersRunnerImageResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


