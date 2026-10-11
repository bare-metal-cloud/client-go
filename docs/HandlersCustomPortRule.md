# HandlersCustomPortRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | CIDR block (default: \&quot;0.0.0.0/0\&quot; for all IPs) | [optional] 
**Description** | Pointer to **string** | Description for the rule | [optional] 
**PortRange** | Pointer to **string** | Port range (e.g., \&quot;80\&quot;, \&quot;443\&quot;, \&quot;3000-3005\&quot;) | [optional] 
**Protocol** | Pointer to **string** | \&quot;tcp\&quot; or \&quot;udp\&quot; | [optional] 

## Methods

### NewHandlersCustomPortRule

`func NewHandlersCustomPortRule() *HandlersCustomPortRule`

NewHandlersCustomPortRule instantiates a new HandlersCustomPortRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCustomPortRuleWithDefaults

`func NewHandlersCustomPortRuleWithDefaults() *HandlersCustomPortRule`

NewHandlersCustomPortRuleWithDefaults instantiates a new HandlersCustomPortRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *HandlersCustomPortRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *HandlersCustomPortRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *HandlersCustomPortRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *HandlersCustomPortRule) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetDescription

`func (o *HandlersCustomPortRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersCustomPortRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersCustomPortRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersCustomPortRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPortRange

`func (o *HandlersCustomPortRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *HandlersCustomPortRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *HandlersCustomPortRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *HandlersCustomPortRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *HandlersCustomPortRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HandlersCustomPortRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HandlersCustomPortRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *HandlersCustomPortRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


