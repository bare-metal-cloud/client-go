# HandlersDetailedHealthCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**map[string]HandlersCheckDetail**](HandlersCheckDetail.md) |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDetailedHealthCheckResponse

`func NewHandlersDetailedHealthCheckResponse() *HandlersDetailedHealthCheckResponse`

NewHandlersDetailedHealthCheckResponse instantiates a new HandlersDetailedHealthCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDetailedHealthCheckResponseWithDefaults

`func NewHandlersDetailedHealthCheckResponseWithDefaults() *HandlersDetailedHealthCheckResponse`

NewHandlersDetailedHealthCheckResponseWithDefaults instantiates a new HandlersDetailedHealthCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *HandlersDetailedHealthCheckResponse) GetChecks() map[string]HandlersCheckDetail`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *HandlersDetailedHealthCheckResponse) GetChecksOk() (*map[string]HandlersCheckDetail, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *HandlersDetailedHealthCheckResponse) SetChecks(v map[string]HandlersCheckDetail)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *HandlersDetailedHealthCheckResponse) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetService

`func (o *HandlersDetailedHealthCheckResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HandlersDetailedHealthCheckResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HandlersDetailedHealthCheckResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HandlersDetailedHealthCheckResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersDetailedHealthCheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersDetailedHealthCheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersDetailedHealthCheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersDetailedHealthCheckResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *HandlersDetailedHealthCheckResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HandlersDetailedHealthCheckResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HandlersDetailedHealthCheckResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HandlersDetailedHealthCheckResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVersion

`func (o *HandlersDetailedHealthCheckResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersDetailedHealthCheckResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersDetailedHealthCheckResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HandlersDetailedHealthCheckResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


