# CloneplanSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointsApproved** | Pointer to **[]string** | EndpointsApproved / EndpointsPending name the binding-used link-external endpoints. | [optional] 
**EndpointsPending** | Pointer to **[]string** |  | [optional] 
**EnvLocalApproved** | Pointer to [**[]CloneplanEnvLocalItem**](CloneplanEnvLocalItem.md) | EnvLocalApproved / EnvLocalPending are the env_local decision items. The approved ones carry the value beside the source value — the confirmation view of what the replay will render; the pending ones are the exceptions&#39; company (they also appear in Exceptions). | [optional] 
**EnvLocalPending** | Pointer to [**[]CloneplanEnvLocalItem**](CloneplanEnvLocalItem.md) |  | [optional] 
**Exceptions** | Pointer to [**[]CloneplanException**](CloneplanException.md) |  | [optional] 
**State** | Pointer to **string** | State is the plan&#39;s own state (planned | nothing-to-clone), echoed so a caller renders the empty plan honestly. | [optional] 
**UnitsAccepted** | Pointer to **int32** |  | [optional] 
**UnitsExcluded** | Pointer to **int32** |  | [optional] 
**UnitsTotal** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloneplanSummary

`func NewCloneplanSummary() *CloneplanSummary`

NewCloneplanSummary instantiates a new CloneplanSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanSummaryWithDefaults

`func NewCloneplanSummaryWithDefaults() *CloneplanSummary`

NewCloneplanSummaryWithDefaults instantiates a new CloneplanSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointsApproved

`func (o *CloneplanSummary) GetEndpointsApproved() []string`

GetEndpointsApproved returns the EndpointsApproved field if non-nil, zero value otherwise.

### GetEndpointsApprovedOk

`func (o *CloneplanSummary) GetEndpointsApprovedOk() (*[]string, bool)`

GetEndpointsApprovedOk returns a tuple with the EndpointsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsApproved

`func (o *CloneplanSummary) SetEndpointsApproved(v []string)`

SetEndpointsApproved sets EndpointsApproved field to given value.

### HasEndpointsApproved

`func (o *CloneplanSummary) HasEndpointsApproved() bool`

HasEndpointsApproved returns a boolean if a field has been set.

### GetEndpointsPending

`func (o *CloneplanSummary) GetEndpointsPending() []string`

GetEndpointsPending returns the EndpointsPending field if non-nil, zero value otherwise.

### GetEndpointsPendingOk

`func (o *CloneplanSummary) GetEndpointsPendingOk() (*[]string, bool)`

GetEndpointsPendingOk returns a tuple with the EndpointsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsPending

`func (o *CloneplanSummary) SetEndpointsPending(v []string)`

SetEndpointsPending sets EndpointsPending field to given value.

### HasEndpointsPending

`func (o *CloneplanSummary) HasEndpointsPending() bool`

HasEndpointsPending returns a boolean if a field has been set.

### GetEnvLocalApproved

`func (o *CloneplanSummary) GetEnvLocalApproved() []CloneplanEnvLocalItem`

GetEnvLocalApproved returns the EnvLocalApproved field if non-nil, zero value otherwise.

### GetEnvLocalApprovedOk

`func (o *CloneplanSummary) GetEnvLocalApprovedOk() (*[]CloneplanEnvLocalItem, bool)`

GetEnvLocalApprovedOk returns a tuple with the EnvLocalApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvLocalApproved

`func (o *CloneplanSummary) SetEnvLocalApproved(v []CloneplanEnvLocalItem)`

SetEnvLocalApproved sets EnvLocalApproved field to given value.

### HasEnvLocalApproved

`func (o *CloneplanSummary) HasEnvLocalApproved() bool`

HasEnvLocalApproved returns a boolean if a field has been set.

### GetEnvLocalPending

`func (o *CloneplanSummary) GetEnvLocalPending() []CloneplanEnvLocalItem`

GetEnvLocalPending returns the EnvLocalPending field if non-nil, zero value otherwise.

### GetEnvLocalPendingOk

`func (o *CloneplanSummary) GetEnvLocalPendingOk() (*[]CloneplanEnvLocalItem, bool)`

GetEnvLocalPendingOk returns a tuple with the EnvLocalPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvLocalPending

`func (o *CloneplanSummary) SetEnvLocalPending(v []CloneplanEnvLocalItem)`

SetEnvLocalPending sets EnvLocalPending field to given value.

### HasEnvLocalPending

`func (o *CloneplanSummary) HasEnvLocalPending() bool`

HasEnvLocalPending returns a boolean if a field has been set.

### GetExceptions

`func (o *CloneplanSummary) GetExceptions() []CloneplanException`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *CloneplanSummary) GetExceptionsOk() (*[]CloneplanException, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *CloneplanSummary) SetExceptions(v []CloneplanException)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *CloneplanSummary) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### GetState

`func (o *CloneplanSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloneplanSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloneplanSummary) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloneplanSummary) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUnitsAccepted

`func (o *CloneplanSummary) GetUnitsAccepted() int32`

GetUnitsAccepted returns the UnitsAccepted field if non-nil, zero value otherwise.

### GetUnitsAcceptedOk

`func (o *CloneplanSummary) GetUnitsAcceptedOk() (*int32, bool)`

GetUnitsAcceptedOk returns a tuple with the UnitsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsAccepted

`func (o *CloneplanSummary) SetUnitsAccepted(v int32)`

SetUnitsAccepted sets UnitsAccepted field to given value.

### HasUnitsAccepted

`func (o *CloneplanSummary) HasUnitsAccepted() bool`

HasUnitsAccepted returns a boolean if a field has been set.

### GetUnitsExcluded

`func (o *CloneplanSummary) GetUnitsExcluded() int32`

GetUnitsExcluded returns the UnitsExcluded field if non-nil, zero value otherwise.

### GetUnitsExcludedOk

`func (o *CloneplanSummary) GetUnitsExcludedOk() (*int32, bool)`

GetUnitsExcludedOk returns a tuple with the UnitsExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsExcluded

`func (o *CloneplanSummary) SetUnitsExcluded(v int32)`

SetUnitsExcluded sets UnitsExcluded field to given value.

### HasUnitsExcluded

`func (o *CloneplanSummary) HasUnitsExcluded() bool`

HasUnitsExcluded returns a boolean if a field has been set.

### GetUnitsTotal

`func (o *CloneplanSummary) GetUnitsTotal() int32`

GetUnitsTotal returns the UnitsTotal field if non-nil, zero value otherwise.

### GetUnitsTotalOk

`func (o *CloneplanSummary) GetUnitsTotalOk() (*int32, bool)`

GetUnitsTotalOk returns a tuple with the UnitsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsTotal

`func (o *CloneplanSummary) SetUnitsTotal(v int32)`

SetUnitsTotal sets UnitsTotal field to given value.

### HasUnitsTotal

`func (o *CloneplanSummary) HasUnitsTotal() bool`

HasUnitsTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


