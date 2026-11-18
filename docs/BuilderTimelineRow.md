# BuilderTimelineRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RunReference** | Pointer to **string** | RunReference is the service&#39;s build run (the attempt ledger&#39;s stable run identity). Empty is legal — a service whose build has not been dispatched yet renders every stage pending. | [optional] 

## Methods

### NewBuilderTimelineRow

`func NewBuilderTimelineRow() *BuilderTimelineRow`

NewBuilderTimelineRow instantiates a new BuilderTimelineRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderTimelineRowWithDefaults

`func NewBuilderTimelineRowWithDefaults() *BuilderTimelineRow`

NewBuilderTimelineRowWithDefaults instantiates a new BuilderTimelineRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BuilderTimelineRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuilderTimelineRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuilderTimelineRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BuilderTimelineRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunReference

`func (o *BuilderTimelineRow) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *BuilderTimelineRow) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *BuilderTimelineRow) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *BuilderTimelineRow) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


