# HandlersInstallChartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | Pointer to **string** | e.g., \&quot;prometheus-community/prometheus\&quot; | [optional] 
**Namespace** | Pointer to **string** | e.g., \&quot;monitoring\&quot; | [optional] 
**Release** | Pointer to **string** | e.g., \&quot;prometheus-prod\&quot; | [optional] 
**Values** | Pointer to **map[string]interface{}** | Optional Helm values override | [optional] 

## Methods

### NewHandlersInstallChartRequest

`func NewHandlersInstallChartRequest() *HandlersInstallChartRequest`

NewHandlersInstallChartRequest instantiates a new HandlersInstallChartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersInstallChartRequestWithDefaults

`func NewHandlersInstallChartRequestWithDefaults() *HandlersInstallChartRequest`

NewHandlersInstallChartRequestWithDefaults instantiates a new HandlersInstallChartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChart

`func (o *HandlersInstallChartRequest) GetChart() string`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *HandlersInstallChartRequest) GetChartOk() (*string, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *HandlersInstallChartRequest) SetChart(v string)`

SetChart sets Chart field to given value.

### HasChart

`func (o *HandlersInstallChartRequest) HasChart() bool`

HasChart returns a boolean if a field has been set.

### GetNamespace

`func (o *HandlersInstallChartRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HandlersInstallChartRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HandlersInstallChartRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HandlersInstallChartRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRelease

`func (o *HandlersInstallChartRequest) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *HandlersInstallChartRequest) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *HandlersInstallChartRequest) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *HandlersInstallChartRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetValues

`func (o *HandlersInstallChartRequest) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HandlersInstallChartRequest) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HandlersInstallChartRequest) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *HandlersInstallChartRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


