# ModelsRunnerPackingHostCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetGb** | Pointer to **int32** | BudgetGB bounds the host&#39;s golden volumes total (§6.3 law 4): a cache fill can never starve the packer. Positive when enabled. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MinFreeGb** | Pointer to **int32** | MinFreeGB is the host free-space floor a cache volume create additionally respects. | [optional] 
**RepoScopes** | Pointer to **[]string** | RepoScopes is the per-repo opt-in list (\&quot;owner/name\&quot; entries); empty means org scopes only. | [optional] 

## Methods

### NewModelsRunnerPackingHostCache

`func NewModelsRunnerPackingHostCache() *ModelsRunnerPackingHostCache`

NewModelsRunnerPackingHostCache instantiates a new ModelsRunnerPackingHostCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRunnerPackingHostCacheWithDefaults

`func NewModelsRunnerPackingHostCacheWithDefaults() *ModelsRunnerPackingHostCache`

NewModelsRunnerPackingHostCacheWithDefaults instantiates a new ModelsRunnerPackingHostCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetGb

`func (o *ModelsRunnerPackingHostCache) GetBudgetGb() int32`

GetBudgetGb returns the BudgetGb field if non-nil, zero value otherwise.

### GetBudgetGbOk

`func (o *ModelsRunnerPackingHostCache) GetBudgetGbOk() (*int32, bool)`

GetBudgetGbOk returns a tuple with the BudgetGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetGb

`func (o *ModelsRunnerPackingHostCache) SetBudgetGb(v int32)`

SetBudgetGb sets BudgetGb field to given value.

### HasBudgetGb

`func (o *ModelsRunnerPackingHostCache) HasBudgetGb() bool`

HasBudgetGb returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsRunnerPackingHostCache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsRunnerPackingHostCache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsRunnerPackingHostCache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsRunnerPackingHostCache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinFreeGb

`func (o *ModelsRunnerPackingHostCache) GetMinFreeGb() int32`

GetMinFreeGb returns the MinFreeGb field if non-nil, zero value otherwise.

### GetMinFreeGbOk

`func (o *ModelsRunnerPackingHostCache) GetMinFreeGbOk() (*int32, bool)`

GetMinFreeGbOk returns a tuple with the MinFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeGb

`func (o *ModelsRunnerPackingHostCache) SetMinFreeGb(v int32)`

SetMinFreeGb sets MinFreeGb field to given value.

### HasMinFreeGb

`func (o *ModelsRunnerPackingHostCache) HasMinFreeGb() bool`

HasMinFreeGb returns a boolean if a field has been set.

### GetRepoScopes

`func (o *ModelsRunnerPackingHostCache) GetRepoScopes() []string`

GetRepoScopes returns the RepoScopes field if non-nil, zero value otherwise.

### GetRepoScopesOk

`func (o *ModelsRunnerPackingHostCache) GetRepoScopesOk() (*[]string, bool)`

GetRepoScopesOk returns a tuple with the RepoScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoScopes

`func (o *ModelsRunnerPackingHostCache) SetRepoScopes(v []string)`

SetRepoScopes sets RepoScopes field to given value.

### HasRepoScopes

`func (o *ModelsRunnerPackingHostCache) HasRepoScopes() bool`

HasRepoScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


