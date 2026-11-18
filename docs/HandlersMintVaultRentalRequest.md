# HandlersMintVaultRentalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | Operation is one of the reviewed spender slugs (the vaultrental allowlist — an unknown operation is a 400). | [optional] 
**Payload** | Pointer to **string** | Payload is the released secret material (the browser&#39;s local decrypt). | [optional] 
**TargetRef** | Pointer to **string** | TargetRef names what the operation spends (\&quot;&lt;type&gt;:&lt;id&gt;\&quot; by convention) — the correlation key the spender handler consumes on. | [optional] 
**TtlMinutes** | Pointer to **int32** | TTLMinutes is optional; 0/absent &#x3D; the 15-minute default, values beyond the 60-minute ceiling are refused (never clamped). | [optional] 

## Methods

### NewHandlersMintVaultRentalRequest

`func NewHandlersMintVaultRentalRequest() *HandlersMintVaultRentalRequest`

NewHandlersMintVaultRentalRequest instantiates a new HandlersMintVaultRentalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMintVaultRentalRequestWithDefaults

`func NewHandlersMintVaultRentalRequestWithDefaults() *HandlersMintVaultRentalRequest`

NewHandlersMintVaultRentalRequestWithDefaults instantiates a new HandlersMintVaultRentalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *HandlersMintVaultRentalRequest) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HandlersMintVaultRentalRequest) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HandlersMintVaultRentalRequest) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HandlersMintVaultRentalRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPayload

`func (o *HandlersMintVaultRentalRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *HandlersMintVaultRentalRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *HandlersMintVaultRentalRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *HandlersMintVaultRentalRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTargetRef

`func (o *HandlersMintVaultRentalRequest) GetTargetRef() string`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *HandlersMintVaultRentalRequest) GetTargetRefOk() (*string, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *HandlersMintVaultRentalRequest) SetTargetRef(v string)`

SetTargetRef sets TargetRef field to given value.

### HasTargetRef

`func (o *HandlersMintVaultRentalRequest) HasTargetRef() bool`

HasTargetRef returns a boolean if a field has been set.

### GetTtlMinutes

`func (o *HandlersMintVaultRentalRequest) GetTtlMinutes() int32`

GetTtlMinutes returns the TtlMinutes field if non-nil, zero value otherwise.

### GetTtlMinutesOk

`func (o *HandlersMintVaultRentalRequest) GetTtlMinutesOk() (*int32, bool)`

GetTtlMinutesOk returns a tuple with the TtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlMinutes

`func (o *HandlersMintVaultRentalRequest) SetTtlMinutes(v int32)`

SetTtlMinutes sets TtlMinutes field to given value.

### HasTtlMinutes

`func (o *HandlersMintVaultRentalRequest) HasTtlMinutes() bool`

HasTtlMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


