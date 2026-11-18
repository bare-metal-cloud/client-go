# ModelsDnsZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CredentialId** | Pointer to **int32** | CredentialID references the registry credential row the zone rides. A credential may serve MANY zones (the fingerprint-sharing law); deletion requires zero referencing zones (named refusal). | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LegacyMigrated** | Pointer to **bool** | LegacyMigrated is the MIGRATION MARKER (R12 round 9): the legacy lane&#39;s resolution key. A stored 133.1-era plan (provider cloudflare, no ZoneRef) resolves ONLY the org&#39;s marker-carrying row — never a First() over all cloudflare rows. Exactly one row per org+provider carries the marker (the 133.1 cardinality); deleting the adopted zone strands the legacy lane into the zone-gone class. | [optional] 
**Nameservers** | Pointer to **string** | Nameservers is the provider&#39;s DECLARED nameserver set at connect (comma-joined, normalized — no trailing dots). The deploy preflight compares the live TLD-parent delegation against this stored set and refuses a mismatch (R15 C3: the green-but-dark trap). | [optional] 
**OrgId** | Pointer to **int32** | OrgID is the tenancy boundary; the (org, zone_fqdn) pair is the natural key. | [optional] 
**Provider** | Pointer to **string** | Provider names the DNS provider (v1: \&quot;cloudflare\&quot;). | [optional] 
**ProviderAccountId** | Pointer to **string** | ProviderAccountID is the provider account the zone lives in (R13 — load-bearing for the same-account preflight). Captured at connect from the zone listing response at zero extra cost; ADOPTED rows carry it empty until the first Cloudflare-tunnel preflight lazily backfills it (an empty account id NEVER passes the preflight silently, R13 round 7). | [optional] 
**ProviderZoneId** | Pointer to **string** | ProviderZoneID is the provider&#39;s opaque zone id. | [optional] 
**Status** | Pointer to **string** | Status is the zone&#39;s health posture (the ZoneStatus* vocabulary above: active | amber | the three drift classes). | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VerifiedAt** | Pointer to **string** | VerifiedAt is when the zone last passed a live classification (connect or preflight). Nil for adopted rows until their first preflight. | [optional] 
**ZoneFqdn** | Pointer to **string** | ZoneFQDN is the zone&#39;s fully qualified domain name (no trailing dot, punycode form). | [optional] 

## Methods

### NewModelsDnsZone

`func NewModelsDnsZone() *ModelsDnsZone`

NewModelsDnsZone instantiates a new ModelsDnsZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsDnsZoneWithDefaults

`func NewModelsDnsZoneWithDefaults() *ModelsDnsZone`

NewModelsDnsZoneWithDefaults instantiates a new ModelsDnsZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelsDnsZone) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsDnsZone) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsDnsZone) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsDnsZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialId

`func (o *ModelsDnsZone) GetCredentialId() int32`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *ModelsDnsZone) GetCredentialIdOk() (*int32, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *ModelsDnsZone) SetCredentialId(v int32)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *ModelsDnsZone) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetId

`func (o *ModelsDnsZone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsDnsZone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsDnsZone) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsDnsZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLegacyMigrated

`func (o *ModelsDnsZone) GetLegacyMigrated() bool`

GetLegacyMigrated returns the LegacyMigrated field if non-nil, zero value otherwise.

### GetLegacyMigratedOk

`func (o *ModelsDnsZone) GetLegacyMigratedOk() (*bool, bool)`

GetLegacyMigratedOk returns a tuple with the LegacyMigrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyMigrated

`func (o *ModelsDnsZone) SetLegacyMigrated(v bool)`

SetLegacyMigrated sets LegacyMigrated field to given value.

### HasLegacyMigrated

`func (o *ModelsDnsZone) HasLegacyMigrated() bool`

HasLegacyMigrated returns a boolean if a field has been set.

### GetNameservers

`func (o *ModelsDnsZone) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ModelsDnsZone) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ModelsDnsZone) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *ModelsDnsZone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsDnsZone) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsDnsZone) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsDnsZone) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsDnsZone) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelsDnsZone) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsDnsZone) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsDnsZone) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelsDnsZone) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderAccountId

`func (o *ModelsDnsZone) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *ModelsDnsZone) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *ModelsDnsZone) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *ModelsDnsZone) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.

### GetProviderZoneId

`func (o *ModelsDnsZone) GetProviderZoneId() string`

GetProviderZoneId returns the ProviderZoneId field if non-nil, zero value otherwise.

### GetProviderZoneIdOk

`func (o *ModelsDnsZone) GetProviderZoneIdOk() (*string, bool)`

GetProviderZoneIdOk returns a tuple with the ProviderZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderZoneId

`func (o *ModelsDnsZone) SetProviderZoneId(v string)`

SetProviderZoneId sets ProviderZoneId field to given value.

### HasProviderZoneId

`func (o *ModelsDnsZone) HasProviderZoneId() bool`

HasProviderZoneId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsDnsZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsDnsZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsDnsZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsDnsZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsDnsZone) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsDnsZone) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsDnsZone) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsDnsZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *ModelsDnsZone) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *ModelsDnsZone) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *ModelsDnsZone) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *ModelsDnsZone) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetZoneFqdn

`func (o *ModelsDnsZone) GetZoneFqdn() string`

GetZoneFqdn returns the ZoneFqdn field if non-nil, zero value otherwise.

### GetZoneFqdnOk

`func (o *ModelsDnsZone) GetZoneFqdnOk() (*string, bool)`

GetZoneFqdnOk returns a tuple with the ZoneFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFqdn

`func (o *ModelsDnsZone) SetZoneFqdn(v string)`

SetZoneFqdn sets ZoneFqdn field to given value.

### HasZoneFqdn

`func (o *ModelsDnsZone) HasZoneFqdn() bool`

HasZoneFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


