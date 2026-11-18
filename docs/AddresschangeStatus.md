# AddresschangeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Change** | Pointer to [**AddresschangeFact**](AddresschangeFact.md) |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**Zones** | Pointer to [**[]AddresschangeZone**](AddresschangeZone.md) |  | [optional] 

## Methods

### NewAddresschangeStatus

`func NewAddresschangeStatus() *AddresschangeStatus`

NewAddresschangeStatus instantiates a new AddresschangeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddresschangeStatusWithDefaults

`func NewAddresschangeStatusWithDefaults() *AddresschangeStatus`

NewAddresschangeStatusWithDefaults instantiates a new AddresschangeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChange

`func (o *AddresschangeStatus) GetChange() AddresschangeFact`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AddresschangeStatus) GetChangeOk() (*AddresschangeFact, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AddresschangeStatus) SetChange(v AddresschangeFact)`

SetChange sets Change field to given value.

### HasChange

`func (o *AddresschangeStatus) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetResolved

`func (o *AddresschangeStatus) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *AddresschangeStatus) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *AddresschangeStatus) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *AddresschangeStatus) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetZones

`func (o *AddresschangeStatus) GetZones() []AddresschangeZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddresschangeStatus) GetZonesOk() (*[]AddresschangeZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddresschangeStatus) SetZones(v []AddresschangeZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddresschangeStatus) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


