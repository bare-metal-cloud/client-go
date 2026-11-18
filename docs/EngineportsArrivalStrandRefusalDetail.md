# EngineportsArrivalStrandRefusalDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ClusterID is the engine&#39;s cluster ref (the numeric cluster row id in string form — the re-target preview&#39;s cluster_id parameter). | [optional] 
**DeadTarget** | Pointer to **string** | DeadTarget is the plan&#39;s stamped prior target — the dead tunnel CNAME the stranded records still point at (the condition the re-target clears). | [optional] 
**Instance** | Pointer to **string** | Instance is the arrival addon instance id (the preview&#39;s instance parameter). | [optional] 
**LiveTarget** | Pointer to **string** | LiveTarget is the recreated tunnel&#39;s current target (the re-target&#39;s destination), best-effort at refusal time — empty when the live read failed (the refusal&#39;s own message still names the tunnel). | [optional] 

## Methods

### NewEngineportsArrivalStrandRefusalDetail

`func NewEngineportsArrivalStrandRefusalDetail() *EngineportsArrivalStrandRefusalDetail`

NewEngineportsArrivalStrandRefusalDetail instantiates a new EngineportsArrivalStrandRefusalDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsArrivalStrandRefusalDetailWithDefaults

`func NewEngineportsArrivalStrandRefusalDetailWithDefaults() *EngineportsArrivalStrandRefusalDetail`

NewEngineportsArrivalStrandRefusalDetailWithDefaults instantiates a new EngineportsArrivalStrandRefusalDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *EngineportsArrivalStrandRefusalDetail) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EngineportsArrivalStrandRefusalDetail) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EngineportsArrivalStrandRefusalDetail) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EngineportsArrivalStrandRefusalDetail) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetDeadTarget

`func (o *EngineportsArrivalStrandRefusalDetail) GetDeadTarget() string`

GetDeadTarget returns the DeadTarget field if non-nil, zero value otherwise.

### GetDeadTargetOk

`func (o *EngineportsArrivalStrandRefusalDetail) GetDeadTargetOk() (*string, bool)`

GetDeadTargetOk returns a tuple with the DeadTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTarget

`func (o *EngineportsArrivalStrandRefusalDetail) SetDeadTarget(v string)`

SetDeadTarget sets DeadTarget field to given value.

### HasDeadTarget

`func (o *EngineportsArrivalStrandRefusalDetail) HasDeadTarget() bool`

HasDeadTarget returns a boolean if a field has been set.

### GetInstance

`func (o *EngineportsArrivalStrandRefusalDetail) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *EngineportsArrivalStrandRefusalDetail) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *EngineportsArrivalStrandRefusalDetail) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *EngineportsArrivalStrandRefusalDetail) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetLiveTarget

`func (o *EngineportsArrivalStrandRefusalDetail) GetLiveTarget() string`

GetLiveTarget returns the LiveTarget field if non-nil, zero value otherwise.

### GetLiveTargetOk

`func (o *EngineportsArrivalStrandRefusalDetail) GetLiveTargetOk() (*string, bool)`

GetLiveTargetOk returns a tuple with the LiveTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveTarget

`func (o *EngineportsArrivalStrandRefusalDetail) SetLiveTarget(v string)`

SetLiveTarget sets LiveTarget field to given value.

### HasLiveTarget

`func (o *EngineportsArrivalStrandRefusalDetail) HasLiveTarget() bool`

HasLiveTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


