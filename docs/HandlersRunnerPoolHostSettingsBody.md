# HandlersRunnerPoolHostSettingsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**HandlersRunnerPoolHostCacheS**](HandlersRunnerPoolHostCacheS.md) |  | [optional] 
**Lids** | Pointer to [**[]HandlersRunnerPoolHostLid**](HandlersRunnerPoolHostLid.md) |  | [optional] 
**MinFreeGb** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersRunnerPoolHostSettingsBody

`func NewHandlersRunnerPoolHostSettingsBody() *HandlersRunnerPoolHostSettingsBody`

NewHandlersRunnerPoolHostSettingsBody instantiates a new HandlersRunnerPoolHostSettingsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerPoolHostSettingsBodyWithDefaults

`func NewHandlersRunnerPoolHostSettingsBodyWithDefaults() *HandlersRunnerPoolHostSettingsBody`

NewHandlersRunnerPoolHostSettingsBodyWithDefaults instantiates a new HandlersRunnerPoolHostSettingsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *HandlersRunnerPoolHostSettingsBody) GetCache() HandlersRunnerPoolHostCacheS`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *HandlersRunnerPoolHostSettingsBody) GetCacheOk() (*HandlersRunnerPoolHostCacheS, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *HandlersRunnerPoolHostSettingsBody) SetCache(v HandlersRunnerPoolHostCacheS)`

SetCache sets Cache field to given value.

### HasCache

`func (o *HandlersRunnerPoolHostSettingsBody) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetLids

`func (o *HandlersRunnerPoolHostSettingsBody) GetLids() []HandlersRunnerPoolHostLid`

GetLids returns the Lids field if non-nil, zero value otherwise.

### GetLidsOk

`func (o *HandlersRunnerPoolHostSettingsBody) GetLidsOk() (*[]HandlersRunnerPoolHostLid, bool)`

GetLidsOk returns a tuple with the Lids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLids

`func (o *HandlersRunnerPoolHostSettingsBody) SetLids(v []HandlersRunnerPoolHostLid)`

SetLids sets Lids field to given value.

### HasLids

`func (o *HandlersRunnerPoolHostSettingsBody) HasLids() bool`

HasLids returns a boolean if a field has been set.

### GetMinFreeGb

`func (o *HandlersRunnerPoolHostSettingsBody) GetMinFreeGb() int32`

GetMinFreeGb returns the MinFreeGb field if non-nil, zero value otherwise.

### GetMinFreeGbOk

`func (o *HandlersRunnerPoolHostSettingsBody) GetMinFreeGbOk() (*int32, bool)`

GetMinFreeGbOk returns a tuple with the MinFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeGb

`func (o *HandlersRunnerPoolHostSettingsBody) SetMinFreeGb(v int32)`

SetMinFreeGb sets MinFreeGb field to given value.

### HasMinFreeGb

`func (o *HandlersRunnerPoolHostSettingsBody) HasMinFreeGb() bool`

HasMinFreeGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


