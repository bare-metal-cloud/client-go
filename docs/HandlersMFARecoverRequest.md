# HandlersMFARecoverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCode** | **string** | Backup code | 
**MfaToken** | **string** | Single-use challenge token | 

## Methods

### NewHandlersMFARecoverRequest

`func NewHandlersMFARecoverRequest(backupCode string, mfaToken string, ) *HandlersMFARecoverRequest`

NewHandlersMFARecoverRequest instantiates a new HandlersMFARecoverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMFARecoverRequestWithDefaults

`func NewHandlersMFARecoverRequestWithDefaults() *HandlersMFARecoverRequest`

NewHandlersMFARecoverRequestWithDefaults instantiates a new HandlersMFARecoverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCode

`func (o *HandlersMFARecoverRequest) GetBackupCode() string`

GetBackupCode returns the BackupCode field if non-nil, zero value otherwise.

### GetBackupCodeOk

`func (o *HandlersMFARecoverRequest) GetBackupCodeOk() (*string, bool)`

GetBackupCodeOk returns a tuple with the BackupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCode

`func (o *HandlersMFARecoverRequest) SetBackupCode(v string)`

SetBackupCode sets BackupCode field to given value.


### GetMfaToken

`func (o *HandlersMFARecoverRequest) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *HandlersMFARecoverRequest) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *HandlersMFARecoverRequest) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


