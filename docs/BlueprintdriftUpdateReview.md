# BlueprintdriftUpdateReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changed** | Pointer to **bool** | Changed is false when the recompile is artifact-identical — the honest \&quot;nothing changed\&quot; answer (the FE renders it, and the apply action offers nothing to confirm). | [optional] 
**CompareUrl** | Pointer to **string** | CompareURL is the upstream commit-compare link (context). | [optional] 
**Diff** | Pointer to [**BlueprintsourceUpdateDiff**](BlueprintsourceUpdateDiff.md) | Diff is the compiled-artifact diff (image digests, field set, resources, component set). | [optional] 
**Gates** | Pointer to **[]string** | Gates are the next compile&#39;s deploy-gate refusals (the card&#39;s gate codes — the FE renders them in wizard language beside the diff; a gated compile also refuses at apply). | [optional] 
**InstalledSha** | Pointer to **string** | InstalledSHA / NextSHA are the review&#39;s two pinned commits; NextDigest is the NEXT compile&#39;s pin (the apply target&#39;s digest — Ref.Digest is the compiled-bytes hash, never an image digest). | [optional] 
**NextDigest** | Pointer to **string** |  | [optional] 
**NextName** | Pointer to **string** | NextName is the compiled stack name (the apply target&#39;s ref name). | [optional] 
**NextSha** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**BlueprintdriftSelectedSource**](BlueprintdriftSelectedSource.md) | Source echoes the selection this review re-resolved (the originally selected source — never silently switched). | [optional] 

## Methods

### NewBlueprintdriftUpdateReview

`func NewBlueprintdriftUpdateReview() *BlueprintdriftUpdateReview`

NewBlueprintdriftUpdateReview instantiates a new BlueprintdriftUpdateReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintdriftUpdateReviewWithDefaults

`func NewBlueprintdriftUpdateReviewWithDefaults() *BlueprintdriftUpdateReview`

NewBlueprintdriftUpdateReviewWithDefaults instantiates a new BlueprintdriftUpdateReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanged

`func (o *BlueprintdriftUpdateReview) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *BlueprintdriftUpdateReview) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *BlueprintdriftUpdateReview) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *BlueprintdriftUpdateReview) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetCompareUrl

`func (o *BlueprintdriftUpdateReview) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *BlueprintdriftUpdateReview) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *BlueprintdriftUpdateReview) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.

### HasCompareUrl

`func (o *BlueprintdriftUpdateReview) HasCompareUrl() bool`

HasCompareUrl returns a boolean if a field has been set.

### GetDiff

`func (o *BlueprintdriftUpdateReview) GetDiff() BlueprintsourceUpdateDiff`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *BlueprintdriftUpdateReview) GetDiffOk() (*BlueprintsourceUpdateDiff, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *BlueprintdriftUpdateReview) SetDiff(v BlueprintsourceUpdateDiff)`

SetDiff sets Diff field to given value.

### HasDiff

`func (o *BlueprintdriftUpdateReview) HasDiff() bool`

HasDiff returns a boolean if a field has been set.

### GetGates

`func (o *BlueprintdriftUpdateReview) GetGates() []string`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *BlueprintdriftUpdateReview) GetGatesOk() (*[]string, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *BlueprintdriftUpdateReview) SetGates(v []string)`

SetGates sets Gates field to given value.

### HasGates

`func (o *BlueprintdriftUpdateReview) HasGates() bool`

HasGates returns a boolean if a field has been set.

### GetInstalledSha

`func (o *BlueprintdriftUpdateReview) GetInstalledSha() string`

GetInstalledSha returns the InstalledSha field if non-nil, zero value otherwise.

### GetInstalledShaOk

`func (o *BlueprintdriftUpdateReview) GetInstalledShaOk() (*string, bool)`

GetInstalledShaOk returns a tuple with the InstalledSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSha

`func (o *BlueprintdriftUpdateReview) SetInstalledSha(v string)`

SetInstalledSha sets InstalledSha field to given value.

### HasInstalledSha

`func (o *BlueprintdriftUpdateReview) HasInstalledSha() bool`

HasInstalledSha returns a boolean if a field has been set.

### GetNextDigest

`func (o *BlueprintdriftUpdateReview) GetNextDigest() string`

GetNextDigest returns the NextDigest field if non-nil, zero value otherwise.

### GetNextDigestOk

`func (o *BlueprintdriftUpdateReview) GetNextDigestOk() (*string, bool)`

GetNextDigestOk returns a tuple with the NextDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDigest

`func (o *BlueprintdriftUpdateReview) SetNextDigest(v string)`

SetNextDigest sets NextDigest field to given value.

### HasNextDigest

`func (o *BlueprintdriftUpdateReview) HasNextDigest() bool`

HasNextDigest returns a boolean if a field has been set.

### GetNextName

`func (o *BlueprintdriftUpdateReview) GetNextName() string`

GetNextName returns the NextName field if non-nil, zero value otherwise.

### GetNextNameOk

`func (o *BlueprintdriftUpdateReview) GetNextNameOk() (*string, bool)`

GetNextNameOk returns a tuple with the NextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextName

`func (o *BlueprintdriftUpdateReview) SetNextName(v string)`

SetNextName sets NextName field to given value.

### HasNextName

`func (o *BlueprintdriftUpdateReview) HasNextName() bool`

HasNextName returns a boolean if a field has been set.

### GetNextSha

`func (o *BlueprintdriftUpdateReview) GetNextSha() string`

GetNextSha returns the NextSha field if non-nil, zero value otherwise.

### GetNextShaOk

`func (o *BlueprintdriftUpdateReview) GetNextShaOk() (*string, bool)`

GetNextShaOk returns a tuple with the NextSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSha

`func (o *BlueprintdriftUpdateReview) SetNextSha(v string)`

SetNextSha sets NextSha field to given value.

### HasNextSha

`func (o *BlueprintdriftUpdateReview) HasNextSha() bool`

HasNextSha returns a boolean if a field has been set.

### GetSource

`func (o *BlueprintdriftUpdateReview) GetSource() BlueprintdriftSelectedSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BlueprintdriftUpdateReview) GetSourceOk() (*BlueprintdriftSelectedSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BlueprintdriftUpdateReview) SetSource(v BlueprintdriftSelectedSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *BlueprintdriftUpdateReview) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


