# NetworkbackendNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ModelsNetworkConfig**](ModelsNetworkConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DraftName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Purpose** | Pointer to [**ModelsNetworkPurpose**](ModelsNetworkPurpose.md) | Draft intent (story 124.17 AC7): &#39;cluster-create&#39; when the wizard created this network as a draft for the cluster named by DraftName. Surfaces the Networks-page resume banner / \&quot;Prepared for a cluster\&quot; chip without a per-row join-status call. | [optional] 
**Status** | Pointer to [**ModelsNetworkStatus**](ModelsNetworkStatus.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkbackendNetwork

`func NewNetworkbackendNetwork() *NetworkbackendNetwork`

NewNetworkbackendNetwork instantiates a new NetworkbackendNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkbackendNetworkWithDefaults

`func NewNetworkbackendNetworkWithDefaults() *NetworkbackendNetwork`

NewNetworkbackendNetworkWithDefaults instantiates a new NetworkbackendNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *NetworkbackendNetwork) GetConfig() ModelsNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkbackendNetwork) GetConfigOk() (*ModelsNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkbackendNetwork) SetConfig(v ModelsNetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkbackendNetwork) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NetworkbackendNetwork) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkbackendNetwork) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkbackendNetwork) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NetworkbackendNetwork) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDraftName

`func (o *NetworkbackendNetwork) GetDraftName() string`

GetDraftName returns the DraftName field if non-nil, zero value otherwise.

### GetDraftNameOk

`func (o *NetworkbackendNetwork) GetDraftNameOk() (*string, bool)`

GetDraftNameOk returns a tuple with the DraftName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftName

`func (o *NetworkbackendNetwork) SetDraftName(v string)`

SetDraftName sets DraftName field to given value.

### HasDraftName

`func (o *NetworkbackendNetwork) HasDraftName() bool`

HasDraftName returns a boolean if a field has been set.

### GetId

`func (o *NetworkbackendNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkbackendNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkbackendNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkbackendNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *NetworkbackendNetwork) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *NetworkbackendNetwork) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *NetworkbackendNetwork) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *NetworkbackendNetwork) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *NetworkbackendNetwork) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NetworkbackendNetwork) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NetworkbackendNetwork) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *NetworkbackendNetwork) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPurpose

`func (o *NetworkbackendNetwork) GetPurpose() ModelsNetworkPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *NetworkbackendNetwork) GetPurposeOk() (*ModelsNetworkPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *NetworkbackendNetwork) SetPurpose(v ModelsNetworkPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *NetworkbackendNetwork) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkbackendNetwork) GetStatus() ModelsNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkbackendNetwork) GetStatusOk() (*ModelsNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkbackendNetwork) SetStatus(v ModelsNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkbackendNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkbackendNetwork) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkbackendNetwork) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkbackendNetwork) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkbackendNetwork) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


