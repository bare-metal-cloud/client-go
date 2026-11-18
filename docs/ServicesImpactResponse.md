# ServicesImpactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ServicesImpactAction**](ServicesImpactAction.md) |  | [optional] 
**Consequences** | Pointer to [**[]ServicesConsequence**](ServicesConsequence.md) |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to [**ServicesImpactSubject**](ServicesImpactSubject.md) |  | [optional] 

## Methods

### NewServicesImpactResponse

`func NewServicesImpactResponse() *ServicesImpactResponse`

NewServicesImpactResponse instantiates a new ServicesImpactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesImpactResponseWithDefaults

`func NewServicesImpactResponseWithDefaults() *ServicesImpactResponse`

NewServicesImpactResponseWithDefaults instantiates a new ServicesImpactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ServicesImpactResponse) GetAction() ServicesImpactAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServicesImpactResponse) GetActionOk() (*ServicesImpactAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServicesImpactResponse) SetAction(v ServicesImpactAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ServicesImpactResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConsequences

`func (o *ServicesImpactResponse) GetConsequences() []ServicesConsequence`

GetConsequences returns the Consequences field if non-nil, zero value otherwise.

### GetConsequencesOk

`func (o *ServicesImpactResponse) GetConsequencesOk() (*[]ServicesConsequence, bool)`

GetConsequencesOk returns a tuple with the Consequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequences

`func (o *ServicesImpactResponse) SetConsequences(v []ServicesConsequence)`

SetConsequences sets Consequences field to given value.

### HasConsequences

`func (o *ServicesImpactResponse) HasConsequences() bool`

HasConsequences returns a boolean if a field has been set.

### GetFingerprint

`func (o *ServicesImpactResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ServicesImpactResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ServicesImpactResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ServicesImpactResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetSteps

`func (o *ServicesImpactResponse) GetSteps() []string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ServicesImpactResponse) GetStepsOk() (*[]string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ServicesImpactResponse) SetSteps(v []string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ServicesImpactResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetSubject

`func (o *ServicesImpactResponse) GetSubject() ServicesImpactSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ServicesImpactResponse) GetSubjectOk() (*ServicesImpactSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ServicesImpactResponse) SetSubject(v ServicesImpactSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ServicesImpactResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


