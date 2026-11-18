# HandlersDeploykitInstallStatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to **string** | Cause is the named reason verbatim — the workflow note or the terminal event&#39;s message; for a terminal failure with no event surviving, the record state&#39;s cause class. | [optional] 
**Code** | Pointer to **string** | Code is the typed refusal code (issue #4493, additive and omitempty): the board decides from this field, never from Cause&#39;s prose. Set only on a terminal failure whose newest coded event survives in the window. | [optional] 
**Kind** | Pointer to **string** | Kind is the detail class: \&quot;failed\&quot; for a terminal failure, \&quot;working\&quot; for a live named note (a backoff among them). | [optional] 
**Refusal** | Pointer to [**EngineportsArrivalStrandRefusalDetail**](EngineportsArrivalStrandRefusalDetail.md) | Refusal is the typed deep-link object the Code names (today: the stranded-records refusal). Absent unless Code names it and its detail decodes. | [optional] 

## Methods

### NewHandlersDeploykitInstallStatusDetail

`func NewHandlersDeploykitInstallStatusDetail() *HandlersDeploykitInstallStatusDetail`

NewHandlersDeploykitInstallStatusDetail instantiates a new HandlersDeploykitInstallStatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDeploykitInstallStatusDetailWithDefaults

`func NewHandlersDeploykitInstallStatusDetailWithDefaults() *HandlersDeploykitInstallStatusDetail`

NewHandlersDeploykitInstallStatusDetailWithDefaults instantiates a new HandlersDeploykitInstallStatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *HandlersDeploykitInstallStatusDetail) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HandlersDeploykitInstallStatusDetail) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HandlersDeploykitInstallStatusDetail) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HandlersDeploykitInstallStatusDetail) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetCode

`func (o *HandlersDeploykitInstallStatusDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HandlersDeploykitInstallStatusDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HandlersDeploykitInstallStatusDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *HandlersDeploykitInstallStatusDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetKind

`func (o *HandlersDeploykitInstallStatusDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HandlersDeploykitInstallStatusDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HandlersDeploykitInstallStatusDetail) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HandlersDeploykitInstallStatusDetail) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRefusal

`func (o *HandlersDeploykitInstallStatusDetail) GetRefusal() EngineportsArrivalStrandRefusalDetail`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *HandlersDeploykitInstallStatusDetail) GetRefusalOk() (*EngineportsArrivalStrandRefusalDetail, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *HandlersDeploykitInstallStatusDetail) SetRefusal(v EngineportsArrivalStrandRefusalDetail)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *HandlersDeploykitInstallStatusDetail) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


