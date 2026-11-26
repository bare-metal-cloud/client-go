# HandlersAppsRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPoint** | Pointer to [**HandlersAppsRestoreClonePoint**](HandlersAppsRestoreClonePoint.md) |  | [optional] 
**Confirm** | Pointer to **string** | Confirm is the user&#39;s fresh consent (the dialog&#39;s confirm). Empty on the first POST — the dialog answer. | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** | Destination + Key are the point&#39;s archive handle (the durability endpoint&#39;s point list carries both). | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**NewInstall** | Pointer to **string** | The clone mode&#39;s fields (BAR-4912): NewInstall is the name of the install to create (required on clone), Data is the explicit with/without choice (required on clone), and BackupPoint is the with-data clone&#39;s point handle (absent on without). | [optional] 
**ParallelInstall** | Pointer to **string** | ParallelInstall is the sibling install a parallel restore lands on. Required on parallel; the record must already exist. | [optional] 
**Target** | Pointer to **string** | Target is in_place (default) or parallel. | [optional] 

## Methods

### NewHandlersAppsRestoreRequest

`func NewHandlersAppsRestoreRequest() *HandlersAppsRestoreRequest`

NewHandlersAppsRestoreRequest instantiates a new HandlersAppsRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAppsRestoreRequestWithDefaults

`func NewHandlersAppsRestoreRequestWithDefaults() *HandlersAppsRestoreRequest`

NewHandlersAppsRestoreRequestWithDefaults instantiates a new HandlersAppsRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPoint

`func (o *HandlersAppsRestoreRequest) GetBackupPoint() HandlersAppsRestoreClonePoint`

GetBackupPoint returns the BackupPoint field if non-nil, zero value otherwise.

### GetBackupPointOk

`func (o *HandlersAppsRestoreRequest) GetBackupPointOk() (*HandlersAppsRestoreClonePoint, bool)`

GetBackupPointOk returns a tuple with the BackupPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPoint

`func (o *HandlersAppsRestoreRequest) SetBackupPoint(v HandlersAppsRestoreClonePoint)`

SetBackupPoint sets BackupPoint field to given value.

### HasBackupPoint

`func (o *HandlersAppsRestoreRequest) HasBackupPoint() bool`

HasBackupPoint returns a boolean if a field has been set.

### GetConfirm

`func (o *HandlersAppsRestoreRequest) GetConfirm() string`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *HandlersAppsRestoreRequest) GetConfirmOk() (*string, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *HandlersAppsRestoreRequest) SetConfirm(v string)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *HandlersAppsRestoreRequest) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetData

`func (o *HandlersAppsRestoreRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HandlersAppsRestoreRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HandlersAppsRestoreRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *HandlersAppsRestoreRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDestination

`func (o *HandlersAppsRestoreRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *HandlersAppsRestoreRequest) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *HandlersAppsRestoreRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *HandlersAppsRestoreRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetKey

`func (o *HandlersAppsRestoreRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HandlersAppsRestoreRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HandlersAppsRestoreRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *HandlersAppsRestoreRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNewInstall

`func (o *HandlersAppsRestoreRequest) GetNewInstall() string`

GetNewInstall returns the NewInstall field if non-nil, zero value otherwise.

### GetNewInstallOk

`func (o *HandlersAppsRestoreRequest) GetNewInstallOk() (*string, bool)`

GetNewInstallOk returns a tuple with the NewInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewInstall

`func (o *HandlersAppsRestoreRequest) SetNewInstall(v string)`

SetNewInstall sets NewInstall field to given value.

### HasNewInstall

`func (o *HandlersAppsRestoreRequest) HasNewInstall() bool`

HasNewInstall returns a boolean if a field has been set.

### GetParallelInstall

`func (o *HandlersAppsRestoreRequest) GetParallelInstall() string`

GetParallelInstall returns the ParallelInstall field if non-nil, zero value otherwise.

### GetParallelInstallOk

`func (o *HandlersAppsRestoreRequest) GetParallelInstallOk() (*string, bool)`

GetParallelInstallOk returns a tuple with the ParallelInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelInstall

`func (o *HandlersAppsRestoreRequest) SetParallelInstall(v string)`

SetParallelInstall sets ParallelInstall field to given value.

### HasParallelInstall

`func (o *HandlersAppsRestoreRequest) HasParallelInstall() bool`

HasParallelInstall returns a boolean if a field has been set.

### GetTarget

`func (o *HandlersAppsRestoreRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HandlersAppsRestoreRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HandlersAppsRestoreRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *HandlersAppsRestoreRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


