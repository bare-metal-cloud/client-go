# HandlersVaultRentalStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumed** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**RentalId** | Pointer to **string** |  | [optional] 
**Reused** | Pointer to **bool** | Reused is true when the mint hit the redelivery-reuse rule (a live, unconsumed rental already existed for this operation + target). | [optional] 
**TargetRef** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersVaultRentalStatusResponse

`func NewHandlersVaultRentalStatusResponse() *HandlersVaultRentalStatusResponse`

NewHandlersVaultRentalStatusResponse instantiates a new HandlersVaultRentalStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVaultRentalStatusResponseWithDefaults

`func NewHandlersVaultRentalStatusResponseWithDefaults() *HandlersVaultRentalStatusResponse`

NewHandlersVaultRentalStatusResponseWithDefaults instantiates a new HandlersVaultRentalStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumed

`func (o *HandlersVaultRentalStatusResponse) GetConsumed() bool`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *HandlersVaultRentalStatusResponse) GetConsumedOk() (*bool, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *HandlersVaultRentalStatusResponse) SetConsumed(v bool)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *HandlersVaultRentalStatusResponse) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersVaultRentalStatusResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersVaultRentalStatusResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersVaultRentalStatusResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersVaultRentalStatusResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *HandlersVaultRentalStatusResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HandlersVaultRentalStatusResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HandlersVaultRentalStatusResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HandlersVaultRentalStatusResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetOperation

`func (o *HandlersVaultRentalStatusResponse) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *HandlersVaultRentalStatusResponse) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *HandlersVaultRentalStatusResponse) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *HandlersVaultRentalStatusResponse) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRentalId

`func (o *HandlersVaultRentalStatusResponse) GetRentalId() string`

GetRentalId returns the RentalId field if non-nil, zero value otherwise.

### GetRentalIdOk

`func (o *HandlersVaultRentalStatusResponse) GetRentalIdOk() (*string, bool)`

GetRentalIdOk returns a tuple with the RentalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRentalId

`func (o *HandlersVaultRentalStatusResponse) SetRentalId(v string)`

SetRentalId sets RentalId field to given value.

### HasRentalId

`func (o *HandlersVaultRentalStatusResponse) HasRentalId() bool`

HasRentalId returns a boolean if a field has been set.

### GetReused

`func (o *HandlersVaultRentalStatusResponse) GetReused() bool`

GetReused returns the Reused field if non-nil, zero value otherwise.

### GetReusedOk

`func (o *HandlersVaultRentalStatusResponse) GetReusedOk() (*bool, bool)`

GetReusedOk returns a tuple with the Reused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReused

`func (o *HandlersVaultRentalStatusResponse) SetReused(v bool)`

SetReused sets Reused field to given value.

### HasReused

`func (o *HandlersVaultRentalStatusResponse) HasReused() bool`

HasReused returns a boolean if a field has been set.

### GetTargetRef

`func (o *HandlersVaultRentalStatusResponse) GetTargetRef() string`

GetTargetRef returns the TargetRef field if non-nil, zero value otherwise.

### GetTargetRefOk

`func (o *HandlersVaultRentalStatusResponse) GetTargetRefOk() (*string, bool)`

GetTargetRefOk returns a tuple with the TargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRef

`func (o *HandlersVaultRentalStatusResponse) SetTargetRef(v string)`

SetTargetRef sets TargetRef field to given value.

### HasTargetRef

`func (o *HandlersVaultRentalStatusResponse) HasTargetRef() bool`

HasTargetRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


