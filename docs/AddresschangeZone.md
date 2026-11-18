# AddresschangeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custody** | Pointer to **string** | Custody is ours | theirs (the verdict the execute arm acts on). | [optional] 
**Provider** | Pointer to **string** | Provider names the custody&#39;s provider (\&quot;platform\&quot; for the seat, the registry row&#39;s provider otherwise; \&quot;\&quot; for theirs). | [optional] 
**Records** | Pointer to [**[]AddresschangeRecord**](AddresschangeRecord.md) |  | [optional] 
**Zone** | Pointer to **string** | Zone is the zone&#39;s FQDN. The platform seat&#39;s zone names the edge records; a registered row names its own. | [optional] 

## Methods

### NewAddresschangeZone

`func NewAddresschangeZone() *AddresschangeZone`

NewAddresschangeZone instantiates a new AddresschangeZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddresschangeZoneWithDefaults

`func NewAddresschangeZoneWithDefaults() *AddresschangeZone`

NewAddresschangeZoneWithDefaults instantiates a new AddresschangeZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustody

`func (o *AddresschangeZone) GetCustody() string`

GetCustody returns the Custody field if non-nil, zero value otherwise.

### GetCustodyOk

`func (o *AddresschangeZone) GetCustodyOk() (*string, bool)`

GetCustodyOk returns a tuple with the Custody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustody

`func (o *AddresschangeZone) SetCustody(v string)`

SetCustody sets Custody field to given value.

### HasCustody

`func (o *AddresschangeZone) HasCustody() bool`

HasCustody returns a boolean if a field has been set.

### GetProvider

`func (o *AddresschangeZone) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AddresschangeZone) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AddresschangeZone) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AddresschangeZone) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRecords

`func (o *AddresschangeZone) GetRecords() []AddresschangeRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *AddresschangeZone) GetRecordsOk() (*[]AddresschangeRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *AddresschangeZone) SetRecords(v []AddresschangeRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *AddresschangeZone) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetZone

`func (o *AddresschangeZone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddresschangeZone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddresschangeZone) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddresschangeZone) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


