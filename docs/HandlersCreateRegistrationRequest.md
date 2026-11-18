# HandlersCreateRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **int32** |  | [optional] 
**HealthPath** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RegistryHost** | Pointer to **string** |  | [optional] 
**RegistryId** | Pointer to **int32** |  | [optional] 
**RegistryRepository** | Pointer to **string** |  | [optional] 
**RepoName** | Pointer to **string** |  | [optional] 
**RepoOwner** | Pointer to **string** |  | [optional] 
**Worker** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersCreateRegistrationRequest

`func NewHandlersCreateRegistrationRequest() *HandlersCreateRegistrationRequest`

NewHandlersCreateRegistrationRequest instantiates a new HandlersCreateRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateRegistrationRequestWithDefaults

`func NewHandlersCreateRegistrationRequestWithDefaults() *HandlersCreateRegistrationRequest`

NewHandlersCreateRegistrationRequestWithDefaults instantiates a new HandlersCreateRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *HandlersCreateRegistrationRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *HandlersCreateRegistrationRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *HandlersCreateRegistrationRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *HandlersCreateRegistrationRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConnectionId

`func (o *HandlersCreateRegistrationRequest) GetConnectionId() int32`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HandlersCreateRegistrationRequest) GetConnectionIdOk() (*int32, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HandlersCreateRegistrationRequest) SetConnectionId(v int32)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *HandlersCreateRegistrationRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetHealthPath

`func (o *HandlersCreateRegistrationRequest) GetHealthPath() string`

GetHealthPath returns the HealthPath field if non-nil, zero value otherwise.

### GetHealthPathOk

`func (o *HandlersCreateRegistrationRequest) GetHealthPathOk() (*string, bool)`

GetHealthPathOk returns a tuple with the HealthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthPath

`func (o *HandlersCreateRegistrationRequest) SetHealthPath(v string)`

SetHealthPath sets HealthPath field to given value.

### HasHealthPath

`func (o *HandlersCreateRegistrationRequest) HasHealthPath() bool`

HasHealthPath returns a boolean if a field has been set.

### GetHost

`func (o *HandlersCreateRegistrationRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HandlersCreateRegistrationRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HandlersCreateRegistrationRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HandlersCreateRegistrationRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *HandlersCreateRegistrationRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersCreateRegistrationRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersCreateRegistrationRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersCreateRegistrationRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRegistryHost

`func (o *HandlersCreateRegistrationRequest) GetRegistryHost() string`

GetRegistryHost returns the RegistryHost field if non-nil, zero value otherwise.

### GetRegistryHostOk

`func (o *HandlersCreateRegistrationRequest) GetRegistryHostOk() (*string, bool)`

GetRegistryHostOk returns a tuple with the RegistryHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryHost

`func (o *HandlersCreateRegistrationRequest) SetRegistryHost(v string)`

SetRegistryHost sets RegistryHost field to given value.

### HasRegistryHost

`func (o *HandlersCreateRegistrationRequest) HasRegistryHost() bool`

HasRegistryHost returns a boolean if a field has been set.

### GetRegistryId

`func (o *HandlersCreateRegistrationRequest) GetRegistryId() int32`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *HandlersCreateRegistrationRequest) GetRegistryIdOk() (*int32, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *HandlersCreateRegistrationRequest) SetRegistryId(v int32)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *HandlersCreateRegistrationRequest) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetRegistryRepository

`func (o *HandlersCreateRegistrationRequest) GetRegistryRepository() string`

GetRegistryRepository returns the RegistryRepository field if non-nil, zero value otherwise.

### GetRegistryRepositoryOk

`func (o *HandlersCreateRegistrationRequest) GetRegistryRepositoryOk() (*string, bool)`

GetRegistryRepositoryOk returns a tuple with the RegistryRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryRepository

`func (o *HandlersCreateRegistrationRequest) SetRegistryRepository(v string)`

SetRegistryRepository sets RegistryRepository field to given value.

### HasRegistryRepository

`func (o *HandlersCreateRegistrationRequest) HasRegistryRepository() bool`

HasRegistryRepository returns a boolean if a field has been set.

### GetRepoName

`func (o *HandlersCreateRegistrationRequest) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *HandlersCreateRegistrationRequest) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *HandlersCreateRegistrationRequest) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.

### HasRepoName

`func (o *HandlersCreateRegistrationRequest) HasRepoName() bool`

HasRepoName returns a boolean if a field has been set.

### GetRepoOwner

`func (o *HandlersCreateRegistrationRequest) GetRepoOwner() string`

GetRepoOwner returns the RepoOwner field if non-nil, zero value otherwise.

### GetRepoOwnerOk

`func (o *HandlersCreateRegistrationRequest) GetRepoOwnerOk() (*string, bool)`

GetRepoOwnerOk returns a tuple with the RepoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoOwner

`func (o *HandlersCreateRegistrationRequest) SetRepoOwner(v string)`

SetRepoOwner sets RepoOwner field to given value.

### HasRepoOwner

`func (o *HandlersCreateRegistrationRequest) HasRepoOwner() bool`

HasRepoOwner returns a boolean if a field has been set.

### GetWorker

`func (o *HandlersCreateRegistrationRequest) GetWorker() bool`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *HandlersCreateRegistrationRequest) GetWorkerOk() (*bool, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *HandlersCreateRegistrationRequest) SetWorker(v bool)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *HandlersCreateRegistrationRequest) HasWorker() bool`

HasWorker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


