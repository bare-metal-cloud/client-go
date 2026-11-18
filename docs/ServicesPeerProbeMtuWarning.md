# ServicesPeerProbeMtuWarning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail names the probed size so the warning is actionable. | [optional] 
**Kind** | Pointer to **string** | Kind is ProbePairKindMtuDegraded (mtu_degraded). | [optional] 
**PacketSize** | Pointer to **int32** |  | [optional] 
**PayloadSize** | Pointer to **int32** | PayloadSize / PacketSize are the measured crossing sizes — the verify_tailnet blocked record&#39;s exact shape (both 0: no DF packet of the probed size crossed; small packets only). | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**SourceVmId** | Pointer to **int32** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**TargetVmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesPeerProbeMtuWarning

`func NewServicesPeerProbeMtuWarning() *ServicesPeerProbeMtuWarning`

NewServicesPeerProbeMtuWarning instantiates a new ServicesPeerProbeMtuWarning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesPeerProbeMtuWarningWithDefaults

`func NewServicesPeerProbeMtuWarningWithDefaults() *ServicesPeerProbeMtuWarning`

NewServicesPeerProbeMtuWarningWithDefaults instantiates a new ServicesPeerProbeMtuWarning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ServicesPeerProbeMtuWarning) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ServicesPeerProbeMtuWarning) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ServicesPeerProbeMtuWarning) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ServicesPeerProbeMtuWarning) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetKind

`func (o *ServicesPeerProbeMtuWarning) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServicesPeerProbeMtuWarning) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServicesPeerProbeMtuWarning) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ServicesPeerProbeMtuWarning) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPacketSize

`func (o *ServicesPeerProbeMtuWarning) GetPacketSize() int32`

GetPacketSize returns the PacketSize field if non-nil, zero value otherwise.

### GetPacketSizeOk

`func (o *ServicesPeerProbeMtuWarning) GetPacketSizeOk() (*int32, bool)`

GetPacketSizeOk returns a tuple with the PacketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketSize

`func (o *ServicesPeerProbeMtuWarning) SetPacketSize(v int32)`

SetPacketSize sets PacketSize field to given value.

### HasPacketSize

`func (o *ServicesPeerProbeMtuWarning) HasPacketSize() bool`

HasPacketSize returns a boolean if a field has been set.

### GetPayloadSize

`func (o *ServicesPeerProbeMtuWarning) GetPayloadSize() int32`

GetPayloadSize returns the PayloadSize field if non-nil, zero value otherwise.

### GetPayloadSizeOk

`func (o *ServicesPeerProbeMtuWarning) GetPayloadSizeOk() (*int32, bool)`

GetPayloadSizeOk returns a tuple with the PayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadSize

`func (o *ServicesPeerProbeMtuWarning) SetPayloadSize(v int32)`

SetPayloadSize sets PayloadSize field to given value.

### HasPayloadSize

`func (o *ServicesPeerProbeMtuWarning) HasPayloadSize() bool`

HasPayloadSize returns a boolean if a field has been set.

### GetSourceName

`func (o *ServicesPeerProbeMtuWarning) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ServicesPeerProbeMtuWarning) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ServicesPeerProbeMtuWarning) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ServicesPeerProbeMtuWarning) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceVmId

`func (o *ServicesPeerProbeMtuWarning) GetSourceVmId() int32`

GetSourceVmId returns the SourceVmId field if non-nil, zero value otherwise.

### GetSourceVmIdOk

`func (o *ServicesPeerProbeMtuWarning) GetSourceVmIdOk() (*int32, bool)`

GetSourceVmIdOk returns a tuple with the SourceVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVmId

`func (o *ServicesPeerProbeMtuWarning) SetSourceVmId(v int32)`

SetSourceVmId sets SourceVmId field to given value.

### HasSourceVmId

`func (o *ServicesPeerProbeMtuWarning) HasSourceVmId() bool`

HasSourceVmId returns a boolean if a field has been set.

### GetTargetName

`func (o *ServicesPeerProbeMtuWarning) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ServicesPeerProbeMtuWarning) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ServicesPeerProbeMtuWarning) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ServicesPeerProbeMtuWarning) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetVmId

`func (o *ServicesPeerProbeMtuWarning) GetTargetVmId() int32`

GetTargetVmId returns the TargetVmId field if non-nil, zero value otherwise.

### GetTargetVmIdOk

`func (o *ServicesPeerProbeMtuWarning) GetTargetVmIdOk() (*int32, bool)`

GetTargetVmIdOk returns a tuple with the TargetVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmId

`func (o *ServicesPeerProbeMtuWarning) SetTargetVmId(v int32)`

SetTargetVmId sets TargetVmId field to given value.

### HasTargetVmId

`func (o *ServicesPeerProbeMtuWarning) HasTargetVmId() bool`

HasTargetVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


