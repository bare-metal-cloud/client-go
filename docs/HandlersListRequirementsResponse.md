# HandlersListRequirementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingCount** | Pointer to **int32** |  | [optional] 
**Requirements** | Pointer to [**[]ServicesUserRequirementStatus**](ServicesUserRequirementStatus.md) |  | [optional] 

## Methods

### NewHandlersListRequirementsResponse

`func NewHandlersListRequirementsResponse() *HandlersListRequirementsResponse`

NewHandlersListRequirementsResponse instantiates a new HandlersListRequirementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersListRequirementsResponseWithDefaults

`func NewHandlersListRequirementsResponseWithDefaults() *HandlersListRequirementsResponse`

NewHandlersListRequirementsResponseWithDefaults instantiates a new HandlersListRequirementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingCount

`func (o *HandlersListRequirementsResponse) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *HandlersListRequirementsResponse) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *HandlersListRequirementsResponse) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *HandlersListRequirementsResponse) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetRequirements

`func (o *HandlersListRequirementsResponse) GetRequirements() []ServicesUserRequirementStatus`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *HandlersListRequirementsResponse) GetRequirementsOk() (*[]ServicesUserRequirementStatus, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *HandlersListRequirementsResponse) SetRequirements(v []ServicesUserRequirementStatus)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *HandlersListRequirementsResponse) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


