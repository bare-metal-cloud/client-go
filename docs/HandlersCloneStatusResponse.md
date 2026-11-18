# HandlersCloneStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | Pointer to [**CloneplanSummary**](CloneplanSummary.md) |  | [optional] 
**Events** | Pointer to [**[]EngineportsInstallProgressEvent**](EngineportsInstallProgressEvent.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Report** | Pointer to [**ClonereplayReport**](ClonereplayReport.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCloneStatusResponse

`func NewHandlersCloneStatusResponse() *HandlersCloneStatusResponse`

NewHandlersCloneStatusResponse instantiates a new HandlersCloneStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCloneStatusResponseWithDefaults

`func NewHandlersCloneStatusResponseWithDefaults() *HandlersCloneStatusResponse`

NewHandlersCloneStatusResponseWithDefaults instantiates a new HandlersCloneStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *HandlersCloneStatusResponse) GetAggregate() CloneplanSummary`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *HandlersCloneStatusResponse) GetAggregateOk() (*CloneplanSummary, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *HandlersCloneStatusResponse) SetAggregate(v CloneplanSummary)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *HandlersCloneStatusResponse) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.

### GetEvents

`func (o *HandlersCloneStatusResponse) GetEvents() []EngineportsInstallProgressEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HandlersCloneStatusResponse) GetEventsOk() (*[]EngineportsInstallProgressEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HandlersCloneStatusResponse) SetEvents(v []EngineportsInstallProgressEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HandlersCloneStatusResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetPlanId

`func (o *HandlersCloneStatusResponse) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *HandlersCloneStatusResponse) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *HandlersCloneStatusResponse) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *HandlersCloneStatusResponse) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetReport

`func (o *HandlersCloneStatusResponse) GetReport() ClonereplayReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *HandlersCloneStatusResponse) GetReportOk() (*ClonereplayReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *HandlersCloneStatusResponse) SetReport(v ClonereplayReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *HandlersCloneStatusResponse) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetState

`func (o *HandlersCloneStatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersCloneStatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersCloneStatusResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersCloneStatusResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkflowId

`func (o *HandlersCloneStatusResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *HandlersCloneStatusResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *HandlersCloneStatusResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *HandlersCloneStatusResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


