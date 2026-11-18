# DnsregistryZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Contains** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameServers** | Pointer to **[]string** |  | [optional] 
**NestedUnder** | Pointer to **string** | NestedUnder / Contains name already-registered zones this zone nests with (the C11 advisory: one heads-up, never a refusal). | [optional] 
**NsMismatch** | Pointer to **bool** | NSMismatch reports the C3 advisory: the internet&#39;s delegation does not (yet) point at the provider&#39;s declared nameservers — a WARNING at connect (a user may be mid-migration), a refusal at preflight. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDnsregistryZoneInfo

`func NewDnsregistryZoneInfo() *DnsregistryZoneInfo`

NewDnsregistryZoneInfo instantiates a new DnsregistryZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsregistryZoneInfoWithDefaults

`func NewDnsregistryZoneInfoWithDefaults() *DnsregistryZoneInfo`

NewDnsregistryZoneInfoWithDefaults instantiates a new DnsregistryZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DnsregistryZoneInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DnsregistryZoneInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DnsregistryZoneInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DnsregistryZoneInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *DnsregistryZoneInfo) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *DnsregistryZoneInfo) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *DnsregistryZoneInfo) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *DnsregistryZoneInfo) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetContains

`func (o *DnsregistryZoneInfo) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *DnsregistryZoneInfo) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *DnsregistryZoneInfo) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *DnsregistryZoneInfo) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetId

`func (o *DnsregistryZoneInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsregistryZoneInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsregistryZoneInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsregistryZoneInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsregistryZoneInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsregistryZoneInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsregistryZoneInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsregistryZoneInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameServers

`func (o *DnsregistryZoneInfo) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *DnsregistryZoneInfo) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *DnsregistryZoneInfo) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *DnsregistryZoneInfo) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetNestedUnder

`func (o *DnsregistryZoneInfo) GetNestedUnder() string`

GetNestedUnder returns the NestedUnder field if non-nil, zero value otherwise.

### GetNestedUnderOk

`func (o *DnsregistryZoneInfo) GetNestedUnderOk() (*string, bool)`

GetNestedUnderOk returns a tuple with the NestedUnder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedUnder

`func (o *DnsregistryZoneInfo) SetNestedUnder(v string)`

SetNestedUnder sets NestedUnder field to given value.

### HasNestedUnder

`func (o *DnsregistryZoneInfo) HasNestedUnder() bool`

HasNestedUnder returns a boolean if a field has been set.

### GetNsMismatch

`func (o *DnsregistryZoneInfo) GetNsMismatch() bool`

GetNsMismatch returns the NsMismatch field if non-nil, zero value otherwise.

### GetNsMismatchOk

`func (o *DnsregistryZoneInfo) GetNsMismatchOk() (*bool, bool)`

GetNsMismatchOk returns a tuple with the NsMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsMismatch

`func (o *DnsregistryZoneInfo) SetNsMismatch(v bool)`

SetNsMismatch sets NsMismatch field to given value.

### HasNsMismatch

`func (o *DnsregistryZoneInfo) HasNsMismatch() bool`

HasNsMismatch returns a boolean if a field has been set.

### GetStatus

`func (o *DnsregistryZoneInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DnsregistryZoneInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DnsregistryZoneInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DnsregistryZoneInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


