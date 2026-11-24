# HandlersExplorerAccessView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanRequest** | Pointer to **bool** | CanRequest is FALSE at launch: the in-product request door is part of the strip (backend#4735 item 1), so the payload never advertises a door it does not open. The field stays in the shape — the v1 console contract reads it, and false keeps the request affordance hidden with no FE change; the field&#39;s removal rides the FE strip card. | [optional] 
**Cluster** | Pointer to [**HandlersExplorerAccessClusterRef**](HandlersExplorerAccessClusterRef.md) |  | [optional] 
**Grant** | Pointer to [**HandlersExplorerGrantView**](HandlersExplorerGrantView.md) | Grant is the ledger row that decided the state (any status) — the expiry clock the boundary renders from. | [optional] 
**PendingRequest** | Pointer to [**HandlersExplorerGrantView**](HandlersExplorerGrantView.md) | PendingRequest is set when a pre-launch request is still queued — \&quot;how to get it\&quot; cites it (the approver tier still decides it). | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State: granted | requested | denied. | [optional] 

## Methods

### NewHandlersExplorerAccessView

`func NewHandlersExplorerAccessView() *HandlersExplorerAccessView`

NewHandlersExplorerAccessView instantiates a new HandlersExplorerAccessView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersExplorerAccessViewWithDefaults

`func NewHandlersExplorerAccessViewWithDefaults() *HandlersExplorerAccessView`

NewHandlersExplorerAccessViewWithDefaults instantiates a new HandlersExplorerAccessView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanRequest

`func (o *HandlersExplorerAccessView) GetCanRequest() bool`

GetCanRequest returns the CanRequest field if non-nil, zero value otherwise.

### GetCanRequestOk

`func (o *HandlersExplorerAccessView) GetCanRequestOk() (*bool, bool)`

GetCanRequestOk returns a tuple with the CanRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRequest

`func (o *HandlersExplorerAccessView) SetCanRequest(v bool)`

SetCanRequest sets CanRequest field to given value.

### HasCanRequest

`func (o *HandlersExplorerAccessView) HasCanRequest() bool`

HasCanRequest returns a boolean if a field has been set.

### GetCluster

`func (o *HandlersExplorerAccessView) GetCluster() HandlersExplorerAccessClusterRef`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HandlersExplorerAccessView) GetClusterOk() (*HandlersExplorerAccessClusterRef, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HandlersExplorerAccessView) SetCluster(v HandlersExplorerAccessClusterRef)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HandlersExplorerAccessView) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetGrant

`func (o *HandlersExplorerAccessView) GetGrant() HandlersExplorerGrantView`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *HandlersExplorerAccessView) GetGrantOk() (*HandlersExplorerGrantView, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *HandlersExplorerAccessView) SetGrant(v HandlersExplorerGrantView)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *HandlersExplorerAccessView) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetPendingRequest

`func (o *HandlersExplorerAccessView) GetPendingRequest() HandlersExplorerGrantView`

GetPendingRequest returns the PendingRequest field if non-nil, zero value otherwise.

### GetPendingRequestOk

`func (o *HandlersExplorerAccessView) GetPendingRequestOk() (*HandlersExplorerGrantView, bool)`

GetPendingRequestOk returns a tuple with the PendingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRequest

`func (o *HandlersExplorerAccessView) SetPendingRequest(v HandlersExplorerGrantView)`

SetPendingRequest sets PendingRequest field to given value.

### HasPendingRequest

`func (o *HandlersExplorerAccessView) HasPendingRequest() bool`

HasPendingRequest returns a boolean if a field has been set.

### GetReason

`func (o *HandlersExplorerAccessView) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersExplorerAccessView) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersExplorerAccessView) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersExplorerAccessView) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *HandlersExplorerAccessView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersExplorerAccessView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersExplorerAccessView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersExplorerAccessView) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


