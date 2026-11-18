# HandlersClonePlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | Pointer to [**CloneplanSummary**](CloneplanSummary.md) |  | [optional] 
**Plan** | Pointer to [**CloneplanPlan**](CloneplanPlan.md) |  | [optional] 
**Review** | Pointer to [**CloneplanReview**](CloneplanReview.md) |  | [optional] 

## Methods

### NewHandlersClonePlanResponse

`func NewHandlersClonePlanResponse() *HandlersClonePlanResponse`

NewHandlersClonePlanResponse instantiates a new HandlersClonePlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersClonePlanResponseWithDefaults

`func NewHandlersClonePlanResponseWithDefaults() *HandlersClonePlanResponse`

NewHandlersClonePlanResponseWithDefaults instantiates a new HandlersClonePlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *HandlersClonePlanResponse) GetAggregate() CloneplanSummary`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *HandlersClonePlanResponse) GetAggregateOk() (*CloneplanSummary, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *HandlersClonePlanResponse) SetAggregate(v CloneplanSummary)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *HandlersClonePlanResponse) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.

### GetPlan

`func (o *HandlersClonePlanResponse) GetPlan() CloneplanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *HandlersClonePlanResponse) GetPlanOk() (*CloneplanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *HandlersClonePlanResponse) SetPlan(v CloneplanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *HandlersClonePlanResponse) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetReview

`func (o *HandlersClonePlanResponse) GetReview() CloneplanReview`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *HandlersClonePlanResponse) GetReviewOk() (*CloneplanReview, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *HandlersClonePlanResponse) SetReview(v CloneplanReview)`

SetReview sets Review field to given value.

### HasReview

`func (o *HandlersClonePlanResponse) HasReview() bool`

HasReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


