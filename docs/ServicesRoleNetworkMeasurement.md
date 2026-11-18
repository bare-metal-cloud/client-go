# ServicesRoleNetworkMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to **[][]int32** | Components groups mutually-reachable VM IDs; CriticalVMIDs are the members sitting on every path — the graph copy&#39;s data. | [optional] 
**CriticalVms** | Pointer to **[]int32** |  | [optional] 
**Expired** | Pointer to **bool** | Expired true: a record exists but its validity lapsed — the re-probe-offer state, distinct from never-measured. | [optional] 
**Measured** | Pointer to **bool** | Measured true: a usable measurement fed the suggestion. | [optional] 
**MeasuredAt** | Pointer to **string** | MeasuredAt is when the feeding probe ran (present only when measured). | [optional] 
**NotMeasuredReason** | Pointer to **string** | NotMeasuredReason is the honest why when Measured is false. | [optional] 
**Plane** | Pointer to **string** | Plane is the selection&#39;s operating plane (Ruling 4): direct for a networkless selection, mesh for a networked one. | [optional] 

## Methods

### NewServicesRoleNetworkMeasurement

`func NewServicesRoleNetworkMeasurement() *ServicesRoleNetworkMeasurement`

NewServicesRoleNetworkMeasurement instantiates a new ServicesRoleNetworkMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesRoleNetworkMeasurementWithDefaults

`func NewServicesRoleNetworkMeasurementWithDefaults() *ServicesRoleNetworkMeasurement`

NewServicesRoleNetworkMeasurementWithDefaults instantiates a new ServicesRoleNetworkMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ServicesRoleNetworkMeasurement) GetComponents() [][]int32`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ServicesRoleNetworkMeasurement) GetComponentsOk() (*[][]int32, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ServicesRoleNetworkMeasurement) SetComponents(v [][]int32)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ServicesRoleNetworkMeasurement) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCriticalVms

`func (o *ServicesRoleNetworkMeasurement) GetCriticalVms() []int32`

GetCriticalVms returns the CriticalVms field if non-nil, zero value otherwise.

### GetCriticalVmsOk

`func (o *ServicesRoleNetworkMeasurement) GetCriticalVmsOk() (*[]int32, bool)`

GetCriticalVmsOk returns a tuple with the CriticalVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalVms

`func (o *ServicesRoleNetworkMeasurement) SetCriticalVms(v []int32)`

SetCriticalVms sets CriticalVms field to given value.

### HasCriticalVms

`func (o *ServicesRoleNetworkMeasurement) HasCriticalVms() bool`

HasCriticalVms returns a boolean if a field has been set.

### GetExpired

`func (o *ServicesRoleNetworkMeasurement) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ServicesRoleNetworkMeasurement) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ServicesRoleNetworkMeasurement) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ServicesRoleNetworkMeasurement) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetMeasured

`func (o *ServicesRoleNetworkMeasurement) GetMeasured() bool`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *ServicesRoleNetworkMeasurement) GetMeasuredOk() (*bool, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *ServicesRoleNetworkMeasurement) SetMeasured(v bool)`

SetMeasured sets Measured field to given value.

### HasMeasured

`func (o *ServicesRoleNetworkMeasurement) HasMeasured() bool`

HasMeasured returns a boolean if a field has been set.

### GetMeasuredAt

`func (o *ServicesRoleNetworkMeasurement) GetMeasuredAt() string`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *ServicesRoleNetworkMeasurement) GetMeasuredAtOk() (*string, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *ServicesRoleNetworkMeasurement) SetMeasuredAt(v string)`

SetMeasuredAt sets MeasuredAt field to given value.

### HasMeasuredAt

`func (o *ServicesRoleNetworkMeasurement) HasMeasuredAt() bool`

HasMeasuredAt returns a boolean if a field has been set.

### GetNotMeasuredReason

`func (o *ServicesRoleNetworkMeasurement) GetNotMeasuredReason() string`

GetNotMeasuredReason returns the NotMeasuredReason field if non-nil, zero value otherwise.

### GetNotMeasuredReasonOk

`func (o *ServicesRoleNetworkMeasurement) GetNotMeasuredReasonOk() (*string, bool)`

GetNotMeasuredReasonOk returns a tuple with the NotMeasuredReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotMeasuredReason

`func (o *ServicesRoleNetworkMeasurement) SetNotMeasuredReason(v string)`

SetNotMeasuredReason sets NotMeasuredReason field to given value.

### HasNotMeasuredReason

`func (o *ServicesRoleNetworkMeasurement) HasNotMeasuredReason() bool`

HasNotMeasuredReason returns a boolean if a field has been set.

### GetPlane

`func (o *ServicesRoleNetworkMeasurement) GetPlane() string`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *ServicesRoleNetworkMeasurement) GetPlaneOk() (*string, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *ServicesRoleNetworkMeasurement) SetPlane(v string)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *ServicesRoleNetworkMeasurement) HasPlane() bool`

HasPlane returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


