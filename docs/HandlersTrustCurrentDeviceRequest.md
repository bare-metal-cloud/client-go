# HandlersTrustCurrentDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Optional: update device name | [optional] 
**EncryptionKey** | **string** |  | 

## Methods

### NewHandlersTrustCurrentDeviceRequest

`func NewHandlersTrustCurrentDeviceRequest(encryptionKey string, ) *HandlersTrustCurrentDeviceRequest`

NewHandlersTrustCurrentDeviceRequest instantiates a new HandlersTrustCurrentDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersTrustCurrentDeviceRequestWithDefaults

`func NewHandlersTrustCurrentDeviceRequestWithDefaults() *HandlersTrustCurrentDeviceRequest`

NewHandlersTrustCurrentDeviceRequestWithDefaults instantiates a new HandlersTrustCurrentDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *HandlersTrustCurrentDeviceRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *HandlersTrustCurrentDeviceRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *HandlersTrustCurrentDeviceRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *HandlersTrustCurrentDeviceRequest) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *HandlersTrustCurrentDeviceRequest) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *HandlersTrustCurrentDeviceRequest) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *HandlersTrustCurrentDeviceRequest) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


