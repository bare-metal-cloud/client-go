# HandlersVMConsoleDevModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled is true when dev mode is being turned ON, false when turned OFF. The FE only fires this endpoint on the ON transition (the story specifies \&quot;one-time audit signal\&quot;); we accept the field anyway so a future caller can record the OFF transition without an API change. | [optional] 

## Methods

### NewHandlersVMConsoleDevModeRequest

`func NewHandlersVMConsoleDevModeRequest() *HandlersVMConsoleDevModeRequest`

NewHandlersVMConsoleDevModeRequest instantiates a new HandlersVMConsoleDevModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVMConsoleDevModeRequestWithDefaults

`func NewHandlersVMConsoleDevModeRequestWithDefaults() *HandlersVMConsoleDevModeRequest`

NewHandlersVMConsoleDevModeRequestWithDefaults instantiates a new HandlersVMConsoleDevModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *HandlersVMConsoleDevModeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HandlersVMConsoleDevModeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HandlersVMConsoleDevModeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HandlersVMConsoleDevModeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


