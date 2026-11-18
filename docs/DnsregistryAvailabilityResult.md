# DnsregistryAvailabilityResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignTarget** | Pointer to **string** | ForeignTarget shows the existing record&#39;s target when State &#x3D;&#x3D; taken (the R6 refusal&#39;s \&quot;the existing target shown\&quot; courtesy). | [optional] 
**Host** | Pointer to **string** | Host is the checked hostname (normalized). | [optional] 
**MigrationEligible** | Pointer to **bool** | MigrationEligible reports, when State &#x3D;&#x3D; ours, whether the tag&#39;s cluster id names a LIVE cluster row of THIS org (issue #2713 — the record-migration flow&#39;s green light: only then can an org admin record a migration consent for the host). Strictly read-only: a failed lookup answers false and never fails the availability answer (the honest degrade — eligible&#x3D;false is exactly what the consent surface would refuse on). | [optional] 
**Ours** | Pointer to [**DnsregistryOwnershipIdentity**](DnsregistryOwnershipIdentity.md) | Ours carries the ownership identity when State &#x3D;&#x3D; ours (the tag&#39;s cluster and arrival — the same grammar the executor composes). | [optional] 
**State** | Pointer to **string** | State: available (no record), ours (a BMC-tagged record), or taken (a foreign record). The R6 law&#39;s read-only face. | [optional] 

## Methods

### NewDnsregistryAvailabilityResult

`func NewDnsregistryAvailabilityResult() *DnsregistryAvailabilityResult`

NewDnsregistryAvailabilityResult instantiates a new DnsregistryAvailabilityResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsregistryAvailabilityResultWithDefaults

`func NewDnsregistryAvailabilityResultWithDefaults() *DnsregistryAvailabilityResult`

NewDnsregistryAvailabilityResultWithDefaults instantiates a new DnsregistryAvailabilityResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForeignTarget

`func (o *DnsregistryAvailabilityResult) GetForeignTarget() string`

GetForeignTarget returns the ForeignTarget field if non-nil, zero value otherwise.

### GetForeignTargetOk

`func (o *DnsregistryAvailabilityResult) GetForeignTargetOk() (*string, bool)`

GetForeignTargetOk returns a tuple with the ForeignTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTarget

`func (o *DnsregistryAvailabilityResult) SetForeignTarget(v string)`

SetForeignTarget sets ForeignTarget field to given value.

### HasForeignTarget

`func (o *DnsregistryAvailabilityResult) HasForeignTarget() bool`

HasForeignTarget returns a boolean if a field has been set.

### GetHost

`func (o *DnsregistryAvailabilityResult) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DnsregistryAvailabilityResult) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DnsregistryAvailabilityResult) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DnsregistryAvailabilityResult) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMigrationEligible

`func (o *DnsregistryAvailabilityResult) GetMigrationEligible() bool`

GetMigrationEligible returns the MigrationEligible field if non-nil, zero value otherwise.

### GetMigrationEligibleOk

`func (o *DnsregistryAvailabilityResult) GetMigrationEligibleOk() (*bool, bool)`

GetMigrationEligibleOk returns a tuple with the MigrationEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationEligible

`func (o *DnsregistryAvailabilityResult) SetMigrationEligible(v bool)`

SetMigrationEligible sets MigrationEligible field to given value.

### HasMigrationEligible

`func (o *DnsregistryAvailabilityResult) HasMigrationEligible() bool`

HasMigrationEligible returns a boolean if a field has been set.

### GetOurs

`func (o *DnsregistryAvailabilityResult) GetOurs() DnsregistryOwnershipIdentity`

GetOurs returns the Ours field if non-nil, zero value otherwise.

### GetOursOk

`func (o *DnsregistryAvailabilityResult) GetOursOk() (*DnsregistryOwnershipIdentity, bool)`

GetOursOk returns a tuple with the Ours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurs

`func (o *DnsregistryAvailabilityResult) SetOurs(v DnsregistryOwnershipIdentity)`

SetOurs sets Ours field to given value.

### HasOurs

`func (o *DnsregistryAvailabilityResult) HasOurs() bool`

HasOurs returns a boolean if a field has been set.

### GetState

`func (o *DnsregistryAvailabilityResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DnsregistryAvailabilityResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DnsregistryAvailabilityResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DnsregistryAvailabilityResult) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


