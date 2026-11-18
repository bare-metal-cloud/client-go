# HandlersRunnerCacheScopeRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to **int32** | Hits/Warms are the scope&#39;s toolchain hit and warm counts (the separate-captioned-values law — never one merged number). | [optional] 
**Posture** | Pointer to **string** | Posture is the scope&#39;s rw/ro posture (the trust decision the server made — the page&#39;s chip, with the trusted-refs sentence as its caption). | [optional] 
**Scope** | Pointer to **string** | Scope is the scope&#39;s name (org- or repo-scoped, the host&#39;s naming). | [optional] 
**SizeBytes** | Pointer to **int32** | SizeBytes is the scope&#39;s current size on the host volume. | [optional] 
**Warms** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersRunnerCacheScopeRow

`func NewHandlersRunnerCacheScopeRow() *HandlersRunnerCacheScopeRow`

NewHandlersRunnerCacheScopeRow instantiates a new HandlersRunnerCacheScopeRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerCacheScopeRowWithDefaults

`func NewHandlersRunnerCacheScopeRowWithDefaults() *HandlersRunnerCacheScopeRow`

NewHandlersRunnerCacheScopeRowWithDefaults instantiates a new HandlersRunnerCacheScopeRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *HandlersRunnerCacheScopeRow) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *HandlersRunnerCacheScopeRow) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *HandlersRunnerCacheScopeRow) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *HandlersRunnerCacheScopeRow) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetPosture

`func (o *HandlersRunnerCacheScopeRow) GetPosture() string`

GetPosture returns the Posture field if non-nil, zero value otherwise.

### GetPostureOk

`func (o *HandlersRunnerCacheScopeRow) GetPostureOk() (*string, bool)`

GetPostureOk returns a tuple with the Posture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosture

`func (o *HandlersRunnerCacheScopeRow) SetPosture(v string)`

SetPosture sets Posture field to given value.

### HasPosture

`func (o *HandlersRunnerCacheScopeRow) HasPosture() bool`

HasPosture returns a boolean if a field has been set.

### GetScope

`func (o *HandlersRunnerCacheScopeRow) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersRunnerCacheScopeRow) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersRunnerCacheScopeRow) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersRunnerCacheScopeRow) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSizeBytes

`func (o *HandlersRunnerCacheScopeRow) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *HandlersRunnerCacheScopeRow) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *HandlersRunnerCacheScopeRow) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *HandlersRunnerCacheScopeRow) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetWarms

`func (o *HandlersRunnerCacheScopeRow) GetWarms() int32`

GetWarms returns the Warms field if non-nil, zero value otherwise.

### GetWarmsOk

`func (o *HandlersRunnerCacheScopeRow) GetWarmsOk() (*int32, bool)`

GetWarmsOk returns a tuple with the Warms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarms

`func (o *HandlersRunnerCacheScopeRow) SetWarms(v int32)`

SetWarms sets Warms field to given value.

### HasWarms

`func (o *HandlersRunnerCacheScopeRow) HasWarms() bool`

HasWarms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


