# HandlersRunnerArtifactSettingsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int32** | Days is the days-mode lifetime; ignored while DeleteWhenRunEnds stands (SetRetention&#39;s law — the caller may send either shape). | [optional] 
**DeleteWhenRunEnds** | Pointer to **bool** | DeleteWhenRunEnds selects the delete-when-run-ends mode. | [optional] 

## Methods

### NewHandlersRunnerArtifactSettingsBody

`func NewHandlersRunnerArtifactSettingsBody() *HandlersRunnerArtifactSettingsBody`

NewHandlersRunnerArtifactSettingsBody instantiates a new HandlersRunnerArtifactSettingsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerArtifactSettingsBodyWithDefaults

`func NewHandlersRunnerArtifactSettingsBodyWithDefaults() *HandlersRunnerArtifactSettingsBody`

NewHandlersRunnerArtifactSettingsBodyWithDefaults instantiates a new HandlersRunnerArtifactSettingsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *HandlersRunnerArtifactSettingsBody) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *HandlersRunnerArtifactSettingsBody) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *HandlersRunnerArtifactSettingsBody) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *HandlersRunnerArtifactSettingsBody) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDeleteWhenRunEnds

`func (o *HandlersRunnerArtifactSettingsBody) GetDeleteWhenRunEnds() bool`

GetDeleteWhenRunEnds returns the DeleteWhenRunEnds field if non-nil, zero value otherwise.

### GetDeleteWhenRunEndsOk

`func (o *HandlersRunnerArtifactSettingsBody) GetDeleteWhenRunEndsOk() (*bool, bool)`

GetDeleteWhenRunEndsOk returns a tuple with the DeleteWhenRunEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteWhenRunEnds

`func (o *HandlersRunnerArtifactSettingsBody) SetDeleteWhenRunEnds(v bool)`

SetDeleteWhenRunEnds sets DeleteWhenRunEnds field to given value.

### HasDeleteWhenRunEnds

`func (o *HandlersRunnerArtifactSettingsBody) HasDeleteWhenRunEnds() bool`

HasDeleteWhenRunEnds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


