# HandlersRunnerConnectionRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the repository name. | [optional] 
**Orphaned** | Pointer to **bool** | Orphaned: the connection&#39;s state is suspended/removed — the mapping is reported, and placement must never consume it. | [optional] 
**ProfileArchived** | Pointer to **bool** | ProfileArchived: the mapped profile row is archived — reported, never silently dropped (the release-archive case stays visible). | [optional] 
**ProfileName** | Pointer to **string** | ProfileName is the mapped profile&#39;s runs-on name. | [optional] 

## Methods

### NewHandlersRunnerConnectionRepositoryResponse

`func NewHandlersRunnerConnectionRepositoryResponse() *HandlersRunnerConnectionRepositoryResponse`

NewHandlersRunnerConnectionRepositoryResponse instantiates a new HandlersRunnerConnectionRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerConnectionRepositoryResponseWithDefaults

`func NewHandlersRunnerConnectionRepositoryResponseWithDefaults() *HandlersRunnerConnectionRepositoryResponse`

NewHandlersRunnerConnectionRepositoryResponseWithDefaults instantiates a new HandlersRunnerConnectionRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HandlersRunnerConnectionRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerConnectionRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerConnectionRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerConnectionRepositoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrphaned

`func (o *HandlersRunnerConnectionRepositoryResponse) GetOrphaned() bool`

GetOrphaned returns the Orphaned field if non-nil, zero value otherwise.

### GetOrphanedOk

`func (o *HandlersRunnerConnectionRepositoryResponse) GetOrphanedOk() (*bool, bool)`

GetOrphanedOk returns a tuple with the Orphaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphaned

`func (o *HandlersRunnerConnectionRepositoryResponse) SetOrphaned(v bool)`

SetOrphaned sets Orphaned field to given value.

### HasOrphaned

`func (o *HandlersRunnerConnectionRepositoryResponse) HasOrphaned() bool`

HasOrphaned returns a boolean if a field has been set.

### GetProfileArchived

`func (o *HandlersRunnerConnectionRepositoryResponse) GetProfileArchived() bool`

GetProfileArchived returns the ProfileArchived field if non-nil, zero value otherwise.

### GetProfileArchivedOk

`func (o *HandlersRunnerConnectionRepositoryResponse) GetProfileArchivedOk() (*bool, bool)`

GetProfileArchivedOk returns a tuple with the ProfileArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileArchived

`func (o *HandlersRunnerConnectionRepositoryResponse) SetProfileArchived(v bool)`

SetProfileArchived sets ProfileArchived field to given value.

### HasProfileArchived

`func (o *HandlersRunnerConnectionRepositoryResponse) HasProfileArchived() bool`

HasProfileArchived returns a boolean if a field has been set.

### GetProfileName

`func (o *HandlersRunnerConnectionRepositoryResponse) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *HandlersRunnerConnectionRepositoryResponse) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *HandlersRunnerConnectionRepositoryResponse) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *HandlersRunnerConnectionRepositoryResponse) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


