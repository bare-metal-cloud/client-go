# RunnerpoolReservationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Draining** | Pointer to **bool** |  | [optional] 
**GrantedAt** | Pointer to **string** |  | [optional] 
**Present** | Pointer to **bool** |  | [optional] 
**Request** | Pointer to **string** | Request is the request affordance (see ReservationRequestAffordance): how an org asks for a reservation — an admin conversation, never a purchase flow. | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**VmIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRunnerpoolReservationState

`func NewRunnerpoolReservationState() *RunnerpoolReservationState`

NewRunnerpoolReservationState instantiates a new RunnerpoolReservationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerpoolReservationStateWithDefaults

`func NewRunnerpoolReservationStateWithDefaults() *RunnerpoolReservationState`

NewRunnerpoolReservationStateWithDefaults instantiates a new RunnerpoolReservationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraining

`func (o *RunnerpoolReservationState) GetDraining() bool`

GetDraining returns the Draining field if non-nil, zero value otherwise.

### GetDrainingOk

`func (o *RunnerpoolReservationState) GetDrainingOk() (*bool, bool)`

GetDrainingOk returns a tuple with the Draining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraining

`func (o *RunnerpoolReservationState) SetDraining(v bool)`

SetDraining sets Draining field to given value.

### HasDraining

`func (o *RunnerpoolReservationState) HasDraining() bool`

HasDraining returns a boolean if a field has been set.

### GetGrantedAt

`func (o *RunnerpoolReservationState) GetGrantedAt() string`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *RunnerpoolReservationState) GetGrantedAtOk() (*string, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *RunnerpoolReservationState) SetGrantedAt(v string)`

SetGrantedAt sets GrantedAt field to given value.

### HasGrantedAt

`func (o *RunnerpoolReservationState) HasGrantedAt() bool`

HasGrantedAt returns a boolean if a field has been set.

### GetPresent

`func (o *RunnerpoolReservationState) GetPresent() bool`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *RunnerpoolReservationState) GetPresentOk() (*bool, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *RunnerpoolReservationState) SetPresent(v bool)`

SetPresent sets Present field to given value.

### HasPresent

`func (o *RunnerpoolReservationState) HasPresent() bool`

HasPresent returns a boolean if a field has been set.

### GetRequest

`func (o *RunnerpoolReservationState) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *RunnerpoolReservationState) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *RunnerpoolReservationState) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *RunnerpoolReservationState) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSize

`func (o *RunnerpoolReservationState) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RunnerpoolReservationState) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RunnerpoolReservationState) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RunnerpoolReservationState) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVmIds

`func (o *RunnerpoolReservationState) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *RunnerpoolReservationState) GetVmIdsOk() (*[]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *RunnerpoolReservationState) SetVmIds(v []string)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *RunnerpoolReservationState) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


