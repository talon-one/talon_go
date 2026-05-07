# Reward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | The name of the reward. | 
**ApiName** | Pointer to **string** | A unique identifier used to reference the reward in API integrations. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications this reward is connected to.   **Note**: Currently, a reward can only be connected to one Application.  | 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox reward. Rewards of a given type can only be connected to Applications of the same type. | 
**Status** | Pointer to **string** | The status of the reward. | 

## Methods

### NewReward

`func NewReward(id int64, created time.Time, accountId int64, name string, apiName string, applicationIds []int64, sandbox bool, status string, ) *Reward`

NewReward instantiates a new Reward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardWithDefaults

`func NewRewardWithDefaults() *Reward`

NewRewardWithDefaults instantiates a new Reward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reward) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reward) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reward) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Reward) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Reward) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Reward) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *Reward) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Reward) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Reward) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *Reward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reward) SetName(v string)`

SetName sets Name field to given value.


### GetApiName

`func (o *Reward) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *Reward) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *Reward) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetDescription

`func (o *Reward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Reward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Reward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Reward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicationIds

`func (o *Reward) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *Reward) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *Reward) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetSandbox

`func (o *Reward) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Reward) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Reward) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetStatus

`func (o *Reward) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Reward) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Reward) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


