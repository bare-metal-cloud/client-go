# HandlersGrantWithStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteMaxExpiresAt** | Pointer to **string** | AbsoluteMaxExpiresAt is the hard cap for sliding-mode grants. Nullable; meaningful only when ExpirationMode &#x3D;&#x3D; \&quot;sliding\&quot;. No amount of activity extends a grant past this time (NFR4 — no uncapped sliding). | [optional] 
**CreatedAt** | Pointer to **string** | Audit trail | [optional] 
**ExpirationMode** | Pointer to **string** | ExpirationMode controls how ExpiresAt behaves (Epic 103 / ADR-078).   - \&quot;fixed\&quot;   : hard deadline at ExpiresAt (default; pre-granted windows).   - \&quot;sliding\&quot; : ExpiresAt resets on each successful CheckAccess, bounded                 by AbsoluteMaxExpiresAt, advancing by IdleTimeout. | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**GrantedById** | Pointer to **int32** | GrantedByID is the admin who created this authorization | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IdleTimeout** | Pointer to [**ModelsDuration**](ModelsDuration.md) | IdleTimeout is the sliding idle window. Nullable; meaningful only when ExpirationMode &#x3D;&#x3D; \&quot;sliding\&quot;. On each successful access the grant&#39;s ExpiresAt advances to now+IdleTimeout (capped at AbsoluteMaxExpiresAt). | [optional] 
**IsJit** | Pointer to **bool** | IsJIT marks grants minted by the JIT MFA-elevation flow (Epic 103 wrapper). It is AUDIT METADATA ONLY — CheckAccess treats JIT-originated and pre-granted rows identically. (The dormant session-blob IsJIT flag from Story 69.1 was removed in Story 103-5; the flag now lives here.) | [optional] 
**LastUsedAt** | Pointer to **string** | LastUsedAt records the most recent successful CheckAccess on a sliding grant. Updated atomically inside the same transaction that extends ExpiresAt. Nullable; only set for sliding grants. | [optional] 
**MemberId** | Pointer to **int32** | MemberID is the ProjectMember receiving access | [optional] 
**OrgId** | Pointer to **int32** | OrgID is the target org for org-scope grants (ADR-075 M1). Nullable: NULL for SA-scope rows. CheckOrgVaultAccess filters on it. | [optional] 
**Reason** | Pointer to **string** | Reason is an optional explanation for the access grant | [optional] 
**RevokedAt** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | Scope discriminates the grant target (ADR-075 M1 / migration 000043):   - \&quot;sa\&quot;  : ServiceAccountID set, OrgID NULL (legacy shape, default)   - \&quot;org\&quot; : OrgID set, ServiceAccountID NULL (org-vault KEK use) Enforced at the DB tier by authorized_accesses_scope_target_chk. | [optional] 
**ServiceAccountId** | Pointer to **int32** | ServiceAccountID is the target service account. Nullable: NULL for org-scope rows (migration 000043 dropped the NOT NULL). | [optional] 
**StartsAt** | Pointer to **string** | Time-bound fields | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersGrantWithStatus

`func NewHandlersGrantWithStatus() *HandlersGrantWithStatus`

NewHandlersGrantWithStatus instantiates a new HandlersGrantWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersGrantWithStatusWithDefaults

`func NewHandlersGrantWithStatusWithDefaults() *HandlersGrantWithStatus`

NewHandlersGrantWithStatusWithDefaults instantiates a new HandlersGrantWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteMaxExpiresAt

`func (o *HandlersGrantWithStatus) GetAbsoluteMaxExpiresAt() string`

GetAbsoluteMaxExpiresAt returns the AbsoluteMaxExpiresAt field if non-nil, zero value otherwise.

### GetAbsoluteMaxExpiresAtOk

`func (o *HandlersGrantWithStatus) GetAbsoluteMaxExpiresAtOk() (*string, bool)`

GetAbsoluteMaxExpiresAtOk returns a tuple with the AbsoluteMaxExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteMaxExpiresAt

`func (o *HandlersGrantWithStatus) SetAbsoluteMaxExpiresAt(v string)`

SetAbsoluteMaxExpiresAt sets AbsoluteMaxExpiresAt field to given value.

### HasAbsoluteMaxExpiresAt

`func (o *HandlersGrantWithStatus) HasAbsoluteMaxExpiresAt() bool`

HasAbsoluteMaxExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersGrantWithStatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersGrantWithStatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersGrantWithStatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersGrantWithStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpirationMode

`func (o *HandlersGrantWithStatus) GetExpirationMode() string`

GetExpirationMode returns the ExpirationMode field if non-nil, zero value otherwise.

### GetExpirationModeOk

`func (o *HandlersGrantWithStatus) GetExpirationModeOk() (*string, bool)`

GetExpirationModeOk returns a tuple with the ExpirationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMode

`func (o *HandlersGrantWithStatus) SetExpirationMode(v string)`

SetExpirationMode sets ExpirationMode field to given value.

### HasExpirationMode

`func (o *HandlersGrantWithStatus) HasExpirationMode() bool`

HasExpirationMode returns a boolean if a field has been set.

### GetExpiresAt

`func (o *HandlersGrantWithStatus) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HandlersGrantWithStatus) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HandlersGrantWithStatus) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HandlersGrantWithStatus) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetGrantedById

`func (o *HandlersGrantWithStatus) GetGrantedById() int32`

GetGrantedById returns the GrantedById field if non-nil, zero value otherwise.

### GetGrantedByIdOk

`func (o *HandlersGrantWithStatus) GetGrantedByIdOk() (*int32, bool)`

GetGrantedByIdOk returns a tuple with the GrantedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedById

`func (o *HandlersGrantWithStatus) SetGrantedById(v int32)`

SetGrantedById sets GrantedById field to given value.

### HasGrantedById

`func (o *HandlersGrantWithStatus) HasGrantedById() bool`

HasGrantedById returns a boolean if a field has been set.

### GetId

`func (o *HandlersGrantWithStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersGrantWithStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersGrantWithStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersGrantWithStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *HandlersGrantWithStatus) GetIdleTimeout() ModelsDuration`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *HandlersGrantWithStatus) GetIdleTimeoutOk() (*ModelsDuration, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *HandlersGrantWithStatus) SetIdleTimeout(v ModelsDuration)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *HandlersGrantWithStatus) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetIsJit

`func (o *HandlersGrantWithStatus) GetIsJit() bool`

GetIsJit returns the IsJit field if non-nil, zero value otherwise.

### GetIsJitOk

`func (o *HandlersGrantWithStatus) GetIsJitOk() (*bool, bool)`

GetIsJitOk returns a tuple with the IsJit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJit

`func (o *HandlersGrantWithStatus) SetIsJit(v bool)`

SetIsJit sets IsJit field to given value.

### HasIsJit

`func (o *HandlersGrantWithStatus) HasIsJit() bool`

HasIsJit returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *HandlersGrantWithStatus) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *HandlersGrantWithStatus) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *HandlersGrantWithStatus) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *HandlersGrantWithStatus) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetMemberId

`func (o *HandlersGrantWithStatus) GetMemberId() int32`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *HandlersGrantWithStatus) GetMemberIdOk() (*int32, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *HandlersGrantWithStatus) SetMemberId(v int32)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *HandlersGrantWithStatus) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetOrgId

`func (o *HandlersGrantWithStatus) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *HandlersGrantWithStatus) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *HandlersGrantWithStatus) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *HandlersGrantWithStatus) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetReason

`func (o *HandlersGrantWithStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersGrantWithStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersGrantWithStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersGrantWithStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRevokedAt

`func (o *HandlersGrantWithStatus) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *HandlersGrantWithStatus) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *HandlersGrantWithStatus) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *HandlersGrantWithStatus) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetScope

`func (o *HandlersGrantWithStatus) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersGrantWithStatus) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersGrantWithStatus) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersGrantWithStatus) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *HandlersGrantWithStatus) GetServiceAccountId() int32`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *HandlersGrantWithStatus) GetServiceAccountIdOk() (*int32, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *HandlersGrantWithStatus) SetServiceAccountId(v int32)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *HandlersGrantWithStatus) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetStartsAt

`func (o *HandlersGrantWithStatus) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *HandlersGrantWithStatus) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *HandlersGrantWithStatus) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *HandlersGrantWithStatus) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersGrantWithStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersGrantWithStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersGrantWithStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersGrantWithStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HandlersGrantWithStatus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlersGrantWithStatus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlersGrantWithStatus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HandlersGrantWithStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


