# ModelsArrivalProviderCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountID is the provider account the token was validated against at install (R13 — the same-account preflight&#39;s arrival half). A rotation whose token answers from a DIFFERENT account refuses by name (reinstall the addon to change accounts). | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Instance** | Pointer to **string** | Instance is the addon instance id the route bindings and arrival facts name (the R1/R13 disambiguator). | [optional] 
**LastEnsuredAt** | Pointer to **string** | LastEnsuredAt is when the arrival-ensure last converged this custody&#39;s connectors (the re-ensure pass&#39;s success stamp). Nil until the first pass lands. A row whose TokenRotatedAt is newer than its LastEnsuredAt is the honest rotation gap the surfaces render. | [optional] 
**OrgId** | Pointer to **int32** | OrgID is the tenancy boundary. Several tunnel addon instances may coexist (R1/R13), so the uniqueness is per (org, provider, instance). | [optional] 
**Provider** | Pointer to **string** | Provider names the tunnel provider (\&quot;cloudflare\&quot; in v1). | [optional] 
**TokenRotatedAt** | Pointer to **string** | TokenRotatedAt is when the custody&#39;s token last ROTATED (RotateSealed — Epic 163 entry 6, FR-7&#39;s convergence law). Nil for install-time captures (Seal): a fresh install&#39;s addon flow ensures its connector right after capture, and there is nothing rotated to converge. The rotation-convergence pass reads this against LastEnsuredAt to find the rotated-since-ensure custody rows. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsArrivalProviderCredential

`func NewModelsArrivalProviderCredential() *ModelsArrivalProviderCredential`

NewModelsArrivalProviderCredential instantiates a new ModelsArrivalProviderCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsArrivalProviderCredentialWithDefaults

`func NewModelsArrivalProviderCredentialWithDefaults() *ModelsArrivalProviderCredential`

NewModelsArrivalProviderCredentialWithDefaults instantiates a new ModelsArrivalProviderCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ModelsArrivalProviderCredential) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ModelsArrivalProviderCredential) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ModelsArrivalProviderCredential) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ModelsArrivalProviderCredential) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsArrivalProviderCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsArrivalProviderCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsArrivalProviderCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsArrivalProviderCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsArrivalProviderCredential) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsArrivalProviderCredential) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsArrivalProviderCredential) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsArrivalProviderCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstance

`func (o *ModelsArrivalProviderCredential) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ModelsArrivalProviderCredential) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ModelsArrivalProviderCredential) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ModelsArrivalProviderCredential) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetLastEnsuredAt

`func (o *ModelsArrivalProviderCredential) GetLastEnsuredAt() string`

GetLastEnsuredAt returns the LastEnsuredAt field if non-nil, zero value otherwise.

### GetLastEnsuredAtOk

`func (o *ModelsArrivalProviderCredential) GetLastEnsuredAtOk() (*string, bool)`

GetLastEnsuredAtOk returns a tuple with the LastEnsuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnsuredAt

`func (o *ModelsArrivalProviderCredential) SetLastEnsuredAt(v string)`

SetLastEnsuredAt sets LastEnsuredAt field to given value.

### HasLastEnsuredAt

`func (o *ModelsArrivalProviderCredential) HasLastEnsuredAt() bool`

HasLastEnsuredAt returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsArrivalProviderCredential) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsArrivalProviderCredential) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsArrivalProviderCredential) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsArrivalProviderCredential) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelsArrivalProviderCredential) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsArrivalProviderCredential) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsArrivalProviderCredential) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelsArrivalProviderCredential) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetTokenRotatedAt

`func (o *ModelsArrivalProviderCredential) GetTokenRotatedAt() string`

GetTokenRotatedAt returns the TokenRotatedAt field if non-nil, zero value otherwise.

### GetTokenRotatedAtOk

`func (o *ModelsArrivalProviderCredential) GetTokenRotatedAtOk() (*string, bool)`

GetTokenRotatedAtOk returns a tuple with the TokenRotatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRotatedAt

`func (o *ModelsArrivalProviderCredential) SetTokenRotatedAt(v string)`

SetTokenRotatedAt sets TokenRotatedAt field to given value.

### HasTokenRotatedAt

`func (o *ModelsArrivalProviderCredential) HasTokenRotatedAt() bool`

HasTokenRotatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsArrivalProviderCredential) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsArrivalProviderCredential) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsArrivalProviderCredential) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsArrivalProviderCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


