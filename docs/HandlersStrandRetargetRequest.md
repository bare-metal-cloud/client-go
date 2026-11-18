# HandlersStrandRetargetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Reviewed** | Pointer to [**[]EngineportsStrandRetargetReview**](EngineportsStrandRetargetReview.md) |  | [optional] 

## Methods

### NewHandlersStrandRetargetRequest

`func NewHandlersStrandRetargetRequest() *HandlersStrandRetargetRequest`

NewHandlersStrandRetargetRequest instantiates a new HandlersStrandRetargetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStrandRetargetRequestWithDefaults

`func NewHandlersStrandRetargetRequestWithDefaults() *HandlersStrandRetargetRequest`

NewHandlersStrandRetargetRequestWithDefaults instantiates a new HandlersStrandRetargetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *HandlersStrandRetargetRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersStrandRetargetRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersStrandRetargetRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersStrandRetargetRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetReviewed

`func (o *HandlersStrandRetargetRequest) GetReviewed() []EngineportsStrandRetargetReview`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *HandlersStrandRetargetRequest) GetReviewedOk() (*[]EngineportsStrandRetargetReview, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *HandlersStrandRetargetRequest) SetReviewed(v []EngineportsStrandRetargetReview)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *HandlersStrandRetargetRequest) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


