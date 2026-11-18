# RunnerpoolMachineDispatchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doc** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to **[]int32** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** | OrgID is the dispatch envelope&#39;s org (AAR.3 item 8): the CP-asserted tenant identity the machine-side refusal gate keys its different-org receipt check on. The pool set it from the admitted job — never from the customer&#39;s metering tags. | [optional] 
**ProviderVmId** | Pointer to **string** | ProviderVMID is the machine&#39;s own dispatch address — an echo of the credential-derived identity, never a caller-chosen value. | [optional] 
**RunHandleId** | Pointer to **string** |  | [optional] 
**Sig** | Pointer to **string** |  | [optional] 

## Methods

### NewRunnerpoolMachineDispatchItem

`func NewRunnerpoolMachineDispatchItem() *RunnerpoolMachineDispatchItem`

NewRunnerpoolMachineDispatchItem instantiates a new RunnerpoolMachineDispatchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerpoolMachineDispatchItemWithDefaults

`func NewRunnerpoolMachineDispatchItemWithDefaults() *RunnerpoolMachineDispatchItem`

NewRunnerpoolMachineDispatchItemWithDefaults instantiates a new RunnerpoolMachineDispatchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoc

`func (o *RunnerpoolMachineDispatchItem) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *RunnerpoolMachineDispatchItem) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *RunnerpoolMachineDispatchItem) SetDoc(v string)`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *RunnerpoolMachineDispatchItem) HasDoc() bool`

HasDoc returns a boolean if a field has been set.

### GetInputs

`func (o *RunnerpoolMachineDispatchItem) GetInputs() []int32`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *RunnerpoolMachineDispatchItem) GetInputsOk() (*[]int32, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *RunnerpoolMachineDispatchItem) SetInputs(v []int32)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *RunnerpoolMachineDispatchItem) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetJobId

`func (o *RunnerpoolMachineDispatchItem) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RunnerpoolMachineDispatchItem) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RunnerpoolMachineDispatchItem) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RunnerpoolMachineDispatchItem) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOrgId

`func (o *RunnerpoolMachineDispatchItem) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RunnerpoolMachineDispatchItem) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RunnerpoolMachineDispatchItem) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RunnerpoolMachineDispatchItem) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProviderVmId

`func (o *RunnerpoolMachineDispatchItem) GetProviderVmId() string`

GetProviderVmId returns the ProviderVmId field if non-nil, zero value otherwise.

### GetProviderVmIdOk

`func (o *RunnerpoolMachineDispatchItem) GetProviderVmIdOk() (*string, bool)`

GetProviderVmIdOk returns a tuple with the ProviderVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVmId

`func (o *RunnerpoolMachineDispatchItem) SetProviderVmId(v string)`

SetProviderVmId sets ProviderVmId field to given value.

### HasProviderVmId

`func (o *RunnerpoolMachineDispatchItem) HasProviderVmId() bool`

HasProviderVmId returns a boolean if a field has been set.

### GetRunHandleId

`func (o *RunnerpoolMachineDispatchItem) GetRunHandleId() string`

GetRunHandleId returns the RunHandleId field if non-nil, zero value otherwise.

### GetRunHandleIdOk

`func (o *RunnerpoolMachineDispatchItem) GetRunHandleIdOk() (*string, bool)`

GetRunHandleIdOk returns a tuple with the RunHandleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunHandleId

`func (o *RunnerpoolMachineDispatchItem) SetRunHandleId(v string)`

SetRunHandleId sets RunHandleId field to given value.

### HasRunHandleId

`func (o *RunnerpoolMachineDispatchItem) HasRunHandleId() bool`

HasRunHandleId returns a boolean if a field has been set.

### GetSig

`func (o *RunnerpoolMachineDispatchItem) GetSig() string`

GetSig returns the Sig field if non-nil, zero value otherwise.

### GetSigOk

`func (o *RunnerpoolMachineDispatchItem) GetSigOk() (*string, bool)`

GetSigOk returns a tuple with the Sig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSig

`func (o *RunnerpoolMachineDispatchItem) SetSig(v string)`

SetSig sets Sig field to given value.

### HasSig

`func (o *RunnerpoolMachineDispatchItem) HasSig() bool`

HasSig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


