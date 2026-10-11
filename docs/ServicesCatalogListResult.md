# ServicesCatalogListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextLast** | Pointer to **string** | pass as last&#x3D; on next page when Link present / truncated | [optional] 
**Repositories** | Pointer to **[]string** |  | [optional] 
**Truncated** | Pointer to **bool** | true when q-filter hit MaxCatalogFilterPages | [optional] 

## Methods

### NewServicesCatalogListResult

`func NewServicesCatalogListResult() *ServicesCatalogListResult`

NewServicesCatalogListResult instantiates a new ServicesCatalogListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesCatalogListResultWithDefaults

`func NewServicesCatalogListResultWithDefaults() *ServicesCatalogListResult`

NewServicesCatalogListResultWithDefaults instantiates a new ServicesCatalogListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextLast

`func (o *ServicesCatalogListResult) GetNextLast() string`

GetNextLast returns the NextLast field if non-nil, zero value otherwise.

### GetNextLastOk

`func (o *ServicesCatalogListResult) GetNextLastOk() (*string, bool)`

GetNextLastOk returns a tuple with the NextLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLast

`func (o *ServicesCatalogListResult) SetNextLast(v string)`

SetNextLast sets NextLast field to given value.

### HasNextLast

`func (o *ServicesCatalogListResult) HasNextLast() bool`

HasNextLast returns a boolean if a field has been set.

### GetRepositories

`func (o *ServicesCatalogListResult) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ServicesCatalogListResult) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ServicesCatalogListResult) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ServicesCatalogListResult) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetTruncated

`func (o *ServicesCatalogListResult) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ServicesCatalogListResult) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ServicesCatalogListResult) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ServicesCatalogListResult) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


