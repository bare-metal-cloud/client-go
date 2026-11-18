# HandlersSecretShareView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**EndReason** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**EndedByUserId** | Pointer to **int32** |  | [optional] 
**ExpiresAtUnix** | Pointer to **int32** |  | [optional] 
**FirstOpenedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**OpenCount** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**RecipientUserId** | Pointer to **int32** |  | [optional] 
**SecretId** | Pointer to **int32** |  | [optional] 
**SecretType** | Pointer to **string** |  | [optional] 
**SharerUserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersSecretShareView

`func NewHandlersSecretShareView() *HandlersSecretShareView`

NewHandlersSecretShareView instantiates a new HandlersSecretShareView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersSecretShareViewWithDefaults

`func NewHandlersSecretShareViewWithDefaults() *HandlersSecretShareView`

NewHandlersSecretShareViewWithDefaults instantiates a new HandlersSecretShareView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HandlersSecretShareView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersSecretShareView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersSecretShareView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersSecretShareView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndReason

`func (o *HandlersSecretShareView) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *HandlersSecretShareView) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *HandlersSecretShareView) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *HandlersSecretShareView) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### GetEndedAt

`func (o *HandlersSecretShareView) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *HandlersSecretShareView) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *HandlersSecretShareView) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *HandlersSecretShareView) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEndedByUserId

`func (o *HandlersSecretShareView) GetEndedByUserId() int32`

GetEndedByUserId returns the EndedByUserId field if non-nil, zero value otherwise.

### GetEndedByUserIdOk

`func (o *HandlersSecretShareView) GetEndedByUserIdOk() (*int32, bool)`

GetEndedByUserIdOk returns a tuple with the EndedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedByUserId

`func (o *HandlersSecretShareView) SetEndedByUserId(v int32)`

SetEndedByUserId sets EndedByUserId field to given value.

### HasEndedByUserId

`func (o *HandlersSecretShareView) HasEndedByUserId() bool`

HasEndedByUserId returns a boolean if a field has been set.

### GetExpiresAtUnix

`func (o *HandlersSecretShareView) GetExpiresAtUnix() int32`

GetExpiresAtUnix returns the ExpiresAtUnix field if non-nil, zero value otherwise.

### GetExpiresAtUnixOk

`func (o *HandlersSecretShareView) GetExpiresAtUnixOk() (*int32, bool)`

GetExpiresAtUnixOk returns a tuple with the ExpiresAtUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUnix

`func (o *HandlersSecretShareView) SetExpiresAtUnix(v int32)`

SetExpiresAtUnix sets ExpiresAtUnix field to given value.

### HasExpiresAtUnix

`func (o *HandlersSecretShareView) HasExpiresAtUnix() bool`

HasExpiresAtUnix returns a boolean if a field has been set.

### GetFirstOpenedAt

`func (o *HandlersSecretShareView) GetFirstOpenedAt() string`

GetFirstOpenedAt returns the FirstOpenedAt field if non-nil, zero value otherwise.

### GetFirstOpenedAtOk

`func (o *HandlersSecretShareView) GetFirstOpenedAtOk() (*string, bool)`

GetFirstOpenedAtOk returns a tuple with the FirstOpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOpenedAt

`func (o *HandlersSecretShareView) SetFirstOpenedAt(v string)`

SetFirstOpenedAt sets FirstOpenedAt field to given value.

### HasFirstOpenedAt

`func (o *HandlersSecretShareView) HasFirstOpenedAt() bool`

HasFirstOpenedAt returns a boolean if a field has been set.

### GetId

`func (o *HandlersSecretShareView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersSecretShareView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersSecretShareView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersSecretShareView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpenCount

`func (o *HandlersSecretShareView) GetOpenCount() int32`

GetOpenCount returns the OpenCount field if non-nil, zero value otherwise.

### GetOpenCountOk

`func (o *HandlersSecretShareView) GetOpenCountOk() (*int32, bool)`

GetOpenCountOk returns a tuple with the OpenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenCount

`func (o *HandlersSecretShareView) SetOpenCount(v int32)`

SetOpenCount sets OpenCount field to given value.

### HasOpenCount

`func (o *HandlersSecretShareView) HasOpenCount() bool`

HasOpenCount returns a boolean if a field has been set.

### GetOrganizationId

`func (o *HandlersSecretShareView) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *HandlersSecretShareView) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *HandlersSecretShareView) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *HandlersSecretShareView) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *HandlersSecretShareView) GetRecipientUserId() int32`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *HandlersSecretShareView) GetRecipientUserIdOk() (*int32, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *HandlersSecretShareView) SetRecipientUserId(v int32)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *HandlersSecretShareView) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetSecretId

`func (o *HandlersSecretShareView) GetSecretId() int32`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *HandlersSecretShareView) GetSecretIdOk() (*int32, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *HandlersSecretShareView) SetSecretId(v int32)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *HandlersSecretShareView) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetSecretType

`func (o *HandlersSecretShareView) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *HandlersSecretShareView) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *HandlersSecretShareView) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *HandlersSecretShareView) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSharerUserId

`func (o *HandlersSecretShareView) GetSharerUserId() int32`

GetSharerUserId returns the SharerUserId field if non-nil, zero value otherwise.

### GetSharerUserIdOk

`func (o *HandlersSecretShareView) GetSharerUserIdOk() (*int32, bool)`

GetSharerUserIdOk returns a tuple with the SharerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharerUserId

`func (o *HandlersSecretShareView) SetSharerUserId(v int32)`

SetSharerUserId sets SharerUserId field to given value.

### HasSharerUserId

`func (o *HandlersSecretShareView) HasSharerUserId() bool`

HasSharerUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


