# HandlersMFAVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCodes** | Pointer to **[]string** | plaintext, ONE-TIME | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersMFAVerifyResponse

`func NewHandlersMFAVerifyResponse() *HandlersMFAVerifyResponse`

NewHandlersMFAVerifyResponse instantiates a new HandlersMFAVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMFAVerifyResponseWithDefaults

`func NewHandlersMFAVerifyResponseWithDefaults() *HandlersMFAVerifyResponse`

NewHandlersMFAVerifyResponseWithDefaults instantiates a new HandlersMFAVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCodes

`func (o *HandlersMFAVerifyResponse) GetBackupCodes() []string`

GetBackupCodes returns the BackupCodes field if non-nil, zero value otherwise.

### GetBackupCodesOk

`func (o *HandlersMFAVerifyResponse) GetBackupCodesOk() (*[]string, bool)`

GetBackupCodesOk returns a tuple with the BackupCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCodes

`func (o *HandlersMFAVerifyResponse) SetBackupCodes(v []string)`

SetBackupCodes sets BackupCodes field to given value.

### HasBackupCodes

`func (o *HandlersMFAVerifyResponse) HasBackupCodes() bool`

HasBackupCodes returns a boolean if a field has been set.

### GetMessage

`func (o *HandlersMFAVerifyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HandlersMFAVerifyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HandlersMFAVerifyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HandlersMFAVerifyResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *HandlersMFAVerifyResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HandlersMFAVerifyResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HandlersMFAVerifyResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HandlersMFAVerifyResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


