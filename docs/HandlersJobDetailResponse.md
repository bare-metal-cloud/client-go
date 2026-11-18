# HandlersJobDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ModelsJobMetadata**](ModelsJobMetadata.md) |  | [optional] 
**Steps** | Pointer to [**[]ModelsJobStep**](ModelsJobStep.md) |  | [optional] 

## Methods

### NewHandlersJobDetailResponse

`func NewHandlersJobDetailResponse() *HandlersJobDetailResponse`

NewHandlersJobDetailResponse instantiates a new HandlersJobDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersJobDetailResponseWithDefaults

`func NewHandlersJobDetailResponseWithDefaults() *HandlersJobDetailResponse`

NewHandlersJobDetailResponseWithDefaults instantiates a new HandlersJobDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *HandlersJobDetailResponse) GetMetadata() ModelsJobMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HandlersJobDetailResponse) GetMetadataOk() (*ModelsJobMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HandlersJobDetailResponse) SetMetadata(v ModelsJobMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HandlersJobDetailResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSteps

`func (o *HandlersJobDetailResponse) GetSteps() []ModelsJobStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *HandlersJobDetailResponse) GetStepsOk() (*[]ModelsJobStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *HandlersJobDetailResponse) SetSteps(v []ModelsJobStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *HandlersJobDetailResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


