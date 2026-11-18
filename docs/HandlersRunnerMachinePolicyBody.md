# HandlersRunnerMachinePolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoexistencePresets** | Pointer to [**[]ModelsMachineCoexistencePreset**](ModelsMachineCoexistencePreset.md) | CoexistencePresets is control 2. | [optional] 
**ConfirmKindRemoval** | Pointer to **bool** | ConfirmKindRemoval is the coexistence consent law&#39;s explicit confirmation — required when the PUT removes a kind the org&#39;s live machines still serve. | [optional] 
**ConsentRecord** | Pointer to [**map[string]ModelsMachineConsentDecision**](ModelsMachineConsentDecision.md) | ConsentRecord is control 5 (consent point → decision); omitted points keep their defaults. | [optional] 
**KindsAllowed** | Pointer to [**[]ModelsRunnerMachineKind**](ModelsRunnerMachineKind.md) | KindsAllowed is control 1. | [optional] 
**OnUnreliability** | Pointer to [**ModelsMachineOnUnreliability**](ModelsMachineOnUnreliability.md) | OnUnreliability is control 4 (quarantine_and_heal | destroy_and_replace). | [optional] 
**OrgDefaultRunnerVersion** | Pointer to **string** | OrgDefaultRunnerVersion is the org-default pin the pinned_org_default posture requires. | [optional] 
**VersionPinPosture** | Pointer to [**ModelsMachineVersionPinPosture**](ModelsMachineVersionPinPosture.md) | VersionPinPosture is control 3 (pinned | pinned_org_default). | [optional] 

## Methods

### NewHandlersRunnerMachinePolicyBody

`func NewHandlersRunnerMachinePolicyBody() *HandlersRunnerMachinePolicyBody`

NewHandlersRunnerMachinePolicyBody instantiates a new HandlersRunnerMachinePolicyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMachinePolicyBodyWithDefaults

`func NewHandlersRunnerMachinePolicyBodyWithDefaults() *HandlersRunnerMachinePolicyBody`

NewHandlersRunnerMachinePolicyBodyWithDefaults instantiates a new HandlersRunnerMachinePolicyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoexistencePresets

`func (o *HandlersRunnerMachinePolicyBody) GetCoexistencePresets() []ModelsMachineCoexistencePreset`

GetCoexistencePresets returns the CoexistencePresets field if non-nil, zero value otherwise.

### GetCoexistencePresetsOk

`func (o *HandlersRunnerMachinePolicyBody) GetCoexistencePresetsOk() (*[]ModelsMachineCoexistencePreset, bool)`

GetCoexistencePresetsOk returns a tuple with the CoexistencePresets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoexistencePresets

`func (o *HandlersRunnerMachinePolicyBody) SetCoexistencePresets(v []ModelsMachineCoexistencePreset)`

SetCoexistencePresets sets CoexistencePresets field to given value.

### HasCoexistencePresets

`func (o *HandlersRunnerMachinePolicyBody) HasCoexistencePresets() bool`

HasCoexistencePresets returns a boolean if a field has been set.

### GetConfirmKindRemoval

`func (o *HandlersRunnerMachinePolicyBody) GetConfirmKindRemoval() bool`

GetConfirmKindRemoval returns the ConfirmKindRemoval field if non-nil, zero value otherwise.

### GetConfirmKindRemovalOk

`func (o *HandlersRunnerMachinePolicyBody) GetConfirmKindRemovalOk() (*bool, bool)`

GetConfirmKindRemovalOk returns a tuple with the ConfirmKindRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmKindRemoval

`func (o *HandlersRunnerMachinePolicyBody) SetConfirmKindRemoval(v bool)`

SetConfirmKindRemoval sets ConfirmKindRemoval field to given value.

### HasConfirmKindRemoval

`func (o *HandlersRunnerMachinePolicyBody) HasConfirmKindRemoval() bool`

HasConfirmKindRemoval returns a boolean if a field has been set.

### GetConsentRecord

`func (o *HandlersRunnerMachinePolicyBody) GetConsentRecord() map[string]ModelsMachineConsentDecision`

GetConsentRecord returns the ConsentRecord field if non-nil, zero value otherwise.

### GetConsentRecordOk

`func (o *HandlersRunnerMachinePolicyBody) GetConsentRecordOk() (*map[string]ModelsMachineConsentDecision, bool)`

GetConsentRecordOk returns a tuple with the ConsentRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRecord

`func (o *HandlersRunnerMachinePolicyBody) SetConsentRecord(v map[string]ModelsMachineConsentDecision)`

SetConsentRecord sets ConsentRecord field to given value.

### HasConsentRecord

`func (o *HandlersRunnerMachinePolicyBody) HasConsentRecord() bool`

HasConsentRecord returns a boolean if a field has been set.

### GetKindsAllowed

`func (o *HandlersRunnerMachinePolicyBody) GetKindsAllowed() []ModelsRunnerMachineKind`

GetKindsAllowed returns the KindsAllowed field if non-nil, zero value otherwise.

### GetKindsAllowedOk

`func (o *HandlersRunnerMachinePolicyBody) GetKindsAllowedOk() (*[]ModelsRunnerMachineKind, bool)`

GetKindsAllowedOk returns a tuple with the KindsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindsAllowed

`func (o *HandlersRunnerMachinePolicyBody) SetKindsAllowed(v []ModelsRunnerMachineKind)`

SetKindsAllowed sets KindsAllowed field to given value.

### HasKindsAllowed

`func (o *HandlersRunnerMachinePolicyBody) HasKindsAllowed() bool`

HasKindsAllowed returns a boolean if a field has been set.

### GetOnUnreliability

`func (o *HandlersRunnerMachinePolicyBody) GetOnUnreliability() ModelsMachineOnUnreliability`

GetOnUnreliability returns the OnUnreliability field if non-nil, zero value otherwise.

### GetOnUnreliabilityOk

`func (o *HandlersRunnerMachinePolicyBody) GetOnUnreliabilityOk() (*ModelsMachineOnUnreliability, bool)`

GetOnUnreliabilityOk returns a tuple with the OnUnreliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnUnreliability

`func (o *HandlersRunnerMachinePolicyBody) SetOnUnreliability(v ModelsMachineOnUnreliability)`

SetOnUnreliability sets OnUnreliability field to given value.

### HasOnUnreliability

`func (o *HandlersRunnerMachinePolicyBody) HasOnUnreliability() bool`

HasOnUnreliability returns a boolean if a field has been set.

### GetOrgDefaultRunnerVersion

`func (o *HandlersRunnerMachinePolicyBody) GetOrgDefaultRunnerVersion() string`

GetOrgDefaultRunnerVersion returns the OrgDefaultRunnerVersion field if non-nil, zero value otherwise.

### GetOrgDefaultRunnerVersionOk

`func (o *HandlersRunnerMachinePolicyBody) GetOrgDefaultRunnerVersionOk() (*string, bool)`

GetOrgDefaultRunnerVersionOk returns a tuple with the OrgDefaultRunnerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefaultRunnerVersion

`func (o *HandlersRunnerMachinePolicyBody) SetOrgDefaultRunnerVersion(v string)`

SetOrgDefaultRunnerVersion sets OrgDefaultRunnerVersion field to given value.

### HasOrgDefaultRunnerVersion

`func (o *HandlersRunnerMachinePolicyBody) HasOrgDefaultRunnerVersion() bool`

HasOrgDefaultRunnerVersion returns a boolean if a field has been set.

### GetVersionPinPosture

`func (o *HandlersRunnerMachinePolicyBody) GetVersionPinPosture() ModelsMachineVersionPinPosture`

GetVersionPinPosture returns the VersionPinPosture field if non-nil, zero value otherwise.

### GetVersionPinPostureOk

`func (o *HandlersRunnerMachinePolicyBody) GetVersionPinPostureOk() (*ModelsMachineVersionPinPosture, bool)`

GetVersionPinPostureOk returns a tuple with the VersionPinPosture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPinPosture

`func (o *HandlersRunnerMachinePolicyBody) SetVersionPinPosture(v ModelsMachineVersionPinPosture)`

SetVersionPinPosture sets VersionPinPosture field to given value.

### HasVersionPinPosture

`func (o *HandlersRunnerMachinePolicyBody) HasVersionPinPosture() bool`

HasVersionPinPosture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


