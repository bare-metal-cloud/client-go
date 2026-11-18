# HandlersRunnerInsightsWindowStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationP50Ms** | Pointer to **int32** | The four percentile values, milliseconds. Nil &#x3D; no qualifying rows for that metric. | [optional] 
**DurationP95Ms** | Pointer to **int32** |  | [optional] 
**PinnedAt** | Pointer to **string** | PinnedAt is the pin stamp on a pinned window; nil on the live open window. | [optional] 
**Profile** | Pointer to **string** | Profile is the stat&#39;s profile (&#39;&#39; &#x3D; unprofiled / org-wide). | [optional] 
**QueueP50Ms** | Pointer to **int32** |  | [optional] 
**QueueP95Ms** | Pointer to **int32** |  | [optional] 
**WindowStart** | Pointer to **string** | WindowStart is the window the stat describes (UTC hour). | [optional] 

## Methods

### NewHandlersRunnerInsightsWindowStat

`func NewHandlersRunnerInsightsWindowStat() *HandlersRunnerInsightsWindowStat`

NewHandlersRunnerInsightsWindowStat instantiates a new HandlersRunnerInsightsWindowStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerInsightsWindowStatWithDefaults

`func NewHandlersRunnerInsightsWindowStatWithDefaults() *HandlersRunnerInsightsWindowStat`

NewHandlersRunnerInsightsWindowStatWithDefaults instantiates a new HandlersRunnerInsightsWindowStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationP50Ms

`func (o *HandlersRunnerInsightsWindowStat) GetDurationP50Ms() int32`

GetDurationP50Ms returns the DurationP50Ms field if non-nil, zero value otherwise.

### GetDurationP50MsOk

`func (o *HandlersRunnerInsightsWindowStat) GetDurationP50MsOk() (*int32, bool)`

GetDurationP50MsOk returns a tuple with the DurationP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationP50Ms

`func (o *HandlersRunnerInsightsWindowStat) SetDurationP50Ms(v int32)`

SetDurationP50Ms sets DurationP50Ms field to given value.

### HasDurationP50Ms

`func (o *HandlersRunnerInsightsWindowStat) HasDurationP50Ms() bool`

HasDurationP50Ms returns a boolean if a field has been set.

### GetDurationP95Ms

`func (o *HandlersRunnerInsightsWindowStat) GetDurationP95Ms() int32`

GetDurationP95Ms returns the DurationP95Ms field if non-nil, zero value otherwise.

### GetDurationP95MsOk

`func (o *HandlersRunnerInsightsWindowStat) GetDurationP95MsOk() (*int32, bool)`

GetDurationP95MsOk returns a tuple with the DurationP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationP95Ms

`func (o *HandlersRunnerInsightsWindowStat) SetDurationP95Ms(v int32)`

SetDurationP95Ms sets DurationP95Ms field to given value.

### HasDurationP95Ms

`func (o *HandlersRunnerInsightsWindowStat) HasDurationP95Ms() bool`

HasDurationP95Ms returns a boolean if a field has been set.

### GetPinnedAt

`func (o *HandlersRunnerInsightsWindowStat) GetPinnedAt() string`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *HandlersRunnerInsightsWindowStat) GetPinnedAtOk() (*string, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *HandlersRunnerInsightsWindowStat) SetPinnedAt(v string)`

SetPinnedAt sets PinnedAt field to given value.

### HasPinnedAt

`func (o *HandlersRunnerInsightsWindowStat) HasPinnedAt() bool`

HasPinnedAt returns a boolean if a field has been set.

### GetProfile

`func (o *HandlersRunnerInsightsWindowStat) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *HandlersRunnerInsightsWindowStat) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *HandlersRunnerInsightsWindowStat) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *HandlersRunnerInsightsWindowStat) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQueueP50Ms

`func (o *HandlersRunnerInsightsWindowStat) GetQueueP50Ms() int32`

GetQueueP50Ms returns the QueueP50Ms field if non-nil, zero value otherwise.

### GetQueueP50MsOk

`func (o *HandlersRunnerInsightsWindowStat) GetQueueP50MsOk() (*int32, bool)`

GetQueueP50MsOk returns a tuple with the QueueP50Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueP50Ms

`func (o *HandlersRunnerInsightsWindowStat) SetQueueP50Ms(v int32)`

SetQueueP50Ms sets QueueP50Ms field to given value.

### HasQueueP50Ms

`func (o *HandlersRunnerInsightsWindowStat) HasQueueP50Ms() bool`

HasQueueP50Ms returns a boolean if a field has been set.

### GetQueueP95Ms

`func (o *HandlersRunnerInsightsWindowStat) GetQueueP95Ms() int32`

GetQueueP95Ms returns the QueueP95Ms field if non-nil, zero value otherwise.

### GetQueueP95MsOk

`func (o *HandlersRunnerInsightsWindowStat) GetQueueP95MsOk() (*int32, bool)`

GetQueueP95MsOk returns a tuple with the QueueP95Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueP95Ms

`func (o *HandlersRunnerInsightsWindowStat) SetQueueP95Ms(v int32)`

SetQueueP95Ms sets QueueP95Ms field to given value.

### HasQueueP95Ms

`func (o *HandlersRunnerInsightsWindowStat) HasQueueP95Ms() bool`

HasQueueP95Ms returns a boolean if a field has been set.

### GetWindowStart

`func (o *HandlersRunnerInsightsWindowStat) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *HandlersRunnerInsightsWindowStat) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *HandlersRunnerInsightsWindowStat) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.

### HasWindowStart

`func (o *HandlersRunnerInsightsWindowStat) HasWindowStart() bool`

HasWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


