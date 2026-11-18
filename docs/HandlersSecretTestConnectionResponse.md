# HandlersSecretTestConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | Pointer to **string** |  | [optional] 
**CheckedAt** | Pointer to **string** | CheckedAt / LastVerifiedAt are RFC3339; LastVerifiedAt is null until a check succeeds (a failed check must not render as a verification). | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**LastVerifiedAt** | Pointer to **string** |  | [optional] 
**LatencyMs** | Pointer to **int32** |  | [optional] 
**Ok** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersSecretTestConnectionResponse

`func NewHandlersSecretTestConnectionResponse() *HandlersSecretTestConnectionResponse`

NewHandlersSecretTestConnectionResponse instantiates a new HandlersSecretTestConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersSecretTestConnectionResponseWithDefaults

`func NewHandlersSecretTestConnectionResponseWithDefaults() *HandlersSecretTestConnectionResponse`

NewHandlersSecretTestConnectionResponseWithDefaults instantiates a new HandlersSecretTestConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *HandlersSecretTestConnectionResponse) GetCall() string`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *HandlersSecretTestConnectionResponse) GetCallOk() (*string, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *HandlersSecretTestConnectionResponse) SetCall(v string)`

SetCall sets Call field to given value.

### HasCall

`func (o *HandlersSecretTestConnectionResponse) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetCheckedAt

`func (o *HandlersSecretTestConnectionResponse) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *HandlersSecretTestConnectionResponse) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *HandlersSecretTestConnectionResponse) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *HandlersSecretTestConnectionResponse) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetDetail

`func (o *HandlersSecretTestConnectionResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *HandlersSecretTestConnectionResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *HandlersSecretTestConnectionResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *HandlersSecretTestConnectionResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetError

`func (o *HandlersSecretTestConnectionResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HandlersSecretTestConnectionResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HandlersSecretTestConnectionResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HandlersSecretTestConnectionResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLastVerifiedAt

`func (o *HandlersSecretTestConnectionResponse) GetLastVerifiedAt() string`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *HandlersSecretTestConnectionResponse) GetLastVerifiedAtOk() (*string, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *HandlersSecretTestConnectionResponse) SetLastVerifiedAt(v string)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *HandlersSecretTestConnectionResponse) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### GetLatencyMs

`func (o *HandlersSecretTestConnectionResponse) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *HandlersSecretTestConnectionResponse) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *HandlersSecretTestConnectionResponse) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *HandlersSecretTestConnectionResponse) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetOk

`func (o *HandlersSecretTestConnectionResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *HandlersSecretTestConnectionResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *HandlersSecretTestConnectionResponse) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *HandlersSecretTestConnectionResponse) HasOk() bool`

HasOk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


