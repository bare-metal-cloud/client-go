# HandlersCreateMultipassHostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** | Remote connection details (required only if IsLocal is false) | [optional] 
**IsLocal** | Pointer to **bool** | true for local machine, false for remote | [optional] 
**Name** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] 
**PrivateSshKey** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCreateMultipassHostRequest

`func NewHandlersCreateMultipassHostRequest(name string, ) *HandlersCreateMultipassHostRequest`

NewHandlersCreateMultipassHostRequest instantiates a new HandlersCreateMultipassHostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateMultipassHostRequestWithDefaults

`func NewHandlersCreateMultipassHostRequestWithDefaults() *HandlersCreateMultipassHostRequest`

NewHandlersCreateMultipassHostRequestWithDefaults instantiates a new HandlersCreateMultipassHostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HandlersCreateMultipassHostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersCreateMultipassHostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersCreateMultipassHostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersCreateMultipassHostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *HandlersCreateMultipassHostRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HandlersCreateMultipassHostRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HandlersCreateMultipassHostRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HandlersCreateMultipassHostRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsLocal

`func (o *HandlersCreateMultipassHostRequest) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *HandlersCreateMultipassHostRequest) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *HandlersCreateMultipassHostRequest) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *HandlersCreateMultipassHostRequest) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetName

`func (o *HandlersCreateMultipassHostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateMultipassHostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateMultipassHostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *HandlersCreateMultipassHostRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersCreateMultipassHostRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersCreateMultipassHostRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersCreateMultipassHostRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateSshKey

`func (o *HandlersCreateMultipassHostRequest) GetPrivateSshKey() string`

GetPrivateSshKey returns the PrivateSshKey field if non-nil, zero value otherwise.

### GetPrivateSshKeyOk

`func (o *HandlersCreateMultipassHostRequest) GetPrivateSshKeyOk() (*string, bool)`

GetPrivateSshKeyOk returns a tuple with the PrivateSshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSshKey

`func (o *HandlersCreateMultipassHostRequest) SetPrivateSshKey(v string)`

SetPrivateSshKey sets PrivateSshKey field to given value.

### HasPrivateSshKey

`func (o *HandlersCreateMultipassHostRequest) HasPrivateSshKey() bool`

HasPrivateSshKey returns a boolean if a field has been set.

### GetScope

`func (o *HandlersCreateMultipassHostRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersCreateMultipassHostRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersCreateMultipassHostRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersCreateMultipassHostRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUsername

`func (o *HandlersCreateMultipassHostRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersCreateMultipassHostRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersCreateMultipassHostRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersCreateMultipassHostRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


