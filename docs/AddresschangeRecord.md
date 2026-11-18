# AddresschangeRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**NewTarget** | Pointer to **string** |  | [optional] 
**OldTarget** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** | Resolved reports the live re-check answer: the name&#39;s public DNS answers the new address. | [optional] 
**Zone** | Pointer to **string** | Zone is the record&#39;s zone FQDN (rides the record so the hand-over list is self-contained). | [optional] 

## Methods

### NewAddresschangeRecord

`func NewAddresschangeRecord() *AddresschangeRecord`

NewAddresschangeRecord instantiates a new AddresschangeRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddresschangeRecordWithDefaults

`func NewAddresschangeRecordWithDefaults() *AddresschangeRecord`

NewAddresschangeRecordWithDefaults instantiates a new AddresschangeRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *AddresschangeRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AddresschangeRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AddresschangeRecord) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AddresschangeRecord) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNewTarget

`func (o *AddresschangeRecord) GetNewTarget() string`

GetNewTarget returns the NewTarget field if non-nil, zero value otherwise.

### GetNewTargetOk

`func (o *AddresschangeRecord) GetNewTargetOk() (*string, bool)`

GetNewTargetOk returns a tuple with the NewTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTarget

`func (o *AddresschangeRecord) SetNewTarget(v string)`

SetNewTarget sets NewTarget field to given value.

### HasNewTarget

`func (o *AddresschangeRecord) HasNewTarget() bool`

HasNewTarget returns a boolean if a field has been set.

### GetOldTarget

`func (o *AddresschangeRecord) GetOldTarget() string`

GetOldTarget returns the OldTarget field if non-nil, zero value otherwise.

### GetOldTargetOk

`func (o *AddresschangeRecord) GetOldTargetOk() (*string, bool)`

GetOldTargetOk returns a tuple with the OldTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTarget

`func (o *AddresschangeRecord) SetOldTarget(v string)`

SetOldTarget sets OldTarget field to given value.

### HasOldTarget

`func (o *AddresschangeRecord) HasOldTarget() bool`

HasOldTarget returns a boolean if a field has been set.

### GetRecordType

`func (o *AddresschangeRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AddresschangeRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AddresschangeRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AddresschangeRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetResolved

`func (o *AddresschangeRecord) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *AddresschangeRecord) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *AddresschangeRecord) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *AddresschangeRecord) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetZone

`func (o *AddresschangeRecord) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddresschangeRecord) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddresschangeRecord) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddresschangeRecord) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


