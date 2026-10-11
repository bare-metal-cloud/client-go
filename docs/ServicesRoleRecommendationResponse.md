# ServicesRoleRecommendationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignment** | Pointer to [**[]RolerecsAssignment**](RolerecsAssignment.md) |  | [optional] 
**DataCoverage** | Pointer to [**ServicesRoleDataCoverage**](ServicesRoleDataCoverage.md) |  | [optional] 
**MasterCount** | Pointer to **int32** |  | [optional] 
**MasterCountGuidance** | Pointer to [**[]RolerecsMasterCountOption**](RolerecsMasterCountOption.md) |  | [optional] 
**SupportedMasterCounts** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewServicesRoleRecommendationResponse

`func NewServicesRoleRecommendationResponse() *ServicesRoleRecommendationResponse`

NewServicesRoleRecommendationResponse instantiates a new ServicesRoleRecommendationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesRoleRecommendationResponseWithDefaults

`func NewServicesRoleRecommendationResponseWithDefaults() *ServicesRoleRecommendationResponse`

NewServicesRoleRecommendationResponseWithDefaults instantiates a new ServicesRoleRecommendationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignment

`func (o *ServicesRoleRecommendationResponse) GetAssignment() []RolerecsAssignment`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *ServicesRoleRecommendationResponse) GetAssignmentOk() (*[]RolerecsAssignment, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *ServicesRoleRecommendationResponse) SetAssignment(v []RolerecsAssignment)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *ServicesRoleRecommendationResponse) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetDataCoverage

`func (o *ServicesRoleRecommendationResponse) GetDataCoverage() ServicesRoleDataCoverage`

GetDataCoverage returns the DataCoverage field if non-nil, zero value otherwise.

### GetDataCoverageOk

`func (o *ServicesRoleRecommendationResponse) GetDataCoverageOk() (*ServicesRoleDataCoverage, bool)`

GetDataCoverageOk returns a tuple with the DataCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCoverage

`func (o *ServicesRoleRecommendationResponse) SetDataCoverage(v ServicesRoleDataCoverage)`

SetDataCoverage sets DataCoverage field to given value.

### HasDataCoverage

`func (o *ServicesRoleRecommendationResponse) HasDataCoverage() bool`

HasDataCoverage returns a boolean if a field has been set.

### GetMasterCount

`func (o *ServicesRoleRecommendationResponse) GetMasterCount() int32`

GetMasterCount returns the MasterCount field if non-nil, zero value otherwise.

### GetMasterCountOk

`func (o *ServicesRoleRecommendationResponse) GetMasterCountOk() (*int32, bool)`

GetMasterCountOk returns a tuple with the MasterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCount

`func (o *ServicesRoleRecommendationResponse) SetMasterCount(v int32)`

SetMasterCount sets MasterCount field to given value.

### HasMasterCount

`func (o *ServicesRoleRecommendationResponse) HasMasterCount() bool`

HasMasterCount returns a boolean if a field has been set.

### GetMasterCountGuidance

`func (o *ServicesRoleRecommendationResponse) GetMasterCountGuidance() []RolerecsMasterCountOption`

GetMasterCountGuidance returns the MasterCountGuidance field if non-nil, zero value otherwise.

### GetMasterCountGuidanceOk

`func (o *ServicesRoleRecommendationResponse) GetMasterCountGuidanceOk() (*[]RolerecsMasterCountOption, bool)`

GetMasterCountGuidanceOk returns a tuple with the MasterCountGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCountGuidance

`func (o *ServicesRoleRecommendationResponse) SetMasterCountGuidance(v []RolerecsMasterCountOption)`

SetMasterCountGuidance sets MasterCountGuidance field to given value.

### HasMasterCountGuidance

`func (o *ServicesRoleRecommendationResponse) HasMasterCountGuidance() bool`

HasMasterCountGuidance returns a boolean if a field has been set.

### GetSupportedMasterCounts

`func (o *ServicesRoleRecommendationResponse) GetSupportedMasterCounts() []int32`

GetSupportedMasterCounts returns the SupportedMasterCounts field if non-nil, zero value otherwise.

### GetSupportedMasterCountsOk

`func (o *ServicesRoleRecommendationResponse) GetSupportedMasterCountsOk() (*[]int32, bool)`

GetSupportedMasterCountsOk returns a tuple with the SupportedMasterCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMasterCounts

`func (o *ServicesRoleRecommendationResponse) SetSupportedMasterCounts(v []int32)`

SetSupportedMasterCounts sets SupportedMasterCounts field to given value.

### HasSupportedMasterCounts

`func (o *ServicesRoleRecommendationResponse) HasSupportedMasterCounts() bool`

HasSupportedMasterCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


