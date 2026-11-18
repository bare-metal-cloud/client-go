# HandlersRunnerBuildDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to **string** |  | [optional] 
**BuildRef** | Pointer to **string** |  | [optional] 
**BuilderImage** | Pointer to **string** |  | [optional] 
**EgressAllowlist** | Pointer to **[]string** |  | [optional] 
**Install** | Pointer to **string** |  | [optional] 
**LogTopicId** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PushRepository** | Pointer to **string** |  | [optional] 
**PushTags** | Pointer to **[]string** |  | [optional] 
**RegistryPushCredential** | Pointer to **string** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**SourceCommit** | Pointer to **string** |  | [optional] 
**SourcePullCredential** | Pointer to **string** |  | [optional] 
**SourceRepo** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersRunnerBuildDispatchRequest

`func NewHandlersRunnerBuildDispatchRequest() *HandlersRunnerBuildDispatchRequest`

NewHandlersRunnerBuildDispatchRequest instantiates a new HandlersRunnerBuildDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerBuildDispatchRequestWithDefaults

`func NewHandlersRunnerBuildDispatchRequestWithDefaults() *HandlersRunnerBuildDispatchRequest`

NewHandlersRunnerBuildDispatchRequestWithDefaults instantiates a new HandlersRunnerBuildDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *HandlersRunnerBuildDispatchRequest) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *HandlersRunnerBuildDispatchRequest) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *HandlersRunnerBuildDispatchRequest) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *HandlersRunnerBuildDispatchRequest) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuildRef

`func (o *HandlersRunnerBuildDispatchRequest) GetBuildRef() string`

GetBuildRef returns the BuildRef field if non-nil, zero value otherwise.

### GetBuildRefOk

`func (o *HandlersRunnerBuildDispatchRequest) GetBuildRefOk() (*string, bool)`

GetBuildRefOk returns a tuple with the BuildRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildRef

`func (o *HandlersRunnerBuildDispatchRequest) SetBuildRef(v string)`

SetBuildRef sets BuildRef field to given value.

### HasBuildRef

`func (o *HandlersRunnerBuildDispatchRequest) HasBuildRef() bool`

HasBuildRef returns a boolean if a field has been set.

### GetBuilderImage

`func (o *HandlersRunnerBuildDispatchRequest) GetBuilderImage() string`

GetBuilderImage returns the BuilderImage field if non-nil, zero value otherwise.

### GetBuilderImageOk

`func (o *HandlersRunnerBuildDispatchRequest) GetBuilderImageOk() (*string, bool)`

GetBuilderImageOk returns a tuple with the BuilderImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderImage

`func (o *HandlersRunnerBuildDispatchRequest) SetBuilderImage(v string)`

SetBuilderImage sets BuilderImage field to given value.

### HasBuilderImage

`func (o *HandlersRunnerBuildDispatchRequest) HasBuilderImage() bool`

HasBuilderImage returns a boolean if a field has been set.

### GetEgressAllowlist

`func (o *HandlersRunnerBuildDispatchRequest) GetEgressAllowlist() []string`

GetEgressAllowlist returns the EgressAllowlist field if non-nil, zero value otherwise.

### GetEgressAllowlistOk

`func (o *HandlersRunnerBuildDispatchRequest) GetEgressAllowlistOk() (*[]string, bool)`

GetEgressAllowlistOk returns a tuple with the EgressAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressAllowlist

`func (o *HandlersRunnerBuildDispatchRequest) SetEgressAllowlist(v []string)`

SetEgressAllowlist sets EgressAllowlist field to given value.

### HasEgressAllowlist

`func (o *HandlersRunnerBuildDispatchRequest) HasEgressAllowlist() bool`

HasEgressAllowlist returns a boolean if a field has been set.

### GetInstall

`func (o *HandlersRunnerBuildDispatchRequest) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HandlersRunnerBuildDispatchRequest) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HandlersRunnerBuildDispatchRequest) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *HandlersRunnerBuildDispatchRequest) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetLogTopicId

`func (o *HandlersRunnerBuildDispatchRequest) GetLogTopicId() string`

GetLogTopicId returns the LogTopicId field if non-nil, zero value otherwise.

### GetLogTopicIdOk

`func (o *HandlersRunnerBuildDispatchRequest) GetLogTopicIdOk() (*string, bool)`

GetLogTopicIdOk returns a tuple with the LogTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTopicId

`func (o *HandlersRunnerBuildDispatchRequest) SetLogTopicId(v string)`

SetLogTopicId sets LogTopicId field to given value.

### HasLogTopicId

`func (o *HandlersRunnerBuildDispatchRequest) HasLogTopicId() bool`

HasLogTopicId returns a boolean if a field has been set.

### GetPort

`func (o *HandlersRunnerBuildDispatchRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersRunnerBuildDispatchRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersRunnerBuildDispatchRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersRunnerBuildDispatchRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPushRepository

`func (o *HandlersRunnerBuildDispatchRequest) GetPushRepository() string`

GetPushRepository returns the PushRepository field if non-nil, zero value otherwise.

### GetPushRepositoryOk

`func (o *HandlersRunnerBuildDispatchRequest) GetPushRepositoryOk() (*string, bool)`

GetPushRepositoryOk returns a tuple with the PushRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushRepository

`func (o *HandlersRunnerBuildDispatchRequest) SetPushRepository(v string)`

SetPushRepository sets PushRepository field to given value.

### HasPushRepository

`func (o *HandlersRunnerBuildDispatchRequest) HasPushRepository() bool`

HasPushRepository returns a boolean if a field has been set.

### GetPushTags

`func (o *HandlersRunnerBuildDispatchRequest) GetPushTags() []string`

GetPushTags returns the PushTags field if non-nil, zero value otherwise.

### GetPushTagsOk

`func (o *HandlersRunnerBuildDispatchRequest) GetPushTagsOk() (*[]string, bool)`

GetPushTagsOk returns a tuple with the PushTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTags

`func (o *HandlersRunnerBuildDispatchRequest) SetPushTags(v []string)`

SetPushTags sets PushTags field to given value.

### HasPushTags

`func (o *HandlersRunnerBuildDispatchRequest) HasPushTags() bool`

HasPushTags returns a boolean if a field has been set.

### GetRegistryPushCredential

`func (o *HandlersRunnerBuildDispatchRequest) GetRegistryPushCredential() string`

GetRegistryPushCredential returns the RegistryPushCredential field if non-nil, zero value otherwise.

### GetRegistryPushCredentialOk

`func (o *HandlersRunnerBuildDispatchRequest) GetRegistryPushCredentialOk() (*string, bool)`

GetRegistryPushCredentialOk returns a tuple with the RegistryPushCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryPushCredential

`func (o *HandlersRunnerBuildDispatchRequest) SetRegistryPushCredential(v string)`

SetRegistryPushCredential sets RegistryPushCredential field to given value.

### HasRegistryPushCredential

`func (o *HandlersRunnerBuildDispatchRequest) HasRegistryPushCredential() bool`

HasRegistryPushCredential returns a boolean if a field has been set.

### GetRuntime

`func (o *HandlersRunnerBuildDispatchRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *HandlersRunnerBuildDispatchRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *HandlersRunnerBuildDispatchRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *HandlersRunnerBuildDispatchRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetSourceCommit

`func (o *HandlersRunnerBuildDispatchRequest) GetSourceCommit() string`

GetSourceCommit returns the SourceCommit field if non-nil, zero value otherwise.

### GetSourceCommitOk

`func (o *HandlersRunnerBuildDispatchRequest) GetSourceCommitOk() (*string, bool)`

GetSourceCommitOk returns a tuple with the SourceCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommit

`func (o *HandlersRunnerBuildDispatchRequest) SetSourceCommit(v string)`

SetSourceCommit sets SourceCommit field to given value.

### HasSourceCommit

`func (o *HandlersRunnerBuildDispatchRequest) HasSourceCommit() bool`

HasSourceCommit returns a boolean if a field has been set.

### GetSourcePullCredential

`func (o *HandlersRunnerBuildDispatchRequest) GetSourcePullCredential() string`

GetSourcePullCredential returns the SourcePullCredential field if non-nil, zero value otherwise.

### GetSourcePullCredentialOk

`func (o *HandlersRunnerBuildDispatchRequest) GetSourcePullCredentialOk() (*string, bool)`

GetSourcePullCredentialOk returns a tuple with the SourcePullCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePullCredential

`func (o *HandlersRunnerBuildDispatchRequest) SetSourcePullCredential(v string)`

SetSourcePullCredential sets SourcePullCredential field to given value.

### HasSourcePullCredential

`func (o *HandlersRunnerBuildDispatchRequest) HasSourcePullCredential() bool`

HasSourcePullCredential returns a boolean if a field has been set.

### GetSourceRepo

`func (o *HandlersRunnerBuildDispatchRequest) GetSourceRepo() string`

GetSourceRepo returns the SourceRepo field if non-nil, zero value otherwise.

### GetSourceRepoOk

`func (o *HandlersRunnerBuildDispatchRequest) GetSourceRepoOk() (*string, bool)`

GetSourceRepoOk returns a tuple with the SourceRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepo

`func (o *HandlersRunnerBuildDispatchRequest) SetSourceRepo(v string)`

SetSourceRepo sets SourceRepo field to given value.

### HasSourceRepo

`func (o *HandlersRunnerBuildDispatchRequest) HasSourceRepo() bool`

HasSourceRepo returns a boolean if a field has been set.

### GetStart

`func (o *HandlersRunnerBuildDispatchRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HandlersRunnerBuildDispatchRequest) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HandlersRunnerBuildDispatchRequest) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *HandlersRunnerBuildDispatchRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStrategy

`func (o *HandlersRunnerBuildDispatchRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *HandlersRunnerBuildDispatchRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *HandlersRunnerBuildDispatchRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *HandlersRunnerBuildDispatchRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *HandlersRunnerBuildDispatchRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *HandlersRunnerBuildDispatchRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *HandlersRunnerBuildDispatchRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *HandlersRunnerBuildDispatchRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


