# HandlersRunnerInstallAckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** |  | [optional] 
**MachineFingerprint** | Pointer to **string** |  | [optional] 
**Preflight** | Pointer to [**RunnerinstallsPreflightReport**](RunnerinstallsPreflightReport.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerInstallAckRequest

`func NewHandlersRunnerInstallAckRequest() *HandlersRunnerInstallAckRequest`

NewHandlersRunnerInstallAckRequest instantiates a new HandlersRunnerInstallAckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerInstallAckRequestWithDefaults

`func NewHandlersRunnerInstallAckRequestWithDefaults() *HandlersRunnerInstallAckRequest`

NewHandlersRunnerInstallAckRequestWithDefaults instantiates a new HandlersRunnerInstallAckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *HandlersRunnerInstallAckRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HandlersRunnerInstallAckRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HandlersRunnerInstallAckRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HandlersRunnerInstallAckRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMachineFingerprint

`func (o *HandlersRunnerInstallAckRequest) GetMachineFingerprint() string`

GetMachineFingerprint returns the MachineFingerprint field if non-nil, zero value otherwise.

### GetMachineFingerprintOk

`func (o *HandlersRunnerInstallAckRequest) GetMachineFingerprintOk() (*string, bool)`

GetMachineFingerprintOk returns a tuple with the MachineFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineFingerprint

`func (o *HandlersRunnerInstallAckRequest) SetMachineFingerprint(v string)`

SetMachineFingerprint sets MachineFingerprint field to given value.

### HasMachineFingerprint

`func (o *HandlersRunnerInstallAckRequest) HasMachineFingerprint() bool`

HasMachineFingerprint returns a boolean if a field has been set.

### GetPreflight

`func (o *HandlersRunnerInstallAckRequest) GetPreflight() RunnerinstallsPreflightReport`

GetPreflight returns the Preflight field if non-nil, zero value otherwise.

### GetPreflightOk

`func (o *HandlersRunnerInstallAckRequest) GetPreflightOk() (*RunnerinstallsPreflightReport, bool)`

GetPreflightOk returns a tuple with the Preflight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreflight

`func (o *HandlersRunnerInstallAckRequest) SetPreflight(v RunnerinstallsPreflightReport)`

SetPreflight sets Preflight field to given value.

### HasPreflight

`func (o *HandlersRunnerInstallAckRequest) HasPreflight() bool`

HasPreflight returns a boolean if a field has been set.

### GetToken

`func (o *HandlersRunnerInstallAckRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *HandlersRunnerInstallAckRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *HandlersRunnerInstallAckRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *HandlersRunnerInstallAckRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


