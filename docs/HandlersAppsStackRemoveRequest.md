# HandlersAppsStackRemoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirm** | Pointer to **string** | Confirm is the purge dialog&#39;s fresh consent. A keep-data removal needs none; the purge refuses its absence (a bare DELETE can never mint the token). | [optional] 
**RetainBackup** | Pointer to **bool** | RetainBackup is Retain Backup: false additionally purges the final archive and destinations after verification. nil means ON. | [optional] 
**RetainData** | Pointer to **bool** | RetainData is Retain Data (ruling 2026-10-12): false rides the purge chain, which needs the fresh confirm token. The pointer is the defaulting law: nil means the ruling&#39;s ON. | [optional] 

## Methods

### NewHandlersAppsStackRemoveRequest

`func NewHandlersAppsStackRemoveRequest() *HandlersAppsStackRemoveRequest`

NewHandlersAppsStackRemoveRequest instantiates a new HandlersAppsStackRemoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAppsStackRemoveRequestWithDefaults

`func NewHandlersAppsStackRemoveRequestWithDefaults() *HandlersAppsStackRemoveRequest`

NewHandlersAppsStackRemoveRequestWithDefaults instantiates a new HandlersAppsStackRemoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirm

`func (o *HandlersAppsStackRemoveRequest) GetConfirm() string`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *HandlersAppsStackRemoveRequest) GetConfirmOk() (*string, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *HandlersAppsStackRemoveRequest) SetConfirm(v string)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *HandlersAppsStackRemoveRequest) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetRetainBackup

`func (o *HandlersAppsStackRemoveRequest) GetRetainBackup() bool`

GetRetainBackup returns the RetainBackup field if non-nil, zero value otherwise.

### GetRetainBackupOk

`func (o *HandlersAppsStackRemoveRequest) GetRetainBackupOk() (*bool, bool)`

GetRetainBackupOk returns a tuple with the RetainBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainBackup

`func (o *HandlersAppsStackRemoveRequest) SetRetainBackup(v bool)`

SetRetainBackup sets RetainBackup field to given value.

### HasRetainBackup

`func (o *HandlersAppsStackRemoveRequest) HasRetainBackup() bool`

HasRetainBackup returns a boolean if a field has been set.

### GetRetainData

`func (o *HandlersAppsStackRemoveRequest) GetRetainData() bool`

GetRetainData returns the RetainData field if non-nil, zero value otherwise.

### GetRetainDataOk

`func (o *HandlersAppsStackRemoveRequest) GetRetainDataOk() (*bool, bool)`

GetRetainDataOk returns a tuple with the RetainData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainData

`func (o *HandlersAppsStackRemoveRequest) SetRetainData(v bool)`

SetRetainData sets RetainData field to given value.

### HasRetainData

`func (o *HandlersAppsStackRemoveRequest) HasRetainData() bool`

HasRetainData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


