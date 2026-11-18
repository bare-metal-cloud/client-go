# EngineBuildInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployedBy** | Pointer to **string** | CI token subject | user id | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Sha** | Pointer to **string** |  | [optional] 
**WorkflowRun** | Pointer to **string** |  | [optional] 

## Methods

### NewEngineBuildInfo

`func NewEngineBuildInfo() *EngineBuildInfo`

NewEngineBuildInfo instantiates a new EngineBuildInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineBuildInfoWithDefaults

`func NewEngineBuildInfoWithDefaults() *EngineBuildInfo`

NewEngineBuildInfoWithDefaults instantiates a new EngineBuildInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployedBy

`func (o *EngineBuildInfo) GetDeployedBy() string`

GetDeployedBy returns the DeployedBy field if non-nil, zero value otherwise.

### GetDeployedByOk

`func (o *EngineBuildInfo) GetDeployedByOk() (*string, bool)`

GetDeployedByOk returns a tuple with the DeployedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedBy

`func (o *EngineBuildInfo) SetDeployedBy(v string)`

SetDeployedBy sets DeployedBy field to given value.

### HasDeployedBy

`func (o *EngineBuildInfo) HasDeployedBy() bool`

HasDeployedBy returns a boolean if a field has been set.

### GetImageDigest

`func (o *EngineBuildInfo) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *EngineBuildInfo) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *EngineBuildInfo) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *EngineBuildInfo) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetRepo

`func (o *EngineBuildInfo) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *EngineBuildInfo) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *EngineBuildInfo) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *EngineBuildInfo) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSha

`func (o *EngineBuildInfo) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *EngineBuildInfo) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *EngineBuildInfo) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *EngineBuildInfo) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetWorkflowRun

`func (o *EngineBuildInfo) GetWorkflowRun() string`

GetWorkflowRun returns the WorkflowRun field if non-nil, zero value otherwise.

### GetWorkflowRunOk

`func (o *EngineBuildInfo) GetWorkflowRunOk() (*string, bool)`

GetWorkflowRunOk returns a tuple with the WorkflowRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRun

`func (o *EngineBuildInfo) SetWorkflowRun(v string)`

SetWorkflowRun sets WorkflowRun field to given value.

### HasWorkflowRun

`func (o *EngineBuildInfo) HasWorkflowRun() bool`

HasWorkflowRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


