# DnsregistryScopeCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Check** | Pointer to **string** | Check names the rung (\&quot;token\&quot;, \&quot;zone-read\&quot;, \&quot;dns-edit\&quot; — or the arrival checklist&#39;s \&quot;account\&quot;, \&quot;tunnel-scope\&quot;). | [optional] 
**Ok** | Pointer to **bool** | OK reports whether the rung passed. | [optional] 

## Methods

### NewDnsregistryScopeCheck

`func NewDnsregistryScopeCheck() *DnsregistryScopeCheck`

NewDnsregistryScopeCheck instantiates a new DnsregistryScopeCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsregistryScopeCheckWithDefaults

`func NewDnsregistryScopeCheckWithDefaults() *DnsregistryScopeCheck`

NewDnsregistryScopeCheckWithDefaults instantiates a new DnsregistryScopeCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheck

`func (o *DnsregistryScopeCheck) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *DnsregistryScopeCheck) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *DnsregistryScopeCheck) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *DnsregistryScopeCheck) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetOk

`func (o *DnsregistryScopeCheck) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *DnsregistryScopeCheck) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *DnsregistryScopeCheck) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *DnsregistryScopeCheck) HasOk() bool`

HasOk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


