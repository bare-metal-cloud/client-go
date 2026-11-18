# EngineportsVerifyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**LatencyMs** | Pointer to **int32** | LatencyMS is the round trip&#39;s milliseconds — the exec&#39;s dispatch, the query, and the stream&#39;s assembly (the wire form; the handler serializes this struct verbatim). | [optional] 
**Pod** | Pointer to **string** |  | [optional] 
**Selector** | Pointer to **string** |  | [optional] 
**Stderr** | Pointer to **string** |  | [optional] 
**Stdout** | Pointer to **string** |  | [optional] 

## Methods

### NewEngineportsVerifyResult

`func NewEngineportsVerifyResult() *EngineportsVerifyResult`

NewEngineportsVerifyResult instantiates a new EngineportsVerifyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsVerifyResultWithDefaults

`func NewEngineportsVerifyResultWithDefaults() *EngineportsVerifyResult`

NewEngineportsVerifyResultWithDefaults instantiates a new EngineportsVerifyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *EngineportsVerifyResult) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *EngineportsVerifyResult) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *EngineportsVerifyResult) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *EngineportsVerifyResult) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetExitCode

`func (o *EngineportsVerifyResult) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *EngineportsVerifyResult) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *EngineportsVerifyResult) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *EngineportsVerifyResult) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetLatencyMs

`func (o *EngineportsVerifyResult) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *EngineportsVerifyResult) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *EngineportsVerifyResult) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *EngineportsVerifyResult) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetPod

`func (o *EngineportsVerifyResult) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *EngineportsVerifyResult) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *EngineportsVerifyResult) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *EngineportsVerifyResult) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetSelector

`func (o *EngineportsVerifyResult) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *EngineportsVerifyResult) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *EngineportsVerifyResult) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *EngineportsVerifyResult) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStderr

`func (o *EngineportsVerifyResult) GetStderr() string`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *EngineportsVerifyResult) GetStderrOk() (*string, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *EngineportsVerifyResult) SetStderr(v string)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *EngineportsVerifyResult) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *EngineportsVerifyResult) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *EngineportsVerifyResult) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *EngineportsVerifyResult) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *EngineportsVerifyResult) HasStdout() bool`

HasStdout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


