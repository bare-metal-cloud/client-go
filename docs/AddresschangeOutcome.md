# AddresschangeOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to **string** | repointed | already_converged | drifted | failed | handed_over | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewAddresschangeOutcome

`func NewAddresschangeOutcome() *AddresschangeOutcome`

NewAddresschangeOutcome instantiates a new AddresschangeOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddresschangeOutcomeWithDefaults

`func NewAddresschangeOutcomeWithDefaults() *AddresschangeOutcome`

NewAddresschangeOutcomeWithDefaults instantiates a new AddresschangeOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AddresschangeOutcome) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AddresschangeOutcome) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AddresschangeOutcome) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AddresschangeOutcome) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *AddresschangeOutcome) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AddresschangeOutcome) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AddresschangeOutcome) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AddresschangeOutcome) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOutcome

`func (o *AddresschangeOutcome) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AddresschangeOutcome) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AddresschangeOutcome) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AddresschangeOutcome) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetZone

`func (o *AddresschangeOutcome) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddresschangeOutcome) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddresschangeOutcome) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddresschangeOutcome) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


