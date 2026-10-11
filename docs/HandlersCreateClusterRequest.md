# HandlersCreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**HandlersCreateClusterRequestData**](HandlersCreateClusterRequestData.md) |  | 
**Name** | **string** |  | 
**Provider** | **string** |  | 

## Methods

### NewHandlersCreateClusterRequest

`func NewHandlersCreateClusterRequest(data HandlersCreateClusterRequestData, name string, provider string, ) *HandlersCreateClusterRequest`

NewHandlersCreateClusterRequest instantiates a new HandlersCreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateClusterRequestWithDefaults

`func NewHandlersCreateClusterRequestWithDefaults() *HandlersCreateClusterRequest`

NewHandlersCreateClusterRequestWithDefaults instantiates a new HandlersCreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HandlersCreateClusterRequest) GetData() HandlersCreateClusterRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HandlersCreateClusterRequest) GetDataOk() (*HandlersCreateClusterRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HandlersCreateClusterRequest) SetData(v HandlersCreateClusterRequestData)`

SetData sets Data field to given value.


### GetName

`func (o *HandlersCreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *HandlersCreateClusterRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersCreateClusterRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersCreateClusterRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


