# DatabasequeryQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Outcome** | Pointer to [**DatabasequeryOutcome**](DatabasequeryOutcome.md) |  | [optional] 
**RowCount** | Pointer to **int32** |  | [optional] 
**Rows** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**TruncatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabasequeryQueryResponse

`func NewDatabasequeryQueryResponse() *DatabasequeryQueryResponse`

NewDatabasequeryQueryResponse instantiates a new DatabasequeryQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabasequeryQueryResponseWithDefaults

`func NewDatabasequeryQueryResponseWithDefaults() *DatabasequeryQueryResponse`

NewDatabasequeryQueryResponseWithDefaults instantiates a new DatabasequeryQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *DatabasequeryQueryResponse) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DatabasequeryQueryResponse) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DatabasequeryQueryResponse) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DatabasequeryQueryResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDurationMs

`func (o *DatabasequeryQueryResponse) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *DatabasequeryQueryResponse) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *DatabasequeryQueryResponse) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *DatabasequeryQueryResponse) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetMessage

`func (o *DatabasequeryQueryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DatabasequeryQueryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DatabasequeryQueryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DatabasequeryQueryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutcome

`func (o *DatabasequeryQueryResponse) GetOutcome() DatabasequeryOutcome`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *DatabasequeryQueryResponse) GetOutcomeOk() (*DatabasequeryOutcome, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *DatabasequeryQueryResponse) SetOutcome(v DatabasequeryOutcome)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *DatabasequeryQueryResponse) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetRowCount

`func (o *DatabasequeryQueryResponse) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *DatabasequeryQueryResponse) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *DatabasequeryQueryResponse) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *DatabasequeryQueryResponse) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetRows

`func (o *DatabasequeryQueryResponse) GetRows() [][]map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DatabasequeryQueryResponse) GetRowsOk() (*[][]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DatabasequeryQueryResponse) SetRows(v [][]map[string]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *DatabasequeryQueryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTruncated

`func (o *DatabasequeryQueryResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *DatabasequeryQueryResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *DatabasequeryQueryResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *DatabasequeryQueryResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetTruncatedBy

`func (o *DatabasequeryQueryResponse) GetTruncatedBy() string`

GetTruncatedBy returns the TruncatedBy field if non-nil, zero value otherwise.

### GetTruncatedByOk

`func (o *DatabasequeryQueryResponse) GetTruncatedByOk() (*string, bool)`

GetTruncatedByOk returns a tuple with the TruncatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedBy

`func (o *DatabasequeryQueryResponse) SetTruncatedBy(v string)`

SetTruncatedBy sets TruncatedBy field to given value.

### HasTruncatedBy

`func (o *DatabasequeryQueryResponse) HasTruncatedBy() bool`

HasTruncatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


