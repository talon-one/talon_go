# NewGiveawaysPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this giveaways pool. | 
**Description** | Pointer to **string** | The description of this giveaways pool. | [optional] 
**SubscribedApplicationsIds** | Pointer to **[]int64** | A list of the IDs of the applications that this giveaways pool is enabled for. | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this program is a live or sandbox program. Programs of a given type can only be connected to Applications of the same type. | 

## Methods

### NewNewGiveawaysPool

`func NewNewGiveawaysPool(name string, sandbox bool, ) *NewGiveawaysPool`

NewNewGiveawaysPool instantiates a new NewGiveawaysPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewGiveawaysPoolWithDefaults

`func NewNewGiveawaysPoolWithDefaults() *NewGiveawaysPool`

NewNewGiveawaysPoolWithDefaults instantiates a new NewGiveawaysPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewGiveawaysPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewGiveawaysPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewGiveawaysPool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewGiveawaysPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewGiveawaysPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewGiveawaysPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewGiveawaysPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubscribedApplicationsIds

`func (o *NewGiveawaysPool) GetSubscribedApplicationsIds() []int64`

GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field if non-nil, zero value otherwise.

### GetSubscribedApplicationsIdsOk

`func (o *NewGiveawaysPool) GetSubscribedApplicationsIdsOk() (*[]int64, bool)`

GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedApplicationsIds

`func (o *NewGiveawaysPool) SetSubscribedApplicationsIds(v []int64)`

SetSubscribedApplicationsIds sets SubscribedApplicationsIds field to given value.

### HasSubscribedApplicationsIds

`func (o *NewGiveawaysPool) HasSubscribedApplicationsIds() bool`

HasSubscribedApplicationsIds returns a boolean if a field has been set.

### GetSandbox

`func (o *NewGiveawaysPool) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewGiveawaysPool) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NewGiveawaysPool) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


