# HandlersRunnerProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**BaseImageId** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**ConcurrencyCap** | Pointer to **int32** |  | [optional] 
**ImageArchived** | Pointer to **bool** | ImageArchived is the B3.9 catalogue&#39;s fact for the profile&#39;s BaseImageID, resolved live at read time (services/runnerimages&#39; ArchivedByImageID, one batched query): a release archiving an image flips every referencing profile&#39;s flag without touching profile rows. Empty BaseImageID → false; an unknown id → false (honest: the id was never in the catalogue — the seed validation makes this unreachable for seeded profiles). Additive field (B3.9): the frontend client type gains it when P1.16 lands. | [optional] 
**JobCount** | Pointer to **int32** | JobCount is 0 in v1, ALWAYS: no join exists today (RunnerJob carries no profile field), and inventing one would misreport. The field ships now so the response shape is stable; the count turns live in B2.5&#39;s placement, which stamps the profile on the job record. | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ManagedBy** | Pointer to **string** |  | [optional] 
**MinCpuCores** | Pointer to **int32** | MinCPUCores/MinRAMMb are the shape floors (0 &#x3D; no floor); ConcurrencyCap is the per-machine cap (0 &#x3D; the no-cap default). | [optional] 
**MinRamMb** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerProfileResponse

`func NewHandlersRunnerProfileResponse() *HandlersRunnerProfileResponse`

NewHandlersRunnerProfileResponse instantiates a new HandlersRunnerProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerProfileResponseWithDefaults

`func NewHandlersRunnerProfileResponseWithDefaults() *HandlersRunnerProfileResponse`

NewHandlersRunnerProfileResponseWithDefaults instantiates a new HandlersRunnerProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *HandlersRunnerProfileResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *HandlersRunnerProfileResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *HandlersRunnerProfileResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *HandlersRunnerProfileResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArchived

`func (o *HandlersRunnerProfileResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *HandlersRunnerProfileResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *HandlersRunnerProfileResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *HandlersRunnerProfileResponse) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetBaseImageId

`func (o *HandlersRunnerProfileResponse) GetBaseImageId() string`

GetBaseImageId returns the BaseImageId field if non-nil, zero value otherwise.

### GetBaseImageIdOk

`func (o *HandlersRunnerProfileResponse) GetBaseImageIdOk() (*string, bool)`

GetBaseImageIdOk returns a tuple with the BaseImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseImageId

`func (o *HandlersRunnerProfileResponse) SetBaseImageId(v string)`

SetBaseImageId sets BaseImageId field to given value.

### HasBaseImageId

`func (o *HandlersRunnerProfileResponse) HasBaseImageId() bool`

HasBaseImageId returns a boolean if a field has been set.

### GetCapacity

`func (o *HandlersRunnerProfileResponse) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *HandlersRunnerProfileResponse) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *HandlersRunnerProfileResponse) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *HandlersRunnerProfileResponse) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConcurrencyCap

`func (o *HandlersRunnerProfileResponse) GetConcurrencyCap() int32`

GetConcurrencyCap returns the ConcurrencyCap field if non-nil, zero value otherwise.

### GetConcurrencyCapOk

`func (o *HandlersRunnerProfileResponse) GetConcurrencyCapOk() (*int32, bool)`

GetConcurrencyCapOk returns a tuple with the ConcurrencyCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyCap

`func (o *HandlersRunnerProfileResponse) SetConcurrencyCap(v int32)`

SetConcurrencyCap sets ConcurrencyCap field to given value.

### HasConcurrencyCap

`func (o *HandlersRunnerProfileResponse) HasConcurrencyCap() bool`

HasConcurrencyCap returns a boolean if a field has been set.

### GetImageArchived

`func (o *HandlersRunnerProfileResponse) GetImageArchived() bool`

GetImageArchived returns the ImageArchived field if non-nil, zero value otherwise.

### GetImageArchivedOk

`func (o *HandlersRunnerProfileResponse) GetImageArchivedOk() (*bool, bool)`

GetImageArchivedOk returns a tuple with the ImageArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageArchived

`func (o *HandlersRunnerProfileResponse) SetImageArchived(v bool)`

SetImageArchived sets ImageArchived field to given value.

### HasImageArchived

`func (o *HandlersRunnerProfileResponse) HasImageArchived() bool`

HasImageArchived returns a boolean if a field has been set.

### GetJobCount

`func (o *HandlersRunnerProfileResponse) GetJobCount() int32`

GetJobCount returns the JobCount field if non-nil, zero value otherwise.

### GetJobCountOk

`func (o *HandlersRunnerProfileResponse) GetJobCountOk() (*int32, bool)`

GetJobCountOk returns a tuple with the JobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCount

`func (o *HandlersRunnerProfileResponse) SetJobCount(v int32)`

SetJobCount sets JobCount field to given value.

### HasJobCount

`func (o *HandlersRunnerProfileResponse) HasJobCount() bool`

HasJobCount returns a boolean if a field has been set.

### GetLabels

`func (o *HandlersRunnerProfileResponse) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HandlersRunnerProfileResponse) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HandlersRunnerProfileResponse) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HandlersRunnerProfileResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetManagedBy

`func (o *HandlersRunnerProfileResponse) GetManagedBy() string`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *HandlersRunnerProfileResponse) GetManagedByOk() (*string, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *HandlersRunnerProfileResponse) SetManagedBy(v string)`

SetManagedBy sets ManagedBy field to given value.

### HasManagedBy

`func (o *HandlersRunnerProfileResponse) HasManagedBy() bool`

HasManagedBy returns a boolean if a field has been set.

### GetMinCpuCores

`func (o *HandlersRunnerProfileResponse) GetMinCpuCores() int32`

GetMinCpuCores returns the MinCpuCores field if non-nil, zero value otherwise.

### GetMinCpuCoresOk

`func (o *HandlersRunnerProfileResponse) GetMinCpuCoresOk() (*int32, bool)`

GetMinCpuCoresOk returns a tuple with the MinCpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpuCores

`func (o *HandlersRunnerProfileResponse) SetMinCpuCores(v int32)`

SetMinCpuCores sets MinCpuCores field to given value.

### HasMinCpuCores

`func (o *HandlersRunnerProfileResponse) HasMinCpuCores() bool`

HasMinCpuCores returns a boolean if a field has been set.

### GetMinRamMb

`func (o *HandlersRunnerProfileResponse) GetMinRamMb() int32`

GetMinRamMb returns the MinRamMb field if non-nil, zero value otherwise.

### GetMinRamMbOk

`func (o *HandlersRunnerProfileResponse) GetMinRamMbOk() (*int32, bool)`

GetMinRamMbOk returns a tuple with the MinRamMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamMb

`func (o *HandlersRunnerProfileResponse) SetMinRamMb(v int32)`

SetMinRamMb sets MinRamMb field to given value.

### HasMinRamMb

`func (o *HandlersRunnerProfileResponse) HasMinRamMb() bool`

HasMinRamMb returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerProfileResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerProfileResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerProfileResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerProfileResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *HandlersRunnerProfileResponse) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *HandlersRunnerProfileResponse) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *HandlersRunnerProfileResponse) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *HandlersRunnerProfileResponse) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


