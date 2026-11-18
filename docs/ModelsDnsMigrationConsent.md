# ModelsDnsMigrationConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int32** | ClusterID is the NEW claimant cluster the consented deploy targets — the deploy whose dns-apply payload stamps this consent consumes it. A consent never matches a deploy from any other cluster. | [optional] 
**ConsentedAt** | Pointer to **string** | ConsentedAt is the recording instant (the wire consent&#39;s recorded_at, RFC3339 UTC on the wire). | [optional] 
**ConsentedByUserId** | Pointer to **int32** | ConsentedByUserID is the org admin/owner who recorded the consent (the audit attribution; the wire consent&#39;s recorded_by rides it). | [optional] 
**ConsumedAt** | Pointer to **string** | ConsumedAt is when a consented deploy took the record over (or a newer consent superseded this one). NULL means the consent is LIVE. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ForeignArrival** | Pointer to **string** | ForeignArrival is the live record&#39;s tag arrival (the identity&#39;s other half — stamped onto the wire consent verbatim, and the claim release&#39;s narrowing term when non-empty). | [optional] 
**ForeignClusterId** | Pointer to **string** | ForeignClusterID is the live record&#39;s tag identity VERBATIM (the numeric cluster row id in string form — the engine&#39;s ownership-tag grammar). The stamped wire consent carries it verbatim, and the executor&#39;s re-assertion compares it to the live tag. | [optional] 
**ForeignClusterRowId** | Pointer to **int32** | ForeignClusterRowID is the tag identity resolved to the live cluster row at record time — the OLD claimant whose host claim the terminal consumption releases (never the whole-cluster release: only this host&#39;s row lets go). | [optional] 
**Host** | Pointer to **string** | Host is the migrated hostname, normalized (lowercase, no trailing dot — the registry&#39;s FQDN form). | [optional] 
**Id** | Pointer to **int32** | ID is the row&#39;s surrogate key (the revoke endpoint&#39;s path id). | [optional] 
**OrgId** | Pointer to **int32** | OrgID is the hard tenancy boundary: every read and write runs org-scoped. It leads both indexes (the live-consent read and the consume&#39;s scan are org-scoped first). | [optional] 
**RecordType** | Pointer to **string** | RecordType is the DNS record type the availability check reported (informational — the engine&#39;s executor reads the live record set). | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int32** | ZoneID is the registered zone row the hostname lives in (the registry&#39;s org-scoped table; the consent is zone-scoped because the record&#39;s custody is). | [optional] 

## Methods

### NewModelsDnsMigrationConsent

`func NewModelsDnsMigrationConsent() *ModelsDnsMigrationConsent`

NewModelsDnsMigrationConsent instantiates a new ModelsDnsMigrationConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsDnsMigrationConsentWithDefaults

`func NewModelsDnsMigrationConsentWithDefaults() *ModelsDnsMigrationConsent`

NewModelsDnsMigrationConsentWithDefaults instantiates a new ModelsDnsMigrationConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ModelsDnsMigrationConsent) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsDnsMigrationConsent) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsDnsMigrationConsent) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsDnsMigrationConsent) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConsentedAt

`func (o *ModelsDnsMigrationConsent) GetConsentedAt() string`

GetConsentedAt returns the ConsentedAt field if non-nil, zero value otherwise.

### GetConsentedAtOk

`func (o *ModelsDnsMigrationConsent) GetConsentedAtOk() (*string, bool)`

GetConsentedAtOk returns a tuple with the ConsentedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedAt

`func (o *ModelsDnsMigrationConsent) SetConsentedAt(v string)`

SetConsentedAt sets ConsentedAt field to given value.

### HasConsentedAt

`func (o *ModelsDnsMigrationConsent) HasConsentedAt() bool`

HasConsentedAt returns a boolean if a field has been set.

### GetConsentedByUserId

`func (o *ModelsDnsMigrationConsent) GetConsentedByUserId() int32`

GetConsentedByUserId returns the ConsentedByUserId field if non-nil, zero value otherwise.

### GetConsentedByUserIdOk

`func (o *ModelsDnsMigrationConsent) GetConsentedByUserIdOk() (*int32, bool)`

GetConsentedByUserIdOk returns a tuple with the ConsentedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedByUserId

`func (o *ModelsDnsMigrationConsent) SetConsentedByUserId(v int32)`

SetConsentedByUserId sets ConsentedByUserId field to given value.

### HasConsentedByUserId

`func (o *ModelsDnsMigrationConsent) HasConsentedByUserId() bool`

HasConsentedByUserId returns a boolean if a field has been set.

### GetConsumedAt

`func (o *ModelsDnsMigrationConsent) GetConsumedAt() string`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *ModelsDnsMigrationConsent) GetConsumedAtOk() (*string, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *ModelsDnsMigrationConsent) SetConsumedAt(v string)`

SetConsumedAt sets ConsumedAt field to given value.

### HasConsumedAt

`func (o *ModelsDnsMigrationConsent) HasConsumedAt() bool`

HasConsumedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsDnsMigrationConsent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsDnsMigrationConsent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsDnsMigrationConsent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsDnsMigrationConsent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetForeignArrival

`func (o *ModelsDnsMigrationConsent) GetForeignArrival() string`

GetForeignArrival returns the ForeignArrival field if non-nil, zero value otherwise.

### GetForeignArrivalOk

`func (o *ModelsDnsMigrationConsent) GetForeignArrivalOk() (*string, bool)`

GetForeignArrivalOk returns a tuple with the ForeignArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignArrival

`func (o *ModelsDnsMigrationConsent) SetForeignArrival(v string)`

SetForeignArrival sets ForeignArrival field to given value.

### HasForeignArrival

`func (o *ModelsDnsMigrationConsent) HasForeignArrival() bool`

HasForeignArrival returns a boolean if a field has been set.

### GetForeignClusterId

`func (o *ModelsDnsMigrationConsent) GetForeignClusterId() string`

GetForeignClusterId returns the ForeignClusterId field if non-nil, zero value otherwise.

### GetForeignClusterIdOk

`func (o *ModelsDnsMigrationConsent) GetForeignClusterIdOk() (*string, bool)`

GetForeignClusterIdOk returns a tuple with the ForeignClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignClusterId

`func (o *ModelsDnsMigrationConsent) SetForeignClusterId(v string)`

SetForeignClusterId sets ForeignClusterId field to given value.

### HasForeignClusterId

`func (o *ModelsDnsMigrationConsent) HasForeignClusterId() bool`

HasForeignClusterId returns a boolean if a field has been set.

### GetForeignClusterRowId

`func (o *ModelsDnsMigrationConsent) GetForeignClusterRowId() int32`

GetForeignClusterRowId returns the ForeignClusterRowId field if non-nil, zero value otherwise.

### GetForeignClusterRowIdOk

`func (o *ModelsDnsMigrationConsent) GetForeignClusterRowIdOk() (*int32, bool)`

GetForeignClusterRowIdOk returns a tuple with the ForeignClusterRowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignClusterRowId

`func (o *ModelsDnsMigrationConsent) SetForeignClusterRowId(v int32)`

SetForeignClusterRowId sets ForeignClusterRowId field to given value.

### HasForeignClusterRowId

`func (o *ModelsDnsMigrationConsent) HasForeignClusterRowId() bool`

HasForeignClusterRowId returns a boolean if a field has been set.

### GetHost

`func (o *ModelsDnsMigrationConsent) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ModelsDnsMigrationConsent) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ModelsDnsMigrationConsent) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ModelsDnsMigrationConsent) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ModelsDnsMigrationConsent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsDnsMigrationConsent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsDnsMigrationConsent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsDnsMigrationConsent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsDnsMigrationConsent) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsDnsMigrationConsent) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsDnsMigrationConsent) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsDnsMigrationConsent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRecordType

`func (o *ModelsDnsMigrationConsent) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ModelsDnsMigrationConsent) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ModelsDnsMigrationConsent) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ModelsDnsMigrationConsent) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsDnsMigrationConsent) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsDnsMigrationConsent) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsDnsMigrationConsent) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsDnsMigrationConsent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetZoneId

`func (o *ModelsDnsMigrationConsent) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ModelsDnsMigrationConsent) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ModelsDnsMigrationConsent) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ModelsDnsMigrationConsent) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


