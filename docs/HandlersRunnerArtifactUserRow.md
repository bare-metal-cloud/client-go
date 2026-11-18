# HandlersRunnerArtifactUserRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DigestPrefix** | Pointer to **string** |  | [optional] 
**FinalizedAt** | Pointer to **string** | FinalizedAt is the finalize stamp (nil while created). | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**JobId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**RetentionDays** | Pointer to **int32** | RetentionDays is the row&#39;s EFFECTIVE lifetime in days — the override when set, else the org policy&#39;s days mode; nil when the org policy is delete-when-run-ends and the row carries no override (the artifact lives until the run ends per the mode). | [optional] 
**RetentionSource** | Pointer to **string** | RetentionSource is \&quot;artifact\&quot; or \&quot;org\&quot; (the vocabulary above). | [optional] 
**RunId** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerArtifactUserRow

`func NewHandlersRunnerArtifactUserRow() *HandlersRunnerArtifactUserRow`

NewHandlersRunnerArtifactUserRow instantiates a new HandlersRunnerArtifactUserRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerArtifactUserRowWithDefaults

`func NewHandlersRunnerArtifactUserRowWithDefaults() *HandlersRunnerArtifactUserRow`

NewHandlersRunnerArtifactUserRowWithDefaults instantiates a new HandlersRunnerArtifactUserRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HandlersRunnerArtifactUserRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersRunnerArtifactUserRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersRunnerArtifactUserRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersRunnerArtifactUserRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDigestPrefix

`func (o *HandlersRunnerArtifactUserRow) GetDigestPrefix() string`

GetDigestPrefix returns the DigestPrefix field if non-nil, zero value otherwise.

### GetDigestPrefixOk

`func (o *HandlersRunnerArtifactUserRow) GetDigestPrefixOk() (*string, bool)`

GetDigestPrefixOk returns a tuple with the DigestPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestPrefix

`func (o *HandlersRunnerArtifactUserRow) SetDigestPrefix(v string)`

SetDigestPrefix sets DigestPrefix field to given value.

### HasDigestPrefix

`func (o *HandlersRunnerArtifactUserRow) HasDigestPrefix() bool`

HasDigestPrefix returns a boolean if a field has been set.

### GetFinalizedAt

`func (o *HandlersRunnerArtifactUserRow) GetFinalizedAt() string`

GetFinalizedAt returns the FinalizedAt field if non-nil, zero value otherwise.

### GetFinalizedAtOk

`func (o *HandlersRunnerArtifactUserRow) GetFinalizedAtOk() (*string, bool)`

GetFinalizedAtOk returns a tuple with the FinalizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedAt

`func (o *HandlersRunnerArtifactUserRow) SetFinalizedAt(v string)`

SetFinalizedAt sets FinalizedAt field to given value.

### HasFinalizedAt

`func (o *HandlersRunnerArtifactUserRow) HasFinalizedAt() bool`

HasFinalizedAt returns a boolean if a field has been set.

### GetId

`func (o *HandlersRunnerArtifactUserRow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersRunnerArtifactUserRow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersRunnerArtifactUserRow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersRunnerArtifactUserRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *HandlersRunnerArtifactUserRow) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HandlersRunnerArtifactUserRow) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HandlersRunnerArtifactUserRow) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HandlersRunnerArtifactUserRow) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerArtifactUserRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerArtifactUserRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerArtifactUserRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerArtifactUserRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepository

`func (o *HandlersRunnerArtifactUserRow) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HandlersRunnerArtifactUserRow) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HandlersRunnerArtifactUserRow) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HandlersRunnerArtifactUserRow) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRetentionDays

`func (o *HandlersRunnerArtifactUserRow) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *HandlersRunnerArtifactUserRow) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *HandlersRunnerArtifactUserRow) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *HandlersRunnerArtifactUserRow) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetRetentionSource

`func (o *HandlersRunnerArtifactUserRow) GetRetentionSource() string`

GetRetentionSource returns the RetentionSource field if non-nil, zero value otherwise.

### GetRetentionSourceOk

`func (o *HandlersRunnerArtifactUserRow) GetRetentionSourceOk() (*string, bool)`

GetRetentionSourceOk returns a tuple with the RetentionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSource

`func (o *HandlersRunnerArtifactUserRow) SetRetentionSource(v string)`

SetRetentionSource sets RetentionSource field to given value.

### HasRetentionSource

`func (o *HandlersRunnerArtifactUserRow) HasRetentionSource() bool`

HasRetentionSource returns a boolean if a field has been set.

### GetRunId

`func (o *HandlersRunnerArtifactUserRow) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *HandlersRunnerArtifactUserRow) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *HandlersRunnerArtifactUserRow) SetRunId(v int32)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *HandlersRunnerArtifactUserRow) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetSize

`func (o *HandlersRunnerArtifactUserRow) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HandlersRunnerArtifactUserRow) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HandlersRunnerArtifactUserRow) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *HandlersRunnerArtifactUserRow) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *HandlersRunnerArtifactUserRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersRunnerArtifactUserRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersRunnerArtifactUserRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersRunnerArtifactUserRow) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


