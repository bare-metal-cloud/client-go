# HandlersMultipassHostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **string** | From SSH connection | [optional] 
**IsLocal** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**MultipassVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SshConnectionId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersMultipassHostResponse

`func NewHandlersMultipassHostResponse() *HandlersMultipassHostResponse`

NewHandlersMultipassHostResponse instantiates a new HandlersMultipassHostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMultipassHostResponseWithDefaults

`func NewHandlersMultipassHostResponseWithDefaults() *HandlersMultipassHostResponse`

NewHandlersMultipassHostResponseWithDefaults instantiates a new HandlersMultipassHostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HandlersMultipassHostResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersMultipassHostResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersMultipassHostResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersMultipassHostResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *HandlersMultipassHostResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersMultipassHostResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersMultipassHostResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersMultipassHostResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *HandlersMultipassHostResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersMultipassHostResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersMultipassHostResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersMultipassHostResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *HandlersMultipassHostResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HandlersMultipassHostResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HandlersMultipassHostResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HandlersMultipassHostResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsLocal

`func (o *HandlersMultipassHostResponse) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *HandlersMultipassHostResponse) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *HandlersMultipassHostResponse) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *HandlersMultipassHostResponse) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetIsVerified

`func (o *HandlersMultipassHostResponse) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *HandlersMultipassHostResponse) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *HandlersMultipassHostResponse) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *HandlersMultipassHostResponse) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetMultipassVersion

`func (o *HandlersMultipassHostResponse) GetMultipassVersion() string`

GetMultipassVersion returns the MultipassVersion field if non-nil, zero value otherwise.

### GetMultipassVersionOk

`func (o *HandlersMultipassHostResponse) GetMultipassVersionOk() (*string, bool)`

GetMultipassVersionOk returns a tuple with the MultipassVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipassVersion

`func (o *HandlersMultipassHostResponse) SetMultipassVersion(v string)`

SetMultipassVersion sets MultipassVersion field to given value.

### HasMultipassVersion

`func (o *HandlersMultipassHostResponse) HasMultipassVersion() bool`

HasMultipassVersion returns a boolean if a field has been set.

### GetName

`func (o *HandlersMultipassHostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersMultipassHostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersMultipassHostResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersMultipassHostResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSshConnectionId

`func (o *HandlersMultipassHostResponse) GetSshConnectionId() int32`

GetSshConnectionId returns the SshConnectionId field if non-nil, zero value otherwise.

### GetSshConnectionIdOk

`func (o *HandlersMultipassHostResponse) GetSshConnectionIdOk() (*int32, bool)`

GetSshConnectionIdOk returns a tuple with the SshConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConnectionId

`func (o *HandlersMultipassHostResponse) SetSshConnectionId(v int32)`

SetSshConnectionId sets SshConnectionId field to given value.

### HasSshConnectionId

`func (o *HandlersMultipassHostResponse) HasSshConnectionId() bool`

HasSshConnectionId returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersMultipassHostResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersMultipassHostResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersMultipassHostResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersMultipassHostResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *HandlersMultipassHostResponse) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *HandlersMultipassHostResponse) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *HandlersMultipassHostResponse) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *HandlersMultipassHostResponse) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


