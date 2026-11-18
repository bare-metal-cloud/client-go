# HandlersCapabilityDefinitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**[]HandlersCapabilityDefinition**](HandlersCapabilityDefinition.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCapabilityDefinitionsResponse

`func NewHandlersCapabilityDefinitionsResponse() *HandlersCapabilityDefinitionsResponse`

NewHandlersCapabilityDefinitionsResponse instantiates a new HandlersCapabilityDefinitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCapabilityDefinitionsResponseWithDefaults

`func NewHandlersCapabilityDefinitionsResponseWithDefaults() *HandlersCapabilityDefinitionsResponse`

NewHandlersCapabilityDefinitionsResponseWithDefaults instantiates a new HandlersCapabilityDefinitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *HandlersCapabilityDefinitionsResponse) GetCapabilities() []HandlersCapabilityDefinition`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HandlersCapabilityDefinitionsResponse) GetCapabilitiesOk() (*[]HandlersCapabilityDefinition, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HandlersCapabilityDefinitionsResponse) SetCapabilities(v []HandlersCapabilityDefinition)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *HandlersCapabilityDefinitionsResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetVersion

`func (o *HandlersCapabilityDefinitionsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersCapabilityDefinitionsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersCapabilityDefinitionsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HandlersCapabilityDefinitionsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


