# EngineportsStrandRecordDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis names the attribution: the lineage tunnel id whose target the record points at (empty on untouched rows). | [optional] 
**CurrentTarget** | Pointer to **string** | CurrentTarget is the record&#39;s stored content (the dead tunnel CNAME). | [optional] 
**Detail** | Pointer to **string** | Detail carries the untouched rows&#39; reason (the unknown-tunnel line or the tag-mismatch line naming the interactive flow). | [optional] 
**Host** | Pointer to **string** | Host is the record&#39;s hostname, lowercased, no trailing dot. | [optional] 
**NewTarget** | Pointer to **string** | NewTarget is the live tunnel&#39;s CNAME — the write phase&#39;s destination (empty on untouched rows: they are never written). | [optional] 
**Outcome** | Pointer to **string** | Outcome is StrandOutcomeCandidate or StrandOutcomeUntouched. | [optional] 
**Proxied** | Pointer to **bool** | Proxied is the record&#39;s preserved orange-cloud flag as read at the preview (the write re-stamps proxied per the tunnel-target law, R16 C17 — a tunnel CNAME routes only orange-clouded). | [optional] 
**RecordType** | Pointer to **string** | RecordType is the provider&#39;s record type (a stranded tunnel record is a CNAME; only CNAMEs are judged). | [optional] 
**Tagged** | Pointer to **bool** | Tagged reports whether the record carries the platform ownership tag. | [optional] 
**Zone** | Pointer to **string** | Zone is the record&#39;s zone FQDN (the registry&#39;s normalized form). | [optional] 
**ZoneId** | Pointer to **int32** | ZoneID is the zone&#39;s registry row id (the report&#39;s deep link). | [optional] 

## Methods

### NewEngineportsStrandRecordDiff

`func NewEngineportsStrandRecordDiff() *EngineportsStrandRecordDiff`

NewEngineportsStrandRecordDiff instantiates a new EngineportsStrandRecordDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsStrandRecordDiffWithDefaults

`func NewEngineportsStrandRecordDiffWithDefaults() *EngineportsStrandRecordDiff`

NewEngineportsStrandRecordDiffWithDefaults instantiates a new EngineportsStrandRecordDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *EngineportsStrandRecordDiff) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *EngineportsStrandRecordDiff) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *EngineportsStrandRecordDiff) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *EngineportsStrandRecordDiff) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetCurrentTarget

`func (o *EngineportsStrandRecordDiff) GetCurrentTarget() string`

GetCurrentTarget returns the CurrentTarget field if non-nil, zero value otherwise.

### GetCurrentTargetOk

`func (o *EngineportsStrandRecordDiff) GetCurrentTargetOk() (*string, bool)`

GetCurrentTargetOk returns a tuple with the CurrentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTarget

`func (o *EngineportsStrandRecordDiff) SetCurrentTarget(v string)`

SetCurrentTarget sets CurrentTarget field to given value.

### HasCurrentTarget

`func (o *EngineportsStrandRecordDiff) HasCurrentTarget() bool`

HasCurrentTarget returns a boolean if a field has been set.

### GetDetail

`func (o *EngineportsStrandRecordDiff) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EngineportsStrandRecordDiff) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EngineportsStrandRecordDiff) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *EngineportsStrandRecordDiff) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetHost

`func (o *EngineportsStrandRecordDiff) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EngineportsStrandRecordDiff) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EngineportsStrandRecordDiff) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EngineportsStrandRecordDiff) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNewTarget

`func (o *EngineportsStrandRecordDiff) GetNewTarget() string`

GetNewTarget returns the NewTarget field if non-nil, zero value otherwise.

### GetNewTargetOk

`func (o *EngineportsStrandRecordDiff) GetNewTargetOk() (*string, bool)`

GetNewTargetOk returns a tuple with the NewTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTarget

`func (o *EngineportsStrandRecordDiff) SetNewTarget(v string)`

SetNewTarget sets NewTarget field to given value.

### HasNewTarget

`func (o *EngineportsStrandRecordDiff) HasNewTarget() bool`

HasNewTarget returns a boolean if a field has been set.

### GetOutcome

`func (o *EngineportsStrandRecordDiff) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *EngineportsStrandRecordDiff) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *EngineportsStrandRecordDiff) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *EngineportsStrandRecordDiff) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetProxied

`func (o *EngineportsStrandRecordDiff) GetProxied() bool`

GetProxied returns the Proxied field if non-nil, zero value otherwise.

### GetProxiedOk

`func (o *EngineportsStrandRecordDiff) GetProxiedOk() (*bool, bool)`

GetProxiedOk returns a tuple with the Proxied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxied

`func (o *EngineportsStrandRecordDiff) SetProxied(v bool)`

SetProxied sets Proxied field to given value.

### HasProxied

`func (o *EngineportsStrandRecordDiff) HasProxied() bool`

HasProxied returns a boolean if a field has been set.

### GetRecordType

`func (o *EngineportsStrandRecordDiff) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *EngineportsStrandRecordDiff) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *EngineportsStrandRecordDiff) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *EngineportsStrandRecordDiff) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTagged

`func (o *EngineportsStrandRecordDiff) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *EngineportsStrandRecordDiff) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *EngineportsStrandRecordDiff) SetTagged(v bool)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *EngineportsStrandRecordDiff) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetZone

`func (o *EngineportsStrandRecordDiff) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *EngineportsStrandRecordDiff) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *EngineportsStrandRecordDiff) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *EngineportsStrandRecordDiff) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZoneId

`func (o *EngineportsStrandRecordDiff) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *EngineportsStrandRecordDiff) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *EngineportsStrandRecordDiff) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *EngineportsStrandRecordDiff) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


