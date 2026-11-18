# DnsregistryRegisterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisories** | Pointer to **[]string** | Advisories carries the C11 nesting heads-ups (\&quot;zone X nests inside registered zone Y\&quot;). | [optional] 
**NsWarnings** | Pointer to **[]string** | NSWarnings lists the zones whose C3 advisory fired (a warning at connect — the registry still records them; the preflight hardens). | [optional] 
**Registered** | Pointer to [**[]ModelsDnsZone**](ModelsDnsZone.md) |  | [optional] 

## Methods

### NewDnsregistryRegisterResult

`func NewDnsregistryRegisterResult() *DnsregistryRegisterResult`

NewDnsregistryRegisterResult instantiates a new DnsregistryRegisterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsregistryRegisterResultWithDefaults

`func NewDnsregistryRegisterResultWithDefaults() *DnsregistryRegisterResult`

NewDnsregistryRegisterResultWithDefaults instantiates a new DnsregistryRegisterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisories

`func (o *DnsregistryRegisterResult) GetAdvisories() []string`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *DnsregistryRegisterResult) GetAdvisoriesOk() (*[]string, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *DnsregistryRegisterResult) SetAdvisories(v []string)`

SetAdvisories sets Advisories field to given value.

### HasAdvisories

`func (o *DnsregistryRegisterResult) HasAdvisories() bool`

HasAdvisories returns a boolean if a field has been set.

### GetNsWarnings

`func (o *DnsregistryRegisterResult) GetNsWarnings() []string`

GetNsWarnings returns the NsWarnings field if non-nil, zero value otherwise.

### GetNsWarningsOk

`func (o *DnsregistryRegisterResult) GetNsWarningsOk() (*[]string, bool)`

GetNsWarningsOk returns a tuple with the NsWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsWarnings

`func (o *DnsregistryRegisterResult) SetNsWarnings(v []string)`

SetNsWarnings sets NsWarnings field to given value.

### HasNsWarnings

`func (o *DnsregistryRegisterResult) HasNsWarnings() bool`

HasNsWarnings returns a boolean if a field has been set.

### GetRegistered

`func (o *DnsregistryRegisterResult) GetRegistered() []ModelsDnsZone`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *DnsregistryRegisterResult) GetRegisteredOk() (*[]ModelsDnsZone, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *DnsregistryRegisterResult) SetRegistered(v []ModelsDnsZone)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *DnsregistryRegisterResult) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


