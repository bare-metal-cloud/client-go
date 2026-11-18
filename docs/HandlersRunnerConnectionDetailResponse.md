# HandlersRunnerConnectionDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedAt** | Pointer to **string** |  | [optional] 
**GithubOrg** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstallationId** | Pointer to **int32** |  | [optional] 
**ProfilesInUse** | Pointer to **[]string** |  | [optional] 
**RepoCount** | Pointer to **int32** |  | [optional] 
**Repositories** | Pointer to [**[]HandlersRunnerConnectionRepositoryResponse**](HandlersRunnerConnectionRepositoryResponse.md) |  | [optional] 
**RunnerIds** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerConnectionDetailResponse

`func NewHandlersRunnerConnectionDetailResponse() *HandlersRunnerConnectionDetailResponse`

NewHandlersRunnerConnectionDetailResponse instantiates a new HandlersRunnerConnectionDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerConnectionDetailResponseWithDefaults

`func NewHandlersRunnerConnectionDetailResponseWithDefaults() *HandlersRunnerConnectionDetailResponse`

NewHandlersRunnerConnectionDetailResponseWithDefaults instantiates a new HandlersRunnerConnectionDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedAt

`func (o *HandlersRunnerConnectionDetailResponse) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *HandlersRunnerConnectionDetailResponse) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *HandlersRunnerConnectionDetailResponse) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *HandlersRunnerConnectionDetailResponse) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetGithubOrg

`func (o *HandlersRunnerConnectionDetailResponse) GetGithubOrg() string`

GetGithubOrg returns the GithubOrg field if non-nil, zero value otherwise.

### GetGithubOrgOk

`func (o *HandlersRunnerConnectionDetailResponse) GetGithubOrgOk() (*string, bool)`

GetGithubOrgOk returns a tuple with the GithubOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubOrg

`func (o *HandlersRunnerConnectionDetailResponse) SetGithubOrg(v string)`

SetGithubOrg sets GithubOrg field to given value.

### HasGithubOrg

`func (o *HandlersRunnerConnectionDetailResponse) HasGithubOrg() bool`

HasGithubOrg returns a boolean if a field has been set.

### GetId

`func (o *HandlersRunnerConnectionDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersRunnerConnectionDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersRunnerConnectionDetailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersRunnerConnectionDetailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallationId

`func (o *HandlersRunnerConnectionDetailResponse) GetInstallationId() int32`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *HandlersRunnerConnectionDetailResponse) GetInstallationIdOk() (*int32, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *HandlersRunnerConnectionDetailResponse) SetInstallationId(v int32)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *HandlersRunnerConnectionDetailResponse) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetProfilesInUse

`func (o *HandlersRunnerConnectionDetailResponse) GetProfilesInUse() []string`

GetProfilesInUse returns the ProfilesInUse field if non-nil, zero value otherwise.

### GetProfilesInUseOk

`func (o *HandlersRunnerConnectionDetailResponse) GetProfilesInUseOk() (*[]string, bool)`

GetProfilesInUseOk returns a tuple with the ProfilesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesInUse

`func (o *HandlersRunnerConnectionDetailResponse) SetProfilesInUse(v []string)`

SetProfilesInUse sets ProfilesInUse field to given value.

### HasProfilesInUse

`func (o *HandlersRunnerConnectionDetailResponse) HasProfilesInUse() bool`

HasProfilesInUse returns a boolean if a field has been set.

### GetRepoCount

`func (o *HandlersRunnerConnectionDetailResponse) GetRepoCount() int32`

GetRepoCount returns the RepoCount field if non-nil, zero value otherwise.

### GetRepoCountOk

`func (o *HandlersRunnerConnectionDetailResponse) GetRepoCountOk() (*int32, bool)`

GetRepoCountOk returns a tuple with the RepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCount

`func (o *HandlersRunnerConnectionDetailResponse) SetRepoCount(v int32)`

SetRepoCount sets RepoCount field to given value.

### HasRepoCount

`func (o *HandlersRunnerConnectionDetailResponse) HasRepoCount() bool`

HasRepoCount returns a boolean if a field has been set.

### GetRepositories

`func (o *HandlersRunnerConnectionDetailResponse) GetRepositories() []HandlersRunnerConnectionRepositoryResponse`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *HandlersRunnerConnectionDetailResponse) GetRepositoriesOk() (*[]HandlersRunnerConnectionRepositoryResponse, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *HandlersRunnerConnectionDetailResponse) SetRepositories(v []HandlersRunnerConnectionRepositoryResponse)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *HandlersRunnerConnectionDetailResponse) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetRunnerIds

`func (o *HandlersRunnerConnectionDetailResponse) GetRunnerIds() []string`

GetRunnerIds returns the RunnerIds field if non-nil, zero value otherwise.

### GetRunnerIdsOk

`func (o *HandlersRunnerConnectionDetailResponse) GetRunnerIdsOk() (*[]string, bool)`

GetRunnerIdsOk returns a tuple with the RunnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerIds

`func (o *HandlersRunnerConnectionDetailResponse) SetRunnerIds(v []string)`

SetRunnerIds sets RunnerIds field to given value.

### HasRunnerIds

`func (o *HandlersRunnerConnectionDetailResponse) HasRunnerIds() bool`

HasRunnerIds returns a boolean if a field has been set.

### GetState

`func (o *HandlersRunnerConnectionDetailResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersRunnerConnectionDetailResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersRunnerConnectionDetailResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersRunnerConnectionDetailResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


