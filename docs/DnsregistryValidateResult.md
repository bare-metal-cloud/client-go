# DnsregistryValidateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]DnsregistryScopeCheck**](DnsregistryScopeCheck.md) |  | [optional] 
**Zones** | Pointer to [**[]DnsregistryZoneInfo**](DnsregistryZoneInfo.md) |  | [optional] 

## Methods

### NewDnsregistryValidateResult

`func NewDnsregistryValidateResult() *DnsregistryValidateResult`

NewDnsregistryValidateResult instantiates a new DnsregistryValidateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsregistryValidateResultWithDefaults

`func NewDnsregistryValidateResultWithDefaults() *DnsregistryValidateResult`

NewDnsregistryValidateResultWithDefaults instantiates a new DnsregistryValidateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *DnsregistryValidateResult) GetChecks() []DnsregistryScopeCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *DnsregistryValidateResult) GetChecksOk() (*[]DnsregistryScopeCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *DnsregistryValidateResult) SetChecks(v []DnsregistryScopeCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *DnsregistryValidateResult) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetZones

`func (o *DnsregistryValidateResult) GetZones() []DnsregistryZoneInfo`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *DnsregistryValidateResult) GetZonesOk() (*[]DnsregistryZoneInfo, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *DnsregistryValidateResult) SetZones(v []DnsregistryZoneInfo)`

SetZones sets Zones field to given value.

### HasZones

`func (o *DnsregistryValidateResult) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


