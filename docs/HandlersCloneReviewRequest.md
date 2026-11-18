# HandlersCloneReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to **[]string** |  | [optional] 
**EnvLocal** | Pointer to [**[]CloneplanEnvLocalApproval**](CloneplanEnvLocalApproval.md) |  | [optional] 
**Excluded** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHandlersCloneReviewRequest

`func NewHandlersCloneReviewRequest() *HandlersCloneReviewRequest`

NewHandlersCloneReviewRequest instantiates a new HandlersCloneReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCloneReviewRequestWithDefaults

`func NewHandlersCloneReviewRequestWithDefaults() *HandlersCloneReviewRequest`

NewHandlersCloneReviewRequestWithDefaults instantiates a new HandlersCloneReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *HandlersCloneReviewRequest) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *HandlersCloneReviewRequest) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *HandlersCloneReviewRequest) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *HandlersCloneReviewRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEnvLocal

`func (o *HandlersCloneReviewRequest) GetEnvLocal() []CloneplanEnvLocalApproval`

GetEnvLocal returns the EnvLocal field if non-nil, zero value otherwise.

### GetEnvLocalOk

`func (o *HandlersCloneReviewRequest) GetEnvLocalOk() (*[]CloneplanEnvLocalApproval, bool)`

GetEnvLocalOk returns a tuple with the EnvLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvLocal

`func (o *HandlersCloneReviewRequest) SetEnvLocal(v []CloneplanEnvLocalApproval)`

SetEnvLocal sets EnvLocal field to given value.

### HasEnvLocal

`func (o *HandlersCloneReviewRequest) HasEnvLocal() bool`

HasEnvLocal returns a boolean if a field has been set.

### GetExcluded

`func (o *HandlersCloneReviewRequest) GetExcluded() []string`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *HandlersCloneReviewRequest) GetExcludedOk() (*[]string, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *HandlersCloneReviewRequest) SetExcluded(v []string)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *HandlersCloneReviewRequest) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


