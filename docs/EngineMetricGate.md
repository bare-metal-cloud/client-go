# EngineMetricGate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureLimit** | Pointer to **int32** | FailureLimit is the tolerated breach count: the gate fails when the breaching-sample count reaches it. Zero &#x3D; DefaultGateFailureLimit. | [optional] 
**Kind** | Pointer to [**EngineGateKind**](EngineGateKind.md) |  | [optional] 
**MaxDelta** | Pointer to **float32** | MaxDelta is the LEGACY form&#39;s ± tolerance vs the baseline. | [optional] 
**Metric** | Pointer to **string** | Metric is the LEGACY form&#39;s metric name (query&#39;s back-compat reading). | [optional] 
**MinSamples** | Pointer to **int32** | MinSamples is the scarcity floor: fewer usable samples in the window is INCONCLUSIVE (pause for a human), never a silent pass. Zero &#x3D; DefaultGateMinSamples. | [optional] 
**Provider** | Pointer to **string** | Provider names a registered source (\&quot;gateway\&quot; &#x3D; the gateway counter-delta provider; story 137.1). Unknown providers are an evaluation-time Inconclusive-with-reason, never a silent pass; the reserved phase-2 names refuse at plan time. | [optional] 
**Query** | Pointer to **string** | Query names the signal in the provider&#39;s grammar (the gateway provider v1: requests | errors | error_rate | latency, optionally \&quot;@&lt;route&gt;\&quot; — the backend workflow documents exactly what it implements). | [optional] 
**Settle** | Pointer to **int32** | Settle delays evaluation from window start — warm-up samples are skipped, never judged. Zero &#x3D; evaluate from the first sample. | [optional] 
**ThresholdRange** | Pointer to [**EngineThresholdRange**](EngineThresholdRange.md) | ThresholdRange is the inclusive [min, max] window a sample must sit inside. Nil + legacy fields &#x3D; the ±MaxDelta reading; neither refuses at plan time (a gate that can never decide). | [optional] 

## Methods

### NewEngineMetricGate

`func NewEngineMetricGate() *EngineMetricGate`

NewEngineMetricGate instantiates a new EngineMetricGate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineMetricGateWithDefaults

`func NewEngineMetricGateWithDefaults() *EngineMetricGate`

NewEngineMetricGateWithDefaults instantiates a new EngineMetricGate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureLimit

`func (o *EngineMetricGate) GetFailureLimit() int32`

GetFailureLimit returns the FailureLimit field if non-nil, zero value otherwise.

### GetFailureLimitOk

`func (o *EngineMetricGate) GetFailureLimitOk() (*int32, bool)`

GetFailureLimitOk returns a tuple with the FailureLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLimit

`func (o *EngineMetricGate) SetFailureLimit(v int32)`

SetFailureLimit sets FailureLimit field to given value.

### HasFailureLimit

`func (o *EngineMetricGate) HasFailureLimit() bool`

HasFailureLimit returns a boolean if a field has been set.

### GetKind

`func (o *EngineMetricGate) GetKind() EngineGateKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngineMetricGate) GetKindOk() (*EngineGateKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngineMetricGate) SetKind(v EngineGateKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EngineMetricGate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaxDelta

`func (o *EngineMetricGate) GetMaxDelta() float32`

GetMaxDelta returns the MaxDelta field if non-nil, zero value otherwise.

### GetMaxDeltaOk

`func (o *EngineMetricGate) GetMaxDeltaOk() (*float32, bool)`

GetMaxDeltaOk returns a tuple with the MaxDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDelta

`func (o *EngineMetricGate) SetMaxDelta(v float32)`

SetMaxDelta sets MaxDelta field to given value.

### HasMaxDelta

`func (o *EngineMetricGate) HasMaxDelta() bool`

HasMaxDelta returns a boolean if a field has been set.

### GetMetric

`func (o *EngineMetricGate) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *EngineMetricGate) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *EngineMetricGate) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *EngineMetricGate) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetMinSamples

`func (o *EngineMetricGate) GetMinSamples() int32`

GetMinSamples returns the MinSamples field if non-nil, zero value otherwise.

### GetMinSamplesOk

`func (o *EngineMetricGate) GetMinSamplesOk() (*int32, bool)`

GetMinSamplesOk returns a tuple with the MinSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSamples

`func (o *EngineMetricGate) SetMinSamples(v int32)`

SetMinSamples sets MinSamples field to given value.

### HasMinSamples

`func (o *EngineMetricGate) HasMinSamples() bool`

HasMinSamples returns a boolean if a field has been set.

### GetProvider

`func (o *EngineMetricGate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EngineMetricGate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EngineMetricGate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EngineMetricGate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetQuery

`func (o *EngineMetricGate) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EngineMetricGate) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EngineMetricGate) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EngineMetricGate) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSettle

`func (o *EngineMetricGate) GetSettle() int32`

GetSettle returns the Settle field if non-nil, zero value otherwise.

### GetSettleOk

`func (o *EngineMetricGate) GetSettleOk() (*int32, bool)`

GetSettleOk returns a tuple with the Settle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettle

`func (o *EngineMetricGate) SetSettle(v int32)`

SetSettle sets Settle field to given value.

### HasSettle

`func (o *EngineMetricGate) HasSettle() bool`

HasSettle returns a boolean if a field has been set.

### GetThresholdRange

`func (o *EngineMetricGate) GetThresholdRange() EngineThresholdRange`

GetThresholdRange returns the ThresholdRange field if non-nil, zero value otherwise.

### GetThresholdRangeOk

`func (o *EngineMetricGate) GetThresholdRangeOk() (*EngineThresholdRange, bool)`

GetThresholdRangeOk returns a tuple with the ThresholdRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdRange

`func (o *EngineMetricGate) SetThresholdRange(v EngineThresholdRange)`

SetThresholdRange sets ThresholdRange field to given value.

### HasThresholdRange

`func (o *EngineMetricGate) HasThresholdRange() bool`

HasThresholdRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


