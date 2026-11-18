# HandlersRunnerInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenWindow** | Pointer to [**HandlersRunnerInsightsWindowStat**](HandlersRunnerInsightsWindowStat.md) | OpenWindow is the open window&#39;s live percentile answer; nil when nothing qualifies (an idle org). | [optional] 
**Profile** | Pointer to **string** | Profile echoes the ?profile&#x3D; filter (&#39;&#39; &#x3D; org-wide). | [optional] 
**Reason** | Pointer to **string** | Reason carries \&quot;no data\&quot; when the requested profile (or the org) holds no insights rows at all; empty otherwise. | [optional] 
**Series** | Pointer to [**[]HandlersRunnerInsightsSeriesPoint**](HandlersRunnerInsightsSeriesPoint.md) | Series is the concurrency series, raw for the most recent 24 hours and hourly for every older rolled hour (newest first). | [optional] 
**Windows** | Pointer to [**[]HandlersRunnerInsightsWindowStat**](HandlersRunnerInsightsWindowStat.md) | Windows are the pinned windows (newest first, capped at the most recent 168 WINDOWS — whole windows, never split across metrics), or only ?window&#x3D;&#39;s hour when given. | [optional] 

## Methods

### NewHandlersRunnerInsightsResponse

`func NewHandlersRunnerInsightsResponse() *HandlersRunnerInsightsResponse`

NewHandlersRunnerInsightsResponse instantiates a new HandlersRunnerInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerInsightsResponseWithDefaults

`func NewHandlersRunnerInsightsResponseWithDefaults() *HandlersRunnerInsightsResponse`

NewHandlersRunnerInsightsResponseWithDefaults instantiates a new HandlersRunnerInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenWindow

`func (o *HandlersRunnerInsightsResponse) GetOpenWindow() HandlersRunnerInsightsWindowStat`

GetOpenWindow returns the OpenWindow field if non-nil, zero value otherwise.

### GetOpenWindowOk

`func (o *HandlersRunnerInsightsResponse) GetOpenWindowOk() (*HandlersRunnerInsightsWindowStat, bool)`

GetOpenWindowOk returns a tuple with the OpenWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWindow

`func (o *HandlersRunnerInsightsResponse) SetOpenWindow(v HandlersRunnerInsightsWindowStat)`

SetOpenWindow sets OpenWindow field to given value.

### HasOpenWindow

`func (o *HandlersRunnerInsightsResponse) HasOpenWindow() bool`

HasOpenWindow returns a boolean if a field has been set.

### GetProfile

`func (o *HandlersRunnerInsightsResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *HandlersRunnerInsightsResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *HandlersRunnerInsightsResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *HandlersRunnerInsightsResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetReason

`func (o *HandlersRunnerInsightsResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersRunnerInsightsResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersRunnerInsightsResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersRunnerInsightsResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSeries

`func (o *HandlersRunnerInsightsResponse) GetSeries() []HandlersRunnerInsightsSeriesPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *HandlersRunnerInsightsResponse) GetSeriesOk() (*[]HandlersRunnerInsightsSeriesPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *HandlersRunnerInsightsResponse) SetSeries(v []HandlersRunnerInsightsSeriesPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *HandlersRunnerInsightsResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetWindows

`func (o *HandlersRunnerInsightsResponse) GetWindows() []HandlersRunnerInsightsWindowStat`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *HandlersRunnerInsightsResponse) GetWindowsOk() (*[]HandlersRunnerInsightsWindowStat, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *HandlersRunnerInsightsResponse) SetWindows(v []HandlersRunnerInsightsWindowStat)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *HandlersRunnerInsightsResponse) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


