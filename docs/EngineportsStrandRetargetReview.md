# EngineportsStrandRetargetReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentTarget** | Pointer to **string** | CurrentTarget is the preview row&#39;s CurrentTarget — the before-state the approval reviewed (reported in the outcome when the live record has since moved to a different dead target). | [optional] 
**Host** | Pointer to **string** | Host is the record&#39;s hostname (the preview row&#39;s Host). | [optional] 
**RecordType** | Pointer to **string** | RecordType is the record&#39;s type (the preview row&#39;s RecordType). | [optional] 
**Zone** | Pointer to **string** | Zone is the record&#39;s zone FQDN (the preview row&#39;s Zone). | [optional] 

## Methods

### NewEngineportsStrandRetargetReview

`func NewEngineportsStrandRetargetReview() *EngineportsStrandRetargetReview`

NewEngineportsStrandRetargetReview instantiates a new EngineportsStrandRetargetReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsStrandRetargetReviewWithDefaults

`func NewEngineportsStrandRetargetReviewWithDefaults() *EngineportsStrandRetargetReview`

NewEngineportsStrandRetargetReviewWithDefaults instantiates a new EngineportsStrandRetargetReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentTarget

`func (o *EngineportsStrandRetargetReview) GetCurrentTarget() string`

GetCurrentTarget returns the CurrentTarget field if non-nil, zero value otherwise.

### GetCurrentTargetOk

`func (o *EngineportsStrandRetargetReview) GetCurrentTargetOk() (*string, bool)`

GetCurrentTargetOk returns a tuple with the CurrentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTarget

`func (o *EngineportsStrandRetargetReview) SetCurrentTarget(v string)`

SetCurrentTarget sets CurrentTarget field to given value.

### HasCurrentTarget

`func (o *EngineportsStrandRetargetReview) HasCurrentTarget() bool`

HasCurrentTarget returns a boolean if a field has been set.

### GetHost

`func (o *EngineportsStrandRetargetReview) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EngineportsStrandRetargetReview) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EngineportsStrandRetargetReview) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EngineportsStrandRetargetReview) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetRecordType

`func (o *EngineportsStrandRetargetReview) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *EngineportsStrandRetargetReview) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *EngineportsStrandRetargetReview) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *EngineportsStrandRetargetReview) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetZone

`func (o *EngineportsStrandRetargetReview) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *EngineportsStrandRetargetReview) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *EngineportsStrandRetargetReview) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *EngineportsStrandRetargetReview) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


