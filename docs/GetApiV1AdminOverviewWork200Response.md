# GetApiV1AdminOverviewWork200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asynq** | Pointer to **map[string]interface{}** |  | [optional] 
**Deadman** | Pointer to [**HandlersOverviewDeadmanBoard**](HandlersOverviewDeadmanBoard.md) |  | [optional] 
**Temporal** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetApiV1AdminOverviewWork200Response

`func NewGetApiV1AdminOverviewWork200Response() *GetApiV1AdminOverviewWork200Response`

NewGetApiV1AdminOverviewWork200Response instantiates a new GetApiV1AdminOverviewWork200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiV1AdminOverviewWork200ResponseWithDefaults

`func NewGetApiV1AdminOverviewWork200ResponseWithDefaults() *GetApiV1AdminOverviewWork200Response`

NewGetApiV1AdminOverviewWork200ResponseWithDefaults instantiates a new GetApiV1AdminOverviewWork200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsynq

`func (o *GetApiV1AdminOverviewWork200Response) GetAsynq() map[string]interface{}`

GetAsynq returns the Asynq field if non-nil, zero value otherwise.

### GetAsynqOk

`func (o *GetApiV1AdminOverviewWork200Response) GetAsynqOk() (*map[string]interface{}, bool)`

GetAsynqOk returns a tuple with the Asynq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynq

`func (o *GetApiV1AdminOverviewWork200Response) SetAsynq(v map[string]interface{})`

SetAsynq sets Asynq field to given value.

### HasAsynq

`func (o *GetApiV1AdminOverviewWork200Response) HasAsynq() bool`

HasAsynq returns a boolean if a field has been set.

### GetDeadman

`func (o *GetApiV1AdminOverviewWork200Response) GetDeadman() HandlersOverviewDeadmanBoard`

GetDeadman returns the Deadman field if non-nil, zero value otherwise.

### GetDeadmanOk

`func (o *GetApiV1AdminOverviewWork200Response) GetDeadmanOk() (*HandlersOverviewDeadmanBoard, bool)`

GetDeadmanOk returns a tuple with the Deadman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadman

`func (o *GetApiV1AdminOverviewWork200Response) SetDeadman(v HandlersOverviewDeadmanBoard)`

SetDeadman sets Deadman field to given value.

### HasDeadman

`func (o *GetApiV1AdminOverviewWork200Response) HasDeadman() bool`

HasDeadman returns a boolean if a field has been set.

### GetTemporal

`func (o *GetApiV1AdminOverviewWork200Response) GetTemporal() map[string]interface{}`

GetTemporal returns the Temporal field if non-nil, zero value otherwise.

### GetTemporalOk

`func (o *GetApiV1AdminOverviewWork200Response) GetTemporalOk() (*map[string]interface{}, bool)`

GetTemporalOk returns a tuple with the Temporal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporal

`func (o *GetApiV1AdminOverviewWork200Response) SetTemporal(v map[string]interface{})`

SetTemporal sets Temporal field to given value.

### HasTemporal

`func (o *GetApiV1AdminOverviewWork200Response) HasTemporal() bool`

HasTemporal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


