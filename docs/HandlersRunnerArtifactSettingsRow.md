# HandlersRunnerArtifactSettingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int32** | Days is the days-mode lifetime in days; null under delete-when-run-ends. | [optional] 
**DeleteWhenRunEnds** | Pointer to **bool** | DeleteWhenRunEnds is the delete-when-run-ends mode&#39;s flag. | [optional] 

## Methods

### NewHandlersRunnerArtifactSettingsRow

`func NewHandlersRunnerArtifactSettingsRow() *HandlersRunnerArtifactSettingsRow`

NewHandlersRunnerArtifactSettingsRow instantiates a new HandlersRunnerArtifactSettingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerArtifactSettingsRowWithDefaults

`func NewHandlersRunnerArtifactSettingsRowWithDefaults() *HandlersRunnerArtifactSettingsRow`

NewHandlersRunnerArtifactSettingsRowWithDefaults instantiates a new HandlersRunnerArtifactSettingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *HandlersRunnerArtifactSettingsRow) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *HandlersRunnerArtifactSettingsRow) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *HandlersRunnerArtifactSettingsRow) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *HandlersRunnerArtifactSettingsRow) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDeleteWhenRunEnds

`func (o *HandlersRunnerArtifactSettingsRow) GetDeleteWhenRunEnds() bool`

GetDeleteWhenRunEnds returns the DeleteWhenRunEnds field if non-nil, zero value otherwise.

### GetDeleteWhenRunEndsOk

`func (o *HandlersRunnerArtifactSettingsRow) GetDeleteWhenRunEndsOk() (*bool, bool)`

GetDeleteWhenRunEndsOk returns a tuple with the DeleteWhenRunEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteWhenRunEnds

`func (o *HandlersRunnerArtifactSettingsRow) SetDeleteWhenRunEnds(v bool)`

SetDeleteWhenRunEnds sets DeleteWhenRunEnds field to given value.

### HasDeleteWhenRunEnds

`func (o *HandlersRunnerArtifactSettingsRow) HasDeleteWhenRunEnds() bool`

HasDeleteWhenRunEnds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


