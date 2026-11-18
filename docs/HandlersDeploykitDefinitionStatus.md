# HandlersDeploykitDefinitionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestDigest** | Pointer to **string** |  | [optional] 
**LatestRevision** | Pointer to **int32** |  | [optional] 
**OrphanedSince** | Pointer to **string** | OrphanedSince carries the orphaned state&#39;s since stamp (zero when connected). | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceRepo** | Pointer to **string** |  | [optional] 
**SourceSha** | Pointer to **string** |  | [optional] 
**SourceState** | Pointer to **string** | SourceState is \&quot;connected\&quot; or the named \&quot;orphaned\&quot; state (D4). | [optional] 

## Methods

### NewHandlersDeploykitDefinitionStatus

`func NewHandlersDeploykitDefinitionStatus() *HandlersDeploykitDefinitionStatus`

NewHandlersDeploykitDefinitionStatus instantiates a new HandlersDeploykitDefinitionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDeploykitDefinitionStatusWithDefaults

`func NewHandlersDeploykitDefinitionStatusWithDefaults() *HandlersDeploykitDefinitionStatus`

NewHandlersDeploykitDefinitionStatusWithDefaults instantiates a new HandlersDeploykitDefinitionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestDigest

`func (o *HandlersDeploykitDefinitionStatus) GetLatestDigest() string`

GetLatestDigest returns the LatestDigest field if non-nil, zero value otherwise.

### GetLatestDigestOk

`func (o *HandlersDeploykitDefinitionStatus) GetLatestDigestOk() (*string, bool)`

GetLatestDigestOk returns a tuple with the LatestDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDigest

`func (o *HandlersDeploykitDefinitionStatus) SetLatestDigest(v string)`

SetLatestDigest sets LatestDigest field to given value.

### HasLatestDigest

`func (o *HandlersDeploykitDefinitionStatus) HasLatestDigest() bool`

HasLatestDigest returns a boolean if a field has been set.

### GetLatestRevision

`func (o *HandlersDeploykitDefinitionStatus) GetLatestRevision() int32`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *HandlersDeploykitDefinitionStatus) GetLatestRevisionOk() (*int32, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *HandlersDeploykitDefinitionStatus) SetLatestRevision(v int32)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *HandlersDeploykitDefinitionStatus) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetOrphanedSince

`func (o *HandlersDeploykitDefinitionStatus) GetOrphanedSince() string`

GetOrphanedSince returns the OrphanedSince field if non-nil, zero value otherwise.

### GetOrphanedSinceOk

`func (o *HandlersDeploykitDefinitionStatus) GetOrphanedSinceOk() (*string, bool)`

GetOrphanedSinceOk returns a tuple with the OrphanedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedSince

`func (o *HandlersDeploykitDefinitionStatus) SetOrphanedSince(v string)`

SetOrphanedSince sets OrphanedSince field to given value.

### HasOrphanedSince

`func (o *HandlersDeploykitDefinitionStatus) HasOrphanedSince() bool`

HasOrphanedSince returns a boolean if a field has been set.

### GetSource

`func (o *HandlersDeploykitDefinitionStatus) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HandlersDeploykitDefinitionStatus) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HandlersDeploykitDefinitionStatus) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HandlersDeploykitDefinitionStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceRepo

`func (o *HandlersDeploykitDefinitionStatus) GetSourceRepo() string`

GetSourceRepo returns the SourceRepo field if non-nil, zero value otherwise.

### GetSourceRepoOk

`func (o *HandlersDeploykitDefinitionStatus) GetSourceRepoOk() (*string, bool)`

GetSourceRepoOk returns a tuple with the SourceRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRepo

`func (o *HandlersDeploykitDefinitionStatus) SetSourceRepo(v string)`

SetSourceRepo sets SourceRepo field to given value.

### HasSourceRepo

`func (o *HandlersDeploykitDefinitionStatus) HasSourceRepo() bool`

HasSourceRepo returns a boolean if a field has been set.

### GetSourceSha

`func (o *HandlersDeploykitDefinitionStatus) GetSourceSha() string`

GetSourceSha returns the SourceSha field if non-nil, zero value otherwise.

### GetSourceShaOk

`func (o *HandlersDeploykitDefinitionStatus) GetSourceShaOk() (*string, bool)`

GetSourceShaOk returns a tuple with the SourceSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSha

`func (o *HandlersDeploykitDefinitionStatus) SetSourceSha(v string)`

SetSourceSha sets SourceSha field to given value.

### HasSourceSha

`func (o *HandlersDeploykitDefinitionStatus) HasSourceSha() bool`

HasSourceSha returns a boolean if a field has been set.

### GetSourceState

`func (o *HandlersDeploykitDefinitionStatus) GetSourceState() string`

GetSourceState returns the SourceState field if non-nil, zero value otherwise.

### GetSourceStateOk

`func (o *HandlersDeploykitDefinitionStatus) GetSourceStateOk() (*string, bool)`

GetSourceStateOk returns a tuple with the SourceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceState

`func (o *HandlersDeploykitDefinitionStatus) SetSourceState(v string)`

SetSourceState sets SourceState field to given value.

### HasSourceState

`func (o *HandlersDeploykitDefinitionStatus) HasSourceState() bool`

HasSourceState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


