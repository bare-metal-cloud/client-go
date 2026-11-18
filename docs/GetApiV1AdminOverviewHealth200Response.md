# GetApiV1AdminOverviewHealth200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]HandlersOverviewServiceHealth**](HandlersOverviewServiceHealth.md) |  | [optional] 

## Methods

### NewGetApiV1AdminOverviewHealth200Response

`func NewGetApiV1AdminOverviewHealth200Response() *GetApiV1AdminOverviewHealth200Response`

NewGetApiV1AdminOverviewHealth200Response instantiates a new GetApiV1AdminOverviewHealth200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiV1AdminOverviewHealth200ResponseWithDefaults

`func NewGetApiV1AdminOverviewHealth200ResponseWithDefaults() *GetApiV1AdminOverviewHealth200Response`

NewGetApiV1AdminOverviewHealth200ResponseWithDefaults instantiates a new GetApiV1AdminOverviewHealth200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneratedAt

`func (o *GetApiV1AdminOverviewHealth200Response) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetApiV1AdminOverviewHealth200Response) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetApiV1AdminOverviewHealth200Response) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GetApiV1AdminOverviewHealth200Response) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetServices

`func (o *GetApiV1AdminOverviewHealth200Response) GetServices() []HandlersOverviewServiceHealth`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetApiV1AdminOverviewHealth200Response) GetServicesOk() (*[]HandlersOverviewServiceHealth, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetApiV1AdminOverviewHealth200Response) SetServices(v []HandlersOverviewServiceHealth)`

SetServices sets Services field to given value.

### HasServices

`func (o *GetApiV1AdminOverviewHealth200Response) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


