# NewReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the reward. | 
**ApiName** | Pointer to **string** | A unique identifier used to reference the reward in API integrations. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications this reward is connected to.   **Note**: Currently, a reward can only be connected to one Application.  | 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox reward. Rewards of a given type can only be connected to Applications of the same type. | 

## Methods

### NewNewReward

`func NewNewReward(name string, apiName string, applicationIds []int64, sandbox bool, ) *NewReward`

NewNewReward instantiates a new NewReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRewardWithDefaults

`func NewNewRewardWithDefaults() *NewReward`

NewNewRewardWithDefaults instantiates a new NewReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewReward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewReward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewReward) SetName(v string)`

SetName sets Name field to given value.


### GetApiName

`func (o *NewReward) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *NewReward) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *NewReward) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetDescription

`func (o *NewReward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewReward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewReward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewReward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicationIds

`func (o *NewReward) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewReward) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewReward) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetSandbox

`func (o *NewReward) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewReward) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NewReward) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


