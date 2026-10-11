# NetworkbackendCandidatesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | Pointer to [**[]NetworkbackendNetworkCandidate**](NetworkbackendNetworkCandidate.md) |  | [optional] 
**Disclosure** | Pointer to [**NetworkbackendNetworkDisclosure**](NetworkbackendNetworkDisclosure.md) |  | [optional] 

## Methods

### NewNetworkbackendCandidatesResult

`func NewNetworkbackendCandidatesResult() *NetworkbackendCandidatesResult`

NewNetworkbackendCandidatesResult instantiates a new NetworkbackendCandidatesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkbackendCandidatesResultWithDefaults

`func NewNetworkbackendCandidatesResultWithDefaults() *NetworkbackendCandidatesResult`

NewNetworkbackendCandidatesResultWithDefaults instantiates a new NetworkbackendCandidatesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *NetworkbackendCandidatesResult) GetCandidates() []NetworkbackendNetworkCandidate`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *NetworkbackendCandidatesResult) GetCandidatesOk() (*[]NetworkbackendNetworkCandidate, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *NetworkbackendCandidatesResult) SetCandidates(v []NetworkbackendNetworkCandidate)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *NetworkbackendCandidatesResult) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetDisclosure

`func (o *NetworkbackendCandidatesResult) GetDisclosure() NetworkbackendNetworkDisclosure`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *NetworkbackendCandidatesResult) GetDisclosureOk() (*NetworkbackendNetworkDisclosure, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *NetworkbackendCandidatesResult) SetDisclosure(v NetworkbackendNetworkDisclosure)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *NetworkbackendCandidatesResult) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


