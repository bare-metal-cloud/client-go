# ServicesClusterLatencyDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is the running-state gate: true only while the cluster is operational (status provisioned). False &#x3D; the panel renders honestly dormant (stale keys may briefly outlive a transition; they are never served as healthy). | [optional] 
**ClusterId** | Pointer to **int32** |  | [optional] 
**ClusterStatus** | Pointer to **string** | ClusterStatus is the cluster&#39;s CURRENT status verbatim. | [optional] 
**MeasuredAt** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]ServicesLatencyDashboardNode**](ServicesLatencyDashboardNode.md) |  | [optional] 
**Pairs** | Pointer to [**[]ServicesLatencyDashboardPair**](ServicesLatencyDashboardPair.md) |  | [optional] 
**Plane** | Pointer to **string** | Plane is the measured operating plane (direct | mesh); empty when no record exists. | [optional] 

## Methods

### NewServicesClusterLatencyDashboard

`func NewServicesClusterLatencyDashboard() *ServicesClusterLatencyDashboard`

NewServicesClusterLatencyDashboard instantiates a new ServicesClusterLatencyDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesClusterLatencyDashboardWithDefaults

`func NewServicesClusterLatencyDashboardWithDefaults() *ServicesClusterLatencyDashboard`

NewServicesClusterLatencyDashboardWithDefaults instantiates a new ServicesClusterLatencyDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ServicesClusterLatencyDashboard) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ServicesClusterLatencyDashboard) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ServicesClusterLatencyDashboard) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ServicesClusterLatencyDashboard) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetClusterId

`func (o *ServicesClusterLatencyDashboard) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ServicesClusterLatencyDashboard) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ServicesClusterLatencyDashboard) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ServicesClusterLatencyDashboard) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterStatus

`func (o *ServicesClusterLatencyDashboard) GetClusterStatus() string`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *ServicesClusterLatencyDashboard) GetClusterStatusOk() (*string, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *ServicesClusterLatencyDashboard) SetClusterStatus(v string)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *ServicesClusterLatencyDashboard) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetMeasuredAt

`func (o *ServicesClusterLatencyDashboard) GetMeasuredAt() string`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *ServicesClusterLatencyDashboard) GetMeasuredAtOk() (*string, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *ServicesClusterLatencyDashboard) SetMeasuredAt(v string)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *ServicesClusterLatencyDashboard) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetNodes

`func (o *ServicesClusterLatencyDashboard) GetNodes() []ServicesLatencyDashboardNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ServicesClusterLatencyDashboard) GetNodesOk() (*[]ServicesLatencyDashboardNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ServicesClusterLatencyDashboard) SetNodes(v []ServicesLatencyDashboardNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ServicesClusterLatencyDashboard) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPairs

`func (o *ServicesClusterLatencyDashboard) GetPairs() []ServicesLatencyDashboardPair`

GetPairs returns the Pairs field if non-nil, zero value otherwise.

### GetPairsOk

`func (o *ServicesClusterLatencyDashboard) GetPairsOk() (*[]ServicesLatencyDashboardPair, bool)`

GetPairsOk returns a tuple with the Pairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairs

`func (o *ServicesClusterLatencyDashboard) SetPairs(v []ServicesLatencyDashboardPair)`

SetPairs sets Pairs field to given value.

### HasPairs

`func (o *ServicesClusterLatencyDashboard) HasPairs() bool`

HasPairs returns a boolean if a field has been set.

### GetPlane

`func (o *ServicesClusterLatencyDashboard) GetPlane() string`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *ServicesClusterLatencyDashboard) GetPlaneOk() (*string, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *ServicesClusterLatencyDashboard) SetPlane(v string)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *ServicesClusterLatencyDashboard) HasPlane() bool`

HasPlane returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


