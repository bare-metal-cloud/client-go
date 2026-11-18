# HandlersJobListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]ServicesJobListItem**](ServicesJobListItem.md) |  | [optional] 
**Pagination** | Pointer to [**ServicesPaginationMetadata**](ServicesPaginationMetadata.md) |  | [optional] 

## Methods

### NewHandlersJobListResponse

`func NewHandlersJobListResponse() *HandlersJobListResponse`

NewHandlersJobListResponse instantiates a new HandlersJobListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersJobListResponseWithDefaults

`func NewHandlersJobListResponseWithDefaults() *HandlersJobListResponse`

NewHandlersJobListResponseWithDefaults instantiates a new HandlersJobListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *HandlersJobListResponse) GetJobs() []ServicesJobListItem`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *HandlersJobListResponse) GetJobsOk() (*[]ServicesJobListItem, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *HandlersJobListResponse) SetJobs(v []ServicesJobListItem)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *HandlersJobListResponse) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetPagination

`func (o *HandlersJobListResponse) GetPagination() ServicesPaginationMetadata`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *HandlersJobListResponse) GetPaginationOk() (*ServicesPaginationMetadata, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *HandlersJobListResponse) SetPagination(v ServicesPaginationMetadata)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *HandlersJobListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


