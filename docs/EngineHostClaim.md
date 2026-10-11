# EngineHostClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the claimed hostname (normalized: lowercase, no trailing dot). | [optional] 
**InstallName** | Pointer to **string** | InstallName names the install holding the claim — the conflicting install a uniqueness refusal names (the backend maps that refusal to HTTP 422; the module&#39;s contract is the sentinel). | [optional] 

## Methods

### NewEngineHostClaim

`func NewEngineHostClaim() *EngineHostClaim`

NewEngineHostClaim instantiates a new EngineHostClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineHostClaimWithDefaults

`func NewEngineHostClaimWithDefaults() *EngineHostClaim`

NewEngineHostClaimWithDefaults instantiates a new EngineHostClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *EngineHostClaim) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EngineHostClaim) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EngineHostClaim) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EngineHostClaim) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInstallName

`func (o *EngineHostClaim) GetInstallName() string`

GetInstallName returns the InstallName field if non-nil, zero value otherwise.

### GetInstallNameOk

`func (o *EngineHostClaim) GetInstallNameOk() (*string, bool)`

GetInstallNameOk returns a tuple with the InstallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallName

`func (o *EngineHostClaim) SetInstallName(v string)`

SetInstallName sets InstallName field to given value.

### HasInstallName

`func (o *EngineHostClaim) HasInstallName() bool`

HasInstallName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


