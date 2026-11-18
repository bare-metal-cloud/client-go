# HandlersElevationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecidedAt** | Pointer to **string** |  | [optional] 
**DecidedBy** | Pointer to **int32** |  | [optional] 
**DenyReason** | Pointer to **string** |  | [optional] 
**DurationMinutes** | Pointer to **int32** |  | [optional] 
**EndReason** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**RequestedAt** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**HandlersElevationUser**](HandlersElevationUser.md) |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersElevationView

`func NewHandlersElevationView() *HandlersElevationView`

NewHandlersElevationView instantiates a new HandlersElevationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersElevationViewWithDefaults

`func NewHandlersElevationViewWithDefaults() *HandlersElevationView`

NewHandlersElevationViewWithDefaults instantiates a new HandlersElevationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecidedAt

`func (o *HandlersElevationView) GetDecidedAt() string`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *HandlersElevationView) GetDecidedAtOk() (*string, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *HandlersElevationView) SetDecidedAt(v string)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *HandlersElevationView) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *HandlersElevationView) GetDecidedBy() int32`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *HandlersElevationView) GetDecidedByOk() (*int32, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *HandlersElevationView) SetDecidedBy(v int32)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *HandlersElevationView) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetDenyReason

`func (o *HandlersElevationView) GetDenyReason() string`

GetDenyReason returns the DenyReason field if non-nil, zero value otherwise.

### GetDenyReasonOk

`func (o *HandlersElevationView) GetDenyReasonOk() (*string, bool)`

GetDenyReasonOk returns a tuple with the DenyReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyReason

`func (o *HandlersElevationView) SetDenyReason(v string)`

SetDenyReason sets DenyReason field to given value.

### HasDenyReason

`func (o *HandlersElevationView) HasDenyReason() bool`

HasDenyReason returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *HandlersElevationView) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *HandlersElevationView) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *HandlersElevationView) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *HandlersElevationView) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetEndReason

`func (o *HandlersElevationView) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *HandlersElevationView) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *HandlersElevationView) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *HandlersElevationView) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### GetEndedAt

`func (o *HandlersElevationView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *HandlersElevationView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *HandlersElevationView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *HandlersElevationView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *HandlersElevationView) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HandlersElevationView) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HandlersElevationView) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *HandlersElevationView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *HandlersElevationView) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersElevationView) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersElevationView) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersElevationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *HandlersElevationView) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *HandlersElevationView) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *HandlersElevationView) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *HandlersElevationView) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetReason

`func (o *HandlersElevationView) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersElevationView) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersElevationView) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersElevationView) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequestedAt

`func (o *HandlersElevationView) GetRequestedAt() string`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *HandlersElevationView) GetRequestedAtOk() (*string, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *HandlersElevationView) SetRequestedAt(v string)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *HandlersElevationView) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartsAt

`func (o *HandlersElevationView) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *HandlersElevationView) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *HandlersElevationView) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *HandlersElevationView) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersElevationView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersElevationView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersElevationView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersElevationView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *HandlersElevationView) GetUser() HandlersElevationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HandlersElevationView) GetUserOk() (*HandlersElevationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HandlersElevationView) SetUser(v HandlersElevationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *HandlersElevationView) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *HandlersElevationView) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlersElevationView) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlersElevationView) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HandlersElevationView) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


