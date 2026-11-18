# GetApiV1AdminOverviewFleet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentsConnected** | Pointer to **int32** |  | [optional] 
**AgentsConnectedError** | Pointer to **string** |  | [optional] 
**AgentsRegistered** | Pointer to **int32** |  | [optional] 
**Clusters** | Pointer to **int32** |  | [optional] 
**Organizations** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to **int32** |  | [optional] 
**UsersMfaEnrolled** | Pointer to **int32** |  | [optional] 
**VirtualMachines** | Pointer to **int32** |  | [optional] 
**VirtualMachinesUnreachable** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetApiV1AdminOverviewFleet200Response

`func NewGetApiV1AdminOverviewFleet200Response() *GetApiV1AdminOverviewFleet200Response`

NewGetApiV1AdminOverviewFleet200Response instantiates a new GetApiV1AdminOverviewFleet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiV1AdminOverviewFleet200ResponseWithDefaults

`func NewGetApiV1AdminOverviewFleet200ResponseWithDefaults() *GetApiV1AdminOverviewFleet200Response`

NewGetApiV1AdminOverviewFleet200ResponseWithDefaults instantiates a new GetApiV1AdminOverviewFleet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentsConnected

`func (o *GetApiV1AdminOverviewFleet200Response) GetAgentsConnected() int32`

GetAgentsConnected returns the AgentsConnected field if non-nil, zero value otherwise.

### GetAgentsConnectedOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetAgentsConnectedOk() (*int32, bool)`

GetAgentsConnectedOk returns a tuple with the AgentsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsConnected

`func (o *GetApiV1AdminOverviewFleet200Response) SetAgentsConnected(v int32)`

SetAgentsConnected sets AgentsConnected field to given value.

### HasAgentsConnected

`func (o *GetApiV1AdminOverviewFleet200Response) HasAgentsConnected() bool`

HasAgentsConnected returns a boolean if a field has been set.

### GetAgentsConnectedError

`func (o *GetApiV1AdminOverviewFleet200Response) GetAgentsConnectedError() string`

GetAgentsConnectedError returns the AgentsConnectedError field if non-nil, zero value otherwise.

### GetAgentsConnectedErrorOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetAgentsConnectedErrorOk() (*string, bool)`

GetAgentsConnectedErrorOk returns a tuple with the AgentsConnectedError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsConnectedError

`func (o *GetApiV1AdminOverviewFleet200Response) SetAgentsConnectedError(v string)`

SetAgentsConnectedError sets AgentsConnectedError field to given value.

### HasAgentsConnectedError

`func (o *GetApiV1AdminOverviewFleet200Response) HasAgentsConnectedError() bool`

HasAgentsConnectedError returns a boolean if a field has been set.

### GetAgentsRegistered

`func (o *GetApiV1AdminOverviewFleet200Response) GetAgentsRegistered() int32`

GetAgentsRegistered returns the AgentsRegistered field if non-nil, zero value otherwise.

### GetAgentsRegisteredOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetAgentsRegisteredOk() (*int32, bool)`

GetAgentsRegisteredOk returns a tuple with the AgentsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsRegistered

`func (o *GetApiV1AdminOverviewFleet200Response) SetAgentsRegistered(v int32)`

SetAgentsRegistered sets AgentsRegistered field to given value.

### HasAgentsRegistered

`func (o *GetApiV1AdminOverviewFleet200Response) HasAgentsRegistered() bool`

HasAgentsRegistered returns a boolean if a field has been set.

### GetClusters

`func (o *GetApiV1AdminOverviewFleet200Response) GetClusters() int32`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetClustersOk() (*int32, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GetApiV1AdminOverviewFleet200Response) SetClusters(v int32)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *GetApiV1AdminOverviewFleet200Response) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetOrganizations

`func (o *GetApiV1AdminOverviewFleet200Response) GetOrganizations() int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetOrganizationsOk() (*int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetApiV1AdminOverviewFleet200Response) SetOrganizations(v int32)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *GetApiV1AdminOverviewFleet200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetUsers

`func (o *GetApiV1AdminOverviewFleet200Response) GetUsers() int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetUsersOk() (*int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetApiV1AdminOverviewFleet200Response) SetUsers(v int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetApiV1AdminOverviewFleet200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUsersMfaEnrolled

`func (o *GetApiV1AdminOverviewFleet200Response) GetUsersMfaEnrolled() int32`

GetUsersMfaEnrolled returns the UsersMfaEnrolled field if non-nil, zero value otherwise.

### GetUsersMfaEnrolledOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetUsersMfaEnrolledOk() (*int32, bool)`

GetUsersMfaEnrolledOk returns a tuple with the UsersMfaEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersMfaEnrolled

`func (o *GetApiV1AdminOverviewFleet200Response) SetUsersMfaEnrolled(v int32)`

SetUsersMfaEnrolled sets UsersMfaEnrolled field to given value.

### HasUsersMfaEnrolled

`func (o *GetApiV1AdminOverviewFleet200Response) HasUsersMfaEnrolled() bool`

HasUsersMfaEnrolled returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *GetApiV1AdminOverviewFleet200Response) GetVirtualMachines() int32`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetVirtualMachinesOk() (*int32, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *GetApiV1AdminOverviewFleet200Response) SetVirtualMachines(v int32)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *GetApiV1AdminOverviewFleet200Response) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.

### GetVirtualMachinesUnreachable

`func (o *GetApiV1AdminOverviewFleet200Response) GetVirtualMachinesUnreachable() int32`

GetVirtualMachinesUnreachable returns the VirtualMachinesUnreachable field if non-nil, zero value otherwise.

### GetVirtualMachinesUnreachableOk

`func (o *GetApiV1AdminOverviewFleet200Response) GetVirtualMachinesUnreachableOk() (*int32, bool)`

GetVirtualMachinesUnreachableOk returns a tuple with the VirtualMachinesUnreachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachinesUnreachable

`func (o *GetApiV1AdminOverviewFleet200Response) SetVirtualMachinesUnreachable(v int32)`

SetVirtualMachinesUnreachable sets VirtualMachinesUnreachable field to given value.

### HasVirtualMachinesUnreachable

`func (o *GetApiV1AdminOverviewFleet200Response) HasVirtualMachinesUnreachable() bool`

HasVirtualMachinesUnreachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


