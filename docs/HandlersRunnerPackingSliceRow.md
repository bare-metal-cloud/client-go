# HandlersRunnerPackingSliceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundJobId** | Pointer to **string** | BoundJobID is the serving job&#39;s reference (empty &#x3D; the slice is idle stock). | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**DestroyPendingAt** | Pointer to **int32** | DestroyPendingAt/Channel are the spot arm&#39;s server-recorded reclaim marks (issue #4254): WHEN the interruption notice landed and WHICH channel delivered it (&#39;eventbridge&#39; — the control-plane SQS watcher — or &#39;metadata_poll&#39;, the guest-reported fallback). The page&#39;s reclaim caption names THIS record, never a guest claim. | [optional] 
**DestroyPendingChannel** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** | ProviderID is the slice&#39;s compound id (\&quot;&lt;host&gt;/&lt;slice-id&gt;\&quot;). | [optional] 
**Status** | Pointer to **string** | Status is the mirror&#39;s lifecycle word (the pool&#39;s state machine). | [optional] 
**Warm** | Pointer to **bool** | Warm marks warm-floor stock (the idle state IS its job). | [optional] 

## Methods

### NewHandlersRunnerPackingSliceRow

`func NewHandlersRunnerPackingSliceRow() *HandlersRunnerPackingSliceRow`

NewHandlersRunnerPackingSliceRow instantiates a new HandlersRunnerPackingSliceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerPackingSliceRowWithDefaults

`func NewHandlersRunnerPackingSliceRowWithDefaults() *HandlersRunnerPackingSliceRow`

NewHandlersRunnerPackingSliceRowWithDefaults instantiates a new HandlersRunnerPackingSliceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundJobId

`func (o *HandlersRunnerPackingSliceRow) GetBoundJobId() string`

GetBoundJobId returns the BoundJobId field if non-nil, zero value otherwise.

### GetBoundJobIdOk

`func (o *HandlersRunnerPackingSliceRow) GetBoundJobIdOk() (*string, bool)`

GetBoundJobIdOk returns a tuple with the BoundJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundJobId

`func (o *HandlersRunnerPackingSliceRow) SetBoundJobId(v string)`

SetBoundJobId sets BoundJobId field to given value.

### HasBoundJobId

`func (o *HandlersRunnerPackingSliceRow) HasBoundJobId() bool`

HasBoundJobId returns a boolean if a field has been set.

### GetClassName

`func (o *HandlersRunnerPackingSliceRow) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *HandlersRunnerPackingSliceRow) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *HandlersRunnerPackingSliceRow) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *HandlersRunnerPackingSliceRow) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersRunnerPackingSliceRow) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersRunnerPackingSliceRow) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersRunnerPackingSliceRow) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersRunnerPackingSliceRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestroyPendingAt

`func (o *HandlersRunnerPackingSliceRow) GetDestroyPendingAt() int32`

GetDestroyPendingAt returns the DestroyPendingAt field if non-nil, zero value otherwise.

### GetDestroyPendingAtOk

`func (o *HandlersRunnerPackingSliceRow) GetDestroyPendingAtOk() (*int32, bool)`

GetDestroyPendingAtOk returns a tuple with the DestroyPendingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyPendingAt

`func (o *HandlersRunnerPackingSliceRow) SetDestroyPendingAt(v int32)`

SetDestroyPendingAt sets DestroyPendingAt field to given value.

### HasDestroyPendingAt

`func (o *HandlersRunnerPackingSliceRow) HasDestroyPendingAt() bool`

HasDestroyPendingAt returns a boolean if a field has been set.

### GetDestroyPendingChannel

`func (o *HandlersRunnerPackingSliceRow) GetDestroyPendingChannel() string`

GetDestroyPendingChannel returns the DestroyPendingChannel field if non-nil, zero value otherwise.

### GetDestroyPendingChannelOk

`func (o *HandlersRunnerPackingSliceRow) GetDestroyPendingChannelOk() (*string, bool)`

GetDestroyPendingChannelOk returns a tuple with the DestroyPendingChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyPendingChannel

`func (o *HandlersRunnerPackingSliceRow) SetDestroyPendingChannel(v string)`

SetDestroyPendingChannel sets DestroyPendingChannel field to given value.

### HasDestroyPendingChannel

`func (o *HandlersRunnerPackingSliceRow) HasDestroyPendingChannel() bool`

HasDestroyPendingChannel returns a boolean if a field has been set.

### GetProviderId

`func (o *HandlersRunnerPackingSliceRow) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *HandlersRunnerPackingSliceRow) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *HandlersRunnerPackingSliceRow) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *HandlersRunnerPackingSliceRow) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersRunnerPackingSliceRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersRunnerPackingSliceRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersRunnerPackingSliceRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersRunnerPackingSliceRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWarm

`func (o *HandlersRunnerPackingSliceRow) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *HandlersRunnerPackingSliceRow) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *HandlersRunnerPackingSliceRow) SetWarm(v bool)`

SetWarm sets Warm field to given value.

### HasWarm

`func (o *HandlersRunnerPackingSliceRow) HasWarm() bool`

HasWarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


