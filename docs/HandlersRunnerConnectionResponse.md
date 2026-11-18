# HandlersRunnerConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedAt** | Pointer to **string** | ConnectedAt is the installation row&#39;s creation, RFC3339. | [optional] 
**GithubOrg** | Pointer to **string** | GithubOrg is the installation&#39;s GitHub account login. | [optional] 
**Id** | Pointer to **string** | ID is the installation ROW id (the connection id the client hands back on the detail/save paths), string on the wire. | [optional] 
**InstallationId** | Pointer to **int32** | InstallationID is GitHub&#39;s id for the installation. | [optional] 
**RepoCount** | Pointer to **int32** | RepoCount is the installation&#39;s mapping rows (its repositories with a saved profile answer). | [optional] 
**State** | Pointer to **string** | State is the derived connection state (the vocabulary above). | [optional] 

## Methods

### NewHandlersRunnerConnectionResponse

`func NewHandlersRunnerConnectionResponse() *HandlersRunnerConnectionResponse`

NewHandlersRunnerConnectionResponse instantiates a new HandlersRunnerConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerConnectionResponseWithDefaults

`func NewHandlersRunnerConnectionResponseWithDefaults() *HandlersRunnerConnectionResponse`

NewHandlersRunnerConnectionResponseWithDefaults instantiates a new HandlersRunnerConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedAt

`func (o *HandlersRunnerConnectionResponse) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *HandlersRunnerConnectionResponse) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *HandlersRunnerConnectionResponse) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *HandlersRunnerConnectionResponse) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetGithubOrg

`func (o *HandlersRunnerConnectionResponse) GetGithubOrg() string`

GetGithubOrg returns the GithubOrg field if non-nil, zero value otherwise.

### GetGithubOrgOk

`func (o *HandlersRunnerConnectionResponse) GetGithubOrgOk() (*string, bool)`

GetGithubOrgOk returns a tuple with the GithubOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubOrg

`func (o *HandlersRunnerConnectionResponse) SetGithubOrg(v string)`

SetGithubOrg sets GithubOrg field to given value.

### HasGithubOrg

`func (o *HandlersRunnerConnectionResponse) HasGithubOrg() bool`

HasGithubOrg returns a boolean if a field has been set.

### GetId

`func (o *HandlersRunnerConnectionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersRunnerConnectionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersRunnerConnectionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersRunnerConnectionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallationId

`func (o *HandlersRunnerConnectionResponse) GetInstallationId() int32`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *HandlersRunnerConnectionResponse) GetInstallationIdOk() (*int32, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *HandlersRunnerConnectionResponse) SetInstallationId(v int32)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *HandlersRunnerConnectionResponse) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetRepoCount

`func (o *HandlersRunnerConnectionResponse) GetRepoCount() int32`

GetRepoCount returns the RepoCount field if non-nil, zero value otherwise.

### GetRepoCountOk

`func (o *HandlersRunnerConnectionResponse) GetRepoCountOk() (*int32, bool)`

GetRepoCountOk returns a tuple with the RepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCount

`func (o *HandlersRunnerConnectionResponse) SetRepoCount(v int32)`

SetRepoCount sets RepoCount field to given value.

### HasRepoCount

`func (o *HandlersRunnerConnectionResponse) HasRepoCount() bool`

HasRepoCount returns a boolean if a field has been set.

### GetState

`func (o *HandlersRunnerConnectionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersRunnerConnectionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersRunnerConnectionResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersRunnerConnectionResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


