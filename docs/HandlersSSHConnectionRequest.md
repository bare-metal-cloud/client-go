# HandlersSSHConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientEnvelope** | Pointer to [**HandlersClientEnvelopeDTO**](HandlersClientEnvelopeDTO.md) | ClientEnvelope is the browser-sealed copy of PrivateSSHKey the server stores (issue #4424, B8); the plaintext is for the install push only. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentConnectionId** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PrivateSshKey** | Pointer to **string** |  | [optional] 
**UseAgentForwarding** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersSSHConnectionRequest

`func NewHandlersSSHConnectionRequest() *HandlersSSHConnectionRequest`

NewHandlersSSHConnectionRequest instantiates a new HandlersSSHConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersSSHConnectionRequestWithDefaults

`func NewHandlersSSHConnectionRequestWithDefaults() *HandlersSSHConnectionRequest`

NewHandlersSSHConnectionRequestWithDefaults instantiates a new HandlersSSHConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientEnvelope

`func (o *HandlersSSHConnectionRequest) GetClientEnvelope() HandlersClientEnvelopeDTO`

GetClientEnvelope returns the ClientEnvelope field if non-nil, zero value otherwise.

### GetClientEnvelopeOk

`func (o *HandlersSSHConnectionRequest) GetClientEnvelopeOk() (*HandlersClientEnvelopeDTO, bool)`

GetClientEnvelopeOk returns a tuple with the ClientEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEnvelope

`func (o *HandlersSSHConnectionRequest) SetClientEnvelope(v HandlersClientEnvelopeDTO)`

SetClientEnvelope sets ClientEnvelope field to given value.

### HasClientEnvelope

`func (o *HandlersSSHConnectionRequest) HasClientEnvelope() bool`

HasClientEnvelope returns a boolean if a field has been set.

### GetDescription

`func (o *HandlersSSHConnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersSSHConnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersSSHConnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersSSHConnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *HandlersSSHConnectionRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HandlersSSHConnectionRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HandlersSSHConnectionRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HandlersSSHConnectionRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *HandlersSSHConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersSSHConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersSSHConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersSSHConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentConnectionId

`func (o *HandlersSSHConnectionRequest) GetParentConnectionId() int32`

GetParentConnectionId returns the ParentConnectionId field if non-nil, zero value otherwise.

### GetParentConnectionIdOk

`func (o *HandlersSSHConnectionRequest) GetParentConnectionIdOk() (*int32, bool)`

GetParentConnectionIdOk returns a tuple with the ParentConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectionId

`func (o *HandlersSSHConnectionRequest) SetParentConnectionId(v int32)`

SetParentConnectionId sets ParentConnectionId field to given value.

### HasParentConnectionId

`func (o *HandlersSSHConnectionRequest) HasParentConnectionId() bool`

HasParentConnectionId returns a boolean if a field has been set.

### GetPort

`func (o *HandlersSSHConnectionRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersSSHConnectionRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersSSHConnectionRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersSSHConnectionRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateSshKey

`func (o *HandlersSSHConnectionRequest) GetPrivateSshKey() string`

GetPrivateSshKey returns the PrivateSshKey field if non-nil, zero value otherwise.

### GetPrivateSshKeyOk

`func (o *HandlersSSHConnectionRequest) GetPrivateSshKeyOk() (*string, bool)`

GetPrivateSshKeyOk returns a tuple with the PrivateSshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSshKey

`func (o *HandlersSSHConnectionRequest) SetPrivateSshKey(v string)`

SetPrivateSshKey sets PrivateSshKey field to given value.

### HasPrivateSshKey

`func (o *HandlersSSHConnectionRequest) HasPrivateSshKey() bool`

HasPrivateSshKey returns a boolean if a field has been set.

### GetUseAgentForwarding

`func (o *HandlersSSHConnectionRequest) GetUseAgentForwarding() bool`

GetUseAgentForwarding returns the UseAgentForwarding field if non-nil, zero value otherwise.

### GetUseAgentForwardingOk

`func (o *HandlersSSHConnectionRequest) GetUseAgentForwardingOk() (*bool, bool)`

GetUseAgentForwardingOk returns a tuple with the UseAgentForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgentForwarding

`func (o *HandlersSSHConnectionRequest) SetUseAgentForwarding(v bool)`

SetUseAgentForwarding sets UseAgentForwarding field to given value.

### HasUseAgentForwarding

`func (o *HandlersSSHConnectionRequest) HasUseAgentForwarding() bool`

HasUseAgentForwarding returns a boolean if a field has been set.

### GetUsername

`func (o *HandlersSSHConnectionRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersSSHConnectionRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersSSHConnectionRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersSSHConnectionRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


