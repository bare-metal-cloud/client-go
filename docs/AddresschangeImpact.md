# AddresschangeImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Change** | Pointer to [**AddresschangeFact**](AddresschangeFact.md) |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to [**AddresschangeSubject**](AddresschangeSubject.md) |  | [optional] 
**Zones** | Pointer to [**[]AddresschangeZone**](AddresschangeZone.md) |  | [optional] 

## Methods

### NewAddresschangeImpact

`func NewAddresschangeImpact() *AddresschangeImpact`

NewAddresschangeImpact instantiates a new AddresschangeImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddresschangeImpactWithDefaults

`func NewAddresschangeImpactWithDefaults() *AddresschangeImpact`

NewAddresschangeImpactWithDefaults instantiates a new AddresschangeImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AddresschangeImpact) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddresschangeImpact) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddresschangeImpact) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AddresschangeImpact) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetChange

`func (o *AddresschangeImpact) GetChange() AddresschangeFact`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AddresschangeImpact) GetChangeOk() (*AddresschangeFact, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AddresschangeImpact) SetChange(v AddresschangeFact)`

SetChange sets Change field to given value.

### HasChange

`func (o *AddresschangeImpact) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetFingerprint

`func (o *AddresschangeImpact) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *AddresschangeImpact) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *AddresschangeImpact) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *AddresschangeImpact) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetSteps

`func (o *AddresschangeImpact) GetSteps() []string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AddresschangeImpact) GetStepsOk() (*[]string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AddresschangeImpact) SetSteps(v []string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AddresschangeImpact) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetSubject

`func (o *AddresschangeImpact) GetSubject() AddresschangeSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AddresschangeImpact) GetSubjectOk() (*AddresschangeSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AddresschangeImpact) SetSubject(v AddresschangeSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AddresschangeImpact) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetZones

`func (o *AddresschangeImpact) GetZones() []AddresschangeZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddresschangeImpact) GetZonesOk() (*[]AddresschangeZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddresschangeImpact) SetZones(v []AddresschangeZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddresschangeImpact) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


