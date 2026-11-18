# BuilderStageState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evidence** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewBuilderStageState

`func NewBuilderStageState() *BuilderStageState`

NewBuilderStageState instantiates a new BuilderStageState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderStageStateWithDefaults

`func NewBuilderStageStateWithDefaults() *BuilderStageState`

NewBuilderStageStateWithDefaults instantiates a new BuilderStageState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidence

`func (o *BuilderStageState) GetEvidence() []string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *BuilderStageState) GetEvidenceOk() (*[]string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *BuilderStageState) SetEvidence(v []string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *BuilderStageState) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetState

`func (o *BuilderStageState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BuilderStageState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BuilderStageState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BuilderStageState) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


