# HandlersOrgVaultMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **int32** |  | [optional] 
**ClusterId** | Pointer to **int32** | ClusterID is the Platform the machine&#39;s VM belongs to. Omitted when the VM is in no Platform (unassigned) or there is no VM row. Never null. | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint is \&quot;sha256:&lt;hex&gt;\&quot; of the enrolled Ed25519 public key — a hash, never key material. Empty for legacy agents enrolled before ADR-065 Phase 2 (rendered honestly as \&quot;no enrolled key\&quot;). | [optional] 
**HeldCount** | Pointer to **int32** |  | [optional] 
**HeldSecrets** | Pointer to [**[]HandlersHeldSecret**](HandlersHeldSecret.md) |  | [optional] 
**LastDeliveryAt** | Pointer to **string** | LastDeliveryAt is the newest ledger row for this machine (nil when the ledger has nothing for it — an honest \&quot;never\&quot;, not a fabricated date). | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | \&quot;trusted\&quot; | \&quot;revoked\&quot; | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**TrustedSince** | Pointer to **string** |  | [optional] 
**VmId** | Pointer to **int32** | VMID is the machine&#39;s VM row id (BAR-1588 / BAR-1600). Omitted when the agent has no VM row (a kubernetes or imported enrollment). | [optional] 

## Methods

### NewHandlersOrgVaultMachine

`func NewHandlersOrgVaultMachine() *HandlersOrgVaultMachine`

NewHandlersOrgVaultMachine instantiates a new HandlersOrgVaultMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultMachineWithDefaults

`func NewHandlersOrgVaultMachineWithDefaults() *HandlersOrgVaultMachine`

NewHandlersOrgVaultMachineWithDefaults instantiates a new HandlersOrgVaultMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *HandlersOrgVaultMachine) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *HandlersOrgVaultMachine) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *HandlersOrgVaultMachine) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *HandlersOrgVaultMachine) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetClusterId

`func (o *HandlersOrgVaultMachine) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersOrgVaultMachine) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersOrgVaultMachine) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersOrgVaultMachine) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetFingerprint

`func (o *HandlersOrgVaultMachine) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *HandlersOrgVaultMachine) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *HandlersOrgVaultMachine) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *HandlersOrgVaultMachine) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHeldCount

`func (o *HandlersOrgVaultMachine) GetHeldCount() int32`

GetHeldCount returns the HeldCount field if non-nil, zero value otherwise.

### GetHeldCountOk

`func (o *HandlersOrgVaultMachine) GetHeldCountOk() (*int32, bool)`

GetHeldCountOk returns a tuple with the HeldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldCount

`func (o *HandlersOrgVaultMachine) SetHeldCount(v int32)`

SetHeldCount sets HeldCount field to given value.

### HasHeldCount

`func (o *HandlersOrgVaultMachine) HasHeldCount() bool`

HasHeldCount returns a boolean if a field has been set.

### GetHeldSecrets

`func (o *HandlersOrgVaultMachine) GetHeldSecrets() []HandlersHeldSecret`

GetHeldSecrets returns the HeldSecrets field if non-nil, zero value otherwise.

### GetHeldSecretsOk

`func (o *HandlersOrgVaultMachine) GetHeldSecretsOk() (*[]HandlersHeldSecret, bool)`

GetHeldSecretsOk returns a tuple with the HeldSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeldSecrets

`func (o *HandlersOrgVaultMachine) SetHeldSecrets(v []HandlersHeldSecret)`

SetHeldSecrets sets HeldSecrets field to given value.

### HasHeldSecrets

`func (o *HandlersOrgVaultMachine) HasHeldSecrets() bool`

HasHeldSecrets returns a boolean if a field has been set.

### GetLastDeliveryAt

`func (o *HandlersOrgVaultMachine) GetLastDeliveryAt() string`

GetLastDeliveryAt returns the LastDeliveryAt field if non-nil, zero value otherwise.

### GetLastDeliveryAtOk

`func (o *HandlersOrgVaultMachine) GetLastDeliveryAtOk() (*string, bool)`

GetLastDeliveryAtOk returns a tuple with the LastDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveryAt

`func (o *HandlersOrgVaultMachine) SetLastDeliveryAt(v string)`

SetLastDeliveryAt sets LastDeliveryAt field to given value.

### HasLastDeliveryAt

`func (o *HandlersOrgVaultMachine) HasLastDeliveryAt() bool`

HasLastDeliveryAt returns a boolean if a field has been set.

### GetName

`func (o *HandlersOrgVaultMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersOrgVaultMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersOrgVaultMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersOrgVaultMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *HandlersOrgVaultMachine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersOrgVaultMachine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersOrgVaultMachine) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersOrgVaultMachine) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubject

`func (o *HandlersOrgVaultMachine) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *HandlersOrgVaultMachine) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *HandlersOrgVaultMachine) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *HandlersOrgVaultMachine) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTrustedSince

`func (o *HandlersOrgVaultMachine) GetTrustedSince() string`

GetTrustedSince returns the TrustedSince field if non-nil, zero value otherwise.

### GetTrustedSinceOk

`func (o *HandlersOrgVaultMachine) GetTrustedSinceOk() (*string, bool)`

GetTrustedSinceOk returns a tuple with the TrustedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedSince

`func (o *HandlersOrgVaultMachine) SetTrustedSince(v string)`

SetTrustedSince sets TrustedSince field to given value.

### HasTrustedSince

`func (o *HandlersOrgVaultMachine) HasTrustedSince() bool`

HasTrustedSince returns a boolean if a field has been set.

### GetVmId

`func (o *HandlersOrgVaultMachine) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *HandlersOrgVaultMachine) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *HandlersOrgVaultMachine) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *HandlersOrgVaultMachine) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


