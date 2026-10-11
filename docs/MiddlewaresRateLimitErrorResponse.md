# MiddlewaresRateLimitErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**LimitType** | Pointer to **string** |  | [optional] 
**MaxRequests** | Pointer to **int32** |  | [optional] 
**RateLimitPeriod** | Pointer to **string** |  | [optional] 
**RetryAfter** | Pointer to **int32** |  | [optional] 

## Methods

### NewMiddlewaresRateLimitErrorResponse

`func NewMiddlewaresRateLimitErrorResponse() *MiddlewaresRateLimitErrorResponse`

NewMiddlewaresRateLimitErrorResponse instantiates a new MiddlewaresRateLimitErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewaresRateLimitErrorResponseWithDefaults

`func NewMiddlewaresRateLimitErrorResponseWithDefaults() *MiddlewaresRateLimitErrorResponse`

NewMiddlewaresRateLimitErrorResponseWithDefaults instantiates a new MiddlewaresRateLimitErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MiddlewaresRateLimitErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MiddlewaresRateLimitErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MiddlewaresRateLimitErrorResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MiddlewaresRateLimitErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLimitType

`func (o *MiddlewaresRateLimitErrorResponse) GetLimitType() string`

GetLimitType returns the LimitType field if non-nil, zero value otherwise.

### GetLimitTypeOk

`func (o *MiddlewaresRateLimitErrorResponse) GetLimitTypeOk() (*string, bool)`

GetLimitTypeOk returns a tuple with the LimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitType

`func (o *MiddlewaresRateLimitErrorResponse) SetLimitType(v string)`

SetLimitType sets LimitType field to given value.

### HasLimitType

`func (o *MiddlewaresRateLimitErrorResponse) HasLimitType() bool`

HasLimitType returns a boolean if a field has been set.

### GetMaxRequests

`func (o *MiddlewaresRateLimitErrorResponse) GetMaxRequests() int32`

GetMaxRequests returns the MaxRequests field if non-nil, zero value otherwise.

### GetMaxRequestsOk

`func (o *MiddlewaresRateLimitErrorResponse) GetMaxRequestsOk() (*int32, bool)`

GetMaxRequestsOk returns a tuple with the MaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequests

`func (o *MiddlewaresRateLimitErrorResponse) SetMaxRequests(v int32)`

SetMaxRequests sets MaxRequests field to given value.

### HasMaxRequests

`func (o *MiddlewaresRateLimitErrorResponse) HasMaxRequests() bool`

HasMaxRequests returns a boolean if a field has been set.

### GetRateLimitPeriod

`func (o *MiddlewaresRateLimitErrorResponse) GetRateLimitPeriod() string`

GetRateLimitPeriod returns the RateLimitPeriod field if non-nil, zero value otherwise.

### GetRateLimitPeriodOk

`func (o *MiddlewaresRateLimitErrorResponse) GetRateLimitPeriodOk() (*string, bool)`

GetRateLimitPeriodOk returns a tuple with the RateLimitPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPeriod

`func (o *MiddlewaresRateLimitErrorResponse) SetRateLimitPeriod(v string)`

SetRateLimitPeriod sets RateLimitPeriod field to given value.

### HasRateLimitPeriod

`func (o *MiddlewaresRateLimitErrorResponse) HasRateLimitPeriod() bool`

HasRateLimitPeriod returns a boolean if a field has been set.

### GetRetryAfter

`func (o *MiddlewaresRateLimitErrorResponse) GetRetryAfter() int32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *MiddlewaresRateLimitErrorResponse) GetRetryAfterOk() (*int32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *MiddlewaresRateLimitErrorResponse) SetRetryAfter(v int32)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *MiddlewaresRateLimitErrorResponse) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


