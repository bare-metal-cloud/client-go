# EdgeinsightsInsightPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cadence** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** | Family / TopK / Cadence echo a metric declaration&#39;s contract. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pattern** | Pointer to **string** | Pattern echoes a logs declaration. | [optional] 
**Presentation** | Pointer to **string** | Presentation echoes the spec&#39;s rendering hint (the per-kind default when the spec left it empty). | [optional] 
**Reason** | Pointer to **string** | Reason carries the degraded/unavailable cause. Empty otherwise. | [optional] 
**Rows** | Pointer to **[]string** | Rows carries the exec collector&#39;s stdout lines (the table fact). | [optional] 
**State** | Pointer to **string** | State is the honest rendering vocabulary (the constants above). | [optional] 
**TopK** | Pointer to **int32** |  | [optional] 

## Methods

### NewEdgeinsightsInsightPanel

`func NewEdgeinsightsInsightPanel() *EdgeinsightsInsightPanel`

NewEdgeinsightsInsightPanel instantiates a new EdgeinsightsInsightPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeinsightsInsightPanelWithDefaults

`func NewEdgeinsightsInsightPanelWithDefaults() *EdgeinsightsInsightPanel`

NewEdgeinsightsInsightPanelWithDefaults instantiates a new EdgeinsightsInsightPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCadence

`func (o *EdgeinsightsInsightPanel) GetCadence() string`

GetCadence returns the Cadence field if non-nil, zero value otherwise.

### GetCadenceOk

`func (o *EdgeinsightsInsightPanel) GetCadenceOk() (*string, bool)`

GetCadenceOk returns a tuple with the Cadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadence

`func (o *EdgeinsightsInsightPanel) SetCadence(v string)`

SetCadence sets Cadence field to given value.

### HasCadence

`func (o *EdgeinsightsInsightPanel) HasCadence() bool`

HasCadence returns a boolean if a field has been set.

### GetFamily

`func (o *EdgeinsightsInsightPanel) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *EdgeinsightsInsightPanel) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *EdgeinsightsInsightPanel) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *EdgeinsightsInsightPanel) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetKind

`func (o *EdgeinsightsInsightPanel) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EdgeinsightsInsightPanel) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EdgeinsightsInsightPanel) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EdgeinsightsInsightPanel) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *EdgeinsightsInsightPanel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeinsightsInsightPanel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeinsightsInsightPanel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EdgeinsightsInsightPanel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPattern

`func (o *EdgeinsightsInsightPanel) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *EdgeinsightsInsightPanel) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *EdgeinsightsInsightPanel) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *EdgeinsightsInsightPanel) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPresentation

`func (o *EdgeinsightsInsightPanel) GetPresentation() string`

GetPresentation returns the Presentation field if non-nil, zero value otherwise.

### GetPresentationOk

`func (o *EdgeinsightsInsightPanel) GetPresentationOk() (*string, bool)`

GetPresentationOk returns a tuple with the Presentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentation

`func (o *EdgeinsightsInsightPanel) SetPresentation(v string)`

SetPresentation sets Presentation field to given value.

### HasPresentation

`func (o *EdgeinsightsInsightPanel) HasPresentation() bool`

HasPresentation returns a boolean if a field has been set.

### GetReason

`func (o *EdgeinsightsInsightPanel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EdgeinsightsInsightPanel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EdgeinsightsInsightPanel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EdgeinsightsInsightPanel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRows

`func (o *EdgeinsightsInsightPanel) GetRows() []string`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EdgeinsightsInsightPanel) GetRowsOk() (*[]string, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EdgeinsightsInsightPanel) SetRows(v []string)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EdgeinsightsInsightPanel) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetState

`func (o *EdgeinsightsInsightPanel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EdgeinsightsInsightPanel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EdgeinsightsInsightPanel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EdgeinsightsInsightPanel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTopK

`func (o *EdgeinsightsInsightPanel) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *EdgeinsightsInsightPanel) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *EdgeinsightsInsightPanel) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *EdgeinsightsInsightPanel) HasTopK() bool`

HasTopK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


