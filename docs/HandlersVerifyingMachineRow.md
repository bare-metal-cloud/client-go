# HandlersVerifyingMachineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryKeyFingerprint** | Pointer to **string** |  | [optional] 
**DeliveryPublicKey** | Pointer to **string** | DeliveryPublicKey is the raw X25519 public key, base64 on the wire. | [optional] 
**PinnedToConnection** | Pointer to **bool** |  | [optional] 
**VerifyingAgentId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersVerifyingMachineRow

`func NewHandlersVerifyingMachineRow() *HandlersVerifyingMachineRow`

NewHandlersVerifyingMachineRow instantiates a new HandlersVerifyingMachineRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVerifyingMachineRowWithDefaults

`func NewHandlersVerifyingMachineRowWithDefaults() *HandlersVerifyingMachineRow`

NewHandlersVerifyingMachineRowWithDefaults instantiates a new HandlersVerifyingMachineRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryKeyFingerprint

`func (o *HandlersVerifyingMachineRow) GetDeliveryKeyFingerprint() string`

GetDeliveryKeyFingerprint returns the DeliveryKeyFingerprint field if non-nil, zero value otherwise.

### GetDeliveryKeyFingerprintOk

`func (o *HandlersVerifyingMachineRow) GetDeliveryKeyFingerprintOk() (*string, bool)`

GetDeliveryKeyFingerprintOk returns a tuple with the DeliveryKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryKeyFingerprint

`func (o *HandlersVerifyingMachineRow) SetDeliveryKeyFingerprint(v string)`

SetDeliveryKeyFingerprint sets DeliveryKeyFingerprint field to given value.

### HasDeliveryKeyFingerprint

`func (o *HandlersVerifyingMachineRow) HasDeliveryKeyFingerprint() bool`

HasDeliveryKeyFingerprint returns a boolean if a field has been set.

### GetDeliveryPublicKey

`func (o *HandlersVerifyingMachineRow) GetDeliveryPublicKey() string`

GetDeliveryPublicKey returns the DeliveryPublicKey field if non-nil, zero value otherwise.

### GetDeliveryPublicKeyOk

`func (o *HandlersVerifyingMachineRow) GetDeliveryPublicKeyOk() (*string, bool)`

GetDeliveryPublicKeyOk returns a tuple with the DeliveryPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPublicKey

`func (o *HandlersVerifyingMachineRow) SetDeliveryPublicKey(v string)`

SetDeliveryPublicKey sets DeliveryPublicKey field to given value.

### HasDeliveryPublicKey

`func (o *HandlersVerifyingMachineRow) HasDeliveryPublicKey() bool`

HasDeliveryPublicKey returns a boolean if a field has been set.

### GetPinnedToConnection

`func (o *HandlersVerifyingMachineRow) GetPinnedToConnection() bool`

GetPinnedToConnection returns the PinnedToConnection field if non-nil, zero value otherwise.

### GetPinnedToConnectionOk

`func (o *HandlersVerifyingMachineRow) GetPinnedToConnectionOk() (*bool, bool)`

GetPinnedToConnectionOk returns a tuple with the PinnedToConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedToConnection

`func (o *HandlersVerifyingMachineRow) SetPinnedToConnection(v bool)`

SetPinnedToConnection sets PinnedToConnection field to given value.

### HasPinnedToConnection

`func (o *HandlersVerifyingMachineRow) HasPinnedToConnection() bool`

HasPinnedToConnection returns a boolean if a field has been set.

### GetVerifyingAgentId

`func (o *HandlersVerifyingMachineRow) GetVerifyingAgentId() int32`

GetVerifyingAgentId returns the VerifyingAgentId field if non-nil, zero value otherwise.

### GetVerifyingAgentIdOk

`func (o *HandlersVerifyingMachineRow) GetVerifyingAgentIdOk() (*int32, bool)`

GetVerifyingAgentIdOk returns a tuple with the VerifyingAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyingAgentId

`func (o *HandlersVerifyingMachineRow) SetVerifyingAgentId(v int32)`

SetVerifyingAgentId sets VerifyingAgentId field to given value.

### HasVerifyingAgentId

`func (o *HandlersVerifyingMachineRow) HasVerifyingAgentId() bool`

HasVerifyingAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


