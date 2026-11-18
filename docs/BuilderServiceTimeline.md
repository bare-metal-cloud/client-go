# BuilderServiceTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to [**BuilderAttemptSummary**](BuilderAttemptSummary.md) |  | [optional] 
**Build** | Pointer to [**BuilderStageState**](BuilderStageState.md) |  | [optional] 
**Health** | Pointer to [**BuilderStageState**](BuilderStageState.md) |  | [optional] 
**Publish** | Pointer to [**BuilderStageState**](BuilderStageState.md) |  | [optional] 
**Revision** | Pointer to [**BuilderRevisionSummary**](BuilderRevisionSummary.md) |  | [optional] 
**RunReference** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewBuilderServiceTimeline

`func NewBuilderServiceTimeline() *BuilderServiceTimeline`

NewBuilderServiceTimeline instantiates a new BuilderServiceTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderServiceTimelineWithDefaults

`func NewBuilderServiceTimelineWithDefaults() *BuilderServiceTimeline`

NewBuilderServiceTimelineWithDefaults instantiates a new BuilderServiceTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *BuilderServiceTimeline) GetAttempt() BuilderAttemptSummary`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *BuilderServiceTimeline) GetAttemptOk() (*BuilderAttemptSummary, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *BuilderServiceTimeline) SetAttempt(v BuilderAttemptSummary)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *BuilderServiceTimeline) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetBuild

`func (o *BuilderServiceTimeline) GetBuild() BuilderStageState`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *BuilderServiceTimeline) GetBuildOk() (*BuilderStageState, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *BuilderServiceTimeline) SetBuild(v BuilderStageState)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *BuilderServiceTimeline) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetHealth

`func (o *BuilderServiceTimeline) GetHealth() BuilderStageState`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *BuilderServiceTimeline) GetHealthOk() (*BuilderStageState, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *BuilderServiceTimeline) SetHealth(v BuilderStageState)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *BuilderServiceTimeline) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetPublish

`func (o *BuilderServiceTimeline) GetPublish() BuilderStageState`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *BuilderServiceTimeline) GetPublishOk() (*BuilderStageState, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *BuilderServiceTimeline) SetPublish(v BuilderStageState)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *BuilderServiceTimeline) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetRevision

`func (o *BuilderServiceTimeline) GetRevision() BuilderRevisionSummary`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BuilderServiceTimeline) GetRevisionOk() (*BuilderRevisionSummary, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BuilderServiceTimeline) SetRevision(v BuilderRevisionSummary)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BuilderServiceTimeline) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRunReference

`func (o *BuilderServiceTimeline) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *BuilderServiceTimeline) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *BuilderServiceTimeline) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *BuilderServiceTimeline) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.

### GetService

`func (o *BuilderServiceTimeline) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BuilderServiceTimeline) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BuilderServiceTimeline) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BuilderServiceTimeline) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


