# EngineportsInstallProgressEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** | Code and Detail are the typed refusal (issue #4493): read back from the stream&#39;s coded event. Empty on every other event. | [optional] 
**Detail** | Pointer to **[]int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OpId** | Pointer to **string** |  | [optional] 
**Plane** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEngineportsInstallProgressEvent

`func NewEngineportsInstallProgressEvent() *EngineportsInstallProgressEvent`

NewEngineportsInstallProgressEvent instantiates a new EngineportsInstallProgressEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsInstallProgressEventWithDefaults

`func NewEngineportsInstallProgressEventWithDefaults() *EngineportsInstallProgressEvent`

NewEngineportsInstallProgressEventWithDefaults instantiates a new EngineportsInstallProgressEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *EngineportsInstallProgressEvent) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *EngineportsInstallProgressEvent) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *EngineportsInstallProgressEvent) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *EngineportsInstallProgressEvent) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCode

`func (o *EngineportsInstallProgressEvent) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EngineportsInstallProgressEvent) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EngineportsInstallProgressEvent) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EngineportsInstallProgressEvent) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *EngineportsInstallProgressEvent) GetDetail() []int32`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EngineportsInstallProgressEvent) GetDetailOk() (*[]int32, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EngineportsInstallProgressEvent) SetDetail(v []int32)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *EngineportsInstallProgressEvent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetId

`func (o *EngineportsInstallProgressEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineportsInstallProgressEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineportsInstallProgressEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EngineportsInstallProgressEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *EngineportsInstallProgressEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EngineportsInstallProgressEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EngineportsInstallProgressEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EngineportsInstallProgressEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOpId

`func (o *EngineportsInstallProgressEvent) GetOpId() string`

GetOpId returns the OpId field if non-nil, zero value otherwise.

### GetOpIdOk

`func (o *EngineportsInstallProgressEvent) GetOpIdOk() (*string, bool)`

GetOpIdOk returns a tuple with the OpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpId

`func (o *EngineportsInstallProgressEvent) SetOpId(v string)`

SetOpId sets OpId field to given value.

### HasOpId

`func (o *EngineportsInstallProgressEvent) HasOpId() bool`

HasOpId returns a boolean if a field has been set.

### GetPlane

`func (o *EngineportsInstallProgressEvent) GetPlane() string`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *EngineportsInstallProgressEvent) GetPlaneOk() (*string, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *EngineportsInstallProgressEvent) SetPlane(v string)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *EngineportsInstallProgressEvent) HasPlane() bool`

HasPlane returns a boolean if a field has been set.

### GetSeq

`func (o *EngineportsInstallProgressEvent) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *EngineportsInstallProgressEvent) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *EngineportsInstallProgressEvent) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *EngineportsInstallProgressEvent) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetState

`func (o *EngineportsInstallProgressEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EngineportsInstallProgressEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EngineportsInstallProgressEvent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EngineportsInstallProgressEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStep

`func (o *EngineportsInstallProgressEvent) GetStep() string`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *EngineportsInstallProgressEvent) GetStepOk() (*string, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *EngineportsInstallProgressEvent) SetStep(v string)`

SetStep sets Step field to given value.

### HasStep

`func (o *EngineportsInstallProgressEvent) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetTs

`func (o *EngineportsInstallProgressEvent) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *EngineportsInstallProgressEvent) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *EngineportsInstallProgressEvent) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *EngineportsInstallProgressEvent) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *EngineportsInstallProgressEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EngineportsInstallProgressEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EngineportsInstallProgressEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EngineportsInstallProgressEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


