# GitopsCreateDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adoptable** | Pointer to [**GitopsAdoptionHint**](GitopsAdoptionHint.md) | Adoptable is the live registration/install this repo + connection already carries (nil &#x3D; none). The FE renders the adoption offer instead of the create form. | [optional] 
**ComposeHealthcheck** | Pointer to **bool** | ComposeHealthcheck reports the compose file declared a healthcheck command — the mapping to an HTTP path is lossy, so health is still asked (never guessed). | [optional] 
**HealthAsked** | Pointer to **bool** | HealthAsked reports that the confirm step must ask for the health path (a serving port is known). Always true for v1 detections. | [optional] 
**IgnoredExpose** | Pointer to **[]int32** | IgnoredEXPOSE lists the Dockerfile EXPOSE ports beyond the first, in file order — the predicate names what it ignored. | [optional] 
**NameSuffixed** | Pointer to **bool** | NameSuffixed reports the collision suffix fired (the surface says why the name differs from the repo name). | [optional] 
**Notes** | Pointer to **[]string** | Notes carries the human-readable trace of the predicate&#39;s decisions. | [optional] 
**Port** | Pointer to **int32** | Port is the PROPOSED serving port. | [optional] 
**PortSource** | Pointer to **string** | PortSource names where the port came from: \&quot;compose\&quot;, \&quot;expose\&quot;, or \&quot;default\&quot; (the default is a proposal the user confirms). | [optional] 
**ProposedInstallName** | Pointer to **string** | ProposedInstallName is the DNS-1123 derivation (with the 8-hex collision suffix when a live install already holds the base name). | [optional] 
**Shape** | Pointer to **string** | Shape is \&quot;dockerfile\&quot; (the only v1 shape; a single-service compose refines a Dockerfile shape, it never replaces it). | [optional] 

## Methods

### NewGitopsCreateDetection

`func NewGitopsCreateDetection() *GitopsCreateDetection`

NewGitopsCreateDetection instantiates a new GitopsCreateDetection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitopsCreateDetectionWithDefaults

`func NewGitopsCreateDetectionWithDefaults() *GitopsCreateDetection`

NewGitopsCreateDetectionWithDefaults instantiates a new GitopsCreateDetection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdoptable

`func (o *GitopsCreateDetection) GetAdoptable() GitopsAdoptionHint`

GetAdoptable returns the Adoptable field if non-nil, zero value otherwise.

### GetAdoptableOk

`func (o *GitopsCreateDetection) GetAdoptableOk() (*GitopsAdoptionHint, bool)`

GetAdoptableOk returns a tuple with the Adoptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptable

`func (o *GitopsCreateDetection) SetAdoptable(v GitopsAdoptionHint)`

SetAdoptable sets Adoptable field to given value.

### HasAdoptable

`func (o *GitopsCreateDetection) HasAdoptable() bool`

HasAdoptable returns a boolean if a field has been set.

### GetComposeHealthcheck

`func (o *GitopsCreateDetection) GetComposeHealthcheck() bool`

GetComposeHealthcheck returns the ComposeHealthcheck field if non-nil, zero value otherwise.

### GetComposeHealthcheckOk

`func (o *GitopsCreateDetection) GetComposeHealthcheckOk() (*bool, bool)`

GetComposeHealthcheckOk returns a tuple with the ComposeHealthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeHealthcheck

`func (o *GitopsCreateDetection) SetComposeHealthcheck(v bool)`

SetComposeHealthcheck sets ComposeHealthcheck field to given value.

### HasComposeHealthcheck

`func (o *GitopsCreateDetection) HasComposeHealthcheck() bool`

HasComposeHealthcheck returns a boolean if a field has been set.

### GetHealthAsked

`func (o *GitopsCreateDetection) GetHealthAsked() bool`

GetHealthAsked returns the HealthAsked field if non-nil, zero value otherwise.

### GetHealthAskedOk

`func (o *GitopsCreateDetection) GetHealthAskedOk() (*bool, bool)`

GetHealthAskedOk returns a tuple with the HealthAsked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthAsked

`func (o *GitopsCreateDetection) SetHealthAsked(v bool)`

SetHealthAsked sets HealthAsked field to given value.

### HasHealthAsked

`func (o *GitopsCreateDetection) HasHealthAsked() bool`

HasHealthAsked returns a boolean if a field has been set.

### GetIgnoredExpose

`func (o *GitopsCreateDetection) GetIgnoredExpose() []int32`

GetIgnoredExpose returns the IgnoredExpose field if non-nil, zero value otherwise.

### GetIgnoredExposeOk

`func (o *GitopsCreateDetection) GetIgnoredExposeOk() (*[]int32, bool)`

GetIgnoredExposeOk returns a tuple with the IgnoredExpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredExpose

`func (o *GitopsCreateDetection) SetIgnoredExpose(v []int32)`

SetIgnoredExpose sets IgnoredExpose field to given value.

### HasIgnoredExpose

`func (o *GitopsCreateDetection) HasIgnoredExpose() bool`

HasIgnoredExpose returns a boolean if a field has been set.

### GetNameSuffixed

`func (o *GitopsCreateDetection) GetNameSuffixed() bool`

GetNameSuffixed returns the NameSuffixed field if non-nil, zero value otherwise.

### GetNameSuffixedOk

`func (o *GitopsCreateDetection) GetNameSuffixedOk() (*bool, bool)`

GetNameSuffixedOk returns a tuple with the NameSuffixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffixed

`func (o *GitopsCreateDetection) SetNameSuffixed(v bool)`

SetNameSuffixed sets NameSuffixed field to given value.

### HasNameSuffixed

`func (o *GitopsCreateDetection) HasNameSuffixed() bool`

HasNameSuffixed returns a boolean if a field has been set.

### GetNotes

`func (o *GitopsCreateDetection) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GitopsCreateDetection) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GitopsCreateDetection) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GitopsCreateDetection) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPort

`func (o *GitopsCreateDetection) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GitopsCreateDetection) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GitopsCreateDetection) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GitopsCreateDetection) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortSource

`func (o *GitopsCreateDetection) GetPortSource() string`

GetPortSource returns the PortSource field if non-nil, zero value otherwise.

### GetPortSourceOk

`func (o *GitopsCreateDetection) GetPortSourceOk() (*string, bool)`

GetPortSourceOk returns a tuple with the PortSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSource

`func (o *GitopsCreateDetection) SetPortSource(v string)`

SetPortSource sets PortSource field to given value.

### HasPortSource

`func (o *GitopsCreateDetection) HasPortSource() bool`

HasPortSource returns a boolean if a field has been set.

### GetProposedInstallName

`func (o *GitopsCreateDetection) GetProposedInstallName() string`

GetProposedInstallName returns the ProposedInstallName field if non-nil, zero value otherwise.

### GetProposedInstallNameOk

`func (o *GitopsCreateDetection) GetProposedInstallNameOk() (*string, bool)`

GetProposedInstallNameOk returns a tuple with the ProposedInstallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedInstallName

`func (o *GitopsCreateDetection) SetProposedInstallName(v string)`

SetProposedInstallName sets ProposedInstallName field to given value.

### HasProposedInstallName

`func (o *GitopsCreateDetection) HasProposedInstallName() bool`

HasProposedInstallName returns a boolean if a field has been set.

### GetShape

`func (o *GitopsCreateDetection) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *GitopsCreateDetection) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *GitopsCreateDetection) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *GitopsCreateDetection) HasShape() bool`

HasShape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


