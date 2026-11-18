# GetApiV1AdminOverviewSensors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsCollector** | Pointer to **map[string]interface{}** |  | [optional] 
**Sensors** | Pointer to [**[]HandlersOverviewSensorStamp**](HandlersOverviewSensorStamp.md) |  | [optional] 
**TransitionDeadman** | Pointer to [**HandlersOverviewDeadmanBoard**](HandlersOverviewDeadmanBoard.md) |  | [optional] 
**VmWatchdogBells** | Pointer to **map[string]interface{}** |  | [optional] 
**WatchdogLastRun** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetApiV1AdminOverviewSensors200Response

`func NewGetApiV1AdminOverviewSensors200Response() *GetApiV1AdminOverviewSensors200Response`

NewGetApiV1AdminOverviewSensors200Response instantiates a new GetApiV1AdminOverviewSensors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiV1AdminOverviewSensors200ResponseWithDefaults

`func NewGetApiV1AdminOverviewSensors200ResponseWithDefaults() *GetApiV1AdminOverviewSensors200Response`

NewGetApiV1AdminOverviewSensors200ResponseWithDefaults instantiates a new GetApiV1AdminOverviewSensors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsCollector

`func (o *GetApiV1AdminOverviewSensors200Response) GetMetricsCollector() map[string]interface{}`

GetMetricsCollector returns the MetricsCollector field if non-nil, zero value otherwise.

### GetMetricsCollectorOk

`func (o *GetApiV1AdminOverviewSensors200Response) GetMetricsCollectorOk() (*map[string]interface{}, bool)`

GetMetricsCollectorOk returns a tuple with the MetricsCollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsCollector

`func (o *GetApiV1AdminOverviewSensors200Response) SetMetricsCollector(v map[string]interface{})`

SetMetricsCollector sets MetricsCollector field to given value.

### HasMetricsCollector

`func (o *GetApiV1AdminOverviewSensors200Response) HasMetricsCollector() bool`

HasMetricsCollector returns a boolean if a field has been set.

### GetSensors

`func (o *GetApiV1AdminOverviewSensors200Response) GetSensors() []HandlersOverviewSensorStamp`

GetSensors returns the Sensors field if non-nil, zero value otherwise.

### GetSensorsOk

`func (o *GetApiV1AdminOverviewSensors200Response) GetSensorsOk() (*[]HandlersOverviewSensorStamp, bool)`

GetSensorsOk returns a tuple with the Sensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensors

`func (o *GetApiV1AdminOverviewSensors200Response) SetSensors(v []HandlersOverviewSensorStamp)`

SetSensors sets Sensors field to given value.

### HasSensors

`func (o *GetApiV1AdminOverviewSensors200Response) HasSensors() bool`

HasSensors returns a boolean if a field has been set.

### GetTransitionDeadman

`func (o *GetApiV1AdminOverviewSensors200Response) GetTransitionDeadman() HandlersOverviewDeadmanBoard`

GetTransitionDeadman returns the TransitionDeadman field if non-nil, zero value otherwise.

### GetTransitionDeadmanOk

`func (o *GetApiV1AdminOverviewSensors200Response) GetTransitionDeadmanOk() (*HandlersOverviewDeadmanBoard, bool)`

GetTransitionDeadmanOk returns a tuple with the TransitionDeadman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionDeadman

`func (o *GetApiV1AdminOverviewSensors200Response) SetTransitionDeadman(v HandlersOverviewDeadmanBoard)`

SetTransitionDeadman sets TransitionDeadman field to given value.

### HasTransitionDeadman

`func (o *GetApiV1AdminOverviewSensors200Response) HasTransitionDeadman() bool`

HasTransitionDeadman returns a boolean if a field has been set.

### GetVmWatchdogBells

`func (o *GetApiV1AdminOverviewSensors200Response) GetVmWatchdogBells() map[string]interface{}`

GetVmWatchdogBells returns the VmWatchdogBells field if non-nil, zero value otherwise.

### GetVmWatchdogBellsOk

`func (o *GetApiV1AdminOverviewSensors200Response) GetVmWatchdogBellsOk() (*map[string]interface{}, bool)`

GetVmWatchdogBellsOk returns a tuple with the VmWatchdogBells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmWatchdogBells

`func (o *GetApiV1AdminOverviewSensors200Response) SetVmWatchdogBells(v map[string]interface{})`

SetVmWatchdogBells sets VmWatchdogBells field to given value.

### HasVmWatchdogBells

`func (o *GetApiV1AdminOverviewSensors200Response) HasVmWatchdogBells() bool`

HasVmWatchdogBells returns a boolean if a field has been set.

### GetWatchdogLastRun

`func (o *GetApiV1AdminOverviewSensors200Response) GetWatchdogLastRun() map[string]interface{}`

GetWatchdogLastRun returns the WatchdogLastRun field if non-nil, zero value otherwise.

### GetWatchdogLastRunOk

`func (o *GetApiV1AdminOverviewSensors200Response) GetWatchdogLastRunOk() (*map[string]interface{}, bool)`

GetWatchdogLastRunOk returns a tuple with the WatchdogLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdogLastRun

`func (o *GetApiV1AdminOverviewSensors200Response) SetWatchdogLastRun(v map[string]interface{})`

SetWatchdogLastRun sets WatchdogLastRun field to given value.

### HasWatchdogLastRun

`func (o *GetApiV1AdminOverviewSensors200Response) HasWatchdogLastRun() bool`

HasWatchdogLastRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


