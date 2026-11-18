# ServicesPeerProbeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **map[string]string** |  | [optional] 
**Cached** | Pointer to **bool** |  | [optional] 
**Components** | Pointer to **[][]int32** | Components groups mutually-reachable VM IDs so the UI can say \&quot;A, B, C reach each other; D is isolated\&quot;. | [optional] 
**CriticalVms** | Pointer to **[]int32** |  | [optional] 
**Disclosure** | Pointer to **string** |  | [optional] 
**FailingPairs** | Pointer to [**[]ServicesPeerProbePairFailure**](ServicesPeerProbePairFailure.md) |  | [optional] 
**MeasuredAt** | Pointer to **string** |  | [optional] 
**MtuWarnings** | Pointer to [**[]ServicesPeerProbeMtuWarning**](ServicesPeerProbeMtuWarning.md) | MtuWarnings carries the mesh plane&#39;s pair-level large-packet failures (issue #3378): WARNING entries, never verdict-changing. A direct-plane response never carries them (the DF component is mesh-plane only). | [optional] 
**Plane** | Pointer to **string** |  | [optional] 
**Reachable** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Residuals** | Pointer to **[]string** |  | [optional] 
**Verdict** | Pointer to **string** |  | [optional] 
**Vms** | Pointer to [**[]ServicesPeerProbeVMResult**](ServicesPeerProbeVMResult.md) |  | [optional] 

## Methods

### NewServicesPeerProbeResponse

`func NewServicesPeerProbeResponse() *ServicesPeerProbeResponse`

NewServicesPeerProbeResponse instantiates a new ServicesPeerProbeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesPeerProbeResponseWithDefaults

`func NewServicesPeerProbeResponseWithDefaults() *ServicesPeerProbeResponse`

NewServicesPeerProbeResponseWithDefaults instantiates a new ServicesPeerProbeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ServicesPeerProbeResponse) GetAddresses() map[string]string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ServicesPeerProbeResponse) GetAddressesOk() (*map[string]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ServicesPeerProbeResponse) SetAddresses(v map[string]string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ServicesPeerProbeResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCached

`func (o *ServicesPeerProbeResponse) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *ServicesPeerProbeResponse) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *ServicesPeerProbeResponse) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *ServicesPeerProbeResponse) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetComponents

`func (o *ServicesPeerProbeResponse) GetComponents() [][]int32`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ServicesPeerProbeResponse) GetComponentsOk() (*[][]int32, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ServicesPeerProbeResponse) SetComponents(v [][]int32)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ServicesPeerProbeResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCriticalVms

`func (o *ServicesPeerProbeResponse) GetCriticalVms() []int32`

GetCriticalVms returns the CriticalVms field if non-nil, zero value otherwise.

### GetCriticalVmsOk

`func (o *ServicesPeerProbeResponse) GetCriticalVmsOk() (*[]int32, bool)`

GetCriticalVmsOk returns a tuple with the CriticalVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalVms

`func (o *ServicesPeerProbeResponse) SetCriticalVms(v []int32)`

SetCriticalVms sets CriticalVms field to given value.

### HasCriticalVms

`func (o *ServicesPeerProbeResponse) HasCriticalVms() bool`

HasCriticalVms returns a boolean if a field has been set.

### GetDisclosure

`func (o *ServicesPeerProbeResponse) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *ServicesPeerProbeResponse) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *ServicesPeerProbeResponse) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *ServicesPeerProbeResponse) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetFailingPairs

`func (o *ServicesPeerProbeResponse) GetFailingPairs() []ServicesPeerProbePairFailure`

GetFailingPairs returns the FailingPairs field if non-nil, zero value otherwise.

### GetFailingPairsOk

`func (o *ServicesPeerProbeResponse) GetFailingPairsOk() (*[]ServicesPeerProbePairFailure, bool)`

GetFailingPairsOk returns a tuple with the FailingPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingPairs

`func (o *ServicesPeerProbeResponse) SetFailingPairs(v []ServicesPeerProbePairFailure)`

SetFailingPairs sets FailingPairs field to given value.

### HasFailingPairs

`func (o *ServicesPeerProbeResponse) HasFailingPairs() bool`

HasFailingPairs returns a boolean if a field has been set.

### GetMeasuredAt

`func (o *ServicesPeerProbeResponse) GetMeasuredAt() string`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *ServicesPeerProbeResponse) GetMeasuredAtOk() (*string, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *ServicesPeerProbeResponse) SetMeasuredAt(v string)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *ServicesPeerProbeResponse) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetMtuWarnings

`func (o *ServicesPeerProbeResponse) GetMtuWarnings() []ServicesPeerProbeMtuWarning`

GetMtuWarnings returns the MtuWarnings field if non-nil, zero value otherwise.

### GetMtuWarningsOk

`func (o *ServicesPeerProbeResponse) GetMtuWarningsOk() (*[]ServicesPeerProbeMtuWarning, bool)`

GetMtuWarningsOk returns a tuple with the MtuWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtuWarnings

`func (o *ServicesPeerProbeResponse) SetMtuWarnings(v []ServicesPeerProbeMtuWarning)`

SetMtuWarnings sets MtuWarnings field to given value.

### HasMtuWarnings

`func (o *ServicesPeerProbeResponse) HasMtuWarnings() bool`

HasMtuWarnings returns a boolean if a field has been set.

### GetPlane

`func (o *ServicesPeerProbeResponse) GetPlane() string`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *ServicesPeerProbeResponse) GetPlaneOk() (*string, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *ServicesPeerProbeResponse) SetPlane(v string)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *ServicesPeerProbeResponse) HasPlane() bool`

HasPlane returns a boolean if a field has been set.

### GetReachable

`func (o *ServicesPeerProbeResponse) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *ServicesPeerProbeResponse) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *ServicesPeerProbeResponse) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *ServicesPeerProbeResponse) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetReason

`func (o *ServicesPeerProbeResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ServicesPeerProbeResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ServicesPeerProbeResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ServicesPeerProbeResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResiduals

`func (o *ServicesPeerProbeResponse) GetResiduals() []string`

GetResiduals returns the Residuals field if non-nil, zero value otherwise.

### GetResidualsOk

`func (o *ServicesPeerProbeResponse) GetResidualsOk() (*[]string, bool)`

GetResidualsOk returns a tuple with the Residuals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResiduals

`func (o *ServicesPeerProbeResponse) SetResiduals(v []string)`

SetResiduals sets Residuals field to given value.

### HasResiduals

`func (o *ServicesPeerProbeResponse) HasResiduals() bool`

HasResiduals returns a boolean if a field has been set.

### GetVerdict

`func (o *ServicesPeerProbeResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ServicesPeerProbeResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ServicesPeerProbeResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ServicesPeerProbeResponse) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetVms

`func (o *ServicesPeerProbeResponse) GetVms() []ServicesPeerProbeVMResult`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *ServicesPeerProbeResponse) GetVmsOk() (*[]ServicesPeerProbeVMResult, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *ServicesPeerProbeResponse) SetVms(v []ServicesPeerProbeVMResult)`

SetVms sets Vms field to given value.

### HasVms

`func (o *ServicesPeerProbeResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


