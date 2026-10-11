# HandlersVaultRotationAdminStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobsNeedingRewrap** | Pointer to **bool** |  | [optional] 
**CurrentKeyState** | Pointer to [**HandlersCurrentKeyState**](HandlersCurrentKeyState.md) |  | [optional] 
**RecentEvents** | Pointer to [**[]HandlersVaultRotationEvent**](HandlersVaultRotationEvent.md) |  | [optional] 

## Methods

### NewHandlersVaultRotationAdminStatusResponse

`func NewHandlersVaultRotationAdminStatusResponse() *HandlersVaultRotationAdminStatusResponse`

NewHandlersVaultRotationAdminStatusResponse instantiates a new HandlersVaultRotationAdminStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVaultRotationAdminStatusResponseWithDefaults

`func NewHandlersVaultRotationAdminStatusResponseWithDefaults() *HandlersVaultRotationAdminStatusResponse`

NewHandlersVaultRotationAdminStatusResponseWithDefaults instantiates a new HandlersVaultRotationAdminStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobsNeedingRewrap

`func (o *HandlersVaultRotationAdminStatusResponse) GetBlobsNeedingRewrap() bool`

GetBlobsNeedingRewrap returns the BlobsNeedingRewrap field if non-nil, zero value otherwise.

### GetBlobsNeedingRewrapOk

`func (o *HandlersVaultRotationAdminStatusResponse) GetBlobsNeedingRewrapOk() (*bool, bool)`

GetBlobsNeedingRewrapOk returns a tuple with the BlobsNeedingRewrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobsNeedingRewrap

`func (o *HandlersVaultRotationAdminStatusResponse) SetBlobsNeedingRewrap(v bool)`

SetBlobsNeedingRewrap sets BlobsNeedingRewrap field to given value.

### HasBlobsNeedingRewrap

`func (o *HandlersVaultRotationAdminStatusResponse) HasBlobsNeedingRewrap() bool`

HasBlobsNeedingRewrap returns a boolean if a field has been set.

### GetCurrentKeyState

`func (o *HandlersVaultRotationAdminStatusResponse) GetCurrentKeyState() HandlersCurrentKeyState`

GetCurrentKeyState returns the CurrentKeyState field if non-nil, zero value otherwise.

### GetCurrentKeyStateOk

`func (o *HandlersVaultRotationAdminStatusResponse) GetCurrentKeyStateOk() (*HandlersCurrentKeyState, bool)`

GetCurrentKeyStateOk returns a tuple with the CurrentKeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKeyState

`func (o *HandlersVaultRotationAdminStatusResponse) SetCurrentKeyState(v HandlersCurrentKeyState)`

SetCurrentKeyState sets CurrentKeyState field to given value.

### HasCurrentKeyState

`func (o *HandlersVaultRotationAdminStatusResponse) HasCurrentKeyState() bool`

HasCurrentKeyState returns a boolean if a field has been set.

### GetRecentEvents

`func (o *HandlersVaultRotationAdminStatusResponse) GetRecentEvents() []HandlersVaultRotationEvent`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *HandlersVaultRotationAdminStatusResponse) GetRecentEventsOk() (*[]HandlersVaultRotationEvent, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *HandlersVaultRotationAdminStatusResponse) SetRecentEvents(v []HandlersVaultRotationEvent)`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *HandlersVaultRotationAdminStatusResponse) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


