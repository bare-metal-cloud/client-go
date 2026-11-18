# BuilderBuildFailureView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitCode** | Pointer to **int32** | ExitCode is the ledger&#39;s exit fact when the output ref carries one (\&quot;exit&#x3D;N\&quot;); nil when the failure landed without one. | [optional] 
**Likely** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to **[]string** | Lines are the real failing lines — the attempt&#39;s last lines, verbatim (the log store&#39;s order). | [optional] 
**OutputRef** | Pointer to **string** | OutputRef is the ledger&#39;s stored output fact, verbatim. | [optional] 
**Saw** | Pointer to **string** | Saw / Likely / ToTry are the read-only diagnosis rows. | [optional] 
**Step** | Pointer to **string** | Step is the failing plan step (the last step the builder&#39;s plan-echo attributed a line to); \&quot;\&quot; when no line carries an attribution. | [optional] 
**ToTry** | Pointer to **string** |  | [optional] 

## Methods

### NewBuilderBuildFailureView

`func NewBuilderBuildFailureView() *BuilderBuildFailureView`

NewBuilderBuildFailureView instantiates a new BuilderBuildFailureView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderBuildFailureViewWithDefaults

`func NewBuilderBuildFailureViewWithDefaults() *BuilderBuildFailureView`

NewBuilderBuildFailureViewWithDefaults instantiates a new BuilderBuildFailureView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitCode

`func (o *BuilderBuildFailureView) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *BuilderBuildFailureView) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *BuilderBuildFailureView) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *BuilderBuildFailureView) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetLikely

`func (o *BuilderBuildFailureView) GetLikely() string`

GetLikely returns the Likely field if non-nil, zero value otherwise.

### GetLikelyOk

`func (o *BuilderBuildFailureView) GetLikelyOk() (*string, bool)`

GetLikelyOk returns a tuple with the Likely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikely

`func (o *BuilderBuildFailureView) SetLikely(v string)`

SetLikely sets Likely field to given value.

### HasLikely

`func (o *BuilderBuildFailureView) HasLikely() bool`

HasLikely returns a boolean if a field has been set.

### GetLines

`func (o *BuilderBuildFailureView) GetLines() []string`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *BuilderBuildFailureView) GetLinesOk() (*[]string, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *BuilderBuildFailureView) SetLines(v []string)`

SetLines sets Lines field to given value.

### HasLines

`func (o *BuilderBuildFailureView) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetOutputRef

`func (o *BuilderBuildFailureView) GetOutputRef() string`

GetOutputRef returns the OutputRef field if non-nil, zero value otherwise.

### GetOutputRefOk

`func (o *BuilderBuildFailureView) GetOutputRefOk() (*string, bool)`

GetOutputRefOk returns a tuple with the OutputRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRef

`func (o *BuilderBuildFailureView) SetOutputRef(v string)`

SetOutputRef sets OutputRef field to given value.

### HasOutputRef

`func (o *BuilderBuildFailureView) HasOutputRef() bool`

HasOutputRef returns a boolean if a field has been set.

### GetSaw

`func (o *BuilderBuildFailureView) GetSaw() string`

GetSaw returns the Saw field if non-nil, zero value otherwise.

### GetSawOk

`func (o *BuilderBuildFailureView) GetSawOk() (*string, bool)`

GetSawOk returns a tuple with the Saw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaw

`func (o *BuilderBuildFailureView) SetSaw(v string)`

SetSaw sets Saw field to given value.

### HasSaw

`func (o *BuilderBuildFailureView) HasSaw() bool`

HasSaw returns a boolean if a field has been set.

### GetStep

`func (o *BuilderBuildFailureView) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *BuilderBuildFailureView) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *BuilderBuildFailureView) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *BuilderBuildFailureView) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetToTry

`func (o *BuilderBuildFailureView) GetToTry() string`

GetToTry returns the ToTry field if non-nil, zero value otherwise.

### GetToTryOk

`func (o *BuilderBuildFailureView) GetToTryOk() (*string, bool)`

GetToTryOk returns a tuple with the ToTry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTry

`func (o *BuilderBuildFailureView) SetToTry(v string)`

SetToTry sets ToTry field to given value.

### HasToTry

`func (o *BuilderBuildFailureView) HasToTry() bool`

HasToTry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


