# HandlersOrgVaultDeliveriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**Deliveries** | Pointer to [**[]HandlersLedgerDelivery**](HandlersLedgerDelivery.md) |  | [optional] 
**Truncated** | Pointer to **bool** | Truncated is true when more rows existed inside the window than the read&#39;s row cap served — the card renders a \&quot;showing the newest N\&quot; note instead of truncating silently (review should-fix 4). | [optional] 

## Methods

### NewHandlersOrgVaultDeliveriesResponse

`func NewHandlersOrgVaultDeliveriesResponse() *HandlersOrgVaultDeliveriesResponse`

NewHandlersOrgVaultDeliveriesResponse instantiates a new HandlersOrgVaultDeliveriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultDeliveriesResponseWithDefaults

`func NewHandlersOrgVaultDeliveriesResponseWithDefaults() *HandlersOrgVaultDeliveriesResponse`

NewHandlersOrgVaultDeliveriesResponseWithDefaults instantiates a new HandlersOrgVaultDeliveriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HandlersOrgVaultDeliveriesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HandlersOrgVaultDeliveriesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HandlersOrgVaultDeliveriesResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HandlersOrgVaultDeliveriesResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDays

`func (o *HandlersOrgVaultDeliveriesResponse) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *HandlersOrgVaultDeliveriesResponse) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *HandlersOrgVaultDeliveriesResponse) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *HandlersOrgVaultDeliveriesResponse) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDeliveries

`func (o *HandlersOrgVaultDeliveriesResponse) GetDeliveries() []HandlersLedgerDelivery`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *HandlersOrgVaultDeliveriesResponse) GetDeliveriesOk() (*[]HandlersLedgerDelivery, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *HandlersOrgVaultDeliveriesResponse) SetDeliveries(v []HandlersLedgerDelivery)`

SetDeliveries sets Deliveries field to given value.

### HasDeliveries

`func (o *HandlersOrgVaultDeliveriesResponse) HasDeliveries() bool`

HasDeliveries returns a boolean if a field has been set.

### GetTruncated

`func (o *HandlersOrgVaultDeliveriesResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *HandlersOrgVaultDeliveriesResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *HandlersOrgVaultDeliveriesResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *HandlersOrgVaultDeliveriesResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


