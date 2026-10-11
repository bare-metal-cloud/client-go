# HandlersJITCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DurationMinutes** | Pointer to **int32** |  | [optional] 
**ExpirationMode** | Pointer to **string** | ExpirationMode + DurationMinutes carry the maintainer&#39;s dual-mode choice (Epic 103 / ADR-078). ExpirationMode is \&quot;fixed\&quot; (default) or \&quot;sliding\&quot;. DurationMinutes sets the fixed deadline or the initial sliding window; it defaults to 15 when unset or out of range. | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**ServiceAccountId** | Pointer to **int32** |  | [optional] 
**TargetResourceId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersJITCreateRequest

`func NewHandlersJITCreateRequest() *HandlersJITCreateRequest`

NewHandlersJITCreateRequest instantiates a new HandlersJITCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersJITCreateRequestWithDefaults

`func NewHandlersJITCreateRequestWithDefaults() *HandlersJITCreateRequest`

NewHandlersJITCreateRequestWithDefaults instantiates a new HandlersJITCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *HandlersJITCreateRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HandlersJITCreateRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HandlersJITCreateRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HandlersJITCreateRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *HandlersJITCreateRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *HandlersJITCreateRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *HandlersJITCreateRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *HandlersJITCreateRequest) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *HandlersJITCreateRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *HandlersJITCreateRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *HandlersJITCreateRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *HandlersJITCreateRequest) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetExpirationMode

`func (o *HandlersJITCreateRequest) GetExpirationMode() string`

GetExpirationMode returns the ExpirationMode field if non-nil, zero value otherwise.

### GetExpirationModeOk

`func (o *HandlersJITCreateRequest) GetExpirationModeOk() (*string, bool)`

GetExpirationModeOk returns a tuple with the ExpirationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMode

`func (o *HandlersJITCreateRequest) SetExpirationMode(v string)`

SetExpirationMode sets ExpirationMode field to given value.

### HasExpirationMode

`func (o *HandlersJITCreateRequest) HasExpirationMode() bool`

HasExpirationMode returns a boolean if a field has been set.

### GetOperation

`func (o *HandlersJITCreateRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HandlersJITCreateRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HandlersJITCreateRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HandlersJITCreateRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *HandlersJITCreateRequest) GetServiceAccountId() int32`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HandlersJITCreateRequest) GetServiceAccountIdOk() (*int32, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HandlersJITCreateRequest) SetServiceAccountId(v int32)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *HandlersJITCreateRequest) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetTargetResourceId

`func (o *HandlersJITCreateRequest) GetTargetResourceId() int32`

GetTargetResourceId returns the TargetResourceId field if non-nil, zero value otherwise.

### GetTargetResourceIdOk

`func (o *HandlersJITCreateRequest) GetTargetResourceIdOk() (*int32, bool)`

GetTargetResourceIdOk returns a tuple with the TargetResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceId

`func (o *HandlersJITCreateRequest) SetTargetResourceId(v int32)`

SetTargetResourceId sets TargetResourceId field to given value.

### HasTargetResourceId

`func (o *HandlersJITCreateRequest) HasTargetResourceId() bool`

HasTargetResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


