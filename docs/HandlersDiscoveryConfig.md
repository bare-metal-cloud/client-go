# HandlersDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckActive** | Pointer to **bool** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Executable** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**ParseRegex** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**SuccessCodes** | Pointer to **[]int32** |  | [optional] 
**Timeout** | Pointer to **int32** | in seconds | [optional] 
**Type** | Pointer to **string** | command, file, service, package, http | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDiscoveryConfig

`func NewHandlersDiscoveryConfig() *HandlersDiscoveryConfig`

NewHandlersDiscoveryConfig instantiates a new HandlersDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDiscoveryConfigWithDefaults

`func NewHandlersDiscoveryConfigWithDefaults() *HandlersDiscoveryConfig`

NewHandlersDiscoveryConfigWithDefaults instantiates a new HandlersDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckActive

`func (o *HandlersDiscoveryConfig) GetCheckActive() bool`

GetCheckActive returns the CheckActive field if non-nil, zero value otherwise.

### GetCheckActiveOk

`func (o *HandlersDiscoveryConfig) GetCheckActiveOk() (*bool, bool)`

GetCheckActiveOk returns a tuple with the CheckActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckActive

`func (o *HandlersDiscoveryConfig) SetCheckActive(v bool)`

SetCheckActive sets CheckActive field to given value.

### HasCheckActive

`func (o *HandlersDiscoveryConfig) HasCheckActive() bool`

HasCheckActive returns a boolean if a field has been set.

### GetCommand

`func (o *HandlersDiscoveryConfig) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *HandlersDiscoveryConfig) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *HandlersDiscoveryConfig) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *HandlersDiscoveryConfig) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetExecutable

`func (o *HandlersDiscoveryConfig) GetExecutable() bool`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *HandlersDiscoveryConfig) GetExecutableOk() (*bool, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *HandlersDiscoveryConfig) SetExecutable(v bool)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *HandlersDiscoveryConfig) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetMetadata

`func (o *HandlersDiscoveryConfig) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HandlersDiscoveryConfig) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HandlersDiscoveryConfig) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HandlersDiscoveryConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *HandlersDiscoveryConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersDiscoveryConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersDiscoveryConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersDiscoveryConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageName

`func (o *HandlersDiscoveryConfig) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *HandlersDiscoveryConfig) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *HandlersDiscoveryConfig) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *HandlersDiscoveryConfig) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetParseRegex

`func (o *HandlersDiscoveryConfig) GetParseRegex() string`

GetParseRegex returns the ParseRegex field if non-nil, zero value otherwise.

### GetParseRegexOk

`func (o *HandlersDiscoveryConfig) GetParseRegexOk() (*string, bool)`

GetParseRegexOk returns a tuple with the ParseRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseRegex

`func (o *HandlersDiscoveryConfig) SetParseRegex(v string)`

SetParseRegex sets ParseRegex field to given value.

### HasParseRegex

`func (o *HandlersDiscoveryConfig) HasParseRegex() bool`

HasParseRegex returns a boolean if a field has been set.

### GetPath

`func (o *HandlersDiscoveryConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HandlersDiscoveryConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HandlersDiscoveryConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HandlersDiscoveryConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSuccessCodes

`func (o *HandlersDiscoveryConfig) GetSuccessCodes() []int32`

GetSuccessCodes returns the SuccessCodes field if non-nil, zero value otherwise.

### GetSuccessCodesOk

`func (o *HandlersDiscoveryConfig) GetSuccessCodesOk() (*[]int32, bool)`

GetSuccessCodesOk returns a tuple with the SuccessCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCodes

`func (o *HandlersDiscoveryConfig) SetSuccessCodes(v []int32)`

SetSuccessCodes sets SuccessCodes field to given value.

### HasSuccessCodes

`func (o *HandlersDiscoveryConfig) HasSuccessCodes() bool`

HasSuccessCodes returns a boolean if a field has been set.

### GetTimeout

`func (o *HandlersDiscoveryConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HandlersDiscoveryConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HandlersDiscoveryConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HandlersDiscoveryConfig) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *HandlersDiscoveryConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlersDiscoveryConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlersDiscoveryConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HandlersDiscoveryConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *HandlersDiscoveryConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HandlersDiscoveryConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HandlersDiscoveryConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HandlersDiscoveryConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


