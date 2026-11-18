# VerifyrelayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **int32** |  | [optional] 
**Blob** | Pointer to **[]int32** | Blob is the machine&#39;s sealed-delivery ciphertext, base64 — SD2&#39;s moment-of-use redelivery riding the verify round-trip. | [optional] 
**Body** | Pointer to **[]int32** | Body is the raw push body, base64. | [optional] 
**DeliveryId** | Pointer to **string** | DeliveryID is the provider&#39;s own delivery id (the ledger&#39;s correlation, surfaced in verdict audit rows). | [optional] 
**EventType** | Pointer to **string** | EventType is the raw event-type header (the machine verifies pings too — the edge drops non-push events only after a verified verdict). | [optional] 
**OrgId** | Pointer to **int32** | OrgID/TriggerID/AgentID are the tenancy and addressing triple. | [optional] 
**Provider** | Pointer to **string** | Provider is the webhook provider slug (github|gitlab) — the machine picks the verify recipe by it. | [optional] 
**ReqId** | Pointer to **string** | ReqID is the request&#39;s correlation id (the verdict queue&#39;s key). | [optional] 
**Signature** | Pointer to **string** | Signature is the raw signature header value (X-Hub-Signature-256 or X-Gitlab-Token — provider vocabulary preserved verbatim). | [optional] 
**TriggerId** | Pointer to **int32** |  | [optional] 

## Methods

### NewVerifyrelayRequest

`func NewVerifyrelayRequest() *VerifyrelayRequest`

NewVerifyrelayRequest instantiates a new VerifyrelayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyrelayRequestWithDefaults

`func NewVerifyrelayRequestWithDefaults() *VerifyrelayRequest`

NewVerifyrelayRequestWithDefaults instantiates a new VerifyrelayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *VerifyrelayRequest) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *VerifyrelayRequest) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *VerifyrelayRequest) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *VerifyrelayRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetBlob

`func (o *VerifyrelayRequest) GetBlob() []int32`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *VerifyrelayRequest) GetBlobOk() (*[]int32, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *VerifyrelayRequest) SetBlob(v []int32)`

SetBlob sets Blob field to given value.

### HasBlob

`func (o *VerifyrelayRequest) HasBlob() bool`

HasBlob returns a boolean if a field has been set.

### GetBody

`func (o *VerifyrelayRequest) GetBody() []int32`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *VerifyrelayRequest) GetBodyOk() (*[]int32, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *VerifyrelayRequest) SetBody(v []int32)`

SetBody sets Body field to given value.

### HasBody

`func (o *VerifyrelayRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDeliveryId

`func (o *VerifyrelayRequest) GetDeliveryId() string`

GetDeliveryId returns the DeliveryId field if non-nil, zero value otherwise.

### GetDeliveryIdOk

`func (o *VerifyrelayRequest) GetDeliveryIdOk() (*string, bool)`

GetDeliveryIdOk returns a tuple with the DeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryId

`func (o *VerifyrelayRequest) SetDeliveryId(v string)`

SetDeliveryId sets DeliveryId field to given value.

### HasDeliveryId

`func (o *VerifyrelayRequest) HasDeliveryId() bool`

HasDeliveryId returns a boolean if a field has been set.

### GetEventType

`func (o *VerifyrelayRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *VerifyrelayRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *VerifyrelayRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *VerifyrelayRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetOrgId

`func (o *VerifyrelayRequest) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *VerifyrelayRequest) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *VerifyrelayRequest) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *VerifyrelayRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProvider

`func (o *VerifyrelayRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VerifyrelayRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VerifyrelayRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *VerifyrelayRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReqId

`func (o *VerifyrelayRequest) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *VerifyrelayRequest) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *VerifyrelayRequest) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *VerifyrelayRequest) HasReqId() bool`

HasReqId returns a boolean if a field has been set.

### GetSignature

`func (o *VerifyrelayRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VerifyrelayRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VerifyrelayRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *VerifyrelayRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetTriggerId

`func (o *VerifyrelayRequest) GetTriggerId() int32`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *VerifyrelayRequest) GetTriggerIdOk() (*int32, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *VerifyrelayRequest) SetTriggerId(v int32)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *VerifyrelayRequest) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


