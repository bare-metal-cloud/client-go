# HandlersDelegationMandateView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOps** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExpiryUnix** | Pointer to **int32** |  | [optional] 
**HarnessFingerprint** | Pointer to **string** |  | [optional] 
**InstallConfirmed** | Pointer to **bool** |  | [optional] 
**MachineId** | Pointer to **string** |  | [optional] 
**MandateId** | Pointer to **string** |  | [optional] 
**RateLimitN** | Pointer to **int32** |  | [optional] 
**RateWindowSeconds** | Pointer to **int32** |  | [optional] 
**SignerUserId** | Pointer to **int32** |  | [optional] 
**SpendLimitCents** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDelegationMandateView

`func NewHandlersDelegationMandateView() *HandlersDelegationMandateView`

NewHandlersDelegationMandateView instantiates a new HandlersDelegationMandateView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDelegationMandateViewWithDefaults

`func NewHandlersDelegationMandateViewWithDefaults() *HandlersDelegationMandateView`

NewHandlersDelegationMandateViewWithDefaults instantiates a new HandlersDelegationMandateView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOps

`func (o *HandlersDelegationMandateView) GetAllowedOps() []string`

GetAllowedOps returns the AllowedOps field if non-nil, zero value otherwise.

### GetAllowedOpsOk

`func (o *HandlersDelegationMandateView) GetAllowedOpsOk() (*[]string, bool)`

GetAllowedOpsOk returns a tuple with the AllowedOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOps

`func (o *HandlersDelegationMandateView) SetAllowedOps(v []string)`

SetAllowedOps sets AllowedOps field to given value.

### HasAllowedOps

`func (o *HandlersDelegationMandateView) HasAllowedOps() bool`

HasAllowedOps returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersDelegationMandateView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersDelegationMandateView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersDelegationMandateView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersDelegationMandateView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiryUnix

`func (o *HandlersDelegationMandateView) GetExpiryUnix() int32`

GetExpiryUnix returns the ExpiryUnix field if non-nil, zero value otherwise.

### GetExpiryUnixOk

`func (o *HandlersDelegationMandateView) GetExpiryUnixOk() (*int32, bool)`

GetExpiryUnixOk returns a tuple with the ExpiryUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryUnix

`func (o *HandlersDelegationMandateView) SetExpiryUnix(v int32)`

SetExpiryUnix sets ExpiryUnix field to given value.

### HasExpiryUnix

`func (o *HandlersDelegationMandateView) HasExpiryUnix() bool`

HasExpiryUnix returns a boolean if a field has been set.

### GetHarnessFingerprint

`func (o *HandlersDelegationMandateView) GetHarnessFingerprint() string`

GetHarnessFingerprint returns the HarnessFingerprint field if non-nil, zero value otherwise.

### GetHarnessFingerprintOk

`func (o *HandlersDelegationMandateView) GetHarnessFingerprintOk() (*string, bool)`

GetHarnessFingerprintOk returns a tuple with the HarnessFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarnessFingerprint

`func (o *HandlersDelegationMandateView) SetHarnessFingerprint(v string)`

SetHarnessFingerprint sets HarnessFingerprint field to given value.

### HasHarnessFingerprint

`func (o *HandlersDelegationMandateView) HasHarnessFingerprint() bool`

HasHarnessFingerprint returns a boolean if a field has been set.

### GetInstallConfirmed

`func (o *HandlersDelegationMandateView) GetInstallConfirmed() bool`

GetInstallConfirmed returns the InstallConfirmed field if non-nil, zero value otherwise.

### GetInstallConfirmedOk

`func (o *HandlersDelegationMandateView) GetInstallConfirmedOk() (*bool, bool)`

GetInstallConfirmedOk returns a tuple with the InstallConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallConfirmed

`func (o *HandlersDelegationMandateView) SetInstallConfirmed(v bool)`

SetInstallConfirmed sets InstallConfirmed field to given value.

### HasInstallConfirmed

`func (o *HandlersDelegationMandateView) HasInstallConfirmed() bool`

HasInstallConfirmed returns a boolean if a field has been set.

### GetMachineId

`func (o *HandlersDelegationMandateView) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *HandlersDelegationMandateView) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *HandlersDelegationMandateView) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *HandlersDelegationMandateView) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetMandateId

`func (o *HandlersDelegationMandateView) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *HandlersDelegationMandateView) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *HandlersDelegationMandateView) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.

### HasMandateId

`func (o *HandlersDelegationMandateView) HasMandateId() bool`

HasMandateId returns a boolean if a field has been set.

### GetRateLimitN

`func (o *HandlersDelegationMandateView) GetRateLimitN() int32`

GetRateLimitN returns the RateLimitN field if non-nil, zero value otherwise.

### GetRateLimitNOk

`func (o *HandlersDelegationMandateView) GetRateLimitNOk() (*int32, bool)`

GetRateLimitNOk returns a tuple with the RateLimitN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitN

`func (o *HandlersDelegationMandateView) SetRateLimitN(v int32)`

SetRateLimitN sets RateLimitN field to given value.

### HasRateLimitN

`func (o *HandlersDelegationMandateView) HasRateLimitN() bool`

HasRateLimitN returns a boolean if a field has been set.

### GetRateWindowSeconds

`func (o *HandlersDelegationMandateView) GetRateWindowSeconds() int32`

GetRateWindowSeconds returns the RateWindowSeconds field if non-nil, zero value otherwise.

### GetRateWindowSecondsOk

`func (o *HandlersDelegationMandateView) GetRateWindowSecondsOk() (*int32, bool)`

GetRateWindowSecondsOk returns a tuple with the RateWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateWindowSeconds

`func (o *HandlersDelegationMandateView) SetRateWindowSeconds(v int32)`

SetRateWindowSeconds sets RateWindowSeconds field to given value.

### HasRateWindowSeconds

`func (o *HandlersDelegationMandateView) HasRateWindowSeconds() bool`

HasRateWindowSeconds returns a boolean if a field has been set.

### GetSignerUserId

`func (o *HandlersDelegationMandateView) GetSignerUserId() int32`

GetSignerUserId returns the SignerUserId field if non-nil, zero value otherwise.

### GetSignerUserIdOk

`func (o *HandlersDelegationMandateView) GetSignerUserIdOk() (*int32, bool)`

GetSignerUserIdOk returns a tuple with the SignerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUserId

`func (o *HandlersDelegationMandateView) SetSignerUserId(v int32)`

SetSignerUserId sets SignerUserId field to given value.

### HasSignerUserId

`func (o *HandlersDelegationMandateView) HasSignerUserId() bool`

HasSignerUserId returns a boolean if a field has been set.

### GetSpendLimitCents

`func (o *HandlersDelegationMandateView) GetSpendLimitCents() int32`

GetSpendLimitCents returns the SpendLimitCents field if non-nil, zero value otherwise.

### GetSpendLimitCentsOk

`func (o *HandlersDelegationMandateView) GetSpendLimitCentsOk() (*int32, bool)`

GetSpendLimitCentsOk returns a tuple with the SpendLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitCents

`func (o *HandlersDelegationMandateView) SetSpendLimitCents(v int32)`

SetSpendLimitCents sets SpendLimitCents field to given value.

### HasSpendLimitCents

`func (o *HandlersDelegationMandateView) HasSpendLimitCents() bool`

HasSpendLimitCents returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersDelegationMandateView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersDelegationMandateView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersDelegationMandateView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersDelegationMandateView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


