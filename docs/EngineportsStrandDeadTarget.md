# EngineportsStrandDeadTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupersededAt** | Pointer to **string** | SupersededAt is when an ensure found a different live id (NULL while the row is the ledger&#39;s active memory — during the refusal window the active row IS the dead id, superseded only by the ensure that clears). | [optional] 
**Target** | Pointer to **string** | Target is the full tunnel CNAME (the string a stranded record&#39;s content equals, by the full-target equality law). | [optional] 
**TunnelId** | Pointer to **string** | TunnelID is the provider&#39;s opaque tunnel id whose CNAME this was. | [optional] 

## Methods

### NewEngineportsStrandDeadTarget

`func NewEngineportsStrandDeadTarget() *EngineportsStrandDeadTarget`

NewEngineportsStrandDeadTarget instantiates a new EngineportsStrandDeadTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsStrandDeadTargetWithDefaults

`func NewEngineportsStrandDeadTargetWithDefaults() *EngineportsStrandDeadTarget`

NewEngineportsStrandDeadTargetWithDefaults instantiates a new EngineportsStrandDeadTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupersededAt

`func (o *EngineportsStrandDeadTarget) GetSupersededAt() string`

GetSupersededAt returns the SupersededAt field if non-nil, zero value otherwise.

### GetSupersededAtOk

`func (o *EngineportsStrandDeadTarget) GetSupersededAtOk() (*string, bool)`

GetSupersededAtOk returns a tuple with the SupersededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersededAt

`func (o *EngineportsStrandDeadTarget) SetSupersededAt(v string)`

SetSupersededAt sets SupersededAt field to given value.

### HasSupersededAt

`func (o *EngineportsStrandDeadTarget) HasSupersededAt() bool`

HasSupersededAt returns a boolean if a field has been set.

### GetTarget

`func (o *EngineportsStrandDeadTarget) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EngineportsStrandDeadTarget) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EngineportsStrandDeadTarget) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *EngineportsStrandDeadTarget) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTunnelId

`func (o *EngineportsStrandDeadTarget) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *EngineportsStrandDeadTarget) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *EngineportsStrandDeadTarget) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *EngineportsStrandDeadTarget) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


