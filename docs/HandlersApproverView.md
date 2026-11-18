# HandlersApproverView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to **int32** |  | [optional] 
**User** | Pointer to [**HandlersElevationUser**](HandlersElevationUser.md) |  | [optional] 

## Methods

### NewHandlersApproverView

`func NewHandlersApproverView() *HandlersApproverView`

NewHandlersApproverView instantiates a new HandlersApproverView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersApproverViewWithDefaults

`func NewHandlersApproverViewWithDefaults() *HandlersApproverView`

NewHandlersApproverViewWithDefaults instantiates a new HandlersApproverView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *HandlersApproverView) GetAddedAt() string`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *HandlersApproverView) GetAddedAtOk() (*string, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *HandlersApproverView) SetAddedAt(v string)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *HandlersApproverView) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetAddedBy

`func (o *HandlersApproverView) GetAddedBy() int32`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *HandlersApproverView) GetAddedByOk() (*int32, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *HandlersApproverView) SetAddedBy(v int32)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *HandlersApproverView) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetUser

`func (o *HandlersApproverView) GetUser() HandlersElevationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HandlersApproverView) GetUserOk() (*HandlersElevationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HandlersApproverView) SetUser(v HandlersElevationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *HandlersApproverView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


