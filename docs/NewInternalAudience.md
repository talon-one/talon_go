# NewInternalAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The human-friendly display name for this audience. | 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox Application. | [optional] 
**Description** | Pointer to **string** | A description of the audience. | [optional] 

## Methods

### NewNewInternalAudience

`func NewNewInternalAudience(name string, ) *NewInternalAudience`

NewNewInternalAudience instantiates a new NewInternalAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewInternalAudienceWithDefaults

`func NewNewInternalAudienceWithDefaults() *NewInternalAudience`

NewNewInternalAudienceWithDefaults instantiates a new NewInternalAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewInternalAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewInternalAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewInternalAudience) SetName(v string)`

SetName sets Name field to given value.


### GetSandbox

`func (o *NewInternalAudience) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewInternalAudience) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NewInternalAudience) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *NewInternalAudience) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetDescription

`func (o *NewInternalAudience) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewInternalAudience) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewInternalAudience) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewInternalAudience) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


